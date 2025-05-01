// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package serving

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func identity[T any](obj *T) (*T, error) {
	return obj, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
	return &s, nil
}

// Helper to strip trailing zeros in fractional part
func rstripZeros(s string) string {
	for len(s) > 0 && s[len(s)-1] == '0' {
		s = s[:len(s)-1]
	}
	if len(s) > 0 && s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}
	return s
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

type Ai21LabsConfig struct {
	// The Databricks secret key reference for an AI21 Labs API key. If you
	// prefer to paste your API key directly, see `ai21labs_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	Ai21labsApiKey string
	// An AI21 Labs API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `ai21labs_api_key`. You
	// must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	Ai21labsApiKeyPlaintext string

	ForceSendFields []string
}

func ai21LabsConfigToPb(st *Ai21LabsConfig) (*ai21LabsConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ai21LabsConfigPb{}
	ai21labsApiKeyPb, err := identity(&st.Ai21labsApiKey)
	if err != nil {
		return nil, err
	}
	if ai21labsApiKeyPb != nil {
		pb.Ai21labsApiKey = *ai21labsApiKeyPb
	}

	ai21labsApiKeyPlaintextPb, err := identity(&st.Ai21labsApiKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if ai21labsApiKeyPlaintextPb != nil {
		pb.Ai21labsApiKeyPlaintext = *ai21labsApiKeyPlaintextPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type ai21LabsConfigPb struct {
	// The Databricks secret key reference for an AI21 Labs API key. If you
	// prefer to paste your API key directly, see `ai21labs_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	Ai21labsApiKey string `json:"ai21labs_api_key,omitempty"`
	// An AI21 Labs API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `ai21labs_api_key`. You
	// must provide an API key using one of the following fields:
	// `ai21labs_api_key` or `ai21labs_api_key_plaintext`.
	Ai21labsApiKeyPlaintext string `json:"ai21labs_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ai21LabsConfigFromPb(pb *ai21LabsConfigPb) (*Ai21LabsConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Ai21LabsConfig{}
	ai21labsApiKeyField, err := identity(&pb.Ai21labsApiKey)
	if err != nil {
		return nil, err
	}
	if ai21labsApiKeyField != nil {
		st.Ai21labsApiKey = *ai21labsApiKeyField
	}
	ai21labsApiKeyPlaintextField, err := identity(&pb.Ai21labsApiKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if ai21labsApiKeyPlaintextField != nil {
		st.Ai21labsApiKeyPlaintext = *ai21labsApiKeyPlaintextField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ai21LabsConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ai21LabsConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AiGatewayConfig struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig *FallbackConfig
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails *AiGatewayGuardrails
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig *AiGatewayInferenceTableConfig
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []AiGatewayRateLimit
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig *AiGatewayUsageTrackingConfig
}

func aiGatewayConfigToPb(st *AiGatewayConfig) (*aiGatewayConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayConfigPb{}
	fallbackConfigPb, err := fallbackConfigToPb(st.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigPb != nil {
		pb.FallbackConfig = fallbackConfigPb
	}

	guardrailsPb, err := aiGatewayGuardrailsToPb(st.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsPb != nil {
		pb.Guardrails = guardrailsPb
	}

	inferenceTableConfigPb, err := aiGatewayInferenceTableConfigToPb(st.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigPb != nil {
		pb.InferenceTableConfig = inferenceTableConfigPb
	}

	var rateLimitsPb []aiGatewayRateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := aiGatewayRateLimitToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rateLimitsPb = append(rateLimitsPb, *itemPb)
		}
	}
	pb.RateLimits = rateLimitsPb

	usageTrackingConfigPb, err := aiGatewayUsageTrackingConfigToPb(st.UsageTrackingConfig)
	if err != nil {
		return nil, err
	}
	if usageTrackingConfigPb != nil {
		pb.UsageTrackingConfig = usageTrackingConfigPb
	}

	return pb, nil
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

type aiGatewayConfigPb struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig *fallbackConfigPb `json:"fallback_config,omitempty"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails *aiGatewayGuardrailsPb `json:"guardrails,omitempty"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig *aiGatewayInferenceTableConfigPb `json:"inference_table_config,omitempty"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []aiGatewayRateLimitPb `json:"rate_limits,omitempty"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig *aiGatewayUsageTrackingConfigPb `json:"usage_tracking_config,omitempty"`
}

func aiGatewayConfigFromPb(pb *aiGatewayConfigPb) (*AiGatewayConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayConfig{}
	fallbackConfigField, err := fallbackConfigFromPb(pb.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigField != nil {
		st.FallbackConfig = fallbackConfigField
	}
	guardrailsField, err := aiGatewayGuardrailsFromPb(pb.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsField != nil {
		st.Guardrails = guardrailsField
	}
	inferenceTableConfigField, err := aiGatewayInferenceTableConfigFromPb(pb.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigField != nil {
		st.InferenceTableConfig = inferenceTableConfigField
	}

	var rateLimitsField []AiGatewayRateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := aiGatewayRateLimitFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rateLimitsField = append(rateLimitsField, *item)
		}
	}
	st.RateLimits = rateLimitsField
	usageTrackingConfigField, err := aiGatewayUsageTrackingConfigFromPb(pb.UsageTrackingConfig)
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
	InvalidKeywords []string
	// Configuration for guardrail PII filter.
	Pii *AiGatewayGuardrailPiiBehavior
	// Indicates whether the safety filter is enabled.
	Safety bool
	// The list of allowed topics. Given a chat request, this guardrail flags
	// the request if its topic is not in the allowed topics.
	ValidTopics []string

	ForceSendFields []string
}

func aiGatewayGuardrailParametersToPb(st *AiGatewayGuardrailParameters) (*aiGatewayGuardrailParametersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayGuardrailParametersPb{}

	var invalidKeywordsPb []string
	for _, item := range st.InvalidKeywords {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			invalidKeywordsPb = append(invalidKeywordsPb, *itemPb)
		}
	}
	pb.InvalidKeywords = invalidKeywordsPb

	piiPb, err := aiGatewayGuardrailPiiBehaviorToPb(st.Pii)
	if err != nil {
		return nil, err
	}
	if piiPb != nil {
		pb.Pii = piiPb
	}

	safetyPb, err := identity(&st.Safety)
	if err != nil {
		return nil, err
	}
	if safetyPb != nil {
		pb.Safety = *safetyPb
	}

	var validTopicsPb []string
	for _, item := range st.ValidTopics {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			validTopicsPb = append(validTopicsPb, *itemPb)
		}
	}
	pb.ValidTopics = validTopicsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type aiGatewayGuardrailParametersPb struct {
	// List of invalid keywords. AI guardrail uses keyword or string matching to
	// decide if the keyword exists in the request or response content.
	InvalidKeywords []string `json:"invalid_keywords,omitempty"`
	// Configuration for guardrail PII filter.
	Pii *aiGatewayGuardrailPiiBehaviorPb `json:"pii,omitempty"`
	// Indicates whether the safety filter is enabled.
	Safety bool `json:"safety,omitempty"`
	// The list of allowed topics. Given a chat request, this guardrail flags
	// the request if its topic is not in the allowed topics.
	ValidTopics []string `json:"valid_topics,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func aiGatewayGuardrailParametersFromPb(pb *aiGatewayGuardrailParametersPb) (*AiGatewayGuardrailParameters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayGuardrailParameters{}

	var invalidKeywordsField []string
	for _, itemPb := range pb.InvalidKeywords {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			invalidKeywordsField = append(invalidKeywordsField, *item)
		}
	}
	st.InvalidKeywords = invalidKeywordsField
	piiField, err := aiGatewayGuardrailPiiBehaviorFromPb(pb.Pii)
	if err != nil {
		return nil, err
	}
	if piiField != nil {
		st.Pii = piiField
	}
	safetyField, err := identity(&pb.Safety)
	if err != nil {
		return nil, err
	}
	if safetyField != nil {
		st.Safety = *safetyField
	}

	var validTopicsField []string
	for _, itemPb := range pb.ValidTopics {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			validTopicsField = append(validTopicsField, *item)
		}
	}
	st.ValidTopics = validTopicsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *aiGatewayGuardrailParametersPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st aiGatewayGuardrailParametersPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AiGatewayGuardrailPiiBehavior struct {
	// Configuration for input guardrail filters.
	Behavior AiGatewayGuardrailPiiBehaviorBehavior
}

func aiGatewayGuardrailPiiBehaviorToPb(st *AiGatewayGuardrailPiiBehavior) (*aiGatewayGuardrailPiiBehaviorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayGuardrailPiiBehaviorPb{}
	behaviorPb, err := identity(&st.Behavior)
	if err != nil {
		return nil, err
	}
	if behaviorPb != nil {
		pb.Behavior = *behaviorPb
	}

	return pb, nil
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

type aiGatewayGuardrailPiiBehaviorPb struct {
	// Configuration for input guardrail filters.
	Behavior AiGatewayGuardrailPiiBehaviorBehavior `json:"behavior,omitempty"`
}

func aiGatewayGuardrailPiiBehaviorFromPb(pb *aiGatewayGuardrailPiiBehaviorPb) (*AiGatewayGuardrailPiiBehavior, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayGuardrailPiiBehavior{}
	behaviorField, err := identity(&pb.Behavior)
	if err != nil {
		return nil, err
	}
	if behaviorField != nil {
		st.Behavior = *behaviorField
	}

	return st, nil
}

type AiGatewayGuardrailPiiBehaviorBehavior string
type aiGatewayGuardrailPiiBehaviorBehaviorPb string

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

// Type always returns AiGatewayGuardrailPiiBehaviorBehavior to satisfy [pflag.Value] interface
func (f *AiGatewayGuardrailPiiBehaviorBehavior) Type() string {
	return "AiGatewayGuardrailPiiBehaviorBehavior"
}

func aiGatewayGuardrailPiiBehaviorBehaviorToPb(st *AiGatewayGuardrailPiiBehaviorBehavior) (*aiGatewayGuardrailPiiBehaviorBehaviorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := aiGatewayGuardrailPiiBehaviorBehaviorPb(*st)
	return &pb, nil
}

func aiGatewayGuardrailPiiBehaviorBehaviorFromPb(pb *aiGatewayGuardrailPiiBehaviorBehaviorPb) (*AiGatewayGuardrailPiiBehaviorBehavior, error) {
	if pb == nil {
		return nil, nil
	}
	st := AiGatewayGuardrailPiiBehaviorBehavior(*pb)
	return &st, nil
}

type AiGatewayGuardrails struct {
	// Configuration for input guardrail filters.
	Input *AiGatewayGuardrailParameters
	// Configuration for output guardrail filters.
	Output *AiGatewayGuardrailParameters
}

func aiGatewayGuardrailsToPb(st *AiGatewayGuardrails) (*aiGatewayGuardrailsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayGuardrailsPb{}
	inputPb, err := aiGatewayGuardrailParametersToPb(st.Input)
	if err != nil {
		return nil, err
	}
	if inputPb != nil {
		pb.Input = inputPb
	}

	outputPb, err := aiGatewayGuardrailParametersToPb(st.Output)
	if err != nil {
		return nil, err
	}
	if outputPb != nil {
		pb.Output = outputPb
	}

	return pb, nil
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

type aiGatewayGuardrailsPb struct {
	// Configuration for input guardrail filters.
	Input *aiGatewayGuardrailParametersPb `json:"input,omitempty"`
	// Configuration for output guardrail filters.
	Output *aiGatewayGuardrailParametersPb `json:"output,omitempty"`
}

func aiGatewayGuardrailsFromPb(pb *aiGatewayGuardrailsPb) (*AiGatewayGuardrails, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayGuardrails{}
	inputField, err := aiGatewayGuardrailParametersFromPb(pb.Input)
	if err != nil {
		return nil, err
	}
	if inputField != nil {
		st.Input = inputField
	}
	outputField, err := aiGatewayGuardrailParametersFromPb(pb.Output)
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
	CatalogName string
	// Indicates whether the inference table is enabled.
	Enabled bool
	// The name of the schema in Unity Catalog. Required when enabling inference
	// tables. NOTE: On update, you have to disable inference table first in
	// order to change the schema name.
	SchemaName string
	// The prefix of the table in Unity Catalog. NOTE: On update, you have to
	// disable inference table first in order to change the prefix name.
	TableNamePrefix string

	ForceSendFields []string
}

func aiGatewayInferenceTableConfigToPb(st *AiGatewayInferenceTableConfig) (*aiGatewayInferenceTableConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayInferenceTableConfigPb{}
	catalogNamePb, err := identity(&st.CatalogName)
	if err != nil {
		return nil, err
	}
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	schemaNamePb, err := identity(&st.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	tableNamePrefixPb, err := identity(&st.TableNamePrefix)
	if err != nil {
		return nil, err
	}
	if tableNamePrefixPb != nil {
		pb.TableNamePrefix = *tableNamePrefixPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type aiGatewayInferenceTableConfigPb struct {
	// The name of the catalog in Unity Catalog. Required when enabling
	// inference tables. NOTE: On update, you have to disable inference table
	// first in order to change the catalog name.
	CatalogName string `json:"catalog_name,omitempty"`
	// Indicates whether the inference table is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog. Required when enabling inference
	// tables. NOTE: On update, you have to disable inference table first in
	// order to change the schema name.
	SchemaName string `json:"schema_name,omitempty"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you have to
	// disable inference table first in order to change the prefix name.
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func aiGatewayInferenceTableConfigFromPb(pb *aiGatewayInferenceTableConfigPb) (*AiGatewayInferenceTableConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayInferenceTableConfig{}
	catalogNameField, err := identity(&pb.CatalogName)
	if err != nil {
		return nil, err
	}
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	schemaNameField, err := identity(&pb.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	tableNamePrefixField, err := identity(&pb.TableNamePrefix)
	if err != nil {
		return nil, err
	}
	if tableNamePrefixField != nil {
		st.TableNamePrefix = *tableNamePrefixField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *aiGatewayInferenceTableConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st aiGatewayInferenceTableConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AiGatewayRateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls int64
	// Key field for a rate limit. Currently, only 'user' and 'endpoint' are
	// supported, with 'endpoint' being the default if not specified.
	Key AiGatewayRateLimitKey
	// Renewal period field for a rate limit. Currently, only 'minute' is
	// supported.
	RenewalPeriod AiGatewayRateLimitRenewalPeriod
}

func aiGatewayRateLimitToPb(st *AiGatewayRateLimit) (*aiGatewayRateLimitPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayRateLimitPb{}
	callsPb, err := identity(&st.Calls)
	if err != nil {
		return nil, err
	}
	if callsPb != nil {
		pb.Calls = *callsPb
	}

	keyPb, err := identity(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}

	renewalPeriodPb, err := identity(&st.RenewalPeriod)
	if err != nil {
		return nil, err
	}
	if renewalPeriodPb != nil {
		pb.RenewalPeriod = *renewalPeriodPb
	}

	return pb, nil
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

type aiGatewayRateLimitPb struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls int64 `json:"calls"`
	// Key field for a rate limit. Currently, only 'user' and 'endpoint' are
	// supported, with 'endpoint' being the default if not specified.
	Key AiGatewayRateLimitKey `json:"key,omitempty"`
	// Renewal period field for a rate limit. Currently, only 'minute' is
	// supported.
	RenewalPeriod AiGatewayRateLimitRenewalPeriod `json:"renewal_period"`
}

func aiGatewayRateLimitFromPb(pb *aiGatewayRateLimitPb) (*AiGatewayRateLimit, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayRateLimit{}
	callsField, err := identity(&pb.Calls)
	if err != nil {
		return nil, err
	}
	if callsField != nil {
		st.Calls = *callsField
	}
	keyField, err := identity(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
	renewalPeriodField, err := identity(&pb.RenewalPeriod)
	if err != nil {
		return nil, err
	}
	if renewalPeriodField != nil {
		st.RenewalPeriod = *renewalPeriodField
	}

	return st, nil
}

type AiGatewayRateLimitKey string
type aiGatewayRateLimitKeyPb string

const AiGatewayRateLimitKeyEndpoint AiGatewayRateLimitKey = `endpoint`

const AiGatewayRateLimitKeyUser AiGatewayRateLimitKey = `user`

// String representation for [fmt.Print]
func (f *AiGatewayRateLimitKey) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AiGatewayRateLimitKey) Set(v string) error {
	switch v {
	case `endpoint`, `user`:
		*f = AiGatewayRateLimitKey(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "endpoint", "user"`, v)
	}
}

// Type always returns AiGatewayRateLimitKey to satisfy [pflag.Value] interface
func (f *AiGatewayRateLimitKey) Type() string {
	return "AiGatewayRateLimitKey"
}

func aiGatewayRateLimitKeyToPb(st *AiGatewayRateLimitKey) (*aiGatewayRateLimitKeyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := aiGatewayRateLimitKeyPb(*st)
	return &pb, nil
}

func aiGatewayRateLimitKeyFromPb(pb *aiGatewayRateLimitKeyPb) (*AiGatewayRateLimitKey, error) {
	if pb == nil {
		return nil, nil
	}
	st := AiGatewayRateLimitKey(*pb)
	return &st, nil
}

type AiGatewayRateLimitRenewalPeriod string
type aiGatewayRateLimitRenewalPeriodPb string

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

// Type always returns AiGatewayRateLimitRenewalPeriod to satisfy [pflag.Value] interface
func (f *AiGatewayRateLimitRenewalPeriod) Type() string {
	return "AiGatewayRateLimitRenewalPeriod"
}

func aiGatewayRateLimitRenewalPeriodToPb(st *AiGatewayRateLimitRenewalPeriod) (*aiGatewayRateLimitRenewalPeriodPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := aiGatewayRateLimitRenewalPeriodPb(*st)
	return &pb, nil
}

func aiGatewayRateLimitRenewalPeriodFromPb(pb *aiGatewayRateLimitRenewalPeriodPb) (*AiGatewayRateLimitRenewalPeriod, error) {
	if pb == nil {
		return nil, nil
	}
	st := AiGatewayRateLimitRenewalPeriod(*pb)
	return &st, nil
}

type AiGatewayUsageTrackingConfig struct {
	// Whether to enable usage tracking.
	Enabled bool

	ForceSendFields []string
}

func aiGatewayUsageTrackingConfigToPb(st *AiGatewayUsageTrackingConfig) (*aiGatewayUsageTrackingConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &aiGatewayUsageTrackingConfigPb{}
	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type aiGatewayUsageTrackingConfigPb struct {
	// Whether to enable usage tracking.
	Enabled bool `json:"enabled,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func aiGatewayUsageTrackingConfigFromPb(pb *aiGatewayUsageTrackingConfigPb) (*AiGatewayUsageTrackingConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AiGatewayUsageTrackingConfig{}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *aiGatewayUsageTrackingConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st aiGatewayUsageTrackingConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AmazonBedrockConfig struct {
	// The Databricks secret key reference for an AWS access key ID with
	// permissions to interact with Bedrock services. If you prefer to paste
	// your API key directly, see `aws_access_key_id_plaintext`. You must
	// provide an API key using one of the following fields: `aws_access_key_id`
	// or `aws_access_key_id_plaintext`.
	AwsAccessKeyId string
	// An AWS access key ID with permissions to interact with Bedrock services
	// provided as a plaintext string. If you prefer to reference your key using
	// Databricks Secrets, see `aws_access_key_id`. You must provide an API key
	// using one of the following fields: `aws_access_key_id` or
	// `aws_access_key_id_plaintext`.
	AwsAccessKeyIdPlaintext string
	// The AWS region to use. Bedrock has to be enabled there.
	AwsRegion string
	// The Databricks secret key reference for an AWS secret access key paired
	// with the access key ID, with permissions to interact with Bedrock
	// services. If you prefer to paste your API key directly, see
	// `aws_secret_access_key_plaintext`. You must provide an API key using one
	// of the following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	AwsSecretAccessKey string
	// An AWS secret access key paired with the access key ID, with permissions
	// to interact with Bedrock services provided as a plaintext string. If you
	// prefer to reference your key using Databricks Secrets, see
	// `aws_secret_access_key`. You must provide an API key using one of the
	// following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	AwsSecretAccessKeyPlaintext string
	// The underlying provider in Amazon Bedrock. Supported values (case
	// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
	BedrockProvider AmazonBedrockConfigBedrockProvider
	// ARN of the instance profile that the external model will use to access
	// AWS resources. You must authenticate using an instance profile or access
	// keys. If you prefer to authenticate using access keys, see
	// `aws_access_key_id`, `aws_access_key_id_plaintext`,
	// `aws_secret_access_key` and `aws_secret_access_key_plaintext`.
	InstanceProfileArn string

	ForceSendFields []string
}

func amazonBedrockConfigToPb(st *AmazonBedrockConfig) (*amazonBedrockConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &amazonBedrockConfigPb{}
	awsAccessKeyIdPb, err := identity(&st.AwsAccessKeyId)
	if err != nil {
		return nil, err
	}
	if awsAccessKeyIdPb != nil {
		pb.AwsAccessKeyId = *awsAccessKeyIdPb
	}

	awsAccessKeyIdPlaintextPb, err := identity(&st.AwsAccessKeyIdPlaintext)
	if err != nil {
		return nil, err
	}
	if awsAccessKeyIdPlaintextPb != nil {
		pb.AwsAccessKeyIdPlaintext = *awsAccessKeyIdPlaintextPb
	}

	awsRegionPb, err := identity(&st.AwsRegion)
	if err != nil {
		return nil, err
	}
	if awsRegionPb != nil {
		pb.AwsRegion = *awsRegionPb
	}

	awsSecretAccessKeyPb, err := identity(&st.AwsSecretAccessKey)
	if err != nil {
		return nil, err
	}
	if awsSecretAccessKeyPb != nil {
		pb.AwsSecretAccessKey = *awsSecretAccessKeyPb
	}

	awsSecretAccessKeyPlaintextPb, err := identity(&st.AwsSecretAccessKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if awsSecretAccessKeyPlaintextPb != nil {
		pb.AwsSecretAccessKeyPlaintext = *awsSecretAccessKeyPlaintextPb
	}

	bedrockProviderPb, err := identity(&st.BedrockProvider)
	if err != nil {
		return nil, err
	}
	if bedrockProviderPb != nil {
		pb.BedrockProvider = *bedrockProviderPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type amazonBedrockConfigPb struct {
	// The Databricks secret key reference for an AWS access key ID with
	// permissions to interact with Bedrock services. If you prefer to paste
	// your API key directly, see `aws_access_key_id_plaintext`. You must
	// provide an API key using one of the following fields: `aws_access_key_id`
	// or `aws_access_key_id_plaintext`.
	AwsAccessKeyId string `json:"aws_access_key_id,omitempty"`
	// An AWS access key ID with permissions to interact with Bedrock services
	// provided as a plaintext string. If you prefer to reference your key using
	// Databricks Secrets, see `aws_access_key_id`. You must provide an API key
	// using one of the following fields: `aws_access_key_id` or
	// `aws_access_key_id_plaintext`.
	AwsAccessKeyIdPlaintext string `json:"aws_access_key_id_plaintext,omitempty"`
	// The AWS region to use. Bedrock has to be enabled there.
	AwsRegion string `json:"aws_region"`
	// The Databricks secret key reference for an AWS secret access key paired
	// with the access key ID, with permissions to interact with Bedrock
	// services. If you prefer to paste your API key directly, see
	// `aws_secret_access_key_plaintext`. You must provide an API key using one
	// of the following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	AwsSecretAccessKey string `json:"aws_secret_access_key,omitempty"`
	// An AWS secret access key paired with the access key ID, with permissions
	// to interact with Bedrock services provided as a plaintext string. If you
	// prefer to reference your key using Databricks Secrets, see
	// `aws_secret_access_key`. You must provide an API key using one of the
	// following fields: `aws_secret_access_key` or
	// `aws_secret_access_key_plaintext`.
	AwsSecretAccessKeyPlaintext string `json:"aws_secret_access_key_plaintext,omitempty"`
	// The underlying provider in Amazon Bedrock. Supported values (case
	// insensitive) include: Anthropic, Cohere, AI21Labs, Amazon.
	BedrockProvider AmazonBedrockConfigBedrockProvider `json:"bedrock_provider"`
	// ARN of the instance profile that the external model will use to access
	// AWS resources. You must authenticate using an instance profile or access
	// keys. If you prefer to authenticate using access keys, see
	// `aws_access_key_id`, `aws_access_key_id_plaintext`,
	// `aws_secret_access_key` and `aws_secret_access_key_plaintext`.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func amazonBedrockConfigFromPb(pb *amazonBedrockConfigPb) (*AmazonBedrockConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AmazonBedrockConfig{}
	awsAccessKeyIdField, err := identity(&pb.AwsAccessKeyId)
	if err != nil {
		return nil, err
	}
	if awsAccessKeyIdField != nil {
		st.AwsAccessKeyId = *awsAccessKeyIdField
	}
	awsAccessKeyIdPlaintextField, err := identity(&pb.AwsAccessKeyIdPlaintext)
	if err != nil {
		return nil, err
	}
	if awsAccessKeyIdPlaintextField != nil {
		st.AwsAccessKeyIdPlaintext = *awsAccessKeyIdPlaintextField
	}
	awsRegionField, err := identity(&pb.AwsRegion)
	if err != nil {
		return nil, err
	}
	if awsRegionField != nil {
		st.AwsRegion = *awsRegionField
	}
	awsSecretAccessKeyField, err := identity(&pb.AwsSecretAccessKey)
	if err != nil {
		return nil, err
	}
	if awsSecretAccessKeyField != nil {
		st.AwsSecretAccessKey = *awsSecretAccessKeyField
	}
	awsSecretAccessKeyPlaintextField, err := identity(&pb.AwsSecretAccessKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if awsSecretAccessKeyPlaintextField != nil {
		st.AwsSecretAccessKeyPlaintext = *awsSecretAccessKeyPlaintextField
	}
	bedrockProviderField, err := identity(&pb.BedrockProvider)
	if err != nil {
		return nil, err
	}
	if bedrockProviderField != nil {
		st.BedrockProvider = *bedrockProviderField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *amazonBedrockConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st amazonBedrockConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AmazonBedrockConfigBedrockProvider string
type amazonBedrockConfigBedrockProviderPb string

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

// Type always returns AmazonBedrockConfigBedrockProvider to satisfy [pflag.Value] interface
func (f *AmazonBedrockConfigBedrockProvider) Type() string {
	return "AmazonBedrockConfigBedrockProvider"
}

func amazonBedrockConfigBedrockProviderToPb(st *AmazonBedrockConfigBedrockProvider) (*amazonBedrockConfigBedrockProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := amazonBedrockConfigBedrockProviderPb(*st)
	return &pb, nil
}

func amazonBedrockConfigBedrockProviderFromPb(pb *amazonBedrockConfigBedrockProviderPb) (*AmazonBedrockConfigBedrockProvider, error) {
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
	AnthropicApiKey string
	// The Anthropic API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `anthropic_api_key`. You
	// must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	AnthropicApiKeyPlaintext string

	ForceSendFields []string
}

func anthropicConfigToPb(st *AnthropicConfig) (*anthropicConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &anthropicConfigPb{}
	anthropicApiKeyPb, err := identity(&st.AnthropicApiKey)
	if err != nil {
		return nil, err
	}
	if anthropicApiKeyPb != nil {
		pb.AnthropicApiKey = *anthropicApiKeyPb
	}

	anthropicApiKeyPlaintextPb, err := identity(&st.AnthropicApiKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if anthropicApiKeyPlaintextPb != nil {
		pb.AnthropicApiKeyPlaintext = *anthropicApiKeyPlaintextPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type anthropicConfigPb struct {
	// The Databricks secret key reference for an Anthropic API key. If you
	// prefer to paste your API key directly, see `anthropic_api_key_plaintext`.
	// You must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	AnthropicApiKey string `json:"anthropic_api_key,omitempty"`
	// The Anthropic API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `anthropic_api_key`. You
	// must provide an API key using one of the following fields:
	// `anthropic_api_key` or `anthropic_api_key_plaintext`.
	AnthropicApiKeyPlaintext string `json:"anthropic_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func anthropicConfigFromPb(pb *anthropicConfigPb) (*AnthropicConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AnthropicConfig{}
	anthropicApiKeyField, err := identity(&pb.AnthropicApiKey)
	if err != nil {
		return nil, err
	}
	if anthropicApiKeyField != nil {
		st.AnthropicApiKey = *anthropicApiKeyField
	}
	anthropicApiKeyPlaintextField, err := identity(&pb.AnthropicApiKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if anthropicApiKeyPlaintextField != nil {
		st.AnthropicApiKeyPlaintext = *anthropicApiKeyPlaintextField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *anthropicConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st anthropicConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ApiKeyAuth struct {
	// The name of the API key parameter used for authentication.
	Key string
	// The Databricks secret key reference for an API Key. If you prefer to
	// paste your token directly, see `value_plaintext`.
	Value string
	// The API Key provided as a plaintext string. If you prefer to reference
	// your token using Databricks Secrets, see `value`.
	ValuePlaintext string

	ForceSendFields []string
}

func apiKeyAuthToPb(st *ApiKeyAuth) (*apiKeyAuthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &apiKeyAuthPb{}
	keyPb, err := identity(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}

	valuePb, err := identity(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

	valuePlaintextPb, err := identity(&st.ValuePlaintext)
	if err != nil {
		return nil, err
	}
	if valuePlaintextPb != nil {
		pb.ValuePlaintext = *valuePlaintextPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type apiKeyAuthPb struct {
	// The name of the API key parameter used for authentication.
	Key string `json:"key"`
	// The Databricks secret key reference for an API Key. If you prefer to
	// paste your token directly, see `value_plaintext`.
	Value string `json:"value,omitempty"`
	// The API Key provided as a plaintext string. If you prefer to reference
	// your token using Databricks Secrets, see `value`.
	ValuePlaintext string `json:"value_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func apiKeyAuthFromPb(pb *apiKeyAuthPb) (*ApiKeyAuth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApiKeyAuth{}
	keyField, err := identity(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
	valueField, err := identity(&pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = *valueField
	}
	valuePlaintextField, err := identity(&pb.ValuePlaintext)
	if err != nil {
		return nil, err
	}
	if valuePlaintextField != nil {
		st.ValuePlaintext = *valuePlaintextField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *apiKeyAuthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st apiKeyAuthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AutoCaptureConfigInput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	CatalogName string
	// Indicates whether the inference table is enabled.
	Enabled bool
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	SchemaName string
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	TableNamePrefix string

	ForceSendFields []string
}

func autoCaptureConfigInputToPb(st *AutoCaptureConfigInput) (*autoCaptureConfigInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &autoCaptureConfigInputPb{}
	catalogNamePb, err := identity(&st.CatalogName)
	if err != nil {
		return nil, err
	}
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	schemaNamePb, err := identity(&st.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	tableNamePrefixPb, err := identity(&st.TableNamePrefix)
	if err != nil {
		return nil, err
	}
	if tableNamePrefixPb != nil {
		pb.TableNamePrefix = *tableNamePrefixPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type autoCaptureConfigInputPb struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	CatalogName string `json:"catalog_name,omitempty"`
	// Indicates whether the inference table is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	SchemaName string `json:"schema_name,omitempty"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func autoCaptureConfigInputFromPb(pb *autoCaptureConfigInputPb) (*AutoCaptureConfigInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoCaptureConfigInput{}
	catalogNameField, err := identity(&pb.CatalogName)
	if err != nil {
		return nil, err
	}
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	schemaNameField, err := identity(&pb.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	tableNamePrefixField, err := identity(&pb.TableNamePrefix)
	if err != nil {
		return nil, err
	}
	if tableNamePrefixField != nil {
		st.TableNamePrefix = *tableNamePrefixField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *autoCaptureConfigInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st autoCaptureConfigInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AutoCaptureConfigOutput struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	CatalogName string
	// Indicates whether the inference table is enabled.
	Enabled bool
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	SchemaName string

	State *AutoCaptureState
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	TableNamePrefix string

	ForceSendFields []string
}

func autoCaptureConfigOutputToPb(st *AutoCaptureConfigOutput) (*autoCaptureConfigOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &autoCaptureConfigOutputPb{}
	catalogNamePb, err := identity(&st.CatalogName)
	if err != nil {
		return nil, err
	}
	if catalogNamePb != nil {
		pb.CatalogName = *catalogNamePb
	}

	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	schemaNamePb, err := identity(&st.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNamePb != nil {
		pb.SchemaName = *schemaNamePb
	}

	statePb, err := autoCaptureStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	tableNamePrefixPb, err := identity(&st.TableNamePrefix)
	if err != nil {
		return nil, err
	}
	if tableNamePrefixPb != nil {
		pb.TableNamePrefix = *tableNamePrefixPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type autoCaptureConfigOutputPb struct {
	// The name of the catalog in Unity Catalog. NOTE: On update, you cannot
	// change the catalog name if the inference table is already enabled.
	CatalogName string `json:"catalog_name,omitempty"`
	// Indicates whether the inference table is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The name of the schema in Unity Catalog. NOTE: On update, you cannot
	// change the schema name if the inference table is already enabled.
	SchemaName string `json:"schema_name,omitempty"`

	State *autoCaptureStatePb `json:"state,omitempty"`
	// The prefix of the table in Unity Catalog. NOTE: On update, you cannot
	// change the prefix name if the inference table is already enabled.
	TableNamePrefix string `json:"table_name_prefix,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func autoCaptureConfigOutputFromPb(pb *autoCaptureConfigOutputPb) (*AutoCaptureConfigOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoCaptureConfigOutput{}
	catalogNameField, err := identity(&pb.CatalogName)
	if err != nil {
		return nil, err
	}
	if catalogNameField != nil {
		st.CatalogName = *catalogNameField
	}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}
	schemaNameField, err := identity(&pb.SchemaName)
	if err != nil {
		return nil, err
	}
	if schemaNameField != nil {
		st.SchemaName = *schemaNameField
	}
	stateField, err := autoCaptureStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	tableNamePrefixField, err := identity(&pb.TableNamePrefix)
	if err != nil {
		return nil, err
	}
	if tableNamePrefixField != nil {
		st.TableNamePrefix = *tableNamePrefixField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *autoCaptureConfigOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st autoCaptureConfigOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AutoCaptureState struct {
	PayloadTable *PayloadTable
}

func autoCaptureStateToPb(st *AutoCaptureState) (*autoCaptureStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &autoCaptureStatePb{}
	payloadTablePb, err := payloadTableToPb(st.PayloadTable)
	if err != nil {
		return nil, err
	}
	if payloadTablePb != nil {
		pb.PayloadTable = payloadTablePb
	}

	return pb, nil
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

type autoCaptureStatePb struct {
	PayloadTable *payloadTablePb `json:"payload_table,omitempty"`
}

func autoCaptureStateFromPb(pb *autoCaptureStatePb) (*AutoCaptureState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AutoCaptureState{}
	payloadTableField, err := payloadTableFromPb(pb.PayloadTable)
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
	Token string
	// The token provided as a plaintext string. If you prefer to reference your
	// token using Databricks Secrets, see `token`.
	TokenPlaintext string

	ForceSendFields []string
}

func bearerTokenAuthToPb(st *BearerTokenAuth) (*bearerTokenAuthPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &bearerTokenAuthPb{}
	tokenPb, err := identity(&st.Token)
	if err != nil {
		return nil, err
	}
	if tokenPb != nil {
		pb.Token = *tokenPb
	}

	tokenPlaintextPb, err := identity(&st.TokenPlaintext)
	if err != nil {
		return nil, err
	}
	if tokenPlaintextPb != nil {
		pb.TokenPlaintext = *tokenPlaintextPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type bearerTokenAuthPb struct {
	// The Databricks secret key reference for a token. If you prefer to paste
	// your token directly, see `token_plaintext`.
	Token string `json:"token,omitempty"`
	// The token provided as a plaintext string. If you prefer to reference your
	// token using Databricks Secrets, see `token`.
	TokenPlaintext string `json:"token_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func bearerTokenAuthFromPb(pb *bearerTokenAuthPb) (*BearerTokenAuth, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BearerTokenAuth{}
	tokenField, err := identity(&pb.Token)
	if err != nil {
		return nil, err
	}
	if tokenField != nil {
		st.Token = *tokenField
	}
	tokenPlaintextField, err := identity(&pb.TokenPlaintext)
	if err != nil {
		return nil, err
	}
	if tokenPlaintextField != nil {
		st.TokenPlaintext = *tokenPlaintextField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *bearerTokenAuthPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st bearerTokenAuthPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get build logs for a served model
type BuildLogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName string
}

func buildLogsRequestToPb(st *BuildLogsRequest) (*buildLogsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &buildLogsRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	servedModelNamePb, err := identity(&st.ServedModelName)
	if err != nil {
		return nil, err
	}
	if servedModelNamePb != nil {
		pb.ServedModelName = *servedModelNamePb
	}

	return pb, nil
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

type buildLogsRequestPb struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" url:"-"`
	// The name of the served model that build logs will be retrieved for. This
	// field is required.
	ServedModelName string `json:"-" url:"-"`
}

func buildLogsRequestFromPb(pb *buildLogsRequestPb) (*BuildLogsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BuildLogsRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	servedModelNameField, err := identity(&pb.ServedModelName)
	if err != nil {
		return nil, err
	}
	if servedModelNameField != nil {
		st.ServedModelName = *servedModelNameField
	}

	return st, nil
}

type BuildLogsResponse struct {
	// The logs associated with building the served entity's environment.
	Logs string
}

func buildLogsResponseToPb(st *BuildLogsResponse) (*buildLogsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &buildLogsResponsePb{}
	logsPb, err := identity(&st.Logs)
	if err != nil {
		return nil, err
	}
	if logsPb != nil {
		pb.Logs = *logsPb
	}

	return pb, nil
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

type buildLogsResponsePb struct {
	// The logs associated with building the served entity's environment.
	Logs string `json:"logs"`
}

func buildLogsResponseFromPb(pb *buildLogsResponsePb) (*BuildLogsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &BuildLogsResponse{}
	logsField, err := identity(&pb.Logs)
	if err != nil {
		return nil, err
	}
	if logsField != nil {
		st.Logs = *logsField
	}

	return st, nil
}

type ChatMessage struct {
	// The content of the message.
	Content string
	// The role of the message. One of [system, user, assistant].
	Role ChatMessageRole

	ForceSendFields []string
}

func chatMessageToPb(st *ChatMessage) (*chatMessagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &chatMessagePb{}
	contentPb, err := identity(&st.Content)
	if err != nil {
		return nil, err
	}
	if contentPb != nil {
		pb.Content = *contentPb
	}

	rolePb, err := identity(&st.Role)
	if err != nil {
		return nil, err
	}
	if rolePb != nil {
		pb.Role = *rolePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type chatMessagePb struct {
	// The content of the message.
	Content string `json:"content,omitempty"`
	// The role of the message. One of [system, user, assistant].
	Role ChatMessageRole `json:"role,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func chatMessageFromPb(pb *chatMessagePb) (*ChatMessage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ChatMessage{}
	contentField, err := identity(&pb.Content)
	if err != nil {
		return nil, err
	}
	if contentField != nil {
		st.Content = *contentField
	}
	roleField, err := identity(&pb.Role)
	if err != nil {
		return nil, err
	}
	if roleField != nil {
		st.Role = *roleField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *chatMessagePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st chatMessagePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The role of the message. One of [system, user, assistant].
type ChatMessageRole string
type chatMessageRolePb string

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

func chatMessageRoleToPb(st *ChatMessageRole) (*chatMessageRolePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := chatMessageRolePb(*st)
	return &pb, nil
}

func chatMessageRoleFromPb(pb *chatMessageRolePb) (*ChatMessageRole, error) {
	if pb == nil {
		return nil, nil
	}
	st := ChatMessageRole(*pb)
	return &st, nil
}

type CohereConfig struct {
	// This is an optional field to provide a customized base URL for the Cohere
	// API. If left unspecified, the standard Cohere base URL is used.
	CohereApiBase string
	// The Databricks secret key reference for a Cohere API key. If you prefer
	// to paste your API key directly, see `cohere_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `cohere_api_key` or
	// `cohere_api_key_plaintext`.
	CohereApiKey string
	// The Cohere API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `cohere_api_key`. You
	// must provide an API key using one of the following fields:
	// `cohere_api_key` or `cohere_api_key_plaintext`.
	CohereApiKeyPlaintext string

	ForceSendFields []string
}

func cohereConfigToPb(st *CohereConfig) (*cohereConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cohereConfigPb{}
	cohereApiBasePb, err := identity(&st.CohereApiBase)
	if err != nil {
		return nil, err
	}
	if cohereApiBasePb != nil {
		pb.CohereApiBase = *cohereApiBasePb
	}

	cohereApiKeyPb, err := identity(&st.CohereApiKey)
	if err != nil {
		return nil, err
	}
	if cohereApiKeyPb != nil {
		pb.CohereApiKey = *cohereApiKeyPb
	}

	cohereApiKeyPlaintextPb, err := identity(&st.CohereApiKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if cohereApiKeyPlaintextPb != nil {
		pb.CohereApiKeyPlaintext = *cohereApiKeyPlaintextPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type cohereConfigPb struct {
	// This is an optional field to provide a customized base URL for the Cohere
	// API. If left unspecified, the standard Cohere base URL is used.
	CohereApiBase string `json:"cohere_api_base,omitempty"`
	// The Databricks secret key reference for a Cohere API key. If you prefer
	// to paste your API key directly, see `cohere_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `cohere_api_key` or
	// `cohere_api_key_plaintext`.
	CohereApiKey string `json:"cohere_api_key,omitempty"`
	// The Cohere API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `cohere_api_key`. You
	// must provide an API key using one of the following fields:
	// `cohere_api_key` or `cohere_api_key_plaintext`.
	CohereApiKeyPlaintext string `json:"cohere_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cohereConfigFromPb(pb *cohereConfigPb) (*CohereConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CohereConfig{}
	cohereApiBaseField, err := identity(&pb.CohereApiBase)
	if err != nil {
		return nil, err
	}
	if cohereApiBaseField != nil {
		st.CohereApiBase = *cohereApiBaseField
	}
	cohereApiKeyField, err := identity(&pb.CohereApiKey)
	if err != nil {
		return nil, err
	}
	if cohereApiKeyField != nil {
		st.CohereApiKey = *cohereApiKeyField
	}
	cohereApiKeyPlaintextField, err := identity(&pb.CohereApiKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if cohereApiKeyPlaintextField != nil {
		st.CohereApiKeyPlaintext = *cohereApiKeyPlaintextField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cohereConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cohereConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway *AiGatewayConfig
	// The budget policy to be applied to the serving endpoint.
	BudgetPolicyId string
	// The core config of the serving endpoint.
	Config *EndpointCoreConfigInput
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name string
	// Rate limits to be applied to the serving endpoint. NOTE: this field is
	// deprecated, please use AI Gateway to manage rate limits.
	RateLimits []RateLimit
	// Enable route optimization for the serving endpoint.
	RouteOptimized bool
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags []EndpointTag

	ForceSendFields []string
}

func createServingEndpointToPb(st *CreateServingEndpoint) (*createServingEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createServingEndpointPb{}
	aiGatewayPb, err := aiGatewayConfigToPb(st.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayPb != nil {
		pb.AiGateway = aiGatewayPb
	}

	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	configPb, err := endpointCoreConfigInputToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	var rateLimitsPb []rateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := rateLimitToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rateLimitsPb = append(rateLimitsPb, *itemPb)
		}
	}
	pb.RateLimits = rateLimitsPb

	routeOptimizedPb, err := identity(&st.RouteOptimized)
	if err != nil {
		return nil, err
	}
	if routeOptimizedPb != nil {
		pb.RouteOptimized = *routeOptimizedPb
	}

	var tagsPb []endpointTagPb
	for _, item := range st.Tags {
		itemPb, err := endpointTagToPb(&item)
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

type createServingEndpointPb struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway *aiGatewayConfigPb `json:"ai_gateway,omitempty"`
	// The budget policy to be applied to the serving endpoint.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The core config of the serving endpoint.
	Config *endpointCoreConfigInputPb `json:"config,omitempty"`
	// The name of the serving endpoint. This field is required and must be
	// unique across a Databricks workspace. An endpoint name can consist of
	// alphanumeric characters, dashes, and underscores.
	Name string `json:"name"`
	// Rate limits to be applied to the serving endpoint. NOTE: this field is
	// deprecated, please use AI Gateway to manage rate limits.
	RateLimits []rateLimitPb `json:"rate_limits,omitempty"`
	// Enable route optimization for the serving endpoint.
	RouteOptimized bool `json:"route_optimized,omitempty"`
	// Tags to be attached to the serving endpoint and automatically propagated
	// to billing logs.
	Tags []endpointTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createServingEndpointFromPb(pb *createServingEndpointPb) (*CreateServingEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateServingEndpoint{}
	aiGatewayField, err := aiGatewayConfigFromPb(pb.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayField != nil {
		st.AiGateway = aiGatewayField
	}
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	configField, err := endpointCoreConfigInputFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	var rateLimitsField []RateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := rateLimitFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rateLimitsField = append(rateLimitsField, *item)
		}
	}
	st.RateLimits = rateLimitsField
	routeOptimizedField, err := identity(&pb.RouteOptimized)
	if err != nil {
		return nil, err
	}
	if routeOptimizedField != nil {
		st.RouteOptimized = *routeOptimizedField
	}

	var tagsField []EndpointTag
	for _, itemPb := range pb.Tags {
		item, err := endpointTagFromPb(&itemPb)
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

func (st *createServingEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createServingEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Configs needed to create a custom provider model route.
type CustomProviderConfig struct {
	// This is a field to provide API key authentication for the custom provider
	// API. You can only specify one authentication method.
	ApiKeyAuth *ApiKeyAuth
	// This is a field to provide bearer token authentication for the custom
	// provider API. You can only specify one authentication method.
	BearerTokenAuth *BearerTokenAuth
	// This is a field to provide the URL of the custom provider API.
	CustomProviderUrl string
}

func customProviderConfigToPb(st *CustomProviderConfig) (*customProviderConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &customProviderConfigPb{}
	apiKeyAuthPb, err := apiKeyAuthToPb(st.ApiKeyAuth)
	if err != nil {
		return nil, err
	}
	if apiKeyAuthPb != nil {
		pb.ApiKeyAuth = apiKeyAuthPb
	}

	bearerTokenAuthPb, err := bearerTokenAuthToPb(st.BearerTokenAuth)
	if err != nil {
		return nil, err
	}
	if bearerTokenAuthPb != nil {
		pb.BearerTokenAuth = bearerTokenAuthPb
	}

	customProviderUrlPb, err := identity(&st.CustomProviderUrl)
	if err != nil {
		return nil, err
	}
	if customProviderUrlPb != nil {
		pb.CustomProviderUrl = *customProviderUrlPb
	}

	return pb, nil
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

type customProviderConfigPb struct {
	// This is a field to provide API key authentication for the custom provider
	// API. You can only specify one authentication method.
	ApiKeyAuth *apiKeyAuthPb `json:"api_key_auth,omitempty"`
	// This is a field to provide bearer token authentication for the custom
	// provider API. You can only specify one authentication method.
	BearerTokenAuth *bearerTokenAuthPb `json:"bearer_token_auth,omitempty"`
	// This is a field to provide the URL of the custom provider API.
	CustomProviderUrl string `json:"custom_provider_url"`
}

func customProviderConfigFromPb(pb *customProviderConfigPb) (*CustomProviderConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CustomProviderConfig{}
	apiKeyAuthField, err := apiKeyAuthFromPb(pb.ApiKeyAuth)
	if err != nil {
		return nil, err
	}
	if apiKeyAuthField != nil {
		st.ApiKeyAuth = apiKeyAuthField
	}
	bearerTokenAuthField, err := bearerTokenAuthFromPb(pb.BearerTokenAuth)
	if err != nil {
		return nil, err
	}
	if bearerTokenAuthField != nil {
		st.BearerTokenAuth = bearerTokenAuthField
	}
	customProviderUrlField, err := identity(&pb.CustomProviderUrl)
	if err != nil {
		return nil, err
	}
	if customProviderUrlField != nil {
		st.CustomProviderUrl = *customProviderUrlField
	}

	return st, nil
}

// Details necessary to query this object's API through the DataPlane APIs.
type DataPlaneInfo struct {
	// Authorization details as a string.
	AuthorizationDetails string
	// The URL of the endpoint for this operation in the dataplane.
	EndpointUrl string

	ForceSendFields []string
}

func dataPlaneInfoToPb(st *DataPlaneInfo) (*dataPlaneInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataPlaneInfoPb{}
	authorizationDetailsPb, err := identity(&st.AuthorizationDetails)
	if err != nil {
		return nil, err
	}
	if authorizationDetailsPb != nil {
		pb.AuthorizationDetails = *authorizationDetailsPb
	}

	endpointUrlPb, err := identity(&st.EndpointUrl)
	if err != nil {
		return nil, err
	}
	if endpointUrlPb != nil {
		pb.EndpointUrl = *endpointUrlPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type dataPlaneInfoPb struct {
	// Authorization details as a string.
	AuthorizationDetails string `json:"authorization_details,omitempty"`
	// The URL of the endpoint for this operation in the dataplane.
	EndpointUrl string `json:"endpoint_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dataPlaneInfoFromPb(pb *dataPlaneInfoPb) (*DataPlaneInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataPlaneInfo{}
	authorizationDetailsField, err := identity(&pb.AuthorizationDetails)
	if err != nil {
		return nil, err
	}
	if authorizationDetailsField != nil {
		st.AuthorizationDetails = *authorizationDetailsField
	}
	endpointUrlField, err := identity(&pb.EndpointUrl)
	if err != nil {
		return nil, err
	}
	if endpointUrlField != nil {
		st.EndpointUrl = *endpointUrlField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dataPlaneInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dataPlaneInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DatabricksModelServingConfig struct {
	// The Databricks secret key reference for a Databricks API token that
	// corresponds to a user or service principal with Can Query access to the
	// model serving endpoint pointed to by this external model. If you prefer
	// to paste your API key directly, see `databricks_api_token_plaintext`. You
	// must provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	DatabricksApiToken string
	// The Databricks API token that corresponds to a user or service principal
	// with Can Query access to the model serving endpoint pointed to by this
	// external model provided as a plaintext string. If you prefer to reference
	// your key using Databricks Secrets, see `databricks_api_token`. You must
	// provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	DatabricksApiTokenPlaintext string
	// The URL of the Databricks workspace containing the model serving endpoint
	// pointed to by this external model.
	DatabricksWorkspaceUrl string

	ForceSendFields []string
}

func databricksModelServingConfigToPb(st *DatabricksModelServingConfig) (*databricksModelServingConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &databricksModelServingConfigPb{}
	databricksApiTokenPb, err := identity(&st.DatabricksApiToken)
	if err != nil {
		return nil, err
	}
	if databricksApiTokenPb != nil {
		pb.DatabricksApiToken = *databricksApiTokenPb
	}

	databricksApiTokenPlaintextPb, err := identity(&st.DatabricksApiTokenPlaintext)
	if err != nil {
		return nil, err
	}
	if databricksApiTokenPlaintextPb != nil {
		pb.DatabricksApiTokenPlaintext = *databricksApiTokenPlaintextPb
	}

	databricksWorkspaceUrlPb, err := identity(&st.DatabricksWorkspaceUrl)
	if err != nil {
		return nil, err
	}
	if databricksWorkspaceUrlPb != nil {
		pb.DatabricksWorkspaceUrl = *databricksWorkspaceUrlPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type databricksModelServingConfigPb struct {
	// The Databricks secret key reference for a Databricks API token that
	// corresponds to a user or service principal with Can Query access to the
	// model serving endpoint pointed to by this external model. If you prefer
	// to paste your API key directly, see `databricks_api_token_plaintext`. You
	// must provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	DatabricksApiToken string `json:"databricks_api_token,omitempty"`
	// The Databricks API token that corresponds to a user or service principal
	// with Can Query access to the model serving endpoint pointed to by this
	// external model provided as a plaintext string. If you prefer to reference
	// your key using Databricks Secrets, see `databricks_api_token`. You must
	// provide an API key using one of the following fields:
	// `databricks_api_token` or `databricks_api_token_plaintext`.
	DatabricksApiTokenPlaintext string `json:"databricks_api_token_plaintext,omitempty"`
	// The URL of the Databricks workspace containing the model serving endpoint
	// pointed to by this external model.
	DatabricksWorkspaceUrl string `json:"databricks_workspace_url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func databricksModelServingConfigFromPb(pb *databricksModelServingConfigPb) (*DatabricksModelServingConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatabricksModelServingConfig{}
	databricksApiTokenField, err := identity(&pb.DatabricksApiToken)
	if err != nil {
		return nil, err
	}
	if databricksApiTokenField != nil {
		st.DatabricksApiToken = *databricksApiTokenField
	}
	databricksApiTokenPlaintextField, err := identity(&pb.DatabricksApiTokenPlaintext)
	if err != nil {
		return nil, err
	}
	if databricksApiTokenPlaintextField != nil {
		st.DatabricksApiTokenPlaintext = *databricksApiTokenPlaintextField
	}
	databricksWorkspaceUrlField, err := identity(&pb.DatabricksWorkspaceUrl)
	if err != nil {
		return nil, err
	}
	if databricksWorkspaceUrlField != nil {
		st.DatabricksWorkspaceUrl = *databricksWorkspaceUrlField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *databricksModelServingConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st databricksModelServingConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DataframeSplitInput struct {
	Columns []any

	Data []any

	Index []int
}

func dataframeSplitInputToPb(st *DataframeSplitInput) (*dataframeSplitInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataframeSplitInputPb{}

	var columnsPb []any
	for _, item := range st.Columns {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			columnsPb = append(columnsPb, *itemPb)
		}
	}
	pb.Columns = columnsPb

	var dataPb []any
	for _, item := range st.Data {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataPb = append(dataPb, *itemPb)
		}
	}
	pb.Data = dataPb

	var indexPb []int
	for _, item := range st.Index {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			indexPb = append(indexPb, *itemPb)
		}
	}
	pb.Index = indexPb

	return pb, nil
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

	var columnsField []any
	for _, itemPb := range pb.Columns {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			columnsField = append(columnsField, *item)
		}
	}
	st.Columns = columnsField

	var dataField []any
	for _, itemPb := range pb.Data {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dataField = append(dataField, *item)
		}
	}
	st.Data = dataField

	var indexField []int
	for _, itemPb := range pb.Index {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			indexField = append(indexField, *item)
		}
	}
	st.Index = indexField

	return st, nil
}

type DeleteResponse struct {
}

func deleteResponseToPb(st *DeleteResponse) (*deleteResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteResponsePb{}

	return pb, nil
}

func (st *DeleteResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Delete a serving endpoint
type DeleteServingEndpointRequest struct {
	Name string
}

func deleteServingEndpointRequestToPb(st *DeleteServingEndpointRequest) (*deleteServingEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteServingEndpointRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
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

type deleteServingEndpointRequestPb struct {
	Name string `json:"-" url:"-"`
}

func deleteServingEndpointRequestFromPb(pb *deleteServingEndpointRequestPb) (*DeleteServingEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteServingEndpointRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

type EmbeddingsV1ResponseEmbeddingElement struct {
	Embedding []float64
	// The index of the embedding in the response.
	Index int
	// This will always be 'embedding'.
	Object EmbeddingsV1ResponseEmbeddingElementObject

	ForceSendFields []string
}

func embeddingsV1ResponseEmbeddingElementToPb(st *EmbeddingsV1ResponseEmbeddingElement) (*embeddingsV1ResponseEmbeddingElementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &embeddingsV1ResponseEmbeddingElementPb{}

	var embeddingPb []float64
	for _, item := range st.Embedding {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			embeddingPb = append(embeddingPb, *itemPb)
		}
	}
	pb.Embedding = embeddingPb

	indexPb, err := identity(&st.Index)
	if err != nil {
		return nil, err
	}
	if indexPb != nil {
		pb.Index = *indexPb
	}

	objectPb, err := identity(&st.Object)
	if err != nil {
		return nil, err
	}
	if objectPb != nil {
		pb.Object = *objectPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type embeddingsV1ResponseEmbeddingElementPb struct {
	Embedding []float64 `json:"embedding,omitempty"`
	// The index of the embedding in the response.
	Index int `json:"index,omitempty"`
	// This will always be 'embedding'.
	Object EmbeddingsV1ResponseEmbeddingElementObject `json:"object,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func embeddingsV1ResponseEmbeddingElementFromPb(pb *embeddingsV1ResponseEmbeddingElementPb) (*EmbeddingsV1ResponseEmbeddingElement, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EmbeddingsV1ResponseEmbeddingElement{}

	var embeddingField []float64
	for _, itemPb := range pb.Embedding {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			embeddingField = append(embeddingField, *item)
		}
	}
	st.Embedding = embeddingField
	indexField, err := identity(&pb.Index)
	if err != nil {
		return nil, err
	}
	if indexField != nil {
		st.Index = *indexField
	}
	objectField, err := identity(&pb.Object)
	if err != nil {
		return nil, err
	}
	if objectField != nil {
		st.Object = *objectField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *embeddingsV1ResponseEmbeddingElementPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st embeddingsV1ResponseEmbeddingElementPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// This will always be 'embedding'.
type EmbeddingsV1ResponseEmbeddingElementObject string
type embeddingsV1ResponseEmbeddingElementObjectPb string

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

func embeddingsV1ResponseEmbeddingElementObjectToPb(st *EmbeddingsV1ResponseEmbeddingElementObject) (*embeddingsV1ResponseEmbeddingElementObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := embeddingsV1ResponseEmbeddingElementObjectPb(*st)
	return &pb, nil
}

func embeddingsV1ResponseEmbeddingElementObjectFromPb(pb *embeddingsV1ResponseEmbeddingElementObjectPb) (*EmbeddingsV1ResponseEmbeddingElementObject, error) {
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
	AutoCaptureConfig *AutoCaptureConfigInput
	// The name of the serving endpoint to update. This field is required.
	Name string
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntityInput
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelInput
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *TrafficConfig
}

func endpointCoreConfigInputToPb(st *EndpointCoreConfigInput) (*endpointCoreConfigInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointCoreConfigInputPb{}
	autoCaptureConfigPb, err := autoCaptureConfigInputToPb(st.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigPb != nil {
		pb.AutoCaptureConfig = autoCaptureConfigPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	var servedEntitiesPb []servedEntityInputPb
	for _, item := range st.ServedEntities {
		itemPb, err := servedEntityInputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedEntitiesPb = append(servedEntitiesPb, *itemPb)
		}
	}
	pb.ServedEntities = servedEntitiesPb

	var servedModelsPb []servedModelInputPb
	for _, item := range st.ServedModels {
		itemPb, err := servedModelInputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedModelsPb = append(servedModelsPb, *itemPb)
		}
	}
	pb.ServedModels = servedModelsPb

	trafficConfigPb, err := trafficConfigToPb(st.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigPb != nil {
		pb.TrafficConfig = trafficConfigPb
	}

	return pb, nil
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

type endpointCoreConfigInputPb struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	AutoCaptureConfig *autoCaptureConfigInputPb `json:"auto_capture_config,omitempty"`
	// The name of the serving endpoint to update. This field is required.
	Name string `json:"-" url:"-"`
	// The list of served entities under the serving endpoint config.
	ServedEntities []servedEntityInputPb `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []servedModelInputPb `json:"served_models,omitempty"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *trafficConfigPb `json:"traffic_config,omitempty"`
}

func endpointCoreConfigInputFromPb(pb *endpointCoreConfigInputPb) (*EndpointCoreConfigInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointCoreConfigInput{}
	autoCaptureConfigField, err := autoCaptureConfigInputFromPb(pb.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigField != nil {
		st.AutoCaptureConfig = autoCaptureConfigField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	var servedEntitiesField []ServedEntityInput
	for _, itemPb := range pb.ServedEntities {
		item, err := servedEntityInputFromPb(&itemPb)
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
		item, err := servedModelInputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedModelsField = append(servedModelsField, *item)
		}
	}
	st.ServedModels = servedModelsField
	trafficConfigField, err := trafficConfigFromPb(pb.TrafficConfig)
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
	AutoCaptureConfig *AutoCaptureConfigOutput
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int64
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntityOutput
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelOutput
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *TrafficConfig

	ForceSendFields []string
}

func endpointCoreConfigOutputToPb(st *EndpointCoreConfigOutput) (*endpointCoreConfigOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointCoreConfigOutputPb{}
	autoCaptureConfigPb, err := autoCaptureConfigOutputToPb(st.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigPb != nil {
		pb.AutoCaptureConfig = autoCaptureConfigPb
	}

	configVersionPb, err := identity(&st.ConfigVersion)
	if err != nil {
		return nil, err
	}
	if configVersionPb != nil {
		pb.ConfigVersion = *configVersionPb
	}

	var servedEntitiesPb []servedEntityOutputPb
	for _, item := range st.ServedEntities {
		itemPb, err := servedEntityOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedEntitiesPb = append(servedEntitiesPb, *itemPb)
		}
	}
	pb.ServedEntities = servedEntitiesPb

	var servedModelsPb []servedModelOutputPb
	for _, item := range st.ServedModels {
		itemPb, err := servedModelOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedModelsPb = append(servedModelsPb, *itemPb)
		}
	}
	pb.ServedModels = servedModelsPb

	trafficConfigPb, err := trafficConfigToPb(st.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigPb != nil {
		pb.TrafficConfig = trafficConfigPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type endpointCoreConfigOutputPb struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	AutoCaptureConfig *autoCaptureConfigOutputPb `json:"auto_capture_config,omitempty"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int64 `json:"config_version,omitempty"`
	// The list of served entities under the serving endpoint config.
	ServedEntities []servedEntityOutputPb `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []servedModelOutputPb `json:"served_models,omitempty"`
	// The traffic configuration associated with the serving endpoint config.
	TrafficConfig *trafficConfigPb `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointCoreConfigOutputFromPb(pb *endpointCoreConfigOutputPb) (*EndpointCoreConfigOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointCoreConfigOutput{}
	autoCaptureConfigField, err := autoCaptureConfigOutputFromPb(pb.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigField != nil {
		st.AutoCaptureConfig = autoCaptureConfigField
	}
	configVersionField, err := identity(&pb.ConfigVersion)
	if err != nil {
		return nil, err
	}
	if configVersionField != nil {
		st.ConfigVersion = *configVersionField
	}

	var servedEntitiesField []ServedEntityOutput
	for _, itemPb := range pb.ServedEntities {
		item, err := servedEntityOutputFromPb(&itemPb)
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
		item, err := servedModelOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedModelsField = append(servedModelsField, *item)
		}
	}
	st.ServedModels = servedModelsField
	trafficConfigField, err := trafficConfigFromPb(pb.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigField != nil {
		st.TrafficConfig = trafficConfigField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointCoreConfigOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointCoreConfigOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointCoreConfigSummary struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities []ServedEntitySpec
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []ServedModelSpec
}

func endpointCoreConfigSummaryToPb(st *EndpointCoreConfigSummary) (*endpointCoreConfigSummaryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointCoreConfigSummaryPb{}

	var servedEntitiesPb []servedEntitySpecPb
	for _, item := range st.ServedEntities {
		itemPb, err := servedEntitySpecToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedEntitiesPb = append(servedEntitiesPb, *itemPb)
		}
	}
	pb.ServedEntities = servedEntitiesPb

	var servedModelsPb []servedModelSpecPb
	for _, item := range st.ServedModels {
		itemPb, err := servedModelSpecToPb(&item)
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

type endpointCoreConfigSummaryPb struct {
	// The list of served entities under the serving endpoint config.
	ServedEntities []servedEntitySpecPb `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models under
	// the serving endpoint config.
	ServedModels []servedModelSpecPb `json:"served_models,omitempty"`
}

func endpointCoreConfigSummaryFromPb(pb *endpointCoreConfigSummaryPb) (*EndpointCoreConfigSummary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointCoreConfigSummary{}

	var servedEntitiesField []ServedEntitySpec
	for _, itemPb := range pb.ServedEntities {
		item, err := servedEntitySpecFromPb(&itemPb)
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
		item, err := servedModelSpecFromPb(&itemPb)
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
	AutoCaptureConfig *AutoCaptureConfigOutput
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int
	// The list of served entities belonging to the last issued update to the
	// serving endpoint.
	ServedEntities []ServedEntityOutput
	// (Deprecated, use served_entities instead) The list of served models
	// belonging to the last issued update to the serving endpoint.
	ServedModels []ServedModelOutput
	// The timestamp when the update to the pending config started.
	StartTime int64
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *TrafficConfig

	ForceSendFields []string
}

func endpointPendingConfigToPb(st *EndpointPendingConfig) (*endpointPendingConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointPendingConfigPb{}
	autoCaptureConfigPb, err := autoCaptureConfigOutputToPb(st.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigPb != nil {
		pb.AutoCaptureConfig = autoCaptureConfigPb
	}

	configVersionPb, err := identity(&st.ConfigVersion)
	if err != nil {
		return nil, err
	}
	if configVersionPb != nil {
		pb.ConfigVersion = *configVersionPb
	}

	var servedEntitiesPb []servedEntityOutputPb
	for _, item := range st.ServedEntities {
		itemPb, err := servedEntityOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedEntitiesPb = append(servedEntitiesPb, *itemPb)
		}
	}
	pb.ServedEntities = servedEntitiesPb

	var servedModelsPb []servedModelOutputPb
	for _, item := range st.ServedModels {
		itemPb, err := servedModelOutputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			servedModelsPb = append(servedModelsPb, *itemPb)
		}
	}
	pb.ServedModels = servedModelsPb

	startTimePb, err := identity(&st.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimePb != nil {
		pb.StartTime = *startTimePb
	}

	trafficConfigPb, err := trafficConfigToPb(st.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigPb != nil {
		pb.TrafficConfig = trafficConfigPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type endpointPendingConfigPb struct {
	// Configuration for Inference Tables which automatically logs requests and
	// responses to Unity Catalog. Note: this field is deprecated for creating
	// new provisioned throughput endpoints, or updating existing provisioned
	// throughput endpoints that never have inference table configured; in these
	// cases please use AI Gateway to manage inference tables.
	AutoCaptureConfig *autoCaptureConfigOutputPb `json:"auto_capture_config,omitempty"`
	// The config version that the serving endpoint is currently serving.
	ConfigVersion int `json:"config_version,omitempty"`
	// The list of served entities belonging to the last issued update to the
	// serving endpoint.
	ServedEntities []servedEntityOutputPb `json:"served_entities,omitempty"`
	// (Deprecated, use served_entities instead) The list of served models
	// belonging to the last issued update to the serving endpoint.
	ServedModels []servedModelOutputPb `json:"served_models,omitempty"`
	// The timestamp when the update to the pending config started.
	StartTime int64 `json:"start_time,omitempty"`
	// The traffic config defining how invocations to the serving endpoint
	// should be routed.
	TrafficConfig *trafficConfigPb `json:"traffic_config,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointPendingConfigFromPb(pb *endpointPendingConfigPb) (*EndpointPendingConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointPendingConfig{}
	autoCaptureConfigField, err := autoCaptureConfigOutputFromPb(pb.AutoCaptureConfig)
	if err != nil {
		return nil, err
	}
	if autoCaptureConfigField != nil {
		st.AutoCaptureConfig = autoCaptureConfigField
	}
	configVersionField, err := identity(&pb.ConfigVersion)
	if err != nil {
		return nil, err
	}
	if configVersionField != nil {
		st.ConfigVersion = *configVersionField
	}

	var servedEntitiesField []ServedEntityOutput
	for _, itemPb := range pb.ServedEntities {
		item, err := servedEntityOutputFromPb(&itemPb)
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
		item, err := servedModelOutputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			servedModelsField = append(servedModelsField, *item)
		}
	}
	st.ServedModels = servedModelsField
	startTimeField, err := identity(&pb.StartTime)
	if err != nil {
		return nil, err
	}
	if startTimeField != nil {
		st.StartTime = *startTimeField
	}
	trafficConfigField, err := trafficConfigFromPb(pb.TrafficConfig)
	if err != nil {
		return nil, err
	}
	if trafficConfigField != nil {
		st.TrafficConfig = trafficConfigField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointPendingConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointPendingConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointState struct {
	// The state of an endpoint's config update. This informs the user if the
	// pending_config is in progress, if the update failed, or if there is no
	// update in progress. Note that if the endpoint's config_update state value
	// is IN_PROGRESS, another update can not be made until the update completes
	// or fails.
	ConfigUpdate EndpointStateConfigUpdate
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served entities in its
	// active configuration are ready. If any of the actively served entities
	// are in a non-ready state, the endpoint state will be NOT_READY.
	Ready EndpointStateReady
}

func endpointStateToPb(st *EndpointState) (*endpointStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointStatePb{}
	configUpdatePb, err := identity(&st.ConfigUpdate)
	if err != nil {
		return nil, err
	}
	if configUpdatePb != nil {
		pb.ConfigUpdate = *configUpdatePb
	}

	readyPb, err := identity(&st.Ready)
	if err != nil {
		return nil, err
	}
	if readyPb != nil {
		pb.Ready = *readyPb
	}

	return pb, nil
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

type endpointStatePb struct {
	// The state of an endpoint's config update. This informs the user if the
	// pending_config is in progress, if the update failed, or if there is no
	// update in progress. Note that if the endpoint's config_update state value
	// is IN_PROGRESS, another update can not be made until the update completes
	// or fails.
	ConfigUpdate EndpointStateConfigUpdate `json:"config_update,omitempty"`
	// The state of an endpoint, indicating whether or not the endpoint is
	// queryable. An endpoint is READY if all of the served entities in its
	// active configuration are ready. If any of the actively served entities
	// are in a non-ready state, the endpoint state will be NOT_READY.
	Ready EndpointStateReady `json:"ready,omitempty"`
}

func endpointStateFromPb(pb *endpointStatePb) (*EndpointState, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointState{}
	configUpdateField, err := identity(&pb.ConfigUpdate)
	if err != nil {
		return nil, err
	}
	if configUpdateField != nil {
		st.ConfigUpdate = *configUpdateField
	}
	readyField, err := identity(&pb.Ready)
	if err != nil {
		return nil, err
	}
	if readyField != nil {
		st.Ready = *readyField
	}

	return st, nil
}

type EndpointStateConfigUpdate string
type endpointStateConfigUpdatePb string

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

// Type always returns EndpointStateConfigUpdate to satisfy [pflag.Value] interface
func (f *EndpointStateConfigUpdate) Type() string {
	return "EndpointStateConfigUpdate"
}

func endpointStateConfigUpdateToPb(st *EndpointStateConfigUpdate) (*endpointStateConfigUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := endpointStateConfigUpdatePb(*st)
	return &pb, nil
}

func endpointStateConfigUpdateFromPb(pb *endpointStateConfigUpdatePb) (*EndpointStateConfigUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointStateConfigUpdate(*pb)
	return &st, nil
}

type EndpointStateReady string
type endpointStateReadyPb string

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

func endpointStateReadyToPb(st *EndpointStateReady) (*endpointStateReadyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := endpointStateReadyPb(*st)
	return &pb, nil
}

func endpointStateReadyFromPb(pb *endpointStateReadyPb) (*EndpointStateReady, error) {
	if pb == nil {
		return nil, nil
	}
	st := EndpointStateReady(*pb)
	return &st, nil
}

type EndpointTag struct {
	// Key field for a serving endpoint tag.
	Key string
	// Optional value field for a serving endpoint tag.
	Value string

	ForceSendFields []string
}

func endpointTagToPb(st *EndpointTag) (*endpointTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointTagPb{}
	keyPb, err := identity(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}

	valuePb, err := identity(&st.Value)
	if err != nil {
		return nil, err
	}
	if valuePb != nil {
		pb.Value = *valuePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type endpointTagPb struct {
	// Key field for a serving endpoint tag.
	Key string `json:"key"`
	// Optional value field for a serving endpoint tag.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func endpointTagFromPb(pb *endpointTagPb) (*EndpointTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTag{}
	keyField, err := identity(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
	valueField, err := identity(&pb.Value)
	if err != nil {
		return nil, err
	}
	if valueField != nil {
		st.Value = *valueField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *endpointTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st endpointTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EndpointTags struct {
	Tags []EndpointTag
}

func endpointTagsToPb(st *EndpointTags) (*endpointTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &endpointTagsPb{}

	var tagsPb []endpointTagPb
	for _, item := range st.Tags {
		itemPb, err := endpointTagToPb(&item)
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

type endpointTagsPb struct {
	Tags []endpointTagPb `json:"tags,omitempty"`
}

func endpointTagsFromPb(pb *endpointTagsPb) (*EndpointTags, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EndpointTags{}

	var tagsField []EndpointTag
	for _, itemPb := range pb.Tags {
		item, err := endpointTagFromPb(&itemPb)
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

// Get metrics of a serving endpoint
type ExportMetricsRequest struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name string
}

func exportMetricsRequestToPb(st *ExportMetricsRequest) (*exportMetricsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportMetricsRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
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

type exportMetricsRequestPb struct {
	// The name of the serving endpoint to retrieve metrics for. This field is
	// required.
	Name string `json:"-" url:"-"`
}

func exportMetricsRequestFromPb(pb *exportMetricsRequestPb) (*ExportMetricsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportMetricsRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

type ExportMetricsResponse struct {
	Contents io.ReadCloser
}

func exportMetricsResponseToPb(st *ExportMetricsResponse) (*exportMetricsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &exportMetricsResponsePb{}
	contentsPb, err := identity(&st.Contents)
	if err != nil {
		return nil, err
	}
	if contentsPb != nil {
		pb.Contents = *contentsPb
	}

	return pb, nil
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

type exportMetricsResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

func exportMetricsResponseFromPb(pb *exportMetricsResponsePb) (*ExportMetricsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExportMetricsResponse{}
	contentsField, err := identity(&pb.Contents)
	if err != nil {
		return nil, err
	}
	if contentsField != nil {
		st.Contents = *contentsField
	}

	return st, nil
}

// Simple Proto message for testing
type ExternalFunctionRequest struct {
	// The connection name to use. This is required to identify the external
	// connection.
	ConnectionName string
	// Additional headers for the request. If not provided, only auth headers
	// from connections would be passed.
	Headers string
	// The JSON payload to send in the request body.
	Json string
	// The HTTP method to use (e.g., 'GET', 'POST').
	Method ExternalFunctionRequestHttpMethod
	// Query parameters for the request.
	Params string
	// The relative path for the API endpoint. This is required.
	Path string

	ForceSendFields []string
}

func externalFunctionRequestToPb(st *ExternalFunctionRequest) (*externalFunctionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalFunctionRequestPb{}
	connectionNamePb, err := identity(&st.ConnectionName)
	if err != nil {
		return nil, err
	}
	if connectionNamePb != nil {
		pb.ConnectionName = *connectionNamePb
	}

	headersPb, err := identity(&st.Headers)
	if err != nil {
		return nil, err
	}
	if headersPb != nil {
		pb.Headers = *headersPb
	}

	jsonPb, err := identity(&st.Json)
	if err != nil {
		return nil, err
	}
	if jsonPb != nil {
		pb.Json = *jsonPb
	}

	methodPb, err := identity(&st.Method)
	if err != nil {
		return nil, err
	}
	if methodPb != nil {
		pb.Method = *methodPb
	}

	paramsPb, err := identity(&st.Params)
	if err != nil {
		return nil, err
	}
	if paramsPb != nil {
		pb.Params = *paramsPb
	}

	pathPb, err := identity(&st.Path)
	if err != nil {
		return nil, err
	}
	if pathPb != nil {
		pb.Path = *pathPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type externalFunctionRequestPb struct {
	// The connection name to use. This is required to identify the external
	// connection.
	ConnectionName string `json:"connection_name"`
	// Additional headers for the request. If not provided, only auth headers
	// from connections would be passed.
	Headers string `json:"headers,omitempty"`
	// The JSON payload to send in the request body.
	Json string `json:"json,omitempty"`
	// The HTTP method to use (e.g., 'GET', 'POST').
	Method ExternalFunctionRequestHttpMethod `json:"method"`
	// Query parameters for the request.
	Params string `json:"params,omitempty"`
	// The relative path for the API endpoint. This is required.
	Path string `json:"path"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalFunctionRequestFromPb(pb *externalFunctionRequestPb) (*ExternalFunctionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalFunctionRequest{}
	connectionNameField, err := identity(&pb.ConnectionName)
	if err != nil {
		return nil, err
	}
	if connectionNameField != nil {
		st.ConnectionName = *connectionNameField
	}
	headersField, err := identity(&pb.Headers)
	if err != nil {
		return nil, err
	}
	if headersField != nil {
		st.Headers = *headersField
	}
	jsonField, err := identity(&pb.Json)
	if err != nil {
		return nil, err
	}
	if jsonField != nil {
		st.Json = *jsonField
	}
	methodField, err := identity(&pb.Method)
	if err != nil {
		return nil, err
	}
	if methodField != nil {
		st.Method = *methodField
	}
	paramsField, err := identity(&pb.Params)
	if err != nil {
		return nil, err
	}
	if paramsField != nil {
		st.Params = *paramsField
	}
	pathField, err := identity(&pb.Path)
	if err != nil {
		return nil, err
	}
	if pathField != nil {
		st.Path = *pathField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalFunctionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalFunctionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExternalFunctionRequestHttpMethod string
type externalFunctionRequestHttpMethodPb string

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

// Type always returns ExternalFunctionRequestHttpMethod to satisfy [pflag.Value] interface
func (f *ExternalFunctionRequestHttpMethod) Type() string {
	return "ExternalFunctionRequestHttpMethod"
}

func externalFunctionRequestHttpMethodToPb(st *ExternalFunctionRequestHttpMethod) (*externalFunctionRequestHttpMethodPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := externalFunctionRequestHttpMethodPb(*st)
	return &pb, nil
}

func externalFunctionRequestHttpMethodFromPb(pb *externalFunctionRequestHttpMethodPb) (*ExternalFunctionRequestHttpMethod, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExternalFunctionRequestHttpMethod(*pb)
	return &st, nil
}

type ExternalModel struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	Ai21labsConfig *Ai21LabsConfig
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	AmazonBedrockConfig *AmazonBedrockConfig
	// Anthropic Config. Only required if the provider is 'anthropic'.
	AnthropicConfig *AnthropicConfig
	// Cohere Config. Only required if the provider is 'cohere'.
	CohereConfig *CohereConfig
	// Custom Provider Config. Only required if the provider is 'custom'.
	CustomProviderConfig *CustomProviderConfig
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	DatabricksModelServingConfig *DatabricksModelServingConfig
	// Google Cloud Vertex AI Config. Only required if the provider is
	// 'google-cloud-vertex-ai'.
	GoogleCloudVertexAiConfig *GoogleCloudVertexAiConfig
	// The name of the external model.
	Name string
	// OpenAI Config. Only required if the provider is 'openai'.
	OpenaiConfig *OpenAiConfig
	// PaLM Config. Only required if the provider is 'palm'.
	PalmConfig *PaLmConfig
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
	// 'databricks-model-serving', 'google-cloud-vertex-ai', 'openai', 'palm',
	// and 'custom'.
	Provider ExternalModelProvider
	// The task type of the external model.
	Task string
}

func externalModelToPb(st *ExternalModel) (*externalModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalModelPb{}
	ai21labsConfigPb, err := ai21LabsConfigToPb(st.Ai21labsConfig)
	if err != nil {
		return nil, err
	}
	if ai21labsConfigPb != nil {
		pb.Ai21labsConfig = ai21labsConfigPb
	}

	amazonBedrockConfigPb, err := amazonBedrockConfigToPb(st.AmazonBedrockConfig)
	if err != nil {
		return nil, err
	}
	if amazonBedrockConfigPb != nil {
		pb.AmazonBedrockConfig = amazonBedrockConfigPb
	}

	anthropicConfigPb, err := anthropicConfigToPb(st.AnthropicConfig)
	if err != nil {
		return nil, err
	}
	if anthropicConfigPb != nil {
		pb.AnthropicConfig = anthropicConfigPb
	}

	cohereConfigPb, err := cohereConfigToPb(st.CohereConfig)
	if err != nil {
		return nil, err
	}
	if cohereConfigPb != nil {
		pb.CohereConfig = cohereConfigPb
	}

	customProviderConfigPb, err := customProviderConfigToPb(st.CustomProviderConfig)
	if err != nil {
		return nil, err
	}
	if customProviderConfigPb != nil {
		pb.CustomProviderConfig = customProviderConfigPb
	}

	databricksModelServingConfigPb, err := databricksModelServingConfigToPb(st.DatabricksModelServingConfig)
	if err != nil {
		return nil, err
	}
	if databricksModelServingConfigPb != nil {
		pb.DatabricksModelServingConfig = databricksModelServingConfigPb
	}

	googleCloudVertexAiConfigPb, err := googleCloudVertexAiConfigToPb(st.GoogleCloudVertexAiConfig)
	if err != nil {
		return nil, err
	}
	if googleCloudVertexAiConfigPb != nil {
		pb.GoogleCloudVertexAiConfig = googleCloudVertexAiConfigPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	openaiConfigPb, err := openAiConfigToPb(st.OpenaiConfig)
	if err != nil {
		return nil, err
	}
	if openaiConfigPb != nil {
		pb.OpenaiConfig = openaiConfigPb
	}

	palmConfigPb, err := paLmConfigToPb(st.PalmConfig)
	if err != nil {
		return nil, err
	}
	if palmConfigPb != nil {
		pb.PalmConfig = palmConfigPb
	}

	providerPb, err := identity(&st.Provider)
	if err != nil {
		return nil, err
	}
	if providerPb != nil {
		pb.Provider = *providerPb
	}

	taskPb, err := identity(&st.Task)
	if err != nil {
		return nil, err
	}
	if taskPb != nil {
		pb.Task = *taskPb
	}

	return pb, nil
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

type externalModelPb struct {
	// AI21Labs Config. Only required if the provider is 'ai21labs'.
	Ai21labsConfig *ai21LabsConfigPb `json:"ai21labs_config,omitempty"`
	// Amazon Bedrock Config. Only required if the provider is 'amazon-bedrock'.
	AmazonBedrockConfig *amazonBedrockConfigPb `json:"amazon_bedrock_config,omitempty"`
	// Anthropic Config. Only required if the provider is 'anthropic'.
	AnthropicConfig *anthropicConfigPb `json:"anthropic_config,omitempty"`
	// Cohere Config. Only required if the provider is 'cohere'.
	CohereConfig *cohereConfigPb `json:"cohere_config,omitempty"`
	// Custom Provider Config. Only required if the provider is 'custom'.
	CustomProviderConfig *customProviderConfigPb `json:"custom_provider_config,omitempty"`
	// Databricks Model Serving Config. Only required if the provider is
	// 'databricks-model-serving'.
	DatabricksModelServingConfig *databricksModelServingConfigPb `json:"databricks_model_serving_config,omitempty"`
	// Google Cloud Vertex AI Config. Only required if the provider is
	// 'google-cloud-vertex-ai'.
	GoogleCloudVertexAiConfig *googleCloudVertexAiConfigPb `json:"google_cloud_vertex_ai_config,omitempty"`
	// The name of the external model.
	Name string `json:"name"`
	// OpenAI Config. Only required if the provider is 'openai'.
	OpenaiConfig *openAiConfigPb `json:"openai_config,omitempty"`
	// PaLM Config. Only required if the provider is 'palm'.
	PalmConfig *paLmConfigPb `json:"palm_config,omitempty"`
	// The name of the provider for the external model. Currently, the supported
	// providers are 'ai21labs', 'anthropic', 'amazon-bedrock', 'cohere',
	// 'databricks-model-serving', 'google-cloud-vertex-ai', 'openai', 'palm',
	// and 'custom'.
	Provider ExternalModelProvider `json:"provider"`
	// The task type of the external model.
	Task string `json:"task"`
}

func externalModelFromPb(pb *externalModelPb) (*ExternalModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalModel{}
	ai21labsConfigField, err := ai21LabsConfigFromPb(pb.Ai21labsConfig)
	if err != nil {
		return nil, err
	}
	if ai21labsConfigField != nil {
		st.Ai21labsConfig = ai21labsConfigField
	}
	amazonBedrockConfigField, err := amazonBedrockConfigFromPb(pb.AmazonBedrockConfig)
	if err != nil {
		return nil, err
	}
	if amazonBedrockConfigField != nil {
		st.AmazonBedrockConfig = amazonBedrockConfigField
	}
	anthropicConfigField, err := anthropicConfigFromPb(pb.AnthropicConfig)
	if err != nil {
		return nil, err
	}
	if anthropicConfigField != nil {
		st.AnthropicConfig = anthropicConfigField
	}
	cohereConfigField, err := cohereConfigFromPb(pb.CohereConfig)
	if err != nil {
		return nil, err
	}
	if cohereConfigField != nil {
		st.CohereConfig = cohereConfigField
	}
	customProviderConfigField, err := customProviderConfigFromPb(pb.CustomProviderConfig)
	if err != nil {
		return nil, err
	}
	if customProviderConfigField != nil {
		st.CustomProviderConfig = customProviderConfigField
	}
	databricksModelServingConfigField, err := databricksModelServingConfigFromPb(pb.DatabricksModelServingConfig)
	if err != nil {
		return nil, err
	}
	if databricksModelServingConfigField != nil {
		st.DatabricksModelServingConfig = databricksModelServingConfigField
	}
	googleCloudVertexAiConfigField, err := googleCloudVertexAiConfigFromPb(pb.GoogleCloudVertexAiConfig)
	if err != nil {
		return nil, err
	}
	if googleCloudVertexAiConfigField != nil {
		st.GoogleCloudVertexAiConfig = googleCloudVertexAiConfigField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	openaiConfigField, err := openAiConfigFromPb(pb.OpenaiConfig)
	if err != nil {
		return nil, err
	}
	if openaiConfigField != nil {
		st.OpenaiConfig = openaiConfigField
	}
	palmConfigField, err := paLmConfigFromPb(pb.PalmConfig)
	if err != nil {
		return nil, err
	}
	if palmConfigField != nil {
		st.PalmConfig = palmConfigField
	}
	providerField, err := identity(&pb.Provider)
	if err != nil {
		return nil, err
	}
	if providerField != nil {
		st.Provider = *providerField
	}
	taskField, err := identity(&pb.Task)
	if err != nil {
		return nil, err
	}
	if taskField != nil {
		st.Task = *taskField
	}

	return st, nil
}

type ExternalModelProvider string
type externalModelProviderPb string

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

// Type always returns ExternalModelProvider to satisfy [pflag.Value] interface
func (f *ExternalModelProvider) Type() string {
	return "ExternalModelProvider"
}

func externalModelProviderToPb(st *ExternalModelProvider) (*externalModelProviderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := externalModelProviderPb(*st)
	return &pb, nil
}

func externalModelProviderFromPb(pb *externalModelProviderPb) (*ExternalModelProvider, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExternalModelProvider(*pb)
	return &st, nil
}

type ExternalModelUsageElement struct {
	// The number of tokens in the chat/completions response.
	CompletionTokens int
	// The number of tokens in the prompt.
	PromptTokens int
	// The total number of tokens in the prompt and response.
	TotalTokens int

	ForceSendFields []string
}

func externalModelUsageElementToPb(st *ExternalModelUsageElement) (*externalModelUsageElementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &externalModelUsageElementPb{}
	completionTokensPb, err := identity(&st.CompletionTokens)
	if err != nil {
		return nil, err
	}
	if completionTokensPb != nil {
		pb.CompletionTokens = *completionTokensPb
	}

	promptTokensPb, err := identity(&st.PromptTokens)
	if err != nil {
		return nil, err
	}
	if promptTokensPb != nil {
		pb.PromptTokens = *promptTokensPb
	}

	totalTokensPb, err := identity(&st.TotalTokens)
	if err != nil {
		return nil, err
	}
	if totalTokensPb != nil {
		pb.TotalTokens = *totalTokensPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type externalModelUsageElementPb struct {
	// The number of tokens in the chat/completions response.
	CompletionTokens int `json:"completion_tokens,omitempty"`
	// The number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens,omitempty"`
	// The total number of tokens in the prompt and response.
	TotalTokens int `json:"total_tokens,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func externalModelUsageElementFromPb(pb *externalModelUsageElementPb) (*ExternalModelUsageElement, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExternalModelUsageElement{}
	completionTokensField, err := identity(&pb.CompletionTokens)
	if err != nil {
		return nil, err
	}
	if completionTokensField != nil {
		st.CompletionTokens = *completionTokensField
	}
	promptTokensField, err := identity(&pb.PromptTokens)
	if err != nil {
		return nil, err
	}
	if promptTokensField != nil {
		st.PromptTokens = *promptTokensField
	}
	totalTokensField, err := identity(&pb.TotalTokens)
	if err != nil {
		return nil, err
	}
	if totalTokensField != nil {
		st.TotalTokens = *totalTokensField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *externalModelUsageElementPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st externalModelUsageElementPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FallbackConfig struct {
	// Whether to enable traffic fallback. When a served entity in the serving
	// endpoint returns specific error codes (e.g. 500), the request will
	// automatically be round-robin attempted with other served entities in the
	// same endpoint, following the order of served entity list, until a
	// successful response is returned. If all attempts fail, return the last
	// response with the error code.
	Enabled bool
}

func fallbackConfigToPb(st *FallbackConfig) (*fallbackConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fallbackConfigPb{}
	enabledPb, err := identity(&st.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledPb != nil {
		pb.Enabled = *enabledPb
	}

	return pb, nil
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

type fallbackConfigPb struct {
	// Whether to enable traffic fallback. When a served entity in the serving
	// endpoint returns specific error codes (e.g. 500), the request will
	// automatically be round-robin attempted with other served entities in the
	// same endpoint, following the order of served entity list, until a
	// successful response is returned. If all attempts fail, return the last
	// response with the error code.
	Enabled bool `json:"enabled"`
}

func fallbackConfigFromPb(pb *fallbackConfigPb) (*FallbackConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FallbackConfig{}
	enabledField, err := identity(&pb.Enabled)
	if err != nil {
		return nil, err
	}
	if enabledField != nil {
		st.Enabled = *enabledField
	}

	return st, nil
}

// All fields are not sensitive as they are hard-coded in the system and made
// available to customers.
type FoundationModel struct {
	Description string

	DisplayName string

	Docs string

	Name string

	ForceSendFields []string
}

func foundationModelToPb(st *FoundationModel) (*foundationModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &foundationModelPb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	docsPb, err := identity(&st.Docs)
	if err != nil {
		return nil, err
	}
	if docsPb != nil {
		pb.Docs = *docsPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	docsField, err := identity(&pb.Docs)
	if err != nil {
		return nil, err
	}
	if docsField != nil {
		st.Docs = *docsField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *foundationModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st foundationModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get the schema for a serving endpoint
type GetOpenApiRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string
}

func getOpenApiRequestToPb(st *GetOpenApiRequest) (*getOpenApiRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getOpenApiRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
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

type getOpenApiRequestPb struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" url:"-"`
}

func getOpenApiRequestFromPb(pb *getOpenApiRequestPb) (*GetOpenApiRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOpenApiRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

type GetOpenApiResponse struct {
	Contents io.ReadCloser
}

func getOpenApiResponseToPb(st *GetOpenApiResponse) (*getOpenApiResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getOpenApiResponsePb{}
	contentsPb, err := identity(&st.Contents)
	if err != nil {
		return nil, err
	}
	if contentsPb != nil {
		pb.Contents = *contentsPb
	}

	return pb, nil
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

type getOpenApiResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

func getOpenApiResponseFromPb(pb *getOpenApiResponsePb) (*GetOpenApiResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOpenApiResponse{}
	contentsField, err := identity(&pb.Contents)
	if err != nil {
		return nil, err
	}
	if contentsField != nil {
		st.Contents = *contentsField
	}

	return st, nil
}

// Get serving endpoint permission levels
type GetServingEndpointPermissionLevelsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string
}

func getServingEndpointPermissionLevelsRequestToPb(st *GetServingEndpointPermissionLevelsRequest) (*getServingEndpointPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServingEndpointPermissionLevelsRequestPb{}
	servingEndpointIdPb, err := identity(&st.ServingEndpointId)
	if err != nil {
		return nil, err
	}
	if servingEndpointIdPb != nil {
		pb.ServingEndpointId = *servingEndpointIdPb
	}

	return pb, nil
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

type getServingEndpointPermissionLevelsRequestPb struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

func getServingEndpointPermissionLevelsRequestFromPb(pb *getServingEndpointPermissionLevelsRequestPb) (*GetServingEndpointPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointPermissionLevelsRequest{}
	servingEndpointIdField, err := identity(&pb.ServingEndpointId)
	if err != nil {
		return nil, err
	}
	if servingEndpointIdField != nil {
		st.ServingEndpointId = *servingEndpointIdField
	}

	return st, nil
}

type GetServingEndpointPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []ServingEndpointPermissionsDescription
}

func getServingEndpointPermissionLevelsResponseToPb(st *GetServingEndpointPermissionLevelsResponse) (*getServingEndpointPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServingEndpointPermissionLevelsResponsePb{}

	var permissionLevelsPb []servingEndpointPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := servingEndpointPermissionsDescriptionToPb(&item)
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

type getServingEndpointPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []servingEndpointPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getServingEndpointPermissionLevelsResponseFromPb(pb *getServingEndpointPermissionLevelsResponsePb) (*GetServingEndpointPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointPermissionLevelsResponse{}

	var permissionLevelsField []ServingEndpointPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := servingEndpointPermissionsDescriptionFromPb(&itemPb)
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

// Get serving endpoint permissions
type GetServingEndpointPermissionsRequest struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string
}

func getServingEndpointPermissionsRequestToPb(st *GetServingEndpointPermissionsRequest) (*getServingEndpointPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServingEndpointPermissionsRequestPb{}
	servingEndpointIdPb, err := identity(&st.ServingEndpointId)
	if err != nil {
		return nil, err
	}
	if servingEndpointIdPb != nil {
		pb.ServingEndpointId = *servingEndpointIdPb
	}

	return pb, nil
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

type getServingEndpointPermissionsRequestPb struct {
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

func getServingEndpointPermissionsRequestFromPb(pb *getServingEndpointPermissionsRequestPb) (*GetServingEndpointPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointPermissionsRequest{}
	servingEndpointIdField, err := identity(&pb.ServingEndpointId)
	if err != nil {
		return nil, err
	}
	if servingEndpointIdField != nil {
		st.ServingEndpointId = *servingEndpointIdField
	}

	return st, nil
}

// Get a single serving endpoint
type GetServingEndpointRequest struct {
	// The name of the serving endpoint. This field is required.
	Name string
}

func getServingEndpointRequestToPb(st *GetServingEndpointRequest) (*getServingEndpointRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getServingEndpointRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
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

type getServingEndpointRequestPb struct {
	// The name of the serving endpoint. This field is required.
	Name string `json:"-" url:"-"`
}

func getServingEndpointRequestFromPb(pb *getServingEndpointRequestPb) (*GetServingEndpointRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetServingEndpointRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

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
	PrivateKey string
	// The private key for the service account which has access to the Google
	// Cloud Vertex AI Service provided as a plaintext secret. See [Best
	// practices for managing service account keys]. If you prefer to reference
	// your key using Databricks Secrets, see `private_key`. You must provide an
	// API key using one of the following fields: `private_key` or
	// `private_key_plaintext`.
	//
	// [Best practices for managing service account keys]:
	// https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys
	PrivateKeyPlaintext string
	// This is the Google Cloud project id that the service account is
	// associated with.
	ProjectId string
	// This is the region for the Google Cloud Vertex AI Service. See [supported
	// regions] for more details. Some models are only available in specific
	// regions.
	//
	// [supported regions]:
	// https://cloud.google.com/vertex-ai/docs/general/locations
	Region string

	ForceSendFields []string
}

func googleCloudVertexAiConfigToPb(st *GoogleCloudVertexAiConfig) (*googleCloudVertexAiConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &googleCloudVertexAiConfigPb{}
	privateKeyPb, err := identity(&st.PrivateKey)
	if err != nil {
		return nil, err
	}
	if privateKeyPb != nil {
		pb.PrivateKey = *privateKeyPb
	}

	privateKeyPlaintextPb, err := identity(&st.PrivateKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if privateKeyPlaintextPb != nil {
		pb.PrivateKeyPlaintext = *privateKeyPlaintextPb
	}

	projectIdPb, err := identity(&st.ProjectId)
	if err != nil {
		return nil, err
	}
	if projectIdPb != nil {
		pb.ProjectId = *projectIdPb
	}

	regionPb, err := identity(&st.Region)
	if err != nil {
		return nil, err
	}
	if regionPb != nil {
		pb.Region = *regionPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type googleCloudVertexAiConfigPb struct {
	// The Databricks secret key reference for a private key for the service
	// account which has access to the Google Cloud Vertex AI Service. See [Best
	// practices for managing service account keys]. If you prefer to paste your
	// API key directly, see `private_key_plaintext`. You must provide an API
	// key using one of the following fields: `private_key` or
	// `private_key_plaintext`
	//
	// [Best practices for managing service account keys]:
	// https://cloud.google.com/iam/docs/best-practices-for-managing-service-account-keys
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
	PrivateKeyPlaintext string `json:"private_key_plaintext,omitempty"`
	// This is the Google Cloud project id that the service account is
	// associated with.
	ProjectId string `json:"project_id"`
	// This is the region for the Google Cloud Vertex AI Service. See [supported
	// regions] for more details. Some models are only available in specific
	// regions.
	//
	// [supported regions]:
	// https://cloud.google.com/vertex-ai/docs/general/locations
	Region string `json:"region"`

	ForceSendFields []string `json:"-" url:"-"`
}

func googleCloudVertexAiConfigFromPb(pb *googleCloudVertexAiConfigPb) (*GoogleCloudVertexAiConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GoogleCloudVertexAiConfig{}
	privateKeyField, err := identity(&pb.PrivateKey)
	if err != nil {
		return nil, err
	}
	if privateKeyField != nil {
		st.PrivateKey = *privateKeyField
	}
	privateKeyPlaintextField, err := identity(&pb.PrivateKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if privateKeyPlaintextField != nil {
		st.PrivateKeyPlaintext = *privateKeyPlaintextField
	}
	projectIdField, err := identity(&pb.ProjectId)
	if err != nil {
		return nil, err
	}
	if projectIdField != nil {
		st.ProjectId = *projectIdField
	}
	regionField, err := identity(&pb.Region)
	if err != nil {
		return nil, err
	}
	if regionField != nil {
		st.Region = *regionField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *googleCloudVertexAiConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st googleCloudVertexAiConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type HttpRequestResponse struct {
	Contents io.ReadCloser
}

func httpRequestResponseToPb(st *HttpRequestResponse) (*httpRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpRequestResponsePb{}
	contentsPb, err := identity(&st.Contents)
	if err != nil {
		return nil, err
	}
	if contentsPb != nil {
		pb.Contents = *contentsPb
	}

	return pb, nil
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

type httpRequestResponsePb struct {
	Contents io.ReadCloser `json:"-"`
}

func httpRequestResponseFromPb(pb *httpRequestResponsePb) (*HttpRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &HttpRequestResponse{}
	contentsField, err := identity(&pb.Contents)
	if err != nil {
		return nil, err
	}
	if contentsField != nil {
		st.Contents = *contentsField
	}

	return st, nil
}

type ListEndpointsResponse struct {
	// The list of endpoints.
	Endpoints []ServingEndpoint
}

func listEndpointsResponseToPb(st *ListEndpointsResponse) (*listEndpointsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listEndpointsResponsePb{}

	var endpointsPb []servingEndpointPb
	for _, item := range st.Endpoints {
		itemPb, err := servingEndpointToPb(&item)
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

type listEndpointsResponsePb struct {
	// The list of endpoints.
	Endpoints []servingEndpointPb `json:"endpoints,omitempty"`
}

func listEndpointsResponseFromPb(pb *listEndpointsResponsePb) (*ListEndpointsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListEndpointsResponse{}

	var endpointsField []ServingEndpoint
	for _, itemPb := range pb.Endpoints {
		item, err := servingEndpointFromPb(&itemPb)
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

// Get the latest logs for a served model
type LogsRequest struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName string
}

func logsRequestToPb(st *LogsRequest) (*logsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logsRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	servedModelNamePb, err := identity(&st.ServedModelName)
	if err != nil {
		return nil, err
	}
	if servedModelNamePb != nil {
		pb.ServedModelName = *servedModelNamePb
	}

	return pb, nil
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

type logsRequestPb struct {
	// The name of the serving endpoint that the served model belongs to. This
	// field is required.
	Name string `json:"-" url:"-"`
	// The name of the served model that logs will be retrieved for. This field
	// is required.
	ServedModelName string `json:"-" url:"-"`
}

func logsRequestFromPb(pb *logsRequestPb) (*LogsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogsRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	servedModelNameField, err := identity(&pb.ServedModelName)
	if err != nil {
		return nil, err
	}
	if servedModelNameField != nil {
		st.ServedModelName = *servedModelNameField
	}

	return st, nil
}

// A representation of all DataPlaneInfo for operations that can be done on a
// model through Data Plane APIs.
type ModelDataPlaneInfo struct {
	// Information required to query DataPlane API 'query' endpoint.
	QueryInfo *DataPlaneInfo
}

func modelDataPlaneInfoToPb(st *ModelDataPlaneInfo) (*modelDataPlaneInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelDataPlaneInfoPb{}
	queryInfoPb, err := dataPlaneInfoToPb(st.QueryInfo)
	if err != nil {
		return nil, err
	}
	if queryInfoPb != nil {
		pb.QueryInfo = queryInfoPb
	}

	return pb, nil
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

type modelDataPlaneInfoPb struct {
	// Information required to query DataPlane API 'query' endpoint.
	QueryInfo *dataPlaneInfoPb `json:"query_info,omitempty"`
}

func modelDataPlaneInfoFromPb(pb *modelDataPlaneInfoPb) (*ModelDataPlaneInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelDataPlaneInfo{}
	queryInfoField, err := dataPlaneInfoFromPb(pb.QueryInfo)
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
	MicrosoftEntraClientId string
	// The Databricks secret key reference for a client secret used for
	// Microsoft Entra ID authentication. If you prefer to paste your client
	// secret directly, see `microsoft_entra_client_secret_plaintext`. You must
	// provide an API key using one of the following fields:
	// `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	MicrosoftEntraClientSecret string
	// The client secret used for Microsoft Entra ID authentication provided as
	// a plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `microsoft_entra_client_secret`. You must provide an API key
	// using one of the following fields: `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	MicrosoftEntraClientSecretPlaintext string
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Tenant ID.
	MicrosoftEntraTenantId string
	// This is a field to provide a customized base URl for the OpenAI API. For
	// Azure OpenAI, this field is required, and is the base URL for the Azure
	// OpenAI API service provided by Azure. For other OpenAI API types, this
	// field is optional, and if left unspecified, the standard OpenAI base URL
	// is used.
	OpenaiApiBase string
	// The Databricks secret key reference for an OpenAI API key using the
	// OpenAI or Azure service. If you prefer to paste your API key directly,
	// see `openai_api_key_plaintext`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	OpenaiApiKey string
	// The OpenAI API key using the OpenAI or Azure service provided as a
	// plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `openai_api_key`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	OpenaiApiKeyPlaintext string
	// This is an optional field to specify the type of OpenAI API to use. For
	// Azure OpenAI, this field is required, and adjust this parameter to
	// represent the preferred security access validation protocol. For access
	// token validation, use azure. For authentication using Azure Active
	// Directory (Azure AD) use, azuread.
	OpenaiApiType string
	// This is an optional field to specify the OpenAI API version. For Azure
	// OpenAI, this field is required, and is the version of the Azure OpenAI
	// service to utilize, specified by a date.
	OpenaiApiVersion string
	// This field is only required for Azure OpenAI and is the name of the
	// deployment resource for the Azure OpenAI service.
	OpenaiDeploymentName string
	// This is an optional field to specify the organization in OpenAI or Azure
	// OpenAI.
	OpenaiOrganization string

	ForceSendFields []string
}

func openAiConfigToPb(st *OpenAiConfig) (*openAiConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &openAiConfigPb{}
	microsoftEntraClientIdPb, err := identity(&st.MicrosoftEntraClientId)
	if err != nil {
		return nil, err
	}
	if microsoftEntraClientIdPb != nil {
		pb.MicrosoftEntraClientId = *microsoftEntraClientIdPb
	}

	microsoftEntraClientSecretPb, err := identity(&st.MicrosoftEntraClientSecret)
	if err != nil {
		return nil, err
	}
	if microsoftEntraClientSecretPb != nil {
		pb.MicrosoftEntraClientSecret = *microsoftEntraClientSecretPb
	}

	microsoftEntraClientSecretPlaintextPb, err := identity(&st.MicrosoftEntraClientSecretPlaintext)
	if err != nil {
		return nil, err
	}
	if microsoftEntraClientSecretPlaintextPb != nil {
		pb.MicrosoftEntraClientSecretPlaintext = *microsoftEntraClientSecretPlaintextPb
	}

	microsoftEntraTenantIdPb, err := identity(&st.MicrosoftEntraTenantId)
	if err != nil {
		return nil, err
	}
	if microsoftEntraTenantIdPb != nil {
		pb.MicrosoftEntraTenantId = *microsoftEntraTenantIdPb
	}

	openaiApiBasePb, err := identity(&st.OpenaiApiBase)
	if err != nil {
		return nil, err
	}
	if openaiApiBasePb != nil {
		pb.OpenaiApiBase = *openaiApiBasePb
	}

	openaiApiKeyPb, err := identity(&st.OpenaiApiKey)
	if err != nil {
		return nil, err
	}
	if openaiApiKeyPb != nil {
		pb.OpenaiApiKey = *openaiApiKeyPb
	}

	openaiApiKeyPlaintextPb, err := identity(&st.OpenaiApiKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if openaiApiKeyPlaintextPb != nil {
		pb.OpenaiApiKeyPlaintext = *openaiApiKeyPlaintextPb
	}

	openaiApiTypePb, err := identity(&st.OpenaiApiType)
	if err != nil {
		return nil, err
	}
	if openaiApiTypePb != nil {
		pb.OpenaiApiType = *openaiApiTypePb
	}

	openaiApiVersionPb, err := identity(&st.OpenaiApiVersion)
	if err != nil {
		return nil, err
	}
	if openaiApiVersionPb != nil {
		pb.OpenaiApiVersion = *openaiApiVersionPb
	}

	openaiDeploymentNamePb, err := identity(&st.OpenaiDeploymentName)
	if err != nil {
		return nil, err
	}
	if openaiDeploymentNamePb != nil {
		pb.OpenaiDeploymentName = *openaiDeploymentNamePb
	}

	openaiOrganizationPb, err := identity(&st.OpenaiOrganization)
	if err != nil {
		return nil, err
	}
	if openaiOrganizationPb != nil {
		pb.OpenaiOrganization = *openaiOrganizationPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type openAiConfigPb struct {
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Client ID.
	MicrosoftEntraClientId string `json:"microsoft_entra_client_id,omitempty"`
	// The Databricks secret key reference for a client secret used for
	// Microsoft Entra ID authentication. If you prefer to paste your client
	// secret directly, see `microsoft_entra_client_secret_plaintext`. You must
	// provide an API key using one of the following fields:
	// `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	MicrosoftEntraClientSecret string `json:"microsoft_entra_client_secret,omitempty"`
	// The client secret used for Microsoft Entra ID authentication provided as
	// a plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `microsoft_entra_client_secret`. You must provide an API key
	// using one of the following fields: `microsoft_entra_client_secret` or
	// `microsoft_entra_client_secret_plaintext`.
	MicrosoftEntraClientSecretPlaintext string `json:"microsoft_entra_client_secret_plaintext,omitempty"`
	// This field is only required for Azure AD OpenAI and is the Microsoft
	// Entra Tenant ID.
	MicrosoftEntraTenantId string `json:"microsoft_entra_tenant_id,omitempty"`
	// This is a field to provide a customized base URl for the OpenAI API. For
	// Azure OpenAI, this field is required, and is the base URL for the Azure
	// OpenAI API service provided by Azure. For other OpenAI API types, this
	// field is optional, and if left unspecified, the standard OpenAI base URL
	// is used.
	OpenaiApiBase string `json:"openai_api_base,omitempty"`
	// The Databricks secret key reference for an OpenAI API key using the
	// OpenAI or Azure service. If you prefer to paste your API key directly,
	// see `openai_api_key_plaintext`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	OpenaiApiKey string `json:"openai_api_key,omitempty"`
	// The OpenAI API key using the OpenAI or Azure service provided as a
	// plaintext string. If you prefer to reference your key using Databricks
	// Secrets, see `openai_api_key`. You must provide an API key using one of
	// the following fields: `openai_api_key` or `openai_api_key_plaintext`.
	OpenaiApiKeyPlaintext string `json:"openai_api_key_plaintext,omitempty"`
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

	ForceSendFields []string `json:"-" url:"-"`
}

func openAiConfigFromPb(pb *openAiConfigPb) (*OpenAiConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OpenAiConfig{}
	microsoftEntraClientIdField, err := identity(&pb.MicrosoftEntraClientId)
	if err != nil {
		return nil, err
	}
	if microsoftEntraClientIdField != nil {
		st.MicrosoftEntraClientId = *microsoftEntraClientIdField
	}
	microsoftEntraClientSecretField, err := identity(&pb.MicrosoftEntraClientSecret)
	if err != nil {
		return nil, err
	}
	if microsoftEntraClientSecretField != nil {
		st.MicrosoftEntraClientSecret = *microsoftEntraClientSecretField
	}
	microsoftEntraClientSecretPlaintextField, err := identity(&pb.MicrosoftEntraClientSecretPlaintext)
	if err != nil {
		return nil, err
	}
	if microsoftEntraClientSecretPlaintextField != nil {
		st.MicrosoftEntraClientSecretPlaintext = *microsoftEntraClientSecretPlaintextField
	}
	microsoftEntraTenantIdField, err := identity(&pb.MicrosoftEntraTenantId)
	if err != nil {
		return nil, err
	}
	if microsoftEntraTenantIdField != nil {
		st.MicrosoftEntraTenantId = *microsoftEntraTenantIdField
	}
	openaiApiBaseField, err := identity(&pb.OpenaiApiBase)
	if err != nil {
		return nil, err
	}
	if openaiApiBaseField != nil {
		st.OpenaiApiBase = *openaiApiBaseField
	}
	openaiApiKeyField, err := identity(&pb.OpenaiApiKey)
	if err != nil {
		return nil, err
	}
	if openaiApiKeyField != nil {
		st.OpenaiApiKey = *openaiApiKeyField
	}
	openaiApiKeyPlaintextField, err := identity(&pb.OpenaiApiKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if openaiApiKeyPlaintextField != nil {
		st.OpenaiApiKeyPlaintext = *openaiApiKeyPlaintextField
	}
	openaiApiTypeField, err := identity(&pb.OpenaiApiType)
	if err != nil {
		return nil, err
	}
	if openaiApiTypeField != nil {
		st.OpenaiApiType = *openaiApiTypeField
	}
	openaiApiVersionField, err := identity(&pb.OpenaiApiVersion)
	if err != nil {
		return nil, err
	}
	if openaiApiVersionField != nil {
		st.OpenaiApiVersion = *openaiApiVersionField
	}
	openaiDeploymentNameField, err := identity(&pb.OpenaiDeploymentName)
	if err != nil {
		return nil, err
	}
	if openaiDeploymentNameField != nil {
		st.OpenaiDeploymentName = *openaiDeploymentNameField
	}
	openaiOrganizationField, err := identity(&pb.OpenaiOrganization)
	if err != nil {
		return nil, err
	}
	if openaiOrganizationField != nil {
		st.OpenaiOrganization = *openaiOrganizationField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *openAiConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st openAiConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PaLmConfig struct {
	// The Databricks secret key reference for a PaLM API key. If you prefer to
	// paste your API key directly, see `palm_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	PalmApiKey string
	// The PaLM API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `palm_api_key`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	PalmApiKeyPlaintext string

	ForceSendFields []string
}

func paLmConfigToPb(st *PaLmConfig) (*paLmConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &paLmConfigPb{}
	palmApiKeyPb, err := identity(&st.PalmApiKey)
	if err != nil {
		return nil, err
	}
	if palmApiKeyPb != nil {
		pb.PalmApiKey = *palmApiKeyPb
	}

	palmApiKeyPlaintextPb, err := identity(&st.PalmApiKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if palmApiKeyPlaintextPb != nil {
		pb.PalmApiKeyPlaintext = *palmApiKeyPlaintextPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type paLmConfigPb struct {
	// The Databricks secret key reference for a PaLM API key. If you prefer to
	// paste your API key directly, see `palm_api_key_plaintext`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	PalmApiKey string `json:"palm_api_key,omitempty"`
	// The PaLM API key provided as a plaintext string. If you prefer to
	// reference your key using Databricks Secrets, see `palm_api_key`. You must
	// provide an API key using one of the following fields: `palm_api_key` or
	// `palm_api_key_plaintext`.
	PalmApiKeyPlaintext string `json:"palm_api_key_plaintext,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func paLmConfigFromPb(pb *paLmConfigPb) (*PaLmConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PaLmConfig{}
	palmApiKeyField, err := identity(&pb.PalmApiKey)
	if err != nil {
		return nil, err
	}
	if palmApiKeyField != nil {
		st.PalmApiKey = *palmApiKeyField
	}
	palmApiKeyPlaintextField, err := identity(&pb.PalmApiKeyPlaintext)
	if err != nil {
		return nil, err
	}
	if palmApiKeyPlaintextField != nil {
		st.PalmApiKeyPlaintext = *palmApiKeyPlaintextField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *paLmConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st paLmConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PatchServingEndpointTags struct {
	// List of endpoint tags to add
	AddTags []EndpointTag
	// List of tag keys to delete
	DeleteTags []string
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name string
}

func patchServingEndpointTagsToPb(st *PatchServingEndpointTags) (*patchServingEndpointTagsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &patchServingEndpointTagsPb{}

	var addTagsPb []endpointTagPb
	for _, item := range st.AddTags {
		itemPb, err := endpointTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			addTagsPb = append(addTagsPb, *itemPb)
		}
	}
	pb.AddTags = addTagsPb

	var deleteTagsPb []string
	for _, item := range st.DeleteTags {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			deleteTagsPb = append(deleteTagsPb, *itemPb)
		}
	}
	pb.DeleteTags = deleteTagsPb

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	return pb, nil
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

type patchServingEndpointTagsPb struct {
	// List of endpoint tags to add
	AddTags []endpointTagPb `json:"add_tags,omitempty"`
	// List of tag keys to delete
	DeleteTags []string `json:"delete_tags,omitempty"`
	// The name of the serving endpoint who's tags to patch. This field is
	// required.
	Name string `json:"-" url:"-"`
}

func patchServingEndpointTagsFromPb(pb *patchServingEndpointTagsPb) (*PatchServingEndpointTags, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PatchServingEndpointTags{}

	var addTagsField []EndpointTag
	for _, itemPb := range pb.AddTags {
		item, err := endpointTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			addTagsField = append(addTagsField, *item)
		}
	}
	st.AddTags = addTagsField

	var deleteTagsField []string
	for _, itemPb := range pb.DeleteTags {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			deleteTagsField = append(deleteTagsField, *item)
		}
	}
	st.DeleteTags = deleteTagsField
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	return st, nil
}

type PayloadTable struct {
	Name string

	Status string

	StatusMessage string

	ForceSendFields []string
}

func payloadTableToPb(st *PayloadTable) (*payloadTablePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &payloadTablePb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	statusPb, err := identity(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	statusMessagePb, err := identity(&st.StatusMessage)
	if err != nil {
		return nil, err
	}
	if statusMessagePb != nil {
		pb.StatusMessage = *statusMessagePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	statusField, err := identity(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	statusMessageField, err := identity(&pb.StatusMessage)
	if err != nil {
		return nil, err
	}
	if statusMessageField != nil {
		st.StatusMessage = *statusMessageField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *payloadTablePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st payloadTablePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PutAiGatewayRequest struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig *FallbackConfig
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails *AiGatewayGuardrails
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig *AiGatewayInferenceTableConfig
	// The name of the serving endpoint whose AI Gateway is being updated. This
	// field is required.
	Name string
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []AiGatewayRateLimit
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig *AiGatewayUsageTrackingConfig
}

func putAiGatewayRequestToPb(st *PutAiGatewayRequest) (*putAiGatewayRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putAiGatewayRequestPb{}
	fallbackConfigPb, err := fallbackConfigToPb(st.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigPb != nil {
		pb.FallbackConfig = fallbackConfigPb
	}

	guardrailsPb, err := aiGatewayGuardrailsToPb(st.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsPb != nil {
		pb.Guardrails = guardrailsPb
	}

	inferenceTableConfigPb, err := aiGatewayInferenceTableConfigToPb(st.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigPb != nil {
		pb.InferenceTableConfig = inferenceTableConfigPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	var rateLimitsPb []aiGatewayRateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := aiGatewayRateLimitToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rateLimitsPb = append(rateLimitsPb, *itemPb)
		}
	}
	pb.RateLimits = rateLimitsPb

	usageTrackingConfigPb, err := aiGatewayUsageTrackingConfigToPb(st.UsageTrackingConfig)
	if err != nil {
		return nil, err
	}
	if usageTrackingConfigPb != nil {
		pb.UsageTrackingConfig = usageTrackingConfigPb
	}

	return pb, nil
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

type putAiGatewayRequestPb struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig *fallbackConfigPb `json:"fallback_config,omitempty"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails *aiGatewayGuardrailsPb `json:"guardrails,omitempty"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig *aiGatewayInferenceTableConfigPb `json:"inference_table_config,omitempty"`
	// The name of the serving endpoint whose AI Gateway is being updated. This
	// field is required.
	Name string `json:"-" url:"-"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []aiGatewayRateLimitPb `json:"rate_limits,omitempty"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig *aiGatewayUsageTrackingConfigPb `json:"usage_tracking_config,omitempty"`
}

func putAiGatewayRequestFromPb(pb *putAiGatewayRequestPb) (*PutAiGatewayRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutAiGatewayRequest{}
	fallbackConfigField, err := fallbackConfigFromPb(pb.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigField != nil {
		st.FallbackConfig = fallbackConfigField
	}
	guardrailsField, err := aiGatewayGuardrailsFromPb(pb.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsField != nil {
		st.Guardrails = guardrailsField
	}
	inferenceTableConfigField, err := aiGatewayInferenceTableConfigFromPb(pb.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigField != nil {
		st.InferenceTableConfig = inferenceTableConfigField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	var rateLimitsField []AiGatewayRateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := aiGatewayRateLimitFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rateLimitsField = append(rateLimitsField, *item)
		}
	}
	st.RateLimits = rateLimitsField
	usageTrackingConfigField, err := aiGatewayUsageTrackingConfigFromPb(pb.UsageTrackingConfig)
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
	FallbackConfig *FallbackConfig
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails *AiGatewayGuardrails
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig *AiGatewayInferenceTableConfig
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []AiGatewayRateLimit
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig *AiGatewayUsageTrackingConfig
}

func putAiGatewayResponseToPb(st *PutAiGatewayResponse) (*putAiGatewayResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putAiGatewayResponsePb{}
	fallbackConfigPb, err := fallbackConfigToPb(st.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigPb != nil {
		pb.FallbackConfig = fallbackConfigPb
	}

	guardrailsPb, err := aiGatewayGuardrailsToPb(st.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsPb != nil {
		pb.Guardrails = guardrailsPb
	}

	inferenceTableConfigPb, err := aiGatewayInferenceTableConfigToPb(st.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigPb != nil {
		pb.InferenceTableConfig = inferenceTableConfigPb
	}

	var rateLimitsPb []aiGatewayRateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := aiGatewayRateLimitToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			rateLimitsPb = append(rateLimitsPb, *itemPb)
		}
	}
	pb.RateLimits = rateLimitsPb

	usageTrackingConfigPb, err := aiGatewayUsageTrackingConfigToPb(st.UsageTrackingConfig)
	if err != nil {
		return nil, err
	}
	if usageTrackingConfigPb != nil {
		pb.UsageTrackingConfig = usageTrackingConfigPb
	}

	return pb, nil
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

type putAiGatewayResponsePb struct {
	// Configuration for traffic fallback which auto fallbacks to other served
	// entities if the request to a served entity fails with certain error
	// codes, to increase availability.
	FallbackConfig *fallbackConfigPb `json:"fallback_config,omitempty"`
	// Configuration for AI Guardrails to prevent unwanted data and unsafe data
	// in requests and responses.
	Guardrails *aiGatewayGuardrailsPb `json:"guardrails,omitempty"`
	// Configuration for payload logging using inference tables. Use these
	// tables to monitor and audit data being sent to and received from model
	// APIs and to improve model quality.
	InferenceTableConfig *aiGatewayInferenceTableConfigPb `json:"inference_table_config,omitempty"`
	// Configuration for rate limits which can be set to limit endpoint traffic.
	RateLimits []aiGatewayRateLimitPb `json:"rate_limits,omitempty"`
	// Configuration to enable usage tracking using system tables. These tables
	// allow you to monitor operational usage on endpoints and their associated
	// costs.
	UsageTrackingConfig *aiGatewayUsageTrackingConfigPb `json:"usage_tracking_config,omitempty"`
}

func putAiGatewayResponseFromPb(pb *putAiGatewayResponsePb) (*PutAiGatewayResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutAiGatewayResponse{}
	fallbackConfigField, err := fallbackConfigFromPb(pb.FallbackConfig)
	if err != nil {
		return nil, err
	}
	if fallbackConfigField != nil {
		st.FallbackConfig = fallbackConfigField
	}
	guardrailsField, err := aiGatewayGuardrailsFromPb(pb.Guardrails)
	if err != nil {
		return nil, err
	}
	if guardrailsField != nil {
		st.Guardrails = guardrailsField
	}
	inferenceTableConfigField, err := aiGatewayInferenceTableConfigFromPb(pb.InferenceTableConfig)
	if err != nil {
		return nil, err
	}
	if inferenceTableConfigField != nil {
		st.InferenceTableConfig = inferenceTableConfigField
	}

	var rateLimitsField []AiGatewayRateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := aiGatewayRateLimitFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			rateLimitsField = append(rateLimitsField, *item)
		}
	}
	st.RateLimits = rateLimitsField
	usageTrackingConfigField, err := aiGatewayUsageTrackingConfigFromPb(pb.UsageTrackingConfig)
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
	Name string
	// The list of endpoint rate limits.
	RateLimits []RateLimit
}

func putRequestToPb(st *PutRequest) (*putRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putRequestPb{}
	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	var rateLimitsPb []rateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := rateLimitToPb(&item)
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

type putRequestPb struct {
	// The name of the serving endpoint whose rate limits are being updated.
	// This field is required.
	Name string `json:"-" url:"-"`
	// The list of endpoint rate limits.
	RateLimits []rateLimitPb `json:"rate_limits,omitempty"`
}

func putRequestFromPb(pb *putRequestPb) (*PutRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutRequest{}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	var rateLimitsField []RateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := rateLimitFromPb(&itemPb)
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
	RateLimits []RateLimit
}

func putResponseToPb(st *PutResponse) (*putResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &putResponsePb{}

	var rateLimitsPb []rateLimitPb
	for _, item := range st.RateLimits {
		itemPb, err := rateLimitToPb(&item)
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

type putResponsePb struct {
	// The list of endpoint rate limits.
	RateLimits []rateLimitPb `json:"rate_limits,omitempty"`
}

func putResponseFromPb(pb *putResponsePb) (*PutResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PutResponse{}

	var rateLimitsField []RateLimit
	for _, itemPb := range pb.RateLimits {
		item, err := rateLimitFromPb(&itemPb)
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
	DataframeRecords []any
	// Pandas Dataframe input in the split orientation.
	DataframeSplit *DataframeSplitInput
	// The extra parameters field used ONLY for __completions, chat,__ and
	// __embeddings external & foundation model__ serving endpoints. This is a
	// map of strings and should only be used with other external/foundation
	// model query fields.
	ExtraParams map[string]string
	// The input string (or array of strings) field used ONLY for __embeddings
	// external & foundation model__ serving endpoints and is the only field
	// (along with extra_params if needed) used by embeddings queries.
	Input any
	// Tensor-based input in columnar format.
	Inputs any
	// Tensor-based input in row format.
	Instances []any
	// The max tokens field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is an integer and should only
	// be used with other chat/completions query fields.
	MaxTokens int
	// The messages field used ONLY for __chat external & foundation model__
	// serving endpoints. This is a map of strings and should only be used with
	// other chat query fields.
	Messages []ChatMessage
	// The n (number of candidates) field used ONLY for __completions__ and
	// __chat external & foundation model__ serving endpoints. This is an
	// integer between 1 and 5 with a default of 1 and should only be used with
	// other chat/completions query fields.
	N int
	// The name of the serving endpoint. This field is required.
	Name string
	// The prompt string (or array of strings) field used ONLY for __completions
	// external & foundation model__ serving endpoints and should only be used
	// with other completions query fields.
	Prompt any
	// The stop sequences field used ONLY for __completions__ and __chat
	// external & foundation model__ serving endpoints. This is a list of
	// strings and should only be used with other chat/completions query fields.
	Stop []string
	// The stream field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a boolean defaulting to
	// false and should only be used with other chat/completions query fields.
	Stream bool
	// The temperature field used ONLY for __completions__ and __chat external &
	// foundation model__ serving endpoints. This is a float between 0.0 and 2.0
	// with a default of 1.0 and should only be used with other chat/completions
	// query fields.
	Temperature float64

	ForceSendFields []string
}

func queryEndpointInputToPb(st *QueryEndpointInput) (*queryEndpointInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryEndpointInputPb{}

	var dataframeRecordsPb []any
	for _, item := range st.DataframeRecords {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataframeRecordsPb = append(dataframeRecordsPb, *itemPb)
		}
	}
	pb.DataframeRecords = dataframeRecordsPb

	dataframeSplitPb, err := dataframeSplitInputToPb(st.DataframeSplit)
	if err != nil {
		return nil, err
	}
	if dataframeSplitPb != nil {
		pb.DataframeSplit = dataframeSplitPb
	}

	extraParamsPb := map[string]string{}
	for k, v := range st.ExtraParams {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			extraParamsPb[k] = *itemPb
		}
	}
	pb.ExtraParams = extraParamsPb

	inputPb, err := identity(&st.Input)
	if err != nil {
		return nil, err
	}
	if inputPb != nil {
		pb.Input = *inputPb
	}

	inputsPb, err := identity(&st.Inputs)
	if err != nil {
		return nil, err
	}
	if inputsPb != nil {
		pb.Inputs = *inputsPb
	}

	var instancesPb []any
	for _, item := range st.Instances {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			instancesPb = append(instancesPb, *itemPb)
		}
	}
	pb.Instances = instancesPb

	maxTokensPb, err := identity(&st.MaxTokens)
	if err != nil {
		return nil, err
	}
	if maxTokensPb != nil {
		pb.MaxTokens = *maxTokensPb
	}

	var messagesPb []chatMessagePb
	for _, item := range st.Messages {
		itemPb, err := chatMessageToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			messagesPb = append(messagesPb, *itemPb)
		}
	}
	pb.Messages = messagesPb

	nPb, err := identity(&st.N)
	if err != nil {
		return nil, err
	}
	if nPb != nil {
		pb.N = *nPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	promptPb, err := identity(&st.Prompt)
	if err != nil {
		return nil, err
	}
	if promptPb != nil {
		pb.Prompt = *promptPb
	}

	var stopPb []string
	for _, item := range st.Stop {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			stopPb = append(stopPb, *itemPb)
		}
	}
	pb.Stop = stopPb

	streamPb, err := identity(&st.Stream)
	if err != nil {
		return nil, err
	}
	if streamPb != nil {
		pb.Stream = *streamPb
	}

	temperaturePb, err := identity(&st.Temperature)
	if err != nil {
		return nil, err
	}
	if temperaturePb != nil {
		pb.Temperature = *temperaturePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryEndpointInputPb struct {
	// Pandas Dataframe input in the records orientation.
	DataframeRecords []any `json:"dataframe_records,omitempty"`
	// Pandas Dataframe input in the split orientation.
	DataframeSplit *dataframeSplitInputPb `json:"dataframe_split,omitempty"`
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
	Messages []chatMessagePb `json:"messages,omitempty"`
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

	ForceSendFields []string `json:"-" url:"-"`
}

func queryEndpointInputFromPb(pb *queryEndpointInputPb) (*QueryEndpointInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryEndpointInput{}

	var dataframeRecordsField []any
	for _, itemPb := range pb.DataframeRecords {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dataframeRecordsField = append(dataframeRecordsField, *item)
		}
	}
	st.DataframeRecords = dataframeRecordsField
	dataframeSplitField, err := dataframeSplitInputFromPb(pb.DataframeSplit)
	if err != nil {
		return nil, err
	}
	if dataframeSplitField != nil {
		st.DataframeSplit = dataframeSplitField
	}

	extraParamsField := map[string]string{}
	for k, v := range pb.ExtraParams {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			extraParamsField[k] = *item
		}
	}
	st.ExtraParams = extraParamsField
	inputField, err := identity(&pb.Input)
	if err != nil {
		return nil, err
	}
	if inputField != nil {
		st.Input = *inputField
	}
	inputsField, err := identity(&pb.Inputs)
	if err != nil {
		return nil, err
	}
	if inputsField != nil {
		st.Inputs = *inputsField
	}

	var instancesField []any
	for _, itemPb := range pb.Instances {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			instancesField = append(instancesField, *item)
		}
	}
	st.Instances = instancesField
	maxTokensField, err := identity(&pb.MaxTokens)
	if err != nil {
		return nil, err
	}
	if maxTokensField != nil {
		st.MaxTokens = *maxTokensField
	}

	var messagesField []ChatMessage
	for _, itemPb := range pb.Messages {
		item, err := chatMessageFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			messagesField = append(messagesField, *item)
		}
	}
	st.Messages = messagesField
	nField, err := identity(&pb.N)
	if err != nil {
		return nil, err
	}
	if nField != nil {
		st.N = *nField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	promptField, err := identity(&pb.Prompt)
	if err != nil {
		return nil, err
	}
	if promptField != nil {
		st.Prompt = *promptField
	}

	var stopField []string
	for _, itemPb := range pb.Stop {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			stopField = append(stopField, *item)
		}
	}
	st.Stop = stopField
	streamField, err := identity(&pb.Stream)
	if err != nil {
		return nil, err
	}
	if streamField != nil {
		st.Stream = *streamField
	}
	temperatureField, err := identity(&pb.Temperature)
	if err != nil {
		return nil, err
	}
	if temperatureField != nil {
		st.Temperature = *temperatureField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryEndpointInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryEndpointInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QueryEndpointResponse struct {
	// The list of choices returned by the __chat or completions
	// external/foundation model__ serving endpoint.
	Choices []V1ResponseChoiceElement
	// The timestamp in seconds when the query was created in Unix time returned
	// by a __completions or chat external/foundation model__ serving endpoint.
	Created int64
	// The list of the embeddings returned by the __embeddings
	// external/foundation model__ serving endpoint.
	Data []EmbeddingsV1ResponseEmbeddingElement
	// The ID of the query that may be returned by a __completions or chat
	// external/foundation model__ serving endpoint.
	Id string
	// The name of the __external/foundation model__ used for querying. This is
	// the name of the model that was specified in the endpoint config.
	Model string
	// The type of object returned by the __external/foundation model__ serving
	// endpoint, one of [text_completion, chat.completion, list (of
	// embeddings)].
	Object QueryEndpointResponseObject
	// The predictions returned by the serving endpoint.
	Predictions []any
	// The name of the served model that served the request. This is useful when
	// there are multiple models behind the same endpoint with traffic split.
	ServedModelName string
	// The usage object that may be returned by the __external/foundation
	// model__ serving endpoint. This contains information about the number of
	// tokens used in the prompt and response.
	Usage *ExternalModelUsageElement

	ForceSendFields []string
}

func queryEndpointResponseToPb(st *QueryEndpointResponse) (*queryEndpointResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &queryEndpointResponsePb{}

	var choicesPb []v1ResponseChoiceElementPb
	for _, item := range st.Choices {
		itemPb, err := v1ResponseChoiceElementToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			choicesPb = append(choicesPb, *itemPb)
		}
	}
	pb.Choices = choicesPb

	createdPb, err := identity(&st.Created)
	if err != nil {
		return nil, err
	}
	if createdPb != nil {
		pb.Created = *createdPb
	}

	var dataPb []embeddingsV1ResponseEmbeddingElementPb
	for _, item := range st.Data {
		itemPb, err := embeddingsV1ResponseEmbeddingElementToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			dataPb = append(dataPb, *itemPb)
		}
	}
	pb.Data = dataPb

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	modelPb, err := identity(&st.Model)
	if err != nil {
		return nil, err
	}
	if modelPb != nil {
		pb.Model = *modelPb
	}

	objectPb, err := identity(&st.Object)
	if err != nil {
		return nil, err
	}
	if objectPb != nil {
		pb.Object = *objectPb
	}

	var predictionsPb []any
	for _, item := range st.Predictions {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			predictionsPb = append(predictionsPb, *itemPb)
		}
	}
	pb.Predictions = predictionsPb

	servedModelNamePb, err := identity(&st.ServedModelName)
	if err != nil {
		return nil, err
	}
	if servedModelNamePb != nil {
		pb.ServedModelName = *servedModelNamePb
	}

	usagePb, err := externalModelUsageElementToPb(st.Usage)
	if err != nil {
		return nil, err
	}
	if usagePb != nil {
		pb.Usage = usagePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type queryEndpointResponsePb struct {
	// The list of choices returned by the __chat or completions
	// external/foundation model__ serving endpoint.
	Choices []v1ResponseChoiceElementPb `json:"choices,omitempty"`
	// The timestamp in seconds when the query was created in Unix time returned
	// by a __completions or chat external/foundation model__ serving endpoint.
	Created int64 `json:"created,omitempty"`
	// The list of the embeddings returned by the __embeddings
	// external/foundation model__ serving endpoint.
	Data []embeddingsV1ResponseEmbeddingElementPb `json:"data,omitempty"`
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
	Usage *externalModelUsageElementPb `json:"usage,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func queryEndpointResponseFromPb(pb *queryEndpointResponsePb) (*QueryEndpointResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QueryEndpointResponse{}

	var choicesField []V1ResponseChoiceElement
	for _, itemPb := range pb.Choices {
		item, err := v1ResponseChoiceElementFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			choicesField = append(choicesField, *item)
		}
	}
	st.Choices = choicesField
	createdField, err := identity(&pb.Created)
	if err != nil {
		return nil, err
	}
	if createdField != nil {
		st.Created = *createdField
	}

	var dataField []EmbeddingsV1ResponseEmbeddingElement
	for _, itemPb := range pb.Data {
		item, err := embeddingsV1ResponseEmbeddingElementFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			dataField = append(dataField, *item)
		}
	}
	st.Data = dataField
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	modelField, err := identity(&pb.Model)
	if err != nil {
		return nil, err
	}
	if modelField != nil {
		st.Model = *modelField
	}
	objectField, err := identity(&pb.Object)
	if err != nil {
		return nil, err
	}
	if objectField != nil {
		st.Object = *objectField
	}

	var predictionsField []any
	for _, itemPb := range pb.Predictions {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			predictionsField = append(predictionsField, *item)
		}
	}
	st.Predictions = predictionsField
	servedModelNameField, err := identity(&pb.ServedModelName)
	if err != nil {
		return nil, err
	}
	if servedModelNameField != nil {
		st.ServedModelName = *servedModelNameField
	}
	usageField, err := externalModelUsageElementFromPb(pb.Usage)
	if err != nil {
		return nil, err
	}
	if usageField != nil {
		st.Usage = usageField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *queryEndpointResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st queryEndpointResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The type of object returned by the __external/foundation model__ serving
// endpoint, one of [text_completion, chat.completion, list (of embeddings)].
type QueryEndpointResponseObject string
type queryEndpointResponseObjectPb string

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

func queryEndpointResponseObjectToPb(st *QueryEndpointResponseObject) (*queryEndpointResponseObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := queryEndpointResponseObjectPb(*st)
	return &pb, nil
}

func queryEndpointResponseObjectFromPb(pb *queryEndpointResponseObjectPb) (*QueryEndpointResponseObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := QueryEndpointResponseObject(*pb)
	return &st, nil
}

type RateLimit struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls int64
	// Key field for a serving endpoint rate limit. Currently, only 'user' and
	// 'endpoint' are supported, with 'endpoint' being the default if not
	// specified.
	Key RateLimitKey
	// Renewal period field for a serving endpoint rate limit. Currently, only
	// 'minute' is supported.
	RenewalPeriod RateLimitRenewalPeriod
}

func rateLimitToPb(st *RateLimit) (*rateLimitPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rateLimitPb{}
	callsPb, err := identity(&st.Calls)
	if err != nil {
		return nil, err
	}
	if callsPb != nil {
		pb.Calls = *callsPb
	}

	keyPb, err := identity(&st.Key)
	if err != nil {
		return nil, err
	}
	if keyPb != nil {
		pb.Key = *keyPb
	}

	renewalPeriodPb, err := identity(&st.RenewalPeriod)
	if err != nil {
		return nil, err
	}
	if renewalPeriodPb != nil {
		pb.RenewalPeriod = *renewalPeriodPb
	}

	return pb, nil
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

type rateLimitPb struct {
	// Used to specify how many calls are allowed for a key within the
	// renewal_period.
	Calls int64 `json:"calls"`
	// Key field for a serving endpoint rate limit. Currently, only 'user' and
	// 'endpoint' are supported, with 'endpoint' being the default if not
	// specified.
	Key RateLimitKey `json:"key,omitempty"`
	// Renewal period field for a serving endpoint rate limit. Currently, only
	// 'minute' is supported.
	RenewalPeriod RateLimitRenewalPeriod `json:"renewal_period"`
}

func rateLimitFromPb(pb *rateLimitPb) (*RateLimit, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RateLimit{}
	callsField, err := identity(&pb.Calls)
	if err != nil {
		return nil, err
	}
	if callsField != nil {
		st.Calls = *callsField
	}
	keyField, err := identity(&pb.Key)
	if err != nil {
		return nil, err
	}
	if keyField != nil {
		st.Key = *keyField
	}
	renewalPeriodField, err := identity(&pb.RenewalPeriod)
	if err != nil {
		return nil, err
	}
	if renewalPeriodField != nil {
		st.RenewalPeriod = *renewalPeriodField
	}

	return st, nil
}

type RateLimitKey string
type rateLimitKeyPb string

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

func rateLimitKeyToPb(st *RateLimitKey) (*rateLimitKeyPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := rateLimitKeyPb(*st)
	return &pb, nil
}

func rateLimitKeyFromPb(pb *rateLimitKeyPb) (*RateLimitKey, error) {
	if pb == nil {
		return nil, nil
	}
	st := RateLimitKey(*pb)
	return &st, nil
}

type RateLimitRenewalPeriod string
type rateLimitRenewalPeriodPb string

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

func rateLimitRenewalPeriodToPb(st *RateLimitRenewalPeriod) (*rateLimitRenewalPeriodPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := rateLimitRenewalPeriodPb(*st)
	return &pb, nil
}

func rateLimitRenewalPeriodFromPb(pb *rateLimitRenewalPeriodPb) (*RateLimitRenewalPeriod, error) {
	if pb == nil {
		return nil, nil
	}
	st := RateLimitRenewalPeriod(*pb)
	return &st, nil
}

type Route struct {
	// The name of the served model this route configures traffic for.
	ServedModelName string
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage int
}

func routeToPb(st *Route) (*routePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &routePb{}
	servedModelNamePb, err := identity(&st.ServedModelName)
	if err != nil {
		return nil, err
	}
	if servedModelNamePb != nil {
		pb.ServedModelName = *servedModelNamePb
	}

	trafficPercentagePb, err := identity(&st.TrafficPercentage)
	if err != nil {
		return nil, err
	}
	if trafficPercentagePb != nil {
		pb.TrafficPercentage = *trafficPercentagePb
	}

	return pb, nil
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

type routePb struct {
	// The name of the served model this route configures traffic for.
	ServedModelName string `json:"served_model_name"`
	// The percentage of endpoint traffic to send to this route. It must be an
	// integer between 0 and 100 inclusive.
	TrafficPercentage int `json:"traffic_percentage"`
}

func routeFromPb(pb *routePb) (*Route, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Route{}
	servedModelNameField, err := identity(&pb.ServedModelName)
	if err != nil {
		return nil, err
	}
	if servedModelNameField != nil {
		st.ServedModelName = *servedModelNameField
	}
	trafficPercentageField, err := identity(&pb.TrafficPercentage)
	if err != nil {
		return nil, err
	}
	if trafficPercentageField != nil {
		st.TrafficPercentage = *trafficPercentageField
	}

	return st, nil
}

type ServedEntityInput struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	EntityName string

	EntityVersion string
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	ExternalModel *ExternalModel
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	WorkloadSize string
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServingModelWorkloadType

	ForceSendFields []string
}

func servedEntityInputToPb(st *ServedEntityInput) (*servedEntityInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedEntityInputPb{}
	entityNamePb, err := identity(&st.EntityName)
	if err != nil {
		return nil, err
	}
	if entityNamePb != nil {
		pb.EntityName = *entityNamePb
	}

	entityVersionPb, err := identity(&st.EntityVersion)
	if err != nil {
		return nil, err
	}
	if entityVersionPb != nil {
		pb.EntityVersion = *entityVersionPb
	}

	environmentVarsPb := map[string]string{}
	for k, v := range st.EnvironmentVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			environmentVarsPb[k] = *itemPb
		}
	}
	pb.EnvironmentVars = environmentVarsPb

	externalModelPb, err := externalModelToPb(st.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelPb != nil {
		pb.ExternalModel = externalModelPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	maxProvisionedThroughputPb, err := identity(&st.MaxProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if maxProvisionedThroughputPb != nil {
		pb.MaxProvisionedThroughput = *maxProvisionedThroughputPb
	}

	minProvisionedThroughputPb, err := identity(&st.MinProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if minProvisionedThroughputPb != nil {
		pb.MinProvisionedThroughput = *minProvisionedThroughputPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	scaleToZeroEnabledPb, err := identity(&st.ScaleToZeroEnabled)
	if err != nil {
		return nil, err
	}
	if scaleToZeroEnabledPb != nil {
		pb.ScaleToZeroEnabled = *scaleToZeroEnabledPb
	}

	workloadSizePb, err := identity(&st.WorkloadSize)
	if err != nil {
		return nil, err
	}
	if workloadSizePb != nil {
		pb.WorkloadSize = *workloadSizePb
	}

	workloadTypePb, err := identity(&st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = *workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servedEntityInputPb struct {
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	EntityName string `json:"entity_name,omitempty"`

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
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	ExternalModel *externalModelPb `json:"external_model,omitempty"`
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
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string `json:"name,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedEntityInputFromPb(pb *servedEntityInputPb) (*ServedEntityInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedEntityInput{}
	entityNameField, err := identity(&pb.EntityName)
	if err != nil {
		return nil, err
	}
	if entityNameField != nil {
		st.EntityName = *entityNameField
	}
	entityVersionField, err := identity(&pb.EntityVersion)
	if err != nil {
		return nil, err
	}
	if entityVersionField != nil {
		st.EntityVersion = *entityVersionField
	}

	environmentVarsField := map[string]string{}
	for k, v := range pb.EnvironmentVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			environmentVarsField[k] = *item
		}
	}
	st.EnvironmentVars = environmentVarsField
	externalModelField, err := externalModelFromPb(pb.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelField != nil {
		st.ExternalModel = externalModelField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	maxProvisionedThroughputField, err := identity(&pb.MaxProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if maxProvisionedThroughputField != nil {
		st.MaxProvisionedThroughput = *maxProvisionedThroughputField
	}
	minProvisionedThroughputField, err := identity(&pb.MinProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if minProvisionedThroughputField != nil {
		st.MinProvisionedThroughput = *minProvisionedThroughputField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	scaleToZeroEnabledField, err := identity(&pb.ScaleToZeroEnabled)
	if err != nil {
		return nil, err
	}
	if scaleToZeroEnabledField != nil {
		st.ScaleToZeroEnabled = *scaleToZeroEnabledField
	}
	workloadSizeField, err := identity(&pb.WorkloadSize)
	if err != nil {
		return nil, err
	}
	if workloadSizeField != nil {
		st.WorkloadSize = *workloadSizeField
	}
	workloadTypeField, err := identity(&pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = *workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedEntityInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedEntityInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedEntityOutput struct {
	CreationTimestamp int64

	Creator string
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	EntityName string

	EntityVersion string
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string
	// The external model to be served. NOTE: Only one of external_model and
	// (entity_name, entity_version, workload_size, workload_type, and
	// scale_to_zero_enabled) can be specified with the latter set being used
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	ExternalModel *ExternalModel
	// All fields are not sensitive as they are hard-coded in the system and
	// made available to customers.
	FoundationModel *FoundationModel
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool

	State *ServedModelState
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	WorkloadSize string
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServingModelWorkloadType

	ForceSendFields []string
}

func servedEntityOutputToPb(st *ServedEntityOutput) (*servedEntityOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedEntityOutputPb{}
	creationTimestampPb, err := identity(&st.CreationTimestamp)
	if err != nil {
		return nil, err
	}
	if creationTimestampPb != nil {
		pb.CreationTimestamp = *creationTimestampPb
	}

	creatorPb, err := identity(&st.Creator)
	if err != nil {
		return nil, err
	}
	if creatorPb != nil {
		pb.Creator = *creatorPb
	}

	entityNamePb, err := identity(&st.EntityName)
	if err != nil {
		return nil, err
	}
	if entityNamePb != nil {
		pb.EntityName = *entityNamePb
	}

	entityVersionPb, err := identity(&st.EntityVersion)
	if err != nil {
		return nil, err
	}
	if entityVersionPb != nil {
		pb.EntityVersion = *entityVersionPb
	}

	environmentVarsPb := map[string]string{}
	for k, v := range st.EnvironmentVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			environmentVarsPb[k] = *itemPb
		}
	}
	pb.EnvironmentVars = environmentVarsPb

	externalModelPb, err := externalModelToPb(st.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelPb != nil {
		pb.ExternalModel = externalModelPb
	}

	foundationModelPb, err := foundationModelToPb(st.FoundationModel)
	if err != nil {
		return nil, err
	}
	if foundationModelPb != nil {
		pb.FoundationModel = foundationModelPb
	}

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	maxProvisionedThroughputPb, err := identity(&st.MaxProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if maxProvisionedThroughputPb != nil {
		pb.MaxProvisionedThroughput = *maxProvisionedThroughputPb
	}

	minProvisionedThroughputPb, err := identity(&st.MinProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if minProvisionedThroughputPb != nil {
		pb.MinProvisionedThroughput = *minProvisionedThroughputPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	scaleToZeroEnabledPb, err := identity(&st.ScaleToZeroEnabled)
	if err != nil {
		return nil, err
	}
	if scaleToZeroEnabledPb != nil {
		pb.ScaleToZeroEnabled = *scaleToZeroEnabledPb
	}

	statePb, err := servedModelStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	workloadSizePb, err := identity(&st.WorkloadSize)
	if err != nil {
		return nil, err
	}
	if workloadSizePb != nil {
		pb.WorkloadSize = *workloadSizePb
	}

	workloadTypePb, err := identity(&st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = *workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servedEntityOutputPb struct {
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	Creator string `json:"creator,omitempty"`
	// The name of the entity to be served. The entity may be a model in the
	// Databricks Model Registry, a model in the Unity Catalog (UC), or a
	// function of type FEATURE_SPEC in the UC. If it is a UC object, the full
	// name of the object should be given in the form of
	// **catalog_name.schema_name.model_name**.
	EntityName string `json:"entity_name,omitempty"`

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
	// for custom model serving for a Databricks registered model. For an
	// existing endpoint with external_model, it cannot be updated to an
	// endpoint without external_model. If the endpoint is created without
	// external_model, users cannot update it to add external_model later. The
	// task type of all external models within an endpoint must be the same.
	ExternalModel *externalModelPb `json:"external_model,omitempty"`
	// All fields are not sensitive as they are hard-coded in the system and
	// made available to customers.
	FoundationModel *foundationModelPb `json:"foundation_model,omitempty"`
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
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string `json:"name,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`

	State *servedModelStatePb `json:"state,omitempty"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedEntityOutputFromPb(pb *servedEntityOutputPb) (*ServedEntityOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedEntityOutput{}
	creationTimestampField, err := identity(&pb.CreationTimestamp)
	if err != nil {
		return nil, err
	}
	if creationTimestampField != nil {
		st.CreationTimestamp = *creationTimestampField
	}
	creatorField, err := identity(&pb.Creator)
	if err != nil {
		return nil, err
	}
	if creatorField != nil {
		st.Creator = *creatorField
	}
	entityNameField, err := identity(&pb.EntityName)
	if err != nil {
		return nil, err
	}
	if entityNameField != nil {
		st.EntityName = *entityNameField
	}
	entityVersionField, err := identity(&pb.EntityVersion)
	if err != nil {
		return nil, err
	}
	if entityVersionField != nil {
		st.EntityVersion = *entityVersionField
	}

	environmentVarsField := map[string]string{}
	for k, v := range pb.EnvironmentVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			environmentVarsField[k] = *item
		}
	}
	st.EnvironmentVars = environmentVarsField
	externalModelField, err := externalModelFromPb(pb.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelField != nil {
		st.ExternalModel = externalModelField
	}
	foundationModelField, err := foundationModelFromPb(pb.FoundationModel)
	if err != nil {
		return nil, err
	}
	if foundationModelField != nil {
		st.FoundationModel = foundationModelField
	}
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	maxProvisionedThroughputField, err := identity(&pb.MaxProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if maxProvisionedThroughputField != nil {
		st.MaxProvisionedThroughput = *maxProvisionedThroughputField
	}
	minProvisionedThroughputField, err := identity(&pb.MinProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if minProvisionedThroughputField != nil {
		st.MinProvisionedThroughput = *minProvisionedThroughputField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	scaleToZeroEnabledField, err := identity(&pb.ScaleToZeroEnabled)
	if err != nil {
		return nil, err
	}
	if scaleToZeroEnabledField != nil {
		st.ScaleToZeroEnabled = *scaleToZeroEnabledField
	}
	stateField, err := servedModelStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	workloadSizeField, err := identity(&pb.WorkloadSize)
	if err != nil {
		return nil, err
	}
	if workloadSizeField != nil {
		st.WorkloadSize = *workloadSizeField
	}
	workloadTypeField, err := identity(&pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = *workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedEntityOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedEntityOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedEntitySpec struct {
	EntityName string

	EntityVersion string

	ExternalModel *ExternalModel
	// All fields are not sensitive as they are hard-coded in the system and
	// made available to customers.
	FoundationModel *FoundationModel

	Name string

	ForceSendFields []string
}

func servedEntitySpecToPb(st *ServedEntitySpec) (*servedEntitySpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedEntitySpecPb{}
	entityNamePb, err := identity(&st.EntityName)
	if err != nil {
		return nil, err
	}
	if entityNamePb != nil {
		pb.EntityName = *entityNamePb
	}

	entityVersionPb, err := identity(&st.EntityVersion)
	if err != nil {
		return nil, err
	}
	if entityVersionPb != nil {
		pb.EntityVersion = *entityVersionPb
	}

	externalModelPb, err := externalModelToPb(st.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelPb != nil {
		pb.ExternalModel = externalModelPb
	}

	foundationModelPb, err := foundationModelToPb(st.FoundationModel)
	if err != nil {
		return nil, err
	}
	if foundationModelPb != nil {
		pb.FoundationModel = foundationModelPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servedEntitySpecPb struct {
	EntityName string `json:"entity_name,omitempty"`

	EntityVersion string `json:"entity_version,omitempty"`

	ExternalModel *externalModelPb `json:"external_model,omitempty"`
	// All fields are not sensitive as they are hard-coded in the system and
	// made available to customers.
	FoundationModel *foundationModelPb `json:"foundation_model,omitempty"`

	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedEntitySpecFromPb(pb *servedEntitySpecPb) (*ServedEntitySpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedEntitySpec{}
	entityNameField, err := identity(&pb.EntityName)
	if err != nil {
		return nil, err
	}
	if entityNameField != nil {
		st.EntityName = *entityNameField
	}
	entityVersionField, err := identity(&pb.EntityVersion)
	if err != nil {
		return nil, err
	}
	if entityVersionField != nil {
		st.EntityVersion = *entityVersionField
	}
	externalModelField, err := externalModelFromPb(pb.ExternalModel)
	if err != nil {
		return nil, err
	}
	if externalModelField != nil {
		st.ExternalModel = externalModelField
	}
	foundationModelField, err := foundationModelFromPb(pb.FoundationModel)
	if err != nil {
		return nil, err
	}
	if foundationModelField != nil {
		st.FoundationModel = foundationModelField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedEntitySpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedEntitySpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedModelInput struct {
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int

	ModelName string

	ModelVersion string
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	WorkloadSize string
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServedModelInputWorkloadType

	ForceSendFields []string
}

func servedModelInputToPb(st *ServedModelInput) (*servedModelInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedModelInputPb{}

	environmentVarsPb := map[string]string{}
	for k, v := range st.EnvironmentVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			environmentVarsPb[k] = *itemPb
		}
	}
	pb.EnvironmentVars = environmentVarsPb

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	maxProvisionedThroughputPb, err := identity(&st.MaxProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if maxProvisionedThroughputPb != nil {
		pb.MaxProvisionedThroughput = *maxProvisionedThroughputPb
	}

	minProvisionedThroughputPb, err := identity(&st.MinProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if minProvisionedThroughputPb != nil {
		pb.MinProvisionedThroughput = *minProvisionedThroughputPb
	}

	modelNamePb, err := identity(&st.ModelName)
	if err != nil {
		return nil, err
	}
	if modelNamePb != nil {
		pb.ModelName = *modelNamePb
	}

	modelVersionPb, err := identity(&st.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionPb != nil {
		pb.ModelVersion = *modelVersionPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	scaleToZeroEnabledPb, err := identity(&st.ScaleToZeroEnabled)
	if err != nil {
		return nil, err
	}
	if scaleToZeroEnabledPb != nil {
		pb.ScaleToZeroEnabled = *scaleToZeroEnabledPb
	}

	workloadSizePb, err := identity(&st.WorkloadSize)
	if err != nil {
		return nil, err
	}
	if workloadSizePb != nil {
		pb.WorkloadSize = *workloadSizePb
	}

	workloadTypePb, err := identity(&st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = *workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servedModelInputPb struct {
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`
	// The maximum tokens per second that the endpoint can scale up to.
	MaxProvisionedThroughput int `json:"max_provisioned_throughput,omitempty"`
	// The minimum tokens per second that the endpoint can scale down to.
	MinProvisionedThroughput int `json:"min_provisioned_throughput,omitempty"`

	ModelName string `json:"model_name"`

	ModelVersion string `json:"model_version"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string `json:"name,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServedModelInputWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedModelInputFromPb(pb *servedModelInputPb) (*ServedModelInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedModelInput{}

	environmentVarsField := map[string]string{}
	for k, v := range pb.EnvironmentVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			environmentVarsField[k] = *item
		}
	}
	st.EnvironmentVars = environmentVarsField
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	maxProvisionedThroughputField, err := identity(&pb.MaxProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if maxProvisionedThroughputField != nil {
		st.MaxProvisionedThroughput = *maxProvisionedThroughputField
	}
	minProvisionedThroughputField, err := identity(&pb.MinProvisionedThroughput)
	if err != nil {
		return nil, err
	}
	if minProvisionedThroughputField != nil {
		st.MinProvisionedThroughput = *minProvisionedThroughputField
	}
	modelNameField, err := identity(&pb.ModelName)
	if err != nil {
		return nil, err
	}
	if modelNameField != nil {
		st.ModelName = *modelNameField
	}
	modelVersionField, err := identity(&pb.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionField != nil {
		st.ModelVersion = *modelVersionField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	scaleToZeroEnabledField, err := identity(&pb.ScaleToZeroEnabled)
	if err != nil {
		return nil, err
	}
	if scaleToZeroEnabledField != nil {
		st.ScaleToZeroEnabled = *scaleToZeroEnabledField
	}
	workloadSizeField, err := identity(&pb.WorkloadSize)
	if err != nil {
		return nil, err
	}
	if workloadSizeField != nil {
		st.WorkloadSize = *workloadSizeField
	}
	workloadTypeField, err := identity(&pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = *workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedModelInputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedModelInputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Please keep this in sync with with workload types in
// InferenceEndpointEntities.scala
type ServedModelInputWorkloadType string
type servedModelInputWorkloadTypePb string

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

func servedModelInputWorkloadTypeToPb(st *ServedModelInputWorkloadType) (*servedModelInputWorkloadTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servedModelInputWorkloadTypePb(*st)
	return &pb, nil
}

func servedModelInputWorkloadTypeFromPb(pb *servedModelInputWorkloadTypePb) (*ServedModelInputWorkloadType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServedModelInputWorkloadType(*pb)
	return &st, nil
}

type ServedModelOutput struct {
	CreationTimestamp int64

	Creator string
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string

	ModelName string

	ModelVersion string
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool

	State *ServedModelState
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	WorkloadSize string
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServingModelWorkloadType

	ForceSendFields []string
}

func servedModelOutputToPb(st *ServedModelOutput) (*servedModelOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedModelOutputPb{}
	creationTimestampPb, err := identity(&st.CreationTimestamp)
	if err != nil {
		return nil, err
	}
	if creationTimestampPb != nil {
		pb.CreationTimestamp = *creationTimestampPb
	}

	creatorPb, err := identity(&st.Creator)
	if err != nil {
		return nil, err
	}
	if creatorPb != nil {
		pb.Creator = *creatorPb
	}

	environmentVarsPb := map[string]string{}
	for k, v := range st.EnvironmentVars {
		itemPb, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			environmentVarsPb[k] = *itemPb
		}
	}
	pb.EnvironmentVars = environmentVarsPb

	instanceProfileArnPb, err := identity(&st.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnPb != nil {
		pb.InstanceProfileArn = *instanceProfileArnPb
	}

	modelNamePb, err := identity(&st.ModelName)
	if err != nil {
		return nil, err
	}
	if modelNamePb != nil {
		pb.ModelName = *modelNamePb
	}

	modelVersionPb, err := identity(&st.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionPb != nil {
		pb.ModelVersion = *modelVersionPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	scaleToZeroEnabledPb, err := identity(&st.ScaleToZeroEnabled)
	if err != nil {
		return nil, err
	}
	if scaleToZeroEnabledPb != nil {
		pb.ScaleToZeroEnabled = *scaleToZeroEnabledPb
	}

	statePb, err := servedModelStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	workloadSizePb, err := identity(&st.WorkloadSize)
	if err != nil {
		return nil, err
	}
	if workloadSizePb != nil {
		pb.WorkloadSize = *workloadSizePb
	}

	workloadTypePb, err := identity(&st.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypePb != nil {
		pb.WorkloadType = *workloadTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servedModelOutputPb struct {
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`

	Creator string `json:"creator,omitempty"`
	// An object containing a set of optional, user-specified environment
	// variable key-value pairs used for serving this entity. Note: this is an
	// experimental feature and subject to change. Example entity environment
	// variables that refer to Databricks secrets: `{"OPENAI_API_KEY":
	// "{{secrets/my_scope/my_key}}", "DATABRICKS_TOKEN":
	// "{{secrets/my_scope2/my_key2}}"}`
	EnvironmentVars map[string]string `json:"environment_vars,omitempty"`
	// ARN of the instance profile that the served entity uses to access AWS
	// resources.
	InstanceProfileArn string `json:"instance_profile_arn,omitempty"`

	ModelName string `json:"model_name,omitempty"`

	ModelVersion string `json:"model_version,omitempty"`
	// The name of a served entity. It must be unique across an endpoint. A
	// served entity name can consist of alphanumeric characters, dashes, and
	// underscores. If not specified for an external model, this field defaults
	// to external_model.name, with '.' and ':' replaced with '-', and if not
	// specified for other entities, it defaults to entity_name-entity_version.
	Name string `json:"name,omitempty"`
	// Whether the compute resources for the served entity should scale down to
	// zero.
	ScaleToZeroEnabled bool `json:"scale_to_zero_enabled,omitempty"`

	State *servedModelStatePb `json:"state,omitempty"`
	// The workload size of the served entity. The workload size corresponds to
	// a range of provisioned concurrency that the compute autoscales between. A
	// single unit of provisioned concurrency can process one request at a time.
	// Valid workload sizes are "Small" (4 - 4 provisioned concurrency),
	// "Medium" (8 - 16 provisioned concurrency), and "Large" (16 - 64
	// provisioned concurrency). Additional custom workload sizes can also be
	// used when available in the workspace. If scale-to-zero is enabled, the
	// lower bound of the provisioned concurrency for each workload size is 0.
	WorkloadSize string `json:"workload_size,omitempty"`
	// The workload type of the served entity. The workload type selects which
	// type of compute to use in the endpoint. The default value for this
	// parameter is "CPU". For deep learning workloads, GPU acceleration is
	// available by selecting workload types like GPU_SMALL and others. See the
	// available [GPU types].
	//
	// [GPU types]: https://docs.databricks.com/en/machine-learning/model-serving/create-manage-serving-endpoints.html#gpu-workload-types
	WorkloadType ServingModelWorkloadType `json:"workload_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedModelOutputFromPb(pb *servedModelOutputPb) (*ServedModelOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedModelOutput{}
	creationTimestampField, err := identity(&pb.CreationTimestamp)
	if err != nil {
		return nil, err
	}
	if creationTimestampField != nil {
		st.CreationTimestamp = *creationTimestampField
	}
	creatorField, err := identity(&pb.Creator)
	if err != nil {
		return nil, err
	}
	if creatorField != nil {
		st.Creator = *creatorField
	}

	environmentVarsField := map[string]string{}
	for k, v := range pb.EnvironmentVars {
		item, err := identity(&v)
		if err != nil {
			return nil, err
		}
		if item != nil {
			environmentVarsField[k] = *item
		}
	}
	st.EnvironmentVars = environmentVarsField
	instanceProfileArnField, err := identity(&pb.InstanceProfileArn)
	if err != nil {
		return nil, err
	}
	if instanceProfileArnField != nil {
		st.InstanceProfileArn = *instanceProfileArnField
	}
	modelNameField, err := identity(&pb.ModelName)
	if err != nil {
		return nil, err
	}
	if modelNameField != nil {
		st.ModelName = *modelNameField
	}
	modelVersionField, err := identity(&pb.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionField != nil {
		st.ModelVersion = *modelVersionField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	scaleToZeroEnabledField, err := identity(&pb.ScaleToZeroEnabled)
	if err != nil {
		return nil, err
	}
	if scaleToZeroEnabledField != nil {
		st.ScaleToZeroEnabled = *scaleToZeroEnabledField
	}
	stateField, err := servedModelStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}
	workloadSizeField, err := identity(&pb.WorkloadSize)
	if err != nil {
		return nil, err
	}
	if workloadSizeField != nil {
		st.WorkloadSize = *workloadSizeField
	}
	workloadTypeField, err := identity(&pb.WorkloadType)
	if err != nil {
		return nil, err
	}
	if workloadTypeField != nil {
		st.WorkloadType = *workloadTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedModelOutputPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedModelOutputPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedModelSpec struct {
	// Only one of model_name and entity_name should be populated
	ModelName string
	// Only one of model_version and entity_version should be populated
	ModelVersion string

	Name string

	ForceSendFields []string
}

func servedModelSpecToPb(st *ServedModelSpec) (*servedModelSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedModelSpecPb{}
	modelNamePb, err := identity(&st.ModelName)
	if err != nil {
		return nil, err
	}
	if modelNamePb != nil {
		pb.ModelName = *modelNamePb
	}

	modelVersionPb, err := identity(&st.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionPb != nil {
		pb.ModelVersion = *modelVersionPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servedModelSpecPb struct {
	// Only one of model_name and entity_name should be populated
	ModelName string `json:"model_name,omitempty"`
	// Only one of model_version and entity_version should be populated
	ModelVersion string `json:"model_version,omitempty"`

	Name string `json:"name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servedModelSpecFromPb(pb *servedModelSpecPb) (*ServedModelSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServedModelSpec{}
	modelNameField, err := identity(&pb.ModelName)
	if err != nil {
		return nil, err
	}
	if modelNameField != nil {
		st.ModelName = *modelNameField
	}
	modelVersionField, err := identity(&pb.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionField != nil {
		st.ModelVersion = *modelVersionField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedModelSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedModelSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedModelState struct {
	Deployment ServedModelStateDeployment

	DeploymentStateMessage string

	ForceSendFields []string
}

func servedModelStateToPb(st *ServedModelState) (*servedModelStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servedModelStatePb{}
	deploymentPb, err := identity(&st.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentPb != nil {
		pb.Deployment = *deploymentPb
	}

	deploymentStateMessagePb, err := identity(&st.DeploymentStateMessage)
	if err != nil {
		return nil, err
	}
	if deploymentStateMessagePb != nil {
		pb.DeploymentStateMessage = *deploymentStateMessagePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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
	deploymentField, err := identity(&pb.Deployment)
	if err != nil {
		return nil, err
	}
	if deploymentField != nil {
		st.Deployment = *deploymentField
	}
	deploymentStateMessageField, err := identity(&pb.DeploymentStateMessage)
	if err != nil {
		return nil, err
	}
	if deploymentStateMessageField != nil {
		st.DeploymentStateMessage = *deploymentStateMessageField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servedModelStatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servedModelStatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServedModelStateDeployment string
type servedModelStateDeploymentPb string

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

func servedModelStateDeploymentToPb(st *ServedModelStateDeployment) (*servedModelStateDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servedModelStateDeploymentPb(*st)
	return &pb, nil
}

func servedModelStateDeploymentFromPb(pb *servedModelStateDeploymentPb) (*ServedModelStateDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServedModelStateDeployment(*pb)
	return &st, nil
}

type ServerLogsResponse struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	Logs string
}

func serverLogsResponseToPb(st *ServerLogsResponse) (*serverLogsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &serverLogsResponsePb{}
	logsPb, err := identity(&st.Logs)
	if err != nil {
		return nil, err
	}
	if logsPb != nil {
		pb.Logs = *logsPb
	}

	return pb, nil
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

type serverLogsResponsePb struct {
	// The most recent log lines of the model server processing invocation
	// requests.
	Logs string `json:"logs"`
}

func serverLogsResponseFromPb(pb *serverLogsResponsePb) (*ServerLogsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServerLogsResponse{}
	logsField, err := identity(&pb.Logs)
	if err != nil {
		return nil, err
	}
	if logsField != nil {
		st.Logs = *logsField
	}

	return st, nil
}

type ServingEndpoint struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway *AiGatewayConfig
	// The budget policy associated with the endpoint.
	BudgetPolicyId string
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigSummary
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp int64
	// The email of the user who created the serving endpoint.
	Creator string
	// System-generated ID of the endpoint, included to be used by the
	// Permissions API.
	Id string
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp int64
	// The name of the serving endpoint.
	Name string
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState
	// Tags attached to the serving endpoint.
	Tags []EndpointTag
	// The task type of the serving endpoint.
	Task string

	ForceSendFields []string
}

func servingEndpointToPb(st *ServingEndpoint) (*servingEndpointPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointPb{}
	aiGatewayPb, err := aiGatewayConfigToPb(st.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayPb != nil {
		pb.AiGateway = aiGatewayPb
	}

	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	configPb, err := endpointCoreConfigSummaryToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}

	creationTimestampPb, err := identity(&st.CreationTimestamp)
	if err != nil {
		return nil, err
	}
	if creationTimestampPb != nil {
		pb.CreationTimestamp = *creationTimestampPb
	}

	creatorPb, err := identity(&st.Creator)
	if err != nil {
		return nil, err
	}
	if creatorPb != nil {
		pb.Creator = *creatorPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lastUpdatedTimestampPb, err := identity(&st.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampPb != nil {
		pb.LastUpdatedTimestamp = *lastUpdatedTimestampPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	statePb, err := endpointStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	var tagsPb []endpointTagPb
	for _, item := range st.Tags {
		itemPb, err := endpointTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	taskPb, err := identity(&st.Task)
	if err != nil {
		return nil, err
	}
	if taskPb != nil {
		pb.Task = *taskPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servingEndpointPb struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway *aiGatewayConfigPb `json:"ai_gateway,omitempty"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The config that is currently being served by the endpoint.
	Config *endpointCoreConfigSummaryPb `json:"config,omitempty"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the serving endpoint.
	Creator string `json:"creator,omitempty"`
	// System-generated ID of the endpoint, included to be used by the
	// Permissions API.
	Id string `json:"id,omitempty"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The name of the serving endpoint.
	Name string `json:"name,omitempty"`
	// Information corresponding to the state of the serving endpoint.
	State *endpointStatePb `json:"state,omitempty"`
	// Tags attached to the serving endpoint.
	Tags []endpointTagPb `json:"tags,omitempty"`
	// The task type of the serving endpoint.
	Task string `json:"task,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointFromPb(pb *servingEndpointPb) (*ServingEndpoint, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpoint{}
	aiGatewayField, err := aiGatewayConfigFromPb(pb.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayField != nil {
		st.AiGateway = aiGatewayField
	}
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	configField, err := endpointCoreConfigSummaryFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	creationTimestampField, err := identity(&pb.CreationTimestamp)
	if err != nil {
		return nil, err
	}
	if creationTimestampField != nil {
		st.CreationTimestamp = *creationTimestampField
	}
	creatorField, err := identity(&pb.Creator)
	if err != nil {
		return nil, err
	}
	if creatorField != nil {
		st.Creator = *creatorField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lastUpdatedTimestampField, err := identity(&pb.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampField != nil {
		st.LastUpdatedTimestamp = *lastUpdatedTimestampField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	stateField, err := endpointStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}

	var tagsField []EndpointTag
	for _, itemPb := range pb.Tags {
		item, err := endpointTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	taskField, err := identity(&pb.Task)
	if err != nil {
		return nil, err
	}
	if taskField != nil {
		st.Task = *taskField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointAccessControlRequest struct {
	// name of the group
	GroupName string
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel
	// application ID of a service principal
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func servingEndpointAccessControlRequestToPb(st *ServingEndpointAccessControlRequest) (*servingEndpointAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointAccessControlRequestPb{}
	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servingEndpointAccessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointAccessControlRequestFromPb(pb *servingEndpointAccessControlRequestPb) (*ServingEndpointAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointAccessControlRequest{}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointAccessControlResponse struct {
	// All permissions.
	AllPermissions []ServingEndpointPermission
	// Display name of the user or service principal.
	DisplayName string
	// name of the group
	GroupName string
	// Name of the service principal.
	ServicePrincipalName string
	// name of the user
	UserName string

	ForceSendFields []string
}

func servingEndpointAccessControlResponseToPb(st *ServingEndpointAccessControlResponse) (*servingEndpointAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointAccessControlResponsePb{}

	var allPermissionsPb []servingEndpointPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := servingEndpointPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb

	displayNamePb, err := identity(&st.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNamePb != nil {
		pb.DisplayName = *displayNamePb
	}

	groupNamePb, err := identity(&st.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNamePb != nil {
		pb.GroupName = *groupNamePb
	}

	servicePrincipalNamePb, err := identity(&st.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNamePb != nil {
		pb.ServicePrincipalName = *servicePrincipalNamePb
	}

	userNamePb, err := identity(&st.UserName)
	if err != nil {
		return nil, err
	}
	if userNamePb != nil {
		pb.UserName = *userNamePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servingEndpointAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []servingEndpointPermissionPb `json:"all_permissions,omitempty"`
	// Display name of the user or service principal.
	DisplayName string `json:"display_name,omitempty"`
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Name of the service principal.
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointAccessControlResponseFromPb(pb *servingEndpointAccessControlResponsePb) (*ServingEndpointAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointAccessControlResponse{}

	var allPermissionsField []ServingEndpointPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := servingEndpointPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	displayNameField, err := identity(&pb.DisplayName)
	if err != nil {
		return nil, err
	}
	if displayNameField != nil {
		st.DisplayName = *displayNameField
	}
	groupNameField, err := identity(&pb.GroupName)
	if err != nil {
		return nil, err
	}
	if groupNameField != nil {
		st.GroupName = *groupNameField
	}
	servicePrincipalNameField, err := identity(&pb.ServicePrincipalName)
	if err != nil {
		return nil, err
	}
	if servicePrincipalNameField != nil {
		st.ServicePrincipalName = *servicePrincipalNameField
	}
	userNameField, err := identity(&pb.UserName)
	if err != nil {
		return nil, err
	}
	if userNameField != nil {
		st.UserName = *userNameField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointDetailed struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway *AiGatewayConfig
	// The budget policy associated with the endpoint.
	BudgetPolicyId string
	// The config that is currently being served by the endpoint.
	Config *EndpointCoreConfigOutput
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp int64
	// The email of the user who created the serving endpoint.
	Creator string
	// Information required to query DataPlane APIs.
	DataPlaneInfo *ModelDataPlaneInfo
	// Endpoint invocation url if route optimization is enabled for endpoint
	EndpointUrl string
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id string
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp int64
	// The name of the serving endpoint.
	Name string
	// The config that the endpoint is attempting to update to.
	PendingConfig *EndpointPendingConfig
	// The permission level of the principal making the request.
	PermissionLevel ServingEndpointDetailedPermissionLevel
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	RouteOptimized bool
	// Information corresponding to the state of the serving endpoint.
	State *EndpointState
	// Tags attached to the serving endpoint.
	Tags []EndpointTag
	// The task type of the serving endpoint.
	Task string

	ForceSendFields []string
}

func servingEndpointDetailedToPb(st *ServingEndpointDetailed) (*servingEndpointDetailedPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointDetailedPb{}
	aiGatewayPb, err := aiGatewayConfigToPb(st.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayPb != nil {
		pb.AiGateway = aiGatewayPb
	}

	budgetPolicyIdPb, err := identity(&st.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdPb != nil {
		pb.BudgetPolicyId = *budgetPolicyIdPb
	}

	configPb, err := endpointCoreConfigOutputToPb(st.Config)
	if err != nil {
		return nil, err
	}
	if configPb != nil {
		pb.Config = configPb
	}

	creationTimestampPb, err := identity(&st.CreationTimestamp)
	if err != nil {
		return nil, err
	}
	if creationTimestampPb != nil {
		pb.CreationTimestamp = *creationTimestampPb
	}

	creatorPb, err := identity(&st.Creator)
	if err != nil {
		return nil, err
	}
	if creatorPb != nil {
		pb.Creator = *creatorPb
	}

	dataPlaneInfoPb, err := modelDataPlaneInfoToPb(st.DataPlaneInfo)
	if err != nil {
		return nil, err
	}
	if dataPlaneInfoPb != nil {
		pb.DataPlaneInfo = dataPlaneInfoPb
	}

	endpointUrlPb, err := identity(&st.EndpointUrl)
	if err != nil {
		return nil, err
	}
	if endpointUrlPb != nil {
		pb.EndpointUrl = *endpointUrlPb
	}

	idPb, err := identity(&st.Id)
	if err != nil {
		return nil, err
	}
	if idPb != nil {
		pb.Id = *idPb
	}

	lastUpdatedTimestampPb, err := identity(&st.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampPb != nil {
		pb.LastUpdatedTimestamp = *lastUpdatedTimestampPb
	}

	namePb, err := identity(&st.Name)
	if err != nil {
		return nil, err
	}
	if namePb != nil {
		pb.Name = *namePb
	}

	pendingConfigPb, err := endpointPendingConfigToPb(st.PendingConfig)
	if err != nil {
		return nil, err
	}
	if pendingConfigPb != nil {
		pb.PendingConfig = pendingConfigPb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	routeOptimizedPb, err := identity(&st.RouteOptimized)
	if err != nil {
		return nil, err
	}
	if routeOptimizedPb != nil {
		pb.RouteOptimized = *routeOptimizedPb
	}

	statePb, err := endpointStateToPb(st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = statePb
	}

	var tagsPb []endpointTagPb
	for _, item := range st.Tags {
		itemPb, err := endpointTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	taskPb, err := identity(&st.Task)
	if err != nil {
		return nil, err
	}
	if taskPb != nil {
		pb.Task = *taskPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servingEndpointDetailedPb struct {
	// The AI Gateway configuration for the serving endpoint. NOTE: External
	// model, provisioned throughput, and pay-per-token endpoints are fully
	// supported; agent endpoints currently only support inference tables.
	AiGateway *aiGatewayConfigPb `json:"ai_gateway,omitempty"`
	// The budget policy associated with the endpoint.
	BudgetPolicyId string `json:"budget_policy_id,omitempty"`
	// The config that is currently being served by the endpoint.
	Config *endpointCoreConfigOutputPb `json:"config,omitempty"`
	// The timestamp when the endpoint was created in Unix time.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// The email of the user who created the serving endpoint.
	Creator string `json:"creator,omitempty"`
	// Information required to query DataPlane APIs.
	DataPlaneInfo *modelDataPlaneInfoPb `json:"data_plane_info,omitempty"`
	// Endpoint invocation url if route optimization is enabled for endpoint
	EndpointUrl string `json:"endpoint_url,omitempty"`
	// System-generated ID of the endpoint. This is used to refer to the
	// endpoint in the Permissions API
	Id string `json:"id,omitempty"`
	// The timestamp when the endpoint was last updated by a user in Unix time.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The name of the serving endpoint.
	Name string `json:"name,omitempty"`
	// The config that the endpoint is attempting to update to.
	PendingConfig *endpointPendingConfigPb `json:"pending_config,omitempty"`
	// The permission level of the principal making the request.
	PermissionLevel ServingEndpointDetailedPermissionLevel `json:"permission_level,omitempty"`
	// Boolean representing if route optimization has been enabled for the
	// endpoint
	RouteOptimized bool `json:"route_optimized,omitempty"`
	// Information corresponding to the state of the serving endpoint.
	State *endpointStatePb `json:"state,omitempty"`
	// Tags attached to the serving endpoint.
	Tags []endpointTagPb `json:"tags,omitempty"`
	// The task type of the serving endpoint.
	Task string `json:"task,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointDetailedFromPb(pb *servingEndpointDetailedPb) (*ServingEndpointDetailed, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointDetailed{}
	aiGatewayField, err := aiGatewayConfigFromPb(pb.AiGateway)
	if err != nil {
		return nil, err
	}
	if aiGatewayField != nil {
		st.AiGateway = aiGatewayField
	}
	budgetPolicyIdField, err := identity(&pb.BudgetPolicyId)
	if err != nil {
		return nil, err
	}
	if budgetPolicyIdField != nil {
		st.BudgetPolicyId = *budgetPolicyIdField
	}
	configField, err := endpointCoreConfigOutputFromPb(pb.Config)
	if err != nil {
		return nil, err
	}
	if configField != nil {
		st.Config = configField
	}
	creationTimestampField, err := identity(&pb.CreationTimestamp)
	if err != nil {
		return nil, err
	}
	if creationTimestampField != nil {
		st.CreationTimestamp = *creationTimestampField
	}
	creatorField, err := identity(&pb.Creator)
	if err != nil {
		return nil, err
	}
	if creatorField != nil {
		st.Creator = *creatorField
	}
	dataPlaneInfoField, err := modelDataPlaneInfoFromPb(pb.DataPlaneInfo)
	if err != nil {
		return nil, err
	}
	if dataPlaneInfoField != nil {
		st.DataPlaneInfo = dataPlaneInfoField
	}
	endpointUrlField, err := identity(&pb.EndpointUrl)
	if err != nil {
		return nil, err
	}
	if endpointUrlField != nil {
		st.EndpointUrl = *endpointUrlField
	}
	idField, err := identity(&pb.Id)
	if err != nil {
		return nil, err
	}
	if idField != nil {
		st.Id = *idField
	}
	lastUpdatedTimestampField, err := identity(&pb.LastUpdatedTimestamp)
	if err != nil {
		return nil, err
	}
	if lastUpdatedTimestampField != nil {
		st.LastUpdatedTimestamp = *lastUpdatedTimestampField
	}
	nameField, err := identity(&pb.Name)
	if err != nil {
		return nil, err
	}
	if nameField != nil {
		st.Name = *nameField
	}
	pendingConfigField, err := endpointPendingConfigFromPb(pb.PendingConfig)
	if err != nil {
		return nil, err
	}
	if pendingConfigField != nil {
		st.PendingConfig = pendingConfigField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	routeOptimizedField, err := identity(&pb.RouteOptimized)
	if err != nil {
		return nil, err
	}
	if routeOptimizedField != nil {
		st.RouteOptimized = *routeOptimizedField
	}
	stateField, err := endpointStateFromPb(pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = stateField
	}

	var tagsField []EndpointTag
	for _, itemPb := range pb.Tags {
		item, err := endpointTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	taskField, err := identity(&pb.Task)
	if err != nil {
		return nil, err
	}
	if taskField != nil {
		st.Task = *taskField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointDetailedPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointDetailedPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointDetailedPermissionLevel string
type servingEndpointDetailedPermissionLevelPb string

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

func servingEndpointDetailedPermissionLevelToPb(st *ServingEndpointDetailedPermissionLevel) (*servingEndpointDetailedPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingEndpointDetailedPermissionLevelPb(*st)
	return &pb, nil
}

func servingEndpointDetailedPermissionLevelFromPb(pb *servingEndpointDetailedPermissionLevelPb) (*ServingEndpointDetailedPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServingEndpointDetailedPermissionLevel(*pb)
	return &st, nil
}

type ServingEndpointPermission struct {
	Inherited bool

	InheritedFromObject []string
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel

	ForceSendFields []string
}

func servingEndpointPermissionToPb(st *ServingEndpointPermission) (*servingEndpointPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointPermissionPb{}
	inheritedPb, err := identity(&st.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedPb != nil {
		pb.Inherited = *inheritedPb
	}

	var inheritedFromObjectPb []string
	for _, item := range st.InheritedFromObject {
		itemPb, err := identity(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			inheritedFromObjectPb = append(inheritedFromObjectPb, *itemPb)
		}
	}
	pb.InheritedFromObject = inheritedFromObjectPb

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servingEndpointPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointPermissionFromPb(pb *servingEndpointPermissionPb) (*ServingEndpointPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermission{}
	inheritedField, err := identity(&pb.Inherited)
	if err != nil {
		return nil, err
	}
	if inheritedField != nil {
		st.Inherited = *inheritedField
	}

	var inheritedFromObjectField []string
	for _, itemPb := range pb.InheritedFromObject {
		item, err := identity(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			inheritedFromObjectField = append(inheritedFromObjectField, *item)
		}
	}
	st.InheritedFromObject = inheritedFromObjectField
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type ServingEndpointPermissionLevel string
type servingEndpointPermissionLevelPb string

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

func servingEndpointPermissionLevelToPb(st *ServingEndpointPermissionLevel) (*servingEndpointPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingEndpointPermissionLevelPb(*st)
	return &pb, nil
}

func servingEndpointPermissionLevelFromPb(pb *servingEndpointPermissionLevelPb) (*ServingEndpointPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServingEndpointPermissionLevel(*pb)
	return &st, nil
}

type ServingEndpointPermissions struct {
	AccessControlList []ServingEndpointAccessControlResponse

	ObjectId string

	ObjectType string

	ForceSendFields []string
}

func servingEndpointPermissionsToPb(st *ServingEndpointPermissions) (*servingEndpointPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointPermissionsPb{}

	var accessControlListPb []servingEndpointAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := servingEndpointAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	objectIdPb, err := identity(&st.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdPb != nil {
		pb.ObjectId = *objectIdPb
	}

	objectTypePb, err := identity(&st.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypePb != nil {
		pb.ObjectType = *objectTypePb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servingEndpointPermissionsPb struct {
	AccessControlList []servingEndpointAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointPermissionsFromPb(pb *servingEndpointPermissionsPb) (*ServingEndpointPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermissions{}

	var accessControlListField []ServingEndpointAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := servingEndpointAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	objectIdField, err := identity(&pb.ObjectId)
	if err != nil {
		return nil, err
	}
	if objectIdField != nil {
		st.ObjectId = *objectIdField
	}
	objectTypeField, err := identity(&pb.ObjectType)
	if err != nil {
		return nil, err
	}
	if objectTypeField != nil {
		st.ObjectType = *objectTypeField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointPermissionsDescription struct {
	Description string
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel

	ForceSendFields []string
}

func servingEndpointPermissionsDescriptionToPb(st *ServingEndpointPermissionsDescription) (*servingEndpointPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointPermissionsDescriptionPb{}
	descriptionPb, err := identity(&st.Description)
	if err != nil {
		return nil, err
	}
	if descriptionPb != nil {
		pb.Description = *descriptionPb
	}

	permissionLevelPb, err := identity(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type servingEndpointPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel ServingEndpointPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func servingEndpointPermissionsDescriptionFromPb(pb *servingEndpointPermissionsDescriptionPb) (*ServingEndpointPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermissionsDescription{}
	descriptionField, err := identity(&pb.Description)
	if err != nil {
		return nil, err
	}
	if descriptionField != nil {
		st.Description = *descriptionField
	}
	permissionLevelField, err := identity(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *servingEndpointPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st servingEndpointPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ServingEndpointPermissionsRequest struct {
	AccessControlList []ServingEndpointAccessControlRequest
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string
}

func servingEndpointPermissionsRequestToPb(st *ServingEndpointPermissionsRequest) (*servingEndpointPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &servingEndpointPermissionsRequestPb{}

	var accessControlListPb []servingEndpointAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := servingEndpointAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	servingEndpointIdPb, err := identity(&st.ServingEndpointId)
	if err != nil {
		return nil, err
	}
	if servingEndpointIdPb != nil {
		pb.ServingEndpointId = *servingEndpointIdPb
	}

	return pb, nil
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

type servingEndpointPermissionsRequestPb struct {
	AccessControlList []servingEndpointAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The serving endpoint for which to get or manage permissions.
	ServingEndpointId string `json:"-" url:"-"`
}

func servingEndpointPermissionsRequestFromPb(pb *servingEndpointPermissionsRequestPb) (*ServingEndpointPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ServingEndpointPermissionsRequest{}

	var accessControlListField []ServingEndpointAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := servingEndpointAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	servingEndpointIdField, err := identity(&pb.ServingEndpointId)
	if err != nil {
		return nil, err
	}
	if servingEndpointIdField != nil {
		st.ServingEndpointId = *servingEndpointIdField
	}

	return st, nil
}

// Please keep this in sync with with workload types in
// InferenceEndpointEntities.scala
type ServingModelWorkloadType string
type servingModelWorkloadTypePb string

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

// Type always returns ServingModelWorkloadType to satisfy [pflag.Value] interface
func (f *ServingModelWorkloadType) Type() string {
	return "ServingModelWorkloadType"
}

func servingModelWorkloadTypeToPb(st *ServingModelWorkloadType) (*servingModelWorkloadTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := servingModelWorkloadTypePb(*st)
	return &pb, nil
}

func servingModelWorkloadTypeFromPb(pb *servingModelWorkloadTypePb) (*ServingModelWorkloadType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ServingModelWorkloadType(*pb)
	return &st, nil
}

type TrafficConfig struct {
	// The list of routes that define traffic to each served entity.
	Routes []Route
}

func trafficConfigToPb(st *TrafficConfig) (*trafficConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &trafficConfigPb{}

	var routesPb []routePb
	for _, item := range st.Routes {
		itemPb, err := routeToPb(&item)
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

type trafficConfigPb struct {
	// The list of routes that define traffic to each served entity.
	Routes []routePb `json:"routes,omitempty"`
}

func trafficConfigFromPb(pb *trafficConfigPb) (*TrafficConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TrafficConfig{}

	var routesField []Route
	for _, itemPb := range pb.Routes {
		item, err := routeFromPb(&itemPb)
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

type V1ResponseChoiceElement struct {
	// The finish reason returned by the endpoint.
	FinishReason string
	// The index of the choice in the __chat or completions__ response.
	Index int
	// The logprobs returned only by the __completions__ endpoint.
	Logprobs int
	// The message response from the __chat__ endpoint.
	Message *ChatMessage
	// The text response from the __completions__ endpoint.
	Text string

	ForceSendFields []string
}

func v1ResponseChoiceElementToPb(st *V1ResponseChoiceElement) (*v1ResponseChoiceElementPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &v1ResponseChoiceElementPb{}
	finishReasonPb, err := identity(&st.FinishReason)
	if err != nil {
		return nil, err
	}
	if finishReasonPb != nil {
		pb.FinishReason = *finishReasonPb
	}

	indexPb, err := identity(&st.Index)
	if err != nil {
		return nil, err
	}
	if indexPb != nil {
		pb.Index = *indexPb
	}

	logprobsPb, err := identity(&st.Logprobs)
	if err != nil {
		return nil, err
	}
	if logprobsPb != nil {
		pb.Logprobs = *logprobsPb
	}

	messagePb, err := chatMessageToPb(st.Message)
	if err != nil {
		return nil, err
	}
	if messagePb != nil {
		pb.Message = messagePb
	}

	textPb, err := identity(&st.Text)
	if err != nil {
		return nil, err
	}
	if textPb != nil {
		pb.Text = *textPb
	}

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
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

type v1ResponseChoiceElementPb struct {
	// The finish reason returned by the endpoint.
	FinishReason string `json:"finishReason,omitempty"`
	// The index of the choice in the __chat or completions__ response.
	Index int `json:"index,omitempty"`
	// The logprobs returned only by the __completions__ endpoint.
	Logprobs int `json:"logprobs,omitempty"`
	// The message response from the __chat__ endpoint.
	Message *chatMessagePb `json:"message,omitempty"`
	// The text response from the __completions__ endpoint.
	Text string `json:"text,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func v1ResponseChoiceElementFromPb(pb *v1ResponseChoiceElementPb) (*V1ResponseChoiceElement, error) {
	if pb == nil {
		return nil, nil
	}
	st := &V1ResponseChoiceElement{}
	finishReasonField, err := identity(&pb.FinishReason)
	if err != nil {
		return nil, err
	}
	if finishReasonField != nil {
		st.FinishReason = *finishReasonField
	}
	indexField, err := identity(&pb.Index)
	if err != nil {
		return nil, err
	}
	if indexField != nil {
		st.Index = *indexField
	}
	logprobsField, err := identity(&pb.Logprobs)
	if err != nil {
		return nil, err
	}
	if logprobsField != nil {
		st.Logprobs = *logprobsField
	}
	messageField, err := chatMessageFromPb(pb.Message)
	if err != nil {
		return nil, err
	}
	if messageField != nil {
		st.Message = messageField
	}
	textField, err := identity(&pb.Text)
	if err != nil {
		return nil, err
	}
	if textField != nil {
		st.Text = *textField
	}

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *v1ResponseChoiceElementPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st v1ResponseChoiceElementPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}
