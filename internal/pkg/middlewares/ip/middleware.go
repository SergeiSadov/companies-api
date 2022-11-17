package ip

import (
	"net"
	"net/http"

	"companies-api/internal/pkg/http_clients/ipapi"

	"go.uber.org/zap"
)

const (
	DefaultAllowedCountry = "CY"
	XForwardedForHeader   = "X-Forwarded-For"
)

type Middleware struct {
	client           ipapi.IClient
	logger           *zap.Logger
	errorAdapter     IErrorAdapter
	allowedCountries []string
}

func NewMiddleware(
	client ipapi.IClient,
	logger *zap.Logger,
	errorAdapter IErrorAdapter,
	allowedCountries ...string,
) *Middleware {
	return &Middleware{
		client:           client,
		logger:           logger,
		errorAdapter:     errorAdapter,
		allowedCountries: allowedCountries,
	}
}

func (m *Middleware) HTTPMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code, err := m.getCountryCode(r)
		if err != nil {
			m.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//bypass empty code as we use free edition of ipapi
		if code == "" {
			h.ServeHTTP(w, r)
			return
		}

		for i := range m.allowedCountries {
			if m.allowedCountries[i] == code {
				h.ServeHTTP(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusForbidden)
	})
}

func (m *Middleware) getCountryCode(r *http.Request) (code string, err error) {
	address := r.Header.Get(XForwardedForHeader)
	if address == "" {
		address = r.RemoteAddr
	}
	ip := net.ParseIP(address)
	if ip == nil {
		return code, ErrInvalidIpAddress
	}

	resp, err := m.client.GetCountryCode(ip.String())
	if err != nil {
		return code, m.errorAdapter.AdaptError(err)
	}

	return resp, nil
}
