// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// RefundCheckFiles defined
type RefundCheckFiles struct {
	// RCFId defined
	RCFId null.Int `xorm:"int(11) pk notnull autoincr 'r_c_f_id'" json:"r_c_f_id" form:"r_c_f_id" xml:"r_c_f_id"`
	// RefCheckId defined
	RefCheckId null.Int `xorm:"int(11) 'ref_check_id'" json:"ref_check_id" form:"ref_check_id" xml:"ref_check_id"`
	// Files defined
	Files null.Int `xorm:"int(11) 'files'" json:"files" form:"files" xml:"files"`
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

// TableName table name of defined RefundCheckFiles
func (m *RefundCheckFiles) TableName() string {
	return "refund_check_files"
}

func (r *RefundCheckFiles) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalRefundCheckFiles(data []byte) (RefundCheckFiles, error) {
	var r RefundCheckFiles
	err := json.Unmarshal(data, &r)
	return r, err
}
