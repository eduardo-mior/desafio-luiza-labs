package services

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ServiceAuth struct{}

func (*ServiceAuth) GenerateTokenJWT() (string, error) {
	now := time.Now().UTC()

	// Gerando um ID unico para o token
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	claims := jwt.MapClaims{
		"jti": uuid,
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(time.Hour * 1).Unix(), // 1 dia
		"sub": 1,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
