package lock

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testBackend struct {
	prepareBackend func(lockId string) error
	prepareLock    func(lockId string) error
	acquireLock    func(leaseId string, contents []byte, duration time.Duration) error
	renewLock      func(leaseId string) error
	releaseLock    func(leaseId string) error
}

func (b *testBackend) PrepareBackend(_ context.Context, lockId string) error {
	return b.prepareBackend(lockId)
}

func (b *testBackend) PrepareLock(_ context.Context, lockId string) error {
	return b.prepareLock(lockId)
}

func (b *testBackend) AcquireLock(_ context.Context, leaseId string, contents []byte, duration time.Duration) error {
	return b.acquireLock(leaseId, contents, duration)
}

func (b *testBackend) RenewLock(_ context.Context, leaseId string) error {
	return b.renewLock(leaseId)
}

func (b *testBackend) ReleaseLock(_ context.Context, leaseId string) error {
	return b.releaseLock(leaseId)
}

func (b *testBackend) RefreshDuration() time.Duration {
	return 10 * time.Millisecond
}

func getOptions(backend *testBackend) []LockOption {
	return []LockOption{
		WithBackend(backend),
		WithLockable(NewLockable("test")),
		WithLeaseDuration(1 * time.Second),
	}
}

func TestAcquire_PrepareBackendFails(t *testing.T) {
	errTest := errors.New("test error")
	backend := &testBackend{
		prepareBackend: func(lockId string) error {
			return errTest
		},
	}
	_, err := Acquire(context.Background(), getOptions(backend)...)
	assert.ErrorIs(t, err, errTest)
}

func TestAcquire_PrepareLockFails(t *testing.T) {
	errTest := errors.New("test error")
	backend := &testBackend{
		prepareBackend: func(lockId string) error {
			return nil
		},
		prepareLock: func(lockId string) error {
			return errTest
		},
	}
	_, err := Acquire(context.Background(), getOptions(backend)...)
	assert.ErrorIs(t, err, errTest)
}

func TestAcquire_AcquireLockFails(t *testing.T) {
	errTest := errors.New("test error")
	backend := &testBackend{
		prepareBackend: func(lockId string) error {
			return nil
		},
		prepareLock: func(lockId string) error {
			return nil
		},
		acquireLock: func(leaseId string, contents []byte, duration time.Duration) error {
			return errTest
		},
	}
	_, err := Acquire(context.Background(), getOptions(backend)...)
	assert.ErrorIs(t, err, errTest)
}

func TestAcquire_UnlockFails(t *testing.T) {
	errTest := errors.New("test error")
	backend := &testBackend{
		prepareBackend: func(lockId string) error {
			return nil
		},
		prepareLock: func(lockId string) error {
			return nil
		},
		acquireLock: func(leaseId string, contents []byte, duration time.Duration) error {
			return nil
		},
		releaseLock: func(leaseId string) error {
			return errTest
		},
		renewLock: func(leaseId string) error {
			return nil
		},
	}
	lock, err := Acquire(context.Background(), getOptions(backend)...)
	assert.NoError(t, err)
	err = lock.Unlock()
	assert.ErrorIs(t, err, errTest)
}

func TestAcquire_UnlockSucceeds(t *testing.T) {
	backend := &testBackend{
		prepareBackend: func(lockId string) error {
			return nil
		},
		prepareLock: func(lockId string) error {
			return nil
		},
		acquireLock: func(leaseId string, contents []byte, duration time.Duration) error {
			return nil
		},
		releaseLock: func(leaseId string) error {
			return nil
		},
		renewLock: func(leaseId string) error {
			return nil
		},
	}
	lock, err := Acquire(context.Background(), getOptions(backend)...)
	assert.NoError(t, err)
	err = lock.Unlock()
	assert.NoError(t, err)
}
