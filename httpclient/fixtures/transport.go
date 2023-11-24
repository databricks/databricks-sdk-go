package fixtures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/databricks/databricks-sdk-go/openapi/code"
)

// HTTPFixture defines request structure for test
type HTTPFixture struct {
	Method          string
	Resource        string
	Response        any
	Status          int
	ExpectedRequest any
	ReuseRequest    bool
	MatchAny        bool
}

func (f HTTPFixture) Match(method, resource string) bool {
	if f.MatchAny {
		return true
	}
	return method == f.Method && resource == f.Resource
}

func (f HTTPFixture) AssertRequest(req *http.Request) error {
	if f.ExpectedRequest == nil {
		return nil
	}
	actualRequestBuf := new(bytes.Buffer)
	_, err := actualRequestBuf.ReadFrom(req.Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}
	// do pretty-priting, so that it's easier to debug failures
	expectedRequestJSON, err := json.MarshalIndent(f.ExpectedRequest, "", "  ")
	if err != nil {
		return fmt.Errorf("unmarshal expected: %w", err)
	}
	var expectedANY, actualANY any
	err = json.Unmarshal(expectedRequestJSON, &expectedANY)
	if err != nil {
		return fmt.Errorf("marshal expected to any: %w", err)
	}
	err = json.Unmarshal(actualRequestBuf.Bytes(), &actualANY)
	if err != nil {
		return fmt.Errorf("marshal actual to any: %w", err)
	}
	if !reflect.DeepEqual(expectedANY, actualANY) {
		actualRequestJSON, _ := json.MarshalIndent(actualANY, "", "  ")
		return fmt.Errorf("want %s, got %s", expectedRequestJSON, actualRequestJSON)
	}
	return nil
}

func (f HTTPFixture) replyWith(req *http.Request, body string) (*http.Response, error) {
	return &http.Response{
		Request:    req,
		StatusCode: f.Status,
		Body:       io.NopCloser(strings.NewReader(body)),
	}, nil
}

func (f HTTPFixture) Reply(req *http.Request) (*http.Response, error) {
	if f.Status == 0 {
		f.Status = 200
	}
	if f.Response == nil {
		return f.replyWith(req, "")
	}
	if alreadyJSON, ok := f.Response.(string); ok {
		return f.replyWith(req, alreadyJSON)
	}
	responseBytes, err := json.Marshal(f.Response)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	return f.replyWith(req, string(responseBytes))
}

var HTTPFailures = []HTTPFixture{
	{
		MatchAny:     true,
		ReuseRequest: true,
		Status:       418,
		Response:     `{"error_code": "NONSENSE", "message": "I'm a teapot"}`,
	},
}

type HTTPFixtures []HTTPFixture

func (fixtures HTTPFixtures) SkipRetryOnIO() bool {
	return true
}

func (fixtures HTTPFixtures) RoundTrip(req *http.Request) (*http.Response, error) {
	resource := req.RequestURI
	if resource == "" {
		query := ""
		if req.URL.RawQuery != "" {
			query = "?" + req.URL.RawQuery
		}
		resource = req.URL.Path + query
	}
	for _, f := range fixtures {
		if !f.Match(req.Method, resource) {
			continue
		}
		if f.Status == 0 {
			f.Status = 200
		}
		err := f.AssertRequest(req)
		if err != nil {
			return nil, fmt.Errorf("expected request: %w", err)
		}
		return f.Reply(req)
	}
	expectedRequest, err := fixtures.bodyStub(req)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	// whitespace in this string is very important for unit tests
	stub := fmt.Sprintf(`	{
		Method:   "%s",
		Resource: "%s",
		%s
		Response: XXX {
			// fill in specific fields...
		},
	},`, req.Method, resource, expectedRequest)
	return nil, fmt.Errorf("missing stub, please add: %s", stub)
}

func (fixtures HTTPFixtures) bodyStub(req *http.Request) (string, error) {
	if req.Body == nil {
		return "", nil
	}
	receivedRequest := map[string]any{}
	actualRequestJSON := new(bytes.Buffer)
	_, err := actualRequestJSON.ReadFrom(req.Body)
	if err != nil {
		return "", fmt.Errorf("read body: %w", err)
	}
	expectedRequest := ""
	if actualRequestJSON.Len() == 0 {
		return "", nil
	}
	err = json.Unmarshal(actualRequestJSON.Bytes(), &receivedRequest)
	if err != nil {
		return "", fmt.Errorf("unmarshal body: %w", err)
	}
	// guessing model name would require going over AST,
	// which is not something i'm willing to write on my weekend
	expectedRequest += "ExpectedRequest: XXX {\n"
	for key, value := range receivedRequest {
		camel := (&code.Named{Name: key}).PascalName()
		expectedRequest += fmt.Sprintf("\t\t\t%s: %#v,\n", camel, value)
	}
	expectedRequest += "\t\t},\n"
	expectedRequest += fmt.Sprintf("\t\t// ExpectedRequest: %#v,\n", receivedRequest)
	return expectedRequest, nil
}
