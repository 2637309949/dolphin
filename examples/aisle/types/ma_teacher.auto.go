// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// MaTeacher defined
type MaTeacher struct {
	// MaTeacherId defined
	MaTeacherId null.Int `xorm:"int(11) pk notnull autoincr 'ma_teacher_id'" json:"ma_teacher_id" form:"ma_teacher_id" xml:"ma_teacher_id"`
	// MaId defined
	MaId null.Int `xorm:"int(11) 'ma_id'" json:"ma_id" form:"ma_id" xml:"ma_id"`
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
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// TeaClassDate defined
	TeaClassDate null.Time `xorm:"datetime 'tea_class_date'" json:"tea_class_date" form:"tea_class_date" xml:"tea_class_date"`
	// TeaBegintime defined
	TeaBegintime null.Time `xorm:"datetime 'tea_begintime'" json:"tea_begintime" form:"tea_begintime" xml:"tea_begintime"`
	// TeaEndtime defined
	TeaEndtime null.Time `xorm:"datetime 'tea_endtime'" json:"tea_endtime" form:"tea_endtime" xml:"tea_endtime"`
	// IfKouHour defined
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" form:"if_kou_hour" xml:"if_kou_hour"`
	// ClassHour defined
	ClassHour null.Float `xorm:"float(50,2) 'class_hour'" json:"class_hour" form:"class_hour" xml:"class_hour"`
}

// TableName table name of defined MaTeacher
func (m *MaTeacher) TableName() string {
	return "ma_teacher"
}

func (r *MaTeacher) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMaTeacher(data []byte) (MaTeacher, error) {
	var r MaTeacher
	err := json.Unmarshal(data, &r)
	return r, err
}
