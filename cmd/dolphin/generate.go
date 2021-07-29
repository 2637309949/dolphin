// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"os/exec"
	"path"
	"reflect"
	"strings"

	"github.com/2637309949/dolphin/cmd/dolphin/modules"
	"github.com/2637309949/dolphin/cmd/dolphin/parser"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

var lines = []parser.Pipe{
	&modules.More{},
	&modules.Dolphin{},
	&modules.Script{},
	&modules.Doc{},
	&modules.SQLTPL{},
	&modules.Boilerplate{},
	&modules.Reverse{},
	&modules.Deploy{},
}

// AddPipe defined addPipe
func AddPipe(p parser.Pipe) {
	lines = append(lines, p)
}

// GetPipesByName defined getbyname
func GetPipesByName(nm ...string) []parser.Pipe {
	return funk.Chain(lines).Filter(func(parser parser.Pipe) bool {
		for i := range nm {
			if parser.Name() == nm[i] {
				return true
			}
		}
		return false
	}).Value().([]parser.Pipe)
}

// Gen struct
type Gen struct {
	Parser *parser.AppParser
	Pipes  []parser.Pipe
}

// New defined gen
func New(parser *parser.AppParser) *Gen {
	return &Gen{Parser: parser}
}

// AddPipe add parser
func (gen *Gen) AddPipe(modules ...parser.Pipe) {
	gen.Pipes = append(gen.Pipes, modules...)
}

// BuildDir defined
func (gen *Gen) BuildDir(dir string, args []string) (err error) {
	defer func() {
		if rErr := recover(); rErr != nil {
			err = fmt.Errorf("%v", rErr)
		}
	}()

	// generate code
	cfgs := []*parser.TmplCfg{}
	for i := range gen.Pipes {
		err := gen.Pipes[i].Pre(gen.Parser)
		if err != nil {
			return err
		}
		items, err := gen.Pipes[i].Build(dir, args, gen.Parser)
		if err != nil {
			return err
		}

		cfgs = append(cfgs, items...)
		for j := range items {
			err = gen.Build(items[j])
			if err != nil {
				return err
			}
		}
		err = gen.Pipes[i].After(gen.Parser, items)
		if err != nil {
			return err
		}
	}

	// fmt code
	filePaths := funk.Keys(funk.Map(funk.Filter(cfgs, func(cfg *parser.TmplCfg) bool { return cfg.GOFmt && path.Ext(cfg.FilePath) == ".go" }), func(cfg *parser.TmplCfg) (string, string) { return path.Dir(cfg.FilePath), path.Dir(cfg.FilePath) })).([]string)
	if len(filePaths) > 0 {
		if !utils.HasBin("goimports") {
			if status := utils.NetWorkStatus(); status {
				if err := utils.InstallPackages("golang.org/x/tools/cmd/goimports"); err != nil {
					logrus.Error(err)
				}
			} else {
				utils.Ensure(errors.New("not goimports found"))
			}
		}
		for i := range filePaths {
			cmd := exec.Command("goimports", "-w", filePaths[i])
			if err := cmd.Run(); err != nil && err != exec.ErrNotFound {
				logrus.Error(err)
			}
		}
	}

	// rpc code
	filePaths = funk.Keys(funk.Map(funk.Filter(cfgs, func(cfg *parser.TmplCfg) bool { return cfg.GOProto && path.Ext(cfg.FilePath) == ".proto" }), func(cfg *parser.TmplCfg) (string, string) { return cfg.FilePath, cfg.FilePath })).([]string)
	if len(filePaths) > 0 {
		if !utils.HasBin("proto") {
			if status := utils.NetWorkStatus(); status {
				if err := utils.InstallPackages("github.com/golang/protobuf/proto", "github.com/golang/protobuf/protoc-gen-go"); err != nil {
					logrus.Error(err)
				}
			} else {
				utils.Ensure(errors.New("not proto found"))
			}
		}
		for i := range filePaths {
			cmd := exec.Command("protoc", "-I", path.Dir(filePaths[i]), filePaths[i], "--go_out=plugins=grpc:"+path.Dir(filePaths[i]))
			if err := cmd.Run(); err != nil && err != exec.ErrNotFound {
				logrus.Error(err)
			}
		}
	}
	return nil
}

// BuildWithCfg defined
func (gen *Gen) Build(cfg *parser.TmplCfg) error {
	var err error
	var genFilePath string
	var tpl *template.Template
	tpl = template.New("template")
	tpl.Funcs(template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
		"safeHTML": func(b string) template.URL {
			return template.URL(b)
		},
	})
	if tpl, err = tpl.Parse(cfg.Text); err != nil {
		return err
	}

	genFilePath = cfg.FilePath
	if _, err = os.Stat(genFilePath); !os.IsNotExist(err) {
		if cfg.Overlap == parser.OverlapInc {
			genFilePath = genFilePath + ".new"
			logrus.Info(fmt.Sprintf("%s inc generate", genFilePath))
		} else if cfg.Overlap == parser.OverlapWrite {
			logrus.Warn(fmt.Sprintf("%s over generate", genFilePath))
		} else if cfg.Overlap == parser.OverlapSkip {
			logrus.Info(fmt.Sprintf("%s skip generate", genFilePath))
			return nil
		}
	}

	var bf bytes.Buffer
	bfw := io.Writer(&bf)
	w, err := utils.OpenFile(genFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	if err = tpl.Execute(bfw, cfg.Data); err != nil {
		return err
	}
	sbt := bf.String()
	sbt = strings.ReplaceAll(sbt, "&#39;", "'")
	_, err = w.WriteString(sbt)
	if err != nil {
		return err
	}
	return nil
}
