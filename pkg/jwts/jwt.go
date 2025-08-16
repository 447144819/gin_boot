package jwts

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 请使用复杂字符串，并从配置中读取！
var (
	secretKey = []byte("tC3!wR0]oW5{tG7(eP8}eK8,fG6?rB8(")
	expiresAt = 24 // Token有效期：小时
)

// 自定义 Claims，可以加入更多字段，
type CustomClaims struct {
	jwt.RegisteredClaims

	//	声明你自己要放进 token 的数据
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
}

type JWTHandler struct{}

func NewJWTHandler() *JWTHandler {
	return &JWTHandler{}
}

func (j *JWTHandler) SetJWTToken(userId int64, username string) (string, error) {
	claims := CustomClaims{
		UserID:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expiresAt) * time.Hour)), // Token有效期：24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ParseToken 解析并验证 JWT Token
func (j *JWTHandler) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
