// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"strings"

	"github.com/2637309949/dolphin/cli/packages/cobra"
	"github.com/2637309949/dolphin/cli/packages/logrus"
	"github.com/2637309949/dolphin/cli/packages/viper"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

var (
	timeFormat = "2006/01/02 15:04:05"
	rootCmd    = &cobra.Command{
		Use:   "Dolphin",
		Short: "Dol",
		Long:  `Dolphin, a cli tools for generate golang code`,
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
			viper.SetDefault("http.port", "8081")
			viper.SetDefault("http.prefix", "/api")
			viper.SetDefault("http.static", "/static")
			viper.SetDefault("dir.sql", "sql")
			viper.SetDefault("dir.sqlmap", "sqlmap")
			viper.SetDefault("dir.app", "app")
			viper.SetDefault("dir.doc", "doc")
			viper.SetDefault("dir.log", "log")
			viper.SetDefault("dir.util", "util")
			viper.SetDefault("dir.model", "model")
			viper.SetDefault("dir.srv", "srv")
			viper.SetDefault("cli.plugins", "")
			viper.SetDefault("license.name", "Apache 2.0")
			viper.SetDefault("license.url", "http://www.apache.org/licenses/LICENSE-2.0.html")
			viper.AutomaticEnv()
			if err := viper.ReadInConfig(); err != nil {
				logrus.Warn("unable to read config file")
			}
			if err := viper.WriteConfig(); err != nil {
				logrus.Warn("unable to save config file")
			}
			if strings.TrimSpace(viper.GetString("host")) == "" {
				viper.SetDefault("host", fmt.Sprintf("127.0.0.1:%v", viper.GetString("http.port")))
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
