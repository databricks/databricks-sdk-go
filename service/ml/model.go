// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package ml

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/databricks/databricks-sdk-go/service/ml/mlpb"
)

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type Activity struct {

	// Wire name: 'activity_type'
	ActivityType ActivityType ``
	// User-provided comment associated with the activity, comment, or
	// transition request.
	// Wire name: 'comment'
	Comment string ``
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``
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
	FromStage string ``
	// Unique identifier for the object.
	// Wire name: 'id'
	Id string ``
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version's stage transition, and
	// may not be returned for some activity types.
	// Wire name: 'system_comment'
	SystemComment string ``
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
	ToStage string ``
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st Activity) MarshalJSON() ([]byte, error) {
	pb, err := ActivityToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Activity) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ActivityPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ActivityFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ActivityToPb(st *Activity) (*mlpb.ActivityPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ActivityPb{}
	activityTypePb, err := ActivityTypeToPb(&st.ActivityType)
	if err != nil {
		return nil, err
	}
	if activityTypePb != nil {
		pb.ActivityType = *activityTypePb
	}
	pb.Comment = st.Comment
	pb.CreationTimestamp = st.CreationTimestamp
	pb.FromStage = st.FromStage
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.SystemComment = st.SystemComment
	pb.ToStage = st.ToStage
	pb.UserId = st.UserId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ActivityFromPb(pb *mlpb.ActivityPb) (*Activity, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Activity{}
	activityTypeField, err := ActivityTypeFromPb(&pb.ActivityType)
	if err != nil {
		return nil, err
	}
	if activityTypeField != nil {
		st.ActivityType = *activityTypeField
	}
	st.Comment = pb.Comment
	st.CreationTimestamp = pb.CreationTimestamp
	st.FromStage = pb.FromStage
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.SystemComment = pb.SystemComment
	st.ToStage = pb.ToStage
	st.UserId = pb.UserId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// An action that a user (with sufficient permissions) could take on an activity
// or comment.
//
// For activities, valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request
//
// * `REJECT_TRANSITION_REQUEST`: Reject a transition request
//
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
//
// For comments, valid values are: * `EDIT_COMMENT`: Edit the comment
//
// * `DELETE_COMMENT`: Delete the comment
type ActivityAction string

// Approve a transition request
const ActivityActionApproveTransitionRequest ActivityAction = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const ActivityActionCancelTransitionRequest ActivityAction = `CANCEL_TRANSITION_REQUEST`

// Delete the comment
const ActivityActionDeleteComment ActivityAction = `DELETE_COMMENT`

// Edit the comment
const ActivityActionEditComment ActivityAction = `EDIT_COMMENT`

// Reject a transition request
const ActivityActionRejectTransitionRequest ActivityAction = `REJECT_TRANSITION_REQUEST`

// String representation for [fmt.Print]
func (f *ActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *ActivityAction) Set(v string) error {
	switch v {
	case `APPROVE_TRANSITION_REQUEST`, `CANCEL_TRANSITION_REQUEST`, `DELETE_COMMENT`, `EDIT_COMMENT`, `REJECT_TRANSITION_REQUEST`:
		*f = ActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVE_TRANSITION_REQUEST", "CANCEL_TRANSITION_REQUEST", "DELETE_COMMENT", "EDIT_COMMENT", "REJECT_TRANSITION_REQUEST"`, v)
	}
}

// Values returns all possible values for ActivityAction.
//
// There is no guarantee on the order of the values in the slice.
func (f *ActivityAction) Values() []ActivityAction {
	return []ActivityAction{
		ActivityActionApproveTransitionRequest,
		ActivityActionCancelTransitionRequest,
		ActivityActionDeleteComment,
		ActivityActionEditComment,
		ActivityActionRejectTransitionRequest,
	}
}

// Type always returns ActivityAction to satisfy [pflag.Value] interface
func (f *ActivityAction) Type() string {
	return "ActivityAction"
}

func ActivityActionToPb(st *ActivityAction) (*mlpb.ActivityActionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.ActivityActionPb(*st)
	return &pb, nil
}

func ActivityActionFromPb(pb *mlpb.ActivityActionPb) (*ActivityAction, error) {
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

// Values returns all possible values for ActivityType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ActivityType) Values() []ActivityType {
	return []ActivityType{
		ActivityTypeAppliedTransition,
		ActivityTypeApprovedRequest,
		ActivityTypeCancelledRequest,
		ActivityTypeNewComment,
		ActivityTypeRejectedRequest,
		ActivityTypeRequestedTransition,
		ActivityTypeSystemTransition,
	}
}

// Type always returns ActivityType to satisfy [pflag.Value] interface
func (f *ActivityType) Type() string {
	return "ActivityType"
}

func ActivityTypeToPb(st *ActivityType) (*mlpb.ActivityTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.ActivityTypePb(*st)
	return &pb, nil
}

func ActivityTypeFromPb(pb *mlpb.ActivityTypePb) (*ActivityType, error) {
	if pb == nil {
		return nil, nil
	}
	st := ActivityType(*pb)
	return &st, nil
}

// Details required to identify and approve a model version stage transition
// request.
type ApproveTransitionRequest struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	// Wire name: 'archive_existing_versions'
	ArchiveExistingVersions bool ``
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string ``
	// Name of the model.
	// Wire name: 'name'
	Name string ``
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
	Stage string ``
	// Version of the model.
	// Wire name: 'version'
	Version         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ApproveTransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := ApproveTransitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ApproveTransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ApproveTransitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ApproveTransitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ApproveTransitionRequestToPb(st *ApproveTransitionRequest) (*mlpb.ApproveTransitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ApproveTransitionRequestPb{}
	pb.ArchiveExistingVersions = st.ArchiveExistingVersions
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Stage = st.Stage
	pb.Version = st.Version

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ApproveTransitionRequestFromPb(pb *mlpb.ApproveTransitionRequestPb) (*ApproveTransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApproveTransitionRequest{}
	st.ArchiveExistingVersions = pb.ArchiveExistingVersions
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ApproveTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	// Wire name: 'activity'
	Activity *Activity ``
}

func (st ApproveTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := ApproveTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ApproveTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ApproveTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ApproveTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ApproveTransitionRequestResponseToPb(st *ApproveTransitionRequestResponse) (*mlpb.ApproveTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ApproveTransitionRequestResponsePb{}
	activityPb, err := ActivityToPb(st.Activity)
	if err != nil {
		return nil, err
	}
	if activityPb != nil {
		pb.Activity = activityPb
	}

	return pb, nil
}

func ApproveTransitionRequestResponseFromPb(pb *mlpb.ApproveTransitionRequestResponsePb) (*ApproveTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ApproveTransitionRequestResponse{}
	activityField, err := ActivityFromPb(pb.Activity)
	if err != nil {
		return nil, err
	}
	if activityField != nil {
		st.Activity = activityField
	}

	return st, nil
}

// An action that a user (with sufficient permissions) could take on an activity
// or comment.
//
// For activities, valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request
//
// * `REJECT_TRANSITION_REQUEST`: Reject a transition request
//
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
//
// For comments, valid values are: * `EDIT_COMMENT`: Edit the comment
//
// * `DELETE_COMMENT`: Delete the comment
type CommentActivityAction string

// Approve a transition request
const CommentActivityActionApproveTransitionRequest CommentActivityAction = `APPROVE_TRANSITION_REQUEST`

// Cancel (delete) a transition request
const CommentActivityActionCancelTransitionRequest CommentActivityAction = `CANCEL_TRANSITION_REQUEST`

// Delete the comment
const CommentActivityActionDeleteComment CommentActivityAction = `DELETE_COMMENT`

// Edit the comment
const CommentActivityActionEditComment CommentActivityAction = `EDIT_COMMENT`

// Reject a transition request
const CommentActivityActionRejectTransitionRequest CommentActivityAction = `REJECT_TRANSITION_REQUEST`

// String representation for [fmt.Print]
func (f *CommentActivityAction) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *CommentActivityAction) Set(v string) error {
	switch v {
	case `APPROVE_TRANSITION_REQUEST`, `CANCEL_TRANSITION_REQUEST`, `DELETE_COMMENT`, `EDIT_COMMENT`, `REJECT_TRANSITION_REQUEST`:
		*f = CommentActivityAction(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "APPROVE_TRANSITION_REQUEST", "CANCEL_TRANSITION_REQUEST", "DELETE_COMMENT", "EDIT_COMMENT", "REJECT_TRANSITION_REQUEST"`, v)
	}
}

// Values returns all possible values for CommentActivityAction.
//
// There is no guarantee on the order of the values in the slice.
func (f *CommentActivityAction) Values() []CommentActivityAction {
	return []CommentActivityAction{
		CommentActivityActionApproveTransitionRequest,
		CommentActivityActionCancelTransitionRequest,
		CommentActivityActionDeleteComment,
		CommentActivityActionEditComment,
		CommentActivityActionRejectTransitionRequest,
	}
}

// Type always returns CommentActivityAction to satisfy [pflag.Value] interface
func (f *CommentActivityAction) Type() string {
	return "CommentActivityAction"
}

func CommentActivityActionToPb(st *CommentActivityAction) (*mlpb.CommentActivityActionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.CommentActivityActionPb(*st)
	return &pb, nil
}

func CommentActivityActionFromPb(pb *mlpb.CommentActivityActionPb) (*CommentActivityAction, error) {
	if pb == nil {
		return nil, nil
	}
	st := CommentActivityAction(*pb)
	return &st, nil
}

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type CommentObject struct {
	// Array of actions on the activity allowed for the current viewer.
	// Wire name: 'available_actions'
	AvailableActions []CommentActivityAction ``
	// User-provided comment associated with the activity, comment, or
	// transition request.
	// Wire name: 'comment'
	Comment string ``
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``
	// Unique identifier for the object.
	// Wire name: 'id'
	Id string ``
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CommentObject) MarshalJSON() ([]byte, error) {
	pb, err := CommentObjectToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CommentObject) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CommentObjectPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CommentObjectFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CommentObjectToPb(st *CommentObject) (*mlpb.CommentObjectPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CommentObjectPb{}

	var availableActionsPb []mlpb.CommentActivityActionPb
	for _, item := range st.AvailableActions {
		itemPb, err := CommentActivityActionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			availableActionsPb = append(availableActionsPb, *itemPb)
		}
	}
	pb.AvailableActions = availableActionsPb
	pb.Comment = st.Comment
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.UserId = st.UserId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CommentObjectFromPb(pb *mlpb.CommentObjectPb) (*CommentObject, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CommentObject{}

	var availableActionsField []CommentActivityAction
	for _, itemPb := range pb.AvailableActions {
		item, err := CommentActivityActionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			availableActionsField = append(availableActionsField, *item)
		}
	}
	st.AvailableActions = availableActionsField
	st.Comment = pb.Comment
	st.CreationTimestamp = pb.CreationTimestamp
	st.Id = pb.Id
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.UserId = pb.UserId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Details required to create a comment on a model version.
type CreateComment struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string ``
	// Name of the model.
	// Wire name: 'name'
	Name string ``
	// Version of the model.
	// Wire name: 'version'
	Version string ``
}

func (st CreateComment) MarshalJSON() ([]byte, error) {
	pb, err := CreateCommentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateComment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateCommentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCommentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCommentToPb(st *CreateComment) (*mlpb.CreateCommentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateCommentPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

func CreateCommentFromPb(pb *mlpb.CreateCommentPb) (*CreateComment, error) {
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
	// New comment object
	// Wire name: 'comment'
	Comment *CommentObject ``
}

func (st CreateCommentResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateCommentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateCommentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateCommentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateCommentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateCommentResponseToPb(st *CreateCommentResponse) (*mlpb.CreateCommentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateCommentResponsePb{}
	commentPb, err := CommentObjectToPb(st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = commentPb
	}

	return pb, nil
}

func CreateCommentResponseFromPb(pb *mlpb.CreateCommentResponsePb) (*CreateCommentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateCommentResponse{}
	commentField, err := CommentObjectFromPb(pb.Comment)
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
	ArtifactLocation string ``
	// Experiment name.
	// Wire name: 'name'
	Name string ``
	// A collection of tags to set on the experiment. Maximum tag size and
	// number of tags per request depends on the storage backend. All storage
	// backends are guaranteed to support tag keys up to 250 bytes in size and
	// tag values up to 5000 bytes in size. All storage backends are also
	// guaranteed to support up to 20 tags per request.
	// Wire name: 'tags'
	Tags            []ExperimentTag ``
	ForceSendFields []string        `tf:"-"`
}

func (st CreateExperiment) MarshalJSON() ([]byte, error) {
	pb, err := CreateExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateExperimentToPb(st *CreateExperiment) (*mlpb.CreateExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateExperimentPb{}
	pb.ArtifactLocation = st.ArtifactLocation
	pb.Name = st.Name

	var tagsPb []mlpb.ExperimentTagPb
	for _, item := range st.Tags {
		itemPb, err := ExperimentTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateExperimentFromPb(pb *mlpb.CreateExperimentPb) (*CreateExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExperiment{}
	st.ArtifactLocation = pb.ArtifactLocation
	st.Name = pb.Name

	var tagsField []ExperimentTag
	for _, itemPb := range pb.Tags {
		item, err := ExperimentTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateExperimentResponse struct {
	// Unique identifier for the experiment.
	// Wire name: 'experiment_id'
	ExperimentId    string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateExperimentResponseToPb(st *CreateExperimentResponse) (*mlpb.CreateExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateExperimentResponsePb{}
	pb.ExperimentId = st.ExperimentId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateExperimentResponseFromPb(pb *mlpb.CreateExperimentResponsePb) (*CreateExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateExperimentResponse{}
	st.ExperimentId = pb.ExperimentId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateFeatureTagRequest struct {

	// Wire name: 'feature_name'
	FeatureName string `tf:"-"`

	// Wire name: 'feature_tag'
	FeatureTag FeatureTag ``

	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func (st CreateFeatureTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateFeatureTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateFeatureTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateFeatureTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateFeatureTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateFeatureTagRequestToPb(st *CreateFeatureTagRequest) (*mlpb.CreateFeatureTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateFeatureTagRequestPb{}
	pb.FeatureName = st.FeatureName
	featureTagPb, err := FeatureTagToPb(&st.FeatureTag)
	if err != nil {
		return nil, err
	}
	if featureTagPb != nil {
		pb.FeatureTag = *featureTagPb
	}
	pb.TableName = st.TableName

	return pb, nil
}

func CreateFeatureTagRequestFromPb(pb *mlpb.CreateFeatureTagRequestPb) (*CreateFeatureTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateFeatureTagRequest{}
	st.FeatureName = pb.FeatureName
	featureTagField, err := FeatureTagFromPb(&pb.FeatureTag)
	if err != nil {
		return nil, err
	}
	if featureTagField != nil {
		st.FeatureTag = *featureTagField
	}
	st.TableName = pb.TableName

	return st, nil
}

type CreateForecastingExperimentRequest struct {
	// The column in the training table used to customize weights for each time
	// series.
	// Wire name: 'custom_weights_column'
	CustomWeightsColumn string ``
	// The path in the workspace to store the created experiment.
	// Wire name: 'experiment_path'
	ExperimentPath string ``
	// The time interval between consecutive rows in the time series data.
	// Possible values include: '1 second', '1 minute', '5 minutes', '10
	// minutes', '15 minutes', '30 minutes', 'Hourly', 'Daily', 'Weekly',
	// 'Monthly', 'Quarterly', 'Yearly'.
	// Wire name: 'forecast_granularity'
	ForecastGranularity string ``
	// The number of time steps into the future to make predictions, calculated
	// as a multiple of forecast_granularity. This value represents how far
	// ahead the model should forecast.
	// Wire name: 'forecast_horizon'
	ForecastHorizon int64 ``
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store future feature data
	// for predictions.
	// Wire name: 'future_feature_data_path'
	FutureFeatureDataPath string ``
	// The region code(s) to automatically add holiday features. Currently
	// supports only one region.
	// Wire name: 'holiday_regions'
	HolidayRegions []string ``
	// Specifies the list of feature columns to include in model training. These
	// columns must exist in the training data and be of type string, numerical,
	// or boolean. If not specified, no additional features will be included.
	// Note: Certain columns are automatically handled: - Automatically
	// excluded: split_column, target_column, custom_weights_column. -
	// Automatically included: time_column.
	// Wire name: 'include_features'
	IncludeFeatures []string ``
	// The maximum duration for the experiment in minutes. The experiment stops
	// automatically if it exceeds this limit.
	// Wire name: 'max_runtime'
	MaxRuntime int64 ``
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used to store predictions.
	// Wire name: 'prediction_data_path'
	PredictionDataPath string ``
	// The evaluation metric used to optimize the forecasting model.
	// Wire name: 'primary_metric'
	PrimaryMetric string ``
	// The fully qualified path of a Unity Catalog model, formatted as
	// catalog_name.schema_name.model_name, used to store the best model.
	// Wire name: 'register_to'
	RegisterTo string ``
	// // The column in the training table used for custom data splits. Values
	// must be 'train', 'validate', or 'test'.
	// Wire name: 'split_column'
	SplitColumn string ``
	// The column in the input training table used as the prediction target for
	// model training. The values in this column are used as the ground truth
	// for model training.
	// Wire name: 'target_column'
	TargetColumn string ``
	// The column in the input training table that represents each row's
	// timestamp.
	// Wire name: 'time_column'
	TimeColumn string ``
	// The column in the training table used to group the dataset for predicting
	// individual time series.
	// Wire name: 'timeseries_identifier_columns'
	TimeseriesIdentifierColumns []string ``
	// The fully qualified path of a Unity Catalog table, formatted as
	// catalog_name.schema_name.table_name, used as training data for the
	// forecasting model.
	// Wire name: 'train_data_path'
	TrainDataPath string ``
	// List of frameworks to include for model tuning. Possible values are
	// 'Prophet', 'ARIMA', 'DeepAR'. An empty list includes all supported
	// frameworks.
	// Wire name: 'training_frameworks'
	TrainingFrameworks []string ``
	ForceSendFields    []string `tf:"-"`
}

func (st CreateForecastingExperimentRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateForecastingExperimentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateForecastingExperimentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateForecastingExperimentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateForecastingExperimentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateForecastingExperimentRequestToPb(st *CreateForecastingExperimentRequest) (*mlpb.CreateForecastingExperimentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateForecastingExperimentRequestPb{}
	pb.CustomWeightsColumn = st.CustomWeightsColumn
	pb.ExperimentPath = st.ExperimentPath
	pb.ForecastGranularity = st.ForecastGranularity
	pb.ForecastHorizon = st.ForecastHorizon
	pb.FutureFeatureDataPath = st.FutureFeatureDataPath
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateForecastingExperimentRequestFromPb(pb *mlpb.CreateForecastingExperimentRequestPb) (*CreateForecastingExperimentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateForecastingExperimentRequest{}
	st.CustomWeightsColumn = pb.CustomWeightsColumn
	st.ExperimentPath = pb.ExperimentPath
	st.ForecastGranularity = pb.ForecastGranularity
	st.ForecastHorizon = pb.ForecastHorizon
	st.FutureFeatureDataPath = pb.FutureFeatureDataPath
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateForecastingExperimentResponse struct {
	// The unique ID of the created forecasting experiment
	// Wire name: 'experiment_id'
	ExperimentId    string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateForecastingExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateForecastingExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateForecastingExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateForecastingExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateForecastingExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateForecastingExperimentResponseToPb(st *CreateForecastingExperimentResponse) (*mlpb.CreateForecastingExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateForecastingExperimentResponsePb{}
	pb.ExperimentId = st.ExperimentId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateForecastingExperimentResponseFromPb(pb *mlpb.CreateForecastingExperimentResponsePb) (*CreateForecastingExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateForecastingExperimentResponse{}
	st.ExperimentId = pb.ExperimentId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateLoggedModelRequest struct {
	// The ID of the experiment that owns the model.
	// Wire name: 'experiment_id'
	ExperimentId string ``
	// The type of the model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	// Wire name: 'model_type'
	ModelType string ``
	// The name of the model (optional). If not specified one will be generated.
	// Wire name: 'name'
	Name string ``
	// Parameters attached to the model.
	// Wire name: 'params'
	Params []LoggedModelParameter ``
	// The ID of the run that created the model.
	// Wire name: 'source_run_id'
	SourceRunId string ``
	// Tags attached to the model.
	// Wire name: 'tags'
	Tags            []LoggedModelTag ``
	ForceSendFields []string         `tf:"-"`
}

func (st CreateLoggedModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateLoggedModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateLoggedModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateLoggedModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateLoggedModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateLoggedModelRequestToPb(st *CreateLoggedModelRequest) (*mlpb.CreateLoggedModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateLoggedModelRequestPb{}
	pb.ExperimentId = st.ExperimentId
	pb.ModelType = st.ModelType
	pb.Name = st.Name

	var paramsPb []mlpb.LoggedModelParameterPb
	for _, item := range st.Params {
		itemPb, err := LoggedModelParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			paramsPb = append(paramsPb, *itemPb)
		}
	}
	pb.Params = paramsPb
	pb.SourceRunId = st.SourceRunId

	var tagsPb []mlpb.LoggedModelTagPb
	for _, item := range st.Tags {
		itemPb, err := LoggedModelTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateLoggedModelRequestFromPb(pb *mlpb.CreateLoggedModelRequestPb) (*CreateLoggedModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateLoggedModelRequest{}
	st.ExperimentId = pb.ExperimentId
	st.ModelType = pb.ModelType
	st.Name = pb.Name

	var paramsField []LoggedModelParameter
	for _, itemPb := range pb.Params {
		item, err := LoggedModelParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			paramsField = append(paramsField, *item)
		}
	}
	st.Params = paramsField
	st.SourceRunId = pb.SourceRunId

	var tagsField []LoggedModelTag
	for _, itemPb := range pb.Tags {
		item, err := LoggedModelTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateLoggedModelResponse struct {
	// The newly created logged model.
	// Wire name: 'model'
	Model *LoggedModel ``
}

func (st CreateLoggedModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateLoggedModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateLoggedModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateLoggedModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateLoggedModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateLoggedModelResponseToPb(st *CreateLoggedModelResponse) (*mlpb.CreateLoggedModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateLoggedModelResponsePb{}
	modelPb, err := LoggedModelToPb(st.Model)
	if err != nil {
		return nil, err
	}
	if modelPb != nil {
		pb.Model = modelPb
	}

	return pb, nil
}

func CreateLoggedModelResponseFromPb(pb *mlpb.CreateLoggedModelResponsePb) (*CreateLoggedModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateLoggedModelResponse{}
	modelField, err := LoggedModelFromPb(pb.Model)
	if err != nil {
		return nil, err
	}
	if modelField != nil {
		st.Model = modelField
	}

	return st, nil
}

type CreateModelRequest struct {
	// Optional description for registered model.
	// Wire name: 'description'
	Description string ``
	// Register models under this name
	// Wire name: 'name'
	Name string ``
	// Additional metadata for registered model.
	// Wire name: 'tags'
	Tags            []ModelTag ``
	ForceSendFields []string   `tf:"-"`
}

func (st CreateModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateModelRequestToPb(st *CreateModelRequest) (*mlpb.CreateModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateModelRequestPb{}
	pb.Description = st.Description
	pb.Name = st.Name

	var tagsPb []mlpb.ModelTagPb
	for _, item := range st.Tags {
		itemPb, err := ModelTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateModelRequestFromPb(pb *mlpb.CreateModelRequestPb) (*CreateModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelRequest{}
	st.Description = pb.Description
	st.Name = pb.Name

	var tagsField []ModelTag
	for _, itemPb := range pb.Tags {
		item, err := ModelTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateModelResponse struct {

	// Wire name: 'registered_model'
	RegisteredModel *Model ``
}

func (st CreateModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateModelResponseToPb(st *CreateModelResponse) (*mlpb.CreateModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateModelResponsePb{}
	registeredModelPb, err := ModelToPb(st.RegisteredModel)
	if err != nil {
		return nil, err
	}
	if registeredModelPb != nil {
		pb.RegisteredModel = registeredModelPb
	}

	return pb, nil
}

func CreateModelResponseFromPb(pb *mlpb.CreateModelResponsePb) (*CreateModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelResponse{}
	registeredModelField, err := ModelFromPb(pb.RegisteredModel)
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
	Description string ``
	// Register model under this name
	// Wire name: 'name'
	Name string ``
	// MLflow run ID for correlation, if `source` was generated by an experiment
	// run in MLflow tracking server
	// Wire name: 'run_id'
	RunId string ``
	// MLflow run link - this is the exact link of the run that generated this
	// model version, potentially hosted at another instance of MLflow.
	// Wire name: 'run_link'
	RunLink string ``
	// URI indicating the location of the model artifacts.
	// Wire name: 'source'
	Source string ``
	// Additional metadata for model version.
	// Wire name: 'tags'
	Tags            []ModelVersionTag ``
	ForceSendFields []string          `tf:"-"`
}

func (st CreateModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateModelVersionRequestToPb(st *CreateModelVersionRequest) (*mlpb.CreateModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateModelVersionRequestPb{}
	pb.Description = st.Description
	pb.Name = st.Name
	pb.RunId = st.RunId
	pb.RunLink = st.RunLink
	pb.Source = st.Source

	var tagsPb []mlpb.ModelVersionTagPb
	for _, item := range st.Tags {
		itemPb, err := ModelVersionTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateModelVersionRequestFromPb(pb *mlpb.CreateModelVersionRequestPb) (*CreateModelVersionRequest, error) {
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
		item, err := ModelVersionTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	// Wire name: 'model_version'
	ModelVersion *ModelVersion ``
}

func (st CreateModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateModelVersionResponseToPb(st *CreateModelVersionResponse) (*mlpb.CreateModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateModelVersionResponsePb{}
	modelVersionPb, err := ModelVersionToPb(st.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionPb != nil {
		pb.ModelVersion = modelVersionPb
	}

	return pb, nil
}

func CreateModelVersionResponseFromPb(pb *mlpb.CreateModelVersionResponsePb) (*CreateModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateModelVersionResponse{}
	modelVersionField, err := ModelVersionFromPb(pb.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionField != nil {
		st.ModelVersion = modelVersionField
	}

	return st, nil
}

type CreateOnlineStoreRequest struct {
	// Online store to create.
	// Wire name: 'online_store'
	OnlineStore OnlineStore ``
}

func (st CreateOnlineStoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateOnlineStoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateOnlineStoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateOnlineStoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateOnlineStoreRequestToPb(st *CreateOnlineStoreRequest) (*mlpb.CreateOnlineStoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateOnlineStoreRequestPb{}
	onlineStorePb, err := OnlineStoreToPb(&st.OnlineStore)
	if err != nil {
		return nil, err
	}
	if onlineStorePb != nil {
		pb.OnlineStore = *onlineStorePb
	}

	return pb, nil
}

func CreateOnlineStoreRequestFromPb(pb *mlpb.CreateOnlineStoreRequestPb) (*CreateOnlineStoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateOnlineStoreRequest{}
	onlineStoreField, err := OnlineStoreFromPb(&pb.OnlineStore)
	if err != nil {
		return nil, err
	}
	if onlineStoreField != nil {
		st.OnlineStore = *onlineStoreField
	}

	return st, nil
}

// Details required to create a registry webhook.
type CreateRegistryWebhook struct {
	// User-specified description for the webhook.
	// Wire name: 'description'
	Description string ``
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
	Events []RegistryWebhookEvent ``
	// External HTTPS URL called on event trigger (by using a POST request).
	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpec ``
	// ID of the job that the webhook runs.
	// Wire name: 'job_spec'
	JobSpec *JobSpec ``
	// If model name is not specified, a registry-wide webhook is created that
	// listens for the specified events across all versions of all registered
	// models.
	// Wire name: 'model_name'
	ModelName string ``
	// Enable or disable triggering the webhook, or put the webhook into test
	// mode. The default is `ACTIVE`: * `ACTIVE`: Webhook is triggered when an
	// associated event happens.
	//
	// * `DISABLED`: Webhook is not triggered.
	//
	// * `TEST_MODE`: Webhook can be triggered through the test endpoint, but is
	// not triggered on a real event.
	// Wire name: 'status'
	Status          RegistryWebhookStatus ``
	ForceSendFields []string              `tf:"-"`
}

func (st CreateRegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := CreateRegistryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateRegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateRegistryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateRegistryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateRegistryWebhookToPb(st *CreateRegistryWebhook) (*mlpb.CreateRegistryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateRegistryWebhookPb{}
	pb.Description = st.Description

	var eventsPb []mlpb.RegistryWebhookEventPb
	for _, item := range st.Events {
		itemPb, err := RegistryWebhookEventToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			eventsPb = append(eventsPb, *itemPb)
		}
	}
	pb.Events = eventsPb
	httpUrlSpecPb, err := HttpUrlSpecToPb(st.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecPb != nil {
		pb.HttpUrlSpec = httpUrlSpecPb
	}
	jobSpecPb, err := JobSpecToPb(st.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecPb != nil {
		pb.JobSpec = jobSpecPb
	}
	pb.ModelName = st.ModelName
	statusPb, err := RegistryWebhookStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateRegistryWebhookFromPb(pb *mlpb.CreateRegistryWebhookPb) (*CreateRegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRegistryWebhook{}
	st.Description = pb.Description

	var eventsField []RegistryWebhookEvent
	for _, itemPb := range pb.Events {
		item, err := RegistryWebhookEventFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			eventsField = append(eventsField, *item)
		}
	}
	st.Events = eventsField
	httpUrlSpecField, err := HttpUrlSpecFromPb(pb.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecField != nil {
		st.HttpUrlSpec = httpUrlSpecField
	}
	jobSpecField, err := JobSpecFromPb(pb.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecField != nil {
		st.JobSpec = jobSpecField
	}
	st.ModelName = pb.ModelName
	statusField, err := RegistryWebhookStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateRun struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string ``
	// The name of the run.
	// Wire name: 'run_name'
	RunName string ``
	// Unix timestamp in milliseconds of when the run started.
	// Wire name: 'start_time'
	StartTime int64 ``
	// Additional metadata for run.
	// Wire name: 'tags'
	Tags []RunTag ``
	// ID of the user executing the run. This field is deprecated as of MLflow
	// 1.0, and will be removed in a future MLflow release. Use 'mlflow.user'
	// tag instead.
	// Wire name: 'user_id'
	UserId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateRun) MarshalJSON() ([]byte, error) {
	pb, err := CreateRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateRunToPb(st *CreateRun) (*mlpb.CreateRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateRunPb{}
	pb.ExperimentId = st.ExperimentId
	pb.RunName = st.RunName
	pb.StartTime = st.StartTime

	var tagsPb []mlpb.RunTagPb
	for _, item := range st.Tags {
		itemPb, err := RunTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb
	pb.UserId = st.UserId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateRunFromPb(pb *mlpb.CreateRunPb) (*CreateRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRun{}
	st.ExperimentId = pb.ExperimentId
	st.RunName = pb.RunName
	st.StartTime = pb.StartTime

	var tagsField []RunTag
	for _, itemPb := range pb.Tags {
		item, err := RunTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.UserId = pb.UserId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateRunResponse struct {
	// The newly created run.
	// Wire name: 'run'
	Run *Run ``
}

func (st CreateRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateRunResponseToPb(st *CreateRunResponse) (*mlpb.CreateRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateRunResponsePb{}
	runPb, err := RunToPb(st.Run)
	if err != nil {
		return nil, err
	}
	if runPb != nil {
		pb.Run = runPb
	}

	return pb, nil
}

func CreateRunResponseFromPb(pb *mlpb.CreateRunResponsePb) (*CreateRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateRunResponse{}
	runField, err := RunFromPb(pb.Run)
	if err != nil {
		return nil, err
	}
	if runField != nil {
		st.Run = runField
	}

	return st, nil
}

// Details required to create a model version stage transition request.
type CreateTransitionRequest struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string ``
	// Name of the model.
	// Wire name: 'name'
	Name string ``
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
	Stage string ``
	// Version of the model.
	// Wire name: 'version'
	Version         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st CreateTransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := CreateTransitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateTransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateTransitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateTransitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateTransitionRequestToPb(st *CreateTransitionRequest) (*mlpb.CreateTransitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateTransitionRequestPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Stage = st.Stage
	pb.Version = st.Version

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func CreateTransitionRequestFromPb(pb *mlpb.CreateTransitionRequestPb) (*CreateTransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTransitionRequest{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type CreateTransitionRequestResponse struct {
	// New activity generated for stage transition request.
	// Wire name: 'request'
	Request *TransitionRequest ``
}

func (st CreateTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateTransitionRequestResponseToPb(st *CreateTransitionRequestResponse) (*mlpb.CreateTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateTransitionRequestResponsePb{}
	requestPb, err := TransitionRequestToPb(st.Request)
	if err != nil {
		return nil, err
	}
	if requestPb != nil {
		pb.Request = requestPb
	}

	return pb, nil
}

func CreateTransitionRequestResponseFromPb(pb *mlpb.CreateTransitionRequestResponsePb) (*CreateTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateTransitionRequestResponse{}
	requestField, err := TransitionRequestFromPb(pb.Request)
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
	Webhook *RegistryWebhook ``
}

func (st CreateWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := CreateWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *CreateWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.CreateWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := CreateWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func CreateWebhookResponseToPb(st *CreateWebhookResponse) (*mlpb.CreateWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.CreateWebhookResponsePb{}
	webhookPb, err := RegistryWebhookToPb(st.Webhook)
	if err != nil {
		return nil, err
	}
	if webhookPb != nil {
		pb.Webhook = webhookPb
	}

	return pb, nil
}

func CreateWebhookResponseFromPb(pb *mlpb.CreateWebhookResponsePb) (*CreateWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &CreateWebhookResponse{}
	webhookField, err := RegistryWebhookFromPb(pb.Webhook)
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
	Digest string ``
	// The name of the dataset. E.g. “my.uc.table@2” “nyc-taxi-dataset”,
	// “fantastic-elk-3”
	// Wire name: 'name'
	Name string ``
	// The profile of the dataset. Summary statistics for the dataset, such as
	// the number of rows in a table, the mean / std / mode of each column in a
	// table, or the number of elements in an array.
	// Wire name: 'profile'
	Profile string ``
	// The schema of the dataset. E.g., MLflow ColSpec JSON for a dataframe,
	// MLflow TensorSpec JSON for an ndarray, or another schema format.
	// Wire name: 'schema'
	Schema string ``
	// Source information for the dataset. Note that the source may not exactly
	// reproduce the dataset if it was transformed / modified before use with
	// MLflow.
	// Wire name: 'source'
	Source string ``
	// The type of the dataset source, e.g. ‘databricks-uc-table’,
	// ‘DBFS’, ‘S3’, ...
	// Wire name: 'source_type'
	SourceType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st Dataset) MarshalJSON() ([]byte, error) {
	pb, err := DatasetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Dataset) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DatasetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DatasetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DatasetToPb(st *Dataset) (*mlpb.DatasetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DatasetPb{}
	pb.Digest = st.Digest
	pb.Name = st.Name
	pb.Profile = st.Profile
	pb.Schema = st.Schema
	pb.Source = st.Source
	pb.SourceType = st.SourceType

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DatasetFromPb(pb *mlpb.DatasetPb) (*Dataset, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// DatasetInput. Represents a dataset and input tags.
type DatasetInput struct {
	// The dataset being used as a Run input.
	// Wire name: 'dataset'
	Dataset Dataset ``
	// A list of tags for the dataset input, e.g. a “context” tag with value
	// “training”
	// Wire name: 'tags'
	Tags []InputTag ``
}

func (st DatasetInput) MarshalJSON() ([]byte, error) {
	pb, err := DatasetInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DatasetInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DatasetInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DatasetInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DatasetInputToPb(st *DatasetInput) (*mlpb.DatasetInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DatasetInputPb{}
	datasetPb, err := DatasetToPb(&st.Dataset)
	if err != nil {
		return nil, err
	}
	if datasetPb != nil {
		pb.Dataset = *datasetPb
	}

	var tagsPb []mlpb.InputTagPb
	for _, item := range st.Tags {
		itemPb, err := InputTagToPb(&item)
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

func DatasetInputFromPb(pb *mlpb.DatasetInputPb) (*DatasetInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DatasetInput{}
	datasetField, err := DatasetFromPb(&pb.Dataset)
	if err != nil {
		return nil, err
	}
	if datasetField != nil {
		st.Dataset = *datasetField
	}

	var tagsField []InputTag
	for _, itemPb := range pb.Tags {
		item, err := InputTagFromPb(&itemPb)
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

type DeleteCommentRequest struct {
	// Unique identifier of an activity
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st DeleteCommentRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteCommentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteCommentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteCommentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteCommentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteCommentRequestToPb(st *DeleteCommentRequest) (*mlpb.DeleteCommentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteCommentRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteCommentRequestFromPb(pb *mlpb.DeleteCommentRequestPb) (*DeleteCommentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteCommentRequest{}
	st.Id = pb.Id

	return st, nil
}

type DeleteExperiment struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string ``
}

func (st DeleteExperiment) MarshalJSON() ([]byte, error) {
	pb, err := DeleteExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteExperimentToPb(st *DeleteExperiment) (*mlpb.DeleteExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteExperimentPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func DeleteExperimentFromPb(pb *mlpb.DeleteExperimentPb) (*DeleteExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteExperiment{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

type DeleteFeatureTagRequest struct {
	// The name of the feature within the feature table.
	// Wire name: 'feature_name'
	FeatureName string `tf:"-"`
	// The key of the tag to delete.
	// Wire name: 'key'
	Key string `tf:"-"`
	// The name of the feature table.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func (st DeleteFeatureTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteFeatureTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteFeatureTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteFeatureTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteFeatureTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteFeatureTagRequestToPb(st *DeleteFeatureTagRequest) (*mlpb.DeleteFeatureTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteFeatureTagRequestPb{}
	pb.FeatureName = st.FeatureName
	pb.Key = st.Key
	pb.TableName = st.TableName

	return pb, nil
}

func DeleteFeatureTagRequestFromPb(pb *mlpb.DeleteFeatureTagRequestPb) (*DeleteFeatureTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteFeatureTagRequest{}
	st.FeatureName = pb.FeatureName
	st.Key = pb.Key
	st.TableName = pb.TableName

	return st, nil
}

type DeleteLoggedModelRequest struct {
	// The ID of the logged model to delete.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
}

func (st DeleteLoggedModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteLoggedModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteLoggedModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteLoggedModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteLoggedModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteLoggedModelRequestToPb(st *DeleteLoggedModelRequest) (*mlpb.DeleteLoggedModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteLoggedModelRequestPb{}
	pb.ModelId = st.ModelId

	return pb, nil
}

func DeleteLoggedModelRequestFromPb(pb *mlpb.DeleteLoggedModelRequestPb) (*DeleteLoggedModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteLoggedModelRequest{}
	st.ModelId = pb.ModelId

	return st, nil
}

type DeleteLoggedModelTagRequest struct {
	// The ID of the logged model to delete the tag from.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
	// The tag key.
	// Wire name: 'tag_key'
	TagKey string `tf:"-"`
}

func (st DeleteLoggedModelTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteLoggedModelTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteLoggedModelTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteLoggedModelTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteLoggedModelTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteLoggedModelTagRequestToPb(st *DeleteLoggedModelTagRequest) (*mlpb.DeleteLoggedModelTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteLoggedModelTagRequestPb{}
	pb.ModelId = st.ModelId
	pb.TagKey = st.TagKey

	return pb, nil
}

func DeleteLoggedModelTagRequestFromPb(pb *mlpb.DeleteLoggedModelTagRequestPb) (*DeleteLoggedModelTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteLoggedModelTagRequest{}
	st.ModelId = pb.ModelId
	st.TagKey = pb.TagKey

	return st, nil
}

type DeleteModelRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func (st DeleteModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteModelRequestToPb(st *DeleteModelRequest) (*mlpb.DeleteModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteModelRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteModelRequestFromPb(pb *mlpb.DeleteModelRequestPb) (*DeleteModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteModelTagRequest struct {
	// Name of the tag. The name must be an exact match; wild-card deletion is
	// not supported. Maximum size is 250 bytes.
	// Wire name: 'key'
	Key string `tf:"-"`
	// Name of the registered model that the tag was logged under.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func (st DeleteModelTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteModelTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteModelTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteModelTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteModelTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteModelTagRequestToPb(st *DeleteModelTagRequest) (*mlpb.DeleteModelTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteModelTagRequestPb{}
	pb.Key = st.Key
	pb.Name = st.Name

	return pb, nil
}

func DeleteModelTagRequestFromPb(pb *mlpb.DeleteModelTagRequestPb) (*DeleteModelTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name

	return st, nil
}

type DeleteModelVersionRequest struct {
	// Name of the registered model
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number
	// Wire name: 'version'
	Version string `tf:"-"`
}

func (st DeleteModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteModelVersionRequestToPb(st *DeleteModelVersionRequest) (*mlpb.DeleteModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteModelVersionRequestPb{}
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

func DeleteModelVersionRequestFromPb(pb *mlpb.DeleteModelVersionRequestPb) (*DeleteModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionRequest{}
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

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

func (st DeleteModelVersionTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteModelVersionTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteModelVersionTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteModelVersionTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteModelVersionTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteModelVersionTagRequestToPb(st *DeleteModelVersionTagRequest) (*mlpb.DeleteModelVersionTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteModelVersionTagRequestPb{}
	pb.Key = st.Key
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

func DeleteModelVersionTagRequestFromPb(pb *mlpb.DeleteModelVersionTagRequestPb) (*DeleteModelVersionTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteModelVersionTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name
	st.Version = pb.Version

	return st, nil
}

type DeleteOnlineStoreRequest struct {
	// Name of the online store to delete.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func (st DeleteOnlineStoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteOnlineStoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteOnlineStoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteOnlineStoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteOnlineStoreRequestToPb(st *DeleteOnlineStoreRequest) (*mlpb.DeleteOnlineStoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteOnlineStoreRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func DeleteOnlineStoreRequestFromPb(pb *mlpb.DeleteOnlineStoreRequestPb) (*DeleteOnlineStoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteOnlineStoreRequest{}
	st.Name = pb.Name

	return st, nil
}

type DeleteRun struct {
	// ID of the run to delete.
	// Wire name: 'run_id'
	RunId string ``
}

func (st DeleteRun) MarshalJSON() ([]byte, error) {
	pb, err := DeleteRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteRunToPb(st *DeleteRun) (*mlpb.DeleteRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteRunPb{}
	pb.RunId = st.RunId

	return pb, nil
}

func DeleteRunFromPb(pb *mlpb.DeleteRunPb) (*DeleteRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRun{}
	st.RunId = pb.RunId

	return st, nil
}

type DeleteRuns struct {
	// The ID of the experiment containing the runs to delete.
	// Wire name: 'experiment_id'
	ExperimentId string ``
	// An optional positive integer indicating the maximum number of runs to
	// delete. The maximum allowed value for max_runs is 10000.
	// Wire name: 'max_runs'
	MaxRuns int ``
	// The maximum creation timestamp in milliseconds since the UNIX epoch for
	// deleting runs. Only runs created prior to or at this timestamp are
	// deleted.
	// Wire name: 'max_timestamp_millis'
	MaxTimestampMillis int64    ``
	ForceSendFields    []string `tf:"-"`
}

func (st DeleteRuns) MarshalJSON() ([]byte, error) {
	pb, err := DeleteRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteRunsToPb(st *DeleteRuns) (*mlpb.DeleteRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteRunsPb{}
	pb.ExperimentId = st.ExperimentId
	pb.MaxRuns = st.MaxRuns
	pb.MaxTimestampMillis = st.MaxTimestampMillis

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteRunsFromPb(pb *mlpb.DeleteRunsPb) (*DeleteRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRuns{}
	st.ExperimentId = pb.ExperimentId
	st.MaxRuns = pb.MaxRuns
	st.MaxTimestampMillis = pb.MaxTimestampMillis

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteRunsResponse struct {
	// The number of runs deleted.
	// Wire name: 'runs_deleted'
	RunsDeleted     int      ``
	ForceSendFields []string `tf:"-"`
}

func (st DeleteRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteRunsResponseToPb(st *DeleteRunsResponse) (*mlpb.DeleteRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteRunsResponsePb{}
	pb.RunsDeleted = st.RunsDeleted

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteRunsResponseFromPb(pb *mlpb.DeleteRunsResponsePb) (*DeleteRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteRunsResponse{}
	st.RunsDeleted = pb.RunsDeleted

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteTag struct {
	// Name of the tag. Maximum size is 255 bytes. Must be provided.
	// Wire name: 'key'
	Key string ``
	// ID of the run that the tag was logged under. Must be provided.
	// Wire name: 'run_id'
	RunId string ``
}

func (st DeleteTag) MarshalJSON() ([]byte, error) {
	pb, err := DeleteTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteTagToPb(st *DeleteTag) (*mlpb.DeleteTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteTagPb{}
	pb.Key = st.Key
	pb.RunId = st.RunId

	return pb, nil
}

func DeleteTagFromPb(pb *mlpb.DeleteTagPb) (*DeleteTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTag{}
	st.Key = pb.Key
	st.RunId = pb.RunId

	return st, nil
}

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
	Stage string `tf:"-"`
	// Version of the model.
	// Wire name: 'version'
	Version         string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st DeleteTransitionRequestRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteTransitionRequestRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteTransitionRequestRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteTransitionRequestRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteTransitionRequestRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteTransitionRequestRequestToPb(st *DeleteTransitionRequestRequest) (*mlpb.DeleteTransitionRequestRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteTransitionRequestRequestPb{}
	pb.Comment = st.Comment
	pb.Creator = st.Creator
	pb.Name = st.Name
	pb.Stage = st.Stage
	pb.Version = st.Version

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func DeleteTransitionRequestRequestFromPb(pb *mlpb.DeleteTransitionRequestRequestPb) (*DeleteTransitionRequestRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTransitionRequestRequest{}
	st.Comment = pb.Comment
	st.Creator = pb.Creator
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type DeleteTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	// Wire name: 'activity'
	Activity *Activity ``
}

func (st DeleteTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := DeleteTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteTransitionRequestResponseToPb(st *DeleteTransitionRequestResponse) (*mlpb.DeleteTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteTransitionRequestResponsePb{}
	activityPb, err := ActivityToPb(st.Activity)
	if err != nil {
		return nil, err
	}
	if activityPb != nil {
		pb.Activity = activityPb
	}

	return pb, nil
}

func DeleteTransitionRequestResponseFromPb(pb *mlpb.DeleteTransitionRequestResponsePb) (*DeleteTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteTransitionRequestResponse{}
	activityField, err := ActivityFromPb(pb.Activity)
	if err != nil {
		return nil, err
	}
	if activityField != nil {
		st.Activity = activityField
	}

	return st, nil
}

type DeleteWebhookRequest struct {
	// Webhook ID required to delete a registry webhook.
	// Wire name: 'id'
	Id string `tf:"-"`
}

func (st DeleteWebhookRequest) MarshalJSON() ([]byte, error) {
	pb, err := DeleteWebhookRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *DeleteWebhookRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.DeleteWebhookRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := DeleteWebhookRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func DeleteWebhookRequestToPb(st *DeleteWebhookRequest) (*mlpb.DeleteWebhookRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.DeleteWebhookRequestPb{}
	pb.Id = st.Id

	return pb, nil
}

func DeleteWebhookRequestFromPb(pb *mlpb.DeleteWebhookRequestPb) (*DeleteWebhookRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &DeleteWebhookRequest{}
	st.Id = pb.Id

	return st, nil
}

// An experiment and its metadata.
type Experiment struct {
	// Location where artifacts for the experiment are stored.
	// Wire name: 'artifact_location'
	ArtifactLocation string ``
	// Creation time
	// Wire name: 'creation_time'
	CreationTime int64 ``
	// Unique identifier for the experiment.
	// Wire name: 'experiment_id'
	ExperimentId string ``
	// Last update time
	// Wire name: 'last_update_time'
	LastUpdateTime int64 ``
	// Current life cycle stage of the experiment: "active" or "deleted".
	// Deleted experiments are not returned by APIs.
	// Wire name: 'lifecycle_stage'
	LifecycleStage string ``
	// Human readable name that identifies the experiment.
	// Wire name: 'name'
	Name string ``
	// Tags: Additional metadata key-value pairs.
	// Wire name: 'tags'
	Tags            []ExperimentTag ``
	ForceSendFields []string        `tf:"-"`
}

func (st Experiment) MarshalJSON() ([]byte, error) {
	pb, err := ExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Experiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExperimentToPb(st *Experiment) (*mlpb.ExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ExperimentPb{}
	pb.ArtifactLocation = st.ArtifactLocation
	pb.CreationTime = st.CreationTime
	pb.ExperimentId = st.ExperimentId
	pb.LastUpdateTime = st.LastUpdateTime
	pb.LifecycleStage = st.LifecycleStage
	pb.Name = st.Name

	var tagsPb []mlpb.ExperimentTagPb
	for _, item := range st.Tags {
		itemPb, err := ExperimentTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ExperimentFromPb(pb *mlpb.ExperimentPb) (*Experiment, error) {
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
		item, err := ExperimentTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ExperimentAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ExperimentAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := ExperimentAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExperimentAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ExperimentAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExperimentAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExperimentAccessControlRequestToPb(st *ExperimentAccessControlRequest) (*mlpb.ExperimentAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ExperimentAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := ExperimentPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ExperimentAccessControlRequestFromPb(pb *mlpb.ExperimentAccessControlRequestPb) (*ExperimentAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := ExperimentPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ExperimentAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []ExperimentPermission ``
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string ``
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ExperimentAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := ExperimentAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExperimentAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ExperimentAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExperimentAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExperimentAccessControlResponseToPb(st *ExperimentAccessControlResponse) (*mlpb.ExperimentAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ExperimentAccessControlResponsePb{}

	var allPermissionsPb []mlpb.ExperimentPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := ExperimentPermissionToPb(&item)
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ExperimentAccessControlResponseFromPb(pb *mlpb.ExperimentAccessControlResponsePb) (*ExperimentAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentAccessControlResponse{}

	var allPermissionsField []ExperimentPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := ExperimentPermissionFromPb(&itemPb)
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ExperimentPermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel ``
	ForceSendFields []string                  `tf:"-"`
}

func (st ExperimentPermission) MarshalJSON() ([]byte, error) {
	pb, err := ExperimentPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExperimentPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ExperimentPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExperimentPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExperimentPermissionToPb(st *ExperimentPermission) (*mlpb.ExperimentPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ExperimentPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := ExperimentPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ExperimentPermissionFromPb(pb *mlpb.ExperimentPermissionPb) (*ExperimentPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := ExperimentPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Permission level
type ExperimentPermissionLevel string

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

// Values returns all possible values for ExperimentPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *ExperimentPermissionLevel) Values() []ExperimentPermissionLevel {
	return []ExperimentPermissionLevel{
		ExperimentPermissionLevelCanEdit,
		ExperimentPermissionLevelCanManage,
		ExperimentPermissionLevelCanRead,
	}
}

// Type always returns ExperimentPermissionLevel to satisfy [pflag.Value] interface
func (f *ExperimentPermissionLevel) Type() string {
	return "ExperimentPermissionLevel"
}

func ExperimentPermissionLevelToPb(st *ExperimentPermissionLevel) (*mlpb.ExperimentPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.ExperimentPermissionLevelPb(*st)
	return &pb, nil
}

func ExperimentPermissionLevelFromPb(pb *mlpb.ExperimentPermissionLevelPb) (*ExperimentPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := ExperimentPermissionLevel(*pb)
	return &st, nil
}

type ExperimentPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []ExperimentAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ExperimentPermissions) MarshalJSON() ([]byte, error) {
	pb, err := ExperimentPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExperimentPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ExperimentPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExperimentPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExperimentPermissionsToPb(st *ExperimentPermissions) (*mlpb.ExperimentPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ExperimentPermissionsPb{}

	var accessControlListPb []mlpb.ExperimentAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := ExperimentAccessControlResponseToPb(&item)
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ExperimentPermissionsFromPb(pb *mlpb.ExperimentPermissionsPb) (*ExperimentPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermissions{}

	var accessControlListField []ExperimentAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := ExperimentAccessControlResponseFromPb(&itemPb)
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ExperimentPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel ExperimentPermissionLevel ``
	ForceSendFields []string                  `tf:"-"`
}

func (st ExperimentPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := ExperimentPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExperimentPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ExperimentPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExperimentPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExperimentPermissionsDescriptionToPb(st *ExperimentPermissionsDescription) (*mlpb.ExperimentPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ExperimentPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := ExperimentPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ExperimentPermissionsDescriptionFromPb(pb *mlpb.ExperimentPermissionsDescriptionPb) (*ExperimentPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := ExperimentPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ExperimentPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []ExperimentAccessControlRequest ``
	// The experiment for which to get or manage permissions.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func (st ExperimentPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ExperimentPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExperimentPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ExperimentPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExperimentPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExperimentPermissionsRequestToPb(st *ExperimentPermissionsRequest) (*mlpb.ExperimentPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ExperimentPermissionsRequestPb{}

	var accessControlListPb []mlpb.ExperimentAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := ExperimentAccessControlRequestToPb(&item)
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

func ExperimentPermissionsRequestFromPb(pb *mlpb.ExperimentPermissionsRequestPb) (*ExperimentPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentPermissionsRequest{}

	var accessControlListField []ExperimentAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := ExperimentAccessControlRequestFromPb(&itemPb)
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
	Key string ``
	// The tag value.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ExperimentTag) MarshalJSON() ([]byte, error) {
	pb, err := ExperimentTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ExperimentTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ExperimentTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ExperimentTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ExperimentTagToPb(st *ExperimentTag) (*mlpb.ExperimentTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ExperimentTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ExperimentTagFromPb(pb *mlpb.ExperimentTagPb) (*ExperimentTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ExperimentTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Feature for model version.
type Feature struct {
	// Feature name
	// Wire name: 'feature_name'
	FeatureName string ``
	// Feature table id
	// Wire name: 'feature_table_id'
	FeatureTableId string ``
	// Feature table name
	// Wire name: 'feature_table_name'
	FeatureTableName string   ``
	ForceSendFields  []string `tf:"-"`
}

func (st Feature) MarshalJSON() ([]byte, error) {
	pb, err := FeatureToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Feature) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.FeaturePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FeatureFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FeatureToPb(st *Feature) (*mlpb.FeaturePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.FeaturePb{}
	pb.FeatureName = st.FeatureName
	pb.FeatureTableId = st.FeatureTableId
	pb.FeatureTableName = st.FeatureTableName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FeatureFromPb(pb *mlpb.FeaturePb) (*Feature, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Feature{}
	st.FeatureName = pb.FeatureName
	st.FeatureTableId = pb.FeatureTableId
	st.FeatureTableName = pb.FeatureTableName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type FeatureLineage struct {
	// List of feature specs that contain this feature.
	// Wire name: 'feature_specs'
	FeatureSpecs []FeatureLineageFeatureSpec ``
	// List of Unity Catalog models that were trained on this feature.
	// Wire name: 'models'
	Models []FeatureLineageModel ``
	// List of online features that use this feature as source.
	// Wire name: 'online_features'
	OnlineFeatures []FeatureLineageOnlineFeature ``
}

func (st FeatureLineage) MarshalJSON() ([]byte, error) {
	pb, err := FeatureLineageToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FeatureLineage) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.FeatureLineagePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FeatureLineageFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FeatureLineageToPb(st *FeatureLineage) (*mlpb.FeatureLineagePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.FeatureLineagePb{}

	var featureSpecsPb []mlpb.FeatureLineageFeatureSpecPb
	for _, item := range st.FeatureSpecs {
		itemPb, err := FeatureLineageFeatureSpecToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			featureSpecsPb = append(featureSpecsPb, *itemPb)
		}
	}
	pb.FeatureSpecs = featureSpecsPb

	var modelsPb []mlpb.FeatureLineageModelPb
	for _, item := range st.Models {
		itemPb, err := FeatureLineageModelToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			modelsPb = append(modelsPb, *itemPb)
		}
	}
	pb.Models = modelsPb

	var onlineFeaturesPb []mlpb.FeatureLineageOnlineFeaturePb
	for _, item := range st.OnlineFeatures {
		itemPb, err := FeatureLineageOnlineFeatureToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onlineFeaturesPb = append(onlineFeaturesPb, *itemPb)
		}
	}
	pb.OnlineFeatures = onlineFeaturesPb

	return pb, nil
}

func FeatureLineageFromPb(pb *mlpb.FeatureLineagePb) (*FeatureLineage, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FeatureLineage{}

	var featureSpecsField []FeatureLineageFeatureSpec
	for _, itemPb := range pb.FeatureSpecs {
		item, err := FeatureLineageFeatureSpecFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			featureSpecsField = append(featureSpecsField, *item)
		}
	}
	st.FeatureSpecs = featureSpecsField

	var modelsField []FeatureLineageModel
	for _, itemPb := range pb.Models {
		item, err := FeatureLineageModelFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			modelsField = append(modelsField, *item)
		}
	}
	st.Models = modelsField

	var onlineFeaturesField []FeatureLineageOnlineFeature
	for _, itemPb := range pb.OnlineFeatures {
		item, err := FeatureLineageOnlineFeatureFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onlineFeaturesField = append(onlineFeaturesField, *item)
		}
	}
	st.OnlineFeatures = onlineFeaturesField

	return st, nil
}

type FeatureLineageFeatureSpec struct {
	// The full name of the feature spec in Unity Catalog.
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st FeatureLineageFeatureSpec) MarshalJSON() ([]byte, error) {
	pb, err := FeatureLineageFeatureSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FeatureLineageFeatureSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.FeatureLineageFeatureSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FeatureLineageFeatureSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FeatureLineageFeatureSpecToPb(st *FeatureLineageFeatureSpec) (*mlpb.FeatureLineageFeatureSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.FeatureLineageFeatureSpecPb{}
	pb.Name = st.Name

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FeatureLineageFeatureSpecFromPb(pb *mlpb.FeatureLineageFeatureSpecPb) (*FeatureLineageFeatureSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FeatureLineageFeatureSpec{}
	st.Name = pb.Name

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type FeatureLineageModel struct {
	// The full name of the model in Unity Catalog.
	// Wire name: 'name'
	Name string ``
	// The version of the model.
	// Wire name: 'version'
	Version         int64    ``
	ForceSendFields []string `tf:"-"`
}

func (st FeatureLineageModel) MarshalJSON() ([]byte, error) {
	pb, err := FeatureLineageModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FeatureLineageModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.FeatureLineageModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FeatureLineageModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FeatureLineageModelToPb(st *FeatureLineageModel) (*mlpb.FeatureLineageModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.FeatureLineageModelPb{}
	pb.Name = st.Name
	pb.Version = st.Version

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FeatureLineageModelFromPb(pb *mlpb.FeatureLineageModelPb) (*FeatureLineageModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FeatureLineageModel{}
	st.Name = pb.Name
	st.Version = pb.Version

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type FeatureLineageOnlineFeature struct {
	// The name of the online feature (column name).
	// Wire name: 'feature_name'
	FeatureName string ``
	// The full name of the online table in Unity Catalog.
	// Wire name: 'table_name'
	TableName       string   ``
	ForceSendFields []string `tf:"-"`
}

func (st FeatureLineageOnlineFeature) MarshalJSON() ([]byte, error) {
	pb, err := FeatureLineageOnlineFeatureToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FeatureLineageOnlineFeature) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.FeatureLineageOnlineFeaturePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FeatureLineageOnlineFeatureFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FeatureLineageOnlineFeatureToPb(st *FeatureLineageOnlineFeature) (*mlpb.FeatureLineageOnlineFeaturePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.FeatureLineageOnlineFeaturePb{}
	pb.FeatureName = st.FeatureName
	pb.TableName = st.TableName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FeatureLineageOnlineFeatureFromPb(pb *mlpb.FeatureLineageOnlineFeaturePb) (*FeatureLineageOnlineFeature, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FeatureLineageOnlineFeature{}
	st.FeatureName = pb.FeatureName
	st.TableName = pb.TableName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Feature list wrap all the features for a model version
type FeatureList struct {

	// Wire name: 'features'
	Features []Feature ``
}

func (st FeatureList) MarshalJSON() ([]byte, error) {
	pb, err := FeatureListToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FeatureList) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.FeatureListPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FeatureListFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FeatureListToPb(st *FeatureList) (*mlpb.FeatureListPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.FeatureListPb{}

	var featuresPb []mlpb.FeaturePb
	for _, item := range st.Features {
		itemPb, err := FeatureToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			featuresPb = append(featuresPb, *itemPb)
		}
	}
	pb.Features = featuresPb

	return pb, nil
}

func FeatureListFromPb(pb *mlpb.FeatureListPb) (*FeatureList, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FeatureList{}

	var featuresField []Feature
	for _, itemPb := range pb.Features {
		item, err := FeatureFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			featuresField = append(featuresField, *item)
		}
	}
	st.Features = featuresField

	return st, nil
}

// Represents a tag on a feature in a feature table.
type FeatureTag struct {

	// Wire name: 'key'
	Key string ``

	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st FeatureTag) MarshalJSON() ([]byte, error) {
	pb, err := FeatureTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FeatureTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.FeatureTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FeatureTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FeatureTagToPb(st *FeatureTag) (*mlpb.FeatureTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.FeatureTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FeatureTagFromPb(pb *mlpb.FeatureTagPb) (*FeatureTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FeatureTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Metadata of a single artifact file or directory.
type FileInfo struct {
	// The size in bytes of the file. Unset for directories.
	// Wire name: 'file_size'
	FileSize int64 ``
	// Whether the path is a directory.
	// Wire name: 'is_dir'
	IsDir bool ``
	// The path relative to the root artifact directory run.
	// Wire name: 'path'
	Path            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st FileInfo) MarshalJSON() ([]byte, error) {
	pb, err := FileInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FileInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.FileInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FileInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FileInfoToPb(st *FileInfo) (*mlpb.FileInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.FileInfoPb{}
	pb.FileSize = st.FileSize
	pb.IsDir = st.IsDir
	pb.Path = st.Path

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func FileInfoFromPb(pb *mlpb.FileInfoPb) (*FileInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FileInfo{}
	st.FileSize = pb.FileSize
	st.IsDir = pb.IsDir
	st.Path = pb.Path

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type FinalizeLoggedModelRequest struct {
	// The ID of the logged model to finalize.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
	// Whether or not the model is ready for use.
	// ``"LOGGED_MODEL_UPLOAD_FAILED"`` indicates that something went wrong when
	// logging the model weights / agent code.
	// Wire name: 'status'
	Status LoggedModelStatus ``
}

func (st FinalizeLoggedModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := FinalizeLoggedModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FinalizeLoggedModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.FinalizeLoggedModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FinalizeLoggedModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FinalizeLoggedModelRequestToPb(st *FinalizeLoggedModelRequest) (*mlpb.FinalizeLoggedModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.FinalizeLoggedModelRequestPb{}
	pb.ModelId = st.ModelId
	statusPb, err := LoggedModelStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	return pb, nil
}

func FinalizeLoggedModelRequestFromPb(pb *mlpb.FinalizeLoggedModelRequestPb) (*FinalizeLoggedModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FinalizeLoggedModelRequest{}
	st.ModelId = pb.ModelId
	statusField, err := LoggedModelStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	return st, nil
}

type FinalizeLoggedModelResponse struct {
	// The updated logged model.
	// Wire name: 'model'
	Model *LoggedModel ``
}

func (st FinalizeLoggedModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := FinalizeLoggedModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *FinalizeLoggedModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.FinalizeLoggedModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := FinalizeLoggedModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func FinalizeLoggedModelResponseToPb(st *FinalizeLoggedModelResponse) (*mlpb.FinalizeLoggedModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.FinalizeLoggedModelResponsePb{}
	modelPb, err := LoggedModelToPb(st.Model)
	if err != nil {
		return nil, err
	}
	if modelPb != nil {
		pb.Model = modelPb
	}

	return pb, nil
}

func FinalizeLoggedModelResponseFromPb(pb *mlpb.FinalizeLoggedModelResponsePb) (*FinalizeLoggedModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &FinalizeLoggedModelResponse{}
	modelField, err := LoggedModelFromPb(pb.Model)
	if err != nil {
		return nil, err
	}
	if modelField != nil {
		st.Model = modelField
	}

	return st, nil
}

// Represents a forecasting experiment with its unique identifier, URL, and
// state.
type ForecastingExperiment struct {
	// The unique ID for the forecasting experiment.
	// Wire name: 'experiment_id'
	ExperimentId string ``
	// The URL to the forecasting experiment page.
	// Wire name: 'experiment_page_url'
	ExperimentPageUrl string ``
	// The current state of the forecasting experiment.
	// Wire name: 'state'
	State           ForecastingExperimentState ``
	ForceSendFields []string                   `tf:"-"`
}

func (st ForecastingExperiment) MarshalJSON() ([]byte, error) {
	pb, err := ForecastingExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ForecastingExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ForecastingExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ForecastingExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ForecastingExperimentToPb(st *ForecastingExperiment) (*mlpb.ForecastingExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ForecastingExperimentPb{}
	pb.ExperimentId = st.ExperimentId
	pb.ExperimentPageUrl = st.ExperimentPageUrl
	statePb, err := ForecastingExperimentStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ForecastingExperimentFromPb(pb *mlpb.ForecastingExperimentPb) (*ForecastingExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ForecastingExperiment{}
	st.ExperimentId = pb.ExperimentId
	st.ExperimentPageUrl = pb.ExperimentPageUrl
	stateField, err := ForecastingExperimentStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ForecastingExperimentState string

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

// Values returns all possible values for ForecastingExperimentState.
//
// There is no guarantee on the order of the values in the slice.
func (f *ForecastingExperimentState) Values() []ForecastingExperimentState {
	return []ForecastingExperimentState{
		ForecastingExperimentStateCancelled,
		ForecastingExperimentStateFailed,
		ForecastingExperimentStatePending,
		ForecastingExperimentStateRunning,
		ForecastingExperimentStateSucceeded,
	}
}

// Type always returns ForecastingExperimentState to satisfy [pflag.Value] interface
func (f *ForecastingExperimentState) Type() string {
	return "ForecastingExperimentState"
}

func ForecastingExperimentStateToPb(st *ForecastingExperimentState) (*mlpb.ForecastingExperimentStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.ForecastingExperimentStatePb(*st)
	return &pb, nil
}

func ForecastingExperimentStateFromPb(pb *mlpb.ForecastingExperimentStatePb) (*ForecastingExperimentState, error) {
	if pb == nil {
		return nil, nil
	}
	st := ForecastingExperimentState(*pb)
	return &st, nil
}

type GetByNameRequest struct {
	// Name of the associated experiment.
	// Wire name: 'experiment_name'
	ExperimentName string `tf:"-"`
}

func (st GetByNameRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetByNameRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetByNameRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetByNameRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetByNameRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetByNameRequestToPb(st *GetByNameRequest) (*mlpb.GetByNameRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetByNameRequestPb{}
	pb.ExperimentName = st.ExperimentName

	return pb, nil
}

func GetByNameRequestFromPb(pb *mlpb.GetByNameRequestPb) (*GetByNameRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetByNameRequest{}
	st.ExperimentName = pb.ExperimentName

	return st, nil
}

type GetExperimentByNameResponse struct {
	// Experiment details.
	// Wire name: 'experiment'
	Experiment *Experiment ``
}

func (st GetExperimentByNameResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetExperimentByNameResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetExperimentByNameResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetExperimentByNameResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetExperimentByNameResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetExperimentByNameResponseToPb(st *GetExperimentByNameResponse) (*mlpb.GetExperimentByNameResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetExperimentByNameResponsePb{}
	experimentPb, err := ExperimentToPb(st.Experiment)
	if err != nil {
		return nil, err
	}
	if experimentPb != nil {
		pb.Experiment = experimentPb
	}

	return pb, nil
}

func GetExperimentByNameResponseFromPb(pb *mlpb.GetExperimentByNameResponsePb) (*GetExperimentByNameResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentByNameResponse{}
	experimentField, err := ExperimentFromPb(pb.Experiment)
	if err != nil {
		return nil, err
	}
	if experimentField != nil {
		st.Experiment = experimentField
	}

	return st, nil
}

type GetExperimentPermissionLevelsRequest struct {
	// The experiment for which to get or manage permissions.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func (st GetExperimentPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetExperimentPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetExperimentPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetExperimentPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetExperimentPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetExperimentPermissionLevelsRequestToPb(st *GetExperimentPermissionLevelsRequest) (*mlpb.GetExperimentPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetExperimentPermissionLevelsRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func GetExperimentPermissionLevelsRequestFromPb(pb *mlpb.GetExperimentPermissionLevelsRequestPb) (*GetExperimentPermissionLevelsRequest, error) {
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
	PermissionLevels []ExperimentPermissionsDescription ``
}

func (st GetExperimentPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetExperimentPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetExperimentPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetExperimentPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetExperimentPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetExperimentPermissionLevelsResponseToPb(st *GetExperimentPermissionLevelsResponse) (*mlpb.GetExperimentPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetExperimentPermissionLevelsResponsePb{}

	var permissionLevelsPb []mlpb.ExperimentPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := ExperimentPermissionsDescriptionToPb(&item)
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

func GetExperimentPermissionLevelsResponseFromPb(pb *mlpb.GetExperimentPermissionLevelsResponsePb) (*GetExperimentPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentPermissionLevelsResponse{}

	var permissionLevelsField []ExperimentPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := ExperimentPermissionsDescriptionFromPb(&itemPb)
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

type GetExperimentPermissionsRequest struct {
	// The experiment for which to get or manage permissions.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func (st GetExperimentPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetExperimentPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetExperimentPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetExperimentPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetExperimentPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetExperimentPermissionsRequestToPb(st *GetExperimentPermissionsRequest) (*mlpb.GetExperimentPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetExperimentPermissionsRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func GetExperimentPermissionsRequestFromPb(pb *mlpb.GetExperimentPermissionsRequestPb) (*GetExperimentPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentPermissionsRequest{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

type GetExperimentRequest struct {
	// ID of the associated experiment.
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func (st GetExperimentRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetExperimentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetExperimentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetExperimentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetExperimentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetExperimentRequestToPb(st *GetExperimentRequest) (*mlpb.GetExperimentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetExperimentRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func GetExperimentRequestFromPb(pb *mlpb.GetExperimentRequestPb) (*GetExperimentRequest, error) {
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
	Experiment *Experiment ``
}

func (st GetExperimentResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetExperimentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetExperimentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetExperimentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetExperimentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetExperimentResponseToPb(st *GetExperimentResponse) (*mlpb.GetExperimentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetExperimentResponsePb{}
	experimentPb, err := ExperimentToPb(st.Experiment)
	if err != nil {
		return nil, err
	}
	if experimentPb != nil {
		pb.Experiment = experimentPb
	}

	return pb, nil
}

func GetExperimentResponseFromPb(pb *mlpb.GetExperimentResponsePb) (*GetExperimentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetExperimentResponse{}
	experimentField, err := ExperimentFromPb(pb.Experiment)
	if err != nil {
		return nil, err
	}
	if experimentField != nil {
		st.Experiment = experimentField
	}

	return st, nil
}

type GetFeatureLineageRequest struct {
	// The name of the feature.
	// Wire name: 'feature_name'
	FeatureName string `tf:"-"`
	// The full name of the feature table in Unity Catalog.
	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func (st GetFeatureLineageRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetFeatureLineageRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetFeatureLineageRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetFeatureLineageRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetFeatureLineageRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetFeatureLineageRequestToPb(st *GetFeatureLineageRequest) (*mlpb.GetFeatureLineageRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetFeatureLineageRequestPb{}
	pb.FeatureName = st.FeatureName
	pb.TableName = st.TableName

	return pb, nil
}

func GetFeatureLineageRequestFromPb(pb *mlpb.GetFeatureLineageRequestPb) (*GetFeatureLineageRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFeatureLineageRequest{}
	st.FeatureName = pb.FeatureName
	st.TableName = pb.TableName

	return st, nil
}

type GetFeatureTagRequest struct {

	// Wire name: 'feature_name'
	FeatureName string `tf:"-"`

	// Wire name: 'key'
	Key string `tf:"-"`

	// Wire name: 'table_name'
	TableName string `tf:"-"`
}

func (st GetFeatureTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetFeatureTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetFeatureTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetFeatureTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetFeatureTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetFeatureTagRequestToPb(st *GetFeatureTagRequest) (*mlpb.GetFeatureTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetFeatureTagRequestPb{}
	pb.FeatureName = st.FeatureName
	pb.Key = st.Key
	pb.TableName = st.TableName

	return pb, nil
}

func GetFeatureTagRequestFromPb(pb *mlpb.GetFeatureTagRequestPb) (*GetFeatureTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetFeatureTagRequest{}
	st.FeatureName = pb.FeatureName
	st.Key = pb.Key
	st.TableName = pb.TableName

	return st, nil
}

type GetForecastingExperimentRequest struct {
	// The unique ID of a forecasting experiment
	// Wire name: 'experiment_id'
	ExperimentId string `tf:"-"`
}

func (st GetForecastingExperimentRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetForecastingExperimentRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetForecastingExperimentRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetForecastingExperimentRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetForecastingExperimentRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetForecastingExperimentRequestToPb(st *GetForecastingExperimentRequest) (*mlpb.GetForecastingExperimentRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetForecastingExperimentRequestPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func GetForecastingExperimentRequestFromPb(pb *mlpb.GetForecastingExperimentRequestPb) (*GetForecastingExperimentRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetForecastingExperimentRequest{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

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
	RunUuid         string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st GetHistoryRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetHistoryRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetHistoryRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetHistoryRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetHistoryRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetHistoryRequestToPb(st *GetHistoryRequest) (*mlpb.GetHistoryRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetHistoryRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.MetricKey = st.MetricKey
	pb.PageToken = st.PageToken
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetHistoryRequestFromPb(pb *mlpb.GetHistoryRequestPb) (*GetHistoryRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetHistoryRequest{}
	st.MaxResults = pb.MaxResults
	st.MetricKey = pb.MetricKey
	st.PageToken = pb.PageToken
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetLatestVersionsRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string ``
	// List of stages.
	// Wire name: 'stages'
	Stages []string ``
}

func (st GetLatestVersionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetLatestVersionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetLatestVersionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetLatestVersionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetLatestVersionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetLatestVersionsRequestToPb(st *GetLatestVersionsRequest) (*mlpb.GetLatestVersionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetLatestVersionsRequestPb{}
	pb.Name = st.Name
	pb.Stages = st.Stages

	return pb, nil
}

func GetLatestVersionsRequestFromPb(pb *mlpb.GetLatestVersionsRequestPb) (*GetLatestVersionsRequest, error) {
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
	ModelVersions []ModelVersion ``
}

func (st GetLatestVersionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetLatestVersionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetLatestVersionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetLatestVersionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetLatestVersionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetLatestVersionsResponseToPb(st *GetLatestVersionsResponse) (*mlpb.GetLatestVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetLatestVersionsResponsePb{}

	var modelVersionsPb []mlpb.ModelVersionPb
	for _, item := range st.ModelVersions {
		itemPb, err := ModelVersionToPb(&item)
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

func GetLatestVersionsResponseFromPb(pb *mlpb.GetLatestVersionsResponsePb) (*GetLatestVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLatestVersionsResponse{}

	var modelVersionsField []ModelVersion
	for _, itemPb := range pb.ModelVersions {
		item, err := ModelVersionFromPb(&itemPb)
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

type GetLoggedModelRequest struct {
	// The ID of the logged model to retrieve.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
}

func (st GetLoggedModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetLoggedModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetLoggedModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetLoggedModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetLoggedModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetLoggedModelRequestToPb(st *GetLoggedModelRequest) (*mlpb.GetLoggedModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetLoggedModelRequestPb{}
	pb.ModelId = st.ModelId

	return pb, nil
}

func GetLoggedModelRequestFromPb(pb *mlpb.GetLoggedModelRequestPb) (*GetLoggedModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLoggedModelRequest{}
	st.ModelId = pb.ModelId

	return st, nil
}

type GetLoggedModelResponse struct {
	// The retrieved logged model.
	// Wire name: 'model'
	Model *LoggedModel ``
}

func (st GetLoggedModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetLoggedModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetLoggedModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetLoggedModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetLoggedModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetLoggedModelResponseToPb(st *GetLoggedModelResponse) (*mlpb.GetLoggedModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetLoggedModelResponsePb{}
	modelPb, err := LoggedModelToPb(st.Model)
	if err != nil {
		return nil, err
	}
	if modelPb != nil {
		pb.Model = modelPb
	}

	return pb, nil
}

func GetLoggedModelResponseFromPb(pb *mlpb.GetLoggedModelResponsePb) (*GetLoggedModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetLoggedModelResponse{}
	modelField, err := LoggedModelFromPb(pb.Model)
	if err != nil {
		return nil, err
	}
	if modelField != nil {
		st.Model = modelField
	}

	return st, nil
}

type GetMetricHistoryResponse struct {
	// All logged values for this metric if `max_results` is not specified in
	// the request or if the total count of metrics returned is less than the
	// service level pagination threshold. Otherwise, this is one page of
	// results.
	// Wire name: 'metrics'
	Metrics []Metric ``
	// A token that can be used to issue a query for the next page of metric
	// history values. A missing token indicates that no additional metrics are
	// available to fetch.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st GetMetricHistoryResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetMetricHistoryResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetMetricHistoryResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetMetricHistoryResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetMetricHistoryResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetMetricHistoryResponseToPb(st *GetMetricHistoryResponse) (*mlpb.GetMetricHistoryResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetMetricHistoryResponsePb{}

	var metricsPb []mlpb.MetricPb
	for _, item := range st.Metrics {
		itemPb, err := MetricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			metricsPb = append(metricsPb, *itemPb)
		}
	}
	pb.Metrics = metricsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetMetricHistoryResponseFromPb(pb *mlpb.GetMetricHistoryResponsePb) (*GetMetricHistoryResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetMetricHistoryResponse{}

	var metricsField []Metric
	for _, itemPb := range pb.Metrics {
		item, err := MetricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			metricsField = append(metricsField, *item)
		}
	}
	st.Metrics = metricsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetModelRequest struct {
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func (st GetModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetModelRequestToPb(st *GetModelRequest) (*mlpb.GetModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetModelRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetModelRequestFromPb(pb *mlpb.GetModelRequestPb) (*GetModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetModelResponse struct {

	// Wire name: 'registered_model_databricks'
	RegisteredModelDatabricks *ModelDatabricks ``
}

func (st GetModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetModelResponseToPb(st *GetModelResponse) (*mlpb.GetModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetModelResponsePb{}
	registeredModelDatabricksPb, err := ModelDatabricksToPb(st.RegisteredModelDatabricks)
	if err != nil {
		return nil, err
	}
	if registeredModelDatabricksPb != nil {
		pb.RegisteredModelDatabricks = registeredModelDatabricksPb
	}

	return pb, nil
}

func GetModelResponseFromPb(pb *mlpb.GetModelResponsePb) (*GetModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelResponse{}
	registeredModelDatabricksField, err := ModelDatabricksFromPb(pb.RegisteredModelDatabricks)
	if err != nil {
		return nil, err
	}
	if registeredModelDatabricksField != nil {
		st.RegisteredModelDatabricks = registeredModelDatabricksField
	}

	return st, nil
}

type GetModelVersionDownloadUriRequest struct {
	// Name of the registered model
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number
	// Wire name: 'version'
	Version string `tf:"-"`
}

func (st GetModelVersionDownloadUriRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetModelVersionDownloadUriRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetModelVersionDownloadUriRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetModelVersionDownloadUriRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetModelVersionDownloadUriRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetModelVersionDownloadUriRequestToPb(st *GetModelVersionDownloadUriRequest) (*mlpb.GetModelVersionDownloadUriRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetModelVersionDownloadUriRequestPb{}
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

func GetModelVersionDownloadUriRequestFromPb(pb *mlpb.GetModelVersionDownloadUriRequestPb) (*GetModelVersionDownloadUriRequest, error) {
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
	ArtifactUri     string   ``
	ForceSendFields []string `tf:"-"`
}

func (st GetModelVersionDownloadUriResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetModelVersionDownloadUriResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetModelVersionDownloadUriResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetModelVersionDownloadUriResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetModelVersionDownloadUriResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetModelVersionDownloadUriResponseToPb(st *GetModelVersionDownloadUriResponse) (*mlpb.GetModelVersionDownloadUriResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetModelVersionDownloadUriResponsePb{}
	pb.ArtifactUri = st.ArtifactUri

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetModelVersionDownloadUriResponseFromPb(pb *mlpb.GetModelVersionDownloadUriResponsePb) (*GetModelVersionDownloadUriResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionDownloadUriResponse{}
	st.ArtifactUri = pb.ArtifactUri

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetModelVersionRequest struct {
	// Name of the registered model
	// Wire name: 'name'
	Name string `tf:"-"`
	// Model version number
	// Wire name: 'version'
	Version string `tf:"-"`
}

func (st GetModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetModelVersionRequestToPb(st *GetModelVersionRequest) (*mlpb.GetModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetModelVersionRequestPb{}
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

func GetModelVersionRequestFromPb(pb *mlpb.GetModelVersionRequestPb) (*GetModelVersionRequest, error) {
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
	ModelVersion *ModelVersion ``
}

func (st GetModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetModelVersionResponseToPb(st *GetModelVersionResponse) (*mlpb.GetModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetModelVersionResponsePb{}
	modelVersionPb, err := ModelVersionToPb(st.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionPb != nil {
		pb.ModelVersion = modelVersionPb
	}

	return pb, nil
}

func GetModelVersionResponseFromPb(pb *mlpb.GetModelVersionResponsePb) (*GetModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetModelVersionResponse{}
	modelVersionField, err := ModelVersionFromPb(pb.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionField != nil {
		st.ModelVersion = modelVersionField
	}

	return st, nil
}

type GetOnlineStoreRequest struct {
	// Name of the online store to get.
	// Wire name: 'name'
	Name string `tf:"-"`
}

func (st GetOnlineStoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetOnlineStoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetOnlineStoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetOnlineStoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetOnlineStoreRequestToPb(st *GetOnlineStoreRequest) (*mlpb.GetOnlineStoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetOnlineStoreRequestPb{}
	pb.Name = st.Name

	return pb, nil
}

func GetOnlineStoreRequestFromPb(pb *mlpb.GetOnlineStoreRequestPb) (*GetOnlineStoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetOnlineStoreRequest{}
	st.Name = pb.Name

	return st, nil
}

type GetRegisteredModelPermissionLevelsRequest struct {
	// The registered model for which to get or manage permissions.
	// Wire name: 'registered_model_id'
	RegisteredModelId string `tf:"-"`
}

func (st GetRegisteredModelPermissionLevelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRegisteredModelPermissionLevelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRegisteredModelPermissionLevelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetRegisteredModelPermissionLevelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRegisteredModelPermissionLevelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetRegisteredModelPermissionLevelsRequestToPb(st *GetRegisteredModelPermissionLevelsRequest) (*mlpb.GetRegisteredModelPermissionLevelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetRegisteredModelPermissionLevelsRequestPb{}
	pb.RegisteredModelId = st.RegisteredModelId

	return pb, nil
}

func GetRegisteredModelPermissionLevelsRequestFromPb(pb *mlpb.GetRegisteredModelPermissionLevelsRequestPb) (*GetRegisteredModelPermissionLevelsRequest, error) {
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
	PermissionLevels []RegisteredModelPermissionsDescription ``
}

func (st GetRegisteredModelPermissionLevelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetRegisteredModelPermissionLevelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRegisteredModelPermissionLevelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetRegisteredModelPermissionLevelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRegisteredModelPermissionLevelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetRegisteredModelPermissionLevelsResponseToPb(st *GetRegisteredModelPermissionLevelsResponse) (*mlpb.GetRegisteredModelPermissionLevelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetRegisteredModelPermissionLevelsResponsePb{}

	var permissionLevelsPb []mlpb.RegisteredModelPermissionsDescriptionPb
	for _, item := range st.PermissionLevels {
		itemPb, err := RegisteredModelPermissionsDescriptionToPb(&item)
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

func GetRegisteredModelPermissionLevelsResponseFromPb(pb *mlpb.GetRegisteredModelPermissionLevelsResponsePb) (*GetRegisteredModelPermissionLevelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRegisteredModelPermissionLevelsResponse{}

	var permissionLevelsField []RegisteredModelPermissionsDescription
	for _, itemPb := range pb.PermissionLevels {
		item, err := RegisteredModelPermissionsDescriptionFromPb(&itemPb)
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

type GetRegisteredModelPermissionsRequest struct {
	// The registered model for which to get or manage permissions.
	// Wire name: 'registered_model_id'
	RegisteredModelId string `tf:"-"`
}

func (st GetRegisteredModelPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRegisteredModelPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRegisteredModelPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetRegisteredModelPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRegisteredModelPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetRegisteredModelPermissionsRequestToPb(st *GetRegisteredModelPermissionsRequest) (*mlpb.GetRegisteredModelPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetRegisteredModelPermissionsRequestPb{}
	pb.RegisteredModelId = st.RegisteredModelId

	return pb, nil
}

func GetRegisteredModelPermissionsRequestFromPb(pb *mlpb.GetRegisteredModelPermissionsRequestPb) (*GetRegisteredModelPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRegisteredModelPermissionsRequest{}
	st.RegisteredModelId = pb.RegisteredModelId

	return st, nil
}

type GetRunRequest struct {
	// ID of the run to fetch. Must be provided.
	// Wire name: 'run_id'
	RunId string `tf:"-"`
	// [Deprecated, use `run_id` instead] ID of the run to fetch. This field
	// will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid         string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st GetRunRequest) MarshalJSON() ([]byte, error) {
	pb, err := GetRunRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRunRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetRunRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRunRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetRunRequestToPb(st *GetRunRequest) (*mlpb.GetRunRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetRunRequestPb{}
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func GetRunRequestFromPb(pb *mlpb.GetRunRequestPb) (*GetRunRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunRequest{}
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type GetRunResponse struct {
	// Run metadata (name, start time, etc) and data (metrics, params, and
	// tags).
	// Wire name: 'run'
	Run *Run ``
}

func (st GetRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := GetRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *GetRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.GetRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := GetRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func GetRunResponseToPb(st *GetRunResponse) (*mlpb.GetRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.GetRunResponsePb{}
	runPb, err := RunToPb(st.Run)
	if err != nil {
		return nil, err
	}
	if runPb != nil {
		pb.Run = runPb
	}

	return pb, nil
}

func GetRunResponseFromPb(pb *mlpb.GetRunResponsePb) (*GetRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &GetRunResponse{}
	runField, err := RunFromPb(pb.Run)
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
	Authorization string ``
	// Enable/disable SSL certificate validation. Default is true. For
	// self-signed certificates, this field must be false AND the destination
	// server must disable certificate validation as well. For security
	// purposes, it is encouraged to perform secret validation with the
	// HMAC-encoded portion of the payload and acknowledge the risk associated
	// with disabling hostname validation whereby it becomes more likely that
	// requests can be maliciously routed to an unintended host.
	// Wire name: 'enable_ssl_verification'
	EnableSslVerification bool ``
	// Shared secret required for HMAC encoding payload. The HMAC-encoded
	// payload will be sent in the header as: { "X-Databricks-Signature":
	// $encoded_payload }.
	// Wire name: 'secret'
	Secret string ``
	// External HTTPS URL called on event trigger (by using a POST request).
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (st HttpUrlSpec) MarshalJSON() ([]byte, error) {
	pb, err := HttpUrlSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *HttpUrlSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.HttpUrlSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := HttpUrlSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func HttpUrlSpecToPb(st *HttpUrlSpec) (*mlpb.HttpUrlSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.HttpUrlSpecPb{}
	pb.Authorization = st.Authorization
	pb.EnableSslVerification = st.EnableSslVerification
	pb.Secret = st.Secret
	pb.Url = st.Url

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func HttpUrlSpecFromPb(pb *mlpb.HttpUrlSpecPb) (*HttpUrlSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &HttpUrlSpec{}
	st.Authorization = pb.Authorization
	st.EnableSslVerification = pb.EnableSslVerification
	st.Secret = pb.Secret
	st.Url = pb.Url

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
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
	EnableSslVerification bool ``
	// External HTTPS URL called on event trigger (by using a POST request).
	// Wire name: 'url'
	Url             string   ``
	ForceSendFields []string `tf:"-"`
}

func (st HttpUrlSpecWithoutSecret) MarshalJSON() ([]byte, error) {
	pb, err := HttpUrlSpecWithoutSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *HttpUrlSpecWithoutSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.HttpUrlSpecWithoutSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := HttpUrlSpecWithoutSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func HttpUrlSpecWithoutSecretToPb(st *HttpUrlSpecWithoutSecret) (*mlpb.HttpUrlSpecWithoutSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.HttpUrlSpecWithoutSecretPb{}
	pb.EnableSslVerification = st.EnableSslVerification
	pb.Url = st.Url

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func HttpUrlSpecWithoutSecretFromPb(pb *mlpb.HttpUrlSpecWithoutSecretPb) (*HttpUrlSpecWithoutSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &HttpUrlSpecWithoutSecret{}
	st.EnableSslVerification = pb.EnableSslVerification
	st.Url = pb.Url

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Tag for a dataset input.
type InputTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string ``
	// The tag value.
	// Wire name: 'value'
	Value string ``
}

func (st InputTag) MarshalJSON() ([]byte, error) {
	pb, err := InputTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *InputTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.InputTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := InputTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func InputTagToPb(st *InputTag) (*mlpb.InputTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.InputTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	return pb, nil
}

func InputTagFromPb(pb *mlpb.InputTagPb) (*InputTag, error) {
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
	AccessToken string ``
	// ID of the job that the webhook runs.
	// Wire name: 'job_id'
	JobId string ``
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	// Wire name: 'workspace_url'
	WorkspaceUrl    string   ``
	ForceSendFields []string `tf:"-"`
}

func (st JobSpec) MarshalJSON() ([]byte, error) {
	pb, err := JobSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.JobSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobSpecToPb(st *JobSpec) (*mlpb.JobSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.JobSpecPb{}
	pb.AccessToken = st.AccessToken
	pb.JobId = st.JobId
	pb.WorkspaceUrl = st.WorkspaceUrl

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobSpecFromPb(pb *mlpb.JobSpecPb) (*JobSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSpec{}
	st.AccessToken = pb.AccessToken
	st.JobId = pb.JobId
	st.WorkspaceUrl = pb.WorkspaceUrl

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type JobSpecWithoutSecret struct {
	// ID of the job that the webhook runs.
	// Wire name: 'job_id'
	JobId string ``
	// URL of the workspace containing the job that this webhook runs. If not
	// specified, the job’s workspace URL is assumed to be the same as the
	// workspace where the webhook is created.
	// Wire name: 'workspace_url'
	WorkspaceUrl    string   ``
	ForceSendFields []string `tf:"-"`
}

func (st JobSpecWithoutSecret) MarshalJSON() ([]byte, error) {
	pb, err := JobSpecWithoutSecretToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *JobSpecWithoutSecret) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.JobSpecWithoutSecretPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := JobSpecWithoutSecretFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func JobSpecWithoutSecretToPb(st *JobSpecWithoutSecret) (*mlpb.JobSpecWithoutSecretPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.JobSpecWithoutSecretPb{}
	pb.JobId = st.JobId
	pb.WorkspaceUrl = st.WorkspaceUrl

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func JobSpecWithoutSecretFromPb(pb *mlpb.JobSpecWithoutSecretPb) (*JobSpecWithoutSecret, error) {
	if pb == nil {
		return nil, nil
	}
	st := &JobSpecWithoutSecret{}
	st.JobId = pb.JobId
	st.WorkspaceUrl = pb.WorkspaceUrl

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

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
	RunUuid         string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListArtifactsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListArtifactsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListArtifactsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListArtifactsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListArtifactsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListArtifactsRequestToPb(st *ListArtifactsRequest) (*mlpb.ListArtifactsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListArtifactsRequestPb{}
	pb.PageToken = st.PageToken
	pb.Path = st.Path
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListArtifactsRequestFromPb(pb *mlpb.ListArtifactsRequestPb) (*ListArtifactsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListArtifactsRequest{}
	st.PageToken = pb.PageToken
	st.Path = pb.Path
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListArtifactsResponse struct {
	// The file location and metadata for artifacts.
	// Wire name: 'files'
	Files []FileInfo ``
	// The token that can be used to retrieve the next page of artifact results.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// The root artifact directory for the run.
	// Wire name: 'root_uri'
	RootUri         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ListArtifactsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListArtifactsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListArtifactsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListArtifactsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListArtifactsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListArtifactsResponseToPb(st *ListArtifactsResponse) (*mlpb.ListArtifactsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListArtifactsResponsePb{}

	var filesPb []mlpb.FileInfoPb
	for _, item := range st.Files {
		itemPb, err := FileInfoToPb(&item)
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListArtifactsResponseFromPb(pb *mlpb.ListArtifactsResponsePb) (*ListArtifactsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListArtifactsResponse{}

	var filesField []FileInfo
	for _, itemPb := range pb.Files {
		item, err := FileInfoFromPb(&itemPb)
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

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
	ViewType        ViewType `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListExperimentsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListExperimentsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListExperimentsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListExperimentsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListExperimentsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListExperimentsRequestToPb(st *ListExperimentsRequest) (*mlpb.ListExperimentsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListExperimentsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken
	viewTypePb, err := ViewTypeToPb(&st.ViewType)
	if err != nil {
		return nil, err
	}
	if viewTypePb != nil {
		pb.ViewType = *viewTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListExperimentsRequestFromPb(pb *mlpb.ListExperimentsRequestPb) (*ListExperimentsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExperimentsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken
	viewTypeField, err := ViewTypeFromPb(&pb.ViewType)
	if err != nil {
		return nil, err
	}
	if viewTypeField != nil {
		st.ViewType = *viewTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListExperimentsResponse struct {
	// Paginated Experiments beginning with the first item on the requested
	// page.
	// Wire name: 'experiments'
	Experiments []Experiment ``
	// Token that can be used to retrieve the next page of experiments. Empty
	// token means no more experiment is available for retrieval.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ListExperimentsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListExperimentsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListExperimentsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListExperimentsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListExperimentsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListExperimentsResponseToPb(st *ListExperimentsResponse) (*mlpb.ListExperimentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListExperimentsResponsePb{}

	var experimentsPb []mlpb.ExperimentPb
	for _, item := range st.Experiments {
		itemPb, err := ExperimentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			experimentsPb = append(experimentsPb, *itemPb)
		}
	}
	pb.Experiments = experimentsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListExperimentsResponseFromPb(pb *mlpb.ListExperimentsResponsePb) (*ListExperimentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListExperimentsResponse{}

	var experimentsField []Experiment
	for _, itemPb := range pb.Experiments {
		item, err := ExperimentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			experimentsField = append(experimentsField, *item)
		}
	}
	st.Experiments = experimentsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListFeatureTagsRequest struct {

	// Wire name: 'feature_name'
	FeatureName string `tf:"-"`
	// The maximum number of results to return.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Pagination token to go to the next page based on a previous query.
	// Wire name: 'page_token'
	PageToken string `tf:"-"`

	// Wire name: 'table_name'
	TableName       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListFeatureTagsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListFeatureTagsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListFeatureTagsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListFeatureTagsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListFeatureTagsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListFeatureTagsRequestToPb(st *ListFeatureTagsRequest) (*mlpb.ListFeatureTagsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListFeatureTagsRequestPb{}
	pb.FeatureName = st.FeatureName
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken
	pb.TableName = st.TableName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListFeatureTagsRequestFromPb(pb *mlpb.ListFeatureTagsRequestPb) (*ListFeatureTagsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFeatureTagsRequest{}
	st.FeatureName = pb.FeatureName
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken
	st.TableName = pb.TableName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Response message for ListFeatureTag.
type ListFeatureTagsResponse struct {

	// Wire name: 'feature_tags'
	FeatureTags []FeatureTag ``
	// Pagination token to request the next page of results for this query.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ListFeatureTagsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListFeatureTagsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListFeatureTagsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListFeatureTagsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListFeatureTagsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListFeatureTagsResponseToPb(st *ListFeatureTagsResponse) (*mlpb.ListFeatureTagsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListFeatureTagsResponsePb{}

	var featureTagsPb []mlpb.FeatureTagPb
	for _, item := range st.FeatureTags {
		itemPb, err := FeatureTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			featureTagsPb = append(featureTagsPb, *itemPb)
		}
	}
	pb.FeatureTags = featureTagsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListFeatureTagsResponseFromPb(pb *mlpb.ListFeatureTagsResponsePb) (*ListFeatureTagsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListFeatureTagsResponse{}

	var featureTagsField []FeatureTag
	for _, itemPb := range pb.FeatureTags {
		item, err := FeatureTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			featureTagsField = append(featureTagsField, *item)
		}
	}
	st.FeatureTags = featureTagsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListModelsRequest struct {
	// Maximum number of registered models desired. Max threshold is 1000.
	// Wire name: 'max_results'
	MaxResults int64 `tf:"-"`
	// Pagination token to go to the next page based on a previous query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListModelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListModelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListModelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListModelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListModelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListModelsRequestToPb(st *ListModelsRequest) (*mlpb.ListModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListModelsRequestPb{}
	pb.MaxResults = st.MaxResults
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListModelsRequestFromPb(pb *mlpb.ListModelsRequestPb) (*ListModelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelsRequest{}
	st.MaxResults = pb.MaxResults
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListModelsResponse struct {
	// Pagination token to request next page of models for the same query.
	// Wire name: 'next_page_token'
	NextPageToken string ``

	// Wire name: 'registered_models'
	RegisteredModels []Model  ``
	ForceSendFields  []string `tf:"-"`
}

func (st ListModelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListModelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListModelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListModelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListModelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListModelsResponseToPb(st *ListModelsResponse) (*mlpb.ListModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListModelsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var registeredModelsPb []mlpb.ModelPb
	for _, item := range st.RegisteredModels {
		itemPb, err := ModelToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			registeredModelsPb = append(registeredModelsPb, *itemPb)
		}
	}
	pb.RegisteredModels = registeredModelsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListModelsResponseFromPb(pb *mlpb.ListModelsResponsePb) (*ListModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListModelsResponse{}
	st.NextPageToken = pb.NextPageToken

	var registeredModelsField []Model
	for _, itemPb := range pb.RegisteredModels {
		item, err := ModelFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			registeredModelsField = append(registeredModelsField, *item)
		}
	}
	st.RegisteredModels = registeredModelsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListOnlineStoresRequest struct {
	// The maximum number of results to return. Defaults to 100 if not
	// specified.
	// Wire name: 'page_size'
	PageSize int `tf:"-"`
	// Pagination token to go to the next page based on a previous query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListOnlineStoresRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListOnlineStoresRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListOnlineStoresRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListOnlineStoresRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListOnlineStoresRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListOnlineStoresRequestToPb(st *ListOnlineStoresRequest) (*mlpb.ListOnlineStoresRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListOnlineStoresRequestPb{}
	pb.PageSize = st.PageSize
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListOnlineStoresRequestFromPb(pb *mlpb.ListOnlineStoresRequestPb) (*ListOnlineStoresRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListOnlineStoresRequest{}
	st.PageSize = pb.PageSize
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListOnlineStoresResponse struct {
	// Pagination token to request the next page of results for this query.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// List of online stores.
	// Wire name: 'online_stores'
	OnlineStores    []OnlineStore ``
	ForceSendFields []string      `tf:"-"`
}

func (st ListOnlineStoresResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListOnlineStoresResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListOnlineStoresResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListOnlineStoresResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListOnlineStoresResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListOnlineStoresResponseToPb(st *ListOnlineStoresResponse) (*mlpb.ListOnlineStoresResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListOnlineStoresResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var onlineStoresPb []mlpb.OnlineStorePb
	for _, item := range st.OnlineStores {
		itemPb, err := OnlineStoreToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			onlineStoresPb = append(onlineStoresPb, *itemPb)
		}
	}
	pb.OnlineStores = onlineStoresPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListOnlineStoresResponseFromPb(pb *mlpb.ListOnlineStoresResponsePb) (*ListOnlineStoresResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListOnlineStoresResponse{}
	st.NextPageToken = pb.NextPageToken

	var onlineStoresField []OnlineStore
	for _, itemPb := range pb.OnlineStores {
		item, err := OnlineStoreFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			onlineStoresField = append(onlineStoresField, *item)
		}
	}
	st.OnlineStores = onlineStoresField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListRegistryWebhooks struct {
	// Token that can be used to retrieve the next page of artifact results
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// Array of registry webhooks.
	// Wire name: 'webhooks'
	Webhooks        []RegistryWebhook ``
	ForceSendFields []string          `tf:"-"`
}

func (st ListRegistryWebhooks) MarshalJSON() ([]byte, error) {
	pb, err := ListRegistryWebhooksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListRegistryWebhooks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListRegistryWebhooksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListRegistryWebhooksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListRegistryWebhooksToPb(st *ListRegistryWebhooks) (*mlpb.ListRegistryWebhooksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListRegistryWebhooksPb{}
	pb.NextPageToken = st.NextPageToken

	var webhooksPb []mlpb.RegistryWebhookPb
	for _, item := range st.Webhooks {
		itemPb, err := RegistryWebhookToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			webhooksPb = append(webhooksPb, *itemPb)
		}
	}
	pb.Webhooks = webhooksPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListRegistryWebhooksFromPb(pb *mlpb.ListRegistryWebhooksPb) (*ListRegistryWebhooks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListRegistryWebhooks{}
	st.NextPageToken = pb.NextPageToken

	var webhooksField []RegistryWebhook
	for _, itemPb := range pb.Webhooks {
		item, err := RegistryWebhookFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			webhooksField = append(webhooksField, *item)
		}
	}
	st.Webhooks = webhooksField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ListTransitionRequestsRequest struct {
	// Name of the registered model.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Version of the model.
	// Wire name: 'version'
	Version string `tf:"-"`
}

func (st ListTransitionRequestsRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListTransitionRequestsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListTransitionRequestsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListTransitionRequestsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListTransitionRequestsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListTransitionRequestsRequestToPb(st *ListTransitionRequestsRequest) (*mlpb.ListTransitionRequestsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListTransitionRequestsRequestPb{}
	pb.Name = st.Name
	pb.Version = st.Version

	return pb, nil
}

func ListTransitionRequestsRequestFromPb(pb *mlpb.ListTransitionRequestsRequestPb) (*ListTransitionRequestsRequest, error) {
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
	Requests []Activity ``
}

func (st ListTransitionRequestsResponse) MarshalJSON() ([]byte, error) {
	pb, err := ListTransitionRequestsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListTransitionRequestsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListTransitionRequestsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListTransitionRequestsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListTransitionRequestsResponseToPb(st *ListTransitionRequestsResponse) (*mlpb.ListTransitionRequestsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListTransitionRequestsResponsePb{}

	var requestsPb []mlpb.ActivityPb
	for _, item := range st.Requests {
		itemPb, err := ActivityToPb(&item)
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

func ListTransitionRequestsResponseFromPb(pb *mlpb.ListTransitionRequestsResponsePb) (*ListTransitionRequestsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListTransitionRequestsResponse{}

	var requestsField []Activity
	for _, itemPb := range pb.Requests {
		item, err := ActivityFromPb(&itemPb)
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

type ListWebhooksRequest struct {
	// Events that trigger the webhook. * `MODEL_VERSION_CREATED`: A new model
	// version was created for the associated model.
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
	//
	// If `events` is specified, any webhook with one or more of the specified
	// trigger events is included in the output. If `events` is not specified,
	// webhooks of all event types are included in the output.
	// Wire name: 'events'
	Events []RegistryWebhookEvent `tf:"-"`

	// Wire name: 'max_results'
	MaxResults int64 `tf:"-"`
	// Registered model name If not specified, all webhooks associated with the
	// specified events are listed, regardless of their associated model.
	// Wire name: 'model_name'
	ModelName string `tf:"-"`
	// Token indicating the page of artifact results to fetch
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st ListWebhooksRequest) MarshalJSON() ([]byte, error) {
	pb, err := ListWebhooksRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ListWebhooksRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ListWebhooksRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ListWebhooksRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ListWebhooksRequestToPb(st *ListWebhooksRequest) (*mlpb.ListWebhooksRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ListWebhooksRequestPb{}

	var eventsPb []mlpb.RegistryWebhookEventPb
	for _, item := range st.Events {
		itemPb, err := RegistryWebhookEventToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			eventsPb = append(eventsPb, *itemPb)
		}
	}
	pb.Events = eventsPb
	pb.MaxResults = st.MaxResults
	pb.ModelName = st.ModelName
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ListWebhooksRequestFromPb(pb *mlpb.ListWebhooksRequestPb) (*ListWebhooksRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ListWebhooksRequest{}

	var eventsField []RegistryWebhookEvent
	for _, itemPb := range pb.Events {
		item, err := RegistryWebhookEventFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			eventsField = append(eventsField, *item)
		}
	}
	st.Events = eventsField
	st.MaxResults = pb.MaxResults
	st.ModelName = pb.ModelName
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type LogBatch struct {
	// Metrics to log. A single request can contain up to 1000 metrics, and up
	// to 1000 metrics, params, and tags in total.
	// Wire name: 'metrics'
	Metrics []Metric ``
	// Params to log. A single request can contain up to 100 params, and up to
	// 1000 metrics, params, and tags in total.
	// Wire name: 'params'
	Params []Param ``
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string ``
	// Tags to log. A single request can contain up to 100 tags, and up to 1000
	// metrics, params, and tags in total.
	// Wire name: 'tags'
	Tags            []RunTag ``
	ForceSendFields []string `tf:"-"`
}

func (st LogBatch) MarshalJSON() ([]byte, error) {
	pb, err := LogBatchToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogBatch) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LogBatchPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogBatchFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogBatchToPb(st *LogBatch) (*mlpb.LogBatchPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LogBatchPb{}

	var metricsPb []mlpb.MetricPb
	for _, item := range st.Metrics {
		itemPb, err := MetricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			metricsPb = append(metricsPb, *itemPb)
		}
	}
	pb.Metrics = metricsPb

	var paramsPb []mlpb.ParamPb
	for _, item := range st.Params {
		itemPb, err := ParamToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			paramsPb = append(paramsPb, *itemPb)
		}
	}
	pb.Params = paramsPb
	pb.RunId = st.RunId

	var tagsPb []mlpb.RunTagPb
	for _, item := range st.Tags {
		itemPb, err := RunTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LogBatchFromPb(pb *mlpb.LogBatchPb) (*LogBatch, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogBatch{}

	var metricsField []Metric
	for _, itemPb := range pb.Metrics {
		item, err := MetricFromPb(&itemPb)
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
		item, err := ParamFromPb(&itemPb)
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
		item, err := RunTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type LogInputs struct {
	// Dataset inputs
	// Wire name: 'datasets'
	Datasets []DatasetInput ``
	// Model inputs
	// Wire name: 'models'
	Models []ModelInput ``
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId string ``
}

func (st LogInputs) MarshalJSON() ([]byte, error) {
	pb, err := LogInputsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogInputs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LogInputsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogInputsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogInputsToPb(st *LogInputs) (*mlpb.LogInputsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LogInputsPb{}

	var datasetsPb []mlpb.DatasetInputPb
	for _, item := range st.Datasets {
		itemPb, err := DatasetInputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			datasetsPb = append(datasetsPb, *itemPb)
		}
	}
	pb.Datasets = datasetsPb

	var modelsPb []mlpb.ModelInputPb
	for _, item := range st.Models {
		itemPb, err := ModelInputToPb(&item)
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

func LogInputsFromPb(pb *mlpb.LogInputsPb) (*LogInputs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogInputs{}

	var datasetsField []DatasetInput
	for _, itemPb := range pb.Datasets {
		item, err := DatasetInputFromPb(&itemPb)
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
		item, err := ModelInputFromPb(&itemPb)
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

type LogLoggedModelParamsRequest struct {
	// The ID of the logged model to log params for.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
	// Parameters to attach to the model.
	// Wire name: 'params'
	Params []LoggedModelParameter ``
}

func (st LogLoggedModelParamsRequest) MarshalJSON() ([]byte, error) {
	pb, err := LogLoggedModelParamsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogLoggedModelParamsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LogLoggedModelParamsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogLoggedModelParamsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogLoggedModelParamsRequestToPb(st *LogLoggedModelParamsRequest) (*mlpb.LogLoggedModelParamsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LogLoggedModelParamsRequestPb{}
	pb.ModelId = st.ModelId

	var paramsPb []mlpb.LoggedModelParameterPb
	for _, item := range st.Params {
		itemPb, err := LoggedModelParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			paramsPb = append(paramsPb, *itemPb)
		}
	}
	pb.Params = paramsPb

	return pb, nil
}

func LogLoggedModelParamsRequestFromPb(pb *mlpb.LogLoggedModelParamsRequestPb) (*LogLoggedModelParamsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogLoggedModelParamsRequest{}
	st.ModelId = pb.ModelId

	var paramsField []LoggedModelParameter
	for _, itemPb := range pb.Params {
		item, err := LoggedModelParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			paramsField = append(paramsField, *item)
		}
	}
	st.Params = paramsField

	return st, nil
}

type LogMetric struct {
	// Dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	// Wire name: 'dataset_digest'
	DatasetDigest string ``
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	// Wire name: 'dataset_name'
	DatasetName string ``
	// Name of the metric.
	// Wire name: 'key'
	Key string ``
	// ID of the logged model associated with the metric, if applicable
	// Wire name: 'model_id'
	ModelId string ``
	// ID of the run under which to log the metric. Must be provided.
	// Wire name: 'run_id'
	RunId string ``
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// metric. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string ``
	// Step at which to log the metric
	// Wire name: 'step'
	Step int64 ``
	// Unix timestamp in milliseconds at the time metric was logged.
	// Wire name: 'timestamp'
	Timestamp int64 ``
	// Double value of the metric being logged.
	// Wire name: 'value'
	Value           float64  ``
	ForceSendFields []string `tf:"-"`
}

func (st LogMetric) MarshalJSON() ([]byte, error) {
	pb, err := LogMetricToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogMetric) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LogMetricPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogMetricFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogMetricToPb(st *LogMetric) (*mlpb.LogMetricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LogMetricPb{}
	pb.DatasetDigest = st.DatasetDigest
	pb.DatasetName = st.DatasetName
	pb.Key = st.Key
	pb.ModelId = st.ModelId
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid
	pb.Step = st.Step
	pb.Timestamp = st.Timestamp
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LogMetricFromPb(pb *mlpb.LogMetricPb) (*LogMetric, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type LogModel struct {
	// MLmodel file in json format.
	// Wire name: 'model_json'
	ModelJson string ``
	// ID of the run to log under
	// Wire name: 'run_id'
	RunId           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st LogModel) MarshalJSON() ([]byte, error) {
	pb, err := LogModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LogModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogModelToPb(st *LogModel) (*mlpb.LogModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LogModelPb{}
	pb.ModelJson = st.ModelJson
	pb.RunId = st.RunId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LogModelFromPb(pb *mlpb.LogModelPb) (*LogModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogModel{}
	st.ModelJson = pb.ModelJson
	st.RunId = pb.RunId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type LogOutputsRequest struct {
	// The model outputs from the Run.
	// Wire name: 'models'
	Models []ModelOutput ``
	// The ID of the Run from which to log outputs.
	// Wire name: 'run_id'
	RunId string ``
}

func (st LogOutputsRequest) MarshalJSON() ([]byte, error) {
	pb, err := LogOutputsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogOutputsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LogOutputsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogOutputsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogOutputsRequestToPb(st *LogOutputsRequest) (*mlpb.LogOutputsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LogOutputsRequestPb{}

	var modelsPb []mlpb.ModelOutputPb
	for _, item := range st.Models {
		itemPb, err := ModelOutputToPb(&item)
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

func LogOutputsRequestFromPb(pb *mlpb.LogOutputsRequestPb) (*LogOutputsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogOutputsRequest{}

	var modelsField []ModelOutput
	for _, itemPb := range pb.Models {
		item, err := ModelOutputFromPb(&itemPb)
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

type LogParam struct {
	// Name of the param. Maximum size is 255 bytes.
	// Wire name: 'key'
	Key string ``
	// ID of the run under which to log the param. Must be provided.
	// Wire name: 'run_id'
	RunId string ``
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// param. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string ``
	// String value of the param being logged. Maximum size is 500 bytes.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st LogParam) MarshalJSON() ([]byte, error) {
	pb, err := LogParamToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LogParam) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LogParamPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LogParamFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LogParamToPb(st *LogParam) (*mlpb.LogParamPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LogParamPb{}
	pb.Key = st.Key
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LogParamFromPb(pb *mlpb.LogParamPb) (*LogParam, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LogParam{}
	st.Key = pb.Key
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// A logged model message includes logged model attributes, tags, registration
// info, params, and linked run metrics.
type LoggedModel struct {
	// The params and metrics attached to the logged model.
	// Wire name: 'data'
	Data *LoggedModelData ``
	// The logged model attributes such as model ID, status, tags, etc.
	// Wire name: 'info'
	Info *LoggedModelInfo ``
}

func (st LoggedModel) MarshalJSON() ([]byte, error) {
	pb, err := LoggedModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LoggedModel) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LoggedModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LoggedModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LoggedModelToPb(st *LoggedModel) (*mlpb.LoggedModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LoggedModelPb{}
	dataPb, err := LoggedModelDataToPb(st.Data)
	if err != nil {
		return nil, err
	}
	if dataPb != nil {
		pb.Data = dataPb
	}
	infoPb, err := LoggedModelInfoToPb(st.Info)
	if err != nil {
		return nil, err
	}
	if infoPb != nil {
		pb.Info = infoPb
	}

	return pb, nil
}

func LoggedModelFromPb(pb *mlpb.LoggedModelPb) (*LoggedModel, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LoggedModel{}
	dataField, err := LoggedModelDataFromPb(pb.Data)
	if err != nil {
		return nil, err
	}
	if dataField != nil {
		st.Data = dataField
	}
	infoField, err := LoggedModelInfoFromPb(pb.Info)
	if err != nil {
		return nil, err
	}
	if infoField != nil {
		st.Info = infoField
	}

	return st, nil
}

// A LoggedModelData message includes logged model params and linked metrics.
type LoggedModelData struct {
	// Performance metrics linked to the model.
	// Wire name: 'metrics'
	Metrics []Metric ``
	// Immutable string key-value pairs of the model.
	// Wire name: 'params'
	Params []LoggedModelParameter ``
}

func (st LoggedModelData) MarshalJSON() ([]byte, error) {
	pb, err := LoggedModelDataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LoggedModelData) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LoggedModelDataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LoggedModelDataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LoggedModelDataToPb(st *LoggedModelData) (*mlpb.LoggedModelDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LoggedModelDataPb{}

	var metricsPb []mlpb.MetricPb
	for _, item := range st.Metrics {
		itemPb, err := MetricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			metricsPb = append(metricsPb, *itemPb)
		}
	}
	pb.Metrics = metricsPb

	var paramsPb []mlpb.LoggedModelParameterPb
	for _, item := range st.Params {
		itemPb, err := LoggedModelParameterToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			paramsPb = append(paramsPb, *itemPb)
		}
	}
	pb.Params = paramsPb

	return pb, nil
}

func LoggedModelDataFromPb(pb *mlpb.LoggedModelDataPb) (*LoggedModelData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LoggedModelData{}

	var metricsField []Metric
	for _, itemPb := range pb.Metrics {
		item, err := MetricFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			metricsField = append(metricsField, *item)
		}
	}
	st.Metrics = metricsField

	var paramsField []LoggedModelParameter
	for _, itemPb := range pb.Params {
		item, err := LoggedModelParameterFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			paramsField = append(paramsField, *item)
		}
	}
	st.Params = paramsField

	return st, nil
}

// A LoggedModelInfo includes logged model attributes, tags, and registration
// info.
type LoggedModelInfo struct {
	// The URI of the directory where model artifacts are stored.
	// Wire name: 'artifact_uri'
	ArtifactUri string ``
	// The timestamp when the model was created in milliseconds since the UNIX
	// epoch.
	// Wire name: 'creation_timestamp_ms'
	CreationTimestampMs int64 ``
	// The ID of the user or principal that created the model.
	// Wire name: 'creator_id'
	CreatorId int64 ``
	// The ID of the experiment that owns the model.
	// Wire name: 'experiment_id'
	ExperimentId string ``
	// The timestamp when the model was last updated in milliseconds since the
	// UNIX epoch.
	// Wire name: 'last_updated_timestamp_ms'
	LastUpdatedTimestampMs int64 ``
	// The unique identifier for the logged model.
	// Wire name: 'model_id'
	ModelId string ``
	// The type of model, such as ``"Agent"``, ``"Classifier"``, ``"LLM"``.
	// Wire name: 'model_type'
	ModelType string ``
	// The name of the model.
	// Wire name: 'name'
	Name string ``
	// The ID of the run that created the model.
	// Wire name: 'source_run_id'
	SourceRunId string ``
	// The status of whether or not the model is ready for use.
	// Wire name: 'status'
	Status LoggedModelStatus ``
	// Details on the current model status.
	// Wire name: 'status_message'
	StatusMessage string ``
	// Mutable string key-value pairs set on the model.
	// Wire name: 'tags'
	Tags            []LoggedModelTag ``
	ForceSendFields []string         `tf:"-"`
}

func (st LoggedModelInfo) MarshalJSON() ([]byte, error) {
	pb, err := LoggedModelInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LoggedModelInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LoggedModelInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LoggedModelInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LoggedModelInfoToPb(st *LoggedModelInfo) (*mlpb.LoggedModelInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LoggedModelInfoPb{}
	pb.ArtifactUri = st.ArtifactUri
	pb.CreationTimestampMs = st.CreationTimestampMs
	pb.CreatorId = st.CreatorId
	pb.ExperimentId = st.ExperimentId
	pb.LastUpdatedTimestampMs = st.LastUpdatedTimestampMs
	pb.ModelId = st.ModelId
	pb.ModelType = st.ModelType
	pb.Name = st.Name
	pb.SourceRunId = st.SourceRunId
	statusPb, err := LoggedModelStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.StatusMessage = st.StatusMessage

	var tagsPb []mlpb.LoggedModelTagPb
	for _, item := range st.Tags {
		itemPb, err := LoggedModelTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LoggedModelInfoFromPb(pb *mlpb.LoggedModelInfoPb) (*LoggedModelInfo, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LoggedModelInfo{}
	st.ArtifactUri = pb.ArtifactUri
	st.CreationTimestampMs = pb.CreationTimestampMs
	st.CreatorId = pb.CreatorId
	st.ExperimentId = pb.ExperimentId
	st.LastUpdatedTimestampMs = pb.LastUpdatedTimestampMs
	st.ModelId = pb.ModelId
	st.ModelType = pb.ModelType
	st.Name = pb.Name
	st.SourceRunId = pb.SourceRunId
	statusField, err := LoggedModelStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.StatusMessage = pb.StatusMessage

	var tagsField []LoggedModelTag
	for _, itemPb := range pb.Tags {
		item, err := LoggedModelTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Parameter associated with a LoggedModel.
type LoggedModelParameter struct {
	// The key identifying this param.
	// Wire name: 'key'
	Key string ``
	// The value of this param.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st LoggedModelParameter) MarshalJSON() ([]byte, error) {
	pb, err := LoggedModelParameterToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LoggedModelParameter) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LoggedModelParameterPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LoggedModelParameterFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LoggedModelParameterToPb(st *LoggedModelParameter) (*mlpb.LoggedModelParameterPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LoggedModelParameterPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LoggedModelParameterFromPb(pb *mlpb.LoggedModelParameterPb) (*LoggedModelParameter, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LoggedModelParameter{}
	st.Key = pb.Key
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// A LoggedModelStatus enum value represents the status of a logged model.
type LoggedModelStatus string

const LoggedModelStatusLoggedModelPending LoggedModelStatus = `LOGGED_MODEL_PENDING`

const LoggedModelStatusLoggedModelReady LoggedModelStatus = `LOGGED_MODEL_READY`

const LoggedModelStatusLoggedModelUploadFailed LoggedModelStatus = `LOGGED_MODEL_UPLOAD_FAILED`

// String representation for [fmt.Print]
func (f *LoggedModelStatus) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *LoggedModelStatus) Set(v string) error {
	switch v {
	case `LOGGED_MODEL_PENDING`, `LOGGED_MODEL_READY`, `LOGGED_MODEL_UPLOAD_FAILED`:
		*f = LoggedModelStatus(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "LOGGED_MODEL_PENDING", "LOGGED_MODEL_READY", "LOGGED_MODEL_UPLOAD_FAILED"`, v)
	}
}

// Values returns all possible values for LoggedModelStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *LoggedModelStatus) Values() []LoggedModelStatus {
	return []LoggedModelStatus{
		LoggedModelStatusLoggedModelPending,
		LoggedModelStatusLoggedModelReady,
		LoggedModelStatusLoggedModelUploadFailed,
	}
}

// Type always returns LoggedModelStatus to satisfy [pflag.Value] interface
func (f *LoggedModelStatus) Type() string {
	return "LoggedModelStatus"
}

func LoggedModelStatusToPb(st *LoggedModelStatus) (*mlpb.LoggedModelStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.LoggedModelStatusPb(*st)
	return &pb, nil
}

func LoggedModelStatusFromPb(pb *mlpb.LoggedModelStatusPb) (*LoggedModelStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := LoggedModelStatus(*pb)
	return &st, nil
}

// Tag for a LoggedModel.
type LoggedModelTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string ``
	// The tag value.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st LoggedModelTag) MarshalJSON() ([]byte, error) {
	pb, err := LoggedModelTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *LoggedModelTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.LoggedModelTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := LoggedModelTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func LoggedModelTagToPb(st *LoggedModelTag) (*mlpb.LoggedModelTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.LoggedModelTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func LoggedModelTagFromPb(pb *mlpb.LoggedModelTagPb) (*LoggedModelTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &LoggedModelTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Metric associated with a run, represented as a key-value pair.
type Metric struct {
	// The dataset digest of the dataset associated with the metric, e.g. an md5
	// hash of the dataset that uniquely identifies it within datasets of the
	// same name.
	// Wire name: 'dataset_digest'
	DatasetDigest string ``
	// The name of the dataset associated with the metric. E.g.
	// “my.uc.table@2” “nyc-taxi-dataset”, “fantastic-elk-3”
	// Wire name: 'dataset_name'
	DatasetName string ``
	// The key identifying the metric.
	// Wire name: 'key'
	Key string ``
	// The ID of the logged model or registered model version associated with
	// the metric, if applicable.
	// Wire name: 'model_id'
	ModelId string ``
	// The ID of the run containing the metric.
	// Wire name: 'run_id'
	RunId string ``
	// The step at which the metric was logged.
	// Wire name: 'step'
	Step int64 ``
	// The timestamp at which the metric was recorded.
	// Wire name: 'timestamp'
	Timestamp int64 ``
	// The value of the metric.
	// Wire name: 'value'
	Value           float64  ``
	ForceSendFields []string `tf:"-"`
}

func (st Metric) MarshalJSON() ([]byte, error) {
	pb, err := MetricToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Metric) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.MetricPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := MetricFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func MetricToPb(st *Metric) (*mlpb.MetricPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.MetricPb{}
	pb.DatasetDigest = st.DatasetDigest
	pb.DatasetName = st.DatasetName
	pb.Key = st.Key
	pb.ModelId = st.ModelId
	pb.RunId = st.RunId
	pb.Step = st.Step
	pb.Timestamp = st.Timestamp
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func MetricFromPb(pb *mlpb.MetricPb) (*Metric, error) {
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type Model struct {
	// Timestamp recorded when this `registered_model` was created.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``
	// Description of this `registered_model`.
	// Wire name: 'description'
	Description string ``
	// Timestamp recorded when metadata for this `registered_model` was last
	// updated.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// Collection of latest model versions for each stage. Only contains models
	// with current `READY` status.
	// Wire name: 'latest_versions'
	LatestVersions []ModelVersion ``
	// Unique name for the model.
	// Wire name: 'name'
	Name string ``
	// Tags: Additional metadata key-value pairs for this `registered_model`.
	// Wire name: 'tags'
	Tags []ModelTag ``
	// User that created this `registered_model`
	// Wire name: 'user_id'
	UserId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st Model) MarshalJSON() ([]byte, error) {
	pb, err := ModelToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Model) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ModelPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ModelFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ModelToPb(st *Model) (*mlpb.ModelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ModelPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Description = st.Description
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	var latestVersionsPb []mlpb.ModelVersionPb
	for _, item := range st.LatestVersions {
		itemPb, err := ModelVersionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			latestVersionsPb = append(latestVersionsPb, *itemPb)
		}
	}
	pb.LatestVersions = latestVersionsPb
	pb.Name = st.Name

	var tagsPb []mlpb.ModelTagPb
	for _, item := range st.Tags {
		itemPb, err := ModelTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb
	pb.UserId = st.UserId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ModelFromPb(pb *mlpb.ModelPb) (*Model, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Model{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Description = pb.Description
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp

	var latestVersionsField []ModelVersion
	for _, itemPb := range pb.LatestVersions {
		item, err := ModelVersionFromPb(&itemPb)
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
		item, err := ModelTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.UserId = pb.UserId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ModelDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``
	// User-specified description for the object.
	// Wire name: 'description'
	Description string ``
	// Unique identifier for the object.
	// Wire name: 'id'
	Id string ``
	// Last update time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// Array of model versions, each the latest version for its stage.
	// Wire name: 'latest_versions'
	LatestVersions []ModelVersion ``
	// Name of the model.
	// Wire name: 'name'
	Name string ``
	// Permission level granted for the requesting user on this registered model
	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel ``
	// Array of tags associated with the model.
	// Wire name: 'tags'
	Tags []ModelTag ``
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ModelDatabricks) MarshalJSON() ([]byte, error) {
	pb, err := ModelDatabricksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ModelDatabricks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ModelDatabricksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ModelDatabricksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ModelDatabricksToPb(st *ModelDatabricks) (*mlpb.ModelDatabricksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ModelDatabricksPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Description = st.Description
	pb.Id = st.Id
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp

	var latestVersionsPb []mlpb.ModelVersionPb
	for _, item := range st.LatestVersions {
		itemPb, err := ModelVersionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			latestVersionsPb = append(latestVersionsPb, *itemPb)
		}
	}
	pb.LatestVersions = latestVersionsPb
	pb.Name = st.Name
	permissionLevelPb, err := PermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	var tagsPb []mlpb.ModelTagPb
	for _, item := range st.Tags {
		itemPb, err := ModelTagToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			tagsPb = append(tagsPb, *itemPb)
		}
	}
	pb.Tags = tagsPb
	pb.UserId = st.UserId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ModelDatabricksFromPb(pb *mlpb.ModelDatabricksPb) (*ModelDatabricks, error) {
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
		item, err := ModelVersionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			latestVersionsField = append(latestVersionsField, *item)
		}
	}
	st.LatestVersions = latestVersionsField
	st.Name = pb.Name
	permissionLevelField, err := PermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	var tagsField []ModelTag
	for _, itemPb := range pb.Tags {
		item, err := ModelTagFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			tagsField = append(tagsField, *item)
		}
	}
	st.Tags = tagsField
	st.UserId = pb.UserId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Represents a LoggedModel or Registered Model Version input to a Run.
type ModelInput struct {
	// The unique identifier of the model.
	// Wire name: 'model_id'
	ModelId string ``
}

func (st ModelInput) MarshalJSON() ([]byte, error) {
	pb, err := ModelInputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ModelInput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ModelInputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ModelInputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ModelInputToPb(st *ModelInput) (*mlpb.ModelInputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ModelInputPb{}
	pb.ModelId = st.ModelId

	return pb, nil
}

func ModelInputFromPb(pb *mlpb.ModelInputPb) (*ModelInput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelInput{}
	st.ModelId = pb.ModelId

	return st, nil
}

// Represents a LoggedModel output of a Run.
type ModelOutput struct {
	// The unique identifier of the model.
	// Wire name: 'model_id'
	ModelId string ``
	// The step at which the model was produced.
	// Wire name: 'step'
	Step int64 ``
}

func (st ModelOutput) MarshalJSON() ([]byte, error) {
	pb, err := ModelOutputToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ModelOutput) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ModelOutputPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ModelOutputFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ModelOutputToPb(st *ModelOutput) (*mlpb.ModelOutputPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ModelOutputPb{}
	pb.ModelId = st.ModelId
	pb.Step = st.Step

	return pb, nil
}

func ModelOutputFromPb(pb *mlpb.ModelOutputPb) (*ModelOutput, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelOutput{}
	st.ModelId = pb.ModelId
	st.Step = pb.Step

	return st, nil
}

// Tag for a registered model
type ModelTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string ``
	// The tag value.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ModelTag) MarshalJSON() ([]byte, error) {
	pb, err := ModelTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ModelTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ModelTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ModelTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ModelTagToPb(st *ModelTag) (*mlpb.ModelTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ModelTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ModelTagFromPb(pb *mlpb.ModelTagPb) (*ModelTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ModelVersion struct {
	// Timestamp recorded when this `model_version` was created.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``
	// Current stage for this `model_version`.
	// Wire name: 'current_stage'
	CurrentStage string ``
	// Description of this `model_version`.
	// Wire name: 'description'
	Description string ``
	// Timestamp recorded when metadata for this `model_version` was last
	// updated.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// Unique name of the model
	// Wire name: 'name'
	Name string ``
	// MLflow run ID used when creating `model_version`, if `source` was
	// generated by an experiment run stored in MLflow tracking server.
	// Wire name: 'run_id'
	RunId string ``
	// Run Link: Direct link to the run that generated this version
	// Wire name: 'run_link'
	RunLink string ``
	// URI indicating the location of the source model artifacts, used when
	// creating `model_version`
	// Wire name: 'source'
	Source string ``
	// Current status of `model_version`
	// Wire name: 'status'
	Status ModelVersionStatus ``
	// Details on current `status`, if it is pending or failed.
	// Wire name: 'status_message'
	StatusMessage string ``
	// Tags: Additional metadata key-value pairs for this `model_version`.
	// Wire name: 'tags'
	Tags []ModelVersionTag ``
	// User that created this `model_version`.
	// Wire name: 'user_id'
	UserId string ``
	// Model's version number.
	// Wire name: 'version'
	Version         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ModelVersion) MarshalJSON() ([]byte, error) {
	pb, err := ModelVersionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ModelVersion) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ModelVersionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ModelVersionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ModelVersionToPb(st *ModelVersion) (*mlpb.ModelVersionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ModelVersionPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.CurrentStage = st.CurrentStage
	pb.Description = st.Description
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.Name = st.Name
	pb.RunId = st.RunId
	pb.RunLink = st.RunLink
	pb.Source = st.Source
	statusPb, err := ModelVersionStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.StatusMessage = st.StatusMessage

	var tagsPb []mlpb.ModelVersionTagPb
	for _, item := range st.Tags {
		itemPb, err := ModelVersionTagToPb(&item)
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ModelVersionFromPb(pb *mlpb.ModelVersionPb) (*ModelVersion, error) {
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
	statusField, err := ModelVersionStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.StatusMessage = pb.StatusMessage

	var tagsField []ModelVersionTag
	for _, itemPb := range pb.Tags {
		item, err := ModelVersionTagFromPb(&itemPb)
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type ModelVersionDatabricks struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``

	// Wire name: 'current_stage'
	CurrentStage string ``
	// User-specified description for the object.
	// Wire name: 'description'
	Description string ``
	// Email Subscription Status: This is the subscription status of the user to
	// the model version Users get subscribed by interacting with the model
	// version.
	// Wire name: 'email_subscription_status'
	EmailSubscriptionStatus RegistryEmailSubscriptionType ``
	// Feature lineage of `model_version`.
	// Wire name: 'feature_list'
	FeatureList *FeatureList ``
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// Name of the model.
	// Wire name: 'name'
	Name string ``
	// Open requests for this `model_versions`. Gap in sequence number is
	// intentional and is done in order to match field sequence numbers of
	// `ModelVersion` proto message
	// Wire name: 'open_requests'
	OpenRequests []Activity ``

	// Wire name: 'permission_level'
	PermissionLevel PermissionLevel ``
	// Unique identifier for the MLflow tracking run associated with the source
	// model artifacts.
	// Wire name: 'run_id'
	RunId string ``
	// URL of the run associated with the model artifacts. This field is set at
	// model version creation time only for model versions whose source run is
	// from a tracking server that is different from the registry server.
	// Wire name: 'run_link'
	RunLink string ``
	// URI that indicates the location of the source model artifacts. This is
	// used when creating the model version.
	// Wire name: 'source'
	Source string ``

	// Wire name: 'status'
	Status Status ``
	// Details on the current status, for example why registration failed.
	// Wire name: 'status_message'
	StatusMessage string ``
	// Array of tags that are associated with the model version.
	// Wire name: 'tags'
	Tags []ModelVersionTag ``
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId string ``
	// Version of the model.
	// Wire name: 'version'
	Version         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ModelVersionDatabricks) MarshalJSON() ([]byte, error) {
	pb, err := ModelVersionDatabricksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ModelVersionDatabricks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ModelVersionDatabricksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ModelVersionDatabricksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ModelVersionDatabricksToPb(st *ModelVersionDatabricks) (*mlpb.ModelVersionDatabricksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ModelVersionDatabricksPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.CurrentStage = st.CurrentStage
	pb.Description = st.Description
	emailSubscriptionStatusPb, err := RegistryEmailSubscriptionTypeToPb(&st.EmailSubscriptionStatus)
	if err != nil {
		return nil, err
	}
	if emailSubscriptionStatusPb != nil {
		pb.EmailSubscriptionStatus = *emailSubscriptionStatusPb
	}
	featureListPb, err := FeatureListToPb(st.FeatureList)
	if err != nil {
		return nil, err
	}
	if featureListPb != nil {
		pb.FeatureList = featureListPb
	}
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.Name = st.Name

	var openRequestsPb []mlpb.ActivityPb
	for _, item := range st.OpenRequests {
		itemPb, err := ActivityToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			openRequestsPb = append(openRequestsPb, *itemPb)
		}
	}
	pb.OpenRequests = openRequestsPb
	permissionLevelPb, err := PermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.RunId = st.RunId
	pb.RunLink = st.RunLink
	pb.Source = st.Source
	statusPb, err := StatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.StatusMessage = st.StatusMessage

	var tagsPb []mlpb.ModelVersionTagPb
	for _, item := range st.Tags {
		itemPb, err := ModelVersionTagToPb(&item)
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ModelVersionDatabricksFromPb(pb *mlpb.ModelVersionDatabricksPb) (*ModelVersionDatabricks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersionDatabricks{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.CurrentStage = pb.CurrentStage
	st.Description = pb.Description
	emailSubscriptionStatusField, err := RegistryEmailSubscriptionTypeFromPb(&pb.EmailSubscriptionStatus)
	if err != nil {
		return nil, err
	}
	if emailSubscriptionStatusField != nil {
		st.EmailSubscriptionStatus = *emailSubscriptionStatusField
	}
	featureListField, err := FeatureListFromPb(pb.FeatureList)
	if err != nil {
		return nil, err
	}
	if featureListField != nil {
		st.FeatureList = featureListField
	}
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.Name = pb.Name

	var openRequestsField []Activity
	for _, itemPb := range pb.OpenRequests {
		item, err := ActivityFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			openRequestsField = append(openRequestsField, *item)
		}
	}
	st.OpenRequests = openRequestsField
	permissionLevelField, err := PermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.RunId = pb.RunId
	st.RunLink = pb.RunLink
	st.Source = pb.Source
	statusField, err := StatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.StatusMessage = pb.StatusMessage

	var tagsField []ModelVersionTag
	for _, itemPb := range pb.Tags {
		item, err := ModelVersionTagFromPb(&itemPb)
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks.
//
// * `FAILED_REGISTRATION`: Request to register a new model version has failed.
//
// * `READY`: Model version is ready for use.
type ModelVersionStatus string

// Request to register a new model version has failed.
const ModelVersionStatusFailedRegistration ModelVersionStatus = `FAILED_REGISTRATION`

// Request to register a new model version is pending as server performs
// background tasks.
const ModelVersionStatusPendingRegistration ModelVersionStatus = `PENDING_REGISTRATION`

// Model version is ready for use.
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

// Values returns all possible values for ModelVersionStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *ModelVersionStatus) Values() []ModelVersionStatus {
	return []ModelVersionStatus{
		ModelVersionStatusFailedRegistration,
		ModelVersionStatusPendingRegistration,
		ModelVersionStatusReady,
	}
}

// Type always returns ModelVersionStatus to satisfy [pflag.Value] interface
func (f *ModelVersionStatus) Type() string {
	return "ModelVersionStatus"
}

func ModelVersionStatusToPb(st *ModelVersionStatus) (*mlpb.ModelVersionStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.ModelVersionStatusPb(*st)
	return &pb, nil
}

func ModelVersionStatusFromPb(pb *mlpb.ModelVersionStatusPb) (*ModelVersionStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := ModelVersionStatus(*pb)
	return &st, nil
}

type ModelVersionTag struct {
	// The tag key.
	// Wire name: 'key'
	Key string ``
	// The tag value.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st ModelVersionTag) MarshalJSON() ([]byte, error) {
	pb, err := ModelVersionTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *ModelVersionTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ModelVersionTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ModelVersionTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ModelVersionTagToPb(st *ModelVersionTag) (*mlpb.ModelVersionTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ModelVersionTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ModelVersionTagFromPb(pb *mlpb.ModelVersionTagPb) (*ModelVersionTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &ModelVersionTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// An OnlineStore is a logical database instance that stores and serves features
// online.
type OnlineStore struct {
	// The capacity of the online store. Valid values are "CU_1", "CU_2",
	// "CU_4", "CU_8".
	// Wire name: 'capacity'
	Capacity string ``
	// The timestamp when the online store was created.
	// Wire name: 'creation_time'
	CreationTime string `` //legacy
	// The email of the creator of the online store.
	// Wire name: 'creator'
	Creator string ``
	// The name of the online store. This is the unique identifier for the
	// online store.
	// Wire name: 'name'
	Name string ``
	// The number of read replicas for the online store. Defaults to 0.
	// Wire name: 'read_replica_count'
	ReadReplicaCount int ``
	// The current state of the online store.
	// Wire name: 'state'
	State           OnlineStoreState ``
	ForceSendFields []string         `tf:"-"`
}

func (st OnlineStore) MarshalJSON() ([]byte, error) {
	pb, err := OnlineStoreToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *OnlineStore) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.OnlineStorePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := OnlineStoreFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func OnlineStoreToPb(st *OnlineStore) (*mlpb.OnlineStorePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.OnlineStorePb{}
	pb.Capacity = st.Capacity
	pb.CreationTime = st.CreationTime
	pb.Creator = st.Creator
	pb.Name = st.Name
	pb.ReadReplicaCount = st.ReadReplicaCount
	statePb, err := OnlineStoreStateToPb(&st.State)
	if err != nil {
		return nil, err
	}
	if statePb != nil {
		pb.State = *statePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func OnlineStoreFromPb(pb *mlpb.OnlineStorePb) (*OnlineStore, error) {
	if pb == nil {
		return nil, nil
	}
	st := &OnlineStore{}
	st.Capacity = pb.Capacity
	st.CreationTime = pb.CreationTime
	st.Creator = pb.Creator
	st.Name = pb.Name
	st.ReadReplicaCount = pb.ReadReplicaCount
	stateField, err := OnlineStoreStateFromPb(&pb.State)
	if err != nil {
		return nil, err
	}
	if stateField != nil {
		st.State = *stateField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type OnlineStoreState string

const OnlineStoreStateAvailable OnlineStoreState = `AVAILABLE`

const OnlineStoreStateDeleting OnlineStoreState = `DELETING`

const OnlineStoreStateFailingOver OnlineStoreState = `FAILING_OVER`

const OnlineStoreStateStarting OnlineStoreState = `STARTING`

const OnlineStoreStateStopped OnlineStoreState = `STOPPED`

const OnlineStoreStateUpdating OnlineStoreState = `UPDATING`

// String representation for [fmt.Print]
func (f *OnlineStoreState) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *OnlineStoreState) Set(v string) error {
	switch v {
	case `AVAILABLE`, `DELETING`, `FAILING_OVER`, `STARTING`, `STOPPED`, `UPDATING`:
		*f = OnlineStoreState(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "AVAILABLE", "DELETING", "FAILING_OVER", "STARTING", "STOPPED", "UPDATING"`, v)
	}
}

// Values returns all possible values for OnlineStoreState.
//
// There is no guarantee on the order of the values in the slice.
func (f *OnlineStoreState) Values() []OnlineStoreState {
	return []OnlineStoreState{
		OnlineStoreStateAvailable,
		OnlineStoreStateDeleting,
		OnlineStoreStateFailingOver,
		OnlineStoreStateStarting,
		OnlineStoreStateStopped,
		OnlineStoreStateUpdating,
	}
}

// Type always returns OnlineStoreState to satisfy [pflag.Value] interface
func (f *OnlineStoreState) Type() string {
	return "OnlineStoreState"
}

func OnlineStoreStateToPb(st *OnlineStoreState) (*mlpb.OnlineStoreStatePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.OnlineStoreStatePb(*st)
	return &pb, nil
}

func OnlineStoreStateFromPb(pb *mlpb.OnlineStoreStatePb) (*OnlineStoreState, error) {
	if pb == nil {
		return nil, nil
	}
	st := OnlineStoreState(*pb)
	return &st, nil
}

// Param associated with a run.
type Param struct {
	// Key identifying this param.
	// Wire name: 'key'
	Key string ``
	// Value associated with this param.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st Param) MarshalJSON() ([]byte, error) {
	pb, err := ParamToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Param) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.ParamPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := ParamFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func ParamToPb(st *Param) (*mlpb.ParamPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.ParamPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func ParamFromPb(pb *mlpb.ParamPb) (*Param, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Param{}
	st.Key = pb.Key
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Permission level of the requesting user on the object. For what is allowed at
// each level, see [MLflow Model permissions](..).
type PermissionLevel string

const PermissionLevelCanCreateRegisteredModel PermissionLevel = `CAN_CREATE_REGISTERED_MODEL`

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
	case `CAN_CREATE_REGISTERED_MODEL`, `CAN_EDIT`, `CAN_MANAGE`, `CAN_MANAGE_PRODUCTION_VERSIONS`, `CAN_MANAGE_STAGING_VERSIONS`, `CAN_READ`:
		*f = PermissionLevel(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CAN_CREATE_REGISTERED_MODEL", "CAN_EDIT", "CAN_MANAGE", "CAN_MANAGE_PRODUCTION_VERSIONS", "CAN_MANAGE_STAGING_VERSIONS", "CAN_READ"`, v)
	}
}

// Values returns all possible values for PermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *PermissionLevel) Values() []PermissionLevel {
	return []PermissionLevel{
		PermissionLevelCanCreateRegisteredModel,
		PermissionLevelCanEdit,
		PermissionLevelCanManage,
		PermissionLevelCanManageProductionVersions,
		PermissionLevelCanManageStagingVersions,
		PermissionLevelCanRead,
	}
}

// Type always returns PermissionLevel to satisfy [pflag.Value] interface
func (f *PermissionLevel) Type() string {
	return "PermissionLevel"
}

func PermissionLevelToPb(st *PermissionLevel) (*mlpb.PermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.PermissionLevelPb(*st)
	return &pb, nil
}

func PermissionLevelFromPb(pb *mlpb.PermissionLevelPb) (*PermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := PermissionLevel(*pb)
	return &st, nil
}

type PublishSpec struct {
	// The name of the target online store.
	// Wire name: 'online_store'
	OnlineStore string ``
	// The full three-part (catalog, schema, table) name of the online table.
	// Wire name: 'online_table_name'
	OnlineTableName string ``
	// The publish mode of the pipeline that syncs the online table with the
	// source table. Defaults to TRIGGERED if not specified. All publish modes
	// require the source table to have Change Data Feed (CDF) enabled.
	// Wire name: 'publish_mode'
	PublishMode PublishSpecPublishMode ``
}

func (st PublishSpec) MarshalJSON() ([]byte, error) {
	pb, err := PublishSpecToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PublishSpec) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.PublishSpecPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PublishSpecFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PublishSpecToPb(st *PublishSpec) (*mlpb.PublishSpecPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.PublishSpecPb{}
	pb.OnlineStore = st.OnlineStore
	pb.OnlineTableName = st.OnlineTableName
	publishModePb, err := PublishSpecPublishModeToPb(&st.PublishMode)
	if err != nil {
		return nil, err
	}
	if publishModePb != nil {
		pb.PublishMode = *publishModePb
	}

	return pb, nil
}

func PublishSpecFromPb(pb *mlpb.PublishSpecPb) (*PublishSpec, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishSpec{}
	st.OnlineStore = pb.OnlineStore
	st.OnlineTableName = pb.OnlineTableName
	publishModeField, err := PublishSpecPublishModeFromPb(&pb.PublishMode)
	if err != nil {
		return nil, err
	}
	if publishModeField != nil {
		st.PublishMode = *publishModeField
	}

	return st, nil
}

type PublishSpecPublishMode string

const PublishSpecPublishModeContinuous PublishSpecPublishMode = `CONTINUOUS`

const PublishSpecPublishModeTriggered PublishSpecPublishMode = `TRIGGERED`

// String representation for [fmt.Print]
func (f *PublishSpecPublishMode) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *PublishSpecPublishMode) Set(v string) error {
	switch v {
	case `CONTINUOUS`, `TRIGGERED`:
		*f = PublishSpecPublishMode(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "CONTINUOUS", "TRIGGERED"`, v)
	}
}

// Values returns all possible values for PublishSpecPublishMode.
//
// There is no guarantee on the order of the values in the slice.
func (f *PublishSpecPublishMode) Values() []PublishSpecPublishMode {
	return []PublishSpecPublishMode{
		PublishSpecPublishModeContinuous,
		PublishSpecPublishModeTriggered,
	}
}

// Type always returns PublishSpecPublishMode to satisfy [pflag.Value] interface
func (f *PublishSpecPublishMode) Type() string {
	return "PublishSpecPublishMode"
}

func PublishSpecPublishModeToPb(st *PublishSpecPublishMode) (*mlpb.PublishSpecPublishModePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.PublishSpecPublishModePb(*st)
	return &pb, nil
}

func PublishSpecPublishModeFromPb(pb *mlpb.PublishSpecPublishModePb) (*PublishSpecPublishMode, error) {
	if pb == nil {
		return nil, nil
	}
	st := PublishSpecPublishMode(*pb)
	return &st, nil
}

type PublishTableRequest struct {
	// The specification for publishing the online table from the source table.
	// Wire name: 'publish_spec'
	PublishSpec PublishSpec ``
	// The full three-part (catalog, schema, table) name of the source table.
	// Wire name: 'source_table_name'
	SourceTableName string `tf:"-"`
}

func (st PublishTableRequest) MarshalJSON() ([]byte, error) {
	pb, err := PublishTableRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PublishTableRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.PublishTableRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PublishTableRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PublishTableRequestToPb(st *PublishTableRequest) (*mlpb.PublishTableRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.PublishTableRequestPb{}
	publishSpecPb, err := PublishSpecToPb(&st.PublishSpec)
	if err != nil {
		return nil, err
	}
	if publishSpecPb != nil {
		pb.PublishSpec = *publishSpecPb
	}
	pb.SourceTableName = st.SourceTableName

	return pb, nil
}

func PublishTableRequestFromPb(pb *mlpb.PublishTableRequestPb) (*PublishTableRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishTableRequest{}
	publishSpecField, err := PublishSpecFromPb(&pb.PublishSpec)
	if err != nil {
		return nil, err
	}
	if publishSpecField != nil {
		st.PublishSpec = *publishSpecField
	}
	st.SourceTableName = pb.SourceTableName

	return st, nil
}

type PublishTableResponse struct {
	// The full three-part (catalog, schema, table) name of the online table.
	// Wire name: 'online_table_name'
	OnlineTableName string ``
	// The ID of the pipeline that syncs the online table with the source table.
	// Wire name: 'pipeline_id'
	PipelineId      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st PublishTableResponse) MarshalJSON() ([]byte, error) {
	pb, err := PublishTableResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *PublishTableResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.PublishTableResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := PublishTableResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func PublishTableResponseToPb(st *PublishTableResponse) (*mlpb.PublishTableResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.PublishTableResponsePb{}
	pb.OnlineTableName = st.OnlineTableName
	pb.PipelineId = st.PipelineId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func PublishTableResponseFromPb(pb *mlpb.PublishTableResponsePb) (*PublishTableResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &PublishTableResponse{}
	st.OnlineTableName = pb.OnlineTableName
	st.PipelineId = pb.PipelineId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RegisteredModelAccessControlRequest struct {
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``

	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel ``
	// application ID of a service principal
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st RegisteredModelAccessControlRequest) MarshalJSON() ([]byte, error) {
	pb, err := RegisteredModelAccessControlRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RegisteredModelAccessControlRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RegisteredModelAccessControlRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RegisteredModelAccessControlRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RegisteredModelAccessControlRequestToPb(st *RegisteredModelAccessControlRequest) (*mlpb.RegisteredModelAccessControlRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RegisteredModelAccessControlRequestPb{}
	pb.GroupName = st.GroupName
	permissionLevelPb, err := RegisteredModelPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}
	pb.ServicePrincipalName = st.ServicePrincipalName
	pb.UserName = st.UserName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RegisteredModelAccessControlRequestFromPb(pb *mlpb.RegisteredModelAccessControlRequestPb) (*RegisteredModelAccessControlRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAccessControlRequest{}
	st.GroupName = pb.GroupName
	permissionLevelField, err := RegisteredModelPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}
	st.ServicePrincipalName = pb.ServicePrincipalName
	st.UserName = pb.UserName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RegisteredModelAccessControlResponse struct {
	// All permissions.
	// Wire name: 'all_permissions'
	AllPermissions []RegisteredModelPermission ``
	// Display name of the user or service principal.
	// Wire name: 'display_name'
	DisplayName string ``
	// name of the group
	// Wire name: 'group_name'
	GroupName string ``
	// Name of the service principal.
	// Wire name: 'service_principal_name'
	ServicePrincipalName string ``
	// name of the user
	// Wire name: 'user_name'
	UserName        string   ``
	ForceSendFields []string `tf:"-"`
}

func (st RegisteredModelAccessControlResponse) MarshalJSON() ([]byte, error) {
	pb, err := RegisteredModelAccessControlResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RegisteredModelAccessControlResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RegisteredModelAccessControlResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RegisteredModelAccessControlResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RegisteredModelAccessControlResponseToPb(st *RegisteredModelAccessControlResponse) (*mlpb.RegisteredModelAccessControlResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RegisteredModelAccessControlResponsePb{}

	var allPermissionsPb []mlpb.RegisteredModelPermissionPb
	for _, item := range st.AllPermissions {
		itemPb, err := RegisteredModelPermissionToPb(&item)
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RegisteredModelAccessControlResponseFromPb(pb *mlpb.RegisteredModelAccessControlResponsePb) (*RegisteredModelAccessControlResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelAccessControlResponse{}

	var allPermissionsField []RegisteredModelPermission
	for _, itemPb := range pb.AllPermissions {
		item, err := RegisteredModelPermissionFromPb(&itemPb)
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RegisteredModelPermission struct {

	// Wire name: 'inherited'
	Inherited bool ``

	// Wire name: 'inherited_from_object'
	InheritedFromObject []string ``

	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel ``
	ForceSendFields []string                       `tf:"-"`
}

func (st RegisteredModelPermission) MarshalJSON() ([]byte, error) {
	pb, err := RegisteredModelPermissionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RegisteredModelPermission) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RegisteredModelPermissionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RegisteredModelPermissionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RegisteredModelPermissionToPb(st *RegisteredModelPermission) (*mlpb.RegisteredModelPermissionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RegisteredModelPermissionPb{}
	pb.Inherited = st.Inherited
	pb.InheritedFromObject = st.InheritedFromObject
	permissionLevelPb, err := RegisteredModelPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RegisteredModelPermissionFromPb(pb *mlpb.RegisteredModelPermissionPb) (*RegisteredModelPermission, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermission{}
	st.Inherited = pb.Inherited
	st.InheritedFromObject = pb.InheritedFromObject
	permissionLevelField, err := RegisteredModelPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Permission level
type RegisteredModelPermissionLevel string

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

// Values returns all possible values for RegisteredModelPermissionLevel.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegisteredModelPermissionLevel) Values() []RegisteredModelPermissionLevel {
	return []RegisteredModelPermissionLevel{
		RegisteredModelPermissionLevelCanEdit,
		RegisteredModelPermissionLevelCanManage,
		RegisteredModelPermissionLevelCanManageProductionVersions,
		RegisteredModelPermissionLevelCanManageStagingVersions,
		RegisteredModelPermissionLevelCanRead,
	}
}

// Type always returns RegisteredModelPermissionLevel to satisfy [pflag.Value] interface
func (f *RegisteredModelPermissionLevel) Type() string {
	return "RegisteredModelPermissionLevel"
}

func RegisteredModelPermissionLevelToPb(st *RegisteredModelPermissionLevel) (*mlpb.RegisteredModelPermissionLevelPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.RegisteredModelPermissionLevelPb(*st)
	return &pb, nil
}

func RegisteredModelPermissionLevelFromPb(pb *mlpb.RegisteredModelPermissionLevelPb) (*RegisteredModelPermissionLevel, error) {
	if pb == nil {
		return nil, nil
	}
	st := RegisteredModelPermissionLevel(*pb)
	return &st, nil
}

type RegisteredModelPermissions struct {

	// Wire name: 'access_control_list'
	AccessControlList []RegisteredModelAccessControlResponse ``

	// Wire name: 'object_id'
	ObjectId string ``

	// Wire name: 'object_type'
	ObjectType      string   ``
	ForceSendFields []string `tf:"-"`
}

func (st RegisteredModelPermissions) MarshalJSON() ([]byte, error) {
	pb, err := RegisteredModelPermissionsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RegisteredModelPermissions) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RegisteredModelPermissionsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RegisteredModelPermissionsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RegisteredModelPermissionsToPb(st *RegisteredModelPermissions) (*mlpb.RegisteredModelPermissionsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RegisteredModelPermissionsPb{}

	var accessControlListPb []mlpb.RegisteredModelAccessControlResponsePb
	for _, item := range st.AccessControlList {
		itemPb, err := RegisteredModelAccessControlResponseToPb(&item)
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

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RegisteredModelPermissionsFromPb(pb *mlpb.RegisteredModelPermissionsPb) (*RegisteredModelPermissions, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermissions{}

	var accessControlListField []RegisteredModelAccessControlResponse
	for _, itemPb := range pb.AccessControlList {
		item, err := RegisteredModelAccessControlResponseFromPb(&itemPb)
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

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RegisteredModelPermissionsDescription struct {

	// Wire name: 'description'
	Description string ``

	// Wire name: 'permission_level'
	PermissionLevel RegisteredModelPermissionLevel ``
	ForceSendFields []string                       `tf:"-"`
}

func (st RegisteredModelPermissionsDescription) MarshalJSON() ([]byte, error) {
	pb, err := RegisteredModelPermissionsDescriptionToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RegisteredModelPermissionsDescription) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RegisteredModelPermissionsDescriptionPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RegisteredModelPermissionsDescriptionFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RegisteredModelPermissionsDescriptionToPb(st *RegisteredModelPermissionsDescription) (*mlpb.RegisteredModelPermissionsDescriptionPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RegisteredModelPermissionsDescriptionPb{}
	pb.Description = st.Description
	permissionLevelPb, err := RegisteredModelPermissionLevelToPb(&st.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelPb != nil {
		pb.PermissionLevel = *permissionLevelPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RegisteredModelPermissionsDescriptionFromPb(pb *mlpb.RegisteredModelPermissionsDescriptionPb) (*RegisteredModelPermissionsDescription, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermissionsDescription{}
	st.Description = pb.Description
	permissionLevelField, err := RegisteredModelPermissionLevelFromPb(&pb.PermissionLevel)
	if err != nil {
		return nil, err
	}
	if permissionLevelField != nil {
		st.PermissionLevel = *permissionLevelField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RegisteredModelPermissionsRequest struct {

	// Wire name: 'access_control_list'
	AccessControlList []RegisteredModelAccessControlRequest ``
	// The registered model for which to get or manage permissions.
	// Wire name: 'registered_model_id'
	RegisteredModelId string `tf:"-"`
}

func (st RegisteredModelPermissionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := RegisteredModelPermissionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RegisteredModelPermissionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RegisteredModelPermissionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RegisteredModelPermissionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RegisteredModelPermissionsRequestToPb(st *RegisteredModelPermissionsRequest) (*mlpb.RegisteredModelPermissionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RegisteredModelPermissionsRequestPb{}

	var accessControlListPb []mlpb.RegisteredModelAccessControlRequestPb
	for _, item := range st.AccessControlList {
		itemPb, err := RegisteredModelAccessControlRequestToPb(&item)
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

func RegisteredModelPermissionsRequestFromPb(pb *mlpb.RegisteredModelPermissionsRequestPb) (*RegisteredModelPermissionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegisteredModelPermissionsRequest{}

	var accessControlListField []RegisteredModelAccessControlRequest
	for _, itemPb := range pb.AccessControlList {
		item, err := RegisteredModelAccessControlRequestFromPb(&itemPb)
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

// .. note:: Experimental: This entity may change or be removed in a future
// release without warning. Email subscription types for registry notifications:
// - `ALL_EVENTS`: Subscribed to all events. - `DEFAULT`: Default subscription
// type. - `SUBSCRIBED`: Subscribed to notifications. - `UNSUBSCRIBED`: Not
// subscribed to notifications.
type RegistryEmailSubscriptionType string

// Subscribed to all events.
const RegistryEmailSubscriptionTypeAllEvents RegistryEmailSubscriptionType = `ALL_EVENTS`

// Default subscription type.
const RegistryEmailSubscriptionTypeDefault RegistryEmailSubscriptionType = `DEFAULT`

// Subscribed to notifications.
const RegistryEmailSubscriptionTypeSubscribed RegistryEmailSubscriptionType = `SUBSCRIBED`

// Not subscribed to notifications.
const RegistryEmailSubscriptionTypeUnsubscribed RegistryEmailSubscriptionType = `UNSUBSCRIBED`

// String representation for [fmt.Print]
func (f *RegistryEmailSubscriptionType) String() string {
	return string(*f)
}

// Set raw string value and validate it against allowed values
func (f *RegistryEmailSubscriptionType) Set(v string) error {
	switch v {
	case `ALL_EVENTS`, `DEFAULT`, `SUBSCRIBED`, `UNSUBSCRIBED`:
		*f = RegistryEmailSubscriptionType(v)
		return nil
	default:
		return fmt.Errorf(`value "%s" is not one of "ALL_EVENTS", "DEFAULT", "SUBSCRIBED", "UNSUBSCRIBED"`, v)
	}
}

// Values returns all possible values for RegistryEmailSubscriptionType.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegistryEmailSubscriptionType) Values() []RegistryEmailSubscriptionType {
	return []RegistryEmailSubscriptionType{
		RegistryEmailSubscriptionTypeAllEvents,
		RegistryEmailSubscriptionTypeDefault,
		RegistryEmailSubscriptionTypeSubscribed,
		RegistryEmailSubscriptionTypeUnsubscribed,
	}
}

// Type always returns RegistryEmailSubscriptionType to satisfy [pflag.Value] interface
func (f *RegistryEmailSubscriptionType) Type() string {
	return "RegistryEmailSubscriptionType"
}

func RegistryEmailSubscriptionTypeToPb(st *RegistryEmailSubscriptionType) (*mlpb.RegistryEmailSubscriptionTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.RegistryEmailSubscriptionTypePb(*st)
	return &pb, nil
}

func RegistryEmailSubscriptionTypeFromPb(pb *mlpb.RegistryEmailSubscriptionTypePb) (*RegistryEmailSubscriptionType, error) {
	if pb == nil {
		return nil, nil
	}
	st := RegistryEmailSubscriptionType(*pb)
	return &st, nil
}

type RegistryWebhook struct {
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``
	// User-specified description for the webhook.
	// Wire name: 'description'
	Description string ``
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
	Events []RegistryWebhookEvent ``

	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpecWithoutSecret ``
	// Webhook ID
	// Wire name: 'id'
	Id string ``

	// Wire name: 'job_spec'
	JobSpec *JobSpecWithoutSecret ``
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	// Wire name: 'last_updated_timestamp'
	LastUpdatedTimestamp int64 ``
	// Name of the model whose events would trigger this webhook.
	// Wire name: 'model_name'
	ModelName string ``

	// Wire name: 'status'
	Status          RegistryWebhookStatus ``
	ForceSendFields []string              `tf:"-"`
}

func (st RegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := RegistryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RegistryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RegistryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RegistryWebhookToPb(st *RegistryWebhook) (*mlpb.RegistryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RegistryWebhookPb{}
	pb.CreationTimestamp = st.CreationTimestamp
	pb.Description = st.Description

	var eventsPb []mlpb.RegistryWebhookEventPb
	for _, item := range st.Events {
		itemPb, err := RegistryWebhookEventToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			eventsPb = append(eventsPb, *itemPb)
		}
	}
	pb.Events = eventsPb
	httpUrlSpecPb, err := HttpUrlSpecWithoutSecretToPb(st.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecPb != nil {
		pb.HttpUrlSpec = httpUrlSpecPb
	}
	pb.Id = st.Id
	jobSpecPb, err := JobSpecWithoutSecretToPb(st.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecPb != nil {
		pb.JobSpec = jobSpecPb
	}
	pb.LastUpdatedTimestamp = st.LastUpdatedTimestamp
	pb.ModelName = st.ModelName
	statusPb, err := RegistryWebhookStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RegistryWebhookFromPb(pb *mlpb.RegistryWebhookPb) (*RegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RegistryWebhook{}
	st.CreationTimestamp = pb.CreationTimestamp
	st.Description = pb.Description

	var eventsField []RegistryWebhookEvent
	for _, itemPb := range pb.Events {
		item, err := RegistryWebhookEventFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			eventsField = append(eventsField, *item)
		}
	}
	st.Events = eventsField
	httpUrlSpecField, err := HttpUrlSpecWithoutSecretFromPb(pb.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecField != nil {
		st.HttpUrlSpec = httpUrlSpecField
	}
	st.Id = pb.Id
	jobSpecField, err := JobSpecWithoutSecretFromPb(pb.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecField != nil {
		st.JobSpec = jobSpecField
	}
	st.LastUpdatedTimestamp = pb.LastUpdatedTimestamp
	st.ModelName = pb.ModelName
	statusField, err := RegistryWebhookStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RegistryWebhookEvent string

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

// Values returns all possible values for RegistryWebhookEvent.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegistryWebhookEvent) Values() []RegistryWebhookEvent {
	return []RegistryWebhookEvent{
		RegistryWebhookEventCommentCreated,
		RegistryWebhookEventModelVersionCreated,
		RegistryWebhookEventModelVersionTagSet,
		RegistryWebhookEventModelVersionTransitionedStage,
		RegistryWebhookEventModelVersionTransitionedToArchived,
		RegistryWebhookEventModelVersionTransitionedToProduction,
		RegistryWebhookEventModelVersionTransitionedToStaging,
		RegistryWebhookEventRegisteredModelCreated,
		RegistryWebhookEventTransitionRequestCreated,
		RegistryWebhookEventTransitionRequestToArchivedCreated,
		RegistryWebhookEventTransitionRequestToProductionCreated,
		RegistryWebhookEventTransitionRequestToStagingCreated,
	}
}

// Type always returns RegistryWebhookEvent to satisfy [pflag.Value] interface
func (f *RegistryWebhookEvent) Type() string {
	return "RegistryWebhookEvent"
}

func RegistryWebhookEventToPb(st *RegistryWebhookEvent) (*mlpb.RegistryWebhookEventPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.RegistryWebhookEventPb(*st)
	return &pb, nil
}

func RegistryWebhookEventFromPb(pb *mlpb.RegistryWebhookEventPb) (*RegistryWebhookEvent, error) {
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

// Values returns all possible values for RegistryWebhookStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *RegistryWebhookStatus) Values() []RegistryWebhookStatus {
	return []RegistryWebhookStatus{
		RegistryWebhookStatusActive,
		RegistryWebhookStatusDisabled,
		RegistryWebhookStatusTestMode,
	}
}

// Type always returns RegistryWebhookStatus to satisfy [pflag.Value] interface
func (f *RegistryWebhookStatus) Type() string {
	return "RegistryWebhookStatus"
}

func RegistryWebhookStatusToPb(st *RegistryWebhookStatus) (*mlpb.RegistryWebhookStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.RegistryWebhookStatusPb(*st)
	return &pb, nil
}

func RegistryWebhookStatusFromPb(pb *mlpb.RegistryWebhookStatusPb) (*RegistryWebhookStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := RegistryWebhookStatus(*pb)
	return &st, nil
}

// Details required to identify and reject a model version stage transition
// request.
type RejectTransitionRequest struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string ``
	// Name of the model.
	// Wire name: 'name'
	Name string ``
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
	Stage string ``
	// Version of the model.
	// Wire name: 'version'
	Version         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st RejectTransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := RejectTransitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RejectTransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RejectTransitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RejectTransitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RejectTransitionRequestToPb(st *RejectTransitionRequest) (*mlpb.RejectTransitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RejectTransitionRequestPb{}
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Stage = st.Stage
	pb.Version = st.Version

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RejectTransitionRequestFromPb(pb *mlpb.RejectTransitionRequestPb) (*RejectTransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RejectTransitionRequest{}
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RejectTransitionRequestResponse struct {
	// New activity generated as a result of this operation.
	// Wire name: 'activity'
	Activity *Activity ``
}

func (st RejectTransitionRequestResponse) MarshalJSON() ([]byte, error) {
	pb, err := RejectTransitionRequestResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RejectTransitionRequestResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RejectTransitionRequestResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RejectTransitionRequestResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RejectTransitionRequestResponseToPb(st *RejectTransitionRequestResponse) (*mlpb.RejectTransitionRequestResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RejectTransitionRequestResponsePb{}
	activityPb, err := ActivityToPb(st.Activity)
	if err != nil {
		return nil, err
	}
	if activityPb != nil {
		pb.Activity = activityPb
	}

	return pb, nil
}

func RejectTransitionRequestResponseFromPb(pb *mlpb.RejectTransitionRequestResponsePb) (*RejectTransitionRequestResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RejectTransitionRequestResponse{}
	activityField, err := ActivityFromPb(pb.Activity)
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
	Name string ``
	// If provided, updates the name for this `registered_model`.
	// Wire name: 'new_name'
	NewName         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st RenameModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := RenameModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RenameModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RenameModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RenameModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RenameModelRequestToPb(st *RenameModelRequest) (*mlpb.RenameModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RenameModelRequestPb{}
	pb.Name = st.Name
	pb.NewName = st.NewName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RenameModelRequestFromPb(pb *mlpb.RenameModelRequestPb) (*RenameModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RenameModelRequest{}
	st.Name = pb.Name
	st.NewName = pb.NewName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RenameModelResponse struct {

	// Wire name: 'registered_model'
	RegisteredModel *Model ``
}

func (st RenameModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := RenameModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RenameModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RenameModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RenameModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RenameModelResponseToPb(st *RenameModelResponse) (*mlpb.RenameModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RenameModelResponsePb{}
	registeredModelPb, err := ModelToPb(st.RegisteredModel)
	if err != nil {
		return nil, err
	}
	if registeredModelPb != nil {
		pb.RegisteredModel = registeredModelPb
	}

	return pb, nil
}

func RenameModelResponseFromPb(pb *mlpb.RenameModelResponsePb) (*RenameModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RenameModelResponse{}
	registeredModelField, err := ModelFromPb(pb.RegisteredModel)
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
	ExperimentId string ``
}

func (st RestoreExperiment) MarshalJSON() ([]byte, error) {
	pb, err := RestoreExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RestoreExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RestoreExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RestoreExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RestoreExperimentToPb(st *RestoreExperiment) (*mlpb.RestoreExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RestoreExperimentPb{}
	pb.ExperimentId = st.ExperimentId

	return pb, nil
}

func RestoreExperimentFromPb(pb *mlpb.RestoreExperimentPb) (*RestoreExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreExperiment{}
	st.ExperimentId = pb.ExperimentId

	return st, nil
}

type RestoreRun struct {
	// ID of the run to restore.
	// Wire name: 'run_id'
	RunId string ``
}

func (st RestoreRun) MarshalJSON() ([]byte, error) {
	pb, err := RestoreRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RestoreRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RestoreRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RestoreRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RestoreRunToPb(st *RestoreRun) (*mlpb.RestoreRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RestoreRunPb{}
	pb.RunId = st.RunId

	return pb, nil
}

func RestoreRunFromPb(pb *mlpb.RestoreRunPb) (*RestoreRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRun{}
	st.RunId = pb.RunId

	return st, nil
}

type RestoreRuns struct {
	// The ID of the experiment containing the runs to restore.
	// Wire name: 'experiment_id'
	ExperimentId string ``
	// An optional positive integer indicating the maximum number of runs to
	// restore. The maximum allowed value for max_runs is 10000.
	// Wire name: 'max_runs'
	MaxRuns int ``
	// The minimum deletion timestamp in milliseconds since the UNIX epoch for
	// restoring runs. Only runs deleted no earlier than this timestamp are
	// restored.
	// Wire name: 'min_timestamp_millis'
	MinTimestampMillis int64    ``
	ForceSendFields    []string `tf:"-"`
}

func (st RestoreRuns) MarshalJSON() ([]byte, error) {
	pb, err := RestoreRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RestoreRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RestoreRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RestoreRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RestoreRunsToPb(st *RestoreRuns) (*mlpb.RestoreRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RestoreRunsPb{}
	pb.ExperimentId = st.ExperimentId
	pb.MaxRuns = st.MaxRuns
	pb.MinTimestampMillis = st.MinTimestampMillis

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RestoreRunsFromPb(pb *mlpb.RestoreRunsPb) (*RestoreRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRuns{}
	st.ExperimentId = pb.ExperimentId
	st.MaxRuns = pb.MaxRuns
	st.MinTimestampMillis = pb.MinTimestampMillis

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type RestoreRunsResponse struct {
	// The number of runs restored.
	// Wire name: 'runs_restored'
	RunsRestored    int      ``
	ForceSendFields []string `tf:"-"`
}

func (st RestoreRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := RestoreRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RestoreRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RestoreRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RestoreRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RestoreRunsResponseToPb(st *RestoreRunsResponse) (*mlpb.RestoreRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RestoreRunsResponsePb{}
	pb.RunsRestored = st.RunsRestored

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RestoreRunsResponseFromPb(pb *mlpb.RestoreRunsResponsePb) (*RestoreRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RestoreRunsResponse{}
	st.RunsRestored = pb.RunsRestored

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// A single run.
type Run struct {
	// Run data.
	// Wire name: 'data'
	Data *RunData ``
	// Run metadata.
	// Wire name: 'info'
	Info *RunInfo ``
	// Run inputs.
	// Wire name: 'inputs'
	Inputs *RunInputs ``
}

func (st Run) MarshalJSON() ([]byte, error) {
	pb, err := RunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *Run) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunToPb(st *Run) (*mlpb.RunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RunPb{}
	dataPb, err := RunDataToPb(st.Data)
	if err != nil {
		return nil, err
	}
	if dataPb != nil {
		pb.Data = dataPb
	}
	infoPb, err := RunInfoToPb(st.Info)
	if err != nil {
		return nil, err
	}
	if infoPb != nil {
		pb.Info = infoPb
	}
	inputsPb, err := RunInputsToPb(st.Inputs)
	if err != nil {
		return nil, err
	}
	if inputsPb != nil {
		pb.Inputs = inputsPb
	}

	return pb, nil
}

func RunFromPb(pb *mlpb.RunPb) (*Run, error) {
	if pb == nil {
		return nil, nil
	}
	st := &Run{}
	dataField, err := RunDataFromPb(pb.Data)
	if err != nil {
		return nil, err
	}
	if dataField != nil {
		st.Data = dataField
	}
	infoField, err := RunInfoFromPb(pb.Info)
	if err != nil {
		return nil, err
	}
	if infoField != nil {
		st.Info = infoField
	}
	inputsField, err := RunInputsFromPb(pb.Inputs)
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
	Metrics []Metric ``
	// Run parameters.
	// Wire name: 'params'
	Params []Param ``
	// Additional metadata key-value pairs.
	// Wire name: 'tags'
	Tags []RunTag ``
}

func (st RunData) MarshalJSON() ([]byte, error) {
	pb, err := RunDataToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunData) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RunDataPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunDataFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunDataToPb(st *RunData) (*mlpb.RunDataPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RunDataPb{}

	var metricsPb []mlpb.MetricPb
	for _, item := range st.Metrics {
		itemPb, err := MetricToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			metricsPb = append(metricsPb, *itemPb)
		}
	}
	pb.Metrics = metricsPb

	var paramsPb []mlpb.ParamPb
	for _, item := range st.Params {
		itemPb, err := ParamToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			paramsPb = append(paramsPb, *itemPb)
		}
	}
	pb.Params = paramsPb

	var tagsPb []mlpb.RunTagPb
	for _, item := range st.Tags {
		itemPb, err := RunTagToPb(&item)
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

func RunDataFromPb(pb *mlpb.RunDataPb) (*RunData, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunData{}

	var metricsField []Metric
	for _, itemPb := range pb.Metrics {
		item, err := MetricFromPb(&itemPb)
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
		item, err := ParamFromPb(&itemPb)
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
		item, err := RunTagFromPb(&itemPb)
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
	ArtifactUri string ``
	// Unix timestamp of when the run ended in milliseconds.
	// Wire name: 'end_time'
	EndTime int64 ``
	// The experiment ID.
	// Wire name: 'experiment_id'
	ExperimentId string ``
	// Current life cycle stage of the experiment : OneOf("active", "deleted")
	// Wire name: 'lifecycle_stage'
	LifecycleStage string ``
	// Unique identifier for the run.
	// Wire name: 'run_id'
	RunId string ``
	// The name of the run.
	// Wire name: 'run_name'
	RunName string ``
	// [Deprecated, use run_id instead] Unique identifier for the run. This
	// field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string ``
	// Unix timestamp of when the run started in milliseconds.
	// Wire name: 'start_time'
	StartTime int64 ``
	// Current status of the run.
	// Wire name: 'status'
	Status RunInfoStatus ``
	// User who initiated the run. This field is deprecated as of MLflow 1.0,
	// and will be removed in a future MLflow release. Use 'mlflow.user' tag
	// instead.
	// Wire name: 'user_id'
	UserId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st RunInfo) MarshalJSON() ([]byte, error) {
	pb, err := RunInfoToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunInfo) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RunInfoPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunInfoFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunInfoToPb(st *RunInfo) (*mlpb.RunInfoPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RunInfoPb{}
	pb.ArtifactUri = st.ArtifactUri
	pb.EndTime = st.EndTime
	pb.ExperimentId = st.ExperimentId
	pb.LifecycleStage = st.LifecycleStage
	pb.RunId = st.RunId
	pb.RunName = st.RunName
	pb.RunUuid = st.RunUuid
	pb.StartTime = st.StartTime
	statusPb, err := RunInfoStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}
	pb.UserId = st.UserId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunInfoFromPb(pb *mlpb.RunInfoPb) (*RunInfo, error) {
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
	statusField, err := RunInfoStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}
	st.UserId = pb.UserId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Status of a run.
type RunInfoStatus string

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

// Values returns all possible values for RunInfoStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *RunInfoStatus) Values() []RunInfoStatus {
	return []RunInfoStatus{
		RunInfoStatusFailed,
		RunInfoStatusFinished,
		RunInfoStatusKilled,
		RunInfoStatusRunning,
		RunInfoStatusScheduled,
	}
}

// Type always returns RunInfoStatus to satisfy [pflag.Value] interface
func (f *RunInfoStatus) Type() string {
	return "RunInfoStatus"
}

func RunInfoStatusToPb(st *RunInfoStatus) (*mlpb.RunInfoStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.RunInfoStatusPb(*st)
	return &pb, nil
}

func RunInfoStatusFromPb(pb *mlpb.RunInfoStatusPb) (*RunInfoStatus, error) {
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
	DatasetInputs []DatasetInput ``
	// Model inputs to the Run.
	// Wire name: 'model_inputs'
	ModelInputs []ModelInput ``
}

func (st RunInputs) MarshalJSON() ([]byte, error) {
	pb, err := RunInputsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunInputs) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RunInputsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunInputsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunInputsToPb(st *RunInputs) (*mlpb.RunInputsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RunInputsPb{}

	var datasetInputsPb []mlpb.DatasetInputPb
	for _, item := range st.DatasetInputs {
		itemPb, err := DatasetInputToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			datasetInputsPb = append(datasetInputsPb, *itemPb)
		}
	}
	pb.DatasetInputs = datasetInputsPb

	var modelInputsPb []mlpb.ModelInputPb
	for _, item := range st.ModelInputs {
		itemPb, err := ModelInputToPb(&item)
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

func RunInputsFromPb(pb *mlpb.RunInputsPb) (*RunInputs, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunInputs{}

	var datasetInputsField []DatasetInput
	for _, itemPb := range pb.DatasetInputs {
		item, err := DatasetInputFromPb(&itemPb)
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
		item, err := ModelInputFromPb(&itemPb)
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
	Key string ``
	// The tag value.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st RunTag) MarshalJSON() ([]byte, error) {
	pb, err := RunTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *RunTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.RunTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := RunTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func RunTagToPb(st *RunTag) (*mlpb.RunTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.RunTagPb{}
	pb.Key = st.Key
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func RunTagFromPb(pb *mlpb.RunTagPb) (*RunTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &RunTag{}
	st.Key = pb.Key
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchExperiments struct {
	// String representing a SQL filter condition (e.g. "name ILIKE
	// 'my-experiment%'")
	// Wire name: 'filter'
	Filter string ``
	// Maximum number of experiments desired. Max threshold is 3000.
	// Wire name: 'max_results'
	MaxResults int64 ``
	// List of columns for ordering search results, which can include experiment
	// name and last updated timestamp with an optional "DESC" or "ASC"
	// annotation, where "ASC" is the default. Tiebreaks are done by experiment
	// id DESC.
	// Wire name: 'order_by'
	OrderBy []string ``
	// Token indicating the page of experiments to fetch
	// Wire name: 'page_token'
	PageToken string ``
	// Qualifier for type of experiments to be returned. If unspecified, return
	// only active experiments.
	// Wire name: 'view_type'
	ViewType        ViewType ``
	ForceSendFields []string `tf:"-"`
}

func (st SearchExperiments) MarshalJSON() ([]byte, error) {
	pb, err := SearchExperimentsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchExperiments) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchExperimentsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchExperimentsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchExperimentsToPb(st *SearchExperiments) (*mlpb.SearchExperimentsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchExperimentsPb{}
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken
	viewTypePb, err := ViewTypeToPb(&st.ViewType)
	if err != nil {
		return nil, err
	}
	if viewTypePb != nil {
		pb.ViewType = *viewTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchExperimentsFromPb(pb *mlpb.SearchExperimentsPb) (*SearchExperiments, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchExperiments{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken
	viewTypeField, err := ViewTypeFromPb(&pb.ViewType)
	if err != nil {
		return nil, err
	}
	if viewTypeField != nil {
		st.ViewType = *viewTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchExperimentsResponse struct {
	// Experiments that match the search criteria
	// Wire name: 'experiments'
	Experiments []Experiment ``
	// Token that can be used to retrieve the next page of experiments. An empty
	// token means that no more experiments are available for retrieval.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SearchExperimentsResponse) MarshalJSON() ([]byte, error) {
	pb, err := SearchExperimentsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchExperimentsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchExperimentsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchExperimentsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchExperimentsResponseToPb(st *SearchExperimentsResponse) (*mlpb.SearchExperimentsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchExperimentsResponsePb{}

	var experimentsPb []mlpb.ExperimentPb
	for _, item := range st.Experiments {
		itemPb, err := ExperimentToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			experimentsPb = append(experimentsPb, *itemPb)
		}
	}
	pb.Experiments = experimentsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchExperimentsResponseFromPb(pb *mlpb.SearchExperimentsResponsePb) (*SearchExperimentsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchExperimentsResponse{}

	var experimentsField []Experiment
	for _, itemPb := range pb.Experiments {
		item, err := ExperimentFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			experimentsField = append(experimentsField, *item)
		}
	}
	st.Experiments = experimentsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchLoggedModelsDataset struct {
	// The digest of the dataset.
	// Wire name: 'dataset_digest'
	DatasetDigest string ``
	// The name of the dataset.
	// Wire name: 'dataset_name'
	DatasetName     string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SearchLoggedModelsDataset) MarshalJSON() ([]byte, error) {
	pb, err := SearchLoggedModelsDatasetToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchLoggedModelsDataset) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchLoggedModelsDatasetPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchLoggedModelsDatasetFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchLoggedModelsDatasetToPb(st *SearchLoggedModelsDataset) (*mlpb.SearchLoggedModelsDatasetPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchLoggedModelsDatasetPb{}
	pb.DatasetDigest = st.DatasetDigest
	pb.DatasetName = st.DatasetName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchLoggedModelsDatasetFromPb(pb *mlpb.SearchLoggedModelsDatasetPb) (*SearchLoggedModelsDataset, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchLoggedModelsDataset{}
	st.DatasetDigest = pb.DatasetDigest
	st.DatasetName = pb.DatasetName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchLoggedModelsOrderBy struct {
	// Whether the search results order is ascending or not.
	// Wire name: 'ascending'
	Ascending bool ``
	// If ``field_name`` refers to a metric, this field specifies the digest of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name and digest will be considered for ordering. This
	// field may only be set if ``dataset_name`` is also set.
	// Wire name: 'dataset_digest'
	DatasetDigest string ``
	// If ``field_name`` refers to a metric, this field specifies the name of
	// the dataset associated with the metric. Only metrics associated with the
	// specified dataset name will be considered for ordering. This field may
	// only be set if ``field_name`` refers to a metric.
	// Wire name: 'dataset_name'
	DatasetName string ``
	// The name of the field to order by, e.g. "metrics.accuracy".
	// Wire name: 'field_name'
	FieldName       string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SearchLoggedModelsOrderBy) MarshalJSON() ([]byte, error) {
	pb, err := SearchLoggedModelsOrderByToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchLoggedModelsOrderBy) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchLoggedModelsOrderByPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchLoggedModelsOrderByFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchLoggedModelsOrderByToPb(st *SearchLoggedModelsOrderBy) (*mlpb.SearchLoggedModelsOrderByPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchLoggedModelsOrderByPb{}
	pb.Ascending = st.Ascending
	pb.DatasetDigest = st.DatasetDigest
	pb.DatasetName = st.DatasetName
	pb.FieldName = st.FieldName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchLoggedModelsOrderByFromPb(pb *mlpb.SearchLoggedModelsOrderByPb) (*SearchLoggedModelsOrderBy, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchLoggedModelsOrderBy{}
	st.Ascending = pb.Ascending
	st.DatasetDigest = pb.DatasetDigest
	st.DatasetName = pb.DatasetName
	st.FieldName = pb.FieldName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchLoggedModelsRequest struct {
	// List of datasets on which to apply the metrics filter clauses. For
	// example, a filter with `metrics.accuracy > 0.9` and dataset info with
	// name "test_dataset" means we will return all logged models with accuracy
	// > 0.9 on the test_dataset. Metric values from ANY dataset matching the
	// criteria are considered. If no datasets are specified, then metrics
	// across all datasets are considered in the filter.
	// Wire name: 'datasets'
	Datasets []SearchLoggedModelsDataset ``
	// The IDs of the experiments in which to search for logged models.
	// Wire name: 'experiment_ids'
	ExperimentIds []string ``
	// A filter expression over logged model info and data that allows returning
	// a subset of logged models. The syntax is a subset of SQL that supports
	// AND'ing together binary operations.
	//
	// Example: ``params.alpha < 0.3 AND metrics.accuracy > 0.9``.
	// Wire name: 'filter'
	Filter string ``
	// The maximum number of Logged Models to return. The maximum limit is 50.
	// Wire name: 'max_results'
	MaxResults int ``
	// The list of columns for ordering the results, with additional fields for
	// sorting criteria.
	// Wire name: 'order_by'
	OrderBy []SearchLoggedModelsOrderBy ``
	// The token indicating the page of logged models to fetch.
	// Wire name: 'page_token'
	PageToken       string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SearchLoggedModelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := SearchLoggedModelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchLoggedModelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchLoggedModelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchLoggedModelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchLoggedModelsRequestToPb(st *SearchLoggedModelsRequest) (*mlpb.SearchLoggedModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchLoggedModelsRequestPb{}

	var datasetsPb []mlpb.SearchLoggedModelsDatasetPb
	for _, item := range st.Datasets {
		itemPb, err := SearchLoggedModelsDatasetToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			datasetsPb = append(datasetsPb, *itemPb)
		}
	}
	pb.Datasets = datasetsPb
	pb.ExperimentIds = st.ExperimentIds
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults

	var orderByPb []mlpb.SearchLoggedModelsOrderByPb
	for _, item := range st.OrderBy {
		itemPb, err := SearchLoggedModelsOrderByToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			orderByPb = append(orderByPb, *itemPb)
		}
	}
	pb.OrderBy = orderByPb
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchLoggedModelsRequestFromPb(pb *mlpb.SearchLoggedModelsRequestPb) (*SearchLoggedModelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchLoggedModelsRequest{}

	var datasetsField []SearchLoggedModelsDataset
	for _, itemPb := range pb.Datasets {
		item, err := SearchLoggedModelsDatasetFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			datasetsField = append(datasetsField, *item)
		}
	}
	st.Datasets = datasetsField
	st.ExperimentIds = pb.ExperimentIds
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults

	var orderByField []SearchLoggedModelsOrderBy
	for _, itemPb := range pb.OrderBy {
		item, err := SearchLoggedModelsOrderByFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			orderByField = append(orderByField, *item)
		}
	}
	st.OrderBy = orderByField
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchLoggedModelsResponse struct {
	// Logged models that match the search criteria.
	// Wire name: 'models'
	Models []LoggedModel ``
	// The token that can be used to retrieve the next page of logged models.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SearchLoggedModelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := SearchLoggedModelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchLoggedModelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchLoggedModelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchLoggedModelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchLoggedModelsResponseToPb(st *SearchLoggedModelsResponse) (*mlpb.SearchLoggedModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchLoggedModelsResponsePb{}

	var modelsPb []mlpb.LoggedModelPb
	for _, item := range st.Models {
		itemPb, err := LoggedModelToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			modelsPb = append(modelsPb, *itemPb)
		}
	}
	pb.Models = modelsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchLoggedModelsResponseFromPb(pb *mlpb.SearchLoggedModelsResponsePb) (*SearchLoggedModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchLoggedModelsResponse{}

	var modelsField []LoggedModel
	for _, itemPb := range pb.Models {
		item, err := LoggedModelFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			modelsField = append(modelsField, *item)
		}
	}
	st.Models = modelsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchModelVersionsRequest struct {
	// String filter condition, like "name='my-model-name'". Must be a single
	// boolean condition, with string values wrapped in single quotes.
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Maximum number of models desired. Max threshold is 10K.
	// Wire name: 'max_results'
	MaxResults int64 `tf:"-"`
	// List of columns to be ordered by including model name, version, stage
	// with an optional "DESC" or "ASC" annotation, where "ASC" is the default.
	// Tiebreaks are done by latest stage transition timestamp, followed by name
	// ASC, followed by version DESC.
	// Wire name: 'order_by'
	OrderBy []string `tf:"-"`
	// Pagination token to go to next page based on previous search query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st SearchModelVersionsRequest) MarshalJSON() ([]byte, error) {
	pb, err := SearchModelVersionsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchModelVersionsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchModelVersionsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchModelVersionsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchModelVersionsRequestToPb(st *SearchModelVersionsRequest) (*mlpb.SearchModelVersionsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchModelVersionsRequestPb{}
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchModelVersionsRequestFromPb(pb *mlpb.SearchModelVersionsRequestPb) (*SearchModelVersionsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelVersionsRequest{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchModelVersionsResponse struct {
	// Models that match the search criteria
	// Wire name: 'model_versions'
	ModelVersions []ModelVersion ``
	// Pagination token to request next page of models for the same search
	// query.
	// Wire name: 'next_page_token'
	NextPageToken   string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SearchModelVersionsResponse) MarshalJSON() ([]byte, error) {
	pb, err := SearchModelVersionsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchModelVersionsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchModelVersionsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchModelVersionsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchModelVersionsResponseToPb(st *SearchModelVersionsResponse) (*mlpb.SearchModelVersionsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchModelVersionsResponsePb{}

	var modelVersionsPb []mlpb.ModelVersionPb
	for _, item := range st.ModelVersions {
		itemPb, err := ModelVersionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			modelVersionsPb = append(modelVersionsPb, *itemPb)
		}
	}
	pb.ModelVersions = modelVersionsPb
	pb.NextPageToken = st.NextPageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchModelVersionsResponseFromPb(pb *mlpb.SearchModelVersionsResponsePb) (*SearchModelVersionsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelVersionsResponse{}

	var modelVersionsField []ModelVersion
	for _, itemPb := range pb.ModelVersions {
		item, err := ModelVersionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			modelVersionsField = append(modelVersionsField, *item)
		}
	}
	st.ModelVersions = modelVersionsField
	st.NextPageToken = pb.NextPageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchModelsRequest struct {
	// String filter condition, like "name LIKE 'my-model-name'". Interpreted in
	// the backend automatically as "name LIKE '%my-model-name%'". Single
	// boolean condition, with string values wrapped in single quotes.
	// Wire name: 'filter'
	Filter string `tf:"-"`
	// Maximum number of models desired. Default is 100. Max threshold is 1000.
	// Wire name: 'max_results'
	MaxResults int64 `tf:"-"`
	// List of columns for ordering search results, which can include model name
	// and last updated timestamp with an optional "DESC" or "ASC" annotation,
	// where "ASC" is the default. Tiebreaks are done by model name ASC.
	// Wire name: 'order_by'
	OrderBy []string `tf:"-"`
	// Pagination token to go to the next page based on a previous search query.
	// Wire name: 'page_token'
	PageToken       string   `tf:"-"`
	ForceSendFields []string `tf:"-"`
}

func (st SearchModelsRequest) MarshalJSON() ([]byte, error) {
	pb, err := SearchModelsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchModelsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchModelsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchModelsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchModelsRequestToPb(st *SearchModelsRequest) (*mlpb.SearchModelsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchModelsRequestPb{}
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchModelsRequestFromPb(pb *mlpb.SearchModelsRequestPb) (*SearchModelsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelsRequest{}
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchModelsResponse struct {
	// Pagination token to request the next page of models.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// Registered Models that match the search criteria.
	// Wire name: 'registered_models'
	RegisteredModels []Model  ``
	ForceSendFields  []string `tf:"-"`
}

func (st SearchModelsResponse) MarshalJSON() ([]byte, error) {
	pb, err := SearchModelsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchModelsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchModelsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchModelsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchModelsResponseToPb(st *SearchModelsResponse) (*mlpb.SearchModelsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchModelsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var registeredModelsPb []mlpb.ModelPb
	for _, item := range st.RegisteredModels {
		itemPb, err := ModelToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			registeredModelsPb = append(registeredModelsPb, *itemPb)
		}
	}
	pb.RegisteredModels = registeredModelsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchModelsResponseFromPb(pb *mlpb.SearchModelsResponsePb) (*SearchModelsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchModelsResponse{}
	st.NextPageToken = pb.NextPageToken

	var registeredModelsField []Model
	for _, itemPb := range pb.RegisteredModels {
		item, err := ModelFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			registeredModelsField = append(registeredModelsField, *item)
		}
	}
	st.RegisteredModels = registeredModelsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchRuns struct {
	// List of experiment IDs to search over.
	// Wire name: 'experiment_ids'
	ExperimentIds []string ``
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
	Filter string ``
	// Maximum number of runs desired. Max threshold is 50000
	// Wire name: 'max_results'
	MaxResults int ``
	// List of columns to be ordered by, including attributes, params, metrics,
	// and tags with an optional `"DESC"` or `"ASC"` annotation, where `"ASC"`
	// is the default. Example: `["params.input DESC", "metrics.alpha ASC",
	// "metrics.rmse"]`. Tiebreaks are done by start_time `DESC` followed by
	// `run_id` for runs with the same start time (and this is the default
	// ordering criterion if order_by is not provided).
	// Wire name: 'order_by'
	OrderBy []string ``
	// Token for the current page of runs.
	// Wire name: 'page_token'
	PageToken string ``
	// Whether to display only active, only deleted, or all runs. Defaults to
	// only active runs.
	// Wire name: 'run_view_type'
	RunViewType     ViewType ``
	ForceSendFields []string `tf:"-"`
}

func (st SearchRuns) MarshalJSON() ([]byte, error) {
	pb, err := SearchRunsToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchRuns) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchRunsPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchRunsFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchRunsToPb(st *SearchRuns) (*mlpb.SearchRunsPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchRunsPb{}
	pb.ExperimentIds = st.ExperimentIds
	pb.Filter = st.Filter
	pb.MaxResults = st.MaxResults
	pb.OrderBy = st.OrderBy
	pb.PageToken = st.PageToken
	runViewTypePb, err := ViewTypeToPb(&st.RunViewType)
	if err != nil {
		return nil, err
	}
	if runViewTypePb != nil {
		pb.RunViewType = *runViewTypePb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchRunsFromPb(pb *mlpb.SearchRunsPb) (*SearchRuns, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchRuns{}
	st.ExperimentIds = pb.ExperimentIds
	st.Filter = pb.Filter
	st.MaxResults = pb.MaxResults
	st.OrderBy = pb.OrderBy
	st.PageToken = pb.PageToken
	runViewTypeField, err := ViewTypeFromPb(&pb.RunViewType)
	if err != nil {
		return nil, err
	}
	if runViewTypeField != nil {
		st.RunViewType = *runViewTypeField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SearchRunsResponse struct {
	// Token for the next page of runs.
	// Wire name: 'next_page_token'
	NextPageToken string ``
	// Runs that match the search criteria.
	// Wire name: 'runs'
	Runs            []Run    ``
	ForceSendFields []string `tf:"-"`
}

func (st SearchRunsResponse) MarshalJSON() ([]byte, error) {
	pb, err := SearchRunsResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SearchRunsResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SearchRunsResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SearchRunsResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SearchRunsResponseToPb(st *SearchRunsResponse) (*mlpb.SearchRunsResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SearchRunsResponsePb{}
	pb.NextPageToken = st.NextPageToken

	var runsPb []mlpb.RunPb
	for _, item := range st.Runs {
		itemPb, err := RunToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			runsPb = append(runsPb, *itemPb)
		}
	}
	pb.Runs = runsPb

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SearchRunsResponseFromPb(pb *mlpb.SearchRunsResponsePb) (*SearchRunsResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SearchRunsResponse{}
	st.NextPageToken = pb.NextPageToken

	var runsField []Run
	for _, itemPb := range pb.Runs {
		item, err := RunFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			runsField = append(runsField, *item)
		}
	}
	st.Runs = runsField

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type SetExperimentTag struct {
	// ID of the experiment under which to log the tag. Must be provided.
	// Wire name: 'experiment_id'
	ExperimentId string ``
	// Name of the tag. Keys up to 250 bytes in size are supported.
	// Wire name: 'key'
	Key string ``
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	// Wire name: 'value'
	Value string ``
}

func (st SetExperimentTag) MarshalJSON() ([]byte, error) {
	pb, err := SetExperimentTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SetExperimentTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SetExperimentTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SetExperimentTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SetExperimentTagToPb(st *SetExperimentTag) (*mlpb.SetExperimentTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SetExperimentTagPb{}
	pb.ExperimentId = st.ExperimentId
	pb.Key = st.Key
	pb.Value = st.Value

	return pb, nil
}

func SetExperimentTagFromPb(pb *mlpb.SetExperimentTagPb) (*SetExperimentTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetExperimentTag{}
	st.ExperimentId = pb.ExperimentId
	st.Key = pb.Key
	st.Value = pb.Value

	return st, nil
}

type SetLoggedModelTagsRequest struct {
	// The ID of the logged model to set the tags on.
	// Wire name: 'model_id'
	ModelId string `tf:"-"`
	// The tags to set on the logged model.
	// Wire name: 'tags'
	Tags []LoggedModelTag ``
}

func (st SetLoggedModelTagsRequest) MarshalJSON() ([]byte, error) {
	pb, err := SetLoggedModelTagsRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SetLoggedModelTagsRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SetLoggedModelTagsRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SetLoggedModelTagsRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SetLoggedModelTagsRequestToPb(st *SetLoggedModelTagsRequest) (*mlpb.SetLoggedModelTagsRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SetLoggedModelTagsRequestPb{}
	pb.ModelId = st.ModelId

	var tagsPb []mlpb.LoggedModelTagPb
	for _, item := range st.Tags {
		itemPb, err := LoggedModelTagToPb(&item)
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

func SetLoggedModelTagsRequestFromPb(pb *mlpb.SetLoggedModelTagsRequestPb) (*SetLoggedModelTagsRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetLoggedModelTagsRequest{}
	st.ModelId = pb.ModelId

	var tagsField []LoggedModelTag
	for _, itemPb := range pb.Tags {
		item, err := LoggedModelTagFromPb(&itemPb)
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

type SetModelTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	// Wire name: 'key'
	Key string ``
	// Unique name of the model.
	// Wire name: 'name'
	Name string ``
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	// Wire name: 'value'
	Value string ``
}

func (st SetModelTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := SetModelTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SetModelTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SetModelTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SetModelTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SetModelTagRequestToPb(st *SetModelTagRequest) (*mlpb.SetModelTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SetModelTagRequestPb{}
	pb.Key = st.Key
	pb.Name = st.Name
	pb.Value = st.Value

	return pb, nil
}

func SetModelTagRequestFromPb(pb *mlpb.SetModelTagRequestPb) (*SetModelTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetModelTagRequest{}
	st.Key = pb.Key
	st.Name = pb.Name
	st.Value = pb.Value

	return st, nil
}

type SetModelVersionTagRequest struct {
	// Name of the tag. Maximum size depends on storage backend. If a tag with
	// this name already exists, its preexisting value will be replaced by the
	// specified `value`. All storage backends are guaranteed to support key
	// values up to 250 bytes in size.
	// Wire name: 'key'
	Key string ``
	// Unique name of the model.
	// Wire name: 'name'
	Name string ``
	// String value of the tag being logged. Maximum size depends on storage
	// backend. All storage backends are guaranteed to support key values up to
	// 5000 bytes in size.
	// Wire name: 'value'
	Value string ``
	// Model version number.
	// Wire name: 'version'
	Version string ``
}

func (st SetModelVersionTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := SetModelVersionTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SetModelVersionTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SetModelVersionTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SetModelVersionTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SetModelVersionTagRequestToPb(st *SetModelVersionTagRequest) (*mlpb.SetModelVersionTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SetModelVersionTagRequestPb{}
	pb.Key = st.Key
	pb.Name = st.Name
	pb.Value = st.Value
	pb.Version = st.Version

	return pb, nil
}

func SetModelVersionTagRequestFromPb(pb *mlpb.SetModelVersionTagRequestPb) (*SetModelVersionTagRequest, error) {
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

type SetTag struct {
	// Name of the tag. Keys up to 250 bytes in size are supported.
	// Wire name: 'key'
	Key string ``
	// ID of the run under which to log the tag. Must be provided.
	// Wire name: 'run_id'
	RunId string ``
	// [Deprecated, use `run_id` instead] ID of the run under which to log the
	// tag. This field will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string ``
	// String value of the tag being logged. Values up to 64KB in size are
	// supported.
	// Wire name: 'value'
	Value           string   ``
	ForceSendFields []string `tf:"-"`
}

func (st SetTag) MarshalJSON() ([]byte, error) {
	pb, err := SetTagToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *SetTag) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.SetTagPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := SetTagFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func SetTagToPb(st *SetTag) (*mlpb.SetTagPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.SetTagPb{}
	pb.Key = st.Key
	pb.RunId = st.RunId
	pb.RunUuid = st.RunUuid
	pb.Value = st.Value

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func SetTagFromPb(pb *mlpb.SetTagPb) (*SetTag, error) {
	if pb == nil {
		return nil, nil
	}
	st := &SetTag{}
	st.Key = pb.Key
	st.RunId = pb.RunId
	st.RunUuid = pb.RunUuid
	st.Value = pb.Value

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// The status of the model version. Valid values are: * `PENDING_REGISTRATION`:
// Request to register a new model version is pending as server performs
// background tasks.
//
// * `FAILED_REGISTRATION`: Request to register a new model version has failed.
//
// * `READY`: Model version is ready for use.
type Status string

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

// Values returns all possible values for Status.
//
// There is no guarantee on the order of the values in the slice.
func (f *Status) Values() []Status {
	return []Status{
		StatusFailedRegistration,
		StatusPendingRegistration,
		StatusReady,
	}
}

// Type always returns Status to satisfy [pflag.Value] interface
func (f *Status) Type() string {
	return "Status"
}

func StatusToPb(st *Status) (*mlpb.StatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.StatusPb(*st)
	return &pb, nil
}

func StatusFromPb(pb *mlpb.StatusPb) (*Status, error) {
	if pb == nil {
		return nil, nil
	}
	st := Status(*pb)
	return &st, nil
}

// Details required to test a registry webhook.
type TestRegistryWebhookRequest struct {
	// If `event` is specified, the test trigger uses the specified event. If
	// `event` is not specified, the test trigger uses a randomly chosen event
	// associated with the webhook.
	// Wire name: 'event'
	Event RegistryWebhookEvent ``
	// Webhook ID
	// Wire name: 'id'
	Id string ``
}

func (st TestRegistryWebhookRequest) MarshalJSON() ([]byte, error) {
	pb, err := TestRegistryWebhookRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TestRegistryWebhookRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.TestRegistryWebhookRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TestRegistryWebhookRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TestRegistryWebhookRequestToPb(st *TestRegistryWebhookRequest) (*mlpb.TestRegistryWebhookRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.TestRegistryWebhookRequestPb{}
	eventPb, err := RegistryWebhookEventToPb(&st.Event)
	if err != nil {
		return nil, err
	}
	if eventPb != nil {
		pb.Event = *eventPb
	}
	pb.Id = st.Id

	return pb, nil
}

func TestRegistryWebhookRequestFromPb(pb *mlpb.TestRegistryWebhookRequestPb) (*TestRegistryWebhookRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TestRegistryWebhookRequest{}
	eventField, err := RegistryWebhookEventFromPb(&pb.Event)
	if err != nil {
		return nil, err
	}
	if eventField != nil {
		st.Event = *eventField
	}
	st.Id = pb.Id

	return st, nil
}

type TestRegistryWebhookResponse struct {
	// Body of the response from the webhook URL
	// Wire name: 'body'
	Body string ``
	// Status code returned by the webhook URL
	// Wire name: 'status_code'
	StatusCode      int      ``
	ForceSendFields []string `tf:"-"`
}

func (st TestRegistryWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := TestRegistryWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TestRegistryWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.TestRegistryWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TestRegistryWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TestRegistryWebhookResponseToPb(st *TestRegistryWebhookResponse) (*mlpb.TestRegistryWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.TestRegistryWebhookResponsePb{}
	pb.Body = st.Body
	pb.StatusCode = st.StatusCode

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TestRegistryWebhookResponseFromPb(pb *mlpb.TestRegistryWebhookResponsePb) (*TestRegistryWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TestRegistryWebhookResponse{}
	st.Body = pb.Body
	st.StatusCode = pb.StatusCode

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// Details required to transition a model version's stage.
type TransitionModelVersionStageDatabricks struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	// Wire name: 'archive_existing_versions'
	ArchiveExistingVersions bool ``
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string ``
	// Name of the model.
	// Wire name: 'name'
	Name string ``
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
	Stage string ``
	// Version of the model.
	// Wire name: 'version'
	Version         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st TransitionModelVersionStageDatabricks) MarshalJSON() ([]byte, error) {
	pb, err := TransitionModelVersionStageDatabricksToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TransitionModelVersionStageDatabricks) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.TransitionModelVersionStageDatabricksPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TransitionModelVersionStageDatabricksFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TransitionModelVersionStageDatabricksToPb(st *TransitionModelVersionStageDatabricks) (*mlpb.TransitionModelVersionStageDatabricksPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.TransitionModelVersionStageDatabricksPb{}
	pb.ArchiveExistingVersions = st.ArchiveExistingVersions
	pb.Comment = st.Comment
	pb.Name = st.Name
	pb.Stage = st.Stage
	pb.Version = st.Version

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TransitionModelVersionStageDatabricksFromPb(pb *mlpb.TransitionModelVersionStageDatabricksPb) (*TransitionModelVersionStageDatabricks, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransitionModelVersionStageDatabricks{}
	st.ArchiveExistingVersions = pb.ArchiveExistingVersions
	st.Comment = pb.Comment
	st.Name = pb.Name
	st.Stage = pb.Stage
	st.Version = pb.Version

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

// For activities, this contains the activity recorded for the action. For
// comments, this contains the comment details. For transition requests, this
// contains the transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	// Wire name: 'available_actions'
	AvailableActions []ActivityAction ``
	// User-provided comment associated with the activity, comment, or
	// transition request.
	// Wire name: 'comment'
	Comment string ``
	// Creation time of the object, as a Unix timestamp in milliseconds.
	// Wire name: 'creation_timestamp'
	CreationTimestamp int64 ``
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
	ToStage string ``
	// The username of the user that created the object.
	// Wire name: 'user_id'
	UserId          string   ``
	ForceSendFields []string `tf:"-"`
}

func (st TransitionRequest) MarshalJSON() ([]byte, error) {
	pb, err := TransitionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TransitionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.TransitionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TransitionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TransitionRequestToPb(st *TransitionRequest) (*mlpb.TransitionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.TransitionRequestPb{}

	var availableActionsPb []mlpb.ActivityActionPb
	for _, item := range st.AvailableActions {
		itemPb, err := ActivityActionToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			availableActionsPb = append(availableActionsPb, *itemPb)
		}
	}
	pb.AvailableActions = availableActionsPb
	pb.Comment = st.Comment
	pb.CreationTimestamp = st.CreationTimestamp
	pb.ToStage = st.ToStage
	pb.UserId = st.UserId

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func TransitionRequestFromPb(pb *mlpb.TransitionRequestPb) (*TransitionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransitionRequest{}

	var availableActionsField []ActivityAction
	for _, itemPb := range pb.AvailableActions {
		item, err := ActivityActionFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			availableActionsField = append(availableActionsField, *item)
		}
	}
	st.AvailableActions = availableActionsField
	st.Comment = pb.Comment
	st.CreationTimestamp = pb.CreationTimestamp
	st.ToStage = pb.ToStage
	st.UserId = pb.UserId

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type TransitionStageResponse struct {
	// Updated model version
	// Wire name: 'model_version_databricks'
	ModelVersionDatabricks *ModelVersionDatabricks ``
}

func (st TransitionStageResponse) MarshalJSON() ([]byte, error) {
	pb, err := TransitionStageResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *TransitionStageResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.TransitionStageResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := TransitionStageResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func TransitionStageResponseToPb(st *TransitionStageResponse) (*mlpb.TransitionStageResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.TransitionStageResponsePb{}
	modelVersionDatabricksPb, err := ModelVersionDatabricksToPb(st.ModelVersionDatabricks)
	if err != nil {
		return nil, err
	}
	if modelVersionDatabricksPb != nil {
		pb.ModelVersionDatabricks = modelVersionDatabricksPb
	}

	return pb, nil
}

func TransitionStageResponseFromPb(pb *mlpb.TransitionStageResponsePb) (*TransitionStageResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &TransitionStageResponse{}
	modelVersionDatabricksField, err := ModelVersionDatabricksFromPb(pb.ModelVersionDatabricks)
	if err != nil {
		return nil, err
	}
	if modelVersionDatabricksField != nil {
		st.ModelVersionDatabricks = modelVersionDatabricksField
	}

	return st, nil
}

// Details required to edit a comment on a model version.
type UpdateComment struct {
	// User-provided comment on the action.
	// Wire name: 'comment'
	Comment string ``
	// Unique identifier of an activity
	// Wire name: 'id'
	Id string ``
}

func (st UpdateComment) MarshalJSON() ([]byte, error) {
	pb, err := UpdateCommentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateComment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateCommentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateCommentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateCommentToPb(st *UpdateComment) (*mlpb.UpdateCommentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateCommentPb{}
	pb.Comment = st.Comment
	pb.Id = st.Id

	return pb, nil
}

func UpdateCommentFromPb(pb *mlpb.UpdateCommentPb) (*UpdateComment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateComment{}
	st.Comment = pb.Comment
	st.Id = pb.Id

	return st, nil
}

type UpdateCommentResponse struct {
	// Updated comment object
	// Wire name: 'comment'
	Comment *CommentObject ``
}

func (st UpdateCommentResponse) MarshalJSON() ([]byte, error) {
	pb, err := UpdateCommentResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateCommentResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateCommentResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateCommentResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateCommentResponseToPb(st *UpdateCommentResponse) (*mlpb.UpdateCommentResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateCommentResponsePb{}
	commentPb, err := CommentObjectToPb(st.Comment)
	if err != nil {
		return nil, err
	}
	if commentPb != nil {
		pb.Comment = commentPb
	}

	return pb, nil
}

func UpdateCommentResponseFromPb(pb *mlpb.UpdateCommentResponsePb) (*UpdateCommentResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateCommentResponse{}
	commentField, err := CommentObjectFromPb(pb.Comment)
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
	ExperimentId string ``
	// If provided, the experiment's name is changed to the new name. The new
	// name must be unique.
	// Wire name: 'new_name'
	NewName         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st UpdateExperiment) MarshalJSON() ([]byte, error) {
	pb, err := UpdateExperimentToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateExperiment) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateExperimentPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateExperimentFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateExperimentToPb(st *UpdateExperiment) (*mlpb.UpdateExperimentPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateExperimentPb{}
	pb.ExperimentId = st.ExperimentId
	pb.NewName = st.NewName

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateExperimentFromPb(pb *mlpb.UpdateExperimentPb) (*UpdateExperiment, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateExperiment{}
	st.ExperimentId = pb.ExperimentId
	st.NewName = pb.NewName

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateFeatureTagRequest struct {

	// Wire name: 'feature_name'
	FeatureName string `tf:"-"`

	// Wire name: 'feature_tag'
	FeatureTag FeatureTag ``

	// Wire name: 'key'
	Key string `tf:"-"`

	// Wire name: 'table_name'
	TableName string `tf:"-"`
	// The list of fields to update.
	// Wire name: 'update_mask'
	UpdateMask      string   `tf:"-"` //legacy
	ForceSendFields []string `tf:"-"`
}

func (st UpdateFeatureTagRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateFeatureTagRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateFeatureTagRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateFeatureTagRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateFeatureTagRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateFeatureTagRequestToPb(st *UpdateFeatureTagRequest) (*mlpb.UpdateFeatureTagRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateFeatureTagRequestPb{}
	pb.FeatureName = st.FeatureName
	featureTagPb, err := FeatureTagToPb(&st.FeatureTag)
	if err != nil {
		return nil, err
	}
	if featureTagPb != nil {
		pb.FeatureTag = *featureTagPb
	}
	pb.Key = st.Key
	pb.TableName = st.TableName
	pb.UpdateMask = st.UpdateMask

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateFeatureTagRequestFromPb(pb *mlpb.UpdateFeatureTagRequestPb) (*UpdateFeatureTagRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateFeatureTagRequest{}
	st.FeatureName = pb.FeatureName
	featureTagField, err := FeatureTagFromPb(&pb.FeatureTag)
	if err != nil {
		return nil, err
	}
	if featureTagField != nil {
		st.FeatureTag = *featureTagField
	}
	st.Key = pb.Key
	st.TableName = pb.TableName
	st.UpdateMask = pb.UpdateMask

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateModelRequest struct {
	// If provided, updates the description for this `registered_model`.
	// Wire name: 'description'
	Description string ``
	// Registered model unique name identifier.
	// Wire name: 'name'
	Name            string   ``
	ForceSendFields []string `tf:"-"`
}

func (st UpdateModelRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateModelRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateModelRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateModelRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateModelRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateModelRequestToPb(st *UpdateModelRequest) (*mlpb.UpdateModelRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateModelRequestPb{}
	pb.Description = st.Description
	pb.Name = st.Name

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateModelRequestFromPb(pb *mlpb.UpdateModelRequestPb) (*UpdateModelRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelRequest{}
	st.Description = pb.Description
	st.Name = pb.Name

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateModelResponse struct {

	// Wire name: 'registered_model'
	RegisteredModel *Model ``
}

func (st UpdateModelResponse) MarshalJSON() ([]byte, error) {
	pb, err := UpdateModelResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateModelResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateModelResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateModelResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateModelResponseToPb(st *UpdateModelResponse) (*mlpb.UpdateModelResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateModelResponsePb{}
	registeredModelPb, err := ModelToPb(st.RegisteredModel)
	if err != nil {
		return nil, err
	}
	if registeredModelPb != nil {
		pb.RegisteredModel = registeredModelPb
	}

	return pb, nil
}

func UpdateModelResponseFromPb(pb *mlpb.UpdateModelResponsePb) (*UpdateModelResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelResponse{}
	registeredModelField, err := ModelFromPb(pb.RegisteredModel)
	if err != nil {
		return nil, err
	}
	if registeredModelField != nil {
		st.RegisteredModel = registeredModelField
	}

	return st, nil
}

type UpdateModelVersionRequest struct {
	// If provided, updates the description for this `registered_model`.
	// Wire name: 'description'
	Description string ``
	// Name of the registered model
	// Wire name: 'name'
	Name string ``
	// Model version number
	// Wire name: 'version'
	Version         string   ``
	ForceSendFields []string `tf:"-"`
}

func (st UpdateModelVersionRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateModelVersionRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateModelVersionRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateModelVersionRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateModelVersionRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateModelVersionRequestToPb(st *UpdateModelVersionRequest) (*mlpb.UpdateModelVersionRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateModelVersionRequestPb{}
	pb.Description = st.Description
	pb.Name = st.Name
	pb.Version = st.Version

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateModelVersionRequestFromPb(pb *mlpb.UpdateModelVersionRequestPb) (*UpdateModelVersionRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelVersionRequest{}
	st.Description = pb.Description
	st.Name = pb.Name
	st.Version = pb.Version

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateModelVersionResponse struct {
	// Return new version number generated for this model in registry.
	// Wire name: 'model_version'
	ModelVersion *ModelVersion ``
}

func (st UpdateModelVersionResponse) MarshalJSON() ([]byte, error) {
	pb, err := UpdateModelVersionResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateModelVersionResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateModelVersionResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateModelVersionResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateModelVersionResponseToPb(st *UpdateModelVersionResponse) (*mlpb.UpdateModelVersionResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateModelVersionResponsePb{}
	modelVersionPb, err := ModelVersionToPb(st.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionPb != nil {
		pb.ModelVersion = modelVersionPb
	}

	return pb, nil
}

func UpdateModelVersionResponseFromPb(pb *mlpb.UpdateModelVersionResponsePb) (*UpdateModelVersionResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateModelVersionResponse{}
	modelVersionField, err := ModelVersionFromPb(pb.ModelVersion)
	if err != nil {
		return nil, err
	}
	if modelVersionField != nil {
		st.ModelVersion = modelVersionField
	}

	return st, nil
}

type UpdateOnlineStoreRequest struct {
	// The name of the online store. This is the unique identifier for the
	// online store.
	// Wire name: 'name'
	Name string `tf:"-"`
	// Online store to update.
	// Wire name: 'online_store'
	OnlineStore OnlineStore ``
	// The list of fields to update.
	// Wire name: 'update_mask'
	UpdateMask string `tf:"-"` //legacy

}

func (st UpdateOnlineStoreRequest) MarshalJSON() ([]byte, error) {
	pb, err := UpdateOnlineStoreRequestToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateOnlineStoreRequest) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateOnlineStoreRequestPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateOnlineStoreRequestFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateOnlineStoreRequestToPb(st *UpdateOnlineStoreRequest) (*mlpb.UpdateOnlineStoreRequestPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateOnlineStoreRequestPb{}
	pb.Name = st.Name
	onlineStorePb, err := OnlineStoreToPb(&st.OnlineStore)
	if err != nil {
		return nil, err
	}
	if onlineStorePb != nil {
		pb.OnlineStore = *onlineStorePb
	}
	pb.UpdateMask = st.UpdateMask

	return pb, nil
}

func UpdateOnlineStoreRequestFromPb(pb *mlpb.UpdateOnlineStoreRequestPb) (*UpdateOnlineStoreRequest, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateOnlineStoreRequest{}
	st.Name = pb.Name
	onlineStoreField, err := OnlineStoreFromPb(&pb.OnlineStore)
	if err != nil {
		return nil, err
	}
	if onlineStoreField != nil {
		st.OnlineStore = *onlineStoreField
	}
	st.UpdateMask = pb.UpdateMask

	return st, nil
}

// Details required to update a registry webhook. Only the fields that need to
// be updated should be specified, and both `http_url_spec` and `job_spec`
// should not be specified in the same request.
type UpdateRegistryWebhook struct {
	// User-specified description for the webhook.
	// Wire name: 'description'
	Description string ``
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
	Events []RegistryWebhookEvent ``

	// Wire name: 'http_url_spec'
	HttpUrlSpec *HttpUrlSpec ``
	// Webhook ID
	// Wire name: 'id'
	Id string ``

	// Wire name: 'job_spec'
	JobSpec *JobSpec ``

	// Wire name: 'status'
	Status          RegistryWebhookStatus ``
	ForceSendFields []string              `tf:"-"`
}

func (st UpdateRegistryWebhook) MarshalJSON() ([]byte, error) {
	pb, err := UpdateRegistryWebhookToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateRegistryWebhook) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateRegistryWebhookPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateRegistryWebhookFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateRegistryWebhookToPb(st *UpdateRegistryWebhook) (*mlpb.UpdateRegistryWebhookPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateRegistryWebhookPb{}
	pb.Description = st.Description

	var eventsPb []mlpb.RegistryWebhookEventPb
	for _, item := range st.Events {
		itemPb, err := RegistryWebhookEventToPb(&item)
		if err != nil {
			return nil, err
		}
		if itemPb != nil {
			eventsPb = append(eventsPb, *itemPb)
		}
	}
	pb.Events = eventsPb
	httpUrlSpecPb, err := HttpUrlSpecToPb(st.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecPb != nil {
		pb.HttpUrlSpec = httpUrlSpecPb
	}
	pb.Id = st.Id
	jobSpecPb, err := JobSpecToPb(st.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecPb != nil {
		pb.JobSpec = jobSpecPb
	}
	statusPb, err := RegistryWebhookStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateRegistryWebhookFromPb(pb *mlpb.UpdateRegistryWebhookPb) (*UpdateRegistryWebhook, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRegistryWebhook{}
	st.Description = pb.Description

	var eventsField []RegistryWebhookEvent
	for _, itemPb := range pb.Events {
		item, err := RegistryWebhookEventFromPb(&itemPb)
		if err != nil {
			return nil, err
		}
		if item != nil {
			eventsField = append(eventsField, *item)
		}
	}
	st.Events = eventsField
	httpUrlSpecField, err := HttpUrlSpecFromPb(pb.HttpUrlSpec)
	if err != nil {
		return nil, err
	}
	if httpUrlSpecField != nil {
		st.HttpUrlSpec = httpUrlSpecField
	}
	st.Id = pb.Id
	jobSpecField, err := JobSpecFromPb(pb.JobSpec)
	if err != nil {
		return nil, err
	}
	if jobSpecField != nil {
		st.JobSpec = jobSpecField
	}
	statusField, err := RegistryWebhookStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateRun struct {
	// Unix timestamp in milliseconds of when the run ended.
	// Wire name: 'end_time'
	EndTime int64 ``
	// ID of the run to update. Must be provided.
	// Wire name: 'run_id'
	RunId string ``
	// Updated name of the run.
	// Wire name: 'run_name'
	RunName string ``
	// [Deprecated, use `run_id` instead] ID of the run to update. This field
	// will be removed in a future MLflow version.
	// Wire name: 'run_uuid'
	RunUuid string ``
	// Updated status of the run.
	// Wire name: 'status'
	Status          UpdateRunStatus ``
	ForceSendFields []string        `tf:"-"`
}

func (st UpdateRun) MarshalJSON() ([]byte, error) {
	pb, err := UpdateRunToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateRun) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateRunPb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateRunFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateRunToPb(st *UpdateRun) (*mlpb.UpdateRunPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateRunPb{}
	pb.EndTime = st.EndTime
	pb.RunId = st.RunId
	pb.RunName = st.RunName
	pb.RunUuid = st.RunUuid
	statusPb, err := UpdateRunStatusToPb(&st.Status)
	if err != nil {
		return nil, err
	}
	if statusPb != nil {
		pb.Status = *statusPb
	}

	if len(st.ForceSendFields) > 0 {
		pb.ForceSendFields = st.ForceSendFields
	}
	return pb, nil
}

func UpdateRunFromPb(pb *mlpb.UpdateRunPb) (*UpdateRun, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRun{}
	st.EndTime = pb.EndTime
	st.RunId = pb.RunId
	st.RunName = pb.RunName
	st.RunUuid = pb.RunUuid
	statusField, err := UpdateRunStatusFromPb(&pb.Status)
	if err != nil {
		return nil, err
	}
	if statusField != nil {
		st.Status = *statusField
	}

	if len(pb.ForceSendFields) > 0 {
		st.ForceSendFields = pb.ForceSendFields
	}
	return st, nil
}

type UpdateRunResponse struct {
	// Updated metadata of the run.
	// Wire name: 'run_info'
	RunInfo *RunInfo ``
}

func (st UpdateRunResponse) MarshalJSON() ([]byte, error) {
	pb, err := UpdateRunResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateRunResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateRunResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateRunResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateRunResponseToPb(st *UpdateRunResponse) (*mlpb.UpdateRunResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateRunResponsePb{}
	runInfoPb, err := RunInfoToPb(st.RunInfo)
	if err != nil {
		return nil, err
	}
	if runInfoPb != nil {
		pb.RunInfo = runInfoPb
	}

	return pb, nil
}

func UpdateRunResponseFromPb(pb *mlpb.UpdateRunResponsePb) (*UpdateRunResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateRunResponse{}
	runInfoField, err := RunInfoFromPb(pb.RunInfo)
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

// Values returns all possible values for UpdateRunStatus.
//
// There is no guarantee on the order of the values in the slice.
func (f *UpdateRunStatus) Values() []UpdateRunStatus {
	return []UpdateRunStatus{
		UpdateRunStatusFailed,
		UpdateRunStatusFinished,
		UpdateRunStatusKilled,
		UpdateRunStatusRunning,
		UpdateRunStatusScheduled,
	}
}

// Type always returns UpdateRunStatus to satisfy [pflag.Value] interface
func (f *UpdateRunStatus) Type() string {
	return "UpdateRunStatus"
}

func UpdateRunStatusToPb(st *UpdateRunStatus) (*mlpb.UpdateRunStatusPb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.UpdateRunStatusPb(*st)
	return &pb, nil
}

func UpdateRunStatusFromPb(pb *mlpb.UpdateRunStatusPb) (*UpdateRunStatus, error) {
	if pb == nil {
		return nil, nil
	}
	st := UpdateRunStatus(*pb)
	return &st, nil
}

type UpdateWebhookResponse struct {

	// Wire name: 'webhook'
	Webhook *RegistryWebhook ``
}

func (st UpdateWebhookResponse) MarshalJSON() ([]byte, error) {
	pb, err := UpdateWebhookResponseToPb(&st)
	if err != nil {
		return nil, err
	}
	return json.Marshal(pb)
}

func (st *UpdateWebhookResponse) UnmarshalJSON(b []byte) error {
	if st == nil {
		return fmt.Errorf("json.Unmarshal on nil pointer")
	}
	pb := &mlpb.UpdateWebhookResponsePb{}
	err := json.Unmarshal(b, pb)
	if err != nil {
		return err
	}
	tmp, err := UpdateWebhookResponseFromPb(pb)
	if err != nil {
		return err
	}
	*st = *tmp
	return nil
}

func UpdateWebhookResponseToPb(st *UpdateWebhookResponse) (*mlpb.UpdateWebhookResponsePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := &mlpb.UpdateWebhookResponsePb{}
	webhookPb, err := RegistryWebhookToPb(st.Webhook)
	if err != nil {
		return nil, err
	}
	if webhookPb != nil {
		pb.Webhook = webhookPb
	}

	return pb, nil
}

func UpdateWebhookResponseFromPb(pb *mlpb.UpdateWebhookResponsePb) (*UpdateWebhookResponse, error) {
	if pb == nil {
		return nil, nil
	}
	st := &UpdateWebhookResponse{}
	webhookField, err := RegistryWebhookFromPb(pb.Webhook)
	if err != nil {
		return nil, err
	}
	if webhookField != nil {
		st.Webhook = webhookField
	}

	return st, nil
}

// Qualifier for the view type.
type ViewType string

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

// Values returns all possible values for ViewType.
//
// There is no guarantee on the order of the values in the slice.
func (f *ViewType) Values() []ViewType {
	return []ViewType{
		ViewTypeActiveOnly,
		ViewTypeAll,
		ViewTypeDeletedOnly,
	}
}

// Type always returns ViewType to satisfy [pflag.Value] interface
func (f *ViewType) Type() string {
	return "ViewType"
}

func ViewTypeToPb(st *ViewType) (*mlpb.ViewTypePb, error) {
	if st == nil {
		return nil, nil
	}
	pb := mlpb.ViewTypePb(*st)
	return &pb, nil
}

func ViewTypeFromPb(pb *mlpb.ViewTypePb) (*ViewType, error) {
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
