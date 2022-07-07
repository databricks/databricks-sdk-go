package implicit

import (
	"fmt"
	"reflect"
	"strings"
)

// ConfigAttribute provides generic way to work with DatabricksClient configuration
// attributes and parses `name`, `env`, and `auth` field tags.
type ConfigAttribute struct {
	Name        string
	Kind        reflect.Kind
	EnvVars     []string
	Credentials []string
	Auth        string
	Sensitive   bool
	Internal    bool
	num         int
}

func (ca *ConfigAttribute) Set(credential interface{}, i interface{}) error {
	rv := reflect.ValueOf(credential)
	field := rv.Elem().Field(ca.num)
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

func attributesOf(credential interface{}, credentialsName string) (attrs []ConfigAttribute) {
	t := reflect.TypeOf(credential)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Struct {
			inner := reflect.New(field.Type).Elem().Interface()
			attrs = append(attrs, attributesOf(inner, credentialsName)...)
			continue
		}
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
			Name:        nameTag,
			Auth:        auth,
			Kind:        field.Type.Kind(),
			Credentials: []string{credentialsName},
			Sensitive:   sensitive,
			Internal:    internal,
			num:         i,
		}
		envTag := field.Tag.Get("env")
		if envTag != "" {
			attr.EnvVars = strings.Split(envTag, ",")
		}
		attrs = append(attrs, attr)
	}
	return
}
