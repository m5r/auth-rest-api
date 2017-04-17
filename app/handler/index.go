package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetIndex(c echo.Context) error {
	return respondJSON(c, http.StatusOK, echo.Map{
		"message": "Hello",
	})
}
