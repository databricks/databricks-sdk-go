package catalog_test

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

func ExampleStorageCredentialsAPI_Create_testUcAccStorageCredentials() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRole{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleStorageCredentialsAPI_Create_testUcAccExternalLocations() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	credential, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRole{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", credential)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, credential.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleStorageCredentialsAPI_Get_testUcAccStorageCredentials() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRole{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byName, err := w.StorageCredentials.GetByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byName)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleStorageCredentialsAPI_ListAll_testUcAccStorageCredentials() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.StorageCredentials.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleStorageCredentialsAPI_Update_testUcAccStorageCredentials() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.StorageCredentials.Create(ctx, catalog.CreateStorageCredential{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRole{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.StorageCredentials.Update(ctx, catalog.UpdateStorageCredential{
		Name:    created.Name,
		Comment: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		AwsIamRole: &catalog.AwsIamRole{
			RoleArn: os.Getenv("TEST_METASTORE_DATA_ACCESS_ARN"),
		},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}
