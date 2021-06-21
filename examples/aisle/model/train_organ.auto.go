// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// TrainOrgan defined
type TrainOrgan struct {
	// TOId defined
	TOId null.Int `xorm:"int(11) pk notnull autoincr 't_o_id'" json:"t_o_id" form:"t_o_id" xml:"t_o_id"`
	// TrainId defined
	TrainId null.Int `xorm:"int(11) 'train_id'" json:"train_id" form:"train_id" xml:"train_id"`
	// OrganId defined
	OrganId null.Int `xorm:"int(11) 'organ_id'" json:"organ_id" form:"organ_id" xml:"organ_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"notnull 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined TrainOrgan
func (m *TrainOrgan) TableName() string {
	return "train_organ"
}
