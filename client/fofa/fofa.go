package fofa

import (
	"github.com/adeljck/AssetSearch/config"
	"github.com/adeljck/AssetSearch/core"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	FoFAKey string
	client  *resty.Client
}

func New(Key string) *Client {
	client := &Client{
		FoFAKey: Key,
		client:  resty.New(),
	}
	client.client.SetBaseURL(config.FoFaApi)
	return client
}
func (F *Client) Name() string {
	return "FoFa"
}
func (F *Client) Query(query string, page int, pageSize int) ([]core.Result, error) {
	return nil, nil
}

func (F *Client) Check() (bool, error) {
	return false, nil
}
