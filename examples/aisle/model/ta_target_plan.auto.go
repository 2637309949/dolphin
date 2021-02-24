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

// TaTargetPlan defined
type TaTargetPlan struct {
	//
	TTPId null.Int `xorm:"int(11) pk notnull autoincr 't_t_p_id'" json:"t_t_p_id" form:"t_t_p_id" xml:"t_t_p_id"`
	//
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	//
	PlanMonth null.Time `xorm:"datetime 'plan_month'" json:"plan_month" form:"plan_month" xml:"plan_month"`
	//
	YxTarget null.Float `xorm:"float(50,2) 'yx_target'" json:"yx_target" form:"yx_target" xml:"yx_target"`
	//
	TqxTarget null.Float `xorm:"float(11,2) 'tqx_target'" json:"tqx_target" form:"tqx_target" xml:"tqx_target"`
	//
	ZbTarget null.Float `xorm:"float(11,2) 'zb_target'" json:"zb_target" form:"zb_target" xml:"zb_target"`
	//
	LtQdTarget null.Float `xorm:"float(11,2) 'lt_qd_target'" json:"lt_qd_target" form:"lt_qd_target" xml:"lt_qd_target"`
	//
	LtTarget null.Float `xorm:"float(11,2) 'lt_target'" json:"lt_target" form:"lt_target" xml:"lt_target"`
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
	//
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	//
	UserName null.String `xorm:"varchar(100) 'user_name'" json:"user_name" form:"user_name" xml:"user_name"`
	//
	UserPhone null.String `xorm:"varchar(11) 'user_phone'" json:"user_phone" form:"user_phone" xml:"user_phone"`
	//
	KxMoneyTarget null.Float `xorm:"float(50,2) 'kx_money_target'" json:"kx_money_target" form:"kx_money_target" xml:"kx_money_target"`
	//
	YxMoenyTarget null.Float `xorm:"float(50,2) 'yx_moeny_target'" json:"yx_moeny_target" form:"yx_moeny_target" xml:"yx_moeny_target"`
	//
	TqxMoenyTarget null.Float `xorm:"float(50,2) 'tqx_moeny_target'" json:"tqx_moeny_target" form:"tqx_moeny_target" xml:"tqx_moeny_target"`
	//
	LtMoenyTarget null.Float `xorm:"float(50,2) 'lt_moeny_target'" json:"lt_moeny_target" form:"lt_moeny_target" xml:"lt_moeny_target"`
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

// TableName table name of defined TaTargetPlan
func (m *TaTargetPlan) TableName() string {
	return "ta_target_plan"
}
