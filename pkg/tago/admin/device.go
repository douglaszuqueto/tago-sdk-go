package admin

import (
	"fmt"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/admin/types"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/api"
)

// Device Device
type Device interface {
	Get()
	List()
	Token()
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
func (d *device) Get() {
	fmt.Println("Get")

	var device types.DeviceGetResponse

	err := d.client.Get("/device/"+"5e83e40caf0d7a001b2b203e", &device)
	if err != nil {
		panic(err)
	}

	fmt.Println(device.Result.Name)
}

// List List
func (d *device) List() {
	fmt.Println("List")

	var device types.DeviceListResponse

	err := d.client.Get("/device", &device)
	if err != nil {
		panic(err)
	}

	fmt.Println(device.Result)
}

// Token Token
func (d *device) Token() {
	fmt.Println("Token")

	var device types.DeviceTokenResponse

	err := d.client.Get("/device/token/"+"5e83e40caf0d7a001b2b203e", &device)
	if err != nil {
		panic(err)
	}

	fmt.Println(device.Result)
}
