package fofa

import "github.com/adeljck/AssetSearch/core"

type Client struct {
	FoFAKey string
}

func New(Key string) *Client {
	return &Client{
		FoFAKey: Key,
	}
}
func (F *Client) Name() string {
	return "FoFa"
}
func (F *Client) Query(query string, page int, pageSize int) ([]core.Result, error) {
	return nil, nil
}

func (F *Client) Check() bool {
	return false
}
