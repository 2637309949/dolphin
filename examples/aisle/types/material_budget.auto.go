// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// MaterialBudget defined
type MaterialBudget struct {
	// MBId defined
	MBId null.Int `xorm:"int(11) pk notnull autoincr 'm_b_id'" json:"m_b_id" form:"m_b_id" xml:"m_b_id"`
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
	// MbRemark defined
	MbRemark null.String `xorm:"varchar(2000) 'mb_remark'" json:"mb_remark" form:"mb_remark" xml:"mb_remark"`
	// TotalMoney defined
	TotalMoney null.Float `xorm:"float(10,2) 'total_money'" json:"total_money" form:"total_money" xml:"total_money"`
	// MbName defined
	MbName null.String `xorm:"varchar(500) 'mb_name'" json:"mb_name" form:"mb_name" xml:"mb_name"`
}

// TableName table name of defined MaterialBudget
func (m *MaterialBudget) TableName() string {
	return "material_budget"
}

func (r *MaterialBudget) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMaterialBudget(data []byte) (MaterialBudget, error) {
	var r MaterialBudget
	err := json.Unmarshal(data, &r)
	return r, err
}
