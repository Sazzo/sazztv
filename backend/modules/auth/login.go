package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/sazzo/sazztv/backend/database"
	"github.com/sazzo/sazztv/backend/model"
	"github.com/sazzo/sazztv/backend/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

func login(c echo.Context) (err error) {
	userCredentials := new(LoginDTO)
	if err = c.Bind(userCredentials); err != nil {
		return c.JSON(http.StatusBadRequest, &util.APIError{
			Message: "Invalid request body",
		})
	}
	if err = c.Validate(userCredentials); err != nil {
      return err
    }

	var user model.User

	db := database.GetConnection()

	findUserErr := db.Model(&model.User{}).Where("username = ?", userCredentials.Username).First(&user).Error
	if errors.Is(findUserErr, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusBadRequest, &util.APIError{
			Message: "Invalid user credentials",
		})
	}

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredentials.Password))
	if comparePasswordErr != nil {
		return c.JSON(http.StatusBadRequest, &util.APIError{
			Message: "Invalid user credentials",
		})
	}

	jwt, jwtEncodeErr := EncodeUserTokenJwt(user.ID, user.Username, user.IsAdmin)
	if jwtEncodeErr != nil {
		log.Fatal(jwtEncodeErr)
	}

	return c.JSON(http.StatusOK, LoginResponse{
		AccessToken: jwt,
	})
}