// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package oauth2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/client"
)

// unexported type that holds implementations of just CustomAppIntegration API methods
type customAppIntegrationImpl struct {
	client *client.DatabricksClient
}

func (a *customAppIntegrationImpl) Create(ctx context.Context, request CreateCustomAppIntegration) (*CreateCustomAppIntegrationOutput, error) {
	var createCustomAppIntegrationOutput CreateCustomAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &createCustomAppIntegrationOutput)
	return &createCustomAppIntegrationOutput, err
}

func (a *customAppIntegrationImpl) Delete(ctx context.Context, request DeleteCustomAppIntegrationRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *customAppIntegrationImpl) Get(ctx context.Context, request GetCustomAppIntegrationRequest) (*GetCustomAppIntegrationOutput, error) {
	var getCustomAppIntegrationOutput GetCustomAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getCustomAppIntegrationOutput)
	return &getCustomAppIntegrationOutput, err
}

func (a *customAppIntegrationImpl) List(ctx context.Context) (*GetCustomAppIntegrationsOutput, error) {
	var getCustomAppIntegrationsOutput GetCustomAppIntegrationsOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getCustomAppIntegrationsOutput)
	return &getCustomAppIntegrationsOutput, err
}

func (a *customAppIntegrationImpl) Update(ctx context.Context, request UpdateCustomAppIntegration) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/custom-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}

// unexported type that holds implementations of just OAuthEnrollment API methods
type oAuthEnrollmentImpl struct {
	client *client.DatabricksClient
}

func (a *oAuthEnrollmentImpl) Create(ctx context.Context, request CreateOAuthEnrollment) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/enrollment", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, nil)
	return err
}

func (a *oAuthEnrollmentImpl) Get(ctx context.Context) (*OAuthEnrollmentStatus, error) {
	var oAuthEnrollmentStatus OAuthEnrollmentStatus
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/enrollment", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &oAuthEnrollmentStatus)
	return &oAuthEnrollmentStatus, err
}

// unexported type that holds implementations of just PublishedAppIntegration API methods
type publishedAppIntegrationImpl struct {
	client *client.DatabricksClient
}

func (a *publishedAppIntegrationImpl) Create(ctx context.Context, request CreatePublishedAppIntegration) (*CreatePublishedAppIntegrationOutput, error) {
	var createPublishedAppIntegrationOutput CreatePublishedAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodPost, path, request, &createPublishedAppIntegrationOutput)
	return &createPublishedAppIntegrationOutput, err
}

func (a *publishedAppIntegrationImpl) Delete(ctx context.Context, request DeletePublishedAppIntegrationRequest) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	err := a.client.Do(ctx, http.MethodDelete, path, request, nil)
	return err
}

func (a *publishedAppIntegrationImpl) Get(ctx context.Context, request GetPublishedAppIntegrationRequest) (*GetPublishedAppIntegrationOutput, error) {
	var getPublishedAppIntegrationOutput GetPublishedAppIntegrationOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	err := a.client.Do(ctx, http.MethodGet, path, request, &getPublishedAppIntegrationOutput)
	return &getPublishedAppIntegrationOutput, err
}

func (a *publishedAppIntegrationImpl) List(ctx context.Context) (*GetPublishedAppIntegrationsOutput, error) {
	var getPublishedAppIntegrationsOutput GetPublishedAppIntegrationsOutput
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations", a.client.ConfiguredAccountID())
	err := a.client.Do(ctx, http.MethodGet, path, nil, &getPublishedAppIntegrationsOutput)
	return &getPublishedAppIntegrationsOutput, err
}

func (a *publishedAppIntegrationImpl) Update(ctx context.Context, request UpdatePublishedAppIntegration) error {
	path := fmt.Sprintf("/api/2.0/accounts/%v/oauth2/published-app-integrations/%v", a.client.ConfiguredAccountID(), request.IntegrationId)
	err := a.client.Do(ctx, http.MethodPatch, path, request, nil)
	return err
}
