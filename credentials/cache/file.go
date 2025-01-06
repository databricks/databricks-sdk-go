package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

const (
	// tokenCacheFile is the path of the default token cache, relative to the
	// user's home directory.
	tokenCacheFile = ".databricks/token-cache.json"

	// only the owner of the file has full execute, read, and write access
	ownerExecReadWrite = 0o700

	// only the owner of the file has full read and write access
	ownerReadWrite = 0o600

	// tokenCacheVersion is the version of the token cache file format.
	//
	// Version 1 format:
	//
	// {
	//   "version": 1,
	//   "tokens": {
	//     "<key>": {
	//       "access_token": "<access_token>",
	//       "token_type": "<token_type>",
	//       "refresh_token": "<refresh
	//       "expiry": "<expiry>"
	//     }
	//   }
	// }
	//
	// The format of "<key>" depends on whether the token is account- or
	// workspace-scoped:
	//  - Account-scoped: "https://<accounts host>/oidc/accounts/<account_id>"
	//  - Workspace-scoped: "https://<workspace host>"
	tokenCacheVersion = 1
)

// FileTokenCache caches tokens in "~/.databricks/token-cache.json". FileTokenCache
// implements the TokenCache interface.
// this implementation requires the calling code to do a machine-wide lock,
// otherwise the file might get corrupt.
type FileTokenCache struct {
	Version int                      `json:"version"`
	Tokens  map[string]*oauth2.Token `json:"tokens"`

	fileLocation string
}

func (c *FileTokenCache) Store(key string, t *oauth2.Token) error {
	err := c.load()
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("load: %w", err)
		}
		dir := filepath.Dir(c.fileLocation)
		if err := os.MkdirAll(dir, ownerExecReadWrite); err != nil {
			return fmt.Errorf("mkdir: %w", err)
		}
	}
	c.Version = tokenCacheVersion
	if c.Tokens == nil {
		c.Tokens = map[string]*oauth2.Token{}
	}
	c.Tokens[key] = t
	raw, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	return os.WriteFile(c.fileLocation, raw, ownerReadWrite)
}

func (c *FileTokenCache) Lookup(key string) (*oauth2.Token, error) {
	err := c.load()
	if errors.Is(err, fs.ErrNotExist) {
		return nil, ErrNotConfigured
	} else if err != nil {
		return nil, fmt.Errorf("load: %w", err)
	}
	t, ok := c.Tokens[key]
	if !ok {
		return nil, ErrNotConfigured
	}
	return t, nil
}

func (c *FileTokenCache) load() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed loading home directory: %w", err)
	}
	loc := filepath.Join(home, tokenCacheFile)
	if err != nil {
		return err
	}
	c.fileLocation = loc
	raw, err := os.ReadFile(loc)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	err = json.Unmarshal(raw, c)
	if err != nil {
		return fmt.Errorf("parse: %w", err)
	}
	if c.Version != tokenCacheVersion {
		// in the later iterations we could do state upgraders,
		// so that we transform token cache from v1 to v2 without
		// losing the tokens and asking the user to re-authenticate.
		return fmt.Errorf("needs version %d, got version %d",
			tokenCacheVersion, c.Version)
	}
	return nil
}

var _ TokenCache = (*FileTokenCache)(nil)
