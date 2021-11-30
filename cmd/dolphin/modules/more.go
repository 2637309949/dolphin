// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"context"
	"html/template"
	"io/ioutil"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/template/dist"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// More struct
type More struct {
}

// Name defined parser name
func (m *More) Name() string {
	return "more"
}

// Pre defined
func (m *More) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (m *More) After(appParser *parser.AppParser, cfgs []*parser.TmplCfg) error {
	appParser.WalkXML(funk.Map(cfgs, func(cfg *parser.TmplCfg) string { return cfg.FilePath }).([]string)...)
	return nil
}

// Build func
func (m *More) Build(dir string, args []string, appParser *parser.AppParser) ([]*parser.TmplCfg, error) {
	var tmplCfgs []*parser.TmplCfg
	ctrByte, err := vfsutil.ReadFile(dist.Assets, "more.ctr.tmpl")
	if err != nil {
		return []*parser.TmplCfg{}, err
	}
	tbByte, err := vfsutil.ReadFile(dist.Assets, "more.tb.tmpl")
	if err != nil {
		return []*parser.TmplCfg{}, err
	}
	if len(args) < 1 {
		logrus.Warn(context.TODO(), "Please give the path to generate the table")
		return tmplCfgs, nil
	}
	ctrPath := utils.SearchFileInDirWithSuffix(dir, ".xml", func(s string) bool {
		data, _ := ioutil.ReadFile(s)
		return strings.Contains(string(data), "<controller")
	})
	tbPath := utils.SearchFileInDirWithSuffix(dir, ".xml", func(s string) bool {
		data, _ := ioutil.ReadFile(s)
		return strings.Contains(string(data), "<table")
	})
	data := map[string]interface{}{
		"PackageName": appParser.PackageName,
		"Name":        appParser.Name,
		"Controllers": appParser.Controllers,
		"Tables":      appParser.Tables,
		"Beans":       appParser.Beans,
		"Viper":       viper.GetViper(),
		"more":        args[0],
		"lt":          template.HTML("<"),
		"gt":          template.HTML(">"),
	}
	tmplCfgs = append(tmplCfgs,
		&parser.TmplCfg{
			Text:     string(ctrByte),
			FilePath: path.Join(path.Dir(ctrPath), args[0]+".xml"),
			Data:     data,
			Overlap:  parser.OverlapSkip,
			GOFmt:    false,
		},
		&parser.TmplCfg{
			Text:     string(tbByte),
			FilePath: path.Join(path.Dir(tbPath), args[0]+".xml"),
			Data:     data,
			Overlap:  parser.OverlapSkip,
			GOFmt:    false,
		},
	)
	return tmplCfgs, nil
}
