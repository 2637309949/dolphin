package main

import (
	"github.com/2637309949/dolphin/packages/web"
	"github.com/2637309949/dolphin/packages/web/core"
	"github.com/gin-gonic/gin"
)

type Context struct {
	core.Context
}

type Context2 struct {
	core.Context
}

func main() {
	web.Handle("GET", "/000", func(ctx Context2) {
		ctx.Success(map[string]interface{}{
			"su": 123,
		})
	})
	web.Handle("GET", "/222", func(ctx *Context2) {
		ctx.Success(map[string]interface{}{
			"su": 123,
		})
	})
	web.Handle("GET", "/111", func(ctx *gin.Context) {
		ctx.JSON(200, 200)
	})
	web.Handle("GET", "/333", func(ctx core.Context) {
		ctx.Success(map[string]interface{}{
			"su": 123,
		})
	})
	web.Run(":8089")
}
