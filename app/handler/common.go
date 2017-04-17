package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type jsonResponse struct {
	Status int `json:"status"`
	Data   interface{} `json:"data"`
}
type jsonErrorResponse struct {
	Status int `json:"status"`
	Error  interface{} `json:"error"`
}

func respondJSON(c echo.Context, status int, payload interface{}) error {
	return c.JSONPretty(status, jsonResponse{
		status,
		payload,
	}, "  ")
}

func respondError(c echo.Context, status int, message string) error {
	return c.JSONPretty(status, jsonErrorResponse{
		status,
		map[string]string{
			"message": message,
		},
	}, "  ")
}

func respondIncorrectCredentials(c echo.Context) error {
	return respondError(c, http.StatusUnauthorized, "Incorrect password or login")
}

func respondMissingCredentials(c echo.Context) error {
	return respondError(c, http.StatusBadRequest, "Login and password are both required")
}

func respondMissingToken(c echo.Context) error {
	return respondError(c, http.StatusBadRequest, "Token is missing")
}

func respondInvalidToken(c echo.Context) error {
	return respondError(c, http.StatusForbidden, "Token is invalid")
}

func respondTakenUsername(c echo.Context) error {
	return respondError(c, http.StatusForbidden, "Username is already taken")
}

func respondTakenEmail(c echo.Context) error {
	return respondError(c, http.StatusForbidden, "Email address is already taken")
}
