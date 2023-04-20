package main

import (
	"log"
	"net/http"

	"github.com/sazzo/sazztv/backend/database"
	"github.com/sazzo/sazztv/backend/modules/auth"
	"github.com/sazzo/sazztv/backend/modules/rtmp"
	"github.com/sazzo/sazztv/backend/modules/users"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	  validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
    // Optionally, you could return the error to give each route more control over the status code
	validationErrors := err.(validator.ValidationErrors)
	println(validationErrors.Error())
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.CreateConnection()

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	rtmp.CreateRoutes(e)
	users.CreateRoutes(e)
	auth.CreateRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
