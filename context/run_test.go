package context

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRun(t *testing.T) {
	Run(func(ctx Context) error {
		assert.Equal(t, DefaultConfFile, ctx.ConfFile())

		cfg := ctx.ServiceConfig()
		assert.Equal(t, DefaultHubAddressTCP, cfg.Hub.Address)
		assert.Equal(t, "info", cfg.Logger.Level)
		assert.Equal(t, 15, cfg.Logger.MaxAge)
		assert.Equal(t, 50, cfg.Logger.MaxSize)
		assert.Equal(t, 15, cfg.Logger.MaxBackups)
		return nil
	})
}
