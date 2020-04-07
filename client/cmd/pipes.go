package cmd

import (
	"github.com/2637309949/dolphin/client/gen"
	"github.com/2637309949/dolphin/client/gen/pipes"
)

// GetPipeByName defined getbyname
func GetPipeByName(nm string) gen.Pipe {
	switch nm {
	case (&pipes.Main{}).Name():
		return &pipes.Main{}
	case (&pipes.App{}).Name():
		return &pipes.App{}
	case (&pipes.Ctr{}).Name():
		return &pipes.Ctr{}
	case (&pipes.Srv{}).Name():
		return &pipes.Srv{}
	case (&pipes.Model{}).Name():
		return &pipes.Model{}
	case (&pipes.Bean{}).Name():
		return &pipes.Bean{}
	case (&pipes.Auto{}).Name():
		return &pipes.Auto{}
	case (&pipes.Tool{}).Name():
		return &pipes.Tool{}
	case (&pipes.SQL{}).Name():
		return &pipes.SQL{}
	case (&pipes.SQLMap{}).Name():
		return &pipes.SQLMap{}
	case (&pipes.OAuth{}).Name():
		return &pipes.OAuth{}
	case (&pipes.Script{}).Name():
		return &pipes.Script{}
	case (&pipes.Doc{}).Name():
		return &pipes.Doc{}
	case (&pipes.SQLTPL{}).Name():
		return &pipes.SQLTPL{}
	}
	return nil
}
