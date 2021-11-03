package gin

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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
