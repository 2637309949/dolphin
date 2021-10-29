// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// LabelInfo defined
type LabelInfo struct {
	// LabelInfoId defined
	LabelInfoId null.Int `xorm:"int(11) pk notnull autoincr 'label_info_id'" json:"label_info_id" form:"label_info_id" xml:"label_info_id"`
	// ParentId defined
	ParentId null.Int `xorm:"int(11) 'parent_id'" json:"parent_id" form:"parent_id" xml:"parent_id"`
	// LabelName defined
	LabelName null.String `xorm:"varchar(100) 'label_name'" json:"label_name" form:"label_name" xml:"label_name"`
	// LabelRemark defined
	LabelRemark null.String `xorm:"varchar(500) 'label_remark'" json:"label_remark" form:"label_remark" xml:"label_remark"`
	// PtId defined
	PtId null.Int `xorm:"int(11) 'pt_id'" json:"pt_id" form:"pt_id" xml:"pt_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined LabelInfo
func (m *LabelInfo) TableName() string {
	return "label_info"
}

func (r *LabelInfo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalLabelInfo(data []byte) (LabelInfo, error) {
	var r LabelInfo
	err := json.Unmarshal(data, &r)
	return r, err
}
