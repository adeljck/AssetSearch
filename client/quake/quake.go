package quake

import "github.com/adeljck/AssetSearch/core"

type Client struct {
	Key string
}

func New(Key string) *Client {
	return &Client{
		Key: Key,
	}
}
func (Q *Client) Name() string {
	return "Quake"
}
func (Q *Client) Check() bool {
	return false
}

func (Q *Client) Query(query string, page int, pageSize int) ([]core.Result, error) {
	return nil, nil
}
