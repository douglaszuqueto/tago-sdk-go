package device

import (
	"fmt"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/device/pubsub"
)

// Device Device
type Device interface {
	Data()
	PubSub() (pubsub.PubSub, error)

	Info()
}

type defaultDevice struct {
	token string
}

// New New
func New(token string) Device {
	return &defaultDevice{
		token: token,
	}
}

func (d *defaultDevice) Info() {
	fmt.Println("Info")
}

func (d *defaultDevice) Data() {
	fmt.Println("Data")
}

func (d *defaultDevice) PubSub() (pubsub.PubSub, error) {
	fmt.Println("PubSub")

	ps := pubsub.New()

	return ps, nil
}
