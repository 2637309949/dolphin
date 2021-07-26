// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"errors"
	"fmt"
	ht "html/template"
	"io/ioutil"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/api/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/api/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Reverse struct
type Reverse struct {
}

// Name defined pipe name
func (app *Reverse) Name() string {
	return "reverse"
}

// Pre defined
func (app *Reverse) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (app *Reverse) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (app *Reverse) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	var tbPath string
	tbByte := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, "table.tmpl")).([]byte)
	tmplCfgs := []*pipe.TmplCfg{}
	engines := []*xorm.Engine{}
	dataSources := []struct {
		DataSource string `xorm:"data_source"`
		DriverName string `xorm:"driver_name"`
	}{}
	if len(args) >= 2 {
		tbPath = path.Join(dir, args[1])
	} else {
		tbPath = utils.SearchFileInDirWithSuffix(dir, ".xml", func(s string) bool {
			data, _ := ioutil.ReadFile(s)
			return strings.Contains(string(data), "<table")
		})
		tbPath = path.Dir(tbPath)
	}
	if tbPath == "" {
		return tmplCfgs, errors.New("please give the path to generate the table")
	}
	engine, err := xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	if err != nil {
		return tmplCfgs, err
	}
	err = engine.SQL("select data_source, driver_name from sys_domain where is_delete=0 and status=1 and app_name=?", viper.GetString("app.name")).Find(&dataSources)
	if err != nil {
		return tmplCfgs, err
	}
	if len(dataSources) == 0 {
		return tmplCfgs, fmt.Errorf("not found any datasource in app_name:%v", viper.GetString("app.name"))
	}
	for i := range dataSources {
		engine, err := xorm.NewEngine(dataSources[i].DriverName, dataSources[i].DataSource)
		if err != nil {
			return tmplCfgs, err
		}
		engines = append(engines, engine)
	}
	for i1 := range engines {
		tables, err := engines[i1].DBMetas()
		if err != nil {
			return tmplCfgs, err
		}
		for i2 := range tables {
			logrus.Infoln(tables[i2].Name)
			meta := schema.Table{}
			meta.Name = tables[i2].Name
			meta.Desc = strings.ReplaceAll(tables[i2].Comment, "\n", "")
			meta.Columns = []*schema.Column{}
			cols := tables[i2].Columns()
			for i3 := range cols {
				c := schema.Column{}
				c.Name = strings.ToLower(cols[i3].Name)
				c.Desc = strings.ReplaceAll(cols[i3].Comment, "\n", "")
				switch cols[i3].SQLType.Name {
				case "VARCHAR", "TEXT", "LONGTEXT", "CHAR", "MEDIUMTEXT", "TINYTEXT":
					c.Type = "null.String"
				case "DATETIME", "TIMESTAMP":
					c.Type = "null.Time"
				case "ENUM", "INT", "SMALLINT", "BIGINT", "TINYINT":
					c.Type = "null.Int"
				case "DECIMAL":
					c.Type = "decimal.Decimal"
				case "BOOLEAN":
					c.Type = "decimal.Bool"
				case "FLOAT", "DOUBLE":
					c.Type = "null.Float"
				case "MEDIUMBLOB", "BLOB":
					c.Type = "[]byte"
				}

				// convert sql type
				switch cols[i3].SQLType.Name {
				case "VARCHAR", "TEXT":
					c.Xorm = fmt.Sprintf("%v", strings.ToLower(cols[i3].SQLType.Name))
					c.Xorm = fmt.Sprintf("%v(%v)", c.Xorm, cols[i3].SQLType.DefaultLength)
				case "DATETIME":
					c.Xorm = fmt.Sprintf("%v", strings.ToLower(cols[i3].SQLType.Name))
				case "ENUM", "INT", "BIGINT", "TINYINT", "SMALLINT", "MEDIUMINT", "INTEGER":
					if cols[i3].SQLType.Name == "ENUM" {
						c.Xorm = fmt.Sprintf("int(%v)", 10)
					} else {
						c.Xorm = fmt.Sprintf("%v", strings.ToLower(cols[i3].SQLType.Name))
						if cols[i3].SQLType.DefaultLength != 0 {
							c.Xorm = fmt.Sprintf("%v(%v)", c.Xorm, cols[i3].SQLType.DefaultLength)
						}
					}
				case "BOOLEAN":
					c.Xorm = fmt.Sprintf("%v", strings.ToLower(cols[i3].SQLType.Name))
				case "MEDIUMBLOB", "BLOB":
					c.Xorm = fmt.Sprintf("%v", strings.ToLower(cols[i3].SQLType.Name))
					if cols[i3].SQLType.DefaultLength != 0 {
						c.Xorm = fmt.Sprintf("%v(%v)", c.Xorm, cols[i3].SQLType.DefaultLength)
					}
				case "FLOAT", "DOUBLE":
					c.Xorm = fmt.Sprintf("%v", strings.ToLower(cols[i3].SQLType.Name))
					if cols[i3].SQLType.DefaultLength != 0 {
						if cols[i3].SQLType.DefaultLength2 != 0 {
							c.Xorm = fmt.Sprintf("%v(%v,%v)", c.Xorm, cols[i3].SQLType.DefaultLength, cols[i3].SQLType.DefaultLength2)
						} else {
							c.Xorm = fmt.Sprintf("%v(%v)", c.Xorm, cols[i3].SQLType.DefaultLength)
						}
					}
				case "DECIMAL":
					c.Xorm = fmt.Sprintf("%v", strings.ToLower(cols[i3].SQLType.Name))
					c.Xorm = fmt.Sprintf("%v(%v,%v)", c.Xorm, cols[i3].SQLType.DefaultLength, cols[i3].SQLType.DefaultLength2)
				}
				if cols[i3].IsPrimaryKey {
					c.Xorm = fmt.Sprintf("%v %v", c.Xorm, "pk")
				}
				if !cols[i3].Nullable {
					c.Xorm = fmt.Sprintf("%v %v", c.Xorm, "notnull")
				}
				if cols[i3].IsAutoIncrement {
					c.Xorm = fmt.Sprintf("%v %v", c.Xorm, "autoincr")
				}
				if cols[i3].Default != "" {
					c.Xorm = fmt.Sprintf(`%v default(%v)`, c.Xorm, cols[i3].Default)
				}
				c.Xorm = fmt.Sprintf("%v", c.Xorm)
				meta.Columns = append(meta.Columns, &c)
			}
			data := map[string]interface{}{
				"PackageName": parser.PackageName,
				"Name":        parser.Name,
				"Controllers": parser.Controllers,
				"Services":    parser.Services,
				"Tables":      parser.Tables,
				"Table":       meta,
				"Beans":       parser.Beans,
				"Viper":       viper.GetViper(),
				"lt":          ht.HTML("<"),
				"gt":          ht.HTML(">"),
			}
			tmplCfg := &pipe.TmplCfg{
				Text:     string(tbByte),
				FilePath: path.Join(tbPath, meta.Name+".xml"),
				Data:     data,
				Overlap:  pipe.OverlapSkip,
				GOFmt:    true,
			}
			tmplCfgs = append(tmplCfgs, tmplCfg)
		}
	}
	return tmplCfgs, nil
}