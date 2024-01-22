package lock

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/blockblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/lease"
)

// azureBackend is a LockBackend implementation that uses Azure Blob Storage.
// It depends on the presence of the following environment variables:
// - LOCK_TOKEN: a SAS token for the storage account
// - LOCK_STORAGE_ACCOUNT_ENDPOINT: the endpoint for the storage account
type azureBackend struct {
	blobClient  *blockblob.Client
	leaseClient *lease.BlobClient
}

func (b *azureBackend) PrepareBackend(_ context.Context, lockId string) error {
	// From the Azure portal, get your Storage account's name and account key.
	token, endpoint := os.Getenv("LOCK_TOKEN"), os.Getenv("LOCK_STORAGE_ACCOUNT_ENDPOINT")
	if token == "" || endpoint == "" {
		return errors.New("LOCK_TOKEN and LOCK_STORAGE_ACCOUNT_ENDPOINT must be set")
	}

	// Create an containerClient object that wraps the container's URL and a default pipeline.
	u := fmt.Sprintf("%s/%s%s", endpoint, lockId, token)
	blobClient, err := blockblob.NewClientWithNoCredential(u, nil)
	if err != nil {
		return err
	}
	b.blobClient = blobClient
	fmt.Printf("Created a container client object with URL %s\n", blobClient.URL())
	return nil
}

func (b *azureBackend) PrepareLock(ctx context.Context, lockId string) error {
	_, err := b.blobClient.UploadBuffer(ctx, nil, nil)
	if err != nil {
		fmt.Printf("Failed to upload empty blob to %s:\n%s\n", b.blobClient.URL(), err)
		return err
	}
	fmt.Printf("Successfully uploaded empty blob to %s\n", b.blobClient.URL())
	return nil
}

func (b *azureBackend) AcquireLock(ctx context.Context, leaseId string, contents []byte, duration time.Duration) error {
	blobLeaseClient, err := lease.NewBlobClient(b.blobClient, &lease.BlobClientOptions{LeaseID: &leaseId})
	if err != nil {
		return err
	}
	b.leaseClient = blobLeaseClient
	fmt.Printf("Created a blob lease client object with lease ID %s\n", leaseId)

	// Now acquire a lease on the container.
	// You can choose to pass an empty string for proposed ID so that the service automatically assigns one for you.
	leaseDuration := int32(math.Min(60, duration.Seconds()))
	acquireLeaseResponse, err := blobLeaseClient.AcquireLease(context.TODO(), leaseDuration, nil)
	if err != nil {
		return err
	}
	fmt.Println("The container is leased for delete operations with lease ID", *acquireLeaseResponse.LeaseID)

	_, err = b.blobClient.UploadBuffer(ctx, []byte(contents), &blockblob.UploadBufferOptions{
		AccessConditions: &blob.AccessConditions{
			LeaseAccessConditions: &blob.LeaseAccessConditions{
				LeaseID: &leaseId,
			},
		},
	})
	if err != nil {
		return err
	}
	fmt.Printf("Successfully uploaded contents %s:\n%s\n", contents, b.blobClient.URL())
	return nil
}

func (b *azureBackend) RenewLock(ctx context.Context, leaseId string) error {
	_, err := b.leaseClient.RenewLease(ctx, nil)
	// Do not call handleError here, as you can only call t.Fatalf() from the main goroutine
	return err
}

func (b *azureBackend) ReleaseLock(ctx context.Context, leaseId string) error {
	_, err := b.blobClient.Delete(ctx, &blob.DeleteOptions{
		AccessConditions: &blob.AccessConditions{
			LeaseAccessConditions: &blob.LeaseAccessConditions{
				LeaseID: &leaseId,
			},
		},
	})
	// Do not call handleError here, as you can only call t.Fatalf() from the main goroutine
	if err == nil {
		fmt.Printf("Successfully deleted blob %s\n", b.blobClient.URL())
	}
	return err
}

func (b *azureBackend) RefreshDuration() time.Duration {
	return 45 * time.Second
}
