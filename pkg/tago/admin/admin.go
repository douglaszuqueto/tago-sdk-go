package admin

import (
	"fmt"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/api"
)

// Manager Manager
type Manager interface {
	Device() (Device, error)
	Bucket() (Bucket, error)

	Info()
}

type manager struct {
	token  string
	client api.Client
}

// New New
func New(token string) Manager {
	m := &manager{
		token: token,
	}

	m.client = *api.NewClient(token)

	return m
}

// Info Info
func (m *manager) Info() {
	fmt.Println("Admin Info:", m.token)
}

// Device Device
func (m *manager) Device() (Device, error) {
	fmt.Println("Device manager")

	e := newDevice(m.client)

	return e, nil
}

// Bucket Bucket
func (m *manager) Bucket() (Bucket, error) {
	fmt.Println("Bucket manager")

	e := newBucket()

	return e, nil
}
