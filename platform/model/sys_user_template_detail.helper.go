// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// sysUserTemplateDetails default
var sysUserTemplateDetails = []SysUserTemplateDetail{
	SysUserTemplateDetail{
		ID:         null.StringFrom("7c3da436-2772-48da-86d8-97b2bd80e391"),
		Name:       null.StringFrom("爱好"),
		Value:      null.StringFrom("Hobby"),
		TempId:     DefaultUserTemplate.ID,
		Type:       null.IntFrom(2),
		Content:    null.StringFrom(""),
		Priority:   null.IntFrom(1),
		CreateBy:   DefaultAdmin.ID,
		CreateTime: null.TimeFrom(time.Now()),
		UpdateBy:   DefaultAdmin.ID,
		UpdateTime: null.TimeFrom(time.Now()),
		DelFlag:    null.IntFrom(0),
	},
}

// InitSysData defined inital system data
func (m *SysUserTemplateDetail) InitSysData(s *xorm.Session) {
	for _, detail := range sysUserTemplateDetails {
		if ct, err := s.Where("id=?", detail.ID.String).Count(new(SysUserTemplateDetail)); ct == 0 || err != nil {
			if err != nil {
				s.Rollback()
				panic(err)
			}
			if _, err := s.InsertOne(&detail); err != nil {
				s.Rollback()
				panic(err)
			}
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
