// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// NetworkDetail defined
type NetworkDetail struct {
	// NDId defined
	NDId null.Int `xorm:"int(11) pk notnull autoincr 'n_d_id'" json:"n_d_id" form:"n_d_id" xml:"n_d_id"`
	// QdDetail defined
	QdDetail null.String `xorm:"varchar(100) 'qd_detail'" json:"qd_detail" form:"qd_detail" xml:"qd_detail"`
	// Remake defined
	Remake null.String `xorm:"varchar(500) 'remake'" json:"remake" form:"remake" xml:"remake"`
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
	// Qudao defined
	Qudao null.Int `xorm:"int(11) 'qudao'" json:"qudao" form:"qudao" xml:"qudao"`
	// ChannelType defined
	ChannelType null.Int `xorm:"int(11) 'channel_type'" json:"channel_type" form:"channel_type" xml:"channel_type"`
	// QueryType defined
	QueryType null.Int `xorm:"int(11) 'query_type'" json:"query_type" form:"query_type" xml:"query_type"`
}

// With defined
func (m *NetworkDetail) With(s interface{}) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return nil, errors.New("ptr required")
	}
	mbt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		return nil, err
	}
	return s, err
}

// Marshal defined
func (m *NetworkDetail) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *NetworkDetail) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *NetworkDetail) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *NetworkDetail) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *NetworkDetail) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *NetworkDetail) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined NetworkDetail
func (m *NetworkDetail) TableName() string {
	return "network_detail"
}
