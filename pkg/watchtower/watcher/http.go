package watcher

import "github.com/labstack/echo"

type HttpWatcher struct {
	e *echo.Echo
}

func (h HttpWatcher) Init() error {
	return nil
}

func (h HttpWatcher) Watch() error {
	return nil
}
