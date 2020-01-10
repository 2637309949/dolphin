// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/gen/tempalte"
	"github.com/2637309949/dolphin/cli/packages/viper"
	"github.com/2637309949/dolphin/cli/schema"
)

// SQLMap struct
type SQLMap struct {
}

// Build func
func (app *SQLMap) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	var tmplCfgs []*gen.TmplCfg
	for _, t := range node.Tables {
		data := map[string]interface{}{
			"PackageName": node.PackageName,
			"Name":        node.Name,
			"Application": node,
			"Table":       t,
		}
		tmplCfg := &gen.TmplCfg{
			Text:     tempalte.TmplSQLMap,
			FilePath: path.Join(dir, viper.GetString("dir.sql"), viper.GetString("dir.sql.sqlmap"), t.Name),
			Data:     data,
			Overlap:  gen.OverlapWrite,
			Suffix:   ".xml",
		}
		tmplCfgs = append(tmplCfgs, tmplCfg)
	}
	return tmplCfgs, nil
}
