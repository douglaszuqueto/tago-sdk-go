package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/admin"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/device"
	"github.com/douglaszuqueto/tago-sdk-go/pkg/util"
)

var cli tago.Tago

func main() {
	util.StatsLoop()

	cli = tago.New()
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
	deviceManager(admin)

	// Bucket manager
	bucketManager(admin)

	// Device
	deviceTest()

	// go func() {
	// 	for {
	// 		select {
	// 		case d := <-data:
	// 			fmt.Println(d.Topic, d.Message.String())
	// 		case b := <-debug:
	// 			fmt.Println(b.Topic, b.Message.String())
	// 		}
	// 	}
	// }()

	fmt.Scanln()
}

//
// Admin Manager
//

func deviceManager(adm admin.Manager) {
	deviceAdm, err := adm.Device()
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
}

func bucketManager(adm admin.Manager) {
	// bucket, err := admin.Bucket()
	// if err != nil {
	// 	panic(err)
	// }

	// bucket.Get()
	// bucket.List()
}

//
// Device
//

func deviceTest() {
	deviceToken := os.Getenv("DEVICE_TOKEN")

	dev, err := cli.Device(deviceToken)
	if err != nil {
		panic(err)
	}

	deviceSendData(dev)
	devicePubSub(dev)
}

func deviceSendData(dev device.Device) {
	payload := util.GeneratePayload()

	res, err := dev.Data(payload)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data:", res)
}

func devicePubSub(dev device.Device) {
	p, err := dev.PubSub()
	if err != nil {
		panic(err)
	}

	payload := util.GeneratePayload()

	msgBytes, err := json.Marshal(payload)
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
	log.Println("Unsubscribing")

	if err := p.UnsubscribeData(); err != nil {
		panic(err)
	}

	if err := p.UnsubscribeDebug(); err != nil {
		panic(err)
	}

	time.Sleep(5 * time.Second)
	log.Println("Closing")

	p.Close()
}
