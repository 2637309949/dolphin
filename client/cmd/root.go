// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"

	"github.com/2637309949/dolphin/packages/cobra"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

var (
	timeFormat = "2006/01/02 15:04:05"
	rootCmd    = &cobra.Command{
		Use:   "dolphin",
		Short: "dol",
		Long:  `dolphin, a cli tools for generate golang code`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !terminal.IsTerminal(unix.Stdout) {
				logrus.SetFormatter(&logrus.JSONFormatter{
					TimestampFormat: timeFormat,
				})
			} else {
				logrus.SetFormatter(&logrus.TextFormatter{
					FullTimestamp:   true,
					TimestampFormat: timeFormat,
				})
			}
			if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
				logrus.SetLevel(logrus.DebugLevel)
			}
			viper.SetConfigName("app")
			viper.AddConfigPath(".")
			viper.AddConfigPath("conf")
			viper.SetDefault("app.name", "dolphin")
			viper.SetDefault("app.version", "1.0")
			viper.SetDefault("app.mode", "release")
			viper.SetDefault("app.viper", false)
			viper.SetDefault("http.hash", "FF61A573-82FC-478B-9AEF-93D6F506DE9A")
			viper.SetDefault("http.port", "8081")
			viper.SetDefault("http.prefix", "/api")
			viper.SetDefault("http.static", "static")
			viper.SetDefault("http.temp", "temp")
			viper.SetDefault("grpc.port", "9081")
			viper.SetDefault("oauth.id", "Y76U9344RABF4")
			viper.SetDefault("oauth.secret", "98UYO6FVB865")
			viper.SetDefault("oauth.login", "/static/web/login.html")
			viper.SetDefault("oauth.affirm", "/static/web/affirm.html")
			viper.SetDefault("db.driver", "mysql")
			viper.SetDefault("db.dataSource", "root:111111@/dolphin?charset=utf8&parseTime=True&loc=Local")
			viper.SetDefault("rd.dataSource", ":@localhost:6379/0")
			viper.SetDefault("dir.app", "app")
			viper.SetDefault("dir.doc", "doc")
			viper.SetDefault("dir.sql", "sql")
			viper.SetDefault("dir.sqlmap", "sqlmap")
			viper.SetDefault("dir.script", "script")
			viper.SetDefault("dir.log", "log")
			viper.SetDefault("dir.util", "util")
			viper.SetDefault("dir.model", "model")
			viper.SetDefault("dir.srv", "srv")
			viper.SetDefault("swag.license.name", "Apache 2.0")
			viper.SetDefault("swag.license.url", "http://www.apache.org/licenses/LICENSE-2.0.html")
			viper.SetDefault("swag.securitydefinitions.oauth2.accessCode", "OAuth2AccessCode")
			viper.SetDefault("swag.tokenUrl", "/api/sys/cas/token")
			viper.SetDefault("swag.authorizationUrl", "/api/sys/cas/authorize")
			viper.SetDefault("swag.scope.read", "Grants read access")
			viper.SetDefault("swag.scope.write", "Grants write access")
			viper.SetDefault("swag.scope.admin", "Grants read and write access to administrative information")
			viper.AutomaticEnv()
			if err := viper.ReadInConfig(); err != nil {
				logrus.Warn("configuration file not found")
			}
			if strings.TrimSpace(viper.GetString("oauth.server")) == "" {
				viper.SetDefault("oauth.server", fmt.Sprintf("http://localhost:%v", viper.GetString("http.port")))
			}
			if strings.TrimSpace(viper.GetString("oauth.cli")) == "" {
				viper.SetDefault("oauth.cli", fmt.Sprintf("http://localhost:%v", viper.GetString("http.port")))
			}
			if strings.TrimSpace(viper.GetString("host")) == "" {
				viper.SetDefault("host", fmt.Sprintf("localhost:%v", viper.GetString("http.port")))
			}
			if viper.GetBool("app.viper") {
				if err := viper.WriteConfig(); err != nil {
					logrus.Warn("failed to save configuration file")
				}
			}
		},
	}
)

// Execute root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
