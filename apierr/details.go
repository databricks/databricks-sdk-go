package apierr

import (
	"encoding/json"
	"time"
)

// ErrorDetails contains the error details of an API error. It is the union of
// known error details types and unknown details.
type ErrorDetails struct {
	ErrorInfo   *ErrorInfo
	RequestInfo *RequestInfo
	RetryInfo   *RetryInfo

	// UnknownDetails contains error details that cannot be unmarshalled into
	// one of the known types above.
	UnknownDetails []any
}

// ErrorInfo describes the cause of the error with structured details.
type ErrorInfo struct {
	// The reason of the error. This is a constant value that identifies the
	// proximate cause of the error.
	Reason string

	// The logical grouping to which the "reason" belongs.
	Domain string

	// Additional structured details about this error.
	Metadata map[string]string
}

// RequestInfo Contains metadata about the request that clients can attach when
// filing a bug or providing other forms of feedback.
type RequestInfo struct {
	// An opaque string that should only be interpreted by the service that
	// generated it. For example, it can be used to identify requests in the
	// service's logs.
	RequestID string

	// Any data that was used to serve this request. For example, an encrypted
	// stack trace that can be sent back to the service provider for debugging.
	ServingData string
}

// RetryInfo describes when the clients can retry a failed request. Clients
// could ignore the recommendation here or retry when this information is
// missing from error responses.
//
// It's always recommended that clients should use exponential backoff when
// retrying.
//
// Clients should wait until `retry_delay` amount of time has passed since
// receiving the error response before retrying.  If retrying requests also
// fail, clients should use an exponential backoff scheme to gradually increase
// the delay between retries based on `retry_delay`, until either a maximum
// number of retries have been reached or a maximum retry delay cap has been
// reached.
type RetryInfo struct {
	// Clients should wait at least this long between retrying the same request.
	RetryDelay time.Duration
}

const errorInfoType string = "type.googleapis.com/google.rpc.ErrorInfo"

// errorInfoPb is the wire-format representation of ErrorInfo. It is used
// internally to unmarshal ErrorInfo from JSON.
type errorInfoPb struct {
	Reason   string            `json:"reason"`
	Domain   string            `json:"domain"`
	Metadata map[string]string `json:"metadata"`
}

const requestInfoType string = "type.googleapis.com/google.rpc.RequestInfo"

// requestInfoPb is the wire-format representation of RequestInfo. It is used
// internally to unmarshal RequestInfo from JSON.
type requestInfoPb struct {
	RequestID   string `json:"request_id"`
	ServingData string `json:"serving_data"`
}

const retryInfoType string = "type.googleapis.com/google.rpc.RetryInfo"

// retryInfoPb is the wire-format representation of RetryInfo. It is used
// internally to unmarshal RetryInfo from JSON.
type retryInfoPb struct {
	RetryDelay durationPb `json:"retry_delay"`
}

type durationPb struct {
	Seconds int64 `json:"seconds"`
	Nanos   int64 `json:"nanos"`
}

func parseErrorDetails(details []any) ErrorDetails {
	var ed ErrorDetails
	for _, d := range details {
		switch d := d.(type) {
		case *ErrorInfo:
			ed.ErrorInfo = d
		case *RequestInfo:
			ed.RequestInfo = d
		case *RetryInfo:
			ed.RetryInfo = d
		default:
			ed.UnknownDetails = append(ed.UnknownDetails, d)
		}
	}
	return ed
}

func unmarshalDetails(details json.RawMessage) (any, error) {
	var a map[string]any
	if err := json.Unmarshal(details, &a); err != nil {
		return nil, err
	}

	if t, ok := a["@type"].(string); ok {
		switch t {
		case errorInfoType:
			var pb errorInfoPb
			if err := json.Unmarshal(details, &pb); err != nil {
				return nil, err
			}
			return ErrorInfo(pb), nil
		case requestInfoType:
			var pb requestInfoPb
			if err := json.Unmarshal(details, &pb); err != nil {
				return nil, err
			}
			return RequestInfo(pb), nil
		case retryInfoType:
			var r retryInfoPb
			if err := json.Unmarshal(details, &r); err != nil {
				return nil, err
			}
			return RetryInfo{RetryDelay: time.Duration(r.RetryDelay.Seconds)*time.Second + time.Duration(r.RetryDelay.Nanos)*time.Nanosecond}, nil
		}
	}

	// If the type is not known, try to unmarshal it as a map.
	return a, nil
}
