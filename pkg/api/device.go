package api

import (
	"net/http"
	"time"
)

// Device Device
type Device struct {
	client *Client
}

// Data Data
type Data struct {
	Variable string      `json:"variable"`
	Value    interface{} `json:"value"`
	Time     time.Time   `json:"time"`
}

// DeviceDataResponse DeviceDataResponse
type DeviceDataResponse struct {
	Message string `json:"message"`
	Result  string `json:"result"`
	Status  bool   `json:"status"`
}

// NewDevice NewDevice
func NewDevice(client *Client) *Device {
	return &Device{
		client: client,
	}
}

// Insert Insert
func (d *Device) Insert(data interface{}) (DeviceDataResponse, error) {
	var payload DeviceDataResponse

	d.client.Do("/data", http.MethodPost, data, &payload)

	return payload, nil
}
