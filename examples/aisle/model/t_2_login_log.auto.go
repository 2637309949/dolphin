// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// T2LoginLog defined
type T2LoginLog struct {
	// LoginLogId defined
	LoginLogId null.Int `xorm:"int(11) pk notnull autoincr 'login_log_id'" json:"login_log_id" form:"login_log_id" xml:"login_log_id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// StartTime defined
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	// EndTime defined
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" form:"end_time" xml:"end_time"`
	// LoginHour defined
	LoginHour null.Float `xorm:"float(11,2) 'login_hour'" json:"login_hour" form:"login_hour" xml:"login_hour"`
	// LoginIp defined
	LoginIp null.String `xorm:"varchar(100) 'login_ip'" json:"login_ip" form:"login_ip" xml:"login_ip"`
}

// Marshal defined
func (m *T2LoginLog) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *T2LoginLog) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *T2LoginLog) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *T2LoginLog) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *T2LoginLog) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *T2LoginLog) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined T2LoginLog
func (m *T2LoginLog) TableName() string {
	return "t_2_login_log"
}
