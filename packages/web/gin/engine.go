package gin

import "github.com/2637309949/dolphin/packages/web/core"

type engine struct {
	*Context
	*handler
}

func NewEngine() core.Engine {
	return &engine{
		Context: NewContext(),
		handler: NewHandle(),
	}
}
func init() {
	core.SetEngine(NewEngine())
}
