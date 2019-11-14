// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"fmt"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/schema"
	"github.com/2637309949/dolphin/cli/tempalte"
	"github.com/spf13/viper"
)

// SQL struct
type SQL struct {
}

// Build func
func (app *SQL) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	var tmplCfgs []*gen.TmplCfg
	for _, c := range node.Controllers {
		for _, api := range c.APIS {
			if strings.TrimSpace(api.Table) != "" {
				data := map[string]interface{}{
					"PackageName": node.PackageName,
					"Name":        node.Name,
					"Controller":  c,
					"Application": node,
					"Api":         api,
				}
				tmplCfg := &gen.TmplCfg{
					Text:     tempalte.TmplSQLCount,
					FilePath: path.Join(dir, viper.GetString("sqlTemplate"), c.Name, fmt.Sprintf("%v_%v_%v", api.Table, "page", "count")),
					Data:     data,
					Overlap:  gen.OverlapSkip,
					Suffix:   ".tpl",
				}
				tmplCfgs = append(tmplCfgs, tmplCfg)
				tmplCfg = &gen.TmplCfg{
					Text:     tempalte.TmplSQLSel,
					FilePath: path.Join(dir, viper.GetString("sqlTemplate"), c.Name, fmt.Sprintf("%v_%v_%v", api.Table, "page", "select")),
					Data:     data,
					Overlap:  gen.OverlapSkip,
					Suffix:   ".tpl",
				}
				tmplCfgs = append(tmplCfgs, tmplCfg)
			}
		}

	}
	return tmplCfgs, nil
}
