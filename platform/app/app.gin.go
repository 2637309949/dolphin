package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/model"
)

// GinProcessMethodOverride check POST request and re-dispatch them if _method
// exists in POST body, this was intent to act as Ruby MethodOverride rack
var GinProcessMethodOverride = func(r *gin.Engine) gin.HandlerFunc {
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

// GinRecovery defined gin recovery middle
var GinRecovery = Recovery(func(ctx *gin.Context, err interface{}) {
	code := 500
	msg := fmt.Sprintf("%v", err)
	if err, ok := err.(model.Error); ok {
		code = err.Code
	}
	ctx.JSON(http.StatusOK, model.Response{
		Code: null.IntFrom(int64(code)),
		Msg:  null.StringFrom(msg),
	})
})

// GinLoggerWithConfig defined logger config
var GinLoggerWithConfig = gin.LoggerWithConfig(gin.LoggerConfig{
	Output:    logrus.GetOutput(),
	Formatter: ginFormatter,
})

func init() {

}
