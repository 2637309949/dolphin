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
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Auto struct
type Auto struct {
}

// Name defined pipe name
func (auto *Auto) Name() string {
	return "auto"
}

// Pre defined
func (auto *Auto) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (auto *Auto) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (auto *Auto) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": parser.PackageName,
		"Name":        parser.Name,
		"Controllers": parser.Controllers,
		"Services":    parser.Services,
		"Tables":      parser.Tables,
		"Beans":       parser.Beans,
		"Viper":       viper.GetViper(),
		"lt":          ht.HTML("<"),
		"gt":          ht.HTML(">"),
	}
	autoByte, _ := vfsutil.ReadFile(template.Assets, "auto.tmpl")
	return []*pipe.TmplCfg{
		{
			Text:     string(autoByte),
			FilePath: path.Join(dir, viper.GetString("dir.app"), "app.auto.go"),
			Data:     data,
			Overlap:  pipe.OverlapWrite,
			GOFmt:    true,
		},
	}, nil
}
