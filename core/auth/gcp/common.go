package gcp

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

func idTokenSource(ctx context.Context, host, serviceAccount string, opts ...option.ClientOption) (oauth2.TokenSource, error) {
	ts, err := impersonate.IDTokenSource(ctx, impersonate.IDTokenConfig{
		Audience:        host,
		TargetPrincipal: serviceAccount,
		IncludeEmail:    true,
	}, opts...)
	if err != nil {
		err = fmt.Errorf("could not obtain OIDC token. %w Running 'gcloud auth application-default login' may help", err)
		return nil, err
	}
	return ts, err
}
