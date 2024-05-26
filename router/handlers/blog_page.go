package handlers

import (
	"dev-diary/utils"
	"html/template"

	"github.com/labstack/echo"
)

func BlogPageHandler(e echo.Context) error {
	template := template.Must(template.ParseFiles(utils.GetTemplateFilePath("blog.html"), utils.GetTemplateFilePath("base.html")))

	return template.ExecuteTemplate(e.Response(), "base", nil)
}
