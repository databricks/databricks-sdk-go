// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/marshal"
)

// Activity recorded for the action.
type Activity struct {
	// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied
	// the corresponding stage transition.
	//
	// * `REQUESTED_TRANSITION`: User requested the corresponding stage
	// transition.
	//
	// * `CANCELLED_REQUEST`: User cancelled an existing transition request.
	//
	// * `APPROVED_REQUEST`: User approved the corresponding stage transition.
	//
	// * `REJECTED_REQUEST`: User rejected the coressponding stage transition.
	//
	// * `SYSTEM_TRANSITION`: For events performed as a side effect, such as
	// archiving existing model versions in a stage.
	// Wire name: 'activity_type'
	ActivityType ActivityType
	// User-provided comment associated with the activity.
	// Wire name: 'comment'
	Comment string
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Source stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'from_stage'
	FromStage Stage
	// Unique identifier for the object.
	// Wire name: 'id'
	Id string
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	// Wire name: 'system_comment'
	SystemComment string
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'to_stage'
	ToStage Stage
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func activityToPb(st *Activity) (*activityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &activityPb{}
	pb.ActivityType = st.ActivityType

	pb.Comment = st.Comment

	pb.CreationTimestamp = st.CreationTimestamp

	pb.FromStage = st.FromStage

	pb.Id = st.Id

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	pb.SystemComment = st.SystemComment

	pb.ToStage = st.ToStage

	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Activity) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &activityPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := activityFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Activity) MarshalJSON() ([]byte, error) {
	pb, err := activityToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type activityPb struct {
	// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied
	// the corresponding stage transition.
	//
	// * `REQUESTED_TRANSITION`: User requested the corresponding stage
	// transition.
	//
	// * `CANCELLED_REQUEST`: User cancelled an existing transition request.
	//
	// * `APPROVED_REQUEST`: User approved the corresponding stage transition.
	//
	// * `REJECTED_REQUEST`: User rejected the coressponding stage transition.
	//
	// * `SYSTEM_TRANSITION`: For events performed as a side effect, such as
	// archiving existing model versions in a stage.
	ActivityType ActivityType `json:"activity_type,omitempty"`
	// User-provided comment associated with the activity.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Source stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	FromStage Stage `json:"from_stage,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	SystemComment string `json:"system_comment,omitempty"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	ToStage Stage `json:"to_stage,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func activityFromPb(pb *activityPb) (*Activity, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Activity{}
	st.ActivityType = pb.ActivityType
	st.Comment = pb.Comment
	st.CreationTimestamp = pb.CreationTimestamp
	st.FromStage = pb.FromStage
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.SystemComment = pb.SystemComment
	st.ToStage = pb.ToStage
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *activityPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st activityPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// An action that a user (with sufficient permissions) could take on an
// activity. Valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request
//
// * `REJECT_TRANSITION_REQUEST`: Reject a transition request
//
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
type ActivityAction string
type activityActionPb string

// Approve a transition request
const ActivityActionApproveTransitionRequest ActivityAction = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const ActivityActionCancelTransitionRequest ActivityAction = `CANCEL_TRANSITION_REQUEST`

// Reject a transition request
const ActivityActionRejectTransitionRequest ActivityAction = `REJECT_TRANSITION_REQUEST`

// String representation for [fmt.Print]
func (f *ActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActivityAction) Set(v string) error {
	switch v {
	case `APPROVE_TRANSITION_REQUEST`, `CANCEL_TRANSITION_REQUEST`, `REJECT_TRANSITION_REQUEST`:
		*f = ActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVE_TRANSITION_REQUEST", "CANCEL_TRANSITION_REQUEST", "REJECT_TRANSITION_REQUEST"`, v)
	}
}

// Type always returns ActivityAction to satisfy [pflag.Value] interface
func (f *ActivityAction) Type() string {
	return "ActivityAction"
}

func activityActionToPb(st *ActivityAction) (*activityActionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := activityActionPb(*st)
	return &pb, nil
}

func activityActionFromPb(pb *activityActionPb) (*ActivityAction, error) {
	if pb == nil {
		return nil, nil
	}
	st := ActivityAction(*pb)
	return &st, nil
}

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition.
//
// * `REQUESTED_TRANSITION`: User requested the corresponding stage transition.
//
// * `CANCELLED_REQUEST`: User cancelled an existing transition request.
//
// * `APPROVED_REQUEST`: User approved the corresponding stage transition.
//
// * `REJECTED_REQUEST`: User rejected the coressponding stage transition.
//
// * `SYSTEM_TRANSITION`: For events performed as a side effect, such as
// archiving existing model versions in a stage.
type ActivityType string
type activityTypePb string

// User applied the corresponding stage transition.
const ActivityTypeAppliedTransition ActivityType = `APPLIED_TRANSITION`

// User approved the corresponding stage transition.
const ActivityTypeApprovedRequest ActivityType = `APPROVED_REQUEST`

// User cancelled an existing transition request.
const ActivityTypeCancelledRequest ActivityType = `CANCELLED_REQUEST`

const ActivityTypeNewComment ActivityType = `NEW_COMMENT`

// User rejected the coressponding stage transition.
const ActivityTypeRejectedRequest ActivityType = `REJECTED_REQUEST`

// User requested the corresponding stage transition.
const ActivityTypeRequestedTransition ActivityType = `REQUESTED_TRANSITION`

// For events performed as a side effect, such as archiving existing model
// versions in a stage.
const ActivityTypeSystemTransition ActivityType = `SYSTEM_TRANSITION`

// String representation for [fmt.Print]
func (f *ActivityType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActivityType) Set(v string) error {
	switch v {
	case `APPLIED_TRANSITION`, `APPROVED_REQUEST`, `CANCELLED_REQUEST`, `NEW_COMMENT`, `REJECTED_REQUEST`, `REQUESTED_TRANSITION`, `SYSTEM_TRANSITION`:
		*f = ActivityType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPLIED_TRANSITION", "APPROVED_REQUEST", "CANCELLED_REQUEST", "NEW_COMMENT", "REJECTED_REQUEST", "REQUESTED_TRANSITION", "SYSTEM_TRANSITION"`, v)
	}
}

// Type always returns ActivityType to satisfy [pflag.Value] interface
func (f *ActivityType) Type() string {
	return "ActivityType"
}

func activityTypeToPb(st *ActivityType) (*activityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := activityTypePb(*st)
	return &pb, nil
}

func activityTypeFromPb(pb *activityTypePb) (*ActivityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ActivityType(*pb)
	return &st, nil
}

type ApproveTransitionRequest struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	// Wire name: 'archive_existing_versions'
	ArchiveExistingVersions bool
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'stage'
	Stage Stage
	// Version of the model.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func approveTransitionRequestToPb(st *ApproveTransitionRequest) (*approveTransitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &approveTransitionRequestPb{}
	pb.ArchiveExistingVersions = st.ArchiveExistingVersions

	pb.Comment = st.Comment

	pb.Name = st.Name

	pb.Stage = st.Stage

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ApproveTransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &approveTransitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := approveTransitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ApproveTransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := approveTransitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type approveTransitionRequestPb struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions bool `json:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage Stage `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func approveTransitionRequestFromPb(pb *approveTransitionRequestPb) (*ApproveTransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApproveTransitionRequest{}
	st.ArchiveExistingVersions = pb.ArchiveExistingVersions
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *approveTransitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st approveTransitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ApproveTransitionRequestResponse struct {
	// Activity recorded for the action.
	// Wire name: 'activity'
	Activity *Activity
}

func approveTransitionRequestResponseToPb(st *ApproveTransitionRequestResponse) (*approveTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &approveTransitionRequestResponsePb{}
	activityPb, err := activityToPb(st.Activity)
	if err != nil {
		return nil, err
	}
	if activityPb != nil {
		pb.Activity = activityPb
	}

	return pb, nil
}

func (st *ApproveTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &approveTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := approveTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ApproveTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := approveTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type approveTransitionRequestResponsePb struct {
	// Activity recorded for the action.
	Activity *activityPb `json:"activity,omitempty"`
}

func approveTransitionRequestResponseFromPb(pb *approveTransitionRequestResponsePb) (*ApproveTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApproveTransitionRequestResponse{}
	activityField, err := activityFromPb(pb.Activity)
	if err != nil {
		return nil, err
	}
	if activityField != nil {
		st.Activity = activityField
	}

	return st, nil
}

type ArtifactCredentialInfo struct {
	// A collection of HTTP headers that should be specified when uploading to
	// or downloading from the specified `signed_uri`.
	// Wire name: 'headers'
	Headers []ArtifactCredentialInfoHttpHeader
	// The path, relative to the Run's artifact root location, of the artifact
	// that can be accessed with the credential.
	// Wire name: 'path'
	Path string
	// The ID of the MLflow Run containing the artifact that can be accessed
	// with the credential.
	// Wire name: 'run_id'
	RunId string
	// The signed URI credential that provides access to the artifact.
	// Wire name: 'signed_uri'
	SignedUri string
	// The type of the signed credential URI (e.g., an AWS presigned URL or an
	// Azure Shared Access Signature URI).
	// Wire name: 'type'
	Type ArtifactCredentialType

	ForceSendFields []string `tf:"-"`
}

func artifactCredentialInfoToPb(st *ArtifactCredentialInfo) (*artifactCredentialInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &artifactCredentialInfoPb{}

	var headersPb []artifactCredentialInfoHttpHeaderPb
	for _, item := range st.Headers {
		itemPb, err := artifactCredentialInfoHttpHeaderToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			headersPb = append(headersPb, *itemPb)
		}
	}
	pb.Headers = headersPb

	pb.Path = st.Path

	pb.RunId = st.RunId

	pb.SignedUri = st.SignedUri

	pb.Type = st.Type

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ArtifactCredentialInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &artifactCredentialInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := artifactCredentialInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ArtifactCredentialInfo) MarshalJSON() ([]byte, error) {
	pb, err := artifactCredentialInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type artifactCredentialInfoPb struct {
	// A collection of HTTP headers that should be specified when uploading to
	// or downloading from the specified `signed_uri`.
	Headers []artifactCredentialInfoHttpHeaderPb `json:"headers,omitempty"`
	// The path, relative to the Run's artifact root location, of the artifact
	// that can be accessed with the credential.
	Path string `json:"path,omitempty"`
	// The ID of the MLflow Run containing the artifact that can be accessed
	// with the credential.
	RunId string `json:"run_id,omitempty"`
	// The signed URI credential that provides access to the artifact.
	SignedUri string `json:"signed_uri,omitempty"`
	// The type of the signed credential URI (e.g., an AWS presigned URL or an
	// Azure Shared Access Signature URI).
	Type ArtifactCredentialType `json:"type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func artifactCredentialInfoFromPb(pb *artifactCredentialInfoPb) (*ArtifactCredentialInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ArtifactCredentialInfo{}

	var headersField []ArtifactCredentialInfoHttpHeader
	for _, itemPb := range pb.Headers {
		item, err := artifactCredentialInfoHttpHeaderFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			headersField = append(headersField, *item)
		}
	}
	st.Headers = headersField
	st.Path = pb.Path
	st.RunId = pb.RunId
	st.SignedUri = pb.SignedUri
	st.Type = pb.Type

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *artifactCredentialInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st artifactCredentialInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ArtifactCredentialInfoHttpHeader struct {
	// The HTTP header name.
	// Wire name: 'name'
	Name string
	// The HTTP header value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func artifactCredentialInfoHttpHeaderToPb(st *ArtifactCredentialInfoHttpHeader) (*artifactCredentialInfoHttpHeaderPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &artifactCredentialInfoHttpHeaderPb{}
	pb.Name = st.Name

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ArtifactCredentialInfoHttpHeader) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &artifactCredentialInfoHttpHeaderPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := artifactCredentialInfoHttpHeaderFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ArtifactCredentialInfoHttpHeader) MarshalJSON() ([]byte, error) {
	pb, err := artifactCredentialInfoHttpHeaderToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type artifactCredentialInfoHttpHeaderPb struct {
	// The HTTP header name.
	Name string `json:"name,omitempty"`
	// The HTTP header value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func artifactCredentialInfoHttpHeaderFromPb(pb *artifactCredentialInfoHttpHeaderPb) (*ArtifactCredentialInfoHttpHeader, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ArtifactCredentialInfoHttpHeader{}
	st.Name = pb.Name
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *artifactCredentialInfoHttpHeaderPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st artifactCredentialInfoHttpHeaderPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// The type of a given artifact access credential
type ArtifactCredentialType string
type artifactCredentialTypePb string

const ArtifactCredentialTypeAwsPresignedUrl ArtifactCredentialType = `AWS_PRESIGNED_URL`

const ArtifactCredentialTypeAzureAdlsGen2SasUri ArtifactCredentialType = `AZURE_ADLS_GEN2_SAS_URI`

const ArtifactCredentialTypeAzureSasUri ArtifactCredentialType = `AZURE_SAS_URI`

const ArtifactCredentialTypeGcpSignedUrl ArtifactCredentialType = `GCP_SIGNED_URL`

// String representation for [fmt.Print]
func (f *ArtifactCredentialType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ArtifactCredentialType) Set(v string) error {
	switch v {
	case `AWS_PRESIGNED_URL`, `AZURE_ADLS_GEN2_SAS_URI`, `AZURE_SAS_URI`, `GCP_SIGNED_URL`:
		*f = ArtifactCredentialType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AWS_PRESIGNED_URL", "AZURE_ADLS_GEN2_SAS_URI", "AZURE_SAS_URI", "GCP_SIGNED_URL"`, v)
	}
}

// Type always returns ArtifactCredentialType to satisfy [pflag.Value] interface
func (f *ArtifactCredentialType) Type() string {
	return "ArtifactCredentialType"
}

func artifactCredentialTypeToPb(st *ArtifactCredentialType) (*artifactCredentialTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := artifactCredentialTypePb(*st)
	return &pb, nil
}

func artifactCredentialTypeFromPb(pb *artifactCredentialTypePb) (*ArtifactCredentialType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ArtifactCredentialType(*pb)
	return &st, nil
}

// An action that a user (with sufficient permissions) could take on a comment.
// Valid values are: * `EDIT_COMMENT`: Edit the comment
//
// * `DELETE_COMMENT`: Delete the comment
type CommentActivityAction string
type commentActivityActionPb string

// Delete the comment
const CommentActivityActionDeleteComment CommentActivityAction = `DELETE_COMMENT`

// Edit the comment
const CommentActivityActionEditComment CommentActivityAction = `EDIT_COMMENT`

// String representation for [fmt.Print]
func (f *CommentActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CommentActivityAction) Set(v string) error {
	switch v {
	case `DELETE_COMMENT`, `EDIT_COMMENT`:
		*f = CommentActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "DELETE_COMMENT", "EDIT_COMMENT"`, v)
	}
}

// Type always returns CommentActivityAction to satisfy [pflag.Value] interface
func (f *CommentActivityAction) Type() string {
	return "CommentActivityAction"
}

func commentActivityActionToPb(st *CommentActivityAction) (*commentActivityActionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := commentActivityActionPb(*st)
	return &pb, nil
}

func commentActivityActionFromPb(pb *commentActivityActionPb) (*CommentActivityAction, error) {
	if pb == nil {
		return nil, nil
	}
	st := CommentActivityAction(*pb)
	return &st, nil
}

// Comment details.
type CommentObject struct {
	// Array of actions on the activity allowed for the current viewer.
	// Wire name: 'available_actions'
	AvailableActions []CommentActivityAction
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Comment ID
	// Wire name: 'id'
	Id string
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func commentObjectToPb(st *CommentObject) (*commentObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &commentObjectPb{}
	pb.AvailableActions = st.AvailableActions

	pb.Comment = st.Comment

	pb.CreationTimestamp = st.CreationTimestamp

	pb.Id = st.Id

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CommentObject) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &commentObjectPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := commentObjectFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CommentObject) MarshalJSON() ([]byte, error) {
	pb, err := commentObjectToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type commentObjectPb struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []CommentActivityAction `json:"available_actions,omitempty"`
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Comment ID
	Id string `json:"id,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func commentObjectFromPb(pb *commentObjectPb) (*CommentObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CommentObject{}
	st.AvailableActions = pb.AvailableActions
	st.Comment = pb.Comment
	st.CreationTimestamp = pb.CreationTimestamp
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *commentObjectPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st commentObjectPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateComment struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Version of the model.
	// Wire name: 'version'
	Version string
}

func createCommentToPb(st *CreateComment) (*createCommentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCommentPb{}
	pb.Comment = st.Comment

	pb.Name = st.Name

	pb.Version = st.Version

	return pb, nil
}

func (st *CreateComment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCommentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCommentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateComment) MarshalJSON() ([]byte, error) {
	pb, err := createCommentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createCommentPb struct {
	// User-provided comment on the action.
	Comment string `json:"comment"`
	// Name of the model.
	Name string `json:"name"`
	// Version of the model.
	Version string `json:"version"`
}

func createCommentFromPb(pb *createCommentPb) (*CreateComment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateComment{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

type CreateCommentResponse struct {
	// Comment details.
	// Wire name: 'comment'
	Comment *CommentObject
}

func createCommentResponseToPb(st *CreateCommentResponse) (*createCommentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createCommentResponsePb{}
	commentPb, err := commentObjectToPb(st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = commentPb
	}

	return pb, nil
}

func (st *CreateCommentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createCommentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createCommentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateCommentResponse) MarshalJSON() ([]byte, error) {
	pb, err := createCommentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createCommentResponsePb struct {
	// Comment details.
	Comment *commentObjectPb `json:"comment,omitempty"`
}

func createCommentResponseFromPb(pb *createCommentResponsePb) (*CreateCommentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCommentResponse{}
	commentField, err := commentObjectFromPb(pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = commentField
	}

	return st, nil
}

type CreateExperiment struct {
	// Location where all artifacts for the experiment are stored. If not
	// provided, the remote server will select an appropriate default.
	// Wire name: 'artifact_location'
	ArtifactLocation string
	// Experiment name.
	// Wire name: 'name'
	Name string
	// A collection of tags to set on the experiment. Maximum tag size and
	// number of tags per request depends on the storage backend. All storage
	// backends are guaranteed to support tag keys up to 250 bytes in size and
	// tag values up to 5000 bytes in size. All storage backends are also
	// guaranteed to support up to 20 tags per request.
	// Wire name: 'tags'
	Tags []ExperimentTag

	ForceSendFields []string `tf:"-"`
}

func createExperimentToPb(st *CreateExperiment) (*createExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExperimentPb{}
	pb.ArtifactLocation = st.ArtifactLocation

	pb.Name = st.Name

	var tagsPb []experimentTagPb
	for _, item := range st.Tags {
		itemPb, err := experimentTagToPb(&item)
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

func (st *CreateExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExperiment) MarshalJSON() ([]byte, error) {
	pb, err := createExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createExperimentPb struct {
	// Location where all artifacts for the experiment are stored. If not
	// provided, the remote server will select an appropriate default.
	ArtifactLocation string `json:"artifact_location,omitempty"`
	// Experiment name.
	Name string `json:"name"`
	// A collection of tags to set on the experiment. Maximum tag size and
	// number of tags per request depends on the storage backend. All storage
	// backends are guaranteed to support tag keys up to 250 bytes in size and
	// tag values up to 5000 bytes in size. All storage backends are also
	// guaranteed to support up to 20 tags per request.
	Tags []experimentTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createExperimentFromPb(pb *createExperimentPb) (*CreateExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExperiment{}
	st.ArtifactLocation = pb.ArtifactLocation
	st.Name = pb.Name

	var tagsField []ExperimentTag
	for _, itemPb := range pb.Tags {
		item, err := experimentTagFromPb(&itemPb)
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

func (st *createExperimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createExperimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	// Wire name: 'experiment_id'
	ExperimentId string

	ForceSendFields []string `tf:"-"`
}

func createExperimentResponseToPb(st *CreateExperimentResponse) (*createExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createExperimentResponsePb{}
	pb.ExperimentId = st.ExperimentId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := createExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createExperimentResponsePb struct {
	// Unique identifier for the experiment.
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createExperimentResponseFromPb(pb *createExperimentResponsePb) (*CreateExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExperimentResponse{}
	st.ExperimentId = pb.ExperimentId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createExperimentResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createExperimentResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateForecastingExperimentRequest struct {
	// The column in the training table used to customize weights for each time
	// series.
	// Wire name: 'custom_weights_column'
	CustomWeightsColumn string
	// The path in the workspace to store the created experiment.
	// Wire name: 'experiment_path'
	ExperimentPath string
	// The time interval between consecutive rows in the time series data.
	// Possible values include: '1 second', '1 minute', '5 minutes', '10
	// minutes', '15 minutes', '30 minutes', 'Hourly', 'Daily', 'Weekly',
	// 'Monthly', 'Quarterly', 'Yearly'.
	// Wire name: 'forecast_granularity'
	ForecastGranularity string
	// The number of time steps into the future to make predictions, calculated
	// as a multiple of forecast_granularity. This value represents how far
	// ahead the model should forecast.
	// Wire name: 'forecast_horizon'
	ForecastHorizon int64
	// The region code(s) to automatically add holiday features. Currently
	// supports only one region.
	// Wire name: 'holiday_regions'
	HolidayRegions []string
	// Specifies the list of feature columns to include in model training. These
	// columns must exist in the training data and be of type string, numerical,
	// or boolean. If not specified, no additional features will be included.
	// Note: Certain columns are automatically handled: - Automatically
	// excluded: split_column, target_column, custom_weights_column. -
	// Automatically included: time_column.
	// Wire name: 'include_features'
	IncludeFeatures []string
	// The maximum duration for the experiment in minutes. The experiment stops
	// automatically if it exceeds this limit.
	// Wire name: 'max_runtime'
	MaxRuntime int64
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store predictions.
	// Wire name: 'prediction_data_path'
	PredictionDataPath string
	// The evaluation metric used to optimize the forecasting model.
	// Wire name: 'primary_metric'
	PrimaryMetric string
	// The fully qualified path of a Unity Catalog model, formatted as
	// catalog_name.schema_name.model_name, used to store the best model.
	// Wire name: 'register_to'
	RegisterTo string
	// // The column in the training table used for custom data splits. Values
	// must be 'train', 'validate', or 'test'.
	// Wire name: 'split_column'
	SplitColumn string
	// The column in the input training table used as the prediction target for
	// model training. The values in this column are used as the ground truth
	// for model training.
	// Wire name: 'target_column'
	TargetColumn string
	// The column in the input training table that represents each row's
	// timestamp.
	// Wire name: 'time_column'
	TimeColumn string
	// The column in the training table used to group the dataset for predicting
	// individual time series.
	// Wire name: 'timeseries_identifier_columns'
	TimeseriesIdentifierColumns []string
	// The fully qualified name of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used as training data for the
	// forecasting model.
	// Wire name: 'train_data_path'
	TrainDataPath string
	// List of frameworks to include for model tuning. Possible values are
	// 'Prophet', 'ARIMA', 'DeepAR'. An empty list includes all supported
	// frameworks.
	// Wire name: 'training_frameworks'
	TrainingFrameworks []string

	ForceSendFields []string `tf:"-"`
}

func createForecastingExperimentRequestToPb(st *CreateForecastingExperimentRequest) (*createForecastingExperimentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createForecastingExperimentRequestPb{}
	pb.CustomWeightsColumn = st.CustomWeightsColumn

	pb.ExperimentPath = st.ExperimentPath

	pb.ForecastGranularity = st.ForecastGranularity

	pb.ForecastHorizon = st.ForecastHorizon

	pb.HolidayRegions = st.HolidayRegions

	pb.IncludeFeatures = st.IncludeFeatures

	pb.MaxRuntime = st.MaxRuntime

	pb.PredictionDataPath = st.PredictionDataPath

	pb.PrimaryMetric = st.PrimaryMetric

	pb.RegisterTo = st.RegisterTo

	pb.SplitColumn = st.SplitColumn

	pb.TargetColumn = st.TargetColumn

	pb.TimeColumn = st.TimeColumn

	pb.TimeseriesIdentifierColumns = st.TimeseriesIdentifierColumns

	pb.TrainDataPath = st.TrainDataPath

	pb.TrainingFrameworks = st.TrainingFrameworks

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateForecastingExperimentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createForecastingExperimentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createForecastingExperimentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateForecastingExperimentRequest) MarshalJSON() ([]byte, error) {
	pb, err := createForecastingExperimentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createForecastingExperimentRequestPb struct {
	// The column in the training table used to customize weights for each time
	// series.
	CustomWeightsColumn string `json:"custom_weights_column,omitempty"`
	// The path in the workspace to store the created experiment.
	ExperimentPath string `json:"experiment_path,omitempty"`
	// The time interval between consecutive rows in the time series data.
	// Possible values include: '1 second', '1 minute', '5 minutes', '10
	// minutes', '15 minutes', '30 minutes', 'Hourly', 'Daily', 'Weekly',
	// 'Monthly', 'Quarterly', 'Yearly'.
	ForecastGranularity string `json:"forecast_granularity"`
	// The number of time steps into the future to make predictions, calculated
	// as a multiple of forecast_granularity. This value represents how far
	// ahead the model should forecast.
	ForecastHorizon int64 `json:"forecast_horizon"`
	// The region code(s) to automatically add holiday features. Currently
	// supports only one region.
	HolidayRegions []string `json:"holiday_regions,omitempty"`
	// Specifies the list of feature columns to include in model training. These
	// columns must exist in the training data and be of type string, numerical,
	// or boolean. If not specified, no additional features will be included.
	// Note: Certain columns are automatically handled: - Automatically
	// excluded: split_column, target_column, custom_weights_column. -
	// Automatically included: time_column.
	IncludeFeatures []string `json:"include_features,omitempty"`
	// The maximum duration for the experiment in minutes. The experiment stops
	// automatically if it exceeds this limit.
	MaxRuntime int64 `json:"max_runtime,omitempty"`
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store predictions.
	PredictionDataPath string `json:"prediction_data_path,omitempty"`
	// The evaluation metric used to optimize the forecasting model.
	PrimaryMetric string `json:"primary_metric,omitempty"`
	// The fully qualified path of a Unity Catalog model, formatted as
	// catalog_name.schema_name.model_name, used to store the best model.
	RegisterTo string `json:"register_to,omitempty"`
	// // The column in the training table used for custom data splits. Values
	// must be 'train', 'validate', or 'test'.
	SplitColumn string `json:"split_column,omitempty"`
	// The column in the input training table used as the prediction target for
	// model training. The values in this column are used as the ground truth
	// for model training.
	TargetColumn string `json:"target_column"`
	// The column in the input training table that represents each row's
	// timestamp.
	TimeColumn string `json:"time_column"`
	// The column in the training table used to group the dataset for predicting
	// individual time series.
	TimeseriesIdentifierColumns []string `json:"timeseries_identifier_columns,omitempty"`
	// The fully qualified name of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used as training data for the
	// forecasting model.
	TrainDataPath string `json:"train_data_path"`
	// List of frameworks to include for model tuning. Possible values are
	// 'Prophet', 'ARIMA', 'DeepAR'. An empty list includes all supported
	// frameworks.
	TrainingFrameworks []string `json:"training_frameworks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createForecastingExperimentRequestFromPb(pb *createForecastingExperimentRequestPb) (*CreateForecastingExperimentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateForecastingExperimentRequest{}
	st.CustomWeightsColumn = pb.CustomWeightsColumn
	st.ExperimentPath = pb.ExperimentPath
	st.ForecastGranularity = pb.ForecastGranularity
	st.ForecastHorizon = pb.ForecastHorizon
	st.HolidayRegions = pb.HolidayRegions
	st.IncludeFeatures = pb.IncludeFeatures
	st.MaxRuntime = pb.MaxRuntime
	st.PredictionDataPath = pb.PredictionDataPath
	st.PrimaryMetric = pb.PrimaryMetric
	st.RegisterTo = pb.RegisterTo
	st.SplitColumn = pb.SplitColumn
	st.TargetColumn = pb.TargetColumn
	st.TimeColumn = pb.TimeColumn
	st.TimeseriesIdentifierColumns = pb.TimeseriesIdentifierColumns
	st.TrainDataPath = pb.TrainDataPath
	st.TrainingFrameworks = pb.TrainingFrameworks

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createForecastingExperimentRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createForecastingExperimentRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateForecastingExperimentResponse struct {
	// The unique ID of the created forecasting experiment
	// Wire name: 'experiment_id'
	ExperimentId string

	ForceSendFields []string `tf:"-"`
}

func createForecastingExperimentResponseToPb(st *CreateForecastingExperimentResponse) (*createForecastingExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createForecastingExperimentResponsePb{}
	pb.ExperimentId = st.ExperimentId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateForecastingExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createForecastingExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createForecastingExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateForecastingExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := createForecastingExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createForecastingExperimentResponsePb struct {
	// The unique ID of the created forecasting experiment
	ExperimentId string `json:"experiment_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createForecastingExperimentResponseFromPb(pb *createForecastingExperimentResponsePb) (*CreateForecastingExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateForecastingExperimentResponse{}
	st.ExperimentId = pb.ExperimentId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createForecastingExperimentResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createForecastingExperimentResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateModelRequest struct {
	// Optional description for registered model.
	// Wire name: 'description'
	Description string
	// Register models under this name
	// Wire name: 'name'
	Name string
	// Additional metadata for registered model.
	// Wire name: 'tags'
	Tags []ModelTag

	ForceSendFields []string `tf:"-"`
}

func createModelRequestToPb(st *CreateModelRequest) (*createModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createModelRequestPb{}
	pb.Description = st.Description

	pb.Name = st.Name

	var tagsPb []modelTagPb
	for _, item := range st.Tags {
		itemPb, err := modelTagToPb(&item)
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

func (st *CreateModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := createModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createModelRequestPb struct {
	// Optional description for registered model.
	Description string `json:"description,omitempty"`
	// Register models under this name
	Name string `json:"name"`
	// Additional metadata for registered model.
	Tags []modelTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createModelRequestFromPb(pb *createModelRequestPb) (*CreateModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelRequest{}
	st.Description = pb.Description
	st.Name = pb.Name

	var tagsField []ModelTag
	for _, itemPb := range pb.Tags {
		item, err := modelTagFromPb(&itemPb)
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

func (st *createModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateModelResponse struct {

	// Wire name: 'registered_model'
	RegisteredModel *Model
}

func createModelResponseToPb(st *CreateModelResponse) (*createModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createModelResponsePb{}
	registeredModelPb, err := modelToPb(st.RegisteredModel)
	if err != nil {
		return nil, err
	}
	if registeredModelPb != nil {
		pb.RegisteredModel = registeredModelPb
	}

	return pb, nil
}

func (st *CreateModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := createModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createModelResponsePb struct {
	RegisteredModel *modelPb `json:"registered_model,omitempty"`
}

func createModelResponseFromPb(pb *createModelResponsePb) (*CreateModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelResponse{}
	registeredModelField, err := modelFromPb(pb.RegisteredModel)
	if err != nil {
		return nil, err
	}
	if registeredModelField != nil {
		st.RegisteredModel = registeredModelField
	}

	return st, nil
}

type CreateModelVersionRequest struct {
	// Optional description for model version.
	// Wire name: 'description'
	Description string
	// Register model under this name
	// Wire name: 'name'
	Name string
	// MLflow run ID for correlation, if `source` was generated by an experiment
	// run in MLflow tracking server
	// Wire name: 'run_id'
	RunId string
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	// Wire name: 'run_link'
	RunLink string
	// URI indicating the location of the model artifacts.
	// Wire name: 'source'
	Source string
	// Additional metadata for model version.
	// Wire name: 'tags'
	Tags []ModelVersionTag

	ForceSendFields []string `tf:"-"`
}

func createModelVersionRequestToPb(st *CreateModelVersionRequest) (*createModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createModelVersionRequestPb{}
	pb.Description = st.Description

	pb.Name = st.Name

	pb.RunId = st.RunId

	pb.RunLink = st.RunLink

	pb.Source = st.Source

	var tagsPb []modelVersionTagPb
	for _, item := range st.Tags {
		itemPb, err := modelVersionTagToPb(&item)
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

func (st *CreateModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := createModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createModelVersionRequestPb struct {
	// Optional description for model version.
	Description string `json:"description,omitempty"`
	// Register model under this name
	Name string `json:"name"`
	// MLflow run ID for correlation, if `source` was generated by an experiment
	// run in MLflow tracking server
	RunId string `json:"run_id,omitempty"`
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	RunLink string `json:"run_link,omitempty"`
	// URI indicating the location of the model artifacts.
	Source string `json:"source"`
	// Additional metadata for model version.
	Tags []modelVersionTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createModelVersionRequestFromPb(pb *createModelVersionRequestPb) (*CreateModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelVersionRequest{}
	st.Description = pb.Description
	st.Name = pb.Name
	st.RunId = pb.RunId
	st.RunLink = pb.RunLink
	st.Source = pb.Source

	var tagsField []ModelVersionTag
	for _, itemPb := range pb.Tags {
		item, err := modelVersionTagFromPb(&itemPb)
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

func (st *createModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	// Wire name: 'model_version'
	ModelVersion *ModelVersion
}

func createModelVersionResponseToPb(st *CreateModelVersionResponse) (*createModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createModelVersionResponsePb{}
	modelVersionPb, err := modelVersionToPb(st.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionPb != nil {
		pb.ModelVersion = modelVersionPb
	}

	return pb, nil
}

func (st *CreateModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := createModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createModelVersionResponsePb struct {
	// Return new version number generated for this model in registry.
	ModelVersion *modelVersionPb `json:"model_version,omitempty"`
}

func createModelVersionResponseFromPb(pb *createModelVersionResponsePb) (*CreateModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelVersionResponse{}
	modelVersionField, err := modelVersionFromPb(pb.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionField != nil {
		st.ModelVersion = modelVersionField
	}

	return st, nil
}

type CreateRegistryWebhook struct {
	// User-specified description for the webhook.
	// Wire name: 'description'
	Description string
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	// Wire name: 'events'
	Events []RegistryWebhookEvent

	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpec

	// Wire name: 'job_spec'
	JobSpec *JobSpec
	// Name of the model whose events would trigger this webhook.
	// Wire name: 'model_name'
	ModelName string
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	// Wire name: 'status'
	Status RegistryWebhookStatus

	ForceSendFields []string `tf:"-"`
}

func createRegistryWebhookToPb(st *CreateRegistryWebhook) (*createRegistryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRegistryWebhookPb{}
	pb.Description = st.Description

	pb.Events = st.Events

	httpUrlSpecPb, err := httpUrlSpecToPb(st.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecPb != nil {
		pb.HttpUrlSpec = httpUrlSpecPb
	}

	jobSpecPb, err := jobSpecToPb(st.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecPb != nil {
		pb.JobSpec = jobSpecPb
	}

	pb.ModelName = st.ModelName

	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateRegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRegistryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRegistryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := createRegistryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createRegistryWebhookPb struct {
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events []RegistryWebhookEvent `json:"events"`

	HttpUrlSpec *httpUrlSpecPb `json:"http_url_spec,omitempty"`

	JobSpec *jobSpecPb `json:"job_spec,omitempty"`
	// Name of the model whose events would trigger this webhook.
	ModelName string `json:"model_name,omitempty"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createRegistryWebhookFromPb(pb *createRegistryWebhookPb) (*CreateRegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRegistryWebhook{}
	st.Description = pb.Description
	st.Events = pb.Events
	httpUrlSpecField, err := httpUrlSpecFromPb(pb.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecField != nil {
		st.HttpUrlSpec = httpUrlSpecField
	}
	jobSpecField, err := jobSpecFromPb(pb.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecField != nil {
		st.JobSpec = jobSpecField
	}
	st.ModelName = pb.ModelName
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createRegistryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createRegistryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateRun struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
	// The name of the run.
	// Wire name: 'run_name'
	RunName string
	// Unix timestamp in milliseconds of when the run started.
	// Wire name: 'start_time'
	StartTime int64
	// Additional metadata for run.
	// Wire name: 'tags'
	Tags []RunTag
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func createRunToPb(st *CreateRun) (*createRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRunPb{}
	pb.ExperimentId = st.ExperimentId

	pb.RunName = st.RunName

	pb.StartTime = st.StartTime

	var tagsPb []runTagPb
	for _, item := range st.Tags {
		itemPb, err := runTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRun) MarshalJSON() ([]byte, error) {
	pb, err := createRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createRunPb struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// The name of the run.
	RunName string `json:"run_name,omitempty"`
	// Unix timestamp in milliseconds of when the run started.
	StartTime int64 `json:"start_time,omitempty"`
	// Additional metadata for run.
	Tags []runTagPb `json:"tags,omitempty"`
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createRunFromPb(pb *createRunPb) (*CreateRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRun{}
	st.ExperimentId = pb.ExperimentId
	st.RunName = pb.RunName
	st.StartTime = pb.StartTime

	var tagsField []RunTag
	for _, itemPb := range pb.Tags {
		item, err := runTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateRunResponse struct {
	// The newly created run.
	// Wire name: 'run'
	Run *Run
}

func createRunResponseToPb(st *CreateRunResponse) (*createRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createRunResponsePb{}
	runPb, err := runToPb(st.Run)
	if err != nil {
		return nil, err
	}
	if runPb != nil {
		pb.Run = runPb
	}

	return pb, nil
}

func (st *CreateRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := createRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createRunResponsePb struct {
	// The newly created run.
	Run *runPb `json:"run,omitempty"`
}

func createRunResponseFromPb(pb *createRunResponsePb) (*CreateRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRunResponse{}
	runField, err := runFromPb(pb.Run)
	if err != nil {
		return nil, err
	}
	if runField != nil {
		st.Run = runField
	}

	return st, nil
}

type CreateTransitionRequest struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'stage'
	Stage Stage
	// Version of the model.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func createTransitionRequestToPb(st *CreateTransitionRequest) (*createTransitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createTransitionRequestPb{}
	pb.Comment = st.Comment

	pb.Name = st.Name

	pb.Stage = st.Stage

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *CreateTransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createTransitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createTransitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateTransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := createTransitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createTransitionRequestPb struct {
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage Stage `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func createTransitionRequestFromPb(pb *createTransitionRequestPb) (*CreateTransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTransitionRequest{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *createTransitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st createTransitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type CreateTransitionRequestResponse struct {
	// Transition request details.
	// Wire name: 'request'
	Request *TransitionRequest
}

func createTransitionRequestResponseToPb(st *CreateTransitionRequestResponse) (*createTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createTransitionRequestResponsePb{}
	requestPb, err := transitionRequestToPb(st.Request)
	if err != nil {
		return nil, err
	}
	if requestPb != nil {
		pb.Request = requestPb
	}

	return pb, nil
}

func (st *CreateTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := createTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createTransitionRequestResponsePb struct {
	// Transition request details.
	Request *transitionRequestPb `json:"request,omitempty"`
}

func createTransitionRequestResponseFromPb(pb *createTransitionRequestResponsePb) (*CreateTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTransitionRequestResponse{}
	requestField, err := transitionRequestFromPb(pb.Request)
	if err != nil {
		return nil, err
	}
	if requestField != nil {
		st.Request = requestField
	}

	return st, nil
}

type CreateWebhookResponse struct {

	// Wire name: 'webhook'
	Webhook *RegistryWebhook
}

func createWebhookResponseToPb(st *CreateWebhookResponse) (*createWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &createWebhookResponsePb{}
	webhookPb, err := registryWebhookToPb(st.Webhook)
	if err != nil {
		return nil, err
	}
	if webhookPb != nil {
		pb.Webhook = webhookPb
	}

	return pb, nil
}

func (st *CreateWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &createWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := createWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st CreateWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := createWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type createWebhookResponsePb struct {
	Webhook *registryWebhookPb `json:"webhook,omitempty"`
}

func createWebhookResponseFromPb(pb *createWebhookResponsePb) (*CreateWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWebhookResponse{}
	webhookField, err := registryWebhookFromPb(pb.Webhook)
	if err != nil {
		return nil, err
	}
	if webhookField != nil {
		st.Webhook = webhookField
	}

	return st, nil
}

// Dataset. Represents a reference to data used for training, testing, or
// evaluation during the model development process.
type Dataset struct {
	// Dataset digest, e.g. an md5 hash of the dataset that uniquely identifies
	// it within datasets of the same name.
	// Wire name: 'digest'
	Digest string
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	// Wire name: 'name'
	Name string
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	// Wire name: 'profile'
	Profile string
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	// Wire name: 'schema'
	Schema string
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	// Wire name: 'source'
	Source string
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	// Wire name: 'source_type'
	SourceType string

	ForceSendFields []string `tf:"-"`
}

func datasetToPb(st *Dataset) (*datasetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &datasetPb{}
	pb.Digest = st.Digest

	pb.Name = st.Name

	pb.Profile = st.Profile

	pb.Schema = st.Schema

	pb.Source = st.Source

	pb.SourceType = st.SourceType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Dataset) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &datasetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := datasetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Dataset) MarshalJSON() ([]byte, error) {
	pb, err := datasetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type datasetPb struct {
	// Dataset digest, e.g. an md5 hash of the dataset that uniquely identifies
	// it within datasets of the same name.
	Digest string `json:"digest"`
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	Name string `json:"name"`
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	Profile string `json:"profile,omitempty"`
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	Schema string `json:"schema,omitempty"`
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	Source string `json:"source"`
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	SourceType string `json:"source_type"`

	ForceSendFields []string `json:"-" url:"-"`
}

func datasetFromPb(pb *datasetPb) (*Dataset, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Dataset{}
	st.Digest = pb.Digest
	st.Name = pb.Name
	st.Profile = pb.Profile
	st.Schema = pb.Schema
	st.Source = pb.Source
	st.SourceType = pb.SourceType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *datasetPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st datasetPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// DatasetInput. Represents a dataset and input tags.
type DatasetInput struct {
	// The dataset being used as a Run input.
	// Wire name: 'dataset'
	Dataset Dataset
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	// Wire name: 'tags'
	Tags []InputTag
}

func datasetInputToPb(st *DatasetInput) (*datasetInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &datasetInputPb{}
	datasetPb, err := datasetToPb(&st.Dataset)
	if err != nil {
		return nil, err
	}
	if datasetPb != nil {
		pb.Dataset = *datasetPb
	}

	var tagsPb []inputTagPb
	for _, item := range st.Tags {
		itemPb, err := inputTagToPb(&item)
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

func (st *DatasetInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &datasetInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := datasetInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DatasetInput) MarshalJSON() ([]byte, error) {
	pb, err := datasetInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type datasetInputPb struct {
	// The dataset being used as a Run input.
	Dataset datasetPb `json:"dataset"`
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	Tags []inputTagPb `json:"tags,omitempty"`
}

func datasetInputFromPb(pb *datasetInputPb) (*DatasetInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatasetInput{}
	datasetField, err := datasetFromPb(&pb.Dataset)
	if err != nil {
		return nil, err
	}
	if datasetField != nil {
		st.Dataset = *datasetField
	}

	var tagsField []InputTag
	for _, itemPb := range pb.Tags {
		item, err := inputTagFromPb(&itemPb)
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

// Delete a comment
type DeleteCommentRequest struct {

	// Wire name: 'id'
	Id string `tf:"-"`
}

func deleteCommentRequestToPb(st *DeleteCommentRequest) (*deleteCommentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCommentRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func (st *DeleteCommentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCommentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCommentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCommentRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteCommentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCommentRequestPb struct {
	Id string `json:"-" url:"id"`
}

func deleteCommentRequestFromPb(pb *deleteCommentRequestPb) (*DeleteCommentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCommentRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteCommentResponse struct {
}

func deleteCommentResponseToPb(st *DeleteCommentResponse) (*deleteCommentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteCommentResponsePb{}

	return pb, nil
}

func (st *DeleteCommentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteCommentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteCommentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteCommentResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteCommentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteCommentResponsePb struct {
}

func deleteCommentResponseFromPb(pb *deleteCommentResponsePb) (*DeleteCommentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCommentResponse{}

	return st, nil
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
}

func deleteExperimentToPb(st *DeleteExperiment) (*deleteExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExperimentPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func (st *DeleteExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExperiment) MarshalJSON() ([]byte, error) {
	pb, err := deleteExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteExperimentPb struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
}

func deleteExperimentFromPb(pb *deleteExperimentPb) (*DeleteExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExperiment{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

type DeleteExperimentResponse struct {
}

func deleteExperimentResponseToPb(st *DeleteExperimentResponse) (*deleteExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteExperimentResponsePb{}

	return pb, nil
}

func (st *DeleteExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteExperimentResponsePb struct {
}

func deleteExperimentResponseFromPb(pb *deleteExperimentResponsePb) (*DeleteExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExperimentResponse{}

	return st, nil
}

// Delete a model
type DeleteModelRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func deleteModelRequestToPb(st *DeleteModelRequest) (*deleteModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *DeleteModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteModelRequestPb struct {
	// Registered model unique name identifier.
	Name string `json:"-" url:"name"`
}

func deleteModelRequestFromPb(pb *deleteModelRequestPb) (*DeleteModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteModelResponse struct {
}

func deleteModelResponseToPb(st *DeleteModelResponse) (*deleteModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelResponsePb{}

	return pb, nil
}

func (st *DeleteModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteModelResponsePb struct {
}

func deleteModelResponseFromPb(pb *deleteModelResponsePb) (*DeleteModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelResponse{}

	return st, nil
}

// Delete a model tag
type DeleteModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	// Wire name: 'key'
	Key string `tf:"-"`
	// Name of the registered model that the tag was logged under.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func deleteModelTagRequestToPb(st *DeleteModelTagRequest) (*deleteModelTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelTagRequestPb{}
	pb.Key = st.Key

	pb.Name = st.Name

	return pb, nil
}

func (st *DeleteModelTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteModelTagRequestPb struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key string `json:"-" url:"key"`
	// Name of the registered model that the tag was logged under.
	Name string `json:"-" url:"name"`
}

func deleteModelTagRequestFromPb(pb *deleteModelTagRequestPb) (*DeleteModelTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name

	return st, nil
}

type DeleteModelTagResponse struct {
}

func deleteModelTagResponseToPb(st *DeleteModelTagResponse) (*deleteModelTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelTagResponsePb{}

	return pb, nil
}

func (st *DeleteModelTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteModelTagResponsePb struct {
}

func deleteModelTagResponseFromPb(pb *deleteModelTagResponsePb) (*DeleteModelTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelTagResponse{}

	return st, nil
}

// Delete a model version.
type DeleteModelVersionRequest struct {
	// Name of the registered model
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number
	// Wire name: 'version'
	Version string `tf:"-"`
}

func deleteModelVersionRequestToPb(st *DeleteModelVersionRequest) (*deleteModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionRequestPb{}
	pb.Name = st.Name

	pb.Version = st.Version

	return pb, nil
}

func (st *DeleteModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteModelVersionRequestPb struct {
	// Name of the registered model
	Name string `json:"-" url:"name"`
	// Model version number
	Version string `json:"-" url:"version"`
}

func deleteModelVersionRequestFromPb(pb *deleteModelVersionRequestPb) (*DeleteModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionRequest{}
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

type DeleteModelVersionResponse struct {
}

func deleteModelVersionResponseToPb(st *DeleteModelVersionResponse) (*deleteModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionResponsePb{}

	return pb, nil
}

func (st *DeleteModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteModelVersionResponsePb struct {
}

func deleteModelVersionResponseFromPb(pb *deleteModelVersionResponsePb) (*DeleteModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionResponse{}

	return st, nil
}

// Delete a model version tag
type DeleteModelVersionTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	// Wire name: 'key'
	Key string `tf:"-"`
	// Name of the registered model that the tag was logged under.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number that the tag was logged under.
	// Wire name: 'version'
	Version string `tf:"-"`
}

func deleteModelVersionTagRequestToPb(st *DeleteModelVersionTagRequest) (*deleteModelVersionTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionTagRequestPb{}
	pb.Key = st.Key

	pb.Name = st.Name

	pb.Version = st.Version

	return pb, nil
}

func (st *DeleteModelVersionTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelVersionTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelVersionTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelVersionTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelVersionTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteModelVersionTagRequestPb struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	Key string `json:"-" url:"key"`
	// Name of the registered model that the tag was logged under.
	Name string `json:"-" url:"name"`
	// Model version number that the tag was logged under.
	Version string `json:"-" url:"version"`
}

func deleteModelVersionTagRequestFromPb(pb *deleteModelVersionTagRequestPb) (*DeleteModelVersionTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

type DeleteModelVersionTagResponse struct {
}

func deleteModelVersionTagResponseToPb(st *DeleteModelVersionTagResponse) (*deleteModelVersionTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteModelVersionTagResponsePb{}

	return pb, nil
}

func (st *DeleteModelVersionTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteModelVersionTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteModelVersionTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteModelVersionTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteModelVersionTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteModelVersionTagResponsePb struct {
}

func deleteModelVersionTagResponseFromPb(pb *deleteModelVersionTagResponsePb) (*DeleteModelVersionTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionTagResponse{}

	return st, nil
}

type DeleteRun struct {
	// ID of the run to delete.
	// Wire name: 'run_id'
	RunId string
}

func deleteRunToPb(st *DeleteRun) (*deleteRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunPb{}
	pb.RunId = st.RunId

	return pb, nil
}

func (st *DeleteRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRun) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteRunPb struct {
	// ID of the run to delete.
	RunId string `json:"run_id"`
}

func deleteRunFromPb(pb *deleteRunPb) (*DeleteRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRun{}
	st.RunId = pb.RunId

	return st, nil
}

type DeleteRunResponse struct {
}

func deleteRunResponseToPb(st *DeleteRunResponse) (*deleteRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunResponsePb{}

	return pb, nil
}

func (st *DeleteRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteRunResponsePb struct {
}

func deleteRunResponseFromPb(pb *deleteRunResponsePb) (*DeleteRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRunResponse{}

	return st, nil
}

type DeleteRuns struct {
	// The ID of the experiment containing the runs to delete.
	// Wire name: 'experiment_id'
	ExperimentId string
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	// Wire name: 'max_runs'
	MaxRuns int
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	// Wire name: 'max_timestamp_millis'
	MaxTimestampMillis int64

	ForceSendFields []string `tf:"-"`
}

func deleteRunsToPb(st *DeleteRuns) (*deleteRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunsPb{}
	pb.ExperimentId = st.ExperimentId

	pb.MaxRuns = st.MaxRuns

	pb.MaxTimestampMillis = st.MaxTimestampMillis

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRuns) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteRunsPb struct {
	// The ID of the experiment containing the runs to delete.
	ExperimentId string `json:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	MaxRuns int `json:"max_runs,omitempty"`
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	MaxTimestampMillis int64 `json:"max_timestamp_millis"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteRunsFromPb(pb *deleteRunsPb) (*DeleteRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRuns{}
	st.ExperimentId = pb.ExperimentId
	st.MaxRuns = pb.MaxRuns
	st.MaxTimestampMillis = pb.MaxTimestampMillis

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteRunsResponse struct {
	// The number of runs deleted.
	// Wire name: 'runs_deleted'
	RunsDeleted int

	ForceSendFields []string `tf:"-"`
}

func deleteRunsResponseToPb(st *DeleteRunsResponse) (*deleteRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteRunsResponsePb{}
	pb.RunsDeleted = st.RunsDeleted

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteRunsResponsePb struct {
	// The number of runs deleted.
	RunsDeleted int `json:"runs_deleted,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteRunsResponseFromPb(pb *deleteRunsResponsePb) (*DeleteRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRunsResponse{}
	st.RunsDeleted = pb.RunsDeleted

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteTag struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	// Wire name: 'key'
	Key string
	// ID of the run that the tag was logged under. Must be provided.
	// Wire name: 'run_id'
	RunId string
}

func deleteTagToPb(st *DeleteTag) (*deleteTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTagPb{}
	pb.Key = st.Key

	pb.RunId = st.RunId

	return pb, nil
}

func (st *DeleteTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTag) MarshalJSON() ([]byte, error) {
	pb, err := deleteTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteTagPb struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	Key string `json:"key"`
	// ID of the run that the tag was logged under. Must be provided.
	RunId string `json:"run_id"`
}

func deleteTagFromPb(pb *deleteTagPb) (*DeleteTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTag{}
	st.Key = pb.Key
	st.RunId = pb.RunId

	return st, nil
}

type DeleteTagResponse struct {
}

func deleteTagResponseToPb(st *DeleteTagResponse) (*deleteTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTagResponsePb{}

	return pb, nil
}

func (st *DeleteTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteTagResponsePb struct {
}

func deleteTagResponseFromPb(pb *deleteTagResponsePb) (*DeleteTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTagResponse{}

	return st, nil
}

// Delete a transition request
type DeleteTransitionRequestRequest struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string `tf:"-"`
	// Username of the user who created this request. Of the transition requests
	// matching the specified details, only the one transition created by this
	// user will be deleted.
	// Wire name: 'creator'
	Creator string `tf:"-"`
	// Name of the model.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Target stage of the transition request. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'stage'
	Stage DeleteTransitionRequestStage `tf:"-"`
	// Version of the model.
	// Wire name: 'version'
	Version string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteTransitionRequestRequestToPb(st *DeleteTransitionRequestRequest) (*deleteTransitionRequestRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTransitionRequestRequestPb{}
	pb.Comment = st.Comment

	pb.Creator = st.Creator

	pb.Name = st.Name

	pb.Stage = st.Stage

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteTransitionRequestRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTransitionRequestRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTransitionRequestRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTransitionRequestRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteTransitionRequestRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteTransitionRequestRequestPb struct {
	// User-provided comment on the action.
	Comment string `json:"-" url:"comment,omitempty"`
	// Username of the user who created this request. Of the transition requests
	// matching the specified details, only the one transition created by this
	// user will be deleted.
	Creator string `json:"-" url:"creator"`
	// Name of the model.
	Name string `json:"-" url:"name"`
	// Target stage of the transition request. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage DeleteTransitionRequestStage `json:"-" url:"stage"`
	// Version of the model.
	Version string `json:"-" url:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteTransitionRequestRequestFromPb(pb *deleteTransitionRequestRequestPb) (*DeleteTransitionRequestRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTransitionRequestRequest{}
	st.Comment = pb.Comment
	st.Creator = pb.Creator
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteTransitionRequestRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteTransitionRequestRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteTransitionRequestResponse struct {
}

func deleteTransitionRequestResponseToPb(st *DeleteTransitionRequestResponse) (*deleteTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteTransitionRequestResponsePb{}

	return pb, nil
}

func (st *DeleteTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteTransitionRequestResponsePb struct {
}

func deleteTransitionRequestResponseFromPb(pb *deleteTransitionRequestResponsePb) (*DeleteTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTransitionRequestResponse{}

	return st, nil
}

type DeleteTransitionRequestStage string
type deleteTransitionRequestStagePb string

const DeleteTransitionRequestStageArchived DeleteTransitionRequestStage = `Archived`

const DeleteTransitionRequestStageNone DeleteTransitionRequestStage = `None`

const DeleteTransitionRequestStageProduction DeleteTransitionRequestStage = `Production`

const DeleteTransitionRequestStageStaging DeleteTransitionRequestStage = `Staging`

// String representation for [fmt.Print]
func (f *DeleteTransitionRequestStage) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *DeleteTransitionRequestStage) Set(v string) error {
	switch v {
	case `Archived`, `None`, `Production`, `Staging`:
		*f = DeleteTransitionRequestStage(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Archived", "None", "Production", "Staging"`, v)
	}
}

// Type always returns DeleteTransitionRequestStage to satisfy [pflag.Value] interface
func (f *DeleteTransitionRequestStage) Type() string {
	return "DeleteTransitionRequestStage"
}

func deleteTransitionRequestStageToPb(st *DeleteTransitionRequestStage) (*deleteTransitionRequestStagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := deleteTransitionRequestStagePb(*st)
	return &pb, nil
}

func deleteTransitionRequestStageFromPb(pb *deleteTransitionRequestStagePb) (*DeleteTransitionRequestStage, error) {
	if pb == nil {
		return nil, nil
	}
	st := DeleteTransitionRequestStage(*pb)
	return &st, nil
}

// Delete a webhook
type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	// Wire name: 'id'
	Id string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func deleteWebhookRequestToPb(st *DeleteWebhookRequest) (*deleteWebhookRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWebhookRequestPb{}
	pb.Id = st.Id

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *DeleteWebhookRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteWebhookRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteWebhookRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteWebhookRequest) MarshalJSON() ([]byte, error) {
	pb, err := deleteWebhookRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteWebhookRequestPb struct {
	// Webhook ID required to delete a registry webhook.
	Id string `json:"-" url:"id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func deleteWebhookRequestFromPb(pb *deleteWebhookRequestPb) (*DeleteWebhookRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWebhookRequest{}
	st.Id = pb.Id

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *deleteWebhookRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st deleteWebhookRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type DeleteWebhookResponse struct {
}

func deleteWebhookResponseToPb(st *DeleteWebhookResponse) (*deleteWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &deleteWebhookResponsePb{}

	return pb, nil
}

func (st *DeleteWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &deleteWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := deleteWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st DeleteWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := deleteWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type deleteWebhookResponsePb struct {
}

func deleteWebhookResponseFromPb(pb *deleteWebhookResponsePb) (*DeleteWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWebhookResponse{}

	return st, nil
}

// An experiment and its metadata.
type Experiment struct {
	// Location where artifacts for the experiment are stored.
	// Wire name: 'artifact_location'
	ArtifactLocation string
	// Creation time
	// Wire name: 'creation_time'
	CreationTime int64
	// Unique identifier for the experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
	// Last update time
	// Wire name: 'last_update_time'
	LastUpdateTime int64
	// Current life cycle stage of the experiment: "active" or "deleted".
	// Deleted experiments are not returned by APIs.
	// Wire name: 'lifecycle_stage'
	LifecycleStage string
	// Human readable name that identifies the experiment.
	// Wire name: 'name'
	Name string
	// Tags: Additional metadata key-value pairs.
	// Wire name: 'tags'
	Tags []ExperimentTag

	ForceSendFields []string `tf:"-"`
}

func experimentToPb(st *Experiment) (*experimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentPb{}
	pb.ArtifactLocation = st.ArtifactLocation

	pb.CreationTime = st.CreationTime

	pb.ExperimentId = st.ExperimentId

	pb.LastUpdateTime = st.LastUpdateTime

	pb.LifecycleStage = st.LifecycleStage

	pb.Name = st.Name

	var tagsPb []experimentTagPb
	for _, item := range st.Tags {
		itemPb, err := experimentTagToPb(&item)
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

func (st *Experiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Experiment) MarshalJSON() ([]byte, error) {
	pb, err := experimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type experimentPb struct {
	// Location where artifacts for the experiment are stored.
	ArtifactLocation string `json:"artifact_location,omitempty"`
	// Creation time
	CreationTime int64 `json:"creation_time,omitempty"`
	// Unique identifier for the experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Last update time
	LastUpdateTime int64 `json:"last_update_time,omitempty"`
	// Current life cycle stage of the experiment: "active" or "deleted".
	// Deleted experiments are not returned by APIs.
	LifecycleStage string `json:"lifecycle_stage,omitempty"`
	// Human readable name that identifies the experiment.
	Name string `json:"name,omitempty"`
	// Tags: Additional metadata key-value pairs.
	Tags []experimentTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentFromPb(pb *experimentPb) (*Experiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Experiment{}
	st.ArtifactLocation = pb.ArtifactLocation
	st.CreationTime = pb.CreationTime
	st.ExperimentId = pb.ExperimentId
	st.LastUpdateTime = pb.LastUpdateTime
	st.LifecycleStage = pb.LifecycleStage
	st.Name = pb.Name

	var tagsField []ExperimentTag
	for _, itemPb := range pb.Tags {
		item, err := experimentTagFromPb(&itemPb)
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

func (st *experimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func experimentAccessControlRequestToPb(st *ExperimentAccessControlRequest) (*experimentAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentAccessControlRequestPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ExperimentAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := experimentAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type experimentAccessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentAccessControlRequestFromPb(pb *experimentAccessControlRequestPb) (*ExperimentAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []ExperimentPermission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func experimentAccessControlResponseToPb(st *ExperimentAccessControlResponse) (*experimentAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentAccessControlResponsePb{}

	var allPermissionsPb []experimentPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := experimentPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb

	pb.DisplayName = st.DisplayName

	pb.GroupName = st.GroupName

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ExperimentAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := experimentAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type experimentAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []experimentPermissionPb `json:"all_permissions,omitempty"`
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

func experimentAccessControlResponseFromPb(pb *experimentAccessControlResponsePb) (*ExperimentAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentAccessControlResponse{}

	var allPermissionsField []ExperimentPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := experimentPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentPermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func experimentPermissionToPb(st *ExperimentPermission) (*experimentPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentPermissionPb{}
	pb.Inherited = st.Inherited

	pb.InheritedFromObject = st.InheritedFromObject

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ExperimentPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentPermission) MarshalJSON() ([]byte, error) {
	pb, err := experimentPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type experimentPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentPermissionFromPb(pb *experimentPermissionPb) (*ExperimentPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type ExperimentPermissionLevel string
type experimentPermissionLevelPb string

const ExperimentPermissionLevelCanEdit ExperimentPermissionLevel = `CAN_EDIT`

const ExperimentPermissionLevelCanManage ExperimentPermissionLevel = `CAN_MANAGE`

const ExperimentPermissionLevelCanRead ExperimentPermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *ExperimentPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ExperimentPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_READ`:
		*f = ExperimentPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_READ"`, v)
	}
}

// Type always returns ExperimentPermissionLevel to satisfy [pflag.Value] interface
func (f *ExperimentPermissionLevel) Type() string {
	return "ExperimentPermissionLevel"
}

func experimentPermissionLevelToPb(st *ExperimentPermissionLevel) (*experimentPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := experimentPermissionLevelPb(*st)
	return &pb, nil
}

func experimentPermissionLevelFromPb(pb *experimentPermissionLevelPb) (*ExperimentPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExperimentPermissionLevel(*pb)
	return &st, nil
}

type ExperimentPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []ExperimentAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func experimentPermissionsToPb(st *ExperimentPermissions) (*experimentPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentPermissionsPb{}

	var accessControlListPb []experimentAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := experimentAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	pb.ObjectId = st.ObjectId

	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ExperimentPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentPermissions) MarshalJSON() ([]byte, error) {
	pb, err := experimentPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type experimentPermissionsPb struct {
	AccessControlList []experimentAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentPermissionsFromPb(pb *experimentPermissionsPb) (*ExperimentPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermissions{}

	var accessControlListField []ExperimentAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := experimentAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentPermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func experimentPermissionsDescriptionToPb(st *ExperimentPermissionsDescription) (*experimentPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentPermissionsDescriptionPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ExperimentPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := experimentPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type experimentPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel ExperimentPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentPermissionsDescriptionFromPb(pb *experimentPermissionsDescriptionPb) (*ExperimentPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ExperimentPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []ExperimentAccessControlRequest
	// The experiment for which to get or manage permissions.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func experimentPermissionsRequestToPb(st *ExperimentPermissionsRequest) (*experimentPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentPermissionsRequestPb{}

	var accessControlListPb []experimentAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := experimentAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func (st *ExperimentPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := experimentPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type experimentPermissionsRequestPb struct {
	AccessControlList []experimentAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" url:"-"`
}

func experimentPermissionsRequestFromPb(pb *experimentPermissionsRequestPb) (*ExperimentPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermissionsRequest{}

	var accessControlListField []ExperimentAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := experimentAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

// A tag for an experiment.
type ExperimentTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func experimentTagToPb(st *ExperimentTag) (*experimentTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &experimentTagPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ExperimentTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &experimentTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := experimentTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ExperimentTag) MarshalJSON() ([]byte, error) {
	pb, err := experimentTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type experimentTagPb struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func experimentTagFromPb(pb *experimentTagPb) (*ExperimentTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *experimentTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st experimentTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Metadata of a single artifact file or directory.
type FileInfo struct {
	// The size in bytes of the file. Unset for directories.
	// Wire name: 'file_size'
	FileSize int64
	// Whether the path is a directory.
	// Wire name: 'is_dir'
	IsDir bool
	// The path relative to the root artifact directory run.
	// Wire name: 'path'
	Path string

	ForceSendFields []string `tf:"-"`
}

func fileInfoToPb(st *FileInfo) (*fileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &fileInfoPb{}
	pb.FileSize = st.FileSize

	pb.IsDir = st.IsDir

	pb.Path = st.Path

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *FileInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &fileInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := fileInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st FileInfo) MarshalJSON() ([]byte, error) {
	pb, err := fileInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type fileInfoPb struct {
	// The size in bytes of the file. Unset for directories.
	FileSize int64 `json:"file_size,omitempty"`
	// Whether the path is a directory.
	IsDir bool `json:"is_dir,omitempty"`
	// The path relative to the root artifact directory run.
	Path string `json:"path,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func fileInfoFromPb(pb *fileInfoPb) (*FileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileInfo{}
	st.FileSize = pb.FileSize
	st.IsDir = pb.IsDir
	st.Path = pb.Path

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *fileInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st fileInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Represents a forecasting experiment with its unique identifier, URL, and
// state.
type ForecastingExperiment struct {
	// The unique ID for the forecasting experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
	// The URL to the forecasting experiment page.
	// Wire name: 'experiment_page_url'
	ExperimentPageUrl string
	// The current state of the forecasting experiment.
	// Wire name: 'state'
	State ForecastingExperimentState

	ForceSendFields []string `tf:"-"`
}

func forecastingExperimentToPb(st *ForecastingExperiment) (*forecastingExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &forecastingExperimentPb{}
	pb.ExperimentId = st.ExperimentId

	pb.ExperimentPageUrl = st.ExperimentPageUrl

	pb.State = st.State

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ForecastingExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &forecastingExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := forecastingExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ForecastingExperiment) MarshalJSON() ([]byte, error) {
	pb, err := forecastingExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type forecastingExperimentPb struct {
	// The unique ID for the forecasting experiment.
	ExperimentId string `json:"experiment_id,omitempty"`
	// The URL to the forecasting experiment page.
	ExperimentPageUrl string `json:"experiment_page_url,omitempty"`
	// The current state of the forecasting experiment.
	State ForecastingExperimentState `json:"state,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func forecastingExperimentFromPb(pb *forecastingExperimentPb) (*ForecastingExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForecastingExperiment{}
	st.ExperimentId = pb.ExperimentId
	st.ExperimentPageUrl = pb.ExperimentPageUrl
	st.State = pb.State

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *forecastingExperimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st forecastingExperimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ForecastingExperimentState string
type forecastingExperimentStatePb string

const ForecastingExperimentStateCancelled ForecastingExperimentState = `CANCELLED`

const ForecastingExperimentStateFailed ForecastingExperimentState = `FAILED`

const ForecastingExperimentStatePending ForecastingExperimentState = `PENDING`

const ForecastingExperimentStateRunning ForecastingExperimentState = `RUNNING`

const ForecastingExperimentStateSucceeded ForecastingExperimentState = `SUCCEEDED`

// String representation for [fmt.Print]
func (f *ForecastingExperimentState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ForecastingExperimentState) Set(v string) error {
	switch v {
	case `CANCELLED`, `FAILED`, `PENDING`, `RUNNING`, `SUCCEEDED`:
		*f = ForecastingExperimentState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CANCELLED", "FAILED", "PENDING", "RUNNING", "SUCCEEDED"`, v)
	}
}

// Type always returns ForecastingExperimentState to satisfy [pflag.Value] interface
func (f *ForecastingExperimentState) Type() string {
	return "ForecastingExperimentState"
}

func forecastingExperimentStateToPb(st *ForecastingExperimentState) (*forecastingExperimentStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := forecastingExperimentStatePb(*st)
	return &pb, nil
}

func forecastingExperimentStateFromPb(pb *forecastingExperimentStatePb) (*ForecastingExperimentState, error) {
	if pb == nil {
		return nil, nil
	}
	st := ForecastingExperimentState(*pb)
	return &st, nil
}

// Get an experiment by name
type GetByNameRequest struct {
	// Name of the associated experiment.
	// Wire name: 'experiment_name'
	ExperimentName string `tf:"-"`
}

func getByNameRequestToPb(st *GetByNameRequest) (*getByNameRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getByNameRequestPb{}
	pb.ExperimentName = st.ExperimentName

	return pb, nil
}

func (st *GetByNameRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getByNameRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getByNameRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetByNameRequest) MarshalJSON() ([]byte, error) {
	pb, err := getByNameRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getByNameRequestPb struct {
	// Name of the associated experiment.
	ExperimentName string `json:"-" url:"experiment_name"`
}

func getByNameRequestFromPb(pb *getByNameRequestPb) (*GetByNameRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetByNameRequest{}
	st.ExperimentName = pb.ExperimentName

	return st, nil
}

// Get credentials to download trace data
type GetCredentialsForTraceDataDownloadRequest struct {
	// The ID of the trace to fetch artifact download credentials for.
	// Wire name: 'request_id'
	RequestId string `tf:"-"`
}

func getCredentialsForTraceDataDownloadRequestToPb(st *GetCredentialsForTraceDataDownloadRequest) (*getCredentialsForTraceDataDownloadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialsForTraceDataDownloadRequestPb{}
	pb.RequestId = st.RequestId

	return pb, nil
}

func (st *GetCredentialsForTraceDataDownloadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialsForTraceDataDownloadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialsForTraceDataDownloadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialsForTraceDataDownloadRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialsForTraceDataDownloadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCredentialsForTraceDataDownloadRequestPb struct {
	// The ID of the trace to fetch artifact download credentials for.
	RequestId string `json:"-" url:"-"`
}

func getCredentialsForTraceDataDownloadRequestFromPb(pb *getCredentialsForTraceDataDownloadRequestPb) (*GetCredentialsForTraceDataDownloadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialsForTraceDataDownloadRequest{}
	st.RequestId = pb.RequestId

	return st, nil
}

type GetCredentialsForTraceDataDownloadResponse struct {
	// The artifact download credentials for the specified trace data.
	// Wire name: 'credential_info'
	CredentialInfo *ArtifactCredentialInfo
}

func getCredentialsForTraceDataDownloadResponseToPb(st *GetCredentialsForTraceDataDownloadResponse) (*getCredentialsForTraceDataDownloadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialsForTraceDataDownloadResponsePb{}
	credentialInfoPb, err := artifactCredentialInfoToPb(st.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoPb != nil {
		pb.CredentialInfo = credentialInfoPb
	}

	return pb, nil
}

func (st *GetCredentialsForTraceDataDownloadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialsForTraceDataDownloadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialsForTraceDataDownloadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialsForTraceDataDownloadResponse) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialsForTraceDataDownloadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCredentialsForTraceDataDownloadResponsePb struct {
	// The artifact download credentials for the specified trace data.
	CredentialInfo *artifactCredentialInfoPb `json:"credential_info,omitempty"`
}

func getCredentialsForTraceDataDownloadResponseFromPb(pb *getCredentialsForTraceDataDownloadResponsePb) (*GetCredentialsForTraceDataDownloadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialsForTraceDataDownloadResponse{}
	credentialInfoField, err := artifactCredentialInfoFromPb(pb.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoField != nil {
		st.CredentialInfo = credentialInfoField
	}

	return st, nil
}

// Get credentials to upload trace data
type GetCredentialsForTraceDataUploadRequest struct {
	// The ID of the trace to fetch artifact upload credentials for.
	// Wire name: 'request_id'
	RequestId string `tf:"-"`
}

func getCredentialsForTraceDataUploadRequestToPb(st *GetCredentialsForTraceDataUploadRequest) (*getCredentialsForTraceDataUploadRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialsForTraceDataUploadRequestPb{}
	pb.RequestId = st.RequestId

	return pb, nil
}

func (st *GetCredentialsForTraceDataUploadRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialsForTraceDataUploadRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialsForTraceDataUploadRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialsForTraceDataUploadRequest) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialsForTraceDataUploadRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCredentialsForTraceDataUploadRequestPb struct {
	// The ID of the trace to fetch artifact upload credentials for.
	RequestId string `json:"-" url:"-"`
}

func getCredentialsForTraceDataUploadRequestFromPb(pb *getCredentialsForTraceDataUploadRequestPb) (*GetCredentialsForTraceDataUploadRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialsForTraceDataUploadRequest{}
	st.RequestId = pb.RequestId

	return st, nil
}

type GetCredentialsForTraceDataUploadResponse struct {
	// The artifact upload credentials for the specified trace data.
	// Wire name: 'credential_info'
	CredentialInfo *ArtifactCredentialInfo
}

func getCredentialsForTraceDataUploadResponseToPb(st *GetCredentialsForTraceDataUploadResponse) (*getCredentialsForTraceDataUploadResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getCredentialsForTraceDataUploadResponsePb{}
	credentialInfoPb, err := artifactCredentialInfoToPb(st.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoPb != nil {
		pb.CredentialInfo = credentialInfoPb
	}

	return pb, nil
}

func (st *GetCredentialsForTraceDataUploadResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getCredentialsForTraceDataUploadResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getCredentialsForTraceDataUploadResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetCredentialsForTraceDataUploadResponse) MarshalJSON() ([]byte, error) {
	pb, err := getCredentialsForTraceDataUploadResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getCredentialsForTraceDataUploadResponsePb struct {
	// The artifact upload credentials for the specified trace data.
	CredentialInfo *artifactCredentialInfoPb `json:"credential_info,omitempty"`
}

func getCredentialsForTraceDataUploadResponseFromPb(pb *getCredentialsForTraceDataUploadResponsePb) (*GetCredentialsForTraceDataUploadResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetCredentialsForTraceDataUploadResponse{}
	credentialInfoField, err := artifactCredentialInfoFromPb(pb.CredentialInfo)
	if err != nil {
		return nil, err
	}
	if credentialInfoField != nil {
		st.CredentialInfo = credentialInfoField
	}

	return st, nil
}

type GetExperimentByNameResponse struct {
	// Experiment details.
	// Wire name: 'experiment'
	Experiment *Experiment
}

func getExperimentByNameResponseToPb(st *GetExperimentByNameResponse) (*getExperimentByNameResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentByNameResponsePb{}
	experimentPb, err := experimentToPb(st.Experiment)
	if err != nil {
		return nil, err
	}
	if experimentPb != nil {
		pb.Experiment = experimentPb
	}

	return pb, nil
}

func (st *GetExperimentByNameResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentByNameResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentByNameResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentByNameResponse) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentByNameResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getExperimentByNameResponsePb struct {
	// Experiment details.
	Experiment *experimentPb `json:"experiment,omitempty"`
}

func getExperimentByNameResponseFromPb(pb *getExperimentByNameResponsePb) (*GetExperimentByNameResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentByNameResponse{}
	experimentField, err := experimentFromPb(pb.Experiment)
	if err != nil {
		return nil, err
	}
	if experimentField != nil {
		st.Experiment = experimentField
	}

	return st, nil
}

// Get experiment permission levels
type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func getExperimentPermissionLevelsRequestToPb(st *GetExperimentPermissionLevelsRequest) (*getExperimentPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentPermissionLevelsRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func (st *GetExperimentPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getExperimentPermissionLevelsRequestPb struct {
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" url:"-"`
}

func getExperimentPermissionLevelsRequestFromPb(pb *getExperimentPermissionLevelsRequestPb) (*GetExperimentPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentPermissionLevelsRequest{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

type GetExperimentPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []ExperimentPermissionsDescription
}

func getExperimentPermissionLevelsResponseToPb(st *GetExperimentPermissionLevelsResponse) (*getExperimentPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentPermissionLevelsResponsePb{}

	var permissionLevelsPb []experimentPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := experimentPermissionsDescriptionToPb(&item)
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

func (st *GetExperimentPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getExperimentPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []experimentPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getExperimentPermissionLevelsResponseFromPb(pb *getExperimentPermissionLevelsResponsePb) (*GetExperimentPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentPermissionLevelsResponse{}

	var permissionLevelsField []ExperimentPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := experimentPermissionsDescriptionFromPb(&itemPb)
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

// Get experiment permissions
type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func getExperimentPermissionsRequestToPb(st *GetExperimentPermissionsRequest) (*getExperimentPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentPermissionsRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func (st *GetExperimentPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getExperimentPermissionsRequestPb struct {
	// The experiment for which to get or manage permissions.
	ExperimentId string `json:"-" url:"-"`
}

func getExperimentPermissionsRequestFromPb(pb *getExperimentPermissionsRequestPb) (*GetExperimentPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentPermissionsRequest{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

// Get an experiment
type GetExperimentRequest struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func getExperimentRequestToPb(st *GetExperimentRequest) (*getExperimentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func (st *GetExperimentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentRequest) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getExperimentRequestPb struct {
	// ID of the associated experiment.
	ExperimentId string `json:"-" url:"experiment_id"`
}

func getExperimentRequestFromPb(pb *getExperimentRequestPb) (*GetExperimentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentRequest{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

type GetExperimentResponse struct {
	// Experiment details.
	// Wire name: 'experiment'
	Experiment *Experiment
}

func getExperimentResponseToPb(st *GetExperimentResponse) (*getExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getExperimentResponsePb{}
	experimentPb, err := experimentToPb(st.Experiment)
	if err != nil {
		return nil, err
	}
	if experimentPb != nil {
		pb.Experiment = experimentPb
	}

	return pb, nil
}

func (st *GetExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := getExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getExperimentResponsePb struct {
	// Experiment details.
	Experiment *experimentPb `json:"experiment,omitempty"`
}

func getExperimentResponseFromPb(pb *getExperimentResponsePb) (*GetExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentResponse{}
	experimentField, err := experimentFromPb(pb.Experiment)
	if err != nil {
		return nil, err
	}
	if experimentField != nil {
		st.Experiment = experimentField
	}

	return st, nil
}

// Get a forecasting experiment
type GetForecastingExperimentRequest struct {
	// The unique ID of a forecasting experiment
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func getForecastingExperimentRequestToPb(st *GetForecastingExperimentRequest) (*getForecastingExperimentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getForecastingExperimentRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func (st *GetForecastingExperimentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getForecastingExperimentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getForecastingExperimentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetForecastingExperimentRequest) MarshalJSON() ([]byte, error) {
	pb, err := getForecastingExperimentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getForecastingExperimentRequestPb struct {
	// The unique ID of a forecasting experiment
	ExperimentId string `json:"-" url:"-"`
}

func getForecastingExperimentRequestFromPb(pb *getForecastingExperimentRequestPb) (*GetForecastingExperimentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetForecastingExperimentRequest{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

// Get metric history for a run
type GetHistoryRequest struct {
	// Maximum number of Metric records to return per paginated request. Default
	// is set to 25,000. If set higher than 25,000, a request Exception will be
	// raised.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Name of the metric.
	// Wire name: 'metric_key'
	MetricKey string `tf:"-"`
	// Token indicating the page of metric histories to fetch.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// ID of the run from which to fetch metric values. Must be provided.
	// Wire name: 'run_id'
	RunId string `tf:"-"`
	// [Deprecated, use `run_id` instead] ID of the run from which to fetch
	// metric values. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getHistoryRequestToPb(st *GetHistoryRequest) (*getHistoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getHistoryRequestPb{}
	pb.MaxResults = st.MaxResults

	pb.MetricKey = st.MetricKey

	pb.PageToken = st.PageToken

	pb.RunId = st.RunId

	pb.RunUuid = st.RunUuid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetHistoryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getHistoryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getHistoryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetHistoryRequest) MarshalJSON() ([]byte, error) {
	pb, err := getHistoryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getHistoryRequestPb struct {
	// Maximum number of Metric records to return per paginated request. Default
	// is set to 25,000. If set higher than 25,000, a request Exception will be
	// raised.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Name of the metric.
	MetricKey string `json:"-" url:"metric_key"`
	// Token indicating the page of metric histories to fetch.
	PageToken string `json:"-" url:"page_token,omitempty"`
	// ID of the run from which to fetch metric values. Must be provided.
	RunId string `json:"-" url:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run from which to fetch
	// metric values. This field will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getHistoryRequestFromPb(pb *getHistoryRequestPb) (*GetHistoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetHistoryRequest{}
	st.MaxResults = pb.MaxResults
	st.MetricKey = pb.MetricKey
	st.PageToken = pb.PageToken
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getHistoryRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getHistoryRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetLatestVersionsRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string
	// List of stages.
	// Wire name: 'stages'
	Stages []string
}

func getLatestVersionsRequestToPb(st *GetLatestVersionsRequest) (*getLatestVersionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLatestVersionsRequestPb{}
	pb.Name = st.Name

	pb.Stages = st.Stages

	return pb, nil
}

func (st *GetLatestVersionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLatestVersionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLatestVersionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLatestVersionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getLatestVersionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getLatestVersionsRequestPb struct {
	// Registered model unique name identifier.
	Name string `json:"name"`
	// List of stages.
	Stages []string `json:"stages,omitempty"`
}

func getLatestVersionsRequestFromPb(pb *getLatestVersionsRequestPb) (*GetLatestVersionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLatestVersionsRequest{}
	st.Name = pb.Name
	st.Stages = pb.Stages

	return st, nil
}

type GetLatestVersionsResponse struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	// Wire name: 'model_versions'
	ModelVersions []ModelVersion
}

func getLatestVersionsResponseToPb(st *GetLatestVersionsResponse) (*getLatestVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getLatestVersionsResponsePb{}

	var modelVersionsPb []modelVersionPb
	for _, item := range st.ModelVersions {
		itemPb, err := modelVersionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			modelVersionsPb = append(modelVersionsPb, *itemPb)
		}
	}
	pb.ModelVersions = modelVersionsPb

	return pb, nil
}

func (st *GetLatestVersionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getLatestVersionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getLatestVersionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetLatestVersionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getLatestVersionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getLatestVersionsResponsePb struct {
	// Latest version models for each requests stage. Only return models with
	// current `READY` status. If no `stages` provided, returns the latest
	// version for each stage, including `"None"`.
	ModelVersions []modelVersionPb `json:"model_versions,omitempty"`
}

func getLatestVersionsResponseFromPb(pb *getLatestVersionsResponsePb) (*GetLatestVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLatestVersionsResponse{}

	var modelVersionsField []ModelVersion
	for _, itemPb := range pb.ModelVersions {
		item, err := modelVersionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			modelVersionsField = append(modelVersionsField, *item)
		}
	}
	st.ModelVersions = modelVersionsField

	return st, nil
}

type GetMetricHistoryResponse struct {
	// All logged values for this metric if `max_results` is not specified in
	// the request or if the total count of metrics returned is less than the
	// service level pagination threshold. Otherwise, this is one page of
	// results.
	// Wire name: 'metrics'
	Metrics []Metric
	// A token that can be used to issue a query for the next page of metric
	// history values. A missing token indicates that no additional metrics are
	// available to fetch.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func getMetricHistoryResponseToPb(st *GetMetricHistoryResponse) (*getMetricHistoryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getMetricHistoryResponsePb{}

	var metricsPb []metricPb
	for _, item := range st.Metrics {
		itemPb, err := metricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			metricsPb = append(metricsPb, *itemPb)
		}
	}
	pb.Metrics = metricsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetMetricHistoryResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getMetricHistoryResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getMetricHistoryResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetMetricHistoryResponse) MarshalJSON() ([]byte, error) {
	pb, err := getMetricHistoryResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getMetricHistoryResponsePb struct {
	// All logged values for this metric if `max_results` is not specified in
	// the request or if the total count of metrics returned is less than the
	// service level pagination threshold. Otherwise, this is one page of
	// results.
	Metrics []metricPb `json:"metrics,omitempty"`
	// A token that can be used to issue a query for the next page of metric
	// history values. A missing token indicates that no additional metrics are
	// available to fetch.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getMetricHistoryResponseFromPb(pb *getMetricHistoryResponsePb) (*GetMetricHistoryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetricHistoryResponse{}

	var metricsField []Metric
	for _, itemPb := range pb.Metrics {
		item, err := metricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			metricsField = append(metricsField, *item)
		}
	}
	st.Metrics = metricsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getMetricHistoryResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getMetricHistoryResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get model
type GetModelRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func getModelRequestToPb(st *GetModelRequest) (*getModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func (st *GetModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := getModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getModelRequestPb struct {
	// Registered model unique name identifier.
	Name string `json:"-" url:"name"`
}

func getModelRequestFromPb(pb *getModelRequestPb) (*GetModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetModelResponse struct {

	// Wire name: 'registered_model_databricks'
	RegisteredModelDatabricks *ModelDatabricks
}

func getModelResponseToPb(st *GetModelResponse) (*getModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelResponsePb{}
	registeredModelDatabricksPb, err := modelDatabricksToPb(st.RegisteredModelDatabricks)
	if err != nil {
		return nil, err
	}
	if registeredModelDatabricksPb != nil {
		pb.RegisteredModelDatabricks = registeredModelDatabricksPb
	}

	return pb, nil
}

func (st *GetModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := getModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getModelResponsePb struct {
	RegisteredModelDatabricks *modelDatabricksPb `json:"registered_model_databricks,omitempty"`
}

func getModelResponseFromPb(pb *getModelResponsePb) (*GetModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelResponse{}
	registeredModelDatabricksField, err := modelDatabricksFromPb(pb.RegisteredModelDatabricks)
	if err != nil {
		return nil, err
	}
	if registeredModelDatabricksField != nil {
		st.RegisteredModelDatabricks = registeredModelDatabricksField
	}

	return st, nil
}

// Get a model version URI
type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number
	// Wire name: 'version'
	Version string `tf:"-"`
}

func getModelVersionDownloadUriRequestToPb(st *GetModelVersionDownloadUriRequest) (*getModelVersionDownloadUriRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionDownloadUriRequestPb{}
	pb.Name = st.Name

	pb.Version = st.Version

	return pb, nil
}

func (st *GetModelVersionDownloadUriRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelVersionDownloadUriRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelVersionDownloadUriRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelVersionDownloadUriRequest) MarshalJSON() ([]byte, error) {
	pb, err := getModelVersionDownloadUriRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getModelVersionDownloadUriRequestPb struct {
	// Name of the registered model
	Name string `json:"-" url:"name"`
	// Model version number
	Version string `json:"-" url:"version"`
}

func getModelVersionDownloadUriRequestFromPb(pb *getModelVersionDownloadUriRequestPb) (*GetModelVersionDownloadUriRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionDownloadUriRequest{}
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

type GetModelVersionDownloadUriResponse struct {
	// URI corresponding to where artifacts for this model version are stored.
	// Wire name: 'artifact_uri'
	ArtifactUri string

	ForceSendFields []string `tf:"-"`
}

func getModelVersionDownloadUriResponseToPb(st *GetModelVersionDownloadUriResponse) (*getModelVersionDownloadUriResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionDownloadUriResponsePb{}
	pb.ArtifactUri = st.ArtifactUri

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetModelVersionDownloadUriResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelVersionDownloadUriResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelVersionDownloadUriResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelVersionDownloadUriResponse) MarshalJSON() ([]byte, error) {
	pb, err := getModelVersionDownloadUriResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getModelVersionDownloadUriResponsePb struct {
	// URI corresponding to where artifacts for this model version are stored.
	ArtifactUri string `json:"artifact_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getModelVersionDownloadUriResponseFromPb(pb *getModelVersionDownloadUriResponsePb) (*GetModelVersionDownloadUriResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionDownloadUriResponse{}
	st.ArtifactUri = pb.ArtifactUri

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getModelVersionDownloadUriResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getModelVersionDownloadUriResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Get a model version
type GetModelVersionRequest struct {
	// Name of the registered model
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number
	// Wire name: 'version'
	Version string `tf:"-"`
}

func getModelVersionRequestToPb(st *GetModelVersionRequest) (*getModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionRequestPb{}
	pb.Name = st.Name

	pb.Version = st.Version

	return pb, nil
}

func (st *GetModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := getModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getModelVersionRequestPb struct {
	// Name of the registered model
	Name string `json:"-" url:"name"`
	// Model version number
	Version string `json:"-" url:"version"`
}

func getModelVersionRequestFromPb(pb *getModelVersionRequestPb) (*GetModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionRequest{}
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

type GetModelVersionResponse struct {

	// Wire name: 'model_version'
	ModelVersion *ModelVersion
}

func getModelVersionResponseToPb(st *GetModelVersionResponse) (*getModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getModelVersionResponsePb{}
	modelVersionPb, err := modelVersionToPb(st.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionPb != nil {
		pb.ModelVersion = modelVersionPb
	}

	return pb, nil
}

func (st *GetModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := getModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getModelVersionResponsePb struct {
	ModelVersion *modelVersionPb `json:"model_version,omitempty"`
}

func getModelVersionResponseFromPb(pb *getModelVersionResponsePb) (*GetModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionResponse{}
	modelVersionField, err := modelVersionFromPb(pb.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionField != nil {
		st.ModelVersion = modelVersionField
	}

	return st, nil
}

// Get registered model permission levels
type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	// Wire name: 'registered_model_id'
	RegisteredModelId string `tf:"-"`
}

func getRegisteredModelPermissionLevelsRequestToPb(st *GetRegisteredModelPermissionLevelsRequest) (*getRegisteredModelPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRegisteredModelPermissionLevelsRequestPb{}
	pb.RegisteredModelId = st.RegisteredModelId

	return pb, nil
}

func (st *GetRegisteredModelPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRegisteredModelPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRegisteredModelPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRegisteredModelPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRegisteredModelPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRegisteredModelPermissionLevelsRequestPb struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" url:"-"`
}

func getRegisteredModelPermissionLevelsRequestFromPb(pb *getRegisteredModelPermissionLevelsRequestPb) (*GetRegisteredModelPermissionLevelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRegisteredModelPermissionLevelsRequest{}
	st.RegisteredModelId = pb.RegisteredModelId

	return st, nil
}

type GetRegisteredModelPermissionLevelsResponse struct {
	// Specific permission levels
	// Wire name: 'permission_levels'
	PermissionLevels []RegisteredModelPermissionsDescription
}

func getRegisteredModelPermissionLevelsResponseToPb(st *GetRegisteredModelPermissionLevelsResponse) (*getRegisteredModelPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRegisteredModelPermissionLevelsResponsePb{}

	var permissionLevelsPb []registeredModelPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := registeredModelPermissionsDescriptionToPb(&item)
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

func (st *GetRegisteredModelPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRegisteredModelPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRegisteredModelPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRegisteredModelPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := getRegisteredModelPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRegisteredModelPermissionLevelsResponsePb struct {
	// Specific permission levels
	PermissionLevels []registeredModelPermissionsDescriptionPb `json:"permission_levels,omitempty"`
}

func getRegisteredModelPermissionLevelsResponseFromPb(pb *getRegisteredModelPermissionLevelsResponsePb) (*GetRegisteredModelPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRegisteredModelPermissionLevelsResponse{}

	var permissionLevelsField []RegisteredModelPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := registeredModelPermissionsDescriptionFromPb(&itemPb)
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

// Get registered model permissions
type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	// Wire name: 'registered_model_id'
	RegisteredModelId string `tf:"-"`
}

func getRegisteredModelPermissionsRequestToPb(st *GetRegisteredModelPermissionsRequest) (*getRegisteredModelPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRegisteredModelPermissionsRequestPb{}
	pb.RegisteredModelId = st.RegisteredModelId

	return pb, nil
}

func (st *GetRegisteredModelPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRegisteredModelPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRegisteredModelPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRegisteredModelPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRegisteredModelPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRegisteredModelPermissionsRequestPb struct {
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" url:"-"`
}

func getRegisteredModelPermissionsRequestFromPb(pb *getRegisteredModelPermissionsRequestPb) (*GetRegisteredModelPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRegisteredModelPermissionsRequest{}
	st.RegisteredModelId = pb.RegisteredModelId

	return st, nil
}

// Get a run
type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	// Wire name: 'run_id'
	RunId string `tf:"-"`
	// [Deprecated, use `run_id` instead] ID of the run to fetch. This field
	// will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func getRunRequestToPb(st *GetRunRequest) (*getRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRunRequestPb{}
	pb.RunId = st.RunId

	pb.RunUuid = st.RunUuid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *GetRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := getRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRunRequestPb struct {
	// ID of the run to fetch. Must be provided.
	RunId string `json:"-" url:"run_id"`
	// [Deprecated, use `run_id` instead] ID of the run to fetch. This field
	// will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func getRunRequestFromPb(pb *getRunRequestPb) (*GetRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunRequest{}
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *getRunRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st getRunRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type GetRunResponse struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	// Wire name: 'run'
	Run *Run
}

func getRunResponseToPb(st *GetRunResponse) (*getRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &getRunResponsePb{}
	runPb, err := runToPb(st.Run)
	if err != nil {
		return nil, err
	}
	if runPb != nil {
		pb.Run = runPb
	}

	return pb, nil
}

func (st *GetRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &getRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := getRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st GetRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := getRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type getRunResponsePb struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	Run *runPb `json:"run,omitempty"`
}

func getRunResponseFromPb(pb *getRunResponsePb) (*GetRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunResponse{}
	runField, err := runFromPb(pb.Run)
	if err != nil {
		return nil, err
	}
	if runField != nil {
		st.Run = runField
	}

	return st, nil
}

type HttpUrlSpec struct {
	// Value of the authorization header that should be sent in the request sent
	// by the wehbook. It should be of the form `"<auth type> <credentials>"`.
	// If set to an empty string, no authorization header will be included in
	// the request.
	// Wire name: 'authorization'
	Authorization string
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	// Wire name: 'enable_ssl_verification'
	EnableSslVerification bool
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	// Wire name: 'secret'
	Secret string
	// External HTTPS URL called on event trigger (by using a POST request).
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func httpUrlSpecToPb(st *HttpUrlSpec) (*httpUrlSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpUrlSpecPb{}
	pb.Authorization = st.Authorization

	pb.EnableSslVerification = st.EnableSslVerification

	pb.Secret = st.Secret

	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *HttpUrlSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &httpUrlSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := httpUrlSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st HttpUrlSpec) MarshalJSON() ([]byte, error) {
	pb, err := httpUrlSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type httpUrlSpecPb struct {
	// Value of the authorization header that should be sent in the request sent
	// by the wehbook. It should be of the form `"<auth type> <credentials>"`.
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
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	Secret string `json:"secret,omitempty"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url string `json:"url"`

	ForceSendFields []string `json:"-" url:"-"`
}

func httpUrlSpecFromPb(pb *httpUrlSpecPb) (*HttpUrlSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &HttpUrlSpec{}
	st.Authorization = pb.Authorization
	st.EnableSslVerification = pb.EnableSslVerification
	st.Secret = pb.Secret
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *httpUrlSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st httpUrlSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type HttpUrlSpecWithoutSecret struct {
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	// Wire name: 'enable_ssl_verification'
	EnableSslVerification bool
	// External HTTPS URL called on event trigger (by using a POST request).
	// Wire name: 'url'
	Url string

	ForceSendFields []string `tf:"-"`
}

func httpUrlSpecWithoutSecretToPb(st *HttpUrlSpecWithoutSecret) (*httpUrlSpecWithoutSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &httpUrlSpecWithoutSecretPb{}
	pb.EnableSslVerification = st.EnableSslVerification

	pb.Url = st.Url

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *HttpUrlSpecWithoutSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &httpUrlSpecWithoutSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := httpUrlSpecWithoutSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st HttpUrlSpecWithoutSecret) MarshalJSON() ([]byte, error) {
	pb, err := httpUrlSpecWithoutSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type httpUrlSpecWithoutSecretPb struct {
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	EnableSslVerification bool `json:"enable_ssl_verification,omitempty"`
	// External HTTPS URL called on event trigger (by using a POST request).
	Url string `json:"url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func httpUrlSpecWithoutSecretFromPb(pb *httpUrlSpecWithoutSecretPb) (*HttpUrlSpecWithoutSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &HttpUrlSpecWithoutSecret{}
	st.EnableSslVerification = pb.EnableSslVerification
	st.Url = pb.Url

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *httpUrlSpecWithoutSecretPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st httpUrlSpecWithoutSecretPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Tag for a dataset input.
type InputTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string
}

func inputTagToPb(st *InputTag) (*inputTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &inputTagPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	return pb, nil
}

func (st *InputTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &inputTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := inputTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st InputTag) MarshalJSON() ([]byte, error) {
	pb, err := inputTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type inputTagPb struct {
	// The tag key.
	Key string `json:"key"`
	// The tag value.
	Value string `json:"value"`
}

func inputTagFromPb(pb *inputTagPb) (*InputTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &InputTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	return st, nil
}

type JobSpec struct {
	// The personal access token used to authorize webhook's job runs.
	// Wire name: 'access_token'
	AccessToken string
	// ID of the job that the webhook runs.
	// Wire name: 'job_id'
	JobId string
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	// Wire name: 'workspace_url'
	WorkspaceUrl string

	ForceSendFields []string `tf:"-"`
}

func jobSpecToPb(st *JobSpec) (*jobSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobSpecPb{}
	pb.AccessToken = st.AccessToken

	pb.JobId = st.JobId

	pb.WorkspaceUrl = st.WorkspaceUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *JobSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobSpec) MarshalJSON() ([]byte, error) {
	pb, err := jobSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobSpecPb struct {
	// The personal access token used to authorize webhook's job runs.
	AccessToken string `json:"access_token"`
	// ID of the job that the webhook runs.
	JobId string `json:"job_id"`
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobSpecFromPb(pb *jobSpecPb) (*JobSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSpec{}
	st.AccessToken = pb.AccessToken
	st.JobId = pb.JobId
	st.WorkspaceUrl = pb.WorkspaceUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobSpecPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobSpecPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type JobSpecWithoutSecret struct {
	// ID of the job that the webhook runs.
	// Wire name: 'job_id'
	JobId string
	// URL of the workspace containing the job that this webhook runs. Defaults
	// to the workspace URL in which the webhook is created. If not specified,
	// the job’s workspace is assumed to be the same as the webhook’s.
	// Wire name: 'workspace_url'
	WorkspaceUrl string

	ForceSendFields []string `tf:"-"`
}

func jobSpecWithoutSecretToPb(st *JobSpecWithoutSecret) (*jobSpecWithoutSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &jobSpecWithoutSecretPb{}
	pb.JobId = st.JobId

	pb.WorkspaceUrl = st.WorkspaceUrl

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *JobSpecWithoutSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &jobSpecWithoutSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := jobSpecWithoutSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st JobSpecWithoutSecret) MarshalJSON() ([]byte, error) {
	pb, err := jobSpecWithoutSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type jobSpecWithoutSecretPb struct {
	// ID of the job that the webhook runs.
	JobId string `json:"job_id,omitempty"`
	// URL of the workspace containing the job that this webhook runs. Defaults
	// to the workspace URL in which the webhook is created. If not specified,
	// the job’s workspace is assumed to be the same as the webhook’s.
	WorkspaceUrl string `json:"workspace_url,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func jobSpecWithoutSecretFromPb(pb *jobSpecWithoutSecretPb) (*JobSpecWithoutSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSpecWithoutSecret{}
	st.JobId = pb.JobId
	st.WorkspaceUrl = pb.WorkspaceUrl

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *jobSpecWithoutSecretPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st jobSpecWithoutSecretPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List artifacts
type ListArtifactsRequest struct {
	// The token indicating the page of artifact results to fetch. `page_token`
	// is not supported when listing artifacts in UC Volumes. A maximum of 1000
	// artifacts will be retrieved for UC Volumes. Please call
	// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
	// Volumes, which supports pagination. See [List directory contents | Files
	// API](/api/workspace/files/listdirectorycontents).
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	// Wire name: 'path'
	Path string `tf:"-"`
	// ID of the run whose artifacts to list. Must be provided.
	// Wire name: 'run_id'
	RunId string `tf:"-"`
	// [Deprecated, use `run_id` instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listArtifactsRequestToPb(st *ListArtifactsRequest) (*listArtifactsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listArtifactsRequestPb{}
	pb.PageToken = st.PageToken

	pb.Path = st.Path

	pb.RunId = st.RunId

	pb.RunUuid = st.RunUuid

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListArtifactsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listArtifactsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listArtifactsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListArtifactsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listArtifactsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listArtifactsRequestPb struct {
	// The token indicating the page of artifact results to fetch. `page_token`
	// is not supported when listing artifacts in UC Volumes. A maximum of 1000
	// artifacts will be retrieved for UC Volumes. Please call
	// `/api/2.0/fs/directories{directory_path}` for listing artifacts in UC
	// Volumes, which supports pagination. See [List directory contents | Files
	// API](/api/workspace/files/listdirectorycontents).
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Filter artifacts matching this path (a relative path from the root
	// artifact directory).
	Path string `json:"-" url:"path,omitempty"`
	// ID of the run whose artifacts to list. Must be provided.
	RunId string `json:"-" url:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run whose artifacts to list.
	// This field will be removed in a future MLflow version.
	RunUuid string `json:"-" url:"run_uuid,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listArtifactsRequestFromPb(pb *listArtifactsRequestPb) (*ListArtifactsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListArtifactsRequest{}
	st.PageToken = pb.PageToken
	st.Path = pb.Path
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listArtifactsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listArtifactsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListArtifactsResponse struct {
	// The file location and metadata for artifacts.
	// Wire name: 'files'
	Files []FileInfo
	// The token that can be used to retrieve the next page of artifact results.
	// Wire name: 'next_page_token'
	NextPageToken string
	// The root artifact directory for the run.
	// Wire name: 'root_uri'
	RootUri string

	ForceSendFields []string `tf:"-"`
}

func listArtifactsResponseToPb(st *ListArtifactsResponse) (*listArtifactsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listArtifactsResponsePb{}

	var filesPb []fileInfoPb
	for _, item := range st.Files {
		itemPb, err := fileInfoToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			filesPb = append(filesPb, *itemPb)
		}
	}
	pb.Files = filesPb

	pb.NextPageToken = st.NextPageToken

	pb.RootUri = st.RootUri

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListArtifactsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listArtifactsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listArtifactsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListArtifactsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listArtifactsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listArtifactsResponsePb struct {
	// The file location and metadata for artifacts.
	Files []fileInfoPb `json:"files,omitempty"`
	// The token that can be used to retrieve the next page of artifact results.
	NextPageToken string `json:"next_page_token,omitempty"`
	// The root artifact directory for the run.
	RootUri string `json:"root_uri,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listArtifactsResponseFromPb(pb *listArtifactsResponsePb) (*ListArtifactsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListArtifactsResponse{}

	var filesField []FileInfo
	for _, itemPb := range pb.Files {
		item, err := fileInfoFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			filesField = append(filesField, *item)
		}
	}
	st.Files = filesField
	st.NextPageToken = pb.NextPageToken
	st.RootUri = pb.RootUri

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listArtifactsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listArtifactsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List experiments
type ListExperimentsRequest struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it'll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	// Wire name: 'max_results'
	MaxResults int64 `tf:"-"`
	// Token indicating the page of experiments to fetch
	// Wire name: 'page_token'
	PageToken string `tf:"-"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	// Wire name: 'view_type'
	ViewType ViewType `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listExperimentsRequestToPb(st *ListExperimentsRequest) (*listExperimentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExperimentsRequestPb{}
	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.ViewType = st.ViewType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListExperimentsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExperimentsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExperimentsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExperimentsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listExperimentsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listExperimentsRequestPb struct {
	// Maximum number of experiments desired. If `max_results` is unspecified,
	// return all experiments. If `max_results` is too large, it'll be
	// automatically capped at 1000. Callers of this endpoint are encouraged to
	// pass max_results explicitly and leverage page_token to iterate through
	// experiments.
	MaxResults int64 `json:"-" url:"max_results,omitempty"`
	// Token indicating the page of experiments to fetch
	PageToken string `json:"-" url:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType ViewType `json:"-" url:"view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExperimentsRequestFromPb(pb *listExperimentsRequestPb) (*ListExperimentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExperimentsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	st.ViewType = pb.ViewType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExperimentsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExperimentsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListExperimentsResponse struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	// Wire name: 'experiments'
	Experiments []Experiment
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func listExperimentsResponseToPb(st *ListExperimentsResponse) (*listExperimentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listExperimentsResponsePb{}

	var experimentsPb []experimentPb
	for _, item := range st.Experiments {
		itemPb, err := experimentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			experimentsPb = append(experimentsPb, *itemPb)
		}
	}
	pb.Experiments = experimentsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListExperimentsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listExperimentsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listExperimentsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListExperimentsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listExperimentsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listExperimentsResponsePb struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	Experiments []experimentPb `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listExperimentsResponseFromPb(pb *listExperimentsResponsePb) (*ListExperimentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExperimentsResponse{}

	var experimentsField []Experiment
	for _, itemPb := range pb.Experiments {
		item, err := experimentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			experimentsField = append(experimentsField, *item)
		}
	}
	st.Experiments = experimentsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listExperimentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listExperimentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List models
type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// Pagination token to go to the next page based on a previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listModelsRequestToPb(st *ListModelsRequest) (*listModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listModelsRequestPb{}
	pb.MaxResults = st.MaxResults

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListModelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listModelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listModelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListModelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listModelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listModelsRequestPb struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// Pagination token to go to the next page based on a previous query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listModelsRequestFromPb(pb *listModelsRequestPb) (*ListModelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListModelsResponse struct {
	// Pagination token to request next page of models for the same query.
	// Wire name: 'next_page_token'
	NextPageToken string

	// Wire name: 'registered_models'
	RegisteredModels []Model

	ForceSendFields []string `tf:"-"`
}

func listModelsResponseToPb(st *ListModelsResponse) (*listModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listModelsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var registeredModelsPb []modelPb
	for _, item := range st.RegisteredModels {
		itemPb, err := modelToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			registeredModelsPb = append(registeredModelsPb, *itemPb)
		}
	}
	pb.RegisteredModels = registeredModelsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListModelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listModelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listModelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListModelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listModelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listModelsResponsePb struct {
	// Pagination token to request next page of models for the same query.
	NextPageToken string `json:"next_page_token,omitempty"`

	RegisteredModels []modelPb `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listModelsResponseFromPb(pb *listModelsResponsePb) (*ListModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelsResponse{}
	st.NextPageToken = pb.NextPageToken

	var registeredModelsField []Model
	for _, itemPb := range pb.RegisteredModels {
		item, err := modelFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			registeredModelsField = append(registeredModelsField, *item)
		}
	}
	st.RegisteredModels = registeredModelsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	// Wire name: 'next_page_token'
	NextPageToken string
	// Array of registry webhooks.
	// Wire name: 'webhooks'
	Webhooks []RegistryWebhook

	ForceSendFields []string `tf:"-"`
}

func listRegistryWebhooksToPb(st *ListRegistryWebhooks) (*listRegistryWebhooksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listRegistryWebhooksPb{}
	pb.NextPageToken = st.NextPageToken

	var webhooksPb []registryWebhookPb
	for _, item := range st.Webhooks {
		itemPb, err := registryWebhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			webhooksPb = append(webhooksPb, *itemPb)
		}
	}
	pb.Webhooks = webhooksPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListRegistryWebhooks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listRegistryWebhooksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listRegistryWebhooksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListRegistryWebhooks) MarshalJSON() ([]byte, error) {
	pb, err := listRegistryWebhooksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listRegistryWebhooksPb struct {
	// Token that can be used to retrieve the next page of artifact results
	NextPageToken string `json:"next_page_token,omitempty"`
	// Array of registry webhooks.
	Webhooks []registryWebhookPb `json:"webhooks,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listRegistryWebhooksFromPb(pb *listRegistryWebhooksPb) (*ListRegistryWebhooks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRegistryWebhooks{}
	st.NextPageToken = pb.NextPageToken

	var webhooksField []RegistryWebhook
	for _, itemPb := range pb.Webhooks {
		item, err := registryWebhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			webhooksField = append(webhooksField, *item)
		}
	}
	st.Webhooks = webhooksField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listRegistryWebhooksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listRegistryWebhooksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// List transition requests
type ListTransitionRequestsRequest struct {
	// Name of the model.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Version of the model.
	// Wire name: 'version'
	Version string `tf:"-"`
}

func listTransitionRequestsRequestToPb(st *ListTransitionRequestsRequest) (*listTransitionRequestsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTransitionRequestsRequestPb{}
	pb.Name = st.Name

	pb.Version = st.Version

	return pb, nil
}

func (st *ListTransitionRequestsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listTransitionRequestsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listTransitionRequestsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListTransitionRequestsRequest) MarshalJSON() ([]byte, error) {
	pb, err := listTransitionRequestsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listTransitionRequestsRequestPb struct {
	// Name of the model.
	Name string `json:"-" url:"name"`
	// Version of the model.
	Version string `json:"-" url:"version"`
}

func listTransitionRequestsRequestFromPb(pb *listTransitionRequestsRequestPb) (*ListTransitionRequestsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTransitionRequestsRequest{}
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

type ListTransitionRequestsResponse struct {
	// Array of open transition requests.
	// Wire name: 'requests'
	Requests []Activity
}

func listTransitionRequestsResponseToPb(st *ListTransitionRequestsResponse) (*listTransitionRequestsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listTransitionRequestsResponsePb{}

	var requestsPb []activityPb
	for _, item := range st.Requests {
		itemPb, err := activityToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			requestsPb = append(requestsPb, *itemPb)
		}
	}
	pb.Requests = requestsPb

	return pb, nil
}

func (st *ListTransitionRequestsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listTransitionRequestsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listTransitionRequestsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListTransitionRequestsResponse) MarshalJSON() ([]byte, error) {
	pb, err := listTransitionRequestsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listTransitionRequestsResponsePb struct {
	// Array of open transition requests.
	Requests []activityPb `json:"requests,omitempty"`
}

func listTransitionRequestsResponseFromPb(pb *listTransitionRequestsResponsePb) (*ListTransitionRequestsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTransitionRequestsResponse{}

	var requestsField []Activity
	for _, itemPb := range pb.Requests {
		item, err := activityFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			requestsField = append(requestsField, *item)
		}
	}
	st.Requests = requestsField

	return st, nil
}

// List registry webhooks
type ListWebhooksRequest struct {
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	// Wire name: 'events'
	Events []RegistryWebhookEvent `tf:"-"`
	// If not specified, all webhooks associated with the specified events are
	// listed, regardless of their associated model.
	// Wire name: 'model_name'
	ModelName string `tf:"-"`
	// Token indicating the page of artifact results to fetch
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func listWebhooksRequestToPb(st *ListWebhooksRequest) (*listWebhooksRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &listWebhooksRequestPb{}
	pb.Events = st.Events

	pb.ModelName = st.ModelName

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ListWebhooksRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &listWebhooksRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := listWebhooksRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ListWebhooksRequest) MarshalJSON() ([]byte, error) {
	pb, err := listWebhooksRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type listWebhooksRequestPb struct {
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	Events []RegistryWebhookEvent `json:"-" url:"events,omitempty"`
	// If not specified, all webhooks associated with the specified events are
	// listed, regardless of their associated model.
	ModelName string `json:"-" url:"model_name,omitempty"`
	// Token indicating the page of artifact results to fetch
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func listWebhooksRequestFromPb(pb *listWebhooksRequestPb) (*ListWebhooksRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWebhooksRequest{}
	st.Events = pb.Events
	st.ModelName = pb.ModelName
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *listWebhooksRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st listWebhooksRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogBatch struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	// Wire name: 'metrics'
	Metrics []Metric
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	// Wire name: 'params'
	Params []Param
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	// Wire name: 'tags'
	Tags []RunTag

	ForceSendFields []string `tf:"-"`
}

func logBatchToPb(st *LogBatch) (*logBatchPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logBatchPb{}

	var metricsPb []metricPb
	for _, item := range st.Metrics {
		itemPb, err := metricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			metricsPb = append(metricsPb, *itemPb)
		}
	}
	pb.Metrics = metricsPb

	var paramsPb []paramPb
	for _, item := range st.Params {
		itemPb, err := paramToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			paramsPb = append(paramsPb, *itemPb)
		}
	}
	pb.Params = paramsPb

	pb.RunId = st.RunId

	var tagsPb []runTagPb
	for _, item := range st.Tags {
		itemPb, err := runTagToPb(&item)
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

func (st *LogBatch) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logBatchPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logBatchFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogBatch) MarshalJSON() ([]byte, error) {
	pb, err := logBatchToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logBatchPb struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	Metrics []metricPb `json:"metrics,omitempty"`
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	Params []paramPb `json:"params,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id,omitempty"`
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	Tags []runTagPb `json:"tags,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logBatchFromPb(pb *logBatchPb) (*LogBatch, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogBatch{}

	var metricsField []Metric
	for _, itemPb := range pb.Metrics {
		item, err := metricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			metricsField = append(metricsField, *item)
		}
	}
	st.Metrics = metricsField

	var paramsField []Param
	for _, itemPb := range pb.Params {
		item, err := paramFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			paramsField = append(paramsField, *item)
		}
	}
	st.Params = paramsField
	st.RunId = pb.RunId

	var tagsField []RunTag
	for _, itemPb := range pb.Tags {
		item, err := runTagFromPb(&itemPb)
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

func (st *logBatchPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logBatchPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogBatchResponse struct {
}

func logBatchResponseToPb(st *LogBatchResponse) (*logBatchResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logBatchResponsePb{}

	return pb, nil
}

func (st *LogBatchResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logBatchResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logBatchResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogBatchResponse) MarshalJSON() ([]byte, error) {
	pb, err := logBatchResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logBatchResponsePb struct {
}

func logBatchResponseFromPb(pb *logBatchResponsePb) (*LogBatchResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogBatchResponse{}

	return st, nil
}

type LogInputs struct {
	// Dataset inputs
	// Wire name: 'datasets'
	Datasets []DatasetInput
	// Model inputs
	// Wire name: 'models'
	Models []ModelInput
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string
}

func logInputsToPb(st *LogInputs) (*logInputsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logInputsPb{}

	var datasetsPb []datasetInputPb
	for _, item := range st.Datasets {
		itemPb, err := datasetInputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			datasetsPb = append(datasetsPb, *itemPb)
		}
	}
	pb.Datasets = datasetsPb

	var modelsPb []modelInputPb
	for _, item := range st.Models {
		itemPb, err := modelInputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			modelsPb = append(modelsPb, *itemPb)
		}
	}
	pb.Models = modelsPb

	pb.RunId = st.RunId

	return pb, nil
}

func (st *LogInputs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logInputsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logInputsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogInputs) MarshalJSON() ([]byte, error) {
	pb, err := logInputsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logInputsPb struct {
	// Dataset inputs
	Datasets []datasetInputPb `json:"datasets,omitempty"`
	// Model inputs
	Models []modelInputPb `json:"models,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id"`
}

func logInputsFromPb(pb *logInputsPb) (*LogInputs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogInputs{}

	var datasetsField []DatasetInput
	for _, itemPb := range pb.Datasets {
		item, err := datasetInputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			datasetsField = append(datasetsField, *item)
		}
	}
	st.Datasets = datasetsField

	var modelsField []ModelInput
	for _, itemPb := range pb.Models {
		item, err := modelInputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			modelsField = append(modelsField, *item)
		}
	}
	st.Models = modelsField
	st.RunId = pb.RunId

	return st, nil
}

type LogInputsResponse struct {
}

func logInputsResponseToPb(st *LogInputsResponse) (*logInputsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logInputsResponsePb{}

	return pb, nil
}

func (st *LogInputsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logInputsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logInputsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogInputsResponse) MarshalJSON() ([]byte, error) {
	pb, err := logInputsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logInputsResponsePb struct {
}

func logInputsResponseFromPb(pb *logInputsResponsePb) (*LogInputsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogInputsResponse{}

	return st, nil
}

type LogMetric struct {
	// Dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	// Wire name: 'dataset_digest'
	DatasetDigest string
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	// Wire name: 'dataset_name'
	DatasetName string
	// Name of the metric.
	// Wire name: 'key'
	Key string
	// ID of the logged model associated with the metric, if applicable
	// Wire name: 'model_id'
	ModelId string
	// ID of the run under which to log the metric. Must be provided.
	// Wire name: 'run_id'
	RunId string
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string
	// Step at which to log the metric
	// Wire name: 'step'
	Step int64
	// Unix timestamp in milliseconds at the time metric was logged.
	// Wire name: 'timestamp'
	Timestamp int64
	// Double value of the metric being logged.
	// Wire name: 'value'
	Value float64

	ForceSendFields []string `tf:"-"`
}

func logMetricToPb(st *LogMetric) (*logMetricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logMetricPb{}
	pb.DatasetDigest = st.DatasetDigest

	pb.DatasetName = st.DatasetName

	pb.Key = st.Key

	pb.ModelId = st.ModelId

	pb.RunId = st.RunId

	pb.RunUuid = st.RunUuid

	pb.Step = st.Step

	pb.Timestamp = st.Timestamp

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *LogMetric) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logMetricPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logMetricFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogMetric) MarshalJSON() ([]byte, error) {
	pb, err := logMetricToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logMetricPb struct {
	// Dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	DatasetName string `json:"dataset_name,omitempty"`
	// Name of the metric.
	Key string `json:"key"`
	// ID of the logged model associated with the metric, if applicable
	ModelId string `json:"model_id,omitempty"`
	// ID of the run under which to log the metric. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Step at which to log the metric
	Step int64 `json:"step,omitempty"`
	// Unix timestamp in milliseconds at the time metric was logged.
	Timestamp int64 `json:"timestamp"`
	// Double value of the metric being logged.
	Value float64 `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logMetricFromPb(pb *logMetricPb) (*LogMetric, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogMetric{}
	st.DatasetDigest = pb.DatasetDigest
	st.DatasetName = pb.DatasetName
	st.Key = pb.Key
	st.ModelId = pb.ModelId
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid
	st.Step = pb.Step
	st.Timestamp = pb.Timestamp
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logMetricPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logMetricPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogMetricResponse struct {
}

func logMetricResponseToPb(st *LogMetricResponse) (*logMetricResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logMetricResponsePb{}

	return pb, nil
}

func (st *LogMetricResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logMetricResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logMetricResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogMetricResponse) MarshalJSON() ([]byte, error) {
	pb, err := logMetricResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logMetricResponsePb struct {
}

func logMetricResponseFromPb(pb *logMetricResponsePb) (*LogMetricResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogMetricResponse{}

	return st, nil
}

type LogModel struct {
	// MLmodel file in json format.
	// Wire name: 'model_json'
	ModelJson string
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string

	ForceSendFields []string `tf:"-"`
}

func logModelToPb(st *LogModel) (*logModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logModelPb{}
	pb.ModelJson = st.ModelJson

	pb.RunId = st.RunId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *LogModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogModel) MarshalJSON() ([]byte, error) {
	pb, err := logModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logModelPb struct {
	// MLmodel file in json format.
	ModelJson string `json:"model_json,omitempty"`
	// ID of the run to log under
	RunId string `json:"run_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logModelFromPb(pb *logModelPb) (*LogModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogModel{}
	st.ModelJson = pb.ModelJson
	st.RunId = pb.RunId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logModelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logModelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogModelResponse struct {
}

func logModelResponseToPb(st *LogModelResponse) (*logModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logModelResponsePb{}

	return pb, nil
}

func (st *LogModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := logModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logModelResponsePb struct {
}

func logModelResponseFromPb(pb *logModelResponsePb) (*LogModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogModelResponse{}

	return st, nil
}

type LogParam struct {
	// Name of the param. Maximum size is 255 bytes.
	// Wire name: 'key'
	Key string
	// ID of the run under which to log the param. Must be provided.
	// Wire name: 'run_id'
	RunId string
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string
	// String value of the param being logged. Maximum size is 500 bytes.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func logParamToPb(st *LogParam) (*logParamPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logParamPb{}
	pb.Key = st.Key

	pb.RunId = st.RunId

	pb.RunUuid = st.RunUuid

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *LogParam) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logParamPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logParamFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogParam) MarshalJSON() ([]byte, error) {
	pb, err := logParamToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logParamPb struct {
	// Name of the param. Maximum size is 255 bytes.
	Key string `json:"key"`
	// ID of the run under which to log the param. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// String value of the param being logged. Maximum size is 500 bytes.
	Value string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func logParamFromPb(pb *logParamPb) (*LogParam, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogParam{}
	st.Key = pb.Key
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *logParamPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st logParamPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type LogParamResponse struct {
}

func logParamResponseToPb(st *LogParamResponse) (*logParamResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &logParamResponsePb{}

	return pb, nil
}

func (st *LogParamResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &logParamResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := logParamResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st LogParamResponse) MarshalJSON() ([]byte, error) {
	pb, err := logParamResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type logParamResponsePb struct {
}

func logParamResponseFromPb(pb *logParamResponsePb) (*LogParamResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogParamResponse{}

	return st, nil
}

// Metric associated with a run, represented as a key-value pair.
type Metric struct {
	// The dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	// Wire name: 'dataset_digest'
	DatasetDigest string
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	// Wire name: 'dataset_name'
	DatasetName string
	// The key identifying the metric.
	// Wire name: 'key'
	Key string
	// The ID of the logged model or registered model version associated with
	// the metric, if applicable.
	// Wire name: 'model_id'
	ModelId string
	// The ID of the run containing the metric.
	// Wire name: 'run_id'
	RunId string
	// The step at which the metric was logged.
	// Wire name: 'step'
	Step int64
	// The timestamp at which the metric was recorded.
	// Wire name: 'timestamp'
	Timestamp int64
	// The value of the metric.
	// Wire name: 'value'
	Value float64

	ForceSendFields []string `tf:"-"`
}

func metricToPb(st *Metric) (*metricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &metricPb{}
	pb.DatasetDigest = st.DatasetDigest

	pb.DatasetName = st.DatasetName

	pb.Key = st.Key

	pb.ModelId = st.ModelId

	pb.RunId = st.RunId

	pb.Step = st.Step

	pb.Timestamp = st.Timestamp

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Metric) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &metricPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := metricFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Metric) MarshalJSON() ([]byte, error) {
	pb, err := metricToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type metricPb struct {
	// The dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	DatasetDigest string `json:"dataset_digest,omitempty"`
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	DatasetName string `json:"dataset_name,omitempty"`
	// The key identifying the metric.
	Key string `json:"key,omitempty"`
	// The ID of the logged model or registered model version associated with
	// the metric, if applicable.
	ModelId string `json:"model_id,omitempty"`
	// The ID of the run containing the metric.
	RunId string `json:"run_id,omitempty"`
	// The step at which the metric was logged.
	Step int64 `json:"step,omitempty"`
	// The timestamp at which the metric was recorded.
	Timestamp int64 `json:"timestamp,omitempty"`
	// The value of the metric.
	Value float64 `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func metricFromPb(pb *metricPb) (*Metric, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Metric{}
	st.DatasetDigest = pb.DatasetDigest
	st.DatasetName = pb.DatasetName
	st.Key = pb.Key
	st.ModelId = pb.ModelId
	st.RunId = pb.RunId
	st.Step = pb.Step
	st.Timestamp = pb.Timestamp
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *metricPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st metricPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type Model struct {
	// Timestamp recorded when this `registered_model` was created.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Description of this `registered_model`.
	// Wire name: 'description'
	Description string
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	// Wire name: 'latest_versions'
	LatestVersions []ModelVersion
	// Unique name for the model.
	// Wire name: 'name'
	Name string
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	// Wire name: 'tags'
	Tags []ModelTag
	// User that created this `registered_model`
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func modelToPb(st *Model) (*modelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelPb{}
	pb.CreationTimestamp = st.CreationTimestamp

	pb.Description = st.Description

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	var latestVersionsPb []modelVersionPb
	for _, item := range st.LatestVersions {
		itemPb, err := modelVersionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			latestVersionsPb = append(latestVersionsPb, *itemPb)
		}
	}
	pb.LatestVersions = latestVersionsPb

	pb.Name = st.Name

	var tagsPb []modelTagPb
	for _, item := range st.Tags {
		itemPb, err := modelTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Model) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Model) MarshalJSON() ([]byte, error) {
	pb, err := modelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type modelPb struct {
	// Timestamp recorded when this `registered_model` was created.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Description of this `registered_model`.
	Description string `json:"description,omitempty"`
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	LatestVersions []modelVersionPb `json:"latest_versions,omitempty"`
	// Unique name for the model.
	Name string `json:"name,omitempty"`
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	Tags []modelTagPb `json:"tags,omitempty"`
	// User that created this `registered_model`
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelFromPb(pb *modelPb) (*Model, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Model{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Description = pb.Description
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp

	var latestVersionsField []ModelVersion
	for _, itemPb := range pb.LatestVersions {
		item, err := modelVersionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			latestVersionsField = append(latestVersionsField, *item)
		}
	}
	st.LatestVersions = latestVersionsField
	st.Name = pb.Name

	var tagsField []ModelTag
	for _, itemPb := range pb.Tags {
		item, err := modelTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// User-specified description for the object.
	// Wire name: 'description'
	Description string
	// Unique identifier for the object.
	// Wire name: 'id'
	Id string
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Array of model versions, each the latest version for its stage.
	// Wire name: 'latest_versions'
	LatestVersions []ModelVersion
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel
	// Array of tags associated with the model.
	// Wire name: 'tags'
	Tags []ModelTag
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func modelDatabricksToPb(st *ModelDatabricks) (*modelDatabricksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelDatabricksPb{}
	pb.CreationTimestamp = st.CreationTimestamp

	pb.Description = st.Description

	pb.Id = st.Id

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	var latestVersionsPb []modelVersionPb
	for _, item := range st.LatestVersions {
		itemPb, err := modelVersionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			latestVersionsPb = append(latestVersionsPb, *itemPb)
		}
	}
	pb.LatestVersions = latestVersionsPb

	pb.Name = st.Name

	pb.PermissionLevel = st.PermissionLevel

	var tagsPb []modelTagPb
	for _, item := range st.Tags {
		itemPb, err := modelTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ModelDatabricks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelDatabricksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelDatabricksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelDatabricks) MarshalJSON() ([]byte, error) {
	pb, err := modelDatabricksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type modelDatabricksPb struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// User-specified description for the object.
	Description string `json:"description,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Array of model versions, each the latest version for its stage.
	LatestVersions []modelVersionPb `json:"latest_versions,omitempty"`
	// Name of the model.
	Name string `json:"name,omitempty"`
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// Array of tags associated with the model.
	Tags []modelTagPb `json:"tags,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelDatabricksFromPb(pb *modelDatabricksPb) (*ModelDatabricks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelDatabricks{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Description = pb.Description
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp

	var latestVersionsField []ModelVersion
	for _, itemPb := range pb.LatestVersions {
		item, err := modelVersionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			latestVersionsField = append(latestVersionsField, *item)
		}
	}
	st.LatestVersions = latestVersionsField
	st.Name = pb.Name
	st.PermissionLevel = pb.PermissionLevel

	var tagsField []ModelTag
	for _, itemPb := range pb.Tags {
		item, err := modelTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelDatabricksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelDatabricksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Represents a LoggedModel or Registered Model Version input to a Run.
type ModelInput struct {
	// The unique identifier of the model.
	// Wire name: 'model_id'
	ModelId string
}

func modelInputToPb(st *ModelInput) (*modelInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelInputPb{}
	pb.ModelId = st.ModelId

	return pb, nil
}

func (st *ModelInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelInput) MarshalJSON() ([]byte, error) {
	pb, err := modelInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type modelInputPb struct {
	// The unique identifier of the model.
	ModelId string `json:"model_id"`
}

func modelInputFromPb(pb *modelInputPb) (*ModelInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelInput{}
	st.ModelId = pb.ModelId

	return st, nil
}

type ModelTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func modelTagToPb(st *ModelTag) (*modelTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelTagPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ModelTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelTag) MarshalJSON() ([]byte, error) {
	pb, err := modelTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type modelTagPb struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelTagFromPb(pb *modelTagPb) (*ModelTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelVersion struct {
	// Timestamp recorded when this `model_version` was created.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Current stage for this `model_version`.
	// Wire name: 'current_stage'
	CurrentStage string
	// Description of this `model_version`.
	// Wire name: 'description'
	Description string
	// Timestamp recorded when metadata for this `model_version` was last
	// updated.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Unique name of the model
	// Wire name: 'name'
	Name string
	// MLflow run ID used when creating `model_version`, if `source` was
	// generated by an experiment run stored in MLflow tracking server.
	// Wire name: 'run_id'
	RunId string
	// Run Link: Direct link to the run that generated this version
	// Wire name: 'run_link'
	RunLink string
	// URI indicating the location of the source model artifacts, used when
	// creating `model_version`
	// Wire name: 'source'
	Source string
	// Current status of `model_version`
	// Wire name: 'status'
	Status ModelVersionStatus
	// Details on current `status`, if it is pending or failed.
	// Wire name: 'status_message'
	StatusMessage string
	// Tags: Additional metadata key-value pairs for this `model_version`.
	// Wire name: 'tags'
	Tags []ModelVersionTag
	// User that created this `model_version`.
	// Wire name: 'user_id'
	UserId string
	// Model's version number.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func modelVersionToPb(st *ModelVersion) (*modelVersionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelVersionPb{}
	pb.CreationTimestamp = st.CreationTimestamp

	pb.CurrentStage = st.CurrentStage

	pb.Description = st.Description

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	pb.Name = st.Name

	pb.RunId = st.RunId

	pb.RunLink = st.RunLink

	pb.Source = st.Source

	pb.Status = st.Status

	pb.StatusMessage = st.StatusMessage

	var tagsPb []modelVersionTagPb
	for _, item := range st.Tags {
		itemPb, err := modelVersionTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.UserId = st.UserId

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ModelVersion) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelVersionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelVersionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelVersion) MarshalJSON() ([]byte, error) {
	pb, err := modelVersionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type modelVersionPb struct {
	// Timestamp recorded when this `model_version` was created.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Current stage for this `model_version`.
	CurrentStage string `json:"current_stage,omitempty"`
	// Description of this `model_version`.
	Description string `json:"description,omitempty"`
	// Timestamp recorded when metadata for this `model_version` was last
	// updated.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Unique name of the model
	Name string `json:"name,omitempty"`
	// MLflow run ID used when creating `model_version`, if `source` was
	// generated by an experiment run stored in MLflow tracking server.
	RunId string `json:"run_id,omitempty"`
	// Run Link: Direct link to the run that generated this version
	RunLink string `json:"run_link,omitempty"`
	// URI indicating the location of the source model artifacts, used when
	// creating `model_version`
	Source string `json:"source,omitempty"`
	// Current status of `model_version`
	Status ModelVersionStatus `json:"status,omitempty"`
	// Details on current `status`, if it is pending or failed.
	StatusMessage string `json:"status_message,omitempty"`
	// Tags: Additional metadata key-value pairs for this `model_version`.
	Tags []modelVersionTagPb `json:"tags,omitempty"`
	// User that created this `model_version`.
	UserId string `json:"user_id,omitempty"`
	// Model's version number.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelVersionFromPb(pb *modelVersionPb) (*ModelVersion, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersion{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.CurrentStage = pb.CurrentStage
	st.Description = pb.Description
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Name = pb.Name
	st.RunId = pb.RunId
	st.RunLink = pb.RunLink
	st.Source = pb.Source
	st.Status = pb.Status
	st.StatusMessage = pb.StatusMessage

	var tagsField []ModelVersionTag
	for _, itemPb := range pb.Tags {
		item, err := modelVersionTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.UserId = pb.UserId
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelVersionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelVersionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type ModelVersionDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Stage of the model version. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'current_stage'
	CurrentStage Stage
	// User-specified description for the object.
	// Wire name: 'description'
	Description string
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	// Wire name: 'run_id'
	RunId string
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	// Wire name: 'run_link'
	RunLink string
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	// Wire name: 'source'
	Source string
	// The status of the model version. Valid values are: *
	// `PENDING_REGISTRATION`: Request to register a new model version is
	// pending as server performs background tasks.
	//
	// * `FAILED_REGISTRATION`: Request to register a new model version has
	// failed.
	//
	// * `READY`: Model version is ready for use.
	// Wire name: 'status'
	Status Status
	// Details on the current status, for example why registration failed.
	// Wire name: 'status_message'
	StatusMessage string
	// Array of tags that are associated with the model version.
	// Wire name: 'tags'
	Tags []ModelVersionTag
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string
	// Version of the model.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func modelVersionDatabricksToPb(st *ModelVersionDatabricks) (*modelVersionDatabricksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelVersionDatabricksPb{}
	pb.CreationTimestamp = st.CreationTimestamp

	pb.CurrentStage = st.CurrentStage

	pb.Description = st.Description

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	pb.Name = st.Name

	pb.PermissionLevel = st.PermissionLevel

	pb.RunId = st.RunId

	pb.RunLink = st.RunLink

	pb.Source = st.Source

	pb.Status = st.Status

	pb.StatusMessage = st.StatusMessage

	var tagsPb []modelVersionTagPb
	for _, item := range st.Tags {
		itemPb, err := modelVersionTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	pb.UserId = st.UserId

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ModelVersionDatabricks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelVersionDatabricksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelVersionDatabricksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelVersionDatabricks) MarshalJSON() ([]byte, error) {
	pb, err := modelVersionDatabricksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type modelVersionDatabricksPb struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Stage of the model version. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	CurrentStage Stage `json:"current_stage,omitempty"`
	// User-specified description for the object.
	Description string `json:"description,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Name of the model.
	Name string `json:"name,omitempty"`
	// Permission level of the requesting user on the object. For what is
	// allowed at each level, see [MLflow Model permissions](..).
	PermissionLevel PermissionLevel `json:"permission_level,omitempty"`
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	RunId string `json:"run_id,omitempty"`
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	RunLink string `json:"run_link,omitempty"`
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	Source string `json:"source,omitempty"`
	// The status of the model version. Valid values are: *
	// `PENDING_REGISTRATION`: Request to register a new model version is
	// pending as server performs background tasks.
	//
	// * `FAILED_REGISTRATION`: Request to register a new model version has
	// failed.
	//
	// * `READY`: Model version is ready for use.
	Status Status `json:"status,omitempty"`
	// Details on the current status, for example why registration failed.
	StatusMessage string `json:"status_message,omitempty"`
	// Array of tags that are associated with the model version.
	Tags []modelVersionTagPb `json:"tags,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`
	// Version of the model.
	Version string `json:"version,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelVersionDatabricksFromPb(pb *modelVersionDatabricksPb) (*ModelVersionDatabricks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersionDatabricks{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.CurrentStage = pb.CurrentStage
	st.Description = pb.Description
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Name = pb.Name
	st.PermissionLevel = pb.PermissionLevel
	st.RunId = pb.RunId
	st.RunLink = pb.RunLink
	st.Source = pb.Source
	st.Status = pb.Status
	st.StatusMessage = pb.StatusMessage

	var tagsField []ModelVersionTag
	for _, itemPb := range pb.Tags {
		item, err := modelVersionTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.UserId = pb.UserId
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelVersionDatabricksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelVersionDatabricksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Current status of `model_version`
type ModelVersionStatus string
type modelVersionStatusPb string

const ModelVersionStatusFailedRegistration ModelVersionStatus = `FAILED_REGISTRATION`

const ModelVersionStatusPendingRegistration ModelVersionStatus = `PENDING_REGISTRATION`

const ModelVersionStatusReady ModelVersionStatus = `READY`

// String representation for [fmt.Print]
func (f *ModelVersionStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ModelVersionStatus) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = ModelVersionStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Type always returns ModelVersionStatus to satisfy [pflag.Value] interface
func (f *ModelVersionStatus) Type() string {
	return "ModelVersionStatus"
}

func modelVersionStatusToPb(st *ModelVersionStatus) (*modelVersionStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := modelVersionStatusPb(*st)
	return &pb, nil
}

func modelVersionStatusFromPb(pb *modelVersionStatusPb) (*ModelVersionStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := ModelVersionStatus(*pb)
	return &st, nil
}

type ModelVersionTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func modelVersionTagToPb(st *ModelVersionTag) (*modelVersionTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &modelVersionTagPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *ModelVersionTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &modelVersionTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := modelVersionTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st ModelVersionTag) MarshalJSON() ([]byte, error) {
	pb, err := modelVersionTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type modelVersionTagPb struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func modelVersionTagFromPb(pb *modelVersionTagPb) (*ModelVersionTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersionTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *modelVersionTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st modelVersionTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Param associated with a run.
type Param struct {
	// Key identifying this param.
	// Wire name: 'key'
	Key string
	// Value associated with this param.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func paramToPb(st *Param) (*paramPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &paramPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *Param) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &paramPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := paramFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Param) MarshalJSON() ([]byte, error) {
	pb, err := paramToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type paramPb struct {
	// Key identifying this param.
	Key string `json:"key,omitempty"`
	// Value associated with this param.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func paramFromPb(pb *paramPb) (*Param, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Param{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *paramPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st paramPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
type PermissionLevel string
type permissionLevelPb string

const PermissionLevelCanEdit PermissionLevel = `CAN_EDIT`

const PermissionLevelCanManage PermissionLevel = `CAN_MANAGE`

const PermissionLevelCanManageProductionVersions PermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const PermissionLevelCanManageStagingVersions PermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const PermissionLevelCanRead PermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *PermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

func permissionLevelToPb(st *PermissionLevel) (*permissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := permissionLevelPb(*st)
	return &pb, nil
}

func permissionLevelFromPb(pb *permissionLevelPb) (*PermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PermissionLevel(*pb)
	return &st, nil
}

type RegisteredModelAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func registeredModelAccessControlRequestToPb(st *RegisteredModelAccessControlRequest) (*registeredModelAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelAccessControlRequestPb{}
	pb.GroupName = st.GroupName

	pb.PermissionLevel = st.PermissionLevel

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RegisteredModelAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type registeredModelAccessControlRequestPb struct {
	// name of the group
	GroupName string `json:"group_name,omitempty"`
	// Permission level
	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`
	// application ID of a service principal
	ServicePrincipalName string `json:"service_principal_name,omitempty"`
	// name of the user
	UserName string `json:"user_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelAccessControlRequestFromPb(pb *registeredModelAccessControlRequestPb) (*RegisteredModelAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAccessControlRequest{}
	st.GroupName = pb.GroupName
	st.PermissionLevel = pb.PermissionLevel
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelAccessControlRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelAccessControlRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []RegisteredModelPermission
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string
	// name of the group
	// Wire name: 'group_name'
	GroupName string
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string
	// name of the user
	// Wire name: 'user_name'
	UserName string

	ForceSendFields []string `tf:"-"`
}

func registeredModelAccessControlResponseToPb(st *RegisteredModelAccessControlResponse) (*registeredModelAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelAccessControlResponsePb{}

	var allPermissionsPb []registeredModelPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := registeredModelPermissionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			allPermissionsPb = append(allPermissionsPb, *itemPb)
		}
	}
	pb.AllPermissions = allPermissionsPb

	pb.DisplayName = st.DisplayName

	pb.GroupName = st.GroupName

	pb.ServicePrincipalName = st.ServicePrincipalName

	pb.UserName = st.UserName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RegisteredModelAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type registeredModelAccessControlResponsePb struct {
	// All permissions.
	AllPermissions []registeredModelPermissionPb `json:"all_permissions,omitempty"`
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

func registeredModelAccessControlResponseFromPb(pb *registeredModelAccessControlResponsePb) (*RegisteredModelAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAccessControlResponse{}

	var allPermissionsField []RegisteredModelPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := registeredModelPermissionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			allPermissionsField = append(allPermissionsField, *item)
		}
	}
	st.AllPermissions = allPermissionsField
	st.DisplayName = pb.DisplayName
	st.GroupName = pb.GroupName
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelAccessControlResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelAccessControlResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelPermission struct {

	// Wire name: 'inherited'
	Inherited bool

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func registeredModelPermissionToPb(st *RegisteredModelPermission) (*registeredModelPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelPermissionPb{}
	pb.Inherited = st.Inherited

	pb.InheritedFromObject = st.InheritedFromObject

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RegisteredModelPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelPermission) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type registeredModelPermissionPb struct {
	Inherited bool `json:"inherited,omitempty"`

	InheritedFromObject []string `json:"inherited_from_object,omitempty"`
	// Permission level
	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelPermissionFromPb(pb *registeredModelPermissionPb) (*RegisteredModelPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelPermissionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelPermissionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Permission level
type RegisteredModelPermissionLevel string
type registeredModelPermissionLevelPb string

const RegisteredModelPermissionLevelCanEdit RegisteredModelPermissionLevel = `CAN_EDIT`

const RegisteredModelPermissionLevelCanManage RegisteredModelPermissionLevel = `CAN_MANAGE`

const RegisteredModelPermissionLevelCanManageProductionVersions RegisteredModelPermissionLevel = `CAN_MANAGE_PRODUCTION_VERSIONS`

const RegisteredModelPermissionLevelCanManageStagingVersions RegisteredModelPermissionLevel = `CAN_MANAGE_STAGING_VERSIONS`

const RegisteredModelPermissionLevelCanRead RegisteredModelPermissionLevel = `CAN_READ`

// String representation for [fmt.Print]
func (f *RegisteredModelPermissionLevel) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegisteredModelPermissionLevel) Set(v string) error {
	switch v {
	case `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*f = RegisteredModelPermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Type always returns RegisteredModelPermissionLevel to satisfy [pflag.Value] interface
func (f *RegisteredModelPermissionLevel) Type() string {
	return "RegisteredModelPermissionLevel"
}

func registeredModelPermissionLevelToPb(st *RegisteredModelPermissionLevel) (*registeredModelPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := registeredModelPermissionLevelPb(*st)
	return &pb, nil
}

func registeredModelPermissionLevelFromPb(pb *registeredModelPermissionLevelPb) (*RegisteredModelPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := RegisteredModelPermissionLevel(*pb)
	return &st, nil
}

type RegisteredModelPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []RegisteredModelAccessControlResponse

	// Wire name: 'object_id'
	ObjectId string

	// Wire name: 'object_type'
	ObjectType string

	ForceSendFields []string `tf:"-"`
}

func registeredModelPermissionsToPb(st *RegisteredModelPermissions) (*registeredModelPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelPermissionsPb{}

	var accessControlListPb []registeredModelAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := registeredModelAccessControlResponseToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	pb.ObjectId = st.ObjectId

	pb.ObjectType = st.ObjectType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RegisteredModelPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelPermissions) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type registeredModelPermissionsPb struct {
	AccessControlList []registeredModelAccessControlResponsePb `json:"access_control_list,omitempty"`

	ObjectId string `json:"object_id,omitempty"`

	ObjectType string `json:"object_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelPermissionsFromPb(pb *registeredModelPermissionsPb) (*RegisteredModelPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermissions{}

	var accessControlListField []RegisteredModelAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := registeredModelAccessControlResponseFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.ObjectId = pb.ObjectId
	st.ObjectType = pb.ObjectType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelPermissionsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelPermissionsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelPermissionsDescription struct {

	// Wire name: 'description'
	Description string
	// Permission level
	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel

	ForceSendFields []string `tf:"-"`
}

func registeredModelPermissionsDescriptionToPb(st *RegisteredModelPermissionsDescription) (*registeredModelPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelPermissionsDescriptionPb{}
	pb.Description = st.Description

	pb.PermissionLevel = st.PermissionLevel

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RegisteredModelPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type registeredModelPermissionsDescriptionPb struct {
	Description string `json:"description,omitempty"`
	// Permission level
	PermissionLevel RegisteredModelPermissionLevel `json:"permission_level,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registeredModelPermissionsDescriptionFromPb(pb *registeredModelPermissionsDescriptionPb) (*RegisteredModelPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermissionsDescription{}
	st.Description = pb.Description
	st.PermissionLevel = pb.PermissionLevel

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registeredModelPermissionsDescriptionPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registeredModelPermissionsDescriptionPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegisteredModelPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []RegisteredModelAccessControlRequest
	// The registered model for which to get or manage permissions.
	// Wire name: 'registered_model_id'
	RegisteredModelId string `tf:"-"`
}

func registeredModelPermissionsRequestToPb(st *RegisteredModelPermissionsRequest) (*registeredModelPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registeredModelPermissionsRequestPb{}

	var accessControlListPb []registeredModelAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := registeredModelAccessControlRequestToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			accessControlListPb = append(accessControlListPb, *itemPb)
		}
	}
	pb.AccessControlList = accessControlListPb

	pb.RegisteredModelId = st.RegisteredModelId

	return pb, nil
}

func (st *RegisteredModelPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registeredModelPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registeredModelPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegisteredModelPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := registeredModelPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type registeredModelPermissionsRequestPb struct {
	AccessControlList []registeredModelAccessControlRequestPb `json:"access_control_list,omitempty"`
	// The registered model for which to get or manage permissions.
	RegisteredModelId string `json:"-" url:"-"`
}

func registeredModelPermissionsRequestFromPb(pb *registeredModelPermissionsRequestPb) (*RegisteredModelPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermissionsRequest{}

	var accessControlListField []RegisteredModelAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := registeredModelAccessControlRequestFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			accessControlListField = append(accessControlListField, *item)
		}
	}
	st.AccessControlList = accessControlListField
	st.RegisteredModelId = pb.RegisteredModelId

	return st, nil
}

type RegistryWebhook struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// User-specified description for the webhook.
	// Wire name: 'description'
	Description string
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	// Wire name: 'events'
	Events []RegistryWebhookEvent

	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpecWithoutSecret
	// Webhook ID
	// Wire name: 'id'
	Id string

	// Wire name: 'job_spec'
	JobSpec *JobSpecWithoutSecret
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64
	// Name of the model whose events would trigger this webhook.
	// Wire name: 'model_name'
	ModelName string
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	// Wire name: 'status'
	Status RegistryWebhookStatus

	ForceSendFields []string `tf:"-"`
}

func registryWebhookToPb(st *RegistryWebhook) (*registryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &registryWebhookPb{}
	pb.CreationTimestamp = st.CreationTimestamp

	pb.Description = st.Description

	pb.Events = st.Events

	httpUrlSpecPb, err := httpUrlSpecWithoutSecretToPb(st.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecPb != nil {
		pb.HttpUrlSpec = httpUrlSpecPb
	}

	pb.Id = st.Id

	jobSpecPb, err := jobSpecWithoutSecretToPb(st.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecPb != nil {
		pb.JobSpec = jobSpecPb
	}

	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	pb.ModelName = st.ModelName

	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &registryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := registryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := registryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type registryWebhookPb struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events []RegistryWebhookEvent `json:"events,omitempty"`

	HttpUrlSpec *httpUrlSpecWithoutSecretPb `json:"http_url_spec,omitempty"`
	// Webhook ID
	Id string `json:"id,omitempty"`

	JobSpec *jobSpecWithoutSecretPb `json:"job_spec,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Name of the model whose events would trigger this webhook.
	ModelName string `json:"model_name,omitempty"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func registryWebhookFromPb(pb *registryWebhookPb) (*RegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegistryWebhook{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Description = pb.Description
	st.Events = pb.Events
	httpUrlSpecField, err := httpUrlSpecWithoutSecretFromPb(pb.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecField != nil {
		st.HttpUrlSpec = httpUrlSpecField
	}
	st.Id = pb.Id
	jobSpecField, err := jobSpecWithoutSecretFromPb(pb.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecField != nil {
		st.JobSpec = jobSpecField
	}
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.ModelName = pb.ModelName
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *registryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st registryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RegistryWebhookEvent string
type registryWebhookEventPb string

const RegistryWebhookEventCommentCreated RegistryWebhookEvent = `COMMENT_CREATED`

const RegistryWebhookEventModelVersionCreated RegistryWebhookEvent = `MODEL_VERSION_CREATED`

const RegistryWebhookEventModelVersionTagSet RegistryWebhookEvent = `MODEL_VERSION_TAG_SET`

const RegistryWebhookEventModelVersionTransitionedStage RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_STAGE`

const RegistryWebhookEventModelVersionTransitionedToArchived RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`

const RegistryWebhookEventModelVersionTransitionedToProduction RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`

const RegistryWebhookEventModelVersionTransitionedToStaging RegistryWebhookEvent = `MODEL_VERSION_TRANSITIONED_TO_STAGING`

const RegistryWebhookEventRegisteredModelCreated RegistryWebhookEvent = `REGISTERED_MODEL_CREATED`

const RegistryWebhookEventTransitionRequestCreated RegistryWebhookEvent = `TRANSITION_REQUEST_CREATED`

const RegistryWebhookEventTransitionRequestToArchivedCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`

const RegistryWebhookEventTransitionRequestToProductionCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`

const RegistryWebhookEventTransitionRequestToStagingCreated RegistryWebhookEvent = `TRANSITION_REQUEST_TO_STAGING_CREATED`

// String representation for [fmt.Print]
func (f *RegistryWebhookEvent) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryWebhookEvent) Set(v string) error {
	switch v {
	case `COMMENT_CREATED`, `MODEL_VERSION_CREATED`, `MODEL_VERSION_TAG_SET`, `MODEL_VERSION_TRANSITIONED_STAGE`, `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`, `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`, `MODEL_VERSION_TRANSITIONED_TO_STAGING`, `REGISTERED_MODEL_CREATED`, `TRANSITION_REQUEST_CREATED`, `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`, `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`, `TRANSITION_REQUEST_TO_STAGING_CREATED`:
		*f = RegistryWebhookEvent(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "COMMENT_CREATED", "MODEL_VERSION_CREATED", "MODEL_VERSION_TAG_SET", "MODEL_VERSION_TRANSITIONED_STAGE", "MODEL_VERSION_TRANSITIONED_TO_ARCHIVED", "MODEL_VERSION_TRANSITIONED_TO_PRODUCTION", "MODEL_VERSION_TRANSITIONED_TO_STAGING", "REGISTERED_MODEL_CREATED", "TRANSITION_REQUEST_CREATED", "TRANSITION_REQUEST_TO_ARCHIVED_CREATED", "TRANSITION_REQUEST_TO_PRODUCTION_CREATED", "TRANSITION_REQUEST_TO_STAGING_CREATED"`, v)
	}
}

// Type always returns RegistryWebhookEvent to satisfy [pflag.Value] interface
func (f *RegistryWebhookEvent) Type() string {
	return "RegistryWebhookEvent"
}

func registryWebhookEventToPb(st *RegistryWebhookEvent) (*registryWebhookEventPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := registryWebhookEventPb(*st)
	return &pb, nil
}

func registryWebhookEventFromPb(pb *registryWebhookEventPb) (*RegistryWebhookEvent, error) {
	if pb == nil {
		return nil, nil
	}
	st := RegistryWebhookEvent(*pb)
	return &st, nil
}

// Enable or disable triggering the webhook, or put the webhook into test mode.
// The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an associated
// event happens.
//
// * `DISABLED`: Webhook is not triggered.
//
// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is not
// triggered on a real event.
type RegistryWebhookStatus string
type registryWebhookStatusPb string

// Webhook is triggered when an associated event happens.
const RegistryWebhookStatusActive RegistryWebhookStatus = `ACTIVE`

// Webhook is not triggered.
const RegistryWebhookStatusDisabled RegistryWebhookStatus = `DISABLED`

// Webhook can be triggered through the test endpoint, but is not triggered on a
// real event.
const RegistryWebhookStatusTestMode RegistryWebhookStatus = `TEST_MODE`

// String representation for [fmt.Print]
func (f *RegistryWebhookStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryWebhookStatus) Set(v string) error {
	switch v {
	case `ACTIVE`, `DISABLED`, `TEST_MODE`:
		*f = RegistryWebhookStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE", "DISABLED", "TEST_MODE"`, v)
	}
}

// Type always returns RegistryWebhookStatus to satisfy [pflag.Value] interface
func (f *RegistryWebhookStatus) Type() string {
	return "RegistryWebhookStatus"
}

func registryWebhookStatusToPb(st *RegistryWebhookStatus) (*registryWebhookStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := registryWebhookStatusPb(*st)
	return &pb, nil
}

func registryWebhookStatusFromPb(pb *registryWebhookStatusPb) (*RegistryWebhookStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := RegistryWebhookStatus(*pb)
	return &st, nil
}

type RejectTransitionRequest struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'stage'
	Stage Stage
	// Version of the model.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func rejectTransitionRequestToPb(st *RejectTransitionRequest) (*rejectTransitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rejectTransitionRequestPb{}
	pb.Comment = st.Comment

	pb.Name = st.Name

	pb.Stage = st.Stage

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RejectTransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rejectTransitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rejectTransitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RejectTransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := rejectTransitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type rejectTransitionRequestPb struct {
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage Stage `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func rejectTransitionRequestFromPb(pb *rejectTransitionRequestPb) (*RejectTransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RejectTransitionRequest{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *rejectTransitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st rejectTransitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RejectTransitionRequestResponse struct {
	// Activity recorded for the action.
	// Wire name: 'activity'
	Activity *Activity
}

func rejectTransitionRequestResponseToPb(st *RejectTransitionRequestResponse) (*rejectTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &rejectTransitionRequestResponsePb{}
	activityPb, err := activityToPb(st.Activity)
	if err != nil {
		return nil, err
	}
	if activityPb != nil {
		pb.Activity = activityPb
	}

	return pb, nil
}

func (st *RejectTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &rejectTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := rejectTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RejectTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := rejectTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type rejectTransitionRequestResponsePb struct {
	// Activity recorded for the action.
	Activity *activityPb `json:"activity,omitempty"`
}

func rejectTransitionRequestResponseFromPb(pb *rejectTransitionRequestResponsePb) (*RejectTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RejectTransitionRequestResponse{}
	activityField, err := activityFromPb(pb.Activity)
	if err != nil {
		return nil, err
	}
	if activityField != nil {
		st.Activity = activityField
	}

	return st, nil
}

type RenameModelRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string
	// If provided, updates the name for this `registered_model`.
	// Wire name: 'new_name'
	NewName string

	ForceSendFields []string `tf:"-"`
}

func renameModelRequestToPb(st *RenameModelRequest) (*renameModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &renameModelRequestPb{}
	pb.Name = st.Name

	pb.NewName = st.NewName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RenameModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &renameModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := renameModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RenameModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := renameModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type renameModelRequestPb struct {
	// Registered model unique name identifier.
	Name string `json:"name"`
	// If provided, updates the name for this `registered_model`.
	NewName string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func renameModelRequestFromPb(pb *renameModelRequestPb) (*RenameModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RenameModelRequest{}
	st.Name = pb.Name
	st.NewName = pb.NewName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *renameModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st renameModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RenameModelResponse struct {

	// Wire name: 'registered_model'
	RegisteredModel *Model
}

func renameModelResponseToPb(st *RenameModelResponse) (*renameModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &renameModelResponsePb{}
	registeredModelPb, err := modelToPb(st.RegisteredModel)
	if err != nil {
		return nil, err
	}
	if registeredModelPb != nil {
		pb.RegisteredModel = registeredModelPb
	}

	return pb, nil
}

func (st *RenameModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &renameModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := renameModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RenameModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := renameModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type renameModelResponsePb struct {
	RegisteredModel *modelPb `json:"registered_model,omitempty"`
}

func renameModelResponseFromPb(pb *renameModelResponsePb) (*RenameModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RenameModelResponse{}
	registeredModelField, err := modelFromPb(pb.RegisteredModel)
	if err != nil {
		return nil, err
	}
	if registeredModelField != nil {
		st.RegisteredModel = registeredModelField
	}

	return st, nil
}

type RestoreExperiment struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
}

func restoreExperimentToPb(st *RestoreExperiment) (*restoreExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreExperimentPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func (st *RestoreExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreExperiment) MarshalJSON() ([]byte, error) {
	pb, err := restoreExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type restoreExperimentPb struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
}

func restoreExperimentFromPb(pb *restoreExperimentPb) (*RestoreExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreExperiment{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

type RestoreExperimentResponse struct {
}

func restoreExperimentResponseToPb(st *RestoreExperimentResponse) (*restoreExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreExperimentResponsePb{}

	return pb, nil
}

func (st *RestoreExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := restoreExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type restoreExperimentResponsePb struct {
}

func restoreExperimentResponseFromPb(pb *restoreExperimentResponsePb) (*RestoreExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreExperimentResponse{}

	return st, nil
}

type RestoreRun struct {
	// ID of the run to restore.
	// Wire name: 'run_id'
	RunId string
}

func restoreRunToPb(st *RestoreRun) (*restoreRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreRunPb{}
	pb.RunId = st.RunId

	return pb, nil
}

func (st *RestoreRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreRun) MarshalJSON() ([]byte, error) {
	pb, err := restoreRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type restoreRunPb struct {
	// ID of the run to restore.
	RunId string `json:"run_id"`
}

func restoreRunFromPb(pb *restoreRunPb) (*RestoreRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRun{}
	st.RunId = pb.RunId

	return st, nil
}

type RestoreRunResponse struct {
}

func restoreRunResponseToPb(st *RestoreRunResponse) (*restoreRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreRunResponsePb{}

	return pb, nil
}

func (st *RestoreRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := restoreRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type restoreRunResponsePb struct {
}

func restoreRunResponseFromPb(pb *restoreRunResponsePb) (*RestoreRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRunResponse{}

	return st, nil
}

type RestoreRuns struct {
	// The ID of the experiment containing the runs to restore.
	// Wire name: 'experiment_id'
	ExperimentId string
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	// Wire name: 'max_runs'
	MaxRuns int
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	// Wire name: 'min_timestamp_millis'
	MinTimestampMillis int64

	ForceSendFields []string `tf:"-"`
}

func restoreRunsToPb(st *RestoreRuns) (*restoreRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreRunsPb{}
	pb.ExperimentId = st.ExperimentId

	pb.MaxRuns = st.MaxRuns

	pb.MinTimestampMillis = st.MinTimestampMillis

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RestoreRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreRuns) MarshalJSON() ([]byte, error) {
	pb, err := restoreRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type restoreRunsPb struct {
	// The ID of the experiment containing the runs to restore.
	ExperimentId string `json:"experiment_id"`
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	MaxRuns int `json:"max_runs,omitempty"`
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	MinTimestampMillis int64 `json:"min_timestamp_millis"`

	ForceSendFields []string `json:"-" url:"-"`
}

func restoreRunsFromPb(pb *restoreRunsPb) (*RestoreRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRuns{}
	st.ExperimentId = pb.ExperimentId
	st.MaxRuns = pb.MaxRuns
	st.MinTimestampMillis = pb.MinTimestampMillis

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *restoreRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st restoreRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type RestoreRunsResponse struct {
	// The number of runs restored.
	// Wire name: 'runs_restored'
	RunsRestored int

	ForceSendFields []string `tf:"-"`
}

func restoreRunsResponseToPb(st *RestoreRunsResponse) (*restoreRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &restoreRunsResponsePb{}
	pb.RunsRestored = st.RunsRestored

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RestoreRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &restoreRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := restoreRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RestoreRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := restoreRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type restoreRunsResponsePb struct {
	// The number of runs restored.
	RunsRestored int `json:"runs_restored,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func restoreRunsResponseFromPb(pb *restoreRunsResponsePb) (*RestoreRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRunsResponse{}
	st.RunsRestored = pb.RunsRestored

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *restoreRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st restoreRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// A single run.
type Run struct {
	// Run data.
	// Wire name: 'data'
	Data *RunData
	// Run metadata.
	// Wire name: 'info'
	Info *RunInfo
	// Run inputs.
	// Wire name: 'inputs'
	Inputs *RunInputs
}

func runToPb(st *Run) (*runPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runPb{}
	dataPb, err := runDataToPb(st.Data)
	if err != nil {
		return nil, err
	}
	if dataPb != nil {
		pb.Data = dataPb
	}

	infoPb, err := runInfoToPb(st.Info)
	if err != nil {
		return nil, err
	}
	if infoPb != nil {
		pb.Info = infoPb
	}

	inputsPb, err := runInputsToPb(st.Inputs)
	if err != nil {
		return nil, err
	}
	if inputsPb != nil {
		pb.Inputs = inputsPb
	}

	return pb, nil
}

func (st *Run) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st Run) MarshalJSON() ([]byte, error) {
	pb, err := runToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runPb struct {
	// Run data.
	Data *runDataPb `json:"data,omitempty"`
	// Run metadata.
	Info *runInfoPb `json:"info,omitempty"`
	// Run inputs.
	Inputs *runInputsPb `json:"inputs,omitempty"`
}

func runFromPb(pb *runPb) (*Run, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Run{}
	dataField, err := runDataFromPb(pb.Data)
	if err != nil {
		return nil, err
	}
	if dataField != nil {
		st.Data = dataField
	}
	infoField, err := runInfoFromPb(pb.Info)
	if err != nil {
		return nil, err
	}
	if infoField != nil {
		st.Info = infoField
	}
	inputsField, err := runInputsFromPb(pb.Inputs)
	if err != nil {
		return nil, err
	}
	if inputsField != nil {
		st.Inputs = inputsField
	}

	return st, nil
}

// Run data (metrics, params, and tags).
type RunData struct {
	// Run metrics.
	// Wire name: 'metrics'
	Metrics []Metric
	// Run parameters.
	// Wire name: 'params'
	Params []Param
	// Additional metadata key-value pairs.
	// Wire name: 'tags'
	Tags []RunTag
}

func runDataToPb(st *RunData) (*runDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runDataPb{}

	var metricsPb []metricPb
	for _, item := range st.Metrics {
		itemPb, err := metricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			metricsPb = append(metricsPb, *itemPb)
		}
	}
	pb.Metrics = metricsPb

	var paramsPb []paramPb
	for _, item := range st.Params {
		itemPb, err := paramToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			paramsPb = append(paramsPb, *itemPb)
		}
	}
	pb.Params = paramsPb

	var tagsPb []runTagPb
	for _, item := range st.Tags {
		itemPb, err := runTagToPb(&item)
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

func (st *RunData) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runDataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runDataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunData) MarshalJSON() ([]byte, error) {
	pb, err := runDataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runDataPb struct {
	// Run metrics.
	Metrics []metricPb `json:"metrics,omitempty"`
	// Run parameters.
	Params []paramPb `json:"params,omitempty"`
	// Additional metadata key-value pairs.
	Tags []runTagPb `json:"tags,omitempty"`
}

func runDataFromPb(pb *runDataPb) (*RunData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunData{}

	var metricsField []Metric
	for _, itemPb := range pb.Metrics {
		item, err := metricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			metricsField = append(metricsField, *item)
		}
	}
	st.Metrics = metricsField

	var paramsField []Param
	for _, itemPb := range pb.Params {
		item, err := paramFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			paramsField = append(paramsField, *item)
		}
	}
	st.Params = paramsField

	var tagsField []RunTag
	for _, itemPb := range pb.Tags {
		item, err := runTagFromPb(&itemPb)
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

// Metadata of a single run.
type RunInfo struct {
	// URI of the directory where artifacts should be uploaded. This can be a
	// local path (starting with "/"), or a distributed file system (DFS) path,
	// like ``s3://bucket/directory`` or ``dbfs:/my/directory``. If not set, the
	// local ``./mlruns`` directory is chosen.
	// Wire name: 'artifact_uri'
	ArtifactUri string
	// Unix timestamp of when the run ended in milliseconds.
	// Wire name: 'end_time'
	EndTime int64
	// The experiment ID.
	// Wire name: 'experiment_id'
	ExperimentId string
	// Current life cycle stage of the experiment : OneOf("active", "deleted")
	// Wire name: 'lifecycle_stage'
	LifecycleStage string
	// Unique identifier for the run.
	// Wire name: 'run_id'
	RunId string
	// The name of the run.
	// Wire name: 'run_name'
	RunName string
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string
	// Unix timestamp of when the run started in milliseconds.
	// Wire name: 'start_time'
	StartTime int64
	// Current status of the run.
	// Wire name: 'status'
	Status RunInfoStatus
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func runInfoToPb(st *RunInfo) (*runInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runInfoPb{}
	pb.ArtifactUri = st.ArtifactUri

	pb.EndTime = st.EndTime

	pb.ExperimentId = st.ExperimentId

	pb.LifecycleStage = st.LifecycleStage

	pb.RunId = st.RunId

	pb.RunName = st.RunName

	pb.RunUuid = st.RunUuid

	pb.StartTime = st.StartTime

	pb.Status = st.Status

	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RunInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunInfo) MarshalJSON() ([]byte, error) {
	pb, err := runInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runInfoPb struct {
	// URI of the directory where artifacts should be uploaded. This can be a
	// local path (starting with "/"), or a distributed file system (DFS) path,
	// like ``s3://bucket/directory`` or ``dbfs:/my/directory``. If not set, the
	// local ``./mlruns`` directory is chosen.
	ArtifactUri string `json:"artifact_uri,omitempty"`
	// Unix timestamp of when the run ended in milliseconds.
	EndTime int64 `json:"end_time,omitempty"`
	// The experiment ID.
	ExperimentId string `json:"experiment_id,omitempty"`
	// Current life cycle stage of the experiment : OneOf("active", "deleted")
	LifecycleStage string `json:"lifecycle_stage,omitempty"`
	// Unique identifier for the run.
	RunId string `json:"run_id,omitempty"`
	// The name of the run.
	RunName string `json:"run_name,omitempty"`
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Unix timestamp of when the run started in milliseconds.
	StartTime int64 `json:"start_time,omitempty"`
	// Current status of the run.
	Status RunInfoStatus `json:"status,omitempty"`
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runInfoFromPb(pb *runInfoPb) (*RunInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunInfo{}
	st.ArtifactUri = pb.ArtifactUri
	st.EndTime = pb.EndTime
	st.ExperimentId = pb.ExperimentId
	st.LifecycleStage = pb.LifecycleStage
	st.RunId = pb.RunId
	st.RunName = pb.RunName
	st.RunUuid = pb.RunUuid
	st.StartTime = pb.StartTime
	st.Status = pb.Status
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runInfoPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runInfoPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Status of a run.
type RunInfoStatus string
type runInfoStatusPb string

const RunInfoStatusFailed RunInfoStatus = `FAILED`

const RunInfoStatusFinished RunInfoStatus = `FINISHED`

const RunInfoStatusKilled RunInfoStatus = `KILLED`

const RunInfoStatusRunning RunInfoStatus = `RUNNING`

const RunInfoStatusScheduled RunInfoStatus = `SCHEDULED`

// String representation for [fmt.Print]
func (f *RunInfoStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RunInfoStatus) Set(v string) error {
	switch v {
	case `FAILED`, `FINISHED`, `KILLED`, `RUNNING`, `SCHEDULED`:
		*f = RunInfoStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "FINISHED", "KILLED", "RUNNING", "SCHEDULED"`, v)
	}
}

// Type always returns RunInfoStatus to satisfy [pflag.Value] interface
func (f *RunInfoStatus) Type() string {
	return "RunInfoStatus"
}

func runInfoStatusToPb(st *RunInfoStatus) (*runInfoStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := runInfoStatusPb(*st)
	return &pb, nil
}

func runInfoStatusFromPb(pb *runInfoStatusPb) (*RunInfoStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := RunInfoStatus(*pb)
	return &st, nil
}

// Run inputs.
type RunInputs struct {
	// Run metrics.
	// Wire name: 'dataset_inputs'
	DatasetInputs []DatasetInput
	// **NOTE**: Experimental: This API field may change or be removed in a
	// future release without warning.
	//
	// Model inputs to the Run.
	// Wire name: 'model_inputs'
	ModelInputs []ModelInput
}

func runInputsToPb(st *RunInputs) (*runInputsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runInputsPb{}

	var datasetInputsPb []datasetInputPb
	for _, item := range st.DatasetInputs {
		itemPb, err := datasetInputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			datasetInputsPb = append(datasetInputsPb, *itemPb)
		}
	}
	pb.DatasetInputs = datasetInputsPb

	var modelInputsPb []modelInputPb
	for _, item := range st.ModelInputs {
		itemPb, err := modelInputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			modelInputsPb = append(modelInputsPb, *itemPb)
		}
	}
	pb.ModelInputs = modelInputsPb

	return pb, nil
}

func (st *RunInputs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runInputsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runInputsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunInputs) MarshalJSON() ([]byte, error) {
	pb, err := runInputsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runInputsPb struct {
	// Run metrics.
	DatasetInputs []datasetInputPb `json:"dataset_inputs,omitempty"`
	// **NOTE**: Experimental: This API field may change or be removed in a
	// future release without warning.
	//
	// Model inputs to the Run.
	ModelInputs []modelInputPb `json:"model_inputs,omitempty"`
}

func runInputsFromPb(pb *runInputsPb) (*RunInputs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunInputs{}

	var datasetInputsField []DatasetInput
	for _, itemPb := range pb.DatasetInputs {
		item, err := datasetInputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			datasetInputsField = append(datasetInputsField, *item)
		}
	}
	st.DatasetInputs = datasetInputsField

	var modelInputsField []ModelInput
	for _, itemPb := range pb.ModelInputs {
		item, err := modelInputFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			modelInputsField = append(modelInputsField, *item)
		}
	}
	st.ModelInputs = modelInputsField

	return st, nil
}

// Tag for a run.
type RunTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string
	// The tag value.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func runTagToPb(st *RunTag) (*runTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &runTagPb{}
	pb.Key = st.Key

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *RunTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &runTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := runTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st RunTag) MarshalJSON() ([]byte, error) {
	pb, err := runTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type runTagPb struct {
	// The tag key.
	Key string `json:"key,omitempty"`
	// The tag value.
	Value string `json:"value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func runTagFromPb(pb *runTagPb) (*RunTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *runTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st runTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchExperiments struct {
	// String representing a SQL filter condition (e.g. "name ILIKE
	// 'my-experiment%'")
	// Wire name: 'filter'
	Filter string
	// Maximum number of experiments desired. Max threshold is 3000.
	// Wire name: 'max_results'
	MaxResults int64
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional "DESC" or "ASC"
	// annotation, where "ASC" is the default. Tiebreaks are done by experiment
	// id DESC.
	// Wire name: 'order_by'
	OrderBy []string
	// Token indicating the page of experiments to fetch
	// Wire name: 'page_token'
	PageToken string
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	// Wire name: 'view_type'
	ViewType ViewType

	ForceSendFields []string `tf:"-"`
}

func searchExperimentsToPb(st *SearchExperiments) (*searchExperimentsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchExperimentsPb{}
	pb.Filter = st.Filter

	pb.MaxResults = st.MaxResults

	pb.OrderBy = st.OrderBy

	pb.PageToken = st.PageToken

	pb.ViewType = st.ViewType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SearchExperiments) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchExperimentsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchExperimentsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchExperiments) MarshalJSON() ([]byte, error) {
	pb, err := searchExperimentsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type searchExperimentsPb struct {
	// String representing a SQL filter condition (e.g. "name ILIKE
	// 'my-experiment%'")
	Filter string `json:"filter,omitempty"`
	// Maximum number of experiments desired. Max threshold is 3000.
	MaxResults int64 `json:"max_results,omitempty"`
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional "DESC" or "ASC"
	// annotation, where "ASC" is the default. Tiebreaks are done by experiment
	// id DESC.
	OrderBy []string `json:"order_by,omitempty"`
	// Token indicating the page of experiments to fetch
	PageToken string `json:"page_token,omitempty"`
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	ViewType ViewType `json:"view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchExperimentsFromPb(pb *searchExperimentsPb) (*SearchExperiments, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchExperiments{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken
	st.ViewType = pb.ViewType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchExperimentsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchExperimentsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	// Wire name: 'experiments'
	Experiments []Experiment
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func searchExperimentsResponseToPb(st *SearchExperimentsResponse) (*searchExperimentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchExperimentsResponsePb{}

	var experimentsPb []experimentPb
	for _, item := range st.Experiments {
		itemPb, err := experimentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			experimentsPb = append(experimentsPb, *itemPb)
		}
	}
	pb.Experiments = experimentsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SearchExperimentsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchExperimentsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchExperimentsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchExperimentsResponse) MarshalJSON() ([]byte, error) {
	pb, err := searchExperimentsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type searchExperimentsResponsePb struct {
	// Experiments that match the search criteria
	Experiments []experimentPb `json:"experiments,omitempty"`
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchExperimentsResponseFromPb(pb *searchExperimentsResponsePb) (*SearchExperimentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchExperimentsResponse{}

	var experimentsField []Experiment
	for _, itemPb := range pb.Experiments {
		item, err := experimentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			experimentsField = append(experimentsField, *item)
		}
	}
	st.Experiments = experimentsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchExperimentsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchExperimentsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Searches model versions
type SearchModelVersionsRequest struct {
	// String filter condition, like "name='my-model-name'". Must be a single
	// boolean condition, with string values wrapped in single quotes.
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Maximum number of models desired. Max threshold is 10K.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// List of columns to be ordered by including model name, version, stage
	// with an optional "DESC" or "ASC" annotation, where "ASC" is the default.
	// Tiebreaks are done by latest stage transition timestamp, followed by name
	// ASC, followed by version DESC.
	// Wire name: 'order_by'
	OrderBy []string `tf:"-"`
	// Pagination token to go to next page based on previous search query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func searchModelVersionsRequestToPb(st *SearchModelVersionsRequest) (*searchModelVersionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchModelVersionsRequestPb{}
	pb.Filter = st.Filter

	pb.MaxResults = st.MaxResults

	pb.OrderBy = st.OrderBy

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SearchModelVersionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchModelVersionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchModelVersionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchModelVersionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := searchModelVersionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type searchModelVersionsRequestPb struct {
	// String filter condition, like "name='my-model-name'". Must be a single
	// boolean condition, with string values wrapped in single quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of models desired. Max threshold is 10K.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// List of columns to be ordered by including model name, version, stage
	// with an optional "DESC" or "ASC" annotation, where "ASC" is the default.
	// Tiebreaks are done by latest stage transition timestamp, followed by name
	// ASC, followed by version DESC.
	OrderBy []string `json:"-" url:"order_by,omitempty"`
	// Pagination token to go to next page based on previous search query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchModelVersionsRequestFromPb(pb *searchModelVersionsRequestPb) (*SearchModelVersionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelVersionsRequest{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchModelVersionsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchModelVersionsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	// Wire name: 'model_versions'
	ModelVersions []ModelVersion
	// Pagination token to request next page of models for the same search
	// query.
	// Wire name: 'next_page_token'
	NextPageToken string

	ForceSendFields []string `tf:"-"`
}

func searchModelVersionsResponseToPb(st *SearchModelVersionsResponse) (*searchModelVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchModelVersionsResponsePb{}

	var modelVersionsPb []modelVersionPb
	for _, item := range st.ModelVersions {
		itemPb, err := modelVersionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			modelVersionsPb = append(modelVersionsPb, *itemPb)
		}
	}
	pb.ModelVersions = modelVersionsPb

	pb.NextPageToken = st.NextPageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SearchModelVersionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchModelVersionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchModelVersionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchModelVersionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := searchModelVersionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type searchModelVersionsResponsePb struct {
	// Models that match the search criteria
	ModelVersions []modelVersionPb `json:"model_versions,omitempty"`
	// Pagination token to request next page of models for the same search
	// query.
	NextPageToken string `json:"next_page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchModelVersionsResponseFromPb(pb *searchModelVersionsResponsePb) (*SearchModelVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelVersionsResponse{}

	var modelVersionsField []ModelVersion
	for _, itemPb := range pb.ModelVersions {
		item, err := modelVersionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			modelVersionsField = append(modelVersionsField, *item)
		}
	}
	st.ModelVersions = modelVersionsField
	st.NextPageToken = pb.NextPageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchModelVersionsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchModelVersionsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Search models
type SearchModelsRequest struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	// Wire name: 'max_results'
	MaxResults int `tf:"-"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	// Wire name: 'order_by'
	OrderBy []string `tf:"-"`
	// Pagination token to go to the next page based on a previous search query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	ForceSendFields []string `tf:"-"`
}

func searchModelsRequestToPb(st *SearchModelsRequest) (*searchModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchModelsRequestPb{}
	pb.Filter = st.Filter

	pb.MaxResults = st.MaxResults

	pb.OrderBy = st.OrderBy

	pb.PageToken = st.PageToken

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SearchModelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchModelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchModelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchModelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := searchModelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type searchModelsRequestPb struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	Filter string `json:"-" url:"filter,omitempty"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	MaxResults int `json:"-" url:"max_results,omitempty"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	OrderBy []string `json:"-" url:"order_by,omitempty"`
	// Pagination token to go to the next page based on a previous search query.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchModelsRequestFromPb(pb *searchModelsRequestPb) (*SearchModelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelsRequest{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchModelsRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchModelsRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	// Wire name: 'next_page_token'
	NextPageToken string
	// Registered Models that match the search criteria.
	// Wire name: 'registered_models'
	RegisteredModels []Model

	ForceSendFields []string `tf:"-"`
}

func searchModelsResponseToPb(st *SearchModelsResponse) (*searchModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchModelsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var registeredModelsPb []modelPb
	for _, item := range st.RegisteredModels {
		itemPb, err := modelToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			registeredModelsPb = append(registeredModelsPb, *itemPb)
		}
	}
	pb.RegisteredModels = registeredModelsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SearchModelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchModelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchModelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchModelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := searchModelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type searchModelsResponsePb struct {
	// Pagination token to request the next page of models.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Registered Models that match the search criteria.
	RegisteredModels []modelPb `json:"registered_models,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchModelsResponseFromPb(pb *searchModelsResponsePb) (*SearchModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelsResponse{}
	st.NextPageToken = pb.NextPageToken

	var registeredModelsField []Model
	for _, itemPb := range pb.RegisteredModels {
		item, err := modelFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			registeredModelsField = append(registeredModelsField, *item)
		}
	}
	st.RegisteredModels = registeredModelsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchModelsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchModelsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchRuns struct {
	// List of experiment IDs to search over.
	// Wire name: 'experiment_ids'
	ExperimentIds []string
	// A filter expression over params, metrics, and tags, that allows returning
	// a subset of runs. The syntax is a subset of SQL that supports ANDing
	// together binary operations between a param, metric, or tag and a
	// constant.
	//
	// Example: `metrics.rmse < 1 and params.model_class = 'LogisticRegression'`
	//
	// You can select columns with special characters (hyphen, space, period,
	// etc.) by using double quotes: `metrics."model class" = 'LinearRegression'
	// and tags."user-name" = 'Tomas'`
	//
	// Supported operators are `=`, `!=`, `>`, `>=`, `<`, and `<=`.
	// Wire name: 'filter'
	Filter string
	// Maximum number of runs desired. Max threshold is 50000
	// Wire name: 'max_results'
	MaxResults int
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional `"DESC"` or `"ASC"` annotation, where `"ASC"`
	// is the default. Example: `["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"]`. Tiebreaks are done by start_time `DESC` followed by
	// `run_id` for runs with the same start time (and this is the default
	// ordering criterion if order_by is not provided).
	// Wire name: 'order_by'
	OrderBy []string
	// Token for the current page of runs.
	// Wire name: 'page_token'
	PageToken string
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	// Wire name: 'run_view_type'
	RunViewType ViewType

	ForceSendFields []string `tf:"-"`
}

func searchRunsToPb(st *SearchRuns) (*searchRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchRunsPb{}
	pb.ExperimentIds = st.ExperimentIds

	pb.Filter = st.Filter

	pb.MaxResults = st.MaxResults

	pb.OrderBy = st.OrderBy

	pb.PageToken = st.PageToken

	pb.RunViewType = st.RunViewType

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SearchRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchRuns) MarshalJSON() ([]byte, error) {
	pb, err := searchRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type searchRunsPb struct {
	// List of experiment IDs to search over.
	ExperimentIds []string `json:"experiment_ids,omitempty"`
	// A filter expression over params, metrics, and tags, that allows returning
	// a subset of runs. The syntax is a subset of SQL that supports ANDing
	// together binary operations between a param, metric, or tag and a
	// constant.
	//
	// Example: `metrics.rmse < 1 and params.model_class = 'LogisticRegression'`
	//
	// You can select columns with special characters (hyphen, space, period,
	// etc.) by using double quotes: `metrics."model class" = 'LinearRegression'
	// and tags."user-name" = 'Tomas'`
	//
	// Supported operators are `=`, `!=`, `>`, `>=`, `<`, and `<=`.
	Filter string `json:"filter,omitempty"`
	// Maximum number of runs desired. Max threshold is 50000
	MaxResults int `json:"max_results,omitempty"`
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional `"DESC"` or `"ASC"` annotation, where `"ASC"`
	// is the default. Example: `["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"]`. Tiebreaks are done by start_time `DESC` followed by
	// `run_id` for runs with the same start time (and this is the default
	// ordering criterion if order_by is not provided).
	OrderBy []string `json:"order_by,omitempty"`
	// Token for the current page of runs.
	PageToken string `json:"page_token,omitempty"`
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	RunViewType ViewType `json:"run_view_type,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchRunsFromPb(pb *searchRunsPb) (*SearchRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchRuns{}
	st.ExperimentIds = pb.ExperimentIds
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken
	st.RunViewType = pb.RunViewType

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchRunsPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchRunsPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SearchRunsResponse struct {
	// Token for the next page of runs.
	// Wire name: 'next_page_token'
	NextPageToken string
	// Runs that match the search criteria.
	// Wire name: 'runs'
	Runs []Run

	ForceSendFields []string `tf:"-"`
}

func searchRunsResponseToPb(st *SearchRunsResponse) (*searchRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &searchRunsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var runsPb []runPb
	for _, item := range st.Runs {
		itemPb, err := runToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			runsPb = append(runsPb, *itemPb)
		}
	}
	pb.Runs = runsPb

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SearchRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &searchRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := searchRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SearchRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := searchRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type searchRunsResponsePb struct {
	// Token for the next page of runs.
	NextPageToken string `json:"next_page_token,omitempty"`
	// Runs that match the search criteria.
	Runs []runPb `json:"runs,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func searchRunsResponseFromPb(pb *searchRunsResponsePb) (*SearchRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchRunsResponse{}
	st.NextPageToken = pb.NextPageToken

	var runsField []Run
	for _, itemPb := range pb.Runs {
		item, err := runFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			runsField = append(runsField, *item)
		}
	}
	st.Runs = runsField

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *searchRunsResponsePb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st searchRunsResponsePb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SetExperimentTag struct {
	// ID of the experiment under which to log the tag. Must be provided.
	// Wire name: 'experiment_id'
	ExperimentId string
	// Name of the tag. Keys up to 250 bytes in size are supported.
	// Wire name: 'key'
	Key string
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	// Wire name: 'value'
	Value string
}

func setExperimentTagToPb(st *SetExperimentTag) (*setExperimentTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setExperimentTagPb{}
	pb.ExperimentId = st.ExperimentId

	pb.Key = st.Key

	pb.Value = st.Value

	return pb, nil
}

func (st *SetExperimentTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setExperimentTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setExperimentTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetExperimentTag) MarshalJSON() ([]byte, error) {
	pb, err := setExperimentTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setExperimentTagPb struct {
	// ID of the experiment under which to log the tag. Must be provided.
	ExperimentId string `json:"experiment_id"`
	// Name of the tag. Keys up to 250 bytes in size are supported.
	Key string `json:"key"`
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	Value string `json:"value"`
}

func setExperimentTagFromPb(pb *setExperimentTagPb) (*SetExperimentTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetExperimentTag{}
	st.ExperimentId = pb.ExperimentId
	st.Key = pb.Key
	st.Value = pb.Value

	return st, nil
}

type SetExperimentTagResponse struct {
}

func setExperimentTagResponseToPb(st *SetExperimentTagResponse) (*setExperimentTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setExperimentTagResponsePb{}

	return pb, nil
}

func (st *SetExperimentTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setExperimentTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setExperimentTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetExperimentTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := setExperimentTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setExperimentTagResponsePb struct {
}

func setExperimentTagResponseFromPb(pb *setExperimentTagResponsePb) (*SetExperimentTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetExperimentTagResponse{}

	return st, nil
}

type SetModelTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	// Wire name: 'key'
	Key string
	// Unique name of the model.
	// Wire name: 'name'
	Name string
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	// Wire name: 'value'
	Value string
}

func setModelTagRequestToPb(st *SetModelTagRequest) (*setModelTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setModelTagRequestPb{}
	pb.Key = st.Key

	pb.Name = st.Name

	pb.Value = st.Value

	return pb, nil
}

func (st *SetModelTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setModelTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setModelTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetModelTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := setModelTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setModelTagRequestPb struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key string `json:"key"`
	// Unique name of the model.
	Name string `json:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value string `json:"value"`
}

func setModelTagRequestFromPb(pb *setModelTagRequestPb) (*SetModelTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetModelTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name
	st.Value = pb.Value

	return st, nil
}

type SetModelTagResponse struct {
}

func setModelTagResponseToPb(st *SetModelTagResponse) (*setModelTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setModelTagResponsePb{}

	return pb, nil
}

func (st *SetModelTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setModelTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setModelTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetModelTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := setModelTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setModelTagResponsePb struct {
}

func setModelTagResponseFromPb(pb *setModelTagResponsePb) (*SetModelTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetModelTagResponse{}

	return st, nil
}

type SetModelVersionTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	// Wire name: 'key'
	Key string
	// Unique name of the model.
	// Wire name: 'name'
	Name string
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	// Wire name: 'value'
	Value string
	// Model version number.
	// Wire name: 'version'
	Version string
}

func setModelVersionTagRequestToPb(st *SetModelVersionTagRequest) (*setModelVersionTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setModelVersionTagRequestPb{}
	pb.Key = st.Key

	pb.Name = st.Name

	pb.Value = st.Value

	pb.Version = st.Version

	return pb, nil
}

func (st *SetModelVersionTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setModelVersionTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setModelVersionTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetModelVersionTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := setModelVersionTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setModelVersionTagRequestPb struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	Key string `json:"key"`
	// Unique name of the model.
	Name string `json:"name"`
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	Value string `json:"value"`
	// Model version number.
	Version string `json:"version"`
}

func setModelVersionTagRequestFromPb(pb *setModelVersionTagRequestPb) (*SetModelVersionTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetModelVersionTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name
	st.Value = pb.Value
	st.Version = pb.Version

	return st, nil
}

type SetModelVersionTagResponse struct {
}

func setModelVersionTagResponseToPb(st *SetModelVersionTagResponse) (*setModelVersionTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setModelVersionTagResponsePb{}

	return pb, nil
}

func (st *SetModelVersionTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setModelVersionTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setModelVersionTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetModelVersionTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := setModelVersionTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setModelVersionTagResponsePb struct {
}

func setModelVersionTagResponseFromPb(pb *setModelVersionTagResponsePb) (*SetModelVersionTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetModelVersionTagResponse{}

	return st, nil
}

type SetTag struct {
	// Name of the tag. Keys up to 250 bytes in size are supported.
	// Wire name: 'key'
	Key string
	// ID of the run under which to log the tag. Must be provided.
	// Wire name: 'run_id'
	RunId string
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	// Wire name: 'value'
	Value string

	ForceSendFields []string `tf:"-"`
}

func setTagToPb(st *SetTag) (*setTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setTagPb{}
	pb.Key = st.Key

	pb.RunId = st.RunId

	pb.RunUuid = st.RunUuid

	pb.Value = st.Value

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *SetTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetTag) MarshalJSON() ([]byte, error) {
	pb, err := setTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setTagPb struct {
	// Name of the tag. Keys up to 250 bytes in size are supported.
	Key string `json:"key"`
	// ID of the run under which to log the tag. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	Value string `json:"value"`

	ForceSendFields []string `json:"-" url:"-"`
}

func setTagFromPb(pb *setTagPb) (*SetTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetTag{}
	st.Key = pb.Key
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid
	st.Value = pb.Value

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *setTagPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st setTagPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type SetTagResponse struct {
}

func setTagResponseToPb(st *SetTagResponse) (*setTagResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &setTagResponsePb{}

	return pb, nil
}

func (st *SetTagResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &setTagResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := setTagResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st SetTagResponse) MarshalJSON() ([]byte, error) {
	pb, err := setTagResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type setTagResponsePb struct {
}

func setTagResponseFromPb(pb *setTagResponsePb) (*SetTagResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetTagResponse{}

	return st, nil
}

// Stage of the model version. Valid values are:
//
// * `None`: The initial stage of a model version.
//
// * `Staging`: Staging or pre-production stage.
//
// * `Production`: Production stage.
//
// * `Archived`: Archived stage.
type Stage string
type stagePb string

// Archived stage.
const StageArchived Stage = `Archived`

// The initial stage of a model version.
const StageNone Stage = `None`

// Production stage.
const StageProduction Stage = `Production`

// Staging or pre-production stage.
const StageStaging Stage = `Staging`

// String representation for [fmt.Print]
func (f *Stage) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Stage) Set(v string) error {
	switch v {
	case `Archived`, `None`, `Production`, `Staging`:
		*f = Stage(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "Archived", "None", "Production", "Staging"`, v)
	}
}

// Type always returns Stage to satisfy [pflag.Value] interface
func (f *Stage) Type() string {
	return "Stage"
}

func stageToPb(st *Stage) (*stagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := stagePb(*st)
	return &pb, nil
}

func stageFromPb(pb *stagePb) (*Stage, error) {
	if pb == nil {
		return nil, nil
	}
	st := Stage(*pb)
	return &st, nil
}

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks.
//
// * `FAILED_REGISTRATION`: Request to register a new model version has failed.
//
// * `READY`: Model version is ready for use.
type Status string
type statusPb string

// Request to register a new model version has failed.
const StatusFailedRegistration Status = `FAILED_REGISTRATION`

// Request to register a new model version is pending as server performs
// background tasks.
const StatusPendingRegistration Status = `PENDING_REGISTRATION`

// Model version is ready for use.
const StatusReady Status = `READY`

// String representation for [fmt.Print]
func (f *Status) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *Status) Set(v string) error {
	switch v {
	case `FAILED_REGISTRATION`, `PENDING_REGISTRATION`, `READY`:
		*f = Status(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED_REGISTRATION", "PENDING_REGISTRATION", "READY"`, v)
	}
}

// Type always returns Status to satisfy [pflag.Value] interface
func (f *Status) Type() string {
	return "Status"
}

func statusToPb(st *Status) (*statusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := statusPb(*st)
	return &pb, nil
}

func statusFromPb(pb *statusPb) (*Status, error) {
	if pb == nil {
		return nil, nil
	}
	st := Status(*pb)
	return &st, nil
}

// Test webhook response object.
type TestRegistryWebhook struct {
	// Body of the response from the webhook URL
	// Wire name: 'body'
	Body string
	// Status code returned by the webhook URL
	// Wire name: 'status_code'
	StatusCode int

	ForceSendFields []string `tf:"-"`
}

func testRegistryWebhookToPb(st *TestRegistryWebhook) (*testRegistryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &testRegistryWebhookPb{}
	pb.Body = st.Body

	pb.StatusCode = st.StatusCode

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TestRegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &testRegistryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := testRegistryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TestRegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := testRegistryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type testRegistryWebhookPb struct {
	// Body of the response from the webhook URL
	Body string `json:"body,omitempty"`
	// Status code returned by the webhook URL
	StatusCode int `json:"status_code,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func testRegistryWebhookFromPb(pb *testRegistryWebhookPb) (*TestRegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TestRegistryWebhook{}
	st.Body = pb.Body
	st.StatusCode = pb.StatusCode

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *testRegistryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st testRegistryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	// Wire name: 'event'
	Event RegistryWebhookEvent
	// Webhook ID
	// Wire name: 'id'
	Id string
}

func testRegistryWebhookRequestToPb(st *TestRegistryWebhookRequest) (*testRegistryWebhookRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &testRegistryWebhookRequestPb{}
	pb.Event = st.Event

	pb.Id = st.Id

	return pb, nil
}

func (st *TestRegistryWebhookRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &testRegistryWebhookRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := testRegistryWebhookRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TestRegistryWebhookRequest) MarshalJSON() ([]byte, error) {
	pb, err := testRegistryWebhookRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type testRegistryWebhookRequestPb struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	Event RegistryWebhookEvent `json:"event,omitempty"`
	// Webhook ID
	Id string `json:"id"`
}

func testRegistryWebhookRequestFromPb(pb *testRegistryWebhookRequestPb) (*TestRegistryWebhookRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TestRegistryWebhookRequest{}
	st.Event = pb.Event
	st.Id = pb.Id

	return st, nil
}

type TestRegistryWebhookResponse struct {
	// Test webhook response object.
	// Wire name: 'webhook'
	Webhook *TestRegistryWebhook
}

func testRegistryWebhookResponseToPb(st *TestRegistryWebhookResponse) (*testRegistryWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &testRegistryWebhookResponsePb{}
	webhookPb, err := testRegistryWebhookToPb(st.Webhook)
	if err != nil {
		return nil, err
	}
	if webhookPb != nil {
		pb.Webhook = webhookPb
	}

	return pb, nil
}

func (st *TestRegistryWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &testRegistryWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := testRegistryWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TestRegistryWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := testRegistryWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type testRegistryWebhookResponsePb struct {
	// Test webhook response object.
	Webhook *testRegistryWebhookPb `json:"webhook,omitempty"`
}

func testRegistryWebhookResponseFromPb(pb *testRegistryWebhookResponsePb) (*TestRegistryWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TestRegistryWebhookResponse{}
	webhookField, err := testRegistryWebhookFromPb(pb.Webhook)
	if err != nil {
		return nil, err
	}
	if webhookField != nil {
		st.Webhook = webhookField
	}

	return st, nil
}

type TransitionModelVersionStageDatabricks struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	// Wire name: 'archive_existing_versions'
	ArchiveExistingVersions bool
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Name of the model.
	// Wire name: 'name'
	Name string
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'stage'
	Stage Stage
	// Version of the model.
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func transitionModelVersionStageDatabricksToPb(st *TransitionModelVersionStageDatabricks) (*transitionModelVersionStageDatabricksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transitionModelVersionStageDatabricksPb{}
	pb.ArchiveExistingVersions = st.ArchiveExistingVersions

	pb.Comment = st.Comment

	pb.Name = st.Name

	pb.Stage = st.Stage

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TransitionModelVersionStageDatabricks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &transitionModelVersionStageDatabricksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := transitionModelVersionStageDatabricksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TransitionModelVersionStageDatabricks) MarshalJSON() ([]byte, error) {
	pb, err := transitionModelVersionStageDatabricksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type transitionModelVersionStageDatabricksPb struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions bool `json:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`
	// Target stage of the transition. Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	Stage Stage `json:"stage"`
	// Version of the model.
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func transitionModelVersionStageDatabricksFromPb(pb *transitionModelVersionStageDatabricksPb) (*TransitionModelVersionStageDatabricks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransitionModelVersionStageDatabricks{}
	st.ArchiveExistingVersions = pb.ArchiveExistingVersions
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *transitionModelVersionStageDatabricksPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st transitionModelVersionStageDatabricksPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

// Transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	// Wire name: 'available_actions'
	AvailableActions []ActivityAction
	// User-provided comment associated with the transition request.
	// Wire name: 'comment'
	Comment string
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	// Wire name: 'to_stage'
	ToStage Stage
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string

	ForceSendFields []string `tf:"-"`
}

func transitionRequestToPb(st *TransitionRequest) (*transitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transitionRequestPb{}
	pb.AvailableActions = st.AvailableActions

	pb.Comment = st.Comment

	pb.CreationTimestamp = st.CreationTimestamp

	pb.ToStage = st.ToStage

	pb.UserId = st.UserId

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *TransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &transitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := transitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := transitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type transitionRequestPb struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []ActivityAction `json:"available_actions,omitempty"`
	// User-provided comment associated with the transition request.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are:
	//
	// * `None`: The initial stage of a model version.
	//
	// * `Staging`: Staging or pre-production stage.
	//
	// * `Production`: Production stage.
	//
	// * `Archived`: Archived stage.
	ToStage Stage `json:"to_stage,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func transitionRequestFromPb(pb *transitionRequestPb) (*TransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransitionRequest{}
	st.AvailableActions = pb.AvailableActions
	st.Comment = pb.Comment
	st.CreationTimestamp = pb.CreationTimestamp
	st.ToStage = pb.ToStage
	st.UserId = pb.UserId

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *transitionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st transitionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type TransitionStageResponse struct {

	// Wire name: 'model_version'
	ModelVersion *ModelVersionDatabricks
}

func transitionStageResponseToPb(st *TransitionStageResponse) (*transitionStageResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &transitionStageResponsePb{}
	modelVersionPb, err := modelVersionDatabricksToPb(st.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionPb != nil {
		pb.ModelVersion = modelVersionPb
	}

	return pb, nil
}

func (st *TransitionStageResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &transitionStageResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := transitionStageResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st TransitionStageResponse) MarshalJSON() ([]byte, error) {
	pb, err := transitionStageResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type transitionStageResponsePb struct {
	ModelVersion *modelVersionDatabricksPb `json:"model_version,omitempty"`
}

func transitionStageResponseFromPb(pb *transitionStageResponsePb) (*TransitionStageResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransitionStageResponse{}
	modelVersionField, err := modelVersionDatabricksFromPb(pb.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionField != nil {
		st.ModelVersion = modelVersionField
	}

	return st, nil
}

type UpdateComment struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string
	// Unique identifier of an activity
	// Wire name: 'id'
	Id string
}

func updateCommentToPb(st *UpdateComment) (*updateCommentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCommentPb{}
	pb.Comment = st.Comment

	pb.Id = st.Id

	return pb, nil
}

func (st *UpdateComment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCommentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCommentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateComment) MarshalJSON() ([]byte, error) {
	pb, err := updateCommentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateCommentPb struct {
	// User-provided comment on the action.
	Comment string `json:"comment"`
	// Unique identifier of an activity
	Id string `json:"id"`
}

func updateCommentFromPb(pb *updateCommentPb) (*UpdateComment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateComment{}
	st.Comment = pb.Comment
	st.Id = pb.Id

	return st, nil
}

type UpdateCommentResponse struct {
	// Comment details.
	// Wire name: 'comment'
	Comment *CommentObject
}

func updateCommentResponseToPb(st *UpdateCommentResponse) (*updateCommentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateCommentResponsePb{}
	commentPb, err := commentObjectToPb(st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = commentPb
	}

	return pb, nil
}

func (st *UpdateCommentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateCommentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateCommentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateCommentResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateCommentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateCommentResponsePb struct {
	// Comment details.
	Comment *commentObjectPb `json:"comment,omitempty"`
}

func updateCommentResponseFromPb(pb *updateCommentResponsePb) (*UpdateCommentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCommentResponse{}
	commentField, err := commentObjectFromPb(pb.Comment)
	if err != nil {
		return nil, err
	}
	if commentField != nil {
		st.Comment = commentField
	}

	return st, nil
}

type UpdateExperiment struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	// Wire name: 'new_name'
	NewName string

	ForceSendFields []string `tf:"-"`
}

func updateExperimentToPb(st *UpdateExperiment) (*updateExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExperimentPb{}
	pb.ExperimentId = st.ExperimentId

	pb.NewName = st.NewName

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExperiment) MarshalJSON() ([]byte, error) {
	pb, err := updateExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateExperimentPb struct {
	// ID of the associated experiment.
	ExperimentId string `json:"experiment_id"`
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	NewName string `json:"new_name,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateExperimentFromPb(pb *updateExperimentPb) (*UpdateExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExperiment{}
	st.ExperimentId = pb.ExperimentId
	st.NewName = pb.NewName

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateExperimentPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateExperimentPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateExperimentResponse struct {
}

func updateExperimentResponseToPb(st *UpdateExperimentResponse) (*updateExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateExperimentResponsePb{}

	return pb, nil
}

func (st *UpdateExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateExperimentResponsePb struct {
}

func updateExperimentResponseFromPb(pb *updateExperimentResponsePb) (*UpdateExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExperimentResponse{}

	return st, nil
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	// Wire name: 'description'
	Description string
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string

	ForceSendFields []string `tf:"-"`
}

func updateModelRequestToPb(st *UpdateModelRequest) (*updateModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelRequestPb{}
	pb.Description = st.Description

	pb.Name = st.Name

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateModelRequestPb struct {
	// If provided, updates the description for this `registered_model`.
	Description string `json:"description,omitempty"`
	// Registered model unique name identifier.
	Name string `json:"name"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateModelRequestFromPb(pb *updateModelRequestPb) (*UpdateModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelRequest{}
	st.Description = pb.Description
	st.Name = pb.Name

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateModelRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateModelRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateModelResponse struct {
}

func updateModelResponseToPb(st *UpdateModelResponse) (*updateModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelResponsePb{}

	return pb, nil
}

func (st *UpdateModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateModelResponsePb struct {
}

func updateModelResponseFromPb(pb *updateModelResponsePb) (*UpdateModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelResponse{}

	return st, nil
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this `registered_model`.
	// Wire name: 'description'
	Description string
	// Name of the registered model
	// Wire name: 'name'
	Name string
	// Model version number
	// Wire name: 'version'
	Version string

	ForceSendFields []string `tf:"-"`
}

func updateModelVersionRequestToPb(st *UpdateModelVersionRequest) (*updateModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelVersionRequestPb{}
	pb.Description = st.Description

	pb.Name = st.Name

	pb.Version = st.Version

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := updateModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateModelVersionRequestPb struct {
	// If provided, updates the description for this `registered_model`.
	Description string `json:"description,omitempty"`
	// Name of the registered model
	Name string `json:"name"`
	// Model version number
	Version string `json:"version"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateModelVersionRequestFromPb(pb *updateModelVersionRequestPb) (*UpdateModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelVersionRequest{}
	st.Description = pb.Description
	st.Name = pb.Name
	st.Version = pb.Version

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateModelVersionRequestPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateModelVersionRequestPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateModelVersionResponse struct {
}

func updateModelVersionResponseToPb(st *UpdateModelVersionResponse) (*updateModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateModelVersionResponsePb{}

	return pb, nil
}

func (st *UpdateModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateModelVersionResponsePb struct {
}

func updateModelVersionResponseFromPb(pb *updateModelVersionResponsePb) (*UpdateModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelVersionResponse{}

	return st, nil
}

type UpdateRegistryWebhook struct {
	// User-specified description for the webhook.
	// Wire name: 'description'
	Description string
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	// Wire name: 'events'
	Events []RegistryWebhookEvent

	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpec
	// Webhook ID
	// Wire name: 'id'
	Id string

	// Wire name: 'job_spec'
	JobSpec *JobSpec
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	// Wire name: 'status'
	Status RegistryWebhookStatus

	ForceSendFields []string `tf:"-"`
}

func updateRegistryWebhookToPb(st *UpdateRegistryWebhook) (*updateRegistryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRegistryWebhookPb{}
	pb.Description = st.Description

	pb.Events = st.Events

	httpUrlSpecPb, err := httpUrlSpecToPb(st.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecPb != nil {
		pb.HttpUrlSpec = httpUrlSpecPb
	}

	pb.Id = st.Id

	jobSpecPb, err := jobSpecToPb(st.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecPb != nil {
		pb.JobSpec = jobSpecPb
	}

	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateRegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRegistryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRegistryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := updateRegistryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateRegistryWebhookPb struct {
	// User-specified description for the webhook.
	Description string `json:"description,omitempty"`
	// Events that can trigger a registry webhook: * `MODEL_VERSION_CREATED`: A
	// new model version was created for the associated model.
	//
	// * `MODEL_VERSION_TRANSITIONED_STAGE`: A model version’s stage was
	// changed.
	//
	// * `TRANSITION_REQUEST_CREATED`: A user requested a model version’s
	// stage be transitioned.
	//
	// * `COMMENT_CREATED`: A user wrote a comment on a registered model.
	//
	// * `REGISTERED_MODEL_CREATED`: A new registered model was created. This
	// event type can only be specified for a registry-wide webhook, which can
	// be created by not specifying a model name in the create request.
	//
	// * `MODEL_VERSION_TAG_SET`: A user set a tag on the model version.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_STAGING`: A model version was
	// transitioned to staging.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_PRODUCTION`: A model version was
	// transitioned to production.
	//
	// * `MODEL_VERSION_TRANSITIONED_TO_ARCHIVED`: A model version was archived.
	//
	// * `TRANSITION_REQUEST_TO_STAGING_CREATED`: A user requested a model
	// version be transitioned to staging.
	//
	// * `TRANSITION_REQUEST_TO_PRODUCTION_CREATED`: A user requested a model
	// version be transitioned to production.
	//
	// * `TRANSITION_REQUEST_TO_ARCHIVED_CREATED`: A user requested a model
	// version be archived.
	Events []RegistryWebhookEvent `json:"events,omitempty"`

	HttpUrlSpec *httpUrlSpecPb `json:"http_url_spec,omitempty"`
	// Webhook ID
	Id string `json:"id"`

	JobSpec *jobSpecPb `json:"job_spec,omitempty"`
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	Status RegistryWebhookStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateRegistryWebhookFromPb(pb *updateRegistryWebhookPb) (*UpdateRegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRegistryWebhook{}
	st.Description = pb.Description
	st.Events = pb.Events
	httpUrlSpecField, err := httpUrlSpecFromPb(pb.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecField != nil {
		st.HttpUrlSpec = httpUrlSpecField
	}
	st.Id = pb.Id
	jobSpecField, err := jobSpecFromPb(pb.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecField != nil {
		st.JobSpec = jobSpecField
	}
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateRegistryWebhookPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateRegistryWebhookPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateRun struct {
	// Unix timestamp in milliseconds of when the run ended.
	// Wire name: 'end_time'
	EndTime int64
	// ID of the run to update. Must be provided.
	// Wire name: 'run_id'
	RunId string
	// Updated name of the run.
	// Wire name: 'run_name'
	RunName string
	// [Deprecated, use `run_id` instead] ID of the run to update. This field
	// will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string
	// Updated status of the run.
	// Wire name: 'status'
	Status UpdateRunStatus

	ForceSendFields []string `tf:"-"`
}

func updateRunToPb(st *UpdateRun) (*updateRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRunPb{}
	pb.EndTime = st.EndTime

	pb.RunId = st.RunId

	pb.RunName = st.RunName

	pb.RunUuid = st.RunUuid

	pb.Status = st.Status

	pb.ForceSendFields = st.ForceSendFields
	return pb, nil
}

func (st *UpdateRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRun) MarshalJSON() ([]byte, error) {
	pb, err := updateRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateRunPb struct {
	// Unix timestamp in milliseconds of when the run ended.
	EndTime int64 `json:"end_time,omitempty"`
	// ID of the run to update. Must be provided.
	RunId string `json:"run_id,omitempty"`
	// Updated name of the run.
	RunName string `json:"run_name,omitempty"`
	// [Deprecated, use `run_id` instead] ID of the run to update. This field
	// will be removed in a future MLflow version.
	RunUuid string `json:"run_uuid,omitempty"`
	// Updated status of the run.
	Status UpdateRunStatus `json:"status,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func updateRunFromPb(pb *updateRunPb) (*UpdateRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRun{}
	st.EndTime = pb.EndTime
	st.RunId = pb.RunId
	st.RunName = pb.RunName
	st.RunUuid = pb.RunUuid
	st.Status = pb.Status

	st.ForceSendFields = pb.ForceSendFields
	return st, nil
}

func (st *updateRunPb) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, st)
}

func (st updateRunPb) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(st)
}

type UpdateRunResponse struct {
	// Updated metadata of the run.
	// Wire name: 'run_info'
	RunInfo *RunInfo
}

func updateRunResponseToPb(st *UpdateRunResponse) (*updateRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateRunResponsePb{}
	runInfoPb, err := runInfoToPb(st.RunInfo)
	if err != nil {
		return nil, err
	}
	if runInfoPb != nil {
		pb.RunInfo = runInfoPb
	}

	return pb, nil
}

func (st *UpdateRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateRunResponsePb struct {
	// Updated metadata of the run.
	RunInfo *runInfoPb `json:"run_info,omitempty"`
}

func updateRunResponseFromPb(pb *updateRunResponsePb) (*UpdateRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRunResponse{}
	runInfoField, err := runInfoFromPb(pb.RunInfo)
	if err != nil {
		return nil, err
	}
	if runInfoField != nil {
		st.RunInfo = runInfoField
	}

	return st, nil
}

// Status of a run.
type UpdateRunStatus string
type updateRunStatusPb string

const UpdateRunStatusFailed UpdateRunStatus = `FAILED`

const UpdateRunStatusFinished UpdateRunStatus = `FINISHED`

const UpdateRunStatusKilled UpdateRunStatus = `KILLED`

const UpdateRunStatusRunning UpdateRunStatus = `RUNNING`

const UpdateRunStatusScheduled UpdateRunStatus = `SCHEDULED`

// String representation for [fmt.Print]
func (f *UpdateRunStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *UpdateRunStatus) Set(v string) error {
	switch v {
	case `FAILED`, `FINISHED`, `KILLED`, `RUNNING`, `SCHEDULED`:
		*f = UpdateRunStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "FAILED", "FINISHED", "KILLED", "RUNNING", "SCHEDULED"`, v)
	}
}

// Type always returns UpdateRunStatus to satisfy [pflag.Value] interface
func (f *UpdateRunStatus) Type() string {
	return "UpdateRunStatus"
}

func updateRunStatusToPb(st *UpdateRunStatus) (*updateRunStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := updateRunStatusPb(*st)
	return &pb, nil
}

func updateRunStatusFromPb(pb *updateRunStatusPb) (*UpdateRunStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpdateRunStatus(*pb)
	return &st, nil
}

type UpdateWebhookResponse struct {
}

func updateWebhookResponseToPb(st *UpdateWebhookResponse) (*updateWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &updateWebhookResponsePb{}

	return pb, nil
}

func (st *UpdateWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &updateWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := updateWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func (st UpdateWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := updateWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

type updateWebhookResponsePb struct {
}

func updateWebhookResponseFromPb(pb *updateWebhookResponsePb) (*UpdateWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWebhookResponse{}

	return st, nil
}

// Qualifier for the view type.
type ViewType string
type viewTypePb string

const ViewTypeActiveOnly ViewType = `ACTIVE_ONLY`

const ViewTypeAll ViewType = `ALL`

const ViewTypeDeletedOnly ViewType = `DELETED_ONLY`

// String representation for [fmt.Print]
func (f *ViewType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ViewType) Set(v string) error {
	switch v {
	case `ACTIVE_ONLY`, `ALL`, `DELETED_ONLY`:
		*f = ViewType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ACTIVE_ONLY", "ALL", "DELETED_ONLY"`, v)
	}
}

// Type always returns ViewType to satisfy [pflag.Value] interface
func (f *ViewType) Type() string {
	return "ViewType"
}

func viewTypeToPb(st *ViewType) (*viewTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := viewTypePb(*st)
	return &pb, nil
}

func viewTypeFromPb(pb *viewTypePb) (*ViewType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ViewType(*pb)
	return &st, nil
}

func durationToPb(d *time.Duration) (*string, error) {
	if d == nil {
		return nil, nil
	}
	s := fmt.Sprintf("%fs", d.Seconds())
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
