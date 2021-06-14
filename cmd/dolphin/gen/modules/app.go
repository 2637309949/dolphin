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

// App struct
type App struct {
}

// Name defined pipe name
func (app *App) Name() string {
	return "app"
}

// Pre defined
func (app *App) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (app *App) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (app *App) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": parser.PackageName,
		"Name":        parser.Name,
		"Application": parser,
		"Viper":       viper.GetViper(),
		"lt":          ht.HTML("<"),
	}

	appByte, _ := vfsutil.ReadFile(template.Assets, "app.tmpl")
	return []*pipe.TmplCfg{
		{
			Text:     string(appByte),
			FilePath: path.Join(dir, viper.GetString("dir.app"), "app.go"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    true,
		},
	}, nil
}
