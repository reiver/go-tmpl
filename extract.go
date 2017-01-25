package tmpl

import (
	"fmt"
	"reflect"
)

const (
	nameKeyShort  = "tmpl"
	nameKeyNormal = "tmpl.name"
)

// extract takes a struct, and returns it into a map[string][]interface{},
// where the keys of the map are the field names from the struct (or the value
// of the `tmpl.name` or `tmpl` tags, in those are available).
//
// Note that the value of the map is the slice []interface{}, and not [interface{}.
// This is so we can detect duplicate values.
func extract(v interface{}) (map[string][]interface{}, error) {

	if nil == v {
		return map[string][]interface{}{}, nil
	}

	var reflectedStructValue reflect.Value
	var reflectedStructType  reflect.Type
	{
		// This needs to get at the struct.
		reflectedValue := reflect.ValueOf(v)
		reflectedType  := reflect.TypeOf(v)
		if nil == reflectedType {
			return nil, errNilReflectedType
		}
		for reflect.Ptr == reflectedValue.Kind() {
			reflectedValue = reflectedValue.Elem()
			reflectedType = reflectedType.Elem()
			if nil == reflectedType {
				return nil, errNilReflectedType
			}
		}
		if reflect.Struct != reflectedType.Kind() {
			return nil, fmt.Errorf("Unsupported Type: %T", v)
		}

		reflectedStructValue = reflectedValue
		reflectedStructType  = reflectedType
	}


	data := map[string][]interface{}{}

	numFields := reflectedStructType.NumField()
	for fieldNumber:=0; fieldNumber<numFields; fieldNumber++ {

		reflectedStructField := reflectedStructType.Field(fieldNumber)

		// Figure out the "name" the user wants to use.
		name, ok := reflectedStructField.Tag.Lookup(nameKeyNormal)
		if !ok {
			name, ok = reflectedStructField.Tag.Lookup(nameKeyShort)
			if !ok {
				name = reflectedStructField.Name
			}
		}

		// Figure out the value the user wants to use.
		reflectedFieldValue := reflectedStructValue.Field(fieldNumber)
		if !reflectedFieldValue.CanInterface() {
			return nil, fmt.Errorf("For type %T and field %q, could not extract value.", v, reflectedStructField.Name)
		}
		value := reflectedFieldValue.Interface()

		// Append the value under the name.
		if _, ok := data[name]; !ok {
			data[name] = []interface{}{}
		}
		data[name] = append(data[name], value)
	}

	return data, nil
}
