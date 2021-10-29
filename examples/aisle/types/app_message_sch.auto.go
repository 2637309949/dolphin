// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// AppMessageSch defined
type AppMessageSch struct {
	// AMSId defined
	AMSId null.Int `xorm:"int(11) pk notnull autoincr 'a_m_s_id'" json:"a_m_s_id" form:"a_m_s_id" xml:"a_m_s_id"`
	// AmnId defined
	AmnId null.Int `xorm:"int(11) 'amn_id'" json:"amn_id" form:"amn_id" xml:"amn_id"`
	// OsId defined
	OsId null.Int `xorm:"int(11) 'os_id'" json:"os_id" form:"os_id" xml:"os_id"`
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

// TableName table name of defined AppMessageSch
func (m *AppMessageSch) TableName() string {
	return "app_message_sch"
}

func (r *AppMessageSch) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAppMessageSch(data []byte) (AppMessageSch, error) {
	var r AppMessageSch
	err := json.Unmarshal(data, &r)
	return r, err
}
