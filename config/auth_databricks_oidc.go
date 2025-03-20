package config

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/databricks/databricks-sdk-go/config/credentials"
	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

type DatabricksOIDCCredentials struct{}

type loggingTransport struct {
	Transport http.RoundTripper
}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Log the request
	dump, err := httputil.DumpRequestOut(req, true)
	if err == nil {
		fmt.Println("HTTP Request:")
		fmt.Println(string(dump))
	}

	// Send the request
	resp, err := t.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Log the response
	dumpResp, err := httputil.DumpResponse(resp, true)
	if err == nil {
		fmt.Println("HTTP Response:")
		fmt.Println(string(dumpResp))
	}

	return resp, nil
}

func print(ctx context.Context, token string) {

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid JWT format")
		return
	}
	// Decode the payload (second part of the JWT)
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		logger.Debugf(ctx, "Error decoding JWT payload: %v", err)
		return
	}

	// Pretty-print JSON
	var prettyPayload map[string]interface{}
	if err := json.Unmarshal(payload, &prettyPayload); err != nil {
		logger.Debugf(ctx, "Error parsing JSON: %v", err)
		return
	}

	prettyJSON, _ := json.MarshalIndent(prettyPayload, "", "  ")
	logger.Debugf(ctx, string(prettyJSON))
}

type databricksOIDCTokenSource struct {
	ctx                 context.Context
	jwtTokenSupplicer   *GithubOIDCTokenSupplier
	workspaceResourceId string
	azureTenantId       string
}

func (d *databricksOIDCTokenSource) Token() *oauth2.Token {
	return nil
}

// Configure implements CredentialsStrategy.
func (d DatabricksOIDCCredentials) Configure(ctx context.Context, cfg *Config) (credentials.CredentialsProvider, error) {
	if cfg.Host == "" || cfg.ClientID == "" || cfg.TokenAudience == "" {
		return nil, nil
	}

	// Get the OIDC token from the environment.
	supplier := GithubOIDCTokenSupplier{
		cfg: cfg,
	}
	token, err := supplier.GetOIDCToken(ctx, cfg.TokenAudience)
	if err != nil {
		return nil, err
	}
	if token == "" {
		logger.Debugf(ctx, "No OIDC token found")
		return nil, nil
	}

	endpoints, err := oidcEndpoints(ctx, cfg)
	if err != nil {
		return nil, err
	}
	logger.Debugf(ctx, "Getting tokken for client %s", cfg.ClientID)

	tsConfig := clientcredentials.Config{
		ClientID:  cfg.ClientID,
		AuthStyle: oauth2.AuthStyleInParams,
		TokenURL:  endpoints.TokenEndpoint,
		Scopes:    []string{"all-apis"},
		EndpointParams: url.Values{
			"subject_token_type": {"urn:ietf:params:oauth:token-type:jwt"},
			"subject_token":      {token},
			"grant_type":         {"urn:ietf:params:oauth:grant-type:token-exchange"},
		},
	}
	client := &http.Client{Transport: &loggingTransport{Transport: http.DefaultTransport}}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, client)

	// Request the token
	_, err = tsConfig.Token(ctx)
	if err != nil {
		return nil, err
	}
	ts := tsConfig.TokenSource(ctx)
	visitor := refreshableVisitor(ts)
	return credentials.NewOAuthCredentialsProvider(visitor, ts.Token), nil
}

// Name implements CredentialsStrategy.
func (d DatabricksOIDCCredentials) Name() string {
	return "databricks-oidc"
}
