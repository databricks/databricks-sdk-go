package marshal

import (
	"reflect"
	"strings"
	"sync"
)

var mutexType sync.Mutex

var typeCache = map[reflect.Type][]cachedType{}

var mutexName sync.Mutex

var nameCache = map[reflect.Type]string{}

type cachedType struct {
	reflect.StructField
	JsonTag       jsonTag
	IndexInStruct int
}

func getTypeFields(structType reflect.Type) []cachedType {
	if res, ok := typeCache[structType]; ok {
		return res
	}
	res := []cachedType{}
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		jsonTag := parseJSONTag(field)
		res = append(res, cachedType{
			StructField:   field,
			JsonTag:       jsonTag,
			IndexInStruct: i,
		})
	}
	mutexType.Lock()
	typeCache[structType] = res
	mutexType.Unlock()
	return res
}

func getTypeName(structType reflect.Type) string {
	if res, ok := nameCache[structType]; ok {
		return res
	}
	name := structType.Name()
	mutexName.Lock()
	nameCache[structType] = name
	mutexName.Unlock()
	return name
}

func parseJSONTag(field reflect.StructField) jsonTag {
	raw := field.Tag.Get("json")
	name := field.Name

	if raw == "-" {
		return jsonTag{ignore: true}
	}

	parts := strings.Split(raw, ",")

	if parts[0] != "" {
		name = parts[0]
	}

	jsonTag := jsonTag{
		name: name,
	}

	for _, v := range parts {
		switch v {
		case "omitempty":
			jsonTag.omitempty = true
		case "string":
			jsonTag.asString = true
		}
	}
	return jsonTag
}
