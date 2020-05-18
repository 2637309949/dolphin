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
func (g *Gen) AddPipe(modules ...pipe.Pipe) {
	g.Pipes = append(g.Pipes, modules...)
}

// Build build code from xml
func (g *Gen) Build(dir string) error {
	for _, pipe := range g.Pipes {
		tmpcfgs, err := pipe.Build(dir, g.App)
		if err != nil {
			return err
		}
		for _, tmpcfg := range tmpcfgs {
			err = g.buildTmpl(tmpcfg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (g *Gen) buildTmpl(tmpcfg *pipe.TmplCfg) error {
	var suffix string
	var tmpl *template.Template
	var err error

	tmpl = template.New("buildTmpl")
	tmpl.Funcs(template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
	})
	suffix = tmpcfg.Suffix
	if tmpl, err = tmpl.Parse(tmpcfg.Text); err != nil {
		return err
	}
	tmpcfg.FilePath = tmpcfg.FilePath + suffix
	if _, err = os.Stat(tmpcfg.FilePath); err == nil {
		if tmpcfg.Overlap == pipe.OverlapInc {
			tmpcfg.FilePath = tmpcfg.FilePath + ".new"
			logrus.Info(fmt.Sprintf("%s inc generate", tmpcfg.FilePath))
		} else if tmpcfg.Overlap == pipe.OverlapWrite {
			logrus.Warn(fmt.Sprintf("%s over generate", tmpcfg.FilePath))
		} else if tmpcfg.Overlap == pipe.OverlapSkip {
			logrus.Info(fmt.Sprintf("%s skip generate", tmpcfg.FilePath))
			return nil
		}
	}

	if err = os.MkdirAll(path.Dir(tmpcfg.FilePath), os.ModePerm); err != nil {
		return err
	}
	w, err := os.OpenFile(tmpcfg.FilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	if err = tmpl.Execute(w, tmpcfg.Data); err != nil {
		return err
	}
	if tmpcfg.GOFmt {
		cmd := exec.Command("gofmt", "-s", "-w", tmpcfg.FilePath)
		if err := cmd.Run(); err != nil && err != exec.ErrNotFound {
			logrus.Fatal(err)
		}
	}
	return nil
}
