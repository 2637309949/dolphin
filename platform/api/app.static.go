// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util"
)

// Static defined
type Static struct {
}

// NewStatic defined
func NewStatic() *Static {
	ctr := &Static{}
	return ctr
}

// StaticRoutes defined
func StaticRoutes(dol *Dolphin) {
	group := dol.Group("/")
	group.Handle("GET", "/domain.js", DomainInstance.Domain)
}

// DomainFormat defined
func DomainFormat(ct string) []byte {
	return []byte(fmt.Sprintf("window.Domain=%v", ct))
}

// Domain defined
func (ctr *Static) Domain(ctx *Context) {
	reg := regexp.MustCompile("^([^:?]+)(:.*)?$")
	contentType := "application/javascript"
	groups := reg.FindAllStringSubmatch(ctx.Request.Host, -1)
	domain := types.SysDomain{}
	ext, err := App.PlatformDB.Where("domain_url=? and is_delete=0 and status=1", groups[0][1]).Get(&domain)
	if err != nil || !ext {
		ctx.Data(200, contentType, DomainFormat("{}"))
		return
	}
	ctx.Data(200, contentType, DomainFormat(string(util.EnsureLeft(json.Marshal(types.Domain{
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
	})).([]byte))))
}

// DomainInstance defined TODO
var DomainInstance = NewStatic()
