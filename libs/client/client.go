package client

import "net/http"

type Client struct {
	c http.Client
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Get(url string) (*http.Response, error) {
	return c.c.Get(url)
}
