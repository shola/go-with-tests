package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.String:
		fn(val.String())
	}

	// numberOfValues := 0
	// var getField func(int) reflect.Value

	// switch val.Kind() {
	// case reflect.Struct:
	// 	numberOfValues = val.NumField()
	// 	getField = val.Field
	// case reflect.Map:
	// 	for _, key := range val.MapKeys() {
	// 		walk(val.MapIndex(key).Interface(), fn)
	// 	}
	// case reflect.Slice, reflect.Array:
	// 	numberOfValues = val.Len()
	// 	getField = val.Index
	// case reflect.String:
	// 	fn(val.String())
	// }

	// for i := 0; i < numberOfValues; i++ {
	// 	walk(getField(i).Interface(), fn)
	// }
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}