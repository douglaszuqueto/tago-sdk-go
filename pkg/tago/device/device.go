package device

import (
	"fmt"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/admin/types"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/api"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/device/pubsub"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/util"
)

// Device Device
type Device interface {
	Data()
	PubSub() (pubsub.PubSub, error)
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

func (d *defaultDevice) Data() {
	var response types.DeviceDataResponse

	err := api.NewClient(d.token).Post("/data", util.GeneratePayload(), &response)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ingesting data: %s\n", response.Result)
}

func (d *defaultDevice) PubSub() (pubsub.PubSub, error) {
	ps := pubsub.New(d.token)

	return ps, nil
}
