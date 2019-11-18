package app

import (
	"html/template"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/2637309949/dolphin/cli/platform/sql"
	"github.com/2637309949/dolphin/cli/platform/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-session/session"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xormplus/xorm"
	oredis "gopkg.in/go-oauth2/redis.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

// Engine defined parse app engine
type Engine struct {
	MSets         MSeti
	Gin           *gin.Engine
	Redis         *redis.Client
	PlatformDB    *xorm.Engine
	BusinessDBSet map[string]*xorm.Engine
	OAuth2        *server.Server
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{Engine: e, RouterGroup: e.Gin.Group(relativePath, handlers...)}
}

// Migration models
func (e *Engine) Migration(name string, db *xorm.Engine) error {
	e.MSets.ForEach(func(n string, m interface{}) {
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
		e.AddBusinessDB(domain.Domain.String, domain.DriverName.String, domain.DataSource.String)
	}
	nset := e.MSets.Name(func(n string) bool {
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
	// only load Platform sql
	e.PlatformDB.SqlTemplate = &xorm.HTMLTemplate{
		Template: make(map[string]*template.Template, 100),
		Funcs:    make(map[string]xorm.FuncMap, 20),
	}
	for k, v := range sql.SQLTPL {
		if err = e.PlatformDB.AddSqlTemplate(k, v); err != nil {
			logrus.Fatal(err)
		}
	}
	// async platform model
	e.Migration(Name, e.PlatformDB)
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
	uri, err := util.Parse(viper.GetString("rd.dataSource"))
	if err != nil {
		logrus.Fatal(err)
	}
	db, err := strconv.Atoi(uri.DbName)
	if err != nil {
		logrus.Fatal(err)
	}
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr:     uri.Laddr,
		Password: uri.Passwd,
		DB:       db,
	}))
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("00000000"), jwt.SigningMethodHS512))
	clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		Secret: "222222",
		Domain: "http://127.0.0.1:8081",
	})
	manager.MapClientStorage(clientStore)
	e.OAuth2 = server.NewServer(server.NewConfig(), manager)
	e.OAuth2.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "test" && password == "test" {
			userID = "test"
		}
		return
	})
	e.OAuth2.SetUserAuthorizationHandler(func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
		store, err := session.Start(nil, w, r)
		if err != nil {
			return
		}
		uid, ok := store.Get("LoggedInUserID")
		if !ok {
			if r.Form == nil {
				r.ParseForm()
			}
			store.Set("ReturnUri", r.Form)
			store.Save()
			w.Header().Set("Location", viper.GetString("oauth.login"))
			w.WriteHeader(http.StatusFound)
			return
		}
		userID = uid.(string)
		store.Delete("LoggedInUserID")
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

// Run booting system
func (e *Engine) Run() {
	e.InitRedis()
	e.InitPlatformDB()
	e.InitBusinessDB()
	e.InitOAuth2()
}

// BuildEngine build engine
// Action after Init Engine
func BuildEngine(build func(*Engine)) func(*Engine) {
	return func(e *Engine) {
		build(e)
	}
}

// NewEngine init Engine
func NewEngine() *Engine {
	gin.SetMode(gin.ReleaseMode)
	e := &Engine{}
	e.MSets = &MSets{m: map[string][]interface{}{}}
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
	return e
}
