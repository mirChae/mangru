package main

import (
	"github.com/mirChae/mangru/pkg/gate"
	"github.com/mirChae/mangru/utils/configure"
)

func main() {
	c, err := configure.New("configs/server.yaml")
	if err != nil {
		panic(err)
	}

	g, err := gate.New(c.GateConfigure())
	if err != nil {
		panic(err)
	}

	if err := g.Open(); err != nil {
		panic(err)
	}
}
