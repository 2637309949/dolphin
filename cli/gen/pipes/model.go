// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/schema"
	"github.com/2637309949/dolphin/cli/tempalte"
)

// Model struct
type Model struct {
}

// Build func
func (m *Model) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	var tmplCfgs []*gen.TmplCfg
	for _, table := range node.Tables {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Table":       table,
		}
		tmplCfg := &gen.TmplCfg{
			Text:     tempalte.TmplModel,
			FilePath: path.Join(dir, "model", table.Name+".auto"),
			Data:     data,
			Overlap:  gen.OverlapWrite,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
