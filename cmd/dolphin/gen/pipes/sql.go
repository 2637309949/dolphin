// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/packages/viper"
)

// SQL struct
type SQL struct {
}

// Name defined pipe name
func (app *SQL) Name() string {
	return "sql"
}

// Build func
func (app *SQL) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	tplCache := map[string]bool{}
	for _, c := range node.Controllers {
		for _, api := range c.APIS {
			if strings.TrimSpace(api.Table) != "" && api.Func == "page" {
				data := map[string]interface{}{
					"PackageName": node.PackageName,
					"Name":        node.Name,
					"Controller":  c,
					"Application": node,
					"Api":         api,
				}
				cpath := path.Join(dir, viper.GetString("dir.sql"), c.Name, fmt.Sprintf("%v_%v_%v", c.Name, api.Name, "count"))
				if _, ok := tplCache[filepath.Base(cpath)]; !ok {
					tmplCfg := &pipe.TmplCfg{
						Text:     template.TmplSQLCount,
						FilePath: cpath,
						Data:     data,
						Overlap:  pipe.OverlapSkip,
						Suffix:   ".tpl",
					}
					tmplCfgs = append(tmplCfgs, tmplCfg)
					tplCache[filepath.Base(cpath)] = true
				}
				spath := path.Join(dir, viper.GetString("dir.sql"), c.Name, fmt.Sprintf("%v_%v_%v", c.Name, api.Name, "select"))
				if _, ok := tplCache[filepath.Base(spath)]; !ok {
					tmplCfg := &pipe.TmplCfg{
						Text:     template.TmplSQLSel,
						FilePath: spath,
						Data:     data,
						Overlap:  pipe.OverlapSkip,
						Suffix:   ".tpl",
					}
					tmplCfgs = append(tmplCfgs, tmplCfg)
					tplCache[filepath.Base(spath)] = true
				}
			}
		}

	}
	return tmplCfgs, nil
}
