package app

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/2637309949/dolphin/platform/util"

	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/fx/cli"
)

func init() {
	// BuildDomain.js
	cli.Invoke(InvokeEngine(func(engine *Engine) {
		engine.Group("/").Handle("GET", "/domain.js", func(ctx *Context) {
			reg := regexp.MustCompile("^([^:?]+)(:.*)?$")
			groups := reg.FindAllStringSubmatch(ctx.Request.Host, -1)
			domain := model.SysDomain{}
			ext, err := App.PlatformDB.Where("domain=? and del_flag=0 and status=1", groups[0][1]).Get(&domain)
			if err != nil || !ext {
				ctx.Data(200, "application/javascript", []byte(fmt.Sprintf("window.Domain={}")))
				return
			}
			ctx.Data(200, "application/javascript", []byte(fmt.Sprintf("window.Domain=%v", string(util.EnsureLeft(json.Marshal(model.Domain{
				Name:          domain.Name,
				FullName:      domain.FullName,
				ContactName:   domain.ContactName,
				ContactEmail:  domain.ContactEmail,
				ContactMobile: domain.ContactMobile,
				LoginUrl:      domain.LoginUrl,
				ApiUrl:        domain.ApiUrl,
				StaticUrl:     domain.StaticUrl,
				Theme:         domain.Theme,
				AuthMode:      domain.AuthMode,
			})).([]byte)))))
		})
	}))
}
