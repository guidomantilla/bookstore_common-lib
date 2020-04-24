package config

import (
	"os"
	"strings"
)

var (
	propertiesConfigFlag bool
	Properties           map[string]string
)

type propertiesConfig struct {
}

func NewPropertiesConfig() *propertiesConfig {
	return &propertiesConfig{}
}

func (propertiesConfig *propertiesConfig) Config() {

	if !propertiesConfigFlag {

		Properties = make(map[string]string, 0)
		for _, env := range os.Environ() {
			pair := strings.SplitN(env, "=", 2)
			Properties[pair[0]] = pair[1]
		}

		propertiesConfigFlag = true
	}
}

func (propertiesConfig *propertiesConfig) Add(prop string, value string) {

	if Properties[prop] == "" {
		Properties[prop] = value
	}
}
