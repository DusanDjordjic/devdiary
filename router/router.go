package router

import (
	"dev-diary/router/handlers"

	"github.com/labstack/echo"
)

func SetupRouter(server *echo.Echo) {
	server.GET("/", handlers.HomePageHandler)
	server.GET("/about", handlers.AboutPageHandler)
	server.GET("/contact", handlers.ContactPageHandler)
	server.GET("/blog", handlers.BlogPageHandler)
}
