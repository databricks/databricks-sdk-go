// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package workspaces

import (
	"context"
)

// These APIs manage workspaces for this account. A Databricks workspace is an
// environment for accessing all of your Databricks assets. The workspace
// organizes objects (notebooks, libraries, and experiments) into folders, and
// provides access to data and computational resources such as clusters and
// jobs. These endpoints are available if your account is on the E2 version of
// the platform or on a select custom plan that allows multiple workspaces per
// account.
type WorkspacesService interface {
	// Create a new workspace using a credential configuration and a storage
	// configuration, an optional network configuration (if using a
	// customer-managed VPC), an optional managed services key configuration (if
	// using customer-managed keys for managed services), and an optional
	// storage key configuration (if using customer-managed keys for storage).
	// The key configurations used for managed services and storage encryption
	// may be the same or different. **Important**: This operation is
	// asynchronous. A response with HTTP status code 200 means the request has
	// been accepted and is in progress, but does not mean that the workspace
	// deployed successfully and is running. The initial workspace status is
	// typically `PROVISIONING`. Use the workspace ID (`workspace_id`) field in
	// the response to identify the new workspace and make repeated `GET`
	// requests with the workspace ID and check its status. The workspace
	// becomes available when the status changes to `RUNNING`. You can share one
	// customer-managed VPC with multiple workspaces in a single account. It is
	// not required to create a new VPC for each workspace. However, you
	// **cannot** reuse subnets or Security Groups between workspaces. If you
	// plan to share one VPC with multiple workspaces, be sure to size your VPC
	// and subnets accordingly. Because a Databricks Account API network
	// configuration encapsulates this information, you cannot reuse a
	// Databricks Account API network configuration across workspaces. For
	// detailed instructions of creating a new workspace with this API
	// **including error handling** see [Create a new workspace using the
	// Account
	// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
	// **Important**: Customer-managed VPCs, PrivateLink, and customer-managed
	// keys are supported on a limited set of deployment and subscription types.
	// If you have questions about availability, contact your Databricks
	// representative. This operation is available only if your account is on
	// the E2 version of the platform or on a select custom plan that allows
	// multiple workspaces per account.
	CreateWorkspace(ctx context.Context, createWorkspaceRequest CreateWorkspaceRequest) (*Workspace, error)
	// Terminate and delete a Databricks workspace. From an API perspective,
	// deletion is immediate. However, it may take a few minutes for all
	// workspaces resources to be deleted, depending on the size and number of
	// workspace resources. This operation is available only if your account is
	// on the E2 version of the platform or on a select custom plan that allows
	// multiple workspaces per account.
	DeleteWorkspace(ctx context.Context, deleteWorkspaceRequest DeleteWorkspaceRequest) error
	DeleteWorkspaceByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) error
	// Get a list of all workspaces associated with an account, specified by ID.
	// This operation is available only if your account is on the E2 version of
	// the platform or on a select custom plan that allows multiple workspaces
	// per account.
	GetAllWorkspaces(ctx context.Context, getAllWorkspacesRequest GetAllWorkspacesRequest) (*ListWorkspacesResponse, error)
	GetAllWorkspacesByAccountId(ctx context.Context, accountId string) (*ListWorkspacesResponse, error)
	// Get information including status for a Databricks workspace, specified by
	// ID. In the response, the `workspace_status` field indicates the current
	// status. After initial workspace creation (which is asynchronous), make
	// repeated `GET` requests with the workspace ID and check its status. The
	// workspace becomes available when the status changes to `RUNNING`. For
	// detailed instructions of creating a new workspace with this API
	// **including error handling** see [Create a new workspace using the
	// Account
	// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
	// This operation is available only if your account is on the E2 version of
	// the platform or on a select custom plan that allows multiple workspaces
	// per account.
	GetWorkspace(ctx context.Context, getWorkspaceRequest GetWorkspaceRequest) (*Workspace, error)
	GetWorkspaceByAccountIdAndWorkspaceId(ctx context.Context, accountId string, workspaceId int64) (*Workspace, error)
	// The `PATCH` operation on this endpoint can update a workspace
	// configuration for either a running workspace or a failed workspace. The
	// elements that can be updated varies between these two use cases. ###
	// Update a failed workspace You can update a Databricks workspace
	// configuration for failed workspace deployment for some but not all
	// fields. This request supports updating only the following fields of a
	// failed workspace: - Credential configuration ID - Storage configuration
	// ID - Network configuration ID. Used only if you use customer-managed VPC.
	// - Key configuration ID for managed services (control plane storage, such
	// as notebook source and Databricks SQL queries). Used only if you use
	// customer-managed keys for managed services. - Key configuration ID for
	// workspace storage (root S3 bucket and optionally EBS volumes). Used only
	// if you use customer-managed keys for workspace storage. IMPORTANT: If the
	// workspace was ever in the running state, even if briefly before becoming
	// a failed workspace, you cannot add a new key configuration ID for
	// workspace storage. After calling the `PATCH` operation to update the
	// workspace configuration, make repeated `GET` requests with the workspace
	// ID and check the workspace status. The workspace is successful if the
	// status changes to `RUNNING`. For detailed instructions of creating a new
	// workspace with this API **including error handling** see [Create a new
	// workspace using the Account
	// API](http://docs.databricks.com/administration-guide/account-api/new-workspace.html).
	// ### Update a running workspace You can update a Databricks workspace
	// configuration for running workspaces for some but not all fields. This
	// request supports updating only the following fields of a running
	// workspace: - Credential configuration ID - Network configuration ID. Used
	// only if you already use use customer-managed VPC. This change is
	// supported only if you specified a network configuration ID in your
	// original workspace creation. In other words, you cannot switch from a
	// Databricks-managed VPC to a customer-managed VPC. Note: You cannot use a
	// network configuration update in this API to add support for PrivateLink
	// (in Public Preview). To add PrivateLink to an existing workspace, contact
	// your Databricks representative. - Key configuration ID for managed
	// services (control plane storage, such as notebook source and Databricks
	// SQL queries). Databricks does not directly encrypt the data with the
	// customer-managed key (CMK). Databricks uses both the CMK and the
	// Databricks managed key (DMK) that is unique to your workspace to encrypt
	// the Data Encryption Key (DEK). Databricks uses the DEK to encrypt your
	// workspace&#39;s managed services persisted data. If the workspace does not
	// already have a CMK for managed services, adding this ID enables managed
	// services encryption for new or updated data. Existing managed services
	// data that existed before adding the key remains not encrypted with the
	// DEK until modified. If the workspace already has customer-managed keys
	// for managed services, this request rotates (changes) the CMK keys and the
	// DEK is re-encrypted with the DMK and the new CMK. - Key configuration ID
	// for workspace storage (root S3 bucket and optionally EBS volumes). You
	// can set this only if the workspace does not already have a
	// customer-managed key configuration for workspace storage. **Important**:
	// For updating running workspaces, this API is unavailable on Mondays,
	// Tuesdays, and Thursdays from 4:30pm-7:30pm PST due to routine
	// maintenance. Plan your workspace updates accordingly. For questions about
	// this schedule, contact your Databricks representative. **Important**: To
	// update a running workspace, your workspace must have no running cluster
	// instances, which includes all-purpose clusters, job clusters, and pools
	// that may have running clusters. Terminate all cluster instances in the
	// workspace before calling this API. ### Wait until changes take effect
	// After calling the `PATCH` operation to update the workspace
	// configuration, make repeated `GET` requests with the workspace ID and
	// check the workspace status and the status of the fields. * For workspaces
	// with a Databricks-managed VPC, the workspace status becomes
	// `PROVISIONING` temporarily (typically under 20 minutes). If the workspace
	// update is successful, the workspace status changes to `RUNNING`. Note
	// that you can also check the workspace status in the [Account
	// Console](https://docs.databricks.com/administration-guide/account-settings-e2/account-console-e2.html).
	// However, you cannot use or create clusters for another 20 minutes after
	// that status change. This results in a total of up to 40 minutes in which
	// you cannot create clusters. If you create or use clusters before this
	// time interval elapses, clusters do not launch successfully, fail, or
	// could cause other unexpected behavior. * For workspaces with a
	// customer-managed VPC, the workspace status stays at status `RUNNING` and
	// the VPC change happens immediately. A change to the storage
	// customer-managed key configuration ID may take a few minutes to update,
	// so continue to check the workspace until you observe it has updated. If
	// the update fails, the workspace may revert silently to its original
	// configuration. Once the workspace has updated, you cannot use or create
	// clusters for another 20 minutes. If you create or use clusters before
	// this time interval elapses, clusters do not launch successfully, fail, or
	// could cause other unexpected behavior. If you update the _storage_
	// customer-managed key configurations, it takes 20 minutes for the changes
	// to fully take effect. During the 20 minute wait, it is important that you
	// stop all REST API calls to the DBFS API. If you are modifying _only the
	// managed services key configuration_, you can omit the 20 minute wait.
	// **Important**: Customer-managed keys and customer-managed VPCs are
	// supported by only some deployment types and subscription types. If you
	// have questions about availability, contact your Databricks
	// representative. This operation is available only if your account is on
	// the E2 version of the platform or on a select custom plan that allows
	// multiple workspaces per account.
	PatchWorkspace(ctx context.Context, updateWorkspaceRequest UpdateWorkspaceRequest) error
}
