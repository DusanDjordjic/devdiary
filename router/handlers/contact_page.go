package handlers

import (
	"dev-diary/utils"
	"html/template"

	"github.com/labstack/echo"
)

func ContactPageHandler(e echo.Context) error {
	template := template.Must(template.ParseFiles(
		utils.GetTemplateFilePath("contact.html"),
		utils.GetTemplateFilePath("base.html"),
	))

	return template.ExecuteTemplate(e.Response(), "base", nil)
}
