package main

import (
	"dev-diary/router"
	"dev-diary/utils"

	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	server.Static("/static", utils.GetStaticFolderPath())
	router.SetupRouter(server)

	if err := server.Start("127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
