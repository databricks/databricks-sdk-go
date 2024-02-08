// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package catalog_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

func ExampleWorkspaceBindingsAPI_Get_catalogWorkspaceBindings() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	bindings, err := w.WorkspaceBindings.GetByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", bindings)

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  created.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}
