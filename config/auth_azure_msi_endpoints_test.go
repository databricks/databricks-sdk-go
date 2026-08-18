package config

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/databricks/databricks-sdk-go/httpclient"
	"github.com/databricks/databricks-sdk-go/httpclient/fixtures"
	"github.com/google/go-cmp/cmp"
)

const (
	testAzureMSIResource = "https://resource.test/"
	testAzureMSIClientID = "test-client-id"
)

func TestAzureMsiTokenSource_endpointSelection(t *testing.T) {
	fixedNow := time.Date(2026, time.August, 17, 12, 0, 0, 0, time.UTC)
	fixedExpiry := time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC)
	federatedTokenFile := filepath.Join(t.TempDir(), "federated-token")
	if err := os.WriteFile(federatedTokenFile, []byte("federated-token"), 0o600); err != nil {
		t.Fatalf("write federated token: %v", err)
	}

	testCases := []struct {
		name       string
		env        map[string]string
		clientID   string
		fixture    fixtures.HTTPFixture
		response   any
		wantToken  string
		wantExpiry time.Time
	}{
		{
			name: "service fabric",
			env: map[string]string{
				azureIdentityEndpointEnv:         "http://service-fabric.test/service-fabric/token",
				azureIdentityHeaderEnv:           "service-fabric-secret",
				azureIdentityServerThumbprintEnv: "thumbprint",
			},
			fixture: fixtures.HTTPFixture{
				Method:   http.MethodGet,
				Resource: "/service-fabric/token?api-version=2019-07-01-preview&resource=https%3A%2F%2Fresource.test%2F",
				ExpectedHeaders: map[string]string{
					"Accept": "application/json",
					"Secret": "service-fabric-secret",
				},
			},
			response:   azureMSITokenResponse("service-fabric-token", fixedExpiry),
			wantToken:  "service-fabric-token",
			wantExpiry: fixedExpiry,
		},
		{
			name: "app service takes precedence over legacy endpoint",
			env: map[string]string{
				azureIdentityEndpointEnv:   "http://app-service.test/app-service/token",
				azureIdentityHeaderEnv:     "app-service-secret",
				azureMSIEndpointEnv:        "http://azure-ml.test/should-not-be-used",
				azureMSISecretEnv:          "legacy-secret",
				azureAuthorityHostEnv:      "http://workload.test/should-not-be-used",
				azureTenantIDEnv:           "tenant",
				azureFederatedTokenFileEnv: federatedTokenFile,
			},
			clientID: testAzureMSIClientID,
			fixture: fixtures.HTTPFixture{
				Method:   http.MethodGet,
				Resource: "/app-service/token?api-version=2019-08-01&client_id=test-client-id&resource=https%3A%2F%2Fresource.test%2F",
				ExpectedHeaders: map[string]string{
					"Accept":            "application/json",
					"X-Identity-Header": "app-service-secret",
				},
			},
			response:   azureMSITokenResponse("app-service-token", fixedExpiry),
			wantToken:  "app-service-token",
			wantExpiry: fixedExpiry,
		},
		{
			name: "azure arc",
			env: map[string]string{
				azureIdentityEndpointEnv: "http://arc.test/arc/token",
				azureIMDSEndpointEnv:     "http://arc.test",
			},
			fixture: fixtures.HTTPFixture{
				Method:   http.MethodGet,
				Resource: "/arc/token?api-version=2020-06-01&resource=https%3A%2F%2Fresource.test%2F",
				ExpectedHeaders: map[string]string{
					"Accept":   "application/json",
					"Metadata": "true",
				},
			},
			response:   azureMSITokenResponse("azure-arc-token", fixedExpiry),
			wantToken:  "azure-arc-token",
			wantExpiry: fixedExpiry,
		},
		{
			name: "azure ml",
			env: map[string]string{
				azureMSIEndpointEnv:              "http://azure-ml.test/azure-ml/token",
				azureMSISecretEnv:                "azure-ml-secret",
				azureAuthorityHostEnv:            "http://workload.test/should-not-be-used",
				azureTenantIDEnv:                 "tenant",
				azureFederatedTokenFileEnv:       federatedTokenFile,
				azurePodIdentityAuthorityHostEnv: "http://pod.test/should-not-be-used",
			},
			clientID: testAzureMSIClientID,
			fixture: fixtures.HTTPFixture{
				Method:   http.MethodGet,
				Resource: "/azure-ml/token?api-version=2017-09-01&clientid=test-client-id&resource=https%3A%2F%2Fresource.test%2F",
				ExpectedHeaders: map[string]string{
					"Accept": "application/json",
					"Secret": "azure-ml-secret",
				},
			},
			response: map[string]any{
				"access_token": "azure-ml-token",
				"expires_on":   "01/01/2030 00:00:00 +00:00",
				"token_type":   "Bearer",
			},
			wantToken:  "azure-ml-token",
			wantExpiry: fixedExpiry,
		},
		{
			name: "cloud shell",
			env: map[string]string{
				azureMSIEndpointEnv: "http://cloud-shell.test/cloud-shell/token",
			},
			fixture: fixtures.HTTPFixture{
				Method:   http.MethodPost,
				Resource: "/cloud-shell/token",
				ExpectedHeaders: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/x-www-form-urlencoded",
					"Metadata":     "true",
				},
				ExpectedRequest: url.Values{
					"resource": []string{testAzureMSIResource},
				},
			},
			response:   azureMSITokenResponse("cloud-shell-token", fixedExpiry),
			wantToken:  "cloud-shell-token",
			wantExpiry: fixedExpiry,
		},
		{
			name: "workload identity",
			env: map[string]string{
				azureAuthorityHostEnv:      "http://workload.test/authority/",
				azureClientIDEnv:           testAzureMSIClientID,
				azureTenantIDEnv:           "test-tenant",
				azureFederatedTokenFileEnv: federatedTokenFile,
			},
			fixture: fixtures.HTTPFixture{
				Method:   http.MethodPost,
				Resource: "/authority/test-tenant/oauth2/v2.0/token",
				ExpectedHeaders: map[string]string{
					"Accept":       "application/json",
					"Content-Type": "application/x-www-form-urlencoded",
				},
				ExpectedRequest: url.Values{
					"client_assertion":      []string{"federated-token"},
					"client_assertion_type": []string{"urn:ietf:params:oauth:client-assertion-type:jwt-bearer"},
					"client_id":             []string{testAzureMSIClientID},
					"grant_type":            []string{"client_credentials"},
					"scope":                 []string{"https://resource.test/.default"},
				},
			},
			response: map[string]any{
				"access_token": "workload-token",
				"expires_in":   3600,
			},
			wantToken:  "workload-token",
			wantExpiry: fixedNow.Add(time.Hour),
		},
		{
			name:     "imds fallback",
			env:      map[string]string{},
			clientID: testAzureMSIClientID,
			fixture: fixtures.HTTPFixture{
				Method:   http.MethodGet,
				Resource: "/metadata/identity/oauth2/token?api-version=2018-02-01&client_id=test-client-id&resource=https%3A%2F%2Fresource.test%2F",
				ExpectedHeaders: map[string]string{
					"Accept":   "application/json",
					"Metadata": "true",
				},
			},
			response:   azureMSITokenResponse("imds-token", fixedExpiry),
			wantToken:  "imds-token",
			wantExpiry: fixedExpiry,
		},
		{
			name: "pod identity authority overrides imds authority",
			env: map[string]string{
				azureAuthorityHostEnv:            "http://incomplete-workload.test",
				azureTenantIDEnv:                 "tenant",
				azurePodIdentityAuthorityHostEnv: "http://pod.test/custom/",
			},
			fixture: fixtures.HTTPFixture{
				Method:   http.MethodGet,
				Resource: "/custom/metadata/identity/oauth2/token?api-version=2018-02-01&resource=https%3A%2F%2Fresource.test%2F",
				ExpectedHeaders: map[string]string{
					"Accept":   "application/json",
					"Metadata": "true",
				},
			},
			response:   azureMSITokenResponse("pod-identity-token", fixedExpiry),
			wantToken:  "pod-identity-token",
			wantExpiry: fixedExpiry,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			withMockEnv(t, tc.env)
			fixture := tc.fixture
			fixture.Response = tc.response
			httpClient := httpclient.NewApiClient(httpclient.ClientConfig{
				Transport: fixtures.SliceTransport{fixture},
			})
			tokenSource := newAzureMSITokenSource(httpClient, testAzureMSIResource, tc.clientID)
			tokenSource.now = func() time.Time { return fixedNow }

			got, gotErr := tokenSource.Token(context.Background())
			if gotErr != nil {
				t.Fatalf("Token() failed: %v", gotErr)
			}
			if got.AccessToken != tc.wantToken {
				t.Errorf("Token().AccessToken = %q, want %q", got.AccessToken, tc.wantToken)
			}
			if got.TokenType != "Bearer" {
				t.Errorf("Token().TokenType = %q, want %q", got.TokenType, "Bearer")
			}
			if diff := cmp.Diff(tc.wantExpiry, got.Expiry); diff != "" {
				t.Errorf("Token().Expiry mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestAzureMsiTokenSource_invalidEnvironment(t *testing.T) {
	testCases := []struct {
		name     string
		env      map[string]string
		clientID string
		wantErr  string
	}{
		{
			name: "identity endpoint is incomplete",
			env: map[string]string{
				azureIdentityEndpointEnv: "http://identity.test/token",
				azureMSIEndpointEnv:      "http://cloud-shell.test/token",
			},
			wantErr: "no managed identity endpoint found: IDENTITY_ENDPOINT requires IDENTITY_HEADER or IMDS_ENDPOINT",
		},
		{
			name: "service fabric rejects a client ID",
			env: map[string]string{
				azureIdentityEndpointEnv:         "http://service-fabric.test/token",
				azureIdentityHeaderEnv:           "secret",
				azureIdentityServerThumbprintEnv: "thumbprint",
			},
			clientID: testAzureMSIClientID,
			wantErr:  "azure_client_id is not supported by \"Service Fabric\" managed identity",
		},
		{
			name: "azure arc rejects a client ID",
			env: map[string]string{
				azureIdentityEndpointEnv: "http://arc.test/token",
				azureIMDSEndpointEnv:     "http://arc.test",
			},
			clientID: testAzureMSIClientID,
			wantErr:  "azure_client_id is not supported by \"Azure Arc\" managed identity",
		},
		{
			name: "cloud shell rejects a client ID",
			env: map[string]string{
				azureMSIEndpointEnv: "http://cloud-shell.test/token",
			},
			clientID: testAzureMSIClientID,
			wantErr:  "azure_client_id is not supported by \"Cloud Shell\" managed identity",
		},
		{
			name: "workload identity requires a client ID",
			env: map[string]string{
				azureAuthorityHostEnv:      "http://workload.test",
				azureTenantIDEnv:           "tenant",
				azureFederatedTokenFileEnv: "token-file",
			},
			wantErr: "azure_client_id is required for workload identity",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			withMockEnv(t, tc.env)
			tokenSource := newAzureMSITokenSource(nil, testAzureMSIResource, tc.clientID)
			got, gotErr := tokenSource.Token(context.Background())
			if got != nil {
				t.Errorf("Token() = %#v, want nil", got)
			}
			if gotErr == nil || gotErr.Error() != tc.wantErr {
				t.Errorf("Token() error = %q, want %q", gotErr, tc.wantErr)
			}
		})
	}
}

func TestAzureMsiTokenSource_azureArcChallenge(t *testing.T) {
	withMockEnv(t, map[string]string{
		azureIdentityEndpointEnv: "http://arc.test/token",
		azureIMDSEndpointEnv:     "http://arc.test",
	})
	fixedExpiry := time.Date(2030, time.January, 1, 0, 0, 0, 0, time.UTC)
	requestResource := "/token?api-version=2020-06-01&resource=https%3A%2F%2Fresource.test%2F"
	httpClient := httpclient.NewApiClient(httpclient.ClientConfig{
		Transport: fixtures.SliceTransport{
			{
				Method:   http.MethodGet,
				Resource: requestResource,
				ExpectedHeaders: map[string]string{
					"Accept":   "application/json",
					"Metadata": "true",
				},
				Status: http.StatusUnauthorized,
				ResponseHeaders: map[string][]string{
					"Www-Authenticate": []string{"Basic realm=/var/opt/azcmagent/tokens/identity.key"},
				},
			},
			{
				Method:   http.MethodGet,
				Resource: requestResource,
				ExpectedHeaders: map[string]string{
					"Accept":        "application/json",
					"Authorization": "Basic arc-secret",
					"Metadata":      "true",
				},
				Response: azureMSITokenResponse("azure-arc-token", fixedExpiry),
			},
		},
	})
	tokenSource := newAzureMSITokenSource(httpClient, testAzureMSIResource, "")
	tokenSource.readAzureArcSecret = func(got string) (string, error) {
		want := "/var/opt/azcmagent/tokens/identity.key"
		if got != want {
			t.Errorf("readAzureArcSecret() path = %q, want %q", got, want)
		}
		return "arc-secret", nil
	}

	got, gotErr := tokenSource.Token(context.Background())
	if gotErr != nil {
		t.Fatalf("Token() failed: %v", gotErr)
	}
	if got.AccessToken != "azure-arc-token" {
		t.Errorf("Token().AccessToken = %q, want %q", got.AccessToken, "azure-arc-token")
	}
}

func TestAzureArcChallengeKeyFile(t *testing.T) {
	testCases := []struct {
		name           string
		status         int
		challenge      string
		wantKeyFile    string
		wantChallenged bool
		wantErr        string
	}{
		{
			name:   "non-challenge response",
			status: http.StatusInternalServerError,
		},
		{
			name:           "missing header",
			status:         http.StatusUnauthorized,
			wantChallenged: true,
			wantErr:        "Azure Arc managed identity response has no WWW-Authenticate header",
		},
		{
			name:           "malformed header",
			status:         http.StatusUnauthorized,
			challenge:      "Basic realm",
			wantChallenged: true,
			wantErr:        `invalid Azure Arc WWW-Authenticate header: "Basic realm"`,
		},
		{
			name:           "valid header",
			status:         http.StatusUnauthorized,
			challenge:      `Basic realm="/var/opt/azcmagent/tokens/identity.key"`,
			wantKeyFile:    "/var/opt/azcmagent/tokens/identity.key",
			wantChallenged: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			header := make(http.Header)
			if tc.challenge != "" {
				header.Set("WWW-Authenticate", tc.challenge)
			}
			err := &tokenError{
				err: &httpclient.HttpError{
					Response: &http.Response{
						StatusCode: tc.status,
						Header:     header,
					},
				},
			}

			gotKeyFile, gotChallenged, gotErr := azureArcChallengeKeyFile(err)
			if gotKeyFile != tc.wantKeyFile {
				t.Errorf("azureArcChallengeKeyFile() key file = %q, want %q", gotKeyFile, tc.wantKeyFile)
			}
			if gotChallenged != tc.wantChallenged {
				t.Errorf("azureArcChallengeKeyFile() challenged = %t, want %t", gotChallenged, tc.wantChallenged)
			}
			if tc.wantErr == "" {
				if gotErr != nil {
					t.Errorf("azureArcChallengeKeyFile() failed: %v", gotErr)
				}
				return
			}
			if gotErr == nil || gotErr.Error() != tc.wantErr {
				t.Errorf("azureArcChallengeKeyFile() error = %q, want %q", gotErr, tc.wantErr)
			}
		})
	}
}

func TestParseAzureMLTokenExpiry(t *testing.T) {
	testCases := []struct {
		name    string
		value   string
		want    time.Time
		wantErr bool
	}{
		{
			name:  "Linux format",
			value: "06/20/2030 02:57:58 +00:00",
			want:  time.Date(2030, time.June, 20, 2, 57, 58, 0, time.UTC),
		},
		{
			name:  "Windows format",
			value: "1/16/2030 5:24:12 AM +00:00",
			want:  time.Date(2030, time.January, 16, 5, 24, 12, 0, time.UTC),
		},
		{
			name:    "invalid format",
			value:   "2030-01-01T00:00:00Z",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, gotErr := parseAzureMLTokenExpiry(tc.value)
			if tc.wantErr {
				if gotErr == nil {
					t.Fatal("parseAzureMLTokenExpiry() succeeded, want error")
				}
				return
			}
			if gotErr != nil {
				t.Fatalf("parseAzureMLTokenExpiry() failed: %v", gotErr)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("parseAzureMLTokenExpiry() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestReadAzureArcSecretFromDirectory_validation(t *testing.T) {
	expectedDirectory := t.TempDir()
	validKeyFile := filepath.Join(expectedDirectory, "identity.key")
	if err := os.WriteFile(validKeyFile, []byte("secret"), 0o600); err != nil {
		t.Fatalf("write valid key file: %v", err)
	}
	invalidExtension := filepath.Join(expectedDirectory, "identity.txt")
	if err := os.WriteFile(invalidExtension, []byte("secret"), 0o600); err != nil {
		t.Fatalf("write invalid-extension key file: %v", err)
	}
	largeKeyFile := filepath.Join(expectedDirectory, "large.key")
	if err := os.WriteFile(largeKeyFile, []byte(strings.Repeat("x", int(azureArcMaximumSecretFileSize+1))), 0o600); err != nil {
		t.Fatalf("write large key file: %v", err)
	}
	otherDirectory := t.TempDir()
	unexpectedKeyFile := filepath.Join(otherDirectory, "identity.key")
	if err := os.WriteFile(unexpectedKeyFile, []byte("secret"), 0o600); err != nil {
		t.Fatalf("write unexpected key file: %v", err)
	}

	testCases := []struct {
		name    string
		keyFile string
		want    string
		wantErr string
	}{
		{
			name:    "valid key file",
			keyFile: validKeyFile,
			want:    "secret",
		},
		{
			name:    "unexpected directory",
			keyFile: unexpectedKeyFile,
			wantErr: fmt.Sprintf("unexpected Azure Arc managed identity secret directory %q", otherDirectory),
		},
		{
			name:    "invalid extension",
			keyFile: invalidExtension,
			wantErr: fmt.Sprintf("Azure Arc managed identity secret file %q must have a .key extension", invalidExtension),
		},
		{
			name:    "file too large",
			keyFile: largeKeyFile,
			wantErr: fmt.Sprintf("Azure Arc managed identity secret file %q exceeds 4096 bytes", largeKeyFile),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, gotErr := readAzureArcSecretFromDirectory(tc.keyFile, expectedDirectory)
			if tc.wantErr != "" {
				if gotErr == nil || gotErr.Error() != tc.wantErr {
					t.Errorf("readAzureArcSecretFromDirectory() error = %q, want %q", gotErr, tc.wantErr)
				}
				return
			}
			if gotErr != nil {
				t.Fatalf("readAzureArcSecretFromDirectory() failed: %v", gotErr)
			}
			if got != tc.want {
				t.Errorf("readAzureArcSecretFromDirectory() = %q, want %q", got, tc.want)
			}
		})
	}
}

func azureMSITokenResponse(accessToken string, expiry time.Time) map[string]any {
	return map[string]any{
		"access_token": accessToken,
		"expires_on":   expiry.Unix(),
		"token_type":   "Bearer",
	}
}
