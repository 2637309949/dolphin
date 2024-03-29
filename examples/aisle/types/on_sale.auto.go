// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// OnSale defined
type OnSale struct {
	// OnSaleId defined
	OnSaleId null.Int `xorm:"int(11) pk notnull autoincr 'on_sale_id'" json:"on_sale_id" form:"on_sale_id" xml:"on_sale_id"`
	// OsName defined
	OsName null.String `xorm:"varchar(100) 'os_name'" json:"os_name" form:"os_name" xml:"os_name"`
	// StartDate defined
	StartDate null.Time `xorm:"datetime 'start_date'" json:"start_date" form:"start_date" xml:"start_date"`
	// EndDate defined
	EndDate null.Time `xorm:"datetime 'end_date'" json:"end_date" form:"end_date" xml:"end_date"`
	// OsMoney defined
	OsMoney null.Float `xorm:"float(10,2) 'os_money'" json:"os_money" form:"os_money" xml:"os_money"`
	// OsDiscount defined
	OsDiscount null.Float `xorm:"float(10,2) 'os_discount'" json:"os_discount" form:"os_discount" xml:"os_discount"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
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
	// OsType defined
	OsType null.Int `xorm:"int(11) 'os_type'" json:"os_type" form:"os_type" xml:"os_type"`
	// OsWay defined
	OsWay null.Int `xorm:"int(11) 'os_way'" json:"os_way" form:"os_way" xml:"os_way"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// MaId defined
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" form:"ma_id" xml:"ma_id"`
	// NaId defined
	NaId null.Int `xorm:"int(11) 'na_id'" json:"na_id" form:"na_id" xml:"na_id"`
	// OsCheck defined
	OsCheck null.Int `xorm:"int(11) 'os_check'" json:"os_check" form:"os_check" xml:"os_check"`
	// Remark defined
	Remark null.String `xorm:"varchar(3000) 'remark'" json:"remark" form:"remark" xml:"remark"`
	// OsBigType defined
	OsBigType null.Int `xorm:"int(11) 'os_big_type'" json:"os_big_type" form:"os_big_type" xml:"os_big_type"`
	// SaleJxType defined
	SaleJxType null.Int `xorm:"int(11) 'sale_jx_type'" json:"sale_jx_type" form:"sale_jx_type" xml:"sale_jx_type"`
}

// TableName table name of defined OnSale
func (m *OnSale) TableName() string {
	return "on_sale"
}

func (r *OnSale) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalOnSale(data []byte) (OnSale, error) {
	var r OnSale
	err := json.Unmarshal(data, &r)
	return r, err
}
