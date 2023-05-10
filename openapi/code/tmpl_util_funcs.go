package code

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"text/template"
)

var ErrSkipThisFile = errors.New("skip generating this file")

var alphanumRE = regexp.MustCompile(`^\w*$`)

var HelperFuncs = template.FuncMap{
	"notLast": func(idx int, a interface{}) bool {
		return idx+1 != reflect.ValueOf(a).Len()
	},
	"lower": strings.ToLower,
	"lowerFirst": func(s string) string {
		return strings.ToLower(s[0:1]) + s[1:]
	},
	"trimPrefix": func(right, left string) string {
		return strings.TrimPrefix(left, right)
	},
	"trimSuffix": func(right, left string) string {
		return strings.TrimSuffix(left, right)
	},
	"replaceAll": func(from, to, str string) string {
		return strings.ReplaceAll(str, from, to)
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
	"alphanumOnly": func(in []Field) (out []Field) {
		for _, v := range in {
			if !alphanumRE.MatchString(v.Name) {
				continue
			}
			out = append(out, v)
		}
		return out
	},
	"list": func(l ...any) []any {
		return l
	},
	"in": func(haystack []any, needle string) bool {
		for _, v := range haystack {
			if needle == fmt.Sprint(v) {
				return true
			}
		}
		return false
	},
	"dict": func(args ...any) map[string]any {
		if len(args)%2 != 0 {
			panic("number of arguments to dict is not even")
		}
		result := map[string]any{}
		for i := 0; i < len(args); i += 2 {
			k := fmt.Sprint(args[i])
			v := args[i+1]
			result[k] = v
		}
		return result
	},
	"getOrDefault": func(dict map[string]any, key string, def any) any {
		v, ok := dict[key]
		if ok {
			return v
		}
		return def
	},
	"fmt": fmt.Sprintf,
	"concat": func(v ...string) string {
		return strings.Join(v, "")
	},
}
