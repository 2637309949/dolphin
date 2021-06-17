// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	ht "html/template"
	"log"
	"os"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Boilerplate struct
type Boilerplate struct {
}

// Name defined pipe name
func (m *Boilerplate) Name() string {
	return "boilerplate"
}

// Pre defined
func (m *Boilerplate) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (m *Boilerplate) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (m *Boilerplate) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	cfgs := []*pipe.TmplCfg{}
	data := map[string]interface{}{
		"PackageName": parser.PackageName,
		"Name":        parser.Name,
		"Desc":        parser.Desc,
		"Controllers": parser.Controllers,
		"Services":    parser.Services,
		"Tables":      parser.Tables,
		"Beans":       parser.Beans,
		"Viper":       viper.GetViper(),
		"lt":          ht.HTML("<"),
		"gt":          ht.HTML(">"),
	}
	walkFn := func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Printf("can't stat file %s: %v\n", p, err)
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		b, err := vfsutil.ReadFile(template.Assets, p)
		if err != nil {
			return err
		}
		filePath := strings.ReplaceAll(p, "/boilerplate/", "")
		cfgs = append(cfgs, &pipe.TmplCfg{
			Text:     string(b),
			FilePath: path.Join(dir, filePath),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    false,
		})
		return nil
	}
	if err := vfsutil.Walk(template.Assets, "/boilerplate", walkFn); err != nil {
		return nil, err
	}
	return cfgs, nil
}
