// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package pipelines

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
	"github.com/databricks/databricks-sdk-go/service/compute"
)

func createPipelineToPb(st *CreatePipeline) (*createPipelinePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPipelinePb{}
	pb.AllowDuplicateNames = st.AllowDuplicateNames
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.Catalog = st.Catalog
	pb.Channel = st.Channel
	pb.Clusters = st.Clusters
	pb.Configuration = st.Configuration
	pb.Continuous = st.Continuous
	pb.Deployment = st.Deployment
	pb.Development = st.Development
	pb.DryRun = st.DryRun
	pb.Edition = st.Edition
	pb.Environment = st.Environment
	pb.EventLog = st.EventLog
	pb.Filters = st.Filters
	pb.GatewayDefinition = st.GatewayDefinition
	pb.Id = st.Id
	pb.IngestionDefinition = st.IngestionDefinition
	pb.Libraries = st.Libraries
	pb.Name = st.Name
	pb.Notifications = st.Notifications
	pb.Photon = st.Photon
	pb.RestartWindow = st.RestartWindow
	pb.RootPath = st.RootPath
	pb.RunAs = st.RunAs
	pb.Schema = st.Schema
	pb.Serverless = st.Serverless
	pb.Storage = st.Storage
	pb.Tags = st.Tags
	pb.Target = st.Target
	pb.Trigger = st.Trigger

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPipelinePb struct {
	AllowDuplicateNames bool                                `json:"allow_duplicate_names,omitempty"`
	BudgetPolicyId      string                              `json:"budget_policy_id,omitempty"`
	Catalog             string                              `json:"catalog,omitempty"`
	Channel             string                              `json:"channel,omitempty"`
	Clusters            []PipelineCluster                   `json:"clusters,omitempty"`
	Configuration       map[string]string                   `json:"configuration,omitempty"`
	Continuous          bool                                `json:"continuous,omitempty"`
	Deployment          *PipelineDeployment                 `json:"deployment,omitempty"`
	Development         bool                                `json:"development,omitempty"`
	DryRun              bool                                `json:"dry_run,omitempty"`
	Edition             string                              `json:"edition,omitempty"`
	Environment         *PipelinesEnvironment               `json:"environment,omitempty"`
	EventLog            *EventLogSpec                       `json:"event_log,omitempty"`
	Filters             *Filters                            `json:"filters,omitempty"`
	GatewayDefinition   *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	Id                  string                              `json:"id,omitempty"`
	IngestionDefinition *IngestionPipelineDefinition        `json:"ingestion_definition,omitempty"`
	Libraries           []PipelineLibrary                   `json:"libraries,omitempty"`
	Name                string                              `json:"name,omitempty"`
	Notifications       []Notifications                     `json:"notifications,omitempty"`
	Photon              bool                                `json:"photon,omitempty"`
	RestartWindow       *RestartWindow                      `json:"restart_window,omitempty"`
	RootPath            string                              `json:"root_path,omitempty"`
	RunAs               *RunAs                              `json:"run_as,omitempty"`
	Schema              string                              `json:"schema,omitempty"`
	Serverless          bool                                `json:"serverless,omitempty"`
	Storage             string                              `json:"storage,omitempty"`
	Tags                map[string]string                   `json:"tags,omitempty"`
	Target              string                              `json:"target,omitempty"`
	Trigger             *PipelineTrigger                    `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPipelineFromPb(pb *createPipelinePb) (*CreatePipeline, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePipeline{}
	st.AllowDuplicateNames = pb.AllowDuplicateNames
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Catalog = pb.Catalog
	st.Channel = pb.Channel
	st.Clusters = pb.Clusters
	st.Configuration = pb.Configuration
	st.Continuous = pb.Continuous
	st.Deployment = pb.Deployment
	st.Development = pb.Development
	st.DryRun = pb.DryRun
	st.Edition = pb.Edition
	st.Environment = pb.Environment
	st.EventLog = pb.EventLog
	st.Filters = pb.Filters
	st.GatewayDefinition = pb.GatewayDefinition
	st.Id = pb.Id
	st.IngestionDefinition = pb.IngestionDefinition
	st.Libraries = pb.Libraries
	st.Name = pb.Name
	st.Notifications = pb.Notifications
	st.Photon = pb.Photon
	st.RestartWindow = pb.RestartWindow
	st.RootPath = pb.RootPath
	st.RunAs = pb.RunAs
	st.Schema = pb.Schema
	st.Serverless = pb.Serverless
	st.Storage = pb.Storage
	st.Tags = pb.Tags
	st.Target = pb.Target
	st.Trigger = pb.Trigger

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPipelinePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPipelinePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createPipelineResponseToPb(st *CreatePipelineResponse) (*createPipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createPipelineResponsePb{}
	pb.EffectiveSettings = st.EffectiveSettings
	pb.PipelineId = st.PipelineId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type createPipelineResponsePb struct {
	EffectiveSettings *PipelineSpec `json:"effective_settings,omitempty"`
	PipelineId        string        `json:"pipeline_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createPipelineResponseFromPb(pb *createPipelineResponsePb) (*CreatePipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreatePipelineResponse{}
	st.EffectiveSettings = pb.EffectiveSettings
	st.PipelineId = pb.PipelineId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createPipelineResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createPipelineResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func cronTriggerToPb(st *CronTrigger) (*cronTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &cronTriggerPb{}
	pb.QuartzCronSchedule = st.QuartzCronSchedule
	pb.TimezoneId = st.TimezoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type cronTriggerPb struct {
	QuartzCronSchedule string `json:"quartz_cron_schedule,omitempty"`
	TimezoneId         string `json:"timezone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func cronTriggerFromPb(pb *cronTriggerPb) (*CronTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CronTrigger{}
	st.QuartzCronSchedule = pb.QuartzCronSchedule
	st.TimezoneId = pb.TimezoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *cronTriggerPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st cronTriggerPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func dataPlaneIdToPb(st *DataPlaneId) (*dataPlaneIdPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &dataPlaneIdPb{}
	pb.Instance = st.Instance
	pb.SeqNo = st.SeqNo

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type dataPlaneIdPb struct {
	Instance string `json:"instance,omitempty"`
	SeqNo    int64  `json:"seq_no,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func dataPlaneIdFromPb(pb *dataPlaneIdPb) (*DataPlaneId, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DataPlaneId{}
	st.Instance = pb.Instance
	st.SeqNo = pb.SeqNo

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *dataPlaneIdPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st dataPlaneIdPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func deletePipelineRequestToPb(st *DeletePipelineRequest) (*deletePipelineRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePipelineRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
}

type deletePipelineRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

func deletePipelineRequestFromPb(pb *deletePipelineRequestPb) (*DeletePipelineRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePipelineRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

func deletePipelineResponseToPb(st *DeletePipelineResponse) (*deletePipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deletePipelineResponsePb{}

	return pb, nil
}

type deletePipelineResponsePb struct {
}

func deletePipelineResponseFromPb(pb *deletePipelineResponsePb) (*DeletePipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeletePipelineResponse{}

	return st, nil
}

func editPipelineToPb(st *EditPipeline) (*editPipelinePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editPipelinePb{}
	pb.AllowDuplicateNames = st.AllowDuplicateNames
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.Catalog = st.Catalog
	pb.Channel = st.Channel
	pb.Clusters = st.Clusters
	pb.Configuration = st.Configuration
	pb.Continuous = st.Continuous
	pb.Deployment = st.Deployment
	pb.Development = st.Development
	pb.Edition = st.Edition
	pb.Environment = st.Environment
	pb.EventLog = st.EventLog
	pb.ExpectedLastModified = st.ExpectedLastModified
	pb.Filters = st.Filters
	pb.GatewayDefinition = st.GatewayDefinition
	pb.Id = st.Id
	pb.IngestionDefinition = st.IngestionDefinition
	pb.Libraries = st.Libraries
	pb.Name = st.Name
	pb.Notifications = st.Notifications
	pb.Photon = st.Photon
	pb.PipelineId = st.PipelineId
	pb.RestartWindow = st.RestartWindow
	pb.RootPath = st.RootPath
	pb.RunAs = st.RunAs
	pb.Schema = st.Schema
	pb.Serverless = st.Serverless
	pb.Storage = st.Storage
	pb.Tags = st.Tags
	pb.Target = st.Target
	pb.Trigger = st.Trigger

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type editPipelinePb struct {
	AllowDuplicateNames  bool                                `json:"allow_duplicate_names,omitempty"`
	BudgetPolicyId       string                              `json:"budget_policy_id,omitempty"`
	Catalog              string                              `json:"catalog,omitempty"`
	Channel              string                              `json:"channel,omitempty"`
	Clusters             []PipelineCluster                   `json:"clusters,omitempty"`
	Configuration        map[string]string                   `json:"configuration,omitempty"`
	Continuous           bool                                `json:"continuous,omitempty"`
	Deployment           *PipelineDeployment                 `json:"deployment,omitempty"`
	Development          bool                                `json:"development,omitempty"`
	Edition              string                              `json:"edition,omitempty"`
	Environment          *PipelinesEnvironment               `json:"environment,omitempty"`
	EventLog             *EventLogSpec                       `json:"event_log,omitempty"`
	ExpectedLastModified int64                               `json:"expected_last_modified,omitempty"`
	Filters              *Filters                            `json:"filters,omitempty"`
	GatewayDefinition    *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	Id                   string                              `json:"id,omitempty"`
	IngestionDefinition  *IngestionPipelineDefinition        `json:"ingestion_definition,omitempty"`
	Libraries            []PipelineLibrary                   `json:"libraries,omitempty"`
	Name                 string                              `json:"name,omitempty"`
	Notifications        []Notifications                     `json:"notifications,omitempty"`
	Photon               bool                                `json:"photon,omitempty"`
	PipelineId           string                              `json:"-" url:"-"`
	RestartWindow        *RestartWindow                      `json:"restart_window,omitempty"`
	RootPath             string                              `json:"root_path,omitempty"`
	RunAs                *RunAs                              `json:"run_as,omitempty"`
	Schema               string                              `json:"schema,omitempty"`
	Serverless           bool                                `json:"serverless,omitempty"`
	Storage              string                              `json:"storage,omitempty"`
	Tags                 map[string]string                   `json:"tags,omitempty"`
	Target               string                              `json:"target,omitempty"`
	Trigger              *PipelineTrigger                    `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func editPipelineFromPb(pb *editPipelinePb) (*EditPipeline, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditPipeline{}
	st.AllowDuplicateNames = pb.AllowDuplicateNames
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Catalog = pb.Catalog
	st.Channel = pb.Channel
	st.Clusters = pb.Clusters
	st.Configuration = pb.Configuration
	st.Continuous = pb.Continuous
	st.Deployment = pb.Deployment
	st.Development = pb.Development
	st.Edition = pb.Edition
	st.Environment = pb.Environment
	st.EventLog = pb.EventLog
	st.ExpectedLastModified = pb.ExpectedLastModified
	st.Filters = pb.Filters
	st.GatewayDefinition = pb.GatewayDefinition
	st.Id = pb.Id
	st.IngestionDefinition = pb.IngestionDefinition
	st.Libraries = pb.Libraries
	st.Name = pb.Name
	st.Notifications = pb.Notifications
	st.Photon = pb.Photon
	st.PipelineId = pb.PipelineId
	st.RestartWindow = pb.RestartWindow
	st.RootPath = pb.RootPath
	st.RunAs = pb.RunAs
	st.Schema = pb.Schema
	st.Serverless = pb.Serverless
	st.Storage = pb.Storage
	st.Tags = pb.Tags
	st.Target = pb.Target
	st.Trigger = pb.Trigger

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *editPipelinePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st editPipelinePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func editPipelineResponseToPb(st *EditPipelineResponse) (*editPipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &editPipelineResponsePb{}

	return pb, nil
}

type editPipelineResponsePb struct {
}

func editPipelineResponseFromPb(pb *editPipelineResponsePb) (*EditPipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EditPipelineResponse{}

	return st, nil
}

func errorDetailToPb(st *ErrorDetail) (*errorDetailPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &errorDetailPb{}
	pb.Exceptions = st.Exceptions
	pb.Fatal = st.Fatal

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type errorDetailPb struct {
	Exceptions []SerializedException `json:"exceptions,omitempty"`
	Fatal      bool                  `json:"fatal,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func errorDetailFromPb(pb *errorDetailPb) (*ErrorDetail, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ErrorDetail{}
	st.Exceptions = pb.Exceptions
	st.Fatal = pb.Fatal

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *errorDetailPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st errorDetailPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func eventLogSpecToPb(st *EventLogSpec) (*eventLogSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &eventLogSpecPb{}
	pb.Catalog = st.Catalog
	pb.Name = st.Name
	pb.Schema = st.Schema

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type eventLogSpecPb struct {
	Catalog string `json:"catalog,omitempty"`
	Name    string `json:"name,omitempty"`
	Schema  string `json:"schema,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func eventLogSpecFromPb(pb *eventLogSpecPb) (*EventLogSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &EventLogSpec{}
	st.Catalog = pb.Catalog
	st.Name = pb.Name
	st.Schema = pb.Schema

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *eventLogSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st eventLogSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func fileLibraryToPb(st *FileLibrary) (*fileLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileLibraryPb{}
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type fileLibraryPb struct {
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileLibraryFromPb(pb *fileLibraryPb) (*FileLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileLibrary{}
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *fileLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st fileLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func filtersToPb(st *Filters) (*filtersPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &filtersPb{}
	pb.Exclude = st.Exclude
	pb.Include = st.Include

	return pb, nil
}

type filtersPb struct {
	Exclude []string `json:"exclude,omitempty"`
	Include []string `json:"include,omitempty"`
}

func filtersFromPb(pb *filtersPb) (*Filters, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Filters{}
	st.Exclude = pb.Exclude
	st.Include = pb.Include

	return st, nil
}

func getPipelinePermissionLevelsRequestToPb(st *GetPipelinePermissionLevelsRequest) (*getPipelinePermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelinePermissionLevelsRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
}

type getPipelinePermissionLevelsRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

func getPipelinePermissionLevelsRequestFromPb(pb *getPipelinePermissionLevelsRequestPb) (*GetPipelinePermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelinePermissionLevelsRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

func getPipelinePermissionLevelsResponseToPb(st *GetPipelinePermissionLevelsResponse) (*getPipelinePermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelinePermissionLevelsResponsePb{}
	pb.PermissionLevels = st.PermissionLevels

	return pb, nil
}

type getPipelinePermissionLevelsResponsePb struct {
	PermissionLevels []PipelinePermissionsDescription `json:"permission_levels,omitempty"`
}

func getPipelinePermissionLevelsResponseFromPb(pb *getPipelinePermissionLevelsResponsePb) (*GetPipelinePermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelinePermissionLevelsResponse{}
	st.PermissionLevels = pb.PermissionLevels

	return st, nil
}

func getPipelinePermissionsRequestToPb(st *GetPipelinePermissionsRequest) (*getPipelinePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelinePermissionsRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
}

type getPipelinePermissionsRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

func getPipelinePermissionsRequestFromPb(pb *getPipelinePermissionsRequestPb) (*GetPipelinePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelinePermissionsRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

func getPipelineRequestToPb(st *GetPipelineRequest) (*getPipelineRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelineRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
}

type getPipelineRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

func getPipelineRequestFromPb(pb *getPipelineRequestPb) (*GetPipelineRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelineRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

func getPipelineResponseToPb(st *GetPipelineResponse) (*getPipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getPipelineResponsePb{}
	pb.Cause = st.Cause
	pb.ClusterId = st.ClusterId
	pb.CreatorUserName = st.CreatorUserName
	pb.EffectiveBudgetPolicyId = st.EffectiveBudgetPolicyId
	pb.Health = st.Health
	pb.LastModified = st.LastModified
	pb.LatestUpdates = st.LatestUpdates
	pb.Name = st.Name
	pb.PipelineId = st.PipelineId
	pb.RunAsUserName = st.RunAsUserName
	pb.Spec = st.Spec
	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type getPipelineResponsePb struct {
	Cause                   string                    `json:"cause,omitempty"`
	ClusterId               string                    `json:"cluster_id,omitempty"`
	CreatorUserName         string                    `json:"creator_user_name,omitempty"`
	EffectiveBudgetPolicyId string                    `json:"effective_budget_policy_id,omitempty"`
	Health                  GetPipelineResponseHealth `json:"health,omitempty"`
	LastModified            int64                     `json:"last_modified,omitempty"`
	LatestUpdates           []UpdateStateInfo         `json:"latest_updates,omitempty"`
	Name                    string                    `json:"name,omitempty"`
	PipelineId              string                    `json:"pipeline_id,omitempty"`
	RunAsUserName           string                    `json:"run_as_user_name,omitempty"`
	Spec                    *PipelineSpec             `json:"spec,omitempty"`
	State                   PipelineState             `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getPipelineResponseFromPb(pb *getPipelineResponsePb) (*GetPipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetPipelineResponse{}
	st.Cause = pb.Cause
	st.ClusterId = pb.ClusterId
	st.CreatorUserName = pb.CreatorUserName
	st.EffectiveBudgetPolicyId = pb.EffectiveBudgetPolicyId
	st.Health = pb.Health
	st.LastModified = pb.LastModified
	st.LatestUpdates = pb.LatestUpdates
	st.Name = pb.Name
	st.PipelineId = pb.PipelineId
	st.RunAsUserName = pb.RunAsUserName
	st.Spec = pb.Spec
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getPipelineResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getPipelineResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func getUpdateRequestToPb(st *GetUpdateRequest) (*getUpdateRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getUpdateRequestPb{}
	pb.PipelineId = st.PipelineId
	pb.UpdateId = st.UpdateId

	return pb, nil
}

type getUpdateRequestPb struct {
	PipelineId string `json:"-" url:"-"`
	UpdateId   string `json:"-" url:"-"`
}

func getUpdateRequestFromPb(pb *getUpdateRequestPb) (*GetUpdateRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetUpdateRequest{}
	st.PipelineId = pb.PipelineId
	st.UpdateId = pb.UpdateId

	return st, nil
}

func getUpdateResponseToPb(st *GetUpdateResponse) (*getUpdateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getUpdateResponsePb{}
	pb.Update = st.Update

	return pb, nil
}

type getUpdateResponsePb struct {
	Update *UpdateInfo `json:"update,omitempty"`
}

func getUpdateResponseFromPb(pb *getUpdateResponsePb) (*GetUpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetUpdateResponse{}
	st.Update = pb.Update

	return st, nil
}

func ingestionConfigToPb(st *IngestionConfig) (*ingestionConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ingestionConfigPb{}
	pb.Report = st.Report
	pb.Schema = st.Schema
	pb.Table = st.Table

	return pb, nil
}

type ingestionConfigPb struct {
	Report *ReportSpec `json:"report,omitempty"`
	Schema *SchemaSpec `json:"schema,omitempty"`
	Table  *TableSpec  `json:"table,omitempty"`
}

func ingestionConfigFromPb(pb *ingestionConfigPb) (*IngestionConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IngestionConfig{}
	st.Report = pb.Report
	st.Schema = pb.Schema
	st.Table = pb.Table

	return st, nil
}

func ingestionGatewayPipelineDefinitionToPb(st *IngestionGatewayPipelineDefinition) (*ingestionGatewayPipelineDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ingestionGatewayPipelineDefinitionPb{}
	pb.ConnectionId = st.ConnectionId
	pb.ConnectionName = st.ConnectionName
	pb.GatewayStorageCatalog = st.GatewayStorageCatalog
	pb.GatewayStorageName = st.GatewayStorageName
	pb.GatewayStorageSchema = st.GatewayStorageSchema

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type ingestionGatewayPipelineDefinitionPb struct {
	ConnectionId          string `json:"connection_id,omitempty"`
	ConnectionName        string `json:"connection_name"`
	GatewayStorageCatalog string `json:"gateway_storage_catalog"`
	GatewayStorageName    string `json:"gateway_storage_name,omitempty"`
	GatewayStorageSchema  string `json:"gateway_storage_schema"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ingestionGatewayPipelineDefinitionFromPb(pb *ingestionGatewayPipelineDefinitionPb) (*IngestionGatewayPipelineDefinition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IngestionGatewayPipelineDefinition{}
	st.ConnectionId = pb.ConnectionId
	st.ConnectionName = pb.ConnectionName
	st.GatewayStorageCatalog = pb.GatewayStorageCatalog
	st.GatewayStorageName = pb.GatewayStorageName
	st.GatewayStorageSchema = pb.GatewayStorageSchema

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ingestionGatewayPipelineDefinitionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ingestionGatewayPipelineDefinitionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func ingestionPipelineDefinitionToPb(st *IngestionPipelineDefinition) (*ingestionPipelineDefinitionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &ingestionPipelineDefinitionPb{}
	pb.ConnectionName = st.ConnectionName
	pb.IngestionGatewayId = st.IngestionGatewayId
	pb.Objects = st.Objects
	pb.SourceType = st.SourceType
	pb.TableConfiguration = st.TableConfiguration

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type ingestionPipelineDefinitionPb struct {
	ConnectionName     string               `json:"connection_name,omitempty"`
	IngestionGatewayId string               `json:"ingestion_gateway_id,omitempty"`
	Objects            []IngestionConfig    `json:"objects,omitempty"`
	SourceType         IngestionSourceType  `json:"source_type,omitempty"`
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func ingestionPipelineDefinitionFromPb(pb *ingestionPipelineDefinitionPb) (*IngestionPipelineDefinition, error) {
	if pb == nil {
		return nil, nil
	}
	st := &IngestionPipelineDefinition{}
	st.ConnectionName = pb.ConnectionName
	st.IngestionGatewayId = pb.IngestionGatewayId
	st.Objects = pb.Objects
	st.SourceType = pb.SourceType
	st.TableConfiguration = pb.TableConfiguration

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *ingestionPipelineDefinitionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ingestionPipelineDefinitionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listPipelineEventsRequestToPb(st *ListPipelineEventsRequest) (*listPipelineEventsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelineEventsRequestPb{}
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken
	pb.PipelineId = st.PipelineId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listPipelineEventsRequestPb struct {
	Filter     string   `json:"-" url:"filter,omitempty"`
	MaxResults int      `json:"-" url:"max_results,omitempty"`
	OrderBy    []string `json:"-" url:"order_by,omitempty"`
	PageToken  string   `json:"-" url:"page_token,omitempty"`
	PipelineId string   `json:"-" url:"-"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPipelineEventsRequestFromPb(pb *listPipelineEventsRequestPb) (*ListPipelineEventsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelineEventsRequest{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken
	st.PipelineId = pb.PipelineId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPipelineEventsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPipelineEventsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listPipelineEventsResponseToPb(st *ListPipelineEventsResponse) (*listPipelineEventsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelineEventsResponsePb{}
	pb.Events = st.Events
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listPipelineEventsResponsePb struct {
	Events        []PipelineEvent `json:"events,omitempty"`
	NextPageToken string          `json:"next_page_token,omitempty"`
	PrevPageToken string          `json:"prev_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPipelineEventsResponseFromPb(pb *listPipelineEventsResponsePb) (*ListPipelineEventsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelineEventsResponse{}
	st.Events = pb.Events
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPipelineEventsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPipelineEventsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listPipelinesRequestToPb(st *ListPipelinesRequest) (*listPipelinesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelinesRequestPb{}
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listPipelinesRequestPb struct {
	Filter     string   `json:"-" url:"filter,omitempty"`
	MaxResults int      `json:"-" url:"max_results,omitempty"`
	OrderBy    []string `json:"-" url:"order_by,omitempty"`
	PageToken  string   `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPipelinesRequestFromPb(pb *listPipelinesRequestPb) (*ListPipelinesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelinesRequest{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPipelinesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPipelinesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listPipelinesResponseToPb(st *ListPipelinesResponse) (*listPipelinesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listPipelinesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.Statuses = st.Statuses

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listPipelinesResponsePb struct {
	NextPageToken string              `json:"next_page_token,omitempty"`
	Statuses      []PipelineStateInfo `json:"statuses,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listPipelinesResponseFromPb(pb *listPipelinesResponsePb) (*ListPipelinesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListPipelinesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.Statuses = pb.Statuses

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listPipelinesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listPipelinesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listUpdatesRequestToPb(st *ListUpdatesRequest) (*listUpdatesRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listUpdatesRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	pb.PipelineId = st.PipelineId
	pb.UntilUpdateId = st.UntilUpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listUpdatesRequestPb struct {
	MaxResults    int    `json:"-" url:"max_results,omitempty"`
	PageToken     string `json:"-" url:"page_token,omitempty"`
	PipelineId    string `json:"-" url:"-"`
	UntilUpdateId string `json:"-" url:"until_update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listUpdatesRequestFromPb(pb *listUpdatesRequestPb) (*ListUpdatesRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListUpdatesRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.PipelineId = pb.PipelineId
	st.UntilUpdateId = pb.UntilUpdateId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listUpdatesRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listUpdatesRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listUpdatesResponseToPb(st *ListUpdatesResponse) (*listUpdatesResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listUpdatesResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.PrevPageToken = st.PrevPageToken
	pb.Updates = st.Updates

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listUpdatesResponsePb struct {
	NextPageToken string       `json:"next_page_token,omitempty"`
	PrevPageToken string       `json:"prev_page_token,omitempty"`
	Updates       []UpdateInfo `json:"updates,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listUpdatesResponseFromPb(pb *listUpdatesResponsePb) (*ListUpdatesResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListUpdatesResponse{}
	st.NextPageToken = pb.NextPageToken
	st.PrevPageToken = pb.PrevPageToken
	st.Updates = pb.Updates

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listUpdatesResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listUpdatesResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func manualTriggerToPb(st *ManualTrigger) (*manualTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &manualTriggerPb{}

	return pb, nil
}

type manualTriggerPb struct {
}

func manualTriggerFromPb(pb *manualTriggerPb) (*ManualTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ManualTrigger{}

	return st, nil
}

func notebookLibraryToPb(st *NotebookLibrary) (*notebookLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notebookLibraryPb{}
	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type notebookLibraryPb struct {
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func notebookLibraryFromPb(pb *notebookLibraryPb) (*NotebookLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &NotebookLibrary{}
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *notebookLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st notebookLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func notificationsToPb(st *Notifications) (*notificationsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &notificationsPb{}
	pb.Alerts = st.Alerts
	pb.EmailRecipients = st.EmailRecipients

	return pb, nil
}

type notificationsPb struct {
	Alerts          []string `json:"alerts,omitempty"`
	EmailRecipients []string `json:"email_recipients,omitempty"`
}

func notificationsFromPb(pb *notificationsPb) (*Notifications, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Notifications{}
	st.Alerts = pb.Alerts
	st.EmailRecipients = pb.EmailRecipients

	return st, nil
}

func originToPb(st *Origin) (*originPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &originPb{}
	pb.BatchId = st.BatchId
	pb.Cloud = st.Cloud
	pb.ClusterId = st.ClusterId
	pb.DatasetName = st.DatasetName
	pb.FlowId = st.FlowId
	pb.FlowName = st.FlowName
	pb.Host = st.Host
	pb.MaintenanceId = st.MaintenanceId
	pb.MaterializationName = st.MaterializationName
	pb.OrgId = st.OrgId
	pb.PipelineId = st.PipelineId
	pb.PipelineName = st.PipelineName
	pb.Region = st.Region
	pb.RequestId = st.RequestId
	pb.TableId = st.TableId
	pb.UcResourceId = st.UcResourceId
	pb.UpdateId = st.UpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type originPb struct {
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

func originFromPb(pb *originPb) (*Origin, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Origin{}
	st.BatchId = pb.BatchId
	st.Cloud = pb.Cloud
	st.ClusterId = pb.ClusterId
	st.DatasetName = pb.DatasetName
	st.FlowId = pb.FlowId
	st.FlowName = pb.FlowName
	st.Host = pb.Host
	st.MaintenanceId = pb.MaintenanceId
	st.MaterializationName = pb.MaterializationName
	st.OrgId = pb.OrgId
	st.PipelineId = pb.PipelineId
	st.PipelineName = pb.PipelineName
	st.Region = pb.Region
	st.RequestId = pb.RequestId
	st.TableId = pb.TableId
	st.UcResourceId = pb.UcResourceId
	st.UpdateId = pb.UpdateId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *originPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st originPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pathPatternToPb(st *PathPattern) (*pathPatternPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pathPatternPb{}
	pb.Include = st.Include

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pathPatternPb struct {
	Include string `json:"include,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pathPatternFromPb(pb *pathPatternPb) (*PathPattern, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PathPattern{}
	st.Include = pb.Include

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pathPatternPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pathPatternPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelineAccessControlRequestToPb(st *PipelineAccessControlRequest) (*pipelineAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	pb.PermissionLevel = st.PermissionLevel
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineAccessControlRequestPb struct {
	GroupName            string                  `json:"group_name,omitempty"`
	PermissionLevel      PipelinePermissionLevel `json:"permission_level,omitempty"`
	ServicePrincipalName string                  `json:"service_principal_name,omitempty"`
	UserName             string                  `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineAccessControlRequestFromPb(pb *pipelineAccessControlRequestPb) (*PipelineAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelineAccessControlResponseToPb(st *PipelineAccessControlResponse) (*pipelineAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineAccessControlResponsePb{}
	pb.AllPermissions = st.AllPermissions
	pb.DisplayName = st.DisplayName
	pb.GroupName = st.GroupName
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineAccessControlResponsePb struct {
	AllPermissions       []PipelinePermission `json:"all_permissions,omitempty"`
	DisplayName          string               `json:"display_name,omitempty"`
	GroupName            string               `json:"group_name,omitempty"`
	ServicePrincipalName string               `json:"service_principal_name,omitempty"`
	UserName             string               `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineAccessControlResponseFromPb(pb *pipelineAccessControlResponsePb) (*PipelineAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineAccessControlResponse{}
	st.AllPermissions = pb.AllPermissions
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelineClusterToPb(st *PipelineCluster) (*pipelineClusterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineClusterPb{}
	pb.ApplyPolicyDefaultValues = st.ApplyPolicyDefaultValues
	pb.Autoscale = st.Autoscale
	pb.AwsAttributes = st.AwsAttributes
	pb.AzureAttributes = st.AzureAttributes
	pb.ClusterLogConf = st.ClusterLogConf
	pb.CustomTags = st.CustomTags
	pb.DriverInstancePoolId = st.DriverInstancePoolId
	pb.DriverNodeTypeId = st.DriverNodeTypeId
	pb.EnableLocalDiskEncryption = st.EnableLocalDiskEncryption
	pb.GcpAttributes = st.GcpAttributes
	pb.InitScripts = st.InitScripts
	pb.InstancePoolId = st.InstancePoolId
	pb.Label = st.Label
	pb.NodeTypeId = st.NodeTypeId
	pb.NumWorkers = st.NumWorkers
	pb.PolicyId = st.PolicyId
	pb.SparkConf = st.SparkConf
	pb.SparkEnvVars = st.SparkEnvVars
	pb.SshPublicKeys = st.SshPublicKeys

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineClusterPb struct {
	ApplyPolicyDefaultValues  bool                      `json:"apply_policy_default_values,omitempty"`
	Autoscale                 *PipelineClusterAutoscale `json:"autoscale,omitempty"`
	AwsAttributes             *compute.AwsAttributes    `json:"aws_attributes,omitempty"`
	AzureAttributes           *compute.AzureAttributes  `json:"azure_attributes,omitempty"`
	ClusterLogConf            *compute.ClusterLogConf   `json:"cluster_log_conf,omitempty"`
	CustomTags                map[string]string         `json:"custom_tags,omitempty"`
	DriverInstancePoolId      string                    `json:"driver_instance_pool_id,omitempty"`
	DriverNodeTypeId          string                    `json:"driver_node_type_id,omitempty"`
	EnableLocalDiskEncryption bool                      `json:"enable_local_disk_encryption,omitempty"`
	GcpAttributes             *compute.GcpAttributes    `json:"gcp_attributes,omitempty"`
	InitScripts               []compute.InitScriptInfo  `json:"init_scripts,omitempty"`
	InstancePoolId            string                    `json:"instance_pool_id,omitempty"`
	Label                     string                    `json:"label,omitempty"`
	NodeTypeId                string                    `json:"node_type_id,omitempty"`
	NumWorkers                int                       `json:"num_workers,omitempty"`
	PolicyId                  string                    `json:"policy_id,omitempty"`
	SparkConf                 map[string]string         `json:"spark_conf,omitempty"`
	SparkEnvVars              map[string]string         `json:"spark_env_vars,omitempty"`
	SshPublicKeys             []string                  `json:"ssh_public_keys,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineClusterFromPb(pb *pipelineClusterPb) (*PipelineCluster, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineCluster{}
	st.ApplyPolicyDefaultValues = pb.ApplyPolicyDefaultValues
	st.Autoscale = pb.Autoscale
	st.AwsAttributes = pb.AwsAttributes
	st.AzureAttributes = pb.AzureAttributes
	st.ClusterLogConf = pb.ClusterLogConf
	st.CustomTags = pb.CustomTags
	st.DriverInstancePoolId = pb.DriverInstancePoolId
	st.DriverNodeTypeId = pb.DriverNodeTypeId
	st.EnableLocalDiskEncryption = pb.EnableLocalDiskEncryption
	st.GcpAttributes = pb.GcpAttributes
	st.InitScripts = pb.InitScripts
	st.InstancePoolId = pb.InstancePoolId
	st.Label = pb.Label
	st.NodeTypeId = pb.NodeTypeId
	st.NumWorkers = pb.NumWorkers
	st.PolicyId = pb.PolicyId
	st.SparkConf = pb.SparkConf
	st.SparkEnvVars = pb.SparkEnvVars
	st.SshPublicKeys = pb.SshPublicKeys

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineClusterPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineClusterPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelineClusterAutoscaleToPb(st *PipelineClusterAutoscale) (*pipelineClusterAutoscalePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineClusterAutoscalePb{}
	pb.MaxWorkers = st.MaxWorkers
	pb.MinWorkers = st.MinWorkers
	pb.Mode = st.Mode

	return pb, nil
}

type pipelineClusterAutoscalePb struct {
	MaxWorkers int                          `json:"max_workers"`
	MinWorkers int                          `json:"min_workers"`
	Mode       PipelineClusterAutoscaleMode `json:"mode,omitempty"`
}

func pipelineClusterAutoscaleFromPb(pb *pipelineClusterAutoscalePb) (*PipelineClusterAutoscale, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineClusterAutoscale{}
	st.MaxWorkers = pb.MaxWorkers
	st.MinWorkers = pb.MinWorkers
	st.Mode = pb.Mode

	return st, nil
}

func pipelineDeploymentToPb(st *PipelineDeployment) (*pipelineDeploymentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineDeploymentPb{}
	pb.Kind = st.Kind
	pb.MetadataFilePath = st.MetadataFilePath

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineDeploymentPb struct {
	Kind             DeploymentKind `json:"kind"`
	MetadataFilePath string         `json:"metadata_file_path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineDeploymentFromPb(pb *pipelineDeploymentPb) (*PipelineDeployment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineDeployment{}
	st.Kind = pb.Kind
	st.MetadataFilePath = pb.MetadataFilePath

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineDeploymentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineDeploymentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelineEventToPb(st *PipelineEvent) (*pipelineEventPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineEventPb{}
	pb.Error = st.Error
	pb.EventType = st.EventType
	pb.Id = st.Id
	pb.Level = st.Level
	pb.MaturityLevel = st.MaturityLevel
	pb.Message = st.Message
	pb.Origin = st.Origin
	pb.Sequence = st.Sequence
	pb.Timestamp = st.Timestamp

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineEventPb struct {
	Error         *ErrorDetail  `json:"error,omitempty"`
	EventType     string        `json:"event_type,omitempty"`
	Id            string        `json:"id,omitempty"`
	Level         EventLevel    `json:"level,omitempty"`
	MaturityLevel MaturityLevel `json:"maturity_level,omitempty"`
	Message       string        `json:"message,omitempty"`
	Origin        *Origin       `json:"origin,omitempty"`
	Sequence      *Sequencing   `json:"sequence,omitempty"`
	Timestamp     string        `json:"timestamp,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineEventFromPb(pb *pipelineEventPb) (*PipelineEvent, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineEvent{}
	st.Error = pb.Error
	st.EventType = pb.EventType
	st.Id = pb.Id
	st.Level = pb.Level
	st.MaturityLevel = pb.MaturityLevel
	st.Message = pb.Message
	st.Origin = pb.Origin
	st.Sequence = pb.Sequence
	st.Timestamp = pb.Timestamp

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineEventPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineEventPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelineLibraryToPb(st *PipelineLibrary) (*pipelineLibraryPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineLibraryPb{}
	pb.File = st.File
	pb.Glob = st.Glob
	pb.Jar = st.Jar
	pb.Maven = st.Maven
	pb.Notebook = st.Notebook
	pb.Whl = st.Whl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineLibraryPb struct {
	File     *FileLibrary          `json:"file,omitempty"`
	Glob     *PathPattern          `json:"glob,omitempty"`
	Jar      string                `json:"jar,omitempty"`
	Maven    *compute.MavenLibrary `json:"maven,omitempty"`
	Notebook *NotebookLibrary      `json:"notebook,omitempty"`
	Whl      string                `json:"whl,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineLibraryFromPb(pb *pipelineLibraryPb) (*PipelineLibrary, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineLibrary{}
	st.File = pb.File
	st.Glob = pb.Glob
	st.Jar = pb.Jar
	st.Maven = pb.Maven
	st.Notebook = pb.Notebook
	st.Whl = pb.Whl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineLibraryPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineLibraryPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelinePermissionToPb(st *PipelinePermission) (*pipelinePermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinePermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelinePermissionPb struct {
	Inherited           bool                    `json:"inherited,omitempty"`
	InheritedFromObject []string                `json:"inherited_from_object,omitempty"`
	PermissionLevel     PipelinePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelinePermissionFromPb(pb *pipelinePermissionPb) (*PipelinePermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelinePermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelinePermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelinePermissionsToPb(st *PipelinePermissions) (*pipelinePermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinePermissionsPb{}
	pb.AccessControlList = st.AccessControlList
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelinePermissionsPb struct {
	AccessControlList []PipelineAccessControlResponse `json:"access_control_list,omitempty"`
	ObjectId          string                          `json:"object_id,omitempty"`
	ObjectType        string                          `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelinePermissionsFromPb(pb *pipelinePermissionsPb) (*PipelinePermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermissions{}
	st.AccessControlList = pb.AccessControlList
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelinePermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelinePermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelinePermissionsDescriptionToPb(st *PipelinePermissionsDescription) (*pipelinePermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinePermissionsDescriptionPb{}
	pb.Description = st.Description
	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelinePermissionsDescriptionPb struct {
	Description     string                  `json:"description,omitempty"`
	PermissionLevel PipelinePermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelinePermissionsDescriptionFromPb(pb *pipelinePermissionsDescriptionPb) (*PipelinePermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelinePermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelinePermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelinePermissionsRequestToPb(st *PipelinePermissionsRequest) (*pipelinePermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinePermissionsRequestPb{}
	pb.AccessControlList = st.AccessControlList
	pb.PipelineId = st.PipelineId

	return pb, nil
}

type pipelinePermissionsRequestPb struct {
	AccessControlList []PipelineAccessControlRequest `json:"access_control_list,omitempty"`
	PipelineId        string                         `json:"-" url:"-"`
}

func pipelinePermissionsRequestFromPb(pb *pipelinePermissionsRequestPb) (*PipelinePermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinePermissionsRequest{}
	st.AccessControlList = pb.AccessControlList
	st.PipelineId = pb.PipelineId

	return st, nil
}

func pipelineSpecToPb(st *PipelineSpec) (*pipelineSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineSpecPb{}
	pb.BudgetPolicyId = st.BudgetPolicyId
	pb.Catalog = st.Catalog
	pb.Channel = st.Channel
	pb.Clusters = st.Clusters
	pb.Configuration = st.Configuration
	pb.Continuous = st.Continuous
	pb.Deployment = st.Deployment
	pb.Development = st.Development
	pb.Edition = st.Edition
	pb.Environment = st.Environment
	pb.EventLog = st.EventLog
	pb.Filters = st.Filters
	pb.GatewayDefinition = st.GatewayDefinition
	pb.Id = st.Id
	pb.IngestionDefinition = st.IngestionDefinition
	pb.Libraries = st.Libraries
	pb.Name = st.Name
	pb.Notifications = st.Notifications
	pb.Photon = st.Photon
	pb.RestartWindow = st.RestartWindow
	pb.RootPath = st.RootPath
	pb.Schema = st.Schema
	pb.Serverless = st.Serverless
	pb.Storage = st.Storage
	pb.Tags = st.Tags
	pb.Target = st.Target
	pb.Trigger = st.Trigger

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineSpecPb struct {
	BudgetPolicyId      string                              `json:"budget_policy_id,omitempty"`
	Catalog             string                              `json:"catalog,omitempty"`
	Channel             string                              `json:"channel,omitempty"`
	Clusters            []PipelineCluster                   `json:"clusters,omitempty"`
	Configuration       map[string]string                   `json:"configuration,omitempty"`
	Continuous          bool                                `json:"continuous,omitempty"`
	Deployment          *PipelineDeployment                 `json:"deployment,omitempty"`
	Development         bool                                `json:"development,omitempty"`
	Edition             string                              `json:"edition,omitempty"`
	Environment         *PipelinesEnvironment               `json:"environment,omitempty"`
	EventLog            *EventLogSpec                       `json:"event_log,omitempty"`
	Filters             *Filters                            `json:"filters,omitempty"`
	GatewayDefinition   *IngestionGatewayPipelineDefinition `json:"gateway_definition,omitempty"`
	Id                  string                              `json:"id,omitempty"`
	IngestionDefinition *IngestionPipelineDefinition        `json:"ingestion_definition,omitempty"`
	Libraries           []PipelineLibrary                   `json:"libraries,omitempty"`
	Name                string                              `json:"name,omitempty"`
	Notifications       []Notifications                     `json:"notifications,omitempty"`
	Photon              bool                                `json:"photon,omitempty"`
	RestartWindow       *RestartWindow                      `json:"restart_window,omitempty"`
	RootPath            string                              `json:"root_path,omitempty"`
	Schema              string                              `json:"schema,omitempty"`
	Serverless          bool                                `json:"serverless,omitempty"`
	Storage             string                              `json:"storage,omitempty"`
	Tags                map[string]string                   `json:"tags,omitempty"`
	Target              string                              `json:"target,omitempty"`
	Trigger             *PipelineTrigger                    `json:"trigger,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineSpecFromPb(pb *pipelineSpecPb) (*PipelineSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineSpec{}
	st.BudgetPolicyId = pb.BudgetPolicyId
	st.Catalog = pb.Catalog
	st.Channel = pb.Channel
	st.Clusters = pb.Clusters
	st.Configuration = pb.Configuration
	st.Continuous = pb.Continuous
	st.Deployment = pb.Deployment
	st.Development = pb.Development
	st.Edition = pb.Edition
	st.Environment = pb.Environment
	st.EventLog = pb.EventLog
	st.Filters = pb.Filters
	st.GatewayDefinition = pb.GatewayDefinition
	st.Id = pb.Id
	st.IngestionDefinition = pb.IngestionDefinition
	st.Libraries = pb.Libraries
	st.Name = pb.Name
	st.Notifications = pb.Notifications
	st.Photon = pb.Photon
	st.RestartWindow = pb.RestartWindow
	st.RootPath = pb.RootPath
	st.Schema = pb.Schema
	st.Serverless = pb.Serverless
	st.Storage = pb.Storage
	st.Tags = pb.Tags
	st.Target = pb.Target
	st.Trigger = pb.Trigger

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelineStateInfoToPb(st *PipelineStateInfo) (*pipelineStateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineStateInfoPb{}
	pb.ClusterId = st.ClusterId
	pb.CreatorUserName = st.CreatorUserName
	pb.Health = st.Health
	pb.LatestUpdates = st.LatestUpdates
	pb.Name = st.Name
	pb.PipelineId = st.PipelineId
	pb.RunAsUserName = st.RunAsUserName
	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type pipelineStateInfoPb struct {
	ClusterId       string                  `json:"cluster_id,omitempty"`
	CreatorUserName string                  `json:"creator_user_name,omitempty"`
	Health          PipelineStateInfoHealth `json:"health,omitempty"`
	LatestUpdates   []UpdateStateInfo       `json:"latest_updates,omitempty"`
	Name            string                  `json:"name,omitempty"`
	PipelineId      string                  `json:"pipeline_id,omitempty"`
	RunAsUserName   string                  `json:"run_as_user_name,omitempty"`
	State           PipelineState           `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func pipelineStateInfoFromPb(pb *pipelineStateInfoPb) (*PipelineStateInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineStateInfo{}
	st.ClusterId = pb.ClusterId
	st.CreatorUserName = pb.CreatorUserName
	st.Health = pb.Health
	st.LatestUpdates = pb.LatestUpdates
	st.Name = pb.Name
	st.PipelineId = pb.PipelineId
	st.RunAsUserName = pb.RunAsUserName
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *pipelineStateInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st pipelineStateInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func pipelineTriggerToPb(st *PipelineTrigger) (*pipelineTriggerPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelineTriggerPb{}
	pb.Cron = st.Cron
	pb.Manual = st.Manual

	return pb, nil
}

type pipelineTriggerPb struct {
	Cron   *CronTrigger   `json:"cron,omitempty"`
	Manual *ManualTrigger `json:"manual,omitempty"`
}

func pipelineTriggerFromPb(pb *pipelineTriggerPb) (*PipelineTrigger, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelineTrigger{}
	st.Cron = pb.Cron
	st.Manual = pb.Manual

	return st, nil
}

func pipelinesEnvironmentToPb(st *PipelinesEnvironment) (*pipelinesEnvironmentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &pipelinesEnvironmentPb{}
	pb.Dependencies = st.Dependencies

	return pb, nil
}

type pipelinesEnvironmentPb struct {
	Dependencies []string `json:"dependencies,omitempty"`
}

func pipelinesEnvironmentFromPb(pb *pipelinesEnvironmentPb) (*PipelinesEnvironment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PipelinesEnvironment{}
	st.Dependencies = pb.Dependencies

	return st, nil
}

func reportSpecToPb(st *ReportSpec) (*reportSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &reportSpecPb{}
	pb.DestinationCatalog = st.DestinationCatalog
	pb.DestinationSchema = st.DestinationSchema
	pb.DestinationTable = st.DestinationTable
	pb.SourceUrl = st.SourceUrl
	pb.TableConfiguration = st.TableConfiguration

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type reportSpecPb struct {
	DestinationCatalog string               `json:"destination_catalog"`
	DestinationSchema  string               `json:"destination_schema"`
	DestinationTable   string               `json:"destination_table,omitempty"`
	SourceUrl          string               `json:"source_url"`
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func reportSpecFromPb(pb *reportSpecPb) (*ReportSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ReportSpec{}
	st.DestinationCatalog = pb.DestinationCatalog
	st.DestinationSchema = pb.DestinationSchema
	st.DestinationTable = pb.DestinationTable
	st.SourceUrl = pb.SourceUrl
	st.TableConfiguration = pb.TableConfiguration

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *reportSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st reportSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func restartWindowToPb(st *RestartWindow) (*restartWindowPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restartWindowPb{}
	pb.DaysOfWeek = st.DaysOfWeek
	pb.StartHour = st.StartHour
	pb.TimeZoneId = st.TimeZoneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type restartWindowPb struct {
	DaysOfWeek []DayOfWeek `json:"days_of_week,omitempty"`
	StartHour  int         `json:"start_hour"`
	TimeZoneId string      `json:"time_zone_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func restartWindowFromPb(pb *restartWindowPb) (*RestartWindow, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestartWindow{}
	st.DaysOfWeek = pb.DaysOfWeek
	st.StartHour = pb.StartHour
	st.TimeZoneId = pb.TimeZoneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *restartWindowPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st restartWindowPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func runAsToPb(st *RunAs) (*runAsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runAsPb{}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type runAsPb struct {
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	UserName             string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runAsFromPb(pb *runAsPb) (*RunAs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunAs{}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runAsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runAsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func schemaSpecToPb(st *SchemaSpec) (*schemaSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &schemaSpecPb{}
	pb.DestinationCatalog = st.DestinationCatalog
	pb.DestinationSchema = st.DestinationSchema
	pb.SourceCatalog = st.SourceCatalog
	pb.SourceSchema = st.SourceSchema
	pb.TableConfiguration = st.TableConfiguration

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type schemaSpecPb struct {
	DestinationCatalog string               `json:"destination_catalog"`
	DestinationSchema  string               `json:"destination_schema"`
	SourceCatalog      string               `json:"source_catalog,omitempty"`
	SourceSchema       string               `json:"source_schema"`
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func schemaSpecFromPb(pb *schemaSpecPb) (*SchemaSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SchemaSpec{}
	st.DestinationCatalog = pb.DestinationCatalog
	st.DestinationSchema = pb.DestinationSchema
	st.SourceCatalog = pb.SourceCatalog
	st.SourceSchema = pb.SourceSchema
	st.TableConfiguration = pb.TableConfiguration

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *schemaSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st schemaSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func sequencingToPb(st *Sequencing) (*sequencingPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &sequencingPb{}
	pb.ControlPlaneSeqNo = st.ControlPlaneSeqNo
	pb.DataPlaneId = st.DataPlaneId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type sequencingPb struct {
	ControlPlaneSeqNo int64        `json:"control_plane_seq_no,omitempty"`
	DataPlaneId       *DataPlaneId `json:"data_plane_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func sequencingFromPb(pb *sequencingPb) (*Sequencing, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Sequencing{}
	st.ControlPlaneSeqNo = pb.ControlPlaneSeqNo
	st.DataPlaneId = pb.DataPlaneId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *sequencingPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st sequencingPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func serializedExceptionToPb(st *SerializedException) (*serializedExceptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &serializedExceptionPb{}
	pb.ClassName = st.ClassName
	pb.Message = st.Message
	pb.Stack = st.Stack

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type serializedExceptionPb struct {
	ClassName string       `json:"class_name,omitempty"`
	Message   string       `json:"message,omitempty"`
	Stack     []StackFrame `json:"stack,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func serializedExceptionFromPb(pb *serializedExceptionPb) (*SerializedException, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SerializedException{}
	st.ClassName = pb.ClassName
	st.Message = pb.Message
	st.Stack = pb.Stack

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *serializedExceptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st serializedExceptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func stackFrameToPb(st *StackFrame) (*stackFramePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stackFramePb{}
	pb.DeclaringClass = st.DeclaringClass
	pb.FileName = st.FileName
	pb.LineNumber = st.LineNumber
	pb.MethodName = st.MethodName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type stackFramePb struct {
	DeclaringClass string `json:"declaring_class,omitempty"`
	FileName       string `json:"file_name,omitempty"`
	LineNumber     int    `json:"line_number,omitempty"`
	MethodName     string `json:"method_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func stackFrameFromPb(pb *stackFramePb) (*StackFrame, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StackFrame{}
	st.DeclaringClass = pb.DeclaringClass
	st.FileName = pb.FileName
	st.LineNumber = pb.LineNumber
	st.MethodName = pb.MethodName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *stackFramePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st stackFramePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func startUpdateToPb(st *StartUpdate) (*startUpdatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startUpdatePb{}
	pb.Cause = st.Cause
	pb.FullRefresh = st.FullRefresh
	pb.FullRefreshSelection = st.FullRefreshSelection
	pb.PipelineId = st.PipelineId
	pb.RefreshSelection = st.RefreshSelection
	pb.ValidateOnly = st.ValidateOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type startUpdatePb struct {
	Cause                StartUpdateCause `json:"cause,omitempty"`
	FullRefresh          bool             `json:"full_refresh,omitempty"`
	FullRefreshSelection []string         `json:"full_refresh_selection,omitempty"`
	PipelineId           string           `json:"-" url:"-"`
	RefreshSelection     []string         `json:"refresh_selection,omitempty"`
	ValidateOnly         bool             `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func startUpdateFromPb(pb *startUpdatePb) (*StartUpdate, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartUpdate{}
	st.Cause = pb.Cause
	st.FullRefresh = pb.FullRefresh
	st.FullRefreshSelection = pb.FullRefreshSelection
	st.PipelineId = pb.PipelineId
	st.RefreshSelection = pb.RefreshSelection
	st.ValidateOnly = pb.ValidateOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *startUpdatePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st startUpdatePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func startUpdateResponseToPb(st *StartUpdateResponse) (*startUpdateResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &startUpdateResponsePb{}
	pb.UpdateId = st.UpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type startUpdateResponsePb struct {
	UpdateId string `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func startUpdateResponseFromPb(pb *startUpdateResponsePb) (*StartUpdateResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StartUpdateResponse{}
	st.UpdateId = pb.UpdateId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *startUpdateResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st startUpdateResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func stopPipelineResponseToPb(st *StopPipelineResponse) (*stopPipelineResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopPipelineResponsePb{}

	return pb, nil
}

type stopPipelineResponsePb struct {
}

func stopPipelineResponseFromPb(pb *stopPipelineResponsePb) (*StopPipelineResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StopPipelineResponse{}

	return st, nil
}

func stopRequestToPb(st *StopRequest) (*stopRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &stopRequestPb{}
	pb.PipelineId = st.PipelineId

	return pb, nil
}

type stopRequestPb struct {
	PipelineId string `json:"-" url:"-"`
}

func stopRequestFromPb(pb *stopRequestPb) (*StopRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &StopRequest{}
	st.PipelineId = pb.PipelineId

	return st, nil
}

func tableSpecToPb(st *TableSpec) (*tableSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableSpecPb{}
	pb.DestinationCatalog = st.DestinationCatalog
	pb.DestinationSchema = st.DestinationSchema
	pb.DestinationTable = st.DestinationTable
	pb.SourceCatalog = st.SourceCatalog
	pb.SourceSchema = st.SourceSchema
	pb.SourceTable = st.SourceTable
	pb.TableConfiguration = st.TableConfiguration

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tableSpecPb struct {
	DestinationCatalog string               `json:"destination_catalog"`
	DestinationSchema  string               `json:"destination_schema"`
	DestinationTable   string               `json:"destination_table,omitempty"`
	SourceCatalog      string               `json:"source_catalog,omitempty"`
	SourceSchema       string               `json:"source_schema,omitempty"`
	SourceTable        string               `json:"source_table"`
	TableConfiguration *TableSpecificConfig `json:"table_configuration,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableSpecFromPb(pb *tableSpecPb) (*TableSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableSpec{}
	st.DestinationCatalog = pb.DestinationCatalog
	st.DestinationSchema = pb.DestinationSchema
	st.DestinationTable = pb.DestinationTable
	st.SourceCatalog = pb.SourceCatalog
	st.SourceSchema = pb.SourceSchema
	st.SourceTable = pb.SourceTable
	st.TableConfiguration = pb.TableConfiguration

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func tableSpecificConfigToPb(st *TableSpecificConfig) (*tableSpecificConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &tableSpecificConfigPb{}
	pb.ExcludeColumns = st.ExcludeColumns
	pb.IncludeColumns = st.IncludeColumns
	pb.PrimaryKeys = st.PrimaryKeys
	pb.SalesforceIncludeFormulaFields = st.SalesforceIncludeFormulaFields
	pb.ScdType = st.ScdType
	pb.SequenceBy = st.SequenceBy

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type tableSpecificConfigPb struct {
	ExcludeColumns                 []string                   `json:"exclude_columns,omitempty"`
	IncludeColumns                 []string                   `json:"include_columns,omitempty"`
	PrimaryKeys                    []string                   `json:"primary_keys,omitempty"`
	SalesforceIncludeFormulaFields bool                       `json:"salesforce_include_formula_fields,omitempty"`
	ScdType                        TableSpecificConfigScdType `json:"scd_type,omitempty"`
	SequenceBy                     []string                   `json:"sequence_by,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func tableSpecificConfigFromPb(pb *tableSpecificConfigPb) (*TableSpecificConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TableSpecificConfig{}
	st.ExcludeColumns = pb.ExcludeColumns
	st.IncludeColumns = pb.IncludeColumns
	st.PrimaryKeys = pb.PrimaryKeys
	st.SalesforceIncludeFormulaFields = pb.SalesforceIncludeFormulaFields
	st.ScdType = pb.ScdType
	st.SequenceBy = pb.SequenceBy

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *tableSpecificConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st tableSpecificConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateInfoToPb(st *UpdateInfo) (*updateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateInfoPb{}
	pb.Cause = st.Cause
	pb.ClusterId = st.ClusterId
	pb.Config = st.Config
	pb.CreationTime = st.CreationTime
	pb.FullRefresh = st.FullRefresh
	pb.FullRefreshSelection = st.FullRefreshSelection
	pb.PipelineId = st.PipelineId
	pb.RefreshSelection = st.RefreshSelection
	pb.State = st.State
	pb.UpdateId = st.UpdateId
	pb.ValidateOnly = st.ValidateOnly

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateInfoPb struct {
	Cause                UpdateInfoCause `json:"cause,omitempty"`
	ClusterId            string          `json:"cluster_id,omitempty"`
	Config               *PipelineSpec   `json:"config,omitempty"`
	CreationTime         int64           `json:"creation_time,omitempty"`
	FullRefresh          bool            `json:"full_refresh,omitempty"`
	FullRefreshSelection []string        `json:"full_refresh_selection,omitempty"`
	PipelineId           string          `json:"pipeline_id,omitempty"`
	RefreshSelection     []string        `json:"refresh_selection,omitempty"`
	State                UpdateInfoState `json:"state,omitempty"`
	UpdateId             string          `json:"update_id,omitempty"`
	ValidateOnly         bool            `json:"validate_only,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateInfoFromPb(pb *updateInfoPb) (*UpdateInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateInfo{}
	st.Cause = pb.Cause
	st.ClusterId = pb.ClusterId
	st.Config = pb.Config
	st.CreationTime = pb.CreationTime
	st.FullRefresh = pb.FullRefresh
	st.FullRefreshSelection = pb.FullRefreshSelection
	st.PipelineId = pb.PipelineId
	st.RefreshSelection = pb.RefreshSelection
	st.State = pb.State
	st.UpdateId = pb.UpdateId
	st.ValidateOnly = pb.ValidateOnly

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func updateStateInfoToPb(st *UpdateStateInfo) (*updateStateInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateStateInfoPb{}
	pb.CreationTime = st.CreationTime
	pb.State = st.State
	pb.UpdateId = st.UpdateId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type updateStateInfoPb struct {
	CreationTime string               `json:"creation_time,omitempty"`
	State        UpdateStateInfoState `json:"state,omitempty"`
	UpdateId     string               `json:"update_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateStateInfoFromPb(pb *updateStateInfoPb) (*UpdateStateInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateStateInfo{}
	st.CreationTime = pb.CreationTime
	st.State = pb.State
	st.UpdateId = pb.UpdateId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateStateInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateStateInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
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
