// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package servingpb

import (
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type Ai21LabsConfigPb struct {
	Ai21labsApiKey          string `json:"ai21labs_api_key,omitempty"`
	Ai21labsApiKeyPlaintext string `json:"ai21labs_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *Ai21LabsConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st Ai21LabsConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AiGatewayConfigPb struct {
	FallbackConfig       *FallbackConfigPb                `json:"fallback_config,omitempty"`
	Guardrails           *AiGatewayGuardrailsPb           `json:"guardrails,omitempty"`
	InferenceTableConfig *AiGatewayInferenceTableConfigPb `json:"inference_table_config,omitempty"`
	RateLimits           []AiGatewayRateLimitPb           `json:"rate_limits,omitempty"`
	UsageTrackingConfig  *AiGatewayUsageTrackingConfigPb  `json:"usage_tracking_config,omitempty"`
}

type AiGatewayGuardrailParametersPb struct {
	InvalidKeywords []string                         `json:"invalid_keywords,omitempty"`
	Pii             *AiGatewayGuardrailPiiBehaviorPb `json:"pii,omitempty"`
	Safety          bool                             `json:"safety,omitempty"`
	ValidTopics     []string                         `json:"valid_topics,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AiGatewayGuardrailParametersPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AiGatewayGuardrailParametersPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AiGatewayGuardrailPiiBehaviorPb struct {
	Behavior AiGatewayGuardrailPiiBehaviorBehaviorPb `json:"behavior,omitempty"`
}

type AiGatewayGuardrailPiiBehaviorBehaviorPb string

const AiGatewayGuardrailPiiBehaviorBehaviorBlock AiGatewayGuardrailPiiBehaviorBehaviorPb = `BLOCK`

const AiGatewayGuardrailPiiBehaviorBehaviorNone AiGatewayGuardrailPiiBehaviorBehaviorPb = `NONE`

type AiGatewayGuardrailsPb struct {
	Input  *AiGatewayGuardrailParametersPb `json:"input,omitempty"`
	Output *AiGatewayGuardrailParametersPb `json:"output,omitempty"`
}

type AiGatewayInferenceTableConfigPb struct {
	CatalogName     string `json:"catalog_name,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	SchemaName      string `json:"schema_name,omitempty"`
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AiGatewayInferenceTableConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AiGatewayInferenceTableConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AiGatewayRateLimitPb struct {
	Calls         int64                             `json:"calls,omitempty"`
	Key           AiGatewayRateLimitKeyPb           `json:"key,omitempty"`
	Principal     string                            `json:"principal,omitempty"`
	RenewalPeriod AiGatewayRateLimitRenewalPeriodPb `json:"renewal_period"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AiGatewayRateLimitPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AiGatewayRateLimitPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AiGatewayRateLimitKeyPb string

const AiGatewayRateLimitKeyEndpoint AiGatewayRateLimitKeyPb = `endpoint`

const AiGatewayRateLimitKeyServicePrincipal AiGatewayRateLimitKeyPb = `service_principal`

const AiGatewayRateLimitKeyUser AiGatewayRateLimitKeyPb = `user`

const AiGatewayRateLimitKeyUserGroup AiGatewayRateLimitKeyPb = `user_group`

type AiGatewayRateLimitRenewalPeriodPb string

const AiGatewayRateLimitRenewalPeriodMinute AiGatewayRateLimitRenewalPeriodPb = `minute`

type AiGatewayUsageTrackingConfigPb struct {
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AiGatewayUsageTrackingConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AiGatewayUsageTrackingConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AmazonBedrockConfigPb struct {
	AwsAccessKeyId              string                               `json:"aws_access_key_id,omitempty"`
	AwsAccessKeyIdPlaintext     string                               `json:"aws_access_key_id_plaintext,omitempty"`
	AwsRegion                   string                               `json:"aws_region"`
	AwsSecretAccessKey          string                               `json:"aws_secret_access_key,omitempty"`
	AwsSecretAccessKeyPlaintext string                               `json:"aws_secret_access_key_plaintext,omitempty"`
	BedrockProvider             AmazonBedrockConfigBedrockProviderPb `json:"bedrock_provider"`
	InstanceProfileArn          string                               `json:"instance_profile_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AmazonBedrockConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AmazonBedrockConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AmazonBedrockConfigBedrockProviderPb string

const AmazonBedrockConfigBedrockProviderAi21labs AmazonBedrockConfigBedrockProviderPb = `ai21labs`

const AmazonBedrockConfigBedrockProviderAmazon AmazonBedrockConfigBedrockProviderPb = `amazon`

const AmazonBedrockConfigBedrockProviderAnthropic AmazonBedrockConfigBedrockProviderPb = `anthropic`

const AmazonBedrockConfigBedrockProviderCohere AmazonBedrockConfigBedrockProviderPb = `cohere`

type AnthropicConfigPb struct {
	AnthropicApiKey          string `json:"anthropic_api_key,omitempty"`
	AnthropicApiKeyPlaintext string `json:"anthropic_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AnthropicConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AnthropicConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ApiKeyAuthPb struct {
	Key            string `json:"key"`
	Value          string `json:"value,omitempty"`
	ValuePlaintext string `json:"value_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ApiKeyAuthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ApiKeyAuthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AutoCaptureConfigInputPb struct {
	CatalogName     string `json:"catalog_name,omitempty"`
	Enabled         bool   `json:"enabled,omitempty"`
	SchemaName      string `json:"schema_name,omitempty"`
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AutoCaptureConfigInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AutoCaptureConfigInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AutoCaptureConfigOutputPb struct {
	CatalogName     string              `json:"catalog_name,omitempty"`
	Enabled         bool                `json:"enabled,omitempty"`
	SchemaName      string              `json:"schema_name,omitempty"`
	State           *AutoCaptureStatePb `json:"state,omitempty"`
	TableNamePrefix string              `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AutoCaptureConfigOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AutoCaptureConfigOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AutoCaptureStatePb struct {
	PayloadTable *PayloadTablePb `json:"payload_table,omitempty"`
}

type BearerTokenAuthPb struct {
	Token          string `json:"token,omitempty"`
	TokenPlaintext string `json:"token_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *BearerTokenAuthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st BearerTokenAuthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type BuildLogsRequestPb struct {
	Name            string `json:"-" url:"-"`
	ServedModelName string `json:"-" url:"-"`
}

type BuildLogsResponsePb struct {
	Logs string `json:"logs"`
}

type ChatMessagePb struct {
	Content string            `json:"content,omitempty"`
	Role    ChatMessageRolePb `json:"role,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ChatMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ChatMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ChatMessageRolePb string

const ChatMessageRoleAssistant ChatMessageRolePb = `assistant`

const ChatMessageRoleSystem ChatMessageRolePb = `system`

const ChatMessageRoleUser ChatMessageRolePb = `user`

type CohereConfigPb struct {
	CohereApiBase         string `json:"cohere_api_base,omitempty"`
	CohereApiKey          string `json:"cohere_api_key,omitempty"`
	CohereApiKeyPlaintext string `json:"cohere_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CohereConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CohereConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePtEndpointRequestPb struct {
	AiGateway      *AiGatewayConfigPb     `json:"ai_gateway,omitempty"`
	BudgetPolicyId string                 `json:"budget_policy_id,omitempty"`
	Config         PtEndpointCoreConfigPb `json:"config"`
	Name           string                 `json:"name"`
	Tags           []EndpointTagPb        `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePtEndpointRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePtEndpointRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateServingEndpointPb struct {
	AiGateway      *AiGatewayConfigPb         `json:"ai_gateway,omitempty"`
	BudgetPolicyId string                     `json:"budget_policy_id,omitempty"`
	Config         *EndpointCoreConfigInputPb `json:"config,omitempty"`
	Description    string                     `json:"description,omitempty"`
	Name           string                     `json:"name"`
	RateLimits     []RateLimitPb              `json:"rate_limits,omitempty"`
	RouteOptimized bool                       `json:"route_optimized,omitempty"`
	Tags           []EndpointTagPb            `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreateServingEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreateServingEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CustomProviderConfigPb struct {
	ApiKeyAuth        *ApiKeyAuthPb      `json:"api_key_auth,omitempty"`
	BearerTokenAuth   *BearerTokenAuthPb `json:"bearer_token_auth,omitempty"`
	CustomProviderUrl string             `json:"custom_provider_url"`
}

type DataPlaneInfoPb struct {
	AuthorizationDetails string `json:"authorization_details,omitempty"`
	EndpointUrl          string `json:"endpoint_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DataPlaneInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DataPlaneInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatabricksModelServingConfigPb struct {
	DatabricksApiToken          string `json:"databricks_api_token,omitempty"`
	DatabricksApiTokenPlaintext string `json:"databricks_api_token_plaintext,omitempty"`
	DatabricksWorkspaceUrl      string `json:"databricks_workspace_url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DatabricksModelServingConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DatabricksModelServingConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DataframeSplitInputPb struct {
	Columns []any `json:"columns,omitempty"`
	Data    []any `json:"data,omitempty"`
	Index   []int `json:"index,omitempty"`
}

type DeleteServingEndpointRequestPb struct {
	Name string `json:"-" url:"-"`
}

type EmbeddingsV1ResponseEmbeddingElementPb struct {
	Embedding []float64                                    `json:"embedding,omitempty"`
	Index     int                                          `json:"index,omitempty"`
	Object    EmbeddingsV1ResponseEmbeddingElementObjectPb `json:"object,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EmbeddingsV1ResponseEmbeddingElementPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EmbeddingsV1ResponseEmbeddingElementPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EmbeddingsV1ResponseEmbeddingElementObjectPb string

const EmbeddingsV1ResponseEmbeddingElementObjectEmbedding EmbeddingsV1ResponseEmbeddingElementObjectPb = `embedding`

type EndpointCoreConfigInputPb struct {
	AutoCaptureConfig *AutoCaptureConfigInputPb `json:"auto_capture_config,omitempty"`
	Name              string                    `json:"-" url:"-"`
	ServedEntities    []ServedEntityInputPb     `json:"served_entities,omitempty"`
	ServedModels      []ServedModelInputPb      `json:"served_models,omitempty"`
	TrafficConfig     *TrafficConfigPb          `json:"traffic_config,omitempty"`
}

type EndpointCoreConfigOutputPb struct {
	AutoCaptureConfig *AutoCaptureConfigOutputPb `json:"auto_capture_config,omitempty"`
	ConfigVersion     int64                      `json:"config_version,omitempty"`
	ServedEntities    []ServedEntityOutputPb     `json:"served_entities,omitempty"`
	ServedModels      []ServedModelOutputPb      `json:"served_models,omitempty"`
	TrafficConfig     *TrafficConfigPb           `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EndpointCoreConfigOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EndpointCoreConfigOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointCoreConfigSummaryPb struct {
	ServedEntities []ServedEntitySpecPb `json:"served_entities,omitempty"`
	ServedModels   []ServedModelSpecPb  `json:"served_models,omitempty"`
}

type EndpointPendingConfigPb struct {
	AutoCaptureConfig *AutoCaptureConfigOutputPb `json:"auto_capture_config,omitempty"`
	ConfigVersion     int                        `json:"config_version,omitempty"`
	ServedEntities    []ServedEntityOutputPb     `json:"served_entities,omitempty"`
	ServedModels      []ServedModelOutputPb      `json:"served_models,omitempty"`
	StartTime         int64                      `json:"start_time,omitempty"`
	TrafficConfig     *TrafficConfigPb           `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EndpointPendingConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EndpointPendingConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointStatePb struct {
	ConfigUpdate EndpointStateConfigUpdatePb `json:"config_update,omitempty"`
	Ready        EndpointStateReadyPb        `json:"ready,omitempty"`
}

type EndpointStateConfigUpdatePb string

const EndpointStateConfigUpdateInProgress EndpointStateConfigUpdatePb = `IN_PROGRESS`

const EndpointStateConfigUpdateNotUpdating EndpointStateConfigUpdatePb = `NOT_UPDATING`

const EndpointStateConfigUpdateUpdateCanceled EndpointStateConfigUpdatePb = `UPDATE_CANCELED`

const EndpointStateConfigUpdateUpdateFailed EndpointStateConfigUpdatePb = `UPDATE_FAILED`

type EndpointStateReadyPb string

const EndpointStateReadyNotReady EndpointStateReadyPb = `NOT_READY`

const EndpointStateReadyReady EndpointStateReadyPb = `READY`

type EndpointTagPb struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EndpointTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EndpointTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointTagsPb struct {
	Tags []EndpointTagPb `json:"tags,omitempty"`
}

type ExportMetricsRequestPb struct {
	Name string `json:"-" url:"-"`
}

type ExportMetricsResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

type ExternalFunctionRequestPb struct {
	ConnectionName string                              `json:"connection_name"`
	Headers        string                              `json:"headers,omitempty"`
	Json           string                              `json:"json,omitempty"`
	Method         ExternalFunctionRequestHttpMethodPb `json:"method"`
	Params         string                              `json:"params,omitempty"`
	Path           string                              `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalFunctionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalFunctionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalFunctionRequestHttpMethodPb string

const ExternalFunctionRequestHttpMethodDelete ExternalFunctionRequestHttpMethodPb = `DELETE`

const ExternalFunctionRequestHttpMethodGet ExternalFunctionRequestHttpMethodPb = `GET`

const ExternalFunctionRequestHttpMethodPatch ExternalFunctionRequestHttpMethodPb = `PATCH`

const ExternalFunctionRequestHttpMethodPost ExternalFunctionRequestHttpMethodPb = `POST`

const ExternalFunctionRequestHttpMethodPut ExternalFunctionRequestHttpMethodPb = `PUT`

type ExternalModelPb struct {
	Ai21labsConfig               *Ai21LabsConfigPb               `json:"ai21labs_config,omitempty"`
	AmazonBedrockConfig          *AmazonBedrockConfigPb          `json:"amazon_bedrock_config,omitempty"`
	AnthropicConfig              *AnthropicConfigPb              `json:"anthropic_config,omitempty"`
	CohereConfig                 *CohereConfigPb                 `json:"cohere_config,omitempty"`
	CustomProviderConfig         *CustomProviderConfigPb         `json:"custom_provider_config,omitempty"`
	DatabricksModelServingConfig *DatabricksModelServingConfigPb `json:"databricks_model_serving_config,omitempty"`
	GoogleCloudVertexAiConfig    *GoogleCloudVertexAiConfigPb    `json:"google_cloud_vertex_ai_config,omitempty"`
	Name                         string                          `json:"name"`
	OpenaiConfig                 *OpenAiConfigPb                 `json:"openai_config,omitempty"`
	PalmConfig                   *PaLmConfigPb                   `json:"palm_config,omitempty"`
	Provider                     ExternalModelProviderPb         `json:"provider"`
	Task                         string                          `json:"task"`
}

type ExternalModelProviderPb string

const ExternalModelProviderAi21labs ExternalModelProviderPb = `ai21labs`

const ExternalModelProviderAmazonBedrock ExternalModelProviderPb = `amazon-bedrock`

const ExternalModelProviderAnthropic ExternalModelProviderPb = `anthropic`

const ExternalModelProviderCohere ExternalModelProviderPb = `cohere`

const ExternalModelProviderCustom ExternalModelProviderPb = `custom`

const ExternalModelProviderDatabricksModelServing ExternalModelProviderPb = `databricks-model-serving`

const ExternalModelProviderGoogleCloudVertexAi ExternalModelProviderPb = `google-cloud-vertex-ai`

const ExternalModelProviderOpenai ExternalModelProviderPb = `openai`

const ExternalModelProviderPalm ExternalModelProviderPb = `palm`

type ExternalModelUsageElementPb struct {
	CompletionTokens int `json:"completion_tokens,omitempty"`
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ExternalModelUsageElementPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ExternalModelUsageElementPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FallbackConfigPb struct {
	Enabled bool `json:"enabled"`
}

type FoundationModelPb struct {
	Description string `json:"description,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Docs        string `json:"docs,omitempty"`
	Name        string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FoundationModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FoundationModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetOpenApiRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GetOpenApiResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

type GetServingEndpointPermissionLevelsRequestPb struct {
	ServingEndpointId string `json:"-" url:"-"`
}

type GetServingEndpointPermissionLevelsResponsePb struct {
	PermissionLevels []ServingEndpointPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetServingEndpointPermissionsRequestPb struct {
	ServingEndpointId string `json:"-" url:"-"`
}

type GetServingEndpointRequestPb struct {
	Name string `json:"-" url:"-"`
}

type GoogleCloudVertexAiConfigPb struct {
	PrivateKey          string `json:"private_key,omitempty"`
	PrivateKeyPlaintext string `json:"private_key_plaintext,omitempty"`
	ProjectId           string `json:"project_id"`
	Region              string `json:"region"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GoogleCloudVertexAiConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GoogleCloudVertexAiConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type HttpRequestResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

type ListEndpointsResponsePb struct {
	Endpoints []ServingEndpointPb `json:"endpoints,omitempty"`
}

type ListServingEndpointsRequestPb struct {
}

type LogsRequestPb struct {
	Name            string `json:"-" url:"-"`
	ServedModelName string `json:"-" url:"-"`
}

type ModelDataPlaneInfoPb struct {
	QueryInfo *DataPlaneInfoPb `json:"query_info,omitempty"`
}

type OpenAiConfigPb struct {
	MicrosoftEntraClientId              string `json:"microsoft_entra_client_id,omitempty"`
	MicrosoftEntraClientSecret          string `json:"microsoft_entra_client_secret,omitempty"`
	MicrosoftEntraClientSecretPlaintext string `json:"microsoft_entra_client_secret_plaintext,omitempty"`
	MicrosoftEntraTenantId              string `json:"microsoft_entra_tenant_id,omitempty"`
	OpenaiApiBase                       string `json:"openai_api_base,omitempty"`
	OpenaiApiKey                        string `json:"openai_api_key,omitempty"`
	OpenaiApiKeyPlaintext               string `json:"openai_api_key_plaintext,omitempty"`
	OpenaiApiType                       string `json:"openai_api_type,omitempty"`
	OpenaiApiVersion                    string `json:"openai_api_version,omitempty"`
	OpenaiDeploymentName                string `json:"openai_deployment_name,omitempty"`
	OpenaiOrganization                  string `json:"openai_organization,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OpenAiConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OpenAiConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PaLmConfigPb struct {
	PalmApiKey          string `json:"palm_api_key,omitempty"`
	PalmApiKeyPlaintext string `json:"palm_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PaLmConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PaLmConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PatchServingEndpointTagsPb struct {
	AddTags    []EndpointTagPb `json:"add_tags,omitempty"`
	DeleteTags []string        `json:"delete_tags,omitempty"`
	Name       string          `json:"-" url:"-"`
}

type PayloadTablePb struct {
	Name          string `json:"name,omitempty"`
	Status        string `json:"status,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PayloadTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PayloadTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PtEndpointCoreConfigPb struct {
	ServedEntities []PtServedModelPb `json:"served_entities,omitempty"`
	TrafficConfig  *TrafficConfigPb  `json:"traffic_config,omitempty"`
}

type PtServedModelPb struct {
	EntityName            string `json:"entity_name"`
	EntityVersion         string `json:"entity_version,omitempty"`
	Name                  string `json:"name,omitempty"`
	ProvisionedModelUnits int64  `json:"provisioned_model_units"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PtServedModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PtServedModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PutAiGatewayRequestPb struct {
	FallbackConfig       *FallbackConfigPb                `json:"fallback_config,omitempty"`
	Guardrails           *AiGatewayGuardrailsPb           `json:"guardrails,omitempty"`
	InferenceTableConfig *AiGatewayInferenceTableConfigPb `json:"inference_table_config,omitempty"`
	Name                 string                           `json:"-" url:"-"`
	RateLimits           []AiGatewayRateLimitPb           `json:"rate_limits,omitempty"`
	UsageTrackingConfig  *AiGatewayUsageTrackingConfigPb  `json:"usage_tracking_config,omitempty"`
}

type PutAiGatewayResponsePb struct {
	FallbackConfig       *FallbackConfigPb                `json:"fallback_config,omitempty"`
	Guardrails           *AiGatewayGuardrailsPb           `json:"guardrails,omitempty"`
	InferenceTableConfig *AiGatewayInferenceTableConfigPb `json:"inference_table_config,omitempty"`
	RateLimits           []AiGatewayRateLimitPb           `json:"rate_limits,omitempty"`
	UsageTrackingConfig  *AiGatewayUsageTrackingConfigPb  `json:"usage_tracking_config,omitempty"`
}

type PutRequestPb struct {
	Name       string        `json:"-" url:"-"`
	RateLimits []RateLimitPb `json:"rate_limits,omitempty"`
}

type PutResponsePb struct {
	RateLimits []RateLimitPb `json:"rate_limits,omitempty"`
}

type QueryEndpointInputPb struct {
	DataframeRecords []any                  `json:"dataframe_records,omitempty"`
	DataframeSplit   *DataframeSplitInputPb `json:"dataframe_split,omitempty"`
	ExtraParams      map[string]string      `json:"extra_params,omitempty"`
	Input            any                    `json:"input,omitempty"`
	Inputs           any                    `json:"inputs,omitempty"`
	Instances        []any                  `json:"instances,omitempty"`
	MaxTokens        int                    `json:"max_tokens,omitempty"`
	Messages         []ChatMessagePb        `json:"messages,omitempty"`
	N                int                    `json:"n,omitempty"`
	Name             string                 `json:"-" url:"-"`
	Prompt           any                    `json:"prompt,omitempty"`
	Stop             []string               `json:"stop,omitempty"`
	Stream           bool                   `json:"stream,omitempty"`
	Temperature      float64                `json:"temperature,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryEndpointInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryEndpointInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryEndpointResponsePb struct {
	Choices         []V1ResponseChoiceElementPb              `json:"choices,omitempty"`
	Created         int64                                    `json:"created,omitempty"`
	Data            []EmbeddingsV1ResponseEmbeddingElementPb `json:"data,omitempty"`
	Id              string                                   `json:"id,omitempty"`
	Model           string                                   `json:"model,omitempty"`
	Object          QueryEndpointResponseObjectPb            `json:"object,omitempty"`
	Predictions     []any                                    `json:"predictions,omitempty"`
	ServedModelName string                                   `json:"-" url:"-" header:"served-model-name,omitempty"`
	Usage           *ExternalModelUsageElementPb             `json:"usage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *QueryEndpointResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st QueryEndpointResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryEndpointResponseObjectPb string

const QueryEndpointResponseObjectChatCompletion QueryEndpointResponseObjectPb = `chat.completion`

const QueryEndpointResponseObjectList QueryEndpointResponseObjectPb = `list`

const QueryEndpointResponseObjectTextCompletion QueryEndpointResponseObjectPb = `text_completion`

type RateLimitPb struct {
	Calls         int64                    `json:"calls"`
	Key           RateLimitKeyPb           `json:"key,omitempty"`
	RenewalPeriod RateLimitRenewalPeriodPb `json:"renewal_period"`
}

type RateLimitKeyPb string

const RateLimitKeyEndpoint RateLimitKeyPb = `endpoint`

const RateLimitKeyUser RateLimitKeyPb = `user`

type RateLimitRenewalPeriodPb string

const RateLimitRenewalPeriodMinute RateLimitRenewalPeriodPb = `minute`

type RoutePb struct {
	ServedEntityName  string `json:"served_entity_name,omitempty"`
	ServedModelName   string `json:"served_model_name,omitempty"`
	TrafficPercentage int    `json:"traffic_percentage"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RoutePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RoutePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedEntityInputPb struct {
	EntityName                string                     `json:"entity_name,omitempty"`
	EntityVersion             string                     `json:"entity_version,omitempty"`
	EnvironmentVars           map[string]string          `json:"environment_vars,omitempty"`
	ExternalModel             *ExternalModelPb           `json:"external_model,omitempty"`
	InstanceProfileArn        string                     `json:"instance_profile_arn,omitempty"`
	MaxProvisionedConcurrency int                        `json:"max_provisioned_concurrency,omitempty"`
	MaxProvisionedThroughput  int                        `json:"max_provisioned_throughput,omitempty"`
	MinProvisionedConcurrency int                        `json:"min_provisioned_concurrency,omitempty"`
	MinProvisionedThroughput  int                        `json:"min_provisioned_throughput,omitempty"`
	Name                      string                     `json:"name,omitempty"`
	ProvisionedModelUnits     int64                      `json:"provisioned_model_units,omitempty"`
	ScaleToZeroEnabled        bool                       `json:"scale_to_zero_enabled,omitempty"`
	WorkloadSize              string                     `json:"workload_size,omitempty"`
	WorkloadType              ServingModelWorkloadTypePb `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServedEntityInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServedEntityInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedEntityOutputPb struct {
	CreationTimestamp         int64                      `json:"creation_timestamp,omitempty"`
	Creator                   string                     `json:"creator,omitempty"`
	EntityName                string                     `json:"entity_name,omitempty"`
	EntityVersion             string                     `json:"entity_version,omitempty"`
	EnvironmentVars           map[string]string          `json:"environment_vars,omitempty"`
	ExternalModel             *ExternalModelPb           `json:"external_model,omitempty"`
	FoundationModel           *FoundationModelPb         `json:"foundation_model,omitempty"`
	InstanceProfileArn        string                     `json:"instance_profile_arn,omitempty"`
	MaxProvisionedConcurrency int                        `json:"max_provisioned_concurrency,omitempty"`
	MaxProvisionedThroughput  int                        `json:"max_provisioned_throughput,omitempty"`
	MinProvisionedConcurrency int                        `json:"min_provisioned_concurrency,omitempty"`
	MinProvisionedThroughput  int                        `json:"min_provisioned_throughput,omitempty"`
	Name                      string                     `json:"name,omitempty"`
	ProvisionedModelUnits     int64                      `json:"provisioned_model_units,omitempty"`
	ScaleToZeroEnabled        bool                       `json:"scale_to_zero_enabled,omitempty"`
	State                     *ServedModelStatePb        `json:"state,omitempty"`
	WorkloadSize              string                     `json:"workload_size,omitempty"`
	WorkloadType              ServingModelWorkloadTypePb `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServedEntityOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServedEntityOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedEntitySpecPb struct {
	EntityName      string             `json:"entity_name,omitempty"`
	EntityVersion   string             `json:"entity_version,omitempty"`
	ExternalModel   *ExternalModelPb   `json:"external_model,omitempty"`
	FoundationModel *FoundationModelPb `json:"foundation_model,omitempty"`
	Name            string             `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServedEntitySpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServedEntitySpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedModelInputPb struct {
	EnvironmentVars           map[string]string              `json:"environment_vars,omitempty"`
	InstanceProfileArn        string                         `json:"instance_profile_arn,omitempty"`
	MaxProvisionedConcurrency int                            `json:"max_provisioned_concurrency,omitempty"`
	MaxProvisionedThroughput  int                            `json:"max_provisioned_throughput,omitempty"`
	MinProvisionedConcurrency int                            `json:"min_provisioned_concurrency,omitempty"`
	MinProvisionedThroughput  int                            `json:"min_provisioned_throughput,omitempty"`
	ModelName                 string                         `json:"model_name"`
	ModelVersion              string                         `json:"model_version"`
	Name                      string                         `json:"name,omitempty"`
	ProvisionedModelUnits     int64                          `json:"provisioned_model_units,omitempty"`
	ScaleToZeroEnabled        bool                           `json:"scale_to_zero_enabled"`
	WorkloadSize              string                         `json:"workload_size,omitempty"`
	WorkloadType              ServedModelInputWorkloadTypePb `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServedModelInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServedModelInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedModelInputWorkloadTypePb string

const ServedModelInputWorkloadTypeCpu ServedModelInputWorkloadTypePb = `CPU`

const ServedModelInputWorkloadTypeGpuLarge ServedModelInputWorkloadTypePb = `GPU_LARGE`

const ServedModelInputWorkloadTypeGpuMedium ServedModelInputWorkloadTypePb = `GPU_MEDIUM`

const ServedModelInputWorkloadTypeGpuSmall ServedModelInputWorkloadTypePb = `GPU_SMALL`

const ServedModelInputWorkloadTypeMultigpuMedium ServedModelInputWorkloadTypePb = `MULTIGPU_MEDIUM`

type ServedModelOutputPb struct {
	CreationTimestamp         int64                      `json:"creation_timestamp,omitempty"`
	Creator                   string                     `json:"creator,omitempty"`
	EnvironmentVars           map[string]string          `json:"environment_vars,omitempty"`
	InstanceProfileArn        string                     `json:"instance_profile_arn,omitempty"`
	MaxProvisionedConcurrency int                        `json:"max_provisioned_concurrency,omitempty"`
	MinProvisionedConcurrency int                        `json:"min_provisioned_concurrency,omitempty"`
	ModelName                 string                     `json:"model_name,omitempty"`
	ModelVersion              string                     `json:"model_version,omitempty"`
	Name                      string                     `json:"name,omitempty"`
	ProvisionedModelUnits     int64                      `json:"provisioned_model_units,omitempty"`
	ScaleToZeroEnabled        bool                       `json:"scale_to_zero_enabled,omitempty"`
	State                     *ServedModelStatePb        `json:"state,omitempty"`
	WorkloadSize              string                     `json:"workload_size,omitempty"`
	WorkloadType              ServingModelWorkloadTypePb `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServedModelOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServedModelOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedModelSpecPb struct {
	ModelName    string `json:"model_name,omitempty"`
	ModelVersion string `json:"model_version,omitempty"`
	Name         string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServedModelSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServedModelSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedModelStatePb struct {
	Deployment             ServedModelStateDeploymentPb `json:"deployment,omitempty"`
	DeploymentStateMessage string                       `json:"deployment_state_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServedModelStatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServedModelStatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedModelStateDeploymentPb string

const ServedModelStateDeploymentAborted ServedModelStateDeploymentPb = `DEPLOYMENT_ABORTED`

const ServedModelStateDeploymentCreating ServedModelStateDeploymentPb = `DEPLOYMENT_CREATING`

const ServedModelStateDeploymentFailed ServedModelStateDeploymentPb = `DEPLOYMENT_FAILED`

const ServedModelStateDeploymentReady ServedModelStateDeploymentPb = `DEPLOYMENT_READY`

const ServedModelStateDeploymentRecovering ServedModelStateDeploymentPb = `DEPLOYMENT_RECOVERING`

type ServerLogsResponsePb struct {
	Logs string `json:"logs"`
}

type ServingEndpointPb struct {
	AiGateway            *AiGatewayConfigPb           `json:"ai_gateway,omitempty"`
	BudgetPolicyId       string                       `json:"budget_policy_id,omitempty"`
	Config               *EndpointCoreConfigSummaryPb `json:"config,omitempty"`
	CreationTimestamp    int64                        `json:"creation_timestamp,omitempty"`
	Creator              string                       `json:"creator,omitempty"`
	Description          string                       `json:"description,omitempty"`
	Id                   string                       `json:"id,omitempty"`
	LastUpdatedTimestamp int64                        `json:"last_updated_timestamp,omitempty"`
	Name                 string                       `json:"name,omitempty"`
	State                *EndpointStatePb             `json:"state,omitempty"`
	Tags                 []EndpointTagPb              `json:"tags,omitempty"`
	Task                 string                       `json:"task,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServingEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServingEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointAccessControlRequestPb struct {
	GroupName            string                           `json:"group_name,omitempty"`
	PermissionLevel      ServingEndpointPermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                           `json:"service_principal_name,omitempty"`
	UserName             string                           `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServingEndpointAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServingEndpointAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointAccessControlResponsePb struct {
	AllPermissions       []ServingEndpointPermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string                        `json:"display_name,omitempty"`
	GroupName            string                        `json:"group_name,omitempty"`
	ServicePrincipalName string                        `json:"service_principal_name,omitempty"`
	UserName             string                        `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServingEndpointAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServingEndpointAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointDetailedPb struct {
	AiGateway            *AiGatewayConfigPb                       `json:"ai_gateway,omitempty"`
	BudgetPolicyId       string                                   `json:"budget_policy_id,omitempty"`
	Config               *EndpointCoreConfigOutputPb              `json:"config,omitempty"`
	CreationTimestamp    int64                                    `json:"creation_timestamp,omitempty"`
	Creator              string                                   `json:"creator,omitempty"`
	DataPlaneInfo        *ModelDataPlaneInfoPb                    `json:"data_plane_info,omitempty"`
	Description          string                                   `json:"description,omitempty"`
	EndpointUrl          string                                   `json:"endpoint_url,omitempty"`
	Id                   string                                   `json:"id,omitempty"`
	LastUpdatedTimestamp int64                                    `json:"last_updated_timestamp,omitempty"`
	Name                 string                                   `json:"name,omitempty"`
	PendingConfig        *EndpointPendingConfigPb                 `json:"pending_config,omitempty"`
	PermissionLevel      ServingEndpointDetailedPermissionLevelPb `json:"permission_level,omitempty"`
	RouteOptimized       bool                                     `json:"route_optimized,omitempty"`
	State                *EndpointStatePb                         `json:"state,omitempty"`
	Tags                 []EndpointTagPb                          `json:"tags,omitempty"`
	Task                 string                                   `json:"task,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServingEndpointDetailedPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServingEndpointDetailedPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointDetailedPermissionLevelPb string

const ServingEndpointDetailedPermissionLevelCanManage ServingEndpointDetailedPermissionLevelPb = `CAN_MANAGE`

const ServingEndpointDetailedPermissionLevelCanQuery ServingEndpointDetailedPermissionLevelPb = `CAN_QUERY`

const ServingEndpointDetailedPermissionLevelCanView ServingEndpointDetailedPermissionLevelPb = `CAN_VIEW`

type ServingEndpointPermissionPb struct {
	Inherited           bool                             `json:"inherited,omitempty"`
	InheritedFromObject []string                         `json:"inherited_from_object,omitempty"`
	PermissionLevel     ServingEndpointPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServingEndpointPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServingEndpointPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointPermissionLevelPb string

const ServingEndpointPermissionLevelCanManage ServingEndpointPermissionLevelPb = `CAN_MANAGE`

const ServingEndpointPermissionLevelCanQuery ServingEndpointPermissionLevelPb = `CAN_QUERY`

const ServingEndpointPermissionLevelCanView ServingEndpointPermissionLevelPb = `CAN_VIEW`

type ServingEndpointPermissionsPb struct {
	AccessControlList []ServingEndpointAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                                   `json:"object_id,omitempty"`
	ObjectType        string                                   `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServingEndpointPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServingEndpointPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointPermissionsDescriptionPb struct {
	Description     string                           `json:"description,omitempty"`
	PermissionLevel ServingEndpointPermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ServingEndpointPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ServingEndpointPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointPermissionsRequestPb struct {
	AccessControlList []ServingEndpointAccessControlRequestPb `json:"access_control_list,omitempty"`
	ServingEndpointId string                                  `json:"-" url:"-"`
}

type ServingModelWorkloadTypePb string

const ServingModelWorkloadTypeCpu ServingModelWorkloadTypePb = `CPU`

const ServingModelWorkloadTypeGpuLarge ServingModelWorkloadTypePb = `GPU_LARGE`

const ServingModelWorkloadTypeGpuMedium ServingModelWorkloadTypePb = `GPU_MEDIUM`

const ServingModelWorkloadTypeGpuSmall ServingModelWorkloadTypePb = `GPU_SMALL`

const ServingModelWorkloadTypeMultigpuMedium ServingModelWorkloadTypePb = `MULTIGPU_MEDIUM`

type TrafficConfigPb struct {
	Routes []RoutePb `json:"routes,omitempty"`
}

type UpdateProvisionedThroughputEndpointConfigRequestPb struct {
	Config PtEndpointCoreConfigPb `json:"config"`
	Name   string                 `json:"-" url:"-"`
}

type V1ResponseChoiceElementPb struct {
	FinishReason string         `json:"finishReason,omitempty"`
	Index        int            `json:"index,omitempty"`
	Logprobs     int            `json:"logprobs,omitempty"`
	Message      *ChatMessagePb `json:"message,omitempty"`
	Text         string         `json:"text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *V1ResponseChoiceElementPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st V1ResponseChoiceElementPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
