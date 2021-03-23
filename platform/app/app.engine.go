// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/2637309949/dolphin/packages/gin"
	"github.com/2637309949/dolphin/packages/go-funk"
	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/oauth2/errors"
	"github.com/2637309949/dolphin/packages/oauth2/generates"
	"github.com/2637309949/dolphin/packages/oauth2/manage"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/packages/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/schemas"
	"github.com/2637309949/dolphin/platform/model"
	"github.com/2637309949/dolphin/platform/plugin"
	"github.com/2637309949/dolphin/platform/sql"
	"github.com/2637309949/dolphin/platform/util"
	"github.com/2637309949/dolphin/platform/util/http"
	"github.com/2637309949/dolphin/platform/util/slice"
	"google.golang.org/grpc"
)

// Engine defined parse app engine
type Engine struct {
	Manager    Manager
	lifecycle  Lifecycle
	Gin        *gin.Engine
	GRPC       *grpc.Server
	OAuth2     *server.Server
	PlatformDB *xorm.Engine
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

func (e *Engine) allocateContext() *Context {
	return &Context{PlatformDB: e.PlatformDB, engine: e, OAuth2: e.OAuth2, AuthInfo: &AuthOAuth2{server: e.OAuth2}}
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{engine: e, RouterGroup: e.Gin.Group(relativePath, handlers...)}
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
		columns = append(columns, funk.Map(colsInfo, func(col *schemas.Column) model.SysTableColumn {
			sc := new(model.SysTableColumn).ColumnInfo(col)
			sc.TbId = tables[len(tables)-1].ID
			return sc
		}).([]model.SysTableColumn)...)
	}, name)
	util.EnsureLeft(session.Exec(fmt.Sprintf("truncate table %v", new(model.SysTable).TableName())))
	util.EnsureLeft(session.Exec(fmt.Sprintf("truncate table %v", new(model.SysTableColumn).TableName())))
	stb, stc := slice.Chunk(tables, 50).([][]model.SysTable), slice.Chunk(columns, 50).([][]model.SysTableColumn)
	for i := range stb {
		util.EnsureLeft(session.Insert(stb[i]))
	}
	for i := range stc {
		util.EnsureLeft(session.Insert(stc[i]))

	}
	session.Commit()
}

func (e *Engine) database() {
	// initPlatformDB
	xlogger := createXLogger()
	logrus.Infoln(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	e.PlatformDB = util.EnsureLeft(xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dataSource"))).(*xorm.Engine)
	util.Ensure(e.PlatformDB.Ping())
	e.PlatformDB.SetLogger(xlogger)
	e.PlatformDB.SetConnMaxLifetime(time.Duration(viper.GetInt("db.connMaxLifetime")) * time.Minute)
	e.PlatformDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	e.PlatformDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	e.RegisterSQLDir(e.PlatformDB, path.Join(".", viper.GetString("dir.sql")))
	e.RegisterSQLMap(e.PlatformDB, sql.SQLTPL)
	{
		(new(model.SysDomain)).Ensure(e.PlatformDB)
		(new(model.SysDomain)).InitSysData(e.PlatformDB.NewSession())
	}

	// initBusinessDB
	domains := []model.SysDomain{}
	asyncOnce := sync.Once{}
	util.Ensure(e.PlatformDB.Where("data_source <> '' and domain <> '' and app_name = ? and is_delete != 1", viper.GetString("app.name")).Find(&domains))
	funk.ForEach(domains, func(domain model.SysDomain) {
		logrus.Infoln(domain.DriverName.String, domain.DataSource.String)
		uri := util.EnsureLeft(http.Parse(domain.DataSource.String)).(*http.URI)
		util.EnsureLeft(e.PlatformDB.SQL(fmt.Sprintf("create database if not exists %v default character set utf8mb4", uri.DbName)).Execute())
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
	funk.ForEach(funk.Filter(domains, func(domain model.SysDomain) bool { return domain.SyncFlag.Int64 != 1 }), func(domain model.SysDomain) {
		// obtain BusinessDB
		db := e.Manager.GetBusinessDB(domain.Domain.String)
		// migration PlatformDB
		asyncOnce.Do(func() {
			e.migration(Name, e.PlatformDB)
			new(model.SysClient).InitSysData(e.PlatformDB.NewSession())
			new(model.SysUser).InitSysData(e.PlatformDB.NewSession())
		})
		// migration BusinessDBSet
		funk.ForEach(platBindModelNames, func(n string) {
			e.migration(n, db)
		})
		// initialize BusinessDBSet data
		new(model.SysRole).InitSysData(db.NewSession())
		new(model.SysOrg).InitSysData(db.NewSession())
		new(model.SysRoleUser).InitSysData(db.NewSession())
		new(model.SysMenu).InitSysData(db.NewSession())
		new(model.SysOptionset).InitSysData(db.NewSession())
		new(model.SysUserTemplate).InitSysData(db.NewSession())
		new(model.SysUserTemplateDetail).InitSysData(db.NewSession())
		// ensure sync_flag
		util.EnsureLeft(e.PlatformDB.ID(domain.ID.String).Cols("sync_flag").Update(&model.SysDomain{SyncFlag: null.IntFrom(1)}))
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

func (e *Engine) authorize() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapTokenStorage(e.Manager.GetTokenStore())
	manager.MapAccessGenerate(generates.NewAccessGenerate())
	manager.MapClientStorage(NewClientStore())
	manager.SetValidateURIHandler(ValidateURIHandler)
	e.OAuth2 = server.NewServer(server.NewConfig(), manager)
	e.OAuth2.SetUserAuthorizationHandler(UserAuthorizationHandler)
	e.OAuth2.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logrus.Error(err)
		return
	})
	e.OAuth2.SetResponseErrorHandler(func(re *errors.Response) {
		logrus.Error(re.Error)
	})
}

// Done returns a channel of signals to block on after starting the
func (e *Engine) done() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}

// lifeCycle start liftcycle hooks
func (e *Engine) lifeCycle() {
	signal := e.done()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := e.lifecycle.Start(ctx); err != nil {
		logrus.Fatal(err)
	}
	<-signal
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := e.lifecycle.Stop(ctx); err != nil {
		logrus.Fatal(err)
	}
}

// Run booting system
func (e *Engine) Run() {
	e.database()
	e.authorize()
	e.lifeCycle()
}

func buildEngine() *Engine {
	e := &Engine{}
	e.Manager = NewDefaultManager()
	e.lifecycle = &lifecycleWrapper{}
	e.GRPC = grpc.NewServer()
	e.Gin = NewGin()
	e.Gin.Use(plugin.Tracker(Tracker(e)))
	e.pool.New = func() interface{} {
		return e.allocateContext()
	}
	e.lifecycle.Append(NewLifeHook(e))
	return e
}

// Run app
func Run() {
	App.Run()
}

// App defined application
var App = buildEngine()
