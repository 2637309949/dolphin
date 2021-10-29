// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// FqorderZx defined
type FqorderZx struct {
	// FqorderZxId defined
	FqorderZxId null.Int `xorm:"int(11) pk notnull autoincr 'fqorder_zx_id'" json:"fqorder_zx_id" form:"fqorder_zx_id" xml:"fqorder_zx_id"`
	// OrderId defined
	OrderId null.Int `xorm:"int(11) 'order_id'" json:"order_id" form:"order_id" xml:"order_id"`
	// ZcSchid defined
	ZcSchid null.Int `xorm:"int(11) 'zc_schid'" json:"zc_schid" form:"zc_schid" xml:"zc_schid"`
	// ZrSchid defined
	ZrSchid null.Int `xorm:"int(11) 'zr_schid'" json:"zr_schid" form:"zr_schid" xml:"zr_schid"`
	// ZxDesc defined
	ZxDesc null.String `xorm:"varchar(1000) 'zx_desc'" json:"zx_desc" form:"zx_desc" xml:"zx_desc"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// CheckUser defined
	CheckUser null.Int `xorm:"int(11) 'check_user'" json:"check_user" form:"check_user" xml:"check_user"`
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
	// AllotTa defined
	AllotTa null.Int `xorm:"int(11) 'allot_ta'" json:"allot_ta" form:"allot_ta" xml:"allot_ta"`
}

// TableName table name of defined FqorderZx
func (m *FqorderZx) TableName() string {
	return "fqorder_zx"
}

func (r *FqorderZx) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFqorderZx(data []byte) (FqorderZx, error) {
	var r FqorderZx
	err := json.Unmarshal(data, &r)
	return r, err
}
