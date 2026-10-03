package catalog

import (
	"context"
	"net/url"

	"github.com/databricks/databricks-sdk-go/listing"
)

// This file is hand-written (not generated). It extends the generated
// EntityTagAssignmentsAPI to percent-encode entity names and governed-tag keys
// when they travel as URL *path* parameters.
//
// Governed tag keys are hierarchical and may contain "/" (e.g.
// "Field/Shared Technical Services"). The generated impl builds the path as
// /api/2.1/unity-catalog/entity-tag-assignments/{entity_type}/{entity_name}/tags/{tag_key},
// so a raw "/" in either value is treated as a path separator and the request
// resolves to no endpoint (404). Encoding each value ("/" -> %2F) makes it route
// as a single path segment; the server decodes it back to the original value.
//
// These methods shadow the promoted (generated) entityTagAssignmentsImpl methods.
// Create is intentionally NOT overridden: it sends the assignment in the JSON
// body, where raw values are correct and must be stored verbatim. The path fields
// encoded below are `json:"-" url:"-"`, so they are used only in the path.

func (a *EntityTagAssignmentsAPI) Get(ctx context.Context, request GetEntityTagAssignmentRequest) (*EntityTagAssignment, error) {
	request.EntityName = url.PathEscape(request.EntityName)
	request.TagKey = url.PathEscape(request.TagKey)
	return a.entityTagAssignmentsImpl.Get(ctx, request)
}

func (a *EntityTagAssignmentsAPI) Delete(ctx context.Context, request DeleteEntityTagAssignmentRequest) error {
	request.EntityName = url.PathEscape(request.EntityName)
	request.TagKey = url.PathEscape(request.TagKey)
	return a.entityTagAssignmentsImpl.Delete(ctx, request)
}

func (a *EntityTagAssignmentsAPI) Update(ctx context.Context, request UpdateEntityTagAssignmentRequest) (*EntityTagAssignment, error) {
	request.EntityName = url.PathEscape(request.EntityName)
	request.TagKey = url.PathEscape(request.TagKey)
	return a.entityTagAssignmentsImpl.Update(ctx, request)
}

func (a *EntityTagAssignmentsAPI) List(ctx context.Context, request ListEntityTagAssignmentsRequest) listing.Iterator[EntityTagAssignment] {
	request.EntityName = url.PathEscape(request.EntityName)
	return a.entityTagAssignmentsImpl.List(ctx, request)
}

func (a *EntityTagAssignmentsAPI) ListAll(ctx context.Context, request ListEntityTagAssignmentsRequest) ([]EntityTagAssignment, error) {
	request.EntityName = url.PathEscape(request.EntityName)
	return a.entityTagAssignmentsImpl.ListAll(ctx, request)
}
