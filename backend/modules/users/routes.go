package users

import (
	"github.com/labstack/echo/v4"
	"github.com/sazzo/sazztv/backend/middleware"
)

func CreateRoutes(e *echo.Echo) {
	router := e.Group("/users")

	router.GET("/:username", getUser)
	router.GET("/@me", getCurrentUser, middleware.GetJWTMiddleware())
}