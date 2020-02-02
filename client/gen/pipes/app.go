// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/packages/viper"

	"github.com/2637309949/dolphin/client/gen"
	"github.com/2637309949/dolphin/client/gen/tempalte"
	"github.com/2637309949/dolphin/client/schema"
)

// App struct
type App struct {
}

// Build func
func (app *App) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
		"Application": node,
	}
	return []*gen.TmplCfg{
		&gen.TmplCfg{
			Text:     tempalte.TmplGin,
			FilePath: path.Join(dir, viper.GetString("dir.app"), "app"),
			Data:     data,
			Overlap:  gen.OverlapSkip,
			Suffix:   ".go",
		},
	}, nil
}
