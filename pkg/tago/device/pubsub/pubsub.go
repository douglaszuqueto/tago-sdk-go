package pubsub

import "fmt"

// PubSub PubSub
type PubSub interface {
	Sub()
	Pub()
	Debug()
}

type ps struct{}

// New New
func New() PubSub {
	return &ps{}
}

func (d *ps) Sub() {
	fmt.Println("Sub")
}

func (d *ps) Pub() {
	fmt.Println("Pub")
}

func (d *ps) Debug() {
	fmt.Println("Debug")
}
