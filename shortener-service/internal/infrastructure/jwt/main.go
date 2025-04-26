package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JWT_SECRET   = "your-secret-key"
	JWT_ISSUER   = "url-shortener"
	JWT_DURATION = 24
)

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"iss": JWT_ISSUER,
		"exp": time.Now().Add(time.Hour * JWT_DURATION).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	tokenString, err := token.SignedString([]byte(JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Verify(token string) bool {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil || !t.Valid {
		return false
	}

	return true
}

func Extract(token string) (jwt.MapClaims, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	claims := t.Claims.(jwt.MapClaims)
	return claims, nil
}
