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

// StudentVisitLanguage defined
type StudentVisitLanguage struct {
	//
	SvlResult null.String `xorm:"varchar(1000) 'svl_result'" json:"svl_result" form:"svl_result" xml:"svl_result"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	SvId null.Int `xorm:"int(11) 'sv_id'" json:"sv_id" form:"sv_id" xml:"sv_id"`
	//
	SlmId null.Int `xorm:"int(11) 'slm_id'" json:"slm_id" form:"slm_id" xml:"slm_id"`
	//
	SVLId null.Int `xorm:"int(11) pk notnull autoincr 's_v_l_id'" json:"s_v_l_id" form:"s_v_l_id" xml:"s_v_l_id"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	//
	StuSystemSta null.Int `xorm:"int(11) 'stu_system_sta'" json:"stu_system_sta" form:"stu_system_sta" xml:"stu_system_sta"`
}

// Parser defined
func (m *SysCommentReply) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *SysCommentReply) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentVisitLanguage
func (m *StudentVisitLanguage) TableName() string {
	return "student_visit_language"
}
