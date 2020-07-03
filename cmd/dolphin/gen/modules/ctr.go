// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"
	"path/filepath"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/shurcooL/httpfs/vfsutil"
)

// Ctr struct
type Ctr struct {
}

// Name defined pipe name
func (ctr *Ctr) Name() string {
	return "ctr"
}

// Build func
func (ctr *Ctr) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	ctrByte, _ := vfsutil.ReadFile(template.Assets, "ctr.tmpl")
	for _, c := range node.Controllers {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Controller":  c,
		}
		extension, filename := filepath.Ext(c.Path), filepath.Base(c.Path)
		filename = filename[0 : len(filename)-len(extension)]

		tmplCfg := &pipe.TmplCfg{
			Text:     string(ctrByte),
			FilePath: path.Join(dir, viper.GetString("dir.app"), filename+".go"),
			Data:     data,
			Overlap:  pipe.OverlapInc,
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
