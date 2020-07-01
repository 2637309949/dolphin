// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/2637309949/dolphin/packages/cobra"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

// InitViper defined
func InitViper(cmd *cobra.Command, args []string) {
	utils.SetFormatter(terminal.IsTerminal(unix.Stdout))
	utils.SetLevel(cmd)
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
	viper.SetDefault("dir.rpc", "rpc")
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
	utils.ViperSetDefault("oauth.server", strings.TrimSpace(viper.GetString("oauth.server")), fmt.Sprintf("http://localhost:%v", viper.GetString("http.port")))
	utils.ViperSetDefault("oauth.cli", strings.TrimSpace(viper.GetString("oauth.cli")), viper.GetString("http.port"))
	utils.ViperSetDefault("host", strings.TrimSpace(viper.GetString("host")), fmt.Sprintf("localhost:%v", viper.GetString("http.port")))
	if viper.GetBool("app.viper") {
		if err := viper.WriteConfig(); err != nil {
			logrus.Warn("failed to save configuration file")
		}
	}
}

var (
	rootCmd = &cobra.Command{
		Use:   "dolphin",
		Short: "dol",
		Long:  `dolphin, a cli tools for generate golang code`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			InitViper(cmd, args)
		},
	}
	build = &cobra.Command{
		Use:   "build",
		Short: "Build project from xml",
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			p := parser.New()
			if err := p.Walk(wd); err != nil {
				return err
			}
			pipes := append(args, "main", "app", "ctr", "proto", "srv", "model", "bean", "auto", "tool", "sql", "sqlmap", "oauth", "script", "doc")
			g := gen.New(p.Application)
			g.AddPipe(gen.GetPipesByName(pipes...)...)
			err = g.BuildWithDir(wd)
			return err
		},
	}
	clean = &cobra.Command{
		Use:   "clean",
		Short: "Removing intermediate files",
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			utils.RemoveFileWithSuffix(wd, ".new")
			return nil
		},
	}
	setup = &cobra.Command{
		Use:   "init",
		Short: "Initialize a empty project",
		RunE: func(cmd *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			if files, _ := utils.WalkFileInDirWithSuffix(wd, ".xml"); len(files) == 0 {
				p := parser.NewTpl(path.Base(wd), path.Base(wd))
				if err := p.Walk(wd); err != nil {
					return err
				}
				g := gen.New(p.Application)
				g.AddPipe(gen.GetPipesByName("boilerplate")...)
				return g.BuildWithDir(wd)
			}
			logrus.Warn("It is not allowed to initialize a non-empty project")
			return nil
		},
	}
)

func main() {
	rootCmd.AddCommand(setup)
	rootCmd.AddCommand(build)
	rootCmd.AddCommand(clean)
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
