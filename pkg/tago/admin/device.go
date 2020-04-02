package admin

import (
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/admin/types"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/api"
)

// Device Device
type Device interface {
	Get(deviceID string) (*types.DeviceGet, error)
	List() ([]types.Device, error)
	Token(deviceID string) ([]types.DeviceToken, error)
}

type device struct {
	client api.Client
}

func newDevice(client api.Client) Device {
	return &device{
		client: client,
	}
}

// Get Get
func (d *device) Get(deviceID string) (*types.DeviceGet, error) {
	var device types.DeviceGetResponse

	err := d.client.Get("/device/"+deviceID, &device)
	if err != nil {
		return nil, err
	}

	return &device.Result, nil
}

// List List
func (d *device) List() ([]types.Device, error) {
	var devices types.DeviceListResponse

	err := d.client.Get("/device", &devices)
	if err != nil {
		return nil, err
	}

	return devices.Result, nil
}

// Token Token
func (d *device) Token(deviceID string) ([]types.DeviceToken, error) {
	var device types.DeviceTokenResponse

	err := d.client.Get("/device/token/"+deviceID, &device)
	if err != nil {
		return nil, err
	}

	return device.Result, nil
}
