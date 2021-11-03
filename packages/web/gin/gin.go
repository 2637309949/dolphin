package gin

import "github.com/2637309949/dolphin/packages/web/core"

func init() {
	core.SetEngine(&struct {
		*Context
		*handler
	}{NewContext(), NewHandle()})
}
