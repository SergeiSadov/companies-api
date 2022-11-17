package definitions

import (
	"net/http"
	"time"

	"companies-api/internal/pkg/http_clients/ipapi"

	lru "github.com/hashicorp/golang-lru"
	"github.com/sarulabs/di"
)

const (
	IPAPIClientDef       = "ipapi_client"
	IPAPICachedClientDef = "ipapi_cached_client"

	defaultCacheSize = 128
	defaultTimeout   = time.Second * 5
)

func GetIPAPIClientDef() di.Def {
	return di.Def{
		Name:  IPAPIClientDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return ipapi.NewClient(&http.Client{
				Timeout: defaultTimeout,
			},
				ipapi.NewErrAdapter(map[int]error{
					ipapi.StatusTooManyRequests: ipapi.ErrToManyRequests,
				}),
			), nil
		},
	}
}

func GetIPAPICachedClientDef() di.Def {
	return di.Def{
		Name:  IPAPICachedClientDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			client := ctn.Get(IPAPIClientDef).(ipapi.IClient)
			cache, err := lru.New(defaultCacheSize)
			if err != nil {
				return nil, err
			}

			return ipapi.NewCachedClient(client, cache), nil
		},
	}
}
