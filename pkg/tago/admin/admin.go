package admin

import "fmt"

// Manager Manager
type Manager interface {
	Device() (Device, error)
	Bucket() (Bucket, error)

	Info()
}

type manager struct {
	token string
}

// New New
func New(token string) Manager {
	return &manager{
		token: token,
	}
}

// Info Info
func (m *manager) Info() {
	fmt.Println("Admin Info:", m.token)
}

// Device Device
func (m *manager) Device() (Device, error) {
	fmt.Println("Device manager")

	e := newDevice()

	return e, nil
}

// Bucket Bucket
func (m *manager) Bucket() (Bucket, error) {
	fmt.Println("Bucket manager")

	e := newBucket()

	return e, nil
}
