// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package groups

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)


type GroupsService interface {
    // Delete one group 
    DeleteGroup(ctx context.Context, deleteGroupRequest DeleteGroupRequest) error
    // Fetch information of one group 
    FetchGroup(ctx context.Context, fetchGroupRequest FetchGroupRequest) (*Group, error)
    // Get multiple groups associated with the &lt;Workspace&gt;. 
    ListGroups(ctx context.Context, listGroupsRequest ListGroupsRequest) (*ListGroupsResponse, error)
    // Create one group in the &lt;Workspace&gt;. 
    NewGroup(ctx context.Context, group Group) (*Group, error)
    // Partially update details of a group 
    PatchGroup(ctx context.Context, partialUpdate PartialUpdate) error
    // Update details of a group by replacing the entire entity 
    ReplaceGroup(ctx context.Context, group Group) error
	FetchGroupById(ctx context.Context, id string) (*Group, error)
	DeleteGroupById(ctx context.Context, id string) error
}

func New(client *client.DatabricksClient) GroupsService {
	return &GroupsAPI{
		client: client,
	}
}

type GroupsAPI struct {
	client *client.DatabricksClient
}

// Delete one group 
func (a *GroupsAPI) DeleteGroup(ctx context.Context, request DeleteGroupRequest) error {
	path := "/api/2.0/preview/scim/v2/Groups/"+request.Id
	err := a.client.Delete(ctx, path, request)
	return err
}

// Fetch information of one group 
func (a *GroupsAPI) FetchGroup(ctx context.Context, request FetchGroupRequest) (*Group, error) {
	var group Group
	path := "/api/2.0/preview/scim/v2/Groups/"+request.Id
	err := a.client.Get(ctx, path, request, &group)
	return &group, err
}

// Get multiple groups associated with the &lt;Workspace&gt;. 
func (a *GroupsAPI) ListGroups(ctx context.Context, request ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Get(ctx, path, request, &listGroupsResponse)
	return &listGroupsResponse, err
}

// Create one group in the &lt;Workspace&gt;. 
func (a *GroupsAPI) NewGroup(ctx context.Context, request Group) (*Group, error) {
	var group Group
	path := "/api/2.0/preview/scim/v2/Groups"
	err := a.client.Post(ctx, path, request, &group)
	return &group, err
}

// Partially update details of a group 
func (a *GroupsAPI) PatchGroup(ctx context.Context, request PartialUpdate) error {
	path := "/api/2.0/preview/scim/v2/Groups/"+request.Id
	err := a.client.Patch(ctx, path, request)
	return err
}

// Update details of a group by replacing the entire entity 
func (a *GroupsAPI) ReplaceGroup(ctx context.Context, request Group) error {
	path := "/api/2.0/preview/scim/v2/Groups/"+request.Id
	err := a.client.Put(ctx, path, request)
	return err
}


func (a *GroupsAPI) FetchGroupById(ctx context.Context, id string) (*Group, error) {
	return a.FetchGroup(ctx, FetchGroupRequest{
		Id: id,
	})
}

func (a *GroupsAPI) DeleteGroupById(ctx context.Context, id string) error {
	return a.DeleteGroup(ctx, DeleteGroupRequest{
		Id: id,
	})
}
