// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package model

import (
	"reflect"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/2637309949/dolphin/packages/xormplus/xorm"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/caches"
	"github.com/2637309949/dolphin/packages/xormplus/xorm/tags"
)

// MarketActivity defined
type MarketActivity struct {
	// MaName defined
	MaName null.String `xorm:"varchar(50) 'ma_name'" json:"ma_name" form:"ma_name" xml:"ma_name"`
	// MaType defined
	MaType null.Int `xorm:"int(11) 'ma_type'" json:"ma_type" form:"ma_type" xml:"ma_type"`
	// StartTime defined
	StartTime null.Time `xorm:"datetime 'start_time'" json:"start_time" form:"start_time" xml:"start_time"`
	// EndTime defined
	EndTime null.Time `xorm:"datetime 'end_time'" json:"end_time" form:"end_time" xml:"end_time"`
	// MaUseMoney defined
	MaUseMoney null.Float `xorm:"float(10,2) 'ma_use_money'" json:"ma_use_money" form:"ma_use_money" xml:"ma_use_money"`
	// Creater defined
	Creater null.String `xorm:"varchar(36) 'creater'" json:"creater" form:"creater" xml:"creater"`
	// Updater defined
	Updater null.String `xorm:"varchar(36) 'updater'" json:"updater" form:"updater" xml:"updater"`
	// MmId defined
	MmId null.Int `xorm:"int(11) 'mm_id'" json:"mm_id" form:"mm_id" xml:"mm_id"`
	// MAId defined
	MAId null.Int `xorm:"int(11) pk notnull autoincr 'm_a_id'" json:"m_a_id" form:"m_a_id" xml:"m_a_id"`
	// CreateDate defined
	CreateDate null.Time `xorm:"datetime 'create_date'" json:"create_date" form:"create_date" xml:"create_date"`
	// UpdateDate defined
	UpdateDate null.Time `xorm:"datetime 'update_date'" json:"update_date" form:"update_date" xml:"update_date"`
	// Isdelete defined
	Isdelete null.Int `xorm:"int(11) 'isdelete'" json:"isdelete" form:"isdelete" xml:"isdelete"`
	// ActivityType defined
	ActivityType null.Int `xorm:"int(11) 'activity_type'" json:"activity_type" form:"activity_type" xml:"activity_type"`
	// ActivityExposure defined
	ActivityExposure null.Float `xorm:"float(10,2) 'activity_exposure'" json:"activity_exposure" form:"activity_exposure" xml:"activity_exposure"`
	// ActivityPlace defined
	ActivityPlace null.String `xorm:"varchar(200) 'activity_place'" json:"activity_place" form:"activity_place" xml:"activity_place"`
	// ActivityContent defined
	ActivityContent null.String `xorm:"varchar(500) 'activity_content'" json:"activity_content" form:"activity_content" xml:"activity_content"`
	// ActivityDesc defined
	ActivityDesc null.String `xorm:"varchar(500) 'activity_desc'" json:"activity_desc" form:"activity_desc" xml:"activity_desc"`
	// ActivityEstimateIncome defined
	ActivityEstimateIncome null.Float `xorm:"float(10,2) 'activity_estimate_income'" json:"activity_estimate_income" form:"activity_estimate_income" xml:"activity_estimate_income"`
	// MupId defined
	MupId null.Int `xorm:"int(11) 'mup_id'" json:"mup_id" form:"mup_id" xml:"mup_id"`
	// BmupId defined
	BmupId null.Int `xorm:"int(11) 'bmup_id'" json:"bmup_id" form:"bmup_id" xml:"bmup_id"`
	// AmupId defined
	AmupId null.Int `xorm:"int(11) 'amup_id'" json:"amup_id" form:"amup_id" xml:"amup_id"`
	// PmupId defined
	PmupId null.Int `xorm:"int(11) 'pmup_id'" json:"pmup_id" form:"pmup_id" xml:"pmup_id"`
	// OmupId defined
	OmupId null.Int `xorm:"int(11) 'omup_id'" json:"omup_id" form:"omup_id" xml:"omup_id"`
	// NupId defined
	NupId null.Int `xorm:"int(11) 'nup_id'" json:"nup_id" form:"nup_id" xml:"nup_id"`
	// MdpId defined
	MdpId null.Int `xorm:"int(11) 'mdp_id'" json:"mdp_id" form:"mdp_id" xml:"mdp_id"`
	// ErdpId defined
	ErdpId null.Int `xorm:"int(11) 'erdp_id'" json:"erdp_id" form:"erdp_id" xml:"erdp_id"`
	// AmdId defined
	AmdId null.Int `xorm:"int(11) 'amd_id'" json:"amd_id" form:"amd_id" xml:"amd_id"`
	// PmdId defined
	PmdId null.Int `xorm:"int(11) 'pmd_id'" json:"pmd_id" form:"pmd_id" xml:"pmd_id"`
	// OmdId defined
	OmdId null.Int `xorm:"int(11) 'omd_id'" json:"omd_id" form:"omd_id" xml:"omd_id"`
}

// Parser defined
func (m *MarketActivity) Parser(db *xorm.Engine) *tags.Parser {
	return tags.NewParser("xorm", db.Dialect(), db.DB().Mapper, db.DB().Mapper, caches.NewManager())
}

// PrimaryKeys defined
func (m *MarketActivity) PrimaryKeys(db *xorm.Engine) ([]string, error) {
	v := reflect.Indirect(reflect.ValueOf(m))
	table, err := m.Parser(db).Parse(v)
	return table.PrimaryKeys, err
}

// TableName table name of defined MarketActivity
func (m *MarketActivity) TableName() string {
	return "market_activity"
}
