package catalog

import (
	"context"
	"testing"

	"github.com/databricks/databricks-sdk-go/listing"
	"github.com/databricks/databricks-sdk-go/qa"
)

func entityTagAssignmentsTestAPI(t *testing.T, method, resource string) *EntityTagAssignmentsAPI {
	t.Helper()
	requestMocks := qa.HTTPFixtures{
		{
			Method:   method,
			Resource: resource,
			Response: EntityTagAssignment{},
		},
	}
	client, server := requestMocks.Client(t)
	t.Cleanup(server.Close)
	return &EntityTagAssignmentsAPI{entityTagAssignmentsImpl: entityTagAssignmentsImpl{client: client}}
}

func TestGetEntityTagAssignmentEncodesPathParameters(t *testing.T) {
	api := entityTagAssignmentsTestAPI(
		t,
		"GET",
		"/api/2.1/unity-catalog/entity-tag-assignments/columns/main.default.events.Inferences%2FSecond/tags/cost%2Fcenter?",
	)

	_, err := api.Get(context.Background(), GetEntityTagAssignmentRequest{
		EntityType: "columns",
		EntityName: "main.default.events.Inferences/Second",
		TagKey:     "cost/center",
	})
	if err != nil {
		t.Fatalf("Get returned error: %v", err)
	}
}

func TestDeleteEntityTagAssignmentEncodesPathParameters(t *testing.T) {
	api := entityTagAssignmentsTestAPI(
		t,
		"DELETE",
		"/api/2.1/unity-catalog/entity-tag-assignments/columns/main.default.events.Inferences%2FSecond/tags/cost%2Fcenter?",
	)

	err := api.Delete(context.Background(), DeleteEntityTagAssignmentRequest{
		EntityType: "columns",
		EntityName: "main.default.events.Inferences/Second",
		TagKey:     "cost/center",
	})
	if err != nil {
		t.Fatalf("Delete returned error: %v", err)
	}
}

func TestUpdateEntityTagAssignmentEncodesPathParameters(t *testing.T) {
	api := entityTagAssignmentsTestAPI(
		t,
		"PATCH",
		"/api/2.1/unity-catalog/entity-tag-assignments/columns/main.default.events.Inferences%2FSecond/tags/cost%2Fcenter",
	)

	_, err := api.Update(context.Background(), UpdateEntityTagAssignmentRequest{
		EntityType: "columns",
		EntityName: "main.default.events.Inferences/Second",
		TagKey:     "cost/center",
	})
	if err != nil {
		t.Fatalf("Update returned error: %v", err)
	}
}

func TestListEntityTagAssignmentsEncodesEntityName(t *testing.T) {
	api := entityTagAssignmentsTestAPI(
		t,
		"GET",
		"/api/2.1/unity-catalog/entity-tag-assignments/columns/main.default.events.Inferences%2FSecond/tags?",
	)

	items, err := listing.ToSlice(context.Background(), api.List(context.Background(), ListEntityTagAssignmentsRequest{
		EntityType: "columns",
		EntityName: "main.default.events.Inferences/Second",
	}))
	if err != nil {
		t.Fatalf("List returned error: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("List returned %d items, want 0", len(items))
	}
}

func TestListAllEntityTagAssignmentsEncodesEntityName(t *testing.T) {
	api := entityTagAssignmentsTestAPI(
		t,
		"GET",
		"/api/2.1/unity-catalog/entity-tag-assignments/columns/main.default.events.Inferences%2FSecond/tags?",
	)

	items, err := api.ListAll(context.Background(), ListEntityTagAssignmentsRequest{
		EntityType: "columns",
		EntityName: "main.default.events.Inferences/Second",
	})
	if err != nil {
		t.Fatalf("ListAll returned error: %v", err)
	}
	if len(items) != 0 {
		t.Fatalf("ListAll returned %d items, want 0", len(items))
	}
}
