// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/client/gen"
	"github.com/2637309949/dolphin/client/gen/template"
	"github.com/2637309949/dolphin/client/schema"
)

// Main struct
type Main struct {
}

// Build func
func (m *Main) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
	}
	return []*gen.TmplCfg{
		&gen.TmplCfg{
			Text:     template.TmplMain,
			FilePath: path.Join(dir, "main"),
			Data:     data,
			Overlap:  gen.OverlapSkip,
			Suffix:   ".go",
		},
	}, nil
}
