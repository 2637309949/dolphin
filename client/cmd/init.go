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
	"github.com/2637309949/dolphin/client/gen/pipes"
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
		g := gen.New(p.Application)
		g.AddPipe(&pipes.Main{})
		g.AddPipe(&pipes.App{})
		g.AddPipe(&pipes.Ctr{})
		g.AddPipe(&pipes.Srv{})
		g.AddPipe(&pipes.Model{})
		g.AddPipe(&pipes.Bean{})
		g.AddPipe(&pipes.Auto{})
		g.AddPipe(&pipes.Tool{})
		g.AddPipe(&pipes.SQL{})
		g.AddPipe(&pipes.SQLMap{})
		g.AddPipe(&pipes.OAuth{})
		g.AddPipe(&pipes.Script{})
		g.AddPipe(&pipes.Doc{})
		for _, v := range args {
			tpl := &pipes.SQLTPL{}
			if v == tpl.Name() {
				g.AddPipe(&pipes.SQLTPL{})
			}
		}
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
