package watcher

type (
	Watcher interface {
		Init() error
		Watch() error
	}
)

func Init(config map[string]interface{}) error {
	return nil
}
