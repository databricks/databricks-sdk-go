// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package transitionrequests

// all definitions in this file are in alphabetical order
// Activity recorded for the action.
type Activity struct {
	ActivityType ActivityType `json:"activity_type,omitempty"`
	// User-provided comment associated with the activity.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Source stage of the transition (if the activity is stage transition
	// related). Valid values are: * `None`: The initial stage of a model
	// version. * `Staging`: Staging or pre-production stage. * `Production`:
	// Production stage. * `Archived`: Archived stage.
	FromStage Stage `json:"from_stage,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id,omitempty"`
	// Time of the object at last update, as a Unix timestamp in milliseconds.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Comment made by system, for example explaining an activity of type
	// `SYSTEM_TRANSITION`. It usually describes a side effect, such as a
	// version being archived as part of another version&#39;s stage transition, and
	// may not be returned for some activity types.
	SystemComment string `json:"system_comment,omitempty"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are: * `None`: The initial stage of a model
	// version. * `Staging`: Staging or pre-production stage. * `Production`:
	// Production stage. * `Archived`: Archived stage.
	ToStage Stage `json:"to_stage,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`
}

// An action that a user (with sufficient permissions) could take on an
// activity. Valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request * `REJECT_TRANSITION_REQUEST`: Reject a transition request
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
type ActivityAction string

// An action that a user (with sufficient permissions) could take on an
// activity. Valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request * `REJECT_TRANSITION_REQUEST`: Reject a transition request
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
const ActivityActionApproveTransitionRequest ActivityAction = `APPROVE_TRANSITION_REQUEST`

// An action that a user (with sufficient permissions) could take on an
// activity. Valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request * `REJECT_TRANSITION_REQUEST`: Reject a transition request
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
const ActivityActionRejectTransitionRequest ActivityAction = `REJECT_TRANSITION_REQUEST`

// An action that a user (with sufficient permissions) could take on an
// activity. Valid values are: * `APPROVE_TRANSITION_REQUEST`: Approve a
// transition request * `REJECT_TRANSITION_REQUEST`: Reject a transition request
// * `CANCEL_TRANSITION_REQUEST`: Cancel (delete) a transition request
const ActivityActionCancelTransitionRequest ActivityAction = `CANCEL_TRANSITION_REQUEST`

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition. * `REQUESTED_TRANSITION`: User requested the
// corresponding stage transition. * `CANCELLED_REQUEST`: User cancelled an
// existing transition request. * `APPROVED_REQUEST`: User approved the
// corresponding stage transition. * `REJECTED_REQUEST`: User rejected the
// coressponding stage transition. * `SYSTEM_TRANSITION`: For events performed
// as a side effect, such as archiving existing model versions in a stage.
type ActivityType string

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition. * `REQUESTED_TRANSITION`: User requested the
// corresponding stage transition. * `CANCELLED_REQUEST`: User cancelled an
// existing transition request. * `APPROVED_REQUEST`: User approved the
// corresponding stage transition. * `REJECTED_REQUEST`: User rejected the
// coressponding stage transition. * `SYSTEM_TRANSITION`: For events performed
// as a side effect, such as archiving existing model versions in a stage.
const ActivityTypeAppliedTransition ActivityType = `APPLIED_TRANSITION`

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition. * `REQUESTED_TRANSITION`: User requested the
// corresponding stage transition. * `CANCELLED_REQUEST`: User cancelled an
// existing transition request. * `APPROVED_REQUEST`: User approved the
// corresponding stage transition. * `REJECTED_REQUEST`: User rejected the
// coressponding stage transition. * `SYSTEM_TRANSITION`: For events performed
// as a side effect, such as archiving existing model versions in a stage.
const ActivityTypeRequestedTransition ActivityType = `REQUESTED_TRANSITION`

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition. * `REQUESTED_TRANSITION`: User requested the
// corresponding stage transition. * `CANCELLED_REQUEST`: User cancelled an
// existing transition request. * `APPROVED_REQUEST`: User approved the
// corresponding stage transition. * `REJECTED_REQUEST`: User rejected the
// coressponding stage transition. * `SYSTEM_TRANSITION`: For events performed
// as a side effect, such as archiving existing model versions in a stage.
const ActivityTypeCancelledRequest ActivityType = `CANCELLED_REQUEST`

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition. * `REQUESTED_TRANSITION`: User requested the
// corresponding stage transition. * `CANCELLED_REQUEST`: User cancelled an
// existing transition request. * `APPROVED_REQUEST`: User approved the
// corresponding stage transition. * `REJECTED_REQUEST`: User rejected the
// coressponding stage transition. * `SYSTEM_TRANSITION`: For events performed
// as a side effect, such as archiving existing model versions in a stage.
const ActivityTypeApprovedRequest ActivityType = `APPROVED_REQUEST`

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition. * `REQUESTED_TRANSITION`: User requested the
// corresponding stage transition. * `CANCELLED_REQUEST`: User cancelled an
// existing transition request. * `APPROVED_REQUEST`: User approved the
// corresponding stage transition. * `REJECTED_REQUEST`: User rejected the
// coressponding stage transition. * `SYSTEM_TRANSITION`: For events performed
// as a side effect, such as archiving existing model versions in a stage.
const ActivityTypeRejectedRequest ActivityType = `REJECTED_REQUEST`

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition. * `REQUESTED_TRANSITION`: User requested the
// corresponding stage transition. * `CANCELLED_REQUEST`: User cancelled an
// existing transition request. * `APPROVED_REQUEST`: User approved the
// corresponding stage transition. * `REJECTED_REQUEST`: User rejected the
// coressponding stage transition. * `SYSTEM_TRANSITION`: For events performed
// as a side effect, such as archiving existing model versions in a stage.
const ActivityTypeNewComment ActivityType = `NEW_COMMENT`

// Type of activity. Valid values are: * `APPLIED_TRANSITION`: User applied the
// corresponding stage transition. * `REQUESTED_TRANSITION`: User requested the
// corresponding stage transition. * `CANCELLED_REQUEST`: User cancelled an
// existing transition request. * `APPROVED_REQUEST`: User approved the
// corresponding stage transition. * `REJECTED_REQUEST`: User rejected the
// coressponding stage transition. * `SYSTEM_TRANSITION`: For events performed
// as a side effect, such as archiving existing model versions in a stage.
const ActivityTypeSystemTransition ActivityType = `SYSTEM_TRANSITION`

type ApproveTransitionRequestRequest struct {
	// Specifies whether to archive all current model versions in the target
	// stage.
	ArchiveExistingVersions bool `json:"archive_existing_versions"`
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`

	Stage Stage `json:"stage"`
	// Version of the model.
	Version string `json:"version"`
}

type ApproveTransitionRequestResponse struct {
	Activity *Activity `json:"activity,omitempty"`
}

type CreateTransitionRequestRequest struct {
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`

	Stage Stage `json:"stage"`
	// Version of the model.
	Version string `json:"version"`
}

type CreateTransitionRequestResponse struct {
	Request *TransitionRequest `json:"request,omitempty"`
}

type DeleteTransitionRequestRequest struct {
	// User-provided comment on the action.
	Comment string ` url:"comment,omitempty"`
	// Username of the user who created this request. Of the transition requests
	// matching the specified details, only the one transition created by this
	// user will be deleted.
	Creator string ` url:"creator,omitempty"`
	// Name of the model.
	Name string ` url:"name,omitempty"`
	// Target stage of the transition request. Valid values are: * `None`: The
	// initial stage of a model version. * `Staging`: Staging or pre-production
	// stage. * `Production`: Production stage. * `Archived`: Archived stage.
	Stage string ` url:"stage,omitempty"`
	// Version of the model.
	Version string ` url:"version,omitempty"`
}

type GetTransitionRequestsRequest struct {
	// Name of the model.
	Name string ` url:"name,omitempty"`
	// Version of the model.
	Version string ` url:"version,omitempty"`
}

type GetTransitionRequestsResponse struct {
	// Array of open transition requests.
	Requests []TransitionRequest `json:"requests,omitempty"`
}

type RejectTransitionRequestRequest struct {
	// User-provided comment on the action.
	Comment string `json:"comment,omitempty"`
	// Name of the model.
	Name string `json:"name"`

	Stage Stage `json:"stage"`
	// Version of the model.
	Version string `json:"version"`
}

type RejectTransitionRequestResponse struct {
	Activity *Activity `json:"activity,omitempty"`
}

// Transition request details.
type TransitionRequest struct {
	// Array of actions on the activity allowed for the current viewer.
	AvailableActions []ActivityAction `json:"available_actions,omitempty"`
	// User-provided comment associated with the transition request.
	Comment string `json:"comment,omitempty"`
	// Creation time of the object, as a Unix timestamp in milliseconds.
	CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
	// Target stage of the transition (if the activity is stage transition
	// related). Valid values are: * `None`: The initial stage of a model
	// version. * `Staging`: Staging or pre-production stage. * `Production`:
	// Production stage. * `Archived`: Archived stage.
	ToStage Stage `json:"to_stage,omitempty"`
	// The username of the user that created the object.
	UserId string `json:"user_id,omitempty"`
}

// Stage of the model version. Valid values are: * `None`: The initial stage of
// a model version. * `Staging`: Staging or pre-production stage. *
// `Production`: Production stage. * `Archived`: Archived stage.
type Stage string

// Stage of the model version. Valid values are: * `None`: The initial stage of
// a model version. * `Staging`: Staging or pre-production stage. *
// `Production`: Production stage. * `Archived`: Archived stage.
const StageNone Stage = `None`

// Stage of the model version. Valid values are: * `None`: The initial stage of
// a model version. * `Staging`: Staging or pre-production stage. *
// `Production`: Production stage. * `Archived`: Archived stage.
const StageStaging Stage = `Staging`

// Stage of the model version. Valid values are: * `None`: The initial stage of
// a model version. * `Staging`: Staging or pre-production stage. *
// `Production`: Production stage. * `Archived`: Archived stage.
const StageProduction Stage = `Production`

// Stage of the model version. Valid values are: * `None`: The initial stage of
// a model version. * `Staging`: Staging or pre-production stage. *
// `Production`: Production stage. * `Archived`: Archived stage.
const StageArchived Stage = `Archived`
