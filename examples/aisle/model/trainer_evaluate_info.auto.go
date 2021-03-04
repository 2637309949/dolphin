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

// TrainerEvaluateInfo defined
type TrainerEvaluateInfo struct {
	// TEIId defined
	TEIId null.Int `xorm:"int(11) pk notnull autoincr 't_e_i_id'" json:"t_e_i_id" form:"t_e_i_id" xml:"t_e_i_id"`
	// TrainerId defined
	TrainerId null.Int `xorm:"int(11) 'trainer_id'" json:"trainer_id" form:"trainer_id" xml:"trainer_id"`
	// Remark defined
	Remark null.String `xorm:"varchar(2000) 'remark'" json:"remark" form:"remark" xml:"remark"`
	// PetId defined
	PetId null.Int `xorm:"int(11) 'pet_id'" json:"pet_id" form:"pet_id" xml:"pet_id"`
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
}

// Marshal defined
func (m *TrainerEvaluateInfo) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *TrainerEvaluateInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *TrainerEvaluateInfo) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *TrainerEvaluateInfo) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *TrainerEvaluateInfo) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *TrainerEvaluateInfo) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TrainerEvaluateInfo
func (m *TrainerEvaluateInfo) TableName() string {
	return "trainer_evaluate_info"
}
