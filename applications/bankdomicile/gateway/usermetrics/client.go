package usermetrics

import "github.com/aronrcihter/monorepo/client"

type Client struct {
	c client.Client
}

func NewClient(c client.Client) *Client {
	return &Client{c: c}
}
