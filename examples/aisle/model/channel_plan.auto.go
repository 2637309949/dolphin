// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// ChannelPlan defined
type ChannelPlan struct {
	//
	CPId null.Int `xorm:"int(11) pk notnull autoincr 'c_p_id'" json:"c_p_id" form:"c_p_id" xml:"c_p_id"`
	//
	PkSch null.Int `xorm:"int(11) 'pk_sch'" json:"pk_sch" form:"pk_sch" xml:"pk_sch"`
	//
	PkUser null.Int `xorm:"int(11) 'pk_user'" json:"pk_user" form:"pk_user" xml:"pk_user"`
	//
	InPaln null.Int `xorm:"int(11) 'in_paln'" json:"in_paln" form:"in_paln" xml:"in_paln"`
	//
	OutPlan null.Int `xorm:"int(11) 'out_plan'" json:"out_plan" form:"out_plan" xml:"out_plan"`
	//
	DemoPlan null.Int `xorm:"int(11) 'demo_plan'" json:"demo_plan" form:"demo_plan" xml:"demo_plan"`
	//
	PalnMonth null.Time `xorm:"datetime 'paln_month'" json:"paln_month" form:"paln_month" xml:"paln_month"`
	//
	AchievementPlan decimal.Decimal `xorm:"decimal(11,2) 'achievement_plan'" json:"achievement_plan" form:"achievement_plan" xml:"achievement_plan"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
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

// TableName table name of defined ChannelPlan
func (m *ChannelPlan) TableName() string {
	return "channel_plan"
}
