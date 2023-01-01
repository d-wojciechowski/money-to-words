package config

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
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

	logPath := getEnvOrDefault(LOG_PATH, "./app.log")
	createFileInPath(logPath)
	return &appConfig{
		LogLevel:   getLogLevel(),
		LogProfile: LogProfile(getEnvOrDefault(LOG_PROFILE, Dev)),
		LogPath:    logPath,
		AppUrl:     getEnvOrDefault(APP_URL, "0.0.0.0:8081"),
	}
}

func createFileInPath(path string) {
	err := os.MkdirAll(filepath.Dir(path), 0770)
	if err != nil {
		panic(fmt.Sprintf("Could not create file with given path: %s. Error message %s ", path, err.Error()))
	}
	f, err := os.Create(path)
	if err != nil {
		panic(fmt.Sprintf("Could not create file with given path: %s. Error message %s ", path, err.Error()))
	}
	f.Close()
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
