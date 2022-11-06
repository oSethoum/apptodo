package routes

import (
	"app/backend/handlers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	e.GET("/playground", handlers.PlaygroundHandler)
	e.Any("/query", handlers.GraphqlHandler)
	e.Any("/subscription", handlers.GraphqlWsHandler)
}
