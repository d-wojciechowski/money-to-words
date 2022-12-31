package main

import (
	"ammount-in-words/internal/config"
	"ammount-in-words/internal/routers"
)

func init() {
	config.InitConfiguration()
	routers.InitRouter()
}

func main() {
	routers.Router.Run("0.0.0.0:8081")
}
