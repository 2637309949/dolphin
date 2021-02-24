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

// TempStorageCz defined
type TempStorageCz struct {
	// TSCId defined
	TSCId null.Int `xorm:"int(11) pk notnull autoincr 't_s_c_id'" json:"t_s_c_id" form:"t_s_c_id" xml:"t_s_c_id"`
	// CzStuId defined
	CzStuId null.Int `xorm:"int(11) 'cz_stu_id'" json:"cz_stu_id" form:"cz_stu_id" xml:"cz_stu_id"`
	// CzRPremium defined
	CzRPremium null.Float `xorm:"float(10,2) 'cz_r_premium'" json:"cz_r_premium" form:"cz_r_premium" xml:"cz_r_premium"`
	// CzWay defined
	CzWay null.Int `xorm:"int(11) 'cz_way'" json:"cz_way" form:"cz_way" xml:"cz_way"`
	// CzCEnter defined
	CzCEnter null.Int `xorm:"int(11) 'cz_c_enter'" json:"cz_c_enter" form:"cz_c_enter" xml:"cz_c_enter"`
	// CzRetDate defined
	CzRetDate null.Time `xorm:"datetime 'cz_ret_date'" json:"cz_ret_date" form:"cz_ret_date" xml:"cz_ret_date"`
	// CzRemark defined
	CzRemark null.String `xorm:"varchar(500) 'cz_remark'" json:"cz_remark" form:"cz_remark" xml:"cz_remark"`
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
	// CzTStorage defined
	CzTStorage null.Int `xorm:"int(11) 'cz_t_storage'" json:"cz_t_storage" form:"cz_t_storage" xml:"cz_t_storage"`
	// FeeId defined
	FeeId null.Int `xorm:"int(11) 'fee_id'" json:"fee_id" form:"fee_id" xml:"fee_id"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// OutType defined
	OutType null.Int `xorm:"int(11) 'out_type'" json:"out_type" form:"out_type" xml:"out_type"`
	// CheckUserId defined
	CheckUserId null.Int `xorm:"int(11) 'check_user_id'" json:"check_user_id" form:"check_user_id" xml:"check_user_id"`
	// CheckTime defined
	CheckTime null.Time `xorm:"datetime 'check_time'" json:"check_time" form:"check_time" xml:"check_time"`
	// ActualRefundTime defined
	ActualRefundTime null.Time `xorm:"datetime 'actual_refund_time'" json:"actual_refund_time" form:"actual_refund_time" xml:"actual_refund_time"`
	// RefType defined
	RefType null.Int `xorm:"int(11) 'ref_type'" json:"ref_type" form:"ref_type" xml:"ref_type"`
	// RefstorageCzFile defined
	RefstorageCzFile null.Int `xorm:"int(11) 'refstorage_cz_file'" json:"refstorage_cz_file" form:"refstorage_cz_file" xml:"refstorage_cz_file"`
	// JzzmFile defined
	JzzmFile null.Int `xorm:"int(11) 'jzzm_file'" json:"jzzm_file" form:"jzzm_file" xml:"jzzm_file"`
	// JzbmFile defined
	JzbmFile null.Int `xorm:"int(11) 'jzbm_file'" json:"jzbm_file" form:"jzbm_file" xml:"jzbm_file"`
	// HmName defined
	HmName null.String `xorm:"varchar(100) 'hm_name'" json:"hm_name" form:"hm_name" xml:"hm_name"`
	// BankCard defined
	BankCard null.String `xorm:"varchar(100) 'bank_card'" json:"bank_card" form:"bank_card" xml:"bank_card"`
	// BankName defined
	BankName null.String `xorm:"varchar(100) 'bank_name'" json:"bank_name" form:"bank_name" xml:"bank_name"`
	// ZdCheckState defined
	ZdCheckState null.Int `xorm:"int(11) 'zd_check_state'" json:"zd_check_state" form:"zd_check_state" xml:"zd_check_state"`
	// CheckDate defined
	CheckDate null.Time `xorm:"datetime 'check_date'" json:"check_date" form:"check_date" xml:"check_date"`
	// PkFlowSet defined
	PkFlowSet null.Int `xorm:"int(11) 'pk_flow_set'" json:"pk_flow_set" form:"pk_flow_set" xml:"pk_flow_set"`
	// NowCheckUser defined
	NowCheckUser null.String `xorm:"varchar(500) 'now_check_user'" json:"now_check_user" form:"now_check_user" xml:"now_check_user"`
	// NowFloor defined
	NowFloor null.Int `xorm:"int(11) 'now_floor'" json:"now_floor" form:"now_floor" xml:"now_floor"`
	// TurnFloor defined
	TurnFloor null.Int `xorm:"int(11) 'turn_floor'" json:"turn_floor" form:"turn_floor" xml:"turn_floor"`
	// HistoryCheckUser defined
	HistoryCheckUser null.String `xorm:"varchar(500) 'history_check_user'" json:"history_check_user" form:"history_check_user" xml:"history_check_user"`
	// PkCfpId defined
	PkCfpId null.Int `xorm:"int(11) 'pk_cfp_id'" json:"pk_cfp_id" form:"pk_cfp_id" xml:"pk_cfp_id"`
	// NowCfid defined
	NowCfid null.Int `xorm:"int(11) 'now_cfid'" json:"now_cfid" form:"now_cfid" xml:"now_cfid"`
	// NextCfid defined
	NextCfid null.Int `xorm:"int(11) 'next_cfid'" json:"next_cfid" form:"next_cfid" xml:"next_cfid"`
	// PkCheckUser defined
	PkCheckUser null.Int `xorm:"int(11) 'pk_check_user'" json:"pk_check_user" form:"pk_check_user" xml:"pk_check_user"`
	// NowCheckid defined
	NowCheckid null.String `xorm:"varchar(100) 'now_checkid'" json:"now_checkid" form:"now_checkid" xml:"now_checkid"`
	// NowCheckname defined
	NowCheckname null.String `xorm:"varchar(100) 'now_checkname'" json:"now_checkname" form:"now_checkname" xml:"now_checkname"`
	// NextCheckid defined
	NextCheckid null.String `xorm:"varchar(100) 'next_checkid'" json:"next_checkid" form:"next_checkid" xml:"next_checkid"`
	// NextCheckname defined
	NextCheckname null.String `xorm:"varchar(100) 'next_checkname'" json:"next_checkname" form:"next_checkname" xml:"next_checkname"`
	// RdrId defined
	RdrId null.Int `xorm:"int(11) 'rdr_id'" json:"rdr_id" form:"rdr_id" xml:"rdr_id"`
}

// Parser defined
func (m *TempStorageCz) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *TempStorageCz) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TempStorageCz
func (m *TempStorageCz) TableName() string {
	return "temp_storage_cz"
}
