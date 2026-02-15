package main

import (
	"log"

	"github.com/abhay786-20/fraud-auth-service/internal/bootstrap"
)

func main() {

	app, err := bootstrap.NewApplication()
	if err != nil {
		log.Fatal(err)
	}

	addr := app.Config.Server.Host + ":" + app.Config.Server.Port
	app.Logger.Info("Server starting on " + addr)

	if err := app.Router.Engine.Run(addr); err != nil {
		log.Fatal("Server failed to start: " + err.Error())
	}
}
