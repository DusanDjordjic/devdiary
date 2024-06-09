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

func GetAllTagsHandler(c echo.Context) error {
	log := logger.Log.Named("[GetAllTagsHandler]")
	log.Info("started")

	tags, total, err := services.GetAllTags()
	if err != nil {
		log.Error("failed to get tags", zap.Error(err))
		return web.Error(http.StatusBadRequest, err)
	}

	web.DataWithCount(c, models.DTOTags(tags), total)
	log.Info("finished")
	return nil
}

func GetTagByIDHandler(c echo.Context) error {
	log := logger.Log.Named("[GetTagByIDHandler]")
	log.Info("started")

	tagID, err := web.ParseID(c)
	if err != nil {
		log.Error("failed to parse id")
		return err
	}

	tag, err := services.GetTagByID(tagID)
	if err != nil {
		log.Error("failed to get tag by id", zap.Error(err))
		return web.Error(http.StatusNotFound, err)
	}

	web.Data(c, tag)
	log.Info("finished")
	return nil
}

func CreateTagHandler(c echo.Context) error {
	log := logger.Log.Named("[CreateTagHandler]")
	log.Info("started")

	var payload services.CreateTagData
	err := c.Bind(&payload)
	if err != nil {
		log.Error("failed to parse body", zap.Error(err))
		return web.Error(http.StatusBadRequest, err)
	}

	tag, err := services.CreateTag(payload)
	if err != nil {
		log.Error("failed to create tag", zap.Error(err))
		return web.InternalServerError()
	}

	web.Data(c, tag)
	log.Info("finished")
	return nil
}

func UpdateTagByIDHandler(c echo.Context) error {
	log := logger.Log.Named("[UpdatePostByIDHandler]")
	log.Info("started")
	log.Info("finished")
	return nil
}

func DeleteTagByIDHandler(c echo.Context) error {
	log := logger.Log.Named("[DeleteTagByIDHandler]")
	log.Info("started")

	tagID, err := web.ParseID(c)
	if err != nil {
		log.Error("failed to parse id")
		return err
	}

	err = services.DeleteTag(tagID)
	if err != nil {
		log.Error("failed to delete tag", zap.Error(err))
		return web.Error(http.StatusBadRequest, err)
	}

	web.NoContent(c, http.StatusOK)
	log.Info("finished")
	return nil
}
