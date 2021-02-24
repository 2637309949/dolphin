// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// OnSale defined
type OnSale struct {
	//
	OnSaleId null.Int `xorm:"int(11) pk notnull autoincr 'on_sale_id'" json:"on_sale_id" form:"on_sale_id" xml:"on_sale_id"`
	//
	OsName null.String `xorm:"varchar(100) 'os_name'" json:"os_name" form:"os_name" xml:"os_name"`
	//
	StartDate null.Time `xorm:"datetime 'start_date'" json:"start_date" form:"start_date" xml:"start_date"`
	//
	EndDate null.Time `xorm:"datetime 'end_date'" json:"end_date" form:"end_date" xml:"end_date"`
	//
	OsMoney null.Float `xorm:"float(10,2) 'os_money'" json:"os_money" form:"os_money" xml:"os_money"`
	//
	OsDiscount null.Float `xorm:"float(10,2) 'os_discount'" json:"os_discount" form:"os_discount" xml:"os_discount"`
	//
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	//
	OrganId null.Int `xorm:"int(11) 'organ_id'" json:"organ_id" form:"organ_id" xml:"organ_id"`
	//
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	//
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	//
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	//
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	//
	OsType null.Int `xorm:"int(11) 'os_type'" json:"os_type" form:"os_type" xml:"os_type"`
	//
	OsWay null.Int `xorm:"int(11) 'os_way'" json:"os_way" form:"os_way" xml:"os_way"`
	//
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	//
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" form:"ma_id" xml:"ma_id"`
	//
	NaId null.Int `xorm:"int(11) 'na_id'" json:"na_id" form:"na_id" xml:"na_id"`
	//
	OsCheck null.Int `xorm:"int(11) 'os_check'" json:"os_check" form:"os_check" xml:"os_check"`
	//
	Remark null.String `xorm:"varchar(3000) 'remark'" json:"remark" form:"remark" xml:"remark"`
	//
	OsBigType null.Int `xorm:"int(11) 'os_big_type'" json:"os_big_type" form:"os_big_type" xml:"os_big_type"`
	//
	SaleJxType null.Int `xorm:"int(11) 'sale_jx_type'" json:"sale_jx_type" form:"sale_jx_type" xml:"sale_jx_type"`
}

// Parser defined
func (m *OnSale) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *OnSale) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined OnSale
func (m *OnSale) TableName() string {
	return "on_sale"
}
