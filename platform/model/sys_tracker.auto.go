// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// SysTracker defined 日志跟踪
type SysTracker struct {
	// ID defined 主键
	ID null.String `xorm:"varchar(36) notnull unique pk comment('主键') 'id'" json:"id" form:"id" xml:"id"`
	// Token defined 令牌
	Token null.String `xorm:"varchar(72) comment('令牌') 'token'" json:"token" form:"token" xml:"token"`
	// UserId defined 用户ID
	UserId null.String `xorm:"varchar(36) comment('用户ID') 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// StatusCode defined 状态
	StatusCode null.Int `xorm:"int comment('状态') 'status_code'" json:"status_code" form:"status_code" xml:"status_code"`
	// Latency defined 反应时间
	Latency null.Float `xorm:"float comment('反应时间') 'latency'" json:"latency" form:"latency" xml:"latency"`
	// ClientIp defined 客户端id
	ClientIp null.String `xorm:"varchar(36) comment('客户端id') 'client_ip'" json:"client_ip" form:"client_ip" xml:"client_ip"`
	// Method defined 请求方法
	Method null.String `xorm:"varchar(36) comment('请求方法') 'method'" json:"method" form:"method" xml:"method"`
	// Path defined 请求路径
	Path null.String `xorm:"varchar(512) comment('请求路径') 'path'" json:"path" form:"path" xml:"path"`
	// Header defined 请求头
	Header []byte `xorm:"blob comment('请求头') 'header'" json:"header" form:"header" xml:"header"`
	// ReqBody defined 请求体
	ReqBody []byte `xorm:"mediumblob comment('请求体') 'req_body'" json:"req_body" form:"req_body" xml:"req_body"`
	// ResBody defined 返回体
	ResBody []byte `xorm:"mediumblob comment('返回体') 'res_body'" json:"res_body" form:"res_body" xml:"res_body"`
	// Creater defined 创建人
	Creater null.String `xorm:"varchar(36) notnull comment('创建人') 'creater'" json:"creater" form:"creater" xml:"creater"`
	// CreateTime defined 创建时间
	CreateTime null.Time `xorm:"datetime notnull comment('创建时间') 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// Updater defined 最后更新人
	Updater null.String `xorm:"varchar(36) notnull comment('最后更新人') 'updater'" json:"updater" form:"updater" xml:"updater"`
	// UpdateTime defined 最后更新时间
	UpdateTime null.Time `xorm:"datetime notnull comment('最后更新时间') 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
	// IsDelete defined 删除标记
	IsDelete null.Int `xorm:"notnull comment('删除标记') 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
	// Remark defined 备注
	Remark null.String `xorm:"varchar(200) comment('备注') 'remark'" json:"remark" form:"remark" xml:"remark"`
}

// With defined
func (m *SysTracker) With(s interface{}) (interface{}, error) {
	if reflect.ValueOf(s).Kind() != reflect.Ptr {
		return nil, errors.New("ptr required")
	}
	mbt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(mbt, s); err != nil {
		return nil, err
	}
	return s, err
}

// Marshal defined
func (m *SysTracker) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Unmarshal defined
func (m *SysTracker) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// ToMap defined
func (m *SysTracker) ToMap() (map[string]interface{}, error) {
	byt, err := m.Marshal()
	if err != nil {
		return nil, err
	}
	itf := map[string]interface{}{}
	err = json.Unmarshal(byt, &itf)
	return itf, err
}

// FromMap defined
func (m *SysTracker) FromMap(fm map[string]interface{}) error {
	byt, err := json.Marshal(fm)
	if err != nil {
		return err
	}
	err = m.Unmarshal(byt)
	return err
}

// Parser defined
func (m *SysTracker) Parser(db *xorm.Engine) *tags.Parser {
	dialect, mapper, cache := db.Dialect(), db.DB().Mapper, caches.NewManager()
	return tags.NewParser("xorm", dialect, mapper, mapper, cache)
}

// PrimaryKeys defined
func (m *SysTracker) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined SysTracker
func (m *SysTracker) TableName() string {
	return "sys_tracker"
}
