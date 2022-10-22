package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"go-psmp/config"
	"go-psmp/pojo/request"
	"go-psmp/pojo/response"
	"go-psmp/utils/date"
	"time"
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(config.JwtConf.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	bf, _ := date.ParseDuration(config.JwtConf.BufferTime)
	ep, _ := date.ParseDuration(config.JwtConf.ExpiresTime)
	claims := request.CustomClaims{

		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.JwtConf.Issuer,                  // 签名的发行者
			NotBefore: jwt.NewNumericDate(time.Now()),         // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间 7天  配置文件
		},
	}
	return claims
}

// CreateToken 创建token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims request.CustomClaims) (string, error) {
	v, err, _ := config.ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, int) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, response.AuthorizationError
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, response.TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, response.TokenNotValidYet
			} else {
				return nil, response.TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, response.SuccessCode
		}
		return nil, response.TokenInvalid

	} else {
		return nil, response.TokenInvalid
	}
}
