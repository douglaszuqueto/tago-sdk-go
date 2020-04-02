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

	adminToken := os.Getenv("ADMIN_TOKEN")

	admin, err := cli.Admin(adminToken)
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

	// Device

	deviceToken := os.Getenv("DEVICE_TOKEN")

	dev, err := cli.Device(deviceToken)
	if err != nil {
		panic(err)
	}

	// dev.Data()

	// Device pubsub

	p, err := dev.PubSub()
	if err != nil {
		panic(err)
	}

	// payload := util.GeneratePayload()

	// msgBytes, err := json.Marshal(payload)
	// if err != nil {
	// 	panic(err)
	// }

	// p.Pub(msgBytes)

	data, err := p.Sub()
	if err != nil {
		panic(err)
	}

	go func() {
		for d := range data {
			fmt.Println(d.Topic, d.Message.String())
		}
	}()

	debug, err := p.Debug()
	if err != nil {
		panic(err)
	}

	for d := range debug {
		fmt.Println(d.Topic, d.Message.String())
	}

	fmt.Scanln()
}
