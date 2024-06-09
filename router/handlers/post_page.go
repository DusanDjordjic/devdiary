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

	post, err := services.GetPostByID(uint(id))
	if err != nil {
		return err
	}

	tmp := template.Must(template.ParseFiles(
		utils.GetTemplateFilePath("post.html"),
		utils.GetTemplateFilePath("base.html"),
	))

	return tmp.ExecuteTemplate(e.Response(), "base", map[string]any{
		"ID":          post.ID,
		"CreatedAt":   post.CreatedAt,
		"Title":       post.Title,
		"Description": post.Description,
		"ImageURL":    post.ImageURL,
		"Content":     template.HTML(post.Content),
	})
}
