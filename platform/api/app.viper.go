package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/json-iterator/go/extra"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func InitViper() {
	extra.RegisterFuzzyDecoders()
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
	viper.SetDefault("dir.log", "")
	viper.SetDefault("dir.types", "types")
	viper.SetDefault("dir.script", "script")
	viper.SetDefault("dir.sql", "sql")
	viper.SetDefault("dir.svc", "svc")
	viper.SetDefault("dir.sqlmap", "sqlmap")
	viper.SetDefault("dir.srv", "srv")
	viper.SetDefault("dir.util", "util")
	viper.SetDefault("http.hash", "FF61A573-82FC-478B-9AEF-93D6F506DE9A")
	viper.SetDefault("http.port", ":8081")
	viper.SetDefault("http.prefix", "/api")
	viper.SetDefault("http.static", "static")
	viper.SetDefault("oauth.id", types.DefaultClient.Client.String)
	viper.SetDefault("oauth.affirm", "/static/web/affirm.html")
	viper.SetDefault("oauth.qrconnect", "/static/web/qrconnect.html")
	viper.SetDefault("oauth.login", "/static/web/login.html")
	viper.SetDefault("oauth.secret", types.DefaultClient.Secret.String)
	viper.SetDefault("redis.addr", []string{})
	viper.SetDefault("redis.username", "")
	viper.SetDefault("redis.password", "111111")
	viper.SetDefault("redis.db", 0)
	viper.SetDefault("redis.network", "")
	viper.SetDefault("redis.mode", "stub")
	viper.SetDefault("redis.max_redirects", 0)
	viper.SetDefault("redis.read_only", true)
	viper.SetDefault("redis.max_retries", 0)
	viper.SetDefault("redis.dial_timeout", 10)
	viper.SetDefault("redis.read_timeout", 10)
	viper.SetDefault("redis.write_timeout", 10)
	viper.SetDefault("redis.pool_size", 15)
	viper.SetDefault("redis.min_idle_conns", 15)
	viper.SetDefault("redis.idle_conns", 10)
	viper.SetDefault("consul.name", "")
	viper.SetDefault("consul.endpoint", "")
	viper.SetDefault("consul.type", "yaml")
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
		logrus.Warn(context.TODO(), "configuration file not found")
	}
	if strings.TrimSpace(viper.GetString("oauth.server")) == "" {
		viper.SetDefault("oauth.server", fmt.Sprintf("http://localhost%v", viper.GetString("http.port")))
	}
	if strings.TrimSpace(viper.GetString("oauth.client")) == "" {
		viper.SetDefault("oauth.client", fmt.Sprintf("http://localhost%v", viper.GetString("http.port")))
	}
	if strings.TrimSpace(viper.GetString("app.host")) == "" {
		viper.SetDefault("app.host", fmt.Sprintf("localhost%v", viper.GetString("http.port")))
	}

	if strings.TrimSpace(viper.GetString("consul.name")) == "" {
		viper.SetDefault("consul.name", viper.GetString("app.name"))
	}

	if viper.GetString("consul.endpoint") != "" && viper.GetString("consul.name") != "" {
		viper.AddRemoteProvider("consul", viper.GetString("consul.endpoint"), viper.GetString("consul.name"))
		viper.SetConfigType(viper.GetString("consul.type"))
		if err := viper.ReadRemoteConfig(); err != nil {
			logrus.Warnf(context.TODO(), "failed to read remote configuration file [%v]", err)
		}
	}

	if viper.GetBool("app.viper") {
		if err := viper.WriteConfig(); err != nil {
			logrus.Warnf(context.TODO(), "failed to save configuration file [%v]", err)
		}
	}
}
