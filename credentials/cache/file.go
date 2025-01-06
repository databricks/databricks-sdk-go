package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/oauth2"
)

const (
	// tokenCacheFile is the path of the default token cache, relative to the
	// user's home directory.
	tokenCacheFilePath = ".databricks/token-cache.json"

	// ownerExecReadWrite is the permission for the .databricks directory.
	ownerExecReadWrite = 0o700

	// ownerReadWrite is the permission for the token-cache.json file.
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

// The format of the token cache file.
type tokenCacheFile struct {
	Version int                      `json:"version"`
	Tokens  map[string]*oauth2.Token `json:"tokens"`
}

// FileTokenCache caches tokens in "~/.databricks/token-cache.json". FileTokenCache
// implements the TokenCache interface.
type FileTokenCache struct {
	fileLocation string

	// mu protects the token cache file from concurrent reads and writes.
	mu *sync.Mutex
}

func NewFileTokenCache() (*FileTokenCache, error) {
	c := &FileTokenCache{
		mu: &sync.Mutex{},
	}
	if err := c.init(); err != nil {
		return nil, err
	}
	return c, nil
}

// Store implements the TokenCache interface.
func (c *FileTokenCache) Store(key string, t *oauth2.Token) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	f, err := c.load()
	if err != nil {
		return fmt.Errorf("load: %w", err)
	}
	if f.Tokens == nil {
		f.Tokens = map[string]*oauth2.Token{}
	}
	f.Tokens[key] = t
	raw, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	return os.WriteFile(c.fileLocation, raw, ownerReadWrite)
}

// Lookup implements the TokenCache interface.
func (c *FileTokenCache) Lookup(key string) (*oauth2.Token, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	f, err := c.load()
	if err != nil {
		return nil, fmt.Errorf("load: %w", err)
	}
	t, ok := f.Tokens[key]
	if !ok {
		return nil, ErrNotConfigured
	}
	return t, nil
}

// init initializes the token cache file. It creates the file and directory if
// they do not already exist.
func (c *FileTokenCache) init() error {
	// set the default file location
	if c.fileLocation == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed loading home directory: %w", err)
		}
		c.fileLocation = filepath.Join(home, tokenCacheFilePath)
	}
	// create the directory if it doesn't already exist
	if _, err := os.Stat(filepath.Dir(c.fileLocation)); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("stat directory: %w", err)
		}
		// create the directory
		if err := os.MkdirAll(filepath.Dir(c.fileLocation), ownerExecReadWrite); err != nil {
			return fmt.Errorf("mkdir: %w", err)
		}
	}
	// create the file if it doesn't already exist
	if _, err := os.Stat(c.fileLocation); err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("stat file: %w", err)
		}
		f := &tokenCacheFile{
			Version: tokenCacheVersion,
			Tokens:  map[string]*oauth2.Token{},
		}
		raw, err := json.MarshalIndent(f, "", "  ")
		if err != nil {
			return fmt.Errorf("marshal: %w", err)
		}
		if err := os.WriteFile(c.fileLocation, raw, ownerReadWrite); err != nil {
			return fmt.Errorf("write: %w", err)
		}
	}
	return nil
}

// load loads the token cache file from disk. If the file is corrupt or if its
// version does not match tokenCacheVersion, it returns an error.
func (c *FileTokenCache) load() (*tokenCacheFile, error) {
	raw, err := os.ReadFile(c.fileLocation)
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}
	f := &tokenCacheFile{}
	err = json.Unmarshal(raw, f)
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	if f.Version != tokenCacheVersion {
		// in the later iterations we could do state upgraders,
		// so that we transform token cache from v1 to v2 without
		// losing the tokens and asking the user to re-authenticate.
		return nil, fmt.Errorf("needs version %d, got version %d", tokenCacheVersion, f.Version)
	}
	return f, nil
}

var _ TokenCache = (*FileTokenCache)(nil)
