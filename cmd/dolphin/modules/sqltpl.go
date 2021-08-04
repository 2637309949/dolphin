// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/template/dist"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// HasSuffix defined suffix of string
func HasSuffix(s string, suffix ...string) bool {
	for _, v := range suffix {
		if strings.HasSuffix(s, v) {
			return true
		}
	}
	return false
}

// SQLTPL struct
type SQLTPL struct {
}

// Name defined parser name
func (app *SQLTPL) Name() string {
	return "sqltpl"
}

// Pre defined
func (app *SQLTPL) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (app *SQLTPL) After(*parser.AppParser, []*parser.TmplCfg) error {
	return nil
}

// Build func
func (app *SQLTPL) Build(dir string, args []string, appParser *parser.AppParser) ([]*parser.TmplCfg, error) {
	var tmplCfgs []*parser.TmplCfg
	var files []struct {
		Name    string
		Content template.HTML
	}
	sqlByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "sql.tmpl")).([]byte)

	if err := filepath.Walk(path.Join(dir, viper.GetString("dir.sql")), func(path string, info os.FileInfo, err error) error {
		if HasSuffix(path, ".xml") {
			ct, _ := ioutil.ReadFile(path)
			var result xorm.XmlSql
			err = xml.Unmarshal(ct, &result)
			if err != nil {
				return err
			}
			trg := regexp.MustCompile("(`[a-zA-Z1-9_]*`)")
			for _, sql := range result.Sql {
				groups := trg.FindAllStringSubmatch(sql.Value, -1)
				repmap := map[string]string{}
				for _, v := range groups {
					repmap[v[0]] = fmt.Sprintf(`%v + "%v" + %v`, "`", v[0], "`")
				}
				ctstr := sql.Value
				for key, v := range repmap {
					ctstr = strings.ReplaceAll(ctstr, key, v)
				}
				files = append(files, struct {
					Name    string
					Content template.HTML
				}{Name: sql.Id, Content: template.HTML(ctstr)})
			}
		} else if HasSuffix(path, ".tpl", ".stpl", ".jet") {
			ct, _ := ioutil.ReadFile(path)
			trg := regexp.MustCompile("(`[a-zA-Z1-9_]*`)")
			groups := trg.FindAllStringSubmatch(string(ct), -1)
			repmap := map[string]string{}
			for _, v := range groups {
				repmap[v[0]] = fmt.Sprintf(`%v+"%v"+%v`, "`", v[0], "`")
			}
			ctstr := string(ct)
			for key, v := range repmap {
				ctstr = strings.ReplaceAll(ctstr, key, v)
			}
			files = append(files, struct {
				Name    string
				Content template.HTML
			}{Name: filepath.Base(path), Content: template.HTML(ctstr)})
		}
		return nil
	}); err != nil {
		return tmplCfgs, err
	}
	data := map[string]interface{}{
		"PackageName": appParser.PackageName,
		"Name":        appParser.Name,
		"Controllers": appParser.Controllers,
		"Services":    appParser.Services,
		"Tables":      appParser.Tables,
		"Beans":       appParser.Beans,
		"Viper":       viper.GetViper(),
		"Files":       files,
		"lt":          template.HTML("<"),
		"gt":          template.HTML(">"),
	}
	tmplCfg := &parser.TmplCfg{
		Text:     string(sqlByte),
		FilePath: path.Join(dir, viper.GetString("dir.sql"), "sql.auto.go"),
		Data:     data,
		Overlap:  parser.OverlapWrite,
	}
	tmplCfgs = append(tmplCfgs, tmplCfg)
	return tmplCfgs, nil
}
