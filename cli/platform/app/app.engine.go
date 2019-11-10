package app

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path"

	"github.com/2637309949/dolphin/cli/platform/sql"

	"github.com/2637309949/dolphin/cli/platform/model"
	"github.com/2637309949/dolphin/cli/platform/util"
	method "github.com/bu/gin-method-override"
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xormplus/xorm"
)

// Engine defined parse app engine
type Engine struct {
	Gin           *gin.Engine
	MSets         MSeti
	PlatformDB    *xorm.Engine
	BusinessDBSet map[string]*xorm.Engine
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

// LoadBusinessDB load BusinessDBSet
func (e *Engine) LoadBusinessDB() {
	domains := []model.Domain{}
	if err := e.PlatformDB.Where("data_source <> '' and domain <> '' and delete_time is null").Find(&domains); err != nil {
		logrus.Fatal(err)
	}
	for _, domain := range domains {
		e.AddBusinessDB(domain.Domain.String, domain.DriverName.String, domain.DataSource.String)
	}

	fmt.Println(e.BusinessDBSet)

	// async domain model
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
	sqlDir := viper.GetString("sqlDir")
	sqlDir = path.Join(".", sqlDir)
	db, err := xorm.NewEngine(driverName, dataSource)
	if err != nil {
		logrus.Fatal(err)
	}
	xLogger := xorm.NewSimpleLogger(logrus.StandardLogger().Out)
	xLogger.ShowSQL(true)
	db.SetLogger(xLogger)

	// RegisterSqlMap
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

// LoadPlatformDB adddb
func (e *Engine) LoadPlatformDB() {
	var err error
	sqlDir := viper.GetString("sqlDir")
	sqlDir = path.Join(".", sqlDir)
	e.PlatformDB, err = xorm.NewEngine(viper.GetString("dbType"), viper.GetString("dbUri"))
	if err != nil {
		logrus.Fatal(err)
	}
	if err = e.PlatformDB.Ping(); err != nil {
		logrus.Fatal(err)
	}
	xLogger := xorm.NewSimpleLogger(logrus.StandardLogger().Out)
	xLogger.ShowSQL(true)
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

// Run booting system
func (e *Engine) Run() {
	e.LoadPlatformDB()
	e.LoadBusinessDB()
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
	e.Gin.Use(nice.Recovery(func(ctx *gin.Context, err interface{}) {
		code := 500
		if err, ok := err.(util.Error); ok {
			code = err.Code
		}
		ctx.JSON(http.StatusInternalServerError, util.M{"code": code, "message": err})
	}))
	e.Gin.Use(method.ProcessMethodOverride(e.Gin))
	return e
}
