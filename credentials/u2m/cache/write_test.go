package cache

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

const (
	numGoroutines = 20
	numWrites     = 20
)

type writerFn func(string, []byte, os.FileMode) error
type readerFn func(string) ([]byte, error)

type writeReadHarness struct {
	t        *testing.T
	path     string
	writer   writerFn
	reader   readerFn
	payloads [][]byte
	valid    map[string]struct{}
}

func prepareHarness(t *testing.T, writer writerFn, reader readerFn) *writeReadHarness {
	f := setup(t)
	if err := os.MkdirAll(filepath.Dir(f), ownerExecReadWrite); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	payloads := make([][]byte, numGoroutines)
	valid := make(map[string]struct{}, numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		body := strings.Repeat(fmt.Sprintf("payload-%d-", i), 8192) // ~80KB per payload
		sum := sha256.Sum256([]byte(body))
		prefix := hex.EncodeToString(sum[:])
		full := []byte(prefix + ":" + body)
		payloads[i] = full
		valid[string(full)] = struct{}{}
	}

	return &writeReadHarness{
		t:        t,
		path:     f,
		writer:   writer,
		reader:   reader,
		payloads: payloads,
		valid:    valid,
	}
}

func validateOutput(h *writeReadHarness, out []byte) error {
	colon := strings.IndexByte(string(out), ':')
	if colon <= 0 {
		return fmt.Errorf("output missing hash prefix: %d bytes", len(out))
	}
	hexHash := string(out[:colon])
	body := out[colon+1:]
	sum := sha256.Sum256(body)
	if hex.EncodeToString(sum[:]) != hexHash {
		return fmt.Errorf("hash mismatch implies corruption/concatenation")
	}
	if _, ok := h.valid[string(out)]; !ok {
		return fmt.Errorf("final file content does not match any valid payload; length=%d", len(out))
	}
	return nil
}

func runConcurrent(h *writeReadHarness) error {
	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)
	for i := 0; i < numGoroutines; i++ {
		// Capture loop variable (we're on Go 1.18).
		// https://go.dev/blog/loopvar-preview
		i := i
		eg.Go(func() error {
			for j := 0; j < numWrites; j++ {
				if err := h.writer(h.path, h.payloads[i], ownerReadWrite); err != nil {
					return fmt.Errorf("write: %w", err)
				}
				out, err := h.reader(h.path)
				if err != nil {
					return fmt.Errorf("read: %v", err)
				}
				if err := validateOutput(h, out); err != nil {
					return err
				}
			}
			return nil
		})
	}
	return eg.Wait()
}

// TestWriteConcurrent simulates many concurrent writes to the same file.
// Each goroutine writes a different payload prefixed with the SHA-256 of the payload
// to allow detecting truncated/concatenated contents.
func TestWriteConcurrent(t *testing.T) {
	t.Run("os.WriteFile", func(t *testing.T) {
		t.Skip("this test fails to illustrate the problem")
		// Fails with:
		// --- FAIL: TestConcurrentWritesRemainValid/os.WriteFile (0.01s)
		//     write_test.go:114: hash mismatch implies corruption/concatenation
		h := prepareHarness(t, os.WriteFile, os.ReadFile)
		if err := runConcurrent(h); err != nil {
			t.Fatalf("%v", err)
		}
	})
	t.Run("writeFileAtomic", func(t *testing.T) {
		h := prepareHarness(t, writeFileAtomic, os.ReadFile)
		if err := runConcurrent(h); err != nil {
			t.Fatalf("%v", err)
		}
	})
}

func TestWriteRetainsPermissions(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")

	// Write file with 0o600 permissions.
	writeFileAtomic(path, []byte("test"), 0o600)

	// Confirm file has 0o600 permissions.
	info, err := os.Stat(path)
	require.NoError(t, err)
	require.Equal(t, fs.FileMode(0o600), info.Mode())

	// Modify file with 0o644 permissions.
	err = os.Chmod(path, 0o644)
	require.NoError(t, err)

	// Write file with 0o644 permissions.
	writeFileAtomic(path, []byte("test"), 0o600)

	// Confirm file has 0o644 permissions.
	info, err = os.Stat(path)
	require.NoError(t, err)
	require.Equal(t, fs.FileMode(0o644), info.Mode())
}
