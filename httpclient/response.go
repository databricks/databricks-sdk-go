package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/common"
	"github.com/databricks/databricks-sdk-go/logger"
)

func WithResponseHeader(key string, value *string) DoOption {
	return DoOption{
		out: func(body *common.ResponseWrapper) error {
			*value = body.Response.Header.Get(key)
			return nil
		},
	}
}

// WithResponseUnmarshal unmarshals the response body into response. The
// supported response types are the following:
//   - *bytes.Buffer,
//   - *io.ReadCloser,
//   - *[]byte,
//   - a pointer to a struct with a Contents io.ReadCloser field,
//   - a pointer to a struct representing a JSON object.
//
// If response is a pointer to a io.ReadCloser or a struct with a io.ReadCloser
// field name "Contents", then the response io.ReadCloser is set to the value of
// the body's reader without actually reading it.
func WithResponseUnmarshal(response any) DoOption {
	return DoOption{
		in: func(r *http.Request) error {
			if r.Header.Get("Accept") != "" {
				return nil
			}
			switch response.(type) {
			case *bytes.Buffer, *io.ReadCloser, *[]byte:
				// don't send Accept header for raw types, even though we have
				// openapi/code/method.go:440#IsResponseByteStream() setting the
				// Accept header explicitly.
				return nil
			default:
				r.Header.Set("Accept", "application/json")
				return nil
			}
		},
		out: func(body *common.ResponseWrapper) error {
			if response == nil {
				return nil
			}
			err := injectHeaders(response, body)
			if err != nil {
				return err
			}

			if field, ok := findContentsField(response); ok {
				field.Set(reflect.ValueOf(body.ReadCloser))
				return nil
			}
			if reader, ok := response.(*io.ReadCloser); ok {
				*reader = body.ReadCloser
				return nil
			}
			if buffer, ok := response.(*bytes.Buffer); ok {
				defer body.ReadCloser.Close()
				_, err := buffer.ReadFrom(body.ReadCloser)
				return err
			}

			// At this point, fully read the content of the body and use it
			// to populate the response object (whether it is a slice of bytes
			// or a JSON object).
			defer body.ReadCloser.Close()
			bodyBytes, err := io.ReadAll(body.ReadCloser)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}
			if len(bodyBytes) == 0 {
				return nil
			}
			if bs, ok := response.(*[]byte); ok {
				*bs = bodyBytes
				return nil
			}
			if err = json.Unmarshal(bodyBytes, &response); err != nil {
				return fmt.Errorf("failed to unmarshal response body: %w. %s", err, apierr.MakeUnexpectedResponse(body.Response, body.RequestBody.DebugBytes, bodyBytes))
			}
			return nil
		},
	}
}

func findContentsField(response any) (*reflect.Value, bool) {
	value := reflect.ValueOf(response)
	value = reflect.Indirect(value)
	if value.Kind() != reflect.Struct {
		return nil, false
	}

	// Check if there is a "Contents" of type io.ReadCloser
	// This is by internal convention with the teams.
	ioType := reflect.TypeOf((*io.ReadCloser)(nil)).Elem()
	contentField := value.FieldByName("Contents")
	if !contentField.IsValid() || !contentField.CanSet() || contentField.Type() != ioType {
		return nil, false
	}
	return &contentField, true
}

func injectHeaders(response any, body *common.ResponseWrapper) error {
	value := reflect.ValueOf(response)
	value = reflect.Indirect(value)
	objectType := value.Type()

	if objectType.Kind() != reflect.Struct {
		return nil
	}
	headers := body.Response.Header

	for i := 0; i < objectType.NumField(); i++ {
		field := objectType.Field(i)
		headerTag := parseHeaderTag(field)
		if headerTag.ignore {
			continue
		}
		// We only support a single value for a header, and of types int and string.
		headerValue := headers.Get(headerTag.name)
		if headerValue == "" {
			continue
		}
		switch field.Type.Kind() {
		case reflect.String:
			value.Field(i).Set(reflect.ValueOf(headerValue))
		case reflect.Int:
			intValue, err := strconv.Atoi(headerValue)
			if err != nil {
				return fmt.Errorf("failed to parse header %s as int: %w", headerTag.name, err)
			}
			value.Field(i).Set(reflect.ValueOf(intValue))
		case reflect.Int64:
			intValue, err := strconv.ParseInt(headerValue, 10, 64)
			if err != nil {
				return fmt.Errorf("failed to parse header %s as int: %w", headerTag.name, err)
			}
			value.Field(i).Set(reflect.ValueOf(intValue))
		default:
			// Do not fail the request if the header is not supported to avoid breaking changes.
			logger.Warnf(context.Background(), "unsupported header type %s for field %s", field.Type.Kind(), field.Name)
		}

	}
	return nil
}

type headerTag struct {
	name      string
	omitempty bool
	ignore    bool
}

func parseHeaderTag(field reflect.StructField) headerTag {
	raw := field.Tag.Get("header")
	name := field.Name

	if raw == "-" || raw == "" {
		return headerTag{ignore: true}
	}

	parts := strings.Split(raw, ",")

	if parts[0] != "" {
		name = parts[0]
	}

	headerTag := headerTag{
		name: name,
	}

	for _, v := range parts {
		switch v {
		case "omitempty":
			headerTag.omitempty = true
		}
	}
	return headerTag
}
