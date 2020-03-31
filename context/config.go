package context

import (
	"github.com/aiicy/aiicy-go/logger"
	"github.com/aiicy/aiicy-go/mqtt"
	"github.com/aiicy/aiicy-go/utils"
)

// ServiceConfig base config of service
type ServiceConfig struct {
	Hub    mqtt.ClientConfig `yaml:"hub" json:"hub"`
	Logger logger.LogConfig  `yaml:"logger" json:"logger"`
}

// LoadConfigFile load config from file
func LoadConfigFile(cfg interface{}, confFile string) error {
	if utils.FileExists(confFile) {
		return utils.LoadYAML(confFile, cfg)
	}
	return utils.UnmarshalYAML(nil, cfg)
}