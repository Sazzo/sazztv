package main

import (
	"log"
	"net/http"

	"github.com/sazzo/sazztv/backend/database"
	"github.com/sazzo/sazztv/backend/modules/auth"
	"github.com/sazzo/sazztv/backend/modules/rtmp"
	"github.com/sazzo/sazztv/backend/modules/users"
	"github.com/sazzo/sazztv/backend/util"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func getErrorMessageForTag(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "Field is required"
	case "email":
		return "Field must be a valid email address"
	default:
		return fieldError.Error()
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		clientValidationErrors := make([]util.ValidationError, len(validationErrors))
		for i, validationError := range validationErrors {
			clientValidationErrors[i] = util.ValidationError{
				Field:   validationError.Field(),
				Message: getErrorMessageForTag(validationError),
			}
		}

		return echo.NewHTTPError(http.StatusBadRequest, &util.ValidationErrors{
			Message: "Validation Error",
			Errors:  clientValidationErrors,
		})
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
	e.Use(middleware.CORS())
	
	e.Validator = &CustomValidator{validator: validator.New()}

	rtmp.CreateRoutes(e)
	users.CreateRoutes(e)
	auth.CreateRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
