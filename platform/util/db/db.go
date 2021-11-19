package db

import (
	"time"
	"path"
	"os"
	"path/filepath"
	"github.com/spf13/viper"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

//LoggerOpt defined TODO
func LoggerOpt(logger interface{}) func(*xorm.Engine) error {
	return func(engine *xorm.Engine) error {
		engine.SetLogger(logger)
		return nil
	}
}

//RegisterSQLOpt defined TODO
func RegisterSQLOpt() func(*xorm.Engine) error {
	return func(engine *xorm.Engine) error{
		sqlDir := path.Join(".", viper.GetString("dir.sql"))
		err := os.MkdirAll(sqlDir, os.ModePerm)
		if err != nil {
			return  err
		}
		err = engine.RegisterSqlMap(xorm.Xml(sqlDir, ".xml"))
		if err != nil {
			return  err
		}
		err = engine.RegisterSqlTemplate(xorm.Pongo2(sqlDir, ".stpl"))
		if err != nil {
			return  err
		}
		err = engine.RegisterSqlTemplate(xorm.Jet(sqlDir, ".jet"))
		if err != nil {
			return  err
		}
		err = engine.RegisterSqlTemplate(xorm.Default(sqlDir, ".tpl"))
		if err != nil {
			return  err
		}
		return nil
	}
}

// RegisterSQLMapOpt defined TODO
func RegisterSQLMapOpt(sqlMap map[string]string) func(*xorm.Engine) error {
	return func(engine *xorm.Engine) error {
		for key, value := range sqlMap {
			if filepath.Ext(key) == "" {
				engine.AddSql(key, value)
			} else {
				err := engine.AddSqlTemplate(key, value)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// InitDB defined TODO
func InitDB(driver, source string, funk ...func(*xorm.Engine) error) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(driver, source)
	if err != nil {
		return nil, err
	}
	engine.SetConnMaxLifetime(time.Duration(viper.GetInt("db.connMaxLifetime")) * time.Minute)
	engine.SetMaxIdleConns(viper.GetInt("db.maxIdleConns"))
	engine.SetMaxOpenConns(viper.GetInt("db.maxOpenConns"))

	for i := range funk {
		err := funk[i](engine)
		if err != nil {
			return nil, err
		}
	}

	if err = engine.Ping(); err != nil{
		return nil, err
	}
	return engine, nil
}