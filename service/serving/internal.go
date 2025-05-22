// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"io"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func ai21LabsConfigToPb(st *Ai21LabsConfig) (*ai21LabsConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ai21LabsConfigPb{}
	pb.Ai21labsApiKey = st.Ai21labsApiKey

	pb.Ai21labsApiKeyPlaintext = st.Ai21labsApiKeyPlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type ai21LabsConfigPb struct {
	Ai21labsApiKey string `json:"ai21labs_api_key,omitempty"`

	Ai21labsApiKeyPlaintext string `json:"ai21labs_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ai21LabsConfigFromPb(pb *ai21LabsConfigPb) (*Ai21LabsConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Ai21LabsConfig{}
	st.Ai21labsApiKey = pb.Ai21labsApiKey
	st.Ai21labsApiKeyPlaintext = pb.Ai21labsApiKeyPlaintext

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ai21LabsConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ai21LabsConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func aiGatewayConfigToPb(st *AiGatewayConfig) (*aiGatewayConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayConfigPb{}
	pb.FallbackConfig = st.FallbackConfig

	pb.Guardrails = st.Guardrails

	pb.InferenceTableConfig = st.InferenceTableConfig

	pb.RateLimits = st.RateLimits

	pb.UsageTrackingConfig = st.UsageTrackingConfig

	return pb, nil
}

type aiGatewayConfigPb struct {
	FallbackConfig *FallbackConfig `json:"fallback_config,omitempty"`

	Guardrails *AiGatewayGuardrails `json:"guardrails,omitempty"`

	InferenceTableConfig *AiGatewayInferenceTableConfig `json:"inference_table_config,omitempty"`

	RateLimits []AiGatewayRateLimit `json:"rate_limits,omitempty"`

	UsageTrackingConfig *AiGatewayUsageTrackingConfig `json:"usage_tracking_config,omitempty"`
}

func aiGatewayConfigFromPb(pb *aiGatewayConfigPb) (*AiGatewayConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayConfig{}
	st.FallbackConfig = pb.FallbackConfig
	st.Guardrails = pb.Guardrails
	st.InferenceTableConfig = pb.InferenceTableConfig
	st.RateLimits = pb.RateLimits
	st.UsageTrackingConfig = pb.UsageTrackingConfig

	return st, nil
}

func aiGatewayGuardrailParametersToPb(st *AiGatewayGuardrailParameters) (*aiGatewayGuardrailParametersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayGuardrailParametersPb{}
	pb.InvalidKeywords = st.InvalidKeywords

	pb.Pii = st.Pii

	pb.Safety = st.Safety

	pb.ValidTopics = st.ValidTopics

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type aiGatewayGuardrailParametersPb struct {
	InvalidKeywords []string `json:"invalid_keywords,omitempty"`

	Pii *AiGatewayGuardrailPiiBehavior `json:"pii,omitempty"`

	Safety bool `json:"safety,omitempty"`

	ValidTopics []string `json:"valid_topics,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func aiGatewayGuardrailParametersFromPb(pb *aiGatewayGuardrailParametersPb) (*AiGatewayGuardrailParameters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayGuardrailParameters{}
	st.InvalidKeywords = pb.InvalidKeywords
	st.Pii = pb.Pii
	st.Safety = pb.Safety
	st.ValidTopics = pb.ValidTopics

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *aiGatewayGuardrailParametersPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st aiGatewayGuardrailParametersPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func aiGatewayGuardrailPiiBehaviorToPb(st *AiGatewayGuardrailPiiBehavior) (*aiGatewayGuardrailPiiBehaviorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayGuardrailPiiBehaviorPb{}
	pb.Behavior = st.Behavior

	return pb, nil
}

type aiGatewayGuardrailPiiBehaviorPb struct {
	Behavior AiGatewayGuardrailPiiBehaviorBehavior `json:"behavior,omitempty"`
}

func aiGatewayGuardrailPiiBehaviorFromPb(pb *aiGatewayGuardrailPiiBehaviorPb) (*AiGatewayGuardrailPiiBehavior, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayGuardrailPiiBehavior{}
	st.Behavior = pb.Behavior

	return st, nil
}

func aiGatewayGuardrailsToPb(st *AiGatewayGuardrails) (*aiGatewayGuardrailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayGuardrailsPb{}
	pb.Input = st.Input

	pb.Output = st.Output

	return pb, nil
}

type aiGatewayGuardrailsPb struct {
	Input *AiGatewayGuardrailParameters `json:"input,omitempty"`

	Output *AiGatewayGuardrailParameters `json:"output,omitempty"`
}

func aiGatewayGuardrailsFromPb(pb *aiGatewayGuardrailsPb) (*AiGatewayGuardrails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayGuardrails{}
	st.Input = pb.Input
	st.Output = pb.Output

	return st, nil
}

func aiGatewayInferenceTableConfigToPb(st *AiGatewayInferenceTableConfig) (*aiGatewayInferenceTableConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayInferenceTableConfigPb{}
	pb.CatalogName = st.CatalogName

	pb.Enabled = st.Enabled

	pb.SchemaName = st.SchemaName

	pb.TableNamePrefix = st.TableNamePrefix

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type aiGatewayInferenceTableConfigPb struct {
	CatalogName string `json:"catalog_name,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	SchemaName string `json:"schema_name,omitempty"`

	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func aiGatewayInferenceTableConfigFromPb(pb *aiGatewayInferenceTableConfigPb) (*AiGatewayInferenceTableConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayInferenceTableConfig{}
	st.CatalogName = pb.CatalogName
	st.Enabled = pb.Enabled
	st.SchemaName = pb.SchemaName
	st.TableNamePrefix = pb.TableNamePrefix

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *aiGatewayInferenceTableConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st aiGatewayInferenceTableConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func aiGatewayRateLimitToPb(st *AiGatewayRateLimit) (*aiGatewayRateLimitPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayRateLimitPb{}
	pb.Calls = st.Calls

	pb.Key = st.Key

	pb.RenewalPeriod = st.RenewalPeriod

	return pb, nil
}

type aiGatewayRateLimitPb struct {
	Calls int64 `json:"calls"`

	Key AiGatewayRateLimitKey `json:"key,omitempty"`

	RenewalPeriod AiGatewayRateLimitRenewalPeriod `json:"renewal_period"`
}

func aiGatewayRateLimitFromPb(pb *aiGatewayRateLimitPb) (*AiGatewayRateLimit, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayRateLimit{}
	st.Calls = pb.Calls
	st.Key = pb.Key
	st.RenewalPeriod = pb.RenewalPeriod

	return st, nil
}

func aiGatewayUsageTrackingConfigToPb(st *AiGatewayUsageTrackingConfig) (*aiGatewayUsageTrackingConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayUsageTrackingConfigPb{}
	pb.Enabled = st.Enabled

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type aiGatewayUsageTrackingConfigPb struct {
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func aiGatewayUsageTrackingConfigFromPb(pb *aiGatewayUsageTrackingConfigPb) (*AiGatewayUsageTrackingConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayUsageTrackingConfig{}
	st.Enabled = pb.Enabled

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *aiGatewayUsageTrackingConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st aiGatewayUsageTrackingConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func amazonBedrockConfigToPb(st *AmazonBedrockConfig) (*amazonBedrockConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &amazonBedrockConfigPb{}
	pb.AwsAccessKeyId = st.AwsAccessKeyId

	pb.AwsAccessKeyIdPlaintext = st.AwsAccessKeyIdPlaintext

	pb.AwsRegion = st.AwsRegion

	pb.AwsSecretAccessKey = st.AwsSecretAccessKey

	pb.AwsSecretAccessKeyPlaintext = st.AwsSecretAccessKeyPlaintext

	pb.BedrockProvider = st.BedrockProvider

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type amazonBedrockConfigPb struct {
	AwsAccessKeyId string `json:"aws_access_key_id,omitempty"`

	AwsAccessKeyIdPlaintext string `json:"aws_access_key_id_plaintext,omitempty"`

	AwsRegion string `json:"aws_region"`

	AwsSecretAccessKey string `json:"aws_secret_access_key,omitempty"`

	AwsSecretAccessKeyPlaintext string `json:"aws_secret_access_key_plaintext,omitempty"`

	BedrockProvider AmazonBedrockConfigBedrockProvider `json:"bedrock_provider"`

	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func amazonBedrockConfigFromPb(pb *amazonBedrockConfigPb) (*AmazonBedrockConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AmazonBedrockConfig{}
	st.AwsAccessKeyId = pb.AwsAccessKeyId
	st.AwsAccessKeyIdPlaintext = pb.AwsAccessKeyIdPlaintext
	st.AwsRegion = pb.AwsRegion
	st.AwsSecretAccessKey = pb.AwsSecretAccessKey
	st.AwsSecretAccessKeyPlaintext = pb.AwsSecretAccessKeyPlaintext
	st.BedrockProvider = pb.BedrockProvider
	st.InstanceProfileArn = pb.InstanceProfileArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *amazonBedrockConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st amazonBedrockConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func anthropicConfigToPb(st *AnthropicConfig) (*anthropicConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &anthropicConfigPb{}
	pb.AnthropicApiKey = st.AnthropicApiKey

	pb.AnthropicApiKeyPlaintext = st.AnthropicApiKeyPlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type anthropicConfigPb struct {
	AnthropicApiKey string `json:"anthropic_api_key,omitempty"`

	AnthropicApiKeyPlaintext string `json:"anthropic_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func anthropicConfigFromPb(pb *anthropicConfigPb) (*AnthropicConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AnthropicConfig{}
	st.AnthropicApiKey = pb.AnthropicApiKey
	st.AnthropicApiKeyPlaintext = pb.AnthropicApiKeyPlaintext

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *anthropicConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st anthropicConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func apiKeyAuthToPb(st *ApiKeyAuth) (*apiKeyAuthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &apiKeyAuthPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ValuePlaintext = st.ValuePlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type apiKeyAuthPb struct {
	Key string `json:"key"`

	Value string `json:"value,omitempty"`

	ValuePlaintext string `json:"value_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func apiKeyAuthFromPb(pb *apiKeyAuthPb) (*ApiKeyAuth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApiKeyAuth{}
	st.Key = pb.Key
	st.Value = pb.Value
	st.ValuePlaintext = pb.ValuePlaintext

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *apiKeyAuthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st apiKeyAuthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func autoCaptureConfigInputToPb(st *AutoCaptureConfigInput) (*autoCaptureConfigInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &autoCaptureConfigInputPb{}
	pb.CatalogName = st.CatalogName

	pb.Enabled = st.Enabled

	pb.SchemaName = st.SchemaName

	pb.TableNamePrefix = st.TableNamePrefix

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type autoCaptureConfigInputPb struct {
	CatalogName string `json:"catalog_name,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	SchemaName string `json:"schema_name,omitempty"`

	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func autoCaptureConfigInputFromPb(pb *autoCaptureConfigInputPb) (*AutoCaptureConfigInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoCaptureConfigInput{}
	st.CatalogName = pb.CatalogName
	st.Enabled = pb.Enabled
	st.SchemaName = pb.SchemaName
	st.TableNamePrefix = pb.TableNamePrefix

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *autoCaptureConfigInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st autoCaptureConfigInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func autoCaptureConfigOutputToPb(st *AutoCaptureConfigOutput) (*autoCaptureConfigOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &autoCaptureConfigOutputPb{}
	pb.CatalogName = st.CatalogName

	pb.Enabled = st.Enabled

	pb.SchemaName = st.SchemaName

	pb.State = st.State

	pb.TableNamePrefix = st.TableNamePrefix

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type autoCaptureConfigOutputPb struct {
	CatalogName string `json:"catalog_name,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	SchemaName string `json:"schema_name,omitempty"`

	State *AutoCaptureState `json:"state,omitempty"`

	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func autoCaptureConfigOutputFromPb(pb *autoCaptureConfigOutputPb) (*AutoCaptureConfigOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoCaptureConfigOutput{}
	st.CatalogName = pb.CatalogName
	st.Enabled = pb.Enabled
	st.SchemaName = pb.SchemaName
	st.State = pb.State
	st.TableNamePrefix = pb.TableNamePrefix

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *autoCaptureConfigOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st autoCaptureConfigOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func autoCaptureStateToPb(st *AutoCaptureState) (*autoCaptureStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &autoCaptureStatePb{}
	pb.PayloadTable = st.PayloadTable

	return pb, nil
}

type autoCaptureStatePb struct {
	PayloadTable *PayloadTable `json:"payload_table,omitempty"`
}

func autoCaptureStateFromPb(pb *autoCaptureStatePb) (*AutoCaptureState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoCaptureState{}
	st.PayloadTable = pb.PayloadTable

	return st, nil
}

func bearerTokenAuthToPb(st *BearerTokenAuth) (*bearerTokenAuthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &bearerTokenAuthPb{}
	pb.Token = st.Token

	pb.TokenPlaintext = st.TokenPlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type bearerTokenAuthPb struct {
	Token string `json:"token,omitempty"`

	TokenPlaintext string `json:"token_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func bearerTokenAuthFromPb(pb *bearerTokenAuthPb) (*BearerTokenAuth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BearerTokenAuth{}
	st.Token = pb.Token
	st.TokenPlaintext = pb.TokenPlaintext

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *bearerTokenAuthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st bearerTokenAuthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func buildLogsRequestToPb(st *BuildLogsRequest) (*buildLogsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &buildLogsRequestPb{}
	pb.Name = st.Name

	pb.ServedModelName = st.ServedModelName

	return pb, nil
}

type buildLogsRequestPb struct {
	Name string `json:"-" url:"-"`

	ServedModelName string `json:"-" url:"-"`
}

func buildLogsRequestFromPb(pb *buildLogsRequestPb) (*BuildLogsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BuildLogsRequest{}
	st.Name = pb.Name
	st.ServedModelName = pb.ServedModelName

	return st, nil
}

func buildLogsResponseToPb(st *BuildLogsResponse) (*buildLogsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &buildLogsResponsePb{}
	pb.Logs = st.Logs

	return pb, nil
}

type buildLogsResponsePb struct {
	Logs string `json:"logs"`
}

func buildLogsResponseFromPb(pb *buildLogsResponsePb) (*BuildLogsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BuildLogsResponse{}
	st.Logs = pb.Logs

	return st, nil
}

func chatMessageToPb(st *ChatMessage) (*chatMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &chatMessagePb{}
	pb.Content = st.Content

	pb.Role = st.Role

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type chatMessagePb struct {
	Content string `json:"content,omitempty"`

	Role ChatMessageRole `json:"role,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func chatMessageFromPb(pb *chatMessagePb) (*ChatMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ChatMessage{}
	st.Content = pb.Content
	st.Role = pb.Role

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *chatMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st chatMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cohereConfigToPb(st *CohereConfig) (*cohereConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cohereConfigPb{}
	pb.CohereApiBase = st.CohereApiBase

	pb.CohereApiKey = st.CohereApiKey

	pb.CohereApiKeyPlaintext = st.CohereApiKeyPlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cohereConfigPb struct {
	CohereApiBase string `json:"cohere_api_base,omitempty"`

	CohereApiKey string `json:"cohere_api_key,omitempty"`

	CohereApiKeyPlaintext string `json:"cohere_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cohereConfigFromPb(pb *cohereConfigPb) (*CohereConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CohereConfig{}
	st.CohereApiBase = pb.CohereApiBase
	st.CohereApiKey = pb.CohereApiKey
	st.CohereApiKeyPlaintext = pb.CohereApiKeyPlaintext

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cohereConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cohereConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createPtEndpointRequestToPb(st *CreatePtEndpointRequest) (*createPtEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPtEndpointRequestPb{}
	pb.AiGateway = st.AiGateway

	pb.BudgetPolicyId = st.BudgetPolicyId

	pb.Config = st.Config

	pb.Name = st.Name

	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPtEndpointRequestPb struct {
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`

	BudgetPolicyId string `json:"budget_policy_id,omitempty"`

	Config PtEndpointCoreConfig `json:"config"`

	Name string `json:"name"`

	Tags []EndpointTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPtEndpointRequestFromPb(pb *createPtEndpointRequestPb) (*CreatePtEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePtEndpointRequest{}
	st.AiGateway = pb.AiGateway
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Config = pb.Config
	st.Name = pb.Name
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPtEndpointRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPtEndpointRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createServingEndpointToPb(st *CreateServingEndpoint) (*createServingEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createServingEndpointPb{}
	pb.AiGateway = st.AiGateway

	pb.BudgetPolicyId = st.BudgetPolicyId

	pb.Config = st.Config

	pb.Name = st.Name

	pb.RateLimits = st.RateLimits

	pb.RouteOptimized = st.RouteOptimized

	pb.Tags = st.Tags

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createServingEndpointPb struct {
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`

	BudgetPolicyId string `json:"budget_policy_id,omitempty"`

	Config *EndpointCoreConfigInput `json:"config,omitempty"`

	Name string `json:"name"`

	RateLimits []RateLimit `json:"rate_limits,omitempty"`

	RouteOptimized bool `json:"route_optimized,omitempty"`

	Tags []EndpointTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createServingEndpointFromPb(pb *createServingEndpointPb) (*CreateServingEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServingEndpoint{}
	st.AiGateway = pb.AiGateway
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Config = pb.Config
	st.Name = pb.Name
	st.RateLimits = pb.RateLimits
	st.RouteOptimized = pb.RouteOptimized
	st.Tags = pb.Tags

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createServingEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createServingEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func customProviderConfigToPb(st *CustomProviderConfig) (*customProviderConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customProviderConfigPb{}
	pb.ApiKeyAuth = st.ApiKeyAuth

	pb.BearerTokenAuth = st.BearerTokenAuth

	pb.CustomProviderUrl = st.CustomProviderUrl

	return pb, nil
}

type customProviderConfigPb struct {
	ApiKeyAuth *ApiKeyAuth `json:"api_key_auth,omitempty"`

	BearerTokenAuth *BearerTokenAuth `json:"bearer_token_auth,omitempty"`

	CustomProviderUrl string `json:"custom_provider_url"`
}

func customProviderConfigFromPb(pb *customProviderConfigPb) (*CustomProviderConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomProviderConfig{}
	st.ApiKeyAuth = pb.ApiKeyAuth
	st.BearerTokenAuth = pb.BearerTokenAuth
	st.CustomProviderUrl = pb.CustomProviderUrl

	return st, nil
}

func dataPlaneInfoToPb(st *DataPlaneInfo) (*dataPlaneInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataPlaneInfoPb{}
	pb.AuthorizationDetails = st.AuthorizationDetails

	pb.EndpointUrl = st.EndpointUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dataPlaneInfoPb struct {
	AuthorizationDetails string `json:"authorization_details,omitempty"`

	EndpointUrl string `json:"endpoint_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dataPlaneInfoFromPb(pb *dataPlaneInfoPb) (*DataPlaneInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataPlaneInfo{}
	st.AuthorizationDetails = pb.AuthorizationDetails
	st.EndpointUrl = pb.EndpointUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dataPlaneInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dataPlaneInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func databricksModelServingConfigToPb(st *DatabricksModelServingConfig) (*databricksModelServingConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databricksModelServingConfigPb{}
	pb.DatabricksApiToken = st.DatabricksApiToken

	pb.DatabricksApiTokenPlaintext = st.DatabricksApiTokenPlaintext

	pb.DatabricksWorkspaceUrl = st.DatabricksWorkspaceUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type databricksModelServingConfigPb struct {
	DatabricksApiToken string `json:"databricks_api_token,omitempty"`

	DatabricksApiTokenPlaintext string `json:"databricks_api_token_plaintext,omitempty"`

	DatabricksWorkspaceUrl string `json:"databricks_workspace_url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databricksModelServingConfigFromPb(pb *databricksModelServingConfigPb) (*DatabricksModelServingConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabricksModelServingConfig{}
	st.DatabricksApiToken = pb.DatabricksApiToken
	st.DatabricksApiTokenPlaintext = pb.DatabricksApiTokenPlaintext
	st.DatabricksWorkspaceUrl = pb.DatabricksWorkspaceUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *databricksModelServingConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databricksModelServingConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func dataframeSplitInputToPb(st *DataframeSplitInput) (*dataframeSplitInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataframeSplitInputPb{}
	pb.Columns = st.Columns

	pb.Data = st.Data

	pb.Index = st.Index

	return pb, nil
}

type dataframeSplitInputPb struct {
	Columns []any `json:"columns,omitempty"`

	Data []any `json:"data,omitempty"`

	Index []int `json:"index,omitempty"`
}

func dataframeSplitInputFromPb(pb *dataframeSplitInputPb) (*DataframeSplitInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataframeSplitInput{}
	st.Columns = pb.Columns
	st.Data = pb.Data
	st.Index = pb.Index

	return st, nil
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
}

type deleteResponsePb struct {
}

func deleteResponseFromPb(pb *deleteResponsePb) (*DeleteResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteResponse{}

	return st, nil
}

func deleteServingEndpointRequestToPb(st *DeleteServingEndpointRequest) (*deleteServingEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteServingEndpointRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type deleteServingEndpointRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteServingEndpointRequestFromPb(pb *deleteServingEndpointRequestPb) (*DeleteServingEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServingEndpointRequest{}
	st.Name = pb.Name

	return st, nil
}

func embeddingsV1ResponseEmbeddingElementToPb(st *EmbeddingsV1ResponseEmbeddingElement) (*embeddingsV1ResponseEmbeddingElementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &embeddingsV1ResponseEmbeddingElementPb{}
	pb.Embedding = st.Embedding

	pb.Index = st.Index

	pb.Object = st.Object

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type embeddingsV1ResponseEmbeddingElementPb struct {
	Embedding []float64 `json:"embedding,omitempty"`

	Index int `json:"index,omitempty"`

	Object EmbeddingsV1ResponseEmbeddingElementObject `json:"object,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func embeddingsV1ResponseEmbeddingElementFromPb(pb *embeddingsV1ResponseEmbeddingElementPb) (*EmbeddingsV1ResponseEmbeddingElement, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EmbeddingsV1ResponseEmbeddingElement{}
	st.Embedding = pb.Embedding
	st.Index = pb.Index
	st.Object = pb.Object

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *embeddingsV1ResponseEmbeddingElementPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st embeddingsV1ResponseEmbeddingElementPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func endpointCoreConfigInputToPb(st *EndpointCoreConfigInput) (*endpointCoreConfigInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointCoreConfigInputPb{}
	pb.AutoCaptureConfig = st.AutoCaptureConfig

	pb.Name = st.Name

	pb.ServedEntities = st.ServedEntities

	pb.ServedModels = st.ServedModels

	pb.TrafficConfig = st.TrafficConfig

	return pb, nil
}

type endpointCoreConfigInputPb struct {
	AutoCaptureConfig *AutoCaptureConfigInput `json:"auto_capture_config,omitempty"`

	Name string `json:"-" url:"-"`

	ServedEntities []ServedEntityInput `json:"served_entities,omitempty"`

	ServedModels []ServedModelInput `json:"served_models,omitempty"`

	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`
}

func endpointCoreConfigInputFromPb(pb *endpointCoreConfigInputPb) (*EndpointCoreConfigInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointCoreConfigInput{}
	st.AutoCaptureConfig = pb.AutoCaptureConfig
	st.Name = pb.Name
	st.ServedEntities = pb.ServedEntities
	st.ServedModels = pb.ServedModels
	st.TrafficConfig = pb.TrafficConfig

	return st, nil
}

func endpointCoreConfigOutputToPb(st *EndpointCoreConfigOutput) (*endpointCoreConfigOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointCoreConfigOutputPb{}
	pb.AutoCaptureConfig = st.AutoCaptureConfig

	pb.ConfigVersion = st.ConfigVersion

	pb.ServedEntities = st.ServedEntities

	pb.ServedModels = st.ServedModels

	pb.TrafficConfig = st.TrafficConfig

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type endpointCoreConfigOutputPb struct {
	AutoCaptureConfig *AutoCaptureConfigOutput `json:"auto_capture_config,omitempty"`

	ConfigVersion int64 `json:"config_version,omitempty"`

	ServedEntities []ServedEntityOutput `json:"served_entities,omitempty"`

	ServedModels []ServedModelOutput `json:"served_models,omitempty"`

	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointCoreConfigOutputFromPb(pb *endpointCoreConfigOutputPb) (*EndpointCoreConfigOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointCoreConfigOutput{}
	st.AutoCaptureConfig = pb.AutoCaptureConfig
	st.ConfigVersion = pb.ConfigVersion
	st.ServedEntities = pb.ServedEntities
	st.ServedModels = pb.ServedModels
	st.TrafficConfig = pb.TrafficConfig

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointCoreConfigOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointCoreConfigOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func endpointCoreConfigSummaryToPb(st *EndpointCoreConfigSummary) (*endpointCoreConfigSummaryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointCoreConfigSummaryPb{}
	pb.ServedEntities = st.ServedEntities

	pb.ServedModels = st.ServedModels

	return pb, nil
}

type endpointCoreConfigSummaryPb struct {
	ServedEntities []ServedEntitySpec `json:"served_entities,omitempty"`

	ServedModels []ServedModelSpec `json:"served_models,omitempty"`
}

func endpointCoreConfigSummaryFromPb(pb *endpointCoreConfigSummaryPb) (*EndpointCoreConfigSummary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointCoreConfigSummary{}
	st.ServedEntities = pb.ServedEntities
	st.ServedModels = pb.ServedModels

	return st, nil
}

func endpointPendingConfigToPb(st *EndpointPendingConfig) (*endpointPendingConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointPendingConfigPb{}
	pb.AutoCaptureConfig = st.AutoCaptureConfig

	pb.ConfigVersion = st.ConfigVersion

	pb.ServedEntities = st.ServedEntities

	pb.ServedModels = st.ServedModels

	pb.StartTime = st.StartTime

	pb.TrafficConfig = st.TrafficConfig

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type endpointPendingConfigPb struct {
	AutoCaptureConfig *AutoCaptureConfigOutput `json:"auto_capture_config,omitempty"`

	ConfigVersion int `json:"config_version,omitempty"`

	ServedEntities []ServedEntityOutput `json:"served_entities,omitempty"`

	ServedModels []ServedModelOutput `json:"served_models,omitempty"`

	StartTime int64 `json:"start_time,omitempty"`

	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointPendingConfigFromPb(pb *endpointPendingConfigPb) (*EndpointPendingConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointPendingConfig{}
	st.AutoCaptureConfig = pb.AutoCaptureConfig
	st.ConfigVersion = pb.ConfigVersion
	st.ServedEntities = pb.ServedEntities
	st.ServedModels = pb.ServedModels
	st.StartTime = pb.StartTime
	st.TrafficConfig = pb.TrafficConfig

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointPendingConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointPendingConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func endpointStateToPb(st *EndpointState) (*endpointStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointStatePb{}
	pb.ConfigUpdate = st.ConfigUpdate

	pb.Ready = st.Ready

	return pb, nil
}

type endpointStatePb struct {
	ConfigUpdate EndpointStateConfigUpdate `json:"config_update,omitempty"`

	Ready EndpointStateReady `json:"ready,omitempty"`
}

func endpointStateFromPb(pb *endpointStatePb) (*EndpointState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointState{}
	st.ConfigUpdate = pb.ConfigUpdate
	st.Ready = pb.Ready

	return st, nil
}

func endpointTagToPb(st *EndpointTag) (*endpointTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointTagPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type endpointTagPb struct {
	Key string `json:"key"`

	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointTagFromPb(pb *endpointTagPb) (*EndpointTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func endpointTagsToPb(st *EndpointTags) (*endpointTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointTagsPb{}
	pb.Tags = st.Tags

	return pb, nil
}

type endpointTagsPb struct {
	Tags []EndpointTag `json:"tags,omitempty"`
}

func endpointTagsFromPb(pb *endpointTagsPb) (*EndpointTags, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTags{}
	st.Tags = pb.Tags

	return st, nil
}

func exportMetricsRequestToPb(st *ExportMetricsRequest) (*exportMetricsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportMetricsRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type exportMetricsRequestPb struct {
	Name string `json:"-" url:"-"`
}

func exportMetricsRequestFromPb(pb *exportMetricsRequestPb) (*ExportMetricsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportMetricsRequest{}
	st.Name = pb.Name

	return st, nil
}

func exportMetricsResponseToPb(st *ExportMetricsResponse) (*exportMetricsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportMetricsResponsePb{}
	pb.Contents = st.Contents

	return pb, nil
}

type exportMetricsResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

func exportMetricsResponseFromPb(pb *exportMetricsResponsePb) (*ExportMetricsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportMetricsResponse{}
	st.Contents = pb.Contents

	return st, nil
}

func externalFunctionRequestToPb(st *ExternalFunctionRequest) (*externalFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalFunctionRequestPb{}
	pb.ConnectionName = st.ConnectionName

	pb.Headers = st.Headers

	pb.Json = st.Json

	pb.Method = st.Method

	pb.Params = st.Params

	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type externalFunctionRequestPb struct {
	ConnectionName string `json:"connection_name"`

	Headers string `json:"headers,omitempty"`

	Json string `json:"json,omitempty"`

	Method ExternalFunctionRequestHttpMethod `json:"method"`

	Params string `json:"params,omitempty"`

	Path string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalFunctionRequestFromPb(pb *externalFunctionRequestPb) (*ExternalFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalFunctionRequest{}
	st.ConnectionName = pb.ConnectionName
	st.Headers = pb.Headers
	st.Json = pb.Json
	st.Method = pb.Method
	st.Params = pb.Params
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalFunctionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalFunctionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func externalModelToPb(st *ExternalModel) (*externalModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalModelPb{}
	pb.Ai21labsConfig = st.Ai21labsConfig

	pb.AmazonBedrockConfig = st.AmazonBedrockConfig

	pb.AnthropicConfig = st.AnthropicConfig

	pb.CohereConfig = st.CohereConfig

	pb.CustomProviderConfig = st.CustomProviderConfig

	pb.DatabricksModelServingConfig = st.DatabricksModelServingConfig

	pb.GoogleCloudVertexAiConfig = st.GoogleCloudVertexAiConfig

	pb.Name = st.Name

	pb.OpenaiConfig = st.OpenaiConfig

	pb.PalmConfig = st.PalmConfig

	pb.Provider = st.Provider

	pb.Task = st.Task

	return pb, nil
}

type externalModelPb struct {
	Ai21labsConfig *Ai21LabsConfig `json:"ai21labs_config,omitempty"`

	AmazonBedrockConfig *AmazonBedrockConfig `json:"amazon_bedrock_config,omitempty"`

	AnthropicConfig *AnthropicConfig `json:"anthropic_config,omitempty"`

	CohereConfig *CohereConfig `json:"cohere_config,omitempty"`

	CustomProviderConfig *CustomProviderConfig `json:"custom_provider_config,omitempty"`

	DatabricksModelServingConfig *DatabricksModelServingConfig `json:"databricks_model_serving_config,omitempty"`

	GoogleCloudVertexAiConfig *GoogleCloudVertexAiConfig `json:"google_cloud_vertex_ai_config,omitempty"`

	Name string `json:"name"`

	OpenaiConfig *OpenAiConfig `json:"openai_config,omitempty"`

	PalmConfig *PaLmConfig `json:"palm_config,omitempty"`

	Provider ExternalModelProvider `json:"provider"`

	Task string `json:"task"`
}

func externalModelFromPb(pb *externalModelPb) (*ExternalModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalModel{}
	st.Ai21labsConfig = pb.Ai21labsConfig
	st.AmazonBedrockConfig = pb.AmazonBedrockConfig
	st.AnthropicConfig = pb.AnthropicConfig
	st.CohereConfig = pb.CohereConfig
	st.CustomProviderConfig = pb.CustomProviderConfig
	st.DatabricksModelServingConfig = pb.DatabricksModelServingConfig
	st.GoogleCloudVertexAiConfig = pb.GoogleCloudVertexAiConfig
	st.Name = pb.Name
	st.OpenaiConfig = pb.OpenaiConfig
	st.PalmConfig = pb.PalmConfig
	st.Provider = pb.Provider
	st.Task = pb.Task

	return st, nil
}

func externalModelUsageElementToPb(st *ExternalModelUsageElement) (*externalModelUsageElementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalModelUsageElementPb{}
	pb.CompletionTokens = st.CompletionTokens

	pb.PromptTokens = st.PromptTokens

	pb.TotalTokens = st.TotalTokens

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type externalModelUsageElementPb struct {
	CompletionTokens int `json:"completion_tokens,omitempty"`

	PromptTokens int `json:"prompt_tokens,omitempty"`

	TotalTokens int `json:"total_tokens,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalModelUsageElementFromPb(pb *externalModelUsageElementPb) (*ExternalModelUsageElement, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalModelUsageElement{}
	st.CompletionTokens = pb.CompletionTokens
	st.PromptTokens = pb.PromptTokens
	st.TotalTokens = pb.TotalTokens

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalModelUsageElementPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalModelUsageElementPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func fallbackConfigToPb(st *FallbackConfig) (*fallbackConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fallbackConfigPb{}
	pb.Enabled = st.Enabled

	return pb, nil
}

type fallbackConfigPb struct {
	Enabled bool `json:"enabled"`
}

func fallbackConfigFromPb(pb *fallbackConfigPb) (*FallbackConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FallbackConfig{}
	st.Enabled = pb.Enabled

	return st, nil
}

func foundationModelToPb(st *FoundationModel) (*foundationModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &foundationModelPb{}
	pb.Description = st.Description

	pb.DisplayName = st.DisplayName

	pb.Docs = st.Docs

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type foundationModelPb struct {
	Description string `json:"description,omitempty"`

	DisplayName string `json:"display_name,omitempty"`

	Docs string `json:"docs,omitempty"`

	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func foundationModelFromPb(pb *foundationModelPb) (*FoundationModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FoundationModel{}
	st.Description = pb.Description
	st.DisplayName = pb.DisplayName
	st.Docs = pb.Docs
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *foundationModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st foundationModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getOpenApiRequestToPb(st *GetOpenApiRequest) (*getOpenApiRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getOpenApiRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getOpenApiRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getOpenApiRequestFromPb(pb *getOpenApiRequestPb) (*GetOpenApiRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOpenApiRequest{}
	st.Name = pb.Name

	return st, nil
}

func getOpenApiResponseToPb(st *GetOpenApiResponse) (*getOpenApiResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getOpenApiResponsePb{}
	pb.Contents = st.Contents

	return pb, nil
}

type getOpenApiResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

func getOpenApiResponseFromPb(pb *getOpenApiResponsePb) (*GetOpenApiResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOpenApiResponse{}
	st.Contents = pb.Contents

	return st, nil
}

func getServingEndpointPermissionLevelsRequestToPb(st *GetServingEndpointPermissionLevelsRequest) (*getServingEndpointPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServingEndpointPermissionLevelsRequestPb{}
	pb.ServingEndpointId = st.ServingEndpointId

	return pb, nil
}

type getServingEndpointPermissionLevelsRequestPb struct {
	ServingEndpointId string `json:"-" url:"-"`
}

func getServingEndpointPermissionLevelsRequestFromPb(pb *getServingEndpointPermissionLevelsRequestPb) (*GetServingEndpointPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointPermissionLevelsRequest{}
	st.ServingEndpointId = pb.ServingEndpointId

	return st, nil
}

func getServingEndpointPermissionLevelsResponseToPb(st *GetServingEndpointPermissionLevelsResponse) (*getServingEndpointPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServingEndpointPermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getServingEndpointPermissionLevelsResponsePb struct {
	PermissionLevels []ServingEndpointPermissionsDescription `json:"permission_levels,omitempty"`
}

func getServingEndpointPermissionLevelsResponseFromPb(pb *getServingEndpointPermissionLevelsResponsePb) (*GetServingEndpointPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointPermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getServingEndpointPermissionsRequestToPb(st *GetServingEndpointPermissionsRequest) (*getServingEndpointPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServingEndpointPermissionsRequestPb{}
	pb.ServingEndpointId = st.ServingEndpointId

	return pb, nil
}

type getServingEndpointPermissionsRequestPb struct {
	ServingEndpointId string `json:"-" url:"-"`
}

func getServingEndpointPermissionsRequestFromPb(pb *getServingEndpointPermissionsRequestPb) (*GetServingEndpointPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointPermissionsRequest{}
	st.ServingEndpointId = pb.ServingEndpointId

	return st, nil
}

func getServingEndpointRequestToPb(st *GetServingEndpointRequest) (*getServingEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServingEndpointRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

type getServingEndpointRequestPb struct {
	Name string `json:"-" url:"-"`
}

func getServingEndpointRequestFromPb(pb *getServingEndpointRequestPb) (*GetServingEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointRequest{}
	st.Name = pb.Name

	return st, nil
}

func googleCloudVertexAiConfigToPb(st *GoogleCloudVertexAiConfig) (*googleCloudVertexAiConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &googleCloudVertexAiConfigPb{}
	pb.PrivateKey = st.PrivateKey

	pb.PrivateKeyPlaintext = st.PrivateKeyPlaintext

	pb.ProjectId = st.ProjectId

	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type googleCloudVertexAiConfigPb struct {
	PrivateKey string `json:"private_key,omitempty"`

	PrivateKeyPlaintext string `json:"private_key_plaintext,omitempty"`

	ProjectId string `json:"project_id"`

	Region string `json:"region"`

	ForceSendFields []string `json:"-" url:"-"`
}

func googleCloudVertexAiConfigFromPb(pb *googleCloudVertexAiConfigPb) (*GoogleCloudVertexAiConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GoogleCloudVertexAiConfig{}
	st.PrivateKey = pb.PrivateKey
	st.PrivateKeyPlaintext = pb.PrivateKeyPlaintext
	st.ProjectId = pb.ProjectId
	st.Region = pb.Region

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *googleCloudVertexAiConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st googleCloudVertexAiConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func httpRequestResponseToPb(st *HttpRequestResponse) (*httpRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpRequestResponsePb{}
	pb.Contents = st.Contents

	return pb, nil
}

type httpRequestResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

func httpRequestResponseFromPb(pb *httpRequestResponsePb) (*HttpRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &HttpRequestResponse{}
	st.Contents = pb.Contents

	return st, nil
}

func listEndpointsResponseToPb(st *ListEndpointsResponse) (*listEndpointsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listEndpointsResponsePb{}
	pb.Endpoints = st.Endpoints

	return pb, nil
}

type listEndpointsResponsePb struct {
	Endpoints []ServingEndpoint `json:"endpoints,omitempty"`
}

func listEndpointsResponseFromPb(pb *listEndpointsResponsePb) (*ListEndpointsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListEndpointsResponse{}
	st.Endpoints = pb.Endpoints

	return st, nil
}

func logsRequestToPb(st *LogsRequest) (*logsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logsRequestPb{}
	pb.Name = st.Name

	pb.ServedModelName = st.ServedModelName

	return pb, nil
}

type logsRequestPb struct {
	Name string `json:"-" url:"-"`

	ServedModelName string `json:"-" url:"-"`
}

func logsRequestFromPb(pb *logsRequestPb) (*LogsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogsRequest{}
	st.Name = pb.Name
	st.ServedModelName = pb.ServedModelName

	return st, nil
}

func modelDataPlaneInfoToPb(st *ModelDataPlaneInfo) (*modelDataPlaneInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelDataPlaneInfoPb{}
	pb.QueryInfo = st.QueryInfo

	return pb, nil
}

type modelDataPlaneInfoPb struct {
	QueryInfo *DataPlaneInfo `json:"query_info,omitempty"`
}

func modelDataPlaneInfoFromPb(pb *modelDataPlaneInfoPb) (*ModelDataPlaneInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelDataPlaneInfo{}
	st.QueryInfo = pb.QueryInfo

	return st, nil
}

func openAiConfigToPb(st *OpenAiConfig) (*openAiConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &openAiConfigPb{}
	pb.MicrosoftEntraClientId = st.MicrosoftEntraClientId

	pb.MicrosoftEntraClientSecret = st.MicrosoftEntraClientSecret

	pb.MicrosoftEntraClientSecretPlaintext = st.MicrosoftEntraClientSecretPlaintext

	pb.MicrosoftEntraTenantId = st.MicrosoftEntraTenantId

	pb.OpenaiApiBase = st.OpenaiApiBase

	pb.OpenaiApiKey = st.OpenaiApiKey

	pb.OpenaiApiKeyPlaintext = st.OpenaiApiKeyPlaintext

	pb.OpenaiApiType = st.OpenaiApiType

	pb.OpenaiApiVersion = st.OpenaiApiVersion

	pb.OpenaiDeploymentName = st.OpenaiDeploymentName

	pb.OpenaiOrganization = st.OpenaiOrganization

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type openAiConfigPb struct {
	MicrosoftEntraClientId string `json:"microsoft_entra_client_id,omitempty"`

	MicrosoftEntraClientSecret string `json:"microsoft_entra_client_secret,omitempty"`

	MicrosoftEntraClientSecretPlaintext string `json:"microsoft_entra_client_secret_plaintext,omitempty"`

	MicrosoftEntraTenantId string `json:"microsoft_entra_tenant_id,omitempty"`

	OpenaiApiBase string `json:"openai_api_base,omitempty"`

	OpenaiApiKey string `json:"openai_api_key,omitempty"`

	OpenaiApiKeyPlaintext string `json:"openai_api_key_plaintext,omitempty"`

	OpenaiApiType string `json:"openai_api_type,omitempty"`

	OpenaiApiVersion string `json:"openai_api_version,omitempty"`

	OpenaiDeploymentName string `json:"openai_deployment_name,omitempty"`

	OpenaiOrganization string `json:"openai_organization,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func openAiConfigFromPb(pb *openAiConfigPb) (*OpenAiConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OpenAiConfig{}
	st.MicrosoftEntraClientId = pb.MicrosoftEntraClientId
	st.MicrosoftEntraClientSecret = pb.MicrosoftEntraClientSecret
	st.MicrosoftEntraClientSecretPlaintext = pb.MicrosoftEntraClientSecretPlaintext
	st.MicrosoftEntraTenantId = pb.MicrosoftEntraTenantId
	st.OpenaiApiBase = pb.OpenaiApiBase
	st.OpenaiApiKey = pb.OpenaiApiKey
	st.OpenaiApiKeyPlaintext = pb.OpenaiApiKeyPlaintext
	st.OpenaiApiType = pb.OpenaiApiType
	st.OpenaiApiVersion = pb.OpenaiApiVersion
	st.OpenaiDeploymentName = pb.OpenaiDeploymentName
	st.OpenaiOrganization = pb.OpenaiOrganization

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *openAiConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st openAiConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func paLmConfigToPb(st *PaLmConfig) (*paLmConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &paLmConfigPb{}
	pb.PalmApiKey = st.PalmApiKey

	pb.PalmApiKeyPlaintext = st.PalmApiKeyPlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type paLmConfigPb struct {
	PalmApiKey string `json:"palm_api_key,omitempty"`

	PalmApiKeyPlaintext string `json:"palm_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func paLmConfigFromPb(pb *paLmConfigPb) (*PaLmConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PaLmConfig{}
	st.PalmApiKey = pb.PalmApiKey
	st.PalmApiKeyPlaintext = pb.PalmApiKeyPlaintext

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *paLmConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st paLmConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func patchServingEndpointTagsToPb(st *PatchServingEndpointTags) (*patchServingEndpointTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchServingEndpointTagsPb{}
	pb.AddTags = st.AddTags

	pb.DeleteTags = st.DeleteTags

	pb.Name = st.Name

	return pb, nil
}

type patchServingEndpointTagsPb struct {
	AddTags []EndpointTag `json:"add_tags,omitempty"`

	DeleteTags []string `json:"delete_tags,omitempty"`

	Name string `json:"-" url:"-"`
}

func patchServingEndpointTagsFromPb(pb *patchServingEndpointTagsPb) (*PatchServingEndpointTags, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PatchServingEndpointTags{}
	st.AddTags = pb.AddTags
	st.DeleteTags = pb.DeleteTags
	st.Name = pb.Name

	return st, nil
}

func payloadTableToPb(st *PayloadTable) (*payloadTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &payloadTablePb{}
	pb.Name = st.Name

	pb.Status = st.Status

	pb.StatusMessage = st.StatusMessage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type payloadTablePb struct {
	Name string `json:"name,omitempty"`

	Status string `json:"status,omitempty"`

	StatusMessage string `json:"status_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func payloadTableFromPb(pb *payloadTablePb) (*PayloadTable, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PayloadTable{}
	st.Name = pb.Name
	st.Status = pb.Status
	st.StatusMessage = pb.StatusMessage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *payloadTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st payloadTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func ptEndpointCoreConfigToPb(st *PtEndpointCoreConfig) (*ptEndpointCoreConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ptEndpointCoreConfigPb{}
	pb.ServedEntities = st.ServedEntities

	pb.TrafficConfig = st.TrafficConfig

	return pb, nil
}

type ptEndpointCoreConfigPb struct {
	ServedEntities []PtServedModel `json:"served_entities,omitempty"`

	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`
}

func ptEndpointCoreConfigFromPb(pb *ptEndpointCoreConfigPb) (*PtEndpointCoreConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PtEndpointCoreConfig{}
	st.ServedEntities = pb.ServedEntities
	st.TrafficConfig = pb.TrafficConfig

	return st, nil
}

func ptServedModelToPb(st *PtServedModel) (*ptServedModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ptServedModelPb{}
	pb.EntityName = st.EntityName

	pb.EntityVersion = st.EntityVersion

	pb.Name = st.Name

	pb.ProvisionedModelUnits = st.ProvisionedModelUnits

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type ptServedModelPb struct {
	EntityName string `json:"entity_name"`

	EntityVersion string `json:"entity_version,omitempty"`

	Name string `json:"name,omitempty"`

	ProvisionedModelUnits int64 `json:"provisioned_model_units"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ptServedModelFromPb(pb *ptServedModelPb) (*PtServedModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PtServedModel{}
	st.EntityName = pb.EntityName
	st.EntityVersion = pb.EntityVersion
	st.Name = pb.Name
	st.ProvisionedModelUnits = pb.ProvisionedModelUnits

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ptServedModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ptServedModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func putAiGatewayRequestToPb(st *PutAiGatewayRequest) (*putAiGatewayRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putAiGatewayRequestPb{}
	pb.FallbackConfig = st.FallbackConfig

	pb.Guardrails = st.Guardrails

	pb.InferenceTableConfig = st.InferenceTableConfig

	pb.Name = st.Name

	pb.RateLimits = st.RateLimits

	pb.UsageTrackingConfig = st.UsageTrackingConfig

	return pb, nil
}

type putAiGatewayRequestPb struct {
	FallbackConfig *FallbackConfig `json:"fallback_config,omitempty"`

	Guardrails *AiGatewayGuardrails `json:"guardrails,omitempty"`

	InferenceTableConfig *AiGatewayInferenceTableConfig `json:"inference_table_config,omitempty"`

	Name string `json:"-" url:"-"`

	RateLimits []AiGatewayRateLimit `json:"rate_limits,omitempty"`

	UsageTrackingConfig *AiGatewayUsageTrackingConfig `json:"usage_tracking_config,omitempty"`
}

func putAiGatewayRequestFromPb(pb *putAiGatewayRequestPb) (*PutAiGatewayRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutAiGatewayRequest{}
	st.FallbackConfig = pb.FallbackConfig
	st.Guardrails = pb.Guardrails
	st.InferenceTableConfig = pb.InferenceTableConfig
	st.Name = pb.Name
	st.RateLimits = pb.RateLimits
	st.UsageTrackingConfig = pb.UsageTrackingConfig

	return st, nil
}

func putAiGatewayResponseToPb(st *PutAiGatewayResponse) (*putAiGatewayResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putAiGatewayResponsePb{}
	pb.FallbackConfig = st.FallbackConfig

	pb.Guardrails = st.Guardrails

	pb.InferenceTableConfig = st.InferenceTableConfig

	pb.RateLimits = st.RateLimits

	pb.UsageTrackingConfig = st.UsageTrackingConfig

	return pb, nil
}

type putAiGatewayResponsePb struct {
	FallbackConfig *FallbackConfig `json:"fallback_config,omitempty"`

	Guardrails *AiGatewayGuardrails `json:"guardrails,omitempty"`

	InferenceTableConfig *AiGatewayInferenceTableConfig `json:"inference_table_config,omitempty"`

	RateLimits []AiGatewayRateLimit `json:"rate_limits,omitempty"`

	UsageTrackingConfig *AiGatewayUsageTrackingConfig `json:"usage_tracking_config,omitempty"`
}

func putAiGatewayResponseFromPb(pb *putAiGatewayResponsePb) (*PutAiGatewayResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutAiGatewayResponse{}
	st.FallbackConfig = pb.FallbackConfig
	st.Guardrails = pb.Guardrails
	st.InferenceTableConfig = pb.InferenceTableConfig
	st.RateLimits = pb.RateLimits
	st.UsageTrackingConfig = pb.UsageTrackingConfig

	return st, nil
}

func putRequestToPb(st *PutRequest) (*putRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putRequestPb{}
	pb.Name = st.Name

	pb.RateLimits = st.RateLimits

	return pb, nil
}

type putRequestPb struct {
	Name string `json:"-" url:"-"`

	RateLimits []RateLimit `json:"rate_limits,omitempty"`
}

func putRequestFromPb(pb *putRequestPb) (*PutRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutRequest{}
	st.Name = pb.Name
	st.RateLimits = pb.RateLimits

	return st, nil
}

func putResponseToPb(st *PutResponse) (*putResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putResponsePb{}
	pb.RateLimits = st.RateLimits

	return pb, nil
}

type putResponsePb struct {
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
}

func putResponseFromPb(pb *putResponsePb) (*PutResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutResponse{}
	st.RateLimits = pb.RateLimits

	return st, nil
}

func queryEndpointInputToPb(st *QueryEndpointInput) (*queryEndpointInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryEndpointInputPb{}
	pb.DataframeRecords = st.DataframeRecords

	pb.DataframeSplit = st.DataframeSplit

	pb.ExtraParams = st.ExtraParams

	pb.Input = st.Input

	pb.Inputs = st.Inputs

	pb.Instances = st.Instances

	pb.MaxTokens = st.MaxTokens

	pb.Messages = st.Messages

	pb.N = st.N

	pb.Name = st.Name

	pb.Prompt = st.Prompt

	pb.Stop = st.Stop

	pb.Stream = st.Stream

	pb.Temperature = st.Temperature

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type queryEndpointInputPb struct {
	DataframeRecords []any `json:"dataframe_records,omitempty"`

	DataframeSplit *DataframeSplitInput `json:"dataframe_split,omitempty"`

	ExtraParams map[string]string `json:"extra_params,omitempty"`

	Input any `json:"input,omitempty"`

	Inputs any `json:"inputs,omitempty"`

	Instances []any `json:"instances,omitempty"`

	MaxTokens int `json:"max_tokens,omitempty"`

	Messages []ChatMessage `json:"messages,omitempty"`

	N int `json:"n,omitempty"`

	Name string `json:"-" url:"-"`

	Prompt any `json:"prompt,omitempty"`

	Stop []string `json:"stop,omitempty"`

	Stream bool `json:"stream,omitempty"`

	Temperature float64 `json:"temperature,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryEndpointInputFromPb(pb *queryEndpointInputPb) (*QueryEndpointInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryEndpointInput{}
	st.DataframeRecords = pb.DataframeRecords
	st.DataframeSplit = pb.DataframeSplit
	st.ExtraParams = pb.ExtraParams
	st.Input = pb.Input
	st.Inputs = pb.Inputs
	st.Instances = pb.Instances
	st.MaxTokens = pb.MaxTokens
	st.Messages = pb.Messages
	st.N = pb.N
	st.Name = pb.Name
	st.Prompt = pb.Prompt
	st.Stop = pb.Stop
	st.Stream = pb.Stream
	st.Temperature = pb.Temperature

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryEndpointInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryEndpointInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func queryEndpointResponseToPb(st *QueryEndpointResponse) (*queryEndpointResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryEndpointResponsePb{}
	pb.Choices = st.Choices

	pb.Created = st.Created

	pb.Data = st.Data

	pb.Id = st.Id

	pb.Model = st.Model

	pb.Object = st.Object

	pb.Predictions = st.Predictions

	pb.ServedModelName = st.ServedModelName

	pb.Usage = st.Usage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type queryEndpointResponsePb struct {
	Choices []V1ResponseChoiceElement `json:"choices,omitempty"`

	Created int64 `json:"created,omitempty"`

	Data []EmbeddingsV1ResponseEmbeddingElement `json:"data,omitempty"`

	Id string `json:"id,omitempty"`

	Model string `json:"model,omitempty"`

	Object QueryEndpointResponseObject `json:"object,omitempty"`

	Predictions []any `json:"predictions,omitempty"`

	ServedModelName string `json:"-" url:"-" header:"served-model-name,omitempty"`

	Usage *ExternalModelUsageElement `json:"usage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryEndpointResponseFromPb(pb *queryEndpointResponsePb) (*QueryEndpointResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryEndpointResponse{}
	st.Choices = pb.Choices
	st.Created = pb.Created
	st.Data = pb.Data
	st.Id = pb.Id
	st.Model = pb.Model
	st.Object = pb.Object
	st.Predictions = pb.Predictions
	st.ServedModelName = pb.ServedModelName
	st.Usage = pb.Usage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryEndpointResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryEndpointResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func rateLimitToPb(st *RateLimit) (*rateLimitPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rateLimitPb{}
	pb.Calls = st.Calls

	pb.Key = st.Key

	pb.RenewalPeriod = st.RenewalPeriod

	return pb, nil
}

type rateLimitPb struct {
	Calls int64 `json:"calls"`

	Key RateLimitKey `json:"key,omitempty"`

	RenewalPeriod RateLimitRenewalPeriod `json:"renewal_period"`
}

func rateLimitFromPb(pb *rateLimitPb) (*RateLimit, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RateLimit{}
	st.Calls = pb.Calls
	st.Key = pb.Key
	st.RenewalPeriod = pb.RenewalPeriod

	return st, nil
}

func routeToPb(st *Route) (*routePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &routePb{}
	pb.ServedModelName = st.ServedModelName

	pb.TrafficPercentage = st.TrafficPercentage

	return pb, nil
}

type routePb struct {
	ServedModelName string `json:"served_model_name"`

	TrafficPercentage int `json:"traffic_percentage"`
}

func routeFromPb(pb *routePb) (*Route, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Route{}
	st.ServedModelName = pb.ServedModelName
	st.TrafficPercentage = pb.TrafficPercentage

	return st, nil
}

func servedEntityInputToPb(st *ServedEntityInput) (*servedEntityInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedEntityInputPb{}
	pb.EntityName = st.EntityName

	pb.EntityVersion = st.EntityVersion

	pb.EnvironmentVars = st.EnvironmentVars

	pb.ExternalModel = st.ExternalModel

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.MaxProvisionedThroughput = st.MaxProvisionedThroughput

	pb.MinProvisionedThroughput = st.MinProvisionedThroughput

	pb.Name = st.Name

	pb.ProvisionedModelUnits = st.ProvisionedModelUnits

	pb.ScaleToZeroEnabled = st.ScaleToZeroEnabled

	pb.WorkloadSize = st.WorkloadSize

	pb.WorkloadType = st.WorkloadType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servedEntityInputPb struct {
	EntityName string `json:"entity_name,omitempty"`

	EntityVersion string `json:"entity_version,omitempty"`

	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`

	ExternalModel *ExternalModel `json:"external_model,omitempty"`

	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`

	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`

	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`

	Name string `json:"name,omitempty"`

	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`

	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`

	WorkloadSize string `json:"workload_size,omitempty"`

	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedEntityInputFromPb(pb *servedEntityInputPb) (*ServedEntityInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedEntityInput{}
	st.EntityName = pb.EntityName
	st.EntityVersion = pb.EntityVersion
	st.EnvironmentVars = pb.EnvironmentVars
	st.ExternalModel = pb.ExternalModel
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.MaxProvisionedThroughput = pb.MaxProvisionedThroughput
	st.MinProvisionedThroughput = pb.MinProvisionedThroughput
	st.Name = pb.Name
	st.ProvisionedModelUnits = pb.ProvisionedModelUnits
	st.ScaleToZeroEnabled = pb.ScaleToZeroEnabled
	st.WorkloadSize = pb.WorkloadSize
	st.WorkloadType = pb.WorkloadType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedEntityInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedEntityInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servedEntityOutputToPb(st *ServedEntityOutput) (*servedEntityOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedEntityOutputPb{}
	pb.CreationTimestamp = st.CreationTimestamp

	pb.Creator = st.Creator

	pb.EntityName = st.EntityName

	pb.EntityVersion = st.EntityVersion

	pb.EnvironmentVars = st.EnvironmentVars

	pb.ExternalModel = st.ExternalModel

	pb.FoundationModel = st.FoundationModel

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.MaxProvisionedThroughput = st.MaxProvisionedThroughput

	pb.MinProvisionedThroughput = st.MinProvisionedThroughput

	pb.Name = st.Name

	pb.ProvisionedModelUnits = st.ProvisionedModelUnits

	pb.ScaleToZeroEnabled = st.ScaleToZeroEnabled

	pb.State = st.State

	pb.WorkloadSize = st.WorkloadSize

	pb.WorkloadType = st.WorkloadType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servedEntityOutputPb struct {
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	Creator string `json:"creator,omitempty"`

	EntityName string `json:"entity_name,omitempty"`

	EntityVersion string `json:"entity_version,omitempty"`

	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`

	ExternalModel *ExternalModel `json:"external_model,omitempty"`

	FoundationModel *FoundationModel `json:"foundation_model,omitempty"`

	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`

	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`

	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`

	Name string `json:"name,omitempty"`

	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`

	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`

	State *ServedModelState `json:"state,omitempty"`

	WorkloadSize string `json:"workload_size,omitempty"`

	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedEntityOutputFromPb(pb *servedEntityOutputPb) (*ServedEntityOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedEntityOutput{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator
	st.EntityName = pb.EntityName
	st.EntityVersion = pb.EntityVersion
	st.EnvironmentVars = pb.EnvironmentVars
	st.ExternalModel = pb.ExternalModel
	st.FoundationModel = pb.FoundationModel
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.MaxProvisionedThroughput = pb.MaxProvisionedThroughput
	st.MinProvisionedThroughput = pb.MinProvisionedThroughput
	st.Name = pb.Name
	st.ProvisionedModelUnits = pb.ProvisionedModelUnits
	st.ScaleToZeroEnabled = pb.ScaleToZeroEnabled
	st.State = pb.State
	st.WorkloadSize = pb.WorkloadSize
	st.WorkloadType = pb.WorkloadType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedEntityOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedEntityOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servedEntitySpecToPb(st *ServedEntitySpec) (*servedEntitySpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedEntitySpecPb{}
	pb.EntityName = st.EntityName

	pb.EntityVersion = st.EntityVersion

	pb.ExternalModel = st.ExternalModel

	pb.FoundationModel = st.FoundationModel

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servedEntitySpecPb struct {
	EntityName string `json:"entity_name,omitempty"`

	EntityVersion string `json:"entity_version,omitempty"`

	ExternalModel *ExternalModel `json:"external_model,omitempty"`

	FoundationModel *FoundationModel `json:"foundation_model,omitempty"`

	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedEntitySpecFromPb(pb *servedEntitySpecPb) (*ServedEntitySpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedEntitySpec{}
	st.EntityName = pb.EntityName
	st.EntityVersion = pb.EntityVersion
	st.ExternalModel = pb.ExternalModel
	st.FoundationModel = pb.FoundationModel
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedEntitySpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedEntitySpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servedModelInputToPb(st *ServedModelInput) (*servedModelInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedModelInputPb{}
	pb.EnvironmentVars = st.EnvironmentVars

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.MaxProvisionedThroughput = st.MaxProvisionedThroughput

	pb.MinProvisionedThroughput = st.MinProvisionedThroughput

	pb.ModelName = st.ModelName

	pb.ModelVersion = st.ModelVersion

	pb.Name = st.Name

	pb.ProvisionedModelUnits = st.ProvisionedModelUnits

	pb.ScaleToZeroEnabled = st.ScaleToZeroEnabled

	pb.WorkloadSize = st.WorkloadSize

	pb.WorkloadType = st.WorkloadType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servedModelInputPb struct {
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`

	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`

	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`

	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`

	ModelName string `json:"model_name"`

	ModelVersion string `json:"model_version"`

	Name string `json:"name,omitempty"`

	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`

	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled"`

	WorkloadSize string `json:"workload_size,omitempty"`

	WorkloadType ServedModelInputWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedModelInputFromPb(pb *servedModelInputPb) (*ServedModelInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedModelInput{}
	st.EnvironmentVars = pb.EnvironmentVars
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.MaxProvisionedThroughput = pb.MaxProvisionedThroughput
	st.MinProvisionedThroughput = pb.MinProvisionedThroughput
	st.ModelName = pb.ModelName
	st.ModelVersion = pb.ModelVersion
	st.Name = pb.Name
	st.ProvisionedModelUnits = pb.ProvisionedModelUnits
	st.ScaleToZeroEnabled = pb.ScaleToZeroEnabled
	st.WorkloadSize = pb.WorkloadSize
	st.WorkloadType = pb.WorkloadType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedModelInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedModelInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servedModelOutputToPb(st *ServedModelOutput) (*servedModelOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedModelOutputPb{}
	pb.CreationTimestamp = st.CreationTimestamp

	pb.Creator = st.Creator

	pb.EnvironmentVars = st.EnvironmentVars

	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.ModelName = st.ModelName

	pb.ModelVersion = st.ModelVersion

	pb.Name = st.Name

	pb.ProvisionedModelUnits = st.ProvisionedModelUnits

	pb.ScaleToZeroEnabled = st.ScaleToZeroEnabled

	pb.State = st.State

	pb.WorkloadSize = st.WorkloadSize

	pb.WorkloadType = st.WorkloadType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servedModelOutputPb struct {
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	Creator string `json:"creator,omitempty"`

	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`

	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`

	ModelName string `json:"model_name,omitempty"`

	ModelVersion string `json:"model_version,omitempty"`

	Name string `json:"name,omitempty"`

	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`

	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`

	State *ServedModelState `json:"state,omitempty"`

	WorkloadSize string `json:"workload_size,omitempty"`

	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedModelOutputFromPb(pb *servedModelOutputPb) (*ServedModelOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedModelOutput{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator
	st.EnvironmentVars = pb.EnvironmentVars
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.ModelName = pb.ModelName
	st.ModelVersion = pb.ModelVersion
	st.Name = pb.Name
	st.ProvisionedModelUnits = pb.ProvisionedModelUnits
	st.ScaleToZeroEnabled = pb.ScaleToZeroEnabled
	st.State = pb.State
	st.WorkloadSize = pb.WorkloadSize
	st.WorkloadType = pb.WorkloadType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedModelOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedModelOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servedModelSpecToPb(st *ServedModelSpec) (*servedModelSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedModelSpecPb{}
	pb.ModelName = st.ModelName

	pb.ModelVersion = st.ModelVersion

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servedModelSpecPb struct {
	ModelName string `json:"model_name,omitempty"`

	ModelVersion string `json:"model_version,omitempty"`

	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedModelSpecFromPb(pb *servedModelSpecPb) (*ServedModelSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedModelSpec{}
	st.ModelName = pb.ModelName
	st.ModelVersion = pb.ModelVersion
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedModelSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedModelSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servedModelStateToPb(st *ServedModelState) (*servedModelStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedModelStatePb{}
	pb.Deployment = st.Deployment

	pb.DeploymentStateMessage = st.DeploymentStateMessage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servedModelStatePb struct {
	Deployment ServedModelStateDeployment `json:"deployment,omitempty"`

	DeploymentStateMessage string `json:"deployment_state_message,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedModelStateFromPb(pb *servedModelStatePb) (*ServedModelState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedModelState{}
	st.Deployment = pb.Deployment
	st.DeploymentStateMessage = pb.DeploymentStateMessage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedModelStatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedModelStatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func serverLogsResponseToPb(st *ServerLogsResponse) (*serverLogsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &serverLogsResponsePb{}
	pb.Logs = st.Logs

	return pb, nil
}

type serverLogsResponsePb struct {
	Logs string `json:"logs"`
}

func serverLogsResponseFromPb(pb *serverLogsResponsePb) (*ServerLogsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServerLogsResponse{}
	st.Logs = pb.Logs

	return st, nil
}

func servingEndpointToPb(st *ServingEndpoint) (*servingEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointPb{}
	pb.AiGateway = st.AiGateway

	pb.BudgetPolicyId = st.BudgetPolicyId

	pb.Config = st.Config

	pb.CreationTimestamp = st.CreationTimestamp

	pb.Creator = st.Creator

	pb.Id = st.Id

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	pb.Name = st.Name

	pb.State = st.State

	pb.Tags = st.Tags

	pb.Task = st.Task

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servingEndpointPb struct {
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`

	BudgetPolicyId string `json:"budget_policy_id,omitempty"`

	Config *EndpointCoreConfigSummary `json:"config,omitempty"`

	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	Creator string `json:"creator,omitempty"`

	Id string `json:"id,omitempty"`

	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`

	Name string `json:"name,omitempty"`

	State *EndpointState `json:"state,omitempty"`

	Tags []EndpointTag `json:"tags,omitempty"`

	Task string `json:"task,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointFromPb(pb *servingEndpointPb) (*ServingEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpoint{}
	st.AiGateway = pb.AiGateway
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Config = pb.Config
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Name = pb.Name
	st.State = pb.State
	st.Tags = pb.Tags
	st.Task = pb.Task

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servingEndpointAccessControlRequestToPb(st *ServingEndpointAccessControlRequest) (*servingEndpointAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointAccessControlRequestPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servingEndpointAccessControlRequestPb struct {
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointAccessControlRequestFromPb(pb *servingEndpointAccessControlRequestPb) (*ServingEndpointAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servingEndpointAccessControlResponseToPb(st *ServingEndpointAccessControlResponse) (*servingEndpointAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions

	pb.DisplayName = st.DisplayName

	pb.GroupName = st.GroupName

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servingEndpointAccessControlResponsePb struct {
	AllPermissions []ServingEndpointPermission `json:"all_permissions,omitempty"`

	DisplayName string `json:"display_name,omitempty"`

	GroupName string `json:"group_name,omitempty"`

	ServicePrincipalName string `json:"service_principal_name,omitempty"`

	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointAccessControlResponseFromPb(pb *servingEndpointAccessControlResponsePb) (*ServingEndpointAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servingEndpointDetailedToPb(st *ServingEndpointDetailed) (*servingEndpointDetailedPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointDetailedPb{}
	pb.AiGateway = st.AiGateway

	pb.BudgetPolicyId = st.BudgetPolicyId

	pb.Config = st.Config

	pb.CreationTimestamp = st.CreationTimestamp

	pb.Creator = st.Creator

	pb.DataPlaneInfo = st.DataPlaneInfo

	pb.EndpointUrl = st.EndpointUrl

	pb.Id = st.Id

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	pb.Name = st.Name

	pb.PendingConfig = st.PendingConfig

	pb.PermissionLevel = st.PermissionLevel

	pb.RouteOptimized = st.RouteOptimized

	pb.State = st.State

	pb.Tags = st.Tags

	pb.Task = st.Task

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servingEndpointDetailedPb struct {
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`

	BudgetPolicyId string `json:"budget_policy_id,omitempty"`

	Config *EndpointCoreConfigOutput `json:"config,omitempty"`

	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	Creator string `json:"creator,omitempty"`

	DataPlaneInfo *ModelDataPlaneInfo `json:"data_plane_info,omitempty"`

	EndpointUrl string `json:"endpoint_url,omitempty"`

	Id string `json:"id,omitempty"`

	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`

	Name string `json:"name,omitempty"`

	PendingConfig *EndpointPendingConfig `json:"pending_config,omitempty"`

	PermissionLevel ServingEndpointDetailedPermissionLevel `json:"permission_level,omitempty"`

	RouteOptimized bool `json:"route_optimized,omitempty"`

	State *EndpointState `json:"state,omitempty"`

	Tags []EndpointTag `json:"tags,omitempty"`

	Task string `json:"task,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointDetailedFromPb(pb *servingEndpointDetailedPb) (*ServingEndpointDetailed, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointDetailed{}
	st.AiGateway = pb.AiGateway
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Config = pb.Config
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator
	st.DataPlaneInfo = pb.DataPlaneInfo
	st.EndpointUrl = pb.EndpointUrl
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Name = pb.Name
	st.PendingConfig = pb.PendingConfig
	st.PermissionLevel = pb.PermissionLevel
	st.RouteOptimized = pb.RouteOptimized
	st.State = pb.State
	st.Tags = pb.Tags
	st.Task = pb.Task

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointDetailedPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointDetailedPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servingEndpointPermissionToPb(st *ServingEndpointPermission) (*servingEndpointPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointPermissionPb{}
	pb.Inherited = st.Inherited

	pb.InheritedFromObject = st.InheritedFromObject

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servingEndpointPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointPermissionFromPb(pb *servingEndpointPermissionPb) (*ServingEndpointPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servingEndpointPermissionsToPb(st *ServingEndpointPermissions) (*servingEndpointPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointPermissionsPb{}
	pb.AccessControlList = st.AccessControlList

	pb.ObjectId = st.ObjectId

	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servingEndpointPermissionsPb struct {
	AccessControlList []ServingEndpointAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointPermissionsFromPb(pb *servingEndpointPermissionsPb) (*ServingEndpointPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermissions{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servingEndpointPermissionsDescriptionToPb(st *ServingEndpointPermissionsDescription) (*servingEndpointPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointPermissionsDescriptionPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type servingEndpointPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`

	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointPermissionsDescriptionFromPb(pb *servingEndpointPermissionsDescriptionPb) (*ServingEndpointPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func servingEndpointPermissionsRequestToPb(st *ServingEndpointPermissionsRequest) (*servingEndpointPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointPermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList

	pb.ServingEndpointId = st.ServingEndpointId

	return pb, nil
}

type servingEndpointPermissionsRequestPb struct {
	AccessControlList []ServingEndpointAccessControlRequest `json:"access_control_list,omitempty"`

	ServingEndpointId string `json:"-" url:"-"`
}

func servingEndpointPermissionsRequestFromPb(pb *servingEndpointPermissionsRequestPb) (*ServingEndpointPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.ServingEndpointId = pb.ServingEndpointId

	return st, nil
}

func trafficConfigToPb(st *TrafficConfig) (*trafficConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trafficConfigPb{}
	pb.Routes = st.Routes

	return pb, nil
}

type trafficConfigPb struct {
	Routes []Route `json:"routes,omitempty"`
}

func trafficConfigFromPb(pb *trafficConfigPb) (*TrafficConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrafficConfig{}
	st.Routes = pb.Routes

	return st, nil
}

func updateProvisionedThroughputEndpointConfigRequestToPb(st *UpdateProvisionedThroughputEndpointConfigRequest) (*updateProvisionedThroughputEndpointConfigRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateProvisionedThroughputEndpointConfigRequestPb{}
	pb.Config = st.Config

	pb.Name = st.Name

	return pb, nil
}

type updateProvisionedThroughputEndpointConfigRequestPb struct {
	Config PtEndpointCoreConfig `json:"config"`

	Name string `json:"-" url:"-"`
}

func updateProvisionedThroughputEndpointConfigRequestFromPb(pb *updateProvisionedThroughputEndpointConfigRequestPb) (*UpdateProvisionedThroughputEndpointConfigRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProvisionedThroughputEndpointConfigRequest{}
	st.Config = pb.Config
	st.Name = pb.Name

	return st, nil
}

func v1ResponseChoiceElementToPb(st *V1ResponseChoiceElement) (*v1ResponseChoiceElementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &v1ResponseChoiceElementPb{}
	pb.FinishReason = st.FinishReason

	pb.Index = st.Index

	pb.Logprobs = st.Logprobs

	pb.Message = st.Message

	pb.Text = st.Text

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type v1ResponseChoiceElementPb struct {
	FinishReason string `json:"finishReason,omitempty"`

	Index int `json:"index,omitempty"`

	Logprobs int `json:"logprobs,omitempty"`

	Message *ChatMessage `json:"message,omitempty"`

	Text string `json:"text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func v1ResponseChoiceElementFromPb(pb *v1ResponseChoiceElementPb) (*V1ResponseChoiceElement, error) {
	if pb == nil {
		return nil, nil
	}
	st := &V1ResponseChoiceElement{}
	st.FinishReason = pb.FinishReason
	st.Index = pb.Index
	st.Logprobs = pb.Logprobs
	st.Message = pb.Message
	st.Text = pb.Text

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *v1ResponseChoiceElementPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st v1ResponseChoiceElementPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
