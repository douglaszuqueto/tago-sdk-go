package admin

import "fmt"

// Bucket Bucket
type Bucket interface {
	Get()
	List()
}

type bucket struct {
}

func newBucket() Bucket {
	return &bucket{}
}

// Get Get
func (d *bucket) Get() {
	fmt.Println("Get")
}

// List List
func (d *bucket) List() {
	fmt.Println("List")
}
