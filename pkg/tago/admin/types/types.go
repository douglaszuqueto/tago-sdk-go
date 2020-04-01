package types

import "time"

// Response Response
type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// DeviceDataResponse DeviceDataResponse
type DeviceDataResponse struct {
	Message string `json:"message"`
	Result  string `json:"result"`
	Status  bool   `json:"status"`
}

// Data Data
type Data struct {
	Variable string      `json:"variable"`
	Value    interface{} `json:"value"`
	Time     time.Time   `json:"time"`
}
