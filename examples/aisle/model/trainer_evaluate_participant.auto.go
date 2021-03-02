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

// TrainerEvaluateParticipant defined
type TrainerEvaluateParticipant struct {
	// TEPId defined
	TEPId null.Int `xorm:"int(11) pk notnull autoincr 't_e_p_id'" json:"t_e_p_id" form:"t_e_p_id" xml:"t_e_p_id"`
	// TrainerId defined
	TrainerId null.Int `xorm:"int(11) 'trainer_id'" json:"trainer_id" form:"trainer_id" xml:"trainer_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// Isdelete defined
	Isdelete null.Int `xorm:"notnull 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// EvaluateRemark defined
	EvaluateRemark null.String `xorm:"varchar(300) 'evaluate_remark'" json:"evaluate_remark" form:"evaluate_remark" xml:"evaluate_remark"`
}

// Marshal defined
func (m *TrainerEvaluateParticipant) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *TrainerEvaluateParticipant) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *TrainerEvaluateParticipant) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *TrainerEvaluateParticipant) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TrainerEvaluateParticipant
func (m *TrainerEvaluateParticipant) TableName() string {
	return "trainer_evaluate_participant"
}
