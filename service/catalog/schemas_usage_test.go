package catalog_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/catalog"
)

func ExampleSchemasAPI_Create_testUcAccSchemas() {
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

	created, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: newCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  newCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, created.FullName)
	if err != nil {
		panic(err)
	}

}

func ExampleSchemasAPI_Get_testUcAccSchemas() {
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

	created, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: newCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Schemas.GetByFullName(ctx, created.FullName)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  newCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, created.FullName)
	if err != nil {
		panic(err)
	}

}

func ExampleSchemasAPI_ListAll_testUcAccSchemas() {
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

	all, err := w.Schemas.ListAll(ctx, catalog.ListSchemasRequest{
		CatalogName: newCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  newCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}

}

func ExampleSchemasAPI_Update_testUcAccSchemas() {
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

	created, err := w.Schemas.Create(ctx, catalog.CreateSchema{
		Name:        fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		CatalogName: newCatalog.Name,
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Schemas.Update(ctx, catalog.UpdateSchema{
		FullName: created.FullName,
		Comment:  fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name:  newCatalog.Name,
		Force: true,
	})
	if err != nil {
		panic(err)
	}
	err = w.Schemas.DeleteByFullName(ctx, created.FullName)
	if err != nil {
		panic(err)
	}

}
