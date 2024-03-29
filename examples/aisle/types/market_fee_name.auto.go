// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// MarketFeeName defined
type MarketFeeName struct {
	// MFNId defined
	MFNId null.Int `xorm:"int(11) pk notnull autoincr 'm_f_n_id'" json:"m_f_n_id" form:"m_f_n_id" xml:"m_f_n_id"`
	// Project defined
	Project null.Int `xorm:"int(11) 'project'" json:"project" form:"project" xml:"project"`
	// Name defined
	Name null.String `xorm:"varchar(200) 'name'" json:"name" form:"name" xml:"name"`
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

// TableName table name of defined MarketFeeName
func (m *MarketFeeName) TableName() string {
	return "market_fee_name"
}

func (r *MarketFeeName) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMarketFeeName(data []byte) (MarketFeeName, error) {
	var r MarketFeeName
	err := json.Unmarshal(data, &r)
	return r, err
}
