package api

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// AdminDevice AdminDevice
type AdminDevice struct {
	client *Client
}

// NewAdminDevice NewAdminDevice
func NewAdminDevice(client *Client) *AdminDevice {
	return &AdminDevice{
		client: client,
	}
}

// AdminDeviceData AdminDeviceData
type AdminDeviceData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Tags []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"tags"`
}

// AdminDeviceResponse AdminDeviceResponse
type AdminDeviceResponse struct {
	Result []DeviceData `json:"result"`
	Status bool         `json:"status"`
}

// DeviceData DeviceData
type DeviceData struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Token struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	} `json:"token"`
}

// TokenData TokenData
type TokenData struct {
	Result []struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	} `json:"result"`
	Status bool `json:"status"`
}

// List List
func (d *AdminDevice) List() ([]DeviceData, error) {
	var payload AdminDeviceResponse

	err := d.client.Do("/device", http.MethodGet, nil, &payload)
	if err != nil {
		return nil, err
	}

	return payload.Result, nil
}

// DeviceToken DeviceToken
func (d *AdminDevice) DeviceToken(deviceID string) (TokenData, error) {
	var payload TokenData

	err := d.client.Do("/device/token/"+deviceID, http.MethodGet, nil, &payload)
	if err != nil {
		log.Println(err)
		return payload, err
	}

	return payload, nil
}

// GetDeviceToken GetDeviceToken
func (d *AdminDevice) GetDeviceToken() ([]DeviceData, error) {
	devices, err := d.List()
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}

	for i, dev := range devices {
		wg.Add(1)
		go func(deviceID string, i int) {
			defer wg.Done()

			token, err := d.DeviceToken(deviceID)
			if err != nil {
				fmt.Println("DeviceToken", err)
				return
			}

			devices[i].Token = token.Result[0]
		}(dev.ID, i)
	}

	wg.Wait()

	return devices, nil
}
