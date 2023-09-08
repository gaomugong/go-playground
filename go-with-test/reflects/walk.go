package reflects

import (
	"reflect"
)

//func walk(x interface{}, fn func(input string)) {
//	val := reflect.ValueOf(x)
//	field := val.Field(0)
//	fn(field.String())
//}

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
			fn(field.String())
		}
		if field.Kind() == reflect.Struct {
			//fmt.Println("find struct field")
			walk(field.Interface(), fn)
		}
	}
}
