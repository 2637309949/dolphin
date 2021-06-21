// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// NetworkActivityExpend defined
type NetworkActivityExpend struct {
	// NAEId defined
	NAEId null.Int `xorm:"int(11) pk notnull autoincr 'n_a_e_id'" json:"n_a_e_id" form:"n_a_e_id" xml:"n_a_e_id"`
	// NaId defined
	NaId null.Int `xorm:"int(11) 'na_id'" json:"na_id" form:"na_id" xml:"na_id"`
	// ExpendMoney defined
	ExpendMoney null.Float `xorm:"float(11,2) 'expend_money'" json:"expend_money" form:"expend_money" xml:"expend_money"`
	// Remark defined
	Remark null.String `xorm:"varchar(500) 'remark'" json:"remark" form:"remark" xml:"remark"`
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
	// FeeType defined
	FeeType null.Int `xorm:"int(11) 'fee_type'" json:"fee_type" form:"fee_type" xml:"fee_type"`
}

// TableName table name of defined NetworkActivityExpend
func (m *NetworkActivityExpend) TableName() string {
	return "network_activity_expend"
}
