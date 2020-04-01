package main

import (
	"fmt"
	"os"
	"time"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/api"
)

func main() {
	tagoAdmin := api.Tago(os.Getenv("ADMIN_TOKEN"))

	device := api.NewAdminDevice(tagoAdmin)

	ms := time.Now()

	res, err := device.GetDeviceToken()
	if err != nil {
		panic(err)
	}

	fmt.Println("Time:", time.Since(ms))

	for _, dev := range res {
		fmt.Println("Device:", dev.Name, dev.Token.Token)
	}

	fmt.Println()

	for _, dev := range res {
		ms := time.Now()

		tagoClient := api.Tago(dev.Token.Token)

		device := api.NewDevice(tagoClient)

		res, err := device.Insert(getPayload())
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Ingesting data on %s | %v => %v\n", dev.Name, res.Status, time.Since(ms))
	}

}

func getPayload() api.Data {
	return api.Data{
		Variable: "temperature",
		Value:    25.5,
		Time:     time.Now(),
	}
}
