package handlers

import (
	"dev-diary/logger"
	"dev-diary/utils"
	"html/template"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func ContactPageHandler(e echo.Context) error {
	log := logger.Log.Named("[ContactPageHandler]")
	log.Info("started")

	template := template.Must(template.ParseFiles(
		utils.GetTemplateFilePath("contact.html"),
		utils.GetTemplateFilePath("base.html"),
	))

	err := template.ExecuteTemplate(e.Response(), "base", nil)
	if err != nil {
		log.Error("failed to execute template", zap.Error(err))
		return err
	}

	log.Info("finished")

	return nil
}
