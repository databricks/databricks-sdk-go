// Package dataplane is an experimental package that provides a token source to
// directly access Databricks data plane.
package dataplane

import (
	"context"
	"sync"

	"github.com/databricks/databricks-sdk-go/config/experimental/auth"
	"golang.org/x/oauth2"
)

// OAuthClient is an interface for Databricks OAuth client.
type OAuthClient interface {
	GetOAuthToken(ctx context.Context, authDetails string, t *oauth2.Token) (*oauth2.Token, error)
}

// EndpointTokenSource is anything that returns tokens given a data plane
// endpoint and authentication details.
type EndpointTokenSource interface {
	Token(ctx context.Context, endpoint string, authDetails string) (*oauth2.Token, error)
}

// NewEndpointTokenSource returns a new EndpointTokenSource that uses the given
// OAuthClient and control plane TokenSource.
func NewEndpointTokenSource(c OAuthClient, cpts auth.TokenSource) *dataPlaneTokenSource {
	return &dataPlaneTokenSource{
		client: c,
		cpts: auth.NewCachedTokenSource(
			cpts,
			auth.WithAsyncRefresh(false), // TODO: Enable async refreshes once the feature is stable.
		),
		tss: make(map[tsKey]auth.TokenSource),
	}
}

type tsKey struct {
	endpoint    string
	authDetails string
}

// dataPlaneTokenSource implements the EndpointTokenSource interface.
//
// For a given endpoint and authentication details, it uses the control plane
// TokenSource to retrieve the control plane token, that is then used to
// retrieve the data plane token through the OAuthClient.
//
// Each token source is cached to avoid unnecessary token requests.
type dataPlaneTokenSource struct {
	client OAuthClient
	cpts   auth.TokenSource
	tss    map[tsKey]auth.TokenSource
	mu     sync.Mutex
}

// Token returns a token for the given endpoint and authentication details.
func (dpts *dataPlaneTokenSource) Token(ctx context.Context, endpoint string, authDetails string) (*oauth2.Token, error) {
	dpts.mu.Lock()
	defer dpts.mu.Unlock()

	key := tsKey{endpoint: endpoint, authDetails: authDetails}

	t, ok := dpts.tss[key]
	if !ok {
		t = auth.NewCachedTokenSource(
			&tokenSource{
				client:      dpts.client,
				cpts:        dpts.cpts,
				authDetails: authDetails,
			},
			auth.WithAsyncRefresh(false), // TODO: Enable async refresh once the feature is stable.
		)
		dpts.tss[key] = t
	}

	return t.Token(ctx)
}

type tokenSource struct {
	client      OAuthClient
	cpts        auth.TokenSource
	authDetails string
}

const JWTGrantType = "urn:ietf:params:oauth:grant-type:jwt-bearer"

func (dpts *tokenSource) Token(ctx context.Context) (*oauth2.Token, error) {
	innerToken, err := dpts.cpts.Token(ctx)
	if err != nil {
		return nil, err
	}
	return dpts.client.GetOAuthToken(ctx, dpts.authDetails, innerToken)
}
