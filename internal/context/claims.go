package context

import (
	"errors"
	"marewood/internal/common"
)

func (r *Context) GetClaims() (*common.Claims, error) {
	jwtClaims, hasClaims := r.c.Get(common.JwtClaimsKey)
	if !hasClaims {
		return &common.Claims{}, errors.New("user information has been lost")
	}

	claims, ok := jwtClaims.(*common.Claims)
	if !ok {
		return &common.Claims{}, errors.New("user information assertion failed")
	}
	return claims, nil
}
