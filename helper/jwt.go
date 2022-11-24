package helper

import (
	"context"
	"fmt"
	"social/domain"
	"social/opentelemetry"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("dfugrshnguregniuiurhgi4oeufidyunhuwbxru3n687aw8oyeiusngwiy")

func GenerateToken(ctx context.Context, userData domain.User) (string, error) {
	tracer := opentelemetry.GetTracer()

	_, span := tracer.Start(ctx, "JWT.GenerateToken")
	defer span.End()

	claims := domain.JWTClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		User: userData,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, err
}

func ParseToken(tokenString string) (*domain.JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %s", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})

	if claims, ok := token.Claims.(*domain.JWTClaim); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
