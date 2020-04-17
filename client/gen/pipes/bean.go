// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/client/gen/pipe"
	"github.com/2637309949/dolphin/client/gen/template"
	"github.com/2637309949/dolphin/client/schema"
	"github.com/2637309949/dolphin/packages/viper"
)

// Bean struct
type Bean struct {
}

// Name defined pipe name
func (m *Bean) Name() string {
	return "bean"
}

// Build func
func (m *Bean) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	for _, bean := range node.Beans {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Bean":        bean,
		}
		tmplCfg := &pipe.TmplCfg{
			Text:     template.TmplBean,
			FilePath: path.Join(dir, viper.GetString("dir.model"), bean.Name+".auto"),
			Data:     data,
			Overlap:  pipe.OverlapWrite,
			Suffix:   ".go",
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
