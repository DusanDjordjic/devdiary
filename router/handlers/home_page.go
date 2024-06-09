package handlers

import (
	"dev-diary/services"
	"dev-diary/utils"
	"html/template"

	"github.com/labstack/echo"
)

func HomePageHandler(e echo.Context) error {
	posts, _, err := services.GetAllPosts()
	if err != nil {
		return err
	}

	template := template.Must(template.ParseFiles(
		utils.GetTemplateFilePath("home.html"),
		utils.GetTemplateFilePath("base.html"),
	))

	return template.ExecuteTemplate(e.Response(), "base", map[string]any{
		"Posts": posts,
	})
}
