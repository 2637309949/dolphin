// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"fmt"
	"html/template"
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/template/dist"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Script struct
type Script struct {
}

// Name defined parser name
func (app *Script) Name() string {
	return "script"
}

// Pre defined
func (app *Script) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (app *Script) After(*parser.AppParser, []*parser.TmplCfg) error {
	return nil
}

// Build func
func (app *Script) Build(dir string, args []string, appParser *parser.AppParser) ([]*parser.TmplCfg, error) {
	var tmplCfgs []*parser.TmplCfg
	tplCache := map[string]bool{}
	axiosByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "request.tmpl")).([]byte)
	tmplCfgs = append(tmplCfgs, &parser.TmplCfg{
		Text:     string(axiosByte),
		FilePath: path.Join(viper.GetString("dir.script"), "request.js"),
		Overlap:  parser.OverlapWrite,
	})
	apiByte, err := vfsutil.ReadFile(dist.Assets, "api.tmpl")
	if err != nil {
		return []*parser.TmplCfg{}, err
	}
	tmplCfgs = append(tmplCfgs, &parser.TmplCfg{
		Text: string(apiByte),
		Data: map[string]interface{}{
			"PackageName": appParser.PackageName,
			"Name":        appParser.Name,
			"Controllers": appParser.Controllers,
			"Services":    appParser.Services,
			"Tables":      appParser.Tables,
			"Beans":       appParser.Beans,
			"Viper":       viper.GetViper(),
			"lt":          template.HTML("<"),
			"gt":          template.HTML(">"),
		},
		FilePath: path.Join(viper.GetString("dir.script"), "apis", "index.js"),
		Overlap:  parser.OverlapWrite,
	})

	apisByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "apis.tmpl")).([]byte)
	for i := range appParser.Controllers {
		filename := utils.FileNameTrimSuffix(appParser.Controllers[i].Path)
		for _, api := range appParser.Controllers[i].APIS {
			data := map[string]interface{}{
				"PackageName": appParser.PackageName,
				"Name":        appParser.Name,
				"Controller":  appParser.Controllers[i],
				"Application": appParser,
				"Api":         api,
				"Viper":       viper.GetViper(),
			}
			cpath := path.Join(dir, viper.GetString("dir.script"), "apis", fmt.Sprintf("%v.js", filename))
			if _, ok := tplCache[cpath]; !ok {
				tmplCfg := &parser.TmplCfg{
					Text:     string(apisByte),
					FilePath: cpath,
					Data:     data,
					Overlap:  parser.OverlapWrite,
				}
				tmplCfgs = append(tmplCfgs, tmplCfg)
				tplCache[cpath] = true
			}
		}

	}
	return tmplCfgs, nil
}
