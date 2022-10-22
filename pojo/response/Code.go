package response

const (
	AuthorizationError   = 10104
	UrlSignError         = 10105
	CacheSetError        = 10106
	CacheGetError        = 10107
	CacheDelError        = 10108
	CacheNotExist        = 10109
	ResubmitError        = 10110
	HashIdsEncodeError   = 10111
	HashIdsDecodeError   = 10112
	AuthorizationExpired = 405
	IllegalAccess        = 403
	TokenExpired         = 406
	TokenNotValidYet     = 407
	TokenMalformed       = 408
	TokenInvalid         = 409
	SuccessCode          = 200
)

func Text(code int) string {

	return zhCNText[code]
}
