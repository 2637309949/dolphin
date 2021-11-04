package echo

import (
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/2637309949/dolphin/packages/web/core"
	"github.com/eriklott/mustache"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var _ core.Context = &Context{}

type Context struct {
	echo.Context
	next func()
}

// Set defined TODO
func (c *Context) Set(k string, v interface{}) {
	c.Context.Set(k, v)
}

// Get defined TODO
func (c *Context) Get(k string) interface{} {
	return c.Context.Get(k)
}

// Param defined TODO
func (c *Context) Param(k string) string {
	return c.Context.Param(k)
}

// Query defined TODO
func (c *Context) Query(k string) string {
	return c.Context.QueryParam(k)
}

// MultipartForm defined TODO
func (c *Context) MultipartForm() (*multipart.Form, error) {
	return c.Context.MultipartForm()
}

// Header defined TODO
func (c *Context) Header(key, value string) {
	c.Context.Response().Writer.Header().Set(key, value)
}

// GetHeader defined TODO
func (c *Context) GetHeader(key string) string {
	return c.Context.Request().Header.Get(key)
}

// Request defined TODO
func (c *Context) SetRequest(r *http.Request) {
	c.Context.SetRequest(r)
}

// Request defined TODO
func (c *Context) Request() *http.Request {
	return c.Context.Request()
}

// ShouldBindWith defined TODO
func (c *Context) ShouldBindWith(v interface{}) error {
	return c.Context.Bind(v)
}

// GetRawData defined TODO
func (c *Context) GetRawData() ([]byte, error) {
	return ioutil.ReadAll(c.Context.Request().Body)
}

// SetCookie defined TODO
func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	if path == "" {
		path = "/"
	}
	c.Context.SetCookie(&http.Cookie{
		Name:     name,
		Value:    url.QueryEscape(value),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
}

// Cookie defined TODO
func (c *Context) Cookie(name string) (string, error) {
	cookie, err := c.Context.Request().Cookie(name)
	if err != nil {
		return "", err
	}
	val, _ := url.QueryUnescape(cookie.Value)
	return val, nil
}

// Success defined TODO
func (c *Context) Success(data interface{}) {
	c.Context.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"data": data,
	})
}

// Next defined TODO
func (c *Context) Next() {
	if c.next != nil {
		c.next()
	}
}

// Abort defined TODO
func (c *Context) Abort() {

}

// Fail defined TODO
func (c *Context) Fail(err error) {
	c.Context.JSON(http.StatusOK, map[string]interface{}{
		"code":   500,
		"detail": err.Error(),
	})
}

// File defined TODO
func (c *Context) File(r io.Reader, filename string, context ...interface{}) {
	c.Context.Response().Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
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
	http.ServeFile(c.Context.Response().Writer, c.Context.Request(), file.Name())
}

// HTML defined TODO
func (c *Context) HTML(r io.Reader, context ...interface{}) {
	c.Context.Response().Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
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
	http.ServeFile(c.Context.Response().Writer, c.Context.Request(), file.Name())
}

// XML defined TODO
func (c *Context) XML(r io.Reader, context ...interface{}) {
	c.Context.Response().Writer.Header().Set("Content-Type", "application/xml; charset=utf-8")
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
	http.ServeFile(c.Context.Response().Writer, c.Context.Request(), file.Name())
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
func (c *Context) ResponseWriter() http.ResponseWriter {
	return c.Context.Response().Writer
}

// Request defined TODO
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.Context.Request().Context().Deadline()
}

// Done defined TODO
func (c *Context) Done() <-chan struct{} {
	return c.Context.Request().Context().Done()
}

// Err defined TODO
func (c *Context) Err() error {
	return c.Context.Request().Context().Err()
}

// Value defined TODO
func (c *Context) Value(key interface{}) interface{} {
	return c.Context.Request().Context().Err()
}

func NewContext(ctx echo.Context, cb ...func()) *Context {
	var cb1 func()
	if len(cb) > 0 {
		cb1 = cb[0]
	}
	return &Context{ctx, cb1}
}
