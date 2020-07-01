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

// Proto struct
type Proto struct {
}

// Name defined pipe name
func (oa *Proto) Name() string {
	return "proto"
}

// Build func
func (oa *Proto) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	ctrByte, _ := vfsutil.ReadFile(template.Assets, "proto.tmpl")
	rpcByte, _ := vfsutil.ReadFile(template.Assets, "rpc.tmpl")
	for _, s := range node.Services {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Service":     s,
		}
		tmplCfg := &pipe.TmplCfg{
			Text:     string(ctrByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), "proto", s.Name+".proto"),
			Data:     data,
			Overlap:  pipe.OverlapInc,
			GOFmt:    false,
			GOProto:  true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
		tmplCfg = &pipe.TmplCfg{
			Text:     string(rpcByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), s.Name+".go"),
			Data:     data,
			Overlap:  pipe.OverlapInc,
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
