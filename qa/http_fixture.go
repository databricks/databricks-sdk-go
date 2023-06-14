package qa

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/client"
	"github.com/databricks/databricks-sdk-go/config"
	"github.com/stretchr/testify/assert"
)

// HTTPFixture defines request structure for test
type HTTPFixture struct {
	Method          string      `json:"method"`
	Resource        string      `json:"resource"`
	Response        interface{} `json:"response,omitempty"`
	ResponseHeaders http.Header `json:"response_headers,omitempty"`
	Status          int         `json:"status"`
	ExpectedHeaders http.Header `json:"expected_headers,omitempty"`
	ExpectedRequest interface{} `json:"expected_request,omitempty"`
	ReuseRequest    bool        `json:"reuse_request,omitempty"`
	MatchAny        bool        `json:"match_any,omitempty"`
}

var HTTPFailures = []HTTPFixture{
	{
		MatchAny:     true,
		ReuseRequest: true,
		Status:       418,
		Response: apierr.APIError{
			ErrorCode:  "NONSENSE",
			StatusCode: 418,
			Message:    "I'm a teapot",
		},
	},
}

type HTTPFixtures []HTTPFixture

func CheckHeadersContain(actual, expected http.Header) error {
	for k, v := range expected {
		actual, ok := actual[k]
		if !ok {
			return fmt.Errorf("header %s not present in request", k)
		}
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		sort.Slice(actual, func(i, j int) bool {
			return actual[i] < actual[j]
		})

		if !reflect.DeepEqual(v, actual) {
			return fmt.Errorf("header %s does not match, expected %v, got %v", k, v, actual)
		}
	}
	return nil
}

func (fixtures HTTPFixtures) NewServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		found := false
		for i, fixture := range fixtures {
			if (req.Method == fixture.Method && req.RequestURI == fixture.Resource) || fixture.MatchAny {
				if fixture.ExpectedRequest != nil {
					buf := new(bytes.Buffer)
					_, err := buf.ReadFrom(req.Body)
					assert.NoError(t, err, err)
					jsonStr, err := json.Marshal(fixture.ExpectedRequest)
					assert.NoError(t, err, err)
					assert.JSONEq(t, string(jsonStr), buf.String(), "json strings do not match")
				}
				// Check headers
				err := CheckHeadersContain(req.Header, fixture.ExpectedHeaders)
				if err != nil {
					t.Fatal(err)
				}
				// Set response headers
				for k, v := range fixture.ResponseHeaders {
					for i := range v {
						rw.Header().Add(k, v[i])
					}
				}
				if fixture.Status == 0 {
					rw.WriteHeader(200)
				} else {
					rw.WriteHeader(fixture.Status)
				}
				// Send response body
				if fixture.Response != nil {
					if alreadyJSON, ok := fixture.Response.(string); ok {
						_, err := rw.Write([]byte(alreadyJSON))
						assert.NoError(t, err, err)
					} else {
						responseBytes, err := json.Marshal(fixture.Response)
						if err != nil {
							assert.NoError(t, err, err)
							t.FailNow()
						}
						_, err = rw.Write(responseBytes)
						assert.NoError(t, err, err)
					}
				}
				found = true
				// Reset the request if it is already used
				if !fixture.ReuseRequest {
					fixtures[i] = HTTPFixture{}
				}
				break
			}
		}
		if !found {
			receivedRequest := map[string]interface{}{}
			buf := new(bytes.Buffer)
			_, err := buf.ReadFrom(req.Body)
			assert.NoError(t, err, err)
			err = json.Unmarshal(buf.Bytes(), &receivedRequest)
			assert.NoError(t, err, err)

			expectedRequest := ""
			if len(receivedRequest) > 0 {
				// guessing model name would require going over AST,
				// which is not something i'm willing to write on my weekend
				expectedRequest += "ExpectedRequest: XXX {\n"
				for key, value := range receivedRequest {
					camel := ""
					for _, part := range strings.Split(key, "_") {
						if len(key) < 4 {
							// golang styles, meh...
							camel += strings.ToUpper(key)
						} else {
							camel += strings.Title(part)
						}
					}
					// best effort prediction of what struct should look like...
					expectedRequest += fmt.Sprintf("					%s: %#v,\n", camel, value)
				}
				expectedRequest += "				},\n"
				expectedRequest += fmt.Sprintf("				// ExpectedRequest: %#v,\n", receivedRequest)
			}
			stub := fmt.Sprintf(`{
				Method:   "%s",
				Resource: "%s",
				%s
				Response: XXX {
					// fill in specific fields...
				},
			},`, req.Method, req.RequestURI, expectedRequest)
			assert.Fail(t, fmt.Sprintf("Missing stub, please add: %s", stub))
			t.FailNow()
		}
	}))
}

// Client creates DatabricksClient for emulated HTTP server
func (fixtures HTTPFixtures) Config(t *testing.T) (*config.Config, *httptest.Server) {
	server := fixtures.NewServer(t)
	return &config.Config{
		Credentials:      config.PatCredentials{},
		Host:             server.URL,
		Token:            "x",
		AzureEnvironment: "public",
	}, server
}

// HTTPFixturesApply is a helper method for executing a callback and closing emulated server after
func (fixtures HTTPFixtures) Apply(t *testing.T, callback func(ctx context.Context, cfg *config.Config)) {
	cfg, server := fixtures.Config(t)
	defer server.Close()
	callback(context.Background(), cfg)
}

func (fixtures HTTPFixtures) Client(t *testing.T) (*client.DatabricksClient, *httptest.Server) {
	cfg, server := fixtures.Config(t)
	client, err := client.New(cfg)
	if err != nil {
		t.Fatalf("client: %s", err)
	}
	return client, server
}

func (fixtures HTTPFixtures) ApplyClient(t *testing.T, callback func(ctx context.Context, client *client.DatabricksClient)) {
	client, server := fixtures.Client(t)
	defer server.Close()
	callback(context.Background(), client)
}
