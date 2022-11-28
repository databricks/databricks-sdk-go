package config

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

var ConfigAttributes = loadAttrs()

// attributes holds meta-representation of Config configuration options
type attributes []ConfigAttribute

func (a attributes) DebugString(cfg *Config) string {
	buf := []string{}
	attrsUsed := []string{}
	envsUsed := []string{}
	for _, attr := range a {
		if attr.IsZero(cfg) {
			continue
		}
		// Don't include internal fields in debug representation.
		if attr.Internal {
			continue
		}
		v := "***"
		if !attr.Sensitive {
			v = attr.GetString(cfg)
		}
		attrsUsed = append(attrsUsed, fmt.Sprintf("%s=%s", attr.Name, v))
		for _, envName := range attr.EnvVars {
			v := os.Getenv(envName)
			if v == "" {
				continue
			}
			envsUsed = append(envsUsed, envName)
		}
	}
	if len(attrsUsed) > 0 {
		buf = append(buf, fmt.Sprintf("Config: %s", strings.Join(attrsUsed, ", ")))
	}
	if len(envsUsed) > 0 {
		buf = append(buf, fmt.Sprintf("Env: %s", strings.Join(envsUsed, ", ")))
	}
	return strings.Join(buf, ". ")
}

func (a attributes) Validate(cfg *Config) error {
	authsUsed := map[string]bool{}
	for _, attr := range a {
		if attr.IsZero(cfg) {
			continue
		}
		if attr.Auth == "" {
			continue
		}
		authsUsed[attr.Auth] = true
	}
	if len(authsUsed) <= 1 {
		return nil
	}
	if cfg.AuthType != "" {
		// client has auth preference set
		return nil
	}
	names := []string{}
	for v := range authsUsed {
		names = append(names, v)
	}
	sort.Strings(names)
	return fmt.Errorf("more than one authorization method configured: %s",
		strings.Join(names, " and "))
}

// Name implements Loader interface for environment variables
func (a attributes) Name() string {
	return "environment"
}

// Configure implements Loader interface for environment variables
func (a attributes) Configure(cfg *Config) error {
	for _, attr := range a {
		if !attr.IsZero(cfg) {
			// don't overwtite a value previously set
			continue
		}
		v := attr.ReadEnv()
		if v == "" {
			continue
		}
		err := attr.SetS(cfg, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a attributes) ResolveFromStringMap(cfg *Config, m map[string]string) error {
	for _, attr := range a {
		if !attr.IsZero(cfg) {
			// don't overwtite a value previously set
			continue
		}
		v, ok := m[attr.Name]
		if !ok || v == "" {
			continue
		}
		err := attr.SetS(cfg, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a attributes) ResolveFromAnyMap(cfg *Config, m map[string]interface{}) error {
	for _, attr := range a {
		if !attr.IsZero(cfg) {
			// don't overwtite a value previously set
			continue
		}
		v, ok := m[attr.Name]
		if !ok {
			continue
		}
		err := attr.Set(cfg, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func loadAttrs() (attrs attributes) {
	t := reflect.TypeOf(Config{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		nameTag := field.Tag.Get("name")
		if nameTag == "" {
			continue
		}
		sensitive := false
		auth := field.Tag.Get("auth")
		authSplit := strings.Split(auth, ",")
		if len(authSplit) == 2 {
			auth = authSplit[0]
			sensitive = authSplit[1] == "sensitive"
		}
		// internal config fields are skipped in debugging
		internal := false
		if auth == "-" {
			auth = ""
			internal = true
		}
		attr := ConfigAttribute{
			Name:      nameTag,
			Auth:      auth,
			Kind:      field.Type.Kind(),
			Sensitive: sensitive,
			Internal:  internal,
			num:       i,
		}
		envTag := field.Tag.Get("env")
		if envTag != "" {
			attr.EnvVars = strings.Split(envTag, ",")
		}
		attrs = append(attrs, attr)
	}
	return
}
