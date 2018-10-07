package gate

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/mirChae/mangru/utils/configure"
)

type HTTPGate struct {
	e *echo.Echo
	s *http.Server
}

func NewHTTPGate(gc configure.GateConfigure) (*HTTPGate, error) {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	s := &http.Server{
		Addr:         ":" + gc.Port,
		ReadTimeout:  gc.ReadTimeout * time.Second,
		WriteTimeout: gc.WriteTimeout * time.Second,
	}

	return &HTTPGate{
		e: e,
		s: s,
	}, nil
}

func (hg HTTPGate) Open() error {
	if err := hg.e.StartServer(hg.s); err != nil {
		return err
	}
	return nil
}
