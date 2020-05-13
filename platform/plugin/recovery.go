package plugin

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/go-errors/errors"
)

// RecoveryWithWriter defined
func RecoveryWithWriter(f func(c *gin.Context, err interface{})) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httprequest, _ := httputil.DumpRequest(c.Request, false)
				goErr := errors.Wrap(err, 3)
				reset := string([]byte{27, 91, 48, 109})
				logrus.Errorf("[Nice Recovery] panic recovered:\n\n%s%s\n\n%s%s", httprequest, goErr.Error(), goErr.Stack(), reset)
				f(c, err)
			}
		}()
		c.Next()
	}
}

// Recovery defined gin recovery middle
func Recovery() gin.HandlerFunc {
	return RecoveryWithWriter(func(ctx *gin.Context, err interface{}) {
		code := 500
		msg := fmt.Sprintf("%v", err)
		if err, ok := err.(model.Error); ok {
			code = err.Code
		}
		ctx.JSON(http.StatusOK, model.Fail{
			Code: code,
			Msg:  msg,
		})
	})
}
