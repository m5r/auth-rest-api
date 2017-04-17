package handler

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/m5r/auth-rest-api/app/model"
)

func GetRestricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.UserClaims)
	name := claims.Name

	return c.String(http.StatusOK, "Welcome "+name+"!")
}
