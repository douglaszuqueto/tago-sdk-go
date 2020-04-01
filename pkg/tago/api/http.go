package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
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

func (c *Client) do(method, path string, data interface{}, payload interface{}) error {
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

// Get Get
func (c *Client) Get(path string, response interface{}) error {
	return c.do(http.MethodGet, path, nil, response)
}

// Post Post
func (c *Client) Post(path string, data interface{}, response interface{}) error {
	return c.do(http.MethodPost, path, data, response)
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
