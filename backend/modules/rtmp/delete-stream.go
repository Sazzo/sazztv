package rtmp

import (
	"errors"
	"net/http"
	"time"

	"github.com/sazzo/sazztv/backend/database"
	"github.com/sazzo/sazztv/backend/model"
	"github.com/sazzo/sazztv/backend/util"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func deleteStream(c echo.Context) error {
	streamKey := c.FormValue("name") // RTMP sends the stream key as the name parameter

	db := database.GetConnection()

	var streamCredentials model.StreamCredentials

	err := db.Model(&model.StreamCredentials{}).Where("stream_key = ?", streamKey).First(&streamCredentials).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusBadRequest, &util.APIError{
			Message: "Invalid stream key",
		})
	}

	var user model.User
	db.Model(&model.User{}).Where("id = ?", streamCredentials.UserID).First(&user)

	user.IsLive = false
	user.LastStreamAt = time.Now()
	db.Save(&user)
	
	return c.NoContent(http.StatusOK)
}