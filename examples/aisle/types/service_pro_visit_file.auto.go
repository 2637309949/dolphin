// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// ServiceProVisitFile defined
type ServiceProVisitFile struct {
	// SPVFId defined
	SPVFId null.Int `xorm:"int(11) pk notnull autoincr 's_p_v_f_id'" json:"s_p_v_f_id" form:"s_p_v_f_id" xml:"s_p_v_f_id"`
	// SpvId defined
	SpvId null.Int `xorm:"int(11) 'spv_id'" json:"spv_id" form:"spv_id" xml:"spv_id"`
	// Files defined
	Files null.Int `xorm:"int(11) 'files'" json:"files" form:"files" xml:"files"`
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

// TableName table name of defined ServiceProVisitFile
func (m *ServiceProVisitFile) TableName() string {
	return "service_pro_visit_file"
}

func (r *ServiceProVisitFile) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalServiceProVisitFile(data []byte) (ServiceProVisitFile, error) {
	var r ServiceProVisitFile
	err := json.Unmarshal(data, &r)
	return r, err
}
