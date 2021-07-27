package reflect

import (
	"bytes"
	"reflect"
)

// IsBlank defined
func IsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// OrOne defined TODO
func OrOne(arr ...interface{}) interface{} {
	for i := range arr {
		item := arr[i]
		if !IsBlank(reflect.ValueOf(item)) {
			return item
		}
	}
	return reflect.Zero(SliceElem(reflect.TypeOf(arr))).Interface()
}

// SliceElem defined
func SliceElem(rtype reflect.Type) reflect.Type {
	for {
		if rtype.Kind() != reflect.Slice && rtype.Kind() != reflect.Array {
			return rtype
		}

		rtype = rtype.Elem()
	}
}

// RedirectValue defined
func RedirectValue(value reflect.Value) reflect.Value {
	for {
		if !value.IsValid() || (value.Kind() != reflect.Ptr && value.Kind() != reflect.Interface) {
			return value
		}

		res := value.Elem()

		// Test for a circular type.
		if res.Kind() == reflect.Ptr && value.Kind() == reflect.Ptr && value.Pointer() == res.Pointer() {
			return value
		}

		if !res.IsValid() && value.Kind() == reflect.Ptr {
			return reflect.Zero(value.Type().Elem())
		}

		value = res
	}
}

// IsFunction returns if the argument is a function.
func IsFunction(in interface{}, num ...int) bool {
	funcType := reflect.TypeOf(in)
	result := funcType != nil && funcType.Kind() == reflect.Func
	if len(num) >= 1 {
		result = result && funcType.NumIn() == num[0]
	}
	if len(num) == 2 {
		result = result && funcType.NumOut() == num[1]
	}
	return result
}

// IsEqual returns if the two objects are equal
func IsEqual(expected interface{}, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}
	if exp, ok := expected.([]byte); ok {
		act, ok := actual.([]byte)
		if !ok {
			return false
		}
		if exp == nil || act == nil {
			return true
		}
		return bytes.Equal(exp, act)
	}
	return reflect.DeepEqual(expected, actual)
}

// IsType returns if the two objects are in the same type
func IsType(expected interface{}, actual interface{}) bool {
	return IsEqual(reflect.TypeOf(expected), reflect.TypeOf(actual))
}

// Equal returns if the two objects are equal
func Equal(expected interface{}, actual interface{}) bool {
	return IsEqual(expected, actual)
}

// NotEqual returns if the two objects are not equal
func NotEqual(expected interface{}, actual interface{}) bool {
	return !IsEqual(expected, actual)
}
