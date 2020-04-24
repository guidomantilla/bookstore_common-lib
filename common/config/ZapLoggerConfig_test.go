package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZapLoggerConfig_Init(testing *testing.T) {

	assert.Nil(testing, ZapLogger, "ZapLogger must be null")
	assert.False(testing, zapLoggerConfigFlag, "zapLoggerConfigFlag must be false")
}

func TestZapLoggerConfig_Config(testing *testing.T) {

	zapLoggerConfig := NewZapLoggerConfig()
	zapLoggerConfig.Config("dev")

	assert.NotNil(testing, ZapLogger, "ZapLogger must not be null")
	assert.True(testing, zapLoggerConfigFlag, "zapLoggerConfigFlag must be true")
}
