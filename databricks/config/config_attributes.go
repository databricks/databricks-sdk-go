package config

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

var ConfigAttributes = loadAttrs()

// Attributes holds meta-representation of Config configuration options
type Attributes []ConfigAttribute

type environmentVariableLoader struct{}

func (l environmentVariableLoader) Name() string {
	return "environment"
}

func (l environmentVariableLoader) Configure(cfg *Config) error {
	for _, attr := range ConfigAttributes {
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

func (a Attributes) DebugString(cfg *Config) string {
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

func (a Attributes) Validate(cfg *Config) error {
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

func (a Attributes) ResolveFromStringMap(cfg *Config, m map[string]string) error {
	for _, attr := range a {
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

// Merge takes attributes from src and sets them on dst.
// Attributes that are already set on dst are skipped.
func (a Attributes) Merge(dst *Config, src *Config) error {
	for _, attr := range a {
		// Ignore if this attribute isn't set at the source.
		if attr.IsZero(src) {
			continue
		}

		// Ignore if this attribute is already set at the destination.
		// If so, the configuration struct was populated with values
		// by the caller. These always take precedence over attributes
		// that were set through the environment.
		if !attr.IsZero(dst) {
			continue
		}

		err := attr.Set(dst, attr.Get(src))
		if err != nil {
			return err
		}
	}
	return nil
}

func loadAttrs() (attrs Attributes) {
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
