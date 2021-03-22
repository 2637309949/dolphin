// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/shurcooL/httpfs/vfsutil"
)

// SQL struct
type SQL struct {
}

// Name defined pipe name
func (app *SQL) Name() string {
	return "sql"
}

// Build func
func (app *SQL) Build(dir string, args []string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	tplCache := map[string]bool{}
	countByte, _ := vfsutil.ReadFile(template.Assets, "count.tmpl")
	selectByte, _ := vfsutil.ReadFile(template.Assets, "select.tmpl")
	treeByte, _ := vfsutil.ReadFile(template.Assets, "tree.tmpl")
	for i := range node.Controllers {
		for _, api := range node.Controllers[i].APIS {
			if strings.TrimSpace(api.Table) != "" && api.Func == "page" {
				data := map[string]interface{}{
					"PackageName": node.PackageName,
					"Name":        node.Name,
					"Controller":  node.Controllers[i],
					"Application": node,
					"Api":         api,
					"Viper":       viper.GetViper(),
				}
				cpath := path.Join(dir, viper.GetString("dir.sql"), node.Controllers[i].Name, fmt.Sprintf("%v_%v_%v.tpl", node.Controllers[i].Name, api.Name, "count"))
				if _, ok := tplCache[filepath.Base(cpath)]; !ok {
					tmplCfg := &pipe.TmplCfg{
						Text:     string(countByte),
						FilePath: cpath,
						Data:     data,
						Overlap:  pipe.OverlapSkip,
					}
					tmplCfgs = append(tmplCfgs, tmplCfg)
					tplCache[filepath.Base(cpath)] = true
				}
				spath := path.Join(dir, viper.GetString("dir.sql"), node.Controllers[i].Name, fmt.Sprintf("%v_%v_%v.tpl", node.Controllers[i].Name, api.Name, "select"))
				if _, ok := tplCache[filepath.Base(spath)]; !ok {
					tmplCfg := &pipe.TmplCfg{
						Text:     string(selectByte),
						FilePath: spath,
						Data:     data,
						Overlap:  pipe.OverlapSkip,
					}
					tmplCfgs = append(tmplCfgs, tmplCfg)
					tplCache[filepath.Base(spath)] = true
				}
			} else if strings.TrimSpace(api.Table) != "" && api.Func == "tree" {
				data := map[string]interface{}{
					"PackageName": node.PackageName,
					"Name":        node.Name,
					"Controller":  node.Controllers[i],
					"Application": node,
					"Api":         api,
					"Viper":       viper.GetViper(),
				}
				cpath := path.Join(dir, viper.GetString("dir.sql"), node.Controllers[i].Name, fmt.Sprintf("%v_%v.tpl", node.Controllers[i].Name, api.Name))
				if _, ok := tplCache[filepath.Base(cpath)]; !ok {
					tmplCfg := &pipe.TmplCfg{
						Text:     string(treeByte),
						FilePath: cpath,
						Data:     data,
						Overlap:  pipe.OverlapSkip,
					}
					tmplCfgs = append(tmplCfgs, tmplCfg)
					tplCache[filepath.Base(cpath)] = true
				}
			}
		}

	}
	return tmplCfgs, nil
}
