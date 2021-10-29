// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// StuScore defined
type StuScore struct {
	// StuScoreId defined
	StuScoreId null.Int `xorm:"int(11) pk notnull autoincr 'stu_score_id'" json:"stu_score_id" form:"stu_score_id" xml:"stu_score_id"`
	// StuId defined
	StuId null.Int `xorm:"int(11) 'stu_id'" json:"stu_id" form:"stu_id" xml:"stu_id"`
	// OfId defined
	OfId null.Int `xorm:"int(11) 'of_id'" json:"of_id" form:"of_id" xml:"of_id"`
	// TmId defined
	TmId null.Int `xorm:"int(11) 'tm_id'" json:"tm_id" form:"tm_id" xml:"tm_id"`
	// LevelId defined
	LevelId null.Int `xorm:"int(11) 'level_id'" json:"level_id" form:"level_id" xml:"level_id"`
	// UnitId defined
	UnitId null.Int `xorm:"int(11) 'unit_id'" json:"unit_id" form:"unit_id" xml:"unit_id"`
	// TestTime defined
	TestTime null.Time `xorm:"datetime 'test_time'" json:"test_time" form:"test_time" xml:"test_time"`
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
	// AllScore defined
	AllScore null.Float `xorm:"float(50,2) 'all_score'" json:"all_score" form:"all_score" xml:"all_score"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// TauserId defined
	TauserId null.Int `xorm:"int(11) 'tauser_id'" json:"tauser_id" form:"tauser_id" xml:"tauser_id"`
}

// TableName table name of defined StuScore
func (m *StuScore) TableName() string {
	return "stu_score"
}

func (r *StuScore) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalStuScore(data []byte) (StuScore, error) {
	var r StuScore
	err := json.Unmarshal(data, &r)
	return r, err
}
