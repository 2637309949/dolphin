// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

var (
	rootCmd = &cobra.Command{
		Use:   "Dolphin",
		Short: "Dol",
		Long:  `Dolphin, a cli tools for generate golang code`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !terminal.IsTerminal(unix.Stdout) {
				logrus.SetFormatter(&logrus.JSONFormatter{
					TimestampFormat: "2006/01/02 15:04:05",
				})
			} else {
				logrus.SetFormatter(&logrus.TextFormatter{
					FullTimestamp:   true,
					TimestampFormat: "2006/01/02 15:04:05",
				})
			}
			if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
				logrus.SetLevel(logrus.DebugLevel)
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
