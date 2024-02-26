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
			// If the field contains a "Content" field of type bytes.Buffer, write the body over there and return.
			if field, ok := findContentsField(response, body); ok {
				// If so, set the value
				field.Set(reflect.ValueOf(body.ReadCloser))
				return nil
			}

			// If the destination is bytes.Buffer, write the body over there
			if raw, ok := response.(*io.ReadCloser); ok {
				*raw = body.ReadCloser
				return nil
			}
			defer body.ReadCloser.Close()
			bs, err := io.ReadAll(body.ReadCloser)
			if err != nil {
				return fmt.Errorf("failed to read response body: %w", err)
			}
			if len(bs) == 0 {
				return nil
			}
			// If the destination is a byte slice or buffer, pass the body verbatim.
			if raw, ok := response.(*[]byte); ok {
				*raw = bs
				return nil
			}
			if raw, ok := response.(*bytes.Buffer); ok {
				_, err := raw.Write(bs)
				return err
			}
			err = json.Unmarshal(bs, &response)
			if err != nil {
				return apierr.MakeUnexpectedError(body.Response, err, body.RequestBody.DebugBytes, bs)
			}
			return nil
		},
	}
}

func findContentsField(response any, body *common.ResponseWrapper) (*reflect.Value, bool) {
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
