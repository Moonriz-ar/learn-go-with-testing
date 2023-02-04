package reflection

import "reflect"

// reflection in computing is the ability of a program to examine its own structure, particularly through types.

func Walk(x interface{}, fn func(input string)) {
	val := GetValue(x)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func GetValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// you can't use NumField on a pointer Value, we need to extrct the underlying value before we can do that by using Elem()
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
