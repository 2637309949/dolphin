// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Model struct
type Model struct {
}

// Name defined pipe name
func (m *Model) Name() string {
	return "model"
}

// Pre defined
func (m *Model) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (m *Model) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (m *Model) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	modelByte, _ := vfsutil.ReadFile(template.Assets, "model.tmpl")
	for i := range parser.Tables {
		data := map[string]interface{}{
			"PackageName": parser.PackageName,
			"Name":        parser.Name,
			"Table":       parser.Tables[i],
			"Viper":       viper.GetViper(),
		}
		filename := utils.FileNameTrimSuffix(parser.Tables[i].Path)
		tmplCfg := &pipe.TmplCfg{
			Text:     string(modelByte),
			FilePath: path.Join(dir, viper.GetString("dir.model"), filename+".auto.go"),
			Data:     data,
			Overlap:  pipe.OverlapWrite,
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
