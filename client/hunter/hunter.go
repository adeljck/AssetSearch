package hunter

import (
	"encoding/json"
	"fmt"
	"github.com/adeljck/AssetSearch/config"
	"github.com/adeljck/AssetSearch/core"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
)

type Client struct {
	Key    string
	client *resty.Client
}

func New(Key string) *Client {
	client := &Client{Key: Key, client: resty.New()}
	return client
}
func (H *Client) Name() string {
	return "Hunter"
}
func (H *Client) Query(query string, page int, pageSize int) (interface{}, error) {
	datas := map[string]string{
		"query":     core.ToBase64(query),
		"page":      strconv.Itoa(page),
		"page_size": strconv.Itoa(pageSize),
		"api-key":   H.Key,
	}
	res, err := H.client.R().SetPathParams(datas).Get(config.HunterSearchPath)
	if err != nil {
		return nil, err
	}
	if res.Header().Get("Content-Type") != "application/json" && res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("%s", "request failed with status code %d", res.StatusCode())
	}
	results := new(Result)
	err = json.Unmarshal(res.Body(), results)
	if err != nil {
		return nil, err
	}
	if results.Code != http.StatusOK {
		return nil, fmt.Errorf("%s", results.Message)
	}
	return results, nil
}
func (H *Client) Check() (bool, error) {
	_, err := H.Query("protocol=\"kkkkkkk\"", 1, 1)
	if err != nil {
		return false, err
	}
	return true, nil
}
