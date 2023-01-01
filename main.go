package main

import (
	"ammount-in-words/internal/config"
	"ammount-in-words/internal/routers"
	"ammount-in-words/pkg/logger"
)

func main() {
	log := logger.Logger
	log.Infow("Currency converter application init procedure: start")
	routers.InitRouter()
	log.Infow("Currency converter application init procedure: end")
	log.Sync()
	err := routers.Router.Run(config.AppConfig.AppUrl)
	if err != nil {
		log.Errorw("Error during web server creation.", "error", err.Error())
	}
}
