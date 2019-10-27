// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/client/gen"
	"github.com/2637309949/dolphin/client/schema"
	"github.com/2637309949/dolphin/client/tempalte"
)

// Invoke struct
type Invoke struct {
}

// Build func
func (auto *Invoke) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Controllers": node.Controllers,
		"Tables":      node.Tables,
	}
	return []*gen.TmplCfg{
		&gen.TmplCfg{
			Text:     tempalte.TmplAuto,
			FilePath: path.Join(dir, "app", "invoke.auto"),
			Data:     data,
			Overlap:  gen.OverlapWrite,
		},
	}, nil
}
