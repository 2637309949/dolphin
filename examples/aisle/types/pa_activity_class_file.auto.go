// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// PaActivityClassFile defined
type PaActivityClassFile struct {
	// PACFId defined
	PACFId null.Int `xorm:"int(11) pk notnull autoincr 'p_a_c_f_id'" json:"p_a_c_f_id" form:"p_a_c_f_id" xml:"p_a_c_f_id"`
	// FileId defined
	FileId null.Int `xorm:"int(11) 'file_id'" json:"file_id" form:"file_id" xml:"file_id"`
	// PacfDesc defined
	PacfDesc null.String `xorm:"varchar(500) 'pacf_desc'" json:"pacf_desc" form:"pacf_desc" xml:"pacf_desc"`
	// ParId defined
	ParId null.Int `xorm:"int(11) 'par_id'" json:"par_id" form:"par_id" xml:"par_id"`
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

// TableName table name of defined PaActivityClassFile
func (m *PaActivityClassFile) TableName() string {
	return "pa_activity_class_file"
}

func (r *PaActivityClassFile) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalPaActivityClassFile(data []byte) (PaActivityClassFile, error) {
	var r PaActivityClassFile
	err := json.Unmarshal(data, &r)
	return r, err
}
