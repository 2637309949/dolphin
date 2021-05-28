// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Tool struct
type Tool struct {
}

// Name defined pipe name
func (tool *Tool) Name() string {
	return "tool"
}

// Pre defined
func (tool *Tool) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (tool *Tool) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (tool *Tool) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": parser.PackageName,
		"Name":        parser.Name,
		"Viper":       viper.GetViper(),
	}
	toolByte, _ := vfsutil.ReadFile(template.Assets, "tool.tmpl")
	return []*pipe.TmplCfg{
		{
			Text:     string(toolByte),
			FilePath: path.Join(dir, viper.GetString("dir.util"), "common.go"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    true,
		},
	}, nil
}
