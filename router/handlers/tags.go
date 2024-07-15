package handlers

import (
	"dev-diary/logger"
	"dev-diary/services"
	"dev-diary/utils"
	"html/template"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func TagsPageHandler(e echo.Context) error {
	log := logger.Log.Named("[TagsPageHandler]")
	log.Info("started")

	template := template.Must(template.ParseFiles(
		utils.GetTemplateFilePath("tags.html"),
		utils.GetTemplateFilePath("base.html"),
	))

	tags, count, err := services.GetAllTags()
	if err != nil {
		log.Error("failed to get tags", zap.Error(err))
		return err
	}

	err = template.ExecuteTemplate(e.Response(), "base", map[string]any{"Tags": tags, "Count": count})
	if err != nil {
		log.Error("failed to execute template", zap.Error(err))
		return err
	}

	log.Info("finished")

	return nil
}
