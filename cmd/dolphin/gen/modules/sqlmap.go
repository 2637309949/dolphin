// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	ht "html/template"
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// SQLMap struct
type SQLMap struct {
}

// Name defined pipe name
func (app *SQLMap) Name() string {
	return "sqlmap"
}

// Pre defined
func (app *SQLMap) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (app *SQLMap) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (app *SQLMap) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	sqlmapByte, _ := vfsutil.ReadFile(template.Assets, "sqlmap.tmpl")
	for i := range parser.Tables {
		data := map[string]interface{}{
			"PackageName": parser.PackageName,
			"Name":        parser.Name,
			"Controllers": parser.Controllers,
			"Services":    parser.Services,
			"Tables":      parser.Tables,
			"Table":       parser.Tables[i],
			"Beans":       parser.Beans,
			"Viper":       viper.GetViper(),
			"lt":          ht.HTML("<"),
			"gt":          ht.HTML(">"),
		}
		filename := utils.FileNameTrimSuffix(parser.Tables[i].Path)
		tmplCfg := &pipe.TmplCfg{
			Text:     string(sqlmapByte),
			FilePath: path.Join(dir, viper.GetString("dir.sql"), viper.GetString("dir.sqlmap"), filename+".xml"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
