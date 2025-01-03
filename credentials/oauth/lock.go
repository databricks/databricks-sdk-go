package oauth

import (
	"fmt"
	"sync"

	"github.com/alexflint/go-filemutex"
)

// Adapts a filemutex.FileMutex to sync.Locker.
type lockerAdaptor struct {
	fileMutex *filemutex.FileMutex
}

// Lock implements sync.Locker.
func (l *lockerAdaptor) Lock() {
	err := l.fileMutex.Lock()
	if err != nil {
		panic(fmt.Errorf("lock: %w", err))
	}
}

// Unlock implements sync.Locker.
func (l *lockerAdaptor) Unlock() {
	err := l.fileMutex.Unlock()
	if err != nil {
		panic(fmt.Errorf("unlock: %w", err))
	}
}

func newLocker(path string) (sync.Locker, error) {
	m, err := filemutex.New(path)
	if err != nil {
		return nil, err
	}
	return &lockerAdaptor{fileMutex: m}, nil
}
