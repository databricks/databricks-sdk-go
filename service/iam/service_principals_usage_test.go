package iam_test

import (
	"context"
	"fmt"
	"time"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/logger"

	"github.com/databricks/databricks-sdk-go/service/iam"
)

func ExampleServicePrincipalsAPI_Create_testMwsAccWorkspaceAssignment() {
	ctx := context.Background()
	a, err := databricks.NewAccountClient()
	if err != nil {
		panic(err)
	}

	spn, err := a.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", spn)

}

func ExampleServicePrincipalsAPI_Create_servicePrincipalsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	// cleanup

	err = w.ServicePrincipals.DeleteById(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleServicePrincipalsAPI_Get_servicePrincipalsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	byId, err := w.ServicePrincipals.GetById(ctx, created.Id)
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", byId)

	// cleanup

	err = w.ServicePrincipals.DeleteById(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}

func ExampleServicePrincipalsAPI_ListAll_servicePrincipalsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	all, err := w.ServicePrincipals.ListAll(ctx, iam.ListServicePrincipalsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", all)

}

func ExampleServicePrincipalsAPI_ServicePrincipalDisplayNameToIdMap_servicePrincipalsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	names, err := w.ServicePrincipals.ServicePrincipalDisplayNameToIdMap(ctx, iam.ListServicePrincipalsRequest{})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", names)

}

func ExampleServicePrincipalsAPI_Update_servicePrincipalsOnAws() {
	ctx := context.Background()
	w, err := databricks.NewWorkspaceClient()
	if err != nil {
		panic(err)
	}

	created, err := w.ServicePrincipals.Create(ctx, iam.ServicePrincipal{
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
	})
	if err != nil {
		panic(err)
	}
	logger.Infof(ctx, "found %v", created)

	err = w.ServicePrincipals.Update(ctx, iam.ServicePrincipal{
		Id:          created.Id,
		DisplayName: fmt.Sprintf("sdk-%x", time.Now().UnixNano()),
		Roles: []iam.ComplexValue{iam.ComplexValue{
			Value: "xyz",
		}},
	})
	if err != nil {
		panic(err)
	}

	// cleanup

	err = w.ServicePrincipals.DeleteById(ctx, created.Id)
	if err != nil {
		panic(err)
	}

}
