// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package tags

import (
	"github.com/databricks/databricks-sdk-go/marshal"
)

type CreateTagAssignmentRequest struct {
	TagAssignment TagAssignment `json:"tag_assignment"`
}

type CreateTagPolicyRequest struct {
	TagPolicy TagPolicy `json:"tag_policy"`
}

type DeleteTagAssignmentRequest struct {
	// The identifier of the entity to which the tag is assigned
	EntityId string `json:"-" url:"-"`
	// The type of entity to which the tag is assigned. Allowed value is
	// dashboards
	EntityType string `json:"-" url:"-"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey string `json:"-" url:"-"`
}

type DeleteTagPolicyRequest struct {
	TagKey string `json:"-" url:"-"`
}

type GetTagAssignmentRequest struct {
	// The identifier of the entity to which the tag is assigned
	EntityId string `json:"-" url:"-"`
	// The type of entity to which the tag is assigned. Allowed value is
	// dashboards
	EntityType string `json:"-" url:"-"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey string `json:"-" url:"-"`
}

type GetTagPolicyRequest struct {
	TagKey string `json:"-" url:"-"`
}

type ListTagAssignmentsRequest struct {
	// The identifier of the entity to which the tag is assigned
	EntityId string `json:"-" url:"-"`
	// The type of entity to which the tag is assigned. Allowed value is
	// dashboards
	EntityType string `json:"-" url:"-"`
	// Optional. Maximum number of tag assignments to return in a single page
	PageSize int `json:"-" url:"page_size,omitempty"`
	// Pagination token to go to the next page of tag assignments. Requests
	// first page if absent.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTagAssignmentsRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTagAssignmentsRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTagAssignmentsResponse struct {
	// Pagination token to request the next page of tag assignments
	NextPageToken string `json:"next_page_token,omitempty"`

	TagAssignments []TagAssignment `json:"tag_assignments,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTagAssignmentsResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTagAssignmentsResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTagPoliciesRequest struct {
	// The maximum number of results to return in this request. Fewer results
	// may be returned than requested. If unspecified or set to 0, this defaults
	// to 1000. The maximum value is 1000; values above 1000 will be coerced
	// down to 1000.
	PageSize int `json:"-" url:"page_size,omitempty"`
	// An optional page token received from a previous list tag policies call.
	PageToken string `json:"-" url:"page_token,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTagPoliciesRequest) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTagPoliciesRequest) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type ListTagPoliciesResponse struct {
	NextPageToken string `json:"next_page_token,omitempty"`

	TagPolicies []TagPolicy `json:"tag_policies,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *ListTagPoliciesResponse) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s ListTagPoliciesResponse) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TagAssignment struct {
	// The identifier of the entity to which the tag is assigned
	EntityId string `json:"entity_id"`
	// The type of entity to which the tag is assigned. Allowed value is
	// dashboards
	EntityType string `json:"entity_type"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey string `json:"tag_key"`
	// The value of the tag
	TagValue string `json:"tag_value,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TagAssignment) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TagAssignment) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type TagPolicy struct {
	// Timestamp when the tag policy was created
	CreateTime string `json:"create_time,omitempty"`

	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`

	TagKey string `json:"tag_key"`
	// Timestamp when the tag policy was last updated
	UpdateTime string `json:"update_time,omitempty"`

	Values []Value `json:"values,omitempty"`

	ForceSendFields []string `json:"-" url:"-"`
}

func (s *TagPolicy) UnmarshalJSON(b []byte) error {
	return marshal.Unmarshal(b, s)
}

func (s TagPolicy) MarshalJSON() ([]byte, error) {
	return marshal.Marshal(s)
}

type UpdateTagAssignmentRequest struct {
	// The identifier of the entity to which the tag is assigned
	EntityId string `json:"-" url:"-"`
	// The type of entity to which the tag is assigned. Allowed value is
	// dashboards
	EntityType string `json:"-" url:"-"`

	TagAssignment TagAssignment `json:"tag_assignment"`
	// The key of the tag. The characters , . : / - = and leading/trailing
	// spaces are not allowed
	TagKey string `json:"-" url:"-"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask string `json:"-" url:"update_mask"`
}

type UpdateTagPolicyRequest struct {
	TagKey string `json:"-" url:"-"`

	TagPolicy TagPolicy `json:"tag_policy"`
	// The field mask must be a single string, with multiple fields separated by
	// commas (no spaces). The field path is relative to the resource object,
	// using a dot (`.`) to navigate sub-fields (e.g., `author.given_name`).
	// Specification of elements in sequence or map fields is not allowed, as
	// only the entire collection field can be specified. Field names must
	// exactly match the resource field names.
	//
	// A field mask of `*` indicates full replacement. It’s recommended to
	// always explicitly list the fields being updated and avoid using `*`
	// wildcards, as it can lead to unintended results if the API changes in the
	// future.
	UpdateMask string `json:"-" url:"update_mask"`
}

type Value struct {
	Name string `json:"name"`
}
