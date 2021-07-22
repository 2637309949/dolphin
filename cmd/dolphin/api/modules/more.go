// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	ht "html/template"
	"io/ioutil"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/api/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/api/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// More struct
type More struct {
}

// Name defined pipe name
func (m *More) Name() string {
	return "more"
}

// Pre defined
func (m *More) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (m *More) After(parser *parser.AppParser, cfgs []*pipe.TmplCfg) error {
	parser.WalkXML(funk.Map(cfgs, func(cfg *pipe.TmplCfg) string { return cfg.FilePath }).([]string)...)
	return nil
}

// Build func
func (m *More) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	ctrByte, err := vfsutil.ReadFile(template.Assets, "more.ctr.tmpl")
	if err != nil {
		return []*pipe.TmplCfg{}, err
	}
	tbByte, err := vfsutil.ReadFile(template.Assets, "more.tb.tmpl")
	if err != nil {
		return []*pipe.TmplCfg{}, err
	}
	if len(args) < 1 {
		logrus.Warn("Please give the path to generate the table")
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
		"PackageName": parser.PackageName,
		"Name":        parser.Name,
		"Controllers": parser.Controllers,
		"Services":    parser.Services,
		"Tables":      parser.Tables,
		"Beans":       parser.Beans,
		"Viper":       viper.GetViper(),
		"more":        args[0],
		"lt":          ht.HTML("<"),
		"gt":          ht.HTML(">"),
	}
	tmplCfgs = append(tmplCfgs,
		&pipe.TmplCfg{
			Text:     string(ctrByte),
			FilePath: path.Join(path.Dir(ctrPath), args[0]+".xml"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    false,
		},
		&pipe.TmplCfg{
			Text:     string(tbByte),
			FilePath: path.Join(path.Dir(tbPath), args[0]+".xml"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    false,
		},
	)
	return tmplCfgs, nil
}
