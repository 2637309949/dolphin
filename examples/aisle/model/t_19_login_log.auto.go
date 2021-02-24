// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// T19LoginLog defined
type T19LoginLog struct {
	//
	LoginLogId null.Int `xorm:"int(11) pk notnull autoincr 'login_log_id'" json:"login_log_id" form:"login_log_id" xml:"login_log_id"`
	//
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	//
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	//
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" form:"end_time" xml:"end_time"`
	//
	LoginHour null.Float `xorm:"float(11,2) 'login_hour'" json:"login_hour" form:"login_hour" xml:"login_hour"`
	//
	LoginIp null.String `xorm:"varchar(100) 'login_ip'" json:"login_ip" form:"login_ip" xml:"login_ip"`
}

// Parser defined
func (m *T19LoginLog) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *T19LoginLog) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined T19LoginLog
func (m *T19LoginLog) TableName() string {
	return "t_19_login_log"
}
