package conf

import (
	"fmt"
	"strings"

	"github.com/2637309949/dolphin/cli/packages/logrus"
	"github.com/2637309949/dolphin/cli/packages/viper"
)

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.AddConfigPath("conf")
	viper.SetDefault("app.mode", "release")
	viper.SetDefault("http.port", "8081")
	viper.SetDefault("http.prefix", "/api")
	viper.SetDefault("http.static", "/static")
	viper.SetDefault("grpc.port", "9081")
	viper.SetDefault("oauth.id", "Y76U9344RABF4")
	viper.SetDefault("oauth.secret", "98UYO6FVB865")
	viper.SetDefault("oauth.login", "/static/h5/login.html")
	viper.SetDefault("oauth.auth", "/static/h5/auth.html")
	viper.SetDefault("db.driver", "mysql")
	viper.SetDefault("db.dataSource", "root:111111@/dolphin?charset=utf8&parseTime=True&loc=Local")
	viper.SetDefault("rd.dataSource", ":@127.0.0.1:6379/0")
	viper.SetDefault("dir.sql", "sql")
	viper.SetDefault("dir.sqlmap", "sqlmap")
	viper.SetDefault("dir.log", "log")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logrus.Warn("unable to read config file")
	}
	if strings.TrimSpace(viper.GetString("oauth.server")) == "" {
		viper.SetDefault("oauth.server", fmt.Sprintf("http://127.0.0.1:%v", viper.GetString("http.port")))
	}
	if strings.TrimSpace(viper.GetString("oauth.cli")) == "" {
		viper.SetDefault("oauth.cli", fmt.Sprintf("http://127.0.0.1:%v", viper.GetString("http.port")))
	}
	if err := viper.WriteConfig(); err != nil {
		logrus.Warn("unable to save config file")
	}
}
