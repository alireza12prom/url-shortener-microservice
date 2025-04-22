package jwt

import (
	"fmt"
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

func VerifyToken(token string) error {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return err
	}

	if !t.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
