// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	ht "html/template"
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
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
	cfgByte, _ := vfsutil.ReadFile(template.Assets, "k8s.cfg.tmpl")
	dpyByte, _ := vfsutil.ReadFile(template.Assets, "k8s.dpy.tmpl")
	ingByte, _ := vfsutil.ReadFile(template.Assets, "k8s.ing.tmpl")
	srvByte, _ := vfsutil.ReadFile(template.Assets, "k8s.srv.tmpl")
	tlsByte, _ := vfsutil.ReadFile(template.Assets, "k8s.tls.tmpl")
	dcrByte, _ := vfsutil.ReadFile(template.Assets, "docker.tmpl")
	return []*pipe.TmplCfg{
		{
			Text:     string(cfgByte),
			FilePath: path.Join(dir, viper.GetString("dir.k8s"), "configmap.yaml"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		},
		{
			Text:     string(dpyByte),
			FilePath: path.Join(dir, viper.GetString("dir.k8s"), "deployment.yaml"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		},
		{
			Text:     string(ingByte),
			FilePath: path.Join(dir, viper.GetString("dir.k8s"), "ingress.yaml"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		},
		{
			Text:     string(srvByte),
			FilePath: path.Join(dir, viper.GetString("dir.k8s"), "service.yaml"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		},
		{
			Text:     string(tlsByte),
			FilePath: path.Join(dir, viper.GetString("dir.k8s"), "tls.yaml"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		},
		{
			Text:     string(dcrByte),
			FilePath: path.Join(dir, "Dockerfile"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		},
	}, nil
}
