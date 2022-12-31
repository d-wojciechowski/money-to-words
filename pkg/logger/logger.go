package logger

import (
	"ammount-in-words/internal/config"
	"fmt"
	"go.uber.org/zap"
)

func CreateLogger() *zap.Logger {
	var logger *zap.Logger
	var err error
	if config.AppConfig.LogOutput == config.Dev {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic(fmt.Sprintf("Could not initialize logger. Error raised: %v", err))
	}
	return logger
}
