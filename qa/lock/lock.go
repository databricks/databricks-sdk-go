package lock

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Acquire acquires a lock according to the provided lock options. If successful,
// the lock is returned. The caller is responsible for calling the Unlock() method
// on the returned lock when the lock is no longer needed.
func Acquire(ctx context.Context, os ...LockOption) (*Lock, error) {
	opts := LockOptions{}
	for _, o := range os {
		o(&opts)
	}
	if opts.Backend == nil {
		opts.Backend = &azureBackend{}
	}

	lockId := opts.Lockable.GetLockId()
	err := opts.Backend.PrepareBackend(ctx, lockId)
	if err != nil {
		return nil, fmt.Errorf("error preparing backend: %w", err)
	}

	// Algorithm:
	//
	// 1. Create the blob with no lease and no content. If the blob already exists, this will fail. Try every 10 seconds up to configured MaxWaitDuration.
	// 2. AcquireLease the blob. If duration < 60, request the lease for the duration. Otherwise, start a goroutine to renew the lease every 45 seconds.
	// 3. Update the blob content to include the current test context.
	// 4. Run the test.
	// 5. Delete the blob. (maybe this removes the lease?)

	timesToPoll := int(opts.LeaseDuration/(10*time.Second)) + 1
	for i := 0; i < timesToPoll; i++ {
		err = opts.Backend.PrepareLock(ctx, lockId)
		if err == nil {
			break
		}
		if i == timesToPoll-1 {
			return nil, fmt.Errorf("max wait duration exceeded, last error: %w", err)
		}
		time.Sleep(10 * time.Second)
	}

	leaseId := uuid.New().String()
	contents, err := opts.getLockContents()
	if err != nil {
		return nil, err
	}
	err = opts.Backend.AcquireLock(ctx, leaseId, contents, 60*time.Second)
	if err != nil {
		return nil, fmt.Errorf("error acquiring lock: %w", err)
	}

	// Channel to signal when the lock should be released
	done := make(chan struct{})

	// Channel to communicate errors renewing the lease or deleting the blob
	errs := make(chan error, 1)

	// Renew the lock every 45 seconds until the lock is released.
	go func() {
		duration := opts.Backend.RefreshDuration()
		timer := time.NewTimer(duration)
		testTimer := time.NewTimer(opts.LeaseDuration)
		defer close(errs)
		for {
			timer.Reset(duration)
			select {
			case <-done:
				err := opts.Backend.ReleaseLock(ctx, leaseId)
				if err != nil {
					errs <- err
				}
				return
			case <-timer.C:
				err := opts.Backend.RenewLock(ctx, leaseId)
				if err != nil {
					errs <- err
					return
				}
				fmt.Printf("Successfully renewed lease %s\n", leaseId)
			case <-testTimer.C:
				errs <- errors.New("max wait duration exceeded")
				return
			}
		}
	}()

	return &Lock{
		done: done,
		errs: errs,
	}, nil
}

// Lock is a lock that is acquired by a single test. Locks can be acquired
// by calling the Acquire() method on a LockOptions struct. The lock is
// released when the Unlock() method is called. Failures to acquire the lock
// in time, to renew the leease, or to release the lock will cause the test
// to fail.
type Lock struct {
	done chan<- struct{}
	errs <-chan error
}

func (l *Lock) Unlock() error {
	close(l.done)
	return <-l.errs
}
