// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// More struct
type More struct {
}

// Name defined pipe name
func (m *More) Name() string {
	return "more"
}

// Build func
func (m *More) Build(dir string, args []string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	ctrByte, _ := vfsutil.ReadFile(template.Assets, "more.ctr.tmpl")
	tbByte, _ := vfsutil.ReadFile(template.Assets, "more.tb.tmpl")
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
		"PackageName": node.PackageName,
		"Name":        node.Name,
		"Viper":       viper.GetViper(),
		"more":        args[0],
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
