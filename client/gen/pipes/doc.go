// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pipes

import (
	"path"

	"github.com/2637309949/dolphin/packages/viper"

	"github.com/2637309949/dolphin/client/gen/pipe"
	"github.com/2637309949/dolphin/client/schema"
	swag "github.com/2637309949/dolphin/packages/swag/gen"
)

// Doc struct
type Doc struct {
}

// Name defined pipe name
func (m *Doc) Name() string {
	return "doc"
}

// Build func
func (m *Doc) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	return []*pipe.TmplCfg{}, swag.New().Build(&swag.Config{
		SearchDir:          dir,
		MainAPIFile:        "main.go",
		PropNamingStrategy: "camelcase",
		OutputDir:          path.Join(dir, viper.GetString("dir.doc")),
		ParseVendor:        true,
		ParseDependency:    true,
		MarkdownFilesDir:   "",
	})
}
