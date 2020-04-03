package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/douglaszuqueto/tago-sdk-go/pkg/tago/admin/types"
)

// Client Client
type Client struct {
	token string
	http  *http.Client
}

// NewClient NewClient
func NewClient(token string) *Client {
	client := &http.Client{}

	return &Client{
		token: token,
		http:  client,
	}
}

func (c *Client) do(method, path string, filter string, data interface{}, payload interface{}) error {
	// p, err := json.Marshal(data)
	// if err != nil {
	// 	return err
	// }

	var p []byte

	switch data.(type) {
	case []byte:
		p = data.([]byte)
	case types.Data:
		var err error

		p, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	client := &http.Client{}

	uri := &url.URL{
		Host:   "api.tago.io",
		Scheme: "https",
		Path:   path,
	}

	if len(filter) > 0 {
		uri.RawQuery = filter
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

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		var response types.Response

		err = dec.Decode(&response)
		if err != nil {
			return err
		}

		// fmt.Println("Error:", resp.StatusCode, response.Message, response.Status)

		return errors.New(response.Message)
	}

	err = dec.Decode(&payload)
	if err != nil {
		return err
	}

	return nil
}

// Get Get
func (c *Client) Get(path string, filter string, response interface{}) error {
	return c.do(http.MethodGet, path, filter, nil, response)
}

// Post Post
func (c *Client) Post(path string, data interface{}, response interface{}) error {
	return c.do(http.MethodPost, path, "", data, response)
}

// Do Do
func (c *Client) Do(path string, data interface{}, payload interface{}) error {
	p, err := json.Marshal(data)
	if err != nil {
		return err
	}

	client := &http.Client{}

	uri := &url.URL{
		Host:   "api.tago.io",
		Scheme: "https",
		Path:   path,
	}

	req, err := http.NewRequest(http.MethodGet, uri.String(), bytes.NewBuffer(p))
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
