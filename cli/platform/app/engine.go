package app

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/2637309949/dolphin/cli/packages/gin"
	"github.com/2637309949/dolphin/cli/packages/logrus"
	"github.com/2637309949/dolphin/cli/packages/null"
	"github.com/2637309949/dolphin/cli/packages/oauth2/errors"
	oaErrors "github.com/2637309949/dolphin/cli/packages/oauth2/errors"
	"github.com/2637309949/dolphin/cli/packages/oauth2/generates"
	"github.com/2637309949/dolphin/cli/packages/oauth2/manage"
	"github.com/2637309949/dolphin/cli/packages/oauth2/models"
	"github.com/2637309949/dolphin/cli/packages/oauth2/server"
	"github.com/2637309949/dolphin/cli/packages/oauth2/store"
	"github.com/2637309949/dolphin/cli/packages/redis"
	"github.com/2637309949/dolphin/cli/packages/viper"
	"github.com/2637309949/dolphin/cli/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/2637309949/dolphin/cli/platform/sql"
	"github.com/2637309949/dolphin/cli/platform/util"
	"github.com/go-session/session"
)

// Engine defined parse app engine
type Engine struct {
	MSet          MSeti
	Gin           *gin.Engine
	Redis         *redis.Client
	PlatformDB    *xorm.Engine
	BusinessDBSet map[string]*xorm.Engine
	OAuth2        *server.Server
	pool          sync.Pool
}

func (e *Engine) allocateContext() *Context {
	return &Context{engine: e, AuthInfo: &AuthOAuth2{server: e.OAuth2}}
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{engine: e, RouterGroup: e.Gin.Group(relativePath, handlers...)}
}

// Migration models
func (e *Engine) Migration(name string, db *xorm.Engine) error {
	e.MSet.ForEach(func(n string, m interface{}) {
		if n == name {
			db.Sync2(m)
		}
	})
	return nil
}

// InitBusinessDB load BusinessDBSet
func (e *Engine) InitBusinessDB() {
	domains := []model.Domain{}
	if err := e.PlatformDB.Where("data_source <> '' and domain <> '' and delete_time is null").Find(&domains); err != nil {
		logrus.Fatal(err)
	}
	for _, domain := range domains {
		uri, err := util.Parse(domain.DataSource.String)
		if err != nil {
			panic(err)
		}
		_, err = e.PlatformDB.Sql(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v DEFAULT CHARACTER SET utf8mb4", uri.DbName)).Execute()
		if err != nil {
			panic(err)
		}
		e.AddBusinessDB(domain.Domain.String, domain.DriverName.String, domain.DataSource.String)
	}
	nset := e.MSet.Name(func(n string) bool {
		return n != Name
	})
	for _, db := range e.BusinessDBSet {
		for _, n := range nset {
			e.Migration(n, db)
		}
	}
}

// GetBusinessDB businessDB
func (e *Engine) GetBusinessDB(domain string) (db *xorm.Engine, ok bool) {
	db, ok = e.BusinessDBSet[domain]
	return
}

// AddBusinessDB adddb
func (e *Engine) AddBusinessDB(domain, driverName, dataSource string) {
	sqlDir := viper.GetString("dir.sql")
	sqlDir = path.Join(".", sqlDir)

	db, err := xorm.NewEngine(driverName, dataSource)
	if err != nil {
		logrus.Fatal(err)
	}

	xLogger := xorm.NewSimpleLogger(logrus.StandardLogger().Out)
	if viper.GetString("app.mode") == "debug" {
		xLogger.ShowSQL(true)
	}
	db.SetLogger(xLogger)

	if err = os.MkdirAll(sqlDir, os.ModePerm); err != nil {
		logrus.Fatal(err)
	}
	if err = db.RegisterSqlMap(xorm.Xml(sqlDir, ".xml")); err != nil {
		logrus.Fatal(err)
	}
	if err = db.RegisterSqlTemplate(xorm.Pongo2(sqlDir, ".stpl")); err != nil {
		logrus.Fatal(err)
	}
	if err = db.RegisterSqlTemplate(xorm.Jet(sqlDir, ".jet")); err != nil {
		logrus.Fatal(err)
	}
	if err = db.RegisterSqlTemplate(xorm.Default(sqlDir, ".tpl")); err != nil {
		logrus.Fatal(err)
	}
	e.BusinessDBSet[domain] = db
}

// InitPlatformDB adddb
func (e *Engine) InitPlatformDB() {
	var err error
	sqlDir := viper.GetString("dir.sql")
	sqlDir = path.Join(".", sqlDir)
	e.PlatformDB, err = xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	if err != nil {
		logrus.Fatal(err)
	}
	if err = e.PlatformDB.Ping(); err != nil {
		logrus.Fatal(err)
	}
	xLogger := xorm.NewSimpleLogger(logrus.StandardLogger().Out)
	if viper.GetString("app.mode") == "debug" {
		xLogger.ShowSQL(true)
	}
	e.PlatformDB.SetLogger(xLogger)
	e.PlatformDB.SqlTemplate = &xorm.HTMLTemplate{
		Template: make(map[string]*template.Template, 100),
		Funcs:    make(map[string]xorm.FuncMap, 20),
	}
	for k, v := range sql.SQLTPL {
		if err = e.PlatformDB.AddSqlTemplate(k, v); err != nil {
			logrus.Fatal(err)
		}
	}
	e.Migration(Name, e.PlatformDB)

	s := e.PlatformDB.NewSession()
	domain := model.Domain{
		ID:         null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		Name:       null.StringFrom("系统默认"),
		FullName:   null.StringFrom("系统默认"),
		DataSource: null.StringFrom(""),
		DriverName: null.StringFrom("mysql"),
		DomainUrl:  null.StringFrom("localhost"),
		LoginUrl:   null.StringFrom("localhost"),
		Type:       null.IntFrom(0),
		Status:     null.IntFrom(1),
		SyncFlag:   null.IntFrom(0),
		Domain:     null.StringFrom("localhost"),
		CreateBy:   null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		CreateTime: null.TimeFrom(time.Now()),
		UpdateBy:   null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		UpdateTime: null.TimeFrom(time.Now()),
	}
	ct, err := s.Where("id=?", domain.ID).Count(new(model.Domain))
	if err != nil {
		s.Rollback()
		panic(err)
	}
	if ct == 0 {
		_, err := e.PlatformDB.InsertOne(&domain)
		if err != nil {
			s.Rollback()
			panic(err)
		}
	}
	admin := model.User{
		ID:         null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		Password:   null.StringFrom("123456"),
		Name:       null.StringFrom("admin"),
		FullName:   null.StringFrom("admin"),
		Status:     null.IntFrom(1),
		Domain:     null.StringFrom("localhost"),
		CreateBy:   null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		CreateTime: null.TimeFrom(time.Now()),
		UpdateBy:   null.StringFrom("6ba7b810-9dad-11d1-80b4-00c04fd430c8"),
		UpdateTime: null.TimeFrom(time.Now()),
	}
	admin.SetPassword(admin.Password.String)
	ct, err = s.Where("id=?", admin.ID).Count(new(model.User))
	if err != nil {
		s.Rollback()
		panic(err)
	}
	if ct == 0 {
		_, err := e.PlatformDB.InsertOne(&admin)
		if err != nil {
			s.Rollback()
			panic(err)
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}

// InitRedis redis
func (e *Engine) InitRedis() {
	uri, err := util.Parse(viper.GetString("rd.dataSource"))
	if err != nil {
		logrus.Fatal(err)
	}
	db, err := strconv.Atoi(uri.DbName)
	if err != nil {
		logrus.Fatal(err)
	}
	rds := redis.NewClient(&redis.Options{
		Addr:     uri.Laddr,
		Password: uri.Passwd,
		DB:       db,
	})
	if _, err := rds.Ping().Result(); err != nil {
		logrus.Fatal(err)
	}
	e.Redis = rds
}

// InitOAuth2 oauth2
func (e *Engine) InitOAuth2() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapTokenStorage(store.NewRedisStoreWithCli(e.Redis))
	manager.MapAccessGenerate(generates.NewAccessGenerate())
	clientStore := store.NewClientStore()
	clientStore.Set(viper.GetString("oauth.id"), &models.Client{
		ID:     viper.GetString("oauth.id"),
		Secret: viper.GetString("oauth.secret"),
		Domain: viper.GetString("oauth.cli"),
	})
	manager.MapClientStorage(clientStore)

	e.OAuth2 = server.NewServer(server.NewConfig(), manager)
	e.OAuth2.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (uid string, dm string, err error) {
		store, err := session.Start(nil, w, r)
		if err != nil {
			return
		}
		userID, uok := store.Get("LoggedInUserID")
		domain, dok := store.Get("LoggedInDomain")
		if !uok || !dok {
			if r.Form == nil {
				r.ParseForm()
			}
			store.Set("ReturnUri", r.Form)
			store.Save()
			w.Header().Set("Location", viper.GetString("oauth.login"))
			w.WriteHeader(http.StatusFound)
			return
		}
		uid = userID.(string)
		dm = domain.(string)
		store.Delete("LoggedInUserID")
		store.Delete("LoggedInDomain")
		store.Save()
		return
	})
	e.OAuth2.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logrus.Error("Internal Error:", err.Error())
		return
	})
	e.OAuth2.SetResponseErrorHandler(func(re *errors.Response) {
		logrus.Error("Response Error:", re.Error.Error())
	})
}

// Auth middles
func (e *Engine) Auth(mode ...AuthType) func(ctx *Context) {
	return func(ctx *Context) {
		if ctx.Auth(ctx.Request) {
			ctx.DB = e.BusinessDBSet[ctx.GetToken().GetDomain()]
			ctx.Set("DB", ctx.DB)
			ctx.Set("AuthInfo", ctx.AuthInfo)
			ctx.Next()
		} else {
			ctx.Fail(oaErrors.ErrInvalidAccessToken)
		}
	}
}

// Run booting system
func (e *Engine) Run() {
	e.InitRedis()
	e.InitPlatformDB()
	e.InitBusinessDB()
	e.InitOAuth2()
}

// BuildEngine build engine
func BuildEngine(build func(*Engine)) func(*Engine) {
	return func(e *Engine) {
		build(e)
	}
}

// NewEngine init Engine
func NewEngine() *Engine {
	gin.SetMode(gin.ReleaseMode)
	e := &Engine{}
	e.MSet = &MSet{m: map[string][]interface{}{}}
	e.BusinessDBSet = map[string]*xorm.Engine{}
	e.Gin = gin.New()
	e.Gin.Use(gin.Logger())
	e.Gin.Use(util.Recovery(func(ctx *gin.Context, err interface{}) {
		code := 500
		if err, ok := err.(model.Error); ok {
			code = err.Code
		}
		ctx.JSON(http.StatusInternalServerError, util.M{"code": code, "message": err})
	}))
	e.Gin.Use(util.ProcessMethodOverride(e.Gin))
	e.Gin.Static(viper.GetString("http.static"), path.Join(util.Getwd(), viper.GetString("http.static")))
	e.pool.New = func() interface{} {
		return e.allocateContext()
	}
	return e
}
