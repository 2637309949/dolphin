// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// Phonehistory defined
type Phonehistory struct {
	// PhId defined
	PhId null.Int `xorm:"int(11) pk notnull autoincr 'ph_id'" json:"ph_id" form:"ph_id" xml:"ph_id"`
	// Callno defined
	Callno null.String `xorm:"varchar(11) 'callno'" json:"callno" form:"callno" xml:"callno"`
	// Calledno defined
	Calledno null.String `xorm:"varchar(11) 'calledno'" json:"calledno" form:"calledno" xml:"calledno"`
	// Begintime defined
	Begintime null.Time `xorm:"datetime 'begintime'" json:"begintime" form:"begintime" xml:"begintime"`
	// Endtime defined
	Endtime null.Time `xorm:"datetime 'endtime'" json:"endtime" form:"endtime" xml:"endtime"`
	// Exten defined
	Exten null.String `xorm:"varchar(100) 'exten'" json:"exten" form:"exten" xml:"exten"`
	// State defined
	State null.String `xorm:"varchar(100) 'state'" json:"state" form:"state" xml:"state"`
	// Recordfile defined
	Recordfile null.String `xorm:"varchar(100) 'recordfile'" json:"recordfile" form:"recordfile" xml:"recordfile"`
	// Fileserver defined
	Fileserver null.String `xorm:"varchar(100) 'fileserver'" json:"fileserver" form:"fileserver" xml:"fileserver"`
	// Province defined
	Province null.String `xorm:"varchar(100) 'province'" json:"province" form:"province" xml:"province"`
	// District defined
	District null.String `xorm:"varchar(100) 'district'" json:"district" form:"district" xml:"district"`
	// Filemongoid defined
	Filemongoid null.String `xorm:"varchar(100) 'filemongoid'" json:"filemongoid" form:"filemongoid" xml:"filemongoid"`
	// Dbfileid defined
	Dbfileid null.Int `xorm:"int(11) 'dbfileid'" json:"dbfileid" form:"dbfileid" xml:"dbfileid"`
	// PhoneTalbeid defined
	PhoneTalbeid null.Float `xorm:"float(11,2) 'phone_talbeid'" json:"phone_talbeid" form:"phone_talbeid" xml:"phone_talbeid"`
	// PhoneDataid defined
	PhoneDataid null.Float `xorm:"float(11,2) 'phone_dataid'" json:"phone_dataid" form:"phone_dataid" xml:"phone_dataid"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
}

// TableName table name of defined Phonehistory
func (m *Phonehistory) TableName() string {
	return "phonehistory"
}

func (r *Phonehistory) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPhonehistory(data []byte) (Phonehistory, error) {
	var r Phonehistory
	err := json.Unmarshal(data, &r)
	return r, err
}
