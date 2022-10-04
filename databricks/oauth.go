package databricks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

func (cfg *Config) OAuthConfig() (*oauth2.Config, error) {
	// if len(cfg.Scopes) == 0 {
	// 	cfg.Scopes = []string{"all-apis"}
	// }
	endpoints, err := cfg.getEndpoints()
	if err != nil {
		return nil, err
	}
	return &oauth2.Config{
		ClientID: cfg.ClientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:   endpoints.AuthorizationEndpoint,
			TokenURL:  endpoints.TokenEndpoint,
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: cfg.RedirectURL,
		Scopes:      cfg.Scopes,
	}, nil
}

type oauthAuthorizationServer struct {
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
}

func (cfg *Config) getEndpoints() (*oauthAuthorizationServer, error) {
	err := cfg.fixHostIfNeeded()
	if err != nil {
		return nil, fmt.Errorf("host: %w", err)
	}
	oidc := fmt.Sprintf("%s/oidc/.well-known/oauth-authorization-server", cfg.Host)
	oidcResponse, err := http.Get(oidc)
	if err != nil {
		return nil, fmt.Errorf("fetch .well-known: %w", err)
	}
	if oidcResponse.Body == nil {
		return nil, fmt.Errorf("fetch .well-known: empty body")
	}
	defer oidcResponse.Body.Close()
	raw, err := io.ReadAll(oidcResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("read .well-known: %w", err)
	}
	var oauthEndpoints oauthAuthorizationServer
	err = json.Unmarshal(raw, &oauthEndpoints)
	if err != nil {
		return nil, fmt.Errorf("parse .well-known: %w", err)
	}
	return &oauthEndpoints, nil
}
