package configure

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type (
	FileConfigure struct {
		GateConfigPath string
		gc             GateConfigure
	}
)

func NewFileConfigure(sc ServerConfigure) *FileConfigure {
	return &FileConfigure{
		GateConfigPath: sc.GateConfigPath,
	}
}

func (fc *FileConfigure) Load() error {
	b, err := ioutil.ReadFile(fc.GateConfigPath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, &fc.gc); err != nil {
		return err
	}

	return nil
}

func (fc *FileConfigure) GateConfigure() GateConfigure {
	return fc.gc
}
