package plugin

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/go-errors/errors"
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

// RecoveryWithWriter defined
func RecoveryWithWriter(f func(c *gin.Context, err interface{}), out io.Writer) gin.HandlerFunc {
	var logger *log.Logger
	if out != nil {
		logger = log.New(out, "\n\n\x1b[31m", log.LstdFlags)
	}
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if logger != nil {
					httprequest, _ := httputil.DumpRequest(c.Request, false)
					goErr := errors.Wrap(err, 3)
					reset := string([]byte{27, 91, 48, 109})
					logger.Printf("[Nice Recovery] panic recovered:\n\n%s%s\n\n%s%s", httprequest, goErr.Error(), goErr.Stack(), reset)
				}

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
		ctx.JSON(http.StatusOK, model.Response{
			Code: null.IntFrom(int64(code)),
			Msg:  null.StringFrom(msg),
		})
	}, gin.DefaultErrorWriter)
}

// Formatter log formatter params
func Formatter(param gin.LogFormatterParams) string {
	if param.Latency > time.Minute {
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("%3d | %13v | %15s |%-7s %s%s",
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
}

// LoggerWithConfig defined logger config
func LoggerWithConfig() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Output:    logrus.GetOutput(),
		Formatter: Formatter,
	})
}
