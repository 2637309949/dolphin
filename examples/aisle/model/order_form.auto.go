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

// OrderForm defined
type OrderForm struct {
	// OrderFormId defined
	OrderFormId null.Int `xorm:"int(11) pk notnull autoincr 'order_form_id'" json:"order_form_id" form:"order_form_id" xml:"order_form_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// OrganId defined
	OrganId null.Int `xorm:"int(11) 'organ_id'" json:"organ_id" form:"organ_id" xml:"organ_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// ScId defined
	ScId null.Int `xorm:"int(11) 'sc_id'" json:"sc_id" form:"sc_id" xml:"sc_id"`
	// AllMoney defined
	AllMoney null.Float `xorm:"float(10,2) 'all_money'" json:"all_money" form:"all_money" xml:"all_money"`
	// SaleMoney defined
	SaleMoney null.Float `xorm:"float(10,2) 'sale_money'" json:"sale_money" form:"sale_money" xml:"sale_money"`
	// MustMoney defined
	MustMoney null.Float `xorm:"float(10,2) 'must_money'" json:"must_money" form:"must_money" xml:"must_money"`
	// GetMoney defined
	GetMoney null.Float `xorm:"float(10,2) 'get_money'" json:"get_money" form:"get_money" xml:"get_money"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// OfState defined
	OfState null.Int `xorm:"int(11) 'of_state'" json:"of_state" form:"of_state" xml:"of_state"`
	// OfType defined
	OfType null.Int `xorm:"int(11) 'of_type'" json:"of_type" form:"of_type" xml:"of_type"`
	// OfNumber defined
	OfNumber null.String `xorm:"varchar(100) 'of_number'" json:"of_number" form:"of_number" xml:"of_number"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
	// RefundMoney defined
	RefundMoney null.Int `xorm:"int(11) 'refund_money'" json:"refund_money" form:"refund_money" xml:"refund_money"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// OfStartDate defined
	OfStartDate null.Time `xorm:"datetime 'of_start_date'" json:"of_start_date" form:"of_start_date" xml:"of_start_date"`
	// OfEndDate defined
	OfEndDate null.Time `xorm:"datetime 'of_end_date'" json:"of_end_date" form:"of_end_date" xml:"of_end_date"`
	// OfRemark defined
	OfRemark null.String `xorm:"varchar(500) 'of_remark'" json:"of_remark" form:"of_remark" xml:"of_remark"`
	// OfJxType defined
	OfJxType null.Int `xorm:"int(11) 'of_jx_type'" json:"of_jx_type" form:"of_jx_type" xml:"of_jx_type"`
	// PtId defined
	PtId null.Int `xorm:"int(11) 'pt_id'" json:"pt_id" form:"pt_id" xml:"pt_id"`
	// OrganSchoolId defined
	OrganSchoolId null.Int `xorm:"int(11) 'organ_school_id'" json:"organ_school_id" form:"organ_school_id" xml:"organ_school_id"`
	// TuitionFee defined
	TuitionFee null.Float `xorm:"float(10,2) 'tuition_fee'" json:"tuition_fee" form:"tuition_fee" xml:"tuition_fee"`
	// TuitionShouldFee defined
	TuitionShouldFee null.Float `xorm:"float(10,2) 'tuition_should_fee'" json:"tuition_should_fee" form:"tuition_should_fee" xml:"tuition_should_fee"`
	// GiftAmount defined
	GiftAmount null.Float `xorm:"float(10,2) 'gift_amount'" json:"gift_amount" form:"gift_amount" xml:"gift_amount"`
	// TeachingMateriaAmount defined
	TeachingMateriaAmount null.Float `xorm:"float(10,2) 'teaching_materia_amount'" json:"teaching_materia_amount" form:"teaching_materia_amount" xml:"teaching_materia_amount"`
	// FmShareNum defined
	FmShareNum null.String `xorm:"varchar(100) 'fm_share_num'" json:"fm_share_num" form:"fm_share_num" xml:"fm_share_num"`
	// ContractNum defined
	ContractNum null.String `xorm:"varchar(50) 'contract_num'" json:"contract_num" form:"contract_num" xml:"contract_num"`
	// AuditState defined
	AuditState null.Int `xorm:"int(11) 'audit_state'" json:"audit_state" form:"audit_state" xml:"audit_state"`
	// FsnId defined
	FsnId null.Int `xorm:"int(11) 'fsn_id'" json:"fsn_id" form:"fsn_id" xml:"fsn_id"`
	// OfBussType defined
	OfBussType null.Int `xorm:"int(11) 'of_buss_type'" json:"of_buss_type" form:"of_buss_type" xml:"of_buss_type"`
	// InvoiceFee defined
	InvoiceFee null.Float `xorm:"float(10,7) 'invoice_fee'" json:"invoice_fee" form:"invoice_fee" xml:"invoice_fee"`
	// IfDedLpfee defined
	IfDedLpfee null.Int `xorm:"int(11) 'if_ded_lpfee'" json:"if_ded_lpfee" form:"if_ded_lpfee" xml:"if_ded_lpfee"`
	// ServiceFee defined
	ServiceFee null.Float `xorm:"float(10,7) 'service_fee'" json:"service_fee" form:"service_fee" xml:"service_fee"`
	// CheckUserId defined
	CheckUserId null.Int `xorm:"int(11) 'check_user_id'" json:"check_user_id" form:"check_user_id" xml:"check_user_id"`
	// CheckTime defined
	CheckTime null.Time `xorm:"datetime 'check_time'" json:"check_time" form:"check_time" xml:"check_time"`
	// TransferSctid defined
	TransferSctid null.Int `xorm:"int(11) 'transfer_sctid'" json:"transfer_sctid" form:"transfer_sctid" xml:"transfer_sctid"`
	// ZxMoney defined
	ZxMoney null.Float `xorm:"float(50,2) 'zx_money'" json:"zx_money" form:"zx_money" xml:"zx_money"`
	// ServiceRatioModifier defined
	ServiceRatioModifier null.Int `xorm:"int(11) 'service_ratio_modifier'" json:"service_ratio_modifier" form:"service_ratio_modifier" xml:"service_ratio_modifier"`
	// FwfkfAllmoney defined
	FwfkfAllmoney null.Float `xorm:"float(50,2) 'fwfkf_allmoney'" json:"fwfkf_allmoney" form:"fwfkf_allmoney" xml:"fwfkf_allmoney"`
	// KfpkfAllmoney defined
	KfpkfAllmoney null.Float `xorm:"float(50,2) 'kfpkf_allmoney'" json:"kfpkf_allmoney" form:"kfpkf_allmoney" xml:"kfpkf_allmoney"`
	// ZxHour defined
	ZxHour null.Float `xorm:"float(50,2) 'zx_hour'" json:"zx_hour" form:"zx_hour" xml:"zx_hour"`
	// FeeCheckState defined
	FeeCheckState null.Int `xorm:"int(11) 'fee_check_state'" json:"fee_check_state" form:"fee_check_state" xml:"fee_check_state"`
	// StkId defined
	StkId null.Int `xorm:"int(11) 'stk_id'" json:"stk_id" form:"stk_id" xml:"stk_id"`
	// ShangkeId defined
	ShangkeId null.Int `xorm:"int(11) 'shangke_id'" json:"shangke_id" form:"shangke_id" xml:"shangke_id"`
	// TrialLessonId defined
	TrialLessonId null.Int `xorm:"int(11) 'trial_lesson_id'" json:"trial_lesson_id" form:"trial_lesson_id" xml:"trial_lesson_id"`
	// IfBookMoney defined
	IfBookMoney null.Int `xorm:"int(11) 'if_book_money'" json:"if_book_money" form:"if_book_money" xml:"if_book_money"`
	// OrderFullTime defined
	OrderFullTime null.Time `xorm:"datetime 'order_full_time'" json:"order_full_time" form:"order_full_time" xml:"order_full_time"`
	// SbtId defined
	SbtId null.Int `xorm:"int(11) 'sbt_id'" json:"sbt_id" form:"sbt_id" xml:"sbt_id"`
	// OrderNozcMoney defined
	OrderNozcMoney null.Float `xorm:"float(50,2) 'order_nozc_money'" json:"order_nozc_money" form:"order_nozc_money" xml:"order_nozc_money"`
	// OrderFeeDate defined
	OrderFeeDate null.Time `xorm:"datetime 'order_fee_date'" json:"order_fee_date" form:"order_fee_date" xml:"order_fee_date"`
	// PkCurrentPay defined
	PkCurrentPay null.Int `xorm:"int(11) 'pk_current_pay'" json:"pk_current_pay" form:"pk_current_pay" xml:"pk_current_pay"`
	// ContractNumNew defined
	ContractNumNew null.String `xorm:"varchar(50) 'contract_num_new'" json:"contract_num_new" form:"contract_num_new" xml:"contract_num_new"`
	// RefHour defined
	RefHour null.Float `xorm:"float(50,2) 'ref_hour'" json:"ref_hour" form:"ref_hour" xml:"ref_hour"`
	// RefMoney defined
	RefMoney null.Float `xorm:"float(50,2) 'ref_money'" json:"ref_money" form:"ref_money" xml:"ref_money"`
	// BookYkMoney defined
	BookYkMoney null.Float `xorm:"float(50,2) 'book_yk_money'" json:"book_yk_money" form:"book_yk_money" xml:"book_yk_money"`
	// StuSign defined
	StuSign null.Int `xorm:"int(11) 'stu_sign'" json:"stu_sign" form:"stu_sign" xml:"stu_sign"`
	// JhrSign defined
	JhrSign null.Int `xorm:"int(11) 'jhr_sign'" json:"jhr_sign" form:"jhr_sign" xml:"jhr_sign"`
	// SqSign defined
	SqSign null.Int `xorm:"int(11) 'sq_sign'" json:"sq_sign" form:"sq_sign" xml:"sq_sign"`
	// StuSignDate defined
	StuSignDate null.Time `xorm:"datetime 'stu_sign_date'" json:"stu_sign_date" form:"stu_sign_date" xml:"stu_sign_date"`
	// YfSignDate defined
	YfSignDate null.Time `xorm:"datetime 'yf_sign_date'" json:"yf_sign_date" form:"yf_sign_date" xml:"yf_sign_date"`
	// Mzks defined
	Mzks null.Float `xorm:"float(50,2) 'mzks'" json:"mzks" form:"mzks" xml:"mzks"`
	// UseMzks defined
	UseMzks null.Float `xorm:"float(50,2) 'use_mzks'" json:"use_mzks" form:"use_mzks" xml:"use_mzks"`
	// SusMzks defined
	SusMzks null.Float `xorm:"float(50,2) 'sus_mzks'" json:"sus_mzks" form:"sus_mzks" xml:"sus_mzks"`
	// EffectiveMonth defined
	EffectiveMonth null.Int `xorm:"int(11) 'effective_month'" json:"effective_month" form:"effective_month" xml:"effective_month"`
	// OfGetMoneyType defined
	OfGetMoneyType null.Int `xorm:"int(11) 'of_get_money_type'" json:"of_get_money_type" form:"of_get_money_type" xml:"of_get_money_type"`
	// OfPeriods defined
	OfPeriods null.Int `xorm:"int(11) 'of_periods'" json:"of_periods" form:"of_periods" xml:"of_periods"`
	// OfNextPeriod defined
	OfNextPeriod null.Int `xorm:"int(11) 'of_next_period'" json:"of_next_period" form:"of_next_period" xml:"of_next_period"`
	// NextPeriodDate defined
	NextPeriodDate null.Time `xorm:"datetime 'next_period_date'" json:"next_period_date" form:"next_period_date" xml:"next_period_date"`
	// OfDeposit defined
	OfDeposit decimal.Decimal `xorm:"decimal(50,3) 'of_deposit'" json:"of_deposit" form:"of_deposit" xml:"of_deposit"`
	// AllViolateMoney defined
	AllViolateMoney decimal.Decimal `xorm:"decimal(50,3) 'all_violate_money'" json:"all_violate_money" form:"all_violate_money" xml:"all_violate_money"`
	// ViolateRatio defined
	ViolateRatio null.Float `xorm:"float(50,5) 'violate_ratio'" json:"violate_ratio" form:"violate_ratio" xml:"violate_ratio"`
	// UseHour defined
	UseHour null.Float `xorm:"float(50,2) 'use_hour'" json:"use_hour" form:"use_hour" xml:"use_hour"`
	// DuePeriod defined
	DuePeriod null.Int `xorm:"int(11) 'due_period'" json:"due_period" form:"due_period" xml:"due_period"`
	// ZxCanhour defined
	ZxCanhour null.Float `xorm:"float(50,2) 'zx_canhour'" json:"zx_canhour" form:"zx_canhour" xml:"zx_canhour"`
	// FxCanhour defined
	FxCanhour null.Float `xorm:"float(50,2) 'fx_canhour'" json:"fx_canhour" form:"fx_canhour" xml:"fx_canhour"`
	// ZxUsehour defined
	ZxUsehour null.Float `xorm:"float(50,2) 'zx_usehour'" json:"zx_usehour" form:"zx_usehour" xml:"zx_usehour"`
	// FxUsehour defined
	FxUsehour null.Float `xorm:"float(50,2) 'fx_usehour'" json:"fx_usehour" form:"fx_usehour" xml:"fx_usehour"`
	// XfsyYypblHour defined
	XfsyYypblHour null.Float `xorm:"float(50,2) 'xfsy_yypbl_hour'" json:"xfsy_yypbl_hour" form:"xfsy_yypbl_hour" xml:"xfsy_yypbl_hour"`
	// XfsyZsHour defined
	XfsyZsHour null.Float `xorm:"float(50,2) 'xfsy_zs_hour'" json:"xfsy_zs_hour" form:"xfsy_zs_hour" xml:"xfsy_zs_hour"`
	// OffqTime defined
	OffqTime null.Time `xorm:"datetime 'offq_time'" json:"offq_time" form:"offq_time" xml:"offq_time"`
	// TmkZxsUsername defined
	TmkZxsUsername null.String `xorm:"varchar(1000) 'tmk_zxs_username'" json:"tmk_zxs_username" form:"tmk_zxs_username" xml:"tmk_zxs_username"`
	// TmkZxsDate defined
	TmkZxsDate null.Time `xorm:"datetime 'tmk_zxs_date'" json:"tmk_zxs_date" form:"tmk_zxs_date" xml:"tmk_zxs_date"`
	// BfOffqTime defined
	BfOffqTime null.Time `xorm:"datetime 'bf_offq_time'" json:"bf_offq_time" form:"bf_offq_time" xml:"bf_offq_time"`
	// PkProtocolType defined
	PkProtocolType null.Int `xorm:"int(11) 'pk_protocol_type'" json:"pk_protocol_type" form:"pk_protocol_type" xml:"pk_protocol_type"`
	// PkDegreeDeposit defined
	PkDegreeDeposit null.Int `xorm:"int(11) 'pk_degree_deposit'" json:"pk_degree_deposit" form:"pk_degree_deposit" xml:"pk_degree_deposit"`
	// DdMoney defined
	DdMoney decimal.Decimal `xorm:"decimal(50,3) 'dd_money'" json:"dd_money" form:"dd_money" xml:"dd_money"`
	// OfVersion defined
	OfVersion null.Int `xorm:"int(11) 'of_version'" json:"of_version" form:"of_version" xml:"of_version"`
	// OfDiscountCoupon defined
	OfDiscountCoupon null.String `xorm:"varchar(200) 'of_discount_coupon'" json:"of_discount_coupon" form:"of_discount_coupon" xml:"of_discount_coupon"`
	// NormalOs defined
	NormalOs decimal.Decimal `xorm:"decimal(50,3) 'normal_os'" json:"normal_os" form:"normal_os" xml:"normal_os"`
	// PkSchOfNorm defined
	PkSchOfNorm null.Int `xorm:"int(11) 'pk_sch_of_norm'" json:"pk_sch_of_norm" form:"pk_sch_of_norm" xml:"pk_sch_of_norm"`
	// DealDate defined
	DealDate null.Time `xorm:"datetime 'deal_date'" json:"deal_date" form:"deal_date" xml:"deal_date"`
	// FqDjmoney defined
	FqDjmoney null.Float `xorm:"float(50,2) 'fq_djmoney'" json:"fq_djmoney" form:"fq_djmoney" xml:"fq_djmoney"`
	// RefundMark defined
	RefundMark null.Int `xorm:"int(11) 'refund_mark'" json:"refund_mark" form:"refund_mark" xml:"refund_mark"`
	// OfBatch defined
	OfBatch null.Int `xorm:"int(11) 'of_batch'" json:"of_batch" form:"of_batch" xml:"of_batch"`
	// IfPartXy defined
	IfPartXy null.Int `xorm:"int(11) 'if_part_xy'" json:"if_part_xy" form:"if_part_xy" xml:"if_part_xy"`
}

// With defined
func (m *OrderForm) With(s interface{}) (interface{}, error) {
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
func (m *OrderForm) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *OrderForm) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *OrderForm) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *OrderForm) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *OrderForm) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *OrderForm) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined OrderForm
func (m *OrderForm) TableName() string {
	return "order_form"
}
