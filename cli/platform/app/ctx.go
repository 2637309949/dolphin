package app

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"github.com/2637309949/dolphin/cli/null"
	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"github.com/xormplus/xorm"
)

// Context defined http handle hook context
type Context struct {
	*gin.Context
	AuthInfo
	engine *Engine
	DB     *xorm.Engine
}

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// RouterGroup defines struct that extend from gin.RouterGroup
type RouterGroup struct {
	*gin.RouterGroup
	engine *Engine
}

// Success defined success result
func (ctx *Context) Success(data interface{}, status ...int) {
	sise, code := http.StatusOK, http.StatusOK
	ctx.JSON(sise, model.Response{
		Code: null.IntFrom(int64(code)),
		Data: data,
	})
}

// Fail defined failt result
func (ctx *Context) Fail(err error, status ...int) {
	sise, code := http.StatusInternalServerError, http.StatusInternalServerError
	msg := fmt.Sprintf("%v", errors.WithStack(err))
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

// HandlerFunc convert to gin.HandlerFunc
func (h HandlerFunc) HandlerFunc(e *Engine) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		c := e.pool.Get().(*Context)
		c.Context = ctx
		// from upper middles dataset
		for _, v := range ctx.Keys {
			switch t := v.(type) {
			case *xorm.Engine:
				c.DB = t
			case AuthInfo:
				c.AuthInfo = t
			}
		}
		h(c)
		e.pool.Put(c)
	})
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return rg.RouterGroup.Handle(httpMethod, relativePath, funk.Map(handlers, func(h HandlerFunc) gin.HandlerFunc {
		return h.HandlerFunc(rg.engine)
	}).([]gin.HandlerFunc)...)
}
