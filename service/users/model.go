package users


// 

type CoreUser string

// 
const CoreUserSchema CoreUser = "urn:ietf:params:scim:schemas:core:2.0:User"



// 
type DeleteUserRequest struct {
	
	// Unique ID for a user in the &lt;Workspace&gt;.
	UserId string `path:"user_id"`
	
}


// 
type DisplayNamePatchOp struct {
	
	// Type of patch operation.
	Op any /* {{op Type of patch operation.} false 0xc000182e40 true false false} */ `json:"op,omitempty"`
	
	// Field to update.
	Path any /* {{path Field to update.} false 0xc000182ea0 true false false} */ `json:"path,omitempty"`
	
	// 
	Value [] `json:"value,omitempty"`
	
}


// 
type Email struct {
	
	// 
	Primary bool `json:"primary,omitempty"`
	
	// Type of email. `work` is an example.
	Type string `json:"type,omitempty"`
	
	// The actual email.
	Value string `json:"value,omitempty"`
	
}


// 
type FetchUserRequest struct {
	
	// 
	UserId string `path:"user_id"`
	
}


// 
type FetchUserResponse struct {
	
	// String that represents a concatenation of given and family names. For example `John Smith`.
	DisplayName string `json:"displayName,omitempty"`
	
	// All the emails associated with the Databricks user.
	Emails []Email `json:"emails,omitempty"`
	
	// Databricks user ID.
	Id string `json:"id,omitempty"`
	
	// 
	Name *Name `json:"name,omitempty"`
	
	// SCIM schema used for the user object.
	Schemas []CoreUser `json:"schemas,omitempty"`
	
	// Email address of the Databricks user.
	UserName string `json:"userName,omitempty"`
	
}


// 

type List string

// 
const ListResponseMessage List = "urn:ietf:params:scim:api:messages:2.0:ListResponse"



// 
type ListUsersRequest struct {
	
	// Comma-separated list of attributes to return in response.
	Attributes string `query:"attributes,omitempty"`
	
	// Desired number of results per page.
	Count int `query:"count,omitempty"`
	
	// Comma-separated list of attributes to exclude in response.
	ExcludedAttributes string `query:"excludedAttributes,omitempty"`
	
	// Query by which the results have to be filtered. Supported operators are equals(`eq`), contains(`co`), starts with(`sw`) and not equals(`ne`). Additionally, simple expressions can be formed using logical operators - `and` and `or`. The [SCIM RFC](https://tools.ietf.org/html/rfc7644#section-3.4.2.2) has more details but we currently only support simple expressions.
	Filter string `query:"filter,omitempty"`
	
	// Attribute to sort the results. Multi-part paths are supported. For example, `userName`, `name.givenName`, and `emails`.
	SortBy string `query:"sortBy,omitempty"`
	
	// The order to sort the results.
	SortOrder any /* {{sortOrder The order to sort the results.} false 0xc000183560 false false true} */ `query:"sortOrder,omitempty"`
	
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `query:"startIndex,omitempty"`
	
}


// 
type ListUsersResponse struct {
	
	// User objects returned in the response.
	Resources []User `json:"Resources,omitempty"`
	
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
type Name struct {
	
	// Family name of the Databricks user.
	FamilyName string `json:"familyName,omitempty"`
	
	// Given name of the Databricks user.
	GivenName string `json:"givenName,omitempty"`
	
}


// 
type NewUserRequest struct {
	
	// String representing a concatenation of given and family names.
	DisplayName string `json:"displayName,omitempty"`
	
	// SCIM schema used for the user object.
	Schemas []CoreUser `json:"schemas,omitempty"`
	
	// Email address of the Databricks user.
	UserName string `json:"userName,omitempty"`
	
}


// 
type NewUserRequest struct {
	
	// String representing a concatenation of given and family names.
	DisplayName string `json:"displayName,omitempty"`
	
	// SCIM schema used for the user object.
	Schemas []CoreUser `json:"schemas,omitempty"`
	
	// Email address of the Databricks user.
	UserName string `json:"userName,omitempty"`
	
}


// 

type Update string

// 
const PatchOpMessage Update = "urn:ietf:params:scim:api:messages:2.0:PatchOp"



// 
type UpdateUserRequest struct {
	
	// Patch operations to be performed on the user object.
	Operations []DisplayNamePatchOp `json:"Operations,omitempty"`
	
	// Type of patch operation.
	Op any /* {{op Type of patch operation.} false 0xc000182d20 true false false} */ `json:"op,omitempty"`
	
	// Path of removing admin status
	Path any /* {{path Path of removing admin status} false 0xc000183020 true false false} */ `json:"path,omitempty"`
	
	// SCIM PatchOp schema used to update user information.
	Schemas []Update `json:"schemas,omitempty"`
	
	// Unique ID for a user in the &lt;Workspace&gt;.
	UserId string `path:"user_id"`
	
	// 
	Value [] `json:"value,omitempty"`
	
}


// 
type User struct {
	
	// String that represents a concatenation of given and family names. For example `John Smith`.
	DisplayName string `json:"displayName,omitempty"`
	
	// All the emails associated with the Databricks user.
	Emails []Email `json:"emails,omitempty"`
	
	// Databricks user ID.
	Id string `json:"id,omitempty"`
	
	// 
	Name *Name `json:"name,omitempty"`
	
	// SCIM schema used for the user object.
	Schemas []CoreUser `json:"schemas,omitempty"`
	
	// Email address of the Databricks user.
	UserName string `json:"userName,omitempty"`
	
}

