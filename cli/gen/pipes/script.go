// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"fmt"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/gen/tempalte"
	"github.com/2637309949/dolphin/cli/packages/viper"
	"github.com/2637309949/dolphin/cli/schema"
)

// Script struct
type Script struct {
}

// Build func
func (app *Script) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	var tmplCfgs []*gen.TmplCfg
	tplCache := map[string]bool{}
	tmplCfgs = append(tmplCfgs, &gen.TmplCfg{
		Text:     tempalte.TmplAxios,
		FilePath: path.Join(viper.GetString("dir.script"), "axios"),
		Overlap:  gen.OverlapWrite,
		Suffix:   ".js",
	})
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
				cpath := path.Join(dir, viper.GetString("dir.script"), "api", fmt.Sprintf("%v", c.Name))
				if _, ok := tplCache[cpath]; !ok {
					tmplCfg := &gen.TmplCfg{
						Text:     tempalte.TmplAPI,
						FilePath: cpath,
						Data:     data,
						Overlap:  gen.OverlapWrite,
						Suffix:   ".js",
					}
					tmplCfgs = append(tmplCfgs, tmplCfg)
					tplCache[cpath] = true
				}
			}
		}

	}
	return tmplCfgs, nil
}
