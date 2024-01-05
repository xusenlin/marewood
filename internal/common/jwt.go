package common

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 72 //Token过期时间
	JwtClaimsKey        = "JwtClaims"    //Claims 储存在 *gin.Context 里的关键字，通过JWTAuth中间件设置。
)

var TokenSecret = []byte("MareWood")

type Claims struct {
	ID       uint
	Username string
	Role     int
	Status   int
	jwt.RegisteredClaims
}

func JwtGenToken(claims *Claims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		Issuer:    "senLin",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(TokenSecret)
}

func JwtParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return TokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid Token, please login again")
}

//
//func JwtGetClaimsByContext(c *gin.Context) (*Claims, error) {
//
//	jwtClaims, hasClaims := c.Get(JwtClaimsKey)
//	if !hasClaims {
//		return &Claims{}, errors.New("user information has been lost")
//	}
//
//	claims, ok := jwtClaims.(*Claims)
//	if !ok {
//		return &Claims{}, errors.New("user information assertion failed")
//	}
//	return claims, nil
//}
//
//func JwtGetUserByContext(c *gin.Context) (*User, error) {
//	claims, err := JwtGetClaimsByContext(c)
//	if err != nil {
//		return nil, err
//	}
//	var user User
//	err = db.Conn.Where("id = ?", claims.ID).First(&user).Error
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
