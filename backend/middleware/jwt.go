package middleware

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sazzo/sazztv/backend/modules/auth"
)

func GetJWTMiddleware() echo.MiddlewareFunc {
	jwtMiddlewareConfig := echojwt.Config{
		NewClaimsFunc:  func(c echo.Context) jwt.Claims {
			return new(auth.JWTClaims)
		},
		SigningKey:     []byte(os.Getenv("JWT_SECRET")),
	}

	return echojwt.WithConfig(jwtMiddlewareConfig)
}