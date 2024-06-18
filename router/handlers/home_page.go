package handlers

import (
	"dev-diary/logger"
	"dev-diary/services"
	"dev-diary/utils"
	"html/template"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func HomePageHandler(e echo.Context) error {
	log := logger.Log.Named("[HomePageHandler]")
	log.Info("started")

	posts, _, err := services.GetAllPosts()
	if err != nil {
		log.Error("failed to get all posts", zap.Error(err))
		return err
	}

	template := template.Must(template.ParseFiles(
		utils.GetTemplateFilePath("home.html"),
		utils.GetTemplateFilePath("base.html"),
	))

	err = template.ExecuteTemplate(e.Response(), "base", map[string]any{
		"Posts": posts,
	})

	if err != nil {
		log.Error("failed to execute template", zap.Error(err))
		return err
	}

	log.Info("finished")

	return nil
}
