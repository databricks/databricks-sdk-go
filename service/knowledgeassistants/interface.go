// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package knowledgeassistants

import (
	"context"
)

// Manage Knowledge Assistants and related resources.
//
// Deprecated: Do not use this interface, it will be removed in a future version of the SDK.
type KnowledgeAssistantsService interface {

	// Creates a Knowledge Assistant.
	CreateKnowledgeAssistant(ctx context.Context, request CreateKnowledgeAssistantRequest) (*KnowledgeAssistant, error)

	// Creates a Knowledge Source under a Knowledge Assistant.
	CreateKnowledgeSource(ctx context.Context, request CreateKnowledgeSourceRequest) (*KnowledgeSource, error)

	// Deletes a Knowledge Assistant.
	DeleteKnowledgeAssistant(ctx context.Context, request DeleteKnowledgeAssistantRequest) error

	// Deletes a Knowledge Source.
	DeleteKnowledgeSource(ctx context.Context, request DeleteKnowledgeSourceRequest) error

	// Gets a Knowledge Assistant.
	GetKnowledgeAssistant(ctx context.Context, request GetKnowledgeAssistantRequest) (*KnowledgeAssistant, error)

	// Gets a Knowledge Source.
	GetKnowledgeSource(ctx context.Context, request GetKnowledgeSourceRequest) (*KnowledgeSource, error)

	// Gets the permission levels that a user can have on an object.
	GetPermissionLevels(ctx context.Context, request GetKnowledgeAssistantPermissionLevelsRequest) (*GetKnowledgeAssistantPermissionLevelsResponse, error)

	// Gets the permissions of a knowledge assistant. Knowledge assistants can
	// inherit permissions from their root object.
	GetPermissions(ctx context.Context, request GetKnowledgeAssistantPermissionsRequest) (*KnowledgeAssistantPermissions, error)

	// List Knowledge Assistants
	ListKnowledgeAssistants(ctx context.Context, request ListKnowledgeAssistantsRequest) (*ListKnowledgeAssistantsResponse, error)

	// Lists Knowledge Sources under a Knowledge Assistant.
	ListKnowledgeSources(ctx context.Context, request ListKnowledgeSourcesRequest) (*ListKnowledgeSourcesResponse, error)

	// Sets permissions on an object, replacing existing permissions if they
	// exist. Deletes all direct permissions if none are specified. Objects can
	// inherit permissions from their root object.
	SetPermissions(ctx context.Context, request KnowledgeAssistantPermissionsRequest) (*KnowledgeAssistantPermissions, error)

	// Sync all non-index Knowledge Sources for a Knowledge Assistant (index
	// sources do not require sync)
	SyncKnowledgeSources(ctx context.Context, request SyncKnowledgeSourcesRequest) error

	// Updates a Knowledge Assistant.
	UpdateKnowledgeAssistant(ctx context.Context, request UpdateKnowledgeAssistantRequest) (*KnowledgeAssistant, error)

	// Updates a Knowledge Source.
	UpdateKnowledgeSource(ctx context.Context, request UpdateKnowledgeSourceRequest) (*KnowledgeSource, error)

	// Updates the permissions on a knowledge assistant. Knowledge assistants
	// can inherit permissions from their root object.
	UpdatePermissions(ctx context.Context, request KnowledgeAssistantPermissionsRequest) (*KnowledgeAssistantPermissions, error)
}
