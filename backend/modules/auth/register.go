package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/sazzo/sazztv/backend/database"
	"github.com/sazzo/sazztv/backend/model"
	"github.com/sazzo/sazztv/backend/util"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

type RegisterDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	AccessToken string `json:"accessToken"`
}

func register(c echo.Context) (err error) {
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
		hashedPassword, hashPasswordErr := bcrypt.GenerateFromPassword([]byte(userCredentials.Password), bcrypt.DefaultCost)
		if hashPasswordErr != nil {
			log.Fatal(hashPasswordErr) 
		}
		
		user = model.User{
			Username: userCredentials.Username,
			Password: string(hashedPassword),
			StreamKey: randstr.Hex(16),
		}

		createUserErr := db.Create(&user).Error
		if createUserErr != nil {
			log.Fatal(createUserErr)
		}
	} else {
		return c.JSON(http.StatusBadRequest, &util.APIError{
			Message: "User already exists",
		})
	}

	jwt, jwtEncodeErr := EncodeUserTokenJwt(user.ID, user.Username, user.IsAdmin)
	if jwtEncodeErr != nil {
		log.Fatal(jwtEncodeErr)
	}

	return c.JSON(http.StatusOK, RegisterResponse{
		AccessToken: jwt,
	})
}