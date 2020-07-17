// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/shurcooL/httpfs/vfsutil"
)

// Model struct
type Model struct {
}

// Name defined pipe name
func (m *Model) Name() string {
	return "model"
}

// Build func
func (m *Model) Build(dir string, args []string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	modelByte, _ := vfsutil.ReadFile(template.Assets, "model.tmpl")
	for _, table := range node.Tables {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Table":       table,
		}
		filename := utils.FileNameTrimSuffix(table.Path)
		tmplCfg := &pipe.TmplCfg{
			Text:     string(modelByte),
			FilePath: path.Join(dir, viper.GetString("dir.model"), filename+".auto.go"),
			Data:     data,
			Overlap:  pipe.OverlapWrite,
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
