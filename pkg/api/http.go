package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

// Client Client
type Client struct {
	token string
	http  *http.Client
}

// Tago Tago
func Tago(token string) *Client {
	client := &http.Client{}

	return &Client{
		token: token,
		http:  client,
	}
}

// Do Do
func (c *Client) Do(path string, method string, data interface{}, payload interface{}) error {
	var p []byte

	switch data.(type) {
	case []byte:
		p = data.([]byte)
	case Data:
		var err error

		p, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	client := &http.Client{}

	uri := &url.URL{
		Host:   "api.tago.io",
		Path:   path,
		Scheme: "https",
	}

	if path == "/device" {
		tagoGw := os.Getenv("TAGO_GW")
		uri.RawQuery = "filter%5Btags%5D%5B0%5D%5Bkey%5D=gateway&filter%5Btags%5D%5B0%5D%5Bvalue%5D=" + tagoGw
	}

	req, err := http.NewRequest(method, uri.String(), bytes.NewBuffer(p))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", c.token)
	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&payload)
	if err != nil {
		return err
	}

	return nil
}
