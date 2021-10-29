// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// CssChageStudent defined
type CssChageStudent struct {
	// CCSId defined
	CCSId null.Int `xorm:"int(11) pk notnull autoincr 'c_c_s_id'" json:"c_c_s_id" form:"c_c_s_id" xml:"c_c_s_id"`
	// PkCss defined
	PkCss null.Int `xorm:"int(11) 'pk_css'" json:"pk_css" form:"pk_css" xml:"pk_css"`
	// PkSct defined
	PkSct null.Int `xorm:"int(11) 'pk_sct'" json:"pk_sct" form:"pk_sct" xml:"pk_sct"`
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
	// PkOldStu defined
	PkOldStu null.Int `xorm:"int(11) 'pk_old_stu'" json:"pk_old_stu" form:"pk_old_stu" xml:"pk_old_stu"`
	// PkNewStu defined
	PkNewStu null.Int `xorm:"int(11) 'pk_new_stu'" json:"pk_new_stu" form:"pk_new_stu" xml:"pk_new_stu"`
	// PkCs defined
	PkCs null.Int `xorm:"int(11) 'pk_cs'" json:"pk_cs" form:"pk_cs" xml:"pk_cs"`
	// PkOldSct defined
	PkOldSct null.Int `xorm:"int(11) 'pk_old_sct'" json:"pk_old_sct" form:"pk_old_sct" xml:"pk_old_sct"`
	// IfKouHour defined
	IfKouHour null.Int `xorm:"int(11) 'if_kou_hour'" json:"if_kou_hour" form:"if_kou_hour" xml:"if_kou_hour"`
	// IfMz defined
	IfMz null.Int `xorm:"int(11) 'if_mz'" json:"if_mz" form:"if_mz" xml:"if_mz"`
	// OldScsId defined
	OldScsId null.Int `xorm:"int(11) 'old_scs_id'" json:"old_scs_id" form:"old_scs_id" xml:"old_scs_id"`
}

// TableName table name of defined CssChageStudent
func (m *CssChageStudent) TableName() string {
	return "css_chage_student"
}

func (r *CssChageStudent) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalCssChageStudent(data []byte) (CssChageStudent, error) {
	var r CssChageStudent
	err := json.Unmarshal(data, &r)
	return r, err
}
