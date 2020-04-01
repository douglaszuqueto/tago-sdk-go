package device

import (
	"fmt"
	"time"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/admin/types"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/api"
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
	client := *api.NewClient(d.token)

	var response types.DeviceDataResponse

	err := client.Post("/data", getPayload(), &response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ingesting data: %s\n", response.Result)
}

func (d *defaultDevice) PubSub() (pubsub.PubSub, error) {
	fmt.Println("PubSub")

	ps := pubsub.New()

	return ps, nil
}

func getPayload() types.Data {
	return types.Data{
		Variable: "temperature",
		Value:    25.5,
		Time:     time.Now(),
	}
}
