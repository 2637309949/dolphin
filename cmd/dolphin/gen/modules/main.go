// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
)

// Main struct
type Main struct {
}

// Name defined pipe name
func (m *Main) Name() string {
	return "main"
}

// Build func
func (m *Main) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
	}
	return []*pipe.TmplCfg{
		&pipe.TmplCfg{
			Text:     template.TmplMain,
			FilePath: path.Join(dir, "main"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			Suffix:   ".go",
			GOFmt:    true,
		},
	}, nil
}
