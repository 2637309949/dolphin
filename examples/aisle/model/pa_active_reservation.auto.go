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

// PaActiveReservation defined
type PaActiveReservation struct {
	//
	PARId null.Int `xorm:"int(11) pk notnull autoincr 'p_a_r_id'" json:"p_a_r_id" form:"p_a_r_id" xml:"p_a_r_id"`
	//
	ActiveTheme null.String `xorm:"varchar(500) 'active_theme'" json:"active_theme" form:"active_theme" xml:"active_theme"`
	//
	ActiveBeginDate null.Time `xorm:"datetime 'active_begin_date'" json:"active_begin_date" form:"active_begin_date" xml:"active_begin_date"`
	//
	ActiveEndDate null.Time `xorm:"datetime 'active_end_date'" json:"active_end_date" form:"active_end_date" xml:"active_end_date"`
	//
	ActiveDesc null.String `xorm:"varchar(500) 'active_desc'" json:"active_desc" form:"active_desc" xml:"active_desc"`
	//
	SyFile null.Int `xorm:"int(11) 'sy_file'" json:"sy_file" form:"sy_file" xml:"sy_file"`
	//
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	//
	MaxNum null.Int `xorm:"int(11) 'max_num'" json:"max_num" form:"max_num" xml:"max_num"`
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
	PblTheme null.Int `xorm:"int(11) 'pbl_theme'" json:"pbl_theme" form:"pbl_theme" xml:"pbl_theme"`
	//
	NowRenshu null.Float `xorm:"float(50,2) 'now_renshu'" json:"now_renshu" form:"now_renshu" xml:"now_renshu"`
	//
	CmId null.Int `xorm:"int(11) 'cm_id'" json:"cm_id" form:"cm_id" xml:"cm_id"`
	//
	KqState null.Int `xorm:"int(11) 'kq_state'" json:"kq_state" form:"kq_state" xml:"kq_state"`
	//
	PaReservation null.String `xorm:"varchar(1000) 'pa_reservation'" json:"pa_reservation" form:"pa_reservation" xml:"pa_reservation"`
}

// Parser defined
func (m *PaActiveReservation) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *PaActiveReservation) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined PaActiveReservation
func (m *PaActiveReservation) TableName() string {
	return "pa_active_reservation"
}
