package ipapi

import (
	"errors"

	lru "github.com/hashicorp/golang-lru"
)

type CachedClient struct {
	client IClient
	cache  *lru.Cache
}

func NewCachedClient(
	client IClient,
	cache *lru.Cache,
) *CachedClient {
	return &CachedClient{
		client: client,
		cache:  cache,
	}
}

func (c CachedClient) GetCountryCode(ip string) (code string, err error) {
	ipVal, found := c.cache.Get(ip)

	var ok bool
	code, ok = ipVal.(string)
	if !found || !ok {
		code, err = c.client.GetCountryCode(ip)
		if err != nil {
			return
		}
		if code == "" {
			return code, errors.New("empty code in response")
		}
		c.cache.Add(ip, code)
		return
	}

	return code, nil
}
