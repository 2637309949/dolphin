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

// TrialLesson defined
type TrialLesson struct {
	// TLId defined
	TLId null.Int `xorm:"int(11) pk notnull autoincr 't_l_id'" json:"t_l_id" form:"t_l_id" xml:"t_l_id"`
	// TlName defined
	TlName null.String `xorm:"varchar(50) 'tl_name'" json:"tl_name" form:"tl_name" xml:"tl_name"`
	// StartTime defined
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	// EndTime defined
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" form:"end_time" xml:"end_time"`
	// TlState defined
	TlState null.Int `xorm:"int(11) 'tl_state'" json:"tl_state" form:"tl_state" xml:"tl_state"`
	// MaxNum defined
	MaxNum null.Int `xorm:"int(11) 'max_num'" json:"max_num" form:"max_num" xml:"max_num"`
	// TeaId defined
	TeaId null.Int `xorm:"int(11) 'tea_id'" json:"tea_id" form:"tea_id" xml:"tea_id"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// TlDate defined
	TlDate null.Time `xorm:"datetime 'tl_date'" json:"tl_date" form:"tl_date" xml:"tl_date"`
	// RoomId defined
	RoomId null.Int `xorm:"int(11) 'room_id'" json:"room_id" form:"room_id" xml:"room_id"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// SchId defined
	SchId null.Int `xorm:"int(11) 'sch_id'" json:"sch_id" form:"sch_id" xml:"sch_id"`
	// LessonType defined
	LessonType null.Int `xorm:"int(11) 'lesson_type'" json:"lesson_type" form:"lesson_type" xml:"lesson_type"`
	// NowNum defined
	NowNum null.Float `xorm:"float(10,2) 'now_num'" json:"now_num" form:"now_num" xml:"now_num"`
	// TlNumState defined
	TlNumState null.Int `xorm:"int(11) 'tl_num_state'" json:"tl_num_state" form:"tl_num_state" xml:"tl_num_state"`
	// Remark defined
	Remark null.String `xorm:"varchar(2000) 'remark'" json:"remark" form:"remark" xml:"remark"`
	// RealTlDate defined
	RealTlDate null.Time `xorm:"datetime 'real_tl_date'" json:"real_tl_date" form:"real_tl_date" xml:"real_tl_date"`
	// RealEndTime defined
	RealEndTime null.Time `xorm:"datetime 'real_end_time'" json:"real_end_time" form:"real_end_time" xml:"real_end_time"`
	// RealStartTime defined
	RealStartTime null.Time `xorm:"datetime 'real_start_time'" json:"real_start_time" form:"real_start_time" xml:"real_start_time"`
	// TlType defined
	TlType null.Int `xorm:"int(11) 'tl_type'" json:"tl_type" form:"tl_type" xml:"tl_type"`
}

// Marshal defined
func (m *TrialLesson) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *TrialLesson) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Parser defined
func (m *TrialLesson) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *TrialLesson) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined TrialLesson
func (m *TrialLesson) TableName() string {
	return "trial_lesson"
}
