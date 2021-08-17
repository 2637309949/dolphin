// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/oauth2/server"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/schemas"
	"github.com/2637309949/dolphin/platform/sql"
	"github.com/2637309949/dolphin/platform/svc"
	"github.com/2637309949/dolphin/platform/types"
	"github.com/2637309949/dolphin/platform/util/file"
	"github.com/2637309949/dolphin/platform/util/http/uri"
	"github.com/2637309949/dolphin/platform/util/slice"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thoas/go-funk"
)

type (
	// Option is a functional option type for Dolphin objects.
	Option func(*Dolphin)
	// Dolphin defined parse app engine
	Dolphin struct {
		http.Handler
		RouterGroup
		Lifecycle
		PlatformDB *xorm.Engine
		Manager    Manager
		OAuth2     *server.Server
		JWT        *JWT
		Http       HttpHandler
		RPC        RpcHandler
		pool       sync.Pool
	}
)

// migration defined TODO
func (dol *Dolphin) migration(name string, db *xorm.Engine) error {
	tables := []types.SysTable{}
	columns := []types.SysTableColumn{}
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

	new(types.SysTable).TruncateTable(db)
	new(types.SysTableColumn).TruncateTable(db)
	stb, stc := slice.Chunk(tables, 50).([][]types.SysTable), slice.Chunk(columns, 50).([][]types.SysTableColumn)

	for i := range stb {
		_, err = db.Insert(stb[i])
		if err != nil {
			return err
		}
	}
	for i := range stc {
		_, err = db.Insert(stc[i])
		if err != nil {
			return err
		}
	}
	return nil
}

// ServeHTTP defined TODO
func (dol *Dolphin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dol.Http.ServeHTTP(w, r)
}

// Reflesh defined init data before bootinh
func (dol *Dolphin) Reflesh() error {
	defer dol.Manager.ModelSet().Release()

	dol.SyncModel()
	dol.SyncController()
	dol.SyncService()
	dol.SyncSrv(svc.NewSvcHepler(CacheStore))

	xlogger := createXLogger()
	logrus.Infoln(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	db, err := xorm.NewEngine(viper.GetString("db.driver"), viper.GetString("db.dataSource"))
	db.SetLogger(xlogger)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	dol.PlatformDB = db
	dol.PlatformDB.SetConnMaxLifetime(time.Duration(viper.GetInt("db.connMaxLifetime")) * time.Minute)
	dol.PlatformDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	dol.PlatformDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))

	err = dol.RegisterSQLDir(dol.PlatformDB, path.Join(".", viper.GetString("dir.sql")))
	if err != nil {
		return err
	}
	// load map
	err = dol.RegisterSQLMap(dol.PlatformDB, sql.SQLTPL)
	if err != nil {
		return err
	}
	new(types.SysDomain).Ensure(db)
	err = new(types.SysDomain).InitSysData(db)
	if err != nil {
		return err
	}
	asyncOnce, domains := sync.Once{}, []types.SysDomain{}
	err = db.Where("data_source <> '' and domain <> '' and app_name = ? and is_delete != 1", viper.GetString("app.name")).Find(&domains)
	if err != nil {
		return err
	}

	s := db.NewSession()
	s.Begin()
	defer s.Close()
	defer func() {
		if err := recover(); err != nil {
			s.Rollback()
			panic(err)
		}
	}()

	for i := range domains {
		domain := domains[i]
		logrus.Infoln(domain.DriverName.String, domain.DataSource.String)
		uri, err := uri.Parse(domain.DataSource.String)
		if err != nil {
			return err
		}
		err = domain.CreateDataBase(dol.PlatformDB, domain.DriverName.String, uri.DbName)
		if err != nil {
			return err
		}
		db, err := xorm.NewEngine(domain.DriverName.String, domain.DataSource.String)
		db.SetLogger(xlogger)
		if err != nil {
			return err
		}
		err = db.Ping()
		if err != nil {
			return err
		}

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

	platBindModelNames := dol.Manager.ModelSet().NameSpaces(Name)
	filtedDomains := funk.Filter(domains, func(domain types.SysDomain) bool { return domain.IsSync.Int64 != 1 }).([]types.SysDomain)
	for i := range filtedDomains {
		asyncOnce.Do(func() {
			err = dol.migration(Name, dol.PlatformDB)
			if err != nil {
				return
			}
			err = new(types.SysClient).InitSysData(s)
			if err != nil {
				return
			}
			err = new(types.SysUser).InitSysData(s)
			if err != nil {
				return
			}
		})
		if err != nil {
			s.Rollback()
			return err
		}

		domain := filtedDomains[i]
		db := dol.Manager.GetBusinessDB(domain.Domain.String)
		s1 := db.NewSession()
		s1.Begin()
		defer s1.Close()
		funk.ForEach(platBindModelNames, func(n string) { dol.migration(n, db) })
		err = new(types.SysRole).InitSysData(s1)
		if err != nil {
			fmt.Println(123, err)
			s1.Rollback()
			return err
		}

		err = new(types.SysOrg).InitSysData(s1)
		if err != nil {
			s1.Rollback()
			return err
		}
		err = new(types.SysRoleUser).InitSysData(s1)
		if err != nil {
			s1.Rollback()
			return err
		}
		err = new(types.SysMenu).InitSysData(s1)
		if err != nil {
			s1.Rollback()
		}
		err = new(types.SysOptionset).InitSysData(s1)
		if err != nil {
			s1.Rollback()
			return err
		}

		err = new(types.SysUserTemplate).InitSysData(s1)
		if err != nil {
			s1.Rollback()
			return err
		}
		err = new(types.SysUserTemplateDetail).InitSysData(s1)
		if err != nil {
			s1.Rollback()
			return err
		}

		_, err = s.ID(domain.ID.Int64).Cols("is_sync").Update(&types.SysDomain{IsSync: null.IntFrom(1)})
		if err != nil {
			s.Rollback()
			return err
		}
		s1.Commit()
	}

	return s.Commit()
}

// RegisterSQLMap defined TODO
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

// RegisterSQLDir defined TODO
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

// LifeCycle start LifeCycle hooks
func (dol *Dolphin) LifeCycle(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	signal := dol.done()
	if err := dol.Start(ctx); err != nil {
		return err
	}
	<-signal
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	if err := dol.Stop(ctx); err != nil {
		return err
	}
	return nil
}

// Run defined TODO
func (dol *Dolphin) Run() error {
	if err := dol.Reflesh(); err != nil {
		logrus.Errorln(err)
		return err
	}
	if err := dol.LifeCycle(context.Background()); err != nil {
		logrus.Errorln(err)
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
	dol.Use(DumpBody(DumpRecv(dol)))
	return dol
}

// NewDolphin defined TODO
func New(options ...Option) *Dolphin {
	dol := &Dolphin{}
	dol.RouterGroup = RouterGroup{Handlers: []HandlerFunc{}, basePath: "/"}
	dol.RouterGroup.dol = dol
	dol.pool.New = func() interface{} {
		return NewContext(dol)
	}
	for i := range options {
		options[i](dol)
	}

	return dol
}
