// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// StudentClassType defined
type StudentClassType struct {
	// SCTId defined
	SCTId null.Int `xorm:"int(11) pk notnull autoincr 's_c_t_id'" json:"s_c_t_id" form:"s_c_t_id" xml:"s_c_t_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// CtId defined
	CtId null.Int `xorm:"int(11) 'ct_id'" json:"ct_id" form:"ct_id" xml:"ct_id"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// SppId defined
	SppId null.Int `xorm:"int(11) 'spp_id'" json:"spp_id" form:"spp_id" xml:"spp_id"`
	// AllPrice defined
	AllPrice null.Float `xorm:"float(11,2) 'all_price'" json:"all_price" form:"all_price" xml:"all_price"`
	// OnePrice defined
	OnePrice null.Float `xorm:"float(10,2) 'one_price'" json:"one_price" form:"one_price" xml:"one_price"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
	// OsMoney defined
	OsMoney null.Float `xorm:"float(10,2) 'os_money'" json:"os_money" form:"os_money" xml:"os_money"`
	// FinalMoney defined
	FinalMoney null.Float `xorm:"float(10,2) 'final_money'" json:"final_money" form:"final_money" xml:"final_money"`
	// RefundOnePrice defined
	RefundOnePrice null.Float `xorm:"float(10,2) 'refund_one_price'" json:"refund_one_price" form:"refund_one_price" xml:"refund_one_price"`
	// FsId defined
	FsId null.Int `xorm:"int(11) 'fs_id'" json:"fs_id" form:"fs_id" xml:"fs_id"`
	// BuyHour defined
	BuyHour null.Float `xorm:"float(11,2) 'buy_hour'" json:"buy_hour" form:"buy_hour" xml:"buy_hour"`
	// RefundHour defined
	RefundHour null.Float `xorm:"float(50,2) 'refund_hour'" json:"refund_hour" form:"refund_hour" xml:"refund_hour"`
	// UseHour defined
	UseHour null.Float `xorm:"float(10,2) 'use_hour'" json:"use_hour" form:"use_hour" xml:"use_hour"`
	// YibaoHour defined
	YibaoHour null.Float `xorm:"float(10,2) 'yibao_hour'" json:"yibao_hour" form:"yibao_hour" xml:"yibao_hour"`
	// SurplusHour defined
	SurplusHour null.Float `xorm:"float(10,2) 'surplus_hour'" json:"surplus_hour" form:"surplus_hour" xml:"surplus_hour"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// WxhMoney defined
	WxhMoney null.Float `xorm:"float(10,2) 'wxh_money'" json:"wxh_money" form:"wxh_money" xml:"wxh_money"`
	// XhMoney defined
	XhMoney null.Float `xorm:"float(10,2) 'xh_money'" json:"xh_money" form:"xh_money" xml:"xh_money"`
	// CwKxPrice defined
	CwKxPrice null.Float `xorm:"float(10,2) 'cw_kx_price'" json:"cw_kx_price" form:"cw_kx_price" xml:"cw_kx_price"`
	// GiveHour defined
	GiveHour null.Int `xorm:"int(11) 'give_hour'" json:"give_hour" form:"give_hour" xml:"give_hour"`
	// ZxHour defined
	ZxHour null.Float `xorm:"float(50,2) 'zx_hour'" json:"zx_hour" form:"zx_hour" xml:"zx_hour"`
	// SctZxMoney defined
	SctZxMoney null.Float `xorm:"float(50,2) 'sct_zx_money'" json:"sct_zx_money" form:"sct_zx_money" xml:"sct_zx_money"`
	// TransferClasstypeHour defined
	TransferClasstypeHour null.Float `xorm:"float(50,2) 'transfer_classtype_hour'" json:"transfer_classtype_hour" form:"transfer_classtype_hour" xml:"transfer_classtype_hour"`
	// TransferClasstypeMoney defined
	TransferClasstypeMoney null.Float `xorm:"float(50,2) 'transfer_classtype_money'" json:"transfer_classtype_money" form:"transfer_classtype_money" xml:"transfer_classtype_money"`
	// OrderAuditState defined
	OrderAuditState null.Int `xorm:"int(11) 'order_audit_state'" json:"order_audit_state" form:"order_audit_state" xml:"order_audit_state"`
	// NowCustomer defined
	NowCustomer null.Int `xorm:"int(11) 'now_customer'" json:"now_customer" form:"now_customer" xml:"now_customer"`
	// HisCustomer defined
	HisCustomer null.String `xorm:"varchar(100) 'his_customer'" json:"his_customer" form:"his_customer" xml:"his_customer"`
	// SbtId defined
	SbtId null.Int `xorm:"int(11) 'sbt_id'" json:"sbt_id" form:"sbt_id" xml:"sbt_id"`
	// GlgiveZsPrice defined
	GlgiveZsPrice null.Float `xorm:"float(50,2) 'glgive_zs_price'" json:"glgive_zs_price" form:"glgive_zs_price" xml:"glgive_zs_price"`
	// SctNosyhourDate defined
	SctNosyhourDate null.Time `xorm:"datetime 'sct_nosyhour_date'" json:"sct_nosyhour_date" form:"sct_nosyhour_date" xml:"sct_nosyhour_date"`
	// RefMoney defined
	RefMoney null.Float `xorm:"float(50,2) 'ref_money'" json:"ref_money" form:"ref_money" xml:"ref_money"`
	// NoKouHour defined
	NoKouHour null.Float `xorm:"float(50,2) 'no_kou_hour'" json:"no_kou_hour" form:"no_kou_hour" xml:"no_kou_hour"`
	// FirstClassDate defined
	FirstClassDate null.Time `xorm:"datetime 'first_class_date'" json:"first_class_date" form:"first_class_date" xml:"first_class_date"`
	// OverdueHour defined
	OverdueHour null.Float `xorm:"float(50,2) 'overdue_hour'" json:"overdue_hour" form:"overdue_hour" xml:"overdue_hour"`
	// OverdueMoney defined
	OverdueMoney null.Float `xorm:"float(50,2) 'overdue_money'" json:"overdue_money" form:"overdue_money" xml:"overdue_money"`
	// PblYuyan defined
	PblYuyan null.Int `xorm:"int(11) 'pbl_yuyan'" json:"pbl_yuyan" form:"pbl_yuyan" xml:"pbl_yuyan"`
	// BplYuyanRemark defined
	BplYuyanRemark null.String `xorm:"varchar(2000) 'bpl_yuyan_remark'" json:"bpl_yuyan_remark" form:"bpl_yuyan_remark" xml:"bpl_yuyan_remark"`
	// FqRefhour defined
	FqRefhour null.Float `xorm:"float(50,2) 'fq_refhour'" json:"fq_refhour" form:"fq_refhour" xml:"fq_refhour"`
	// FqRefmoney defined
	FqRefmoney null.Float `xorm:"float(50,2) 'fq_refmoney'" json:"fq_refmoney" form:"fq_refmoney" xml:"fq_refmoney"`
}

// Marshal defined
func (m *StudentClassType) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StudentClassType) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *StudentClassType) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *StudentClassType) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StudentClassType
func (m *StudentClassType) TableName() string {
	return "student_class_type"
}
