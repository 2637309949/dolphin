// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"encoding/xml"
	"fmt"
	ht "html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/shurcooL/httpfs/vfsutil"
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

// Name defined pipe name
func Name() string {
	return "sqltmp"
}

// Name defined pipe name
func (app *SQLTPL) Name() string {
	return "sqltpl"
}

// Build func
func (app *SQLTPL) Build(dir string, args []string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	var tmplCfgs []*pipe.TmplCfg
	var files []struct {
		Name    string
		Content ht.HTML
	}
	sqlByte, _ := vfsutil.ReadFile(template.Assets, "sql.tmpl")
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
					Content ht.HTML
				}{Name: sql.Id, Content: ht.HTML(ctstr)})
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
				Content ht.HTML
			}{Name: filepath.Base(path), Content: ht.HTML(ctstr)})
		}
		return nil
	}); err != nil {
		return tmplCfgs, err
	}
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
		"Application": node,
		"Files":       files,
	}
	tmplCfg := &pipe.TmplCfg{
		Text:     string(sqlByte),
		FilePath: path.Join(dir, viper.GetString("dir.sql"), "sql.auto.go"),
		Data:     data,
		Overlap:  pipe.OverlapWrite,
	}
	tmplCfgs = append(tmplCfgs, tmplCfg)
	return tmplCfgs, nil
}
