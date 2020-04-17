package pipe

import "github.com/2637309949/dolphin/client/schema"

type (
	// Overlap int
	Overlap int
	// TmplCfg struct
	TmplCfg struct {
		Text     string
		FilePath string
		Data     interface{}
		Suffix   string
		Overlap  Overlap
	}
	// Pipe interface
	Pipe interface {
		Name() string
		Build(string, *schema.Application) ([]*TmplCfg, error)
	}
)

// OverlapSkip Overlap
const (
	OverlapSkip Overlap = iota + 1
	OverlapWrite
	OverlapInc
)
