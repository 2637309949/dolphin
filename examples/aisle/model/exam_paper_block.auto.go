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

// ExamPaperBlock defined
type ExamPaperBlock struct {
	// EPBId defined
	EPBId null.Int `xorm:"int(11) pk notnull autoincr 'e_p_b_id'" json:"e_p_b_id" form:"e_p_b_id" xml:"e_p_b_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// EpId defined
	EpId null.Int `xorm:"int(11) 'ep_id'" json:"ep_id" form:"ep_id" xml:"ep_id"`
	// BlockName defined
	BlockName null.String `xorm:"varchar(50) 'block_name'" json:"block_name" form:"block_name" xml:"block_name"`
	// BlockTimeLength defined
	BlockTimeLength null.Int `xorm:"int(11) 'block_time_length'" json:"block_time_length" form:"block_time_length" xml:"block_time_length"`
	// BlockQuesNum defined
	BlockQuesNum null.Int `xorm:"int(11) 'block_ques_num'" json:"block_ques_num" form:"block_ques_num" xml:"block_ques_num"`
	// BlockScore defined
	BlockScore null.Int `xorm:"int(11) 'block_score'" json:"block_score" form:"block_score" xml:"block_score"`
	// BlockOrder defined
	BlockOrder null.Int `xorm:"int(11) 'block_order'" json:"block_order" form:"block_order" xml:"block_order"`
	// BlockRemark defined
	BlockRemark null.String `xorm:"varchar(200) 'block_remark'" json:"block_remark" form:"block_remark" xml:"block_remark"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Marshal defined
func (m *ExamPaperBlock) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ExamPaperBlock) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *ExamPaperBlock) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ExamPaperBlock) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ExamPaperBlock
func (m *ExamPaperBlock) TableName() string {
	return "exam_paper_block"
}
