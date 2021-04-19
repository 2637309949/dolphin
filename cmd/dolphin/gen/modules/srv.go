// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	ht "html/template"
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Srv struct
type Srv struct {
}

// Name defined pipe name
func (app *Srv) Name() string {
	return "srv"
}

// Build func
func (app *Srv) Build(dir string, args []string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	srvByte, _ := vfsutil.ReadFile(template.Assets, "srv.tmpl")
	for i := range node.Controllers {
		data := map[string]interface{}{
			"lt":          ht.HTML("<"),
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Controller":  node.Controllers[i],
			"Viper":       viper.GetViper(),
		}
		filename := utils.FileNameTrimSuffix(node.Controllers[i].Path)
		tmplCfg := &pipe.TmplCfg{
			Text:     string(srvByte),
			FilePath: path.Join(dir, viper.GetString("dir.srv"), filename+".go"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
