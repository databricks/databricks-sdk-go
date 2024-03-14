package useragent

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/databricks/databricks-sdk-go/version"
	"golang.org/x/mod/semver"
)

const (
	RuntimeKey = "runtime"
	CicdKey    = "cicd"
	AuthKey    = "auth"
)

// WithProduct sets the product name and product version globally.
// It should be called by developers to differentiate their application from others.
func WithProduct(name, version string) {
	if err := matchAlphanum(name); err != nil {
		panic(err)
	}
	if err := matchSemVer(version); err != nil {
		panic(err)
	}
	productName = name
	productVersion = version
}

var productName = "unknown"
var productVersion = "0.0.0"

const sdkName = "databricks-sdk-go"
const sdkVersion = version.Version

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
	extra = extra.With(key, value)
}

func getUpstreamUserAgentInfo() []info {
	product := os.Getenv("DATABRICKS_SDK_UPSTREAM")
	version := os.Getenv("DATABRICKS_SDK_UPSTREAM_VERSION")
	if product == "" || version == "" {
		return nil
	}
	return []info{
		{"upstream", product},
		{"upstream-version", version},
	}
}

// InContext populates context with user agent dimension,
// usually to differentiate subsets of functionality and
// agreed with Databricks.
func InContext(ctx context.Context, key, value string) context.Context {
	uac, _ := ctx.Value(ctxAgent).(data)
	uac = uac.With(key, value)
	return context.WithValue(ctx, ctxAgent, uac)
}

// FromContext gets compliant user-agent string in a given context
func FromContext(ctx context.Context) string {
	base := data{
		{productName, productVersion},
		{sdkName, sdkVersion},
		{"go", goVersion()},
		{"os", runtime.GOOS},
	}
	base = append(base, extra...)
	base = append(base, getUpstreamUserAgentInfo()...)
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

// Sanitize replaces all non-alphanumeric characters with a hyphen. Use this to
// ensure that the user agent value is valid. This is useful when the value is not
// ensured to be valid at compile time.
//
// Example: You want to avoid having '/' and ' ' in the value because it will
// make downstream applications fail.
//
// Note: Semver strings are comprised of alphanumeric characters, hyphens, periods
// and plus signs. This function will not remove these characters.
// see:
// 1. https://semver.org/#spec-item-9
// 2. https://semver.org/#spec-item-10
func Sanitize(s string) string {
	return regexp.MustCompile(`[^0-9A-Za-z_\.\+-]`).ReplaceAllString(s, "-")
}

// With always uses the latest value for a given alphanumeric key.
// Panics if key or value don't satisfy alphanumeric or semver format.
func (d data) With(key, value string) data {
	if err := matchAlphanum(key); err != nil {
		panic(err)
	}
	if err := matchAlphanumOrSemVer(value); err != nil {
		panic(err)
	}
	var c data
	var found bool
	for _, i := range d {
		if i.Key == key {
			i.Value = value
			found = true
		}
		c = append(c, i)
	}
	if !found {
		c = append(c, info{key, value})
	}
	return c
}

func (d data) String() string {
	pairs := []string{}
	for _, v := range d {
		pairs = append(pairs, v.String())
	}
	return strings.Join(pairs, " ")
}
