package cache

import (
	"fmt"
	"os"
	"path/filepath"
)

// writeFileAtomic writes data to a temporary file in the same directory and then
// renames it to the target path.
//
// Atomicity analysis:
//   - On POSIX-like systems, renaming a file within the same directory is an
//     atomic operation. This guarantees that readers either see the old file or
//     the fully written new file, never a partially written one.
//   - On non-Unix platforms, including Windows, rename is not guaranteed to be
//     atomic (see os.Rename docs). This implementation still significantly
//     reduces the risk of partial reads because the target is only updated after
//     the temporary file has been fully written, but it is not a strict atomic
//     replace on those platforms.
//   - We deliberately omit fsyncs; durability is delegated to the filesystem.
func writeFileAtomic(path string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(path)

	// Stat the target file to read its current permissions.
	fi, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("stat: %w", err)
	}

	// Use permissions of the target file if it exists.
	if fi != nil {
		perm = fi.Mode() & 0o777
	}

	// Create a temp file in the same directory to ensure the subsequent rename
	// happens within the same directory (required for atomicity on POSIX).
	tmp, err := os.CreateTemp(dir, ".tmp-*")
	if err != nil {
		return fmt.Errorf("createtemp: %w", err)
	}

	tmpName := tmp.Name()

	// Ensure the final permissions on the replacement match the desired mode.
	if err := os.Chmod(tmpName, perm); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return fmt.Errorf("chmod: %w", err)
	}

	// Write contents.
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return fmt.Errorf("write: %w", err)
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpName)
		return fmt.Errorf("close: %w", err)
	}

	// Replace the target.
	if err := os.Rename(tmpName, path); err != nil {
		os.Remove(tmpName)
		return fmt.Errorf("rename: %w", err)
	}

	return nil
}
