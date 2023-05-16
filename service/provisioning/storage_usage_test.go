package provisioning_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/provisioning"
)

func ExampleStorageAPI_Create_testMwsAccLogDelivery() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	bucket, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", bucket)

	// cleanup

	err = a.Storage.DeleteByStorageConfigurationId(ctx, bucket.StorageConfigurationId)
	if err != nil {
		panic(err)
	}

}

func ExampleStorageAPI_Create_testMwsAccStorage() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", storage)

}

func ExampleStorageAPI_Create_testMwsAccWorkspaces() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: os.Getenv("TEST_ROOT_BUCKET"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", storage)

	// cleanup

	err = a.Storage.DeleteByStorageConfigurationId(ctx, storage.StorageConfigurationId)
	if err != nil {
		panic(err)
	}

}

func ExampleStorageAPI_Get_testMwsAccStorage() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	storage, err := a.Storage.Create(ctx, provisioning.CreateStorageConfigurationRequest{
		StorageConfigurationName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		RootBucketInfo: provisioning.RootBucketInfo{
			BucketName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", storage)

	byId, err := a.Storage.GetByStorageConfigurationId(ctx, storage.StorageConfigurationId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

}

func ExampleStorageAPI_ListAll_testMwsAccStorage() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	configs, err := a.Storage.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", configs)

}
