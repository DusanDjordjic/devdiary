package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ParseID(c echo.Context) (uint, error) {
	s := c.Param("id")
	id, err := strconv.ParseUint(s, 10, 64)

	if err != nil {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "invalid id param")
	}

	return uint(id), nil
}
