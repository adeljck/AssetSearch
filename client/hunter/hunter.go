package hunter

import (
	"github.com/adeljck/AssetSearch/core"
)

type Client struct {
	Key string
}

func New(Key string) *Client {
	return &Client{Key: Key}
}
func (H *Client) Name() string {
	return "Hunter"
}

func (H *Client) Query(query string, page int, pageSize int) ([]core.Result, error) {
	return nil, nil
}
func (H *Client) Check() (bool, error) {
	return false, nil
}
