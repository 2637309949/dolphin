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
	tplCache := map[string]bool{}
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
				cpath := path.Join(dir, viper.GetString("dir.sql"), c.Name, fmt.Sprintf("%v_%v_%v", c.Name, "list", "count"))
				if _, ok := tplCache[cpath]; !ok {
					tmplCfg := &gen.TmplCfg{
						Text:     tempalte.TmplSQLCount,
						FilePath: cpath,
						Data:     data,
						Overlap:  gen.OverlapSkip,
						Suffix:   ".tpl",
					}
					tmplCfgs = append(tmplCfgs, tmplCfg)
					tplCache[cpath] = true
				}
				spath := path.Join(dir, viper.GetString("dir.sql"), c.Name, fmt.Sprintf("%v_%v_%v", c.Name, "list", "select"))
				if _, ok := tplCache[spath]; !ok {
					tmplCfg := &gen.TmplCfg{
						Text:     tempalte.TmplSQLSel,
						FilePath: spath,
						Data:     data,
						Overlap:  gen.OverlapSkip,
						Suffix:   ".tpl",
					}
					tmplCfgs = append(tmplCfgs, tmplCfg)
					tplCache[spath] = true
				}
			}
		}

	}
	return tmplCfgs, nil
}
