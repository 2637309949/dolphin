// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// T2LoginLog defined
type T2LoginLog struct {
	// LoginLogId defined
	LoginLogId null.Int `xorm:"int(11) pk notnull autoincr 'login_log_id'" json:"login_log_id" form:"login_log_id" xml:"login_log_id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// StartTime defined
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	// EndTime defined
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" form:"end_time" xml:"end_time"`
	// LoginHour defined
	LoginHour null.Float `xorm:"float(11,2) 'login_hour'" json:"login_hour" form:"login_hour" xml:"login_hour"`
	// LoginIp defined
	LoginIp null.String `xorm:"varchar(100) 'login_ip'" json:"login_ip" form:"login_ip" xml:"login_ip"`
}

// TableName table name of defined T2LoginLog
func (m *T2LoginLog) TableName() string {
	return "t_2_login_log"
}

func (r *T2LoginLog) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalT2LoginLog(data []byte) (T2LoginLog, error) {
	var r T2LoginLog
	err := json.Unmarshal(data, &r)
	return r, err
}
