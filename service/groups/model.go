package groups


// 

type CoreGroup string

// 
const CoreGroupSchema CoreGroup = "urn:ietf:params:scim:schemas:core:2.0:Group"



// 
type DeleteGroupRequest struct {
	
	// Unique ID for a group in the &lt;Workspace&gt;.
	GroupId string `path:"group_id"`
	
}


// 
type FetchGroupRequest struct {
	
	// 
	GroupId string `path:"group_id"`
	
}


// 
type FetchGroupResponse struct {
	
	// String that represents a human-readable group name
	DisplayName string `json:"displayName,omitempty"`
	
	// Databricks group ID
	Id string `json:"id,omitempty"`
	
	// All the members associated with the group
	Members []GroupMember `json:"members,omitempty"`
	
	// SCIM schema used for the group object.
	Schemas []CoreGroup `json:"schemas,omitempty"`
	
}


// 
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


// 
type GroupMember struct {
	
	// A url reference to the principal
	$Ref string `json:"$ref,omitempty"`
	
	// The display name of the principal
	Display string `json:"display,omitempty"`
	
	// The id of the principal
	Value string `json:"value,omitempty"`
	
}


// 
type GroupRemovePatchOp struct {
	
	// Type of patch operation.
	Op any /* {{op Type of patch operation.} false 0xc0002ac780 true false false} */ `json:"op,omitempty"`
	
	// Path of remove group member.
	Path any /* {{path Path of remove group member.} false 0xc0002ac7e0 true false false} */ `json:"path,omitempty"`
	
}


// 

type List string

// 
const ListResponseMessage List = "urn:ietf:params:scim:api:messages:2.0:ListResponse"



// 
type ListGroupsRequest struct {
	
	// Comma-separated list of attributes to return in response.
	Attributes string `query:"attributes,omitempty"`
	
	// Desired number of results per page.
	Count int `query:"count,omitempty"`
	
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `query:"excludedAttributes,omitempty"`
	
	// Query by which the results have to be filtered. Supported operators are equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`). Additionally, simple expressions can be formed using logical operators - `and` and `or`. The [SCIM RFC](https://tools.ietf.org/html/rfc7644#section-3.4.2.2) has more details but we currently only support simple expressions.
	Filter string `query:"filter,omitempty"`
	
	// Attribute to sort the results.
	SortBy string `query:"sortBy,omitempty"`
	
	// The order to sort the results.
	SortOrder any /* {{sortOrder The order to sort the results.} false 0xc0002acfc0 false false true} */ `query:"sortOrder,omitempty"`
	
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `query:"startIndex,omitempty"`
	
}


// 
type ListGroupsResponse struct {
	
	// User objects returned in the response.
	Resources []Group `json:"Resources,omitempty"`
	
	// Total results returned in the response.
	ItemsPerPage int64 `json:"itemsPerPage,omitempty"`
	
	// SCIM ListOp schema used in the response.
	Schemas []List `json:"schemas,omitempty"`
	
	// Starting index of all the results that matched the request filters. First item is number 1.
	StartIndex int64 `json:"startIndex,omitempty"`
	
	// Total results that match the request filters.
	TotalResults int64 `json:"totalResults,omitempty"`
	
}


// 
type NewGroupRequest struct {
	
	// String that represents a human-readable group name
	DisplayName string `json:"displayName,omitempty"`
	
	// All the members associated with the group
	Members [] `json:"members,omitempty"`
	
	// SCIM schema used for the group object.
	Schemas []CoreGroup `json:"schemas,omitempty"`
	
}


// 
type NewGroupRequest struct {
	
	// String that represents a human-readable group name
	DisplayName string `json:"displayName,omitempty"`
	
	// All the members associated with the group
	Members [] `json:"members,omitempty"`
	
	// SCIM schema used for the group object.
	Schemas []CoreGroup `json:"schemas,omitempty"`
	
}


// 

type Update string

// 
const PatchOpMessage Update = "urn:ietf:params:scim:api:messages:2.0:PatchOp"



// 
type UpdateGroupRequest struct {
	
	// Patch operations to be performed on the user object.
	Operations []GroupRemovePatchOp `json:"Operations,omitempty"`
	
	// Unique ID for a group in the &lt;Workspace&gt;.
	GroupId string `path:"group_id"`
	
	// Type of patch operation.
	Op any /* {{op Type of patch operation.} false 0xc0002aca80 true false false} */ `json:"op,omitempty"`
	
	// Path of removing admin status
	Path any /* {{path Path of removing admin status} false 0xc0002ac840 true false false} */ `json:"path,omitempty"`
	
	// SCIM PatchOp schema used to update user information.
	Schemas []Update `json:"schemas,omitempty"`
	
	// 
	Value [] `json:"value,omitempty"`
	
}

