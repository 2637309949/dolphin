// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ClassChangeTeaHis defined
type ClassChangeTeaHis struct {
	// CCTHId defined
	CCTHId null.Int `xorm:"int(11) pk notnull autoincr 'c_c_t_h_id'" json:"c_c_t_h_id" form:"c_c_t_h_id" xml:"c_c_t_h_id"`
	// PkCct defined
	PkCct null.Int `xorm:"int(11) 'pk_cct'" json:"pk_cct" form:"pk_cct" xml:"pk_cct"`
	// PkHisTea defined
	PkHisTea null.Int `xorm:"int(11) 'pk_his_tea'" json:"pk_his_tea" form:"pk_his_tea" xml:"pk_his_tea"`
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

// TableName table name of defined ClassChangeTeaHis
func (m *ClassChangeTeaHis) TableName() string {
	return "class_change_tea_his"
}

func (r *ClassChangeTeaHis) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalClassChangeTeaHis(data []byte) (ClassChangeTeaHis, error) {
	var r ClassChangeTeaHis
	err := json.Unmarshal(data, &r)
	return r, err
}
