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
	"golang.org/x/oauth2"
)

func getTestServer(host string, token string) *httptest.Server {
	counter := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get(MetadataServiceVersionHeader) != MetadataServiceVersion {
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(serverResponse{
			Host: host,
			Token: oauth2.Token{
				AccessToken: fmt.Sprintf("%s-%d", token, counter),
				Expiry:      time.Now().Add(1 * time.Second),
				TokenType:   "Bearer",
			},
		})

		counter++
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
