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
	autoByte, _ := vfsutil.ReadFile(template.Assets, "auto.tmpl")
	return []*pipe.TmplCfg{
		&pipe.TmplCfg{
			Text:     string(autoByte),
			FilePath: path.Join(dir, viper.GetString("dir.app"), "app.auto.go"),
			Data:     data,
			Overlap:  pipe.OverlapWrite,
			GOFmt:    true,
		},
	}, nil
}
