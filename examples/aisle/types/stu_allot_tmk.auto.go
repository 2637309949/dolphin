// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// StuAllotTmk defined
type StuAllotTmk struct {
	// SATId defined
	SATId null.Int `xorm:"int(11) pk notnull autoincr 's_a_t_id'" json:"s_a_t_id" form:"s_a_t_id" xml:"s_a_t_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// TmkUserid defined
	TmkUserid null.Int `xorm:"int(11) 'tmk_userid'" json:"tmk_userid" form:"tmk_userid" xml:"tmk_userid"`
	// AllotState defined
	AllotState null.Int `xorm:"int(11) 'allot_state'" json:"allot_state" form:"allot_state" xml:"allot_state"`
	// AllotDate defined
	AllotDate null.Time `xorm:"datetime 'allot_date'" json:"allot_date" form:"allot_date" xml:"allot_date"`
	// QxAllotDate defined
	QxAllotDate null.Time `xorm:"datetime 'qx_allot_date'" json:"qx_allot_date" form:"qx_allot_date" xml:"qx_allot_date"`
	// HisTmkUserid defined
	HisTmkUserid null.Int `xorm:"int(11) 'his_tmk_userid'" json:"his_tmk_userid" form:"his_tmk_userid" xml:"his_tmk_userid"`
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
	// AllotDesc defined
	AllotDesc null.String `xorm:"varchar(5000) 'allot_desc'" json:"allot_desc" form:"allot_desc" xml:"allot_desc"`
	// IfPl defined
	IfPl null.Int `xorm:"int(11) 'if_pl'" json:"if_pl" form:"if_pl" xml:"if_pl"`
}

// TableName table name of defined StuAllotTmk
func (m *StuAllotTmk) TableName() string {
	return "stu_allot_tmk"
}

func (r *StuAllotTmk) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStuAllotTmk(data []byte) (StuAllotTmk, error) {
	var r StuAllotTmk
	err := json.Unmarshal(data, &r)
	return r, err
}
