// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"errors"
	ht "html/template"
	"log"
	"os"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
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
	if len(args) < 1 {
		return cfgs, errors.New("please provide the project name")
	}
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return cfgs, errors.New("the project folder already exists")
	}

	if err := vfsutil.Walk(template.Assets, "/boilerplate", func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			log.Printf("can't stat file %s: %v\n", p, err)
			return nil
		}
		if fi.IsDir() {
			return nil
		}
		b := utils.EnsureLeft(vfsutil.ReadFile(template.Assets, p)).([]byte)
		filePath := strings.ReplaceAll(p, "/boilerplate/", "")
		cfgs = append(cfgs, &pipe.TmplCfg{
			Text:     string(b),
			FilePath: path.Join(dir, filePath),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
			GOFmt:    false,
		})
		return nil
	}); err != nil {
		return nil, err
	}
	return cfgs, nil
}
