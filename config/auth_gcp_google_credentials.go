package config

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/databricks/databricks-sdk-go/logger"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
)

type GoogleCredentials struct {
}

func (c GoogleCredentials) Name() string {
	return "google-credentials"
}

func (c GoogleCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if cfg.GoogleCredentials == "" || !cfg.IsGcp() {
		return nil, nil
	}
	json, err := readCredentials(cfg.GoogleCredentials)
	if err != nil {
		err = fmt.Errorf("could not read GoogleCredentials. "+
			"Make sure the file exists, or the JSON content is valid: %w", err)
		return nil, err
	}
	inner, err := idtoken.NewTokenSource(ctx, cfg.Host, option.WithCredentialsJSON(json))
	if err != nil {
		return nil, fmt.Errorf("could not obtain OIDC token from JSON: %w", err)
	}
	// Obtain token source for creating Google Cloud Platform token.
	creds, err := google.CredentialsFromJSON(ctx, json,
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/compute")
	if err != nil {
		return nil, fmt.Errorf("could not obtain OAuth2 token from JSON: %w", err)
	}
	logger.Infof(ctx, "Using Google Credentials")
	return serviceToServiceVisitor(inner, creds.TokenSource, "X-Databricks-GCP-SA-Access-Token"), nil
}

// Reads credentials as JSON. Credentials can be either a path to JSON file,
// or actual JSON string.
func readCredentials(credentials string) ([]byte, error) {
	// Try to read credentials as file path.
	if _, err := os.Stat(credentials); err == nil {
		jsonContents, err := ioutil.ReadFile(credentials)
		if err != nil {
			return jsonContents, err
		}
		return jsonContents, nil
	}
	// Assume that credential is actually JSON string.
	return []byte(credentials), nil
}
