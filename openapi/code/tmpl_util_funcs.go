package code

import (
	"reflect"
	"strings"
	"text/template"
)

var HelperFuncs = template.FuncMap{
	"notLast": func(idx int, a interface{}) bool {
		return idx+1 != reflect.ValueOf(a).Len()
	},
	"trim_prefix": func(left, right string) string {
		return strings.TrimPrefix(left, right)
	},
	"without": func(left, right string) string {
		return strings.ReplaceAll(right, left, "")
	},
}
