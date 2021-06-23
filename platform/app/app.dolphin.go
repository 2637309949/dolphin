// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package app

import (
	"context"
	nhttp "net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"regexp"
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
	"github.com/2637309949/dolphin/platform/util/file"
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
		dol      *Dolphin
		Handlers []HandlerFunc
		basePath string
	}
	// Dolphin defined parse app engine
	Dolphin struct {
		RouterGroup
		PlatformDB *xorm.Engine
		lifecycle  Lifecycle
		Manager    Manager
		OAuth2     *server.Server
		Http       HttpHandler
		RPC        RPCHandler
		pool       sync.Pool
	}
	onlyFilesFS struct {
		fs nhttp.FileSystem
	}
	neuteredReaddirFile struct {
		nhttp.File
	}
)

func Dir(root string, listDirectory bool) nhttp.FileSystem {
	fs := nhttp.Dir(root)
	if listDirectory {
		return fs
	}
	return &onlyFilesFS{fs}
}

// Open defined TODO
func (fs onlyFilesFS) Open(name string) (nhttp.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	return neuteredReaddirFile{f}, nil
}

// Readdir defined TODO
func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

// Handle overwrite RouterGroup.Handle
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlerFuncs ...HandlerFunc) {
	for i, methods := 0, strings.Split(httpMethod, ","); i < len(methods); i++ {
		method := methods[i]
		re, err := regexp.Compile("^[A-Z]+$")
		if matches := re.MatchString(method); !matches || err != nil {
			panic("http method " + method + " is not valid")
		}
		group.handle(method, relativePath, handlerFuncs...)
	}
}

// Static defined TODO
func (group *RouterGroup) Static(relativePath, root string) {
	group.StaticFS(relativePath, Dir(root, false))
}

// StaticFS defined TODO
func (group *RouterGroup) StaticFS(relativePath string, fs nhttp.FileSystem) {
	if strings.Contains(relativePath, ":") || strings.Contains(relativePath, "*") {
		panic("URL parameters can not be used when serving a static folder")
	}
	handler := group.createStaticHandler(relativePath, fs)
	urlPattern := path.Join(relativePath, "/*filepath")

	group.Handle("GET", urlPattern, handler)
	group.Handle("HEAD", urlPattern, handler)
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs nhttp.FileSystem) HandlerFunc {
	absolutePath := group.calculateAbsolutePath(relativePath)
	fileServer := nhttp.StripPrefix(absolutePath, nhttp.FileServer(fs))
	return func(c *Context) {
		if _, noListing := fs.(*onlyFilesFS); noListing {
			c.Writer.WriteHeader(nhttp.StatusNotFound)
		}
		file := c.Param("filepath")
		f, err := fs.Open(file)
		if err != nil {
			c.Writer.WriteHeader(nhttp.StatusNotFound)
			return
		}
		f.Close()
		fileServer.ServeHTTP(c.Writer, c.Request)
	}
}

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	group.dol.Http.Handle(httpMethod, absolutePath, handlers...)
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

// Use defined TODO
func (group *RouterGroup) Use(middleware ...HandlerFunc) {
	group.Handlers = append(group.Handlers, middleware...)
}

// Group defined TODO
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(relativePath),
		dol:      group.dol,
	}
}

// HandlerFunc convert to gin.HandlerFunc
func (dol *Dolphin) HandlerFunc(h HandlerFunc) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		c := dol.pool.Get().(*Context)
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
		dol.pool.Put(c)
	})
}

// allocateContext defined new context
func (dol *Dolphin) allocateContext() *Context {
	return &Context{PlatformDB: dol.PlatformDB, AuthInfo: &AuthOAuth2{server: dol.OAuth2}}
}

// Migration models
func (dol *Dolphin) migration(name string, db *xorm.Engine) error {
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
	err = dol.Manager.MSet().ForEach(func(n string, m interface{}) error {
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
func (dol *Dolphin) Reflesh() error {
	logrus.Infoln(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	xlogger := createXLogger()
	db, err := xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	if err != nil {
		return err
	}
	dol.PlatformDB = db
	err = dol.PlatformDB.Ping()
	if err != nil {
		return err
	}
	dol.PlatformDB.SetLogger(xlogger)
	dol.PlatformDB.SetConnMaxLifetime(time.Duration(viper.GetInt("db.connMaxLifetime")) * time.Minute)
	dol.PlatformDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	dol.PlatformDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))
	err = dol.RegisterSQLDir(dol.PlatformDB, path.Join(".", viper.GetString("dir.sql")))
	if err != nil {
		return err
	}
	err = dol.RegisterSQLMap(dol.PlatformDB, sql.SQLTPL)
	if err != nil {
		return err
	}
	new(model.SysDomain).Ensure(dol.PlatformDB)
	new(model.SysDomain).InitSysData(dol.PlatformDB.NewSession())

	asyncOnce, domains := sync.Once{}, []model.SysDomain{}
	err = dol.PlatformDB.Where("data_source <> '' and domain <> '' and app_name = ? and is_delete != 1", viper.GetString("app.name")).Find(&domains)
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
		err = domain.CreateDataBase(dol.PlatformDB, domain.DriverName.String, uri.DbName)
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
		err = dol.RegisterSQLDir(db, path.Join(".", viper.GetString("dir.sql")))
		if err != nil {
			return err
		}
		err = dol.RegisterSQLMap(db, sql.SQLTPL)
		if err != nil {
			return err
		}
		dol.Manager.AddBusinessDB(domain.Domain.String, db)
	}

	platBindModelNames := dol.Manager.MSet().Name(func(name string) bool { return name != Name })
	filtedDomains := funk.Filter(domains, func(domain model.SysDomain) bool { return domain.IsSync.Int64 != 1 }).([]model.SysDomain)
	for i := range filtedDomains {
		domain := filtedDomains[i]
		db := dol.Manager.GetBusinessDB(domain.Domain.String)
		asyncOnce.Do(func() {
			dol.migration(Name, dol.PlatformDB)
			new(model.SysClient).InitSysData(dol.PlatformDB.NewSession())
			new(model.SysUser).InitSysData(dol.PlatformDB.NewSession())
		})
		funk.ForEach(platBindModelNames, func(n string) { dol.migration(n, db) })
		new(model.SysRole).InitSysData(db.NewSession())
		new(model.SysOrg).InitSysData(db.NewSession())
		new(model.SysRoleUser).InitSysData(db.NewSession())
		new(model.SysMenu).InitSysData(db.NewSession())
		new(model.SysOptionset).InitSysData(db.NewSession())
		new(model.SysUserTemplate).InitSysData(db.NewSession())
		new(model.SysUserTemplateDetail).InitSysData(db.NewSession())
		_, err = dol.PlatformDB.ID(domain.ID.Int64).Cols("is_sync").Update(&model.SysDomain{IsSync: null.IntFrom(1)})
		if err != nil {
			return err
		}
	}
	dol.Manager.MSet().Release()
	return err
}

// RegisterSQLMap defined sql
func (dol *Dolphin) RegisterSQLMap(db *xorm.Engine, sqlMap map[string]string) error {
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
func (dol *Dolphin) RegisterSQLDir(db *xorm.Engine, sqlDir string) error {
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
func (dol *Dolphin) done() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}

// lifeCycle start liftcycle hooks
func (dol *Dolphin) lifeCycle(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	signal := dol.done()
	if err := dol.lifecycle.Start(ctx); err != nil {
		return err
	}
	<-signal
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	if err := dol.lifecycle.Stop(ctx); err != nil {
		return err
	}
	return nil
}

// Run booting system
func (dol *Dolphin) Run() {
	util.Ensure(dol.Reflesh())
	util.Ensure(dol.lifeCycle(context.Background()))
}

func NewDolphin() *Dolphin {
	dol := &Dolphin{
		RouterGroup: RouterGroup{
			Handlers: nil,
			basePath: "/",
		},
	}
	dol.RouterGroup.dol = dol
	dol.Manager = NewDefaultManager()
	dol.lifecycle = &lifecycleWrapper{}
	dol.Http = NewGinHandler(dol)
	dol.RPC = NewGRPCHandler(dol)
	dol.pool.New = func() interface{} { return dol.allocateContext() }
	dol.lifecycle.Append(NewLifeHook(dol))

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MapTokenStorage(dol.Manager.GetTokenStore())
	manager.MapAccessGenerate(generates.NewAccessGenerate())
	manager.MapClientStorage(NewClientStore())
	manager.SetValidateURIHandler(ValidateURIHandler)
	dol.OAuth2 = server.NewServer(server.NewConfig(), manager)
	dol.OAuth2.SetUserAuthorizationHandler(UserAuthorizationHandler)
	dol.OAuth2.SetInternalErrorHandler(func(err error) (re *errors.Response) { logrus.Error(err); return })
	dol.OAuth2.SetResponseErrorHandler(func(re *errors.Response) { logrus.Error(re.Error) })

	dol.Use(Recovery())
	dol.Use(HttpTrace())
	dol.Use(Cors())
	dol.Static(viper.GetString("http.static"), path.Join(file.Getwd(), viper.GetString("http.static")))
	dol.Use(Tracker(TrackerOpts(dol)))
	return dol
}

var (
	// App defined
	App = NewDolphin()
	// Run defined
	Run = App.Run
)
