package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Source struct {
	Type SourceType `json:"type"`
	Name string     `json:"name,omitempty"`
}

func (s *Source) String() string {
	if s.Name == "" {
		return string(s.Type)
	}
	return fmt.Sprintf("%s %s", s.Name, s.Type)
}

type SourceType string

const (
	SourceEnv           SourceType = "environment variable"
	SourceFile          SourceType = "config file"
	SourceDynamicConfig SourceType = "dynamic configuration"
)

// ConfigAttribute provides generic way to work with Config configuration
// attributes and parses `name`, `env`, and `auth` field tags.
//
// Internal: this field can become unexported in the future
type ConfigAttribute struct {
	Name      string
	Kind      reflect.Kind
	EnvVars   []string
	Auth      string
	AuthTypes []string
	Sensitive bool
	Internal  bool
	num       int
}

func (a *ConfigAttribute) ReadEnv() (string, string) {
	for _, envName := range a.EnvVars {
		v := os.Getenv(envName)
		if v == "" {
			continue
		}
		return v, envName
	}
	return "", ""
}

func (a *ConfigAttribute) SetS(cfg *Config, v string) error {
	switch a.Kind {
	case reflect.String:
		return a.Set(cfg, v)
	case reflect.Bool:
		vv, err := strconv.ParseBool(v)
		if err != nil {
			return err
		}
		return a.Set(cfg, vv)
	case reflect.Int:
		vv, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		return a.Set(cfg, vv)
	case reflect.Map:
		// Parse map from string format: key1=value1;key2=value2
		vv, err := parseMapFromString(v)
		if err != nil {
			return fmt.Errorf("cannot parse %s: %w", a.Name, err)
		}
		return a.Set(cfg, vv)
	default:
		return fmt.Errorf("cannot set %s of unknown type %s",
			a.Name, reflectKind(a.Kind))
	}
}

// parseMapFromString parses a string of format "key1=value1;key2=value2" into a map.
// Empty keys or values are allowed. If the string is empty, an empty map is returned.
func parseMapFromString(s string) (map[string]string, error) {
	result := make(map[string]string)
	if s == "" {
		return result, nil
	}
	pairs := strings.Split(s, ";")
	for _, pair := range pairs {
		pair = strings.TrimSpace(pair)
		if pair == "" {
			continue
		}
		// Split only on the first '=' to allow '=' in values
		idx := strings.Index(pair, "=")
		if idx == -1 {
			return nil, fmt.Errorf("invalid key-value pair %q: missing '='", pair)
		}
		key := pair[:idx]
		value := pair[idx+1:]
		result[key] = value
	}
	return result, nil
}

func (a *ConfigAttribute) Set(cfg *Config, i interface{}) error {
	rv := reflect.ValueOf(cfg)
	field := rv.Elem().Field(a.num)
	switch a.Kind {
	case reflect.String:
		field.SetString(i.(string))
	case reflect.Bool:
		field.SetBool(i.(bool))
	case reflect.Int:
		field.SetInt(int64(i.(int)))
	case reflect.Map:
		field.Set(reflect.ValueOf(i))
	default:
		// must extensively test with providerFixture to avoid this one
		return fmt.Errorf("cannot set %s of unknown type %s", a.Name, reflectKind(a.Kind))
	}
	return nil
}

func (a *ConfigAttribute) IsZero(cfg *Config) bool {
	rv := reflect.ValueOf(cfg)
	field := rv.Elem().Field(a.num)
	return field.IsZero()
}

func (a *ConfigAttribute) GetString(cfg *Config) string {
	rv := reflect.ValueOf(cfg)
	field := rv.Elem().Field(a.num)
	return fmt.Sprintf("%v", field.Interface())
}

func (a *ConfigAttribute) HasAuthAttribute() bool {
	return a.Auth != ""
}
