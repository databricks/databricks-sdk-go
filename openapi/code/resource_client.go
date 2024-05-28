package code

import "github.com/databricks/databricks-sdk-go/openapi"

type ResourceClient struct {
	Named

	PathStyle           openapi.PathStyle
	IsAccounts          bool
	Package             *Package
	methods             map[string]*Method
	ByPathParamsMethods []*Shortcut
	tag                 *openapi.Tag
}
