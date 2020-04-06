package context

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestContext(t *testing.T) {
	os.Setenv(EnvKeyConfFile, "file")
	os.Setenv(EnvKeyServiceInstanceName, "instance")
	os.Setenv(EnvKeyServiceName, "service")

	ctx := NewContext("")
	assert.Equal(t, "service", ctx.ServiceName())
	assert.Equal(t, "instance", ctx.InstanceName())
	assert.Equal(t, "file", ctx.ConfFile())
	cfg := ctx.ServiceConfig()
	assert.Equal(t, DefaultHubAddressTCP, cfg.Hub.Address)
	assert.Equal(t, 15, cfg.Logger.MaxAge)
	assert.Equal(t, 50, cfg.Logger.MaxSize)
	assert.Equal(t, 15, cfg.Logger.MaxBackups)

	ctx = NewContext("../example/etc/aiicy/service.yml")
	assert.Equal(t, "instance", ctx.InstanceName())
	assert.Equal(t, "service", ctx.ServiceName())
	assert.Equal(t, "../example/etc/aiicy/service.yml", ctx.ConfFile())
	cfg = ctx.ServiceConfig()
	assert.Equal(t, DefaultHubAddressSSL, cfg.Hub.Address)
	assert.Equal(t, "debug", cfg.Logger.Level)
	assert.Equal(t, "console", cfg.Logger.Encoding)
	assert.Equal(t, 15, cfg.Logger.MaxAge)
	assert.Equal(t, 50, cfg.Logger.MaxSize)
	assert.Equal(t, 15, cfg.Logger.MaxBackups)
}
