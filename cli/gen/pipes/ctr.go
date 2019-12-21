// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/schema"
	"github.com/2637309949/dolphin/cli/tempalte"
	"github.com/2637309949/dolphin/cli/packages/viper"
)

// Ctr struct
type Ctr struct {
}

// Build func
func (ctr *Ctr) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	var tmplCfgs []*gen.TmplCfg
	for _, c := range node.Controllers {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Controller":  c,
		}
		tmplCfg := &gen.TmplCfg{
			Text:     tempalte.TmplCtr,
			FilePath: path.Join(dir, viper.GetString("dir.app"), c.Name),
			Data:     data,
			Overlap:  gen.OverlapInc,
			Suffix:   ".go",
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
