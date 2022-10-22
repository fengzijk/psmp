package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

// CustomClaims Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UUID        string
	UserId      uint
	Username    string
	NickName    string
	AuthorityId uint
}
