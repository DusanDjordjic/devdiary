package handlers

import (
	"dev-diary/services"
	"dev-diary/utils"
	"html/template"
	"strconv"

	"github.com/labstack/echo"
)

func PostPageHandler(e echo.Context) error {
	idString := e.Param("id")

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return err
	}

	blog, err := services.GetBlogByID(uint(id))
	if err != nil {
		return err
	}

	template := template.Must(template.ParseFiles(utils.GetTemplateFilePath("post.html"), utils.GetTemplateFilePath("base.html")))

	return template.ExecuteTemplate(e.Response(), "base", blog)
}
