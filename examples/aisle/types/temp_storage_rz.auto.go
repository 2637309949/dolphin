// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// TempStorageRz defined
type TempStorageRz struct {
	// TSRId defined
	TSRId null.Int `xorm:"int(11) pk notnull autoincr 't_s_r_id'" json:"t_s_r_id" form:"t_s_r_id" xml:"t_s_r_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// TollMoney defined
	TollMoney null.Float `xorm:"float(11,2) 'toll_money'" json:"toll_money" form:"toll_money" xml:"toll_money"`
	// FeeWay defined
	FeeWay null.Int `xorm:"int(11) 'fee_way'" json:"fee_way" form:"fee_way" xml:"fee_way"`
	// ComeEnter defined
	ComeEnter null.Int `xorm:"int(11) 'come_enter'" json:"come_enter" form:"come_enter" xml:"come_enter"`
	// TollDate defined
	TollDate null.Time `xorm:"datetime 'toll_date'" json:"toll_date" form:"toll_date" xml:"toll_date"`
	// Remark defined
	Remark null.String `xorm:"varchar(500) 'remark'" json:"remark" form:"remark" xml:"remark"`
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
	// TsrStuId defined
	TsrStuId null.Int `xorm:"int(11) 'tsr_stu_id'" json:"tsr_stu_id" form:"tsr_stu_id" xml:"tsr_stu_id"`
	// RefId defined
	RefId null.Int `xorm:"int(11) 'ref_id'" json:"ref_id" form:"ref_id" xml:"ref_id"`
	// ZcSource defined
	ZcSource null.Int `xorm:"int(11) 'zc_source'" json:"zc_source" form:"zc_source" xml:"zc_source"`
	// BussType defined
	BussType null.Int `xorm:"int(11) 'buss_type'" json:"buss_type" form:"buss_type" xml:"buss_type"`
	// CheckState defined
	CheckState null.Int `xorm:"int(11) 'check_state'" json:"check_state" form:"check_state" xml:"check_state"`
	// RefRzOrderid defined
	RefRzOrderid null.Int `xorm:"int(11) 'ref_rz_orderid'" json:"ref_rz_orderid" form:"ref_rz_orderid" xml:"ref_rz_orderid"`
}

// TableName table name of defined TempStorageRz
func (m *TempStorageRz) TableName() string {
	return "temp_storage_rz"
}

func (r *TempStorageRz) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTempStorageRz(data []byte) (TempStorageRz, error) {
	var r TempStorageRz
	err := json.Unmarshal(data, &r)
	return r, err
}
