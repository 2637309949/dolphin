package app

import (
	"net/http"

	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"github.com/xormplus/xorm"
	"gopkg.in/guregu/null.v3"
)

// Context defined http handle hook context
type Context struct {
	*gin.Context
	User *model.User
	DB   *xorm.Engine
}

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// RouterGroup defines struct that extend from gin.RouterGroup
type RouterGroup struct {
	*gin.RouterGroup
	Engine *Engine
}

// Success defined success result
func (ctx *Context) Success(data interface{}, status ...int) {
	code := http.StatusInternalServerError
	ctx.JSON(code, model.Response{
		Code: null.IntFrom(200),
		Data: data,
	})
}

// Fail defined failt result
func (ctx *Context) Fail(err error, status ...int) {
	sise, code := http.StatusInternalServerError, http.StatusInternalServerError
	msg := err.Error()
	if cusErr, ok := err.(model.Error); ok {
		code = cusErr.Code
	}
	if len(status) > 0 {
		sise = status[0]
	}
	ctx.JSON(sise, model.Response{
		Code: null.IntFrom(int64(code)),
		Msg:  null.StringFrom(msg),
	})
}

// WithUser defined User
func (ctx *Context) WithUser(user model.User) {
	ctx.User = &user
	return
}

// HandlerFunc convert to gin.HandlerFunc
func (h HandlerFunc) HandlerFunc(e *Engine) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		c := &Context{Context: ctx}
		c.WithUser(model.User{})
		c.DB = e.BusinessDBSet["localhost"]
		h(c)
	})
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return rg.RouterGroup.Handle(httpMethod, relativePath, funk.Map(handlers, func(h HandlerFunc) gin.HandlerFunc {
		return h.HandlerFunc(rg.Engine)
	}).([]gin.HandlerFunc)...)
}
