package app

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"sync"

	"google.golang.org/grpc"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/go-session/session"
	"github.com/2637309949/dolphin/packages/logrus"
	oError "github.com/2637309949/dolphin/packages/oauth2/errors"
	"github.com/2637309949/dolphin/packages/oauth2/generates"
	"github.com/2637309949/dolphin/packages/oauth2/manage"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/packages/oauth2/store"
	"github.com/2637309949/dolphin/packages/redis"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/plugin"
	"github.com/2637309949/dolphin/platform/sql"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/file"

	httpUtil "github.com/2637309949/dolphin/platform/util/http"
)

// Engine defined parse app engine
type Engine struct {
	MSet          MSeti
	Gin           *gin.Engine
	GRPC          *grpc.Server
	OAuth2        *server.Server
	Redis         *redis.Client
	PlatformDB    *xorm.Engine
	BusinessDBSet map[string]*xorm.Engine
	pool          sync.Pool
}

// HandlerFunc convert to gin.HandlerFunc
func (e *Engine) HandlerFunc(h HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		c := e.pool.Get().(*Context)
		c.Context = ctx
		for _, v := range ctx.Keys {
			switch t := v.(type) {
			case *xorm.Engine:
				c.DB = t
			case AuthInfo:
				c.AuthInfo = t
			}
		}
		h(c)
		e.pool.Put(c)
	})
}

func (e *Engine) allocateContext() *Context {
	return &Context{PlatformDB: e.PlatformDB, engine: e, AuthInfo: &AuthOAuth2{server: e.OAuth2}}
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{engine: e, RouterGroup: e.Gin.Group(relativePath, handlers...)}
}

// Migration models
func (e *Engine) migration(name string, db *xorm.Engine) {
	e.MSet.ForEach(func(n string, m interface{}) {
		if n == name {
			if err := db.Sync2(m); err != nil {
				panic(err)
			}
		}
	})
}

func (e *Engine) initBusinessDB() {
	domains := []model.SysDomain{}
	if err := e.PlatformDB.Where("data_source <> '' and domain <> '' and del_flag = 0").Find(&domains); err != nil {
		logrus.Fatal(err)
	}

	nset := e.MSet.Name(func(n string) bool {
		return n != Name
	})

	for _, domain := range domains {
		uri, err := httpUtil.Parse(domain.DataSource.String)
		if err != nil {
			panic(err)
		}
		if _, err = e.PlatformDB.Sql(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v DEFAULT CHARACTER SET utf8mb4", uri.DbName)).Execute(); err != nil {
			panic(err)
		}
		db, err := xorm.NewEngine(domain.DriverName.String, domain.DataSource.String)
		if err != nil {
			logrus.Fatal(err)
		}
		xLogger := xorm.NewSimpleLogger(logrus.StandardLogger().Out)
		if viper.GetString("app.mode") == "debug" {
			xLogger.ShowSQL(true)
		}
		db.SetLogger(xLogger)
		e.RegisterSQLDir(db, path.Join(".", viper.GetString("dir.sql")))
		e.RegisterSQLMap(db, sql.SQLTPL)

		// migration db
		if domain.SyncFlag.Int64 == 0 {
			for _, n := range nset {
				e.migration(n, db)
			}
		}
		domain.SyncFlag.SetValid(1)
		_, err = e.PlatformDB.Cols("sync_flag").Update(&domain)
		if err != nil {
			logrus.Fatal(err)
		}

		// initialize data
		(new(model.SysRole)).InitSysData(db.NewSession())
		(new(model.SysRoleUser)).InitSysData(db.NewSession())
		(new(model.SysMenu)).InitSysData(db.NewSession())

		fmt.Println("----------------------", domain.Domain.String)
		e.AddBusinessDB(domain.Domain.String, db)
	}
}

func (e *Engine) initPlatformDB() {
	var err error
	sqlDir := viper.GetString("dir.sql")
	sqlDir = path.Join(".", sqlDir)
	if e.PlatformDB, err = xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dataSource")); err != nil {
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
	e.RegisterSQLDir(e.PlatformDB, path.Join(".", viper.GetString("dir.sql")))
	e.RegisterSQLMap(e.PlatformDB, sql.SQLTPL)
	e.migration(Name, e.PlatformDB)

	// initialize data
	(new(model.SysDomain)).InitSysData(e.PlatformDB.NewSession())
	(new(model.SysClient)).InitSysData(e.PlatformDB.NewSession())
	(new(model.SysUser)).InitSysData(e.PlatformDB.NewSession())
}

func (e *Engine) initDB() {
	e.initPlatformDB()
	e.initBusinessDB()
}

// GetBusinessDB businessDB
func (e *Engine) GetBusinessDB(domain string) (db *xorm.Engine, ok bool) {
	db, ok = e.BusinessDBSet[domain]
	return
}

// RegisterSQLMap defined sql
func (e *Engine) RegisterSQLMap(db *xorm.Engine, sqlMap map[string]string) {
	for key, value := range sqlMap {
		if filepath.Ext(key) == "" {
			db.AddSql(key, value)
		} else {
			db.AddSqlTemplate(key, value)
		}
	}
}

// RegisterSQLDir defined sql
func (e *Engine) RegisterSQLDir(db *xorm.Engine, sqlDir string) {
	if err := os.MkdirAll(sqlDir, os.ModePerm); err != nil {
		logrus.Fatal(err)
	}
	if err := db.RegisterSqlMap(xorm.Xml(sqlDir, ".xml")); err != nil {
		logrus.Fatal(err)
	}
	if err := db.RegisterSqlTemplate(xorm.Pongo2(sqlDir, ".stpl")); err != nil {
		logrus.Fatal(err)
	}
	if err := db.RegisterSqlTemplate(xorm.Jet(sqlDir, ".jet")); err != nil {
		logrus.Fatal(err)
	}
	if err := db.RegisterSqlTemplate(xorm.Default(sqlDir, ".tpl")); err != nil {
		logrus.Fatal(err)
	}
}

// AddBusinessDB adddb
func (e *Engine) AddBusinessDB(domain string, db *xorm.Engine) {
	if e.BusinessDBSet[domain] != nil {
		panic(fmt.Errorf("domain %v has exited", domain))
	}
	e.BusinessDBSet[domain] = db
}

// InitRedis redis
func (e *Engine) initRedis() {
	uri, err := httpUtil.Parse(viper.GetString("rd.dataSource"))
	if err != nil {
		panic(err)
	}
	db, err := strconv.Atoi(uri.DbName)
	if err != nil {
		panic(err)
	}
	rds := redis.NewClient(&redis.Options{
		Addr:     uri.Laddr,
		Password: uri.Passwd,
		DB:       db,
	})
	if _, err := rds.Ping().Result(); err != nil {
		panic(err)
	}
	e.Redis = rds
}

func (e *Engine) initOAuth() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapTokenStorage(store.NewRedisStoreWithCli(e.Redis, TokenkeyNamespace))
	manager.MapAccessGenerate(generates.NewAccessGenerate())
	manager.MapClientStorage(NewClientStore())

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
		store.Save()
		return
	})
	e.OAuth2.SetInternalErrorHandler(func(err error) (re *oError.Response) {
		logrus.Error("internal error:", err.Error())
		return
	})
	e.OAuth2.SetResponseErrorHandler(func(re *oError.Response) {
		logrus.Error("response error:", re.Error.Error())
	})
}

// Auth middles
func (e *Engine) Auth(mode ...AuthType) func(ctx *Context) {
	return func(ctx *Context) {
		if ctx.Auth(ctx.Request) {
			fmt.Println("-----------", ctx.GetToken().GetDomain())
			ctx.DB = e.BusinessDBSet[ctx.GetToken().GetDomain()]
			if ctx.DB == nil {
				ctx.Fail(util.ErrInvalidDomain)
				ctx.Abort()
			} else {
				ctx.Set("DB", ctx.DB)
				ctx.Set("AuthInfo", ctx.AuthInfo)
				ctx.Next()
			}
		} else {
			ctx.Fail(util.ErrInvalidAccessToken, 401)
			ctx.Abort()
		}
	}
}

// RegisterHandler register handler
func (e *Engine) RegisterHandler(name string, h func(ctx *Context)) func(ctx *Context) {
	return h
}

// Run booting system
func (e *Engine) Run() {
	e.initRedis()
	e.initDB()
	e.initOAuth()
	e.MSet.Release()
}

// InvokeEngine build engine
func InvokeEngine(build func(*Engine)) func(*Engine) {
	return func(e *Engine) {
		build(e)
	}
}

// InvokeContext build context
func InvokeContext(httpMethod string, relativePath string, handlers ...HandlerFunc) func(*Engine) {
	return func(e *Engine) {
		group := e.Group(viper.GetString("http.prefix"))
		group.Handle(httpMethod, relativePath, handlers...)
	}
}

func buildEngine() *Engine {
	e := &Engine{}
	e.MSet = &MSet{m: map[string][]interface{}{}}
	e.BusinessDBSet = map[string]*xorm.Engine{}
	e.GRPC = grpc.NewServer()
	if viper.GetString("app.mode") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	e.Gin = gin.New()
	e.Gin.Static(viper.GetString("http.static"), path.Join(file.Getwd(), viper.GetString("http.static")))
	e.Gin.Use(plugin.Tracker(gin.LoggerConfig{
		Output:    logrus.GetOutput(),
		Formatter: plugin.Formatter,
	}, Tracker(e)))
	e.Gin.Use(plugin.Recovery())
	e.Gin.Use(plugin.ProcessMethodOverride(e.Gin))
	e.pool.New = func() interface{} {
		return e.allocateContext()
	}
	return e
}

// App defined application
var App = buildEngine()
