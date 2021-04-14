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
			ID:         null.StringFrom("c6a7bt50-7dad-31d1-81b5-10c34fd460e2"),
			Name:       null.StringFrom("域名启用状态"),
			Code:       null.StringFrom("sys_domain_status"),
			Value:      null.StringFrom(`[{"text":"禁用","value":0},{"text":"正常","value":1}]`),
			Creater:    DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:         null.StringFrom("b6a7ct5q-2dad-31d1-81b5-10c34fd460c3"),
			Name:       null.StringFrom("域名认证模式"),
			Code:       null.StringFrom("sys_domain_auth_mode"),
			Value:      null.StringFrom(`[{"text":"集成登录","value":0},{"text":"单点登录","value":1}]`),
			Creater:    DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:         null.StringFrom("c6a7bt50-7dad-31d1-81b5-10c34fd460c2"),
			Name:       null.StringFrom("域名数据库标志"),
			Code:       null.StringFrom("sys_domain_is_sync"),
			Value:      null.StringFrom(`[{"text":"未同步","value":0},{"text":"已同步","value":1}]`),
			Creater:    DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:         null.StringFrom("b4a7bt50-7dad-12d1-81b5-10c34fd460e4"),
			Name:       null.StringFrom("帐号启用状态"),
			Code:       null.StringFrom("sys_user_status"),
			Value:      null.StringFrom(`[{"text":"禁用","value":0},{"text":"正常","value":1}]`),
			Creater:    DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:         null.StringFrom("a1b7b150-6dad-12d1-81b5-10c34fd460e5"),
			Name:       null.StringFrom("菜单类型"),
			Code:       null.StringFrom("sys_menu_type"),
			Value:      null.StringFrom(`[{"text":"Dir","value":0},{"text":"Menu","value":1},{"text":"Button","value":2}]`),
			Creater:    DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
		{
			ID:         null.StringFrom("c2b5b260-7jad-12d1-81b5-10c34fd460e6"),
			Name:       null.StringFrom("持久类型"),
			Code:       null.StringFrom("sys_attachment_durable"),
			Value:      null.StringFrom(`[{"text":"否","value":0},{"text":"是","value":1}]`),
			Creater:    DefaultAdmin.ID,
			CreateTime: null.TimeFrom(time.Now()),
			Updater:    DefaultAdmin.ID,
			UpdateTime: null.TimeFrom(time.Now()),
			IsDelete:   null.IntFrom(0),
		},
	}
	for _, option := range options {
		if ct, err := s.Where("id=?", option.ID.String).Count(new(SysOptionset)); ct == 0 || err != nil {
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
