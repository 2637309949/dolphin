// Code generated by dol build. DO NOT EDIT.
// source: auto.go

package types

import (
	"encoding/json"

	"github.com/2637309949/dolphin/packages/null"
)

// TransReq defined trans info
type TransReq struct {
	// amount
	Amount null.Int `json:"amount" xml:"amount"`
	// TransInResult
	TransInResult null.String `json:"trans_in_result" xml:"trans_in_result"`
	// TransOutResult
	TransOutResult null.String `json:"trans_out_result" xml:"trans_out_result"`
}

func (r *TransReq) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func UnmarshalTransReq(data []byte) (TransReq, error) {
	var r TransReq
	err := json.Unmarshal(data, &r)
	return r, err
}