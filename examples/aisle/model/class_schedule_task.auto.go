// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// ClassScheduleTask defined
type ClassScheduleTask struct {
	// CSTId defined
	CSTId null.Int `xorm:"int(11) pk notnull autoincr 'c_s_t_id'" json:"c_s_t_id" form:"c_s_t_id" xml:"c_s_t_id"`
	// CsId defined
	CsId null.Int `xorm:"int(11) 'cs_id'" json:"cs_id" form:"cs_id" xml:"cs_id"`
	// FeedContent defined
	FeedContent null.String `xorm:"'feed_content'" json:"feed_content" form:"feed_content" xml:"feed_content"`
	// FeedRequier defined
	FeedRequier null.String `xorm:"'feed_requier'" json:"feed_requier" form:"feed_requier" xml:"feed_requier"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// Isdelete defined
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) default(54) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// CheckTime defined
	CheckTime null.Time `xorm:"datetime 'check_time'" json:"check_time" form:"check_time" xml:"check_time"`
	// CheckUser defined
	CheckUser null.Int `xorm:"int(11) 'check_user'" json:"check_user" form:"check_user" xml:"check_user"`
	// CheckNoReason defined
	CheckNoReason null.String `xorm:"varchar(5000) 'check_no_reason'" json:"check_no_reason" form:"check_no_reason" xml:"check_no_reason"`
	// AddUserType defined
	AddUserType null.String `xorm:"varchar(1000) 'add_user_type'" json:"add_user_type" form:"add_user_type" xml:"add_user_type"`
	// IfSendStu defined
	IfSendStu null.Int `xorm:"int(11) default(3) 'if_send_stu'" json:"if_send_stu" form:"if_send_stu" xml:"if_send_stu"`
	// IfSendTa defined
	IfSendTa null.Int `xorm:"int(11) default(3) 'if_send_ta'" json:"if_send_ta" form:"if_send_ta" xml:"if_send_ta"`
}

// With defined
func (m *ClassScheduleTask) With(s interface{}) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return nil, errors.New("ptr required")
	}
	mbt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		return nil, err
	}
	return s, err
}

// Marshal defined
func (m *ClassScheduleTask) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ClassScheduleTask) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ClassScheduleTask) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ClassScheduleTask) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ClassScheduleTask) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *ClassScheduleTask) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ClassScheduleTask
func (m *ClassScheduleTask) TableName() string {
	return "class_schedule_task"
}
