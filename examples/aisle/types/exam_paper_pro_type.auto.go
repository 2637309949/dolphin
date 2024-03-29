// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ExamPaperProType defined
type ExamPaperProType struct {
	// EPPTId defined
	EPPTId null.Int `xorm:"int(11) pk notnull autoincr 'e_p_p_t_id'" json:"e_p_p_t_id" form:"e_p_p_t_id" xml:"e_p_p_t_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// EpId defined
	EpId null.Int `xorm:"int(11) 'ep_id'" json:"ep_id" form:"ep_id" xml:"ep_id"`
	// PtId defined
	PtId null.Int `xorm:"int(11) 'pt_id'" json:"pt_id" form:"pt_id" xml:"pt_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined ExamPaperProType
func (m *ExamPaperProType) TableName() string {
	return "exam_paper_pro_type"
}

func (r *ExamPaperProType) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalExamPaperProType(data []byte) (ExamPaperProType, error) {
	var r ExamPaperProType
	err := json.Unmarshal(data, &r)
	return r, err
}
