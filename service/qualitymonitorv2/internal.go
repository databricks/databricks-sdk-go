// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package qualitymonitorv2

import (
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

func anomalyDetectionConfigToPb(st *AnomalyDetectionConfig) (*anomalyDetectionConfigPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &anomalyDetectionConfigPb{}
	pb.LastRunId = st.LastRunId
	pb.LatestRunStatus = st.LatestRunStatus

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type anomalyDetectionConfigPb struct {
	LastRunId       string                    `json:"last_run_id,omitempty"`
	LatestRunStatus AnomalyDetectionRunStatus `json:"latest_run_status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func anomalyDetectionConfigFromPb(pb *anomalyDetectionConfigPb) (*AnomalyDetectionConfig, error) {
	if pb == nil {
		return nil, nil
	}
	st := &AnomalyDetectionConfig{}
	st.LastRunId = pb.LastRunId
	st.LatestRunStatus = pb.LatestRunStatus

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *anomalyDetectionConfigPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st anomalyDetectionConfigPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func createQualityMonitorRequestToPb(st *CreateQualityMonitorRequest) (*createQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createQualityMonitorRequestPb{}
	pb.QualityMonitor = st.QualityMonitor

	return pb, nil
}

type createQualityMonitorRequestPb struct {
	QualityMonitor QualityMonitor `json:"quality_monitor"`
}

func createQualityMonitorRequestFromPb(pb *createQualityMonitorRequestPb) (*CreateQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateQualityMonitorRequest{}
	st.QualityMonitor = pb.QualityMonitor

	return st, nil
}

func deleteQualityMonitorRequestToPb(st *DeleteQualityMonitorRequest) (*deleteQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQualityMonitorRequestPb{}
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	return pb, nil
}

type deleteQualityMonitorRequestPb struct {
	ObjectId   string `json:"-" url:"-"`
	ObjectType string `json:"-" url:"-"`
}

func deleteQualityMonitorRequestFromPb(pb *deleteQualityMonitorRequestPb) (*DeleteQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteQualityMonitorRequest{}
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	return st, nil
}

func deleteQualityMonitorResponseToPb(st *DeleteQualityMonitorResponse) (*deleteQualityMonitorResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteQualityMonitorResponsePb{}

	return pb, nil
}

type deleteQualityMonitorResponsePb struct {
}

func deleteQualityMonitorResponseFromPb(pb *deleteQualityMonitorResponsePb) (*DeleteQualityMonitorResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteQualityMonitorResponse{}

	return st, nil
}

func getQualityMonitorRequestToPb(st *GetQualityMonitorRequest) (*getQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getQualityMonitorRequestPb{}
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	return pb, nil
}

type getQualityMonitorRequestPb struct {
	ObjectId   string `json:"-" url:"-"`
	ObjectType string `json:"-" url:"-"`
}

func getQualityMonitorRequestFromPb(pb *getQualityMonitorRequestPb) (*GetQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetQualityMonitorRequest{}
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	return st, nil
}

func listQualityMonitorRequestToPb(st *ListQualityMonitorRequest) (*listQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQualityMonitorRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listQualityMonitorRequestPb struct {
	PageSize  int    `json:"-" url:"page_size,omitempty"`
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQualityMonitorRequestFromPb(pb *listQualityMonitorRequestPb) (*ListQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQualityMonitorRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQualityMonitorRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQualityMonitorRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func listQualityMonitorResponseToPb(st *ListQualityMonitorResponse) (*listQualityMonitorResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listQualityMonitorResponsePb{}
	pb.NextPageToken = st.NextPageToken
	pb.QualityMonitors = st.QualityMonitors

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

type listQualityMonitorResponsePb struct {
	NextPageToken   string           `json:"next_page_token,omitempty"`
	QualityMonitors []QualityMonitor `json:"quality_monitors,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listQualityMonitorResponseFromPb(pb *listQualityMonitorResponsePb) (*ListQualityMonitorResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListQualityMonitorResponse{}
	st.NextPageToken = pb.NextPageToken
	st.QualityMonitors = pb.QualityMonitors

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listQualityMonitorResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listQualityMonitorResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

func qualityMonitorToPb(st *QualityMonitor) (*qualityMonitorPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &qualityMonitorPb{}
	pb.AnomalyDetectionConfig = st.AnomalyDetectionConfig
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType

	return pb, nil
}

type qualityMonitorPb struct {
	AnomalyDetectionConfig *AnomalyDetectionConfig `json:"anomaly_detection_config,omitempty"`
	ObjectId               string                  `json:"object_id"`
	ObjectType             string                  `json:"object_type"`
}

func qualityMonitorFromPb(pb *qualityMonitorPb) (*QualityMonitor, error) {
	if pb == nil {
		return nil, nil
	}
	st := &QualityMonitor{}
	st.AnomalyDetectionConfig = pb.AnomalyDetectionConfig
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	return st, nil
}

func updateQualityMonitorRequestToPb(st *UpdateQualityMonitorRequest) (*updateQualityMonitorRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateQualityMonitorRequestPb{}
	pb.ObjectId = st.ObjectId
	pb.ObjectType = st.ObjectType
	pb.QualityMonitor = st.QualityMonitor

	return pb, nil
}

type updateQualityMonitorRequestPb struct {
	ObjectId       string         `json:"-" url:"-"`
	ObjectType     string         `json:"-" url:"-"`
	QualityMonitor QualityMonitor `json:"quality_monitor"`
}

func updateQualityMonitorRequestFromPb(pb *updateQualityMonitorRequestPb) (*UpdateQualityMonitorRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateQualityMonitorRequest{}
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType
	st.QualityMonitor = pb.QualityMonitor

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
