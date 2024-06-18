package api

import (
	"dev-diary/logger"
	"dev-diary/utils"
	"fmt"
	"html/template"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func SendMessageHandler(c echo.Context) error {
	log := logger.Log.Named("[SendMessageHandler]")
	log.Info("started")

	name := c.FormValue("name")
	email := c.FormValue("email")
	message := c.FormValue("message")

	fmt.Println(name, email, message)

	template := template.Must(template.ParseFiles(
		utils.GetTemplateFilePath("received_message.html"),
		utils.GetTemplateFilePath("base.html"),
	))

	err := template.ExecuteTemplate(c.Response(), "base", map[string]any{
		"Message": "Successfully received message",
	})

	if err != nil {
		log.Error("failed to execute template", zap.Error(err))
		return err
	}

	log.Info("finished")

	return nil
}
