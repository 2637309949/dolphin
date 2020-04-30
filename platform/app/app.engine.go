package app

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/2637309949/dolphin/packages/oauth2"

	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/null"
	"google.golang.org/grpc"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/go-session/session"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/oauth2/errors"
	"github.com/2637309949/dolphin/packages/oauth2/generates"
	"github.com/2637309949/dolphin/packages/oauth2/manage"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/plugin"
	"github.com/2637309949/dolphin/platform/sql"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/file"

	httpUtil "github.com/2637309949/dolphin/platform/util/http"
)

type (
	// Manager Engine management interface
	Manager interface {
		GetMSet() MSeti
		GetBusinessDBSet() map[string]*xorm.Engine
		GetBusinessDB(string) *xorm.Engine
		AddBusinessDB(string, *xorm.Engine)
		GetTokenStore() oauth2.TokenStore
	}
	// Engine defined parse app engine
	Engine struct {
		Manager    Manager
		Gin        *gin.Engine
		GRPC       *grpc.Server
		OAuth2     *server.Server
		PlatformDB *xorm.Engine
		pool       sync.Pool
	}
)

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
	return &Context{PlatformDB: e.PlatformDB, engine: e, OAuth2: e.OAuth2, AuthInfo: &AuthOAuth2{server: e.OAuth2}}
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{engine: e, RouterGroup: e.Gin.Group(relativePath, handlers...)}
}

// Migration models
func (e *Engine) migration(name string, db *xorm.Engine) {
	e.Manager.GetMSet().ForEach(func(n string, m interface{}) {
		util.Ensure(db.Sync2(m))
	}, name)
}

func (e *Engine) initDB() {
	xLogger := xorm.NewSimpleLogger(logrus.StandardLogger().Out)
	xLogger.ShowSQL(viper.GetString("app.mode") == "debug")

	// initPlatformDB
	e.PlatformDB = util.EnsureLeft(xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dataSource"))).(*xorm.Engine)
	util.Ensure(e.PlatformDB.Ping())
	e.PlatformDB.SetLogger(xLogger)
	e.RegisterSQLDir(e.PlatformDB, path.Join(".", viper.GetString("dir.sql")))
	e.RegisterSQLMap(e.PlatformDB, sql.SQLTPL)
	(new(model.SysDomain)).Ensure(e.PlatformDB)
	(new(model.SysDomain)).InitSysData(e.PlatformDB.NewSession())

	// initBusinessDB
	domains := []model.SysDomain{}
	util.Ensure(e.PlatformDB.Where("data_source <> '' and domain <> '' and app_name = ? and del_flag = 0", viper.GetString("app.name")).Find(&domains))
	funk.ForEach(domains, func(domain model.SysDomain) {
		uri := util.EnsureLeft(httpUtil.Parse(domain.DataSource.String)).(*httpUtil.URI)
		util.EnsureLeft(e.PlatformDB.Sql(fmt.Sprintf("create database if not exists %v default character set utf8mb4", uri.DbName)).Execute())
		db := util.EnsureLeft(xorm.NewEngine(domain.DriverName.String, domain.DataSource.String)).(*xorm.Engine)
		db.SetLogger(xLogger)

		e.RegisterSQLDir(db, path.Join(".", viper.GetString("dir.sql")))
		e.RegisterSQLMap(db, sql.SQLTPL)
		e.Manager.AddBusinessDB(domain.Domain.String, db)
	})

	// migration db
	zmap := map[string]*xorm.Engine{}
	if util.EnsureLeft(e.PlatformDB.Where("sync_flag=0").Count(new(model.SysDomain))).(int64) > 0 {
		defer func() {
			util.EnsureLeft(e.PlatformDB.Where("sync_flag=0 and app_name = ?", viper.GetString("app.name")).Cols("sync_flag").Update(&model.SysDomain{SyncFlag: null.IntFrom(1)}))
			e.Manager.GetMSet().Release()
		}()
		nset := e.Manager.GetMSet().Name(func(n string) bool {
			return n != Name
		})
		for _, v := range e.Manager.GetBusinessDBSet() {
			zmap[v.DataSourceName()] = v
		}
		// migration PlatformDB
		e.migration(Name, e.PlatformDB)
		// migration BusinessDBSet
		funk.ForEach(nset, func(n string) {
			for _, v := range zmap {
				e.migration(n, v)
			}
		})
	}

	// initialize PlatformDB data
	(new(model.SysClient)).InitSysData(e.PlatformDB.NewSession())
	(new(model.SysUser)).InitSysData(e.PlatformDB.NewSession())

	// initialize BusinessDBSet data
	for _, v := range zmap {
		(new(model.SysRole)).InitSysData(v.NewSession())
		(new(model.SysRoleUser)).InitSysData(v.NewSession())
		(new(model.SysMenu)).InitSysData(v.NewSession())
		(new(model.SysOptionset)).InitSysData(v.NewSession())
	}
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
	util.Ensure(os.MkdirAll(sqlDir, os.ModePerm))
	util.Ensure(db.RegisterSqlMap(xorm.Xml(sqlDir, ".xml")))
	util.Ensure(db.RegisterSqlTemplate(xorm.Pongo2(sqlDir, ".stpl")))
	util.Ensure(db.RegisterSqlTemplate(xorm.Jet(sqlDir, ".jet")))
	util.Ensure(db.RegisterSqlTemplate(xorm.Default(sqlDir, ".tpl")))
}

func (e *Engine) initOAuth() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapTokenStorage(e.Manager.GetTokenStore())
	manager.MapAccessGenerate(generates.NewAccessGenerate())
	manager.MapClientStorage(NewClientStore())
	manager.SetValidateURIHandler(func(baseURI string, redirectURI string) error {
		reg := regexp.MustCompile("^(http://|https://)?([^/?:]+)(:[0-9]*)?(/[^?]*)?(\\?.*)?$")
		base := reg.FindAllStringSubmatch(baseURI, -1)
		redirect := reg.FindAllStringSubmatch(redirectURI, -1)
		if base[0][2] != redirect[0][2] {
			return errors.ErrInvalidRedirectURI
		}
		return nil
	})

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
	e.OAuth2.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logrus.Errorf("internal error:%v", err.Error())
		return
	})
	e.OAuth2.SetResponseErrorHandler(func(re *errors.Response) {
		logrus.Errorf("response error:%v", re.Error.Error())
	})
}

// Auth middles
func (e *Engine) Auth() func(ctx *Context) {
	return func(ctx *Context) {
		if !ctx.Auth(ctx.Request) {
			ctx.Fail(util.ErrInvalidAccessToken, 401)
			ctx.Abort()
			return
		}
		if ctx.DB = e.Manager.GetBusinessDB(ctx.GetToken().GetDomain()); ctx.DB == nil {
			ctx.Fail(util.ErrInvalidDomain)
			ctx.Abort()
			return
		}
		ctx.Set("DB", ctx.DB)
		ctx.Set("AuthInfo", ctx.AuthInfo)
		ctx.Next()
	}
}

// Run booting system
func (e *Engine) Run() {
	e.initDB()
	e.initOAuth()
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
		e.Group(viper.GetString("http.prefix")).Handle(httpMethod, relativePath, handlers...)
	}
}

func buildEngine() *Engine {
	e := &Engine{}
	e.Manager = NewDefaultManager()
	e.GRPC = grpc.NewServer()
	if viper.GetString("app.mode") != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	e.Gin = gin.New()
	e.Gin.Use(plugin.CORS())
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
