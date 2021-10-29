// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// StudentTypeJournal defined
type StudentTypeJournal struct {
	// STJId defined
	STJId null.Int `xorm:"int(11) pk notnull autoincr 's_t_j_id'" json:"s_t_j_id" form:"s_t_j_id" xml:"s_t_j_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// StuType defined
	StuType null.Int `xorm:"int(11) 'stu_type'" json:"stu_type" form:"stu_type" xml:"stu_type"`
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
	// StuTypeName defined
	StuTypeName null.String `xorm:"varchar(100) 'stu_type_name'" json:"stu_type_name" form:"stu_type_name" xml:"stu_type_name"`
	// Stutypedesc defined
	Stutypedesc null.String `xorm:"varchar(500) 'stutypedesc'" json:"stutypedesc" form:"stutypedesc" xml:"stutypedesc"`
	// OldStuType defined
	OldStuType null.Int `xorm:"int(11) 'old_stu_type'" json:"old_stu_type" form:"old_stu_type" xml:"old_stu_type"`
	// OldStuTypeName defined
	OldStuTypeName null.String `xorm:"varchar(1000) 'old_stu_type_name'" json:"old_stu_type_name" form:"old_stu_type_name" xml:"old_stu_type_name"`
}

// TableName table name of defined StudentTypeJournal
func (m *StudentTypeJournal) TableName() string {
	return "student_type_journal"
}

func (r *StudentTypeJournal) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStudentTypeJournal(data []byte) (StudentTypeJournal, error) {
	var r StudentTypeJournal
	err := json.Unmarshal(data, &r)
	return r, err
}
