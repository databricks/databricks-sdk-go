package catalog_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

func ExampleCatalogsAPI_Create_testUcAccCatalogs() {
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

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  created.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleCatalogsAPI_Create_testUcAccSchemas() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	newCatalog, err := w.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", newCatalog)

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  newCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleCatalogsAPI_Get_testUcAccCatalogs() {
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

	_, err = w.Catalogs.GetByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  created.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleCatalogsAPI_ListAll_testUcAccCatalogs() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Catalogs.ListAll(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleCatalogsAPI_Update_testUcAccCatalogs() {
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

	_, err = w.Catalogs.Update(ctx, catalog.UpdateCatalog{
		Name:    created.Name,
		Comment: "updated",
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  created.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}
