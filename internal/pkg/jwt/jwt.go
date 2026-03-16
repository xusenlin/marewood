package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 72 // Token过期时间
	JwtClaimsKey        = "JwtClaims"    // Claims 储存在 *gin.Context 里的关键字
)

var TokenSecret = []byte("MareWood")

// Claims JWT Claims 结构
type Claims struct {
	ID       uint
	Username string
	Role     int
	Status   int
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT Token
func GenerateToken(claims *Claims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		Issuer:    "senLin",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(TokenSecret)
}

// ParseToken 解析 JWT Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return TokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid Token, please login again")
}
