// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package supervisoragents

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/common/types/fieldmask"
	"github.com/databricks/databricks-sdk-go/common/types/time"
	"github.com/databricks/databricks-sdk-go/marshal"
)

// Databricks app. Supported app: custom mcp, custom agent.
type App struct {
	// App name
	Name string `json:"name"`
}

type CreateExampleRequest struct {
	// The example to create under the parent Supervisor Agent.
	Example Example `json:"example"`
	// Parent resource where this example will be created. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent string `json:"-" url:"-"`
}

type CreateSupervisorAgentRequest struct {
	// The Supervisor Agent to create.
	SupervisorAgent SupervisorAgent `json:"supervisor_agent"`
}

type CreateToolRequest struct {
	// Parent resource where this tool will be created. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent string `json:"-" url:"-"`

	Tool Tool `json:"tool"`
	// The ID to use for the tool, which will become the final component of the
	// tool's resource name.
	ToolId string `json:"-" url:"tool_id"`
}

type DeleteExampleRequest struct {
	// The resource name of the example to delete. Format:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name string `json:"-" url:"-"`
}

type DeleteSupervisorAgentRequest struct {
	// The resource name of the Supervisor Agent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name string `json:"-" url:"-"`
}

type DeleteToolRequest struct {
	// The resource name of the Tool. Format:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name string `json:"-" url:"-"`
}

// An example associated with a Supervisor Agent. Contains a question and
// guidelines for how the agent should respond.
type Example struct {
	// The universally unique identifier (UUID) of the example.
	ExampleId string `json:"example_id,omitempty"`
	// Guidelines for answering the question.
	Guidelines []string `json:"guidelines"`
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name string `json:"name,omitempty"`
	// The example question.
	Question string `json:"question"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Example) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Example) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type GenieSpace struct {
	// The ID of the genie space.
	Id string `json:"id"`
}

type GetExampleRequest struct {
	// The resource name of the example. Format:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name string `json:"-" url:"-"`
}

type GetSupervisorAgentPermissionLevelsRequest struct {
	// The supervisor agent for which to get or manage permissions.
	SupervisorAgentId string `json:"-" url:"-"`
}

type GetSupervisorAgentPermissionLevelsResponse struct {
	// Specific permission levels
	PermissionLevels []SupervisorAgentPermissionsDescription `json:"permission_levels,omitempty"`
}

type GetSupervisorAgentPermissionsRequest struct {
	// The supervisor agent for which to get or manage permissions.
	SupervisorAgentId string `json:"-" url:"-"`
}

type GetSupervisorAgentRequest struct {
	// The resource name of the Supervisor Agent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name string `json:"-" url:"-"`
}

type GetToolRequest struct {
	// The resource name of the Tool. Format:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name string `json:"-" url:"-"`
}

type KnowledgeAssistant struct {
	// The ID of the knowledge assistant.
	KnowledgeAssistantId string `json:"knowledge_assistant_id"`
	// Deprecated: use knowledge_assistant_id instead.
	ServingEndpointName string `json:"serving_endpoint_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *KnowledgeAssistant) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s KnowledgeAssistant) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListExamplesRequest struct {
	// The maximum number of examples to return. If unspecified, at most 100
	// examples will be returned. The maximum value is 100; values above 100
	// will be coerced to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListExamples` call. Provide this
	// to retrieve the subsequent page. If unspecified, the first page will be
	// returned.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent resource to list from. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExamplesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExamplesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// A list of Supervisor Agent examples.
type ListExamplesResponse struct {
	Examples []Example `json:"examples,omitempty"`

	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListExamplesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListExamplesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSupervisorAgentsRequest struct {
	// The maximum number of supervisor agents to return. If unspecified, at
	// most 100 supervisor agents will be returned. The maximum value is 100;
	// values above 100 will be coerced to 100.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// A page token, received from a previous `ListSupervisorAgents` call.
	// Provide this to retrieve the subsequent page. If unspecified, the first
	// page will be returned.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSupervisorAgentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSupervisorAgentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListSupervisorAgentsResponse struct {
	// A token that can be sent as `page_token` to retrieve the next page. If
	// this field is omitted, there are no subsequent pages.
	NextPageToken string `json:"next_page_token,omitempty"`

	SupervisorAgents []SupervisorAgent `json:"supervisor_agents,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListSupervisorAgentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListSupervisorAgentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListToolsRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`
	// Parent resource to list from. Format:
	// supervisor-agents/{supervisor_agent_id}
	Parent string `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListToolsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListToolsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListToolsResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	Tools []Tool `json:"tools,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListToolsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListToolsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgent struct {
	// Creation timestamp.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// The creator of the Supervisor Agent.
	Creator string `json:"creator,omitempty"`
	// Description of what this agent can do (user-facing).
	Description string `json:"description,omitempty"`
	// The display name of the Supervisor Agent, unique at workspace level.
	DisplayName string `json:"display_name"`
	// The name of the supervisor agent's serving endpoint.
	EndpointName string `json:"endpoint_name,omitempty"`
	// The MLflow experiment ID.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Deprecated: Use supervisor_agent_id instead.
	Id string `json:"id,omitempty"`
	// Optional natural-language instructions for the supervisor agent.
	Instructions string `json:"instructions,omitempty"`
	// The resource name of the SupervisorAgent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name string `json:"name,omitempty"`
	// The universally unique identifier (UUID) of the Supervisor Agent.
	SupervisorAgentId string `json:"supervisor_agent_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgent) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgent) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgentAccessControlRequest struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`

	PermissionLevel SupervisorAgentPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgentAccessControlRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgentAccessControlRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgentAccessControlResponse struct {
	// All permissions.
	AllPermissions []SupervisorAgentPermission `json:"all_permissions,omitempty"`
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

func (s *SupervisorAgentAccessControlResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgentAccessControlResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgentPermission struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`

	PermissionLevel SupervisorAgentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgentPermission) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgentPermission) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Permission level
type SupervisorAgentPermissionLevel string

const SupervisorAgentPermissionLevelCanManage SupervisorAgentPermissionLevel = `CAN_MANAGE`

const SupervisorAgentPermissionLevelCanQuery SupervisorAgentPermissionLevel = `CAN_QUERY`

// String representation for [fmt.Print]
func (f *SupervisorAgentPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *SupervisorAgentPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_MANAGE`, `CAN_QUERY`:
		*f = SupervisorAgentPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_MANAGE", "CAN_QUERY"`, v)
	}
}

// Values returns all possible values for SupervisorAgentPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *SupervisorAgentPermissionLevel) Values() []SupervisorAgentPermissionLevel {
	return []SupervisorAgentPermissionLevel{
		SupervisorAgentPermissionLevelCanManage,
		SupervisorAgentPermissionLevelCanQuery,
	}
}

// Type always returns SupervisorAgentPermissionLevel to satisfy [pflag.Value] interface
func (f *SupervisorAgentPermissionLevel) Type() string {
	return "SupervisorAgentPermissionLevel"
}

type SupervisorAgentPermissions struct {
	AccessControlList []SupervisorAgentAccessControlResponse `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgentPermissions) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgentPermissions) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgentPermissionsDescription struct {
	Description string `json:"description,omitempty"`

	PermissionLevel SupervisorAgentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *SupervisorAgentPermissionsDescription) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s SupervisorAgentPermissionsDescription) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type SupervisorAgentPermissionsRequest struct {
	AccessControlList []SupervisorAgentAccessControlRequest `json:"access_control_list,omitempty"`
	// The supervisor agent for which to get or manage permissions.
	SupervisorAgentId string `json:"-" url:"-"`
}

type Tool struct {
	App *App `json:"app,omitempty"`
	// Description of what this tool does (user-facing).
	Description string `json:"description,omitempty"`

	GenieSpace *GenieSpace `json:"genie_space,omitempty"`
	// Deprecated: Use tool_id instead.
	Id string `json:"id,omitempty"`

	KnowledgeAssistant *KnowledgeAssistant `json:"knowledge_assistant,omitempty"`
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name string `json:"name,omitempty"`
	// User specified id of the Tool.
	ToolId string `json:"tool_id,omitempty"`
	// Tool type. Must be one of: "genie_space", "knowledge_assistant",
	// "uc_function", "uc_connection", "app", "volume", "lakeview_dashboard",
	// "serving_endpoint", "uc_table", "vector_search_index", "catalog",
	// "schema", "supervisor_agent".
	ToolType string `json:"tool_type"`

	UcConnection *UcConnection `json:"uc_connection,omitempty"`

	UcFunction *UcFunction `json:"uc_function,omitempty"`

	Volume *Volume `json:"volume,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *Tool) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s Tool) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

// Databricks UC connection. Supported connection: external mcp server.
type UcConnection struct {
	Name string `json:"name"`
}

type UcFunction struct {
	// Full uc function name
	Name string `json:"name"`
}

type UpdateExampleRequest struct {
	Example Example `json:"example"`
	// The resource name of the example to update. Format:
	// supervisor-agents/{supervisor_agent_id}/examples/{example_id}
	Name string `json:"-" url:"-"`
	// Comma-delimited list of fields to update on the example. Allowed values:
	// `question`, `guidelines`. Examples: - `question` - `question,guidelines`
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

type UpdateSupervisorAgentRequest struct {
	// The resource name of the SupervisorAgent. Format:
	// supervisor-agents/{supervisor_agent_id}
	Name string `json:"-" url:"-"`
	// The SupervisorAgent to update.
	SupervisorAgent SupervisorAgent `json:"supervisor_agent"`
	// Field mask for fields to be updated.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

type UpdateToolRequest struct {
	// Full resource name:
	// supervisor-agents/{supervisor_agent_id}/tools/{tool_id}
	Name string `json:"-" url:"-"`
	// The Tool to update.
	Tool Tool `json:"tool"`
	// Field mask for fields to be updated.
	UpdateMask fieldmask.FieldMask `json:"-" url:"update_mask"`
}

type Volume struct {
	// Full uc volume name
	Name string `json:"name"`
}
