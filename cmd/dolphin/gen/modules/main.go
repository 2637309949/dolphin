// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Main struct
type Main struct {
}

// Name defined pipe name
func (m *Main) Name() string {
	return "main"
}

// Build func
func (m *Main) Build(dir string, args []string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
		"Viper":       viper.GetViper(),
	}
	mainByte, _ := vfsutil.ReadFile(template.Assets, "main.tmpl")
	return []*pipe.TmplCfg{
		{
			Text:     string(mainByte),
			FilePath: path.Join(dir, "main.go"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    true,
		},
	}, nil
}
