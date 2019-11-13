// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cmd

import (
	"os"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/gen/pipes"
	"github.com/2637309949/dolphin/cli/parser"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "build",
	Short: "build a empty project",
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
		g.AddPipe(&pipes.Model{})
		g.AddPipe(&pipes.Bean{})
		g.AddPipe(&pipes.Auto{})
		g.AddPipe(&pipes.Tool{})
		g.AddPipe(&pipes.Doc{})
		if err := g.Build(wd); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
