package errors

// known errors
var (
	ErrNotFound = New(Code(100002), "record has not found", "not found")
)
