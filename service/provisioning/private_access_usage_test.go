// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

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

func ExamplePrivateAccessAPI_Create_privateAccess() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.PrivateAccess.Create(ctx, provisioning.UpsertPrivateAccessSettingsRequest{
		PrivateAccessSettingsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Region:                    os.Getenv("AWS_REGION"),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = a.PrivateAccess.DeleteByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	if err != nil {
		panic(err)
	}

}

func ExamplePrivateAccessAPI_Get_privateAccess() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.PrivateAccess.Create(ctx, provisioning.UpsertPrivateAccessSettingsRequest{
		PrivateAccessSettingsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Region:                    os.Getenv("AWS_REGION"),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := a.PrivateAccess.GetByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = a.PrivateAccess.DeleteByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	if err != nil {
		panic(err)
	}

}

func ExamplePrivateAccessAPI_ListAll_privateAccess() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	all, err := a.PrivateAccess.List(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExamplePrivateAccessAPI_Replace_privateAccess() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	created, err := a.PrivateAccess.Create(ctx, provisioning.UpsertPrivateAccessSettingsRequest{
		PrivateAccessSettingsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Region:                    os.Getenv("AWS_REGION"),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = a.PrivateAccess.Replace(ctx, provisioning.UpsertPrivateAccessSettingsRequest{
		PrivateAccessSettingsId:   created.PrivateAccessSettingsId,
		PrivateAccessSettingsName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Region:                    os.Getenv("AWS_REGION"),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = a.PrivateAccess.DeleteByPrivateAccessSettingsId(ctx, created.PrivateAccessSettingsId)
	if err != nil {
		panic(err)
	}

}
