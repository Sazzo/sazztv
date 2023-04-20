package users

import (
	"errors"
	"net/http"

	"github.com/sazzo/sazztv/backend/database"
	"github.com/sazzo/sazztv/backend/model"
	"github.com/sazzo/sazztv/backend/util"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func getUser(c echo.Context) error {
	db := database.GetConnection()
	username := c.Param("username")

	var user model.User

	err := db.Model(&model.User{}).Where("username = ?", username).Select("id", "username", "is_live", "last_stream_at", "created_at", "updated_at").First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &util.APIError{
			Message: "User not found",
		})
	}

	return c.JSON(http.StatusOK, user)
}