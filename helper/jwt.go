package helper

import (
	"FEDeployService/config"
	"FEDeployService/models"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenToken 生成JWT
func GenToken(userID uint, username string, role int, status int) (string, error) {
	c := models.Claims{
		ID:       userID,
		Username: username, // 自定义字段
		Role:   role,
		Status:   status,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(models.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    config.Cfg.AppName,                                // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
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
	return nil, errors.New("无效的 token")
}
