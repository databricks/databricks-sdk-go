package config

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/databricks/useragent"
	"github.com/stretchr/testify/assert"
)

func TestAuthInUserAgentVisitorDefault(t *testing.T) {
	request := &http.Request{}
	err := authInUserAgentVisitor("PAT")(request)
	assert.NoError(t, err)
	// tests may be developed and run on different versions of different things
	uac := strings.Split(useragent.FromContext(request.Context()), " ")
	assert.Contains(t, uac, "auth/PAT")
}

func TestAuthInUserAgentVisitorCustom(t *testing.T) {
	request := &http.Request{}
	ctx := context.Background()
	ctx = context.WithValue(ctx, AuthType{}, "oath")
	request = request.WithContext(ctx)
	err := authInUserAgentVisitor("PAT")(request)
	assert.NoError(t, err)
	// tests may be developed and run on different versions of different things
	uac := strings.Split(useragent.FromContext(request.Context()), " ")
	assert.NotContains(t, uac, "auth/PAT")
	assert.Contains(t, uac, "auth/oath")
}
