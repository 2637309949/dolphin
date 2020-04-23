// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/packages/viper"
)

// Model struct
type Model struct {
}

// Name defined pipe name
func (m *Model) Name() string {
	return "model"
}

// Build func
func (m *Model) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	for _, table := range node.Tables {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Table":       table,
		}
		tmplCfg := &pipe.TmplCfg{
			Text:     template.TmplModel,
			FilePath: path.Join(dir, viper.GetString("dir.model"), table.Name+".auto"),
			Data:     data,
			Overlap:  pipe.OverlapWrite,
			Suffix:   ".go",
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
