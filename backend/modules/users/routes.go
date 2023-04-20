package users

import "github.com/labstack/echo/v4"

func CreateRoutes(e *echo.Echo) {
	router := e.Group("/users")

	router.GET("/:username", getUser)
}