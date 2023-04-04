package config

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func getTestServer(host string, token string) *httptest.Server {
	counter := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(MetadataServiceVersionHeader) != MetadataServiceVersion {
			w.WriteHeader(http.StatusBadRequest)
		}
		if r.Header.Get(MetadataServiceHostHeader) != host {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(serverResponse{
			AccessToken: fmt.Sprintf("%s-%d", token, counter),
			ExpiresOn:   json.Number(fmt.Sprintf("%d", time.Now().Add(1*time.Second).Unix())),
			TokenType:   "Bearer",
		})

		counter++
	}))
	return ts
}

func TestAuthServerCheckHost(t *testing.T) {
	host := "ZZZ"
	token := "XXX"

	ts := getTestServer(host, token)
	defer ts.Close()

	sc := MetadataServiceCredentials{}
	config := &Config{
		Host:               "YYY",
		MetadataServiceURL: ts.URL,
	}
	_, err := sc.Configure(context.Background(), config)
	require.Empty(t, err)
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

	require.Equal(t, fmt.Sprintf("Bearer %s-1", token), req.Header.Get("Authorization"))
}

func TestAuthServerRefresh(t *testing.T) {
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

	require.Equal(t, fmt.Sprintf("Bearer %s-1", token), req.Header.Get("Authorization"))

	req = &http.Request{
		Header: http.Header{},
	}
	err = authProvider(req)
	require.NoError(t, err)

	require.Equal(t, fmt.Sprintf("Bearer %s-2", token), req.Header.Get("Authorization"))
}
