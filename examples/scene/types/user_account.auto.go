// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
	"github.com/shopspring/decimal"
)

// UserAccount defined
type UserAccount struct {
	// ID defined
	ID null.Int `xorm:"int(11) pk notnull autoincr 'id'" json:"id" form:"id" xml:"id"`
	// UserId defined
	UserId null.Int `xorm:"int(11) notnull 'user_id'" json:"user_id" form:"user_id" xml:"user_id"`
	// Balance defined
	Balance decimal.Decimal `xorm:"decimal(10,2) notnull default(0.00) 'balance'" json:"balance" form:"balance" xml:"balance"`
	// CreateTime defined
	CreateTime null.Time `xorm:"datetime default(CURRENT_TIMESTAMP) 'create_time'" json:"create_time" form:"create_time" xml:"create_time"`
	// UpdateTime defined
	UpdateTime null.Time `xorm:"datetime default(CURRENT_TIMESTAMP) 'update_time'" json:"update_time" form:"update_time" xml:"update_time"`
}

// TableName table name of defined UserAccount
func (m *UserAccount) TableName() string {
	return "user_account"
}

func (r *UserAccount) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalUserAccount(data []byte) (UserAccount, error) {
	var r UserAccount
	err := json.Unmarshal(data, &r)
	return r, err
}
