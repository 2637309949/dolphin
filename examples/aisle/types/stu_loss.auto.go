// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// StuLoss defined
type StuLoss struct {
	// StuLossId defined
	StuLossId null.Int `xorm:"int(11) pk notnull autoincr 'stu_loss_id'" json:"stu_loss_id" form:"stu_loss_id" xml:"stu_loss_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// LossDesc defined
	LossDesc null.String `xorm:"varchar(5000) 'loss_desc'" json:"loss_desc" form:"loss_desc" xml:"loss_desc"`
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

// TableName table name of defined StuLoss
func (m *StuLoss) TableName() string {
	return "stu_loss"
}

func (r *StuLoss) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStuLoss(data []byte) (StuLoss, error) {
	var r StuLoss
	err := json.Unmarshal(data, &r)
	return r, err
}
