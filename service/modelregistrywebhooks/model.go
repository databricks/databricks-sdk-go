// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package modelregistrywebhooks

// all definitions in this file are in alphabetical order

type CreateRegistryWebhookRequest struct {
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`

	Events []RegistryWebhookEvent `json:"events"`

	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"`

	JobSpec *JobSpec `json:"job_spec,omitempty"`
	// Name of the model whose events would trigger this webhook.
	ModelName string `json:"model_name,omitempty"`

	Status RegistryWebhookStatus `json:"status,omitempty"`
}

type CreateRegistryWebhookResponse struct {
	Webhook any /* ERROR */ `json:"webhook,omitempty"`
}

type DeleteRegistryWebhookRequest struct {
	// Webhook ID
	Id string `json:"id"`
}

type HttpUrlSpec struct {
	// Value of the authorization header that should be sent in the request sent
	// by the wehbook. It should be of the form `&#34;&lt;auth type&gt; &lt;credentials&gt;&#34;`.
	// If set to an empty string, no authorization header will be included in
	// the request.
	Authorization string `json:"authorization,omitempty"`
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification bool `json:"enable_ssl_verification,omitempty"`
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { &#34;X-Databricks-Signature&#34;:
	// $encoded_payload }.
	Secret string `json:"secret,omitempty"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url string `json:"url"`
}

type JobSpec struct {
	// The personal access token used to authorize webhook&#39;s job runs.
	AccessToken string `json:"access_token"`
	// ID of the job that the webhook runs.
	JobId string `json:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job?s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl string `json:"workspace_url,omitempty"`
}

type ListRegistryWebhooksRequest struct {
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	Events []RegistryWebhookEvent `json:"events,omitempty"`
	// Name of the model whose events would trigger this webhook.
	ModelName string `json:"model_name,omitempty"`
}

type ListRegistryWebhooksResponse struct {
	// Array of registry webhooks.
	Webhooks any /* MISSING TYPE */ `json:"webhooks,omitempty"`
}

type RegistryWebhookEvent string

const RegistryWebhookEventModelVersionCreated RegistryWebhookEvent = `MODEL_VERSION_CREATED`

const RegistryWebhookEventModelVersionTransitionedStage RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_STAGE`

const RegistryWebhookEventTransitionRequestCreated RegistryWebhookEvent = `TRANSITION_REQUEST_CREATED`

const RegistryWebhookEventCommentCreated RegistryWebhookEvent = `COMMENT_CREATED`

const RegistryWebhookEventRegisteredModelCreated RegistryWebhookEvent = `REGISTERED_MODEL_CREATED`

const RegistryWebhookEventModelVersionTagSet RegistryWebhookEvent = `MODEL_VERSION_TAG_SET`

const RegistryWebhookEventModelVersionTransitionedToStaging RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_STAGING`

const RegistryWebhookEventModelVersionTransitionedToProduction RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`

const RegistryWebhookEventModelVersionTransitionedToArchived RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`

const RegistryWebhookEventTransitionRequestToStagingCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_STAGING_CREATED`

const RegistryWebhookEventTransitionRequestToProductionCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`

const RegistryWebhookEventTransitionRequestToArchivedCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`

// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A new
// model version was created for the associated model. *
// `MODEL_VERSION_TRANSITIONED_STAGE`: A model version?s stage was changed. *
// `TRANSITION_REQUEST_CREATED`: A user requested a model version?s stage be
// transitioned. * `COMMENT_CREATED`: A user wrote a comment on a registered
// model. * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
// event type can only be specified for a registry-wide webhook, which can be
// created by not specifying a model name in the create request. *
// `MODEL_VERSION_TAG_SET`: A user set a tag on the model version. *
// `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was transitioned to
// staging. * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
// transitioned to production. * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A
// model version was archived. * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user
// requested a model version be transitioned to staging. *
// `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model version
// be transitioned to production. * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A
// user requested a model version be archived.
type RegistryWebhookEvents []RegistryWebhookEvent

// Enable or disable triggering the webhook, or put the webhook into test mode.
// The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an associated
// event happens. * `DISABLED`: Webhook is not triggered. * `TEST_MODE`: Webhook
// can be triggered through the test endpoint, but is not triggered on a real
// event.
type RegistryWebhookStatus string

// Enable or disable triggering the webhook, or put the webhook into test mode.
// The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an associated
// event happens. * `DISABLED`: Webhook is not triggered. * `TEST_MODE`: Webhook
// can be triggered through the test endpoint, but is not triggered on a real
// event.
const RegistryWebhookStatusActive RegistryWebhookStatus = `ACTIVE`

// Enable or disable triggering the webhook, or put the webhook into test mode.
// The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an associated
// event happens. * `DISABLED`: Webhook is not triggered. * `TEST_MODE`: Webhook
// can be triggered through the test endpoint, but is not triggered on a real
// event.
const RegistryWebhookStatusDisabled RegistryWebhookStatus = `DISABLED`

// Enable or disable triggering the webhook, or put the webhook into test mode.
// The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an associated
// event happens. * `DISABLED`: Webhook is not triggered. * `TEST_MODE`: Webhook
// can be triggered through the test endpoint, but is not triggered on a real
// event.
const RegistryWebhookStatusTestMode RegistryWebhookStatus = `TEST_MODE`

// Test webhook response object.
type TestRegistryWebhook struct {
	// Body of the response from the webhook URL
	Body string `json:"body,omitempty"`
	// Status code returned by the webhook URL
	StatusCode int `json:"status_code,omitempty"`
}

type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event RegistryWebhookEvent `json:"event,omitempty"`
	// Webhook ID
	Id string `json:"id"`
}

type TestRegistryWebhookResponse struct {
	Webhook *TestRegistryWebhook `json:"webhook,omitempty"`
}

type UpdateRegistryWebhookRequest struct {
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`

	Events []RegistryWebhookEvent `json:"events,omitempty"`

	HttpUrlSpec *HttpUrlSpec `json:"http_url_spec,omitempty"`
	// Webhook ID
	Id string `json:"id"`

	JobSpec *JobSpec `json:"job_spec,omitempty"`

	Status RegistryWebhookStatus `json:"status,omitempty"`
}
