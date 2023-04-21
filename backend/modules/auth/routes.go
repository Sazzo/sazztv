package auth

import "github.com/labstack/echo/v4"

func CreateRoutes(e *echo.Echo) {
	router := e.Group("/auth")

	router.POST("/login", login)
	router.POST("/register", register)
}