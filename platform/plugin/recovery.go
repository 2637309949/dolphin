package plugin

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/2637309949/dolphin/platform/model"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
)

// Recovery defined gin recovery middle
func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httprequest, _ := httputil.DumpRequest(ctx.Request, false)
				goErr := errors.Wrap(err, 3)
				reset := string([]byte{27, 91, 48, 109})
				logrus.Errorf("[Nice Recovery] panic recovered:\n\n%s%s\n\n%s%s", httprequest, goErr.Error(), goErr.Stack(), reset)
				code := 500
				msg := fmt.Sprintf("%v", err)
				if err, ok := err.(model.Error); ok {
					code = err.Code
				}
				ctx.JSON(http.StatusOK, model.Fail{
					Code: code,
					Msg:  msg,
				})
			}
		}()
		ctx.Next()
	}
}
