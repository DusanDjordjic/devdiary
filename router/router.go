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
	server.GET("/tags", handlers.TagsPageHandler)
	server.GET("/tags/:name", handlers.TagPageHandler)
	server.GET("/posts/:id", handlers.PostPageHandler)
	server.POST("/contact", api.SendMessageHandler)

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

	// Catch all page to redirect to "/"
	server.GET("/*", handlers.CatchAllPage)

}
