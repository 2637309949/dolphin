package slice

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/2637309949/dolphin/platform/util/snaker"
)

// StrSliceContains check string exists in array
func StrSliceContains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}

// CamelToSnake change array string to snaker
func CamelToSnake(ss []string) []string {
	result := make([]string, 0)
	for _, s := range ss {
		result = append(result, snaker.CamelToSnake(s))
	}
	return result
}

// GetFieldSliceByName defined
func GetFieldSliceByName(arr interface{}, name string, format ...string) interface{} {
	if !IsIteratee(arr) {
		panic("First parameter must be an iteratee")
	}

	var (
		ft = "%v"
	)
	if len(format) > 0 {
		ft = format[0]
	}

	sm := []map[string]interface{}{}
	bte, err := json.Marshal(arr)
	if err != nil {
		return nil
	}
	if err = json.Unmarshal(bte, &sm); err != nil || len(sm) == 0 {
		return nil
	}
	var (
		arrElemType     = reflect.TypeOf(sm[0][name])
		resultSliceType = reflect.SliceOf(arrElemType)
		resultSlice     = reflect.MakeSlice(resultSliceType, 0, 0)
	)

	for i := 0; i < len(sm); i++ {
		item := sm[i][name]
		switch item.(type) {
		case string:
			item = fmt.Sprintf(ft, item)
		}
		iv := reflect.ValueOf(item)
		resultSlice = reflect.Append(resultSlice, iv)
	}
	return resultSlice.Interface()
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
