// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package users

// all definitions in this file are in alphabetical order


type CoreUser string

const CoreUserCoreUserSchema CoreUser = `urn:ietf:params:scim:schemas:core:2.0:User`

type CreateUserRequest struct {
    // String representing a concatenation of given and family names. 
    DisplayName string `json:"displayName,omitempty"`
    // SCIM schema used for the user object. 
    Schemas []CoreUser `json:"schemas,omitempty"`
    // Email address of the Databricks user. 
    UserName string `json:"userName,omitempty"`
}


type DeleteUserRequest struct {
    // Unique ID for a user in the &lt;Workspace&gt;. 
    UserId string `path:"user_id"`
}


type DisplayNamePatchOp struct {
    // Type of patch operation. 
    Op DisplayNamePatchOpOp `json:"op,omitempty"`
    // Field to update. 
    Path DisplayNamePatchOpPath `json:"path,omitempty"`
    
    Value []DisplayNamePatchOpValue `json:"value,omitempty"`
}

// Type of patch operation. 
// Type of patch operation. 
type DisplayNamePatchOpOp string
// Type of patch operation. 
const DisplayNamePatchOpOpReplace DisplayNamePatchOpOp = `replace`
// Field to update. 
// Field to update. 
type DisplayNamePatchOpPath string
// Field to update. 
const DisplayNamePatchOpPathDisplayname DisplayNamePatchOpPath = `displayName`

type DisplayNamePatchOpValue struct {
    // New value for the field specified in path. 
    Value string `json:"value,omitempty"`
}


type Email struct {
    
    Primary bool `json:"primary,omitempty"`
    // Type of email. `work` is an example. 
    Type string `json:"type,omitempty"`
    // The actual email. 
    Value string `json:"value,omitempty"`
}


type FetchUserRequest struct {
    // Unique ID for a user in the &lt;Workspace&gt;. 
    UserId string `path:"user_id"`
}



type List string

const ListListResponseMessage List = `urn:ietf:params:scim:api:messages:2.0:ListResponse`

type ListUsersRequest struct {
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
    // Attribute to sort the results. Multi-part paths are supported. For 
    // example, `userName`, `name.givenName`, and `emails`. 
    SortBy string `query:"sortBy,omitempty"`
    // The order to sort the results. 
    SortOrder ListUsersSortOrder `query:"sortOrder,omitempty"`
    // Specifies the index of the first result. First item is number 1. 
    StartIndex int `query:"startIndex,omitempty"`
}


type ListUsersResponse struct {
    // User objects returned in the response. 
    Resources []User `json:"Resources,omitempty"`
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



type ListUsersSortOrder string

const ListUsersSortOrderAscending ListUsersSortOrder = `ascending`
const ListUsersSortOrderDescending ListUsersSortOrder = `descending`

type Name struct {
    // Family name of the Databricks user. 
    FamilyName string `json:"familyName,omitempty"`
    // Given name of the Databricks user. 
    GivenName string `json:"givenName,omitempty"`
}



type Update string

const UpdatePatchOpMessage Update = `urn:ietf:params:scim:api:messages:2.0:PatchOp`
// Type of patch operation. 
// Type of patch operation. 
type UpdateUserOp string
// Type of patch operation. 
const UpdateUserOpRemove UpdateUserOp = `remove`
// Path of removing admin status 
// Path of removing admin status 
type UpdateUserPath string
// Path of removing admin status 
const UpdateUserPathRolesValueEqAccountAdmin UpdateUserPath = `roles[value eq \&#34;account_admin\&#34;]`

type UpdateUserRequest struct {
    // Patch operations to be performed on the user object. 
    Operations []DisplayNamePatchOp `json:"Operations,omitempty"`
    // Type of patch operation. 
    Op UpdateUserOp `json:"op,omitempty"`
    // Path of removing admin status 
    Path UpdateUserPath `json:"path,omitempty"`
    // SCIM PatchOp schema used to update user information. 
    Schemas []Update `json:"schemas,omitempty"`
    // Unique ID for a user in the &lt;Workspace&gt;. 
    UserId string `path:"user_id"`
    
    Value []UpdateUserValue `json:"value,omitempty"`
}


type UpdateUserValue struct {
    // Account admin field. 
    Value UpdateUserValueValue `json:"value,omitempty"`
}

// Account admin field. 
// Account admin field. 
type UpdateUserValueValue string
// Account admin field. 
const UpdateUserValueValueAccountAdmin UpdateUserValueValue = `account_admin`

type User struct {
    // String that represents a concatenation of given and family names. For 
    // example `John Smith`. 
    DisplayName string `json:"displayName,omitempty"`
    // All the emails associated with the Databricks user. 
    Emails []Email `json:"emails,omitempty"`
    // Databricks user ID. 
    Id string `json:"id,omitempty"`
    
    Name *Name `json:"name,omitempty"`
    // SCIM schema used for the user object. 
    Schemas []CoreUser `json:"schemas,omitempty"`
    // Email address of the Databricks user. 
    UserName string `json:"userName,omitempty"`
}

