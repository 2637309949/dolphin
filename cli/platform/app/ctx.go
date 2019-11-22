package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"

	"github.com/2637309949/dolphin/cli/null"
	"github.com/2637309949/dolphin/cli/oauth2"
	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"github.com/xormplus/xorm"
)

// AuthType authorization model
type AuthType string

// define authorization model
const (
	OAuth2 AuthType = "oauth2"
)

// Context defined http handle hook context
type Context struct {
	*gin.Context
	DB    *xorm.Engine
	Token oauth2.TokenInfo
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
		c := &Context{Context: ctx}
		c.Request.ParseForm()
		if accessToken, ok := e.OAuth2.BearerAuth(ctx.Request); ok {
			token, err := e.OAuth2.Manager.LoadAccessToken(accessToken)
			if err != nil {
				logrus.WithError(err).Warning("load accessToken failed.")
			} else {
				c.Token = token
				if domain := strings.TrimSpace(c.Token.GetDomain()); domain != "" {
					c.DB = e.BusinessDBSet[domain]
				}
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
