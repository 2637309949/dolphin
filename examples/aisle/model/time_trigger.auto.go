// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// TimeTrigger defined
type TimeTrigger struct {
	//
	TgTjNamestr null.String `xorm:"varchar(500) notnull 'tg_tj_namestr'" json:"tg_tj_namestr" xml:"tg_tj_namestr"`
	//
	T3240 null.Int `xorm:"int(11) pk notnull autoincr 't_324_0'" json:"t_324_0" xml:"t_324_0"`
	//
	TgName null.String `xorm:"varchar(500) notnull 'tg_name'" json:"tg_name" xml:"tg_name"`
	//
	Tablecnname null.String `xorm:"varchar(500) notnull 'tablecnname'" json:"tablecnname" xml:"tablecnname"`
	//
	ExeBetweentime null.String `xorm:"varchar(500) notnull 'exe_betweentime'" json:"exe_betweentime" xml:"exe_betweentime"`
	//
	ExeTime null.String `xorm:"varchar(500) notnull 'exe_time'" json:"exe_time" xml:"exe_time"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" xml:"isdelete"`
}

// TableName table name of defined TimeTrigger
func (m *TimeTrigger) TableName() string {
	return "time_trigger"
}