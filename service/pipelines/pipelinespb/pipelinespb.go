// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelinespb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute/computepb"
)

type CreatePipelinePb struct {
	AllowDuplicateNames bool                                  `json:"allow_duplicate_names,omitempty"`
	BudgetPolicyId      string                                `json:"budget_policy_id,omitempty"`
	Catalog             string                                `json:"catalog,omitempty"`
	Channel             string                                `json:"channel,omitempty"`
	Clusters            []PipelineClusterPb                   `json:"clusters,omitempty"`
	Configuration       map[string]string                     `json:"configuration,omitempty"`
	Continuous          bool                                  `json:"continuous,omitempty"`
	Deployment          *PipelineDeploymentPb                 `json:"deployment,omitempty"`
	Development         bool                                  `json:"development,omitempty"`
	DryRun              bool                                  `json:"dry_run,omitempty"`
	Edition             string                                `json:"edition,omitempty"`
	Environment         *PipelinesEnvironmentPb               `json:"environment,omitempty"`
	EventLog            *EventLogSpecPb                       `json:"event_log,omitempty"`
	Filters             *FiltersPb                            `json:"filters,omitempty"`
	GatewayDefinition   *IngestionGatewayPipelineDefinitionPb `json:"gateway_definition,omitempty"`
	Id                  string                                `json:"id,omitempty"`
	IngestionDefinition *IngestionPipelineDefinitionPb        `json:"ingestion_definition,omitempty"`
	Libraries           []PipelineLibraryPb                   `json:"libraries,omitempty"`
	Name                string                                `json:"name,omitempty"`
	Notifications       []NotificationsPb                     `json:"notifications,omitempty"`
	Photon              bool                                  `json:"photon,omitempty"`
	RestartWindow       *RestartWindowPb                      `json:"restart_window,omitempty"`
	RootPath            string                                `json:"root_path,omitempty"`
	RunAs               *RunAsPb                              `json:"run_as,omitempty"`
	Schema              string                                `json:"schema,omitempty"`
	Serverless          bool                                  `json:"serverless,omitempty"`
	Storage             string                                `json:"storage,omitempty"`
	Tags                map[string]string                     `json:"tags,omitempty"`
	Target              string                                `json:"target,omitempty"`
	Trigger             *PipelineTriggerPb                    `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePipelinePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePipelinePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreatePipelineResponsePb struct {
	EffectiveSettings *PipelineSpecPb `json:"effective_settings,omitempty"`
	PipelineId        string          `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CreatePipelineResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CreatePipelineResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CronTriggerPb struct {
	QuartzCronSchedule string `json:"quartz_cron_schedule,omitempty"`
	TimezoneId         string `json:"timezone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *CronTriggerPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st CronTriggerPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DataPlaneIdPb struct {
	Instance string `json:"instance,omitempty"`
	SeqNo    int64  `json:"seq_no,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *DataPlaneIdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st DataPlaneIdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DayOfWeekPb string

const DayOfWeekFriday DayOfWeekPb = `FRIDAY`

const DayOfWeekMonday DayOfWeekPb = `MONDAY`

const DayOfWeekSaturday DayOfWeekPb = `SATURDAY`

const DayOfWeekSunday DayOfWeekPb = `SUNDAY`

const DayOfWeekThursday DayOfWeekPb = `THURSDAY`

const DayOfWeekTuesday DayOfWeekPb = `TUESDAY`

const DayOfWeekWednesday DayOfWeekPb = `WEDNESDAY`

type DeletePipelineRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

type DeletePipelineResponsePb struct {
}

type DeploymentKindPb string

const DeploymentKindBundle DeploymentKindPb = `BUNDLE`

type EditPipelinePb struct {
	AllowDuplicateNames  bool                                  `json:"allow_duplicate_names,omitempty"`
	BudgetPolicyId       string                                `json:"budget_policy_id,omitempty"`
	Catalog              string                                `json:"catalog,omitempty"`
	Channel              string                                `json:"channel,omitempty"`
	Clusters             []PipelineClusterPb                   `json:"clusters,omitempty"`
	Configuration        map[string]string                     `json:"configuration,omitempty"`
	Continuous           bool                                  `json:"continuous,omitempty"`
	Deployment           *PipelineDeploymentPb                 `json:"deployment,omitempty"`
	Development          bool                                  `json:"development,omitempty"`
	Edition              string                                `json:"edition,omitempty"`
	Environment          *PipelinesEnvironmentPb               `json:"environment,omitempty"`
	EventLog             *EventLogSpecPb                       `json:"event_log,omitempty"`
	ExpectedLastModified int64                                 `json:"expected_last_modified,omitempty"`
	Filters              *FiltersPb                            `json:"filters,omitempty"`
	GatewayDefinition    *IngestionGatewayPipelineDefinitionPb `json:"gateway_definition,omitempty"`
	Id                   string                                `json:"id,omitempty"`
	IngestionDefinition  *IngestionPipelineDefinitionPb        `json:"ingestion_definition,omitempty"`
	Libraries            []PipelineLibraryPb                   `json:"libraries,omitempty"`
	Name                 string                                `json:"name,omitempty"`
	Notifications        []NotificationsPb                     `json:"notifications,omitempty"`
	Photon               bool                                  `json:"photon,omitempty"`
	PipelineId           string                                `json:"-" url:"-"`
	RestartWindow        *RestartWindowPb                      `json:"restart_window,omitempty"`
	RootPath             string                                `json:"root_path,omitempty"`
	RunAs                *RunAsPb                              `json:"run_as,omitempty"`
	Schema               string                                `json:"schema,omitempty"`
	Serverless           bool                                  `json:"serverless,omitempty"`
	Storage              string                                `json:"storage,omitempty"`
	Tags                 map[string]string                     `json:"tags,omitempty"`
	Target               string                                `json:"target,omitempty"`
	Trigger              *PipelineTriggerPb                    `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EditPipelinePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EditPipelinePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EditPipelineResponsePb struct {
}

type ErrorDetailPb struct {
	Exceptions []SerializedExceptionPb `json:"exceptions,omitempty"`
	Fatal      bool                    `json:"fatal,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ErrorDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ErrorDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type EventLevelPb string

const EventLevelError EventLevelPb = `ERROR`

const EventLevelInfo EventLevelPb = `INFO`

const EventLevelMetrics EventLevelPb = `METRICS`

const EventLevelWarn EventLevelPb = `WARN`

type EventLogSpecPb struct {
	Catalog string `json:"catalog,omitempty"`
	Name    string `json:"name,omitempty"`
	Schema  string `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *EventLogSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st EventLogSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FileLibraryPb struct {
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *FileLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st FileLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type FiltersPb struct {
	Exclude []string `json:"exclude,omitempty"`
	Include []string `json:"include,omitempty"`
}

type GetPipelinePermissionLevelsRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

type GetPipelinePermissionLevelsResponsePb struct {
	PermissionLevels []PipelinePermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

type GetPipelinePermissionsRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

type GetPipelineRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

type GetPipelineResponsePb struct {
	Cause                   string                      `json:"cause,omitempty"`
	ClusterId               string                      `json:"cluster_id,omitempty"`
	CreatorUserName         string                      `json:"creator_user_name,omitempty"`
	EffectiveBudgetPolicyId string                      `json:"effective_budget_policy_id,omitempty"`
	Health                  GetPipelineResponseHealthPb `json:"health,omitempty"`
	LastModified            int64                       `json:"last_modified,omitempty"`
	LatestUpdates           []UpdateStateInfoPb         `json:"latest_updates,omitempty"`
	Name                    string                      `json:"name,omitempty"`
	PipelineId              string                      `json:"pipeline_id,omitempty"`
	RunAs                   *RunAsPb                    `json:"run_as,omitempty"`
	RunAsUserName           string                      `json:"run_as_user_name,omitempty"`
	Spec                    *PipelineSpecPb             `json:"spec,omitempty"`
	State                   PipelineStatePb             `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *GetPipelineResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st GetPipelineResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetPipelineResponseHealthPb string

const GetPipelineResponseHealthHealthy GetPipelineResponseHealthPb = `HEALTHY`

const GetPipelineResponseHealthUnhealthy GetPipelineResponseHealthPb = `UNHEALTHY`

type GetUpdateRequestPb struct {
	PipelineId string `json:"-" url:"-"`
	UpdateId   string `json:"-" url:"-"`
}

type GetUpdateResponsePb struct {
	Update *UpdateInfoPb `json:"update,omitempty"`
}

type IngestionConfigPb struct {
	Report *ReportSpecPb `json:"report,omitempty"`
	Schema *SchemaSpecPb `json:"schema,omitempty"`
	Table  *TableSpecPb  `json:"table,omitempty"`
}

type IngestionGatewayPipelineDefinitionPb struct {
	ConnectionId          string `json:"connection_id,omitempty"`
	ConnectionName        string `json:"connection_name"`
	GatewayStorageCatalog string `json:"gateway_storage_catalog"`
	GatewayStorageName    string `json:"gateway_storage_name,omitempty"`
	GatewayStorageSchema  string `json:"gateway_storage_schema"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *IngestionGatewayPipelineDefinitionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st IngestionGatewayPipelineDefinitionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type IngestionPipelineDefinitionPb struct {
	ConnectionName     string                 `json:"connection_name,omitempty"`
	IngestionGatewayId string                 `json:"ingestion_gateway_id,omitempty"`
	Objects            []IngestionConfigPb    `json:"objects,omitempty"`
	SourceType         IngestionSourceTypePb  `json:"source_type,omitempty"`
	TableConfiguration *TableSpecificConfigPb `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *IngestionPipelineDefinitionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st IngestionPipelineDefinitionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigPb struct {
	CursorColumns                        []string `json:"cursor_columns,omitempty"`
	DeletionCondition                    string   `json:"deletion_condition,omitempty"`
	HardDeletionSyncMinIntervalInSeconds int64    `json:"hard_deletion_sync_min_interval_in_seconds,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type IngestionSourceTypePb string

const IngestionSourceTypeBigquery IngestionSourceTypePb = `BIGQUERY`

const IngestionSourceTypeConfluence IngestionSourceTypePb = `CONFLUENCE`

const IngestionSourceTypeDynamics365 IngestionSourceTypePb = `DYNAMICS365`

const IngestionSourceTypeGa4RawData IngestionSourceTypePb = `GA4_RAW_DATA`

const IngestionSourceTypeManagedPostgresql IngestionSourceTypePb = `MANAGED_POSTGRESQL`

const IngestionSourceTypeMetaMarketing IngestionSourceTypePb = `META_MARKETING`

const IngestionSourceTypeMysql IngestionSourceTypePb = `MYSQL`

const IngestionSourceTypeNetsuite IngestionSourceTypePb = `NETSUITE`

const IngestionSourceTypeOracle IngestionSourceTypePb = `ORACLE`

const IngestionSourceTypePostgresql IngestionSourceTypePb = `POSTGRESQL`

const IngestionSourceTypeRedshift IngestionSourceTypePb = `REDSHIFT`

const IngestionSourceTypeSalesforce IngestionSourceTypePb = `SALESFORCE`

const IngestionSourceTypeServicenow IngestionSourceTypePb = `SERVICENOW`

const IngestionSourceTypeSharepoint IngestionSourceTypePb = `SHAREPOINT`

const IngestionSourceTypeSqldw IngestionSourceTypePb = `SQLDW`

const IngestionSourceTypeSqlserver IngestionSourceTypePb = `SQLSERVER`

const IngestionSourceTypeTeradata IngestionSourceTypePb = `TERADATA`

const IngestionSourceTypeWorkdayRaas IngestionSourceTypePb = `WORKDAY_RAAS`

type ListPipelineEventsRequestPb struct {
	Filter     string   `json:"-" url:"filter,omitempty"`
	MaxResults int      `json:"-" url:"max_results,omitempty"`
	OrderBy    []string `json:"-" url:"order_by,omitempty"`
	PageToken  string   `json:"-" url:"page_token,omitempty"`
	PipelineId string   `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListPipelineEventsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListPipelineEventsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPipelineEventsResponsePb struct {
	Events        []PipelineEventPb `json:"events,omitempty"`
	NextPageToken string            `json:"next_page_token,omitempty"`
	PrevPageToken string            `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListPipelineEventsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListPipelineEventsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPipelinesRequestPb struct {
	Filter     string   `json:"-" url:"filter,omitempty"`
	MaxResults int      `json:"-" url:"max_results,omitempty"`
	OrderBy    []string `json:"-" url:"order_by,omitempty"`
	PageToken  string   `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListPipelinesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListPipelinesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListPipelinesResponsePb struct {
	NextPageToken string                `json:"next_page_token,omitempty"`
	Statuses      []PipelineStateInfoPb `json:"statuses,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListPipelinesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListPipelinesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListUpdatesRequestPb struct {
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	PipelineId    string `json:"-" url:"-"`
	UntilUpdateId string `json:"-" url:"until_update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListUpdatesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListUpdatesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListUpdatesResponsePb struct {
	NextPageToken string         `json:"next_page_token,omitempty"`
	PrevPageToken string         `json:"prev_page_token,omitempty"`
	Updates       []UpdateInfoPb `json:"updates,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListUpdatesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListUpdatesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ManualTriggerPb struct {
}

type MaturityLevelPb string

const MaturityLevelDeprecated MaturityLevelPb = `DEPRECATED`

const MaturityLevelEvolving MaturityLevelPb = `EVOLVING`

const MaturityLevelStable MaturityLevelPb = `STABLE`

type NotebookLibraryPb struct {
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *NotebookLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st NotebookLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type NotificationsPb struct {
	Alerts          []string `json:"alerts,omitempty"`
	EmailRecipients []string `json:"email_recipients,omitempty"`
}

type OriginPb struct {
	BatchId             int64  `json:"batch_id,omitempty"`
	Cloud               string `json:"cloud,omitempty"`
	ClusterId           string `json:"cluster_id,omitempty"`
	DatasetName         string `json:"dataset_name,omitempty"`
	FlowId              string `json:"flow_id,omitempty"`
	FlowName            string `json:"flow_name,omitempty"`
	Host                string `json:"host,omitempty"`
	MaintenanceId       string `json:"maintenance_id,omitempty"`
	MaterializationName string `json:"materialization_name,omitempty"`
	OrgId               int64  `json:"org_id,omitempty"`
	PipelineId          string `json:"pipeline_id,omitempty"`
	PipelineName        string `json:"pipeline_name,omitempty"`
	Region              string `json:"region,omitempty"`
	RequestId           string `json:"request_id,omitempty"`
	TableId             string `json:"table_id,omitempty"`
	UcResourceId        string `json:"uc_resource_id,omitempty"`
	UpdateId            string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *OriginPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st OriginPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PathPatternPb struct {
	Include string `json:"include,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PathPatternPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PathPatternPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineAccessControlRequestPb struct {
	GroupName            string                    `json:"group_name,omitempty"`
	PermissionLevel      PipelinePermissionLevelPb `json:"permission_level,omitempty"`
	ServicePrincipalName string                    `json:"service_principal_name,omitempty"`
	UserName             string                    `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineAccessControlResponsePb struct {
	AllPermissions       []PipelinePermissionPb `json:"all_permissions,omitempty"`
	DisplayName          string                 `json:"display_name,omitempty"`
	GroupName            string                 `json:"group_name,omitempty"`
	ServicePrincipalName string                 `json:"service_principal_name,omitempty"`
	UserName             string                 `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineClusterPb struct {
	ApplyPolicyDefaultValues  bool                         `json:"apply_policy_default_values,omitempty"`
	Autoscale                 *PipelineClusterAutoscalePb  `json:"autoscale,omitempty"`
	AwsAttributes             *computepb.AwsAttributesPb   `json:"aws_attributes,omitempty"`
	AzureAttributes           *computepb.AzureAttributesPb `json:"azure_attributes,omitempty"`
	ClusterLogConf            *computepb.ClusterLogConfPb  `json:"cluster_log_conf,omitempty"`
	CustomTags                map[string]string            `json:"custom_tags,omitempty"`
	DriverInstancePoolId      string                       `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId          string                       `json:"driver_node_type_id,omitempty"`
	EnableLocalDiskEncryption bool                         `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes             *computepb.GcpAttributesPb   `json:"gcp_attributes,omitempty"`
	InitScripts               []computepb.InitScriptInfoPb `json:"init_scripts,omitempty"`
	InstancePoolId            string                       `json:"instance_pool_id,omitempty"`
	Label                     string                       `json:"label,omitempty"`
	NodeTypeId                string                       `json:"node_type_id,omitempty"`
	NumWorkers                int                          `json:"num_workers,omitempty"`
	PolicyId                  string                       `json:"policy_id,omitempty"`
	SparkConf                 map[string]string            `json:"spark_conf,omitempty"`
	SparkEnvVars              map[string]string            `json:"spark_env_vars,omitempty"`
	SshPublicKeys             []string                     `json:"ssh_public_keys,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineClusterAutoscalePb struct {
	MaxWorkers int                            `json:"max_workers"`
	MinWorkers int                            `json:"min_workers"`
	Mode       PipelineClusterAutoscaleModePb `json:"mode,omitempty"`
}

type PipelineClusterAutoscaleModePb string

const PipelineClusterAutoscaleModeEnhanced PipelineClusterAutoscaleModePb = `ENHANCED`

const PipelineClusterAutoscaleModeLegacy PipelineClusterAutoscaleModePb = `LEGACY`

type PipelineDeploymentPb struct {
	Kind             DeploymentKindPb `json:"kind"`
	MetadataFilePath string           `json:"metadata_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineDeploymentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineDeploymentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineEventPb struct {
	Error         *ErrorDetailPb  `json:"error,omitempty"`
	EventType     string          `json:"event_type,omitempty"`
	Id            string          `json:"id,omitempty"`
	Level         EventLevelPb    `json:"level,omitempty"`
	MaturityLevel MaturityLevelPb `json:"maturity_level,omitempty"`
	Message       string          `json:"message,omitempty"`
	Origin        *OriginPb       `json:"origin,omitempty"`
	Sequence      *SequencingPb   `json:"sequence,omitempty"`
	Timestamp     string          `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineEventPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineEventPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineLibraryPb struct {
	File     *FileLibraryPb            `json:"file,omitempty"`
	Glob     *PathPatternPb            `json:"glob,omitempty"`
	Jar      string                    `json:"jar,omitempty"`
	Maven    *computepb.MavenLibraryPb `json:"maven,omitempty"`
	Notebook *NotebookLibraryPb        `json:"notebook,omitempty"`
	Whl      string                    `json:"whl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelinePermissionPb struct {
	Inherited           bool                      `json:"inherited,omitempty"`
	InheritedFromObject []string                  `json:"inherited_from_object,omitempty"`
	PermissionLevel     PipelinePermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelinePermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelinePermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelinePermissionLevelPb string

const PipelinePermissionLevelCanManage PipelinePermissionLevelPb = `CAN_MANAGE`

const PipelinePermissionLevelCanRun PipelinePermissionLevelPb = `CAN_RUN`

const PipelinePermissionLevelCanView PipelinePermissionLevelPb = `CAN_VIEW`

const PipelinePermissionLevelIsOwner PipelinePermissionLevelPb = `IS_OWNER`

type PipelinePermissionsPb struct {
	AccessControlList []PipelineAccessControlResponsePb `json:"access_control_list,omitempty"`
	ObjectId          string                            `json:"object_id,omitempty"`
	ObjectType        string                            `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelinePermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelinePermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelinePermissionsDescriptionPb struct {
	Description     string                    `json:"description,omitempty"`
	PermissionLevel PipelinePermissionLevelPb `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelinePermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelinePermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelinePermissionsRequestPb struct {
	AccessControlList []PipelineAccessControlRequestPb `json:"access_control_list,omitempty"`
	PipelineId        string                           `json:"-" url:"-"`
}

type PipelineSpecPb struct {
	BudgetPolicyId      string                                `json:"budget_policy_id,omitempty"`
	Catalog             string                                `json:"catalog,omitempty"`
	Channel             string                                `json:"channel,omitempty"`
	Clusters            []PipelineClusterPb                   `json:"clusters,omitempty"`
	Configuration       map[string]string                     `json:"configuration,omitempty"`
	Continuous          bool                                  `json:"continuous,omitempty"`
	Deployment          *PipelineDeploymentPb                 `json:"deployment,omitempty"`
	Development         bool                                  `json:"development,omitempty"`
	Edition             string                                `json:"edition,omitempty"`
	Environment         *PipelinesEnvironmentPb               `json:"environment,omitempty"`
	EventLog            *EventLogSpecPb                       `json:"event_log,omitempty"`
	Filters             *FiltersPb                            `json:"filters,omitempty"`
	GatewayDefinition   *IngestionGatewayPipelineDefinitionPb `json:"gateway_definition,omitempty"`
	Id                  string                                `json:"id,omitempty"`
	IngestionDefinition *IngestionPipelineDefinitionPb        `json:"ingestion_definition,omitempty"`
	Libraries           []PipelineLibraryPb                   `json:"libraries,omitempty"`
	Name                string                                `json:"name,omitempty"`
	Notifications       []NotificationsPb                     `json:"notifications,omitempty"`
	Photon              bool                                  `json:"photon,omitempty"`
	RestartWindow       *RestartWindowPb                      `json:"restart_window,omitempty"`
	RootPath            string                                `json:"root_path,omitempty"`
	Schema              string                                `json:"schema,omitempty"`
	Serverless          bool                                  `json:"serverless,omitempty"`
	Storage             string                                `json:"storage,omitempty"`
	Tags                map[string]string                     `json:"tags,omitempty"`
	Target              string                                `json:"target,omitempty"`
	Trigger             *PipelineTriggerPb                    `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineStatePb string

const PipelineStateDeleted PipelineStatePb = `DELETED`

const PipelineStateDeploying PipelineStatePb = `DEPLOYING`

const PipelineStateFailed PipelineStatePb = `FAILED`

const PipelineStateIdle PipelineStatePb = `IDLE`

const PipelineStateRecovering PipelineStatePb = `RECOVERING`

const PipelineStateResetting PipelineStatePb = `RESETTING`

const PipelineStateRunning PipelineStatePb = `RUNNING`

const PipelineStateStarting PipelineStatePb = `STARTING`

const PipelineStateStopping PipelineStatePb = `STOPPING`

type PipelineStateInfoPb struct {
	ClusterId       string                    `json:"cluster_id,omitempty"`
	CreatorUserName string                    `json:"creator_user_name,omitempty"`
	Health          PipelineStateInfoHealthPb `json:"health,omitempty"`
	LatestUpdates   []UpdateStateInfoPb       `json:"latest_updates,omitempty"`
	Name            string                    `json:"name,omitempty"`
	PipelineId      string                    `json:"pipeline_id,omitempty"`
	RunAsUserName   string                    `json:"run_as_user_name,omitempty"`
	State           PipelineStatePb           `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *PipelineStateInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st PipelineStateInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type PipelineStateInfoHealthPb string

const PipelineStateInfoHealthHealthy PipelineStateInfoHealthPb = `HEALTHY`

const PipelineStateInfoHealthUnhealthy PipelineStateInfoHealthPb = `UNHEALTHY`

type PipelineTriggerPb struct {
	Cron   *CronTriggerPb   `json:"cron,omitempty"`
	Manual *ManualTriggerPb `json:"manual,omitempty"`
}

type PipelinesEnvironmentPb struct {
	Dependencies []string `json:"dependencies,omitempty"`
}

type ReportSpecPb struct {
	DestinationCatalog string                 `json:"destination_catalog"`
	DestinationSchema  string                 `json:"destination_schema"`
	DestinationTable   string                 `json:"destination_table,omitempty"`
	SourceUrl          string                 `json:"source_url"`
	TableConfiguration *TableSpecificConfigPb `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ReportSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ReportSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RestartWindowPb struct {
	DaysOfWeek []DayOfWeekPb `json:"days_of_week,omitempty"`
	StartHour  int           `json:"start_hour"`
	TimeZoneId string        `json:"time_zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RestartWindowPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RestartWindowPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RunAsPb struct {
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	UserName             string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *RunAsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st RunAsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SchemaSpecPb struct {
	DestinationCatalog string                 `json:"destination_catalog"`
	DestinationSchema  string                 `json:"destination_schema"`
	SourceCatalog      string                 `json:"source_catalog,omitempty"`
	SourceSchema       string                 `json:"source_schema"`
	TableConfiguration *TableSpecificConfigPb `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SchemaSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SchemaSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SequencingPb struct {
	ControlPlaneSeqNo int64          `json:"control_plane_seq_no,omitempty"`
	DataPlaneId       *DataPlaneIdPb `json:"data_plane_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SequencingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SequencingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SerializedExceptionPb struct {
	ClassName string         `json:"class_name,omitempty"`
	Message   string         `json:"message,omitempty"`
	Stack     []StackFramePb `json:"stack,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *SerializedExceptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st SerializedExceptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StackFramePb struct {
	DeclaringClass string `json:"declaring_class,omitempty"`
	FileName       string `json:"file_name,omitempty"`
	LineNumber     int    `json:"line_number,omitempty"`
	MethodName     string `json:"method_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *StackFramePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st StackFramePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StartUpdatePb struct {
	Cause                StartUpdateCausePb `json:"cause,omitempty"`
	FullRefresh          bool               `json:"full_refresh,omitempty"`
	FullRefreshSelection []string           `json:"full_refresh_selection,omitempty"`
	PipelineId           string             `json:"-" url:"-"`
	RefreshSelection     []string           `json:"refresh_selection,omitempty"`
	ValidateOnly         bool               `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *StartUpdatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st StartUpdatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StartUpdateCausePb string

const StartUpdateCauseApiCall StartUpdateCausePb = `API_CALL`

const StartUpdateCauseInfrastructureMaintenance StartUpdateCausePb = `INFRASTRUCTURE_MAINTENANCE`

const StartUpdateCauseJobTask StartUpdateCausePb = `JOB_TASK`

const StartUpdateCauseRetryOnFailure StartUpdateCausePb = `RETRY_ON_FAILURE`

const StartUpdateCauseSchemaChange StartUpdateCausePb = `SCHEMA_CHANGE`

const StartUpdateCauseServiceUpgrade StartUpdateCausePb = `SERVICE_UPGRADE`

const StartUpdateCauseUserAction StartUpdateCausePb = `USER_ACTION`

type StartUpdateResponsePb struct {
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *StartUpdateResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st StartUpdateResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type StopPipelineResponsePb struct {
}

type StopRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

type TableSpecPb struct {
	DestinationCatalog string                 `json:"destination_catalog"`
	DestinationSchema  string                 `json:"destination_schema"`
	DestinationTable   string                 `json:"destination_table,omitempty"`
	SourceCatalog      string                 `json:"source_catalog,omitempty"`
	SourceSchema       string                 `json:"source_schema,omitempty"`
	SourceTable        string                 `json:"source_table"`
	TableConfiguration *TableSpecificConfigPb `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TableSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TableSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableSpecificConfigPb struct {
	ExcludeColumns                 []string                                                                   `json:"exclude_columns,omitempty"`
	IncludeColumns                 []string                                                                   `json:"include_columns,omitempty"`
	PrimaryKeys                    []string                                                                   `json:"primary_keys,omitempty"`
	QueryBasedConnectorConfig      *IngestionPipelineDefinitionTableSpecificConfigQueryBasedConnectorConfigPb `json:"query_based_connector_config,omitempty"`
	SalesforceIncludeFormulaFields bool                                                                       `json:"salesforce_include_formula_fields,omitempty"`
	ScdType                        TableSpecificConfigScdTypePb                                               `json:"scd_type,omitempty"`
	SequenceBy                     []string                                                                   `json:"sequence_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *TableSpecificConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st TableSpecificConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TableSpecificConfigScdTypePb string

const TableSpecificConfigScdTypeAppendOnly TableSpecificConfigScdTypePb = `APPEND_ONLY`

const TableSpecificConfigScdTypeScdType1 TableSpecificConfigScdTypePb = `SCD_TYPE_1`

const TableSpecificConfigScdTypeScdType2 TableSpecificConfigScdTypePb = `SCD_TYPE_2`

type UpdateInfoPb struct {
	Cause                UpdateInfoCausePb `json:"cause,omitempty"`
	ClusterId            string            `json:"cluster_id,omitempty"`
	Config               *PipelineSpecPb   `json:"config,omitempty"`
	CreationTime         int64             `json:"creation_time,omitempty"`
	FullRefresh          bool              `json:"full_refresh,omitempty"`
	FullRefreshSelection []string          `json:"full_refresh_selection,omitempty"`
	PipelineId           string            `json:"pipeline_id,omitempty"`
	RefreshSelection     []string          `json:"refresh_selection,omitempty"`
	State                UpdateInfoStatePb `json:"state,omitempty"`
	UpdateId             string            `json:"update_id,omitempty"`
	ValidateOnly         bool              `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateInfoCausePb string

const UpdateInfoCauseApiCall UpdateInfoCausePb = `API_CALL`

const UpdateInfoCauseInfrastructureMaintenance UpdateInfoCausePb = `INFRASTRUCTURE_MAINTENANCE`

const UpdateInfoCauseJobTask UpdateInfoCausePb = `JOB_TASK`

const UpdateInfoCauseRetryOnFailure UpdateInfoCausePb = `RETRY_ON_FAILURE`

const UpdateInfoCauseSchemaChange UpdateInfoCausePb = `SCHEMA_CHANGE`

const UpdateInfoCauseServiceUpgrade UpdateInfoCausePb = `SERVICE_UPGRADE`

const UpdateInfoCauseUserAction UpdateInfoCausePb = `USER_ACTION`

type UpdateInfoStatePb string

const UpdateInfoStateCanceled UpdateInfoStatePb = `CANCELED`

const UpdateInfoStateCompleted UpdateInfoStatePb = `COMPLETED`

const UpdateInfoStateCreated UpdateInfoStatePb = `CREATED`

const UpdateInfoStateFailed UpdateInfoStatePb = `FAILED`

const UpdateInfoStateInitializing UpdateInfoStatePb = `INITIALIZING`

const UpdateInfoStateQueued UpdateInfoStatePb = `QUEUED`

const UpdateInfoStateResetting UpdateInfoStatePb = `RESETTING`

const UpdateInfoStateRunning UpdateInfoStatePb = `RUNNING`

const UpdateInfoStateSettingUpTables UpdateInfoStatePb = `SETTING_UP_TABLES`

const UpdateInfoStateStopping UpdateInfoStatePb = `STOPPING`

const UpdateInfoStateWaitingForResources UpdateInfoStatePb = `WAITING_FOR_RESOURCES`

type UpdateStateInfoPb struct {
	CreationTime string                 `json:"creation_time,omitempty"`
	State        UpdateStateInfoStatePb `json:"state,omitempty"`
	UpdateId     string                 `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *UpdateStateInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st UpdateStateInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateStateInfoStatePb string

const UpdateStateInfoStateCanceled UpdateStateInfoStatePb = `CANCELED`

const UpdateStateInfoStateCompleted UpdateStateInfoStatePb = `COMPLETED`

const UpdateStateInfoStateCreated UpdateStateInfoStatePb = `CREATED`

const UpdateStateInfoStateFailed UpdateStateInfoStatePb = `FAILED`

const UpdateStateInfoStateInitializing UpdateStateInfoStatePb = `INITIALIZING`

const UpdateStateInfoStateQueued UpdateStateInfoStatePb = `QUEUED`

const UpdateStateInfoStateResetting UpdateStateInfoStatePb = `RESETTING`

const UpdateStateInfoStateRunning UpdateStateInfoStatePb = `RUNNING`

const UpdateStateInfoStateSettingUpTables UpdateStateInfoStatePb = `SETTING_UP_TABLES`

const UpdateStateInfoStateStopping UpdateStateInfoStatePb = `STOPPING`

const UpdateStateInfoStateWaitingForResources UpdateStateInfoStatePb = `WAITING_FOR_RESOURCES`
