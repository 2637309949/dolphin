// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	ht "html/template"
	"path"
	"path/filepath"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/template"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// OAuth struct
type OAuth struct {
}

// Name defined pipe name
func (oa *OAuth) Name() string {
	return "oauth"
}

// Pre defined
func (oa *OAuth) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (oa *OAuth) After(*parser.AppParser, []*pipe.TmplCfg) error {
	return nil
}

// Build func
func (oa *OAuth) Build(dir string, args []string, parser *parser.AppParser) ([]*pipe.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": parser.PackageName,
		"Name":        parser.Name,
		"Controllers": parser.Controllers,
		"Services":    parser.Services,
		"Tables":      parser.Tables,
		"Beans":       parser.Beans,
		"Viper":       viper.GetViper(),
		"lt":          ht.HTML("<"),
		"gt":          ht.HTML(">"),
	}

	affirm, login := strings.Join(strings.Split(viper.GetString("oauth.affirm"), "/"), "/"), strings.Join(strings.Split(viper.GetString("oauth.login"), "/"), "/")
	affirmPath, loginPath := affirm[0:len(affirm)-len(filepath.Ext(affirm))], login[0:len(login)-len(filepath.Ext(login))]

	affirmByte, _ := vfsutil.ReadFile(template.Assets, "static/web/affirm.html")
	loginByte, _ := vfsutil.ReadFile(template.Assets, "static/web/login.html")
	return []*pipe.TmplCfg{
		{
			Text:     string(affirmByte),
			FilePath: path.Join(dir, affirmPath+".html"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		},
		{
			Text:     string(loginByte),
			FilePath: path.Join(dir, loginPath+".html"),
			Data:     data,
			Overlap:  pipe.OverlapSkip,
		},
	}, nil
}
