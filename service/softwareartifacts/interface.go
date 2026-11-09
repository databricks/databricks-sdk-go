// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package softwareartifacts

import (
	"context"
)

// Manage software artifacts and their versions in Unity Catalog.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type SoftwareArtifactsService interface {

	// Gets a software artifact.
	GetSoftwareArtifact(ctx context.Context, request GetSoftwareArtifactRequest) (*SoftwareArtifact, error)

	// Gets a software artifact version.
	GetSoftwareArtifactVersion(ctx context.Context, request GetSoftwareArtifactVersionRequest) (*SoftwareArtifactVersion, error)

	// Lists the versions of a software artifact.
	ListSoftwareArtifactVersions(ctx context.Context, request ListSoftwareArtifactVersionsRequest) (*ListSoftwareArtifactVersionsResponse, error)

	// Lists software artifacts in a schema.
	ListSoftwareArtifacts(ctx context.Context, request ListSoftwareArtifactsRequest) (*ListSoftwareArtifactsResponse, error)

	// Resolves a software artifact version from its tag alias.
	ResolveSoftwareArtifactVersion(ctx context.Context, request ResolveSoftwareArtifactVersionRequest) (*ResolveSoftwareArtifactVersionResponse, error)

	// Updates a software artifact. The name is immutable; only the owner may be
	// changed. The destination owner must exist in the account. The caller
	// needs MANAGE_ACCESS_CONTROL on the artifact. The response contains its
	// name and the requested owner; audit fields are omitted.
	UpdateSoftwareArtifact(ctx context.Context, request UpdateSoftwareArtifactRequest) (*SoftwareArtifact, error)
}
