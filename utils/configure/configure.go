package configure

import (
	"errors"
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type (
	Configure interface {
		Load() error
		GateConfigure() GateConfigure
	}

	ServerConfigure struct {
		ConfigType     string `yaml:"config_type"`
		GateConfigPath string `yaml:"gate_config_path"`
	}

	GateConfigure struct {
		Type         string        `json:"type" yaml:"type"`
		Port         string        `json:"port" yaml:"port"`
		ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout"`
	}

	PostHorseConfigure struct {
	}
)

func New(path string) (Configure, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var sc ServerConfigure
	if err := yaml.Unmarshal(b, &sc); err != nil {
		return nil, err
	}

	switch sc.ConfigType {
	case "file":
		fc := NewFileConfigure(sc)
		fc.Load()

		return fc, nil

	default:
		return nil, errors.New("Invalid configuration type.")
	}
}
