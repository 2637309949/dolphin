// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import "github.com/2637309949/dolphin/packages/null"

// TrackerInfo defined 日志信息
type TrackerInfo struct {
	*SysTracker
	// 域名
	Domain null.String `json:"domain" xml:"domain"`
}
