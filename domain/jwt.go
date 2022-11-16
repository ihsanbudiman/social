package domain

import "github.com/golang-jwt/jwt/v4"

type JWTClaim struct {
	User
	jwt.RegisteredClaims
}
