package api

import (
	"dev-diary/logger"
	"dev-diary/models"
	"dev-diary/router/web"
	"dev-diary/services"
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func GetAllPostsHandler(c echo.Context) error {
	log := logger.Log.Named("[GetAllPostsHandler]")
	log.Info("started")

	posts, total, err := services.GetAllPosts()
	if err != nil {
		log.Error("failed to get posts", zap.Error(err))
		return web.Error(http.StatusBadRequest, err)
	}

	web.DataWithCount(c, models.DTOPosts(posts), total)
	log.Info("finished")
	return nil
}

func GetPostByIDHandler(c echo.Context) error {
	log := logger.Log.Named("[GetPostByIDHandler]")
	log.Info("started")

	postID, err := web.ParseID(c)
	if err != nil {
		log.Error("failed to parse id")
		return err
	}

	post, err := services.GetPostByID(postID)
	if err != nil {
		log.Error("failed to get post by id", zap.Error(err))
		return web.Error(http.StatusNotFound, err)
	}

	web.Data(c, post)
	log.Info("finished")
	return nil
}

func CreatePostHandler(c echo.Context) error {
	log := logger.Log.Named("[CreatePostHandler]")
	log.Info("started")

	var payload services.CreatePostData
	err := c.Bind(&payload)
	if err != nil {
		log.Error("failed to parse body", zap.Error(err))
		return web.Error(http.StatusBadRequest, err)
	}

	post, err := services.CreatePost(payload)
	if err != nil {
		log.Error("failed to create post", zap.Error(err))
		return web.InternalServerError()
	}

	web.Data(c, post)
	log.Info("finished")
	return nil
}

func UpdatePostByIDHandler(c echo.Context) error {
	log := logger.Log.Named("[UpdatePostByIDHandler]")
	log.Info("started")
	log.Info("finished")
	return nil
}

func DeletePostByIDHandler(c echo.Context) error {
	log := logger.Log.Named("[DeletePostByIDHandler]")
	log.Info("started")

	postID, err := web.ParseID(c)
	if err != nil {
		log.Error("failed to parse id")
		return err
	}

	err = services.DeletePost(postID)
	if err != nil {
		log.Error("failed to delete post", zap.Error(err))
		return web.Error(http.StatusBadRequest, err)
	}

	web.NoContent(c, http.StatusOK)
	log.Info("finished")
	return nil
}
