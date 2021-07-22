// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	ht "html/template"
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/api/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/api/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Deploy struct
type Deploy struct {
}

// Name defined pipe name
func (dp *Deploy) Name() string {
	return "deploy"
}

// Pre defined
func (dp *Deploy) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (dp *Deploy) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (dp *Deploy) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": parser.PackageName,
		"Name":        parser.Name,
		"Controllers": parser.Controllers,
		"Services":    parser.Services,
		"Tables":      parser.Tables,
		"Beans":       parser.Beans,
		"Viper":       viper.GetViper(),
		"lt":          ht.HTML("<"),
		"gt":          ht.HTML(">"),
	}
	tmpl2file := map[string]string{
		"k8s.configmap.tmpl":       "configmap.yaml",
		"k8s.statefulset.tmpl":     "statefulset.yaml",
		"k8s.ingress.tmpl":         "ingress.yaml",
		"k8s.service.tmpl":         "service.yaml",
		"k8s.tls.tmpl":             "tls.yaml",
		"k8s.pvc.tmpl":             "pvc.yaml",
		"k8s.gateway.tmpl":         "gateway.yaml",
		"k8s.virtualservice.tmpl":  "virtualservice.yaml",
		"k8s.destinationrule.tmpl": "destinationrule.yaml",
	}
	cfgs := []*pipe.TmplCfg{}
	for key, value := range tmpl2file {
		dByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, key)).([]byte)
		cfgs = append(cfgs, &pipe.TmplCfg{
			Text:     string(dByte),
			FilePath: path.Join(dir, viper.GetString("dir.k8s"), value),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		})
	}
	return cfgs, nil
}
