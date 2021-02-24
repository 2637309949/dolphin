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

// Invoice defined
type Invoice struct {
	// InvoiceId defined
	InvoiceId null.Int `xorm:"int(11) pk notnull autoincr 'invoice_id'" json:"invoice_id" form:"invoice_id" xml:"invoice_id"`
	// InvoiceNumber defined
	InvoiceNumber null.String `xorm:"varchar(100) 'invoice_number'" json:"invoice_number" form:"invoice_number" xml:"invoice_number"`
	// InvoiceMoney defined
	InvoiceMoney null.Int `xorm:"int(11) 'invoice_money'" json:"invoice_money" form:"invoice_money" xml:"invoice_money"`
	// InvoiceCompany defined
	InvoiceCompany null.String `xorm:"varchar(100) 'invoice_company'" json:"invoice_company" form:"invoice_company" xml:"invoice_company"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// InvoiceType defined
	InvoiceType null.Int `xorm:"int(11) 'invoice_type'" json:"invoice_type" form:"invoice_type" xml:"invoice_type"`
	// FeeId defined
	FeeId null.Int `xorm:"int(11) 'fee_id'" json:"fee_id" form:"fee_id" xml:"fee_id"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// BillType defined
	BillType null.Int `xorm:"int(11) 'bill_type'" json:"bill_type" form:"bill_type" xml:"bill_type"`
	// OrId defined
	OrId null.Int `xorm:"int(11) 'or_id'" json:"or_id" form:"or_id" xml:"or_id"`
}

// Parser defined
func (m *Invoice) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *Invoice) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined Invoice
func (m *Invoice) TableName() string {
	return "invoice"
}
