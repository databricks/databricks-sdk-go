// Code generated from Databricks SDK for Go integration tests by openapi.roll.TestRegenerateExamples. DO NOT EDIT.

package iam_test

import (
	"context"

	"github.com/xuxiaoshuo/databricks-sdk-go"
	"github.com/xuxiaoshuo/databricks-sdk-go/logger"
)

func ExampleCurrentUserAPI_Me_currentUser() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	me, err := w.CurrentUser.Me(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", me)

}

func ExampleCurrentUserAPI_Me_tokens() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	me2, err := w.CurrentUser.Me(ctx)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", me2)

}
