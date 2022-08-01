package databricks

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ConfigAttribute provides generic way to work with DatabricksClient configuration
// attributes and parses `name`, `env`, and `auth` field tags.
type ConfigAttribute struct {
	Name      string
	Kind      reflect.Kind
	EnvVars   []string
	Sensitive bool
	where     string
	num       int
}

func (ca *ConfigAttribute) ReadEnv() (string, bool) {
	for _, envName := range ca.EnvVars {
		v := os.Getenv(envName)
		if v == "" {
			continue
		}
		return v, true
	}
	return "", false
}

func (ca *ConfigAttribute) SetS(rv reflect.Value, v string) error {
	switch ca.Kind {
	case reflect.String:
		return ca.Set(rv, v)
	case reflect.Bool:
		vv, err := strconv.ParseBool(v)
		if err != nil {
			return err
		}
		return ca.Set(rv, vv)
	case reflect.Int:
		vv, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		return ca.Set(rv, vv)
	default:
		return fmt.Errorf("cannot set %s of unknown type %s", ca.Name, reflectKind(ca.Kind))
	}
}

func (ca *ConfigAttribute) Set(rv reflect.Value, i interface{}) error {
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem() // pointer deref
	}
	field := rv.Field(ca.num)
	switch ca.Kind {
	case reflect.String:
		field.SetString(i.(string))
	case reflect.Bool:
		field.SetBool(i.(bool))
	case reflect.Int:
		field.SetInt(int64(i.(int)))
	default:
		// must extensively test with providerFixture to avoid this one
		return fmt.Errorf("cannot set %s of unknown type %s", ca.Name, reflectKind(ca.Kind))
	}
	return nil
}

func (ca *ConfigAttribute) IsZero(credential interface{}) bool {
	rv := reflect.ValueOf(credential)
	field := rv.Elem().Field(ca.num)
	return field.IsZero()
}

func (ca *ConfigAttribute) GetString(client interface{}) string {
	rv := reflect.ValueOf(client)
	field := rv.Elem().Field(ca.num)
	return fmt.Sprintf("%v", field.Interface())
}

func getTemplateType(template interface{}) reflect.Type {
	t := reflect.TypeOf(template)
	if t.Kind() == reflect.Pointer {
		return t.Elem()
	}
	return t
}

func attributesOf(template interface{}, where string) (attrs []ConfigAttribute) {
	t := getTemplateType(template)
	// t := reflect.ValueOf(template).Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		nameTag := field.Tag.Get("name")
		if nameTag == "" {
			continue
		}
		attr := ConfigAttribute{
			Name:      nameTag,
			Kind:      field.Type.Kind(),
			Sensitive: strings.Contains(field.Tag.Get("auth"), "sensitive"),
			where:     where,
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
