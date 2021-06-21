// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"context"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/oauth2/errors"
	"github.com/2637309949/dolphin/packages/oauth2/generates"
	"github.com/2637309949/dolphin/packages/oauth2/manage"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/schemas"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/sql"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/http"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

type (
	// HandlerFunc defined
	HandlerFunc func(ctx *Context)
	// HandlersChain defined
	HandlersChain []HandlerFunc
	// RouterGroup defines
	RouterGroup struct {
		engine   *Engine
		Handlers []HandlerFunc
		basePath string
	}
	// Engine defined parse app engine
	Engine struct {
		RouterGroup
		PlatformDB *xorm.Engine
		lifecycle  Lifecycle
		Manager    Manager
		OAuth2     *server.Server
		Http       HttpHandler
		RPC        RPCHandler
		pool       sync.Pool
	}
)

// Handle overwrite RouterGroup.Handle
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlerFuncs ...HandlerFunc) {
	for i, methods := 0, strings.Split(httpMethod, ","); i < len(methods); i++ {
		method := methods[i]
		absPath := path.Join(group.basePath, relativePath)
		group.engine.Http.Handle(method, absPath, handlerFuncs...)
	}
}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	if finalSize >= int(63) {
		panic("too many handlers")
	}
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	if relativePath == "" {
		return group.basePath
	}

	finalPath := path.Join(group.basePath, relativePath)
	if util.LastChar(relativePath) == '/' && util.LastChar(finalPath) != '/' {
		return finalPath + "/"
	}
	return finalPath
}

func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(relativePath),
		engine:   group.engine,
	}
}

// HandlerFunc convert to gin.HandlerFunc
func (e *Engine) HandlerFunc(h HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		c := e.pool.Get().(*Context)
		c.reset()
		c.Context = ctx
		for k := range ctx.Keys {
			switch t := ctx.Keys[k].(type) {
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

// allocateContext defined new context
func (e *Engine) allocateContext() *Context {
	return &Context{PlatformDB: e.PlatformDB, AuthInfo: &AuthOAuth2{server: e.OAuth2}}
}

// Migration models
func (e *Engine) migration(name string, db *xorm.Engine) error {
	tables := []model.SysTable{}
	columns := []model.SysTableColumn{}
	session := db.NewSession()
	session.Begin()
	defer session.Close()
	defer func() {
		if err := recover(); err != nil {
			session.Rollback()
			panic(err)
		}
	}()
	err := db.Sync2(new(model.SysTable), new(model.SysTableColumn))
	if err != nil {
		return err
	}
	err = e.Manager.MSet().ForEach(func(n string, m interface{}) error {
		if viper.GetString("app.mode") == "debug" {
			if tbn, ok := m.(interface{ TableName() string }); ok {
				logrus.Infof("Sync Model[%v]:%v", n, tbn.TableName())
			}
		}
		err = db.Sync2(m)
		if err != nil {
			return err
		}
		tableInfo, err := db.TableInfo(m)
		if err != nil {
			return err
		}
		colsInfo := tableInfo.Columns()
		tables = append(tables, new(model.SysTable).TableInfo(tableInfo))
		tables[len(tables)-1].ID = null.IntFrom(int64(len(tables)))
		columns = append(columns, funk.Map(colsInfo, func(col *schemas.Column) model.SysTableColumn {
			sc := new(model.SysTableColumn).ColumnInfo(col)
			sc.TbId = tables[len(tables)-1].ID
			return sc
		}).([]model.SysTableColumn)...)
		return nil
	}, name)
	if err != nil {
		return err
	}
	new(model.SysTable).TruncateTable(session, session.Engine().DriverName())
	new(model.SysTableColumn).TruncateTable(session, session.Engine().DriverName())
	stb, stc := slice.Chunk(tables, 50).([][]model.SysTable), slice.Chunk(columns, 50).([][]model.SysTableColumn)
	for i := range stb {
		_, err = session.Insert(stb[i])
		if err != nil {
			session.Rollback()
			return err
		}
	}
	for i := range stc {
		session.Insert(stc[i])
		if err != nil {
			session.Rollback()
			return err
		}
	}
	return session.Commit()
}

// Reflesh defined init data before bootinh
func (e *Engine) Reflesh() error {
	logrus.Infoln(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	xlogger := createXLogger()
	db, err := xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	if err != nil {
		return err
	}
	e.PlatformDB = db
	err = e.PlatformDB.Ping()
	if err != nil {
		return err
	}
	e.PlatformDB.SetLogger(xlogger)
	e.PlatformDB.SetConnMaxLifetime(time.Duration(viper.GetInt("db.connMaxLifetime")) * time.Minute)
	e.PlatformDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	e.PlatformDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	err = e.RegisterSQLDir(e.PlatformDB, path.Join(".", viper.GetString("dir.sql")))
	if err != nil {
		return err
	}
	err = e.RegisterSQLMap(e.PlatformDB, sql.SQLTPL)
	if err != nil {
		return err
	}
	new(model.SysDomain).Ensure(e.PlatformDB)
	new(model.SysDomain).InitSysData(e.PlatformDB.NewSession())

	asyncOnce, domains := sync.Once{}, []model.SysDomain{}
	err = e.PlatformDB.Where("data_source <> '' and domain <> '' and app_name = ? and is_delete != 1", viper.GetString("app.name")).Find(&domains)
	if err != nil {
		return err
	}

	for i := range domains {
		domain := domains[i]
		logrus.Infoln(domain.DriverName.String, domain.DataSource.String)
		uri, err := http.Parse(domain.DataSource.String)
		if err != nil {
			return err
		}
		err = domain.CreateDataBase(e.PlatformDB, domain.DriverName.String, uri.DbName)
		if err != nil {
			return err
		}
		db, err := xorm.NewEngine(domain.DriverName.String, domain.DataSource.String)
		if err != nil {
			return err
		}
		db.SetLogger(xlogger)
		db.SetConnMaxLifetime(time.Duration(viper.GetInt("db.connMaxLifetime")) * time.Minute)
		db.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
		db.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
		err = e.RegisterSQLDir(db, path.Join(".", viper.GetString("dir.sql")))
		if err != nil {
			return err
		}
		err = e.RegisterSQLMap(db, sql.SQLTPL)
		if err != nil {
			return err
		}
		e.Manager.AddBusinessDB(domain.Domain.String, db)
	}

	platBindModelNames := e.Manager.MSet().Name(func(name string) bool { return name != Name })
	filtedDomains := funk.Filter(domains, func(domain model.SysDomain) bool { return domain.IsSync.Int64 != 1 }).([]model.SysDomain)
	for i := range filtedDomains {
		domain := filtedDomains[i]
		db := e.Manager.GetBusinessDB(domain.Domain.String)
		asyncOnce.Do(func() {
			e.migration(Name, e.PlatformDB)
			new(model.SysClient).InitSysData(e.PlatformDB.NewSession())
			new(model.SysUser).InitSysData(e.PlatformDB.NewSession())
		})
		funk.ForEach(platBindModelNames, func(n string) { e.migration(n, db) })
		new(model.SysRole).InitSysData(db.NewSession())
		new(model.SysOrg).InitSysData(db.NewSession())
		new(model.SysRoleUser).InitSysData(db.NewSession())
		new(model.SysMenu).InitSysData(db.NewSession())
		new(model.SysOptionset).InitSysData(db.NewSession())
		new(model.SysUserTemplate).InitSysData(db.NewSession())
		new(model.SysUserTemplateDetail).InitSysData(db.NewSession())
		_, err = e.PlatformDB.ID(domain.ID.Int64).Cols("is_sync").Update(&model.SysDomain{IsSync: null.IntFrom(1)})
		if err != nil {
			return err
		}
	}
	e.Manager.MSet().Release()
	return err
}

// RegisterSQLMap defined sql
func (e *Engine) RegisterSQLMap(db *xorm.Engine, sqlMap map[string]string) error {
	for key, value := range sqlMap {
		if filepath.Ext(key) == "" {
			db.AddSql(key, value)
		} else {
			err := db.AddSqlTemplate(key, value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// RegisterSQLDir defined sql
func (e *Engine) RegisterSQLDir(db *xorm.Engine, sqlDir string) error {
	err := os.MkdirAll(sqlDir, os.ModePerm)
	if err != nil {
		return err
	}
	err = db.RegisterSqlMap(xorm.Xml(sqlDir, ".xml"))
	if err != nil {
		return err
	}
	err = db.RegisterSqlTemplate(xorm.Pongo2(sqlDir, ".stpl"))
	if err != nil {
		return err
	}
	err = db.RegisterSqlTemplate(xorm.Jet(sqlDir, ".jet"))
	if err != nil {
		return err
	}
	err = db.RegisterSqlTemplate(xorm.Default(sqlDir, ".tpl"))
	return err
}

// Done returns a channel of signals to block on after starting the
func (e *Engine) done() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}

// lifeCycle start liftcycle hooks
func (e *Engine) lifeCycle(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	signal := e.done()
	if err := e.lifecycle.Start(ctx); err != nil {
		return err
	}
	<-signal
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	if err := e.lifecycle.Stop(ctx); err != nil {
		return err
	}
	return nil
}

// Run booting system
func (e *Engine) Run() {
	util.Ensure(e.Reflesh())
	util.Ensure(e.lifeCycle(context.Background()))
}

func NewEngine() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
		},
	}
	engine.RouterGroup.engine = engine
	engine.Manager = NewDefaultManager()
	engine.lifecycle = &lifecycleWrapper{}
	engine.Http = NewGinHandler(engine)
	engine.RPC = NewGRPCHandler(engine)
	engine.pool.New = func() interface{} { return engine.allocateContext() }
	engine.lifecycle.Append(NewLifeHook(engine))

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapTokenStorage(engine.Manager.GetTokenStore())
	manager.MapAccessGenerate(generates.NewAccessGenerate())
	manager.MapClientStorage(NewClientStore())
	manager.SetValidateURIHandler(ValidateURIHandler)
	engine.OAuth2 = server.NewServer(server.NewConfig(), manager)
	engine.OAuth2.SetUserAuthorizationHandler(UserAuthorizationHandler)
	engine.OAuth2.SetInternalErrorHandler(func(err error) (re *errors.Response) { logrus.Error(err); return })
	engine.OAuth2.SetResponseErrorHandler(func(re *errors.Response) { logrus.Error(re.Error) })
	return engine
}

var (
	// App defined
	App = NewEngine()
	// Run defined
	Run = App.Run
)
