// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package api

import (
	"scene/proto"
	"scene/rpc"
	"scene/srv"
	"scene/svc"
	"scene/types"
	"time"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// Name project
var Name = "scene"
var NopHandlerFunc = func(ctx *Context) { ctx.Next() }

// Controller defined
type Controller struct {
	Method       string
	RelativePath string
	Auth,
	Roles,
	Cache,
	Interceptor,
	Handler HandlerFunc
}

// Article defined
type Article struct {
	Name string
	Srv  *srv.Article
	Add,
	BatchAdd,
	Del,
	BatchDel,
	Update,
	BatchUpdate,
	Page,
	Get,
	Payment Controller
}

// NewArticle defined
func NewArticle() *Article {
	ctr := &Article{Name: "article", Srv: srv.NewArticle()}
	ctr.Add.Method = "POST"
	ctr.Add.RelativePath = "/article/add"
	ctr.Add.Auth = Auth("token")
	ctr.Add.Roles = NopHandlerFunc
	ctr.Add.Cache = NopHandlerFunc
	ctr.Add.Interceptor = NopHandlerFunc
	ctr.Add.Handler = ctr.ArticleAdd
	ctr.BatchAdd.Method = "POST"
	ctr.BatchAdd.RelativePath = "/article/batch_add"
	ctr.BatchAdd.Auth = Auth("token")
	ctr.BatchAdd.Roles = NopHandlerFunc
	ctr.BatchAdd.Cache = NopHandlerFunc
	ctr.BatchAdd.Interceptor = NopHandlerFunc
	ctr.BatchAdd.Handler = ctr.ArticleBatchAdd
	ctr.Del.Method = "DELETE"
	ctr.Del.RelativePath = "/article/del"
	ctr.Del.Auth = Auth("token")
	ctr.Del.Roles = NopHandlerFunc
	ctr.Del.Cache = NopHandlerFunc
	ctr.Del.Interceptor = NopHandlerFunc
	ctr.Del.Handler = ctr.ArticleDel
	ctr.BatchDel.Method = "PUT"
	ctr.BatchDel.RelativePath = "/article/batch_del"
	ctr.BatchDel.Auth = Auth("token")
	ctr.BatchDel.Roles = NopHandlerFunc
	ctr.BatchDel.Cache = NopHandlerFunc
	ctr.BatchDel.Interceptor = NopHandlerFunc
	ctr.BatchDel.Handler = ctr.ArticleBatchDel
	ctr.Update.Method = "PUT"
	ctr.Update.RelativePath = "/article/update"
	ctr.Update.Auth = Auth("token")
	ctr.Update.Roles = NopHandlerFunc
	ctr.Update.Cache = NopHandlerFunc
	ctr.Update.Interceptor = NopHandlerFunc
	ctr.Update.Handler = ctr.ArticleUpdate
	ctr.BatchUpdate.Method = "PUT"
	ctr.BatchUpdate.RelativePath = "/article/batch_update"
	ctr.BatchUpdate.Auth = Auth("token")
	ctr.BatchUpdate.Roles = NopHandlerFunc
	ctr.BatchUpdate.Cache = NopHandlerFunc
	ctr.BatchUpdate.Interceptor = NopHandlerFunc
	ctr.BatchUpdate.Handler = ctr.ArticleBatchUpdate
	ctr.Page.Method = "GET"
	ctr.Page.RelativePath = "/article/page"
	ctr.Page.Auth = Auth("token")
	ctr.Page.Roles = NopHandlerFunc
	ctr.Page.Cache = NopHandlerFunc
	ctr.Page.Interceptor = NopHandlerFunc
	ctr.Page.Handler = ctr.ArticlePage
	ctr.Get.Method = "GET"
	ctr.Get.RelativePath = "/article/get"
	ctr.Get.Auth = Auth("token")
	ctr.Get.Roles = NopHandlerFunc
	ctr.Get.Cache = NopHandlerFunc
	ctr.Get.Interceptor = NopHandlerFunc
	ctr.Get.Handler = ctr.ArticleGet
	ctr.Payment.Method = "POST"
	ctr.Payment.RelativePath = "/article/payment"
	ctr.Payment.Auth = Auth("token")
	ctr.Payment.Roles = NopHandlerFunc
	ctr.Payment.Cache = NopHandlerFunc
	ctr.Payment.Interceptor = NopHandlerFunc
	ctr.Payment.Handler = ctr.ArticlePayment
	return ctr
}

// ArticleRoutes defined
func ArticleRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), ArticleInstance
	group.Handle(instance.Add.Method, instance.Add.RelativePath, instance.Add.Auth, instance.Add.Roles, instance.Add.Cache, instance.Add.Interceptor, instance.Add.Handler)
	group.Handle(instance.BatchAdd.Method, instance.BatchAdd.RelativePath, instance.BatchAdd.Auth, instance.BatchAdd.Roles, instance.BatchAdd.Cache, instance.BatchAdd.Interceptor, instance.BatchAdd.Handler)
	group.Handle(instance.Del.Method, instance.Del.RelativePath, instance.Del.Auth, instance.Del.Roles, instance.Del.Cache, instance.Del.Interceptor, instance.Del.Handler)
	group.Handle(instance.BatchDel.Method, instance.BatchDel.RelativePath, instance.BatchDel.Auth, instance.BatchDel.Roles, instance.BatchDel.Cache, instance.BatchDel.Interceptor, instance.BatchDel.Handler)
	group.Handle(instance.Update.Method, instance.Update.RelativePath, instance.Update.Auth, instance.Update.Roles, instance.Update.Cache, instance.Update.Interceptor, instance.Update.Handler)
	group.Handle(instance.BatchUpdate.Method, instance.BatchUpdate.RelativePath, instance.BatchUpdate.Auth, instance.BatchUpdate.Roles, instance.BatchUpdate.Cache, instance.BatchUpdate.Interceptor, instance.BatchUpdate.Handler)
	group.Handle(instance.Page.Method, instance.Page.RelativePath, instance.Page.Auth, instance.Page.Roles, instance.Page.Cache, instance.Page.Interceptor, instance.Page.Handler)
	group.Handle(instance.Get.Method, instance.Get.RelativePath, instance.Get.Auth, instance.Get.Roles, instance.Get.Cache, instance.Get.Interceptor, instance.Get.Handler)
	group.Handle(instance.Payment.Method, instance.Payment.RelativePath, instance.Payment.Auth, instance.Payment.Roles, instance.Payment.Cache, instance.Payment.Interceptor, instance.Payment.Handler)
}

// ArticleInstance defined
var ArticleInstance = NewArticle()

// Caching defined
type Caching struct {
	Name string
	Srv  *srv.Caching
	Info Controller
}

// NewCaching defined
func NewCaching() *Caching {
	ctr := &Caching{Name: "caching", Srv: srv.NewCaching()}
	ctr.Info.Method = "GET"
	ctr.Info.RelativePath = "/caching/info"
	ctr.Info.Auth = NopHandlerFunc
	ctr.Info.Roles = NopHandlerFunc
	ctr.Info.Cache = Cache(5 * time.Second)
	ctr.Info.Interceptor = NopHandlerFunc
	ctr.Info.Handler = ctr.CachingInfo
	return ctr
}

// CachingRoutes defined
func CachingRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), CachingInstance
	group.Handle(instance.Info.Method, instance.Info.RelativePath, instance.Info.Auth, instance.Info.Roles, instance.Info.Cache, instance.Info.Interceptor, instance.Info.Handler)
}

// CachingInstance defined
var CachingInstance = NewCaching()

// Dtm defined
type Dtm struct {
	Name string
	Srv  *srv.Dtm
	Tcc,
	TransOut,
	TransOutConfirm,
	TransOutRevert,
	TransIn,
	TransInConfirm,
	TransInRevert Controller
}

// NewDtm defined
func NewDtm() *Dtm {
	ctr := &Dtm{Name: "dtm", Srv: srv.NewDtm()}
	ctr.Tcc.Method = "GET"
	ctr.Tcc.RelativePath = "/dtm/tcc"
	ctr.Tcc.Auth = NopHandlerFunc
	ctr.Tcc.Roles = NopHandlerFunc
	ctr.Tcc.Cache = NopHandlerFunc
	ctr.Tcc.Interceptor = NopHandlerFunc
	ctr.Tcc.Handler = ctr.DtmTcc
	ctr.TransOut.Method = "POST"
	ctr.TransOut.RelativePath = "/dtm/trans_out"
	ctr.TransOut.Auth = NopHandlerFunc
	ctr.TransOut.Roles = NopHandlerFunc
	ctr.TransOut.Cache = NopHandlerFunc
	ctr.TransOut.Interceptor = NopHandlerFunc
	ctr.TransOut.Handler = ctr.DtmTransOut
	ctr.TransOutConfirm.Method = "POST"
	ctr.TransOutConfirm.RelativePath = "/dtm/trans_out_confirm"
	ctr.TransOutConfirm.Auth = NopHandlerFunc
	ctr.TransOutConfirm.Roles = NopHandlerFunc
	ctr.TransOutConfirm.Cache = NopHandlerFunc
	ctr.TransOutConfirm.Interceptor = NopHandlerFunc
	ctr.TransOutConfirm.Handler = ctr.DtmTransOutConfirm
	ctr.TransOutRevert.Method = "POST"
	ctr.TransOutRevert.RelativePath = "/dtm/trans_out_revert"
	ctr.TransOutRevert.Auth = NopHandlerFunc
	ctr.TransOutRevert.Roles = NopHandlerFunc
	ctr.TransOutRevert.Cache = NopHandlerFunc
	ctr.TransOutRevert.Interceptor = NopHandlerFunc
	ctr.TransOutRevert.Handler = ctr.DtmTransOutRevert
	ctr.TransIn.Method = "POST"
	ctr.TransIn.RelativePath = "/dtm/trans_in"
	ctr.TransIn.Auth = NopHandlerFunc
	ctr.TransIn.Roles = NopHandlerFunc
	ctr.TransIn.Cache = NopHandlerFunc
	ctr.TransIn.Interceptor = NopHandlerFunc
	ctr.TransIn.Handler = ctr.DtmTransIn
	ctr.TransInConfirm.Method = "POST"
	ctr.TransInConfirm.RelativePath = "/dtm/trans_in_confirm"
	ctr.TransInConfirm.Auth = NopHandlerFunc
	ctr.TransInConfirm.Roles = NopHandlerFunc
	ctr.TransInConfirm.Cache = NopHandlerFunc
	ctr.TransInConfirm.Interceptor = NopHandlerFunc
	ctr.TransInConfirm.Handler = ctr.DtmTransInConfirm
	ctr.TransInRevert.Method = "POST"
	ctr.TransInRevert.RelativePath = "/dtm/trans_in_revert"
	ctr.TransInRevert.Auth = NopHandlerFunc
	ctr.TransInRevert.Roles = NopHandlerFunc
	ctr.TransInRevert.Cache = NopHandlerFunc
	ctr.TransInRevert.Interceptor = NopHandlerFunc
	ctr.TransInRevert.Handler = ctr.DtmTransInRevert
	return ctr
}

// DtmRoutes defined
func DtmRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), DtmInstance
	group.Handle(instance.Tcc.Method, instance.Tcc.RelativePath, instance.Tcc.Auth, instance.Tcc.Roles, instance.Tcc.Cache, instance.Tcc.Interceptor, instance.Tcc.Handler)
	group.Handle(instance.TransOut.Method, instance.TransOut.RelativePath, instance.TransOut.Auth, instance.TransOut.Roles, instance.TransOut.Cache, instance.TransOut.Interceptor, instance.TransOut.Handler)
	group.Handle(instance.TransOutConfirm.Method, instance.TransOutConfirm.RelativePath, instance.TransOutConfirm.Auth, instance.TransOutConfirm.Roles, instance.TransOutConfirm.Cache, instance.TransOutConfirm.Interceptor, instance.TransOutConfirm.Handler)
	group.Handle(instance.TransOutRevert.Method, instance.TransOutRevert.RelativePath, instance.TransOutRevert.Auth, instance.TransOutRevert.Roles, instance.TransOutRevert.Cache, instance.TransOutRevert.Interceptor, instance.TransOutRevert.Handler)
	group.Handle(instance.TransIn.Method, instance.TransIn.RelativePath, instance.TransIn.Auth, instance.TransIn.Roles, instance.TransIn.Cache, instance.TransIn.Interceptor, instance.TransIn.Handler)
	group.Handle(instance.TransInConfirm.Method, instance.TransInConfirm.RelativePath, instance.TransInConfirm.Auth, instance.TransInConfirm.Roles, instance.TransInConfirm.Cache, instance.TransInConfirm.Interceptor, instance.TransInConfirm.Handler)
	group.Handle(instance.TransInRevert.Method, instance.TransInRevert.RelativePath, instance.TransInRevert.Auth, instance.TransInRevert.Roles, instance.TransInRevert.Cache, instance.TransInRevert.Interceptor, instance.TransInRevert.Handler)
}

// DtmInstance defined
var DtmInstance = NewDtm()

// Encrypt defined
type Encrypt struct {
	Name string
	Srv  *srv.Encrypt
	Add,
	Info Controller
}

// NewEncrypt defined
func NewEncrypt() *Encrypt {
	ctr := &Encrypt{Name: "encrypt", Srv: srv.NewEncrypt()}
	ctr.Add.Method = "POST"
	ctr.Add.RelativePath = "/encrypt/add"
	ctr.Add.Auth = Auth("encrypt")
	ctr.Add.Roles = NopHandlerFunc
	ctr.Add.Cache = NopHandlerFunc
	ctr.Add.Interceptor = NopHandlerFunc
	ctr.Add.Handler = ctr.EncryptAdd
	ctr.Info.Method = "GET"
	ctr.Info.RelativePath = "/encrypt/info"
	ctr.Info.Auth = NopHandlerFunc
	ctr.Info.Roles = NopHandlerFunc
	ctr.Info.Cache = NopHandlerFunc
	ctr.Info.Interceptor = NopHandlerFunc
	ctr.Info.Handler = ctr.EncryptInfo
	return ctr
}

// EncryptRoutes defined
func EncryptRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), EncryptInstance
	group.Handle(instance.Add.Method, instance.Add.RelativePath, instance.Add.Auth, instance.Add.Roles, instance.Add.Cache, instance.Add.Interceptor, instance.Add.Handler)
	group.Handle(instance.Info.Method, instance.Info.RelativePath, instance.Info.Auth, instance.Info.Roles, instance.Info.Cache, instance.Info.Interceptor, instance.Info.Handler)
}

// EncryptInstance defined
var EncryptInstance = NewEncrypt()

// Jwt defined
type Jwt struct {
	Name string
	Srv  *srv.Jwt
	Login,
	Check Controller
}

// NewJwt defined
func NewJwt() *Jwt {
	ctr := &Jwt{Name: "jwt", Srv: srv.NewJwt()}
	ctr.Login.Method = "POST"
	ctr.Login.RelativePath = "/jwt/login"
	ctr.Login.Auth = NopHandlerFunc
	ctr.Login.Roles = NopHandlerFunc
	ctr.Login.Cache = NopHandlerFunc
	ctr.Login.Interceptor = NopHandlerFunc
	ctr.Login.Handler = ctr.JwtLogin
	ctr.Check.Method = "GET"
	ctr.Check.RelativePath = "/jwt/check"
	ctr.Check.Auth = Auth("jwt")
	ctr.Check.Roles = NopHandlerFunc
	ctr.Check.Cache = NopHandlerFunc
	ctr.Check.Interceptor = NopHandlerFunc
	ctr.Check.Handler = ctr.JwtCheck
	return ctr
}

// JwtRoutes defined
func JwtRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), JwtInstance
	group.Handle(instance.Login.Method, instance.Login.RelativePath, instance.Login.Auth, instance.Login.Roles, instance.Login.Cache, instance.Login.Interceptor, instance.Login.Handler)
	group.Handle(instance.Check.Method, instance.Check.RelativePath, instance.Check.Auth, instance.Check.Roles, instance.Check.Cache, instance.Check.Interceptor, instance.Check.Handler)
}

// JwtInstance defined
var JwtInstance = NewJwt()

// Kafka defined
type Kafka struct {
	Name string
	Srv  *srv.Kafka
	Add,
	Get Controller
}

// NewKafka defined
func NewKafka() *Kafka {
	ctr := &Kafka{Name: "kafka", Srv: srv.NewKafka()}
	ctr.Add.Method = "POST"
	ctr.Add.RelativePath = "/kafka/add"
	ctr.Add.Auth = Auth("token")
	ctr.Add.Roles = NopHandlerFunc
	ctr.Add.Cache = NopHandlerFunc
	ctr.Add.Interceptor = NopHandlerFunc
	ctr.Add.Handler = ctr.KafkaAdd
	ctr.Get.Method = "GET"
	ctr.Get.RelativePath = "/kafka/get"
	ctr.Get.Auth = Auth("token")
	ctr.Get.Roles = NopHandlerFunc
	ctr.Get.Cache = NopHandlerFunc
	ctr.Get.Interceptor = NopHandlerFunc
	ctr.Get.Handler = ctr.KafkaGet
	return ctr
}

// KafkaRoutes defined
func KafkaRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), KafkaInstance
	group.Handle(instance.Add.Method, instance.Add.RelativePath, instance.Add.Auth, instance.Add.Roles, instance.Add.Cache, instance.Add.Interceptor, instance.Add.Handler)
	group.Handle(instance.Get.Method, instance.Get.RelativePath, instance.Get.Auth, instance.Get.Roles, instance.Get.Cache, instance.Get.Interceptor, instance.Get.Handler)
}

// KafkaInstance defined
var KafkaInstance = NewKafka()

// Middles defined
type Middles struct {
	Name string
	Srv  *srv.Middles
	GlobalTest,
	LocalTest Controller
}

// NewMiddles defined
func NewMiddles() *Middles {
	ctr := &Middles{Name: "middles", Srv: srv.NewMiddles()}
	ctr.GlobalTest.Method = "GET"
	ctr.GlobalTest.RelativePath = "/middles/globalTest"
	ctr.GlobalTest.Auth = NopHandlerFunc
	ctr.GlobalTest.Roles = NopHandlerFunc
	ctr.GlobalTest.Cache = NopHandlerFunc
	ctr.GlobalTest.Interceptor = NopHandlerFunc
	ctr.GlobalTest.Handler = ctr.MiddlesGlobalTest
	ctr.LocalTest.Method = "GET"
	ctr.LocalTest.RelativePath = "/middles/localTest"
	ctr.LocalTest.Auth = NopHandlerFunc
	ctr.LocalTest.Roles = NopHandlerFunc
	ctr.LocalTest.Cache = NopHandlerFunc
	ctr.LocalTest.Interceptor = NopHandlerFunc
	ctr.LocalTest.Handler = ctr.MiddlesLocalTest
	return ctr
}

// MiddlesRoutes defined
func MiddlesRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), MiddlesInstance
	group.Handle(instance.GlobalTest.Method, instance.GlobalTest.RelativePath, instance.GlobalTest.Auth, instance.GlobalTest.Roles, instance.GlobalTest.Cache, instance.GlobalTest.Interceptor, instance.GlobalTest.Handler)
	group.Handle(instance.LocalTest.Method, instance.LocalTest.RelativePath, instance.LocalTest.Auth, instance.LocalTest.Roles, instance.LocalTest.Cache, instance.LocalTest.Interceptor, instance.LocalTest.Handler)
}

// MiddlesInstance defined
var MiddlesInstance = NewMiddles()

// Nsq defined
type Nsq struct {
	Name string
	Srv  *srv.Nsq
	Add,
	Get Controller
}

// NewNsq defined
func NewNsq() *Nsq {
	ctr := &Nsq{Name: "nsq", Srv: srv.NewNsq()}
	ctr.Add.Method = "POST"
	ctr.Add.RelativePath = "/nsq/add"
	ctr.Add.Auth = Auth("token")
	ctr.Add.Roles = NopHandlerFunc
	ctr.Add.Cache = NopHandlerFunc
	ctr.Add.Interceptor = NopHandlerFunc
	ctr.Add.Handler = ctr.NsqAdd
	ctr.Get.Method = "GET"
	ctr.Get.RelativePath = "/nsq/get"
	ctr.Get.Auth = Auth("token")
	ctr.Get.Roles = NopHandlerFunc
	ctr.Get.Cache = NopHandlerFunc
	ctr.Get.Interceptor = NopHandlerFunc
	ctr.Get.Handler = ctr.NsqGet
	return ctr
}

// NsqRoutes defined
func NsqRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), NsqInstance
	group.Handle(instance.Add.Method, instance.Add.RelativePath, instance.Add.Auth, instance.Add.Roles, instance.Add.Cache, instance.Add.Interceptor, instance.Add.Handler)
	group.Handle(instance.Get.Method, instance.Get.RelativePath, instance.Get.Auth, instance.Get.Roles, instance.Get.Cache, instance.Get.Interceptor, instance.Get.Handler)
}

// NsqInstance defined
var NsqInstance = NewNsq()

// RedisLock defined
type RedisLock struct {
	Name string
	Srv  *srv.RedisLock
	Lock,
	Unlock Controller
}

// NewRedisLock defined
func NewRedisLock() *RedisLock {
	ctr := &RedisLock{Name: "redis_lock", Srv: srv.NewRedisLock()}
	ctr.Lock.Method = "GET"
	ctr.Lock.RelativePath = "/redis/lock/lock"
	ctr.Lock.Auth = Auth("token")
	ctr.Lock.Roles = NopHandlerFunc
	ctr.Lock.Cache = NopHandlerFunc
	ctr.Lock.Interceptor = NopHandlerFunc
	ctr.Lock.Handler = ctr.RedisLockLock
	ctr.Unlock.Method = "GET"
	ctr.Unlock.RelativePath = "/redis/lock/unlock"
	ctr.Unlock.Auth = Auth("token")
	ctr.Unlock.Roles = NopHandlerFunc
	ctr.Unlock.Cache = NopHandlerFunc
	ctr.Unlock.Interceptor = NopHandlerFunc
	ctr.Unlock.Handler = ctr.RedisLockUnlock
	return ctr
}

// RedisLockRoutes defined
func RedisLockRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), RedisLockInstance
	group.Handle(instance.Lock.Method, instance.Lock.RelativePath, instance.Lock.Auth, instance.Lock.Roles, instance.Lock.Cache, instance.Lock.Interceptor, instance.Lock.Handler)
	group.Handle(instance.Unlock.Method, instance.Unlock.RelativePath, instance.Unlock.Auth, instance.Unlock.Roles, instance.Unlock.Cache, instance.Unlock.Interceptor, instance.Unlock.Handler)
}

// RedisLockInstance defined
var RedisLockInstance = NewRedisLock()

// RedisMq defined
type RedisMq struct {
	Name string
	Srv  *srv.RedisMq
	Add,
	Get Controller
}

// NewRedisMq defined
func NewRedisMq() *RedisMq {
	ctr := &RedisMq{Name: "redis_mq", Srv: srv.NewRedisMq()}
	ctr.Add.Method = "POST"
	ctr.Add.RelativePath = "/redis/mq/add"
	ctr.Add.Auth = Auth("token")
	ctr.Add.Roles = NopHandlerFunc
	ctr.Add.Cache = NopHandlerFunc
	ctr.Add.Interceptor = NopHandlerFunc
	ctr.Add.Handler = ctr.RedisMqAdd
	ctr.Get.Method = "GET"
	ctr.Get.RelativePath = "/redis/mq/get"
	ctr.Get.Auth = Auth("token")
	ctr.Get.Roles = NopHandlerFunc
	ctr.Get.Cache = NopHandlerFunc
	ctr.Get.Interceptor = NopHandlerFunc
	ctr.Get.Handler = ctr.RedisMqGet
	return ctr
}

// RedisMqRoutes defined
func RedisMqRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), RedisMqInstance
	group.Handle(instance.Add.Method, instance.Add.RelativePath, instance.Add.Auth, instance.Add.Roles, instance.Add.Cache, instance.Add.Interceptor, instance.Add.Handler)
	group.Handle(instance.Get.Method, instance.Get.RelativePath, instance.Get.Auth, instance.Get.Roles, instance.Get.Cache, instance.Get.Interceptor, instance.Get.Handler)
}

// RedisMqInstance defined
var RedisMqInstance = NewRedisMq()

// RPC defined
type RPC struct {
	Name    string
	Srv     *srv.RPC
	Message Controller
}

// NewRPC defined
func NewRPC() *RPC {
	ctr := &RPC{Name: "rpc", Srv: srv.NewRPC()}
	ctr.Message.Method = "GET"
	ctr.Message.RelativePath = "/rpc/message"
	ctr.Message.Auth = Auth("token")
	ctr.Message.Roles = NopHandlerFunc
	ctr.Message.Cache = NopHandlerFunc
	ctr.Message.Interceptor = NopHandlerFunc
	ctr.Message.Handler = ctr.RPCMessage
	return ctr
}

// RPCRoutes defined
func RPCRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), RPCInstance
	group.Handle(instance.Message.Method, instance.Message.RelativePath, instance.Message.Auth, instance.Message.Roles, instance.Message.Cache, instance.Message.Interceptor, instance.Message.Handler)
}

// RPCInstance defined
var RPCInstance = NewRPC()

// Sqlmap defined
type Sqlmap struct {
	Name      string
	Srv       *srv.Sqlmap
	Selectone Controller
}

// NewSqlmap defined
func NewSqlmap() *Sqlmap {
	ctr := &Sqlmap{Name: "sqlmap", Srv: srv.NewSqlmap()}
	ctr.Selectone.Method = "GET"
	ctr.Selectone.RelativePath = "/sqlmap/selectone"
	ctr.Selectone.Auth = NopHandlerFunc
	ctr.Selectone.Roles = NopHandlerFunc
	ctr.Selectone.Cache = NopHandlerFunc
	ctr.Selectone.Interceptor = NopHandlerFunc
	ctr.Selectone.Handler = ctr.SqlmapSelectone
	return ctr
}

// SqlmapRoutes defined
func SqlmapRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), SqlmapInstance
	group.Handle(instance.Selectone.Method, instance.Selectone.RelativePath, instance.Selectone.Auth, instance.Selectone.Roles, instance.Selectone.Cache, instance.Selectone.Interceptor, instance.Selectone.Handler)
}

// SqlmapInstance defined
var SqlmapInstance = NewSqlmap()

// User defined
type User struct {
	Name string
	Srv  *srv.User
	Info Controller
}

// NewUser defined
func NewUser() *User {
	ctr := &User{Name: "user", Srv: srv.NewUser()}
	ctr.Info.Method = "GET"
	ctr.Info.RelativePath = "/user/info"
	ctr.Info.Auth = Auth("token")
	ctr.Info.Roles = NopHandlerFunc
	ctr.Info.Cache = NopHandlerFunc
	ctr.Info.Interceptor = NopHandlerFunc
	ctr.Info.Handler = ctr.UserInfo
	return ctr
}

// UserRoutes defined
func UserRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), UserInstance
	group.Handle(instance.Info.Method, instance.Info.RelativePath, instance.Info.Auth, instance.Info.Roles, instance.Info.Cache, instance.Info.Interceptor, instance.Info.Handler)
}

// UserInstance defined
var UserInstance = NewUser()

// View defined
type View struct {
	Name string
	Srv  *srv.View
	File,
	HTML,
	XML Controller
}

// NewView defined
func NewView() *View {
	ctr := &View{Name: "view", Srv: srv.NewView()}
	ctr.File.Method = "GET"
	ctr.File.RelativePath = "/view/file"
	ctr.File.Auth = NopHandlerFunc
	ctr.File.Roles = NopHandlerFunc
	ctr.File.Cache = NopHandlerFunc
	ctr.File.Interceptor = NopHandlerFunc
	ctr.File.Handler = ctr.ViewFile
	ctr.HTML.Method = "GET"
	ctr.HTML.RelativePath = "/view/html"
	ctr.HTML.Auth = NopHandlerFunc
	ctr.HTML.Roles = NopHandlerFunc
	ctr.HTML.Cache = NopHandlerFunc
	ctr.HTML.Interceptor = NopHandlerFunc
	ctr.HTML.Handler = ctr.ViewHTML
	ctr.XML.Method = "GET"
	ctr.XML.RelativePath = "/view/xml"
	ctr.XML.Auth = NopHandlerFunc
	ctr.XML.Roles = NopHandlerFunc
	ctr.XML.Cache = NopHandlerFunc
	ctr.XML.Interceptor = NopHandlerFunc
	ctr.XML.Handler = ctr.ViewXML
	return ctr
}

// ViewRoutes defined
func ViewRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), ViewInstance
	group.Handle(instance.File.Method, instance.File.RelativePath, instance.File.Auth, instance.File.Roles, instance.File.Cache, instance.File.Interceptor, instance.File.Handler)
	group.Handle(instance.HTML.Method, instance.HTML.RelativePath, instance.HTML.Auth, instance.HTML.Roles, instance.HTML.Cache, instance.HTML.Interceptor, instance.HTML.Handler)
	group.Handle(instance.XML.Method, instance.XML.RelativePath, instance.XML.Auth, instance.XML.Roles, instance.XML.Cache, instance.XML.Interceptor, instance.XML.Handler)
}

// ViewInstance defined
var ViewInstance = NewView()

// Vote defined
type Vote struct {
	Name string
	Srv  *srv.Vote
	Like Controller
}

// NewVote defined
func NewVote() *Vote {
	ctr := &Vote{Name: "vote", Srv: srv.NewVote()}
	ctr.Like.Method = "POST"
	ctr.Like.RelativePath = "/vote/like"
	ctr.Like.Auth = Auth("token")
	ctr.Like.Roles = NopHandlerFunc
	ctr.Like.Cache = NopHandlerFunc
	ctr.Like.Interceptor = NopHandlerFunc
	ctr.Like.Handler = ctr.VoteLike
	return ctr
}

// VoteRoutes defined
func VoteRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), VoteInstance
	group.Handle(instance.Like.Method, instance.Like.RelativePath, instance.Like.Auth, instance.Like.Roles, instance.Like.Cache, instance.Like.Interceptor, instance.Like.Handler)
}

// VoteInstance defined
var VoteInstance = NewVote()

// Xlsx defined
type Xlsx struct {
	Name   string
	Srv    *srv.Xlsx
	Import Controller
}

// NewXlsx defined
func NewXlsx() *Xlsx {
	ctr := &Xlsx{Name: "xlsx", Srv: srv.NewXlsx()}
	ctr.Import.Method = "POST"
	ctr.Import.RelativePath = "/xlsx/import"
	ctr.Import.Auth = NopHandlerFunc
	ctr.Import.Roles = NopHandlerFunc
	ctr.Import.Cache = NopHandlerFunc
	ctr.Import.Interceptor = NopHandlerFunc
	ctr.Import.Handler = ctr.XlsxImport
	return ctr
}

// XlsxRoutes defined
func XlsxRoutes(rg *RouterGroup) {
	group, instance := rg.Group(viper.GetString("http.prefix")), XlsxInstance
	group.Handle(instance.Import.Method, instance.Import.RelativePath, instance.Import.Auth, instance.Import.Roles, instance.Import.Cache, instance.Import.Interceptor, instance.Import.Handler)
}

// XlsxInstance defined
var XlsxInstance = NewXlsx()

// MessageSrv defined
func MessageSrvService(dol *Dolphin) {
	dol.Remote.RegisterServer(func(gs *grpc.Server) { proto.RegisterMessageSrvServer(gs, &rpc.MessageSrv{}) })
}

// SyncModel defined
func (dol *Dolphin) SyncModel() error {
	mseti := dol.Manager.ModelSet()
	mseti.Add(new(types.Article))
	mseti.Add(new(types.PostLike))
	mseti.Add(new(types.SysAppFun))
	mseti.Add(new(types.SysArea))
	mseti.Add(new(types.SysAreaTemplate))
	mseti.Add(new(types.SysAreaTemplateDetail))
	mseti.Add(new(types.SysAttachment))
	mseti.Add(new(types.SysComment))
	mseti.Add(new(types.SysCommentReply))
	mseti.Add(new(types.SysDataPermission))
	mseti.Add(new(types.SysDataPermissionDetail))
	mseti.Add(new(types.SysEmailToken))
	mseti.Add(new(types.SysMenu))
	mseti.Add(new(types.SysNotification))
	mseti.Add(new(types.SysOptionset))
	mseti.Add(new(types.SysOrg))
	mseti.Add(new(types.SysPermission))
	mseti.Add(new(types.SysRole))
	mseti.Add(new(types.SysRoleAppFun))
	mseti.Add(new(types.SysRoleDataPermission))
	mseti.Add(new(types.SysRoleMenu))
	mseti.Add(new(types.SysRolePermission))
	mseti.Add(new(types.SysRoleUser))
	mseti.Add(new(types.SysSchedule))
	mseti.Add(new(types.SysScheduleHistory))
	mseti.Add(new(types.SysSetting))
	mseti.Add(new(types.SysTable))
	mseti.Add(new(types.SysTableColUser))
	mseti.Add(new(types.SysTableColumn))
	mseti.Add(new(types.SysTag))
	mseti.Add(new(types.SysTagGroup))
	mseti.Add(new(types.SysTracker))
	mseti.Add(new(types.SysUserBinding))
	mseti.Add(new(types.SysUserTag))
	mseti.Add(new(types.SysUserTemplate))
	mseti.Add(new(types.SysUserTemplateDetail))
	mseti.Add(new(types.UserAccount))
	mseti.Add(new(types.UserAccountTrading))
	mseti.Add(new(types.UserLikePost))
	return nil
}

// SyncController defined
func (dol *Dolphin) SyncController() error {
	ArticleRoutes(&dol.RouterGroup)
	CachingRoutes(&dol.RouterGroup)
	DtmRoutes(&dol.RouterGroup)
	EncryptRoutes(&dol.RouterGroup)
	JwtRoutes(&dol.RouterGroup)
	KafkaRoutes(&dol.RouterGroup)
	MiddlesRoutes(&dol.RouterGroup)
	NsqRoutes(&dol.RouterGroup)
	RedisLockRoutes(&dol.RouterGroup)
	RedisMqRoutes(&dol.RouterGroup)
	RPCRoutes(&dol.RouterGroup)
	SqlmapRoutes(&dol.RouterGroup)
	UserRoutes(&dol.RouterGroup)
	ViewRoutes(&dol.RouterGroup)
	VoteRoutes(&dol.RouterGroup)
	XlsxRoutes(&dol.RouterGroup)
	return nil
}

// SyncSrv defined
func (dol *Dolphin) SyncSrv(svc *svc.ServiceContext) error {
	ArticleInstance.Srv.SetServiceContext(svc)
	CachingInstance.Srv.SetServiceContext(svc)
	DtmInstance.Srv.SetServiceContext(svc)
	EncryptInstance.Srv.SetServiceContext(svc)
	JwtInstance.Srv.SetServiceContext(svc)
	KafkaInstance.Srv.SetServiceContext(svc)
	MiddlesInstance.Srv.SetServiceContext(svc)
	NsqInstance.Srv.SetServiceContext(svc)
	RedisLockInstance.Srv.SetServiceContext(svc)
	RedisMqInstance.Srv.SetServiceContext(svc)
	RPCInstance.Srv.SetServiceContext(svc)
	SqlmapInstance.Srv.SetServiceContext(svc)
	UserInstance.Srv.SetServiceContext(svc)
	ViewInstance.Srv.SetServiceContext(svc)
	VoteInstance.Srv.SetServiceContext(svc)
	XlsxInstance.Srv.SetServiceContext(svc)
	return nil
}

// SyncService defined
func (dol *Dolphin) SyncService() error {
	MessageSrvService(dol)
	return nil
}
