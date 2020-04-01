package main

import (
	"fmt"
	"os"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago"
)

func main() {
	cli := tago.New()
	cli.Info()

	fmt.Println()

	// Admin manager

	admin, err := cli.Admin(os.Getenv("ADMIN_TOKEN"))
	if err != nil {
		panic(err)
	}

	admin.Info()

	// Device manager

	device, err := admin.Device()
	if err != nil {
		panic(err)
	}

	device.Get()
	device.List()
	device.Token()

	// Bucket manager

	bucket, err := admin.Bucket()
	if err != nil {
		panic(err)
	}

	bucket.Get()
	bucket.List()
}
