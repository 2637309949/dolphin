// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"github.com/2637309949/dolphin/packages/null"
)

// TextBook defined
type TextBook struct {
	// TextBookId defined
	TextBookId null.Int `xorm:"int(11) pk notnull autoincr 'text_book_id'" json:"text_book_id" form:"text_book_id" xml:"text_book_id"`
	// TbName defined
	TbName null.String `xorm:"varchar(100) 'tb_name'" json:"tb_name" form:"tb_name" xml:"tb_name"`
	// TbPrice defined
	TbPrice null.Float `xorm:"float(11,2) 'tb_price'" json:"tb_price" form:"tb_price" xml:"tb_price"`
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
	// TbType defined
	TbType null.Int `xorm:"int(11) 'tb_type'" json:"tb_type" form:"tb_type" xml:"tb_type"`
}

// TableName table name of defined TextBook
func (m *TextBook) TableName() string {
	return "text_book"
}