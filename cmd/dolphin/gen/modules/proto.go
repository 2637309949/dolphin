// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Proto struct
type Proto struct {
}

// Name defined pipe name
func (oa *Proto) Name() string {
	return "proto"
}

// Pre defined
func (oa *Proto) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (oa *Proto) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (oa *Proto) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	ctrByte, _ := vfsutil.ReadFile(template.Assets, "proto.tmpl")
	rpcByte, _ := vfsutil.ReadFile(template.Assets, "rpc.tmpl")
	rpcCliByte, _ := vfsutil.ReadFile(template.Assets, "rpc.cli.tmpl")
	for _, s := range parser.Services {
		data := map[string]interface{}{
			"PackageName": parser.PackageName,
			"Name":        parser.Name,
			"Service":     s,
			"Viper":       viper.GetViper(),
		}
		filename := utils.FileNameTrimSuffix(s.Path)
		tmplCfg := &pipe.TmplCfg{
			Text:     string(ctrByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), "proto", filename+".proto"),
			Data:     data,
			Overlap:  pipe.OverlapInc,
			GOFmt:    false,
			GOProto:  true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
		tmplCfg = &pipe.TmplCfg{
			Text:     string(rpcByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), filename+".go"),
			Data:     data,
			Overlap:  pipe.OverlapInc,
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
		tmplCfg = &pipe.TmplCfg{
			Text:     string(rpcCliByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), filename+".cli.go"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    true,
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
