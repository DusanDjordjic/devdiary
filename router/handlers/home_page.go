package handlers

import (
	"dev-diary/utils"
	"html/template"

	"github.com/labstack/echo"
)

func HomePageHandler(e echo.Context) error {
	template := template.Must(template.ParseFiles(utils.GetTemplateFilePath("home.html"), utils.GetTemplateFilePath("base.html")))

	return template.ExecuteTemplate(e.Response(), "base", nil)
}
