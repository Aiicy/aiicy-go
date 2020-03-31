package mqtt

import (
	"github.com/aiicy/aiicy-go/utils"
	"time"
)

// QOSTopic topic and qos
type QOSTopic struct {
	QOS   uint32 `yaml:"qos" json:"qos" validate:"min=0, max=1"`
	Topic string `yaml:"topic" json:"topic" validate:"nonzero"`
}

// Subscriptions subscriptions
type Subscriptions []QOSTopic

// ClientConfig client config
type ClientConfig struct {
	Address           string `yaml:"address" json:"address"`
	Username          string `yaml:"username" json:"username"`
	Password          string `yaml:"password" json:"password"`
	utils.Certificate `yaml:",inline" json:",inline"`
	ClientID          string        `yaml:"clientid" json:"clientid"`
	CleanSession      bool          `yaml:"cleansession" json:"cleansession"`
	Timeout           time.Duration `yaml:"timeout" json:"timeout" default:"30s"`
	Interval          time.Duration `yaml:"interval" json:"interval" default:"1m"`
	KeepAlive         time.Duration `yaml:"keepalive" json:"keepalive" default:"10m"`
	BufferSize        int           `yaml:"buffersize" json:"buffersize" default:"10"`
	ValidateSubs      bool          `yaml:"validatesubs" json:"validatesubs"`
	Subscriptions     Subscriptions `yaml:"subscriptions" json:"subscriptions" default:"[]"`
}
