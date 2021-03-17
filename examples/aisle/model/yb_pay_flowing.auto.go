// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/decimal"
	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// YbPayFlowing defined
type YbPayFlowing struct {
	// YPFId defined
	YPFId null.Int `xorm:"int(11) pk notnull autoincr 'y_p_f_id'" json:"y_p_f_id" form:"y_p_f_id" xml:"y_p_f_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// OrderId defined
	OrderId null.Int `xorm:"int(11) 'order_id'" json:"order_id" form:"order_id" xml:"order_id"`
	// PayMoney defined
	PayMoney decimal.Decimal `xorm:"decimal(50,3) 'pay_money'" json:"pay_money" form:"pay_money" xml:"pay_money"`
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
	// YbFlowNumber defined
	YbFlowNumber null.String `xorm:"varchar(1000) 'yb_flow_number'" json:"yb_flow_number" form:"yb_flow_number" xml:"yb_flow_number"`
	// SuccessFailure defined
	SuccessFailure null.Int `xorm:"int(11) 'success_failure'" json:"success_failure" form:"success_failure" xml:"success_failure"`
	// SourceType defined
	SourceType null.Int `xorm:"int(11) 'source_type'" json:"source_type" form:"source_type" xml:"source_type"`
	// Posmenu defined
	Posmenu null.Int `xorm:"int(11) 'posmenu'" json:"posmenu" form:"posmenu" xml:"posmenu"`
	// PayTime defined
	PayTime null.Time `xorm:"datetime 'pay_time'" json:"pay_time" form:"pay_time" xml:"pay_time"`
	// LoginNum defined
	LoginNum null.String `xorm:"varchar(50) 'login_num'" json:"login_num" form:"login_num" xml:"login_num"`
	// PosNum defined
	PosNum null.String `xorm:"varchar(20) 'pos_num'" json:"pos_num" form:"pos_num" xml:"pos_num"`
	// Posrequestid defined
	Posrequestid null.String `xorm:"varchar(6) 'posrequestid'" json:"posrequestid" form:"posrequestid" xml:"posrequestid"`
	// Referno defined
	Referno null.String `xorm:"varchar(50) 'referno'" json:"referno" form:"referno" xml:"referno"`
	// Chequeno defined
	Chequeno null.String `xorm:"varchar(50) 'chequeno'" json:"chequeno" form:"chequeno" xml:"chequeno"`
	// Bankcardno defined
	Bankcardno null.String `xorm:"varchar(50) 'bankcardno'" json:"bankcardno" form:"bankcardno" xml:"bankcardno"`
	// Bankcardname defined
	Bankcardname null.String `xorm:"varchar(50) 'bankcardname'" json:"bankcardname" form:"bankcardname" xml:"bankcardname"`
	// Bankcardtype defined
	Bankcardtype null.String `xorm:"varchar(50) 'bankcardtype'" json:"bankcardtype" form:"bankcardtype" xml:"bankcardtype"`
	// Bankorderno defined
	Bankorderno null.String `xorm:"varchar(50) 'bankorderno'" json:"bankorderno" form:"bankorderno" xml:"bankorderno"`
	// Yeepayorderno defined
	Yeepayorderno null.String `xorm:"varchar(50) 'yeepayorderno'" json:"yeepayorderno" form:"yeepayorderno" xml:"yeepayorderno"`
	// Customerno defined
	Customerno null.String `xorm:"varchar(50) 'customerno'" json:"customerno" form:"customerno" xml:"customerno"`
	// Paytypecode defined
	Paytypecode null.Int `xorm:"int(11) 'paytypecode'" json:"paytypecode" form:"paytypecode" xml:"paytypecode"`
	// PayState defined
	PayState null.Int `xorm:"int(11) 'pay_state'" json:"pay_state" form:"pay_state" xml:"pay_state"`
	// PkSchId defined
	PkSchId null.Int `xorm:"int(11) 'pk_sch_id'" json:"pk_sch_id" form:"pk_sch_id" xml:"pk_sch_id"`
	// YeePayNum defined
	YeePayNum null.String `xorm:"varchar(20) 'yee_pay_num'" json:"yee_pay_num" form:"yee_pay_num" xml:"yee_pay_num"`
	// Amount defined
	Amount decimal.Decimal `xorm:"decimal(50,3) 'amount'" json:"amount" form:"amount" xml:"amount"`
	// AccountState defined
	AccountState null.Int `xorm:"int(11) 'account_state'" json:"account_state" form:"account_state" xml:"account_state"`
	// ReconciliationsState defined
	ReconciliationsState null.Int `xorm:"int(11) 'reconciliations_state'" json:"reconciliations_state" form:"reconciliations_state" xml:"reconciliations_state"`
	// ErrorMsg defined
	ErrorMsg null.String `xorm:"varchar(500) 'error_msg'" json:"error_msg" form:"error_msg" xml:"error_msg"`
	// PkYrrId defined
	PkYrrId null.Int `xorm:"int(11) 'pk_yrr_id'" json:"pk_yrr_id" form:"pk_yrr_id" xml:"pk_yrr_id"`
	// ServiceCharge defined
	ServiceCharge decimal.Decimal `xorm:"decimal(50,3) 'service_charge'" json:"service_charge" form:"service_charge" xml:"service_charge"`
	// GetMoney defined
	GetMoney decimal.Decimal `xorm:"decimal(50,3) 'get_money'" json:"get_money" form:"get_money" xml:"get_money"`
	// UpdateStuybzfDesc defined
	UpdateStuybzfDesc null.String `xorm:"varchar(1000) 'update_stuybzf_desc'" json:"update_stuybzf_desc" form:"update_stuybzf_desc" xml:"update_stuybzf_desc"`
	// YpfId defined
	YpfId null.Int `xorm:"int(11) 'ypf_id'" json:"ypf_id" form:"ypf_id" xml:"ypf_id"`
}

// Marshal defined
func (m *YbPayFlowing) With(s interface{}) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return nil, errors.New("ptr required")
	}
	mbt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		return nil, err
	}
	return s, err
}

// Marshal defined
func (m *YbPayFlowing) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *YbPayFlowing) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *YbPayFlowing) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *YbPayFlowing) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *YbPayFlowing) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *YbPayFlowing) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined YbPayFlowing
func (m *YbPayFlowing) TableName() string {
	return "yb_pay_flowing"
}
