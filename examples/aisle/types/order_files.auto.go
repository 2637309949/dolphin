// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// OrderFiles defined
type OrderFiles struct {
	// OFId defined
	OFId null.Int `xorm:"int(11) pk notnull autoincr 'o_f_id'" json:"o_f_id" form:"o_f_id" xml:"o_f_id"`
	// OrderId defined
	OrderId null.Int `xorm:"int(11) 'order_id'" json:"order_id" form:"order_id" xml:"order_id"`
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

// TableName table name of defined OrderFiles
func (m *OrderFiles) TableName() string {
	return "order_files"
}

func (r *OrderFiles) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOrderFiles(data []byte) (OrderFiles, error) {
	var r OrderFiles
	err := json.Unmarshal(data, &r)
	return r, err
}
