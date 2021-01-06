package serviceUser

import (
	"MareWood/config"
	"MareWood/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

// GenToken 生成JWT
func GenToken(claims *models.Claims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(models.TokenExpireDuration).Unix(), // 过期时间
		Issuer:    config.Cfg.AppName,                                // 签发人
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(models.TokenSecret)
}

func ParseToken(tokenString string) (*models.Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return models.TokenSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("Token invalid ")
}

func GetJwtClaimsByContext(c *gin.Context) (*models.Claims, error) {

	jwtClaims, hasClaims := c.Get(models.JwtClaimsKey)
	if !hasClaims {
		return &models.Claims{}, errors.New(" User information is lost")
	}

	claims, ok := jwtClaims.(*models.Claims)
	if !ok {
		return &models.Claims{}, errors.New(" User information assertion failed")
	}
	return claims, nil
}
