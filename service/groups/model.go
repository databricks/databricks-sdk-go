// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package groups

// all definitions in this file are in alphabetical order


type CoreGroup string

const CoreGroupCoreGroupSchema CoreGroup = `urn:ietf:params:scim:schemas:core:2.0:Group`

type CreateGroupRequest struct {
	// String that represents a human-readable group name 
	DisplayName string `json:"displayName,omitempty"`
	// All the members associated with the group 
	Members []CreateGroupRequestMembers `json:"members,omitempty"`
	// SCIM schema used for the group object. 
	Schemas []CoreGroup `json:"schemas,omitempty"`
}


type CreateGroupRequestMembers struct {
	// The id of the principal 
	Value string `json:"value,omitempty"`
}


type DeleteGroupRequest struct {
	// Unique ID for a group in the &lt;Workspace&gt;. 
	GroupId string `path:"group_id"`
}


type FetchGroupRequest struct {
	
	GroupId string `path:"group_id"`
}


type Group struct {
	// String that represents a human-readable group name 
	DisplayName string `json:"displayName,omitempty"`
	// Databricks group ID 
	Id string `json:"id,omitempty"`
	// All the members associated with the group 
	Members []GroupMember `json:"members,omitempty"`
	// SCIM schema used for the group object. 
	Schemas []CoreGroup `json:"schemas,omitempty"`
}


type GroupMember struct {
	// A url reference to the principal 
	Ref string `json:"$ref,omitempty"`
	// The display name of the principal 
	Display string `json:"display,omitempty"`
	// The id of the principal 
	Value string `json:"value,omitempty"`
}


type GroupRemovePatchOp struct {
	// Type of patch operation. 
	Op GroupRemovePatchOpOp `json:"op,omitempty"`
	// Path of remove group member. 
	Path GroupRemovePatchOpPath `json:"path,omitempty"`
}

// Type of patch operation. 
// Type of patch operation. 
type GroupRemovePatchOpOp string
// Type of patch operation. 
const GroupRemovePatchOpOpRemove GroupRemovePatchOpOp = `remove`
// Path of remove group member. 
// Path of remove group member. 
type GroupRemovePatchOpPath string
// Path of remove group member. 
const GroupRemovePatchOpPathMembersValueEqPrincipalid GroupRemovePatchOpPath = `members[value eq &#34;principalId&#34;]`


type List string

const ListListResponseMessage List = `urn:ietf:params:scim:api:messages:2.0:ListResponse`

type ListGroupsRequest struct {
	// Comma-separated list of attributes to return in response. 
	Attributes string `query:"attributes,omitempty"`
	// Desired number of results per page. 
	Count int `query:"count,omitempty"`
	// Comma-separated list of attributes to exclude in response. 
	ExcludedAttributes string `query:"excludedAttributes,omitempty"`
	// Query by which the results have to be filtered. Supported operators are 
    // equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`). 
    // Additionally, simple expressions can be formed using logical operators - 
    // `and` and `or`. The [SCIM 
    // RFC](https://tools.ietf.org/html/rfc7644#section-3.4.2.2) has more 
    // details but we currently only support simple expressions. 
	Filter string `query:"filter,omitempty"`
	// Attribute to sort the results. 
	SortBy string `query:"sortBy,omitempty"`
	// The order to sort the results. 
	SortOrder ListGroupsSortOrder `query:"sortOrder,omitempty"`
	// Specifies the index of the first result. First item is number 1. 
	StartIndex int `query:"startIndex,omitempty"`
}


type ListGroupsResponse struct {
	// User objects returned in the response. 
	Resources []Group `json:"Resources,omitempty"`
	// Total results returned in the response. 
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	// SCIM ListOp schema used in the response. 
	Schemas []List `json:"schemas,omitempty"`
	// Starting index of all the results that matched the request filters. 
    // First item is number 1. 
	StartIndex int64 `json:"startIndex,omitempty"`
	// Total results that match the request filters. 
	TotalResults int64 `json:"totalResults,omitempty"`
}



type ListGroupsSortOrder string

const ListGroupsSortOrderAscending ListGroupsSortOrder = `ascending`
const ListGroupsSortOrderDescending ListGroupsSortOrder = `descending`


type Update string

const UpdatePatchOpMessage Update = `urn:ietf:params:scim:api:messages:2.0:PatchOp`
// Type of patch operation. 
// Type of patch operation. 
type UpdateGroupOp string
// Type of patch operation. 
const UpdateGroupOpRemove UpdateGroupOp = `remove`
// Path of removing admin status 
// Path of removing admin status 
type UpdateGroupPath string
// Path of removing admin status 
const UpdateGroupPathRolesValueEqAccountAdmin UpdateGroupPath = `roles[value eq \&#34;account_admin\&#34;]`

type UpdateGroupRequest struct {
	// Patch operations to be performed on the user object. 
	Operations []GroupRemovePatchOp `json:"Operations,omitempty"`
	// Unique ID for a group in the &lt;Workspace&gt;. 
	GroupId string `path:"group_id"`
	// Type of patch operation. 
	Op UpdateGroupOp `json:"op,omitempty"`
	// Path of removing admin status 
	Path UpdateGroupPath `json:"path,omitempty"`
	// SCIM PatchOp schema used to update user information. 
	Schemas []Update `json:"schemas,omitempty"`
	
	Value []UpdateGroupValue `json:"value,omitempty"`
}


type UpdateGroupValue struct {
	// Account admin field. 
	Value UpdateGroupValueValue `json:"value,omitempty"`
}

// Account admin field. 
// Account admin field. 
type UpdateGroupValueValue string
// Account admin field. 
const UpdateGroupValueValueAccountAdmin UpdateGroupValueValue = `account_admin`
