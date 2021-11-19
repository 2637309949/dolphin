package gin

import (
	"net/http"
	"strings"
	"reflect"

	"github.com/gin-gonic/gin"
)

// IsFunction defined TODO
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

// JsonCheck defined TODO
func JsonCheck(req *http.Request) bool {
	jsonType := "application/json"
	if req.ContentLength > 0 && strings.Contains(req.Header.Get("Content-Type"), jsonType) {
		return true
	}
	return false
}

// UriCheck defined TODO
func UriCheck(params gin.Params) bool {
	return len(params) > 0
}

// QueryCheck defined TODO
func QueryCheck(req *http.Request) bool {
	return req.Method == http.MethodGet || req.Method == http.MethodDelete
}
