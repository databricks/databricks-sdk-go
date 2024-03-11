package useragent

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"unicode"

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

// Sanitize the user agent value. This is useful when the value is not ensured to be
// to be valid at compile time. Having this sanitization then ensures downstream
// applications can correctly parse the full user agent header, by making sure
// characters like '/' and ' ' are not present in the value.
func Sanitize(s string) string {
	allowList := func(r rune) rune {
		if r == '.' || unicode.IsLetter(r) || unicode.IsDigit(r) {
			return r
		}
		return '-'
	}

	return strings.Map(allowList, s)
}

func (u info) String() string {
	return fmt.Sprintf("%s/%s", u.Key, u.Value)
}

type data []info

// Validate the key value pair being set in the user agent. Error if invalid.
func validate(key, value string) error {
	// DBR versions as set in the `DATABRICKS_RUNTIME_VERSION` environment variable
	// are not valid semver strings. Eg: 15.1 or client.0. Thus we allow arbitrary
	// values for the runtime key.
	if key == RuntimeKey {
		return nil
	}
	if !isAlphanum(key) {
		return fmt.Errorf("expected user agent key to be alphanumeric: %s", key)
	}
	if !isAlphanum(value) && !isSemVer(value) {
		return fmt.Errorf("expected user agent value for key %q to be alphanumeric or semver: %s", key, value)
	}
	return nil
}

// With always uses the latest value for a given alphanumeric key.
// Panics if key or value don't satisfy alphanumeric or semver format.
func (d data) With(key, value string) data {
	if err := validate(key, value); err != nil {
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
