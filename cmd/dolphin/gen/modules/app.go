// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/shurcooL/httpfs/vfsutil"
)

// App struct
type App struct {
}

// Name defined pipe name
func (app *App) Name() string {
	return "app"
}

// Build func
func (app *App) Build(dir string, args []string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
		"Application": node,
		"Viper":       viper.GetViper(),
	}

	appByte, _ := vfsutil.ReadFile(template.Assets, "app.tmpl")
	return []*pipe.TmplCfg{
		&pipe.TmplCfg{
			Text:     string(appByte),
			FilePath: path.Join(dir, viper.GetString("dir.app"), "app.go"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    true,
		},
	}, nil
}
