// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/web/core"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/schemas"
	"github.com/2637309949/dolphin/platform/sql"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/file"
	"github.com/2637309949/dolphin/platform/util/http/uri"
	"github.com/2637309949/dolphin/platform/util/slice"
	"github.com/2637309949/dolphin/platform/util/db"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

// Option is a functional option type for Dolphin objects.
type Option func(*Dolphin)

// Dolphin defined parse app engine
type Dolphin struct {
	Lifecycle
	PlatformDB *xorm.Engine
	Manager    Manager
	Identity   *Identity
	Web        *core.Web
}

// migration defined TODO
func (dol *Dolphin) migration(name string, db *xorm.Engine) error {
	tables, columns := []types.SysTable{}, []types.SysTableColumn{}
	err := db.Sync2(new(types.SysTable), new(types.SysTableColumn))
	if err != nil {
		return err
	}

	models := dol.Manager.ModelSet().ByName(name)
	for i := range models {
		if IsDebugging() {
			logrus.Infof("Sync Model[%v]:%v", name, models[i].TableName())
		}
		err = db.Sync2(models[i])
		if err != nil {
			return err
		}
		tableInfo, err := db.TableInfo(models[i])
		if err != nil {
			return err
		}
		colsInfo := tableInfo.Columns()
		tables = append(tables, new(types.SysTable).TableInfo(tableInfo))
		tables[len(tables)-1].ID = null.IntFrom(int64(len(tables)))
		columns = append(columns, funk.Map(colsInfo, func(col *schemas.Column) types.SysTableColumn {
			sc := new(types.SysTableColumn).ColumnInfo(col)
			sc.TbId = tables[len(tables)-1].ID
			return sc
		}).([]types.SysTableColumn)...)
	}

	dol.truncateTable(db.NewSession(), new(types.SysTable), new(types.SysTableColumn))
	for _, v := range []interface{}{tables, columns} {
		chunk := slice.Chunk(v, 50)
		switch ck := chunk.(type) {
		case  [][]types.SysTable :
			for i := range ck {
				_, err = db.Insert(ck[i])
				if err != nil {
					return err
				}
			}
		case [][]types.SysTableColumn:
			for i := range ck {
				_, err = db.Insert(ck[i])
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// initSysData defined TODO
func (dol *Dolphin) initSysData(s *xorm.Session, bean ...interface{InitSysData(*xorm.Session)error}) error {
	for i := range bean {
		err := bean[i].InitSysData(s)
		if err != nil {
			return err
		}
	}
	return nil
}

// truncateTable defined TODO
func (dol *Dolphin) truncateTable(s *xorm.Session, bean ...interface{TruncateTable(*xorm.Session)error}) error {
	for i := range bean {
		err := bean[i].TruncateTable(s)
		if err != nil {
			return err
		}
	}
	return nil
}

// initDomain defined TODO
func (dol *Dolphin) initDomain(domain types.SysDomain) error {
	platBindModelNames := dol.Manager.ModelSet().NameSpaces(Name)
	db := dol.Manager.GetBusinessDB(domain.Domain.String)
	s1 := db.NewSession()
	s1.Begin()
	defer s1.Close()

	funk.ForEach(platBindModelNames, func(n string) { dol.migration(n, db) })
	err := dol.initSysData(s1, new(types.SysRole), new(types.SysOrg), new(types.SysRoleUser), new(types.SysMenu), new(types.SysOptionset), new(types.SysUserTemplate), new(types.SysUserTemplateDetail))
	if err != nil {
		s1.Rollback()
		return err
	}
	dol.PlatformDB.ID(domain.ID.Int64).Cols("is_sync").Update(&types.SysDomain{IsSync: null.IntFrom(1)})
	return s1.Commit()
}

// initBusinessDB defined TODO
func (dol *Dolphin) initBusinessDB(domain types.SysDomain) error {
	logrus.Infoln(domain.DriverName.String, domain.DataSource.String)
	xlogger := createXLogger()
	uri, err := uri.Parse(domain.DataSource.String)
	if err != nil {
		return err
	}
	err = domain.CreateDataBase(dol.PlatformDB, domain.DriverName.String, uri.DbName)
	if err != nil {
		return err
	}
	engine, err :=  db.InitDB(domain.DriverName.String, domain.DataSource.String, db.LoggerOpt(xlogger), db.RegisterSQLOpt(), db.RegisterSQLMapOpt(sql.SQLTPL))
	if err != nil {
		return err
	}
	dol.Manager.AddBusinessDB(domain.Domain.String, engine)
	return nil
}

// ServeHTTP defined TODO
func (dol *Dolphin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dol.Web.ServeHTTP(w, r)
}

// Initialize defined init data before bootinh
func (dol *Dolphin) Initialize() error {
	defer dol.Manager.ModelSet().Release()

	SyncModel()
	SyncController()
	SyncSrv(svc.NewServiceContext(CacheStore))

	xlogger := createXLogger()
	logrus.Infoln(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	engine, err := db.InitDB(viper.GetString("db.driver"), viper.GetString("db.dataSource"), db.LoggerOpt(xlogger), db.RegisterSQLOpt(), db.RegisterSQLMapOpt(sql.SQLTPL))
	if err != nil {
		return err
	}
	dol.PlatformDB = engine
	new(types.SysDomain).Ensure(dol.PlatformDB)
	err = new(types.SysDomain).InitSysData(dol.PlatformDB)
	if err != nil {
		return err
	}

    domains := []types.SysDomain{}
	err = dol.PlatformDB.Where("data_source <> '' and domain <> '' and app_name = ? and is_delete != 1", viper.GetString("app.name")).Find(&domains)
	if err != nil {
		return err 
	}

	funk.ForEach(domains, func(d types.SysDomain) { dol.initBusinessDB(d) })
	{
		s := dol.PlatformDB.NewSession()
		s.Begin()
		defer s.Close()
		defer s.Commit()
		err = dol.migration(Name, dol.PlatformDB)
		if err != nil {
			return err 
		}
		err = dol.initSysData(s, new(types.SysClient), new(types.SysUser))
		if err != nil {
			s.Rollback()
			return err
		}
	}

	syncDomain := funk.Filter(domains, func(domain types.SysDomain) bool { return domain.IsSync.Int64 != 1 }).([]types.SysDomain)
	funk.ForEach(syncDomain, func(d types.SysDomain) { dol.initDomain(d) })
	return nil
}

// Done returns a channel of signals to block on after starting the
func (dol *Dolphin) done() <-chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return c
}

// LifeCycle defined TODO
func (dol *Dolphin) LifeCycle(ctx context.Context) error {
	if err := dol.Start(ctx); err != nil {
		return err
	}
	<-dol.done()
	if err := dol.Stop(ctx); err != nil {
		return err
	}
	return nil
}

// Use defined TODO
func (dol *Dolphin) Use(middleware ...core.HandlerFunc) {
	dol.Web.Use(middleware...)
}

// Group defined TODO
func (dol *Dolphin) Group(relativePath string, handlers ...core.HandlerFunc) *core.RouterGroup {
	return dol.Web.Group(relativePath, handlers...)
}

// Static defined TODO
func (dol *Dolphin) Static(relativePath, root string) {
	dol.Web.Static(relativePath, root)
}

// Run defined TODO
func (dol *Dolphin) Run() error {
	if err := dol.Initialize(); err != nil {
		logrus.Error(err)
		return err
	}
	if err := dol.LifeCycle(context.Background()); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

// NewDefault defined TODO
func NewDefault(options ...Option) *Dolphin {
	dol := New(options...)
	dol.Use(Recovery())
	dol.Use(HttpTrace())
	dol.Use(Cors("*", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"))
	dol.Static(viper.GetString("http.static"), path.Join(file.Getwd(), viper.GetString("http.static")))
	dol.Use(DumpBody(DumpRecv()))
	return dol
}

// NewDolphin defined TODO
func New(options ...Option) *Dolphin {
	dol := &Dolphin{}

	for i := range options {
		options[i](dol)
	}
	return dol
}
