package users

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sazzo/sazztv/backend/database"
	"github.com/sazzo/sazztv/backend/model"
	"github.com/sazzo/sazztv/backend/modules/auth"
	"github.com/sazzo/sazztv/backend/util"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getCurrentUser(c echo.Context) error {
	db := database.GetConnection()
	userToken := c.Get("user").(*jwt.Token)
	userTokenClaims := userToken.Claims.(*auth.JWTClaims)

	user := new(struct {
		ID string `json:"id"`
		Username string `json:"username"`
		StreamSettings model.StreamSettings `gorm:"foreignKey:UserID" json:"stream_settings"`
		IsAdmin bool `json:"is_admin"`
		IsLive bool `json:"is_live"`
		LastStreamAt string `json:"last_stream_at"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	})

	err := db.Model(&model.User{}).Where("id = ?", userTokenClaims.Id).Preload("StreamSettings").First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &util.APIError{
			Message: "User not found",
		})
	}

	return c.JSON(http.StatusOK, user)
}