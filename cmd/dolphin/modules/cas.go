// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package modules

import (
	"html/template"
	"path"
	"path/filepath"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/template/dist"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/shurcooL/httpfs/vfsutil"
	"github.com/spf13/viper"
)

// Cas struct
type Cas struct {
}

// Name defined parser name
func (dp *Cas) Name() string {
	return "cas"
}

// Pre defined
func (dp *Cas) Pre(*parser.AppParser) error {
	return nil
}

// After defined
func (dp *Cas) After(*parser.AppParser, []*parser.TmplCfg) error {
	return nil
}

// Build func
func (dp *Cas) Build(dir string, args []string, appParser *parser.AppParser) ([]*parser.TmplCfg, error) {
	data := map[string]interface{}{
		"PackageName": appParser.PackageName,
		"Name":        appParser.Name,
		"Controllers": appParser.Controllers,
		"Services":    appParser.Services,
		"Tables":      appParser.Tables,
		"Beans":       appParser.Beans,
		"Viper":       viper.GetViper(),
		"lt":          template.HTML("<"),
		"gt":          template.HTML(">"),
	}
	cfgs := []*parser.TmplCfg{}
	affirm, login, qrconnect := strings.Join(strings.Split(viper.GetString("oauth.affirm"), "/"), "/"), strings.Join(strings.Split(viper.GetString("oauth.login"), "/"), "/"), strings.Join(strings.Split(viper.GetString("oauth.qrconnect"), "/"), "/")
	affirmPath, loginPath, qrconnectPath := affirm[0:len(affirm)-len(filepath.Ext(affirm))], login[0:len(login)-len(filepath.Ext(login))], qrconnect[0:len(qrconnect)-len(filepath.Ext(qrconnect))]
	affirmByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "static/web/affirm.html")).([]byte)
	qrconnectByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "static/web/qrconnect.html")).([]byte)
	loginByte := utils.EnsureLeft(vfsutil.ReadFile(dist.Assets, "static/web/login.html")).([]byte)
	cfgs = append(cfgs, &parser.TmplCfg{
		Text:     string(affirmByte),
		FilePath: path.Join(dir, affirmPath+".html"),
		Data:     data,
		Overlap:  parser.OverlapSkip,
	})
	cfgs = append(cfgs, &parser.TmplCfg{
		Text:     string(loginByte),
		FilePath: path.Join(dir, loginPath+".html"),
		Data:     data,
		Overlap:  parser.OverlapSkip,
	})
	cfgs = append(cfgs, &parser.TmplCfg{
		Text:     string(qrconnectByte),
		FilePath: path.Join(dir, qrconnectPath+".html"),
		Data:     data,
		Overlap:  parser.OverlapSkip,
	})
	return cfgs, nil
}
