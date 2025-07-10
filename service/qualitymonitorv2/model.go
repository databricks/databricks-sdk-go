// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitorv2

import (
	"fmt"

	"github.com/databricks/databricks-sdk-go/marshal"
)

type AnomalyDetectionConfig struct {
	// Run id of the last run of the workflow
	LastRunId string `json:"last_run_id,omitempty"`
	// The status of the last run of the workflow.
	LatestRunStatus AnomalyDetectionRunStatus `json:"latest_run_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *AnomalyDetectionConfig) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s AnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
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

type CreateQualityMonitorRequest struct {
	QualityMonitor QualityMonitor `json:"quality_monitor"`
}

type DeleteQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType string `json:"-" url:"-"`
}

type GetQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType string `json:"-" url:"-"`
}

type ListQualityMonitorRequest struct {
	PageSize int `json:"-" url:"page_size,omitempty"`

	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQualityMonitorRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQualityMonitorRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListQualityMonitorResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	QualityMonitors []QualityMonitor `json:"quality_monitors,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListQualityMonitorResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListQualityMonitorResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type QualityMonitor struct {
	AnomalyDetectionConfig *AnomalyDetectionConfig `json:"anomaly_detection_config,omitempty"`
	// The uuid of the request object. For example, schema id.
	ObjectId string `json:"object_id"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType string `json:"object_type"`
}

type UpdateQualityMonitorRequest struct {
	// The uuid of the request object. For example, schema id.
	ObjectId string `json:"-" url:"-"`
	// The type of the monitored object. Can be one of the following: schema.
	ObjectType string `json:"-" url:"-"`

	QualityMonitor QualityMonitor `json:"quality_monitor"`
}
