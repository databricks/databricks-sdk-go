// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitorv2

import (
	"encoding/json"
	"fmt"
)

type AnomalyDetectionConfig struct {
	// Run id of the last run of the workflow
	// Wire name: 'last_run_id'
	LastRunId string `json:"last_run_id,omitempty"`
	// The status of the last run of the workflow.
	// Wire name: 'latest_run_status'
	LatestRunStatus AnomalyDetectionRunStatus `json:"latest_run_status,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *AnomalyDetectionConfig) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &anomalyDetectionConfigPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := anomalyDetectionConfigFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st AnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	pb, err := anomalyDetectionConfigToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
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

// Create a quality monitor
type CreateQualityMonitorRequest struct {

	// Wire name: 'quality_monitor'
	QualityMonitor QualityMonitor `json:"quality_monitor"`
}

func (st *CreateQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := createQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Delete a quality monitor
type DeleteQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	ObjectId string `json:"-" tf:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType string `json:"-" tf:"-"`
}

func (st *DeleteQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type DeleteQualityMonitorResponse struct {
}

func (st *DeleteQualityMonitorResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteQualityMonitorResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteQualityMonitorResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteQualityMonitorResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteQualityMonitorResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Read a quality monitor
type GetQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	ObjectId string `json:"-" tf:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType string `json:"-" tf:"-"`
}

func (st *GetQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := getQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// List quality monitors
type ListQualityMonitorRequest struct {
	PageSize int `json:"-" tf:"-"`

	PageToken string `json:"-" tf:"-"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := listQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type ListQualityMonitorResponse struct {

	// Wire name: 'next_page_token'
	NextPageToken string `json:"next_page_token,omitempty"`

	// Wire name: 'quality_monitors'
	QualityMonitors []QualityMonitor `json:"quality_monitors,omitempty"`

	ForceSendFields []string `json:"-" tf:"-"`
}

func (st *ListQualityMonitorResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listQualityMonitorResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listQualityMonitorResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListQualityMonitorResponse) MarshalJSON() ([]byte, error) {
	pb, err := listQualityMonitorResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type QualityMonitor struct {

	// Wire name: 'anomaly_detection_config'
	AnomalyDetectionConfig *AnomalyDetectionConfig `json:"anomaly_detection_config,omitempty"`
	// The uuid of the request object. For example, schema id.
	// Wire name: 'object_id'
	ObjectId string `json:"object_id"`
	// The type of the monitored object. Can be one of the following: schema.
	// Wire name: 'object_type'
	ObjectType string `json:"object_type"`
}

func (st *QualityMonitor) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &qualityMonitorPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := qualityMonitorFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st QualityMonitor) MarshalJSON() ([]byte, error) {
	pb, err := qualityMonitorToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

// Update a quality monitor
type UpdateQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	ObjectId string `json:"-" tf:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType string `json:"-" tf:"-"`

	// Wire name: 'quality_monitor'
	QualityMonitor QualityMonitor `json:"quality_monitor"`
}

func (st *UpdateQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateQualityMonitorRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateQualityMonitorRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateQualityMonitorRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}
