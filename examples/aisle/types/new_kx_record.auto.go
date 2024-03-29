// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// NewKxRecord defined
type NewKxRecord struct {
	// NKRId defined
	NKRId null.Int `xorm:"int(11) pk notnull autoincr 'n_k_r_id'" json:"n_k_r_id" form:"n_k_r_id" xml:"n_k_r_id"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// KxMonth defined
	KxMonth null.Time `xorm:"datetime 'kx_month'" json:"kx_month" form:"kx_month" xml:"kx_month"`
	// KxMoney defined
	KxMoney null.Float `xorm:"float(11,2) 'kx_money'" json:"kx_money" form:"kx_money" xml:"kx_money"`
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
	// NkrType defined
	NkrType null.String `xorm:"varchar(100) 'nkr_type'" json:"nkr_type" form:"nkr_type" xml:"nkr_type"`
	// OldOfId defined
	OldOfId null.Float `xorm:"float(50,2) 'old_of_id'" json:"old_of_id" form:"old_of_id" xml:"old_of_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// SchId defined
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	// OrganId defined
	OrganId null.Int `xorm:"int(11) 'organ_id'" json:"organ_id" form:"organ_id" xml:"organ_id"`
}

// TableName table name of defined NewKxRecord
func (m *NewKxRecord) TableName() string {
	return "new_kx_record"
}

func (r *NewKxRecord) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalNewKxRecord(data []byte) (NewKxRecord, error) {
	var r NewKxRecord
	err := json.Unmarshal(data, &r)
	return r, err
}
