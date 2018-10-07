package watchtower

import "github.com/mirChae/mangru/pkg/watchtower/watcher"

type (
	Watchtower struct {
		fromGate     <-chan map[string]interface{}
		toPigeonGram chan<- map[string]interface{}
		watchers     []watcher.Watcher
	}
)
