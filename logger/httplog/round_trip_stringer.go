package httplog

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RoundTripStringer struct {
	Response           *http.Response
	Err                error
	RequestBody        []byte
	ResponseBody       []byte
	DebugHeaders       bool
	DebugTruncateBytes int
}

func (r RoundTripStringer) String() string {
	request := r.Response.Request
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s %s", request.Method,
		escapeNewLines(request.URL.Path)))
	if request.URL.RawQuery != "" {
		sb.WriteString("?")
		q, _ := url.QueryUnescape(request.URL.RawQuery)
		sb.WriteString(q)
	}
	sb.WriteString("\n")
	if r.DebugHeaders {
		sb.WriteString("> * Host: ")
		sb.WriteString(escapeNewLines(request.Host))
		sb.WriteString("\n")
		for k, v := range request.Header {
			trunc := onlyNBytes(strings.Join(v, ""), r.DebugTruncateBytes)
			sb.WriteString(fmt.Sprintf("> * %s: %s\n", k, escapeNewLines(trunc)))
		}
	}
	if len(r.RequestBody) > 0 {
		sb.WriteString(r.redactedDump("> ", r.RequestBody))
		sb.WriteString("\n")
	}
	sb.WriteString("< ")
	if r.Response != nil {
		sb.WriteString(fmt.Sprintf("%s %s", r.Response.Proto, r.Response.Status))
		// Only display error on this line if the response body is empty.
		// Otherwise the response body will include details about the error.
		if len(r.ResponseBody) == 0 && r.Err != nil {
			sb.WriteString(fmt.Sprintf(" (Error: %s)", r.Err))
		}
	} else {
		sb.WriteString(fmt.Sprintf("Error: %s", r.Err))
	}
	sb.WriteString("\n")
	if r.DebugHeaders {
		for k, v := range r.Response.Header {
			trunc := onlyNBytes(strings.Join(v, ""), r.DebugTruncateBytes)
			sb.WriteString(fmt.Sprintf("< * %s: %s\n", k, escapeNewLines(trunc)))
		}
	}
	if len(r.ResponseBody) > 0 {
		sb.WriteString(r.redactedDump("< ", r.ResponseBody))
	}
	return sb.String()
}

func (r RoundTripStringer) redactedDump(prefix string, body []byte) (res string) {
	return bodyLogger{
		debugTruncateBytes: r.DebugTruncateBytes,
	}.redactedDump(prefix, body)
}

// CWE-117 prevention
func escapeNewLines(in string) string {
	in = strings.Replace(in, "\n", "", -1)
	in = strings.Replace(in, "\r", "", -1)
	return in
}
