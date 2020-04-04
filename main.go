package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/util"
)

func main() {
	// util.StatsLoop()

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

	deviceAdm, err := admin.Device()
	if err != nil {
		panic(err)
	}

	device, err := deviceAdm.Get("5e83e40caf0d7a001b2b203e")
	if err != nil {
		panic(err)
	}

	fmt.Println(device.Name)

	q, _ := url.ParseQuery("")
	q.Add("filter[tags][0][key]", "gw")
	q.Add("filter[tags][0][value]", "gw-01")

	deviceList, err := deviceAdm.List(q.Encode())
	if err != nil {
		panic(err)
	}

	for _, d := range deviceList {
		fmt.Println(d.Name, d.Tags)
	}

	token, err := deviceAdm.Token("5e83e40caf0d7a001b2b203e")
	if err != nil {
		panic(err)
	}

	fmt.Println(token[0].Name, token[0].Token)

	// device.List()
	// device.Token()

	// Bucket manager

	// bucket, err := admin.Bucket()
	// if err != nil {
	// 	panic(err)
	// }

	// bucket.Get()
	// bucket.List()

	// Device

	deviceToken := os.Getenv("DEVICE_TOKEN")

	dev, err := cli.Device(deviceToken)
	if err != nil {
		panic(err)
	}

	payload := util.GeneratePayload()

	msgBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	res, err := dev.Data(payload)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data:", res)

	// Device pubsub

	p, err := dev.PubSub()
	if err != nil {
		panic(err)
	}

	payload = util.GeneratePayload()

	msgBytes, err = json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	err = p.Pub(msgBytes)
	if err != nil {
		panic(err)
	}

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

	go func() {
		for d := range debug {
			fmt.Println(d.Topic, d.Message.String())
		}
	}()

	time.Sleep(5 * time.Second)

	if err := p.UnsubscribeData(); err != nil {
		panic(err)
	}

	if err := p.UnsubscribeDebug(); err != nil {
		panic(err)
	}

	p.Close()

	fmt.Scanln()

	// p.Close()
}
