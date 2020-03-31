package context

import (
	"github.com/aiicy/aiicy-go/logger"
	"github.com/aiicy/aiicy-go/utils"
)

type ctx struct {
	cfg ServiceConfig
	log *logger.Logger

	nodeName    string
	appName     string
	serviceName string
	confFile    string
	httpAddress string
	mqttAddress string
	linkAddress string
}

func (c *ctx) LoadCustomConfig(cfg interface{}, files ...string) error {
	f := c.confFile
	if len(files) > 0 && len(files[0]) > 0 {
		f = files[0]
	}
	if utils.FileExists(f) {
		return utils.LoadYAML(f, cfg)
	}
	return utils.UnmarshalYAML(nil, cfg)
}
