// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package modelversioncomments

// all definitions in this file are in alphabetical order
// An action that a user (with sufficient permissions) could take on a comment.
// Valid values are: * `EDIT_COMMENT`: Edit the comment * `DELETE_COMMENT`:
// Delete the comment
type CommentActivityAction string

// An action that a user (with sufficient permissions) could take on a comment.
// Valid values are: * `EDIT_COMMENT`: Edit the comment * `DELETE_COMMENT`:
// Delete the comment
const CommentActivityActionEditComment CommentActivityAction = `EDIT_COMMENT`
// An action that a user (with sufficient permissions) could take on a comment.
// Valid values are: * `EDIT_COMMENT`: Edit the comment * `DELETE_COMMENT`:
// Delete the comment
const CommentActivityActionDeleteComment CommentActivityAction = `DELETE_COMMENT`
// Comment details.
type CommentObject struct {
    // Array of actions on the activity allowed for the current viewer.
    AvailableActions []CommentActivityAction `json:"available_actions,omitempty"`
    // User-provided comment on the action.
    Comment string `json:"comment,omitempty"`
    // Creation time of the object, as a Unix timestamp in milliseconds.
    CreationTimestamp int64 `json:"creation_timestamp,omitempty"`
    // Time of the object at last update, as a Unix timestamp in milliseconds.
    LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
    // The username of the user that created the object.
    UserId string `json:"user_id,omitempty"`
}


type CreateCommentRequest struct {
    // User-provided comment on the action.
    Comment string `json:"comment"`
    // Name of the model.
    Name string `json:"name"`
    // Version of the model.
    Version string `json:"version"`
}


type CreateCommentResponse struct {
    
    Comment *CommentObject `json:"comment,omitempty"`
}


type DeleteCommentRequest struct {
    // Unique identifier of an activity
    Id string `json:"id"`
}


type UpdateCommentRequest struct {
    // User-provided comment on the action.
    Comment string `json:"comment"`
    // Unique identifier of an activity
    Id string `json:"id"`
}


type UpdateCommentResponse struct {
    
    Comment *CommentObject `json:"comment,omitempty"`
}

