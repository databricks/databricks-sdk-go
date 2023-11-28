package fixtures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/databricks/databricks-sdk-go/openapi/code"
)

func resourceFromRequest(req *http.Request) string {
	resource := req.RequestURI
	if resource == "" {
		query := ""
		if req.URL.RawQuery != "" {
			query = "?" + req.URL.RawQuery
		}
		resource = req.URL.Path + query
	}
	return resource
}

func bodyStub(req *http.Request) (string, error) {
	if req.Body == nil {
		return "", nil
	}
	receivedRequest := map[string]any{}
	actualRequestJSON := new(bytes.Buffer)
	_, err := actualRequestJSON.ReadFrom(req.Body)
	if err != nil {
		return "", fmt.Errorf("read: %w", err)
	}
	expectedRequest := ""
	if actualRequestJSON.Len() == 0 {
		return "", nil
	}
	err = json.Unmarshal(actualRequestJSON.Bytes(), &receivedRequest)
	if err != nil {
		return "", fmt.Errorf("unmarshal: %w", err)
	}
	// guessing model name would require going over AST,
	// which is not something i'm willing to write on my weekend
	expectedRequest += "ExpectedRequest: XXX {\n"
	for key, value := range receivedRequest {
		camel := (&code.Named{Name: key}).PascalName()
		expectedRequest += fmt.Sprintf("\t\t\t%s: %#v,\n", camel, value)
	}
	expectedRequest += "\t\t},\n"
	expectedRequest += fmt.Sprintf("\t\t// ExpectedRequest: %#v,\n\t\t", receivedRequest)
	return expectedRequest, nil
}
