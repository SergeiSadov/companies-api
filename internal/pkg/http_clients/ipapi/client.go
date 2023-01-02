package ipapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	UrlTpl = "https://ipapi.co/%s/json/"
)

type Client struct {
	client     HttpClient
	errAdapter IErrorAdapter
}

func NewClient(client HttpClient, errAdapter IErrorAdapter) *Client {
	return &Client{
		client:     client,
		errAdapter: errAdapter,
	}
}

func (c *Client) GetCountryCode(ip string) (code string, err error) {
	resp, err := c.client.Get(fmt.Sprintf(UrlTpl, ip))
	if err != nil || resp.StatusCode != http.StatusOK {
		return code, c.errAdapter.AdaptStatusToErr(resp.StatusCode, err)
	}

	var countryResp Response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &countryResp); err != nil {
		return
	}

	return countryResp.CountryCode, nil
}
