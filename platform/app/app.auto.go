// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package app

import (
	"github.com/2637309949/dolphin/platform/model"

	"github.com/2637309949/dolphin/packages/fx/cli"
	"github.com/2637309949/dolphin/packages/viper"
)

// Name project
var Name = "platform"

// SysAppFun defined
type SysAppFun struct {
	Add,
	Del,
	Update,
	Page,
	Tree,
	Get func(ctx *Context)
}

// NewSysAppFun defined
func NewSysAppFun() *SysAppFun {
	ctr := &SysAppFun{}
	ctr.Add = SysAppFunAdd
	ctr.Del = SysAppFunDel
	ctr.Update = SysAppFunUpdate
	ctr.Page = SysAppFunPage
	ctr.Tree = SysAppFunTree
	ctr.Get = SysAppFunGet
	return ctr
}

// SysAppFunRoutes defined
func SysAppFunRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/app/fun/add", engine.Auth(), SysAppFunInstance.Add)
	group.Handle("DELETE", "/sys/app/fun/del", engine.Auth(), SysAppFunInstance.Del)
	group.Handle("PUT", "/sys/app/fun/update", engine.Auth(), SysAppFunInstance.Update)
	group.Handle("GET", "/sys/app/fun/page", engine.Auth(), SysAppFunInstance.Page)
	group.Handle("GET", "/sys/app/fun/tree", engine.Auth(), SysAppFunInstance.Tree)
	group.Handle("GET", "/sys/app/fun/get", engine.Auth(), SysAppFunInstance.Get)
}

// SysAppFunInstance defined
var SysAppFunInstance = NewSysAppFun()

// SysArea defined
type SysArea struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysArea defined
func NewSysArea() *SysArea {
	ctr := &SysArea{}
	ctr.Add = SysAreaAdd
	ctr.Del = SysAreaDel
	ctr.Update = SysAreaUpdate
	ctr.Page = SysAreaPage
	ctr.Get = SysAreaGet
	return ctr
}

// SysAreaRoutes defined
func SysAreaRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/area/add", engine.Auth(), SysAreaInstance.Add)
	group.Handle("DELETE", "/sys/area/del", engine.Auth(), SysAreaInstance.Del)
	group.Handle("PUT", "/sys/area/update", engine.Auth(), SysAreaInstance.Update)
	group.Handle("GET", "/sys/area/page", engine.Auth(), SysAreaInstance.Page)
	group.Handle("GET", "/sys/area/get", engine.Auth(), SysAreaInstance.Get)
}

// SysAreaInstance defined
var SysAreaInstance = NewSysArea()

// SysAttachment defined
type SysAttachment struct {
	Add,
	Upload,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysAttachment defined
func NewSysAttachment() *SysAttachment {
	ctr := &SysAttachment{}
	ctr.Add = SysAttachmentAdd
	ctr.Upload = SysAttachmentUpload
	ctr.Del = SysAttachmentDel
	ctr.Update = SysAttachmentUpdate
	ctr.Page = SysAttachmentPage
	ctr.Get = SysAttachmentGet
	return ctr
}

// SysAttachmentRoutes defined
func SysAttachmentRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/attachment/add", engine.Auth(), SysAttachmentInstance.Add)
	group.Handle("POST", "/sys/attachment/upload", engine.Auth(), SysAttachmentInstance.Upload)
	group.Handle("DELETE", "/sys/attachment/del", engine.Auth(), SysAttachmentInstance.Del)
	group.Handle("PUT", "/sys/attachment/update", engine.Auth(), SysAttachmentInstance.Update)
	group.Handle("GET", "/sys/attachment/page", engine.Auth(), SysAttachmentInstance.Page)
	group.Handle("GET", "/sys/attachment/get", engine.Auth(), SysAttachmentInstance.Get)
}

// SysAttachmentInstance defined
var SysAttachmentInstance = NewSysAttachment()

// SysCas defined
type SysCas struct {
	Login,
	Logout,
	Affirm,
	Authorize,
	Token,
	URL,
	Oauth2,
	Refresh,
	Check,
	Profile,
	Qrcode func(ctx *Context)
}

// NewSysCas defined
func NewSysCas() *SysCas {
	ctr := &SysCas{}
	ctr.Login = SysCasLogin
	ctr.Logout = SysCasLogout
	ctr.Affirm = SysCasAffirm
	ctr.Authorize = SysCasAuthorize
	ctr.Token = SysCasToken
	ctr.URL = SysCasURL
	ctr.Oauth2 = SysCasOauth2
	ctr.Refresh = SysCasRefresh
	ctr.Check = SysCasCheck
	ctr.Profile = SysCasProfile
	ctr.Qrcode = SysCasQrcode
	return ctr
}

// SysCasRoutes defined
func SysCasRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/cas/login", SysCasInstance.Login)
	group.Handle("GET", "/sys/cas/logout", SysCasInstance.Logout)
	group.Handle("POST", "/sys/cas/affirm", SysCasInstance.Affirm)
	group.Handle("GET", "/sys/cas/authorize", SysCasInstance.Authorize)
	group.Handle("POST", "/sys/cas/token", SysCasInstance.Token)
	group.Handle("GET", "/sys/cas/url", SysCasInstance.URL)
	group.Handle("GET", "/sys/cas/oauth2", SysCasInstance.Oauth2)
	group.Handle("GET", "/sys/cas/refresh", SysCasInstance.Refresh)
	group.Handle("GET", "/sys/cas/check", SysCasInstance.Check)
	group.Handle("GET", "/sys/cas/profile", engine.Auth(), SysCasInstance.Profile)
	group.Handle("GET", "/sys/cas/qrcode", SysCasInstance.Qrcode)
}

// SysCasInstance defined
var SysCasInstance = NewSysCas()

// SysClient defined
type SysClient struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysClient defined
func NewSysClient() *SysClient {
	ctr := &SysClient{}
	ctr.Add = SysClientAdd
	ctr.Del = SysClientDel
	ctr.Update = SysClientUpdate
	ctr.Page = SysClientPage
	ctr.Get = SysClientGet
	return ctr
}

// SysClientRoutes defined
func SysClientRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/client/add", engine.Auth(), SysClientInstance.Add)
	group.Handle("DELETE", "/sys/client/del", engine.Auth(), SysClientInstance.Del)
	group.Handle("PUT", "/sys/client/update", engine.Auth(), SysClientInstance.Update)
	group.Handle("GET", "/sys/client/page", engine.Auth(), SysClientInstance.Page)
	group.Handle("GET", "/sys/client/get", engine.Auth(), SysClientInstance.Get)
}

// SysClientInstance defined
var SysClientInstance = NewSysClient()

// SysDataPermission defined
type SysDataPermission struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysDataPermission defined
func NewSysDataPermission() *SysDataPermission {
	ctr := &SysDataPermission{}
	ctr.Add = SysDataPermissionAdd
	ctr.Del = SysDataPermissionDel
	ctr.Update = SysDataPermissionUpdate
	ctr.Page = SysDataPermissionPage
	ctr.Get = SysDataPermissionGet
	return ctr
}

// SysDataPermissionRoutes defined
func SysDataPermissionRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/data/permission/add", engine.Auth(), SysDataPermissionInstance.Add)
	group.Handle("DELETE", "/sys/data/permission/del", engine.Auth(), SysDataPermissionInstance.Del)
	group.Handle("PUT", "/sys/data/permission/update", engine.Auth(), SysDataPermissionInstance.Update)
	group.Handle("GET", "/sys/data/permission/page", engine.Auth(), SysDataPermissionInstance.Page)
	group.Handle("GET", "/sys/data/permission/get", engine.Auth(), SysDataPermissionInstance.Get)
}

// SysDataPermissionInstance defined
var SysDataPermissionInstance = NewSysDataPermission()

// SysDingtalk defined
type SysDingtalk struct {
	Oauth2 func(ctx *Context)
}

// NewSysDingtalk defined
func NewSysDingtalk() *SysDingtalk {
	ctr := &SysDingtalk{}
	ctr.Oauth2 = SysDingtalkOauth2
	return ctr
}

// SysDingtalkRoutes defined
func SysDingtalkRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("GET", "/sys/dingtalk/oauth2", SysDingtalkInstance.Oauth2)
}

// SysDingtalkInstance defined
var SysDingtalkInstance = NewSysDingtalk()

// SysDomain defined
type SysDomain struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysDomain defined
func NewSysDomain() *SysDomain {
	ctr := &SysDomain{}
	ctr.Add = SysDomainAdd
	ctr.Del = SysDomainDel
	ctr.Update = SysDomainUpdate
	ctr.Page = SysDomainPage
	ctr.Get = SysDomainGet
	return ctr
}

// SysDomainRoutes defined
func SysDomainRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/domain/add", engine.Auth(), SysDomainInstance.Add)
	group.Handle("DELETE", "/sys/domain/del", engine.Auth(), SysDomainInstance.Del)
	group.Handle("PUT", "/sys/domain/update", engine.Auth(), SysDomainInstance.Update)
	group.Handle("GET", "/sys/domain/page", engine.Auth(), SysDomainInstance.Page)
	group.Handle("GET", "/sys/domain/get", engine.Auth(), SysDomainInstance.Get)
}

// SysDomainInstance defined
var SysDomainInstance = NewSysDomain()

// SysMenu defined
type SysMenu struct {
	Add,
	Del,
	BatchDel,
	Update,
	Sidebar,
	Page,
	Tree,
	Get func(ctx *Context)
}

// NewSysMenu defined
func NewSysMenu() *SysMenu {
	ctr := &SysMenu{}
	ctr.Add = SysMenuAdd
	ctr.Del = SysMenuDel
	ctr.BatchDel = SysMenuBatchDel
	ctr.Update = SysMenuUpdate
	ctr.Sidebar = SysMenuSidebar
	ctr.Page = SysMenuPage
	ctr.Tree = SysMenuTree
	ctr.Get = SysMenuGet
	return ctr
}

// SysMenuRoutes defined
func SysMenuRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/menu/add", engine.Auth(), SysMenuInstance.Add)
	group.Handle("DELETE", "/sys/menu/del", engine.Auth(), SysMenuInstance.Del)
	group.Handle("DELETE", "/sys/menu/batch_del", engine.Auth(), SysMenuInstance.BatchDel)
	group.Handle("PUT", "/sys/menu/update", engine.Auth(), SysMenuInstance.Update)
	group.Handle("GET", "/sys/menu/sidebar", engine.Auth(), SysMenuInstance.Sidebar)
	group.Handle("GET", "/sys/menu/page", engine.Auth(), SysMenuInstance.Page)
	group.Handle("GET", "/sys/menu/tree", engine.Auth(), SysMenuInstance.Tree)
	group.Handle("GET", "/sys/menu/get", engine.Auth(), SysMenuInstance.Get)
}

// SysMenuInstance defined
var SysMenuInstance = NewSysMenu()

// SysOptionset defined
type SysOptionset struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysOptionset defined
func NewSysOptionset() *SysOptionset {
	ctr := &SysOptionset{}
	ctr.Add = SysOptionsetAdd
	ctr.Del = SysOptionsetDel
	ctr.Update = SysOptionsetUpdate
	ctr.Page = SysOptionsetPage
	ctr.Get = SysOptionsetGet
	return ctr
}

// SysOptionsetRoutes defined
func SysOptionsetRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/optionset/add", engine.Auth(), SysOptionsetInstance.Add)
	group.Handle("DELETE", "/sys/optionset/del", engine.Auth(), SysOptionsetInstance.Del)
	group.Handle("PUT", "/sys/optionset/update", engine.Auth(), SysOptionsetInstance.Update)
	group.Handle("GET", "/sys/optionset/page", engine.Auth(), SysOptionsetInstance.Page)
	group.Handle("GET", "/sys/optionset/get", engine.Auth(), SysOptionsetInstance.Get)
}

// SysOptionsetInstance defined
var SysOptionsetInstance = NewSysOptionset()

// SysOrg defined
type SysOrg struct {
	Add,
	Del,
	BatchDel,
	Update,
	Page,
	Tree,
	Get func(ctx *Context)
}

// NewSysOrg defined
func NewSysOrg() *SysOrg {
	ctr := &SysOrg{}
	ctr.Add = SysOrgAdd
	ctr.Del = SysOrgDel
	ctr.BatchDel = SysOrgBatchDel
	ctr.Update = SysOrgUpdate
	ctr.Page = SysOrgPage
	ctr.Tree = SysOrgTree
	ctr.Get = SysOrgGet
	return ctr
}

// SysOrgRoutes defined
func SysOrgRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/org/add", engine.Auth(), SysOrgInstance.Add)
	group.Handle("DELETE", "/sys/org/del", engine.Auth(), SysOrgInstance.Del)
	group.Handle("DELETE", "/sys/org/batch_del", engine.Auth(), SysOrgInstance.BatchDel)
	group.Handle("PUT", "/sys/org/update", engine.Auth(), SysOrgInstance.Update)
	group.Handle("GET", "/sys/org/page", engine.Auth(), SysOrgInstance.Page)
	group.Handle("GET", "/sys/org/tree", engine.Auth(), SysOrgInstance.Tree)
	group.Handle("GET", "/sys/org/get", engine.Auth(), SysOrgInstance.Get)
}

// SysOrgInstance defined
var SysOrgInstance = NewSysOrg()

// SysPermission defined
type SysPermission struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysPermission defined
func NewSysPermission() *SysPermission {
	ctr := &SysPermission{}
	ctr.Add = SysPermissionAdd
	ctr.Del = SysPermissionDel
	ctr.Update = SysPermissionUpdate
	ctr.Page = SysPermissionPage
	ctr.Get = SysPermissionGet
	return ctr
}

// SysPermissionRoutes defined
func SysPermissionRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/permission/add", engine.Auth(), SysPermissionInstance.Add)
	group.Handle("DELETE", "/sys/permission/del", engine.Auth(), SysPermissionInstance.Del)
	group.Handle("PUT", "/sys/permission/update", engine.Auth(), SysPermissionInstance.Update)
	group.Handle("GET", "/sys/permission/page", engine.Auth(), SysPermissionInstance.Page)
	group.Handle("GET", "/sys/permission/get", engine.Auth(), SysPermissionInstance.Get)
}

// SysPermissionInstance defined
var SysPermissionInstance = NewSysPermission()

// SysRole defined
type SysRole struct {
	Add,
	Del,
	Update,
	Page,
	RoleMenuTree,
	RoleAppFunTree,
	Get func(ctx *Context)
}

// NewSysRole defined
func NewSysRole() *SysRole {
	ctr := &SysRole{}
	ctr.Add = SysRoleAdd
	ctr.Del = SysRoleDel
	ctr.Update = SysRoleUpdate
	ctr.Page = SysRolePage
	ctr.RoleMenuTree = SysRoleRoleMenuTree
	ctr.RoleAppFunTree = SysRoleRoleAppFunTree
	ctr.Get = SysRoleGet
	return ctr
}

// SysRoleRoutes defined
func SysRoleRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/role/add", engine.Auth(), SysRoleInstance.Add)
	group.Handle("DELETE", "/sys/role/del", engine.Auth(), SysRoleInstance.Del)
	group.Handle("PUT", "/sys/role/update", engine.Auth(), SysRoleInstance.Update)
	group.Handle("GET", "/sys/role/page", engine.Auth(), SysRoleInstance.Page)
	group.Handle("GET", "/sys/role/role_menu_tree", engine.Auth(), SysRoleInstance.RoleMenuTree)
	group.Handle("GET", "/sys/role/role_app_fun_tree", engine.Auth(), SysRoleInstance.RoleAppFunTree)
	group.Handle("GET", "/sys/role/get", engine.Auth(), SysRoleInstance.Get)
}

// SysRoleInstance defined
var SysRoleInstance = NewSysRole()

// SysScheduling defined
type SysScheduling struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysScheduling defined
func NewSysScheduling() *SysScheduling {
	ctr := &SysScheduling{}
	ctr.Add = SysSchedulingAdd
	ctr.Del = SysSchedulingDel
	ctr.Update = SysSchedulingUpdate
	ctr.Page = SysSchedulingPage
	ctr.Get = SysSchedulingGet
	return ctr
}

// SysSchedulingRoutes defined
func SysSchedulingRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/scheduling/add", engine.Auth(), SysSchedulingInstance.Add)
	group.Handle("DELETE", "/sys/scheduling/del", engine.Auth(), SysSchedulingInstance.Del)
	group.Handle("PUT", "/sys/scheduling/update", engine.Auth(), SysSchedulingInstance.Update)
	group.Handle("GET", "/sys/scheduling/page", engine.Auth(), SysSchedulingInstance.Page)
	group.Handle("GET", "/sys/scheduling/get", engine.Auth(), SysSchedulingInstance.Get)
}

// SysSchedulingInstance defined
var SysSchedulingInstance = NewSysScheduling()

// SysTag defined
type SysTag struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysTag defined
func NewSysTag() *SysTag {
	ctr := &SysTag{}
	ctr.Add = SysTagAdd
	ctr.Del = SysTagDel
	ctr.Update = SysTagUpdate
	ctr.Page = SysTagPage
	ctr.Get = SysTagGet
	return ctr
}

// SysTagRoutes defined
func SysTagRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/tag/add", engine.Auth(), SysTagInstance.Add)
	group.Handle("DELETE", "/sys/tag/del", engine.Auth(), SysTagInstance.Del)
	group.Handle("PUT", "/sys/tag/update", engine.Auth(), SysTagInstance.Update)
	group.Handle("GET", "/sys/tag/page", engine.Auth(), SysTagInstance.Page)
	group.Handle("GET", "/sys/tag/get", engine.Auth(), SysTagInstance.Get)
}

// SysTagInstance defined
var SysTagInstance = NewSysTag()

// SysTagGroup defined
type SysTagGroup struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysTagGroup defined
func NewSysTagGroup() *SysTagGroup {
	ctr := &SysTagGroup{}
	ctr.Add = SysTagGroupAdd
	ctr.Del = SysTagGroupDel
	ctr.Update = SysTagGroupUpdate
	ctr.Page = SysTagGroupPage
	ctr.Get = SysTagGroupGet
	return ctr
}

// SysTagGroupRoutes defined
func SysTagGroupRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/tag/group/add", engine.Auth(), SysTagGroupInstance.Add)
	group.Handle("DELETE", "/sys/tag/group/del", engine.Auth(), SysTagGroupInstance.Del)
	group.Handle("PUT", "/sys/tag/group/update", engine.Auth(), SysTagGroupInstance.Update)
	group.Handle("GET", "/sys/tag/group/page", engine.Auth(), SysTagGroupInstance.Page)
	group.Handle("GET", "/sys/tag/group/get", engine.Auth(), SysTagGroupInstance.Get)
}

// SysTagGroupInstance defined
var SysTagGroupInstance = NewSysTagGroup()

// SysTracker defined
type SysTracker struct {
	Page,
	Get func(ctx *Context)
}

// NewSysTracker defined
func NewSysTracker() *SysTracker {
	ctr := &SysTracker{}
	ctr.Page = SysTrackerPage
	ctr.Get = SysTrackerGet
	return ctr
}

// SysTrackerRoutes defined
func SysTrackerRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("GET", "/sys/tracker/page", engine.Auth(), SysTrackerInstance.Page)
	group.Handle("GET", "/sys/tracker/get", engine.Auth(), SysTrackerInstance.Get)
}

// SysTrackerInstance defined
var SysTrackerInstance = NewSysTracker()

// SysUser defined
type SysUser struct {
	Add,
	Del,
	Update,
	Page,
	Get,
	Login,
	Logout func(ctx *Context)
}

// NewSysUser defined
func NewSysUser() *SysUser {
	ctr := &SysUser{}
	ctr.Add = SysUserAdd
	ctr.Del = SysUserDel
	ctr.Update = SysUserUpdate
	ctr.Page = SysUserPage
	ctr.Get = SysUserGet
	ctr.Login = SysUserLogin
	ctr.Logout = SysUserLogout
	return ctr
}

// SysUserRoutes defined
func SysUserRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/user/add", engine.Auth(), SysUserInstance.Add)
	group.Handle("DELETE", "/sys/user/del", engine.Auth(), SysUserInstance.Del)
	group.Handle("PUT", "/sys/user/update", engine.Auth(), SysUserInstance.Update)
	group.Handle("GET", "/sys/user/page", engine.Auth(), SysUserInstance.Page)
	group.Handle("GET", "/sys/user/get", engine.Auth(), SysUserInstance.Get)
	group.Handle("POST", "/sys/user/login", SysUserInstance.Login)
	group.Handle("GET", "/sys/user/logout", engine.Auth(), SysUserInstance.Logout)
}

// SysUserInstance defined
var SysUserInstance = NewSysUser()

// SysUserTemplate defined
type SysUserTemplate struct {
	Add,
	Del,
	Update,
	Page,
	Get func(ctx *Context)
}

// NewSysUserTemplate defined
func NewSysUserTemplate() *SysUserTemplate {
	ctr := &SysUserTemplate{}
	ctr.Add = SysUserTemplateAdd
	ctr.Del = SysUserTemplateDel
	ctr.Update = SysUserTemplateUpdate
	ctr.Page = SysUserTemplatePage
	ctr.Get = SysUserTemplateGet
	return ctr
}

// SysUserTemplateRoutes defined
func SysUserTemplateRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("POST", "/sys/user/template/add", engine.Auth(), SysUserTemplateInstance.Add)
	group.Handle("DELETE", "/sys/user/template/del", engine.Auth(), SysUserTemplateInstance.Del)
	group.Handle("PUT", "/sys/user/template/update", engine.Auth(), SysUserTemplateInstance.Update)
	group.Handle("GET", "/sys/user/template/page", engine.Auth(), SysUserTemplateInstance.Page)
	group.Handle("GET", "/sys/user/template/get", engine.Auth(), SysUserTemplateInstance.Get)
}

// SysUserTemplateInstance defined
var SysUserTemplateInstance = NewSysUserTemplate()

// SysWechat defined
type SysWechat struct {
	Oauth2 func(ctx *Context)
}

// NewSysWechat defined
func NewSysWechat() *SysWechat {
	ctr := &SysWechat{}
	ctr.Oauth2 = SysWechatOauth2
	return ctr
}

// SysWechatRoutes defined
func SysWechatRoutes(engine *Engine) {
	group := engine.Group(viper.GetString("http.prefix"))
	group.Handle("GET", "/sys/wechat/oauth2", SysWechatInstance.Oauth2)
}

// SysWechatInstance defined
var SysWechatInstance = NewSysWechat()

func init() {
	// Sync models
	cli.Invoke(InvokeEngine(func(e *Engine) {
		e.Manager.GetMSet().Add(new(model.SysAppFun))
		e.Manager.GetMSet().Add(new(model.SysArea))
		e.Manager.GetMSet().Add(new(model.SysAreaTemplate))
		e.Manager.GetMSet().Add(new(model.SysAreaTemplateDetail))
		e.Manager.GetMSet().Add(new(model.SysAttachment))
		e.Manager.GetMSet().Add(new(model.SysClient), "platform")
		e.Manager.GetMSet().Add(new(model.SysDataPermission))
		e.Manager.GetMSet().Add(new(model.SysDataPermissionDetail))
		e.Manager.GetMSet().Add(new(model.SysDomain), "platform")
		e.Manager.GetMSet().Add(new(model.SysMenu))
		e.Manager.GetMSet().Add(new(model.SysOptionset))
		e.Manager.GetMSet().Add(new(model.SysOrg))
		e.Manager.GetMSet().Add(new(model.SysPermission))
		e.Manager.GetMSet().Add(new(model.SysRole))
		e.Manager.GetMSet().Add(new(model.SysRoleAppFun))
		e.Manager.GetMSet().Add(new(model.SysRoleMenu))
		e.Manager.GetMSet().Add(new(model.SysRolePermission))
		e.Manager.GetMSet().Add(new(model.SysRoleUser))
		e.Manager.GetMSet().Add(new(model.SysTableColUser))
		e.Manager.GetMSet().Add(new(model.SysTag))
		e.Manager.GetMSet().Add(new(model.SysTagGroup))
		e.Manager.GetMSet().Add(new(model.SysTracker), "platform")
		e.Manager.GetMSet().Add(new(model.SysUser), "platform")
		e.Manager.GetMSet().Add(new(model.SysUserBinding))
		e.Manager.GetMSet().Add(new(model.SysUserTag))
		e.Manager.GetMSet().Add(new(model.SysUserTemplate))
		e.Manager.GetMSet().Add(new(model.SysUserTemplateDetail))
	}))
	// Async Ctr
	cli.Invoke(InvokeEngine(func(engine *Engine) {
		SysAppFunRoutes(engine)
		SysAreaRoutes(engine)
		SysAttachmentRoutes(engine)
		SysCasRoutes(engine)
		SysClientRoutes(engine)
		SysDataPermissionRoutes(engine)
		SysDingtalkRoutes(engine)
		SysDomainRoutes(engine)
		SysMenuRoutes(engine)
		SysOptionsetRoutes(engine)
		SysOrgRoutes(engine)
		SysPermissionRoutes(engine)
		SysRoleRoutes(engine)
		SysSchedulingRoutes(engine)
		SysTagRoutes(engine)
		SysTagGroupRoutes(engine)
		SysTrackerRoutes(engine)
		SysUserRoutes(engine)
		SysUserTemplateRoutes(engine)
		SysWechatRoutes(engine)
	}))
	// Booting system
	cli.Invoke(InvokeEngine(func(e *Engine) {
		if viper.GetString("app.name") == Name {
			e.Run()
		}
	}))
}
