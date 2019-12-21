// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/cli/gen"
	"github.com/2637309949/dolphin/cli/schema"
	"github.com/2637309949/dolphin/cli/tempalte"
	"github.com/2637309949/dolphin/cli/packages/viper"
)

// OAuth struct
type OAuth struct {
}

// Build func
func (oa *OAuth) Build(dir string, node *schema.Application) ([]*gen.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
	}
	return []*gen.TmplCfg{
		&gen.TmplCfg{
			Text:     tempalte.TmplAuth,
			FilePath: path.Join(dir, viper.GetString("http.static"), "auth"),
			Data:     data,
			Overlap:  gen.OverlapSkip,
			Suffix:   ".html",
		},
		&gen.TmplCfg{
			Text:     tempalte.TmplLogin,
			FilePath: path.Join(dir, viper.GetString("http.static"), "login"),
			Data:     data,
			Overlap:  gen.OverlapSkip,
			Suffix:   ".html",
		},
	}, nil
}
