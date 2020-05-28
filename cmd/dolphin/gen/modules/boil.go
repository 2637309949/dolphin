// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/shurcooL/httpfs/vfsutil"
)

// Boilerplate struct
type Boilerplate struct {
}

// Name defined pipe name
func (m *Boilerplate) Name() string {
	return "boilerplate"
}

// Build func
func (m *Boilerplate) Build(dir string, node *schema.Application) ([]*pipe.TmplCfg, error) {
	cfgs := []*pipe.TmplCfg{}
	data := map[string]interface{}{
		"PackageName": node.PackageName,
		"Name":        node.Name,
		"Desc":        node.Desc,
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
