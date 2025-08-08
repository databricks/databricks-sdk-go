// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitorv2pb

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type AnomalyDetectionConfigPb struct {
	LastRunId       string                      `json:"last_run_id,omitempty"`
	LatestRunStatus AnomalyDetectionRunStatusPb `json:"latest_run_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *AnomalyDetectionConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st AnomalyDetectionConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type AnomalyDetectionRunStatusPb string

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusCanceled AnomalyDetectionRunStatusPb = `ANOMALY_DETECTION_RUN_STATUS_CANCELED`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusFailed AnomalyDetectionRunStatusPb = `ANOMALY_DETECTION_RUN_STATUS_FAILED`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusJobDeleted AnomalyDetectionRunStatusPb = `ANOMALY_DETECTION_RUN_STATUS_JOB_DELETED`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusPending AnomalyDetectionRunStatusPb = `ANOMALY_DETECTION_RUN_STATUS_PENDING`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusRunning AnomalyDetectionRunStatusPb = `ANOMALY_DETECTION_RUN_STATUS_RUNNING`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusSuccess AnomalyDetectionRunStatusPb = `ANOMALY_DETECTION_RUN_STATUS_SUCCESS`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusUnknown AnomalyDetectionRunStatusPb = `ANOMALY_DETECTION_RUN_STATUS_UNKNOWN`

const AnomalyDetectionRunStatusAnomalyDetectionRunStatusWorkspaceMismatchError AnomalyDetectionRunStatusPb = `ANOMALY_DETECTION_RUN_STATUS_WORKSPACE_MISMATCH_ERROR`

type CreateQualityMonitorRequestPb struct {
	QualityMonitor QualityMonitorPb `json:"quality_monitor"`
}

type DeleteQualityMonitorRequestPb struct {
	ObjectId   string `json:"-" url:"-"`
	ObjectType string `json:"-" url:"-"`
}

type GetQualityMonitorRequestPb struct {
	ObjectId   string `json:"-" url:"-"`
	ObjectType string `json:"-" url:"-"`
}

type ListQualityMonitorRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListQualityMonitorRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListQualityMonitorRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListQualityMonitorResponsePb struct {
	NextPageToken   string             `json:"next_page_token,omitempty"`
	QualityMonitors []QualityMonitorPb `json:"quality_monitors,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (st *ListQualityMonitorResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st ListQualityMonitorResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type QualityMonitorPb struct {
	AnomalyDetectionConfig *AnomalyDetectionConfigPb `json:"anomaly_detection_config,omitempty"`
	ObjectId               string                    `json:"object_id"`
	ObjectType             string                    `json:"object_type"`
}

type UpdateQualityMonitorRequestPb struct {
	ObjectId       string           `json:"-" url:"-"`
	ObjectType     string           `json:"-" url:"-"`
	QualityMonitor QualityMonitorPb `json:"quality_monitor"`
}
