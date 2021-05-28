// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"fmt"
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Script struct
type Script struct {
}

// Name defined pipe name
func (app *Script) Name() string {
	return "script"
}

// Pre defined
func (app *Script) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (app *Script) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (app *Script) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	tplCache := map[string]bool{}
	axiosByte, _ := vfsutil.ReadFile(template.Assets, "request.tmpl")
	tmplCfgs = append(tmplCfgs, &pipe.TmplCfg{
		Text:     string(axiosByte),
		FilePath: path.Join(viper.GetString("dir.script"), "request.js"),
		Overlap:  pipe.OverlapWrite,
	})

	apiByte, _ := vfsutil.ReadFile(template.Assets, "api.tmpl")
	tmplCfgs = append(tmplCfgs, &pipe.TmplCfg{
		Text: string(apiByte),
		Data: map[string]interface{}{
			"PackageName": parser.PackageName,
			"Name":        parser.Name,
			"Controllers": parser.Controllers,
			"Application": parser,
			"Viper":       viper.GetViper(),
		},
		FilePath: path.Join(viper.GetString("dir.script"), "apis", "index.js"),
		Overlap:  pipe.OverlapWrite,
	})

	apisByte, _ := vfsutil.ReadFile(template.Assets, "apis.tmpl")
	for i := range parser.Controllers {
		filename := utils.FileNameTrimSuffix(parser.Controllers[i].Path)
		for _, api := range parser.Controllers[i].APIS {
			data := map[string]interface{}{
				"PackageName": parser.PackageName,
				"Name":        parser.Name,
				"Controller":  parser.Controllers[i],
				"Application": parser,
				"Api":         api,
				"Viper":       viper.GetViper(),
			}
			cpath := path.Join(dir, viper.GetString("dir.script"), "apis", fmt.Sprintf("%v.js", filename))
			if _, ok := tplCache[cpath]; !ok {
				tmplCfg := &pipe.TmplCfg{
					Text:     string(apisByte),
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
