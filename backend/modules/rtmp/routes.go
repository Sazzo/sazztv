package rtmp

import "github.com/labstack/echo/v4"

func CreateRoutes(e *echo.Echo) {
	router := e.Group("/rtmp")

	router.POST("/create-stream", createStream)
	router.POST("/delete-stream", deleteStream)
}