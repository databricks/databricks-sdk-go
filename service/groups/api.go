// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package groups

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks"
	"github.com/databricks/databricks-sdk-go/databricks/client"
)

func New(cfg *databricks.Config) GroupsService {
	return &GroupsAPI{
		client: client.New(cfg),
	}
}


type GroupsService interface {
    // Delete one group 
    DeleteGroup(ctx context.Context, deleteGroupRequest DeleteGroupRequest) error
    // Fetch information of one group 
    FetchGroup(ctx context.Context, fetchGroupRequest FetchGroupRequest) (*Group, error)
    // Get multiple groups associated with the &lt;Workspace&gt;. 
    ListGroups(ctx context.Context, listGroupsRequest ListGroupsRequest) (*ListGroupsResponse, error)
    // Create one group in the &lt;Workspace&gt;. 
    NewGroup(ctx context.Context, createGroupRequest CreateGroupRequest) (*Group, error)
    // Update details of a group 
    UpdateGroup(ctx context.Context, updateGroupRequest UpdateGroupRequest) error
	FetchGroupByGroupId(ctx context.Context, groupId string) (*Group, error)
	DeleteGroupByGroupId(ctx context.Context, groupId string) error
}

type GroupsAPI struct {
	client *client.DatabricksClient
}

// Delete one group 
func (a *GroupsAPI) DeleteGroup(ctx context.Context, in DeleteGroupRequest) error {
	
	err := a.client.Delete(ctx, "/scim/v2/Groups/{group_id}", in)
	return err
}

// Fetch information of one group 
func (a *GroupsAPI) FetchGroup(ctx context.Context, in FetchGroupRequest) (*Group, error) {
	var group Group
	err := a.client.Get(ctx, "/scim/v2/Groups/{group_id}", in, &group)
	return &group, err
}

// Get multiple groups associated with the &lt;Workspace&gt;. 
func (a *GroupsAPI) ListGroups(ctx context.Context, in ListGroupsRequest) (*ListGroupsResponse, error) {
	var listGroupsResponse ListGroupsResponse
	err := a.client.Get(ctx, "/scim/v2/Groups", in, &listGroupsResponse)
	return &listGroupsResponse, err
}

// Create one group in the &lt;Workspace&gt;. 
func (a *GroupsAPI) NewGroup(ctx context.Context, in CreateGroupRequest) (*Group, error) {
	var group Group
	err := a.client.Post(ctx, "/scim/v2/Groups", in, &group)
	return &group, err
}

// Update details of a group 
func (a *GroupsAPI) UpdateGroup(ctx context.Context, in UpdateGroupRequest) error {
	
	err := a.client.Patch(ctx, "/scim/v2/Groups/{group_id}", in)
	return err
}


func (a *GroupsAPI) FetchGroupByGroupId(ctx context.Context, groupId string) (*Group, error) {
	return a.FetchGroup(ctx, FetchGroupRequest{
		GroupId: groupId,
	})
}

func (a *GroupsAPI) DeleteGroupByGroupId(ctx context.Context, groupId string) error {
	return a.DeleteGroup(ctx, DeleteGroupRequest{
		GroupId: groupId,
	})
}
