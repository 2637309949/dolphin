package util

import (
	"strings"

	"github.com/2637309949/dolphin/packages/gin"
)

// ProcessMethodOverride check POST request and re-dispatch them if _method
// exists in POST body, this was intent to act as Ruby MethodOverride rack
func ProcessMethodOverride(r *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		// only need to check POST method
		if c.Request.Method != "POST" {
			return
		}
		// only deal with request where _method exists
		method := c.PostForm("_method")
		if method == "" {
			return
		}
		// these are known verb used frequently
		method = strings.ToLower(method)
		switch method {
		case "delete":
			c.Request.Method = "DELETE"

		case "put":
			c.Request.Method = "PUT"

		case "patch":
			c.Request.Method = "PATCH"
		// if user request method we don't know, then just pass this request to
		// next middle-ware, then this will treat as plain POST request
		default:
			return
		}
		// after we rewrite this request, we need pass to gin engine for routing again, otherwise, this rewrite route will fail to 404
		r.HandleContext(c)
		return
	}
}
