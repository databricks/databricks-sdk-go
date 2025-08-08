// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitorv2

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/qualitymonitorv2/qualitymonitorv2pb"
)

type AnomalyDetectionConfig struct {
	// Run id of the last run of the workflow
	// Wire name: 'last_run_id'
	LastRunId string ``
	// The status of the last run of the workflow.
	// Wire name: 'latest_run_status'
	LatestRunStatus AnomalyDetectionRunStatus ``
	ForceSendFields []string                  `tf:"-"`
}

func (st AnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	pb, err := AnomalyDetectionConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *AnomalyDetectionConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &qualitymonitorv2pb.AnomalyDetectionConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := AnomalyDetectionConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func AnomalyDetectionConfigToPb(st *AnomalyDetectionConfig) (*qualitymonitorv2pb.AnomalyDetectionConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &qualitymonitorv2pb.AnomalyDetectionConfigPb{}
	pb.LastRunId = st.LastRunId
	latestRunStatusPb, err := AnomalyDetectionRunStatusToPb(&st.LatestRunStatus)
	if err != nil {
		return nil, err
	}
	if latestRunStatusPb != nil {
		pb.LatestRunStatus = *latestRunStatusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func AnomalyDetectionConfigFromPb(pb *qualitymonitorv2pb.AnomalyDetectionConfigPb) (*AnomalyDetectionConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AnomalyDetectionConfig{}
	st.LastRunId = pb.LastRunId
	latestRunStatusField, err := AnomalyDetectionRunStatusFromPb(&pb.LatestRunStatus)
	if err != nil {
		return nil, err
	}
	if latestRunStatusField != nil {
		st.LatestRunStatus = *latestRunStatusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Status of Anomaly Detection Job Run
type AnomalyDetectionRunStatus string

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusCanceled AnomalyDetectionRunStatus = `ANOMALY_DETECTION_RUN_STATUS_CANCELED`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusFailed AnomalyDetectionRunStatus = `ANOMALY_DETECTION_RUN_STATUS_FAILED`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusJobDeleted AnomalyDetectionRunStatus = `ANOMALY_DETECTION_RUN_STATUS_JOB_DELETED`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusPending AnomalyDetectionRunStatus = `ANOMALY_DETECTION_RUN_STATUS_PENDING`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusRunning AnomalyDetectionRunStatus = `ANOMALY_DETECTION_RUN_STATUS_RUNNING`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusSuccess AnomalyDetectionRunStatus = `ANOMALY_DETECTION_RUN_STATUS_SUCCESS`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusUnknown AnomalyDetectionRunStatus = `ANOMALY_DETECTION_RUN_STATUS_UNKNOWN`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusWorkspaceMismatchError AnomalyDetectionRunStatus = `ANOMALY_DETECTION_RUN_STATUS_WORKSPACE_MISMATCH_ERROR`

// String representation for [fmt.Print]
func (f *AnomalyDetectionRunStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *AnomalyDetectionRunStatus) Set(v string) error {
	switch v {
	case `ANOMALY_DETECTION_RUN_STATUS_CANCELED`, `ANOMALY_DETECTION_RUN_STATUS_FAILED`, `ANOMALY_DETECTION_RUN_STATUS_JOB_DELETED`, `ANOMALY_DETECTION_RUN_STATUS_PENDING`, `ANOMALY_DETECTION_RUN_STATUS_RUNNING`, `ANOMALY_DETECTION_RUN_STATUS_SUCCESS`, `ANOMALY_DETECTION_RUN_STATUS_UNKNOWN`, `ANOMALY_DETECTION_RUN_STATUS_WORKSPACE_MISMATCH_ERROR`:
		*f = AnomalyDetectionRunStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ANOMALY_DETECTION_RUN_STATUS_CANCELED", "ANOMALY_DETECTION_RUN_STATUS_FAILED", "ANOMALY_DETECTION_RUN_STATUS_JOB_DELETED", "ANOMALY_DETECTION_RUN_STATUS_PENDING", "ANOMALY_DETECTION_RUN_STATUS_RUNNING", "ANOMALY_DETECTION_RUN_STATUS_SUCCESS", "ANOMALY_DETECTION_RUN_STATUS_UNKNOWN", "ANOMALY_DETECTION_RUN_STATUS_WORKSPACE_MISMATCH_ERROR"`, v)
	}
}

// Values returns all possible values for AnomalyDetectionRunStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *AnomalyDetectionRunStatus) Values() []AnomalyDetectionRunStatus {
	return []AnomalyDetectionRunStatus{
		AnomalyDetectionRunStatusAnomalyDetectionRunStatusCanceled,
		AnomalyDetectionRunStatusAnomalyDetectionRunStatusFailed,
		AnomalyDetectionRunStatusAnomalyDetectionRunStatusJobDeleted,
		AnomalyDetectionRunStatusAnomalyDetectionRunStatusPending,
		AnomalyDetectionRunStatusAnomalyDetectionRunStatusRunning,
		AnomalyDetectionRunStatusAnomalyDetectionRunStatusSuccess,
		AnomalyDetectionRunStatusAnomalyDetectionRunStatusUnknown,
		AnomalyDetectionRunStatusAnomalyDetectionRunStatusWorkspaceMismatchError,
	}
}

// Type always returns AnomalyDetectionRunStatus to satisfy [pflag.Value] interface
func (f *AnomalyDetectionRunStatus) Type() string {
	return "AnomalyDetectionRunStatus"
}

func AnomalyDetectionRunStatusToPb(st *AnomalyDetectionRunStatus) (*qualitymonitorv2pb.AnomalyDetectionRunStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := qualitymonitorv2pb.AnomalyDetectionRunStatusPb(*st)
	return &pb, nil
}

func AnomalyDetectionRunStatusFromPb(pb *qualitymonitorv2pb.AnomalyDetectionRunStatusPb) (*AnomalyDetectionRunStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := AnomalyDetectionRunStatus(*pb)
	return &st, nil
}

type CreateQualityMonitorRequest struct {

	// Wire name: 'quality_monitor'
	QualityMonitor QualityMonitor ``
}

func (st CreateQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &qualitymonitorv2pb.CreateQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateQualityMonitorRequestToPb(st *CreateQualityMonitorRequest) (*qualitymonitorv2pb.CreateQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &qualitymonitorv2pb.CreateQualityMonitorRequestPb{}
	qualityMonitorPb, err := QualityMonitorToPb(&st.QualityMonitor)
	if err != nil {
		return nil, err
	}
	if qualityMonitorPb != nil {
		pb.QualityMonitor = *qualityMonitorPb
	}

	return pb, nil
}

func CreateQualityMonitorRequestFromPb(pb *qualitymonitorv2pb.CreateQualityMonitorRequestPb) (*CreateQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateQualityMonitorRequest{}
	qualityMonitorField, err := QualityMonitorFromPb(&pb.QualityMonitor)
	if err != nil {
		return nil, err
	}
	if qualityMonitorField != nil {
		st.QualityMonitor = *qualityMonitorField
	}

	return st, nil
}

type DeleteQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	// Wire name: 'object_id'
	ObjectId string `tf:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	// Wire name: 'object_type'
	ObjectType string `tf:"-"`
}

func (st DeleteQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &qualitymonitorv2pb.DeleteQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteQualityMonitorRequestToPb(st *DeleteQualityMonitorRequest) (*qualitymonitorv2pb.DeleteQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &qualitymonitorv2pb.DeleteQualityMonitorRequestPb{}
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	return pb, nil
}

func DeleteQualityMonitorRequestFromPb(pb *qualitymonitorv2pb.DeleteQualityMonitorRequestPb) (*DeleteQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteQualityMonitorRequest{}
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	return st, nil
}

type GetQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	// Wire name: 'object_id'
	ObjectId string `tf:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	// Wire name: 'object_type'
	ObjectType string `tf:"-"`
}

func (st GetQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &qualitymonitorv2pb.GetQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetQualityMonitorRequestToPb(st *GetQualityMonitorRequest) (*qualitymonitorv2pb.GetQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &qualitymonitorv2pb.GetQualityMonitorRequestPb{}
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	return pb, nil
}

func GetQualityMonitorRequestFromPb(pb *qualitymonitorv2pb.GetQualityMonitorRequestPb) (*GetQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQualityMonitorRequest{}
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	return st, nil
}

type ListQualityMonitorRequest struct {

	// Wire name: 'page_size'
	PageSize int `tf:"-"`

	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &qualitymonitorv2pb.ListQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListQualityMonitorRequestToPb(st *ListQualityMonitorRequest) (*qualitymonitorv2pb.ListQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &qualitymonitorv2pb.ListQualityMonitorRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListQualityMonitorRequestFromPb(pb *qualitymonitorv2pb.ListQualityMonitorRequestPb) (*ListQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQualityMonitorRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListQualityMonitorResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'quality_monitors'
	QualityMonitors []QualityMonitor ``
	ForceSendFields []string         `tf:"-"`
}

func (st ListQualityMonitorResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListQualityMonitorResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListQualityMonitorResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &qualitymonitorv2pb.ListQualityMonitorResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListQualityMonitorResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListQualityMonitorResponseToPb(st *ListQualityMonitorResponse) (*qualitymonitorv2pb.ListQualityMonitorResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &qualitymonitorv2pb.ListQualityMonitorResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var qualityMonitorsPb []qualitymonitorv2pb.QualityMonitorPb
	for _, item := range st.QualityMonitors {
		itemPb, err := QualityMonitorToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			qualityMonitorsPb = append(qualityMonitorsPb, *itemPb)
		}
	}
	pb.QualityMonitors = qualityMonitorsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListQualityMonitorResponseFromPb(pb *qualitymonitorv2pb.ListQualityMonitorResponsePb) (*ListQualityMonitorResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQualityMonitorResponse{}
	st.NextPageToken = pb.NextPageToken

	var qualityMonitorsField []QualityMonitor
	for _, itemPb := range pb.QualityMonitors {
		item, err := QualityMonitorFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			qualityMonitorsField = append(qualityMonitorsField, *item)
		}
	}
	st.QualityMonitors = qualityMonitorsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type QualityMonitor struct {

	// Wire name: 'anomaly_detection_config'
	AnomalyDetectionConfig *AnomalyDetectionConfig ``
	// The uuid of the request object. For example, schema id.
	// Wire name: 'object_id'
	ObjectId string ``
	// The type of the monitored object. Can be one of the following: schema.
	// Wire name: 'object_type'
	ObjectType string ``
}

func (st QualityMonitor) MarshalJSON() ([]byte, error) {
	pb, err := QualityMonitorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *QualityMonitor) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &qualitymonitorv2pb.QualityMonitorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := QualityMonitorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func QualityMonitorToPb(st *QualityMonitor) (*qualitymonitorv2pb.QualityMonitorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &qualitymonitorv2pb.QualityMonitorPb{}
	anomalyDetectionConfigPb, err := AnomalyDetectionConfigToPb(st.AnomalyDetectionConfig)
	if err != nil {
		return nil, err
	}
	if anomalyDetectionConfigPb != nil {
		pb.AnomalyDetectionConfig = anomalyDetectionConfigPb
	}
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	return pb, nil
}

func QualityMonitorFromPb(pb *qualitymonitorv2pb.QualityMonitorPb) (*QualityMonitor, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QualityMonitor{}
	anomalyDetectionConfigField, err := AnomalyDetectionConfigFromPb(pb.AnomalyDetectionConfig)
	if err != nil {
		return nil, err
	}
	if anomalyDetectionConfigField != nil {
		st.AnomalyDetectionConfig = anomalyDetectionConfigField
	}
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	return st, nil
}

type UpdateQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	// Wire name: 'object_id'
	ObjectId string `tf:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	// Wire name: 'object_type'
	ObjectType string `tf:"-"`

	// Wire name: 'quality_monitor'
	QualityMonitor QualityMonitor ``
}

func (st UpdateQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &qualitymonitorv2pb.UpdateQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateQualityMonitorRequestToPb(st *UpdateQualityMonitorRequest) (*qualitymonitorv2pb.UpdateQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &qualitymonitorv2pb.UpdateQualityMonitorRequestPb{}
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType
	qualityMonitorPb, err := QualityMonitorToPb(&st.QualityMonitor)
	if err != nil {
		return nil, err
	}
	if qualityMonitorPb != nil {
		pb.QualityMonitor = *qualityMonitorPb
	}

	return pb, nil
}

func UpdateQualityMonitorRequestFromPb(pb *qualitymonitorv2pb.UpdateQualityMonitorRequestPb) (*UpdateQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateQualityMonitorRequest{}
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType
	qualityMonitorField, err := QualityMonitorFromPb(&pb.QualityMonitor)
	if err != nil {
		return nil, err
	}
	if qualityMonitorField != nil {
		st.QualityMonitor = *qualityMonitorField
	}

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
