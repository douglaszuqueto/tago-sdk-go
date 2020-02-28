package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/api"
)

var (
	accessToken = flag.String("token", "", "Account token or Device token")
	payloadPath = flag.String("payload", "", "JSON payoad path")
)

func main() {
	flag.Parse()

	if len(*accessToken) == 0 {
		panic("Account token or Device token is required!")
	}

	if len(*payloadPath) == 0 {
		panic("Payload path is required!")
	}

	file, err := os.Open(*payloadPath)
	if err != nil {
		panic("Invalid file:" + err.Error())
	}

	by, err := ioutil.ReadAll(file)
	if err != nil {
		panic("Invalid file")
	}

	var p api.Data

	if err := json.Unmarshal(by, &p); err != nil {
		panic("JSON error:" + err.Error())
	}

	tagoClient := api.Tago(*accessToken)

	device := api.NewDevice(tagoClient)

	// payload := api.Data{
	// 	Variable: "temperature",
	// 	Value:    25.5,
	// 	Time:     time.Now(),
	// }

	dev, err := device.Insert(by)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Status:", dev.Status)
	fmt.Println("Result:", dev.Result)
	fmt.Println("Message:", dev.Message)
}
