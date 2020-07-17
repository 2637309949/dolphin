// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"fmt"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/shurcooL/httpfs/vfsutil"
)

// Table struct
type Table struct {
}

// Name defined pipe name
func (app *Table) Name() string {
	return "table"
}

// Build func
func (app *Table) Build(dir string, args []string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	tbByte, _ := vfsutil.ReadFile(template.Assets, "table.tmpl")
	tmplCfgs := []*pipe.TmplCfg{}
	engines := []*xorm.Engine{}
	dataSources := []struct {
		DataSource string `xorm:"data_source"`
		DriverName string `xorm:"driver_name"`
	}{}
	if len(args) < 2 {
		logrus.Warn("Please give the path to generate the table")
		return tmplCfgs, nil
	}
	engine, err := xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	if err != nil {
		return tmplCfgs, err
	}
	err = engine.SQL("select data_source, driver_name from sys_domain where del_flag=0 and status=1").Find(&dataSources)
	if err != nil {
		return tmplCfgs, err
	}
	for _, ds := range dataSources {
		engine, err := xorm.NewEngine(ds.DriverName, ds.DataSource)
		if err != nil {
			return tmplCfgs, err
		}
		engines = append(engines, engine)
	}
	for _, engine := range engines {
		tables, _ := engine.DBMetas()
		for _, tb := range tables {
			meta := schema.Table{}
			meta.Name = tb.Name
			meta.Desc = tb.Comment
			meta.Columns = []*schema.Column{}
			for _, col := range tb.Columns() {
				c := schema.Column{}
				c.Name = col.Name
				c.Desc = col.Comment
				// convert golang type
				switch col.SQLType.Name {
				case "VARCHAR", "TEXT":
					c.Type = "null.String"
				case "DATETIME":
					c.Type = "null.Time"
				case "INT":
					c.Type = "null.Int"
				case "DECIMAL":
					c.Type = "decimal.Decimal"
				case "FLOAT":
					c.Type = "null.Float"
				}
				// convert sql type
				switch col.SQLType.Name {
				case "VARCHAR", "TEXT":
					c.Xorm = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
					if col.SQLType.DefaultLength != 0 {
						c.Xorm = fmt.Sprintf("%v(%v)", c.Xorm, col.SQLType.DefaultLength)
					}
				case "DATETIME":
					c.Xorm = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
				case "INT":
					c.Xorm = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
					if col.SQLType.DefaultLength != 0 {
						c.Xorm = fmt.Sprintf("%v(%v)", c.Xorm, col.SQLType.DefaultLength)
					}
				case "FLOAT":
					if col.SQLType.DefaultLength != 0 {
						c.Xorm = fmt.Sprintf("%v(%v, %v)", c.Xorm, col.SQLType.DefaultLength, col.SQLType.DefaultLength2)
					}
				case "DECIMAL":
					c.Xorm = fmt.Sprintf("%v", strings.ToLower(col.SQLType.Name))
					if col.SQLType.DefaultLength != 0 {
						c.Xorm = fmt.Sprintf("%v(%v, %v)", c.Xorm, col.SQLType.DefaultLength, col.SQLType.DefaultLength2)
					}
				}
				if col.IsPrimaryKey {
					c.Xorm = fmt.Sprintf("%v %v", c.Xorm, "pk")
				}
				if col.Nullable {
					c.Xorm = fmt.Sprintf("%v %v", c.Xorm, "notnull")
				}
				if col.IsAutoIncrement {
					c.Xorm = fmt.Sprintf("%v %v", c.Xorm, "autoincr")
				}
				meta.Columns = append(meta.Columns, &c)
			}
			data := map[string]interface{}{
				"PackageName": node.PackageName,
				"Name":        node.Name,
				"Table":       meta,
			}
			tmplCfg := &pipe.TmplCfg{
				Text:     string(tbByte),
				FilePath: path.Join(dir, args[1], meta.Name+".xml"),
				Data:     data,
				Overlap:  pipe.OverlapSkip,
				GOFmt:    true,
			}
			tmplCfgs = append(tmplCfgs, tmplCfg)
		}
	}
	return tmplCfgs, nil
}
