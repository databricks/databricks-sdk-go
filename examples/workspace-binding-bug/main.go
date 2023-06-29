package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"
	"github.com/databricks/databricks-sdk-go/service/catalog"
)

func makeCatalogName() string {
	// Generate a name with a random suffix of 16 letters
	const chars = "abcdefghijklmnopqrstuvwxyz"

	// Generate the random string
	b := make([]byte, 16)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return "miles-test-" + string(b)
}

func main() {
	// Set DATABRICKS_HOST, DATABRICKS_USERNAME, DATABRICKS_PASSWORD
	cfg := databricks.Config{}
	logger.DefaultLogger = &logger.SimpleLogger{
		Level: logger.LevelTrace,
	}
	c, err := databricks.NewWorkspaceClient(&cfg)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	cat, err := c.Catalogs.Create(ctx, catalog.CreateCatalog{
		Name: makeCatalogName(),
	})
	if err != nil {
		panic(err)
	}
	if err := c.Schemas.DeleteByFullName(ctx, cat.Name+".default"); err != nil {
		panic(fmt.Errorf("cannot remove new catalog default schema: %w", err))
	}
	c.Catalogs.Update(ctx, catalog.UpdateCatalog{
		Name:          cat.Name,
		IsolationMode: catalog.IsolationModeIsolated,
	})
	currentMetastoreAssignment, err := c.Metastores.Current(ctx)
	if err != nil {
		panic(err)
	}
	_, err = c.WorkspaceBindings.Update(ctx, catalog.UpdateWorkspaceBindings{
		Name:             cat.Name,
		AssignWorkspaces: []int64{currentMetastoreAssignment.WorkspaceId},
	})
	if err != nil {
		panic(err)
	}
	_, err = c.WorkspaceBindings.Update(ctx, catalog.UpdateWorkspaceBindings{
		Name:               cat.Name,
		UnassignWorkspaces: []int64{currentMetastoreAssignment.WorkspaceId},
	})
	if err != nil {
		panic(err)
	}
	err = c.Catalogs.Delete(ctx, catalog.DeleteCatalogRequest{
		Name: cat.Name,
	})
	if err != nil {
		panic(err)
	}
}
