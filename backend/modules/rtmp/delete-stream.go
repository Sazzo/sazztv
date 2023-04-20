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

	var user model.User

	err := db.Model(&model.User{}).Where("stream_key = ?", streamKey).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusBadRequest, &util.APIError{
			Message: "Invalid stream key",
		})
	}

	user.IsLive = false
	user.LastStreamAt = time.Now()
	db.Save(&user)
	
	// Send a 303 redirect to the user username, so RTMP can transform the stream key into a username (so we can use it in the stream URL)
	return c.NoContent(http.StatusOK)
}