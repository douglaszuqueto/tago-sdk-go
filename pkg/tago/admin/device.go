package admin

import "fmt"

// Device Device
type Device interface {
	Get()
	List()
	Token()
}

type device struct {
}

func newDevice() Device {
	return &device{}
}

// Get Get
func (d *device) Get() {
	fmt.Println("Get")
}

// List List
func (d *device) List() {
	fmt.Println("List")
}

// Token Token
func (d *device) Token() {
	fmt.Println("Token")
}
