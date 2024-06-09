package web

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

func Error(statusCode int, err error) *echo.HTTPError {
	return echo.NewHTTPError(statusCode, err)
}

func InternalServerError() *echo.HTTPError {
	return echo.NewHTTPError(http.StatusInternalServerError, errors.New("internal server error"))
}

func Data(c echo.Context, data any) {
	c.JSON(http.StatusOK, data)
}

type DataWithCountResponse struct {
	Data  any   `json:"data"`
	Total int64 `json:"total"`
}

func DataWithCount(c echo.Context, data any, total int64) {
	c.JSON(http.StatusOK, DataWithCountResponse{Data: data, Total: total})
}

func NoContent(c echo.Context, statusCode int) {
	c.NoContent(statusCode)
}
