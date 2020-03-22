package plugin

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/go-errors/errors"
	"github.com/mattn/go-isatty"
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

// defaultLogFormatter is the default log format function Logger middleware uses.
var defaultLogFormatter = func(param gin.LogFormatterParams) string {
	var methodColor, resetColor string
	if param.IsOutputColor() {
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	if param.Latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		param.Latency = param.Latency - param.Latency%time.Second
	}
	return fmt.Sprintf("%3d | %13v | %15s |%s %-7s %s %s%s",
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.ErrorMessage,
	)
}

// LogFormatterParams is the structure any formatter will be handed when time to log comes
type LogFormatterParams struct {
	gin.LogFormatterParams
	Domain  string
	Token   string
	UserID  string
	Header  []byte
	ReqBody []byte
	ResBody []byte
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// Tracker instance a Logger middleware with config.
func Tracker(conf gin.LoggerConfig, cb ...func(*gin.Context, *LogFormatterParams)) gin.HandlerFunc {
	formatter := conf.Formatter
	if formatter == nil {
		formatter = defaultLogFormatter
	}

	out := conf.Output
	if out == nil {
		out = gin.DefaultWriter
	}

	notlogged := conf.SkipPaths

	isTerm := true

	if w, ok := out.(*os.File); !ok || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(w.Fd()) && !isatty.IsCygwinTerminal(w.Fd())) {
		isTerm = false
	}

	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		req := c.Request.Clone(c)
		writer := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		// Request byte
		body, _ := ioutil.ReadAll(req.Body)
		c.Writer = writer

		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// resove header
		hr := bytes.Buffer{}
		req.Write(&hr)

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			param := LogFormatterParams{
				LogFormatterParams: gin.LogFormatterParams{
					Request: c.Request,
					IsTerm:  isTerm,
					Keys:    c.Keys,
				},
				Header:  hr.Bytes(),
				ReqBody: body,
				ResBody: writer.body.Bytes(),
			}

			// Stop timer
			param.TimeStamp = time.Now()
			param.Latency = param.TimeStamp.Sub(start)

			param.ClientIP = c.ClientIP()
			param.Method = c.Request.Method
			param.StatusCode = c.Writer.Status()
			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()

			param.BodySize = c.Writer.Size()

			if raw != "" {
				path = path + "?" + raw
			}

			param.Path = path

			logrus.Info(formatter(param.LogFormatterParams))

			// LogFormatterParams defined recorder
			if len(cb) > 0 {
				cb[0](c, &param)
			}
		}
	}
}
