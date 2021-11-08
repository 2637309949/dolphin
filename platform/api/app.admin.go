package api

import "github.com/eriklott/mustache"

// Admin defined TODO
type Admin struct {
}

// NewAdmin defined TODO
func NewAdmin() *Admin {
	ctr := &Admin{}
	return ctr
}

// AdminRoutes defined TODO
func AdminRoutes(dol *Dolphin) {
	g, i := dol.Group("/"), AdminInstance
	g.Handle("GET", "/admin", i.Admin)
}

// Admin defined TODO
func (ctr *Admin) Admin(ctx *Context) {
	m := mustache.NewTemplate()
	m.Parse("admin", "table todo")
	str, err := m.Render("admin", "")
	if err != nil {
		ctx.Fail(err)
		return
	}
	ctx.Data(200, "text/html; charset=utf-8", []byte(str))
}

// AdminInstance defined TODO
var AdminInstance = NewAdmin()
