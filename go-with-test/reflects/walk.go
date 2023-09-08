package reflects

import (
	"reflect"
)

//func walk(x interface{}, fn func(input string)) {
//	val := reflect.ValueOf(x)
//	field := val.Field(0)
//	fn(field.String())
//}

//func walk(x interface{}, fn func(input string)) {
//	val := reflect.ValueOf(x)
//
//	for i := 0; i < val.NumField(); i++ {
//		field := val.Field(i)
//
//		if field.Kind() == reflect.String {
//			fn(field.String())
//		}
//		if field.Kind() == reflect.Struct {
//			//fmt.Println("find struct field")
//			walk(field.Interface(), fn)
//		}
//	}
//}

// walk 递归提取struct中所有字符串字段
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		// Elem() 提取底层值
		val = val.Elem()
	}

	return val
}
