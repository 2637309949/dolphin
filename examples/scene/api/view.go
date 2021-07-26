// Code generated by dol build. Only Generate by tools if not existed.
// source: view.go

package api

// ViewFile api implementation
// @Summary 文件下载
// @Tags 视图
// @Param Authorization header string false "认证令牌"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/view/file [get]
func (ctr *View) ViewFile(ctx *Context) {
	ctx.RenderFile("static/tmpl/view.tmpl", "view.txt", map[string]interface{}{"app": 100})
}

// ViewHTML api implementation
// @Summary HTML显示
// @Tags 视图
// @Param Authorization header string false "认证令牌"
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/view/html [get]
func (ctr *View) ViewHTML(ctx *Context) {
	ctx.RenderHTML("static/tmpl/view.tmpl", map[string]interface{}{"app": 100})
}

// ViewXML api implementation
// @Summary XML显示
// @Tags 视图
// @Produce application/json
// @Failure 403 {object} types.Fail
// @Success 200 {object} types.Success
// @Failure 500 {object} types.Fail
// @Router /api/view/xml [get]
func (ctr *View) ViewXML(ctx *Context) {
	ctx.RenderXML("static/tmpl/view.tmpl", map[string]interface{}{"app": 100})
}
