package gate

import (
	"errors"
	"strings"

	"github.com/mirChae/mangru/utils/configure"
)

type (
	Gate interface {
		Open() error
	}
)

func New(gc configure.GateConfigure) (Gate, error) {
	switch strings.ToUpper(gc.Type) {
	case "HTTP":
		ht, err := NewHTTPGate(gc)
		if err != nil {
			return nil, err
		}

		return ht, nil

	default:
		return nil, errors.New("InValid gate type.")
	}
}
