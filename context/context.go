package context

import (
	"github.com/aiicy/aiicy-go/logger"
	"github.com/aiicy/aiicy-go/utils"
	"os"
	"os/signal"
	"syscall"
)

// Env keys
const (
	EnvKeyConfFile            = "AIICY_CONF_FILE"
	EnvKeyServiceName         = "AIICY_SERVICE_NAME"
	EnvKeyServiceInstanceName = "AIICY_SERVICE_INSTANCE_NAME"
)

// API keys
const (
	DefaultHubAddressTCP = "tcp://aiicy-hub:1883"
	DefaultHubAddressSSL = "ssl://aiicy-hub:8883"
)

// Path keys
const (
	// DefaultConfFile config path of the service by default
	DefaultConfFile = "etc/aiicy/service.yml"
)

// Context of service
type Context interface {
	// InstanceName returns instance name.
	InstanceName() string
	// ServiceName returns service name.
	ServiceName() string
	// ConfFile returns config file.
	ConfFile() string
	// ServiceConfig returns service config.
	ServiceConfig() ServiceConfig
	// LoadCustomConfig loads custom config, if path is empty, will load config from default path.
	LoadCustomConfig(cfg interface{}, files ...string) error
	// returns logger interface
	Log() *logger.Logger
	// waiting to exit, receiving SIGTERM and SIGINT signals
	Wait()
	// returns wait channel
	WaitChan() <-chan os.Signal
}

type ctx struct {
	cfg ServiceConfig
	log *logger.Logger

	instanceName string
	serviceName  string
	confFile     string
	httpAddress  string
	mqttAddress  string
	linkAddress  string
}

// NewContext creates a new context
func NewContext(confFile string) Context {
	if confFile == "" {
		confFile = os.Getenv(EnvKeyConfFile)
	}
	c := &ctx{
		confFile:     confFile,
		serviceName:  os.Getenv(EnvKeyServiceName),
		instanceName: os.Getenv(EnvKeyServiceInstanceName),
	}
	var fs []logger.Field
	if c.serviceName != "" {
		fs = append(fs, logger.String("service", c.serviceName))
	}
	if c.instanceName != "" {
		fs = append(fs, logger.String("instance", c.instanceName))
	}
	c.log = logger.With(fs...)

	err := c.LoadCustomConfig(&c.cfg)
	if err != nil {
		c.log.Error("failed to load service config, to use default config", logger.Error(err))
		utils.UnmarshalYAML(nil, &c.cfg)
	}

	c.log = logger.New(c.cfg.Logger, fs...)

	if c.cfg.Hub.Address == "" {
		c.log.Error("hub not configured, to use default config")
		if c.cfg.Hub.Key == "" {
			c.cfg.Hub.Address = DefaultHubAddressTCP
		} else {
			c.cfg.Hub.Address = DefaultHubAddressSSL
		}
	}

	c.log.Debug("context is created", logger.String("file", confFile), logger.Any("conf", c.cfg))
	return c
}

func (c *ctx) InstanceName() string {
	return c.instanceName
}

func (c *ctx) ServiceName() string {
	return c.serviceName
}

func (c *ctx) ConfFile() string {
	return c.confFile
}

func (c *ctx) ServiceConfig() ServiceConfig {
	return c.cfg
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

func (c *ctx) Log() *logger.Logger {
	return c.log
}

func (c *ctx) Wait() {
	<-c.WaitChan()
}

func (c *ctx) WaitChan() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	signal.Ignore(syscall.SIGPIPE)
	return sig
}
