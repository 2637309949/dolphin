// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/shurcooL/httpfs/vfsutil"
)

// Tool struct
type Tool struct {
}

// Name defined pipe name
func (tool *Tool) Name() string {
	return "tool"
}

// Build func
func (tool *Tool) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
	}
	toolByte, _ := vfsutil.ReadFile(template.Assets, "tool.tmpl")
	return []*pipe.TmplCfg{
		&pipe.TmplCfg{
			Text:     string(toolByte),
			FilePath: path.Join(dir, viper.GetString("dir.util"), "tool.go"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    true,
		},
	}, nil
}
