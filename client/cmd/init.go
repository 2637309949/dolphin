// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/2637309949/dolphin/packages/go-funk"

	"github.com/2637309949/dolphin/client/gen"
	"github.com/2637309949/dolphin/client/parser"
	"github.com/2637309949/dolphin/packages/cobra"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "build a initialized project",
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		p := parser.New()
		if err := p.Walk(wd); err != nil {
			return err
		}

		args = append(args, "main", "app", "ctr", "srv", "model", "bean", "auto", "tool", "sql", "sqlmap", "oauth", "script", "doc")
		g := gen.New(p.Application)
		funk.ForEach(args, func(name string) {
			g.AddPipe(GetPipeByName(name))
		})
		if err := g.Build(wd); err != nil {
			return err
		}
		return nil
	},
}

var clean = &cobra.Command{
	Use:   "clean",
	Short: "clear all cached files",
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		var files []string
		if err := filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".new") {
				files = append(files, path)
			}
			return nil
		}); err != nil {
			return err
		}
		funk.ForEach(files, func(newFile string) {
			os.Remove(newFile)
		})
		return nil
	},
}

func init() {
	rootCmd.AddCommand(build)
	rootCmd.AddCommand(clean)
}
