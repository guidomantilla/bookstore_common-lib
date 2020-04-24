package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPropertiesConfig_Init(testing *testing.T) {

	assert.Nil(testing, Properties, "Properties must be null")
	assert.False(testing, propertiesConfigFlag, "propertiesConfigFlag must be false")
}

func TestPropertiesConfig_Config(testing *testing.T) {
	propertiesConfig := NewPropertiesConfig()
	propertiesConfig.Config()

	assert.NotNil(testing, Properties, "Properties must not be null")
	assert.True(testing, propertiesConfigFlag, "propertiesConfigFlag must be true")
}
