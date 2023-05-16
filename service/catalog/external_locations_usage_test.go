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

func ExampleExternalLocationsAPI_Create_testUcAccExternalLocations() {
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

	created, err := w.ExternalLocations.Create(ctx, catalog.CreateExternalLocation{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, credential.Name)
	if err != nil {
		panic(err)
	}
	err = w.ExternalLocations.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleExternalLocationsAPI_Get_testUcAccExternalLocations() {
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

	created, err := w.ExternalLocations.Create(ctx, catalog.CreateExternalLocation{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.ExternalLocations.GetByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, credential.Name)
	if err != nil {
		panic(err)
	}
	err = w.ExternalLocations.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleExternalLocationsAPI_ListAll_testUcAccExternalLocations() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.ExternalLocations.ListAll(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleExternalLocationsAPI_Update_testUcAccExternalLocations() {
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

	created, err := w.ExternalLocations.Create(ctx, catalog.CreateExternalLocation{
		Name:           fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.ExternalLocations.Update(ctx, catalog.UpdateExternalLocation{
		Name:           created.Name,
		CredentialName: credential.Name,
		Url:            fmt.Sprintf("s3://%s/%s", os.Getenv("TEST_BUCKET"), fmt.Sprintf("sdk-%x", time.Now().UnixNano())),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.StorageCredentials.DeleteByName(ctx, credential.Name)
	if err != nil {
		panic(err)
	}
	err = w.ExternalLocations.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}
