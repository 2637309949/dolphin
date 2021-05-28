// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"path"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	swag "github.com/2637309949/dolphin/packages/swag/gen"
	"github.com/spf13/viper"
)

// Doc struct
type Doc struct {
}

// Name defined pipe name
func (m *Doc) Name() string {
	return "doc"
}

// Pre defined
func (m *Doc) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (m *Doc) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (m *Doc) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	return []*pipe.TmplCfg{}, swag.New().Build(&swag.Config{
		SearchDir:          dir,
		MainAPIFile:        "main.go",
		PropNamingStrategy: "camelcase",
		MarkdownFilesDir:   "",
		OutputDir:          path.Join(dir, viper.GetString("dir.doc")),
		ParseVendor:        true,
		ParseDependency:    true,
		// ParseDepth:         2,
	})
}
