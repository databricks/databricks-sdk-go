// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// all definitions in this file are in alphabetical order

type Ai21LabsConfig struct {
	// The Databricks secret key reference for an AI21Labs API key.
	Ai21labsApiKey string `json:"ai21labs_api_key"`
}

type AnthropicConfig struct {
	// The Databricks secret key reference for an Anthropic API key.
	AnthropicApiKey string `json:"anthropic_api_key"`
}

type AppEvents struct {
	EventName string `json:"event_name,omitempty"`

	EventTime string `json:"event_time,omitempty"`

	EventType string `json:"event_type,omitempty"`

	Message string `json:"message,omitempty"`

	ServiceName string `json:"service_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *AppEvents) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppEvents) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppManifest struct {
	// Workspace dependencies.
	Dependencies []any `json:"dependencies,omitempty"`
	// application description
	Description string `json:"description,omitempty"`
	// Ingress rules for app public endpoints
	Ingress any `json:"ingress,omitempty"`
	// Only a-z and dashes (-). Max length of 30.
	Name string `json:"name,omitempty"`
	// Container private registry
	Registry any `json:"registry,omitempty"`
	// list of app services. Restricted to one for now.
	Services any `json:"services,omitempty"`
	// The manifest format version. Must be set to 1.
	Version any `json:"version,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *AppManifest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppManifest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AppServiceStatus struct {
	Deployment any `json:"deployment,omitempty"`

	Name string `json:"name,omitempty"`

	Template any `json:"template,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *AppServiceStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AppServiceStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AutoCaptureConfigInput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if it was already set.
	CatalogName string `json:"catalog_name,omitempty"`
	// If inference tables are enabled or not. NOTE: If you have already
	// disabled payload logging once, you cannot enable again.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if it was already set.
	SchemaName string `json:"schema_name,omitempty"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if it was already set.
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *AutoCaptureConfigInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoCaptureConfigInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AutoCaptureConfigOutput struct {
	// The name of the catalog in Unity Catalog.
	CatalogName string `json:"catalog_name,omitempty"`
	// If inference tables are enabled or not.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog.
	SchemaName string `json:"schema_name,omitempty"`

	State *AutoCaptureState `json:"state,omitempty"`
	// The prefix of the table in Unity Catalog.
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *AutoCaptureConfigOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoCaptureConfigOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type AutoCaptureState struct {
	PayloadTable *PayloadTable `json:"payload_table,omitempty"`
}

type AwsBedrockConfig struct {
	// The Databricks secret key reference for an AWS Access Key ID with
	// permissions to interact with Bedrock services.
	AwsAccessKeyId string `json:"aws_access_key_id"`
	// The AWS region to use. Bedrock has to be enabled there.
	AwsRegion string `json:"aws_region"`
	// The Databricks secret key reference for an AWS Secret Access Key paired
	// with the access key ID, with permissions to interact with Bedrock
	// services.
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	// The underlying provider in AWS Bedrock. Supported values (case
	// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
	BedrockProvider AwsBedrockConfigBedrockProvider `json:"bedrock_provider"`
}

// The underlying provider in AWS Bedrock. Supported values (case insensitive)
// include: Anthropic, Cohere, AI21Labs, Amazon.
type AwsBedrockConfigBedrockProvider string

const AwsBedrockConfigBedrockProviderAi21labs AwsBedrockConfigBedrockProvider = `ai21labs`

const AwsBedrockConfigBedrockProviderAmazon AwsBedrockConfigBedrockProvider = `amazon`

const AwsBedrockConfigBedrockProviderAnthropic AwsBedrockConfigBedrockProvider = `anthropic`

const AwsBedrockConfigBedrockProviderCohere AwsBedrockConfigBedrockProvider = `cohere`

// String representation for [fmt.Print]
func (f *AwsBedrockConfigBedrockProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AwsBedrockConfigBedrockProvider) Set(v string) error {
	switch v {
	case `ai21labs`, `amazon`, `anthropic`, `cohere`:
		*f = AwsBedrockConfigBedrockProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ai21labs", "amazon", "anthropic", "cohere"`, v)
	}
}

// Type always returns AwsBedrockConfigBedrockProvider to satisfy [pflag.Value] interface
func (f *AwsBedrockConfigBedrockProvider) Type() string {
	return "AwsBedrockConfigBedrockProvider"
}

// Retrieve the logs associated with building the model's environment for a
// given serving endpoint's served model.
type BuildLogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" url:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName string `json:"-" url:"-"`
}

type BuildLogsResponse struct {
	// The logs associated with building the served entity's environment.
	Logs string `json:"logs"`
}

type ChatMessage struct {
	// The content of the message.
	Content string `json:"content,omitempty"`
	// The role of the message. One of [system, user, assistant].
	Role ChatMessageRole `json:"role,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ChatMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ChatMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The role of the message. One of [system, user, assistant].
type ChatMessageRole string

const ChatMessageRoleAssistant ChatMessageRole = `assistant`

const ChatMessageRoleSystem ChatMessageRole = `system`

const ChatMessageRoleUser ChatMessageRole = `user`

// String representation for [fmt.Print]
func (f *ChatMessageRole) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ChatMessageRole) Set(v string) error {
	switch v {
	case `assistant`, `system`, `user`:
		*f = ChatMessageRole(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "assistant", "system", "user"`, v)
	}
}

// Type always returns ChatMessageRole to satisfy [pflag.Value] interface
func (f *ChatMessageRole) Type() string {
	return "ChatMessageRole"
}

type CohereConfig struct {
	// The Databricks secret key reference for a Cohere API key.
	CohereApiKey string `json:"cohere_api_key"`
}

type CreateServingEndpoint struct {
	// The core config of the serving endpoint.
	Config EndpointCoreConfigInput `json:"config"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name string `json:"name"`
	// Rate limits to be applied to the serving endpoint. NOTE: only external
	// and foundation model endpoints are supported as of now.
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags []EndpointTag `json:"tags,omitempty"`
}

type DatabricksModelServingConfig struct {
	// The Databricks secret key reference for a Databricks API token that
	// corresponds to a user or service principal with Can Query access to the
	// model serving endpoint pointed to by this external model.
	DatabricksApiToken string `json:"databricks_api_token"`
	// The URL of the Databricks workspace containing the model serving endpoint
	// pointed to by this external model.
	DatabricksWorkspaceUrl string `json:"databricks_workspace_url"`
}

type DataframeSplitInput struct {
	Columns []any `json:"columns,omitempty"`

	Data []any `json:"data,omitempty"`

	Index []int `json:"index,omitempty"`
}

// Delete an application
type DeleteAppRequest struct {
	// The name of an application. This field is required.
	Name string `json:"-" url:"-"`
}

type DeleteAppResponse struct {
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DeleteAppResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeleteAppResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Delete a serving endpoint
type DeleteServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
}

type DeployAppRequest struct {
	// Manifest that specifies the application requirements
	Manifest AppManifest `json:"manifest"`
	// Information passed at app deployment time to fulfill app dependencies
	Resources any `json:"resources,omitempty"`
}

type DeploymentStatus struct {
	// Container logs.
	ContainerLogs []any `json:"container_logs,omitempty"`
	// description
	DeploymentId string `json:"deployment_id,omitempty"`
	// Supplementary information about pod
	ExtraInfo string `json:"extra_info,omitempty"`
	// State: one of DEPLOYING,SUCCESS, FAILURE, DEPLOYMENT_STATE_UNSPECIFIED
	State DeploymentStatusState `json:"state,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *DeploymentStatus) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DeploymentStatus) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// State: one of DEPLOYING,SUCCESS, FAILURE, DEPLOYMENT_STATE_UNSPECIFIED
type DeploymentStatusState string

const DeploymentStatusStateDeploying DeploymentStatusState = `DEPLOYING`

const DeploymentStatusStateDeploymentStateUnspecified DeploymentStatusState = `DEPLOYMENT_STATE_UNSPECIFIED`

const DeploymentStatusStateFailure DeploymentStatusState = `FAILURE`

const DeploymentStatusStateSuccess DeploymentStatusState = `SUCCESS`

// String representation for [fmt.Print]
func (f *DeploymentStatusState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeploymentStatusState) Set(v string) error {
	switch v {
	case `DEPLOYING`, `DEPLOYMENT_STATE_UNSPECIFIED`, `FAILURE`, `SUCCESS`:
		*f = DeploymentStatusState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPLOYING", "DEPLOYMENT_STATE_UNSPECIFIED", "FAILURE", "SUCCESS"`, v)
	}
}

// Type always returns DeploymentStatusState to satisfy [pflag.Value] interface
func (f *DeploymentStatusState) Type() string {
	return "DeploymentStatusState"
}

type EmbeddingsV1ResponseEmbeddingElement struct {
	Embedding []float64 `json:"embedding,omitempty"`
	// The index of the embedding in the response.
	Index int `json:"index,omitempty"`
	// This will always be 'embedding'.
	Object EmbeddingsV1ResponseEmbeddingElementObject `json:"object,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EmbeddingsV1ResponseEmbeddingElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingsV1ResponseEmbeddingElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// This will always be 'embedding'.
type EmbeddingsV1ResponseEmbeddingElementObject string

const EmbeddingsV1ResponseEmbeddingElementObjectEmbedding EmbeddingsV1ResponseEmbeddingElementObject = `embedding`

// String representation for [fmt.Print]
func (f *EmbeddingsV1ResponseEmbeddingElementObject) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EmbeddingsV1ResponseEmbeddingElementObject) Set(v string) error {
	switch v {
	case `embedding`:
		*f = EmbeddingsV1ResponseEmbeddingElementObject(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "embedding"`, v)
	}
}

// Type always returns EmbeddingsV1ResponseEmbeddingElementObject to satisfy [pflag.Value] interface
func (f *EmbeddingsV1ResponseEmbeddingElementObject) Type() string {
	return "EmbeddingsV1ResponseEmbeddingElementObject"
}

type EndpointCoreConfigInput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig *AutoCaptureConfigInput `json:"auto_capture_config,omitempty"`
	// The name of the serving endpoint to update. This field is required.
	Name string `json:"-" url:"-"`
	// A list of served entities for the endpoint to serve. A serving endpoint
	// can have up to 10 served entities.
	ServedEntities []ServedEntityInput `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) A list of served models for the
	// endpoint to serve. A serving endpoint can have up to 10 served models.
	ServedModels []ServedModelInput `json:"served_models,omitempty"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`
}

type EndpointCoreConfigOutput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog.
	AutoCaptureConfig *AutoCaptureConfigOutput `json:"auto_capture_config,omitempty"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int `json:"config_version,omitempty"`
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntityOutput `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelOutput `json:"served_models,omitempty"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EndpointCoreConfigOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointCoreConfigOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointCoreConfigSummary struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntitySpec `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelSpec `json:"served_models,omitempty"`
}

type EndpointPendingConfig struct {
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int `json:"config_version,omitempty"`
	// The list of served entities belonging to the last issued update to the
	// serving endpoint.
	ServedEntities []ServedEntityOutput `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models
	// belonging to the last issued update to the serving endpoint.
	ServedModels []ServedModelOutput `json:"served_models,omitempty"`
	// The timestamp when the update to the pending config started.
	StartTime int64 `json:"start_time,omitempty"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EndpointPendingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointPendingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type EndpointState struct {
	// The state of an endpoint's config update. This informs the user if the
	// pending_config is in progress, if the update failed, or if there is no
	// update in progress. Note that if the endpoint's config_update state value
	// is IN_PROGRESS, another update can not be made until the update completes
	// or fails."
	ConfigUpdate EndpointStateConfigUpdate `json:"config_update,omitempty"`
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served entities in its
	// active configuration are ready. If any of the actively served entities
	// are in a non-ready state, the endpoint state will be NOT_READY.
	Ready EndpointStateReady `json:"ready,omitempty"`
}

// The state of an endpoint's config update. This informs the user if the
// pending_config is in progress, if the update failed, or if there is no update
// in progress. Note that if the endpoint's config_update state value is
// IN_PROGRESS, another update can not be made until the update completes or
// fails."
type EndpointStateConfigUpdate string

const EndpointStateConfigUpdateInProgress EndpointStateConfigUpdate = `IN_PROGRESS`

const EndpointStateConfigUpdateNotUpdating EndpointStateConfigUpdate = `NOT_UPDATING`

const EndpointStateConfigUpdateUpdateFailed EndpointStateConfigUpdate = `UPDATE_FAILED`

// String representation for [fmt.Print]
func (f *EndpointStateConfigUpdate) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStateConfigUpdate) Set(v string) error {
	switch v {
	case `IN_PROGRESS`, `NOT_UPDATING`, `UPDATE_FAILED`:
		*f = EndpointStateConfigUpdate(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN_PROGRESS", "NOT_UPDATING", "UPDATE_FAILED"`, v)
	}
}

// Type always returns EndpointStateConfigUpdate to satisfy [pflag.Value] interface
func (f *EndpointStateConfigUpdate) Type() string {
	return "EndpointStateConfigUpdate"
}

// The state of an endpoint, indicating whether or not the endpoint is
// queryable. An endpoint is READY if all of the served entities in its active
// configuration are ready. If any of the actively served entities are in a
// non-ready state, the endpoint state will be NOT_READY.
type EndpointStateReady string

const EndpointStateReadyNotReady EndpointStateReady = `NOT_READY`

const EndpointStateReadyReady EndpointStateReady = `READY`

// String representation for [fmt.Print]
func (f *EndpointStateReady) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStateReady) Set(v string) error {
	switch v {
	case `NOT_READY`, `READY`:
		*f = EndpointStateReady(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "NOT_READY", "READY"`, v)
	}
}

// Type always returns EndpointStateReady to satisfy [pflag.Value] interface
func (f *EndpointStateReady) Type() string {
	return "EndpointStateReady"
}

type EndpointTag struct {
	// Key field for a serving endpoint tag.
	Key string `json:"key"`
	// Optional value field for a serving endpoint tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *EndpointTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Retrieve the metrics associated with a serving endpoint
type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name string `json:"-" url:"-"`
}

type ExternalModel struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	Ai21labsConfig *Ai21LabsConfig `json:"ai21labs_config,omitempty"`
	// Anthropic Config. Only required if the provider is 'anthropic'.
	AnthropicConfig *AnthropicConfig `json:"anthropic_config,omitempty"`
	// AWS Bedrock Config. Only required if the provider is 'aws-bedrock'.
	AwsBedrockConfig *AwsBedrockConfig `json:"aws_bedrock_config,omitempty"`
	// Cohere Config. Only required if the provider is 'cohere'.
	CohereConfig *CohereConfig `json:"cohere_config,omitempty"`
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	DatabricksModelServingConfig *DatabricksModelServingConfig `json:"databricks_model_serving_config,omitempty"`
	// The name of the external model.
	Name string `json:"name"`
	// OpenAI Config. Only required if the provider is 'openai'.
	OpenaiConfig *OpenAiConfig `json:"openai_config,omitempty"`
	// PaLM Config. Only required if the provider is 'palm'.
	PalmConfig *PaLmConfig `json:"palm_config,omitempty"`
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'aws-bedrock', 'cohere',
	// 'databricks-model-serving', 'openai', and 'palm'.",
	Provider ExternalModelProvider `json:"provider"`
	// The task type of the external model.
	Task string `json:"task"`
}

// The name of the provider for the external model. Currently, the supported
// providers are 'ai21labs', 'anthropic', 'aws-bedrock', 'cohere',
// 'databricks-model-serving', 'openai', and 'palm'.",
type ExternalModelProvider string

const ExternalModelProviderAi21labs ExternalModelProvider = `ai21labs`

const ExternalModelProviderAnthropic ExternalModelProvider = `anthropic`

const ExternalModelProviderAwsBedrock ExternalModelProvider = `aws-bedrock`

const ExternalModelProviderCohere ExternalModelProvider = `cohere`

const ExternalModelProviderDatabricksModelServing ExternalModelProvider = `databricks-model-serving`

const ExternalModelProviderOpenai ExternalModelProvider = `openai`

const ExternalModelProviderPalm ExternalModelProvider = `palm`

// String representation for [fmt.Print]
func (f *ExternalModelProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExternalModelProvider) Set(v string) error {
	switch v {
	case `ai21labs`, `anthropic`, `aws-bedrock`, `cohere`, `databricks-model-serving`, `openai`, `palm`:
		*f = ExternalModelProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ai21labs", "anthropic", "aws-bedrock", "cohere", "databricks-model-serving", "openai", "palm"`, v)
	}
}

// Type always returns ExternalModelProvider to satisfy [pflag.Value] interface
func (f *ExternalModelProvider) Type() string {
	return "ExternalModelProvider"
}

type ExternalModelUsageElement struct {
	// The number of tokens in the chat/completions response.
	CompletionTokens int `json:"completion_tokens,omitempty"`
	// The number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens,omitempty"`
	// The total number of tokens in the prompt and response.
	TotalTokens int `json:"total_tokens,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ExternalModelUsageElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalModelUsageElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type FoundationModel struct {
	// The description of the foundation model.
	Description string `json:"description,omitempty"`
	// The display name of the foundation model.
	DisplayName string `json:"display_name,omitempty"`
	// The URL to the documentation of the foundation model.
	Docs string `json:"docs,omitempty"`
	// The name of the foundation model.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *FoundationModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FoundationModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get deployment status for an application
type GetAppDeploymentStatusRequest struct {
	// The deployment id for an application. This field is required.
	DeploymentId string `json:"-" url:"-"`
	// Boolean flag to include application logs
	IncludeAppLog string `json:"-" url:"include_app_log,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAppDeploymentStatusRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAppDeploymentStatusRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get definition for an application
type GetAppRequest struct {
	// The name of an application. This field is required.
	Name string `json:"-" url:"-"`
}

type GetAppResponse struct {
	CurrentServices []AppServiceStatus `json:"current_services,omitempty"`

	Name string `json:"name,omitempty"`

	PendingServices []AppServiceStatus `json:"pending_services,omitempty"`

	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *GetAppResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GetAppResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Get deployment events for an application
type GetEventsRequest struct {
	// The name of an application. This field is required.
	Name string `json:"-" url:"-"`
}

// Get serving endpoint permission levels
type GetServingEndpointPermissionLevelsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ServingEndpointPermissionsDescription `json:"permission_levels,omitempty"`
}

// Get serving endpoint permissions
type GetServingEndpointPermissionsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

// Get a single serving endpoint
type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
}

type ListAppEventsResponse struct {
	// App events
	Events []AppEvents `json:"events,omitempty"`
}

type ListAppsResponse struct {
	// Available apps.
	Apps []any `json:"apps,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ListAppsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListAppsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListEndpointsResponse struct {
	// The list of endpoints.
	Endpoints []ServingEndpoint `json:"endpoints,omitempty"`
}

// Retrieve the most recent log lines associated with a given serving endpoint's
// served model
type LogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" url:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName string `json:"-" url:"-"`
}

type OpenAiConfig struct {
	// This is the base URL for the OpenAI API (default:
	// "https://api.openai.com/v1"). For Azure OpenAI, this field is required,
	// and is the base URL for the Azure OpenAI API service provided by Azure.
	OpenaiApiBase string `json:"openai_api_base,omitempty"`
	// The Databricks secret key reference for an OpenAI or Azure OpenAI API
	// key.
	OpenaiApiKey string `json:"openai_api_key"`
	// This is an optional field to specify the type of OpenAI API to use. For
	// Azure OpenAI, this field is required, and adjust this parameter to
	// represent the preferred security access validation protocol. For access
	// token validation, use azure. For authentication using Azure Active
	// Directory (Azure AD) use, azuread.
	OpenaiApiType string `json:"openai_api_type,omitempty"`
	// This is an optional field to specify the OpenAI API version. For Azure
	// OpenAI, this field is required, and is the version of the Azure OpenAI
	// service to utilize, specified by a date.
	OpenaiApiVersion string `json:"openai_api_version,omitempty"`
	// This field is only required for Azure OpenAI and is the name of the
	// deployment resource for the Azure OpenAI service.
	OpenaiDeploymentName string `json:"openai_deployment_name,omitempty"`
	// This is an optional field to specify the organization in OpenAI or Azure
	// OpenAI.
	OpenaiOrganization string `json:"openai_organization,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *OpenAiConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OpenAiConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type PaLmConfig struct {
	// The Databricks secret key reference for a PaLM API key.
	PalmApiKey string `json:"palm_api_key"`
}

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	AddTags []EndpointTag `json:"add_tags,omitempty"`
	// List of tag keys to delete
	DeleteTags []string `json:"delete_tags,omitempty"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name string `json:"-" url:"-"`
}

type PayloadTable struct {
	// The name of the payload table.
	Name string `json:"name,omitempty"`
	// The status of the payload table.
	Status string `json:"status,omitempty"`
	// The status message of the payload table.
	StatusMessage string `json:"status_message,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *PayloadTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PayloadTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Update the rate limits of a serving endpoint
type PutRequest struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name string `json:"-" url:"-"`
	// The list of endpoint rate limits.
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
}

type PutResponse struct {
	// The list of endpoint rate limits.
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
}

type QueryEndpointInput struct {
	// Pandas Dataframe input in the records orientation.
	DataframeRecords []any `json:"dataframe_records,omitempty"`
	// Pandas Dataframe input in the split orientation.
	DataframeSplit *DataframeSplitInput `json:"dataframe_split,omitempty"`
	// The extra parameters field used ONLY for __completions, chat,__ and
	// __embeddings external & foundation model__ serving endpoints. This is a
	// map of strings and should only be used with other external/foundation
	// model query fields.
	ExtraParams map[string]string `json:"extra_params,omitempty"`
	// The input string (or array of strings) field used ONLY for __embeddings
	// external & foundation model__ serving endpoints and is the only field
	// (along with extra_params if needed) used by embeddings queries.
	Input any `json:"input,omitempty"`
	// Tensor-based input in columnar format.
	Inputs any `json:"inputs,omitempty"`
	// Tensor-based input in row format.
	Instances []any `json:"instances,omitempty"`
	// The max tokens field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is an integer and should only
	// be used with other chat/completions query fields.
	MaxTokens int `json:"max_tokens,omitempty"`
	// The messages field used ONLY for __chat external & foundation model__
	// serving endpoints. This is a map of strings and should only be used with
	// other chat query fields.
	Messages []ChatMessage `json:"messages,omitempty"`
	// The n (number of candidates) field used ONLY for __completions__ and
	// __chat external & foundation model__ serving endpoints. This is an
	// integer between 1 and 5 with a default of 1 and should only be used with
	// other chat/completions query fields.
	N int `json:"n,omitempty"`
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
	// The prompt string (or array of strings) field used ONLY for __completions
	// external & foundation model__ serving endpoints and should only be used
	// with other completions query fields.
	Prompt any `json:"prompt,omitempty"`
	// The stop sequences field used ONLY for __completions__ and __chat
	// external & foundation model__ serving endpoints. This is a list of
	// strings and should only be used with other chat/completions query fields.
	Stop []string `json:"stop,omitempty"`
	// The stream field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a boolean defaulting to
	// false and should only be used with other chat/completions query fields.
	Stream bool `json:"stream,omitempty"`
	// The temperature field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a float between 0.0 and 2.0
	// with a default of 1.0 and should only be used with other chat/completions
	// query fields.
	Temperature float64 `json:"temperature,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *QueryEndpointInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryEndpointInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QueryEndpointResponse struct {
	// The list of choices returned by the __chat or completions
	// external/foundation model__ serving endpoint.
	Choices []V1ResponseChoiceElement `json:"choices,omitempty"`
	// The timestamp in seconds when the query was created in Unix time returned
	// by a __completions or chat external/foundation model__ serving endpoint.
	Created int64 `json:"created,omitempty"`
	// The list of the embeddings returned by the __embeddings
	// external/foundation model__ serving endpoint.
	Data []EmbeddingsV1ResponseEmbeddingElement `json:"data,omitempty"`
	// The ID of the query that may be returned by a __completions or chat
	// external/foundation model__ serving endpoint.
	Id string `json:"id,omitempty"`
	// The name of the __external/foundation model__ used for querying. This is
	// the name of the model that was specified in the endpoint config.
	Model string `json:"model,omitempty"`
	// The type of object returned by the __external/foundation model__ serving
	// endpoint, one of [text_completion, chat.completion, list (of
	// embeddings)].
	Object QueryEndpointResponseObject `json:"object,omitempty"`
	// The predictions returned by the serving endpoint.
	Predictions []any `json:"predictions,omitempty"`
	// The name of the served model that served the request. This is useful when
	// there are multiple models behind the same endpoint with traffic split.
	ServedModelName string `json:"-" url:"-" header:"served-model-name,omitempty"`
	// The usage object that may be returned by the __external/foundation
	// model__ serving endpoint. This contains information about the number of
	// tokens used in the prompt and response.
	Usage *ExternalModelUsageElement `json:"usage,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *QueryEndpointResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryEndpointResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The type of object returned by the __external/foundation model__ serving
// endpoint, one of [text_completion, chat.completion, list (of embeddings)].
type QueryEndpointResponseObject string

const QueryEndpointResponseObjectChatCompletion QueryEndpointResponseObject = `chat.completion`

const QueryEndpointResponseObjectList QueryEndpointResponseObject = `list`

const QueryEndpointResponseObjectTextCompletion QueryEndpointResponseObject = `text_completion`

// String representation for [fmt.Print]
func (f *QueryEndpointResponseObject) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *QueryEndpointResponseObject) Set(v string) error {
	switch v {
	case `chat.completion`, `list`, `text_completion`:
		*f = QueryEndpointResponseObject(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "chat.completion", "list", "text_completion"`, v)
	}
}

// Type always returns QueryEndpointResponseObject to satisfy [pflag.Value] interface
func (f *QueryEndpointResponseObject) Type() string {
	return "QueryEndpointResponseObject"
}

type RateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls int `json:"calls"`
	// Key field for a serving endpoint rate limit. Currently, only 'user' and
	// 'endpoint' are supported, with 'endpoint' being the default if not
	// specified.
	Key RateLimitKey `json:"key,omitempty"`
	// Renewal period field for a serving endpoint rate limit. Currently, only
	// 'minute' is supported.
	RenewalPeriod RateLimitRenewalPeriod `json:"renewal_period"`
}

// Key field for a serving endpoint rate limit. Currently, only 'user' and
// 'endpoint' are supported, with 'endpoint' being the default if not specified.
type RateLimitKey string

const RateLimitKeyEndpoint RateLimitKey = `endpoint`

const RateLimitKeyUser RateLimitKey = `user`

// String representation for [fmt.Print]
func (f *RateLimitKey) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RateLimitKey) Set(v string) error {
	switch v {
	case `endpoint`, `user`:
		*f = RateLimitKey(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "endpoint", "user"`, v)
	}
}

// Type always returns RateLimitKey to satisfy [pflag.Value] interface
func (f *RateLimitKey) Type() string {
	return "RateLimitKey"
}

// Renewal period field for a serving endpoint rate limit. Currently, only
// 'minute' is supported.
type RateLimitRenewalPeriod string

const RateLimitRenewalPeriodMinute RateLimitRenewalPeriod = `minute`

// String representation for [fmt.Print]
func (f *RateLimitRenewalPeriod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RateLimitRenewalPeriod) Set(v string) error {
	switch v {
	case `minute`:
		*f = RateLimitRenewalPeriod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "minute"`, v)
	}
}

// Type always returns RateLimitRenewalPeriod to satisfy [pflag.Value] interface
func (f *RateLimitRenewalPeriod) Type() string {
	return "RateLimitRenewalPeriod"
}

type Route struct {
	// The name of the served model this route configures traffic for.
	ServedModelName string `json:"served_model_name"`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage int `json:"traffic_percentage"`
}

type ServedEntityInput struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	EntityName string `json:"entity_name,omitempty"`
	// The version of the model in Databricks Model Registry to be served or
	// empty if the entity is a FEATURE_SPEC.
	EntityVersion string `json:"entity_version,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. When an
	// external_model is present, the served entities list can only have one
	// served_entity object. For an existing endpoint with external_model, it
	// can not be updated to an endpoint without external_model. If the endpoint
	// is created without external_model, users cannot update it to add
	// external_model later.
	ExternalModel *ExternalModel `json:"external_model,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to
	// <entity-name>-<entity-version>.
	Name string `json:"name,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size is 0.
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType string `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedEntityInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntityInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedEntityOutput struct {
	// The creation timestamp of the served entity in Unix time.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the served entity.
	Creator string `json:"creator,omitempty"`
	// The name of the entity served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object is given in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	EntityName string `json:"entity_name,omitempty"`
	// The version of the served entity in Databricks Model Registry or empty if
	// the entity is a FEATURE_SPEC.
	EntityVersion string `json:"entity_version,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// The external model that is served. NOTE: Only one of external_model,
	// foundation_model, and (entity_name, entity_version, workload_size,
	// workload_type, and scale_to_zero_enabled) is returned based on the
	// endpoint type.
	ExternalModel *ExternalModel `json:"external_model,omitempty"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version, workload_size,
	// workload_type, and scale_to_zero_enabled) is returned based on the
	// endpoint type.
	FoundationModel *FoundationModel `json:"foundation_model,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`
	// The name of the served entity.
	Name string `json:"name,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`
	// Information corresponding to the state of the served entity.
	State *ServedModelState `json:"state,omitempty"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType string `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedEntityOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntityOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedEntitySpec struct {
	// The name of the entity served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object is given in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	EntityName string `json:"entity_name,omitempty"`
	// The version of the served entity in Databricks Model Registry or empty if
	// the entity is a FEATURE_SPEC.
	EntityVersion string `json:"entity_version,omitempty"`
	// The external model that is served. NOTE: Only one of external_model,
	// foundation_model, and (entity_name, entity_version) is returned based on
	// the endpoint type.
	ExternalModel *ExternalModel `json:"external_model,omitempty"`
	// The foundation model that is served. NOTE: Only one of foundation_model,
	// external_model, and (entity_name, entity_version) is returned based on
	// the endpoint type.
	FoundationModel *FoundationModel `json:"foundation_model,omitempty"`
	// The name of the served entity.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedEntitySpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntitySpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelInput struct {
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this model. Note: this is an
	// experimental feature and subject to change. Example model environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// ARN of the instance profile that the served model will use to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The name of the model in Databricks Model Registry to be served or if the
	// model resides in Unity Catalog, the full name of model, in the form of
	// __catalog_name__.__schema_name__.__model_name__.
	ModelName string `json:"model_name"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `json:"model_version"`
	// The name of a served model. It must be unique across an endpoint. If not
	// specified, this field will default to <model-name>-<model-version>. A
	// served model name can consist of alphanumeric characters, dashes, and
	// underscores.
	Name string `json:"name,omitempty"`
	// Whether the compute resources for the served model should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled"`
	// The workload size of the served model. The workload size corresponds to a
	// range of provisioned concurrency that the compute will autoscale between.
	// A single unit of provisioned concurrency can process one request at a
	// time. Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize ServedModelInputWorkloadSize `json:"workload_size"`
	// The workload type of the served model. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServedModelInputWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedModelInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The workload size of the served model. The workload size corresponds to a
// range of provisioned concurrency that the compute will autoscale between. A
// single unit of provisioned concurrency can process one request at a time.
// Valid workload sizes are "Small" (4 - 4 provisioned concurrency), "Medium" (8
// - 16 provisioned concurrency), and "Large" (16 - 64 provisioned concurrency).
// If scale-to-zero is enabled, the lower bound of the provisioned concurrency
// for each workload size will be 0.
type ServedModelInputWorkloadSize string

const ServedModelInputWorkloadSizeLarge ServedModelInputWorkloadSize = `Large`

const ServedModelInputWorkloadSizeMedium ServedModelInputWorkloadSize = `Medium`

const ServedModelInputWorkloadSizeSmall ServedModelInputWorkloadSize = `Small`

// String representation for [fmt.Print]
func (f *ServedModelInputWorkloadSize) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServedModelInputWorkloadSize) Set(v string) error {
	switch v {
	case `Large`, `Medium`, `Small`:
		*f = ServedModelInputWorkloadSize(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Large", "Medium", "Small"`, v)
	}
}

// Type always returns ServedModelInputWorkloadSize to satisfy [pflag.Value] interface
func (f *ServedModelInputWorkloadSize) Type() string {
	return "ServedModelInputWorkloadSize"
}

// The workload type of the served model. The workload type selects which type
// of compute to use in the endpoint. The default value for this parameter is
// "CPU". For deep learning workloads, GPU acceleration is available by
// selecting workload types like GPU_SMALL and others. See the available [GPU
// types].
//
// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
type ServedModelInputWorkloadType string

const ServedModelInputWorkloadTypeCpu ServedModelInputWorkloadType = `CPU`

const ServedModelInputWorkloadTypeGpuLarge ServedModelInputWorkloadType = `GPU_LARGE`

const ServedModelInputWorkloadTypeGpuMedium ServedModelInputWorkloadType = `GPU_MEDIUM`

const ServedModelInputWorkloadTypeGpuSmall ServedModelInputWorkloadType = `GPU_SMALL`

const ServedModelInputWorkloadTypeMultigpuMedium ServedModelInputWorkloadType = `MULTIGPU_MEDIUM`

// String representation for [fmt.Print]
func (f *ServedModelInputWorkloadType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServedModelInputWorkloadType) Set(v string) error {
	switch v {
	case `CPU`, `GPU_LARGE`, `GPU_MEDIUM`, `GPU_SMALL`, `MULTIGPU_MEDIUM`:
		*f = ServedModelInputWorkloadType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CPU", "GPU_LARGE", "GPU_MEDIUM", "GPU_SMALL", "MULTIGPU_MEDIUM"`, v)
	}
}

// Type always returns ServedModelInputWorkloadType to satisfy [pflag.Value] interface
func (f *ServedModelInputWorkloadType) Type() string {
	return "ServedModelInputWorkloadType"
}

type ServedModelOutput struct {
	// The creation timestamp of the served model in Unix time.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the served model.
	Creator string `json:"creator,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this model. Note: this is an
	// experimental feature and subject to change. Example model environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// ARN of the instance profile that the served model will use to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The name of the model in Databricks Model Registry or the full name of
	// the model in Unity Catalog.
	ModelName string `json:"model_name,omitempty"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `json:"model_version,omitempty"`
	// The name of the served model.
	Name string `json:"name,omitempty"`
	// Whether the compute resources for the Served Model should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`
	// Information corresponding to the state of the Served Model.
	State *ServedModelState `json:"state,omitempty"`
	// The workload size of the served model. The workload size corresponds to a
	// range of provisioned concurrency that the compute will autoscale between.
	// A single unit of provisioned concurrency can process one request at a
	// time. Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). If scale-to-zero is enabled, the lower bound of
	// the provisioned concurrency for each workload size will be 0.
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served model. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType string `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedModelOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelSpec struct {
	// The name of the model in Databricks Model Registry or the full name of
	// the model in Unity Catalog.
	ModelName string `json:"model_name,omitempty"`
	// The version of the model in Databricks Model Registry or Unity Catalog to
	// be served.
	ModelVersion string `json:"model_version,omitempty"`
	// The name of the served model.
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedModelSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServedModelState struct {
	// The state of the served entity deployment. DEPLOYMENT_CREATING indicates
	// that the served entity is not ready yet because the deployment is still
	// being created (i.e container image is building, model server is deploying
	// for the first time, etc.). DEPLOYMENT_RECOVERING indicates that the
	// served entity was previously in a ready state but no longer is and is
	// attempting to recover. DEPLOYMENT_READY indicates that the served entity
	// is ready to receive traffic. DEPLOYMENT_FAILED indicates that there was
	// an error trying to bring up the served entity (e.g container image build
	// failed, the model server failed to start due to a model loading error,
	// etc.) DEPLOYMENT_ABORTED indicates that the deployment was terminated
	// likely due to a failure in bringing up another served entity under the
	// same endpoint and config version.
	Deployment ServedModelStateDeployment `json:"deployment,omitempty"`
	// More information about the state of the served entity, if available.
	DeploymentStateMessage string `json:"deployment_state_message,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServedModelState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The state of the served entity deployment. DEPLOYMENT_CREATING indicates that
// the served entity is not ready yet because the deployment is still being
// created (i.e container image is building, model server is deploying for the
// first time, etc.). DEPLOYMENT_RECOVERING indicates that the served entity was
// previously in a ready state but no longer is and is attempting to recover.
// DEPLOYMENT_READY indicates that the served entity is ready to receive
// traffic. DEPLOYMENT_FAILED indicates that there was an error trying to bring
// up the served entity (e.g container image build failed, the model server
// failed to start due to a model loading error, etc.) DEPLOYMENT_ABORTED
// indicates that the deployment was terminated likely due to a failure in
// bringing up another served entity under the same endpoint and config version.
type ServedModelStateDeployment string

const ServedModelStateDeploymentAborted ServedModelStateDeployment = `DEPLOYMENT_ABORTED`

const ServedModelStateDeploymentCreating ServedModelStateDeployment = `DEPLOYMENT_CREATING`

const ServedModelStateDeploymentFailed ServedModelStateDeployment = `DEPLOYMENT_FAILED`

const ServedModelStateDeploymentReady ServedModelStateDeployment = `DEPLOYMENT_READY`

const ServedModelStateDeploymentRecovering ServedModelStateDeployment = `DEPLOYMENT_RECOVERING`

// String representation for [fmt.Print]
func (f *ServedModelStateDeployment) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServedModelStateDeployment) Set(v string) error {
	switch v {
	case `DEPLOYMENT_ABORTED`, `DEPLOYMENT_CREATING`, `DEPLOYMENT_FAILED`, `DEPLOYMENT_READY`, `DEPLOYMENT_RECOVERING`:
		*f = ServedModelStateDeployment(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DEPLOYMENT_ABORTED", "DEPLOYMENT_CREATING", "DEPLOYMENT_FAILED", "DEPLOYMENT_READY", "DEPLOYMENT_RECOVERING"`, v)
	}
}

// Type always returns ServedModelStateDeployment to satisfy [pflag.Value] interface
func (f *ServedModelStateDeployment) Type() string {
	return "ServedModelStateDeployment"
}

type ServerLogsResponse struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	Logs string `json:"logs"`
}

type ServingEndpoint struct {
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigSummary `json:"config,omitempty"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the serving endpoint.
	Creator string `json:"creator,omitempty"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id string `json:"id,omitempty"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The name of the serving endpoint.
	Name string `json:"name,omitempty"`
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState `json:"state,omitempty"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `json:"tags,omitempty"`
	// The task type of the serving endpoint.
	Task string `json:"task,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointAccessControlResponse struct {
	// All permissions.
	AllPermissions []ServingEndpointPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointDetailed struct {
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigOutput `json:"config,omitempty"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the serving endpoint.
	Creator string `json:"creator,omitempty"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id string `json:"id,omitempty"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The name of the serving endpoint.
	Name string `json:"name,omitempty"`
	// The config that the endpoint is attempting to update to.
	PendingConfig *EndpointPendingConfig `json:"pending_config,omitempty"`
	// The permission level of the principal making the request.
	PermissionLevel ServingEndpointDetailedPermissionLevel `json:"permission_level,omitempty"`
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState `json:"state,omitempty"`
	// Tags attached to the serving endpoint.
	Tags []EndpointTag `json:"tags,omitempty"`
	// The task type of the serving endpoint.
	Task string `json:"task,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointDetailed) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointDetailed) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// The permission level of the principal making the request.
type ServingEndpointDetailedPermissionLevel string

const ServingEndpointDetailedPermissionLevelCanManage ServingEndpointDetailedPermissionLevel = `CAN_MANAGE`

const ServingEndpointDetailedPermissionLevelCanQuery ServingEndpointDetailedPermissionLevel = `CAN_QUERY`

const ServingEndpointDetailedPermissionLevelCanView ServingEndpointDetailedPermissionLevel = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *ServingEndpointDetailedPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServingEndpointDetailedPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`:
		*f = ServingEndpointDetailedPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW"`, v)
	}
}

// Type always returns ServingEndpointDetailedPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointDetailedPermissionLevel) Type() string {
	return "ServingEndpointDetailedPermissionLevel"
}

type ServingEndpointPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type ServingEndpointPermissionLevel string

const ServingEndpointPermissionLevelCanManage ServingEndpointPermissionLevel = `CAN_MANAGE`

const ServingEndpointPermissionLevelCanQuery ServingEndpointPermissionLevel = `CAN_QUERY`

const ServingEndpointPermissionLevelCanView ServingEndpointPermissionLevel = `CAN_VIEW`

// String representation for [fmt.Print]
func (f *ServingEndpointPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServingEndpointPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`, `CAN_VIEW`:
		*f = ServingEndpointPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY", "CAN_VIEW"`, v)
	}
}

// Type always returns ServingEndpointPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointPermissionLevel) Type() string {
	return "ServingEndpointPermissionLevel"
}

type ServingEndpointPermissions struct {
	AccessControlList []ServingEndpointAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointPermissionsDescription struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *ServingEndpointPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ServingEndpointPermissionsRequest struct {
	AccessControlList []ServingEndpointAccessControlRequest `json:"access_control_list,omitempty"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served entity.
	Routes []Route `json:"routes,omitempty"`
}

type V1ResponseChoiceElement struct {
	// The finish reason returned by the endpoint.
	FinishReason string `json:"finishReason,omitempty"`
	// The index of the choice in the __chat or completions__ response.
	Index int `json:"index,omitempty"`
	// The logprobs returned only by the __completions__ endpoint.
	Logprobs int `json:"logprobs,omitempty"`
	// The message response from the __chat__ endpoint.
	Message *ChatMessage `json:"message,omitempty"`
	// The text response from the __completions__ endpoint.
	Text string `json:"text,omitempty"`

	ForceSendFields []string `json:"-"`
}

func (s *V1ResponseChoiceElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s V1ResponseChoiceElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}
