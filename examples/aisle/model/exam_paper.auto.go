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

// ExamPaper defined
type ExamPaper struct {
	// ExamPaperId defined
	ExamPaperId null.Int `xorm:"int(11) pk notnull autoincr 'exam_paper_id'" json:"exam_paper_id" form:"exam_paper_id" xml:"exam_paper_id"`
	// Remark defined
	Remark null.String `xorm:"varchar(200) 'remark'" json:"remark" form:"remark" xml:"remark"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// EpName defined
	EpName null.String `xorm:"varchar(50) 'ep_name'" json:"ep_name" form:"ep_name" xml:"ep_name"`
	// EpType defined
	EpType null.Int `xorm:"int(11) 'ep_type'" json:"ep_type" form:"ep_type" xml:"ep_type"`
	// EpDuration defined
	EpDuration null.Float `xorm:"float(11,2) 'ep_duration'" json:"ep_duration" form:"ep_duration" xml:"ep_duration"`
	// EpTotalScore defined
	EpTotalScore null.Int `xorm:"int(11) 'ep_total_score'" json:"ep_total_score" form:"ep_total_score" xml:"ep_total_score"`
	// PassingScore defined
	PassingScore null.Int `xorm:"int(11) 'passing_score'" json:"passing_score" form:"passing_score" xml:"passing_score"`
	// EqNumber defined
	EqNumber null.Int `xorm:"int(11) 'eq_number'" json:"eq_number" form:"eq_number" xml:"eq_number"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Marshal defined
func (m *ExamPaper) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *ExamPaper) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *ExamPaper) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *ExamPaper) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *ExamPaper) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *ExamPaper) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined ExamPaper
func (m *ExamPaper) TableName() string {
	return "exam_paper"
}
