package gin

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/2637309949/dolphin/packages/web/core"
	"github.com/eriklott/mustache"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
)

var _ core.Context = &Context{}

type Context struct {
	*gin.Context
}

// Set defined TODO
func (c *Context) Set(k string, v interface{}) {
	c.Context.Set(k, v)
}

// Get defined TODO
func (c *Context) Get(k string) (interface{}, bool) {
	return c.Context.Get(k)
}

// Param defined TODO
func (c *Context) Param(k string) string {
	return c.Context.Param(k)
}

// Query defined TODO
func (c *Context) Query(k string) string {
	return c.Context.Query(k)
}

// MultipartForm defined TODO
func (c *Context) MultipartForm() (*multipart.Form, error) {
	return c.Context.MultipartForm()
}

// Header defined TODO
func (c *Context) Header(key, value string) {
	c.Context.Header(key, value)
}

// GetHeader defined TODO
func (c *Context) GetHeader(key string) string {
	return c.Context.GetHeader(key)
}

// ShouldBindWith defined TODO
func (c *Context) ShouldBindWith(v interface{}) error {
	if UriCheck(c.Context.Params) {
		if err := c.Context.ShouldBindUri(v); err != nil {
			return err
		}
	}
	if QueryCheck(c.Context.Request) {
		if err := c.Context.ShouldBindQuery(v); err != nil {
			return err
		}
	}
	if JsonCheck(c.Context.Request) {
		if err := c.ShouldBindBody(v); err != nil {
			return err
		}
	}
	return nil
}

// ShouldBindBody defined TODO
func (c *Context) ShouldBindBody(ptr interface{}) error {
	var err error
	var body []byte

	if ptr == nil {
		return errors.New("first parameter must be an pointer")
	}
	if cb, ok := c.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			body = cbb
		}
	}
	if body == nil {
		body, err = ioutil.ReadAll(c.Context.Request.Body)
		if err != nil {
			return err
		}
		c.Set(gin.BodyBytesKey, body)
	}
	return binding.JSON.BindBody(body, ptr)
}

// GetRawData defined TODO
func (c *Context) GetRawData() ([]byte, error) {
	return c.Context.GetRawData()
}

// SetCookie defined TODO
func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	c.Context.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

// Cookie defined TODO
func (c *Context) Cookie(name string) (string, error) {
	return c.Context.Cookie(name)
}

// Next defined TODO
func (c *Context) Next() {
	c.Context.Next()
}

// Abort defined TODO
func (c *Context) Abort() {
	c.Context.Abort()
}

// IsAborted defined TODO
func (c *Context) IsAborted() bool {
	return c.Context.IsAborted()
}

// Success defined TODO
func (c *Context) Success(data interface{}) {
	c.Context.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"data": data,
	})
}

// Fail defined TODO
func (c *Context) Fail(err error) {
	c.Context.JSON(http.StatusOK, map[string]interface{}{
		"code":   500,
		"detail": err.Error(),
	})
}

// HTML defined TODO
func (c *Context) HTML(r io.Reader, context ...interface{}) {
	c.Context.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	file, err := ioutil.TempFile("", "*")
	if err != nil {

		return
	}
	defer os.Remove(file.Name())
	bte, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	str, err := mustache.NewTemplate().Render(string(bte), context...)
	if err != nil {
		return
	}
	if _, err = file.WriteString(str); err != nil {
		return
	}
	http.ServeFile(c.Context.Writer, c.Context.Request, file.Name())
}

// XML defined TODO
func (c *Context) XML(r io.Reader, context ...interface{}) {
	c.Context.Writer.Header().Set("Content-Type", "application/xml; charset=utf-8")
	file, err := ioutil.TempFile("", "*")
	if err != nil {
		return
	}
	defer os.Remove(file.Name())
	bte, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	str, err := mustache.NewTemplate().Render(string(bte), context...)
	if err != nil {
		return
	}
	if _, err = file.WriteString(str); err != nil {
		return
	}
	http.ServeFile(c.Context.Writer, c.Context.Request, file.Name())
}

// XML defined TODO
func (c *Context) String(data string, context ...interface{}) {
	str, err := mustache.NewTemplate().Render(data, context...)
	if err != nil {
		return
	}
	c.Context.String(200, str)
}

// Request defined TODO
func (c *Context) Request() *http.Request {
	return c.Context.Request
}

// Request defined TODO
func (c *Context) ResponseWriter() http.ResponseWriter {
	return c.Context.Writer
}

// Request defined TODO
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Deadline()
}

// Done defined TODO
func (c *Context) Done() <-chan struct{} {
	return c.Context.Done()
}

// Err defined TODO
func (c *Context) Err() error {
	return c.Context.Err()
}

// Value defined TODO
func (c *Context) Value(key interface{}) interface{} {
	return c.Context.Value(key)
}

// File defined TODO
func (c *Context) File(r io.Reader, filename string, context ...interface{}) {
	c.Context.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	file, err := ioutil.TempFile("", "*")
	if err != nil {
		return
	}
	defer os.Remove(file.Name())
	bte, err := ioutil.ReadAll(r)
	if err != nil {
		logrus.Error(err)
		return
	}
	str, err := mustache.NewTemplate().Render(string(bte), context...)
	if err != nil {
		logrus.Error(err)
		return
	}
	if _, err = file.WriteString(str); err != nil {
		logrus.Error(err)
		return
	}
	http.ServeFile(c.Context.Writer, c.Context.Request, file.Name())
}

func NewContext(ctx *gin.Context) *Context {
	return &Context{ctx}
}
