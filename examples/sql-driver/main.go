package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"

	"github.com/databricks/databricks-sdk-go/service/warehouses"
	"github.com/databricks/databricks-sdk-go/workspaces"
	dbsql "github.com/databricks/databricks-sql-go"
)

func do(ctx context.Context) (driver.Connector, error) {
	w, err := workspaces.NewClient()
	if err != nil {
		return nil, err
	}
	wh, err := w.Warehouses.CreateWarehouseAndWait(ctx, warehouses.CreateWarehouseRequest{
		Name:        "Test from SDK Example",
		ClusterSize: "2X-Small",
	})
	if err != nil {
		return nil, err
	}
	return dbsql.NewConnector(
		dbsql.WithServerHostname(wh.OdbcParams.Hostname),
		dbsql.WithPort(wh.OdbcParams.Port),
		dbsql.WithHTTPPath(wh.OdbcParams.Path),
		//dbsql.WithAccessToken(os.Getenv("DATABRICKS_ACCESSTOKEN")
	)
}

func main() {
	ctx := context.Background()
	connector, err := do(ctx)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db := sql.OpenDB(connector)
	defer db.Close()
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	var res float64
	err1 := db.QueryRowContext(ctx, `select max(carat) from default.diamonds`).Scan(res)

	if err1 != nil {
		if err1 == sql.ErrNoRows {
			fmt.Println("not found")
			return
		} else {
			fmt.Printf("err: %v\n", err1)
		}
	}
	fmt.Println(res)
}
