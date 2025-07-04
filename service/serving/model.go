// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/google/go-querystring/query"
)

type Ai21LabsConfig struct {
	// The Databricks secret key reference for an AI21 Labs API key. If you
	// prefer to paste your API key directly, see `ai21labs_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	// Wire name: 'ai21labs_api_key'
	Ai21labsApiKey string `json:"ai21labs_api_key,omitempty"`
	// An AI21 Labs API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `ai21labs_api_key`. You
	// must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	// Wire name: 'ai21labs_api_key_plaintext'
	Ai21labsApiKeyPlaintext string `json:"ai21labs_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Ai21LabsConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := ai21LabsConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Ai21LabsConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ai21LabsConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ai21LabsConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Ai21LabsConfig) MarshalJSON() ([]byte, error) {
	pb, err := ai21LabsConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AiGatewayConfig struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	// Wire name: 'fallback_config'
	FallbackConfig *FallbackConfig `json:"fallback_config,omitempty"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	// Wire name: 'guardrails'
	Guardrails *AiGatewayGuardrails `json:"guardrails,omitempty"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	// Wire name: 'inference_table_config'
	InferenceTableConfig *AiGatewayInferenceTableConfig `json:"inference_table_config,omitempty"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	// Wire name: 'rate_limits'
	RateLimits []AiGatewayRateLimit `json:"rate_limits,omitempty"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	// Wire name: 'usage_tracking_config'
	UsageTrackingConfig *AiGatewayUsageTrackingConfig `json:"usage_tracking_config,omitempty"`
}

func (st AiGatewayConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := aiGatewayConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AiGatewayConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aiGatewayConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aiGatewayConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AiGatewayConfig) MarshalJSON() ([]byte, error) {
	pb, err := aiGatewayConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AiGatewayGuardrailParameters struct {
	// List of invalid keywords. AI guardrail uses keyword or string matching to
	// decide if the keyword exists in the request or response content.
	// Wire name: 'invalid_keywords'
	InvalidKeywords []string `json:"invalid_keywords,omitempty"`
	// Configuration for guardrail PII filter.
	// Wire name: 'pii'
	Pii *AiGatewayGuardrailPiiBehavior `json:"pii,omitempty"`
	// Indicates whether the safety filter is enabled.
	// Wire name: 'safety'
	Safety bool `json:"safety,omitempty"`
	// The list of allowed topics. Given a chat request, this guardrail flags
	// the request if its topic is not in the allowed topics.
	// Wire name: 'valid_topics'
	ValidTopics []string `json:"valid_topics,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AiGatewayGuardrailParameters) EncodeValues(key string, v *url.Values) error {
	pb, err := aiGatewayGuardrailParametersToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AiGatewayGuardrailParameters) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aiGatewayGuardrailParametersPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aiGatewayGuardrailParametersFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AiGatewayGuardrailParameters) MarshalJSON() ([]byte, error) {
	pb, err := aiGatewayGuardrailParametersToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AiGatewayGuardrailPiiBehavior struct {
	// Configuration for input guardrail filters.
	// Wire name: 'behavior'
	Behavior AiGatewayGuardrailPiiBehaviorBehavior `json:"behavior,omitempty"`
}

func (st AiGatewayGuardrailPiiBehavior) EncodeValues(key string, v *url.Values) error {
	pb, err := aiGatewayGuardrailPiiBehaviorToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AiGatewayGuardrailPiiBehavior) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aiGatewayGuardrailPiiBehaviorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aiGatewayGuardrailPiiBehaviorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AiGatewayGuardrailPiiBehavior) MarshalJSON() ([]byte, error) {
	pb, err := aiGatewayGuardrailPiiBehaviorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AiGatewayGuardrailPiiBehaviorBehavior string

const AiGatewayGuardrailPiiBehaviorBehaviorBlock AiGatewayGuardrailPiiBehaviorBehavior = `BLOCK`

const AiGatewayGuardrailPiiBehaviorBehaviorNone AiGatewayGuardrailPiiBehaviorBehavior = `NONE`

// String representation for [fmt.Print]
func (f *AiGatewayGuardrailPiiBehaviorBehavior) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AiGatewayGuardrailPiiBehaviorBehavior) Set(v string) error {
	switch v {
	case `BLOCK`, `NONE`:
		*f = AiGatewayGuardrailPiiBehaviorBehavior(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "BLOCK", "NONE"`, v)
	}
}

// Values returns all possible values for AiGatewayGuardrailPiiBehaviorBehavior.
//
// There is no guarantee on the order of the values in the slice.
func (f *AiGatewayGuardrailPiiBehaviorBehavior) Values() []AiGatewayGuardrailPiiBehaviorBehavior {
	return []AiGatewayGuardrailPiiBehaviorBehavior{
		AiGatewayGuardrailPiiBehaviorBehaviorBlock,
		AiGatewayGuardrailPiiBehaviorBehaviorNone,
	}
}

// Type always returns AiGatewayGuardrailPiiBehaviorBehavior to satisfy [pflag.Value] interface
func (f *AiGatewayGuardrailPiiBehaviorBehavior) Type() string {
	return "AiGatewayGuardrailPiiBehaviorBehavior"
}

type AiGatewayGuardrails struct {
	// Configuration for input guardrail filters.
	// Wire name: 'input'
	Input *AiGatewayGuardrailParameters `json:"input,omitempty"`
	// Configuration for output guardrail filters.
	// Wire name: 'output'
	Output *AiGatewayGuardrailParameters `json:"output,omitempty"`
}

func (st AiGatewayGuardrails) EncodeValues(key string, v *url.Values) error {
	pb, err := aiGatewayGuardrailsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AiGatewayGuardrails) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aiGatewayGuardrailsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aiGatewayGuardrailsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AiGatewayGuardrails) MarshalJSON() ([]byte, error) {
	pb, err := aiGatewayGuardrailsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AiGatewayInferenceTableConfig struct {
	// The name of the catalog in Unity Catalog. Required when enabling
	// inference tables. NOTE: On update, you have to disable inference table
	// first in order to change the catalog name.
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// Indicates whether the inference table is enabled.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog. Required when enabling inference
	// tables. NOTE: On update, you have to disable inference table first in
	// order to change the schema name.
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you have to
	// disable inference table first in order to change the prefix name.
	// Wire name: 'table_name_prefix'
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AiGatewayInferenceTableConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := aiGatewayInferenceTableConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AiGatewayInferenceTableConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aiGatewayInferenceTableConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aiGatewayInferenceTableConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AiGatewayInferenceTableConfig) MarshalJSON() ([]byte, error) {
	pb, err := aiGatewayInferenceTableConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AiGatewayRateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	// Wire name: 'calls'
	Calls int64 `json:"calls"`
	// Key field for a rate limit. Currently, 'user', 'user_group,
	// 'service_principal', and 'endpoint' are supported, with 'endpoint' being
	// the default if not specified.
	// Wire name: 'key'
	Key AiGatewayRateLimitKey `json:"key,omitempty"`
	// Principal field for a user, user group, or service principal to apply
	// rate limiting to. Accepts a user email, group name, or service principal
	// application ID.
	// Wire name: 'principal'
	Principal string `json:"principal,omitempty"`
	// Renewal period field for a rate limit. Currently, only 'minute' is
	// supported.
	// Wire name: 'renewal_period'
	RenewalPeriod AiGatewayRateLimitRenewalPeriod `json:"renewal_period"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AiGatewayRateLimit) EncodeValues(key string, v *url.Values) error {
	pb, err := aiGatewayRateLimitToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AiGatewayRateLimit) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aiGatewayRateLimitPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aiGatewayRateLimitFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AiGatewayRateLimit) MarshalJSON() ([]byte, error) {
	pb, err := aiGatewayRateLimitToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AiGatewayRateLimitKey string

const AiGatewayRateLimitKeyEndpoint AiGatewayRateLimitKey = `endpoint`

const AiGatewayRateLimitKeyServicePrincipal AiGatewayRateLimitKey = `service_principal`

const AiGatewayRateLimitKeyUser AiGatewayRateLimitKey = `user`

const AiGatewayRateLimitKeyUserGroup AiGatewayRateLimitKey = `user_group`

// String representation for [fmt.Print]
func (f *AiGatewayRateLimitKey) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AiGatewayRateLimitKey) Set(v string) error {
	switch v {
	case `endpoint`, `service_principal`, `user`, `user_group`:
		*f = AiGatewayRateLimitKey(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "endpoint", "service_principal", "user", "user_group"`, v)
	}
}

// Values returns all possible values for AiGatewayRateLimitKey.
//
// There is no guarantee on the order of the values in the slice.
func (f *AiGatewayRateLimitKey) Values() []AiGatewayRateLimitKey {
	return []AiGatewayRateLimitKey{
		AiGatewayRateLimitKeyEndpoint,
		AiGatewayRateLimitKeyServicePrincipal,
		AiGatewayRateLimitKeyUser,
		AiGatewayRateLimitKeyUserGroup,
	}
}

// Type always returns AiGatewayRateLimitKey to satisfy [pflag.Value] interface
func (f *AiGatewayRateLimitKey) Type() string {
	return "AiGatewayRateLimitKey"
}

type AiGatewayRateLimitRenewalPeriod string

const AiGatewayRateLimitRenewalPeriodMinute AiGatewayRateLimitRenewalPeriod = `minute`

// String representation for [fmt.Print]
func (f *AiGatewayRateLimitRenewalPeriod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AiGatewayRateLimitRenewalPeriod) Set(v string) error {
	switch v {
	case `minute`:
		*f = AiGatewayRateLimitRenewalPeriod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "minute"`, v)
	}
}

// Values returns all possible values for AiGatewayRateLimitRenewalPeriod.
//
// There is no guarantee on the order of the values in the slice.
func (f *AiGatewayRateLimitRenewalPeriod) Values() []AiGatewayRateLimitRenewalPeriod {
	return []AiGatewayRateLimitRenewalPeriod{
		AiGatewayRateLimitRenewalPeriodMinute,
	}
}

// Type always returns AiGatewayRateLimitRenewalPeriod to satisfy [pflag.Value] interface
func (f *AiGatewayRateLimitRenewalPeriod) Type() string {
	return "AiGatewayRateLimitRenewalPeriod"
}

type AiGatewayUsageTrackingConfig struct {
	// Whether to enable usage tracking.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AiGatewayUsageTrackingConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := aiGatewayUsageTrackingConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AiGatewayUsageTrackingConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &aiGatewayUsageTrackingConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := aiGatewayUsageTrackingConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AiGatewayUsageTrackingConfig) MarshalJSON() ([]byte, error) {
	pb, err := aiGatewayUsageTrackingConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AmazonBedrockConfig struct {
	// The Databricks secret key reference for an AWS access key ID with
	// permissions to interact with Bedrock services. If you prefer to paste
	// your API key directly, see `aws_access_key_id_plaintext`. You must
	// provide an API key using one of the following fields: `aws_access_key_id`
	// or `aws_access_key_id_plaintext`.
	// Wire name: 'aws_access_key_id'
	AwsAccessKeyId string `json:"aws_access_key_id,omitempty"`
	// An AWS access key ID with permissions to interact with Bedrock services
	// provided as a plaintext string. If you prefer to reference your key using
	// Databricks Secrets, see `aws_access_key_id`. You must provide an API key
	// using one of the following fields: `aws_access_key_id` or
	// `aws_access_key_id_plaintext`.
	// Wire name: 'aws_access_key_id_plaintext'
	AwsAccessKeyIdPlaintext string `json:"aws_access_key_id_plaintext,omitempty"`
	// The AWS region to use. Bedrock has to be enabled there.
	// Wire name: 'aws_region'
	AwsRegion string `json:"aws_region"`
	// The Databricks secret key reference for an AWS secret access key paired
	// with the access key ID, with permissions to interact with Bedrock
	// services. If you prefer to paste your API key directly, see
	// `aws_secret_access_key_plaintext`. You must provide an API key using one
	// of the following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	// Wire name: 'aws_secret_access_key'
	AwsSecretAccessKey string `json:"aws_secret_access_key,omitempty"`
	// An AWS secret access key paired with the access key ID, with permissions
	// to interact with Bedrock services provided as a plaintext string. If you
	// prefer to reference your key using Databricks Secrets, see
	// `aws_secret_access_key`. You must provide an API key using one of the
	// following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	// Wire name: 'aws_secret_access_key_plaintext'
	AwsSecretAccessKeyPlaintext string `json:"aws_secret_access_key_plaintext,omitempty"`
	// The underlying provider in Amazon Bedrock. Supported values (case
	// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
	// Wire name: 'bedrock_provider'
	BedrockProvider AmazonBedrockConfigBedrockProvider `json:"bedrock_provider"`
	// ARN of the instance profile that the external model will use to access
	// AWS resources. You must authenticate using an instance profile or access
	// keys. If you prefer to authenticate using access keys, see
	// `aws_access_key_id`, `aws_access_key_id_plaintext`,
	// `aws_secret_access_key` and `aws_secret_access_key_plaintext`.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AmazonBedrockConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := amazonBedrockConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AmazonBedrockConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &amazonBedrockConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := amazonBedrockConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AmazonBedrockConfig) MarshalJSON() ([]byte, error) {
	pb, err := amazonBedrockConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AmazonBedrockConfigBedrockProvider string

const AmazonBedrockConfigBedrockProviderAi21labs AmazonBedrockConfigBedrockProvider = `ai21labs`

const AmazonBedrockConfigBedrockProviderAmazon AmazonBedrockConfigBedrockProvider = `amazon`

const AmazonBedrockConfigBedrockProviderAnthropic AmazonBedrockConfigBedrockProvider = `anthropic`

const AmazonBedrockConfigBedrockProviderCohere AmazonBedrockConfigBedrockProvider = `cohere`

// String representation for [fmt.Print]
func (f *AmazonBedrockConfigBedrockProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AmazonBedrockConfigBedrockProvider) Set(v string) error {
	switch v {
	case `ai21labs`, `amazon`, `anthropic`, `cohere`:
		*f = AmazonBedrockConfigBedrockProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ai21labs", "amazon", "anthropic", "cohere"`, v)
	}
}

// Values returns all possible values for AmazonBedrockConfigBedrockProvider.
//
// There is no guarantee on the order of the values in the slice.
func (f *AmazonBedrockConfigBedrockProvider) Values() []AmazonBedrockConfigBedrockProvider {
	return []AmazonBedrockConfigBedrockProvider{
		AmazonBedrockConfigBedrockProviderAi21labs,
		AmazonBedrockConfigBedrockProviderAmazon,
		AmazonBedrockConfigBedrockProviderAnthropic,
		AmazonBedrockConfigBedrockProviderCohere,
	}
}

// Type always returns AmazonBedrockConfigBedrockProvider to satisfy [pflag.Value] interface
func (f *AmazonBedrockConfigBedrockProvider) Type() string {
	return "AmazonBedrockConfigBedrockProvider"
}

type AnthropicConfig struct {
	// The Databricks secret key reference for an Anthropic API key. If you
	// prefer to paste your API key directly, see `anthropic_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	// Wire name: 'anthropic_api_key'
	AnthropicApiKey string `json:"anthropic_api_key,omitempty"`
	// The Anthropic API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `anthropic_api_key`. You
	// must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	// Wire name: 'anthropic_api_key_plaintext'
	AnthropicApiKeyPlaintext string `json:"anthropic_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AnthropicConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := anthropicConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AnthropicConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &anthropicConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := anthropicConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AnthropicConfig) MarshalJSON() ([]byte, error) {
	pb, err := anthropicConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ApiKeyAuth struct {
	// The name of the API key parameter used for authentication.
	// Wire name: 'key'
	Key string `json:"key"`
	// The Databricks secret key reference for an API Key. If you prefer to
	// paste your token directly, see `value_plaintext`.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`
	// The API Key provided as a plaintext string. If you prefer to reference
	// your token using Databricks Secrets, see `value`.
	// Wire name: 'value_plaintext'
	ValuePlaintext string `json:"value_plaintext,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ApiKeyAuth) EncodeValues(key string, v *url.Values) error {
	pb, err := apiKeyAuthToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ApiKeyAuth) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &apiKeyAuthPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := apiKeyAuthFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ApiKeyAuth) MarshalJSON() ([]byte, error) {
	pb, err := apiKeyAuthToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AutoCaptureConfigInput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// Indicates whether the inference table is enabled.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	// Wire name: 'table_name_prefix'
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AutoCaptureConfigInput) EncodeValues(key string, v *url.Values) error {
	pb, err := autoCaptureConfigInputToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AutoCaptureConfigInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &autoCaptureConfigInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := autoCaptureConfigInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AutoCaptureConfigInput) MarshalJSON() ([]byte, error) {
	pb, err := autoCaptureConfigInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AutoCaptureConfigOutput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	// Wire name: 'catalog_name'
	CatalogName string `json:"catalog_name,omitempty"`
	// Indicates whether the inference table is enabled.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	// Wire name: 'schema_name'
	SchemaName string `json:"schema_name,omitempty"`

	// Wire name: 'state'
	State *AutoCaptureState `json:"state,omitempty"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	// Wire name: 'table_name_prefix'
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st AutoCaptureConfigOutput) EncodeValues(key string, v *url.Values) error {
	pb, err := autoCaptureConfigOutputToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AutoCaptureConfigOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &autoCaptureConfigOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := autoCaptureConfigOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AutoCaptureConfigOutput) MarshalJSON() ([]byte, error) {
	pb, err := autoCaptureConfigOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type AutoCaptureState struct {

	// Wire name: 'payload_table'
	PayloadTable *PayloadTable `json:"payload_table,omitempty"`
}

func (st AutoCaptureState) EncodeValues(key string, v *url.Values) error {
	pb, err := autoCaptureStateToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *AutoCaptureState) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &autoCaptureStatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := autoCaptureStateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AutoCaptureState) MarshalJSON() ([]byte, error) {
	pb, err := autoCaptureStateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type BearerTokenAuth struct {
	// The Databricks secret key reference for a token. If you prefer to paste
	// your token directly, see `token_plaintext`.
	// Wire name: 'token'
	Token string `json:"token,omitempty"`
	// The token provided as a plaintext string. If you prefer to reference your
	// token using Databricks Secrets, see `token`.
	// Wire name: 'token_plaintext'
	TokenPlaintext string `json:"token_plaintext,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st BearerTokenAuth) EncodeValues(key string, v *url.Values) error {
	pb, err := bearerTokenAuthToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *BearerTokenAuth) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &bearerTokenAuthPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := bearerTokenAuthFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BearerTokenAuth) MarshalJSON() ([]byte, error) {
	pb, err := bearerTokenAuthToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type BuildLogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" tf:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName string `json:"-" tf:"-"`
}

func (st BuildLogsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := buildLogsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *BuildLogsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &buildLogsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := buildLogsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BuildLogsRequest) MarshalJSON() ([]byte, error) {
	pb, err := buildLogsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type BuildLogsResponse struct {
	// The logs associated with building the served entity's environment.
	// Wire name: 'logs'
	Logs string `json:"logs"`
}

func (st BuildLogsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := buildLogsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *BuildLogsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &buildLogsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := buildLogsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st BuildLogsResponse) MarshalJSON() ([]byte, error) {
	pb, err := buildLogsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ChatMessage struct {
	// The content of the message.
	// Wire name: 'content'
	Content string `json:"content,omitempty"`
	// The role of the message. One of [system, user, assistant].
	// Wire name: 'role'
	Role ChatMessageRole `json:"role,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ChatMessage) EncodeValues(key string, v *url.Values) error {
	pb, err := chatMessageToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ChatMessage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &chatMessagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := chatMessageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ChatMessage) MarshalJSON() ([]byte, error) {
	pb, err := chatMessageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for ChatMessageRole.
//
// There is no guarantee on the order of the values in the slice.
func (f *ChatMessageRole) Values() []ChatMessageRole {
	return []ChatMessageRole{
		ChatMessageRoleAssistant,
		ChatMessageRoleSystem,
		ChatMessageRoleUser,
	}
}

// Type always returns ChatMessageRole to satisfy [pflag.Value] interface
func (f *ChatMessageRole) Type() string {
	return "ChatMessageRole"
}

type CohereConfig struct {
	// This is an optional field to provide a customized base URL for the Cohere
	// API. If left unspecified, the standard Cohere base URL is used.
	// Wire name: 'cohere_api_base'
	CohereApiBase string `json:"cohere_api_base,omitempty"`
	// The Databricks secret key reference for a Cohere API key. If you prefer
	// to paste your API key directly, see `cohere_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `cohere_api_key` or
	// `cohere_api_key_plaintext`.
	// Wire name: 'cohere_api_key'
	CohereApiKey string `json:"cohere_api_key,omitempty"`
	// The Cohere API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `cohere_api_key`. You
	// must provide an API key using one of the following fields:
	// `cohere_api_key` or `cohere_api_key_plaintext`.
	// Wire name: 'cohere_api_key_plaintext'
	CohereApiKeyPlaintext string `json:"cohere_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CohereConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := cohereConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CohereConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &cohereConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := cohereConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CohereConfig) MarshalJSON() ([]byte, error) {
	pb, err := cohereConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreatePtEndpointRequest struct {
	// The AI Gateway configuration for the serving endpoint.
	// Wire name: 'ai_gateway'
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`
	// The budget policy associated with the endpoint.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The core config of the serving endpoint.
	// Wire name: 'config'
	Config PtEndpointCoreConfig `json:"config"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	// Wire name: 'name'
	Name string `json:"name"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	// Wire name: 'tags'
	Tags []EndpointTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreatePtEndpointRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := createPtEndpointRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreatePtEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createPtEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createPtEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreatePtEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := createPtEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type CreateServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	// Wire name: 'ai_gateway'
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`
	// The budget policy to be applied to the serving endpoint.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The core config of the serving endpoint.
	// Wire name: 'config'
	Config *EndpointCoreConfigInput `json:"config,omitempty"`

	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	// Wire name: 'name'
	Name string `json:"name"`
	// Rate limits to be applied to the serving endpoint. NOTE: this field is
	// deprecated, please use AI Gateway to manage rate limits.
	// Wire name: 'rate_limits'
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
	// Enable route optimization for the serving endpoint.
	// Wire name: 'route_optimized'
	RouteOptimized bool `json:"route_optimized,omitempty"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	// Wire name: 'tags'
	Tags []EndpointTag `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st CreateServingEndpoint) EncodeValues(key string, v *url.Values) error {
	pb, err := createServingEndpointToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CreateServingEndpoint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createServingEndpointPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createServingEndpointFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateServingEndpoint) MarshalJSON() ([]byte, error) {
	pb, err := createServingEndpointToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Configs needed to create a custom provider model route.
type CustomProviderConfig struct {
	// This is a field to provide API key authentication for the custom provider
	// API. You can only specify one authentication method.
	// Wire name: 'api_key_auth'
	ApiKeyAuth *ApiKeyAuth `json:"api_key_auth,omitempty"`
	// This is a field to provide bearer token authentication for the custom
	// provider API. You can only specify one authentication method.
	// Wire name: 'bearer_token_auth'
	BearerTokenAuth *BearerTokenAuth `json:"bearer_token_auth,omitempty"`
	// This is a field to provide the URL of the custom provider API.
	// Wire name: 'custom_provider_url'
	CustomProviderUrl string `json:"custom_provider_url"`
}

func (st CustomProviderConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := customProviderConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *CustomProviderConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &customProviderConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := customProviderConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CustomProviderConfig) MarshalJSON() ([]byte, error) {
	pb, err := customProviderConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Details necessary to query this object's API through the DataPlane APIs.
type DataPlaneInfo struct {
	// Authorization details as a string.
	// Wire name: 'authorization_details'
	AuthorizationDetails string `json:"authorization_details,omitempty"`
	// The URL of the endpoint for this operation in the dataplane.
	// Wire name: 'endpoint_url'
	EndpointUrl string `json:"endpoint_url,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DataPlaneInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := dataPlaneInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DataPlaneInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dataPlaneInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dataPlaneInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DataPlaneInfo) MarshalJSON() ([]byte, error) {
	pb, err := dataPlaneInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DatabricksModelServingConfig struct {
	// The Databricks secret key reference for a Databricks API token that
	// corresponds to a user or service principal with Can Query access to the
	// model serving endpoint pointed to by this external model. If you prefer
	// to paste your API key directly, see `databricks_api_token_plaintext`. You
	// must provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	// Wire name: 'databricks_api_token'
	DatabricksApiToken string `json:"databricks_api_token,omitempty"`
	// The Databricks API token that corresponds to a user or service principal
	// with Can Query access to the model serving endpoint pointed to by this
	// external model provided as a plaintext string. If you prefer to reference
	// your key using Databricks Secrets, see `databricks_api_token`. You must
	// provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	// Wire name: 'databricks_api_token_plaintext'
	DatabricksApiTokenPlaintext string `json:"databricks_api_token_plaintext,omitempty"`
	// The URL of the Databricks workspace containing the model serving endpoint
	// pointed to by this external model.
	// Wire name: 'databricks_workspace_url'
	DatabricksWorkspaceUrl string `json:"databricks_workspace_url"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st DatabricksModelServingConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := databricksModelServingConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DatabricksModelServingConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &databricksModelServingConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := databricksModelServingConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatabricksModelServingConfig) MarshalJSON() ([]byte, error) {
	pb, err := databricksModelServingConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DataframeSplitInput struct {

	// Wire name: 'columns'
	Columns []any `json:"columns,omitempty"`

	// Wire name: 'data'
	Data []any `json:"data,omitempty"`

	// Wire name: 'index'
	Index []int `json:"index,omitempty"`
}

func (st DataframeSplitInput) EncodeValues(key string, v *url.Values) error {
	pb, err := dataframeSplitInputToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DataframeSplitInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &dataframeSplitInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := dataframeSplitInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DataframeSplitInput) MarshalJSON() ([]byte, error) {
	pb, err := dataframeSplitInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteServingEndpointRequest struct {
	Name string `json:"-" tf:"-"`
}

func (st DeleteServingEndpointRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := deleteServingEndpointRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *DeleteServingEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteServingEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteServingEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteServingEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteServingEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EmbeddingsV1ResponseEmbeddingElement struct {

	// Wire name: 'embedding'
	Embedding []float64 `json:"embedding,omitempty"`
	// The index of the embedding in the response.
	// Wire name: 'index'
	Index int `json:"index,omitempty"`
	// This will always be 'embedding'.
	// Wire name: 'object'
	Object EmbeddingsV1ResponseEmbeddingElementObject `json:"object,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EmbeddingsV1ResponseEmbeddingElement) EncodeValues(key string, v *url.Values) error {
	pb, err := embeddingsV1ResponseEmbeddingElementToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EmbeddingsV1ResponseEmbeddingElement) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &embeddingsV1ResponseEmbeddingElementPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := embeddingsV1ResponseEmbeddingElementFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EmbeddingsV1ResponseEmbeddingElement) MarshalJSON() ([]byte, error) {
	pb, err := embeddingsV1ResponseEmbeddingElementToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for EmbeddingsV1ResponseEmbeddingElementObject.
//
// There is no guarantee on the order of the values in the slice.
func (f *EmbeddingsV1ResponseEmbeddingElementObject) Values() []EmbeddingsV1ResponseEmbeddingElementObject {
	return []EmbeddingsV1ResponseEmbeddingElementObject{
		EmbeddingsV1ResponseEmbeddingElementObjectEmbedding,
	}
}

// Type always returns EmbeddingsV1ResponseEmbeddingElementObject to satisfy [pflag.Value] interface
func (f *EmbeddingsV1ResponseEmbeddingElementObject) Type() string {
	return "EmbeddingsV1ResponseEmbeddingElementObject"
}

type EndpointCoreConfigInput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	// Wire name: 'auto_capture_config'
	AutoCaptureConfig *AutoCaptureConfigInput `json:"auto_capture_config,omitempty"`
	// The name of the serving endpoint to update. This field is required.
	Name string `json:"-" tf:"-"`
	// The list of served entities under the serving endpoint config.
	// Wire name: 'served_entities'
	ServedEntities []ServedEntityInput `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	// Wire name: 'served_models'
	ServedModels []ServedModelInput `json:"served_models,omitempty"`
	// The traffic configuration associated with the serving endpoint config.
	// Wire name: 'traffic_config'
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`
}

func (st EndpointCoreConfigInput) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointCoreConfigInputToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointCoreConfigInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointCoreConfigInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointCoreConfigInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointCoreConfigInput) MarshalJSON() ([]byte, error) {
	pb, err := endpointCoreConfigInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EndpointCoreConfigOutput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	// Wire name: 'auto_capture_config'
	AutoCaptureConfig *AutoCaptureConfigOutput `json:"auto_capture_config,omitempty"`
	// The config version that the serving endpoint is currently serving.
	// Wire name: 'config_version'
	ConfigVersion int64 `json:"config_version,omitempty"`
	// The list of served entities under the serving endpoint config.
	// Wire name: 'served_entities'
	ServedEntities []ServedEntityOutput `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	// Wire name: 'served_models'
	ServedModels []ServedModelOutput `json:"served_models,omitempty"`
	// The traffic configuration associated with the serving endpoint config.
	// Wire name: 'traffic_config'
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EndpointCoreConfigOutput) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointCoreConfigOutputToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointCoreConfigOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointCoreConfigOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointCoreConfigOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointCoreConfigOutput) MarshalJSON() ([]byte, error) {
	pb, err := endpointCoreConfigOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EndpointCoreConfigSummary struct {
	// The list of served entities under the serving endpoint config.
	// Wire name: 'served_entities'
	ServedEntities []ServedEntitySpec `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	// Wire name: 'served_models'
	ServedModels []ServedModelSpec `json:"served_models,omitempty"`
}

func (st EndpointCoreConfigSummary) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointCoreConfigSummaryToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointCoreConfigSummary) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointCoreConfigSummaryPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointCoreConfigSummaryFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointCoreConfigSummary) MarshalJSON() ([]byte, error) {
	pb, err := endpointCoreConfigSummaryToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EndpointPendingConfig struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	// Wire name: 'auto_capture_config'
	AutoCaptureConfig *AutoCaptureConfigOutput `json:"auto_capture_config,omitempty"`
	// The config version that the serving endpoint is currently serving.
	// Wire name: 'config_version'
	ConfigVersion int `json:"config_version,omitempty"`
	// The list of served entities belonging to the last issued update to the
	// serving endpoint.
	// Wire name: 'served_entities'
	ServedEntities []ServedEntityOutput `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models
	// belonging to the last issued update to the serving endpoint.
	// Wire name: 'served_models'
	ServedModels []ServedModelOutput `json:"served_models,omitempty"`
	// The timestamp when the update to the pending config started.
	// Wire name: 'start_time'
	StartTime int64 `json:"start_time,omitempty"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	// Wire name: 'traffic_config'
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EndpointPendingConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointPendingConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointPendingConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointPendingConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointPendingConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointPendingConfig) MarshalJSON() ([]byte, error) {
	pb, err := endpointPendingConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EndpointState struct {
	// The state of an endpoint's config update. This informs the user if the
	// pending_config is in progress, if the update failed, or if there is no
	// update in progress. Note that if the endpoint's config_update state value
	// is IN_PROGRESS, another update can not be made until the update completes
	// or fails.
	// Wire name: 'config_update'
	ConfigUpdate EndpointStateConfigUpdate `json:"config_update,omitempty"`
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served entities in its
	// active configuration are ready. If any of the actively served entities
	// are in a non-ready state, the endpoint state will be NOT_READY.
	// Wire name: 'ready'
	Ready EndpointStateReady `json:"ready,omitempty"`
}

func (st EndpointState) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointStateToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointState) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointStatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointStateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointState) MarshalJSON() ([]byte, error) {
	pb, err := endpointStateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EndpointStateConfigUpdate string

const EndpointStateConfigUpdateInProgress EndpointStateConfigUpdate = `IN_PROGRESS`

const EndpointStateConfigUpdateNotUpdating EndpointStateConfigUpdate = `NOT_UPDATING`

const EndpointStateConfigUpdateUpdateCanceled EndpointStateConfigUpdate = `UPDATE_CANCELED`

const EndpointStateConfigUpdateUpdateFailed EndpointStateConfigUpdate = `UPDATE_FAILED`

// String representation for [fmt.Print]
func (f *EndpointStateConfigUpdate) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *EndpointStateConfigUpdate) Set(v string) error {
	switch v {
	case `IN_PROGRESS`, `NOT_UPDATING`, `UPDATE_CANCELED`, `UPDATE_FAILED`:
		*f = EndpointStateConfigUpdate(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "IN_PROGRESS", "NOT_UPDATING", "UPDATE_CANCELED", "UPDATE_FAILED"`, v)
	}
}

// Values returns all possible values for EndpointStateConfigUpdate.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointStateConfigUpdate) Values() []EndpointStateConfigUpdate {
	return []EndpointStateConfigUpdate{
		EndpointStateConfigUpdateInProgress,
		EndpointStateConfigUpdateNotUpdating,
		EndpointStateConfigUpdateUpdateCanceled,
		EndpointStateConfigUpdateUpdateFailed,
	}
}

// Type always returns EndpointStateConfigUpdate to satisfy [pflag.Value] interface
func (f *EndpointStateConfigUpdate) Type() string {
	return "EndpointStateConfigUpdate"
}

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

// Values returns all possible values for EndpointStateReady.
//
// There is no guarantee on the order of the values in the slice.
func (f *EndpointStateReady) Values() []EndpointStateReady {
	return []EndpointStateReady{
		EndpointStateReadyNotReady,
		EndpointStateReadyReady,
	}
}

// Type always returns EndpointStateReady to satisfy [pflag.Value] interface
func (f *EndpointStateReady) Type() string {
	return "EndpointStateReady"
}

type EndpointTag struct {
	// Key field for a serving endpoint tag.
	// Wire name: 'key'
	Key string `json:"key"`
	// Optional value field for a serving endpoint tag.
	// Wire name: 'value'
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st EndpointTag) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointTagToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointTag) MarshalJSON() ([]byte, error) {
	pb, err := endpointTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type EndpointTags struct {

	// Wire name: 'tags'
	Tags []EndpointTag `json:"tags,omitempty"`
}

func (st EndpointTags) EncodeValues(key string, v *url.Values) error {
	pb, err := endpointTagsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *EndpointTags) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &endpointTagsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := endpointTagsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st EndpointTags) MarshalJSON() ([]byte, error) {
	pb, err := endpointTagsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name string `json:"-" tf:"-"`
}

func (st ExportMetricsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := exportMetricsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExportMetricsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exportMetricsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exportMetricsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExportMetricsRequest) MarshalJSON() ([]byte, error) {
	pb, err := exportMetricsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExportMetricsResponse struct {
	Contents io.ReadCloser `json:"-" tf:"-"`
}

func (st ExportMetricsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := exportMetricsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExportMetricsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &exportMetricsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := exportMetricsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExportMetricsResponse) MarshalJSON() ([]byte, error) {
	pb, err := exportMetricsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Simple Proto message for testing
type ExternalFunctionRequest struct {
	// The connection name to use. This is required to identify the external
	// connection.
	// Wire name: 'connection_name'
	ConnectionName string `json:"connection_name"`
	// Additional headers for the request. If not provided, only auth headers
	// from connections would be passed.
	// Wire name: 'headers'
	Headers string `json:"headers,omitempty"`
	// The JSON payload to send in the request body.
	// Wire name: 'json'
	Json string `json:"json,omitempty"`
	// The HTTP method to use (e.g., 'GET', 'POST').
	// Wire name: 'method'
	Method ExternalFunctionRequestHttpMethod `json:"method"`
	// Query parameters for the request.
	// Wire name: 'params'
	Params string `json:"params,omitempty"`
	// The relative path for the API endpoint. This is required.
	// Wire name: 'path'
	Path string `json:"path"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalFunctionRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := externalFunctionRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalFunctionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalFunctionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalFunctionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalFunctionRequest) MarshalJSON() ([]byte, error) {
	pb, err := externalFunctionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalFunctionRequestHttpMethod string

const ExternalFunctionRequestHttpMethodDelete ExternalFunctionRequestHttpMethod = `DELETE`

const ExternalFunctionRequestHttpMethodGet ExternalFunctionRequestHttpMethod = `GET`

const ExternalFunctionRequestHttpMethodPatch ExternalFunctionRequestHttpMethod = `PATCH`

const ExternalFunctionRequestHttpMethodPost ExternalFunctionRequestHttpMethod = `POST`

const ExternalFunctionRequestHttpMethodPut ExternalFunctionRequestHttpMethod = `PUT`

// String representation for [fmt.Print]
func (f *ExternalFunctionRequestHttpMethod) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExternalFunctionRequestHttpMethod) Set(v string) error {
	switch v {
	case `DELETE`, `GET`, `PATCH`, `POST`, `PUT`:
		*f = ExternalFunctionRequestHttpMethod(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETE", "GET", "PATCH", "POST", "PUT"`, v)
	}
}

// Values returns all possible values for ExternalFunctionRequestHttpMethod.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExternalFunctionRequestHttpMethod) Values() []ExternalFunctionRequestHttpMethod {
	return []ExternalFunctionRequestHttpMethod{
		ExternalFunctionRequestHttpMethodDelete,
		ExternalFunctionRequestHttpMethodGet,
		ExternalFunctionRequestHttpMethodPatch,
		ExternalFunctionRequestHttpMethodPost,
		ExternalFunctionRequestHttpMethodPut,
	}
}

// Type always returns ExternalFunctionRequestHttpMethod to satisfy [pflag.Value] interface
func (f *ExternalFunctionRequestHttpMethod) Type() string {
	return "ExternalFunctionRequestHttpMethod"
}

type ExternalModel struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	// Wire name: 'ai21labs_config'
	Ai21labsConfig *Ai21LabsConfig `json:"ai21labs_config,omitempty"`
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	// Wire name: 'amazon_bedrock_config'
	AmazonBedrockConfig *AmazonBedrockConfig `json:"amazon_bedrock_config,omitempty"`
	// Anthropic Config. Only required if the provider is 'anthropic'.
	// Wire name: 'anthropic_config'
	AnthropicConfig *AnthropicConfig `json:"anthropic_config,omitempty"`
	// Cohere Config. Only required if the provider is 'cohere'.
	// Wire name: 'cohere_config'
	CohereConfig *CohereConfig `json:"cohere_config,omitempty"`
	// Custom Provider Config. Only required if the provider is 'custom'.
	// Wire name: 'custom_provider_config'
	CustomProviderConfig *CustomProviderConfig `json:"custom_provider_config,omitempty"`
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	// Wire name: 'databricks_model_serving_config'
	DatabricksModelServingConfig *DatabricksModelServingConfig `json:"databricks_model_serving_config,omitempty"`
	// Google Cloud Vertex AI Config. Only required if the provider is
	// 'google-cloud-vertex-ai'.
	// Wire name: 'google_cloud_vertex_ai_config'
	GoogleCloudVertexAiConfig *GoogleCloudVertexAiConfig `json:"google_cloud_vertex_ai_config,omitempty"`
	// The name of the external model.
	// Wire name: 'name'
	Name string `json:"name"`
	// OpenAI Config. Only required if the provider is 'openai'.
	// Wire name: 'openai_config'
	OpenaiConfig *OpenAiConfig `json:"openai_config,omitempty"`
	// PaLM Config. Only required if the provider is 'palm'.
	// Wire name: 'palm_config'
	PalmConfig *PaLmConfig `json:"palm_config,omitempty"`
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
	// 'databricks-model-serving', 'google-cloud-vertex-ai', 'openai', 'palm',
	// and 'custom'.
	// Wire name: 'provider'
	Provider ExternalModelProvider `json:"provider"`
	// The task type of the external model.
	// Wire name: 'task'
	Task string `json:"task"`
}

func (st ExternalModel) EncodeValues(key string, v *url.Values) error {
	pb, err := externalModelToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalModel) MarshalJSON() ([]byte, error) {
	pb, err := externalModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ExternalModelProvider string

const ExternalModelProviderAi21labs ExternalModelProvider = `ai21labs`

const ExternalModelProviderAmazonBedrock ExternalModelProvider = `amazon-bedrock`

const ExternalModelProviderAnthropic ExternalModelProvider = `anthropic`

const ExternalModelProviderCohere ExternalModelProvider = `cohere`

const ExternalModelProviderCustom ExternalModelProvider = `custom`

const ExternalModelProviderDatabricksModelServing ExternalModelProvider = `databricks-model-serving`

const ExternalModelProviderGoogleCloudVertexAi ExternalModelProvider = `google-cloud-vertex-ai`

const ExternalModelProviderOpenai ExternalModelProvider = `openai`

const ExternalModelProviderPalm ExternalModelProvider = `palm`

// String representation for [fmt.Print]
func (f *ExternalModelProvider) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExternalModelProvider) Set(v string) error {
	switch v {
	case `ai21labs`, `amazon-bedrock`, `anthropic`, `cohere`, `custom`, `databricks-model-serving`, `google-cloud-vertex-ai`, `openai`, `palm`:
		*f = ExternalModelProvider(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ai21labs", "amazon-bedrock", "anthropic", "cohere", "custom", "databricks-model-serving", "google-cloud-vertex-ai", "openai", "palm"`, v)
	}
}

// Values returns all possible values for ExternalModelProvider.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExternalModelProvider) Values() []ExternalModelProvider {
	return []ExternalModelProvider{
		ExternalModelProviderAi21labs,
		ExternalModelProviderAmazonBedrock,
		ExternalModelProviderAnthropic,
		ExternalModelProviderCohere,
		ExternalModelProviderCustom,
		ExternalModelProviderDatabricksModelServing,
		ExternalModelProviderGoogleCloudVertexAi,
		ExternalModelProviderOpenai,
		ExternalModelProviderPalm,
	}
}

// Type always returns ExternalModelProvider to satisfy [pflag.Value] interface
func (f *ExternalModelProvider) Type() string {
	return "ExternalModelProvider"
}

type ExternalModelUsageElement struct {
	// The number of tokens in the chat/completions response.
	// Wire name: 'completion_tokens'
	CompletionTokens int `json:"completion_tokens,omitempty"`
	// The number of tokens in the prompt.
	// Wire name: 'prompt_tokens'
	PromptTokens int `json:"prompt_tokens,omitempty"`
	// The total number of tokens in the prompt and response.
	// Wire name: 'total_tokens'
	TotalTokens int `json:"total_tokens,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ExternalModelUsageElement) EncodeValues(key string, v *url.Values) error {
	pb, err := externalModelUsageElementToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ExternalModelUsageElement) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &externalModelUsageElementPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := externalModelUsageElementFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExternalModelUsageElement) MarshalJSON() ([]byte, error) {
	pb, err := externalModelUsageElementToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type FallbackConfig struct {
	// Whether to enable traffic fallback. When a served entity in the serving
	// endpoint returns specific error codes (e.g. 500), the request will
	// automatically be round-robin attempted with other served entities in the
	// same endpoint, following the order of served entity list, until a
	// successful response is returned. If all attempts fail, return the last
	// response with the error code.
	// Wire name: 'enabled'
	Enabled bool `json:"enabled"`
}

func (st FallbackConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := fallbackConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *FallbackConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fallbackConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fallbackConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FallbackConfig) MarshalJSON() ([]byte, error) {
	pb, err := fallbackConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// All fields are not sensitive as they are hard-coded in the system and made
// available to customers.
type FoundationModel struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`

	// Wire name: 'docs'
	Docs string `json:"docs,omitempty"`

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st FoundationModel) EncodeValues(key string, v *url.Values) error {
	pb, err := foundationModelToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *FoundationModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &foundationModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := foundationModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FoundationModel) MarshalJSON() ([]byte, error) {
	pb, err := foundationModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetOpenApiRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" tf:"-"`
}

func (st GetOpenApiRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getOpenApiRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetOpenApiRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getOpenApiRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getOpenApiRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetOpenApiRequest) MarshalJSON() ([]byte, error) {
	pb, err := getOpenApiRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetOpenApiResponse struct {
	Contents io.ReadCloser `json:"-" tf:"-"`
}

func (st GetOpenApiResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getOpenApiResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetOpenApiResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getOpenApiResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getOpenApiResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetOpenApiResponse) MarshalJSON() ([]byte, error) {
	pb, err := getOpenApiResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetServingEndpointPermissionLevelsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" tf:"-"`
}

func (st GetServingEndpointPermissionLevelsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getServingEndpointPermissionLevelsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetServingEndpointPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getServingEndpointPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getServingEndpointPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetServingEndpointPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getServingEndpointPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []ServingEndpointPermissionsDescription `json:"permission_levels,omitempty"`
}

func (st GetServingEndpointPermissionLevelsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := getServingEndpointPermissionLevelsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetServingEndpointPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getServingEndpointPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getServingEndpointPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetServingEndpointPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getServingEndpointPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetServingEndpointPermissionsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" tf:"-"`
}

func (st GetServingEndpointPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getServingEndpointPermissionsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetServingEndpointPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getServingEndpointPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getServingEndpointPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetServingEndpointPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getServingEndpointPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" tf:"-"`
}

func (st GetServingEndpointRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := getServingEndpointRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GetServingEndpointRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getServingEndpointRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getServingEndpointRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetServingEndpointRequest) MarshalJSON() ([]byte, error) {
	pb, err := getServingEndpointRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type GoogleCloudVertexAiConfig struct {
	// The Databricks secret key reference for a private key for the service
	// account which has access to the Google Cloud Vertex AI Service. See [Best
	// practices for managing service account keys]. If you prefer to paste your
	// API key directly, see `private_key_plaintext`. You must provide an API
	// key using one of the following fields: `private_key` or
	// `private_key_plaintext`
	//
	// [Best practices for managing service account keys]:
	// https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys
	// Wire name: 'private_key'
	PrivateKey string `json:"private_key,omitempty"`
	// The private key for the service account which has access to the Google
	// Cloud Vertex AI Service provided as a plaintext secret. See [Best
	// practices for managing service account keys]. If you prefer to reference
	// your key using Databricks Secrets, see `private_key`. You must provide an
	// API key using one of the following fields: `private_key` or
	// `private_key_plaintext`.
	//
	// [Best practices for managing service account keys]:
	// https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys
	// Wire name: 'private_key_plaintext'
	PrivateKeyPlaintext string `json:"private_key_plaintext,omitempty"`
	// This is the Google Cloud project id that the service account is
	// associated with.
	// Wire name: 'project_id'
	ProjectId string `json:"project_id"`
	// This is the region for the Google Cloud Vertex AI Service. See [supported
	// regions] for more details. Some models are only available in specific
	// regions.
	//
	// [supported regions]:
	// https://cloud.google.com/vertex-ai/docs/general/locations
	// Wire name: 'region'
	Region string `json:"region"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st GoogleCloudVertexAiConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := googleCloudVertexAiConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *GoogleCloudVertexAiConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &googleCloudVertexAiConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := googleCloudVertexAiConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GoogleCloudVertexAiConfig) MarshalJSON() ([]byte, error) {
	pb, err := googleCloudVertexAiConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type HttpRequestResponse struct {
	Contents io.ReadCloser `json:"-" tf:"-"`
}

func (st HttpRequestResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := httpRequestResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *HttpRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &httpRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := httpRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st HttpRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := httpRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListEndpointsResponse struct {
	// The list of endpoints.
	// Wire name: 'endpoints'
	Endpoints []ServingEndpoint `json:"endpoints,omitempty"`
}

func (st ListEndpointsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := listEndpointsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ListEndpointsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listEndpointsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listEndpointsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListEndpointsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listEndpointsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type LogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" tf:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName string `json:"-" tf:"-"`
}

func (st LogsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := logsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *LogsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogsRequest) MarshalJSON() ([]byte, error) {
	pb, err := logsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// A representation of all DataPlaneInfo for operations that can be done on a
// model through Data Plane APIs.
type ModelDataPlaneInfo struct {
	// Information required to query DataPlane API 'query' endpoint.
	// Wire name: 'query_info'
	QueryInfo *DataPlaneInfo `json:"query_info,omitempty"`
}

func (st ModelDataPlaneInfo) EncodeValues(key string, v *url.Values) error {
	pb, err := modelDataPlaneInfoToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ModelDataPlaneInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelDataPlaneInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelDataPlaneInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelDataPlaneInfo) MarshalJSON() ([]byte, error) {
	pb, err := modelDataPlaneInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Configs needed to create an OpenAI model route.
type OpenAiConfig struct {
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Client ID.
	// Wire name: 'microsoft_entra_client_id'
	MicrosoftEntraClientId string `json:"microsoft_entra_client_id,omitempty"`
	// The Databricks secret key reference for a client secret used for
	// Microsoft Entra ID authentication. If you prefer to paste your client
	// secret directly, see `microsoft_entra_client_secret_plaintext`. You must
	// provide an API key using one of the following fields:
	// `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	// Wire name: 'microsoft_entra_client_secret'
	MicrosoftEntraClientSecret string `json:"microsoft_entra_client_secret,omitempty"`
	// The client secret used for Microsoft Entra ID authentication provided as
	// a plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `microsoft_entra_client_secret`. You must provide an API key
	// using one of the following fields: `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	// Wire name: 'microsoft_entra_client_secret_plaintext'
	MicrosoftEntraClientSecretPlaintext string `json:"microsoft_entra_client_secret_plaintext,omitempty"`
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Tenant ID.
	// Wire name: 'microsoft_entra_tenant_id'
	MicrosoftEntraTenantId string `json:"microsoft_entra_tenant_id,omitempty"`
	// This is a field to provide a customized base URl for the OpenAI API. For
	// Azure OpenAI, this field is required, and is the base URL for the Azure
	// OpenAI API service provided by Azure. For other OpenAI API types, this
	// field is optional, and if left unspecified, the standard OpenAI base URL
	// is used.
	// Wire name: 'openai_api_base'
	OpenaiApiBase string `json:"openai_api_base,omitempty"`
	// The Databricks secret key reference for an OpenAI API key using the
	// OpenAI or Azure service. If you prefer to paste your API key directly,
	// see `openai_api_key_plaintext`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	// Wire name: 'openai_api_key'
	OpenaiApiKey string `json:"openai_api_key,omitempty"`
	// The OpenAI API key using the OpenAI or Azure service provided as a
	// plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `openai_api_key`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	// Wire name: 'openai_api_key_plaintext'
	OpenaiApiKeyPlaintext string `json:"openai_api_key_plaintext,omitempty"`
	// This is an optional field to specify the type of OpenAI API to use. For
	// Azure OpenAI, this field is required, and adjust this parameter to
	// represent the preferred security access validation protocol. For access
	// token validation, use azure. For authentication using Azure Active
	// Directory (Azure AD) use, azuread.
	// Wire name: 'openai_api_type'
	OpenaiApiType string `json:"openai_api_type,omitempty"`
	// This is an optional field to specify the OpenAI API version. For Azure
	// OpenAI, this field is required, and is the version of the Azure OpenAI
	// service to utilize, specified by a date.
	// Wire name: 'openai_api_version'
	OpenaiApiVersion string `json:"openai_api_version,omitempty"`
	// This field is only required for Azure OpenAI and is the name of the
	// deployment resource for the Azure OpenAI service.
	// Wire name: 'openai_deployment_name'
	OpenaiDeploymentName string `json:"openai_deployment_name,omitempty"`
	// This is an optional field to specify the organization in OpenAI or Azure
	// OpenAI.
	// Wire name: 'openai_organization'
	OpenaiOrganization string `json:"openai_organization,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st OpenAiConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := openAiConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *OpenAiConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &openAiConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := openAiConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st OpenAiConfig) MarshalJSON() ([]byte, error) {
	pb, err := openAiConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PaLmConfig struct {
	// The Databricks secret key reference for a PaLM API key. If you prefer to
	// paste your API key directly, see `palm_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	// Wire name: 'palm_api_key'
	PalmApiKey string `json:"palm_api_key,omitempty"`
	// The PaLM API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `palm_api_key`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	// Wire name: 'palm_api_key_plaintext'
	PalmApiKeyPlaintext string `json:"palm_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PaLmConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := paLmConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PaLmConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &paLmConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := paLmConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PaLmConfig) MarshalJSON() ([]byte, error) {
	pb, err := paLmConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	// Wire name: 'add_tags'
	AddTags []EndpointTag `json:"add_tags,omitempty"`
	// List of tag keys to delete
	// Wire name: 'delete_tags'
	DeleteTags []string `json:"delete_tags,omitempty"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name string `json:"-" tf:"-"`
}

func (st PatchServingEndpointTags) EncodeValues(key string, v *url.Values) error {
	pb, err := patchServingEndpointTagsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PatchServingEndpointTags) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &patchServingEndpointTagsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := patchServingEndpointTagsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PatchServingEndpointTags) MarshalJSON() ([]byte, error) {
	pb, err := patchServingEndpointTagsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PayloadTable struct {

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	// Wire name: 'status'
	Status string `json:"status,omitempty"`

	// Wire name: 'status_message'
	StatusMessage string `json:"status_message,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PayloadTable) EncodeValues(key string, v *url.Values) error {
	pb, err := payloadTableToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PayloadTable) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &payloadTablePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := payloadTableFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PayloadTable) MarshalJSON() ([]byte, error) {
	pb, err := payloadTableToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PtEndpointCoreConfig struct {
	// The list of served entities under the serving endpoint config.
	// Wire name: 'served_entities'
	ServedEntities []PtServedModel `json:"served_entities,omitempty"`

	// Wire name: 'traffic_config'
	TrafficConfig *TrafficConfig `json:"traffic_config,omitempty"`
}

func (st PtEndpointCoreConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := ptEndpointCoreConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PtEndpointCoreConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ptEndpointCoreConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ptEndpointCoreConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PtEndpointCoreConfig) MarshalJSON() ([]byte, error) {
	pb, err := ptEndpointCoreConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PtServedModel struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	// Wire name: 'entity_name'
	EntityName string `json:"entity_name"`

	// Wire name: 'entity_version'
	EntityVersion string `json:"entity_version,omitempty"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The number of model units to be provisioned.
	// Wire name: 'provisioned_model_units'
	ProvisionedModelUnits int64 `json:"provisioned_model_units"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st PtServedModel) EncodeValues(key string, v *url.Values) error {
	pb, err := ptServedModelToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PtServedModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &ptServedModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ptServedModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PtServedModel) MarshalJSON() ([]byte, error) {
	pb, err := ptServedModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PutAiGatewayRequest struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	// Wire name: 'fallback_config'
	FallbackConfig *FallbackConfig `json:"fallback_config,omitempty"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	// Wire name: 'guardrails'
	Guardrails *AiGatewayGuardrails `json:"guardrails,omitempty"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	// Wire name: 'inference_table_config'
	InferenceTableConfig *AiGatewayInferenceTableConfig `json:"inference_table_config,omitempty"`
	// The name of the serving endpoint whose AI Gateway is being updated. This
	// field is required.
	Name string `json:"-" tf:"-"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	// Wire name: 'rate_limits'
	RateLimits []AiGatewayRateLimit `json:"rate_limits,omitempty"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	// Wire name: 'usage_tracking_config'
	UsageTrackingConfig *AiGatewayUsageTrackingConfig `json:"usage_tracking_config,omitempty"`
}

func (st PutAiGatewayRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := putAiGatewayRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PutAiGatewayRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &putAiGatewayRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := putAiGatewayRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PutAiGatewayRequest) MarshalJSON() ([]byte, error) {
	pb, err := putAiGatewayRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PutAiGatewayResponse struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	// Wire name: 'fallback_config'
	FallbackConfig *FallbackConfig `json:"fallback_config,omitempty"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	// Wire name: 'guardrails'
	Guardrails *AiGatewayGuardrails `json:"guardrails,omitempty"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	// Wire name: 'inference_table_config'
	InferenceTableConfig *AiGatewayInferenceTableConfig `json:"inference_table_config,omitempty"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	// Wire name: 'rate_limits'
	RateLimits []AiGatewayRateLimit `json:"rate_limits,omitempty"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	// Wire name: 'usage_tracking_config'
	UsageTrackingConfig *AiGatewayUsageTrackingConfig `json:"usage_tracking_config,omitempty"`
}

func (st PutAiGatewayResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := putAiGatewayResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PutAiGatewayResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &putAiGatewayResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := putAiGatewayResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PutAiGatewayResponse) MarshalJSON() ([]byte, error) {
	pb, err := putAiGatewayResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PutRequest struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name string `json:"-" tf:"-"`
	// The list of endpoint rate limits.
	// Wire name: 'rate_limits'
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
}

func (st PutRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := putRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PutRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &putRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := putRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PutRequest) MarshalJSON() ([]byte, error) {
	pb, err := putRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type PutResponse struct {
	// The list of endpoint rate limits.
	// Wire name: 'rate_limits'
	RateLimits []RateLimit `json:"rate_limits,omitempty"`
}

func (st PutResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := putResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *PutResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &putResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := putResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st PutResponse) MarshalJSON() ([]byte, error) {
	pb, err := putResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryEndpointInput struct {
	// Pandas Dataframe input in the records orientation.
	// Wire name: 'dataframe_records'
	DataframeRecords []any `json:"dataframe_records,omitempty"`
	// Pandas Dataframe input in the split orientation.
	// Wire name: 'dataframe_split'
	DataframeSplit *DataframeSplitInput `json:"dataframe_split,omitempty"`
	// The extra parameters field used ONLY for __completions, chat,__ and
	// __embeddings external & foundation model__ serving endpoints. This is a
	// map of strings and should only be used with other external/foundation
	// model query fields.
	// Wire name: 'extra_params'
	ExtraParams map[string]string `json:"extra_params,omitempty"`
	// The input string (or array of strings) field used ONLY for __embeddings
	// external & foundation model__ serving endpoints and is the only field
	// (along with extra_params if needed) used by embeddings queries.
	// Wire name: 'input'
	Input any `json:"input,omitempty"`
	// Tensor-based input in columnar format.
	// Wire name: 'inputs'
	Inputs any `json:"inputs,omitempty"`
	// Tensor-based input in row format.
	// Wire name: 'instances'
	Instances []any `json:"instances,omitempty"`
	// The max tokens field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is an integer and should only
	// be used with other chat/completions query fields.
	// Wire name: 'max_tokens'
	MaxTokens int `json:"max_tokens,omitempty"`
	// The messages field used ONLY for __chat external & foundation model__
	// serving endpoints. This is a map of strings and should only be used with
	// other chat query fields.
	// Wire name: 'messages'
	Messages []ChatMessage `json:"messages,omitempty"`
	// The n (number of candidates) field used ONLY for __completions__ and
	// __chat external & foundation model__ serving endpoints. This is an
	// integer between 1 and 5 with a default of 1 and should only be used with
	// other chat/completions query fields.
	// Wire name: 'n'
	N int `json:"n,omitempty"`
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" tf:"-"`
	// The prompt string (or array of strings) field used ONLY for __completions
	// external & foundation model__ serving endpoints and should only be used
	// with other completions query fields.
	// Wire name: 'prompt'
	Prompt any `json:"prompt,omitempty"`
	// The stop sequences field used ONLY for __completions__ and __chat
	// external & foundation model__ serving endpoints. This is a list of
	// strings and should only be used with other chat/completions query fields.
	// Wire name: 'stop'
	Stop []string `json:"stop,omitempty"`
	// The stream field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a boolean defaulting to
	// false and should only be used with other chat/completions query fields.
	// Wire name: 'stream'
	Stream bool `json:"stream,omitempty"`
	// The temperature field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a float between 0.0 and 2.0
	// with a default of 1.0 and should only be used with other chat/completions
	// query fields.
	// Wire name: 'temperature'
	Temperature float64 `json:"temperature,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QueryEndpointInput) EncodeValues(key string, v *url.Values) error {
	pb, err := queryEndpointInputToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryEndpointInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryEndpointInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryEndpointInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryEndpointInput) MarshalJSON() ([]byte, error) {
	pb, err := queryEndpointInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QueryEndpointResponse struct {
	// The list of choices returned by the __chat or completions
	// external/foundation model__ serving endpoint.
	// Wire name: 'choices'
	Choices []V1ResponseChoiceElement `json:"choices,omitempty"`
	// The timestamp in seconds when the query was created in Unix time returned
	// by a __completions or chat external/foundation model__ serving endpoint.
	// Wire name: 'created'
	Created int64 `json:"created,omitempty"`
	// The list of the embeddings returned by the __embeddings
	// external/foundation model__ serving endpoint.
	// Wire name: 'data'
	Data []EmbeddingsV1ResponseEmbeddingElement `json:"data,omitempty"`
	// The ID of the query that may be returned by a __completions or chat
	// external/foundation model__ serving endpoint.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The name of the __external/foundation model__ used for querying. This is
	// the name of the model that was specified in the endpoint config.
	// Wire name: 'model'
	Model string `json:"model,omitempty"`
	// The type of object returned by the __external/foundation model__ serving
	// endpoint, one of [text_completion, chat.completion, list (of
	// embeddings)].
	// Wire name: 'object'
	Object QueryEndpointResponseObject `json:"object,omitempty"`
	// The predictions returned by the serving endpoint.
	// Wire name: 'predictions'
	Predictions []any `json:"predictions,omitempty"`
	// The name of the served model that served the request. This is useful when
	// there are multiple models behind the same endpoint with traffic split.
	ServedModelName string `json:"-" tf:"-"`
	// The usage object that may be returned by the __external/foundation
	// model__ serving endpoint. This contains information about the number of
	// tokens used in the prompt and response.
	// Wire name: 'usage'
	Usage *ExternalModelUsageElement `json:"usage,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st QueryEndpointResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := queryEndpointResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *QueryEndpointResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &queryEndpointResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := queryEndpointResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QueryEndpointResponse) MarshalJSON() ([]byte, error) {
	pb, err := queryEndpointResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for QueryEndpointResponseObject.
//
// There is no guarantee on the order of the values in the slice.
func (f *QueryEndpointResponseObject) Values() []QueryEndpointResponseObject {
	return []QueryEndpointResponseObject{
		QueryEndpointResponseObjectChatCompletion,
		QueryEndpointResponseObjectList,
		QueryEndpointResponseObjectTextCompletion,
	}
}

// Type always returns QueryEndpointResponseObject to satisfy [pflag.Value] interface
func (f *QueryEndpointResponseObject) Type() string {
	return "QueryEndpointResponseObject"
}

type RateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	// Wire name: 'calls'
	Calls int64 `json:"calls"`
	// Key field for a serving endpoint rate limit. Currently, only 'user' and
	// 'endpoint' are supported, with 'endpoint' being the default if not
	// specified.
	// Wire name: 'key'
	Key RateLimitKey `json:"key,omitempty"`
	// Renewal period field for a serving endpoint rate limit. Currently, only
	// 'minute' is supported.
	// Wire name: 'renewal_period'
	RenewalPeriod RateLimitRenewalPeriod `json:"renewal_period"`
}

func (st RateLimit) EncodeValues(key string, v *url.Values) error {
	pb, err := rateLimitToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *RateLimit) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rateLimitPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rateLimitFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RateLimit) MarshalJSON() ([]byte, error) {
	pb, err := rateLimitToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

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

// Values returns all possible values for RateLimitKey.
//
// There is no guarantee on the order of the values in the slice.
func (f *RateLimitKey) Values() []RateLimitKey {
	return []RateLimitKey{
		RateLimitKeyEndpoint,
		RateLimitKeyUser,
	}
}

// Type always returns RateLimitKey to satisfy [pflag.Value] interface
func (f *RateLimitKey) Type() string {
	return "RateLimitKey"
}

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

// Values returns all possible values for RateLimitRenewalPeriod.
//
// There is no guarantee on the order of the values in the slice.
func (f *RateLimitRenewalPeriod) Values() []RateLimitRenewalPeriod {
	return []RateLimitRenewalPeriod{
		RateLimitRenewalPeriodMinute,
	}
}

// Type always returns RateLimitRenewalPeriod to satisfy [pflag.Value] interface
func (f *RateLimitRenewalPeriod) Type() string {
	return "RateLimitRenewalPeriod"
}

type Route struct {

	// Wire name: 'served_entity_name'
	ServedEntityName string `json:"served_entity_name,omitempty"`
	// The name of the served model this route configures traffic for.
	// Wire name: 'served_model_name'
	ServedModelName string `json:"served_model_name,omitempty"`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	// Wire name: 'traffic_percentage'
	TrafficPercentage int `json:"traffic_percentage"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st Route) EncodeValues(key string, v *url.Values) error {
	pb, err := routeToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *Route) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &routePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := routeFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Route) MarshalJSON() ([]byte, error) {
	pb, err := routeToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServedEntityInput struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	// Wire name: 'entity_name'
	EntityName string `json:"entity_name,omitempty"`

	// Wire name: 'entity_version'
	EntityVersion string `json:"entity_version,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	// Wire name: 'environment_vars'
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	// Wire name: 'external_model'
	ExternalModel *ExternalModel `json:"external_model,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	// Wire name: 'max_provisioned_concurrency'
	MaxProvisionedConcurrency int `json:"max_provisioned_concurrency,omitempty"`
	// The maximum tokens per second that the endpoint can scale up to.
	// Wire name: 'max_provisioned_throughput'
	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	// Wire name: 'min_provisioned_concurrency'
	MinProvisionedConcurrency int `json:"min_provisioned_concurrency,omitempty"`
	// The minimum tokens per second that the endpoint can scale down to.
	// Wire name: 'min_provisioned_throughput'
	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The number of model units provisioned.
	// Wire name: 'provisioned_model_units'
	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	// Wire name: 'scale_to_zero_enabled'
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	// Do not use if min_provisioned_concurrency and max_provisioned_concurrency
	// are specified.
	// Wire name: 'workload_size'
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	// Wire name: 'workload_type'
	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServedEntityInput) EncodeValues(key string, v *url.Values) error {
	pb, err := servedEntityInputToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServedEntityInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servedEntityInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servedEntityInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServedEntityInput) MarshalJSON() ([]byte, error) {
	pb, err := servedEntityInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServedEntityOutput struct {

	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	// Wire name: 'entity_name'
	EntityName string `json:"entity_name,omitempty"`

	// Wire name: 'entity_version'
	EntityVersion string `json:"entity_version,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	// Wire name: 'environment_vars'
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	// Wire name: 'external_model'
	ExternalModel *ExternalModel `json:"external_model,omitempty"`

	// Wire name: 'foundation_model'
	FoundationModel *FoundationModel `json:"foundation_model,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	// Wire name: 'max_provisioned_concurrency'
	MaxProvisionedConcurrency int `json:"max_provisioned_concurrency,omitempty"`
	// The maximum tokens per second that the endpoint can scale up to.
	// Wire name: 'max_provisioned_throughput'
	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	// Wire name: 'min_provisioned_concurrency'
	MinProvisionedConcurrency int `json:"min_provisioned_concurrency,omitempty"`
	// The minimum tokens per second that the endpoint can scale down to.
	// Wire name: 'min_provisioned_throughput'
	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The number of model units provisioned.
	// Wire name: 'provisioned_model_units'
	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	// Wire name: 'scale_to_zero_enabled'
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`

	// Wire name: 'state'
	State *ServedModelState `json:"state,omitempty"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	// Do not use if min_provisioned_concurrency and max_provisioned_concurrency
	// are specified.
	// Wire name: 'workload_size'
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	// Wire name: 'workload_type'
	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServedEntityOutput) EncodeValues(key string, v *url.Values) error {
	pb, err := servedEntityOutputToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServedEntityOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servedEntityOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servedEntityOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServedEntityOutput) MarshalJSON() ([]byte, error) {
	pb, err := servedEntityOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServedEntitySpec struct {

	// Wire name: 'entity_name'
	EntityName string `json:"entity_name,omitempty"`

	// Wire name: 'entity_version'
	EntityVersion string `json:"entity_version,omitempty"`

	// Wire name: 'external_model'
	ExternalModel *ExternalModel `json:"external_model,omitempty"`

	// Wire name: 'foundation_model'
	FoundationModel *FoundationModel `json:"foundation_model,omitempty"`

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServedEntitySpec) EncodeValues(key string, v *url.Values) error {
	pb, err := servedEntitySpecToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServedEntitySpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servedEntitySpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servedEntitySpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServedEntitySpec) MarshalJSON() ([]byte, error) {
	pb, err := servedEntitySpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServedModelInput struct {
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	// Wire name: 'environment_vars'
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	// Wire name: 'max_provisioned_concurrency'
	MaxProvisionedConcurrency int `json:"max_provisioned_concurrency,omitempty"`
	// The maximum tokens per second that the endpoint can scale up to.
	// Wire name: 'max_provisioned_throughput'
	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	// Wire name: 'min_provisioned_concurrency'
	MinProvisionedConcurrency int `json:"min_provisioned_concurrency,omitempty"`
	// The minimum tokens per second that the endpoint can scale down to.
	// Wire name: 'min_provisioned_throughput'
	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`

	// Wire name: 'model_name'
	ModelName string `json:"model_name"`

	// Wire name: 'model_version'
	ModelVersion string `json:"model_version"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The number of model units provisioned.
	// Wire name: 'provisioned_model_units'
	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	// Wire name: 'scale_to_zero_enabled'
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	// Do not use if min_provisioned_concurrency and max_provisioned_concurrency
	// are specified.
	// Wire name: 'workload_size'
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	// Wire name: 'workload_type'
	WorkloadType ServedModelInputWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServedModelInput) EncodeValues(key string, v *url.Values) error {
	pb, err := servedModelInputToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServedModelInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servedModelInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servedModelInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServedModelInput) MarshalJSON() ([]byte, error) {
	pb, err := servedModelInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Please keep this in sync with with workload types in
// InferenceEndpointEntities.scala
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

// Values returns all possible values for ServedModelInputWorkloadType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServedModelInputWorkloadType) Values() []ServedModelInputWorkloadType {
	return []ServedModelInputWorkloadType{
		ServedModelInputWorkloadTypeCpu,
		ServedModelInputWorkloadTypeGpuLarge,
		ServedModelInputWorkloadTypeGpuMedium,
		ServedModelInputWorkloadTypeGpuSmall,
		ServedModelInputWorkloadTypeMultigpuMedium,
	}
}

// Type always returns ServedModelInputWorkloadType to satisfy [pflag.Value] interface
func (f *ServedModelInputWorkloadType) Type() string {
	return "ServedModelInputWorkloadType"
}

type ServedModelOutput struct {

	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	// Wire name: 'environment_vars'
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	// Wire name: 'max_provisioned_concurrency'
	MaxProvisionedConcurrency int `json:"max_provisioned_concurrency,omitempty"`
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	// Wire name: 'min_provisioned_concurrency'
	MinProvisionedConcurrency int `json:"min_provisioned_concurrency,omitempty"`

	// Wire name: 'model_name'
	ModelName string `json:"model_name,omitempty"`

	// Wire name: 'model_version'
	ModelVersion string `json:"model_version,omitempty"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The number of model units provisioned.
	// Wire name: 'provisioned_model_units'
	ProvisionedModelUnits int64 `json:"provisioned_model_units,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	// Wire name: 'scale_to_zero_enabled'
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`

	// Wire name: 'state'
	State *ServedModelState `json:"state,omitempty"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	// Do not use if min_provisioned_concurrency and max_provisioned_concurrency
	// are specified.
	// Wire name: 'workload_size'
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	// Wire name: 'workload_type'
	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServedModelOutput) EncodeValues(key string, v *url.Values) error {
	pb, err := servedModelOutputToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServedModelOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servedModelOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servedModelOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServedModelOutput) MarshalJSON() ([]byte, error) {
	pb, err := servedModelOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServedModelSpec struct {
	// Only one of model_name and entity_name should be populated
	// Wire name: 'model_name'
	ModelName string `json:"model_name,omitempty"`
	// Only one of model_version and entity_version should be populated
	// Wire name: 'model_version'
	ModelVersion string `json:"model_version,omitempty"`

	// Wire name: 'name'
	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServedModelSpec) EncodeValues(key string, v *url.Values) error {
	pb, err := servedModelSpecToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServedModelSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servedModelSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servedModelSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServedModelSpec) MarshalJSON() ([]byte, error) {
	pb, err := servedModelSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServedModelState struct {

	// Wire name: 'deployment'
	Deployment ServedModelStateDeployment `json:"deployment,omitempty"`

	// Wire name: 'deployment_state_message'
	DeploymentStateMessage string `json:"deployment_state_message,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServedModelState) EncodeValues(key string, v *url.Values) error {
	pb, err := servedModelStateToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServedModelState) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servedModelStatePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servedModelStateFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServedModelState) MarshalJSON() ([]byte, error) {
	pb, err := servedModelStateToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

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

// Values returns all possible values for ServedModelStateDeployment.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServedModelStateDeployment) Values() []ServedModelStateDeployment {
	return []ServedModelStateDeployment{
		ServedModelStateDeploymentAborted,
		ServedModelStateDeploymentCreating,
		ServedModelStateDeploymentFailed,
		ServedModelStateDeploymentReady,
		ServedModelStateDeploymentRecovering,
	}
}

// Type always returns ServedModelStateDeployment to satisfy [pflag.Value] interface
func (f *ServedModelStateDeployment) Type() string {
	return "ServedModelStateDeployment"
}

type ServerLogsResponse struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	// Wire name: 'logs'
	Logs string `json:"logs"`
}

func (st ServerLogsResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := serverLogsResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServerLogsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &serverLogsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := serverLogsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServerLogsResponse) MarshalJSON() ([]byte, error) {
	pb, err := serverLogsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	// Wire name: 'ai_gateway'
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`
	// The budget policy associated with the endpoint.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The config that is currently being served by the endpoint.
	// Wire name: 'config'
	Config *EndpointCoreConfigSummary `json:"config,omitempty"`
	// The timestamp when the endpoint was created in Unix time.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the serving endpoint.
	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// Description of the endpoint
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// System-generated ID of the endpoint, included to be used by the
	// Permissions API.
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The name of the serving endpoint.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// Information corresponding to the state of the serving endpoint.
	// Wire name: 'state'
	State *EndpointState `json:"state,omitempty"`
	// Tags attached to the serving endpoint.
	// Wire name: 'tags'
	Tags []EndpointTag `json:"tags,omitempty"`
	// The task type of the serving endpoint.
	// Wire name: 'task'
	Task string `json:"task,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServingEndpoint) EncodeValues(key string, v *url.Values) error {
	pb, err := servingEndpointToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServingEndpoint) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servingEndpointPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servingEndpointFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServingEndpoint) MarshalJSON() ([]byte, error) {
	pb, err := servingEndpointToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServingEndpointAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServingEndpointAccessControlRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := servingEndpointAccessControlRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServingEndpointAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servingEndpointAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servingEndpointAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServingEndpointAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := servingEndpointAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServingEndpointAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []ServingEndpointPermission `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	// Wire name: 'group_name'
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	// Wire name: 'user_name'
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServingEndpointAccessControlResponse) EncodeValues(key string, v *url.Values) error {
	pb, err := servingEndpointAccessControlResponseToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServingEndpointAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servingEndpointAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servingEndpointAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServingEndpointAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := servingEndpointAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServingEndpointDetailed struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	// Wire name: 'ai_gateway'
	AiGateway *AiGatewayConfig `json:"ai_gateway,omitempty"`
	// The budget policy associated with the endpoint.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The config that is currently being served by the endpoint.
	// Wire name: 'config'
	Config *EndpointCoreConfigOutput `json:"config,omitempty"`
	// The timestamp when the endpoint was created in Unix time.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the serving endpoint.
	// Wire name: 'creator'
	Creator string `json:"creator,omitempty"`
	// Information required to query DataPlane APIs.
	// Wire name: 'data_plane_info'
	DataPlaneInfo *ModelDataPlaneInfo `json:"data_plane_info,omitempty"`
	// Description of the serving model
	// Wire name: 'description'
	Description string `json:"description,omitempty"`
	// Endpoint invocation url if route optimization is enabled for endpoint
	// Wire name: 'endpoint_url'
	EndpointUrl string `json:"endpoint_url,omitempty"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	// Wire name: 'id'
	Id string `json:"id,omitempty"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The name of the serving endpoint.
	// Wire name: 'name'
	Name string `json:"name,omitempty"`
	// The config that the endpoint is attempting to update to.
	// Wire name: 'pending_config'
	PendingConfig *EndpointPendingConfig `json:"pending_config,omitempty"`
	// The permission level of the principal making the request.
	// Wire name: 'permission_level'
	PermissionLevel ServingEndpointDetailedPermissionLevel `json:"permission_level,omitempty"`
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	// Wire name: 'route_optimized'
	RouteOptimized bool `json:"route_optimized,omitempty"`
	// Information corresponding to the state of the serving endpoint.
	// Wire name: 'state'
	State *EndpointState `json:"state,omitempty"`
	// Tags attached to the serving endpoint.
	// Wire name: 'tags'
	Tags []EndpointTag `json:"tags,omitempty"`
	// The task type of the serving endpoint.
	// Wire name: 'task'
	Task string `json:"task,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServingEndpointDetailed) EncodeValues(key string, v *url.Values) error {
	pb, err := servingEndpointDetailedToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServingEndpointDetailed) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servingEndpointDetailedPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servingEndpointDetailedFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServingEndpointDetailed) MarshalJSON() ([]byte, error) {
	pb, err := servingEndpointDetailedToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

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

// Values returns all possible values for ServingEndpointDetailedPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServingEndpointDetailedPermissionLevel) Values() []ServingEndpointDetailedPermissionLevel {
	return []ServingEndpointDetailedPermissionLevel{
		ServingEndpointDetailedPermissionLevelCanManage,
		ServingEndpointDetailedPermissionLevelCanQuery,
		ServingEndpointDetailedPermissionLevelCanView,
	}
}

// Type always returns ServingEndpointDetailedPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointDetailedPermissionLevel) Type() string {
	return "ServingEndpointDetailedPermissionLevel"
}

type ServingEndpointPermission struct {

	// Wire name: 'inherited'
	Inherited bool `json:"inherited,omitempty"`

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServingEndpointPermission) EncodeValues(key string, v *url.Values) error {
	pb, err := servingEndpointPermissionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServingEndpointPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servingEndpointPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servingEndpointPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServingEndpointPermission) MarshalJSON() ([]byte, error) {
	pb, err := servingEndpointPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Values returns all possible values for ServingEndpointPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServingEndpointPermissionLevel) Values() []ServingEndpointPermissionLevel {
	return []ServingEndpointPermissionLevel{
		ServingEndpointPermissionLevelCanManage,
		ServingEndpointPermissionLevelCanQuery,
		ServingEndpointPermissionLevelCanView,
	}
}

// Type always returns ServingEndpointPermissionLevel to satisfy [pflag.Value] interface
func (f *ServingEndpointPermissionLevel) Type() string {
	return "ServingEndpointPermissionLevel"
}

type ServingEndpointPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []ServingEndpointAccessControlResponse `json:"access_control_list,omitempty"`

	// Wire name: 'object_id'
	ObjectId string `json:"object_id,omitempty"`

	// Wire name: 'object_type'
	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServingEndpointPermissions) EncodeValues(key string, v *url.Values) error {
	pb, err := servingEndpointPermissionsToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServingEndpointPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servingEndpointPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servingEndpointPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServingEndpointPermissions) MarshalJSON() ([]byte, error) {
	pb, err := servingEndpointPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServingEndpointPermissionsDescription struct {

	// Wire name: 'description'
	Description string `json:"description,omitempty"`

	// Wire name: 'permission_level'
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st ServingEndpointPermissionsDescription) EncodeValues(key string, v *url.Values) error {
	pb, err := servingEndpointPermissionsDescriptionToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServingEndpointPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servingEndpointPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servingEndpointPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServingEndpointPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := servingEndpointPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ServingEndpointPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []ServingEndpointAccessControlRequest `json:"access_control_list,omitempty"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" tf:"-"`
}

func (st ServingEndpointPermissionsRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := servingEndpointPermissionsRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *ServingEndpointPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &servingEndpointPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := servingEndpointPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ServingEndpointPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := servingEndpointPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Please keep this in sync with with workload types in
// InferenceEndpointEntities.scala
type ServingModelWorkloadType string

const ServingModelWorkloadTypeCpu ServingModelWorkloadType = `CPU`

const ServingModelWorkloadTypeGpuLarge ServingModelWorkloadType = `GPU_LARGE`

const ServingModelWorkloadTypeGpuMedium ServingModelWorkloadType = `GPU_MEDIUM`

const ServingModelWorkloadTypeGpuSmall ServingModelWorkloadType = `GPU_SMALL`

const ServingModelWorkloadTypeMultigpuMedium ServingModelWorkloadType = `MULTIGPU_MEDIUM`

// String representation for [fmt.Print]
func (f *ServingModelWorkloadType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ServingModelWorkloadType) Set(v string) error {
	switch v {
	case `CPU`, `GPU_LARGE`, `GPU_MEDIUM`, `GPU_SMALL`, `MULTIGPU_MEDIUM`:
		*f = ServingModelWorkloadType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CPU", "GPU_LARGE", "GPU_MEDIUM", "GPU_SMALL", "MULTIGPU_MEDIUM"`, v)
	}
}

// Values returns all possible values for ServingModelWorkloadType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ServingModelWorkloadType) Values() []ServingModelWorkloadType {
	return []ServingModelWorkloadType{
		ServingModelWorkloadTypeCpu,
		ServingModelWorkloadTypeGpuLarge,
		ServingModelWorkloadTypeGpuMedium,
		ServingModelWorkloadTypeGpuSmall,
		ServingModelWorkloadTypeMultigpuMedium,
	}
}

// Type always returns ServingModelWorkloadType to satisfy [pflag.Value] interface
func (f *ServingModelWorkloadType) Type() string {
	return "ServingModelWorkloadType"
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served entity.
	// Wire name: 'routes'
	Routes []Route `json:"routes,omitempty"`
}

func (st TrafficConfig) EncodeValues(key string, v *url.Values) error {
	pb, err := trafficConfigToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *TrafficConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &trafficConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := trafficConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TrafficConfig) MarshalJSON() ([]byte, error) {
	pb, err := trafficConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type UpdateProvisionedThroughputEndpointConfigRequest struct {

	// Wire name: 'config'
	Config PtEndpointCoreConfig `json:"config"`
	// The name of the pt endpoint to update. This field is required.
	Name string `json:"-" tf:"-"`
}

func (st UpdateProvisionedThroughputEndpointConfigRequest) EncodeValues(key string, v *url.Values) error {
	pb, err := updateProvisionedThroughputEndpointConfigRequestToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *UpdateProvisionedThroughputEndpointConfigRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateProvisionedThroughputEndpointConfigRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateProvisionedThroughputEndpointConfigRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateProvisionedThroughputEndpointConfigRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateProvisionedThroughputEndpointConfigRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type V1ResponseChoiceElement struct {
	// The finish reason returned by the endpoint.
	// Wire name: 'finishReason'
	FinishReason string `json:"finishReason,omitempty"`
	// The index of the choice in the __chat or completions__ response.
	// Wire name: 'index'
	Index int `json:"index,omitempty"`
	// The logprobs returned only by the __completions__ endpoint.
	// Wire name: 'logprobs'
	Logprobs int `json:"logprobs,omitempty"`
	// The message response from the __chat__ endpoint.
	// Wire name: 'message'
	Message *ChatMessage `json:"message,omitempty"`
	// The text response from the __completions__ endpoint.
	// Wire name: 'text'
	Text string `json:"text,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st V1ResponseChoiceElement) EncodeValues(key string, v *url.Values) error {
	pb, err := v1ResponseChoiceElementToPb(&st)
	if err != nil {
		return err
	}
	vals, err := query.Values(pb)
	if err != nil {
		return err
	}
	for k, vs := range vals {
		for _, val := range vs {
			v.Add(fmt.Sprintf("%s[%s]", key, k), val)
		}
	}
	return nil
}

func (st *V1ResponseChoiceElement) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &v1ResponseChoiceElementPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := v1ResponseChoiceElementFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st V1ResponseChoiceElement) MarshalJSON() ([]byte, error) {
	pb, err := v1ResponseChoiceElementToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
