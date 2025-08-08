// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/serving/servingpb"
)

type Ai21LabsConfig struct {
	// The Databricks secret key reference for an AI21 Labs API key. If you
	// prefer to paste your API key directly, see `ai21labs_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	// Wire name: 'ai21labs_api_key'
	Ai21labsApiKey string ``
	// An AI21 Labs API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `ai21labs_api_key`. You
	// must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	// Wire name: 'ai21labs_api_key_plaintext'
	Ai21labsApiKeyPlaintext string   ``
	ForceSendFields         []string `tf:"-"`
}

func (s *Ai21LabsConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Ai21LabsConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func Ai21LabsConfigToPb(st *Ai21LabsConfig) (*servingpb.Ai21LabsConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.Ai21LabsConfigPb{}
	pb.Ai21labsApiKey = st.Ai21labsApiKey
	pb.Ai21labsApiKeyPlaintext = st.Ai21labsApiKeyPlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func Ai21LabsConfigFromPb(pb *servingpb.Ai21LabsConfigPb) (*Ai21LabsConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Ai21LabsConfig{}
	st.Ai21labsApiKey = pb.Ai21labsApiKey
	st.Ai21labsApiKeyPlaintext = pb.Ai21labsApiKeyPlaintext

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AiGatewayConfig struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	// Wire name: 'fallback_config'
	FallbackConfig *FallbackConfig ``
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	// Wire name: 'guardrails'
	Guardrails *AiGatewayGuardrails ``
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	// Wire name: 'inference_table_config'
	InferenceTableConfig *AiGatewayInferenceTableConfig ``
	// Configuration for rate limits which can be set to limit endpoint traffic.
	// Wire name: 'rate_limits'
	RateLimits []AiGatewayRateLimit ``
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	// Wire name: 'usage_tracking_config'
	UsageTrackingConfig *AiGatewayUsageTrackingConfig ``
}

func AiGatewayConfigToPb(st *AiGatewayConfig) (*servingpb.AiGatewayConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AiGatewayConfigPb{}
	fallbackConfigPb, err := FallbackConfigToPb(st.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigPb != nil {
		pb.FallbackConfig = fallbackConfigPb
	}
	guardrailsPb, err := AiGatewayGuardrailsToPb(st.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsPb != nil {
		pb.Guardrails = guardrailsPb
	}
	inferenceTableConfigPb, err := AiGatewayInferenceTableConfigToPb(st.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigPb != nil {
		pb.InferenceTableConfig = inferenceTableConfigPb
	}

	var rateLimitsPb []servingpb.AiGatewayRateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := AiGatewayRateLimitToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rateLimitsPb = append(rateLimitsPb, *itemPb)
		}
	}
	pb.RateLimits = rateLimitsPb
	usageTrackingConfigPb, err := AiGatewayUsageTrackingConfigToPb(st.UsageTrackingConfig)
	if err != nil {
		return nil, err
	}
	if usageTrackingConfigPb != nil {
		pb.UsageTrackingConfig = usageTrackingConfigPb
	}

	return pb, nil
}

func AiGatewayConfigFromPb(pb *servingpb.AiGatewayConfigPb) (*AiGatewayConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayConfig{}
	fallbackConfigField, err := FallbackConfigFromPb(pb.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigField != nil {
		st.FallbackConfig = fallbackConfigField
	}
	guardrailsField, err := AiGatewayGuardrailsFromPb(pb.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsField != nil {
		st.Guardrails = guardrailsField
	}
	inferenceTableConfigField, err := AiGatewayInferenceTableConfigFromPb(pb.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigField != nil {
		st.InferenceTableConfig = inferenceTableConfigField
	}

	var rateLimitsField []AiGatewayRateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := AiGatewayRateLimitFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rateLimitsField = append(rateLimitsField, *item)
		}
	}
	st.RateLimits = rateLimitsField
	usageTrackingConfigField, err := AiGatewayUsageTrackingConfigFromPb(pb.UsageTrackingConfig)
	if err != nil {
		return nil, err
	}
	if usageTrackingConfigField != nil {
		st.UsageTrackingConfig = usageTrackingConfigField
	}

	return st, nil
}

type AiGatewayGuardrailParameters struct {
	// List of invalid keywords. AI guardrail uses keyword or string matching to
	// decide if the keyword exists in the request or response content.
	// Wire name: 'invalid_keywords'
	InvalidKeywords []string ``
	// Configuration for guardrail PII filter.
	// Wire name: 'pii'
	Pii *AiGatewayGuardrailPiiBehavior ``
	// Indicates whether the safety filter is enabled.
	// Wire name: 'safety'
	Safety bool ``
	// The list of allowed topics. Given a chat request, this guardrail flags
	// the request if its topic is not in the allowed topics.
	// Wire name: 'valid_topics'
	ValidTopics     []string ``
	ForceSendFields []string `tf:"-"`
}

func (s *AiGatewayGuardrailParameters) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiGatewayGuardrailParameters) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AiGatewayGuardrailParametersToPb(st *AiGatewayGuardrailParameters) (*servingpb.AiGatewayGuardrailParametersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AiGatewayGuardrailParametersPb{}
	pb.InvalidKeywords = st.InvalidKeywords
	piiPb, err := AiGatewayGuardrailPiiBehaviorToPb(st.Pii)
	if err != nil {
		return nil, err
	}
	if piiPb != nil {
		pb.Pii = piiPb
	}
	pb.Safety = st.Safety
	pb.ValidTopics = st.ValidTopics

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AiGatewayGuardrailParametersFromPb(pb *servingpb.AiGatewayGuardrailParametersPb) (*AiGatewayGuardrailParameters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayGuardrailParameters{}
	st.InvalidKeywords = pb.InvalidKeywords
	piiField, err := AiGatewayGuardrailPiiBehaviorFromPb(pb.Pii)
	if err != nil {
		return nil, err
	}
	if piiField != nil {
		st.Pii = piiField
	}
	st.Safety = pb.Safety
	st.ValidTopics = pb.ValidTopics

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AiGatewayGuardrailPiiBehavior struct {
	// Configuration for input guardrail filters.
	// Wire name: 'behavior'
	Behavior AiGatewayGuardrailPiiBehaviorBehavior ``
}

func AiGatewayGuardrailPiiBehaviorToPb(st *AiGatewayGuardrailPiiBehavior) (*servingpb.AiGatewayGuardrailPiiBehaviorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AiGatewayGuardrailPiiBehaviorPb{}
	behaviorPb, err := AiGatewayGuardrailPiiBehaviorBehaviorToPb(&st.Behavior)
	if err != nil {
		return nil, err
	}
	if behaviorPb != nil {
		pb.Behavior = *behaviorPb
	}

	return pb, nil
}

func AiGatewayGuardrailPiiBehaviorFromPb(pb *servingpb.AiGatewayGuardrailPiiBehaviorPb) (*AiGatewayGuardrailPiiBehavior, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayGuardrailPiiBehavior{}
	behaviorField, err := AiGatewayGuardrailPiiBehaviorBehaviorFromPb(&pb.Behavior)
	if err != nil {
		return nil, err
	}
	if behaviorField != nil {
		st.Behavior = *behaviorField
	}

	return st, nil
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

func AiGatewayGuardrailPiiBehaviorBehaviorToPb(st *AiGatewayGuardrailPiiBehaviorBehavior) (*servingpb.AiGatewayGuardrailPiiBehaviorBehaviorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.AiGatewayGuardrailPiiBehaviorBehaviorPb(*st)
	return &pb, nil
}

func AiGatewayGuardrailPiiBehaviorBehaviorFromPb(pb *servingpb.AiGatewayGuardrailPiiBehaviorBehaviorPb) (*AiGatewayGuardrailPiiBehaviorBehavior, error) {
	if pb == nil {
		return nil, nil
	}
	st := AiGatewayGuardrailPiiBehaviorBehavior(*pb)
	return &st, nil
}

type AiGatewayGuardrails struct {
	// Configuration for input guardrail filters.
	// Wire name: 'input'
	Input *AiGatewayGuardrailParameters ``
	// Configuration for output guardrail filters.
	// Wire name: 'output'
	Output *AiGatewayGuardrailParameters ``
}

func AiGatewayGuardrailsToPb(st *AiGatewayGuardrails) (*servingpb.AiGatewayGuardrailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AiGatewayGuardrailsPb{}
	inputPb, err := AiGatewayGuardrailParametersToPb(st.Input)
	if err != nil {
		return nil, err
	}
	if inputPb != nil {
		pb.Input = inputPb
	}
	outputPb, err := AiGatewayGuardrailParametersToPb(st.Output)
	if err != nil {
		return nil, err
	}
	if outputPb != nil {
		pb.Output = outputPb
	}

	return pb, nil
}

func AiGatewayGuardrailsFromPb(pb *servingpb.AiGatewayGuardrailsPb) (*AiGatewayGuardrails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayGuardrails{}
	inputField, err := AiGatewayGuardrailParametersFromPb(pb.Input)
	if err != nil {
		return nil, err
	}
	if inputField != nil {
		st.Input = inputField
	}
	outputField, err := AiGatewayGuardrailParametersFromPb(pb.Output)
	if err != nil {
		return nil, err
	}
	if outputField != nil {
		st.Output = outputField
	}

	return st, nil
}

type AiGatewayInferenceTableConfig struct {
	// The name of the catalog in Unity Catalog. Required when enabling
	// inference tables. NOTE: On update, you have to disable inference table
	// first in order to change the catalog name.
	// Wire name: 'catalog_name'
	CatalogName string ``
	// Indicates whether the inference table is enabled.
	// Wire name: 'enabled'
	Enabled bool ``
	// The name of the schema in Unity Catalog. Required when enabling inference
	// tables. NOTE: On update, you have to disable inference table first in
	// order to change the schema name.
	// Wire name: 'schema_name'
	SchemaName string ``
	// The prefix of the table in Unity Catalog. NOTE: On update, you have to
	// disable inference table first in order to change the prefix name.
	// Wire name: 'table_name_prefix'
	TableNamePrefix string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AiGatewayInferenceTableConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiGatewayInferenceTableConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AiGatewayInferenceTableConfigToPb(st *AiGatewayInferenceTableConfig) (*servingpb.AiGatewayInferenceTableConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AiGatewayInferenceTableConfigPb{}
	pb.CatalogName = st.CatalogName
	pb.Enabled = st.Enabled
	pb.SchemaName = st.SchemaName
	pb.TableNamePrefix = st.TableNamePrefix

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AiGatewayInferenceTableConfigFromPb(pb *servingpb.AiGatewayInferenceTableConfigPb) (*AiGatewayInferenceTableConfig, error) {
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

type AiGatewayRateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	// Wire name: 'calls'
	Calls int64 ``
	// Key field for a rate limit. Currently, 'user', 'user_group,
	// 'service_principal', and 'endpoint' are supported, with 'endpoint' being
	// the default if not specified.
	// Wire name: 'key'
	Key AiGatewayRateLimitKey ``
	// Principal field for a user, user group, or service principal to apply
	// rate limiting to. Accepts a user email, group name, or service principal
	// application ID.
	// Wire name: 'principal'
	Principal string ``
	// Renewal period field for a rate limit. Currently, only 'minute' is
	// supported.
	// Wire name: 'renewal_period'
	RenewalPeriod   AiGatewayRateLimitRenewalPeriod ``
	ForceSendFields []string                        `tf:"-"`
}

func (s *AiGatewayRateLimit) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiGatewayRateLimit) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AiGatewayRateLimitToPb(st *AiGatewayRateLimit) (*servingpb.AiGatewayRateLimitPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AiGatewayRateLimitPb{}
	pb.Calls = st.Calls
	keyPb, err := AiGatewayRateLimitKeyToPb(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}
	pb.Principal = st.Principal
	renewalPeriodPb, err := AiGatewayRateLimitRenewalPeriodToPb(&st.RenewalPeriod)
	if err != nil {
		return nil, err
	}
	if renewalPeriodPb != nil {
		pb.RenewalPeriod = *renewalPeriodPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AiGatewayRateLimitFromPb(pb *servingpb.AiGatewayRateLimitPb) (*AiGatewayRateLimit, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayRateLimit{}
	st.Calls = pb.Calls
	keyField, err := AiGatewayRateLimitKeyFromPb(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
	st.Principal = pb.Principal
	renewalPeriodField, err := AiGatewayRateLimitRenewalPeriodFromPb(&pb.RenewalPeriod)
	if err != nil {
		return nil, err
	}
	if renewalPeriodField != nil {
		st.RenewalPeriod = *renewalPeriodField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func AiGatewayRateLimitKeyToPb(st *AiGatewayRateLimitKey) (*servingpb.AiGatewayRateLimitKeyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.AiGatewayRateLimitKeyPb(*st)
	return &pb, nil
}

func AiGatewayRateLimitKeyFromPb(pb *servingpb.AiGatewayRateLimitKeyPb) (*AiGatewayRateLimitKey, error) {
	if pb == nil {
		return nil, nil
	}
	st := AiGatewayRateLimitKey(*pb)
	return &st, nil
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

func AiGatewayRateLimitRenewalPeriodToPb(st *AiGatewayRateLimitRenewalPeriod) (*servingpb.AiGatewayRateLimitRenewalPeriodPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.AiGatewayRateLimitRenewalPeriodPb(*st)
	return &pb, nil
}

func AiGatewayRateLimitRenewalPeriodFromPb(pb *servingpb.AiGatewayRateLimitRenewalPeriodPb) (*AiGatewayRateLimitRenewalPeriod, error) {
	if pb == nil {
		return nil, nil
	}
	st := AiGatewayRateLimitRenewalPeriod(*pb)
	return &st, nil
}

type AiGatewayUsageTrackingConfig struct {
	// Whether to enable usage tracking.
	// Wire name: 'enabled'
	Enabled         bool     ``
	ForceSendFields []string `tf:"-"`
}

func (s *AiGatewayUsageTrackingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AiGatewayUsageTrackingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AiGatewayUsageTrackingConfigToPb(st *AiGatewayUsageTrackingConfig) (*servingpb.AiGatewayUsageTrackingConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AiGatewayUsageTrackingConfigPb{}
	pb.Enabled = st.Enabled

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AiGatewayUsageTrackingConfigFromPb(pb *servingpb.AiGatewayUsageTrackingConfigPb) (*AiGatewayUsageTrackingConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayUsageTrackingConfig{}
	st.Enabled = pb.Enabled

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AmazonBedrockConfig struct {
	// The Databricks secret key reference for an AWS access key ID with
	// permissions to interact with Bedrock services. If you prefer to paste
	// your API key directly, see `aws_access_key_id_plaintext`. You must
	// provide an API key using one of the following fields: `aws_access_key_id`
	// or `aws_access_key_id_plaintext`.
	// Wire name: 'aws_access_key_id'
	AwsAccessKeyId string ``
	// An AWS access key ID with permissions to interact with Bedrock services
	// provided as a plaintext string. If you prefer to reference your key using
	// Databricks Secrets, see `aws_access_key_id`. You must provide an API key
	// using one of the following fields: `aws_access_key_id` or
	// `aws_access_key_id_plaintext`.
	// Wire name: 'aws_access_key_id_plaintext'
	AwsAccessKeyIdPlaintext string ``
	// The AWS region to use. Bedrock has to be enabled there.
	// Wire name: 'aws_region'
	AwsRegion string ``
	// The Databricks secret key reference for an AWS secret access key paired
	// with the access key ID, with permissions to interact with Bedrock
	// services. If you prefer to paste your API key directly, see
	// `aws_secret_access_key_plaintext`. You must provide an API key using one
	// of the following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	// Wire name: 'aws_secret_access_key'
	AwsSecretAccessKey string ``
	// An AWS secret access key paired with the access key ID, with permissions
	// to interact with Bedrock services provided as a plaintext string. If you
	// prefer to reference your key using Databricks Secrets, see
	// `aws_secret_access_key`. You must provide an API key using one of the
	// following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	// Wire name: 'aws_secret_access_key_plaintext'
	AwsSecretAccessKeyPlaintext string ``
	// The underlying provider in Amazon Bedrock. Supported values (case
	// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
	// Wire name: 'bedrock_provider'
	BedrockProvider AmazonBedrockConfigBedrockProvider ``
	// ARN of the instance profile that the external model will use to access
	// AWS resources. You must authenticate using an instance profile or access
	// keys. If you prefer to authenticate using access keys, see
	// `aws_access_key_id`, `aws_access_key_id_plaintext`,
	// `aws_secret_access_key` and `aws_secret_access_key_plaintext`.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string   ``
	ForceSendFields    []string `tf:"-"`
}

func (s *AmazonBedrockConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AmazonBedrockConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AmazonBedrockConfigToPb(st *AmazonBedrockConfig) (*servingpb.AmazonBedrockConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AmazonBedrockConfigPb{}
	pb.AwsAccessKeyId = st.AwsAccessKeyId
	pb.AwsAccessKeyIdPlaintext = st.AwsAccessKeyIdPlaintext
	pb.AwsRegion = st.AwsRegion
	pb.AwsSecretAccessKey = st.AwsSecretAccessKey
	pb.AwsSecretAccessKeyPlaintext = st.AwsSecretAccessKeyPlaintext
	bedrockProviderPb, err := AmazonBedrockConfigBedrockProviderToPb(&st.BedrockProvider)
	if err != nil {
		return nil, err
	}
	if bedrockProviderPb != nil {
		pb.BedrockProvider = *bedrockProviderPb
	}
	pb.InstanceProfileArn = st.InstanceProfileArn

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AmazonBedrockConfigFromPb(pb *servingpb.AmazonBedrockConfigPb) (*AmazonBedrockConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AmazonBedrockConfig{}
	st.AwsAccessKeyId = pb.AwsAccessKeyId
	st.AwsAccessKeyIdPlaintext = pb.AwsAccessKeyIdPlaintext
	st.AwsRegion = pb.AwsRegion
	st.AwsSecretAccessKey = pb.AwsSecretAccessKey
	st.AwsSecretAccessKeyPlaintext = pb.AwsSecretAccessKeyPlaintext
	bedrockProviderField, err := AmazonBedrockConfigBedrockProviderFromPb(&pb.BedrockProvider)
	if err != nil {
		return nil, err
	}
	if bedrockProviderField != nil {
		st.BedrockProvider = *bedrockProviderField
	}
	st.InstanceProfileArn = pb.InstanceProfileArn

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func AmazonBedrockConfigBedrockProviderToPb(st *AmazonBedrockConfigBedrockProvider) (*servingpb.AmazonBedrockConfigBedrockProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.AmazonBedrockConfigBedrockProviderPb(*st)
	return &pb, nil
}

func AmazonBedrockConfigBedrockProviderFromPb(pb *servingpb.AmazonBedrockConfigBedrockProviderPb) (*AmazonBedrockConfigBedrockProvider, error) {
	if pb == nil {
		return nil, nil
	}
	st := AmazonBedrockConfigBedrockProvider(*pb)
	return &st, nil
}

type AnthropicConfig struct {
	// The Databricks secret key reference for an Anthropic API key. If you
	// prefer to paste your API key directly, see `anthropic_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	// Wire name: 'anthropic_api_key'
	AnthropicApiKey string ``
	// The Anthropic API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `anthropic_api_key`. You
	// must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	// Wire name: 'anthropic_api_key_plaintext'
	AnthropicApiKeyPlaintext string   ``
	ForceSendFields          []string `tf:"-"`
}

func (s *AnthropicConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AnthropicConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AnthropicConfigToPb(st *AnthropicConfig) (*servingpb.AnthropicConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AnthropicConfigPb{}
	pb.AnthropicApiKey = st.AnthropicApiKey
	pb.AnthropicApiKeyPlaintext = st.AnthropicApiKeyPlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AnthropicConfigFromPb(pb *servingpb.AnthropicConfigPb) (*AnthropicConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AnthropicConfig{}
	st.AnthropicApiKey = pb.AnthropicApiKey
	st.AnthropicApiKeyPlaintext = pb.AnthropicApiKeyPlaintext

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ApiKeyAuth struct {
	// The name of the API key parameter used for authentication.
	// Wire name: 'key'
	Key string ``
	// The Databricks secret key reference for an API Key. If you prefer to
	// paste your token directly, see `value_plaintext`.
	// Wire name: 'value'
	Value string ``
	// The API Key provided as a plaintext string. If you prefer to reference
	// your token using Databricks Secrets, see `value`.
	// Wire name: 'value_plaintext'
	ValuePlaintext  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ApiKeyAuth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ApiKeyAuth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ApiKeyAuthToPb(st *ApiKeyAuth) (*servingpb.ApiKeyAuthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ApiKeyAuthPb{}
	pb.Key = st.Key
	pb.Value = st.Value
	pb.ValuePlaintext = st.ValuePlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ApiKeyAuthFromPb(pb *servingpb.ApiKeyAuthPb) (*ApiKeyAuth, error) {
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

type AutoCaptureConfigInput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	// Wire name: 'catalog_name'
	CatalogName string ``
	// Indicates whether the inference table is enabled.
	// Wire name: 'enabled'
	Enabled bool ``
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	// Wire name: 'schema_name'
	SchemaName string ``
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	// Wire name: 'table_name_prefix'
	TableNamePrefix string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AutoCaptureConfigInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoCaptureConfigInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AutoCaptureConfigInputToPb(st *AutoCaptureConfigInput) (*servingpb.AutoCaptureConfigInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AutoCaptureConfigInputPb{}
	pb.CatalogName = st.CatalogName
	pb.Enabled = st.Enabled
	pb.SchemaName = st.SchemaName
	pb.TableNamePrefix = st.TableNamePrefix

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AutoCaptureConfigInputFromPb(pb *servingpb.AutoCaptureConfigInputPb) (*AutoCaptureConfigInput, error) {
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

type AutoCaptureConfigOutput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	// Wire name: 'catalog_name'
	CatalogName string ``
	// Indicates whether the inference table is enabled.
	// Wire name: 'enabled'
	Enabled bool ``
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	// Wire name: 'schema_name'
	SchemaName string ``

	// Wire name: 'state'
	State *AutoCaptureState ``
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	// Wire name: 'table_name_prefix'
	TableNamePrefix string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *AutoCaptureConfigOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AutoCaptureConfigOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func AutoCaptureConfigOutputToPb(st *AutoCaptureConfigOutput) (*servingpb.AutoCaptureConfigOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AutoCaptureConfigOutputPb{}
	pb.CatalogName = st.CatalogName
	pb.Enabled = st.Enabled
	pb.SchemaName = st.SchemaName
	statePb, err := AutoCaptureStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}
	pb.TableNamePrefix = st.TableNamePrefix

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func AutoCaptureConfigOutputFromPb(pb *servingpb.AutoCaptureConfigOutputPb) (*AutoCaptureConfigOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoCaptureConfigOutput{}
	st.CatalogName = pb.CatalogName
	st.Enabled = pb.Enabled
	st.SchemaName = pb.SchemaName
	stateField, err := AutoCaptureStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	st.TableNamePrefix = pb.TableNamePrefix

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type AutoCaptureState struct {

	// Wire name: 'payload_table'
	PayloadTable *PayloadTable ``
}

func AutoCaptureStateToPb(st *AutoCaptureState) (*servingpb.AutoCaptureStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.AutoCaptureStatePb{}
	payloadTablePb, err := PayloadTableToPb(st.PayloadTable)
	if err != nil {
		return nil, err
	}
	if payloadTablePb != nil {
		pb.PayloadTable = payloadTablePb
	}

	return pb, nil
}

func AutoCaptureStateFromPb(pb *servingpb.AutoCaptureStatePb) (*AutoCaptureState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoCaptureState{}
	payloadTableField, err := PayloadTableFromPb(pb.PayloadTable)
	if err != nil {
		return nil, err
	}
	if payloadTableField != nil {
		st.PayloadTable = payloadTableField
	}

	return st, nil
}

type BearerTokenAuth struct {
	// The Databricks secret key reference for a token. If you prefer to paste
	// your token directly, see `token_plaintext`.
	// Wire name: 'token'
	Token string ``
	// The token provided as a plaintext string. If you prefer to reference your
	// token using Databricks Secrets, see `token`.
	// Wire name: 'token_plaintext'
	TokenPlaintext  string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *BearerTokenAuth) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s BearerTokenAuth) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func BearerTokenAuthToPb(st *BearerTokenAuth) (*servingpb.BearerTokenAuthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.BearerTokenAuthPb{}
	pb.Token = st.Token
	pb.TokenPlaintext = st.TokenPlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func BearerTokenAuthFromPb(pb *servingpb.BearerTokenAuthPb) (*BearerTokenAuth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BearerTokenAuth{}
	st.Token = pb.Token
	st.TokenPlaintext = pb.TokenPlaintext

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type BuildLogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	// Wire name: 'name'
	Name string `tf:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	// Wire name: 'served_model_name'
	ServedModelName string `tf:"-"`
}

func BuildLogsRequestToPb(st *BuildLogsRequest) (*servingpb.BuildLogsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.BuildLogsRequestPb{}
	pb.Name = st.Name
	pb.ServedModelName = st.ServedModelName

	return pb, nil
}

func BuildLogsRequestFromPb(pb *servingpb.BuildLogsRequestPb) (*BuildLogsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BuildLogsRequest{}
	st.Name = pb.Name
	st.ServedModelName = pb.ServedModelName

	return st, nil
}

type BuildLogsResponse struct {
	// The logs associated with building the served entity's environment.
	// Wire name: 'logs'
	Logs string ``
}

func BuildLogsResponseToPb(st *BuildLogsResponse) (*servingpb.BuildLogsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.BuildLogsResponsePb{}
	pb.Logs = st.Logs

	return pb, nil
}

func BuildLogsResponseFromPb(pb *servingpb.BuildLogsResponsePb) (*BuildLogsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BuildLogsResponse{}
	st.Logs = pb.Logs

	return st, nil
}

type ChatMessage struct {
	// The content of the message.
	// Wire name: 'content'
	Content string ``
	// The role of the message. One of [system, user, assistant].
	// Wire name: 'role'
	Role            ChatMessageRole ``
	ForceSendFields []string        `tf:"-"`
}

func (s *ChatMessage) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ChatMessage) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ChatMessageToPb(st *ChatMessage) (*servingpb.ChatMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ChatMessagePb{}
	pb.Content = st.Content
	rolePb, err := ChatMessageRoleToPb(&st.Role)
	if err != nil {
		return nil, err
	}
	if rolePb != nil {
		pb.Role = *rolePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ChatMessageFromPb(pb *servingpb.ChatMessagePb) (*ChatMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ChatMessage{}
	st.Content = pb.Content
	roleField, err := ChatMessageRoleFromPb(&pb.Role)
	if err != nil {
		return nil, err
	}
	if roleField != nil {
		st.Role = *roleField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ChatMessageRoleToPb(st *ChatMessageRole) (*servingpb.ChatMessageRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.ChatMessageRolePb(*st)
	return &pb, nil
}

func ChatMessageRoleFromPb(pb *servingpb.ChatMessageRolePb) (*ChatMessageRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := ChatMessageRole(*pb)
	return &st, nil
}

type CohereConfig struct {
	// This is an optional field to provide a customized base URL for the Cohere
	// API. If left unspecified, the standard Cohere base URL is used.
	// Wire name: 'cohere_api_base'
	CohereApiBase string ``
	// The Databricks secret key reference for a Cohere API key. If you prefer
	// to paste your API key directly, see `cohere_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `cohere_api_key` or
	// `cohere_api_key_plaintext`.
	// Wire name: 'cohere_api_key'
	CohereApiKey string ``
	// The Cohere API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `cohere_api_key`. You
	// must provide an API key using one of the following fields:
	// `cohere_api_key` or `cohere_api_key_plaintext`.
	// Wire name: 'cohere_api_key_plaintext'
	CohereApiKeyPlaintext string   ``
	ForceSendFields       []string `tf:"-"`
}

func (s *CohereConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CohereConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CohereConfigToPb(st *CohereConfig) (*servingpb.CohereConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.CohereConfigPb{}
	pb.CohereApiBase = st.CohereApiBase
	pb.CohereApiKey = st.CohereApiKey
	pb.CohereApiKeyPlaintext = st.CohereApiKeyPlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CohereConfigFromPb(pb *servingpb.CohereConfigPb) (*CohereConfig, error) {
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

type CreatePtEndpointRequest struct {
	// The AI Gateway configuration for the serving endpoint.
	// Wire name: 'ai_gateway'
	AiGateway *AiGatewayConfig ``
	// The budget policy associated with the endpoint.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// The core config of the serving endpoint.
	// Wire name: 'config'
	Config PtEndpointCoreConfig ``
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	// Wire name: 'name'
	Name string ``
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	// Wire name: 'tags'
	Tags            []EndpointTag ``
	ForceSendFields []string      `tf:"-"`
}

func (s *CreatePtEndpointRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreatePtEndpointRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreatePtEndpointRequestToPb(st *CreatePtEndpointRequest) (*servingpb.CreatePtEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.CreatePtEndpointRequestPb{}
	aiGatewayPb, err := AiGatewayConfigToPb(st.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayPb != nil {
		pb.AiGateway = aiGatewayPb
	}
	pb.BudgetPolicyId = st.BudgetPolicyId
	configPb, err := PtEndpointCoreConfigToPb(&st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = *configPb
	}
	pb.Name = st.Name

	var tagsPb []servingpb.EndpointTagPb
	for _, item := range st.Tags {
		itemPb, err := EndpointTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreatePtEndpointRequestFromPb(pb *servingpb.CreatePtEndpointRequestPb) (*CreatePtEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePtEndpointRequest{}
	aiGatewayField, err := AiGatewayConfigFromPb(pb.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayField != nil {
		st.AiGateway = aiGatewayField
	}
	st.BudgetPolicyId = pb.BudgetPolicyId
	configField, err := PtEndpointCoreConfigFromPb(&pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = *configField
	}
	st.Name = pb.Name

	var tagsField []EndpointTag
	for _, itemPb := range pb.Tags {
		item, err := EndpointTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type CreateServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	// Wire name: 'ai_gateway'
	AiGateway *AiGatewayConfig ``
	// The budget policy to be applied to the serving endpoint.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// The core config of the serving endpoint.
	// Wire name: 'config'
	Config *EndpointCoreConfigInput ``

	// Wire name: 'description'
	Description string ``
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	// Wire name: 'name'
	Name string ``
	// Rate limits to be applied to the serving endpoint. NOTE: this field is
	// deprecated, please use AI Gateway to manage rate limits.
	// Wire name: 'rate_limits'
	RateLimits []RateLimit ``
	// Enable route optimization for the serving endpoint.
	// Wire name: 'route_optimized'
	RouteOptimized bool ``
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	// Wire name: 'tags'
	Tags            []EndpointTag ``
	ForceSendFields []string      `tf:"-"`
}

func (s *CreateServingEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s CreateServingEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func CreateServingEndpointToPb(st *CreateServingEndpoint) (*servingpb.CreateServingEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.CreateServingEndpointPb{}
	aiGatewayPb, err := AiGatewayConfigToPb(st.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayPb != nil {
		pb.AiGateway = aiGatewayPb
	}
	pb.BudgetPolicyId = st.BudgetPolicyId
	configPb, err := EndpointCoreConfigInputToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}
	pb.Description = st.Description
	pb.Name = st.Name

	var rateLimitsPb []servingpb.RateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := RateLimitToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rateLimitsPb = append(rateLimitsPb, *itemPb)
		}
	}
	pb.RateLimits = rateLimitsPb
	pb.RouteOptimized = st.RouteOptimized

	var tagsPb []servingpb.EndpointTagPb
	for _, item := range st.Tags {
		itemPb, err := EndpointTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func CreateServingEndpointFromPb(pb *servingpb.CreateServingEndpointPb) (*CreateServingEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServingEndpoint{}
	aiGatewayField, err := AiGatewayConfigFromPb(pb.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayField != nil {
		st.AiGateway = aiGatewayField
	}
	st.BudgetPolicyId = pb.BudgetPolicyId
	configField, err := EndpointCoreConfigInputFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	st.Description = pb.Description
	st.Name = pb.Name

	var rateLimitsField []RateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := RateLimitFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rateLimitsField = append(rateLimitsField, *item)
		}
	}
	st.RateLimits = rateLimitsField
	st.RouteOptimized = pb.RouteOptimized

	var tagsField []EndpointTag
	for _, itemPb := range pb.Tags {
		item, err := EndpointTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

// Configs needed to create a custom provider model route.
type CustomProviderConfig struct {
	// This is a field to provide API key authentication for the custom provider
	// API. You can only specify one authentication method.
	// Wire name: 'api_key_auth'
	ApiKeyAuth *ApiKeyAuth ``
	// This is a field to provide bearer token authentication for the custom
	// provider API. You can only specify one authentication method.
	// Wire name: 'bearer_token_auth'
	BearerTokenAuth *BearerTokenAuth ``
	// This is a field to provide the URL of the custom provider API.
	// Wire name: 'custom_provider_url'
	CustomProviderUrl string ``
}

func CustomProviderConfigToPb(st *CustomProviderConfig) (*servingpb.CustomProviderConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.CustomProviderConfigPb{}
	apiKeyAuthPb, err := ApiKeyAuthToPb(st.ApiKeyAuth)
	if err != nil {
		return nil, err
	}
	if apiKeyAuthPb != nil {
		pb.ApiKeyAuth = apiKeyAuthPb
	}
	bearerTokenAuthPb, err := BearerTokenAuthToPb(st.BearerTokenAuth)
	if err != nil {
		return nil, err
	}
	if bearerTokenAuthPb != nil {
		pb.BearerTokenAuth = bearerTokenAuthPb
	}
	pb.CustomProviderUrl = st.CustomProviderUrl

	return pb, nil
}

func CustomProviderConfigFromPb(pb *servingpb.CustomProviderConfigPb) (*CustomProviderConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomProviderConfig{}
	apiKeyAuthField, err := ApiKeyAuthFromPb(pb.ApiKeyAuth)
	if err != nil {
		return nil, err
	}
	if apiKeyAuthField != nil {
		st.ApiKeyAuth = apiKeyAuthField
	}
	bearerTokenAuthField, err := BearerTokenAuthFromPb(pb.BearerTokenAuth)
	if err != nil {
		return nil, err
	}
	if bearerTokenAuthField != nil {
		st.BearerTokenAuth = bearerTokenAuthField
	}
	st.CustomProviderUrl = pb.CustomProviderUrl

	return st, nil
}

// Details necessary to query this object's API through the DataPlane APIs.
type DataPlaneInfo struct {
	// Authorization details as a string.
	// Wire name: 'authorization_details'
	AuthorizationDetails string ``
	// The URL of the endpoint for this operation in the dataplane.
	// Wire name: 'endpoint_url'
	EndpointUrl     string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *DataPlaneInfo) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DataPlaneInfo) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DataPlaneInfoToPb(st *DataPlaneInfo) (*servingpb.DataPlaneInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.DataPlaneInfoPb{}
	pb.AuthorizationDetails = st.AuthorizationDetails
	pb.EndpointUrl = st.EndpointUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DataPlaneInfoFromPb(pb *servingpb.DataPlaneInfoPb) (*DataPlaneInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataPlaneInfo{}
	st.AuthorizationDetails = pb.AuthorizationDetails
	st.EndpointUrl = pb.EndpointUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type DatabricksModelServingConfig struct {
	// The Databricks secret key reference for a Databricks API token that
	// corresponds to a user or service principal with Can Query access to the
	// model serving endpoint pointed to by this external model. If you prefer
	// to paste your API key directly, see `databricks_api_token_plaintext`. You
	// must provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	// Wire name: 'databricks_api_token'
	DatabricksApiToken string ``
	// The Databricks API token that corresponds to a user or service principal
	// with Can Query access to the model serving endpoint pointed to by this
	// external model provided as a plaintext string. If you prefer to reference
	// your key using Databricks Secrets, see `databricks_api_token`. You must
	// provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	// Wire name: 'databricks_api_token_plaintext'
	DatabricksApiTokenPlaintext string ``
	// The URL of the Databricks workspace containing the model serving endpoint
	// pointed to by this external model.
	// Wire name: 'databricks_workspace_url'
	DatabricksWorkspaceUrl string   ``
	ForceSendFields        []string `tf:"-"`
}

func (s *DatabricksModelServingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s DatabricksModelServingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func DatabricksModelServingConfigToPb(st *DatabricksModelServingConfig) (*servingpb.DatabricksModelServingConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.DatabricksModelServingConfigPb{}
	pb.DatabricksApiToken = st.DatabricksApiToken
	pb.DatabricksApiTokenPlaintext = st.DatabricksApiTokenPlaintext
	pb.DatabricksWorkspaceUrl = st.DatabricksWorkspaceUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func DatabricksModelServingConfigFromPb(pb *servingpb.DatabricksModelServingConfigPb) (*DatabricksModelServingConfig, error) {
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

type DataframeSplitInput struct {

	// Wire name: 'columns'
	Columns []any ``

	// Wire name: 'data'
	Data []any ``

	// Wire name: 'index'
	Index []int ``
}

func DataframeSplitInputToPb(st *DataframeSplitInput) (*servingpb.DataframeSplitInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.DataframeSplitInputPb{}
	pb.Columns = st.Columns
	pb.Data = st.Data
	pb.Index = st.Index

	return pb, nil
}

func DataframeSplitInputFromPb(pb *servingpb.DataframeSplitInputPb) (*DataframeSplitInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataframeSplitInput{}
	st.Columns = pb.Columns
	st.Data = pb.Data
	st.Index = pb.Index

	return st, nil
}

type DeleteServingEndpointRequest struct {

	// Wire name: 'name'
	Name string `tf:"-"`
}

func DeleteServingEndpointRequestToPb(st *DeleteServingEndpointRequest) (*servingpb.DeleteServingEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.DeleteServingEndpointRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteServingEndpointRequestFromPb(pb *servingpb.DeleteServingEndpointRequestPb) (*DeleteServingEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServingEndpointRequest{}
	st.Name = pb.Name

	return st, nil
}

type EmbeddingsV1ResponseEmbeddingElement struct {

	// Wire name: 'embedding'
	Embedding []float64 ``
	// The index of the embedding in the response.
	// Wire name: 'index'
	Index int ``
	// This will always be 'embedding'.
	// Wire name: 'object'
	Object          EmbeddingsV1ResponseEmbeddingElementObject ``
	ForceSendFields []string                                   `tf:"-"`
}

func (s *EmbeddingsV1ResponseEmbeddingElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EmbeddingsV1ResponseEmbeddingElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EmbeddingsV1ResponseEmbeddingElementToPb(st *EmbeddingsV1ResponseEmbeddingElement) (*servingpb.EmbeddingsV1ResponseEmbeddingElementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.EmbeddingsV1ResponseEmbeddingElementPb{}
	pb.Embedding = st.Embedding
	pb.Index = st.Index
	objectPb, err := EmbeddingsV1ResponseEmbeddingElementObjectToPb(&st.Object)
	if err != nil {
		return nil, err
	}
	if objectPb != nil {
		pb.Object = *objectPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EmbeddingsV1ResponseEmbeddingElementFromPb(pb *servingpb.EmbeddingsV1ResponseEmbeddingElementPb) (*EmbeddingsV1ResponseEmbeddingElement, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EmbeddingsV1ResponseEmbeddingElement{}
	st.Embedding = pb.Embedding
	st.Index = pb.Index
	objectField, err := EmbeddingsV1ResponseEmbeddingElementObjectFromPb(&pb.Object)
	if err != nil {
		return nil, err
	}
	if objectField != nil {
		st.Object = *objectField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func EmbeddingsV1ResponseEmbeddingElementObjectToPb(st *EmbeddingsV1ResponseEmbeddingElementObject) (*servingpb.EmbeddingsV1ResponseEmbeddingElementObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.EmbeddingsV1ResponseEmbeddingElementObjectPb(*st)
	return &pb, nil
}

func EmbeddingsV1ResponseEmbeddingElementObjectFromPb(pb *servingpb.EmbeddingsV1ResponseEmbeddingElementObjectPb) (*EmbeddingsV1ResponseEmbeddingElementObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := EmbeddingsV1ResponseEmbeddingElementObject(*pb)
	return &st, nil
}

type EndpointCoreConfigInput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	// Wire name: 'auto_capture_config'
	AutoCaptureConfig *AutoCaptureConfigInput ``
	// The name of the serving endpoint to update. This field is required.
	// Wire name: 'name'
	Name string `tf:"-"`
	// The list of served entities under the serving endpoint config.
	// Wire name: 'served_entities'
	ServedEntities []ServedEntityInput ``
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	// Wire name: 'served_models'
	ServedModels []ServedModelInput ``
	// The traffic configuration associated with the serving endpoint config.
	// Wire name: 'traffic_config'
	TrafficConfig *TrafficConfig ``
}

func EndpointCoreConfigInputToPb(st *EndpointCoreConfigInput) (*servingpb.EndpointCoreConfigInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.EndpointCoreConfigInputPb{}
	autoCaptureConfigPb, err := AutoCaptureConfigInputToPb(st.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigPb != nil {
		pb.AutoCaptureConfig = autoCaptureConfigPb
	}
	pb.Name = st.Name

	var servedEntitiesPb []servingpb.ServedEntityInputPb
	for _, item := range st.ServedEntities {
		itemPb, err := ServedEntityInputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedEntitiesPb = append(servedEntitiesPb, *itemPb)
		}
	}
	pb.ServedEntities = servedEntitiesPb

	var servedModelsPb []servingpb.ServedModelInputPb
	for _, item := range st.ServedModels {
		itemPb, err := ServedModelInputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedModelsPb = append(servedModelsPb, *itemPb)
		}
	}
	pb.ServedModels = servedModelsPb
	trafficConfigPb, err := TrafficConfigToPb(st.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigPb != nil {
		pb.TrafficConfig = trafficConfigPb
	}

	return pb, nil
}

func EndpointCoreConfigInputFromPb(pb *servingpb.EndpointCoreConfigInputPb) (*EndpointCoreConfigInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointCoreConfigInput{}
	autoCaptureConfigField, err := AutoCaptureConfigInputFromPb(pb.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigField != nil {
		st.AutoCaptureConfig = autoCaptureConfigField
	}
	st.Name = pb.Name

	var servedEntitiesField []ServedEntityInput
	for _, itemPb := range pb.ServedEntities {
		item, err := ServedEntityInputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedEntitiesField = append(servedEntitiesField, *item)
		}
	}
	st.ServedEntities = servedEntitiesField

	var servedModelsField []ServedModelInput
	for _, itemPb := range pb.ServedModels {
		item, err := ServedModelInputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedModelsField = append(servedModelsField, *item)
		}
	}
	st.ServedModels = servedModelsField
	trafficConfigField, err := TrafficConfigFromPb(pb.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigField != nil {
		st.TrafficConfig = trafficConfigField
	}

	return st, nil
}

type EndpointCoreConfigOutput struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	// Wire name: 'auto_capture_config'
	AutoCaptureConfig *AutoCaptureConfigOutput ``
	// The config version that the serving endpoint is currently serving.
	// Wire name: 'config_version'
	ConfigVersion int64 ``
	// The list of served entities under the serving endpoint config.
	// Wire name: 'served_entities'
	ServedEntities []ServedEntityOutput ``
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	// Wire name: 'served_models'
	ServedModels []ServedModelOutput ``
	// The traffic configuration associated with the serving endpoint config.
	// Wire name: 'traffic_config'
	TrafficConfig   *TrafficConfig ``
	ForceSendFields []string       `tf:"-"`
}

func (s *EndpointCoreConfigOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointCoreConfigOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EndpointCoreConfigOutputToPb(st *EndpointCoreConfigOutput) (*servingpb.EndpointCoreConfigOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.EndpointCoreConfigOutputPb{}
	autoCaptureConfigPb, err := AutoCaptureConfigOutputToPb(st.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigPb != nil {
		pb.AutoCaptureConfig = autoCaptureConfigPb
	}
	pb.ConfigVersion = st.ConfigVersion

	var servedEntitiesPb []servingpb.ServedEntityOutputPb
	for _, item := range st.ServedEntities {
		itemPb, err := ServedEntityOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedEntitiesPb = append(servedEntitiesPb, *itemPb)
		}
	}
	pb.ServedEntities = servedEntitiesPb

	var servedModelsPb []servingpb.ServedModelOutputPb
	for _, item := range st.ServedModels {
		itemPb, err := ServedModelOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedModelsPb = append(servedModelsPb, *itemPb)
		}
	}
	pb.ServedModels = servedModelsPb
	trafficConfigPb, err := TrafficConfigToPb(st.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigPb != nil {
		pb.TrafficConfig = trafficConfigPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EndpointCoreConfigOutputFromPb(pb *servingpb.EndpointCoreConfigOutputPb) (*EndpointCoreConfigOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointCoreConfigOutput{}
	autoCaptureConfigField, err := AutoCaptureConfigOutputFromPb(pb.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigField != nil {
		st.AutoCaptureConfig = autoCaptureConfigField
	}
	st.ConfigVersion = pb.ConfigVersion

	var servedEntitiesField []ServedEntityOutput
	for _, itemPb := range pb.ServedEntities {
		item, err := ServedEntityOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedEntitiesField = append(servedEntitiesField, *item)
		}
	}
	st.ServedEntities = servedEntitiesField

	var servedModelsField []ServedModelOutput
	for _, itemPb := range pb.ServedModels {
		item, err := ServedModelOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedModelsField = append(servedModelsField, *item)
		}
	}
	st.ServedModels = servedModelsField
	trafficConfigField, err := TrafficConfigFromPb(pb.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigField != nil {
		st.TrafficConfig = trafficConfigField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type EndpointCoreConfigSummary struct {
	// The list of served entities under the serving endpoint config.
	// Wire name: 'served_entities'
	ServedEntities []ServedEntitySpec ``
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	// Wire name: 'served_models'
	ServedModels []ServedModelSpec ``
}

func EndpointCoreConfigSummaryToPb(st *EndpointCoreConfigSummary) (*servingpb.EndpointCoreConfigSummaryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.EndpointCoreConfigSummaryPb{}

	var servedEntitiesPb []servingpb.ServedEntitySpecPb
	for _, item := range st.ServedEntities {
		itemPb, err := ServedEntitySpecToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedEntitiesPb = append(servedEntitiesPb, *itemPb)
		}
	}
	pb.ServedEntities = servedEntitiesPb

	var servedModelsPb []servingpb.ServedModelSpecPb
	for _, item := range st.ServedModels {
		itemPb, err := ServedModelSpecToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedModelsPb = append(servedModelsPb, *itemPb)
		}
	}
	pb.ServedModels = servedModelsPb

	return pb, nil
}

func EndpointCoreConfigSummaryFromPb(pb *servingpb.EndpointCoreConfigSummaryPb) (*EndpointCoreConfigSummary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointCoreConfigSummary{}

	var servedEntitiesField []ServedEntitySpec
	for _, itemPb := range pb.ServedEntities {
		item, err := ServedEntitySpecFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedEntitiesField = append(servedEntitiesField, *item)
		}
	}
	st.ServedEntities = servedEntitiesField

	var servedModelsField []ServedModelSpec
	for _, itemPb := range pb.ServedModels {
		item, err := ServedModelSpecFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedModelsField = append(servedModelsField, *item)
		}
	}
	st.ServedModels = servedModelsField

	return st, nil
}

type EndpointPendingConfig struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	// Wire name: 'auto_capture_config'
	AutoCaptureConfig *AutoCaptureConfigOutput ``
	// The config version that the serving endpoint is currently serving.
	// Wire name: 'config_version'
	ConfigVersion int ``
	// The list of served entities belonging to the last issued update to the
	// serving endpoint.
	// Wire name: 'served_entities'
	ServedEntities []ServedEntityOutput ``
	// (Deprecated, use served_entities instead) The list of served models
	// belonging to the last issued update to the serving endpoint.
	// Wire name: 'served_models'
	ServedModels []ServedModelOutput ``
	// The timestamp when the update to the pending config started.
	// Wire name: 'start_time'
	StartTime int64 ``
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	// Wire name: 'traffic_config'
	TrafficConfig   *TrafficConfig ``
	ForceSendFields []string       `tf:"-"`
}

func (s *EndpointPendingConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointPendingConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EndpointPendingConfigToPb(st *EndpointPendingConfig) (*servingpb.EndpointPendingConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.EndpointPendingConfigPb{}
	autoCaptureConfigPb, err := AutoCaptureConfigOutputToPb(st.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigPb != nil {
		pb.AutoCaptureConfig = autoCaptureConfigPb
	}
	pb.ConfigVersion = st.ConfigVersion

	var servedEntitiesPb []servingpb.ServedEntityOutputPb
	for _, item := range st.ServedEntities {
		itemPb, err := ServedEntityOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedEntitiesPb = append(servedEntitiesPb, *itemPb)
		}
	}
	pb.ServedEntities = servedEntitiesPb

	var servedModelsPb []servingpb.ServedModelOutputPb
	for _, item := range st.ServedModels {
		itemPb, err := ServedModelOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedModelsPb = append(servedModelsPb, *itemPb)
		}
	}
	pb.ServedModels = servedModelsPb
	pb.StartTime = st.StartTime
	trafficConfigPb, err := TrafficConfigToPb(st.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigPb != nil {
		pb.TrafficConfig = trafficConfigPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EndpointPendingConfigFromPb(pb *servingpb.EndpointPendingConfigPb) (*EndpointPendingConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointPendingConfig{}
	autoCaptureConfigField, err := AutoCaptureConfigOutputFromPb(pb.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigField != nil {
		st.AutoCaptureConfig = autoCaptureConfigField
	}
	st.ConfigVersion = pb.ConfigVersion

	var servedEntitiesField []ServedEntityOutput
	for _, itemPb := range pb.ServedEntities {
		item, err := ServedEntityOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedEntitiesField = append(servedEntitiesField, *item)
		}
	}
	st.ServedEntities = servedEntitiesField

	var servedModelsField []ServedModelOutput
	for _, itemPb := range pb.ServedModels {
		item, err := ServedModelOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedModelsField = append(servedModelsField, *item)
		}
	}
	st.ServedModels = servedModelsField
	st.StartTime = pb.StartTime
	trafficConfigField, err := TrafficConfigFromPb(pb.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigField != nil {
		st.TrafficConfig = trafficConfigField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type EndpointState struct {
	// The state of an endpoint's config update. This informs the user if the
	// pending_config is in progress, if the update failed, or if there is no
	// update in progress. Note that if the endpoint's config_update state value
	// is IN_PROGRESS, another update can not be made until the update completes
	// or fails.
	// Wire name: 'config_update'
	ConfigUpdate EndpointStateConfigUpdate ``
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served entities in its
	// active configuration are ready. If any of the actively served entities
	// are in a non-ready state, the endpoint state will be NOT_READY.
	// Wire name: 'ready'
	Ready EndpointStateReady ``
}

func EndpointStateToPb(st *EndpointState) (*servingpb.EndpointStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.EndpointStatePb{}
	configUpdatePb, err := EndpointStateConfigUpdateToPb(&st.ConfigUpdate)
	if err != nil {
		return nil, err
	}
	if configUpdatePb != nil {
		pb.ConfigUpdate = *configUpdatePb
	}
	readyPb, err := EndpointStateReadyToPb(&st.Ready)
	if err != nil {
		return nil, err
	}
	if readyPb != nil {
		pb.Ready = *readyPb
	}

	return pb, nil
}

func EndpointStateFromPb(pb *servingpb.EndpointStatePb) (*EndpointState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointState{}
	configUpdateField, err := EndpointStateConfigUpdateFromPb(&pb.ConfigUpdate)
	if err != nil {
		return nil, err
	}
	if configUpdateField != nil {
		st.ConfigUpdate = *configUpdateField
	}
	readyField, err := EndpointStateReadyFromPb(&pb.Ready)
	if err != nil {
		return nil, err
	}
	if readyField != nil {
		st.Ready = *readyField
	}

	return st, nil
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

func EndpointStateConfigUpdateToPb(st *EndpointStateConfigUpdate) (*servingpb.EndpointStateConfigUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.EndpointStateConfigUpdatePb(*st)
	return &pb, nil
}

func EndpointStateConfigUpdateFromPb(pb *servingpb.EndpointStateConfigUpdatePb) (*EndpointStateConfigUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointStateConfigUpdate(*pb)
	return &st, nil
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

func EndpointStateReadyToPb(st *EndpointStateReady) (*servingpb.EndpointStateReadyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.EndpointStateReadyPb(*st)
	return &pb, nil
}

func EndpointStateReadyFromPb(pb *servingpb.EndpointStateReadyPb) (*EndpointStateReady, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointStateReady(*pb)
	return &st, nil
}

type EndpointTag struct {
	// Key field for a serving endpoint tag.
	// Wire name: 'key'
	Key string ``
	// Optional value field for a serving endpoint tag.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *EndpointTag) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s EndpointTag) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func EndpointTagToPb(st *EndpointTag) (*servingpb.EndpointTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.EndpointTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func EndpointTagFromPb(pb *servingpb.EndpointTagPb) (*EndpointTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type EndpointTags struct {

	// Wire name: 'tags'
	Tags []EndpointTag ``
}

func EndpointTagsToPb(st *EndpointTags) (*servingpb.EndpointTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.EndpointTagsPb{}

	var tagsPb []servingpb.EndpointTagPb
	for _, item := range st.Tags {
		itemPb, err := EndpointTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	return pb, nil
}

func EndpointTagsFromPb(pb *servingpb.EndpointTagsPb) (*EndpointTags, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTags{}

	var tagsField []EndpointTag
	for _, itemPb := range pb.Tags {
		item, err := EndpointTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	return st, nil
}

type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func ExportMetricsRequestToPb(st *ExportMetricsRequest) (*servingpb.ExportMetricsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ExportMetricsRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func ExportMetricsRequestFromPb(pb *servingpb.ExportMetricsRequestPb) (*ExportMetricsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportMetricsRequest{}
	st.Name = pb.Name

	return st, nil
}

type ExportMetricsResponse struct {

	// Wire name: 'contents'
	Contents io.ReadCloser `tf:"-"`
}

func ExportMetricsResponseToPb(st *ExportMetricsResponse) (*servingpb.ExportMetricsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ExportMetricsResponsePb{}
	pb.Contents = st.Contents

	return pb, nil
}

func ExportMetricsResponseFromPb(pb *servingpb.ExportMetricsResponsePb) (*ExportMetricsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportMetricsResponse{}
	st.Contents = pb.Contents

	return st, nil
}

// Simple Proto message for testing
type ExternalFunctionRequest struct {
	// The connection name to use. This is required to identify the external
	// connection.
	// Wire name: 'connection_name'
	ConnectionName string ``
	// Additional headers for the request. If not provided, only auth headers
	// from connections would be passed.
	// Wire name: 'headers'
	Headers string ``
	// The JSON payload to send in the request body.
	// Wire name: 'json'
	Json string ``
	// The HTTP method to use (e.g., 'GET', 'POST').
	// Wire name: 'method'
	Method ExternalFunctionRequestHttpMethod ``
	// Query parameters for the request.
	// Wire name: 'params'
	Params string ``
	// The relative path for the API endpoint. This is required.
	// Wire name: 'path'
	Path            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalFunctionRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalFunctionRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalFunctionRequestToPb(st *ExternalFunctionRequest) (*servingpb.ExternalFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ExternalFunctionRequestPb{}
	pb.ConnectionName = st.ConnectionName
	pb.Headers = st.Headers
	pb.Json = st.Json
	methodPb, err := ExternalFunctionRequestHttpMethodToPb(&st.Method)
	if err != nil {
		return nil, err
	}
	if methodPb != nil {
		pb.Method = *methodPb
	}
	pb.Params = st.Params
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalFunctionRequestFromPb(pb *servingpb.ExternalFunctionRequestPb) (*ExternalFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalFunctionRequest{}
	st.ConnectionName = pb.ConnectionName
	st.Headers = pb.Headers
	st.Json = pb.Json
	methodField, err := ExternalFunctionRequestHttpMethodFromPb(&pb.Method)
	if err != nil {
		return nil, err
	}
	if methodField != nil {
		st.Method = *methodField
	}
	st.Params = pb.Params
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ExternalFunctionRequestHttpMethodToPb(st *ExternalFunctionRequestHttpMethod) (*servingpb.ExternalFunctionRequestHttpMethodPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.ExternalFunctionRequestHttpMethodPb(*st)
	return &pb, nil
}

func ExternalFunctionRequestHttpMethodFromPb(pb *servingpb.ExternalFunctionRequestHttpMethodPb) (*ExternalFunctionRequestHttpMethod, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExternalFunctionRequestHttpMethod(*pb)
	return &st, nil
}

type ExternalModel struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	// Wire name: 'ai21labs_config'
	Ai21labsConfig *Ai21LabsConfig ``
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	// Wire name: 'amazon_bedrock_config'
	AmazonBedrockConfig *AmazonBedrockConfig ``
	// Anthropic Config. Only required if the provider is 'anthropic'.
	// Wire name: 'anthropic_config'
	AnthropicConfig *AnthropicConfig ``
	// Cohere Config. Only required if the provider is 'cohere'.
	// Wire name: 'cohere_config'
	CohereConfig *CohereConfig ``
	// Custom Provider Config. Only required if the provider is 'custom'.
	// Wire name: 'custom_provider_config'
	CustomProviderConfig *CustomProviderConfig ``
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	// Wire name: 'databricks_model_serving_config'
	DatabricksModelServingConfig *DatabricksModelServingConfig ``
	// Google Cloud Vertex AI Config. Only required if the provider is
	// 'google-cloud-vertex-ai'.
	// Wire name: 'google_cloud_vertex_ai_config'
	GoogleCloudVertexAiConfig *GoogleCloudVertexAiConfig ``
	// The name of the external model.
	// Wire name: 'name'
	Name string ``
	// OpenAI Config. Only required if the provider is 'openai'.
	// Wire name: 'openai_config'
	OpenaiConfig *OpenAiConfig ``
	// PaLM Config. Only required if the provider is 'palm'.
	// Wire name: 'palm_config'
	PalmConfig *PaLmConfig ``
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
	// 'databricks-model-serving', 'google-cloud-vertex-ai', 'openai', 'palm',
	// and 'custom'.
	// Wire name: 'provider'
	Provider ExternalModelProvider ``
	// The task type of the external model.
	// Wire name: 'task'
	Task string ``
}

func ExternalModelToPb(st *ExternalModel) (*servingpb.ExternalModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ExternalModelPb{}
	ai21labsConfigPb, err := Ai21LabsConfigToPb(st.Ai21labsConfig)
	if err != nil {
		return nil, err
	}
	if ai21labsConfigPb != nil {
		pb.Ai21labsConfig = ai21labsConfigPb
	}
	amazonBedrockConfigPb, err := AmazonBedrockConfigToPb(st.AmazonBedrockConfig)
	if err != nil {
		return nil, err
	}
	if amazonBedrockConfigPb != nil {
		pb.AmazonBedrockConfig = amazonBedrockConfigPb
	}
	anthropicConfigPb, err := AnthropicConfigToPb(st.AnthropicConfig)
	if err != nil {
		return nil, err
	}
	if anthropicConfigPb != nil {
		pb.AnthropicConfig = anthropicConfigPb
	}
	cohereConfigPb, err := CohereConfigToPb(st.CohereConfig)
	if err != nil {
		return nil, err
	}
	if cohereConfigPb != nil {
		pb.CohereConfig = cohereConfigPb
	}
	customProviderConfigPb, err := CustomProviderConfigToPb(st.CustomProviderConfig)
	if err != nil {
		return nil, err
	}
	if customProviderConfigPb != nil {
		pb.CustomProviderConfig = customProviderConfigPb
	}
	databricksModelServingConfigPb, err := DatabricksModelServingConfigToPb(st.DatabricksModelServingConfig)
	if err != nil {
		return nil, err
	}
	if databricksModelServingConfigPb != nil {
		pb.DatabricksModelServingConfig = databricksModelServingConfigPb
	}
	googleCloudVertexAiConfigPb, err := GoogleCloudVertexAiConfigToPb(st.GoogleCloudVertexAiConfig)
	if err != nil {
		return nil, err
	}
	if googleCloudVertexAiConfigPb != nil {
		pb.GoogleCloudVertexAiConfig = googleCloudVertexAiConfigPb
	}
	pb.Name = st.Name
	openaiConfigPb, err := OpenAiConfigToPb(st.OpenaiConfig)
	if err != nil {
		return nil, err
	}
	if openaiConfigPb != nil {
		pb.OpenaiConfig = openaiConfigPb
	}
	palmConfigPb, err := PaLmConfigToPb(st.PalmConfig)
	if err != nil {
		return nil, err
	}
	if palmConfigPb != nil {
		pb.PalmConfig = palmConfigPb
	}
	providerPb, err := ExternalModelProviderToPb(&st.Provider)
	if err != nil {
		return nil, err
	}
	if providerPb != nil {
		pb.Provider = *providerPb
	}
	pb.Task = st.Task

	return pb, nil
}

func ExternalModelFromPb(pb *servingpb.ExternalModelPb) (*ExternalModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalModel{}
	ai21labsConfigField, err := Ai21LabsConfigFromPb(pb.Ai21labsConfig)
	if err != nil {
		return nil, err
	}
	if ai21labsConfigField != nil {
		st.Ai21labsConfig = ai21labsConfigField
	}
	amazonBedrockConfigField, err := AmazonBedrockConfigFromPb(pb.AmazonBedrockConfig)
	if err != nil {
		return nil, err
	}
	if amazonBedrockConfigField != nil {
		st.AmazonBedrockConfig = amazonBedrockConfigField
	}
	anthropicConfigField, err := AnthropicConfigFromPb(pb.AnthropicConfig)
	if err != nil {
		return nil, err
	}
	if anthropicConfigField != nil {
		st.AnthropicConfig = anthropicConfigField
	}
	cohereConfigField, err := CohereConfigFromPb(pb.CohereConfig)
	if err != nil {
		return nil, err
	}
	if cohereConfigField != nil {
		st.CohereConfig = cohereConfigField
	}
	customProviderConfigField, err := CustomProviderConfigFromPb(pb.CustomProviderConfig)
	if err != nil {
		return nil, err
	}
	if customProviderConfigField != nil {
		st.CustomProviderConfig = customProviderConfigField
	}
	databricksModelServingConfigField, err := DatabricksModelServingConfigFromPb(pb.DatabricksModelServingConfig)
	if err != nil {
		return nil, err
	}
	if databricksModelServingConfigField != nil {
		st.DatabricksModelServingConfig = databricksModelServingConfigField
	}
	googleCloudVertexAiConfigField, err := GoogleCloudVertexAiConfigFromPb(pb.GoogleCloudVertexAiConfig)
	if err != nil {
		return nil, err
	}
	if googleCloudVertexAiConfigField != nil {
		st.GoogleCloudVertexAiConfig = googleCloudVertexAiConfigField
	}
	st.Name = pb.Name
	openaiConfigField, err := OpenAiConfigFromPb(pb.OpenaiConfig)
	if err != nil {
		return nil, err
	}
	if openaiConfigField != nil {
		st.OpenaiConfig = openaiConfigField
	}
	palmConfigField, err := PaLmConfigFromPb(pb.PalmConfig)
	if err != nil {
		return nil, err
	}
	if palmConfigField != nil {
		st.PalmConfig = palmConfigField
	}
	providerField, err := ExternalModelProviderFromPb(&pb.Provider)
	if err != nil {
		return nil, err
	}
	if providerField != nil {
		st.Provider = *providerField
	}
	st.Task = pb.Task

	return st, nil
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

func ExternalModelProviderToPb(st *ExternalModelProvider) (*servingpb.ExternalModelProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.ExternalModelProviderPb(*st)
	return &pb, nil
}

func ExternalModelProviderFromPb(pb *servingpb.ExternalModelProviderPb) (*ExternalModelProvider, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExternalModelProvider(*pb)
	return &st, nil
}

type ExternalModelUsageElement struct {
	// The number of tokens in the chat/completions response.
	// Wire name: 'completion_tokens'
	CompletionTokens int ``
	// The number of tokens in the prompt.
	// Wire name: 'prompt_tokens'
	PromptTokens int ``
	// The total number of tokens in the prompt and response.
	// Wire name: 'total_tokens'
	TotalTokens     int      ``
	ForceSendFields []string `tf:"-"`
}

func (s *ExternalModelUsageElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ExternalModelUsageElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ExternalModelUsageElementToPb(st *ExternalModelUsageElement) (*servingpb.ExternalModelUsageElementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ExternalModelUsageElementPb{}
	pb.CompletionTokens = st.CompletionTokens
	pb.PromptTokens = st.PromptTokens
	pb.TotalTokens = st.TotalTokens

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ExternalModelUsageElementFromPb(pb *servingpb.ExternalModelUsageElementPb) (*ExternalModelUsageElement, error) {
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

type FallbackConfig struct {
	// Whether to enable traffic fallback. When a served entity in the serving
	// endpoint returns specific error codes (e.g. 500), the request will
	// automatically be round-robin attempted with other served entities in the
	// same endpoint, following the order of served entity list, until a
	// successful response is returned. If all attempts fail, return the last
	// response with the error code.
	// Wire name: 'enabled'
	Enabled bool ``
}

func FallbackConfigToPb(st *FallbackConfig) (*servingpb.FallbackConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.FallbackConfigPb{}
	pb.Enabled = st.Enabled

	return pb, nil
}

func FallbackConfigFromPb(pb *servingpb.FallbackConfigPb) (*FallbackConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FallbackConfig{}
	st.Enabled = pb.Enabled

	return st, nil
}

// All fields are not sensitive as they are hard-coded in the system and made
// available to customers.
type FoundationModel struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'display_name'
	DisplayName string ``

	// Wire name: 'docs'
	Docs string ``

	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *FoundationModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s FoundationModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func FoundationModelToPb(st *FoundationModel) (*servingpb.FoundationModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.FoundationModelPb{}
	pb.Description = st.Description
	pb.DisplayName = st.DisplayName
	pb.Docs = st.Docs
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func FoundationModelFromPb(pb *servingpb.FoundationModelPb) (*FoundationModel, error) {
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

type GetOpenApiRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetOpenApiRequestToPb(st *GetOpenApiRequest) (*servingpb.GetOpenApiRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.GetOpenApiRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetOpenApiRequestFromPb(pb *servingpb.GetOpenApiRequestPb) (*GetOpenApiRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOpenApiRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetOpenApiResponse struct {

	// Wire name: 'contents'
	Contents io.ReadCloser `tf:"-"`
}

func GetOpenApiResponseToPb(st *GetOpenApiResponse) (*servingpb.GetOpenApiResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.GetOpenApiResponsePb{}
	pb.Contents = st.Contents

	return pb, nil
}

func GetOpenApiResponseFromPb(pb *servingpb.GetOpenApiResponsePb) (*GetOpenApiResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOpenApiResponse{}
	st.Contents = pb.Contents

	return st, nil
}

type GetServingEndpointPermissionLevelsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	// Wire name: 'serving_endpoint_id'
	ServingEndpointId string `tf:"-"`
}

func GetServingEndpointPermissionLevelsRequestToPb(st *GetServingEndpointPermissionLevelsRequest) (*servingpb.GetServingEndpointPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.GetServingEndpointPermissionLevelsRequestPb{}
	pb.ServingEndpointId = st.ServingEndpointId

	return pb, nil
}

func GetServingEndpointPermissionLevelsRequestFromPb(pb *servingpb.GetServingEndpointPermissionLevelsRequestPb) (*GetServingEndpointPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointPermissionLevelsRequest{}
	st.ServingEndpointId = pb.ServingEndpointId

	return st, nil
}

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []ServingEndpointPermissionsDescription ``
}

func GetServingEndpointPermissionLevelsResponseToPb(st *GetServingEndpointPermissionLevelsResponse) (*servingpb.GetServingEndpointPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.GetServingEndpointPermissionLevelsResponsePb{}

	var permissionLevelsPb []servingpb.ServingEndpointPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := ServingEndpointPermissionsDescriptionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			permissionLevelsPb = append(permissionLevelsPb, *itemPb)
		}
	}
	pb.PermissionLevels = permissionLevelsPb

	return pb, nil
}

func GetServingEndpointPermissionLevelsResponseFromPb(pb *servingpb.GetServingEndpointPermissionLevelsResponsePb) (*GetServingEndpointPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointPermissionLevelsResponse{}

	var permissionLevelsField []ServingEndpointPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := ServingEndpointPermissionsDescriptionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			permissionLevelsField = append(permissionLevelsField, *item)
		}
	}
	st.PermissionLevels = permissionLevelsField

	return st, nil
}

type GetServingEndpointPermissionsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	// Wire name: 'serving_endpoint_id'
	ServingEndpointId string `tf:"-"`
}

func GetServingEndpointPermissionsRequestToPb(st *GetServingEndpointPermissionsRequest) (*servingpb.GetServingEndpointPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.GetServingEndpointPermissionsRequestPb{}
	pb.ServingEndpointId = st.ServingEndpointId

	return pb, nil
}

func GetServingEndpointPermissionsRequestFromPb(pb *servingpb.GetServingEndpointPermissionsRequestPb) (*GetServingEndpointPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointPermissionsRequest{}
	st.ServingEndpointId = pb.ServingEndpointId

	return st, nil
}

type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func GetServingEndpointRequestToPb(st *GetServingEndpointRequest) (*servingpb.GetServingEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.GetServingEndpointRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetServingEndpointRequestFromPb(pb *servingpb.GetServingEndpointRequestPb) (*GetServingEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointRequest{}
	st.Name = pb.Name

	return st, nil
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
	PrivateKey string ``
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
	PrivateKeyPlaintext string ``
	// This is the Google Cloud project id that the service account is
	// associated with.
	// Wire name: 'project_id'
	ProjectId string ``
	// This is the region for the Google Cloud Vertex AI Service. See [supported
	// regions] for more details. Some models are only available in specific
	// regions.
	//
	// [supported regions]:
	// https://cloud.google.com/vertex-ai/docs/general/locations
	// Wire name: 'region'
	Region          string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *GoogleCloudVertexAiConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s GoogleCloudVertexAiConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func GoogleCloudVertexAiConfigToPb(st *GoogleCloudVertexAiConfig) (*servingpb.GoogleCloudVertexAiConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.GoogleCloudVertexAiConfigPb{}
	pb.PrivateKey = st.PrivateKey
	pb.PrivateKeyPlaintext = st.PrivateKeyPlaintext
	pb.ProjectId = st.ProjectId
	pb.Region = st.Region

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func GoogleCloudVertexAiConfigFromPb(pb *servingpb.GoogleCloudVertexAiConfigPb) (*GoogleCloudVertexAiConfig, error) {
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

type HttpRequestResponse struct {

	// Wire name: 'contents'
	Contents io.ReadCloser `tf:"-"`
}

func HttpRequestResponseToPb(st *HttpRequestResponse) (*servingpb.HttpRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.HttpRequestResponsePb{}
	pb.Contents = st.Contents

	return pb, nil
}

func HttpRequestResponseFromPb(pb *servingpb.HttpRequestResponsePb) (*HttpRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &HttpRequestResponse{}
	st.Contents = pb.Contents

	return st, nil
}

type ListEndpointsResponse struct {
	// The list of endpoints.
	// Wire name: 'endpoints'
	Endpoints []ServingEndpoint ``
}

func ListEndpointsResponseToPb(st *ListEndpointsResponse) (*servingpb.ListEndpointsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ListEndpointsResponsePb{}

	var endpointsPb []servingpb.ServingEndpointPb
	for _, item := range st.Endpoints {
		itemPb, err := ServingEndpointToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			endpointsPb = append(endpointsPb, *itemPb)
		}
	}
	pb.Endpoints = endpointsPb

	return pb, nil
}

func ListEndpointsResponseFromPb(pb *servingpb.ListEndpointsResponsePb) (*ListEndpointsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListEndpointsResponse{}

	var endpointsField []ServingEndpoint
	for _, itemPb := range pb.Endpoints {
		item, err := ServingEndpointFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			endpointsField = append(endpointsField, *item)
		}
	}
	st.Endpoints = endpointsField

	return st, nil
}

type LogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	// Wire name: 'name'
	Name string `tf:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	// Wire name: 'served_model_name'
	ServedModelName string `tf:"-"`
}

func LogsRequestToPb(st *LogsRequest) (*servingpb.LogsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.LogsRequestPb{}
	pb.Name = st.Name
	pb.ServedModelName = st.ServedModelName

	return pb, nil
}

func LogsRequestFromPb(pb *servingpb.LogsRequestPb) (*LogsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogsRequest{}
	st.Name = pb.Name
	st.ServedModelName = pb.ServedModelName

	return st, nil
}

// A representation of all DataPlaneInfo for operations that can be done on a
// model through Data Plane APIs.
type ModelDataPlaneInfo struct {
	// Information required to query DataPlane API 'query' endpoint.
	// Wire name: 'query_info'
	QueryInfo *DataPlaneInfo ``
}

func ModelDataPlaneInfoToPb(st *ModelDataPlaneInfo) (*servingpb.ModelDataPlaneInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ModelDataPlaneInfoPb{}
	queryInfoPb, err := DataPlaneInfoToPb(st.QueryInfo)
	if err != nil {
		return nil, err
	}
	if queryInfoPb != nil {
		pb.QueryInfo = queryInfoPb
	}

	return pb, nil
}

func ModelDataPlaneInfoFromPb(pb *servingpb.ModelDataPlaneInfoPb) (*ModelDataPlaneInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelDataPlaneInfo{}
	queryInfoField, err := DataPlaneInfoFromPb(pb.QueryInfo)
	if err != nil {
		return nil, err
	}
	if queryInfoField != nil {
		st.QueryInfo = queryInfoField
	}

	return st, nil
}

// Configs needed to create an OpenAI model route.
type OpenAiConfig struct {
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Client ID.
	// Wire name: 'microsoft_entra_client_id'
	MicrosoftEntraClientId string ``
	// The Databricks secret key reference for a client secret used for
	// Microsoft Entra ID authentication. If you prefer to paste your client
	// secret directly, see `microsoft_entra_client_secret_plaintext`. You must
	// provide an API key using one of the following fields:
	// `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	// Wire name: 'microsoft_entra_client_secret'
	MicrosoftEntraClientSecret string ``
	// The client secret used for Microsoft Entra ID authentication provided as
	// a plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `microsoft_entra_client_secret`. You must provide an API key
	// using one of the following fields: `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	// Wire name: 'microsoft_entra_client_secret_plaintext'
	MicrosoftEntraClientSecretPlaintext string ``
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Tenant ID.
	// Wire name: 'microsoft_entra_tenant_id'
	MicrosoftEntraTenantId string ``
	// This is a field to provide a customized base URl for the OpenAI API. For
	// Azure OpenAI, this field is required, and is the base URL for the Azure
	// OpenAI API service provided by Azure. For other OpenAI API types, this
	// field is optional, and if left unspecified, the standard OpenAI base URL
	// is used.
	// Wire name: 'openai_api_base'
	OpenaiApiBase string ``
	// The Databricks secret key reference for an OpenAI API key using the
	// OpenAI or Azure service. If you prefer to paste your API key directly,
	// see `openai_api_key_plaintext`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	// Wire name: 'openai_api_key'
	OpenaiApiKey string ``
	// The OpenAI API key using the OpenAI or Azure service provided as a
	// plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `openai_api_key`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	// Wire name: 'openai_api_key_plaintext'
	OpenaiApiKeyPlaintext string ``
	// This is an optional field to specify the type of OpenAI API to use. For
	// Azure OpenAI, this field is required, and adjust this parameter to
	// represent the preferred security access validation protocol. For access
	// token validation, use azure. For authentication using Azure Active
	// Directory (Azure AD) use, azuread.
	// Wire name: 'openai_api_type'
	OpenaiApiType string ``
	// This is an optional field to specify the OpenAI API version. For Azure
	// OpenAI, this field is required, and is the version of the Azure OpenAI
	// service to utilize, specified by a date.
	// Wire name: 'openai_api_version'
	OpenaiApiVersion string ``
	// This field is only required for Azure OpenAI and is the name of the
	// deployment resource for the Azure OpenAI service.
	// Wire name: 'openai_deployment_name'
	OpenaiDeploymentName string ``
	// This is an optional field to specify the organization in OpenAI or Azure
	// OpenAI.
	// Wire name: 'openai_organization'
	OpenaiOrganization string   ``
	ForceSendFields    []string `tf:"-"`
}

func (s *OpenAiConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s OpenAiConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func OpenAiConfigToPb(st *OpenAiConfig) (*servingpb.OpenAiConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.OpenAiConfigPb{}
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

func OpenAiConfigFromPb(pb *servingpb.OpenAiConfigPb) (*OpenAiConfig, error) {
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

type PaLmConfig struct {
	// The Databricks secret key reference for a PaLM API key. If you prefer to
	// paste your API key directly, see `palm_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	// Wire name: 'palm_api_key'
	PalmApiKey string ``
	// The PaLM API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `palm_api_key`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	// Wire name: 'palm_api_key_plaintext'
	PalmApiKeyPlaintext string   ``
	ForceSendFields     []string `tf:"-"`
}

func (s *PaLmConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PaLmConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PaLmConfigToPb(st *PaLmConfig) (*servingpb.PaLmConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.PaLmConfigPb{}
	pb.PalmApiKey = st.PalmApiKey
	pb.PalmApiKeyPlaintext = st.PalmApiKeyPlaintext

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PaLmConfigFromPb(pb *servingpb.PaLmConfigPb) (*PaLmConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PaLmConfig{}
	st.PalmApiKey = pb.PalmApiKey
	st.PalmApiKeyPlaintext = pb.PalmApiKeyPlaintext

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	// Wire name: 'add_tags'
	AddTags []EndpointTag ``
	// List of tag keys to delete
	// Wire name: 'delete_tags'
	DeleteTags []string ``
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func PatchServingEndpointTagsToPb(st *PatchServingEndpointTags) (*servingpb.PatchServingEndpointTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.PatchServingEndpointTagsPb{}

	var addTagsPb []servingpb.EndpointTagPb
	for _, item := range st.AddTags {
		itemPb, err := EndpointTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			addTagsPb = append(addTagsPb, *itemPb)
		}
	}
	pb.AddTags = addTagsPb
	pb.DeleteTags = st.DeleteTags
	pb.Name = st.Name

	return pb, nil
}

func PatchServingEndpointTagsFromPb(pb *servingpb.PatchServingEndpointTagsPb) (*PatchServingEndpointTags, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PatchServingEndpointTags{}

	var addTagsField []EndpointTag
	for _, itemPb := range pb.AddTags {
		item, err := EndpointTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			addTagsField = append(addTagsField, *item)
		}
	}
	st.AddTags = addTagsField
	st.DeleteTags = pb.DeleteTags
	st.Name = pb.Name

	return st, nil
}

type PayloadTable struct {

	// Wire name: 'name'
	Name string ``

	// Wire name: 'status'
	Status string ``

	// Wire name: 'status_message'
	StatusMessage   string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *PayloadTable) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PayloadTable) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PayloadTableToPb(st *PayloadTable) (*servingpb.PayloadTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.PayloadTablePb{}
	pb.Name = st.Name
	pb.Status = st.Status
	pb.StatusMessage = st.StatusMessage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PayloadTableFromPb(pb *servingpb.PayloadTablePb) (*PayloadTable, error) {
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

type PtEndpointCoreConfig struct {
	// The list of served entities under the serving endpoint config.
	// Wire name: 'served_entities'
	ServedEntities []PtServedModel ``

	// Wire name: 'traffic_config'
	TrafficConfig *TrafficConfig ``
}

func PtEndpointCoreConfigToPb(st *PtEndpointCoreConfig) (*servingpb.PtEndpointCoreConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.PtEndpointCoreConfigPb{}

	var servedEntitiesPb []servingpb.PtServedModelPb
	for _, item := range st.ServedEntities {
		itemPb, err := PtServedModelToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedEntitiesPb = append(servedEntitiesPb, *itemPb)
		}
	}
	pb.ServedEntities = servedEntitiesPb
	trafficConfigPb, err := TrafficConfigToPb(st.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigPb != nil {
		pb.TrafficConfig = trafficConfigPb
	}

	return pb, nil
}

func PtEndpointCoreConfigFromPb(pb *servingpb.PtEndpointCoreConfigPb) (*PtEndpointCoreConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PtEndpointCoreConfig{}

	var servedEntitiesField []PtServedModel
	for _, itemPb := range pb.ServedEntities {
		item, err := PtServedModelFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedEntitiesField = append(servedEntitiesField, *item)
		}
	}
	st.ServedEntities = servedEntitiesField
	trafficConfigField, err := TrafficConfigFromPb(pb.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigField != nil {
		st.TrafficConfig = trafficConfigField
	}

	return st, nil
}

type PtServedModel struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	// Wire name: 'entity_name'
	EntityName string ``

	// Wire name: 'entity_version'
	EntityVersion string ``
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	// Wire name: 'name'
	Name string ``
	// The number of model units to be provisioned.
	// Wire name: 'provisioned_model_units'
	ProvisionedModelUnits int64    ``
	ForceSendFields       []string `tf:"-"`
}

func (s *PtServedModel) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s PtServedModel) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func PtServedModelToPb(st *PtServedModel) (*servingpb.PtServedModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.PtServedModelPb{}
	pb.EntityName = st.EntityName
	pb.EntityVersion = st.EntityVersion
	pb.Name = st.Name
	pb.ProvisionedModelUnits = st.ProvisionedModelUnits

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func PtServedModelFromPb(pb *servingpb.PtServedModelPb) (*PtServedModel, error) {
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

type PutAiGatewayRequest struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	// Wire name: 'fallback_config'
	FallbackConfig *FallbackConfig ``
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	// Wire name: 'guardrails'
	Guardrails *AiGatewayGuardrails ``
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	// Wire name: 'inference_table_config'
	InferenceTableConfig *AiGatewayInferenceTableConfig ``
	// The name of the serving endpoint whose AI Gateway is being updated. This
	// field is required.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	// Wire name: 'rate_limits'
	RateLimits []AiGatewayRateLimit ``
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	// Wire name: 'usage_tracking_config'
	UsageTrackingConfig *AiGatewayUsageTrackingConfig ``
}

func PutAiGatewayRequestToPb(st *PutAiGatewayRequest) (*servingpb.PutAiGatewayRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.PutAiGatewayRequestPb{}
	fallbackConfigPb, err := FallbackConfigToPb(st.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigPb != nil {
		pb.FallbackConfig = fallbackConfigPb
	}
	guardrailsPb, err := AiGatewayGuardrailsToPb(st.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsPb != nil {
		pb.Guardrails = guardrailsPb
	}
	inferenceTableConfigPb, err := AiGatewayInferenceTableConfigToPb(st.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigPb != nil {
		pb.InferenceTableConfig = inferenceTableConfigPb
	}
	pb.Name = st.Name

	var rateLimitsPb []servingpb.AiGatewayRateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := AiGatewayRateLimitToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rateLimitsPb = append(rateLimitsPb, *itemPb)
		}
	}
	pb.RateLimits = rateLimitsPb
	usageTrackingConfigPb, err := AiGatewayUsageTrackingConfigToPb(st.UsageTrackingConfig)
	if err != nil {
		return nil, err
	}
	if usageTrackingConfigPb != nil {
		pb.UsageTrackingConfig = usageTrackingConfigPb
	}

	return pb, nil
}

func PutAiGatewayRequestFromPb(pb *servingpb.PutAiGatewayRequestPb) (*PutAiGatewayRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutAiGatewayRequest{}
	fallbackConfigField, err := FallbackConfigFromPb(pb.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigField != nil {
		st.FallbackConfig = fallbackConfigField
	}
	guardrailsField, err := AiGatewayGuardrailsFromPb(pb.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsField != nil {
		st.Guardrails = guardrailsField
	}
	inferenceTableConfigField, err := AiGatewayInferenceTableConfigFromPb(pb.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigField != nil {
		st.InferenceTableConfig = inferenceTableConfigField
	}
	st.Name = pb.Name

	var rateLimitsField []AiGatewayRateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := AiGatewayRateLimitFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rateLimitsField = append(rateLimitsField, *item)
		}
	}
	st.RateLimits = rateLimitsField
	usageTrackingConfigField, err := AiGatewayUsageTrackingConfigFromPb(pb.UsageTrackingConfig)
	if err != nil {
		return nil, err
	}
	if usageTrackingConfigField != nil {
		st.UsageTrackingConfig = usageTrackingConfigField
	}

	return st, nil
}

type PutAiGatewayResponse struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	// Wire name: 'fallback_config'
	FallbackConfig *FallbackConfig ``
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	// Wire name: 'guardrails'
	Guardrails *AiGatewayGuardrails ``
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	// Wire name: 'inference_table_config'
	InferenceTableConfig *AiGatewayInferenceTableConfig ``
	// Configuration for rate limits which can be set to limit endpoint traffic.
	// Wire name: 'rate_limits'
	RateLimits []AiGatewayRateLimit ``
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	// Wire name: 'usage_tracking_config'
	UsageTrackingConfig *AiGatewayUsageTrackingConfig ``
}

func PutAiGatewayResponseToPb(st *PutAiGatewayResponse) (*servingpb.PutAiGatewayResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.PutAiGatewayResponsePb{}
	fallbackConfigPb, err := FallbackConfigToPb(st.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigPb != nil {
		pb.FallbackConfig = fallbackConfigPb
	}
	guardrailsPb, err := AiGatewayGuardrailsToPb(st.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsPb != nil {
		pb.Guardrails = guardrailsPb
	}
	inferenceTableConfigPb, err := AiGatewayInferenceTableConfigToPb(st.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigPb != nil {
		pb.InferenceTableConfig = inferenceTableConfigPb
	}

	var rateLimitsPb []servingpb.AiGatewayRateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := AiGatewayRateLimitToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rateLimitsPb = append(rateLimitsPb, *itemPb)
		}
	}
	pb.RateLimits = rateLimitsPb
	usageTrackingConfigPb, err := AiGatewayUsageTrackingConfigToPb(st.UsageTrackingConfig)
	if err != nil {
		return nil, err
	}
	if usageTrackingConfigPb != nil {
		pb.UsageTrackingConfig = usageTrackingConfigPb
	}

	return pb, nil
}

func PutAiGatewayResponseFromPb(pb *servingpb.PutAiGatewayResponsePb) (*PutAiGatewayResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutAiGatewayResponse{}
	fallbackConfigField, err := FallbackConfigFromPb(pb.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigField != nil {
		st.FallbackConfig = fallbackConfigField
	}
	guardrailsField, err := AiGatewayGuardrailsFromPb(pb.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsField != nil {
		st.Guardrails = guardrailsField
	}
	inferenceTableConfigField, err := AiGatewayInferenceTableConfigFromPb(pb.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigField != nil {
		st.InferenceTableConfig = inferenceTableConfigField
	}

	var rateLimitsField []AiGatewayRateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := AiGatewayRateLimitFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rateLimitsField = append(rateLimitsField, *item)
		}
	}
	st.RateLimits = rateLimitsField
	usageTrackingConfigField, err := AiGatewayUsageTrackingConfigFromPb(pb.UsageTrackingConfig)
	if err != nil {
		return nil, err
	}
	if usageTrackingConfigField != nil {
		st.UsageTrackingConfig = usageTrackingConfigField
	}

	return st, nil
}

type PutRequest struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	// Wire name: 'name'
	Name string `tf:"-"`
	// The list of endpoint rate limits.
	// Wire name: 'rate_limits'
	RateLimits []RateLimit ``
}

func PutRequestToPb(st *PutRequest) (*servingpb.PutRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.PutRequestPb{}
	pb.Name = st.Name

	var rateLimitsPb []servingpb.RateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := RateLimitToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rateLimitsPb = append(rateLimitsPb, *itemPb)
		}
	}
	pb.RateLimits = rateLimitsPb

	return pb, nil
}

func PutRequestFromPb(pb *servingpb.PutRequestPb) (*PutRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutRequest{}
	st.Name = pb.Name

	var rateLimitsField []RateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := RateLimitFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rateLimitsField = append(rateLimitsField, *item)
		}
	}
	st.RateLimits = rateLimitsField

	return st, nil
}

type PutResponse struct {
	// The list of endpoint rate limits.
	// Wire name: 'rate_limits'
	RateLimits []RateLimit ``
}

func PutResponseToPb(st *PutResponse) (*servingpb.PutResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.PutResponsePb{}

	var rateLimitsPb []servingpb.RateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := RateLimitToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rateLimitsPb = append(rateLimitsPb, *itemPb)
		}
	}
	pb.RateLimits = rateLimitsPb

	return pb, nil
}

func PutResponseFromPb(pb *servingpb.PutResponsePb) (*PutResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutResponse{}

	var rateLimitsField []RateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := RateLimitFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rateLimitsField = append(rateLimitsField, *item)
		}
	}
	st.RateLimits = rateLimitsField

	return st, nil
}

type QueryEndpointInput struct {
	// Pandas Dataframe input in the records orientation.
	// Wire name: 'dataframe_records'
	DataframeRecords []any ``
	// Pandas Dataframe input in the split orientation.
	// Wire name: 'dataframe_split'
	DataframeSplit *DataframeSplitInput ``
	// The extra parameters field used ONLY for __completions, chat,__ and
	// __embeddings external & foundation model__ serving endpoints. This is a
	// map of strings and should only be used with other external/foundation
	// model query fields.
	// Wire name: 'extra_params'
	ExtraParams map[string]string ``
	// The input string (or array of strings) field used ONLY for __embeddings
	// external & foundation model__ serving endpoints and is the only field
	// (along with extra_params if needed) used by embeddings queries.
	// Wire name: 'input'
	Input any ``
	// Tensor-based input in columnar format.
	// Wire name: 'inputs'
	Inputs any ``
	// Tensor-based input in row format.
	// Wire name: 'instances'
	Instances []any ``
	// The max tokens field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is an integer and should only
	// be used with other chat/completions query fields.
	// Wire name: 'max_tokens'
	MaxTokens int ``
	// The messages field used ONLY for __chat external & foundation model__
	// serving endpoints. This is a map of strings and should only be used with
	// other chat query fields.
	// Wire name: 'messages'
	Messages []ChatMessage ``
	// The n (number of candidates) field used ONLY for __completions__ and
	// __chat external & foundation model__ serving endpoints. This is an
	// integer between 1 and 5 with a default of 1 and should only be used with
	// other chat/completions query fields.
	// Wire name: 'n'
	N int ``
	// The name of the serving endpoint. This field is required.
	// Wire name: 'name'
	Name string `tf:"-"`
	// The prompt string (or array of strings) field used ONLY for __completions
	// external & foundation model__ serving endpoints and should only be used
	// with other completions query fields.
	// Wire name: 'prompt'
	Prompt any ``
	// The stop sequences field used ONLY for __completions__ and __chat
	// external & foundation model__ serving endpoints. This is a list of
	// strings and should only be used with other chat/completions query fields.
	// Wire name: 'stop'
	Stop []string ``
	// The stream field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a boolean defaulting to
	// false and should only be used with other chat/completions query fields.
	// Wire name: 'stream'
	Stream bool ``
	// The temperature field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a float between 0.0 and 2.0
	// with a default of 1.0 and should only be used with other chat/completions
	// query fields.
	// Wire name: 'temperature'
	Temperature     float64  ``
	ForceSendFields []string `tf:"-"`
}

func (s *QueryEndpointInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryEndpointInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryEndpointInputToPb(st *QueryEndpointInput) (*servingpb.QueryEndpointInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.QueryEndpointInputPb{}
	pb.DataframeRecords = st.DataframeRecords
	dataframeSplitPb, err := DataframeSplitInputToPb(st.DataframeSplit)
	if err != nil {
		return nil, err
	}
	if dataframeSplitPb != nil {
		pb.DataframeSplit = dataframeSplitPb
	}
	pb.ExtraParams = st.ExtraParams
	pb.Input = st.Input
	pb.Inputs = st.Inputs
	pb.Instances = st.Instances
	pb.MaxTokens = st.MaxTokens

	var messagesPb []servingpb.ChatMessagePb
	for _, item := range st.Messages {
		itemPb, err := ChatMessageToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			messagesPb = append(messagesPb, *itemPb)
		}
	}
	pb.Messages = messagesPb
	pb.N = st.N
	pb.Name = st.Name
	pb.Prompt = st.Prompt
	pb.Stop = st.Stop
	pb.Stream = st.Stream
	pb.Temperature = st.Temperature

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func QueryEndpointInputFromPb(pb *servingpb.QueryEndpointInputPb) (*QueryEndpointInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryEndpointInput{}
	st.DataframeRecords = pb.DataframeRecords
	dataframeSplitField, err := DataframeSplitInputFromPb(pb.DataframeSplit)
	if err != nil {
		return nil, err
	}
	if dataframeSplitField != nil {
		st.DataframeSplit = dataframeSplitField
	}
	st.ExtraParams = pb.ExtraParams
	st.Input = pb.Input
	st.Inputs = pb.Inputs
	st.Instances = pb.Instances
	st.MaxTokens = pb.MaxTokens

	var messagesField []ChatMessage
	for _, itemPb := range pb.Messages {
		item, err := ChatMessageFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			messagesField = append(messagesField, *item)
		}
	}
	st.Messages = messagesField
	st.N = pb.N
	st.Name = pb.Name
	st.Prompt = pb.Prompt
	st.Stop = pb.Stop
	st.Stream = pb.Stream
	st.Temperature = pb.Temperature

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type QueryEndpointResponse struct {
	// The list of choices returned by the __chat or completions
	// external/foundation model__ serving endpoint.
	// Wire name: 'choices'
	Choices []V1ResponseChoiceElement ``
	// The timestamp in seconds when the query was created in Unix time returned
	// by a __completions or chat external/foundation model__ serving endpoint.
	// Wire name: 'created'
	Created int64 ``
	// The list of the embeddings returned by the __embeddings
	// external/foundation model__ serving endpoint.
	// Wire name: 'data'
	Data []EmbeddingsV1ResponseEmbeddingElement ``
	// The ID of the query that may be returned by a __completions or chat
	// external/foundation model__ serving endpoint.
	// Wire name: 'id'
	Id string ``
	// The name of the __external/foundation model__ used for querying. This is
	// the name of the model that was specified in the endpoint config.
	// Wire name: 'model'
	Model string ``
	// The type of object returned by the __external/foundation model__ serving
	// endpoint, one of [text_completion, chat.completion, list (of
	// embeddings)].
	// Wire name: 'object'
	Object QueryEndpointResponseObject ``
	// The predictions returned by the serving endpoint.
	// Wire name: 'predictions'
	Predictions []any ``
	// The name of the served model that served the request. This is useful when
	// there are multiple models behind the same endpoint with traffic split.
	// Wire name: 'served-model-name'
	ServedModelName string `tf:"-"`
	// The usage object that may be returned by the __external/foundation
	// model__ serving endpoint. This contains information about the number of
	// tokens used in the prompt and response.
	// Wire name: 'usage'
	Usage           *ExternalModelUsageElement ``
	ForceSendFields []string                   `tf:"-"`
}

func (s *QueryEndpointResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s QueryEndpointResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func QueryEndpointResponseToPb(st *QueryEndpointResponse) (*servingpb.QueryEndpointResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.QueryEndpointResponsePb{}

	var choicesPb []servingpb.V1ResponseChoiceElementPb
	for _, item := range st.Choices {
		itemPb, err := V1ResponseChoiceElementToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			choicesPb = append(choicesPb, *itemPb)
		}
	}
	pb.Choices = choicesPb
	pb.Created = st.Created

	var dataPb []servingpb.EmbeddingsV1ResponseEmbeddingElementPb
	for _, item := range st.Data {
		itemPb, err := EmbeddingsV1ResponseEmbeddingElementToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataPb = append(dataPb, *itemPb)
		}
	}
	pb.Data = dataPb
	pb.Id = st.Id
	pb.Model = st.Model
	objectPb, err := QueryEndpointResponseObjectToPb(&st.Object)
	if err != nil {
		return nil, err
	}
	if objectPb != nil {
		pb.Object = *objectPb
	}
	pb.Predictions = st.Predictions
	pb.ServedModelName = st.ServedModelName
	usagePb, err := ExternalModelUsageElementToPb(st.Usage)
	if err != nil {
		return nil, err
	}
	if usagePb != nil {
		pb.Usage = usagePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func QueryEndpointResponseFromPb(pb *servingpb.QueryEndpointResponsePb) (*QueryEndpointResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryEndpointResponse{}

	var choicesField []V1ResponseChoiceElement
	for _, itemPb := range pb.Choices {
		item, err := V1ResponseChoiceElementFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			choicesField = append(choicesField, *item)
		}
	}
	st.Choices = choicesField
	st.Created = pb.Created

	var dataField []EmbeddingsV1ResponseEmbeddingElement
	for _, itemPb := range pb.Data {
		item, err := EmbeddingsV1ResponseEmbeddingElementFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dataField = append(dataField, *item)
		}
	}
	st.Data = dataField
	st.Id = pb.Id
	st.Model = pb.Model
	objectField, err := QueryEndpointResponseObjectFromPb(&pb.Object)
	if err != nil {
		return nil, err
	}
	if objectField != nil {
		st.Object = *objectField
	}
	st.Predictions = pb.Predictions
	st.ServedModelName = pb.ServedModelName
	usageField, err := ExternalModelUsageElementFromPb(pb.Usage)
	if err != nil {
		return nil, err
	}
	if usageField != nil {
		st.Usage = usageField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func QueryEndpointResponseObjectToPb(st *QueryEndpointResponseObject) (*servingpb.QueryEndpointResponseObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.QueryEndpointResponseObjectPb(*st)
	return &pb, nil
}

func QueryEndpointResponseObjectFromPb(pb *servingpb.QueryEndpointResponseObjectPb) (*QueryEndpointResponseObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := QueryEndpointResponseObject(*pb)
	return &st, nil
}

type RateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	// Wire name: 'calls'
	Calls int64 ``
	// Key field for a serving endpoint rate limit. Currently, only 'user' and
	// 'endpoint' are supported, with 'endpoint' being the default if not
	// specified.
	// Wire name: 'key'
	Key RateLimitKey ``
	// Renewal period field for a serving endpoint rate limit. Currently, only
	// 'minute' is supported.
	// Wire name: 'renewal_period'
	RenewalPeriod RateLimitRenewalPeriod ``
}

func RateLimitToPb(st *RateLimit) (*servingpb.RateLimitPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.RateLimitPb{}
	pb.Calls = st.Calls
	keyPb, err := RateLimitKeyToPb(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}
	renewalPeriodPb, err := RateLimitRenewalPeriodToPb(&st.RenewalPeriod)
	if err != nil {
		return nil, err
	}
	if renewalPeriodPb != nil {
		pb.RenewalPeriod = *renewalPeriodPb
	}

	return pb, nil
}

func RateLimitFromPb(pb *servingpb.RateLimitPb) (*RateLimit, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RateLimit{}
	st.Calls = pb.Calls
	keyField, err := RateLimitKeyFromPb(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
	renewalPeriodField, err := RateLimitRenewalPeriodFromPb(&pb.RenewalPeriod)
	if err != nil {
		return nil, err
	}
	if renewalPeriodField != nil {
		st.RenewalPeriod = *renewalPeriodField
	}

	return st, nil
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

func RateLimitKeyToPb(st *RateLimitKey) (*servingpb.RateLimitKeyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.RateLimitKeyPb(*st)
	return &pb, nil
}

func RateLimitKeyFromPb(pb *servingpb.RateLimitKeyPb) (*RateLimitKey, error) {
	if pb == nil {
		return nil, nil
	}
	st := RateLimitKey(*pb)
	return &st, nil
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

func RateLimitRenewalPeriodToPb(st *RateLimitRenewalPeriod) (*servingpb.RateLimitRenewalPeriodPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.RateLimitRenewalPeriodPb(*st)
	return &pb, nil
}

func RateLimitRenewalPeriodFromPb(pb *servingpb.RateLimitRenewalPeriodPb) (*RateLimitRenewalPeriod, error) {
	if pb == nil {
		return nil, nil
	}
	st := RateLimitRenewalPeriod(*pb)
	return &st, nil
}

type Route struct {

	// Wire name: 'served_entity_name'
	ServedEntityName string ``
	// The name of the served model this route configures traffic for.
	// Wire name: 'served_model_name'
	ServedModelName string ``
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	// Wire name: 'traffic_percentage'
	TrafficPercentage int      ``
	ForceSendFields   []string `tf:"-"`
}

func (s *Route) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Route) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func RouteToPb(st *Route) (*servingpb.RoutePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.RoutePb{}
	pb.ServedEntityName = st.ServedEntityName
	pb.ServedModelName = st.ServedModelName
	pb.TrafficPercentage = st.TrafficPercentage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func RouteFromPb(pb *servingpb.RoutePb) (*Route, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Route{}
	st.ServedEntityName = pb.ServedEntityName
	st.ServedModelName = pb.ServedModelName
	st.TrafficPercentage = pb.TrafficPercentage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ServedEntityInput struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	// Wire name: 'entity_name'
	EntityName string ``

	// Wire name: 'entity_version'
	EntityVersion string ``
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	// Wire name: 'environment_vars'
	EnvironmentVars map[string]string ``
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	// Wire name: 'external_model'
	ExternalModel *ExternalModel ``
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	// Wire name: 'max_provisioned_concurrency'
	MaxProvisionedConcurrency int ``
	// The maximum tokens per second that the endpoint can scale up to.
	// Wire name: 'max_provisioned_throughput'
	MaxProvisionedThroughput int ``
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	// Wire name: 'min_provisioned_concurrency'
	MinProvisionedConcurrency int ``
	// The minimum tokens per second that the endpoint can scale down to.
	// Wire name: 'min_provisioned_throughput'
	MinProvisionedThroughput int ``
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	// Wire name: 'name'
	Name string ``
	// The number of model units provisioned.
	// Wire name: 'provisioned_model_units'
	ProvisionedModelUnits int64 ``
	// Whether the compute resources for the served entity should scale down to
	// zero.
	// Wire name: 'scale_to_zero_enabled'
	ScaleToZeroEnabled bool ``
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
	WorkloadSize string ``
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	// Wire name: 'workload_type'
	WorkloadType    ServingModelWorkloadType ``
	ForceSendFields []string                 `tf:"-"`
}

func (s *ServedEntityInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntityInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServedEntityInputToPb(st *ServedEntityInput) (*servingpb.ServedEntityInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServedEntityInputPb{}
	pb.EntityName = st.EntityName
	pb.EntityVersion = st.EntityVersion
	pb.EnvironmentVars = st.EnvironmentVars
	externalModelPb, err := ExternalModelToPb(st.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelPb != nil {
		pb.ExternalModel = externalModelPb
	}
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.MaxProvisionedConcurrency = st.MaxProvisionedConcurrency
	pb.MaxProvisionedThroughput = st.MaxProvisionedThroughput
	pb.MinProvisionedConcurrency = st.MinProvisionedConcurrency
	pb.MinProvisionedThroughput = st.MinProvisionedThroughput
	pb.Name = st.Name
	pb.ProvisionedModelUnits = st.ProvisionedModelUnits
	pb.ScaleToZeroEnabled = st.ScaleToZeroEnabled
	pb.WorkloadSize = st.WorkloadSize
	workloadTypePb, err := ServingModelWorkloadTypeToPb(&st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = *workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServedEntityInputFromPb(pb *servingpb.ServedEntityInputPb) (*ServedEntityInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedEntityInput{}
	st.EntityName = pb.EntityName
	st.EntityVersion = pb.EntityVersion
	st.EnvironmentVars = pb.EnvironmentVars
	externalModelField, err := ExternalModelFromPb(pb.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelField != nil {
		st.ExternalModel = externalModelField
	}
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.MaxProvisionedConcurrency = pb.MaxProvisionedConcurrency
	st.MaxProvisionedThroughput = pb.MaxProvisionedThroughput
	st.MinProvisionedConcurrency = pb.MinProvisionedConcurrency
	st.MinProvisionedThroughput = pb.MinProvisionedThroughput
	st.Name = pb.Name
	st.ProvisionedModelUnits = pb.ProvisionedModelUnits
	st.ScaleToZeroEnabled = pb.ScaleToZeroEnabled
	st.WorkloadSize = pb.WorkloadSize
	workloadTypeField, err := ServingModelWorkloadTypeFromPb(&pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = *workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ServedEntityOutput struct {

	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``

	// Wire name: 'creator'
	Creator string ``
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	// Wire name: 'entity_name'
	EntityName string ``

	// Wire name: 'entity_version'
	EntityVersion string ``
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	// Wire name: 'environment_vars'
	EnvironmentVars map[string]string ``
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	// Wire name: 'external_model'
	ExternalModel *ExternalModel ``

	// Wire name: 'foundation_model'
	FoundationModel *FoundationModel ``
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	// Wire name: 'max_provisioned_concurrency'
	MaxProvisionedConcurrency int ``
	// The maximum tokens per second that the endpoint can scale up to.
	// Wire name: 'max_provisioned_throughput'
	MaxProvisionedThroughput int ``
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	// Wire name: 'min_provisioned_concurrency'
	MinProvisionedConcurrency int ``
	// The minimum tokens per second that the endpoint can scale down to.
	// Wire name: 'min_provisioned_throughput'
	MinProvisionedThroughput int ``
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	// Wire name: 'name'
	Name string ``
	// The number of model units provisioned.
	// Wire name: 'provisioned_model_units'
	ProvisionedModelUnits int64 ``
	// Whether the compute resources for the served entity should scale down to
	// zero.
	// Wire name: 'scale_to_zero_enabled'
	ScaleToZeroEnabled bool ``

	// Wire name: 'state'
	State *ServedModelState ``
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
	WorkloadSize string ``
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	// Wire name: 'workload_type'
	WorkloadType    ServingModelWorkloadType ``
	ForceSendFields []string                 `tf:"-"`
}

func (s *ServedEntityOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntityOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServedEntityOutputToPb(st *ServedEntityOutput) (*servingpb.ServedEntityOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServedEntityOutputPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Creator = st.Creator
	pb.EntityName = st.EntityName
	pb.EntityVersion = st.EntityVersion
	pb.EnvironmentVars = st.EnvironmentVars
	externalModelPb, err := ExternalModelToPb(st.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelPb != nil {
		pb.ExternalModel = externalModelPb
	}
	foundationModelPb, err := FoundationModelToPb(st.FoundationModel)
	if err != nil {
		return nil, err
	}
	if foundationModelPb != nil {
		pb.FoundationModel = foundationModelPb
	}
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.MaxProvisionedConcurrency = st.MaxProvisionedConcurrency
	pb.MaxProvisionedThroughput = st.MaxProvisionedThroughput
	pb.MinProvisionedConcurrency = st.MinProvisionedConcurrency
	pb.MinProvisionedThroughput = st.MinProvisionedThroughput
	pb.Name = st.Name
	pb.ProvisionedModelUnits = st.ProvisionedModelUnits
	pb.ScaleToZeroEnabled = st.ScaleToZeroEnabled
	statePb, err := ServedModelStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}
	pb.WorkloadSize = st.WorkloadSize
	workloadTypePb, err := ServingModelWorkloadTypeToPb(&st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = *workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServedEntityOutputFromPb(pb *servingpb.ServedEntityOutputPb) (*ServedEntityOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedEntityOutput{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator
	st.EntityName = pb.EntityName
	st.EntityVersion = pb.EntityVersion
	st.EnvironmentVars = pb.EnvironmentVars
	externalModelField, err := ExternalModelFromPb(pb.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelField != nil {
		st.ExternalModel = externalModelField
	}
	foundationModelField, err := FoundationModelFromPb(pb.FoundationModel)
	if err != nil {
		return nil, err
	}
	if foundationModelField != nil {
		st.FoundationModel = foundationModelField
	}
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.MaxProvisionedConcurrency = pb.MaxProvisionedConcurrency
	st.MaxProvisionedThroughput = pb.MaxProvisionedThroughput
	st.MinProvisionedConcurrency = pb.MinProvisionedConcurrency
	st.MinProvisionedThroughput = pb.MinProvisionedThroughput
	st.Name = pb.Name
	st.ProvisionedModelUnits = pb.ProvisionedModelUnits
	st.ScaleToZeroEnabled = pb.ScaleToZeroEnabled
	stateField, err := ServedModelStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	st.WorkloadSize = pb.WorkloadSize
	workloadTypeField, err := ServingModelWorkloadTypeFromPb(&pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = *workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ServedEntitySpec struct {

	// Wire name: 'entity_name'
	EntityName string ``

	// Wire name: 'entity_version'
	EntityVersion string ``

	// Wire name: 'external_model'
	ExternalModel *ExternalModel ``

	// Wire name: 'foundation_model'
	FoundationModel *FoundationModel ``

	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ServedEntitySpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedEntitySpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServedEntitySpecToPb(st *ServedEntitySpec) (*servingpb.ServedEntitySpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServedEntitySpecPb{}
	pb.EntityName = st.EntityName
	pb.EntityVersion = st.EntityVersion
	externalModelPb, err := ExternalModelToPb(st.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelPb != nil {
		pb.ExternalModel = externalModelPb
	}
	foundationModelPb, err := FoundationModelToPb(st.FoundationModel)
	if err != nil {
		return nil, err
	}
	if foundationModelPb != nil {
		pb.FoundationModel = foundationModelPb
	}
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServedEntitySpecFromPb(pb *servingpb.ServedEntitySpecPb) (*ServedEntitySpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedEntitySpec{}
	st.EntityName = pb.EntityName
	st.EntityVersion = pb.EntityVersion
	externalModelField, err := ExternalModelFromPb(pb.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelField != nil {
		st.ExternalModel = externalModelField
	}
	foundationModelField, err := FoundationModelFromPb(pb.FoundationModel)
	if err != nil {
		return nil, err
	}
	if foundationModelField != nil {
		st.FoundationModel = foundationModelField
	}
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ServedModelInput struct {
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	// Wire name: 'environment_vars'
	EnvironmentVars map[string]string ``
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	// Wire name: 'max_provisioned_concurrency'
	MaxProvisionedConcurrency int ``
	// The maximum tokens per second that the endpoint can scale up to.
	// Wire name: 'max_provisioned_throughput'
	MaxProvisionedThroughput int ``
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	// Wire name: 'min_provisioned_concurrency'
	MinProvisionedConcurrency int ``
	// The minimum tokens per second that the endpoint can scale down to.
	// Wire name: 'min_provisioned_throughput'
	MinProvisionedThroughput int ``

	// Wire name: 'model_name'
	ModelName string ``

	// Wire name: 'model_version'
	ModelVersion string ``
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	// Wire name: 'name'
	Name string ``
	// The number of model units provisioned.
	// Wire name: 'provisioned_model_units'
	ProvisionedModelUnits int64 ``
	// Whether the compute resources for the served entity should scale down to
	// zero.
	// Wire name: 'scale_to_zero_enabled'
	ScaleToZeroEnabled bool ``
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
	WorkloadSize string ``
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	// Wire name: 'workload_type'
	WorkloadType    ServedModelInputWorkloadType ``
	ForceSendFields []string                     `tf:"-"`
}

func (s *ServedModelInput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelInput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServedModelInputToPb(st *ServedModelInput) (*servingpb.ServedModelInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServedModelInputPb{}
	pb.EnvironmentVars = st.EnvironmentVars
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.MaxProvisionedConcurrency = st.MaxProvisionedConcurrency
	pb.MaxProvisionedThroughput = st.MaxProvisionedThroughput
	pb.MinProvisionedConcurrency = st.MinProvisionedConcurrency
	pb.MinProvisionedThroughput = st.MinProvisionedThroughput
	pb.ModelName = st.ModelName
	pb.ModelVersion = st.ModelVersion
	pb.Name = st.Name
	pb.ProvisionedModelUnits = st.ProvisionedModelUnits
	pb.ScaleToZeroEnabled = st.ScaleToZeroEnabled
	pb.WorkloadSize = st.WorkloadSize
	workloadTypePb, err := ServedModelInputWorkloadTypeToPb(&st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = *workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServedModelInputFromPb(pb *servingpb.ServedModelInputPb) (*ServedModelInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedModelInput{}
	st.EnvironmentVars = pb.EnvironmentVars
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.MaxProvisionedConcurrency = pb.MaxProvisionedConcurrency
	st.MaxProvisionedThroughput = pb.MaxProvisionedThroughput
	st.MinProvisionedConcurrency = pb.MinProvisionedConcurrency
	st.MinProvisionedThroughput = pb.MinProvisionedThroughput
	st.ModelName = pb.ModelName
	st.ModelVersion = pb.ModelVersion
	st.Name = pb.Name
	st.ProvisionedModelUnits = pb.ProvisionedModelUnits
	st.ScaleToZeroEnabled = pb.ScaleToZeroEnabled
	st.WorkloadSize = pb.WorkloadSize
	workloadTypeField, err := ServedModelInputWorkloadTypeFromPb(&pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = *workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ServedModelInputWorkloadTypeToPb(st *ServedModelInputWorkloadType) (*servingpb.ServedModelInputWorkloadTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.ServedModelInputWorkloadTypePb(*st)
	return &pb, nil
}

func ServedModelInputWorkloadTypeFromPb(pb *servingpb.ServedModelInputWorkloadTypePb) (*ServedModelInputWorkloadType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServedModelInputWorkloadType(*pb)
	return &st, nil
}

type ServedModelOutput struct {

	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``

	// Wire name: 'creator'
	Creator string ``
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	// Wire name: 'environment_vars'
	EnvironmentVars map[string]string ``
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	// Wire name: 'instance_profile_arn'
	InstanceProfileArn string ``
	// The maximum provisioned concurrency that the endpoint can scale up to. Do
	// not use if workload_size is specified.
	// Wire name: 'max_provisioned_concurrency'
	MaxProvisionedConcurrency int ``
	// The minimum provisioned concurrency that the endpoint can scale down to.
	// Do not use if workload_size is specified.
	// Wire name: 'min_provisioned_concurrency'
	MinProvisionedConcurrency int ``

	// Wire name: 'model_name'
	ModelName string ``

	// Wire name: 'model_version'
	ModelVersion string ``
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	// Wire name: 'name'
	Name string ``
	// The number of model units provisioned.
	// Wire name: 'provisioned_model_units'
	ProvisionedModelUnits int64 ``
	// Whether the compute resources for the served entity should scale down to
	// zero.
	// Wire name: 'scale_to_zero_enabled'
	ScaleToZeroEnabled bool ``

	// Wire name: 'state'
	State *ServedModelState ``
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
	WorkloadSize string ``
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	// Wire name: 'workload_type'
	WorkloadType    ServingModelWorkloadType ``
	ForceSendFields []string                 `tf:"-"`
}

func (s *ServedModelOutput) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelOutput) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServedModelOutputToPb(st *ServedModelOutput) (*servingpb.ServedModelOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServedModelOutputPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Creator = st.Creator
	pb.EnvironmentVars = st.EnvironmentVars
	pb.InstanceProfileArn = st.InstanceProfileArn
	pb.MaxProvisionedConcurrency = st.MaxProvisionedConcurrency
	pb.MinProvisionedConcurrency = st.MinProvisionedConcurrency
	pb.ModelName = st.ModelName
	pb.ModelVersion = st.ModelVersion
	pb.Name = st.Name
	pb.ProvisionedModelUnits = st.ProvisionedModelUnits
	pb.ScaleToZeroEnabled = st.ScaleToZeroEnabled
	statePb, err := ServedModelStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}
	pb.WorkloadSize = st.WorkloadSize
	workloadTypePb, err := ServingModelWorkloadTypeToPb(&st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = *workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServedModelOutputFromPb(pb *servingpb.ServedModelOutputPb) (*ServedModelOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedModelOutput{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator
	st.EnvironmentVars = pb.EnvironmentVars
	st.InstanceProfileArn = pb.InstanceProfileArn
	st.MaxProvisionedConcurrency = pb.MaxProvisionedConcurrency
	st.MinProvisionedConcurrency = pb.MinProvisionedConcurrency
	st.ModelName = pb.ModelName
	st.ModelVersion = pb.ModelVersion
	st.Name = pb.Name
	st.ProvisionedModelUnits = pb.ProvisionedModelUnits
	st.ScaleToZeroEnabled = pb.ScaleToZeroEnabled
	stateField, err := ServedModelStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	st.WorkloadSize = pb.WorkloadSize
	workloadTypeField, err := ServingModelWorkloadTypeFromPb(&pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = *workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ServedModelSpec struct {
	// Only one of model_name and entity_name should be populated
	// Wire name: 'model_name'
	ModelName string ``
	// Only one of model_version and entity_version should be populated
	// Wire name: 'model_version'
	ModelVersion string ``

	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ServedModelSpec) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelSpec) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServedModelSpecToPb(st *ServedModelSpec) (*servingpb.ServedModelSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServedModelSpecPb{}
	pb.ModelName = st.ModelName
	pb.ModelVersion = st.ModelVersion
	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServedModelSpecFromPb(pb *servingpb.ServedModelSpecPb) (*ServedModelSpec, error) {
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

type ServedModelState struct {

	// Wire name: 'deployment'
	Deployment ServedModelStateDeployment ``

	// Wire name: 'deployment_state_message'
	DeploymentStateMessage string   ``
	ForceSendFields        []string `tf:"-"`
}

func (s *ServedModelState) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServedModelState) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServedModelStateToPb(st *ServedModelState) (*servingpb.ServedModelStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServedModelStatePb{}
	deploymentPb, err := ServedModelStateDeploymentToPb(&st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = *deploymentPb
	}
	pb.DeploymentStateMessage = st.DeploymentStateMessage

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServedModelStateFromPb(pb *servingpb.ServedModelStatePb) (*ServedModelState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedModelState{}
	deploymentField, err := ServedModelStateDeploymentFromPb(&pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = *deploymentField
	}
	st.DeploymentStateMessage = pb.DeploymentStateMessage

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ServedModelStateDeploymentToPb(st *ServedModelStateDeployment) (*servingpb.ServedModelStateDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.ServedModelStateDeploymentPb(*st)
	return &pb, nil
}

func ServedModelStateDeploymentFromPb(pb *servingpb.ServedModelStateDeploymentPb) (*ServedModelStateDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServedModelStateDeployment(*pb)
	return &st, nil
}

type ServerLogsResponse struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	// Wire name: 'logs'
	Logs string ``
}

func ServerLogsResponseToPb(st *ServerLogsResponse) (*servingpb.ServerLogsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServerLogsResponsePb{}
	pb.Logs = st.Logs

	return pb, nil
}

func ServerLogsResponseFromPb(pb *servingpb.ServerLogsResponsePb) (*ServerLogsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServerLogsResponse{}
	st.Logs = pb.Logs

	return st, nil
}

type ServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	// Wire name: 'ai_gateway'
	AiGateway *AiGatewayConfig ``
	// The budget policy associated with the endpoint.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// The config that is currently being served by the endpoint.
	// Wire name: 'config'
	Config *EndpointCoreConfigSummary ``
	// The timestamp when the endpoint was created in Unix time.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``
	// The email of the user who created the serving endpoint.
	// Wire name: 'creator'
	Creator string ``
	// Description of the endpoint
	// Wire name: 'description'
	Description string ``
	// System-generated ID of the endpoint, included to be used by the
	// Permissions API.
	// Wire name: 'id'
	Id string ``
	// The timestamp when the endpoint was last updated by a user in Unix time.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// The name of the serving endpoint.
	// Wire name: 'name'
	Name string ``
	// Information corresponding to the state of the serving endpoint.
	// Wire name: 'state'
	State *EndpointState ``
	// Tags attached to the serving endpoint.
	// Wire name: 'tags'
	Tags []EndpointTag ``
	// The task type of the serving endpoint.
	// Wire name: 'task'
	Task            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ServingEndpoint) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpoint) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServingEndpointToPb(st *ServingEndpoint) (*servingpb.ServingEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServingEndpointPb{}
	aiGatewayPb, err := AiGatewayConfigToPb(st.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayPb != nil {
		pb.AiGateway = aiGatewayPb
	}
	pb.BudgetPolicyId = st.BudgetPolicyId
	configPb, err := EndpointCoreConfigSummaryToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Creator = st.Creator
	pb.Description = st.Description
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.Name = st.Name
	statePb, err := EndpointStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	var tagsPb []servingpb.EndpointTagPb
	for _, item := range st.Tags {
		itemPb, err := EndpointTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb
	pb.Task = st.Task

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServingEndpointFromPb(pb *servingpb.ServingEndpointPb) (*ServingEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpoint{}
	aiGatewayField, err := AiGatewayConfigFromPb(pb.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayField != nil {
		st.AiGateway = aiGatewayField
	}
	st.BudgetPolicyId = pb.BudgetPolicyId
	configField, err := EndpointCoreConfigSummaryFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator
	st.Description = pb.Description
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Name = pb.Name
	stateField, err := EndpointStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}

	var tagsField []EndpointTag
	for _, itemPb := range pb.Tags {
		item, err := EndpointTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.Task = pb.Task

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ServingEndpointAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel ServingEndpointPermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ServingEndpointAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServingEndpointAccessControlRequestToPb(st *ServingEndpointAccessControlRequest) (*servingpb.ServingEndpointAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServingEndpointAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := ServingEndpointPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServingEndpointAccessControlRequestFromPb(pb *servingpb.ServingEndpointAccessControlRequestPb) (*ServingEndpointAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := ServingEndpointPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ServingEndpointAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []ServingEndpointPermission ``
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string ``
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ServingEndpointAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServingEndpointAccessControlResponseToPb(st *ServingEndpointAccessControlResponse) (*servingpb.ServingEndpointAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServingEndpointAccessControlResponsePb{}

	var allPermissionsPb []servingpb.ServingEndpointPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := ServingEndpointPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServingEndpointAccessControlResponseFromPb(pb *servingpb.ServingEndpointAccessControlResponsePb) (*ServingEndpointAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointAccessControlResponse{}

	var allPermissionsField []ServingEndpointPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := ServingEndpointPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ServingEndpointDetailed struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	// Wire name: 'ai_gateway'
	AiGateway *AiGatewayConfig ``
	// The budget policy associated with the endpoint.
	// Wire name: 'budget_policy_id'
	BudgetPolicyId string ``
	// The config that is currently being served by the endpoint.
	// Wire name: 'config'
	Config *EndpointCoreConfigOutput ``
	// The timestamp when the endpoint was created in Unix time.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``
	// The email of the user who created the serving endpoint.
	// Wire name: 'creator'
	Creator string ``
	// Information required to query DataPlane APIs.
	// Wire name: 'data_plane_info'
	DataPlaneInfo *ModelDataPlaneInfo ``
	// Description of the serving model
	// Wire name: 'description'
	Description string ``
	// Endpoint invocation url if route optimization is enabled for endpoint
	// Wire name: 'endpoint_url'
	EndpointUrl string ``
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	// Wire name: 'id'
	Id string ``
	// The timestamp when the endpoint was last updated by a user in Unix time.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// The name of the serving endpoint.
	// Wire name: 'name'
	Name string ``
	// The config that the endpoint is attempting to update to.
	// Wire name: 'pending_config'
	PendingConfig *EndpointPendingConfig ``
	// The permission level of the principal making the request.
	// Wire name: 'permission_level'
	PermissionLevel ServingEndpointDetailedPermissionLevel ``
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	// Wire name: 'route_optimized'
	RouteOptimized bool ``
	// Information corresponding to the state of the serving endpoint.
	// Wire name: 'state'
	State *EndpointState ``
	// Tags attached to the serving endpoint.
	// Wire name: 'tags'
	Tags []EndpointTag ``
	// The task type of the serving endpoint.
	// Wire name: 'task'
	Task            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ServingEndpointDetailed) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointDetailed) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServingEndpointDetailedToPb(st *ServingEndpointDetailed) (*servingpb.ServingEndpointDetailedPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServingEndpointDetailedPb{}
	aiGatewayPb, err := AiGatewayConfigToPb(st.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayPb != nil {
		pb.AiGateway = aiGatewayPb
	}
	pb.BudgetPolicyId = st.BudgetPolicyId
	configPb, err := EndpointCoreConfigOutputToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Creator = st.Creator
	dataPlaneInfoPb, err := ModelDataPlaneInfoToPb(st.DataPlaneInfo)
	if err != nil {
		return nil, err
	}
	if dataPlaneInfoPb != nil {
		pb.DataPlaneInfo = dataPlaneInfoPb
	}
	pb.Description = st.Description
	pb.EndpointUrl = st.EndpointUrl
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.Name = st.Name
	pendingConfigPb, err := EndpointPendingConfigToPb(st.PendingConfig)
	if err != nil {
		return nil, err
	}
	if pendingConfigPb != nil {
		pb.PendingConfig = pendingConfigPb
	}
	permissionLevelPb, err := ServingEndpointDetailedPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.RouteOptimized = st.RouteOptimized
	statePb, err := EndpointStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	var tagsPb []servingpb.EndpointTagPb
	for _, item := range st.Tags {
		itemPb, err := EndpointTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb
	pb.Task = st.Task

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServingEndpointDetailedFromPb(pb *servingpb.ServingEndpointDetailedPb) (*ServingEndpointDetailed, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointDetailed{}
	aiGatewayField, err := AiGatewayConfigFromPb(pb.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayField != nil {
		st.AiGateway = aiGatewayField
	}
	st.BudgetPolicyId = pb.BudgetPolicyId
	configField, err := EndpointCoreConfigOutputFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Creator = pb.Creator
	dataPlaneInfoField, err := ModelDataPlaneInfoFromPb(pb.DataPlaneInfo)
	if err != nil {
		return nil, err
	}
	if dataPlaneInfoField != nil {
		st.DataPlaneInfo = dataPlaneInfoField
	}
	st.Description = pb.Description
	st.EndpointUrl = pb.EndpointUrl
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Name = pb.Name
	pendingConfigField, err := EndpointPendingConfigFromPb(pb.PendingConfig)
	if err != nil {
		return nil, err
	}
	if pendingConfigField != nil {
		st.PendingConfig = pendingConfigField
	}
	permissionLevelField, err := ServingEndpointDetailedPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.RouteOptimized = pb.RouteOptimized
	stateField, err := EndpointStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}

	var tagsField []EndpointTag
	for _, itemPb := range pb.Tags {
		item, err := EndpointTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.Task = pb.Task

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ServingEndpointDetailedPermissionLevelToPb(st *ServingEndpointDetailedPermissionLevel) (*servingpb.ServingEndpointDetailedPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.ServingEndpointDetailedPermissionLevelPb(*st)
	return &pb, nil
}

func ServingEndpointDetailedPermissionLevelFromPb(pb *servingpb.ServingEndpointDetailedPermissionLevelPb) (*ServingEndpointDetailedPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServingEndpointDetailedPermissionLevel(*pb)
	return &st, nil
}

type ServingEndpointPermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel ServingEndpointPermissionLevel ``
	ForceSendFields []string                       `tf:"-"`
}

func (s *ServingEndpointPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServingEndpointPermissionToPb(st *ServingEndpointPermission) (*servingpb.ServingEndpointPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServingEndpointPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := ServingEndpointPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServingEndpointPermissionFromPb(pb *servingpb.ServingEndpointPermissionPb) (*ServingEndpointPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := ServingEndpointPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
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

func ServingEndpointPermissionLevelToPb(st *ServingEndpointPermissionLevel) (*servingpb.ServingEndpointPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.ServingEndpointPermissionLevelPb(*st)
	return &pb, nil
}

func ServingEndpointPermissionLevelFromPb(pb *servingpb.ServingEndpointPermissionLevelPb) (*ServingEndpointPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServingEndpointPermissionLevel(*pb)
	return &st, nil
}

type ServingEndpointPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []ServingEndpointAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *ServingEndpointPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServingEndpointPermissionsToPb(st *ServingEndpointPermissions) (*servingpb.ServingEndpointPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServingEndpointPermissionsPb{}

	var accessControlListPb []servingpb.ServingEndpointAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := ServingEndpointAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServingEndpointPermissionsFromPb(pb *servingpb.ServingEndpointPermissionsPb) (*ServingEndpointPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermissions{}

	var accessControlListField []ServingEndpointAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := ServingEndpointAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ServingEndpointPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel ServingEndpointPermissionLevel ``
	ForceSendFields []string                       `tf:"-"`
}

func (s *ServingEndpointPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ServingEndpointPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func ServingEndpointPermissionsDescriptionToPb(st *ServingEndpointPermissionsDescription) (*servingpb.ServingEndpointPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServingEndpointPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := ServingEndpointPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func ServingEndpointPermissionsDescriptionFromPb(pb *servingpb.ServingEndpointPermissionsDescriptionPb) (*ServingEndpointPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := ServingEndpointPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

type ServingEndpointPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []ServingEndpointAccessControlRequest ``
	// The serving endpoint for which to get or manage permissions.
	// Wire name: 'serving_endpoint_id'
	ServingEndpointId string `tf:"-"`
}

func ServingEndpointPermissionsRequestToPb(st *ServingEndpointPermissionsRequest) (*servingpb.ServingEndpointPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.ServingEndpointPermissionsRequestPb{}

	var accessControlListPb []servingpb.ServingEndpointAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := ServingEndpointAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb
	pb.ServingEndpointId = st.ServingEndpointId

	return pb, nil
}

func ServingEndpointPermissionsRequestFromPb(pb *servingpb.ServingEndpointPermissionsRequestPb) (*ServingEndpointPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermissionsRequest{}

	var accessControlListField []ServingEndpointAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := ServingEndpointAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ServingEndpointId = pb.ServingEndpointId

	return st, nil
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

func ServingModelWorkloadTypeToPb(st *ServingModelWorkloadType) (*servingpb.ServingModelWorkloadTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingpb.ServingModelWorkloadTypePb(*st)
	return &pb, nil
}

func ServingModelWorkloadTypeFromPb(pb *servingpb.ServingModelWorkloadTypePb) (*ServingModelWorkloadType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServingModelWorkloadType(*pb)
	return &st, nil
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served entity.
	// Wire name: 'routes'
	Routes []Route ``
}

func TrafficConfigToPb(st *TrafficConfig) (*servingpb.TrafficConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.TrafficConfigPb{}

	var routesPb []servingpb.RoutePb
	for _, item := range st.Routes {
		itemPb, err := RouteToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			routesPb = append(routesPb, *itemPb)
		}
	}
	pb.Routes = routesPb

	return pb, nil
}

func TrafficConfigFromPb(pb *servingpb.TrafficConfigPb) (*TrafficConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrafficConfig{}

	var routesField []Route
	for _, itemPb := range pb.Routes {
		item, err := RouteFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			routesField = append(routesField, *item)
		}
	}
	st.Routes = routesField

	return st, nil
}

type UpdateProvisionedThroughputEndpointConfigRequest struct {

	// Wire name: 'config'
	Config PtEndpointCoreConfig ``
	// The name of the pt endpoint to update. This field is required.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func UpdateProvisionedThroughputEndpointConfigRequestToPb(st *UpdateProvisionedThroughputEndpointConfigRequest) (*servingpb.UpdateProvisionedThroughputEndpointConfigRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.UpdateProvisionedThroughputEndpointConfigRequestPb{}
	configPb, err := PtEndpointCoreConfigToPb(&st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = *configPb
	}
	pb.Name = st.Name

	return pb, nil
}

func UpdateProvisionedThroughputEndpointConfigRequestFromPb(pb *servingpb.UpdateProvisionedThroughputEndpointConfigRequestPb) (*UpdateProvisionedThroughputEndpointConfigRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateProvisionedThroughputEndpointConfigRequest{}
	configField, err := PtEndpointCoreConfigFromPb(&pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = *configField
	}
	st.Name = pb.Name

	return st, nil
}

type V1ResponseChoiceElement struct {
	// The finish reason returned by the endpoint.
	// Wire name: 'finishReason'
	FinishReason string ``
	// The index of the choice in the __chat or completions__ response.
	// Wire name: 'index'
	Index int ``
	// The logprobs returned only by the __completions__ endpoint.
	// Wire name: 'logprobs'
	Logprobs int ``
	// The message response from the __chat__ endpoint.
	// Wire name: 'message'
	Message *ChatMessage ``
	// The text response from the __completions__ endpoint.
	// Wire name: 'text'
	Text            string   ``
	ForceSendFields []string `tf:"-"`
}

func (s *V1ResponseChoiceElement) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s V1ResponseChoiceElement) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

func V1ResponseChoiceElementToPb(st *V1ResponseChoiceElement) (*servingpb.V1ResponseChoiceElementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingpb.V1ResponseChoiceElementPb{}
	pb.FinishReason = st.FinishReason
	pb.Index = st.Index
	pb.Logprobs = st.Logprobs
	messagePb, err := ChatMessageToPb(st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = messagePb
	}
	pb.Text = st.Text

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func V1ResponseChoiceElementFromPb(pb *servingpb.V1ResponseChoiceElementPb) (*V1ResponseChoiceElement, error) {
	if pb == nil {
		return nil, nil
	}
	st := &V1ResponseChoiceElement{}
	st.FinishReason = pb.FinishReason
	st.Index = pb.Index
	st.Logprobs = pb.Logprobs
	messageField, err := ChatMessageFromPb(pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = messageField
	}
	st.Text = pb.Text

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%.9fs", d.Seconds())
	return &s, nil
}

func durationFromPb(s *string) (*time.Duration, error) {
	if s == nil {
		return nil, nil
	}
	d, err := time.ParseDuration(*s)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

func timestampToPb(t *time.Time) (*string, error) {
	if t == nil {
		return nil, nil
	}
	s := t.Format(time.RFC3339)
	return &s, nil
}

func timestampFromPb(s *string) (*time.Time, error) {
	if s == nil {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, *s)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func fieldMaskToPb(fm *[]string) (*string, error) {
	if fm == nil {
		return nil, nil
	}
	s := strings.Join(*fm, ",")
	return &s, nil
}

func fieldMaskFromPb(s *string) (*[]string, error) {
	if s == nil {
		return nil, nil
	}
	fm := strings.Split(*s, ",")
	return &fm, nil
}
