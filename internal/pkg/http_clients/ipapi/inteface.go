//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package ipapi

import (
	"net/http"
)

type IClient interface {
	GetCountryCode(ip string) (code string, err error)
}

type HttpClient interface {
	Get(url string) (resp *http.Response, err error)
}
