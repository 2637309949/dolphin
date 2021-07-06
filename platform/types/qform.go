package types

import "encoding/json"

// QueryForm defined
type QueryForm map[string]interface{}

// Marshal defined
func (f *QueryForm) Marshal() ([]byte, error) {
	return json.Marshal(f)
}

// Unmarshal defined
func (f *QueryForm) Unmarshal(v interface{}) error {
	byt, err := f.Marshal()
	if err != nil {
		return err
	}
	return json.Unmarshal(byt, v)
}
