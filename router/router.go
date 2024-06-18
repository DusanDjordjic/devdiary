package router

import (
	"dev-diary/router/handlers"
	"dev-diary/router/handlers/api"

	"github.com/labstack/echo"
)

func SetupRouter(server *echo.Echo) {
	server.GET("/", handlers.HomePageHandler)
	server.GET("/about", handlers.AboutPageHandler)
	server.GET("/contact", handlers.ContactPageHandler)
	server.GET("/post/:id", handlers.PostPageHandler)
	server.POST("/send_message", api.SendMessageHandler)

	// API Handlers
	server.GET("/api/posts", api.GetAllPostsHandler)
	server.GET("/api/posts/:id", api.GetPostByIDHandler)
	server.POST("/api/posts", api.CreatePostHandler)
	server.PUT("/api/posts/:id", api.UpdatePostByIDHandler)
	server.DELETE("/api/posts/:id", api.DeletePostByIDHandler)

	server.GET("/api/tags", api.GetAllTagsHandler)
	server.GET("/api/tags/:id", api.GetTagByIDHandler)
	server.POST("/api/tags", api.CreateTagHandler)
	server.PUT("/api/tags/:id", api.UpdateTagByIDHandler)
	server.DELETE("/api/tags/:id", api.DeleteTagByIDHandler)

}
