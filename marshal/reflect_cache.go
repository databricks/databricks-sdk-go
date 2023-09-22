package marshal

import "reflect"

var typeCache = map[reflect.Type]*[]cachedType{}

var nameCache = map[reflect.Type]string{}

type cachedType struct {
	reflect.StructField
	JsonTag string
}

func getTypeFields(structType reflect.Type) *[]cachedType {
	if res, ok := typeCache[structType]; ok {
		return res
	}
	res := []cachedType{}
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		res = append(res, cachedType{
			StructField: field,
			JsonTag:     field.Tag.Get("json"),
		})
	}
	typeCache[structType] = &res
	return &res
}

func getTypeName(structType reflect.Type) string {
	if res, ok := nameCache[structType]; ok {
		return res
	}
	name := structType.Name()
	nameCache[structType] = name
	return name
}
