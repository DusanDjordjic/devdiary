package api

import (
	"dev-diary/logger"

	"github.com/labstack/echo"
)

func GetAllTagsHandler(c echo.Context) error {
	log := logger.Log.Named("[GetAllTagsHandler]")
	log.Info("started")
	log.Info("finished")
	return nil
}

func GetTagByIDHandler(c echo.Context) error {
	log := logger.Log.Named("[GetTagByIDHandler]")
	log.Info("started")
	log.Info("finished")
	return nil
}

func CreateTagHandler(c echo.Context) error {
	log := logger.Log.Named("[CreateTagHandler]")
	log.Info("started")
	log.Info("finished")
	return nil
}

func UpdateTagByIDHandler(c echo.Context) error {
	log := logger.Log.Named("[UpdateTagByIDHandler]")
	log.Info("started")
	log.Info("finished")
	return nil
}

func DeleteTagByIDHandler(c echo.Context) error {
	log := logger.Log.Named("[DeleteTagByIDHandler]")
	log.Info("started")
	log.Info("finished")
	return nil
}
