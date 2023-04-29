package users

import (
	"github.com/labstack/echo/v4"
	"github.com/sazzo/sazztv/backend/middleware"
)

func CreateRoutes(e *echo.Echo) {
	router := e.Group("/users")

	router.GET("/:username", getUser)

	// Auth-required routes

	router.GET("/@me", getCurrentUser, middleware.JWT())
	router.GET("/@me/stream-credentials", getUserStreamCredentials, middleware.JWT())

	router.PATCH("/@me/username", changeUsername, middleware.JWT())
}