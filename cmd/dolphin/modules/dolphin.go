// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"fmt"
	"html/template"
	"path"
	"path/filepath"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/template/dist"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Dolphin struct
type Dolphin struct {
}

// Name defined parser name
func (m *Dolphin) Name() string {
	return "dol"
}

// Pre defined
func (m *Dolphin) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (m *Dolphin) After(*parser.AppParser, []*parser.TmplCfg) error {
	return nil
}

// Build func
func (m *Dolphin) Build(dir string, args []string, appParser *parser.AppParser) ([]*parser.TmplCfg, error) {
	tmpls := []*parser.TmplCfg{}
	tmplArgs := map[string]interface{}{
		"PackageName": appParser.PackageName,
		"Name":        appParser.Name,
		"Controllers": appParser.Controllers,
		"Services":    appParser.Services,
		"Tables":      appParser.Tables,
		"Beans":       appParser.Beans,
		"Viper":       viper.GetViper(),
		"lt":          template.HTML("<"),
		"gt":          template.HTML(">"),
	}

	// main template
	mainByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "main.tmpl")).([]byte)
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(mainByte),
		FilePath: path.Join(dir, "main.go"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapSkip,
		GOFmt:    true,
	})

	// app template
	appByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "app.tmpl")).([]byte)
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(appByte),
		FilePath: path.Join(dir, viper.GetString("dir.api"), "app.go"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapSkip,
		GOFmt:    true,
	})

	// auto template
	autoByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "auto.tmpl")).([]byte)
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(autoByte),
		FilePath: path.Join(dir, viper.GetString("dir.api"), "app.auto.go"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapWrite,
		GOFmt:    true,
	})

	// docker template
	drByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "docker.tmpl")).([]byte)
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(drByte),
		FilePath: path.Join(dir, "Dockerfile"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapSkip,
	})

	// html template
	affirm, login := strings.Join(strings.Split(viper.GetString("oauth.affirm"), "/"), "/"), strings.Join(strings.Split(viper.GetString("oauth.login"), "/"), "/")
	affirmPath, loginPath := affirm[0:len(affirm)-len(filepath.Ext(affirm))], login[0:len(login)-len(filepath.Ext(login))]
	affirmByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "static/web/affirm.html")).([]byte)
	loginByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "static/web/login.html")).([]byte)
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(affirmByte),
		FilePath: path.Join(dir, affirmPath+".html"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapSkip,
	})
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(loginByte),
		FilePath: path.Join(dir, loginPath+".html"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapSkip,
	})

	// tool template
	toolByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "tool.tmpl")).([]byte)
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(toolByte),
		FilePath: path.Join(dir, viper.GetString("dir.util"), "common.go"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapSkip,
	})

	// bean template
	beanByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "bean.tmpl")).([]byte)
	for i := range appParser.Beans {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Bean"] = appParser.Beans[i]
		tmplArgs["Viper"] = viper.GetViper()
		filename := utils.FileNameTrimSuffix(appParser.Beans[i].Path)
		tmpls = append(tmpls, &parser.TmplCfg{
			Text:     string(beanByte),
			FilePath: path.Join(dir, viper.GetString("dir.types"), filename+".auto.go"),
			Data:     tmplArgs,
			Overlap:  parser.OverlapWrite,
			GOFmt:    true,
		})
	}

	// ctr template
	ctrByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "ctr.tmpl")).([]byte)
	for i := range appParser.Controllers {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Controller"] = appParser.Controllers[i]
		tmplArgs["Viper"] = viper.GetViper()
		filename := utils.FileNameTrimSuffix(appParser.Controllers[i].Path)
		tmpls = append(tmpls, &parser.TmplCfg{
			Text:     string(ctrByte),
			FilePath: path.Join(dir, viper.GetString("dir.api"), filename+".go"),
			Data:     tmplArgs,
			Overlap:  parser.OverlapInc,
			GOFmt:    true,
		})
	}

	// table template
	typesByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "types.tmpl")).([]byte)
	for i := range appParser.Tables {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Table"] = appParser.Tables[i]
		tmplArgs["Viper"] = viper.GetViper()
		filename := utils.FileNameTrimSuffix(appParser.Tables[i].Path)
		tmpls = append(tmpls, &parser.TmplCfg{
			Text:     string(typesByte),
			FilePath: path.Join(dir, viper.GetString("dir.types"), filename+".auto.go"),
			Data:     tmplArgs,
			Overlap:  parser.OverlapWrite,
			GOFmt:    true,
		})
	}

	// errors template
	errByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "errors.tmpl")).([]byte)
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(errByte),
		FilePath: path.Join(dir, viper.GetString("dir.types"), "errors.go"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapSkip,
	})

	// proto template
	protoByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "proto.tmpl")).([]byte)
	rpcByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "rpc.tmpl")).([]byte)
	rpcCliByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "rpc.cli.tmpl")).([]byte)
	for i := range appParser.Services {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Service"] = appParser.Services[i]
		tmplArgs["Viper"] = viper.GetViper()
		filename := utils.FileNameTrimSuffix(appParser.Services[i].Path)
		tmpls = append(tmpls, &parser.TmplCfg{
			Text:     string(protoByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), "proto", filename+".proto"),
			Data:     tmplArgs,
			Overlap:  parser.OverlapInc,
			GOFmt:    false,
			GOProto:  true,
		})
		tmpls = append(tmpls, &parser.TmplCfg{
			Text:     string(rpcByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), filename+".go"),
			Data:     tmplArgs,
			Overlap:  parser.OverlapInc,
			GOFmt:    true,
		})
		tmpls = append(tmpls, &parser.TmplCfg{
			Text:     string(rpcCliByte),
			FilePath: path.Join(dir, viper.GetString("dir.rpc"), filename+".cli.go"),
			Data:     tmplArgs,
			Overlap:  parser.OverlapSkip,
			GOFmt:    true,
		})
	}

	// sql template
	tplCache := map[string]bool{}
	countByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "count.tmpl")).([]byte)
	selectByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "select.tmpl")).([]byte)
	treeByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "tree.tmpl")).([]byte)
	for i := range appParser.Controllers {
		for j := range appParser.Controllers[i].APIS {
			ctr, table, funk := appParser.Controllers[i], appParser.Controllers[i].APIS[j].Table, appParser.Controllers[i].APIS[j].Func
			tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
			tmplArgs["Controller"] = ctr
			tmplArgs["Api"] = ctr.APIS[j]
			tmplArgs["Viper"] = viper.GetViper()
			if strings.TrimSpace(table) != "" && funk == "page" {
				cpath := path.Join(dir, viper.GetString("dir.sql"), ctr.Name, fmt.Sprintf("%v_%v_%v.tpl", ctr.Name, ctr.APIS[j].Name, "count"))
				if _, ok := tplCache[filepath.Base(cpath)]; !ok {
					tmpls = append(tmpls, &parser.TmplCfg{
						Text:     string(countByte),
						FilePath: cpath,
						Data:     tmplArgs,
						Overlap:  parser.OverlapSkip,
					})
					tplCache[filepath.Base(cpath)] = true
				}
				spath := path.Join(dir, viper.GetString("dir.sql"), ctr.Name, fmt.Sprintf("%v_%v_%v.tpl", ctr.Name, ctr.APIS[j].Name, "select"))
				if _, ok := tplCache[filepath.Base(spath)]; !ok {
					tmpls = append(tmpls, &parser.TmplCfg{
						Text:     string(selectByte),
						FilePath: spath,
						Data:     tmplArgs,
						Overlap:  parser.OverlapSkip,
					})
					tplCache[filepath.Base(spath)] = true
				}
			} else if strings.TrimSpace(table) != "" && funk == "tree" {
				cpath := path.Join(dir, viper.GetString("dir.sql"), appParser.Controllers[i].Name, fmt.Sprintf("%v_%v.tpl", ctr.Name, ctr.APIS[j].Name))
				if _, ok := tplCache[filepath.Base(cpath)]; !ok {
					tmpls = append(tmpls, &parser.TmplCfg{
						Text:     string(treeByte),
						FilePath: cpath,
						Data:     tmplArgs,
						Overlap:  parser.OverlapSkip,
					})
					tplCache[filepath.Base(cpath)] = true
				}
			}
		}
	}

	// sqlmap template
	sqlmapByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "sqlmap.tmpl")).([]byte)
	for i := range appParser.Tables {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Table"] = appParser.Tables[i]
		tmplArgs["Viper"] = viper.GetViper()
		filename := utils.FileNameTrimSuffix(appParser.Tables[i].Path)
		tmpls = append(tmpls, &parser.TmplCfg{
			Text:     string(sqlmapByte),
			FilePath: path.Join(dir, viper.GetString("dir.sql"), viper.GetString("dir.sqlmap"), filename+".xml"),
			Data:     tmplArgs,
			Overlap:  parser.OverlapSkip,
		})
	}

	// srv template
	srvByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "srv.tmpl")).([]byte)
	for i := range appParser.Controllers {
		tmplArgs := utils.Copy(tmplArgs).(map[string]interface{})
		tmplArgs["Controller"] = appParser.Controllers[i]
		tmplArgs["Viper"] = viper.GetViper()
		filename := utils.FileNameTrimSuffix(appParser.Controllers[i].Path)
		tmpls = append(tmpls, &parser.TmplCfg{
			Text:     string(srvByte),
			FilePath: path.Join(dir, viper.GetString("dir.srv"), filename+".go"),
			Data:     tmplArgs,
			Overlap:  parser.OverlapSkip,
			GOFmt:    true,
		})
	}

	// svc template
	svcByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "svc.tmpl")).([]byte)
	helperByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "svc.helper.tmpl")).([]byte)
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(svcByte),
		FilePath: path.Join(dir, viper.GetString("dir.svc"), "svc.go"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapSkip,
		GOFmt:    true,
	})
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(helperByte),
		FilePath: path.Join(dir, viper.GetString("dir.svc"), "svc.helper.go"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapSkip,
		GOFmt:    true,
	})

	// x_test template
	xTestByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "x_test.tmpl")).([]byte)
	xPlatformByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "x_platform_test.tmpl")).([]byte)
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(xTestByte),
		FilePath: path.Join(dir, "x_test.go"),
		Data:     tmplArgs,
		Overlap:  parser.OverlapWrite,
		GOFmt:    true,
	})
	tmpls = append(tmpls, &parser.TmplCfg{
		Text:     string(xPlatformByte),
		FilePath: path.Join(dir, fmt.Sprintf("x_%v_test.go", appParser.Name)),
		Data:     tmplArgs,
		Overlap:  parser.OverlapInc,
		GOFmt:    true,
	})
	return tmpls, nil
}
