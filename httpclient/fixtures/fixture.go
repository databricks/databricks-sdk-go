package fixtures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
)

// HTTPFixture defines request structure for test
type HTTPFixture struct {
	Method       string
	Resource     string
	MatchAny     bool
	ReuseRequest bool

	Response        any
	Status          int
	ExpectedRequest any
	PassFile        string
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
		Proto:      "HTTP/1.1",
		StatusCode: f.Status,
		Status:     http.StatusText(f.Status),
		Body:       io.NopCloser(strings.NewReader(body)),
	}, nil
}

func (f HTTPFixture) Reply(req *http.Request) (*http.Response, error) {
	if f.Status == 0 {
		f.Status = 200
	}
	if f.PassFile != "" {
		raw, err := os.ReadFile(f.PassFile)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", f.PassFile, err)
		}
		return f.replyWith(req, string(raw))
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
