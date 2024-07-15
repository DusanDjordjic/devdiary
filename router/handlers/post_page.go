package handlers

import (
	"dev-diary/logger"
	"dev-diary/services"
	"dev-diary/utils"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func PostPageHandler(e echo.Context) error {
	log := logger.Log.Named("[PostPageHandler]")
	log.Info("started")

	idString := e.Param("id")

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		log.Error("failed to parse id", zap.Error(err))
		return e.Redirect(http.StatusTemporaryRedirect, "/")
	}

	post, err := services.GetPostByID(uint(id))
	if err != nil {
		log.Error("failed to get post by id", zap.Uint64("id", id), zap.Error(err))
		return e.Redirect(http.StatusTemporaryRedirect, "/")
	}

	tmp := template.Must(template.ParseFiles(
		utils.GetTemplateFilePath("post.html"),
		utils.GetTemplateFilePath("base.html"),
	))

	err = tmp.ExecuteTemplate(e.Response(), "base", map[string]any{
		"ID":              post.ID,
		"CreatedAt":       post.CreatedAt,
		"Title":           post.Title,
		"Description":     post.Description,
		"ImageURL":        post.ImageURL,
		"ImageCaption":    post.ImageCaption,
		"ImageCaptionURL": post.ImageCaptionURL,
		"Content":         template.HTML(post.Content),
		"Tags":            post.Tags,
	})

	if err != nil {
		log.Error("failed to execute template", zap.Error(err))
		return err
	}

	log.Info("finished")

	return nil
}
