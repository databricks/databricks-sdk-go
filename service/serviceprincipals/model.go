package serviceprincipals


// 

type CoreServicePrincipal string

// 
const CoreServicePrincipalSchema CoreServicePrincipal = "urn:ietf:params:scim:schemas:core:2.0:ServicePrincipal"



// 
type DeleteServicePrincipalRequest struct {
	
	// Unique ID for a service principal in the &lt;Workspace&gt;.
	ServicePrincipalId string `path:"service_principal_id"`
	
}


// 
type DisplayNamePatchOp struct {
	
	// Type of patch operation.
	Op any /* {{op Type of patch operation.} false 0xc0002b27e0 true false false} */ `json:"op,omitempty"`
	
	// Field to update.
	Path any /* {{path Field to update.} false 0xc0002b2660 true false false} */ `json:"path,omitempty"`
	
	// 
	Value [] `json:"value,omitempty"`
	
}


// 
type FetchServicePrincipalRequest struct {
	
	// 
	ServicePrincipalId string `path:"service_principal_id"`
	
}


// 

type List string

// 
const ListResponseMessage List = "urn:ietf:params:scim:api:messages:2.0:ListResponse"



// 
type ListServicePrincipalsRequest struct {
	
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
	SortOrder any /* {{sortOrder The order to sort the results.} false 0xc0002adce0 false false true} */ `query:"sortOrder,omitempty"`
	
	// Specifies the index of the first result. First item is number 1.
	StartIndex int `query:"startIndex,omitempty"`
	
}


// 
type ListServicePrincipalsResponse struct {
	
	// User objects returned in the response.
	Resources []ServicePrincipal `json:"Resources,omitempty"`
	
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
type ListServicePrincipalsResponse struct {
	
	// User objects returned in the response.
	Resources []ServicePrincipal `json:"Resources,omitempty"`
	
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
type NewServicePrincipalRequest struct {
	
	// String that represents a human-readable group name
	DisplayName string `json:"displayName,omitempty"`
	
	// SCIM schema used for the group object.
	Schemas []CoreServicePrincipal `json:"schemas,omitempty"`
	
}


// 
type NewServicePrincipalRequest struct {
	
	// String that represents a human-readable group name
	DisplayName string `json:"displayName,omitempty"`
	
	// SCIM schema used for the group object.
	Schemas []CoreServicePrincipal `json:"schemas,omitempty"`
	
}


// 
type ServicePrincipal struct {
	
	// UUID relating to the service principal
	ApplicationId string `json:"applicationId,omitempty"`
	
	// String that represents a human-readable name
	DisplayName string `json:"displayName,omitempty"`
	
	// Databricks service principal ID.
	Id string `json:"id,omitempty"`
	
	// SCIM schema used for the user object.
	Schemas []CoreServicePrincipal `json:"schemas,omitempty"`
	
}


// 

type Update string

// 
const PatchOpMessage Update = "urn:ietf:params:scim:api:messages:2.0:PatchOp"



// 
type UpdateDisplayNameRequest struct {
	
	// Patch operations to be performed on the user object.
	Operations []DisplayNamePatchOp `json:"Operations,omitempty"`
	
	// SCIM PatchOp schema used to update user information.
	Schemas []Update `json:"schemas,omitempty"`
	
	// Unique ID for a service principal in the &lt;Workspace&gt;.
	ServicePrincipalId string `path:"service_principal_id"`
	
}

