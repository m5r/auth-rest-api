package handler

import (
	"time"
	"net/http"
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	_ "github.com/lib/pq"
	"github.com/m5r/auth-rest-api/app/model"
)

func PostRefresh(c echo.Context, jwtSigningKey string) error {
	var (
		err      error
		req      map[string]string
		decoder  *json.Decoder
		claims   *model.UserClaims
		oldToken *jwt.Token
		newToken *jwt.Token
		t        string
	)

	req = make(map[string]string)

	decoder = json.NewDecoder(c.Request().Body)
	if err = decoder.Decode(&req); err != nil {
		return respondError(c, http.StatusBadRequest, err.Error())
	}
	defer c.Request().Body.Close()

	if len(req["token"]) == 0 {
		return respondMissingToken(c)
	}

	oldToken, err = jwt.ParseWithClaims(req["token"], &model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSigningKey), nil
	})
	if err != nil {
		return respondError(c, http.StatusBadRequest, err.Error())
	}

	claims, ok := oldToken.Claims.(*model.UserClaims)
	if !ok || !oldToken.Valid {
		return respondInvalidToken(c)
	}

	claims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()

	newToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if t, err = newToken.SignedString([]byte(jwtSigningKey)); err != nil {
		return respondError(c, http.StatusInternalServerError, "Error during token generation, please try again later")
	}

	return respondJSON(c, http.StatusOK, echo.Map{
		"token": t,
	})
}
