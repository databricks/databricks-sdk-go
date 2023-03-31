package config

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func getTestServer(host string, token string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Databricks-Metadata-Version") != MetadataServiceVersion {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{
			"host": "%s",
			"token": {
				"access_token": "%s",
				"expiry": "%s",
				"token_type": "Bearer"	
			}	
		}`, host, token, time.Now().Add(30*time.Second).Format(time.RFC3339))))
	}))
	return ts
}

func TestAuthServerSetHost(t *testing.T) {
	host := "ZZZ"
	token := "XXX"

	ts := getTestServer(host, token)
	defer ts.Close()

	sc := MetadataServiceCredentials{}
	config := &Config{
		MetadataServiceURL: ts.URL,
	}
	authProvider, err := sc.Configure(context.Background(), config)
	require.NoError(t, err)
	require.NotEmpty(t, authProvider)

	require.Equal(t, config.Host, host)
}

func TestAuthServerAuthorize(t *testing.T) {
	host := "ZZZ"
	token := "XXX"

	ts := getTestServer(host, token)
	defer ts.Close()

	sc := MetadataServiceCredentials{}
	authProvider, err := sc.Configure(context.Background(), &Config{
		MetadataServiceURL: ts.URL,
		Host:               host,
	})
	require.NoError(t, err)
	require.NotEmpty(t, authProvider)

	req := &http.Request{
		Header: http.Header{},
	}

	err = authProvider(req)
	require.NoError(t, err)

	require.Equal(t, req.Header.Get("Authorization"), fmt.Sprintf("Bearer %s", token))
}
