package config

import (
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap/zapcore"
	"os"
)

type LogProfile string

const (
	Prod LogProfile = "prod"
	Dev             = "dev"
)

type envFileKey string

const (
	LOG_PROFILE envFileKey = "LOG_PROFILE"
	LOG_PATH               = "LOG_PATH"
	LOG_LEVEL              = "LOG_LEVEL"
	APP_URL                = "APP_URL"
)

type appConfig struct {
	LogLevel   zapcore.Level
	LogProfile LogProfile
	LogPath    string
	AppUrl     string
}

var AppConfig = getConfiguration()

func getConfiguration() *appConfig {

	return &appConfig{
		LogLevel:   getLogLevel(),
		LogProfile: LogProfile(getEnvOrDefault(LOG_PROFILE, Dev)),
		LogPath:    getEnvOrDefault(LOG_PATH, "./app.log"),
		AppUrl:     getEnvOrDefault(APP_URL, "0.0.0.0:8081"),
	}
}

func getLogLevel() zapcore.Level {
	logLevel := getEnvOrDefault(LOG_LEVEL, "info")
	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		panic("Could not initialize app configuration. LogLevel is invalid! Provided level: " + logLevel)
	}
	return level
}

func getEnvOrDefault(key envFileKey, def string) string {
	envVal := os.Getenv(string(key))
	if envVal == "" {
		return def
	}
	return envVal
}
