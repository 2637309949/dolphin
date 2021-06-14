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

// Engine defined parse app engine
type Engine struct {
	PlatformDB *xorm.Engine
	lifecycle  Lifecycle
	Manager    Manager
	OAuth2     *server.Server
	Http       HttpHandler
	RPC        RPCHandler
	pool       sync.Pool
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

// Group handlers
func (e *Engine) Group(relativePath string, routes ...Route) *RouterGroup {
	return &RouterGroup{Routes: routes, basePath: relativePath, engine: e}
}

// Migration models
func (e *Engine) migration(name string, db *xorm.Engine) {
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
	util.Ensure(db.Sync2(new(model.SysTable), new(model.SysTableColumn)))
	e.Manager.MSet().ForEach(func(n string, m interface{}) {
		if mode := viper.GetString("app.mode"); mode == "debug" {
			if tbn, ok := m.(interface{ TableName() string }); ok {
				logrus.Infof("Sync Model[%v]:%v", n, tbn.TableName())
			}
		}
		util.Ensure(db.Sync2(m))
		tableInfo := util.EnsureLeft(db.TableInfo(m)).(*schemas.Table)
		colsInfo := tableInfo.Columns()
		tables = append(tables, new(model.SysTable).TableInfo(tableInfo))
		tables[len(tables)-1].ID = null.IntFrom(int64(len(tables)))
		columns = append(columns, funk.Map(colsInfo, func(col *schemas.Column) model.SysTableColumn {
			sc := new(model.SysTableColumn).ColumnInfo(col)
			sc.TbId = tables[len(tables)-1].ID
			return sc
		}).([]model.SysTableColumn)...)
	}, name)
	new(model.SysTable).TruncateTable(session, session.Engine().DriverName())
	new(model.SysTableColumn).TruncateTable(session, session.Engine().DriverName())
	stb, stc := slice.Chunk(tables, 50).([][]model.SysTable), slice.Chunk(columns, 50).([][]model.SysTableColumn)
	for i := range stb {
		util.EnsureLeft(session.Insert(stb[i]))
	}
	for i := range stc {
		util.EnsureLeft(session.Insert(stc[i]))

	}
	session.Commit()
}

// Reflesh defined init data before bootinh
func (e *Engine) Reflesh() {
	// initPlatformDB
	logrus.Infoln(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	xlogger := createXLogger()
	e.PlatformDB = util.EnsureLeft(xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dataSource"))).(*xorm.Engine)
	util.Ensure(e.PlatformDB.Ping())
	e.PlatformDB.SetLogger(xlogger)
	e.PlatformDB.SetConnMaxLifetime(time.Duration(viper.GetInt("db.connMaxLifetime")) * time.Minute)
	e.PlatformDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	e.PlatformDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	e.RegisterSQLDir(e.PlatformDB, path.Join(".", viper.GetString("dir.sql")))
	e.RegisterSQLMap(e.PlatformDB, sql.SQLTPL)
	new(model.SysDomain).Ensure(e.PlatformDB)
	new(model.SysDomain).InitSysData(e.PlatformDB.NewSession())

	// initBusinessDB
	asyncOnce, domains := sync.Once{}, []model.SysDomain{}
	util.Ensure(e.PlatformDB.Where("data_source <> '' and domain <> '' and app_name = ? and is_delete != 1", viper.GetString("app.name")).Find(&domains))
	funk.ForEach(domains, func(domain model.SysDomain) {
		logrus.Infoln(domain.DriverName.String, domain.DataSource.String)
		uri := util.EnsureLeft(http.Parse(domain.DataSource.String)).(*http.URI)
		domain.CreateDataBase(e.PlatformDB, domain.DriverName.String, uri.DbName)
		db := util.EnsureLeft(xorm.NewEngine(domain.DriverName.String, domain.DataSource.String)).(*xorm.Engine)
		db.SetLogger(xlogger)
		db.SetConnMaxLifetime(time.Duration(viper.GetInt("db.connMaxLifetime")) * time.Minute)
		db.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
		db.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
		e.RegisterSQLDir(db, path.Join(".", viper.GetString("dir.sql")))
		e.RegisterSQLMap(db, sql.SQLTPL)
		e.Manager.AddBusinessDB(domain.Domain.String, db)
	})

	// fetch all not bind in PlatformDB
	platBindModelNames := e.Manager.MSet().Name(func(name string) bool { return name != Name })
	funk.ForEach(funk.Filter(domains, func(domain model.SysDomain) bool { return domain.IsSync.Int64 != 1 }), func(domain model.SysDomain) {
		// obtain BusinessDB
		db := e.Manager.GetBusinessDB(domain.Domain.String)
		// migration PlatformDB, ensure domain.IsSync.Int64 != 1
		asyncOnce.Do(func() {
			e.migration(Name, e.PlatformDB)
			new(model.SysClient).InitSysData(e.PlatformDB.NewSession())
			new(model.SysUser).InitSysData(e.PlatformDB.NewSession())
		})
		// migration BusinessDBSet
		funk.ForEach(platBindModelNames, func(n string) { e.migration(n, db) })
		// initialize BusinessDBSet data
		new(model.SysRole).InitSysData(db.NewSession())
		new(model.SysOrg).InitSysData(db.NewSession())
		new(model.SysRoleUser).InitSysData(db.NewSession())
		new(model.SysMenu).InitSysData(db.NewSession())
		new(model.SysOptionset).InitSysData(db.NewSession())
		new(model.SysUserTemplate).InitSysData(db.NewSession())
		new(model.SysUserTemplateDetail).InitSysData(db.NewSession())
		// ensure is_sync
		util.EnsureLeft(e.PlatformDB.ID(domain.ID.Int64).Cols("is_sync").Update(&model.SysDomain{IsSync: null.IntFrom(1)}))
	})
	// release model sets
	e.Manager.MSet().Release()
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

// Done returns a channel of signals to block on after starting the
func (e *Engine) done() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}

// lifeCycle start liftcycle hooks
func (e *Engine) lifeCycle() error {
	signal := e.done()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := e.lifecycle.Start(ctx); err != nil {
		return err
	}
	<-signal
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := e.lifecycle.Stop(ctx); err != nil {
		return err
	}
	return nil
}

// Run booting system
func (e *Engine) Run() error {
	e.Reflesh()
	return e.lifeCycle()
}

func NewEngine() *Engine {
	e := &Engine{}
	e.Manager = NewDefaultManager()
	e.lifecycle = &lifecycleWrapper{}
	e.Http = NewGinHandler(e)
	e.RPC = NewGRPCHandler(e)
	e.pool.New = func() interface{} { return e.allocateContext() }
	e.lifecycle.Append(NewLifeHook(e))

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapTokenStorage(e.Manager.GetTokenStore())
	manager.MapAccessGenerate(generates.NewAccessGenerate())
	manager.MapClientStorage(NewClientStore())
	manager.SetValidateURIHandler(ValidateURIHandler)
	e.OAuth2 = server.NewServer(server.NewConfig(), manager)
	e.OAuth2.SetUserAuthorizationHandler(UserAuthorizationHandler)
	e.OAuth2.SetInternalErrorHandler(func(err error) (re *errors.Response) { logrus.Error(err); return })
	e.OAuth2.SetResponseErrorHandler(func(re *errors.Response) { logrus.Error(re.Error) })
	return e
}

var (
	// App defined
	App = NewEngine()
	// Run defined
	Run = App.Run
)
