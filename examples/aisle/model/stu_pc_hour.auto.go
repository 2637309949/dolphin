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

// StuPcHour defined
type StuPcHour struct {
	// SPHId defined
	SPHId null.Int `xorm:"int(11) pk notnull autoincr 's_p_h_id'" json:"s_p_h_id" form:"s_p_h_id" xml:"s_p_h_id"`
	// PcNumber defined
	PcNumber null.Int `xorm:"int(11) 'pc_number'" json:"pc_number" form:"pc_number" xml:"pc_number"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// Notpblsyhour defined
	Notpblsyhour null.Float `xorm:"float(11,2) 'notpblsyhour'" json:"notpblsyhour" form:"notpblsyhour" xml:"notpblsyhour"`
	// Pblsyhour defined
	Pblsyhour null.Float `xorm:"float(11,2) 'pblsyhour'" json:"pblsyhour" form:"pblsyhour" xml:"pblsyhour"`
	// Notpblgivesyhour defined
	Notpblgivesyhour null.Float `xorm:"float(11,2) 'notpblgivesyhour'" json:"notpblgivesyhour" form:"notpblgivesyhour" xml:"notpblgivesyhour"`
	// Pblgivesyhour defined
	Pblgivesyhour null.Float `xorm:"float(11,2) 'pblgivesyhour'" json:"pblgivesyhour" form:"pblgivesyhour" xml:"pblgivesyhour"`
	// BlGiveHour defined
	BlGiveHour null.Float `xorm:"float(11,2) 'bl_give_hour'" json:"bl_give_hour" form:"bl_give_hour" xml:"bl_give_hour"`
	// Notpblwxhmoney defined
	Notpblwxhmoney null.Float `xorm:"float(11,2) 'notpblwxhmoney'" json:"notpblwxhmoney" form:"notpblwxhmoney" xml:"notpblwxhmoney"`
	// Pblwxhmoney defined
	Pblwxhmoney null.Float `xorm:"float(11,2) 'pblwxhmoney'" json:"pblwxhmoney" form:"pblwxhmoney" xml:"pblwxhmoney"`
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
	// YessHour defined
	YessHour null.Float `xorm:"float(50,2) 'yess_hour'" json:"yess_hour" form:"yess_hour" xml:"yess_hour"`
	// YessGivehour defined
	YessGivehour null.Float `xorm:"float(50,2) 'yess_givehour'" json:"yess_givehour" form:"yess_givehour" xml:"yess_givehour"`
	// YessMoney defined
	YessMoney null.Float `xorm:"float(50,2) 'yess_money'" json:"yess_money" form:"yess_money" xml:"yess_money"`
	// PblZhuanYyhour defined
	PblZhuanYyhour null.Float `xorm:"float(50,2) 'pbl_zhuan_yyhour'" json:"pbl_zhuan_yyhour" form:"pbl_zhuan_yyhour" xml:"pbl_zhuan_yyhour"`
	// PcStateDate defined
	PcStateDate null.Time `xorm:"datetime 'pc_state_date'" json:"pc_state_date" form:"pc_state_date" xml:"pc_state_date"`
	// PcEndDate defined
	PcEndDate null.Time `xorm:"datetime 'pc_end_date'" json:"pc_end_date" form:"pc_end_date" xml:"pc_end_date"`
}

// With defined
func (m *StuPcHour) With(s interface{}) (interface{}, error) {
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
func (m *StuPcHour) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *StuPcHour) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *StuPcHour) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *StuPcHour) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *StuPcHour) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *StuPcHour) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined StuPcHour
func (m *StuPcHour) TableName() string {
	return "stu_pc_hour"
}
