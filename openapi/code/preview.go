package code

import (
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi"
)

func isPrivatePreview(n *openapi.Node) bool {
	return strings.ToLower(n.Preview) == "private"
}

func isPublicPreview(n *openapi.Node) bool {
	return strings.ToLower(n.Preview) == "public"
}
