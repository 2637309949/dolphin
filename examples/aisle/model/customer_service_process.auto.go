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

// CustomerServiceProcess defined
type CustomerServiceProcess struct {
	// CSPId defined
	CSPId null.Int `xorm:"int(11) pk notnull autoincr 'c_s_p_id'" json:"c_s_p_id" form:"c_s_p_id" xml:"c_s_p_id"`
	// CspRemark defined
	CspRemark null.String `xorm:"varchar(2000) 'csp_remark'" json:"csp_remark" form:"csp_remark" xml:"csp_remark"`
	// CspOrder defined
	CspOrder null.Int `xorm:"int(11) 'csp_order'" json:"csp_order" form:"csp_order" xml:"csp_order"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// Isdelete defined
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// CpcId defined
	CpcId null.Int `xorm:"int(11) 'cpc_id'" json:"cpc_id" form:"cpc_id" xml:"cpc_id"`
	// CspNumber defined
	CspNumber null.Int `xorm:"int(11) 'csp_number'" json:"csp_number" form:"csp_number" xml:"csp_number"`
}

// Parser defined
func (m *CustomerServiceProcess) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *CustomerServiceProcess) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined CustomerServiceProcess
func (m *CustomerServiceProcess) TableName() string {
	return "customer_service_process"
}
