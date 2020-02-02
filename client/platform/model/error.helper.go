package model

// Error defined error helpers
func (e Error) Error() string {
	return e.Msg
}
