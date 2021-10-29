// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// CourseVideo defined
type CourseVideo struct {
	// CVId defined
	CVId null.Int `xorm:"int(11) pk notnull autoincr 'c_v_id'" json:"c_v_id" form:"c_v_id" xml:"c_v_id"`
	// CsId defined
	CsId null.Int `xorm:"int(11) 'cs_id'" json:"cs_id" form:"cs_id" xml:"cs_id"`
	// CourseVideoDesc defined
	CourseVideoDesc null.String `xorm:"varchar(10000) 'course_video_desc'" json:"course_video_desc" form:"course_video_desc" xml:"course_video_desc"`
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

// TableName table name of defined CourseVideo
func (m *CourseVideo) TableName() string {
	return "course_video"
}

func (r *CourseVideo) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCourseVideo(data []byte) (CourseVideo, error) {
	var r CourseVideo
	err := json.Unmarshal(data, &r)
	return r, err
}
