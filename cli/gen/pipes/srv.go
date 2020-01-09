// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/gen/tempalte"
	"github.com/2637309949/dolphin/cli/packages/viper"
	"github.com/2637309949/dolphin/cli/schema"
)

// Srv struct
type Srv struct {
}

// Build func
func (app *Srv) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	var tmplCfgs []*gen.TmplCfg
	for _, c := range node.Controllers {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Controller":  c,
		}
		tmplCfg := &gen.TmplCfg{
			Text:     tempalte.TmplSrv,
			FilePath: path.Join(dir, viper.GetString("dir.srv"), c.Name),
			Data:     data,
			Overlap:  gen.OverlapSkip,
			Suffix:   ".go",
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
