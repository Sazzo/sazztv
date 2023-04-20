package auth

import (
	"net/http"

	"github.com/sazzo/sazztv/backend/util"

	"github.com/labstack/echo/v4"
)

type LoginDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
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


	return c.JSON(http.StatusOK, userCredentials)
}