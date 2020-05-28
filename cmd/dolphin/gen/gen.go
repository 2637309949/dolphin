// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"path"
	"reflect"

	"github.com/2637309949/dolphin/cmd/dolphin/gen/modules"
	"github.com/2637309949/dolphin/cmd/dolphin/gen/pipe"
	"github.com/2637309949/dolphin/cmd/dolphin/schema"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrus"
)

var lines = []pipe.Pipe{
	&modules.Main{},
	&modules.App{},
	&modules.Ctr{},
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
}

// AddPipe defined addPipe
func AddPipe(p pipe.Pipe) {
	lines = append(lines, p)
}

// GetPipeByName defined getbyname
func GetPipeByName(nm string) pipe.Pipe {
	return funk.Find(lines, func(pipe pipe.Pipe) bool {
		return pipe.Name() == nm
	}).(pipe.Pipe)
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

// BuildWithDir defined
func (gen *Gen) BuildWithDir(dir string) error {
	for _, pipe := range gen.Pipes {
		cfgs, err := pipe.Build(dir, gen.App)
		if err != nil {
			return err
		}
		for _, cfg := range cfgs {
			err = gen.BuildWithCfg(cfg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// BuildWithCfg defined
func (gen *Gen) BuildWithCfg(cfg *pipe.TmplCfg) error {
	var err error
	var tpl *template.Template

	tpl = template.New("template")
	tpl.Funcs(template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
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

	w, err := OpenFile(cfg.FilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	if err = tpl.Execute(w, cfg.Data); err != nil {
		return err
	}
	if cfg.GOFmt {
		cmd := exec.Command("gofmt", "-s", "-w", cfg.FilePath)
		if err := cmd.Run(); err != nil && err != exec.ErrNotFound {
			logrus.Fatal(err)
		}
	}
	return nil
}

// OpenFile defiend
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	if err := os.MkdirAll(path.Dir(name), os.ModePerm); err != nil {
		return nil, err
	}
	return os.OpenFile(name, flag, perm)
}
