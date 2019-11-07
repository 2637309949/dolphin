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

// Query defined
func (e *Engine) Query(ctx *Context) *Query {
	return &Query{m: util.M{}, ctx: ctx}
}

// PageSearch defined
func (e *Engine) PageSearch(db *xorm.Engine, controller, api, table string, q map[string]interface{}) (interface{}, error) {
	page := q["page"].(int)
	size := q["size"].(int)
	q["offset"] = (page - 1) * size

	sqlTagName := fmt.Sprintf("%s_%s_select.tpl", controller, api)
	result, err := db.SqlTemplateClient(sqlTagName, &q).Query().List()

	if err != nil {
		return nil, err
	}

	sqlTagName = fmt.Sprintf("%s_%s_count.tpl", controller, api)
	cresult, err := db.SqlTemplateClient(sqlTagName, &q).Query().List()
	if err != nil {
		return nil, err
	}

	if result == nil {
		ret := util.M{}
		ret["page"] = page
		ret["size"] = size
		ret["data"] = []interface{}{}
		ret["totalrecords"] = 0
		ret["totalpages"] = 0
		ret[""] = 0
		return &ret, nil
	}

	records := cresult[0]["records"].(int64)
	var totalpages int64 = 0
	// if records < int64(size) {
	if records < int64(size) {
		totalpages = 1
	} else if records%int64(size) == 0 {
		totalpages = records / int64(size)
	} else {
		totalpages = records/int64(size) + 1
	}

	ret := util.M{}
	ret["page"] = page
	ret["size"] = size
	ret["data"] = result
	ret["totalrecords"] = records
	ret["totalpages"] = totalpages
	return &ret, nil
}

// Group handlers
func (e *Engine) Group(relativePath string, handlers ...gin.HandlerFunc) *RouterGroup {
	return &RouterGroup{Engine: e, RouterGroup: e.Gin.Group(relativePath, handlers...)}
}

// Sync models
func (e *Engine) Sync(beans ...interface{}) error {
	e.MSets.ForEach(func(m interface{}) {
		for _, db := range e.BusinessDBSet {
			db.Sync2(beans...)
		}
		e.PlatformDB.Sync2(beans...)
	})
	return nil
}

// LoadBusinessDB load BusinessDBSet
func (e *Engine) LoadBusinessDB() {
	domains := []model.Domain{}
	if err := e.PlatformDB.Where("delete_time is null and data_source <> null and domain <> null").Find(&domains); err != nil {
		logrus.Fatal(err)
	}
	for _, domain := range domains {
		e.AddBusinessDB(domain.Domain.String, domain.DriverName.String, domain.DataSource.String)
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
	// Sync model
	e.Sync()
}

// StartUp booting system
func (e *Engine) StartUp() {
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
	e.MSets = &MSets{m: []interface{}{}}
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
