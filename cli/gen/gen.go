// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gen

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/sirupsen/logrus"

	"github.com/2637309949/dolphin/cli/schema"
)

type (
	// Overlap int
	Overlap int
	// TmplCfg struct
	TmplCfg struct {
		Text     string
		FilePath string
		Data     interface{}
		Overlap  Overlap
	}
	// Pipe interface
	Pipe interface {
		Build(string, *schema.Application) ([]*TmplCfg, error)
	}
	// Gen struct
	Gen struct {
		App   *schema.Application
		Pipes []Pipe
	}
)

// OverlapSkip Overlap
const (
	OverlapSkip Overlap = iota + 1
	OverlapWrite
	OverlapInc
)

// New defined gen
func New(app *schema.Application) *Gen {
	return &Gen{App: app}
}

// AddPipe add pipe
func (g *Gen) AddPipe(pipes ...Pipe) {
	g.Pipes = append(g.Pipes, pipes...)
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

func (g *Gen) buildTmpl(tmpcfg *TmplCfg) error {
	var suffix string
	var tmpl *template.Template
	var err error

	tmpl = template.New("buildTmpl")
	suffix = ".go"
	if tmpl, err = tmpl.Parse(tmpcfg.Text); err != nil {
		return err
	}
	tmpcfg.FilePath = tmpcfg.FilePath + suffix
	if _, err = os.Stat(tmpcfg.FilePath); err == nil {
		if tmpcfg.Overlap == OverlapInc {
			tmpcfg.FilePath = tmpcfg.FilePath + ".new"
			logrus.Warn(fmt.Sprintf("file:%s has exitsed, inc generate", tmpcfg.FilePath))
		} else if tmpcfg.Overlap == OverlapWrite {
			logrus.Warn(fmt.Sprintf("file:%s has exitsed, over generate", tmpcfg.FilePath))
		} else if tmpcfg.Overlap == OverlapSkip {
			logrus.Warn(fmt.Sprintf("file:%s has exitsed, skip generate", tmpcfg.FilePath))
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
	return nil
}
