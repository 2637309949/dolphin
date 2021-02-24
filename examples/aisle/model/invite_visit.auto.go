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

// InviteVisit defined
type InviteVisit struct {
	// PlanTime defined
	PlanTime null.Time `xorm:"datetime 'plan_time'" json:"plan_time" form:"plan_time" xml:"plan_time"`
	// PlanContent defined
	PlanContent null.String `xorm:"varchar(500) 'plan_content'" json:"plan_content" form:"plan_content" xml:"plan_content"`
	// RealTime defined
	RealTime null.Time `xorm:"datetime 'real_time'" json:"real_time" form:"real_time" xml:"real_time"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// IvState defined
	IvState null.Int `xorm:"int(11) 'iv_state'" json:"iv_state" form:"iv_state" xml:"iv_state"`
	// SvId defined
	SvId null.Int `xorm:"int(11) 'sv_id'" json:"sv_id" form:"sv_id" xml:"sv_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// ScId defined
	ScId null.Int `xorm:"int(11) 'sc_id'" json:"sc_id" form:"sc_id" xml:"sc_id"`
	// IVId defined
	IVId null.Int `xorm:"int(11) pk notnull autoincr 'i_v_id'" json:"i_v_id" form:"i_v_id" xml:"i_v_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// RoomId defined
	RoomId null.Int `xorm:"int(11) 'room_id'" json:"room_id" form:"room_id" xml:"room_id"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// IvType defined
	IvType null.Int `xorm:"int(11) 'iv_type'" json:"iv_type" form:"iv_type" xml:"iv_type"`
	// StuDepartment defined
	StuDepartment null.Int `xorm:"int(11) 'stu_department'" json:"stu_department" form:"stu_department" xml:"stu_department"`
}

// Parser defined
func (m *InviteVisit) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *InviteVisit) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined InviteVisit
func (m *InviteVisit) TableName() string {
	return "invite_visit"
}
