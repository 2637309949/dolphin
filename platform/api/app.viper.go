package api

import (
	"fmt"
	"strings"

	"github.com/2637309949/dolphin/platform/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.AddConfigPath("conf")
	viper.SetDefault("app.mode", "release")
	viper.SetDefault("app.name", "dolphin")
	viper.SetDefault("app.version", "1.0")
	viper.SetDefault("app.viper", false)
	viper.SetDefault("app.auth", "token")
	viper.SetDefault("jwt.secret", "6hy789iu87")
	viper.SetDefault("jwt.expire", 604800)
	viper.SetDefault("db.connMaxLifetime", 10)
	viper.SetDefault("db.dataSource", "root:111111@/dolphin?charset=utf8&parseTime=True&loc=Local")
	viper.SetDefault("db.driver", "mysql")
	viper.SetDefault("db.maxIdleConns", 15)
	viper.SetDefault("db.maxOpenConns", 50)
	viper.SetDefault("dir.api", "api")
	viper.SetDefault("dir.doc", "")
	viper.SetDefault("dir.log", "log")
	viper.SetDefault("dir.types", "types")
	viper.SetDefault("dir.rpc", "rpc")
	viper.SetDefault("dir.script", "script")
	viper.SetDefault("dir.sql", "sql")
	viper.SetDefault("dir.svc", "svc")
	viper.SetDefault("dir.sqlmap", "sqlmap")
	viper.SetDefault("dir.srv", "srv")
	viper.SetDefault("dir.util", "util")
	viper.SetDefault("rpc.port", "9081")
	viper.SetDefault("rpc.user_srv", "localhost:9081")
	viper.SetDefault("rpc.domain_srv", "localhost:9081")
	viper.SetDefault("http.hash", "FF61A573-82FC-478B-9AEF-93D6F506DE9A")
	viper.SetDefault("http.port", "8081")
	viper.SetDefault("http.prefix", "/api")
	viper.SetDefault("http.static", "static")
	viper.SetDefault("oauth.affirm", "/static/web/affirm.html")
	viper.SetDefault("oauth.id", types.DefaultClient.Client.String)
	viper.SetDefault("oauth.login", "/static/web/login.html")
	viper.SetDefault("oauth.secret", types.DefaultClient.Secret.String)
	viper.SetDefault("redis.addr", "localhost:6379")
	viper.SetDefault("redis.password", "111111")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("swag.authorizationUrl", "/api/sys/cas/authorize")
	viper.SetDefault("swag.license.name", "Apache 2.0")
	viper.SetDefault("swag.license.url", "http://www.apache.org/licenses/LICENSE-2.0.html")
	viper.SetDefault("swag.scope.admin", "Grants read and write access to administrative information")
	viper.SetDefault("swag.scope.read", "Grants read access")
	viper.SetDefault("swag.scope.write", "Grants write access")
	viper.SetDefault("swag.securitydefinitions.oauth2.accessCode", "OAuth2AccessCode")
	viper.SetDefault("swag.tokenUrl", "/api/sys/cas/token")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.Warn("configuration file not found")
	}
	if strings.TrimSpace(viper.GetString("oauth.server")) == "" {
		viper.SetDefault("oauth.server", fmt.Sprintf("http://localhost:%v", viper.GetString("http.port")))
	}
	if strings.TrimSpace(viper.GetString("oauth.client")) == "" {
		viper.SetDefault("oauth.client", fmt.Sprintf("http://localhost:%v", viper.GetString("http.port")))
	}
	if strings.TrimSpace(viper.GetString("app.host")) == "" {
		viper.SetDefault("app.host", fmt.Sprintf("localhost:%v", viper.GetString("http.port")))
	}
	if viper.GetBool("app.viper") {
		if err := viper.WriteConfig(); err != nil {
			logrus.Warn("failed to save configuration file")
		}
	}
}
