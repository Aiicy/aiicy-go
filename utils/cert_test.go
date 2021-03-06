package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTLSConfigServer(t *testing.T) {
	tls, err := NewTLSConfigServer(Certificate{Key: "../example/var/lib/aiicy/testcert/server.key"})
	assert.Error(t, err)

	tls, err = NewTLSConfigServer(Certificate{Cert: "../example/var/lib/aiicy/testcert/server.pem"})
	assert.Error(t, err)

	c := Certificate{
		Key:  "../example/var/lib/aiicy/testcert/server.key",
		Cert: "../example/var/lib/aiicy/testcert/server.pem",
	}

	tls, err = NewTLSConfigServer(c)
	assert.NoError(t, err)
	assert.NotEmpty(t, tls)
}

func TestNewTLSConfigClient(t *testing.T) {
	tls, err := NewTLSConfigClient(Certificate{Key: "../example/var/lib/aiicy/testcert/client.key"})
	assert.Error(t, err)

	tls, err = NewTLSConfigClient(Certificate{Cert: "../example/var/lib/aiicy/testcert/client.pem"})
	assert.Error(t, err)
	assert.Empty(t, tls)

	c := Certificate{
		Key:  "../example/var/lib/aiicy/testcert/client.key",
		Cert: "../example/var/lib/aiicy/testcert/client.pem",
	}
	tls, err = NewTLSConfigClient(c)
	assert.NoError(t, err)
	assert.NotEmpty(t, tls)
}
