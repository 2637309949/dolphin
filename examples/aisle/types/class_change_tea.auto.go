// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ClassChangeTea defined
type ClassChangeTea struct {
	// CCTId defined
	CCTId null.Int `xorm:"int(11) pk notnull autoincr 'c_c_t_id'" json:"c_c_t_id" form:"c_c_t_id" xml:"c_c_t_id"`
	// ChangeReason defined
	ChangeReason null.String `xorm:"varchar(200) 'change_reason'" json:"change_reason" form:"change_reason" xml:"change_reason"`
	// ChangeTime defined
	ChangeTime null.Time `xorm:"datetime 'change_time'" json:"change_time" form:"change_time" xml:"change_time"`
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
	// PkClass defined
	PkClass null.Int `xorm:"int(11) 'pk_class'" json:"pk_class" form:"pk_class" xml:"pk_class"`
}

// TableName table name of defined ClassChangeTea
func (m *ClassChangeTea) TableName() string {
	return "class_change_tea"
}

func (r *ClassChangeTea) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalClassChangeTea(data []byte) (ClassChangeTea, error) {
	var r ClassChangeTea
	err := json.Unmarshal(data, &r)
	return r, err
}
