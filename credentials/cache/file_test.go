package cache

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

func setup(t *testing.T) string {
	tempHomeDir := t.TempDir()
	return filepath.Join(tempHomeDir, "token-cache.json")
}

func TestStoreAndLookup(t *testing.T) {
	c := &FileTokenCache{
		fileLocation: setup(t),
	}
	assert.NoError(t, c.init())
	err := c.Store("x", &oauth2.Token{
		AccessToken: "abc",
	})
	require.NoError(t, err)

	err = c.Store("y", &oauth2.Token{
		AccessToken: "bcd",
	})
	require.NoError(t, err)

	l := &FileTokenCache{}
	tok, err := l.Lookup("x")
	require.NoError(t, err)
	assert.Equal(t, "abc", tok.AccessToken)

	_, err = l.Lookup("z")
	assert.Equal(t, ErrNotConfigured, err)
}

func TestNoCacheFileReturnsErrNotConfigured(t *testing.T) {
	l := &FileTokenCache{
		fileLocation: setup(t),
	}
	assert.NoError(t, l.init())
	_, err := l.Lookup("x")
	assert.Equal(t, ErrNotConfigured, err)
}

func TestLoadCorruptFile(t *testing.T) {
	f := setup(t)
	err := os.MkdirAll(filepath.Dir(f), ownerExecReadWrite)
	require.NoError(t, err)
	err = os.WriteFile(f, []byte("abc"), ownerExecReadWrite)
	require.NoError(t, err)

	l := &FileTokenCache{
		fileLocation: f,
	}
	assert.EqualError(t, l.init(), "load: parse: invalid character 'a' looking for beginning of value")
}

func TestLoadWrongVersion(t *testing.T) {
	f := setup(t)
	err := os.MkdirAll(filepath.Dir(f), ownerExecReadWrite)
	require.NoError(t, err)
	err = os.WriteFile(f, []byte(`{"version": 823, "things": []}`), ownerExecReadWrite)
	require.NoError(t, err)

	l := &FileTokenCache{}
	assert.EqualError(t, l.init(), "load: needs version 1, got version 823")
}
