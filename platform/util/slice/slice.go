package slice

import "github.com/2637309949/dolphin/platform/util/snaker"

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
