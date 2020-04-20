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

// Auto struct
type Auto struct {
}

// Name defined pipe name
func (auto *Auto) Name() string {
	return "auto"
}

// Build func
func (auto *Auto) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
		"Controllers": node.Controllers,
		"Tables":      node.Tables,
	}
	return []*pipe.TmplCfg{
		&pipe.TmplCfg{
			Text:     template.TmplAuto,
			FilePath: path.Join(dir, viper.GetString("dir.app"), "app"),
			Data:     data,
			Overlap:  pipe.OverlapWrite,
			Suffix:   ".auto.go",
			GOFmt:    true,
		},
	}, nil
}
