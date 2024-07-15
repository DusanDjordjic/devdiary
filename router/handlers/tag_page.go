package handlers

import (
	"dev-diary/logger"
	"dev-diary/services"
	"dev-diary/utils"
	"html/template"
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func TagPageHandler(e echo.Context) error {
	log := logger.Log.Named("[TagPageHandler]")
	log.Info("started")

	template := template.Must(template.ParseFiles(
		utils.GetTemplateFilePath("tag.html"),
		utils.GetTemplateFilePath("base.html"),
	))

	tagName := e.Param("name")
	if tagName == "" {
		log.Error("tag is empty")
		return e.Redirect(http.StatusTemporaryRedirect, "/tags")
	}

	tag, err := services.GetTagByNameWithPosts(tagName)
	if err != nil {
		log.Error("failed to get tag", zap.Error(err))
		return e.Redirect(http.StatusTemporaryRedirect, "/tags")
	}

	err = template.ExecuteTemplate(e.Response(), "base", tag)
	if err != nil {
		log.Error("failed to execute template", zap.Error(err))
		return err
	}

	log.Info("finished")

	return nil
}
