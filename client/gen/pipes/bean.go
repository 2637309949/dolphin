// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/client/gen"
	"github.com/2637309949/dolphin/client/schema"
	"github.com/2637309949/dolphin/client/tempalte"
)

// Bean struct
type Bean struct {
}

// Build func
func (m *Bean) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	var tmplCfgs []*gen.TmplCfg
	for _, bean := range node.Beans {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Bean":        bean,
		}
		tmplCfg := &gen.TmplCfg{
			Text:     tempalte.TmplBean,
			FilePath: path.Join(dir, "model", bean.Name+".auto"),
			Data:     data,
			Overlap:  gen.OverlapWrite,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
