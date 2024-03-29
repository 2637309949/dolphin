// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// FailStudentVisit defined
type FailStudentVisit struct {
	// FSVId defined
	FSVId null.Int `xorm:"int(11) pk notnull autoincr 'f_s_v_id'" json:"f_s_v_id" form:"f_s_v_id" xml:"f_s_v_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// VisitContent defined
	VisitContent null.String `xorm:"varchar(500) 'visit_content'" json:"visit_content" form:"visit_content" xml:"visit_content"`
	// VisitTime defined
	VisitTime null.Time `xorm:"datetime 'visit_time'" json:"visit_time" form:"visit_time" xml:"visit_time"`
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
}

// TableName table name of defined FailStudentVisit
func (m *FailStudentVisit) TableName() string {
	return "fail_student_visit"
}

func (r *FailStudentVisit) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalFailStudentVisit(data []byte) (FailStudentVisit, error) {
	var r FailStudentVisit
	err := json.Unmarshal(data, &r)
	return r, err
}
