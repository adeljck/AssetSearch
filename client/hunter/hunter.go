package hunter

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adeljck/AssetSearch/config"
	"github.com/adeljck/AssetSearch/core"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
	"strings"
)

type Client struct {
	Key    string
	client *resty.Client
}

func New(Key string) *Client {
	client := &Client{Key: Key, client: resty.New()}
	client.client.SetBaseURL(config.HunterApi)
	return client
}
func (H *Client) Name() string {
	return "Hunter"
}
func (H *Client) Query(query string, page int, pageSize int) (interface{}, error) {
	if query == "" {
		return nil, errors.New("Query Error")
	}
	if page <= 0 {
		return nil, errors.New("Page Error")
	}
	if pageSize < 10 {
		return nil, errors.New("PageSize Error")
	}
	datas := map[string]string{
		"api-key":   H.Key,
		"query":     core.ToBase64(query),
		"page":      strconv.Itoa(page),
		"page_size": strconv.Itoa(pageSize),
	}
	res, err := H.client.R().SetQueryParams(datas).Get(config.HunterSearchPath)
	if err != nil {
		return nil, err
	}
	if !strings.Contains(res.Header().Get("Content-Type"), "application/json") || res.StatusCode() != http.StatusOK {
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
	_, err := H.Query("protocol=\"kkkkkkk\"", 1, 10)
	if err != nil {
		return false, err
	}
	return true, nil
}
