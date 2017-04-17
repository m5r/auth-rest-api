package handler

import (
	"time"
	"net/http"
	"encoding/json"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/m5r/auth-rest-api/app/model"
)

func PostAuth(db *gorm.DB, c echo.Context, jwtSigningKey string) error {
	var (
		err     error
		req     map[string]string
		decoder *json.Decoder
		claims  *model.UserClaims
		user    model.User
		token   *jwt.Token
		t       string
	)

	req = make(map[string]string)

	decoder = json.NewDecoder(c.Request().Body)
	if err = decoder.Decode(&req); err != nil {
		return respondError(c, http.StatusBadRequest, err.Error())
	}
	defer c.Request().Body.Close()

	if len(req["password"]) == 0 || len(req["login"]) == 0 {
		return respondMissingCredentials(c)
	}

	db.Where("name = ?", strings.ToLower(req["login"])).First(&user)

	if user.Name == "" {
		db.Where("email = ?", strings.ToLower(req["login"])).First(&user)

		if user.Email == "" {
			return respondIncorrectCredentials(c)
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req["password"])) != nil {
		return respondIncorrectCredentials(c)
	}

	claims = &model.UserClaims{
		Name:  user.Name,
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if t, err = token.SignedString([]byte(jwtSigningKey)); err != nil {
		return respondError(c, http.StatusInternalServerError, "Error during token generation, please try again later")
	}

	return respondJSON(c, http.StatusOK, echo.Map{
		"token": t,
	})
}
