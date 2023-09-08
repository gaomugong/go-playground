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

//// walk 递归提取struct中所有字符串字段
//func walk(x interface{}, fn func(input string)) {
//	val := getValue(x)
//
//	if val.Kind() == reflect.Slice {
//		for i := 0; i < val.Len(); i++ {
//			walk(val.Index(i).Interface(), fn)
//		}
//		return
//	}
//
//	for i := 0; i < val.NumField(); i++ {
//		field := val.Field(i)
//		switch field.Kind() {
//		case reflect.String:
//			fn(field.String())
//		case reflect.Struct:
//			walk(field.Interface(), fn)
//		}
//	}
//}

//// walk 递归提取struct中所有字符串字段
//func walk(x interface{}, fn func(input string)) {
//	val := getValue(x)
//
//	// 如果是 struct 或切片，我们会遍历它的值，并对每个值调用 walk 函数。
//	// 如果是 reflect.String，我们就调用 fn
//	switch val.Kind() {
//	case reflect.Struct:
//		for i := 0; i < val.NumField(); i++ {
//			walk(val.Field(i).Interface(), fn)
//		}
//	case reflect.Slice:
//		for i := 0; i < val.Len(); i++ {
//			walk(val.Index(i).Interface(), fn)
//		}
//	case reflect.String:
//		fn(val.String())
//	}
//}

//func walk(x interface{}, fn func(input string)) {
//	val := getValue(x)
//
//	fieldLen := 0
//	var getField func(i int) reflect.Value
//
//	// 如果是 struct 或切片，我们会遍历它的值，并对每个值调用 walk 函数。
//	// 如果是 reflect.String，我们就调用 fn
//	switch val.Kind() {
//	case reflect.Map:
//		for _, key := range val.MapKeys() {
//			walk(val.MapIndex(key).Interface(), fn)
//		}
//	case reflect.Struct:
//		fieldLen, getField = val.NumField(), val.Field
//	case reflect.Slice, reflect.Array:
//		fieldLen, getField = val.Len(), val.Index
//	case reflect.String:
//		fn(val.String())
//	}
//
//	for i := 0; i < fieldLen; i++ {
//		walk(getField(i).Interface(), fn)
//	}
//}

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	// 如果是 struct 或切片，我们会遍历它的值，并对每个值调用 walk 函数。
	// 如果是 reflect.String，我们就调用 fn
	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
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
