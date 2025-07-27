package fofa

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
	FoFAKey string
	client  *resty.Client
	fields  []string
}

func New(Key string) *Client {
	client := &Client{
		FoFAKey: Key,
		client:  resty.New(),
		fields:  []string{"ip", "port", "protocol", "host", "domain", "icp", "title"},
	}
	client.client.SetBaseURL(config.FoFaApi)
	return client
}
func (F *Client) Name() string {
	return "FoFa"
}
func (F *Client) Query(query string, page int, pageSize int, fields []string) (interface{}, error) {
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
		"qbase64": core.ToBase64(query),
		"page":    strconv.Itoa(page),
		"size":    strconv.Itoa(pageSize),
		"fields":  strings.Join(F.fields, ","),
		"key":     F.FoFAKey,
	}
	if fields != nil {
		datas["fields"] = strings.Join(fields, ",")
	}
	res, err := F.client.R().SetQueryParams(datas).Get(config.FoFaSearchPath)
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != http.StatusOK || !strings.Contains(res.Header().Get("Content-Type"), "application/json") {
		return nil, fmt.Errorf("invalid key Or Requests Error")
	}
	results := new(Results)
	err = json.Unmarshal(res.Body(), &results)
	if err != nil {
		return nil, err
	}
	if !results.Error {
		return nil, fmt.Errorf("invalid key Or Requests Error")
	}
	return results, nil
}

func (F *Client) Check() (bool, error) {
	res, err := F.client.R().SetPathParam("key", F.FoFAKey).Get(config.FoFaUserInfoPath)
	if err != nil {
		return false, err
	}
	if res.StatusCode() != http.StatusOK || res.Header().Get("Content-Type") != "application/json" {
		return false, errors.New("invalid Key Or Request Error")
	}
	results := new(FoFaAccountInfo)
	err = json.Unmarshal(res.Body(), results)
	if err != nil {
		return false, err
	}
	if results.Error {
		return false, errors.New(results.ErrMsg)
	}
	return true, nil
}
