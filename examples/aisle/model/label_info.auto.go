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

// LabelInfo defined
type LabelInfo struct {
	// LabelInfoId defined
	LabelInfoId null.Int `xorm:"int(11) pk notnull autoincr 'label_info_id'" json:"label_info_id" form:"label_info_id" xml:"label_info_id"`
	// ParentId defined
	ParentId null.Int `xorm:"int(11) 'parent_id'" json:"parent_id" form:"parent_id" xml:"parent_id"`
	// LabelName defined
	LabelName null.String `xorm:"varchar(100) 'label_name'" json:"label_name" form:"label_name" xml:"label_name"`
	// LabelRemark defined
	LabelRemark null.String `xorm:"varchar(500) 'label_remark'" json:"label_remark" form:"label_remark" xml:"label_remark"`
	// PtId defined
	PtId null.Int `xorm:"int(11) 'pt_id'" json:"pt_id" form:"pt_id" xml:"pt_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
}

// Marshal defined
func (m *LabelInfo) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *LabelInfo) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *LabelInfo) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *LabelInfo) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *LabelInfo) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *LabelInfo) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined LabelInfo
func (m *LabelInfo) TableName() string {
	return "label_info"
}
