// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// MenuPicture defined
type MenuPicture struct {
	// T2850 defined
	T2850 null.Int `xorm:"int(11) pk notnull autoincr 't_285_0'" json:"t_285_0" form:"t_285_0" xml:"t_285_0"`
	// MenupicName defined
	MenupicName null.String `xorm:"varchar(500) notnull 'menupic_name'" json:"menupic_name" form:"menupic_name" xml:"menupic_name"`
	// MenupicPicture defined
	MenupicPicture null.Int `xorm:"int(11) notnull 'menupic_picture'" json:"menupic_picture" form:"menupic_picture" xml:"menupic_picture"`
	// ProId defined
	ProId null.Int `xorm:"int(11) 'pro_id'" json:"pro_id" form:"pro_id" xml:"pro_id"`
	// IsDelete defined
	IsDelete null.Int `xorm:"int(11) 'is_delete'" json:"is_delete" form:"is_delete" xml:"is_delete"`
}

// TableName table name of defined MenuPicture
func (m *MenuPicture) TableName() string {
	return "menu_picture"
}

func (r *MenuPicture) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalMenuPicture(data []byte) (MenuPicture, error) {
	var r MenuPicture
	err := json.Unmarshal(data, &r)
	return r, err
}
