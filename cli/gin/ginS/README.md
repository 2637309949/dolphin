# Gin Default Server

This is API experiment for Gin.

```go
package main

import (
	"github.com/2637309949/dolphin/cli/gin"
	"github.com/2637309949/dolphin/cli/gin/ginS"
)

func main() {
	ginS.GET("/", func(c *gin.Context) { c.String(200, "Hello World") })
	ginS.Run()
}
```
