package app

import (
	"net/http"

	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"github.com/xormplus/xorm"
	"github.com/2637309949/dolphin/cli/null"
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

// WithUser add user
func (ctx *Context) WithUser(u *model.User) {
	ctx.User = u
}

// Success defined success result
func (ctx *Context) Success(data interface{}, status ...int) {
	code := http.StatusOK
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

// HandlerFunc convert to gin.HandlerFunc
func (h HandlerFunc) HandlerFunc(e *Engine) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		c := &Context{Context: ctx}
		c.Request.ParseForm()
		if accessToken, ok := e.OAuth2.BearerAuth(ctx.Request); ok {
			if token, err := e.OAuth2.Manager.LoadAccessToken(accessToken); err == nil {
				account := model.User{
					ID: null.StringFrom(token.GetUserID()),
				}
				e.PlatformDB.Where("delete_time is null").Get(&account)
				c.WithUser(&account)
				c.DB = e.BusinessDBSet[account.Domain.String]
			}
		}
		h(c)
	})
}

// Handle overwrite RouterGroup.Handle
func (rg *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return rg.RouterGroup.Handle(httpMethod, relativePath, funk.Map(handlers, func(h HandlerFunc) gin.HandlerFunc {
		return h.HandlerFunc(rg.Engine)
	}).([]gin.HandlerFunc)...)
}
