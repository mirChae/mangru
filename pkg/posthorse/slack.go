package posthorse

import "github.com/mirChae/mangru/utils/configure"

type (
	SlackPostHorse struct {
		URL string
	}
)

func NewSlackPostHorse(pc configure.PostHorseConfigure) (PostHorse, error) {
	return &SlackPostHorse{}, nil
}

func (sph *SlackPostHorse) SendMessage() error {
	return nil
}
