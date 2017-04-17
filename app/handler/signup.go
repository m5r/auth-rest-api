package handler

import (
	"net/http"
	"encoding/json"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/m5r/auth-rest-api/app/model"
	"strings"
)

func PostSignup(db *gorm.DB, c echo.Context) error {
	var (
		err     error
		req     map[string]string
		decoder *json.Decoder
		user    model.User
		encPwd  []byte
	)

	req = make(map[string]string)

	decoder = json.NewDecoder(c.Request().Body)
	if err = decoder.Decode(&req); err != nil {
		return respondError(c, http.StatusBadRequest, err.Error())
	}
	defer c.Request().Body.Close()

	if len(req["password"]) == 0 || len(req["name"]) == 0 {
		return respondMissingCredentials(c)
	}

	db.Where("name = ?", strings.ToLower(req["name"])).First(&user)
	if user.Name != "" {
		return respondTakenUsername(c)
	}

	db.Where("email = ?", strings.ToLower(req["email"])).First(&user)
	if user.Email != "" {
		return respondTakenEmail(c)
	}

	encPwd, _ = bcrypt.GenerateFromPassword([]byte(req["password"]), bcrypt.DefaultCost)
	user = model.User{
		Name:     strings.ToLower(req["name"]),
		Password: string(encPwd),
		Email:    req["email"],
	}

	db.NewRecord(user)
	db.Create(&user)

	return respondJSON(c, http.StatusOK, echo.Map{
		"message": "Account successfully created",
	})
}
