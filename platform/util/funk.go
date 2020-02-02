package util

import (
	"reflect"
)

// IsFunction returns if the argument is a function.
func IsFunction(in interface{}, num ...int) bool {
	funcType := reflect.TypeOf(in)
	result := funcType.Kind() == reflect.Func
	if len(num) >= 1 {
		result = result && funcType.NumIn() == num[0]
	}
	if len(num) == 2 {
		result = result && funcType.NumOut() == num[1]
	}
	return result
}

// IsIteratee returns if the argument is an iteratee.
func IsIteratee(in interface{}) bool {
	if in == nil {
		return false
	}
	arrType := reflect.TypeOf(in)
	kind := arrType.Kind()
	return kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map
}

// Filter iterates over elements of collection, returning an array of
// all elements predicate returns truthy for.
func Filter(arr interface{}, predicate interface{}) interface{} {
	if !IsIteratee(arr) {
		panic("First parameter must be an iteratee")
	}
	if !IsFunction(predicate, 1, 1) {
		panic("Second argument must be function")
	}
	funcValue := reflect.ValueOf(predicate)
	funcType := funcValue.Type()
	if funcType.Out(0).Kind() != reflect.Bool {
		panic("Return argument should be a boolean")
	}
	arrValue := reflect.ValueOf(arr)
	arrType := arrValue.Type()
	// Get slice type corresponding to array type
	resultSliceType := reflect.SliceOf(arrType.Elem())
	// MakeSlice takes a slice kind type, and makes a slice.
	resultSlice := reflect.MakeSlice(resultSliceType, 0, 0)
	for i := 0; i < arrValue.Len(); i++ {
		elem := arrValue.Index(i)
		result := funcValue.Call([]reflect.Value{elem})[0].Interface().(bool)
		if result {
			resultSlice = reflect.Append(resultSlice, elem)
		}
	}
	return resultSlice.Interface()
}
