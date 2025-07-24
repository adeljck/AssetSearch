package quake

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adeljck/AssetSearch/config"
	"github.com/adeljck/AssetSearch/core"
	"github.com/go-resty/resty/v2"
	"net/http"
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
	client.client.SetBaseURL(config.QuakeApi)
	client.client.SetHeader("X-QuakeToken", Key)
	client.client.SetHeader("Content-Type", "application/json")
	return client
}
func (Q *Client) Name() string {
	return "Quake"
}
func (Q *Client) Check() (bool, error) {
	res, err := Q.client.R().Get(config.QuakeUserInfoPath)
	if err != nil {
		return false, err
	}
	if res.StatusCode() != http.StatusOK || res.Header().Get("Content-Type") != "application/json" {
		return false, errors.New("invalid Key Or Request Error")
	}
	results := new(QuakeUserInfo)
	err = json.Unmarshal(res.Body(), results)
	if err != nil {
		return false, err
	}
	if results.Code != 0 || results.Data.Baned {
		return false, errors.New(results.Message)
	}
	return true, nil
}
func (Q *Client) Query(query string, page int, pageSize int) ([]core.Result, error) {
	if query == "" {
		return nil, errors.New("Query Error")
	}
	if page <= 0 {
		return nil, errors.New("Page Error")
	}
	if pageSize < 10 {
		return nil, errors.New("PageSize Error")
	}
	//datas := map[string]interface{}{
	//	"query": query,
	//	"start": strconv.Itoa((page - 1) * pageSize),
	//	"size":  strconv.Itoa(pageSize),
	//}
	datas := struct {
		Query string `json:"query"`
		Start int    `json:"start"`
		Size  int    `json:"size"`
	}{
		Query: query,
		Start: (page - 1) * pageSize,
		Size:  pageSize,
	}
	res, err := Q.client.R().SetBody(datas).Post(config.QuakeSearchPath)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != http.StatusOK || res.Header().Get("Content-Type") != "application/json" {
		return nil, errors.New("invalid Key Or Request Error")
	}

	fmt.Println(string(res.Body()))
	return nil, nil
}
