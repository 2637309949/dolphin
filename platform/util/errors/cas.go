package errors

import "errors"

// known errors
var (
	ErrDomainNotFound         = errors.New("domain not found")
	ErrRedirectUriNotFound    = errors.New("redirect_uri not found")
	ErrCodeNotFound           = errors.New("code not found")
	ErrUUIDNotFound           = errors.New("uuid not found")
	ErrInvalidUUID            = errors.New("invalid not found")
	ErrNotFoundState          = errors.New("not found state")
	ErrAccountOrPasswordError = errors.New("account doesn't exist or password error")
	ErrRecordExisted          = errors.New("the record already exists")
)
