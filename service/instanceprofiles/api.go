// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package instanceprofiles

import (
	"context"

	"github.com/databricks/databricks-sdk-go/databricks/client"
)

type InstanceprofilesService interface {
	// Registers an instance profile in Databricks. In the UI, you can then
	// give users the permission to use this instance profile when launching
	// clusters. This API is only available to admin users. An example request:
	// .. code:: { &#34;instance_profile_arn&#34;:
	// &#34;arn:aws:iam::123456789:instance-profile/datascience-role&#34; }
	AddInstanceProfile(ctx context.Context, addInstanceProfileRequest AddInstanceProfileRequest) error
	// Lists the instance profiles that the calling user can use to launch a
	// cluster. This API is available to all users. An example response: ..
	// code:: { &#34;instance_profiles&#34;:
	// [&#34;arn:aws:iam::123456789:instance-profile/datascience-role&#34;] }
	ListInstanceProfiles(ctx context.Context) (*ListInstanceProfilesResponse, error)
	// Removes the instance profile with the provided ARN. Existing clusters
	// with this instance profile will continue to function. This API is only
	// accessible to admin users. An example request: .. code:: {
	// &#34;instance_profile_arn&#34;:
	// &#34;arn:aws:iam::123456789:instance-profile/datascience-role&#34; }
	RemoveInstanceProfile(ctx context.Context, removeInstanceProfileRequest RemoveInstanceProfileRequest) error
}

func New(client *client.DatabricksClient) InstanceprofilesService {
	return &InstanceprofilesAPI{
		client: client,
	}
}

type InstanceprofilesAPI struct {
	client *client.DatabricksClient
}

// Registers an instance profile in Databricks. In the UI, you can then give
// users the permission to use this instance profile when launching clusters.
// This API is only available to admin users. An example request: .. code:: {
// &#34;instance_profile_arn&#34;:
// &#34;arn:aws:iam::123456789:instance-profile/datascience-role&#34; }
func (a *InstanceprofilesAPI) AddInstanceProfile(ctx context.Context, request AddInstanceProfileRequest) error {
	path := "/api/2.0/instance-profiles/add"
	err := a.client.Post(ctx, path, request, nil)
	return err
}

// Lists the instance profiles that the calling user can use to launch a
// cluster. This API is available to all users. An example response: .. code::
// { &#34;instance_profiles&#34;:
// [&#34;arn:aws:iam::123456789:instance-profile/datascience-role&#34;] }
func (a *InstanceprofilesAPI) ListInstanceProfiles(ctx context.Context) (*ListInstanceProfilesResponse, error) {
	var listInstanceProfilesResponse ListInstanceProfilesResponse
	path := "/api/2.0/instance-profiles/list"
	err := a.client.Get(ctx, path, nil, &listInstanceProfilesResponse)
	return &listInstanceProfilesResponse, err
}

// Removes the instance profile with the provided ARN. Existing clusters with
// this instance profile will continue to function. This API is only accessible
// to admin users. An example request: .. code:: { &#34;instance_profile_arn&#34;:
// &#34;arn:aws:iam::123456789:instance-profile/datascience-role&#34; }
func (a *InstanceprofilesAPI) RemoveInstanceProfile(ctx context.Context, request RemoveInstanceProfileRequest) error {
	path := "/api/2.0/instance-profiles/remove"
	err := a.client.Post(ctx, path, request, nil)
	return err
}
