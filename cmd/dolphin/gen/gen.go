// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gen

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

	"github.com/2637309949/dolphin/cmd/dolphin/gen/modules"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/cmd/dolphin/utils"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrus"
)

var lines = []pipe.Pipe{
	&modules.Main{},
	&modules.App{},
	&modules.Ctr{},
	&modules.Proto{},
	&modules.Srv{},
	&modules.Model{},
	&modules.Bean{},
	&modules.Auto{},
	&modules.Tool{},
	&modules.SQL{},
	&modules.SQLMap{},
	&modules.OAuth{},
	&modules.Script{},
	&modules.Doc{},
	&modules.SQLTPL{},
	&modules.Boilerplate{},
	&modules.Table{},
	&modules.Deploy{},
}

// AddPipe defined addPipe
func AddPipe(p pipe.Pipe) {
	lines = append(lines, p)
}

// GetPipesByName defined getbyname
func GetPipesByName(nm ...string) []pipe.Pipe {
	return funk.Filter(lines, func(pipe pipe.Pipe) bool {
		for _, v := range nm {
			if pipe.Name() == v {
				return true
			}
		}
		return false
	}).([]pipe.Pipe)
}

// Gen struct
type Gen struct {
	App   *schema.Application
	Pipes []pipe.Pipe
}

// New defined gen
func New(app *schema.Application) *Gen {
	return &Gen{App: app}
}

// AddPipe add pipe
func (gen *Gen) AddPipe(modules ...pipe.Pipe) {
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
	cfgs := []*pipe.TmplCfg{}
	for i := range gen.Pipes {
		items, err := gen.Pipes[i].Build(dir, args, gen.App)
		cfgs = append(cfgs, items...)
		if err != nil {
			return err
		}
		for j := range items {
			err = gen.Build(cfgs[j])
			if err != nil {
				panic(err)
			}
		}
	}
	// fmt code
	filePaths := funk.Keys(funk.Map(funk.Filter(cfgs, func(cfg *pipe.TmplCfg) bool { return cfg.GOFmt && path.Ext(cfg.FilePath) == ".go" }), func(cfg *pipe.TmplCfg) (string, string) { return path.Dir(cfg.FilePath), path.Dir(cfg.FilePath) })).([]string)
	if len(filePaths) > 0 {
		if err := utils.InstallPackages("golang.org/x/tools/cmd/goimports"); err != nil {
			logrus.Error(err)
		}
		for i := range filePaths {
			cmd := exec.Command("goimports", "-w", filePaths[i])
			if err := cmd.Run(); err != nil && err != exec.ErrNotFound {
				logrus.Error(err)
			}
		}
	}
	// rpc code
	filePaths = funk.Keys(funk.Map(funk.Filter(cfgs, func(cfg *pipe.TmplCfg) bool { return cfg.GOProto && path.Ext(cfg.FilePath) == ".proto" }), func(cfg *pipe.TmplCfg) (string, string) { return cfg.FilePath, cfg.FilePath })).([]string)
	if len(filePaths) > 0 {
		if err := utils.InstallPackages("github.com/golang/protobuf/proto", "github.com/golang/protobuf/protoc-gen-go"); err != nil {
			logrus.Error(err)
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
func (gen *Gen) Build(cfg *pipe.TmplCfg) error {
	var err error
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
	if _, err = os.Stat(cfg.FilePath); err == nil {
		if cfg.Overlap == pipe.OverlapInc {
			cfg.FilePath = cfg.FilePath + ".new"
			logrus.Info(fmt.Sprintf("%s inc generate", cfg.FilePath))
		} else if cfg.Overlap == pipe.OverlapWrite {
			logrus.Warn(fmt.Sprintf("%s over generate", cfg.FilePath))
		} else if cfg.Overlap == pipe.OverlapSkip {
			logrus.Info(fmt.Sprintf("%s skip generate", cfg.FilePath))
			return nil
		}
	}
	var bf bytes.Buffer
	bfw := io.Writer(&bf)
	w, err := utils.OpenFile(cfg.FilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
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
