package ipapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	UrlTpl = "https://ipapi.co/%s/json/"
)

type Client struct {
	client HttpClient
}

func NewClient(client HttpClient) *Client {
	return &Client{
		client: client,
	}
}

func (c Client) GetCountryCode(ip string) (code string, err error) {
	resp, err := c.client.Get(fmt.Sprintf(UrlTpl, ip))
	if err != nil {
		return
	}
	var countryResp Response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if err = json.Unmarshal(body, &countryResp); err != nil {
		return
	}

	return countryResp.CountryCode, nil
}
