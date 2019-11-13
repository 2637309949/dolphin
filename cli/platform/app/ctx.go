package app

import (
	"net/http"

	"gopkg.in/guregu/null.v3"

	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/gin-gonic/gin"
	"github.com/xormplus/xorm"
)

// Context defined http handle hook context
type Context struct {
	*gin.Context
	DB   *xorm.Engine
	User *model.User
}

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// RouterGroup defines struct that extend from gin.RouterGroup
type RouterGroup struct {
	*gin.RouterGroup
	Engine *Engine
}

// Success defined success result
func (ctx *Context) Success(data interface{}) {
	ctx.JSON(http.StatusOK, model.Response{
		Code: null.IntFrom(200),
		Data: data,
	})
}

// Fail defined failt result
func (ctx *Context) Fail(err error) {
	code := 500
	msg := err.Error()
	if cusErr, ok := err.(model.Error); ok {
		code = cusErr.Code
	}
	ctx.JSON(http.StatusInternalServerError, model.Response{
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
		// with user
		c.WithUser(model.User{})
		// find db by PlatformDB
		c.DB = e.PlatformDB
		h(c)
	})
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	var newHandlers []gin.HandlerFunc
	for _, h := range handlers {
		newHandlers = append(newHandlers, h.HandlerFunc(rg.Engine))
	}
	return rg.RouterGroup.Handle(httpMethod, relativePath, newHandlers...)
}
