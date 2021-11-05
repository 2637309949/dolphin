package errors

var (
	ErrDomainNotFound         = New(Code(200001), "domain not found", "domain not found")
	ErrRedirectUriNotFound    = New(Code(200002), "redirect_uri not found", "redirect_uri not found")
	ErrCodeNotFound           = New(Code(200003), "code not found", "code not found")
	ErrUUIDNotFound           = New(Code(200004), "uuid not found", "uuid not found")
	ErrInvalidUUID            = New(Code(200005), "invalid not found", "invalid not found")
	ErrNotFoundState          = New(Code(200006), "not found state", "not found state")
	ErrAccountOrPasswordError = New(Code(200007), "account doesn't exist or password error", "account doesn't exist or password error")
	ErrRecordExisted          = New(Code(200008), "the record already exists", "the record already exists")
)
