package aiicy

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadComposeAppConfigCompatible(t *testing.T) {
	cfg := ComposeAppConfig{
		Version:    "3",
		AppVersion: "v2",
		Services: map[string]ComposeService{
			"test-hub": ComposeService{
				Image: "hub.test.com/aiicy/aiicy-hub",
				Networks: NetworksInfo{
					ServiceNetworks: map[string]ServiceNetwork{
						"test-network": ServiceNetwork{
							Aliases:     []string{"test-net"},
							Ipv4Address: "192.168.0.3",
						},
					},
				},
				Replica:   1,
				Ports:     []string{"1883:1883"},
				Devices:   []string{},
				DependsOn: []string{},
				Restart: RestartPolicyInfo{
					Policy: "always",
					Backoff: BackoffInfo{
						Min:    time.Second,
						Max:    time.Minute * 5,
						Factor: 2,
					},
				},
				Volumes: []ServiceVolume{
					ServiceVolume{
						Source:   "var/db/aiicy/test-hub-conf",
						Target:   "/etc/aiicy",
						ReadOnly: false,
					},
				},
				Command: Command{
					Cmd: []string{"-c", "conf/conf.yml"},
				},
				Environment: Environment{
					Envs: map[string]string{
						"version": "v1",
					},
				},
			},
			"test-timer": ComposeService{
				Image:     "hub.test.com/aiicy/aiicy-timer",
				DependsOn: []string{"test-hub"},
				Replica:   1,
				Ports:     []string{},
				Devices:   []string{},
				Environment: Environment{
					Envs: map[string]string{
						"version": "v2",
					},
				},
				Networks: NetworksInfo{
					ServiceNetworks: map[string]ServiceNetwork{
						"test-network": ServiceNetwork{},
					},
				},
				Volumes: []ServiceVolume{
					ServiceVolume{
						Source:   "var/db/aiicy/test-timer-conf",
						Target:   "/etc/aiicy",
						ReadOnly: true,
					},
				},
				Restart: RestartPolicyInfo{
					Policy: "always",
					Backoff: BackoffInfo{
						Min:    time.Second,
						Max:    time.Minute * 5,
						Factor: 2,
					},
				},
				Command: Command{
					Cmd: []string{"/bin/sh"},
				},
			},
		},
		Volumes: map[string]ComposeVolume{},
		Networks: map[string]ComposeNetwork{
			"test-network": ComposeNetwork{
				Driver:     "bridge",
				DriverOpts: map[string]string{},
				Labels:     map[string]string{},
			},
		},
	}

	composeConfString := `
version: '3'
app_version: v2
services:
  test-hub:
    image: hub.test.com/aiicy/aiicy-hub
    networks:
      test-network:
        aliases:
          - test-net
        ipv4_address: 192.168.0.3
    replica: 1
    ports:
      - 1883:1883
    volumes:
      - var/db/aiicy/test-hub-conf:/etc/aiicy
    command:
      - '-c'
      - conf/conf.yml
    environment:
      - version=v1
  test-timer:
    image: hub.test.com/aiicy/aiicy-timer
    depends_on:
      - test-hub
    replica: 1
    networks:
      - test-network
    volumes:
      - source: var/db/aiicy/test-timer-conf
        target: /etc/aiicy
        read_only: true
    environment:
      version: v2
    command: '/bin/sh'

networks:
  test-network:
`

	dir, err := ioutil.TempDir("", "template")
	assert.NoError(t, err)
	fileName1 := "compose_conf"
	f, err := os.Create(filepath.Join(dir, fileName1))
	defer f.Close()
	_, err = io.WriteString(f, composeConfString)
	assert.NoError(t, err)
	cfg2, err := LoadComposeAppConfigCompatible(filepath.Join(dir, fileName1))
	assert.NoError(t, err)
	assert.Equal(t, cfg, cfg2)
}
