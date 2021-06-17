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

// Bean struct
type Bean struct {
}

// Name defined pipe name
func (m *Bean) Name() string {
	return "bean"
}

// Pre defined
func (m *Bean) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (m *Bean) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (m *Bean) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	beanByte, _ := vfsutil.ReadFile(template.Assets, "bean.tmpl")
	for i := range parser.Beans {
		data := map[string]interface{}{
			"PackageName": parser.PackageName,
			"Name":        parser.Name,
			"Controllers": parser.Controllers,
			"Services":    parser.Services,
			"Tables":      parser.Tables,
			"Beans":       parser.Beans,
			"Bean":        parser.Beans[i],
			"Viper":       viper.GetViper(),
			"lt":          ht.HTML("<"),
			"gt":          ht.HTML(">"),
		}
		filename := utils.FileNameTrimSuffix(parser.Beans[i].Path)
		tmplCfg := &pipe.TmplCfg{
			Text:     string(beanByte),
			FilePath: path.Join(dir, viper.GetString("dir.model"), filename+".auto.go"),
			Data:     data,
			Overlap:  pipe.OverlapWrite,
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
