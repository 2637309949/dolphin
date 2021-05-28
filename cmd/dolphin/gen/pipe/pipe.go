package pipe

import "github.com/2637309949/dolphin/cmd/dolphin/parser"

type (
	// Overlap int
	Overlap int
	// TmplCfg struct
	TmplCfg struct {
		Text     string
		FilePath string
		Data     interface{}
		Overlap  Overlap
		GOFmt    bool
		GOProto  bool
		Format   func(string) string
	}
	// Pipe interface
	Pipe interface {
		Pre(*parser.AppParser) error
		After(*parser.AppParser, []*TmplCfg) error
		Name() string
		Build(string, []string, *parser.AppParser) ([]*TmplCfg, error)
	}
)

// OverlapSkip Overlap
const (
	OverlapSkip Overlap = iota + 1
	OverlapWrite
	OverlapInc
)
