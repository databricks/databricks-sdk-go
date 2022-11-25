package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
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
	Sensitive bool
	Internal  bool
	num       int
}

func (a *ConfigAttribute) ReadEnv() string {
	for _, envName := range a.EnvVars {
		v := os.Getenv(envName)
		if v == "" {
			continue
		}
		return v
	}
	return ""
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
	default:
		return fmt.Errorf("cannot set %s of unknown type %s",
			a.Name, reflectKind(a.Kind))
	}
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
