package context

import (
	"errors"
	"marewood/internal/pkg/jwt"
)

func (r *Context) GetClaims() (*jwt.Claims, error) {
	jwtClaims, hasClaims := r.c.Get(jwt.JwtClaimsKey)
	if !hasClaims {
		return &jwt.Claims{}, errors.New("user information has been lost")
	}

	claims, ok := jwtClaims.(*jwt.Claims)
	if !ok {
		return &jwt.Claims{}, errors.New("user information assertion failed")
	}
	return claims, nil
}
