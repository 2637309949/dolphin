package app

import (
	"path"

	"github.com/2637309949/dolphin/platform/plugin"
	"github.com/2637309949/dolphin/platform/util/file"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewGin defined
func NewGin() *gin.Engine {
	gn := gin.New()
	gn.Use(plugin.CORS())
	gn.Static(viper.GetString("http.static"), path.Join(file.Getwd(), viper.GetString("http.static")))
	gn.Use(plugin.Recovery())
	gn.Use(plugin.Override(gn.HandleContext))
	return gn
}

func init() {
	gin.DefaultWriter = logrus.StandardLogger().Out
	gin.SetMode(viper.GetString("app.mode"))
}
