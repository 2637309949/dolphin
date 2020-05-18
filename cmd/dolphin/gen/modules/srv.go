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
)

// Srv struct
type Srv struct {
}

// Name defined pipe name
func (app *Srv) Name() string {
	return "srv"
}

// Build func
func (app *Srv) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	for _, c := range node.Controllers {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Controller":  c,
		}
		tmplCfg := &pipe.TmplCfg{
			Text:     template.TmplSrv,
			FilePath: path.Join(dir, viper.GetString("dir.srv"), c.Name),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			Suffix:   ".go",
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
