package handlers

import (
	"dev-diary/services"
	"dev-diary/utils"
	"html/template"

	"github.com/labstack/echo"
)

func BlogPageHandler(e echo.Context) error {
	blogs, _, err := services.GetAllBlogs()
	if err != nil {
		return err
	}

	template := template.Must(template.ParseFiles(utils.GetTemplateFilePath("blog.html"), utils.GetTemplateFilePath("base.html")))

	return template.ExecuteTemplate(e.Response(), "base", map[string]any{
		"Posts": blogs,
	})
}
