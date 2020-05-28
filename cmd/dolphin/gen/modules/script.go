// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"fmt"
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/packages/viper"
)

// Script struct
type Script struct {
}

// Name defined pipe name
func (app *Script) Name() string {
	return "script"
}

// Build func
func (app *Script) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	tplCache := map[string]bool{}
	tmplCfgs = append(tmplCfgs, &pipe.TmplCfg{
		Text:     template.TmplAxios,
		FilePath: path.Join(viper.GetString("dir.script"), "axios.js"),
		Overlap:  pipe.OverlapWrite,
	})
	tmplCfgs = append(tmplCfgs, &pipe.TmplCfg{
		Text: template.TmplAPI,
		Data: map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Controllers": node.Controllers,
			"Application": node,
		},
		FilePath: path.Join(viper.GetString("dir.script"), "apis", "index.js"),
		Overlap:  pipe.OverlapWrite,
	})
	for _, c := range node.Controllers {
		for _, api := range c.APIS {
			data := map[string]interface{}{
				"PackageName": node.PackageName,
				"Name":        node.Name,
				"Controller":  c,
				"Application": node,
				"Api":         api,
			}
			cpath := path.Join(dir, viper.GetString("dir.script"), "apis", fmt.Sprintf("%v.js", c.Name))
			if _, ok := tplCache[cpath]; !ok {
				tmplCfg := &pipe.TmplCfg{
					Text:     template.TmplAPIS,
					FilePath: cpath,
					Data:     data,
					Overlap:  pipe.OverlapWrite,
				}
				tmplCfgs = append(tmplCfgs, tmplCfg)
				tplCache[cpath] = true
			}
		}

	}
	return tmplCfgs, nil
}
