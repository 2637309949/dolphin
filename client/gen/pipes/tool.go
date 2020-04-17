// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/client/gen/pipe"
	"github.com/2637309949/dolphin/client/gen/template"
	"github.com/2637309949/dolphin/client/schema"
	"github.com/2637309949/dolphin/packages/viper"
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
	return []*pipe.TmplCfg{
		&pipe.TmplCfg{
			Text:     template.TmplTool,
			FilePath: path.Join(dir, viper.GetString("dir.util"), "tool"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			Suffix:   ".go",
		},
	}, nil
}
