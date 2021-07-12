package slice

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

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
	if err = json.Unmarshal(bte, &sm); err != nil || len(sm) == 0 || sm[0][name] == nil {
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

// PatchSliceByField defined
func PatchSliceByField(target interface{}, source interface{}, targetMatchKey string, sourceMatchKey string, key ...string) ([]map[string]interface{}, error) {
	if !IsIteratee(target) {
		panic("First parameter must be an iteratee")
	}

	if !IsIteratee(source) {
		panic("Second parameter must be an iteratee")
	}

	tsm := []map[string]interface{}{}
	bte, err := json.Marshal(target)
	if err != nil {
		return []map[string]interface{}{}, err
	}

	if err = json.Unmarshal(bte, &tsm); err != nil || len(tsm) == 0 {
		return []map[string]interface{}{}, err
	}

	ssm := []map[string]interface{}{}
	ste, err := json.Marshal(source)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	if err = json.Unmarshal(ste, &ssm); err != nil {
		return []map[string]interface{}{}, err
	}

	for i := range tsm {
		for i2 := range ssm {
			if tsm[i][targetMatchKey] == ssm[i2][sourceMatchKey] {
				for i3 := range key {
					keys := strings.Split(key[i3], "#")
					if len(keys) > 1 {
						tsm[i][keys[0]] = ssm[i2][keys[1]]
					} else {
						tsm[i][keys[0]] = ssm[i2][keys[0]]
					}
				}
			}
		}
	}
	return tsm, err
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

// StringInSlice defined
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// RemoveStringDuplicates defined
func RemoveStringDuplicates(s []string) ([]string, []string) {
	encountered := make(map[string]struct{})
	result := make([]string, 0)
	duplicate := make([]string, 0)
	for _, v := range s {
		if _, ok := encountered[v]; ok {
			duplicate = append(duplicate, v)
			continue
		} else {
			encountered[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result, duplicate
}

func redirectValue(value reflect.Value) reflect.Value {
	for {
		if !value.IsValid() || value.Kind() != reflect.Ptr {
			return value
		}

		res := reflect.Indirect(value)

		// Test for a circular type.
		if res.Kind() == reflect.Ptr && value.Pointer() == res.Pointer() {
			return value
		}

		value = res
	}
}

func sliceElem(rtype reflect.Type) reflect.Type {
	for {
		if rtype.Kind() != reflect.Slice && rtype.Kind() != reflect.Array {
			return rtype
		}

		rtype = rtype.Elem()
	}
}

// Chunk defined
func Chunk(slice interface{}, chunkSize int) interface{} {
	arrValue := redirectValue(reflect.ValueOf(slice))
	if !IsIteratee(arrValue.Interface()) {
		panic("First parameter must be an iteratee")
	}
	accValue := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(slice)), 0, 0)
	for i := 0; i < arrValue.Len(); i += chunkSize {
		end := i + chunkSize
		if end > arrValue.Len() {
			end = arrValue.Len()
		}
		accValue = reflect.Append(accValue, arrValue.Slice(i, end))
	}
	return accValue.Interface()
}
