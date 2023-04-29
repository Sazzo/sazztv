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

func getUserStreamCredentials(c echo.Context) error {
	db := database.GetConnection()
	authUser := c.Get("user").(model.User)
	
	streamCredentials := new(struct {
		RTMPUrl string `json:"rtmp_url"`
		StreamKey string `json:"stream_key"`
	})

	err := db.Model(&model.StreamCredentials{}).Where("user_id = ?", authUser.ID).First(&streamCredentials).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, &util.APIError{
			Message: "Stream credentials for the user not found",
		})
	}

	return c.JSON(http.StatusOK, streamCredentials)
}