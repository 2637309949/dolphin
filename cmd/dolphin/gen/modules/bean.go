// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/shurcooL/httpfs/vfsutil"
)

// Bean struct
type Bean struct {
}

// Name defined pipe name
func (m *Bean) Name() string {
	return "bean"
}

// Build func
func (m *Bean) Build(dir string, args []string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	beanByte, _ := vfsutil.ReadFile(template.Assets, "bean.tmpl")
	for i := range node.Beans {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Bean":        node.Beans[i],
			"Viper":       viper.GetViper(),
		}
		filename := utils.FileNameTrimSuffix(node.Beans[i].Path)
		tmplCfg := &pipe.TmplCfg{
			Text:     string(beanByte),
			FilePath: path.Join(dir, viper.GetString("dir.model"), filename+".auto.go"),
			Data:     data,
			Overlap:  pipe.OverlapWrite,
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
