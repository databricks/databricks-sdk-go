package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/credentials"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/oauth2"
)

var mockTokenProvider = func() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: "dummy",
	}, nil
}

var mockInfoGetter = func() (*DataPlaneInfo, error) {
	return &DataPlaneInfo{
		EndpointUrl:          "endpoint_url",
		AuthorizationDetails: "details",
	}, nil
}

func TestGetOAuthTokenForDataPlane(t *testing.T) {
	tokenResponse := &credentials.OAuthToken{
		AccessToken: "dummyDataPlane",
		ExpiresIn:   3600,
		Scope:       "scope",
		TokenType:   "Bearer",
	}
	marshalled, err := json.Marshal(tokenResponse)
	require.NoError(t, err)
	c := NewApiClient(ClientConfig{
		Transport: hc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(string(marshalled))),
				Request:    r,
			}, nil
		}),
	})
	token, err := c.GetOAuthTokenForDataPlane("service", "path", mockInfoGetter, mockTokenProvider)
	assert.NoError(t, err)
	assert.Equal(t, "dummyDataPlane", token.AccessToken)
	cachedData, ok := c.dataPlaneCache[DataPlaneInfoKey{
		ServiceName: "service",
		Path:        "path",
	}]
	assert.True(t, ok)
	assert.Equal(t, "details", cachedData.DataPlaneInfo.AuthorizationDetails)
	assert.Equal(t, "endpoint_url", cachedData.DataPlaneInfo.EndpointUrl)
	assert.Equal(t, "dummyDataPlane", cachedData.Token.AccessToken)
	assert.False(t, cachedData.Expired())
}
