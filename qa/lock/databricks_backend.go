package lock

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/workspace"
)

// databricksBackend stores locks in the Databricks workspace. All locks are stored
// in a fixed directory, and each lock is stored in a file named after the lock ID.
// The workspace client is shared by all locks.
type databricksBackend struct {
	lockClient *databricks.WorkspaceClient
	lockDir    string
	lockName   string
}

// Configure the backend by setting the workspace client and lock file.
func (d *databricksBackend) PrepareBackend(ctx context.Context, lockId string) error {
	configJson, lockDirectory := os.Getenv("LOCK_CONFIG"), os.Getenv("LOCK_DIRECTORY")
	if configJson == "" || lockDirectory == "" {
		panic("LOCK_CONFIG and LOCK_DIRECTORY must be set")
	}

	c := &databricks.Config{}
	err := json.Unmarshal([]byte(configJson), c)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}

	w, err := databricks.NewWorkspaceClient(c)
	if err != nil {
		panic(fmt.Errorf("failed to create workspace client: %w", err))
	}

	d.lockClient = w
	d.lockDir = lockDirectory
	d.lockName = lockId
	return nil
}

func (d *databricksBackend) getLockPath() string {
	return fmt.Sprintf("%s/%s.lock", d.lockDir, d.lockName)
}

func (d *databricksBackend) getCurrentLockState(ctx context.Context) (*lockState, error) {
	resp, err := d.lockClient.Workspace.Download(ctx, d.getLockPath())
	if errors.Is(err, apierr.ErrNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to read lock %s state: %w", d.lockName, err)
	}
	defer resp.Close()
	bs, err := io.ReadAll(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to read lock %s state: %w", d.lockName, err)
	}

	var state lockState
	err = json.Unmarshal(bs, &state)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal lock %s state: %w", d.lockName, err)
	}

	return &state, nil
}

func (d *databricksBackend) uploadLockState(ctx context.Context, contents *lockState, overwrite bool) error {
	buf, err := json.Marshal(contents)
	if err != nil {
		return fmt.Errorf("failed to marshal contents: %w", err)
	}

	opts := []workspace.UploadOption{}
	if overwrite {
		opts = append(opts, workspace.UploadOverwrite())
	}

	return d.lockClient.Workspace.Upload(ctx, d.getLockPath(), bytes.NewReader(buf), opts...)
}

// Check if the lock file exists, and if it does, check if the lease is still valid.
// If it doesn't exist, create it.
// If it is valid, return an error.
// If it is not valid, delete the lock file and create a new one.
func (d *databricksBackend) AcquireLock(ctx context.Context, contents *lockState) error {
	lockState, err := d.getCurrentLockState(ctx)
	if err != nil {
		return fmt.Errorf("failed to get lock state during acquire: %w", err)
	}

	if lockState == nil {
		return d.uploadLockState(ctx, contents, false)
	}

	if !lockState.IsValid() {
		fmt.Printf("Lock %s is expired, deleting and recreating. Previous lock state: %#v\n", d.lockName, lockState)
		err = d.lockClient.Workspace.Delete(ctx, workspace.Delete{Path: d.getLockPath()})
		if err != nil {
			return fmt.Errorf("failed to delete lock file: %w", err)
		}

		return d.uploadLockState(ctx, contents, false)
	}

	if lockState.LeaseId != contents.LeaseId {
		return fmt.Errorf("lock %s is held by another lease, current lock state: %#v", d.lockName, lockState)
	}

	fmt.Printf("Lock %s is already held by %s\n", d.lockName, lockState.LeaseId)
	return nil
}

// Check if the lock file exists, and if it does, check if the lease is still valid.
// If it doesn't exist or is invalid, fail (the lock should be kept valid the entire time it is in use).
// If it is valid and held by a different lease, return an error.
// Otherwise, extend the lease.
func (d *databricksBackend) RenewLock(ctx context.Context, leaseId string) error {
	lockState, err := d.getCurrentLockState(ctx)
	if err != nil {
		return fmt.Errorf("failed to get lock state during renew: %w", err)
	}
	if lockState == nil {
		return fmt.Errorf("lock %s does not exist", d.lockName)
	}
	if !lockState.IsValid() {
		return fmt.Errorf("lease %s on lock %s expired at %s", leaseId, d.lockName, lockState.Expiry)
	}
	if lockState.LeaseId != leaseId {
		return fmt.Errorf("lock %s is held by %s, not %s", d.lockName, lockState.LeaseId, leaseId)
	}
	lockState.Extend(2 * d.RefreshDuration())
	return d.uploadLockState(ctx, lockState, true)
}

func (d *databricksBackend) ReleaseLock(ctx context.Context, leaseId string) error {
	lockState, err := d.getCurrentLockState(ctx)
	if err != nil {
		return fmt.Errorf("failed to get lock state during release: %w", err)
	}
	if lockState == nil {
		return nil
	}
	if lockState.IsValid() && lockState.LeaseId != leaseId {
		return fmt.Errorf("lock %s is held by %s, not %s", d.lockName, lockState.LeaseId, leaseId)
	}

	return d.lockClient.Workspace.Delete(ctx, workspace.Delete{Path: d.getLockPath()})
}

func (d *databricksBackend) RefreshDuration() time.Duration {
	return time.Minute
}

var _ LockBackend = &databricksBackend{}
