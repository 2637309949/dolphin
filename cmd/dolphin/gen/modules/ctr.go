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

// Ctr struct
type Ctr struct {
}

// Name defined pipe name
func (ctr *Ctr) Name() string {
	return "ctr"
}

// Pre defined
func (ctr *Ctr) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (ctr *Ctr) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (ctr *Ctr) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	ctrByte, err := vfsutil.ReadFile(template.Assets, "ctr.tmpl")
	if err != nil {
		return []*pipe.TmplCfg{}, err
	}
	for i := range parser.Controllers {
		data := map[string]interface{}{
			"PackageName": parser.PackageName,
			"Name":        parser.Name,
			"Controllers": parser.Controllers,
			"Controller":  parser.Controllers[i],
			"Services":    parser.Services,
			"Tables":      parser.Tables,
			"Beans":       parser.Beans,
			"Viper":       viper.GetViper(),
			"lt":          ht.HTML("<"),
			"gt":          ht.HTML(">"),
		}
		filename := utils.FileNameTrimSuffix(parser.Controllers[i].Path)
		tmplCfg := &pipe.TmplCfg{
			Text:     string(ctrByte),
			FilePath: path.Join(dir, viper.GetString("dir.app"), filename+".go"),
			Data:     data,
			Overlap:  pipe.OverlapInc,
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
