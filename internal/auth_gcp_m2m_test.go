package internal

import (
	"net/http"
	"testing"

	"github.com/databricks/databricks-sdk-go"
	"github.com/databricks/databricks-sdk-go/service/iam"
	"github.com/stretchr/testify/require"
)

// headerCapturingTransport records the headers of the most recent outgoing
// request and delegates to an inner RoundTripper.
type headerCapturingTransport struct {
	inner http.RoundTripper
	last  http.Header
}

func (t *headerCapturingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.last = req.Header.Clone()
	return t.inner.RoundTrip(req)
}

// TestMwsAccAccountOAuthM2MGcpAuth verifies the oauth-m2m-gcp strategy end to
// end on a GCP account: a Databricks OAuth service-principal token authenticates
// the account API call (Authorization header), and a Google Cloud access token
// is attached in the X-Databricks-GCP-SA-Access-Token passthrough header on the
// real outgoing request.
//
// Gated on DATABRICKS_GOOGLE_SERVICE_ACCOUNT, which is only set in GCP CI, so
// the test is skipped on other clouds.
func TestMwsAccAccountOAuthM2MGcpAuth(t *testing.T) {
	ctx, _ := accountTest(t)
	t.Log(GetEnvOrSkipTest(t, "CLOUD_ENV"))

	host := GetEnvOrSkipTest(t, "DATABRICKS_HOST")
	accountID := GetEnvOrSkipTest(t, "DATABRICKS_ACCOUNT_ID")
	clientID := GetEnvOrSkipTest(t, "TEST_DATABRICKS_CLIENT_ID")
	clientSecret := GetEnvOrSkipTest(t, "TEST_DATABRICKS_CLIENT_SECRET")
	// The Google service account whose cloud-platform access token is passed
	// through. Only present in GCP CI, so this skips elsewhere.
	googleServiceAccount := GetEnvOrSkipTest(t, "DATABRICKS_GOOGLE_SERVICE_ACCOUNT")

	recorder := &headerCapturingTransport{inner: http.DefaultTransport}

	accCfg := &databricks.Config{
		Host:                 host,
		AccountID:            accountID,
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		GoogleServiceAccount: googleServiceAccount,
		AuthType:             "oauth-m2m-gcp",
		HTTPTransport:        recorder,
	}

	accClient, err := databricks.NewAccountClient(accCfg)
	require.NoError(t, err)

	// A real account API call proves the Databricks OAuth identity authenticates
	// and that minting the Google access token does not break the request.
	it := accClient.ServicePrincipals.List(ctx, iam.ListAccountServicePrincipalsRequest{})
	_, err = it.Next(ctx)
	require.NoError(t, err)

	// Both credentials must have been attached to the actual outgoing request.
	require.NotNil(t, recorder.last, "no request was captured")
	require.Contains(t, recorder.last.Get("Authorization"), "Bearer ")
	require.NotEmpty(t, recorder.last.Get("X-Databricks-GCP-SA-Access-Token"),
		"expected the GCP SA access token passthrough header to be set")
}
