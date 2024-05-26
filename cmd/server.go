package main

import (
	"dev-diary/db"
	"dev-diary/router"
	"dev-diary/utils"
	"log"

	"github.com/labstack/echo"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	server := echo.New()
	server.Static("/static", utils.GetStaticFolderPath())
	router.SetupRouter(server)

	if err := server.Start("127.0.0.1:8080"); err != nil {
		panic(err)
	}

	if err := db.Close(); err != nil {
		panic(err)
	}

}
