// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ClassroomTimeResource defined
type ClassroomTimeResource struct {
	// CTRId defined
	CTRId null.Int `xorm:"int(11) pk notnull autoincr 'c_t_r_id'" json:"c_t_r_id" form:"c_t_r_id" xml:"c_t_r_id"`
	// DataId defined
	DataId null.Float `xorm:"float(11,2) 'data_id'" json:"data_id" form:"data_id" xml:"data_id"`
	// TableId defined
	TableId null.Float `xorm:"float(11,2) 'table_id'" json:"table_id" form:"table_id" xml:"table_id"`
	// MainDataId defined
	MainDataId null.Int `xorm:"int(11) 'main_data_id'" json:"main_data_id" form:"main_data_id" xml:"main_data_id"`
	// StartTime defined
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	// EndTime defined
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" form:"end_time" xml:"end_time"`
	// ConflictContent defined
	ConflictContent null.String `xorm:"varchar(500) 'conflict_content'" json:"conflict_content" form:"conflict_content" xml:"conflict_content"`
	// ShowContent defined
	ShowContent null.String `xorm:"varchar(500) 'show_content'" json:"show_content" form:"show_content" xml:"show_content"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// CourseType defined
	CourseType null.Int `xorm:"int(11) 'course_type'" json:"course_type" form:"course_type" xml:"course_type"`
	// CourseState defined
	CourseState null.Int `xorm:"int(11) 'course_state'" json:"course_state" form:"course_state" xml:"course_state"`
}

// TableName table name of defined ClassroomTimeResource
func (m *ClassroomTimeResource) TableName() string {
	return "classroom_time_resource"
}

func (r *ClassroomTimeResource) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalClassroomTimeResource(data []byte) (ClassroomTimeResource, error) {
	var r ClassroomTimeResource
	err := json.Unmarshal(data, &r)
	return r, err
}
