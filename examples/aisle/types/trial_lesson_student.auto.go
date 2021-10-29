// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// TrialLessonStudent defined
type TrialLessonStudent struct {
	// TLSId defined
	TLSId null.Int `xorm:"int(11) pk notnull autoincr 't_l_s_id'" json:"t_l_s_id" form:"t_l_s_id" xml:"t_l_s_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// TlId defined
	TlId null.Int `xorm:"int(11) 'tl_id'" json:"tl_id" form:"tl_id" xml:"tl_id"`
	// TlsState defined
	TlsState null.Int `xorm:"int(11) 'tls_state'" json:"tls_state" form:"tls_state" xml:"tls_state"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// SvId defined
	SvId null.Int `xorm:"int(11) 'sv_id'" json:"sv_id" form:"sv_id" xml:"sv_id"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// TlDate defined
	TlDate null.Time `xorm:"datetime 'tl_date'" json:"tl_date" form:"tl_date" xml:"tl_date"`
	// IvId defined
	IvId null.Int `xorm:"int(11) 'iv_id'" json:"iv_id" form:"iv_id" xml:"iv_id"`
	// YjQdDate defined
	YjQdDate null.Time `xorm:"datetime 'yj_qd_date'" json:"yj_qd_date" form:"yj_qd_date" xml:"yj_qd_date"`
}

// TableName table name of defined TrialLessonStudent
func (m *TrialLessonStudent) TableName() string {
	return "trial_lesson_student"
}

func (r *TrialLessonStudent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTrialLessonStudent(data []byte) (TrialLessonStudent, error) {
	var r TrialLessonStudent
	err := json.Unmarshal(data, &r)
	return r, err
}
