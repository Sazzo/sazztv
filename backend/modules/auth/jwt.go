package auth

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	Id	   	 string `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func EncodeUserTokenJwt(id string, username string, isAdmin bool) (string, error) {
	claims := JWTClaims{
		Id:       id,
		Username: username,
		IsAdmin:  isAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	return token.SignedString(jwtSecret)
}
