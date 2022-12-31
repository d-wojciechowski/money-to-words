package config

import (
	"github.com/joho/godotenv"
)

type LogType int

const (
	Prod int = 0
	Dev      = 1
)

type appConfig struct {
	LogLevel  string
	LogOutput LogType
}

var AppConfig = getConfiguration()

func getConfiguration() *appConfig {
	_ = godotenv.Load() //todo consider panic or other halt of app on dotEnv fail

	return &appConfig{
		LogLevel:  "debug",
		LogOutput: Dev,
	}
}
