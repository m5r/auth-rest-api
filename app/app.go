package app

import (
	"log"
	"fmt"

	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"

	"github.com/m5r/auth-rest-api/config"
	"github.com/m5r/auth-rest-api/app/handler"
	"github.com/m5r/auth-rest-api/app/model"
)

type App struct {
	Router        *echo.Echo
	DB            *gorm.DB
	JWTSigningKey string
	ListeningPort int
}

func (a *App) Initialize(config *config.Config) {
	var (
		err    error
		dbArgs string
	)

	a.JWTSigningKey = config.JWTSigningKey
	a.ListeningPort = config.ListeningPort

	dbArgs = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.DB.Host,
		config.DB.Username,
		config.DB.Name,
		config.DB.Password,
	)

	a.DB, err = gorm.Open(config.DB.Dialect, dbArgs)
	if err != nil {
		log.Fatal("Could not connect to database. " + err.Error())
	}

	a.DB = model.DBMigrate(a.DB)
	a.DB.LogMode(config.Debug)

	a.Router = echo.New()
	a.SetRouter()
}

func (a *App) SetRouter() {
	a.Router.Use(middleware.Logger())
	a.Router.Use(middleware.Gzip())
	a.Router.Use(middleware.Recover())

	a.Router.GET("/", a.GetIndex)
	a.Router.POST("/auth", a.PostAuth)
	a.Router.POST("/signup", a.PostSignup)
	a.Router.POST("/refresh", a.PostRefresh)

	r := a.Router.Group("/restricted")

	jwtConfig := middleware.JWTConfig{
		Claims:     &model.UserClaims{},
		SigningKey: []byte(a.JWTSigningKey),
	}

	r.Use(middleware.JWTWithConfig(jwtConfig))
	r.GET("", handler.GetRestricted)
}

func (a *App) GetIndex(c echo.Context) error {
	return handler.GetIndex(c)
}

func (a *App) PostAuth(c echo.Context) error {
	return handler.PostAuth(a.DB, c, a.JWTSigningKey)
}

func (a *App) PostSignup(c echo.Context) error {
	return handler.PostSignup(a.DB, c)
}

func (a *App) PostRefresh(c echo.Context) error {
	return handler.PostRefresh(c, a.JWTSigningKey)
}

func (a *App) Run(port string) {
	a.Router.Logger.Fatal(a.Router.Start(port))
}
