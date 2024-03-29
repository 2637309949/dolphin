// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// Sheng defined
type Sheng struct {
	// ShengId defined
	ShengId null.Int `xorm:"int(11) pk notnull autoincr 'sheng_id'" json:"sheng_id" form:"sheng_id" xml:"sheng_id"`
	// ShengName defined
	ShengName null.String `xorm:"varchar(20) 'sheng_name'" json:"sheng_name" form:"sheng_name" xml:"sheng_name"`
	// ShengCode defined
	ShengCode null.String `xorm:"varchar(30) 'sheng_code'" json:"sheng_code" form:"sheng_code" xml:"sheng_code"`
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
	// ShengArea defined
	ShengArea null.Int `xorm:"int(11) 'sheng_area'" json:"sheng_area" form:"sheng_area" xml:"sheng_area"`
}

// TableName table name of defined Sheng
func (m *Sheng) TableName() string {
	return "sheng"
}

func (r *Sheng) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalSheng(data []byte) (Sheng, error) {
	var r Sheng
	err := json.Unmarshal(data, &r)
	return r, err
}
