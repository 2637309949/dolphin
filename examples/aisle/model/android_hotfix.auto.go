// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"github.com/2637309949/dolphin/packages/null"
)

// AndroidHotfix defined
type AndroidHotfix struct {
	// AHId defined
	AHId null.Int `xorm:"int(11) pk notnull autoincr 'a_h_id'" json:"a_h_id" form:"a_h_id" xml:"a_h_id"`
	// AppName defined
	AppName null.String `xorm:"varchar(200) 'app_name'" json:"app_name" form:"app_name" xml:"app_name"`
	// Applicationid defined
	Applicationid null.String `xorm:"varchar(200) 'applicationid'" json:"applicationid" form:"applicationid" xml:"applicationid"`
	// VersionName defined
	VersionName null.String `xorm:"varchar(200) 'version_name'" json:"version_name" form:"version_name" xml:"version_name"`
	// Descriptions defined
	Descriptions null.String `xorm:"varchar(500) 'descriptions'" json:"descriptions" form:"descriptions" xml:"descriptions"`
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

// TableName table name of defined AndroidHotfix
func (m *AndroidHotfix) TableName() string {
	return "android_hotfix"
}
