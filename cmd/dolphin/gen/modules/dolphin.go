// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"fmt"
	ht "html/template"
	"path"
	"path/filepath"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Dolphin struct
type Dolphin struct {
}

// Name defined pipe name
func (m *Dolphin) Name() string {
	return "dol"
}

// Pre defined
func (m *Dolphin) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (m *Dolphin) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (m *Dolphin) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	tmpls := []*pipe.TmplCfg{}
	tmplArgs := map[string]interface{}{
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

	// main template
	mainByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "main.tmpl")).([]byte)
	tmpls = append(tmpls, &pipe.TmplCfg{
		Text:     string(mainByte),
		FilePath: path.Join(dir, "main.go"),
		Data:     tmplArgs,
		Overlap:  pipe.OverlapSkip,
		GOFmt:    true,
	})

	// app template
	appByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "app.tmpl")).([]byte)
	tmpls = append(tmpls, &pipe.TmplCfg{
		Text:     string(appByte),
		FilePath: path.Join(dir, viper.GetString("dir.app"), "app.go"),
		Data:     tmplArgs,
		Overlap:  pipe.OverlapSkip,
		GOFmt:    true,
	})

	// auto template
	autoByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "auto.tmpl")).([]byte)
	tmpls = append(tmpls, &pipe.TmplCfg{
		Text:     string(autoByte),
		FilePath: path.Join(dir, viper.GetString("dir.app"), "app.auto.go"),
		Data:     tmplArgs,
		Overlap:  pipe.OverlapWrite,
		GOFmt:    true,
	})

	// bean template
	beanByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "bean.tmpl")).([]byte)
	for i := range parser.Beans {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Bean"] = parser.Beans[i]
		filename := utils.FileNameTrimSuffix(parser.Beans[i].Path)
		tmpls = append(tmpls, &pipe.TmplCfg{
			Text:     string(beanByte),
			FilePath: path.Join(dir, viper.GetString("dir.model"), filename+".auto.go"),
			Data:     tmplArgs,
			Overlap:  pipe.OverlapWrite,
			GOFmt:    true,
		})
	}

	// ctr template
	ctrByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "ctr.tmpl")).([]byte)
	for i := range parser.Controllers {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Controller"] = parser.Controllers[i]
		filename := utils.FileNameTrimSuffix(parser.Controllers[i].Path)
		tmpls = append(tmpls, &pipe.TmplCfg{
			Text:     string(ctrByte),
			FilePath: path.Join(dir, viper.GetString("dir.app"), filename+".go"),
			Data:     tmplArgs,
			Overlap:  pipe.OverlapInc,
			GOFmt:    true,
		})
	}

	// table template
	modelByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "model.tmpl")).([]byte)
	for i := range parser.Tables {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Table"] = parser.Tables[i]
		filename := utils.FileNameTrimSuffix(parser.Tables[i].Path)
		tmpls = append(tmpls, &pipe.TmplCfg{
			Text:     string(modelByte),
			FilePath: path.Join(dir, viper.GetString("dir.model"), filename+".auto.go"),
			Data:     tmplArgs,
			Overlap:  pipe.OverlapWrite,
			GOFmt:    true,
		})
	}

	// proto template
	protoByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "proto.tmpl")).([]byte)
	rpcByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "rpc.tmpl")).([]byte)
	rpcCliByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "rpc.cli.tmpl")).([]byte)
	for i := range parser.Services {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Service"] = parser.Services[i]
		filename := utils.FileNameTrimSuffix(parser.Services[i].Path)
		tmpls = append(tmpls, &pipe.TmplCfg{
			Text:     string(protoByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), "proto", filename+".proto"),
			Data:     tmplArgs,
			Overlap:  pipe.OverlapInc,
			GOFmt:    false,
			GOProto:  true,
		})
		tmpls = append(tmpls, &pipe.TmplCfg{
			Text:     string(rpcByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), filename+".go"),
			Data:     tmplArgs,
			Overlap:  pipe.OverlapInc,
			GOFmt:    true,
		})
		tmpls = append(tmpls, &pipe.TmplCfg{
			Text:     string(rpcCliByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), filename+".cli.go"),
			Data:     tmplArgs,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    true,
		})
	}

	// sql template
	tplCache := map[string]bool{}
	countByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "count.tmpl")).([]byte)
	selectByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "select.tmpl")).([]byte)
	treeByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "tree.tmpl")).([]byte)
	for i := range parser.Controllers {
		for j := range parser.Controllers[i].APIS {
			ctr, table, funk := parser.Controllers[i], parser.Controllers[i].APIS[j].Table, parser.Controllers[i].APIS[j].Func
			tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
			tmplArgs["Controller"] = ctr
			tmplArgs["Api"] = ctr.APIS[j]
			if strings.TrimSpace(table) != "" && funk == "page" {
				cpath := path.Join(dir, viper.GetString("dir.sql"), ctr.Name, fmt.Sprintf("%v_%v_%v.tpl", ctr.Name, ctr.APIS[j].Name, "count"))
				if _, ok := tplCache[filepath.Base(cpath)]; !ok {
					tmpls = append(tmpls, &pipe.TmplCfg{
						Text:     string(countByte),
						FilePath: cpath,
						Data:     tmplArgs,
						Overlap:  pipe.OverlapSkip,
					})
					tplCache[filepath.Base(cpath)] = true
				}
				spath := path.Join(dir, viper.GetString("dir.sql"), ctr.Name, fmt.Sprintf("%v_%v_%v.tpl", ctr.Name, ctr.APIS[j].Name, "select"))
				if _, ok := tplCache[filepath.Base(spath)]; !ok {
					tmpls = append(tmpls, &pipe.TmplCfg{
						Text:     string(selectByte),
						FilePath: spath,
						Data:     tmplArgs,
						Overlap:  pipe.OverlapSkip,
					})
					tplCache[filepath.Base(spath)] = true
				}
			} else if strings.TrimSpace(table) != "" && funk == "tree" {
				cpath := path.Join(dir, viper.GetString("dir.sql"), parser.Controllers[i].Name, fmt.Sprintf("%v_%v.tpl", ctr.Name, ctr.APIS[j].Name))
				if _, ok := tplCache[filepath.Base(cpath)]; !ok {
					tmpls = append(tmpls, &pipe.TmplCfg{
						Text:     string(treeByte),
						FilePath: cpath,
						Data:     tmplArgs,
						Overlap:  pipe.OverlapSkip,
					})
					tplCache[filepath.Base(cpath)] = true
				}
			}
		}
	}

	// sqlmap template
	sqlmapByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "sqlmap.tmpl")).([]byte)
	for i := range parser.Tables {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Table"] = parser.Tables[i]
		filename := utils.FileNameTrimSuffix(parser.Tables[i].Path)
		tmpls = append(tmpls, &pipe.TmplCfg{
			Text:     string(sqlmapByte),
			FilePath: path.Join(dir, viper.GetString("dir.sql"), viper.GetString("dir.sqlmap"), filename+".xml"),
			Data:     tmplArgs,
			Overlap:  pipe.OverlapSkip,
		})
	}

	// srv template
	srvByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "srv.tmpl")).([]byte)
	for i := range parser.Controllers {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Controller"] = parser.Controllers[i]
		filename := utils.FileNameTrimSuffix(parser.Controllers[i].Path)
		tmpls = append(tmpls, &pipe.TmplCfg{
			Text:     string(srvByte),
			FilePath: path.Join(dir, viper.GetString("dir.srv"), filename+".go"),
			Data:     tmplArgs,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    true,
		})
	}

	// svc template
	svcByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "svc.tmpl")).([]byte)
	svcDbByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "svc.db.tmpl")).([]byte)
	svcCache := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "svc.cache.tmpl")).([]byte)
	svcHttp := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "svc.http.tmpl")).([]byte)
	tmpls = append(tmpls, &pipe.TmplCfg{
		Text:     string(svcDbByte),
		FilePath: path.Join(dir, viper.GetString("dir.svc"), "db.go"),
		Data:     tmplArgs,
		Overlap:  pipe.OverlapSkip,
		GOFmt:    true,
	})
	tmpls = append(tmpls, &pipe.TmplCfg{
		Text:     string(svcCache),
		FilePath: path.Join(dir, viper.GetString("dir.svc"), "cache.go"),
		Data:     tmplArgs,
		Overlap:  pipe.OverlapSkip,
		GOFmt:    true,
	})
	tmpls = append(tmpls, &pipe.TmplCfg{
		Text:     string(svcByte),
		FilePath: path.Join(dir, viper.GetString("dir.svc"), "svc.go"),
		Data:     tmplArgs,
		Overlap:  pipe.OverlapSkip,
		GOFmt:    true,
	})
	tmpls = append(tmpls, &pipe.TmplCfg{
		Text:     string(svcHttp),
		FilePath: path.Join(dir, viper.GetString("dir.svc"), "http.go"),
		Data:     tmplArgs,
		Overlap:  pipe.OverlapSkip,
		GOFmt:    true,
	})
	return tmpls, nil
}
