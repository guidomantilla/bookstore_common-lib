package config

import (
	"go.uber.org/zap"
)

var (
	zapLoggerConfigFlag bool
	ZapLogger           *zap.Logger
)

type zapLoggerConfig struct {
}

func NewZapLoggerConfig() *zapLoggerConfig {
	return &zapLoggerConfig{}
}

func (zapLoggerConfig *zapLoggerConfig) Config(environment string) {

	if !zapLoggerConfigFlag {

		var logConfig zap.Config
		if environment == "pro" {
			logConfig = zap.NewProductionConfig()
		} else {
			logConfig = zap.NewDevelopmentConfig()
		}

		var err error
		ZapLogger, err = logConfig.Build()
		if err != nil {
			panic(err)
		}

		zapLoggerConfigFlag = true
	}
}
