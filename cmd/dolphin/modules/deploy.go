// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"html/template"
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/template/dist"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Deploy struct
type Deploy struct {
}

// Name defined parser name
func (dp *Deploy) Name() string {
	return "deploy"
}

// Pre defined
func (dp *Deploy) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (dp *Deploy) After(*parser.AppParser, []*parser.TmplCfg) error {
	return nil
}

// Build func
func (dp *Deploy) Build(dir string, args []string, appParser *parser.AppParser) ([]*parser.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": appParser.PackageName,
		"Name":        appParser.Name,
		"Controllers": appParser.Controllers,
		"Tables":      appParser.Tables,
		"Beans":       appParser.Beans,
		"Viper":       viper.GetViper(),
		"lt":          template.HTML("<"),
		"gt":          template.HTML(">"),
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
	cfgs := []*parser.TmplCfg{}
	for key, value := range tmpl2file {
		dByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, key)).([]byte)
		cfgs = append(cfgs, &parser.TmplCfg{
			Text:     string(dByte),
			FilePath: path.Join(dir, viper.GetString("dir.k8s"), value),
			Data:     data,
			Overlap:  parser.OverlapSkip,
		})
	}
	return cfgs, nil
}
