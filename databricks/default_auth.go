package databricks

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"sort"
	"strings"
)

var (
	defaultConfig = Config{
		Credentials:        DefaultCredentials{},
		RateLimitPerSecond: 15,
		DebugTruncateBytes: 1024,
	}

	authProviders = []CredentialsProvider{
		PatCredentials{},
		BasicCredentials{},
		AzureClientSecretCredentials{},
		AzureCliCredentials{},
		GoogleDefaultCredentials{},
		GoogleCredentials{},
		DatabricksCliCredentials{},
	}

	// ConfigAttributes holds all configuration attributes required for Databricks SDK
	ConfigAttributes = discoverConfigurations(authProviders, defaultConfig)
)

func discoverConfigurations(providers []CredentialsProvider, cfg Config) (res []ConfigAttribute) {
	overlap := map[string]string{}
	res = append(res, discoverAttributesOf(&cfg, "config", overlap)...)
	for _, credentials := range providers {
		res = append(res, discoverAttributesOf(credentials, credentials.Name(), overlap)...)
	}
	return
}

func discoverAttributesOf(instance interface{}, name string, overlap map[string]string) (res []ConfigAttribute) {
	for _, attr := range attributesOf(instance, name) {
		where, exists := overlap[attr.Name]
		if exists {
			panic(fmt.Errorf("sanity check failed: %s previously defined in %s", attr.Name, where))
		}
		overlap[attr.Name] = attr.where
		res = append(res, attr)
	}
	return
}

type DefaultCredentials struct {
	Explicit map[string]string
	skip     string
}

func (c DefaultCredentials) Name() string {
	return "default"
}

func (c DefaultCredentials) Configure(ctx context.Context, cfg *Config) (func(*http.Request) error, error) {
	if c.Explicit == nil {
		c.Explicit = map[string]string{}
	}
	effective, err := c.effectiveCredentials(cfg, c.Explicit)
	if err != nil {
		return nil, err
	}
	for _, v := range effective {
		if v.Name() == c.skip {
			// in case we're re-wrapping Default Credentials as Databricks CLI profile
			continue
		}
		visitor, err := v.Configure(ctx, cfg)
		if err != nil {
			return nil, err
		}
		if visitor == nil {
			continue
		}
		return visitor, nil
	}
	return nil, fmt.Errorf("cannot configure default credentials")
}

func (c DefaultCredentials) effectiveCredentials(cfg *Config, explicit map[string]string) ([]CredentialsProvider, error) {
	placeholders := map[string]reflect.Value{
		"config": reflect.ValueOf(cfg), // TODO: figure out what to do with mutexes...
	}
	refOf := func(cp CredentialsProvider) reflect.Value {
		// field :=
		// //field.Elem().Set(reflect.ValueOf(b))
		// println(field.String())
		// return reflect.ValueOf(&cp)
		return reflect.New(reflect.TypeOf(cp))
	}
	for _, provider := range authProviders { // TODO: this does need to be in the same order.
		placeholders[provider.Name()] = refOf(provider)
	}
	authsUsed := map[string]bool{}
	for _, attr := range ConfigAttributes {
		if c.skip != "" && attr.where == c.skip {
			continue
		}
		target, ok := placeholders[attr.where]
		if !ok {
			return nil, fmt.Errorf("cannot find where to put %s", attr.Name)
		}
		var raw string
		// explicit attributes have precedence
		expl, okExpl := explicit[attr.Name]
		env, okEnv := attr.ReadEnv()
		if expl == "" {
			okExpl = false
		}
		if env == "" {
			okEnv = false
		}
		if !okExpl && !okEnv {
			continue
		}
		if okExpl {
			raw = expl
		} else {
			raw = env
		}
		err := attr.SetS(target, raw)
		if err != nil {
			return nil, err
		}
		if attr.where != "config" {
			authsUsed[attr.where] = true
		}
	}
	configuredCreds := []string{}
	for v := range authsUsed {
		configuredCreds = append(configuredCreds, v)
	}
	sort.Strings(configuredCreds)
	if len(configuredCreds) > 1 {
		return nil, fmt.Errorf("more than one authorization method configured: %s",
			strings.Join(configuredCreds, " and "))
	}
	creds := []CredentialsProvider{}
	for k, v := range placeholders {
		if k == "config" {
			continue // TODO: or more?
		}
		creds = append(creds, v.Elem().Interface().(CredentialsProvider))
	}
	return creds, nil
}
