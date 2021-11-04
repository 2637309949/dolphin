package main

import (
	"github.com/2637309949/dolphin/packages/web"
	"github.com/2637309949/dolphin/packages/web/echo"
)

type Context struct {
	*echo.Context
}

func main() {
	web.Handle("GET", "/000", func(ctx *Context) {
		ctx.Success(map[string]interface{}{
			"su": 123,
		})
	})
	web.Run(":8089")
}
