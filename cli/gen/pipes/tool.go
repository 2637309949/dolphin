// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/schema"
	"github.com/2637309949/dolphin/cli/tempalte"
	"github.com/spf13/viper"
)

// Tool struct
type Tool struct {
}

// Build func
func (tool *Tool) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
	}
	return []*gen.TmplCfg{
		&gen.TmplCfg{
			Text:     tempalte.TmplTool,
			FilePath: path.Join(dir, viper.GetString("dir.util"), "tool"),
			Data:     data,
			Overlap:  gen.OverlapSkip,
			Suffix:   ".go",
		},
	}, nil
}
