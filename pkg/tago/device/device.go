package device

import (
	"errors"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/admin/types"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/api"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/device/pubsub"
)

// Device Device
type Device interface {
	Data(msg interface{}) (bool, error)
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

func (d *defaultDevice) Data(msg interface{}) (bool, error) {
	var response types.DeviceDataResponse

	err := api.NewClient(d.token).Post("/data", msg, &response)
	if err != nil {
		return false, err
	}

	if !response.Status {
		return false, errors.New(response.Result)
	}

	return response.Status, nil
}

func (d *defaultDevice) PubSub() (pubsub.PubSub, error) {
	ps := pubsub.New(d.token)

	return ps, nil
}
