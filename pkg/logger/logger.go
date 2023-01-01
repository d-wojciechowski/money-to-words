package logger

import (
	"ammount-in-words/internal/config"
	"fmt"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

func init() {
	Logger = CreateLogger().Sugar()
}

func CreateLogger() *zap.Logger {
	var logger *zap.Logger
	var err error
	if config.AppConfig.LogProfile == config.Dev {
		logger, err = NewCustomDevelopmentConfig().Build()
	} else {
		logger, err = NewCustomProductionConfig().Build()
	}
	if err != nil {
		panic(fmt.Sprintf("Could not initialize logger. Error raised: %v", err))
	}
	return logger
}

func NewCustomProductionConfig() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr", config.AppConfig.LogPath},
		ErrorOutputPaths: []string{"stderr", config.AppConfig.LogPath},
	}
}

func NewCustomDevelopmentConfig() zap.Config {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
		OutputPaths:      []string{"stderr", config.AppConfig.LogPath},
		ErrorOutputPaths: []string{"stderr", config.AppConfig.LogPath},
	}
}
