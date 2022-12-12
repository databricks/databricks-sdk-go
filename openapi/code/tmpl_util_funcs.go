package code

import (
	"errors"
	"reflect"
	"strings"
	"text/template"
)

var ErrSkipThisFile = errors.New("skip generating this file")

var HelperFuncs = template.FuncMap{
	"notLast": func(idx int, a interface{}) bool {
		return idx+1 != reflect.ValueOf(a).Len()
	},
	"lower": strings.ToLower,
	"trimPrefix": func(right, left string) string {
		return strings.TrimPrefix(left, right)
	},
	"trimSuffix": func(right, left string) string {
		return strings.TrimSuffix(left, right)
	},
	"without": func(left, right string) string {
		return strings.ReplaceAll(right, left, "")
	},
	"skipThisFile": func() error {
		// error is rendered as string in the resulting file, so we must panic,
		// so that we handle this error in [gen.Pass[T].File] gracefully
		// via errors.Is(err, code.ErrSkipThisFile)
		panic(ErrSkipThisFile)
	},
}
