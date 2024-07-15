package handlers

import (
	"dev-diary/logger"
	"net/http"

	"github.com/labstack/echo"
	"go.uber.org/zap"
)

func CatchAllPage(e echo.Context) error {
	log := logger.Log.Named("[CatchAllPage")
	log.Info("started")
	log.Info("Url not found", zap.String("URL", e.Request().URL.String()))
	log.Info("finished")
	return e.Redirect(http.StatusTemporaryRedirect, "/")
}
