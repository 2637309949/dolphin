// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/template/dist"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/gin-gonic/gin"
	"github.com/shurcooL/vfsgen"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	// "github.com/go-sql-driver/mysql" init
	_ "github.com/go-sql-driver/mysql"
	// "github.com/lib/pq" init
	_ "github.com/lib/pq"
	// "github.com/mattn/go-sqlite3" init
	_ "github.com/mattn/go-sqlite3"

	"golang.org/x/sys/unix"
	"golang.org/x/term"
)

// AssetsFileSystem defined TODO
type AssetsFileSystem struct {
	RelativePath string
	http.FileSystem
}

// Open defined TODO
func (assert *AssetsFileSystem) Open(name string) (http.File, error) {
	if strings.Contains(name, "swagger.yaml") {
		return os.Open(path.Join(viper.GetString("dir.doc"), "swagger.yaml"))
	}
	return assert.FileSystem.Open(path.Join(assert.RelativePath, name))
}

// Open calls the OS default program for uri
func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}

// InitViper defined TODO
func InitViper(cmd *cobra.Command, args []string) {
	utils.SetFormatter(term.IsTerminal(unix.Stdout))
	utils.SetLevel(cmd)
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.AddConfigPath("conf")
	viper.SetDefault("app.name", "dolphin")
	viper.SetDefault("app.version", "1.1.2")
	viper.SetDefault("app.mode", "release")
	viper.SetDefault("app.auth", "token")
	viper.SetDefault("app.viper", false)
	viper.SetDefault("http.hash", "FF61A573-82FC-478B-9AEF-93D6F506DE9A")
	viper.SetDefault("http.port", "8081")
	viper.SetDefault("http.prefix", "/api")
	viper.SetDefault("http.static", "static")
	viper.SetDefault("http.temp", "temp")
	viper.SetDefault("rpc.port", "9081")
	viper.SetDefault("oauth.id", "Y76U9344RABF4")
	viper.SetDefault("oauth.secret", "8UYO6FVB8UYO6FVB")
	viper.SetDefault("oauth.login", "/static/web/login.html")
	viper.SetDefault("oauth.affirm", "/static/web/affirm.html")
	viper.SetDefault("oauth.qrconnect", "/static/web/qrconnect.html")
	viper.SetDefault("db.driver", "mysql")
	viper.SetDefault("db.dataSource", "root:111111@/dolphin?charset=utf8&parseTime=True&loc=Local")
	viper.SetDefault("rd.dataSource", ":@localhost:6379/0")
	viper.SetDefault("dir.api", "api")
	viper.SetDefault("dir.doc", "")
	viper.SetDefault("dir.sql", "sql")
	viper.SetDefault("dir.sqlmap", "sqlmap")
	viper.SetDefault("dir.script", "script")
	viper.SetDefault("dir.log", "log")
	viper.SetDefault("dir.util", "util")
	viper.SetDefault("dir.types", "types")
	viper.SetDefault("dir.srv", "srv")
	viper.SetDefault("dir.svc", "svc")
	viper.SetDefault("dir.proto", "proto")
	viper.SetDefault("dir.rpc", "rpc")
	viper.SetDefault("dir.xml", "xml")
	viper.SetDefault("dir.k8s", "k8s")
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
	utils.ViperSetDefault("oauth.client", strings.TrimSpace(viper.GetString("oauth.client")), viper.GetString("http.port"))
	utils.ViperSetDefault("app.host", strings.TrimSpace(viper.GetString("app.host")), fmt.Sprintf("localhost:%v", viper.GetString("http.port")))
	if viper.GetBool("app.viper") {
		if err := viper.WriteConfig(); err != nil {
			logrus.Warn("failed to save configuration file")
		}
	}
}

var (
	commands = map[string]string{
		"windows": "cmd /c start",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
	rootCmd = &cobra.Command{
		Use:  "dolphin",
		Long: `Code generation tool for golang`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmd.Use == "assert" || cmd.Use == "new" || cmd.Use == "clean" {
				return
			}
			InitViper(cmd, args)
		},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	build = &cobra.Command{
		Use:   "build",
		Short: "Build from the configuration file",
		RunE: func(_ *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			p := parser.New()
			if err := p.Walk(wd); err != nil {
				return err
			}
			if len(args) == 0 {
				args = append(args, "dolphin")
			}
			g := New(p)
			g.AddPipe(GetPipesByName(args...)...)
			if err = g.BuildDir(wd, args); err != nil {
				return err
			}
			return err
		},
	}
	clean = &cobra.Command{
		Use:   "clean",
		Short: "Remove temp file, such as *.go.new",
		RunE: func(*cobra.Command, []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			utils.RemoveFileWithSuffix(wd, ".new")
			return nil
		},
	}
	more = &cobra.Command{
		Use:   "more",
		Short: "Add controller and table",
		RunE: func(_ *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			p := parser.NewTpl(path.Base(wd), path.Base(wd))
			g := New(p)
			g.AddPipe(GetPipesByName("more", "dolphin")...)
			return g.BuildDir(wd, args)
		},
	}
	reverse = &cobra.Command{
		Use:   "reverse",
		Short: "Inversion of the data model",
		RunE: func(_ *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			p := parser.NewTpl(path.Base(wd), path.Base(wd))
			g := New(p)
			g.AddPipe(GetPipesByName("reverse", "dolphin")...)
			return g.BuildDir(wd, args)
		},
	}
	new = &cobra.Command{
		Use:   "new",
		Short: "New a empty project",
		RunE: func(_ *cobra.Command, args []string) error {
			wd, err := os.Getwd()
			if err != nil {
				return err
			}
			wd = path.Join(wd, args[0])
			p := parser.NewTpl(path.Base(wd), path.Base(wd))
			if err := p.Walk(wd); err != nil {
				return err
			}
			g := New(p)
			g.AddPipe(GetPipesByName("boilerplate", "dolphin")...)
			err = g.BuildDir(wd, args)
			if err != nil {
				return err
			}
			logrus.Infof("new project success, cd to %v dir and run `dolphin build`", args[0])
			return nil
		},
	}
	serve = &cobra.Command{
		Use:   "serve",
		Short: "Serve api document",
		RunE: func(_ *cobra.Command, _ []string) error {
			gin.SetMode("release")
			router := gin.New()
			router.StaticFS("/", &AssetsFileSystem{RelativePath: "/swagger", FileSystem: dist.Assets})
			Open("http://localhost:7799")
			router.Run(":7799")
			return nil
		},
	}
	assert = &cobra.Command{
		Use:   "assert",
		Short: "Generate Go code that statically implements input filesystem",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := os.MkdirAll("../dist", os.ModePerm); err != nil {
				if err != nil {
					return err
				}
			}
			err := vfsgen.Generate(http.Dir("./"), vfsgen.Options{
				Filename:     "../dist/assets.go",
				PackageName:  "dist",
				BuildTags:    "!dev",
				VariableName: "Assets",
			})
			if err != nil {
				return err
			}
			return nil
		},
	}
)

func main() {
	rootCmd.AddCommand(new)
	rootCmd.AddCommand(reverse)
	rootCmd.AddCommand(build)
	rootCmd.AddCommand(more)
	rootCmd.AddCommand(clean)
	rootCmd.AddCommand(serve)
	rootCmd.AddCommand(assert)
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
