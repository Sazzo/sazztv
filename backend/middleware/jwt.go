package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sazzo/sazztv/backend/database"
	"github.com/sazzo/sazztv/backend/model"
	"github.com/sazzo/sazztv/backend/modules/auth"
	"gorm.io/gorm"
)

func JWT() echo.MiddlewareFunc {
	return func (next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizationHeader := c.Request().Header.Get("Authorization")
			if authorizationHeader == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized.")
			}

			// remove the "Bearer " prefix
			jwtToken := strings.Split(authorizationHeader, " ")[1]
			if jwtToken == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized.")
			}

			jwtClaims, err := auth.DecodeUserTokenJwt(jwtToken)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized.")
			}

			c.Set("userTokenClaims", jwtClaims)
		
			db := database.GetConnection()

			var user model.User
			findUserErr := db.Model(&model.User{}).Where("id = ?", jwtClaims.Id).First(&user).Error
			if errors.Is(findUserErr, gorm.ErrRecordNotFound) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized.")
			}

			c.Set("user", user)
			return next(c)
		}
	}
}