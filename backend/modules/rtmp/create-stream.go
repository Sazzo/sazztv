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

func createStream(c echo.Context) error {
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

	user.IsLive = true
	user.LastStreamAt = time.Now()
	db.Save(&user)
	
	// Send a 303 redirect to the user username, so RTMP can transform the stream key into a username (so we can use it in the stream URL)
	return c.Redirect(http.StatusSeeOther, user.Username)
}