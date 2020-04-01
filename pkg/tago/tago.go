package tago

import (
	"fmt"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/admin"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/device"
)

// Tago Tago
type Tago interface {
	Admin(token string) (admin.Manager, error)
	Device(token string) (device.Device, error)
	Info()
}

// Default Default
type Default struct {
}

var version = "0.0.1"

// New New
func New() Tago {
	return &Default{}
}

// Info Info
func (t *Default) Info() {
	fmt.Println("Tago SDK version:", version)
}

// Admin Admin
func (t *Default) Admin(token string) (admin.Manager, error) {
	manager := admin.New(token)

	return manager, nil
}

// Device Device
func (t *Default) Device(token string) (device.Device, error) {
	e := device.New(token)

	return e, nil
}
