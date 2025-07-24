package quake

import (
	"github.com/adeljck/AssetSearch/core"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Key    string
	client *resty.Client
}

func New(Key string) *Client {
	client := &Client{
		Key:    Key,
		client: resty.New(),
	}
	return client
}
func (Q *Client) Name() string {
	return "Quake"
}
func (Q *Client) Check() (bool, error) {
	return false, nil
}

func (Q *Client) Query(query string, page int, pageSize int) ([]core.Result, error) {
	return nil, nil
}
