// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/schema"
	"github.com/2637309949/dolphin/cli/tempalte"
)

// Doc struct
type Doc struct {
}

// Build func
func (m *Doc) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	var tmplCfgs []*gen.TmplCfg
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
	}
	tmplCfg := &gen.TmplCfg{
		Text:     tempalte.TmplDoc,
		FilePath: path.Join(dir, "docs", "build"),
		Data:     data,
		Overlap:  gen.OverlapSkip,
		Suffix:   ".sh",
	}
	tmplCfgs = append(tmplCfgs, tmplCfg)
	return tmplCfgs, nil
}
