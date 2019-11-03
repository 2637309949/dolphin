package app

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/2637309949/dolphin/cli/platform/sql"
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

// Sync2 handlers
func (e *Engine) Sync2(beans ...interface{}) error {
	for _, db := range e.BusinessDBSet {
		db.Sync2(beans...)
	}
	return nil
}

// NewEngine init Engine
func NewEngine() *Engine {
	var err error
	// set logger
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	// read config
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.AddConfigPath("conf")
	viper.AutomaticEnv()
	viper.SetDefault("sqlDir", "sql")
	viper.SetDefault("dbType", "mysql")
	viper.SetDefault("port", "8089")
	viper.SetDefault("dbUri", "root:111111@/dolphin?charset=utf8&parseTime=True&loc=Local")
	if err = viper.ReadInConfig(); err != nil {
		logrus.Warn("unable to read config file")
	}

	// init xorm
	e := &Engine{}
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
	// RegisterSqlMap
	sqlDir := path.Join(".", viper.GetString("sqlDir"))
	if err = os.MkdirAll(sqlDir, os.ModePerm); err != nil {
		logrus.Fatal(err)
	}
	if err = e.PlatformDB.RegisterSqlMap(xorm.Xml(sqlDir, ".xml")); err != nil {
		logrus.Fatal(err)
	}
	if err = e.PlatformDB.RegisterSqlTemplate(xorm.Pongo2(sqlDir, ".stpl")); err != nil {
		logrus.Fatal(err)
	}
	if err = e.PlatformDB.RegisterSqlTemplate(xorm.Jet(sqlDir, ".jet")); err != nil {
		logrus.Fatal(err)
	}
	if err = e.PlatformDB.RegisterSqlTemplate(xorm.Default(sqlDir, ".tpl")); err != nil {
		logrus.Fatal(err)
	}
	for k, v := range sql.SQLTPL {
		logrus.Info(k, v)
		if err = e.PlatformDB.AddSqlTemplate(k, v); err != nil {
			logrus.Fatal(err)
		}
	}

	// init gin
	gin.SetMode(gin.ReleaseMode)
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
