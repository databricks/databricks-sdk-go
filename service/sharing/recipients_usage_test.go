package sharing_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/sharing"
)

func ExampleRecipientsAPI_Create_testUcAccRecipients() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Recipients.Create(ctx, sharing.CreateRecipient{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.Recipients.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleRecipientsAPI_Get_testUcAccRecipients() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Recipients.Create(ctx, sharing.CreateRecipient{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	_, err = w.Recipients.GetByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Recipients.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}

func ExampleRecipientsAPI_ListAll_testUcAccRecipients() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.Recipients.ListAll(ctx, sharing.ListRecipientsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleRecipientsAPI_Update_testUcAccRecipients() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.Recipients.Create(ctx, sharing.CreateRecipient{
		Name: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.Recipients.Update(ctx, sharing.UpdateRecipient{
		Name:    created.Name,
		Comment: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.Recipients.DeleteByName(ctx, created.Name)
	if err != nil {
		panic(err)
	}

}
