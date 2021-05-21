package model

import (
	"encoding/json"
	"time"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
)

// Values defined
func (m *SysOptionset) Values() []struct {
	Text  string      `json:"text"`
	Value interface{} `json:"value"`
} {
	kv := []struct {
		Text  string      `json:"text"`
		Value interface{} `json:"value"`
	}{}
	json.Unmarshal([]byte(m.Value.String), &kv)
	return kv
}

// InitSysData defined
func (m *SysOptionset) InitSysData(s *xorm.Session) {
	options := []SysOptionset{
		{
			Name:       null.StringFrom("域名启用状态"),
			Code:       null.StringFrom("sys_domain_status"),
			Value:      null.StringFrom(`[{"text":"禁用","value":0},{"text":"正常","value":1}]`),
			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			Name:       null.StringFrom("域名认证模式"),
			Code:       null.StringFrom("sys_domain_auth_mode"),
			Value:      null.StringFrom(`[{"text":"集成登录","value":0},{"text":"单点登录","value":1}]`),
			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			Name:       null.StringFrom("域名数据库标志"),
			Code:       null.StringFrom("sys_domain_is_sync"),
			Value:      null.StringFrom(`[{"text":"未同步","value":0},{"text":"已同步","value":1}]`),
			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			Name:       null.StringFrom("帐号启用状态"),
			Code:       null.StringFrom("sys_user_status"),
			Value:      null.StringFrom(`[{"text":"禁用","value":0},{"text":"正常","value":1}]`),
			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			Name:       null.StringFrom("角色启用状态"),
			Code:       null.StringFrom("sys_role_status"),
			Value:      null.StringFrom(`[{"text":"禁用","value":0},{"text":"正常","value":1}]`),
			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			Name:       null.StringFrom("菜单类型"),
			Code:       null.StringFrom("sys_menu_type"),
			Value:      null.StringFrom(`[{"text":"Dir","value":0},{"text":"Menu","value":1},{"text":"Button","value":2}]`),
			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			Name:       null.StringFrom("持久类型"),
			Code:       null.StringFrom("sys_attachment_durable"),
			Value:      null.StringFrom(`[{"text":"否","value":0},{"text":"是","value":1}]`),
			Creater:    null.IntFrom(1),
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    null.IntFrom(1),
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
	}
	for _, option := range options {
		if ct, err := s.Where("code=?", option.Code.String).Count(new(SysOptionset)); ct == 0 || err != nil {
			if err != nil {
				s.Rollback()
				panic(err)
			}
			if _, err := s.InsertOne(&option); err != nil {
				s.Rollback()
				panic(err)
			}
		}
	}
	if err := s.Commit(); err != nil {
		panic(err)
	}
}
