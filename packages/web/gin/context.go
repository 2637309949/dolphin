package gin

import (
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"reflect"

	"github.com/2637309949/dolphin/packages/web/core"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Context struct {
	engine  *gin.Engine
	context *gin.Context
}

// Set defined TODO
func (c *Context) Set(k string, v interface{}) {
	c.context.Set(k, v)
}

// Get defined TODO
func (c *Context) Get(k string) (interface{}, bool) {
	return c.context.Get(k)
}

// Param defined TODO
func (c *Context) Param(k string) string {
	return c.context.Param(k)
}

// Query defined TODO
func (c *Context) Query(k string) string {
	return c.context.Query(k)
}

// MultipartForm defined TODO
func (c *Context) MultipartForm() (*multipart.Form, error) {
	return c.context.MultipartForm()
}

// Header defined TODO
func (c *Context) Header(key, value string) {
	c.context.Header(key, value)
}

// GetHeader defined TODO
func (c *Context) GetHeader(key string) string {
	return c.context.GetHeader(key)
}

// ShouldBindWith defined TODO
func (c *Context) ShouldBindWith(v interface{}) error {
	if UriCheck(c.context.Params) {
		if err := c.context.ShouldBindUri(v); err != nil {
			return err
		}
	}
	if QueryCheck(c.context.Request) {
		if err := c.context.ShouldBindQuery(v); err != nil {
			return err
		}
	}
	if JsonCheck(c.context.Request) {
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
		body, err = ioutil.ReadAll(c.context.Request.Body)
		if err != nil {
			return err
		}
		c.Set(gin.BodyBytesKey, body)
	}
	return binding.JSON.BindBody(body, ptr)
}

// GetRawData defined TODO
func (c *Context) GetRawData() ([]byte, error) {
	return c.context.GetRawData()
}

// SetCookie defined TODO
func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	c.context.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

// Cookie defined TODO
func (c *Context) Cookie(name string) (string, error) {
	return c.context.Cookie(name)
}

// Next defined TODO
func (c *Context) Next() {
	c.context.Next()
}

// Success defined TODO
func (c *Context) Success(data interface{}) {
	c.context.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"data": data,
	})
}

// Fail defined TODO
func (c *Context) Fail(err error) {
	c.context.JSON(http.StatusOK, map[string]interface{}{
		"code":   500,
		"detail": err.Error(),
	})
}

// HTML defined TODO
func (c *Context) HTML(io.Reader, ...interface{}) {
}

// XML defined TODO
func (c *Context) XML(io.Reader, ...interface{}) {
}

// File defined TODO
func (c *Context) File(io.Reader, string, ...interface{}) {
}

// Handler defined TODO
func (c *Context) Handler() http.Handler {
	return c.engine
}

// Handler defined TODO
func (c *Context) Handle(httpMethod string, relativePath string, handlers ...core.HandlerFunc) {
	hfc := []gin.HandlerFunc{}
	for i := range handlers {
		fv := reflect.ValueOf(handlers[i])
		ft := fv.Type()
		if ft.NumIn() != 1 {
			panic("invalid handler: wrong numIn")
		}
		in := reflect.New(ft.In(0))
		_, ok := in.Interface().(core.Context)
		if ft.In(0).String() != "core.Context" && !ok {
			panic("invalid handler context: should extend core.Context")
		}
		hfc = append(hfc, func(v reflect.Value) func(ct *gin.Context) {
			funcType := v.Type()
			return func(ctx *gin.Context) {
				switch true {
				case ft.In(0).String() == "core.Context":
					v.Call([]reflect.Value{reflect.ValueOf(&Context{context: ctx})})
				default:
					in := reflect.New(funcType.In(0)).Elem()
					ct := in.FieldByName("Context")
					if ct.CanSet() {
						ct.Set(reflect.ValueOf(&Context{context: ctx}))
						v.Call([]reflect.Value{in})
						return
					}
				}
			}
		}(fv))
	}
	c.engine.Handle(httpMethod, relativePath, hfc...)
}

func NewContext() *Context {
	return &Context{
		engine: gin.New(),
	}
}

func init() {
	core.SetEngine(NewContext())
}
