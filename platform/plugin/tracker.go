package plugin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/mattn/go-isatty"
)

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
func Tracker(receiver ...func(*gin.Context, *LogFormatterParams)) gin.HandlerFunc {
	notlogged := []string{}
	isTerm := true
	out := logrus.GetOutput()
	formatter := Formatter
	recv := func(*gin.Context, *LogFormatterParams) {}

	if len(receiver) > 0 {
		recv = receiver[0]
	}

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
		// bytes buffer
		buf, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(buf)))
		writer := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// resove header
		hr := bytes.Buffer{}
		c.Request.Header.Write(&hr)

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
				ReqBody: buf,
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
			recv(c, &param)
		}
	}
}
