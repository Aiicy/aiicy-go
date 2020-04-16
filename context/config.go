package context

import (
	"github.com/aiicy/aiicy-go/logger"
	"github.com/aiicy/aiicy-go/mqtt"
)

// ServiceConfig base config of service
type ServiceConfig struct {
	Hub    mqtt.ClientConfig `yaml:"hub" json:"hub"`
	Logger logger.Config     `yaml:"logger" json:"logger"`
}
