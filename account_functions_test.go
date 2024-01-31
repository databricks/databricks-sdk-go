package databricks_test

import (
	"context"
	"fmt"

	"github.com/databricks/databricks-sdk-go"
)

func ExampleAccountClient_GetWorkspaceClient() {
	// Note: To run this example, your Databricks CLI default profile must be
	// configured to your Databricks account. You must have at least one running
	// workspace in your account.
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	workspaces, err := a.Workspaces.List(ctx)
	if err != nil {
		panic(err)
	}
	w, err := a.GetWorkspaceClient(workspaces[0])
	if err != nil {
		panic(err)
	}
	me, err := w.CurrentUser.Me(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(me.Active)
	// Output: true
}
