// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// AboutUs defined
type AboutUs struct {
	// AboutUsId defined
	AboutUsId null.Int `xorm:"int(11) pk notnull autoincr 'about_us_id'" json:"about_us_id" form:"about_us_id" xml:"about_us_id"`
	// AboutUsPicture defined
	AboutUsPicture null.Int `xorm:"int(11) 'about_us_picture'" json:"about_us_picture" form:"about_us_picture" xml:"about_us_picture"`
	// AboutUsContent defined
	AboutUsContent null.String `xorm:"varchar(10000) 'about_us_content'" json:"about_us_content" form:"about_us_content" xml:"about_us_content"`
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
	// AboutUsVideo defined
	AboutUsVideo null.Int `xorm:"int(11) 'about_us_video'" json:"about_us_video" form:"about_us_video" xml:"about_us_video"`
	// AboutUsCntitle defined
	AboutUsCntitle null.String `xorm:"varchar(1000) 'about_us_cntitle'" json:"about_us_cntitle" form:"about_us_cntitle" xml:"about_us_cntitle"`
	// AboutUsEntitle defined
	AboutUsEntitle null.String `xorm:"varchar(1000) 'about_us_entitle'" json:"about_us_entitle" form:"about_us_entitle" xml:"about_us_entitle"`
}

// TableName table name of defined AboutUs
func (m *AboutUs) TableName() string {
	return "about_us"
}

func (r *AboutUs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalAboutUs(data []byte) (AboutUs, error) {
	var r AboutUs
	err := json.Unmarshal(data, &r)
	return r, err
}
