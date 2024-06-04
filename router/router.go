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
	server.GET("/blog/:id", handlers.PostPageHandler)

	// TODO Create an API for creating and managing posts
	// Uploading pictures and other data
	server.GET("/api/posts", handlers.PostPageHandler)
	server.POST("/api/posts", handlers.PostPageHandler)
}
