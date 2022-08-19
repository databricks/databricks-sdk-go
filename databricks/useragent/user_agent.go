package useragent

import (
	"context"
	"fmt"
	"runtime"
	"strings"

	"github.com/databricks/databricks-sdk-go/databricks/internal"
	"golang.org/x/mod/semver"
)

const SdkName = "databricks-sdk-go"

// WithProduct is expected to be set by developers to differentiate
// their app from others. usage of `databricks.WithProduc` is preferred.
func WithProduct(name, version string) {
	// TODO: validate fields
	product = name
	productVersion = version
}

// product holds product name
var product = "unknown"

// productVersion holds different release version
var productVersion = "0.0.0"

// extra holds per-process static extra information
// bits, that are agreed upfront with Databricks.
var extra data

// private type for user agent info in the context
type key int

// context key for holding volatile part of user agent data
const ctxAgent key = 5

// WithUserAgentExtra sets per-process extra user agent data,
// which integration developers have agreed with Databricks.
func WithUserAgentExtra(key, value string) {
	// TODO: validate fields
	extra = append(extra, info{key, value})
}

// InContext populates context with user agent dimension,
// usually to differentiate subsets of functionality and
// agreed with Databricks.
func InContext(ctx context.Context, k, v string) context.Context {
	uac, _ := ctx.Value(ctxAgent).(data)
	// TODO: validate fields
	uac = append(uac, info{k, v})
	return context.WithValue(ctx, ctxAgent, uac)
}

// FromContext gets compliant user-agent string in a given context
func FromContext(ctx context.Context) string {
	base := data{
		{product, productVersion},
		{SdkName, internal.Version},
		{"go", goVersion()},
		{"os", runtime.GOOS},
	}
	base = append(base, extra...)
	uac, _ := ctx.Value(ctxAgent).(data)
	return append(base, uac...).String()
}

func goVersion() string {
	gv := runtime.Version()
	ssv := strings.ReplaceAll(gv, "go", "v")
	sv := semver.Canonical(ssv)
	return strings.TrimPrefix(sv, "v")
}

type info struct {
	Key   string
	Value string
}

func (u info) String() string {
	return fmt.Sprintf("%s/%s", u.Key, u.Value)
}

type data []info

func (d data) String() string {
	pairs := []string{}
	for _, v := range d {
		pairs = append(pairs, v.String())
	}
	return strings.Join(pairs, " ")
}
