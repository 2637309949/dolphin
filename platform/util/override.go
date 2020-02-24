package util

import (
	"strings"

	"github.com/2637309949/dolphin/packages/gin"
)

// ProcessMethodOverride check POST request and re-dispatch them if _method
// exists in POST body, this was intent to act as Ruby MethodOverride rack
func ProcessMethodOverride(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != "POST" {
			return
		}
		method := c.PostForm("_method")
		if method == "" {
			return
		}
		method = strings.ToLower(method)
		switch method {
		case "delete":
			c.Request.Method = "DELETE"
		case "put":
			c.Request.Method = "PUT"
		case "patch":
			c.Request.Method = "PATCH"
		default:
			return
		}
		r.HandleContext(c)
		return
	}
}
