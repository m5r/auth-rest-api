package main

import (
	"github.com/m5r/auth-rest-api/app"
	"github.com/m5r/auth-rest-api/config"
	"fmt"
)

func main() {
	conf := config.GetConfig()

	api := &app.App{}
	api.Initialize(conf)
	defer api.DB.Close()
	api.Run(fmt.Sprintf(":%d", api.ListeningPort))
}
