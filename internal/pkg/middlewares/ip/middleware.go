package ip

import (
	"errors"
	"net"
	"net/http"

	"companies-api/internal/pkg/http_clients/ipapi"

	"go.uber.org/zap"
)

const (
	DefaultAllowedCountry = "CY"
)

type Middleware struct {
	client           ipapi.IClient
	logger           *zap.Logger
	allowedCountries []string
}

func NewMiddleware(
	client ipapi.IClient,
	logger *zap.Logger,
	allowedCountries ...string,
) *Middleware {
	return &Middleware{
		client:           client,
		logger:           logger,
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
	address := r.Header.Get("X-Forwarded-For")
	if address == "" {
		address = r.RemoteAddr
	}
	ip := net.ParseIP(address)
	if ip == nil {
		return code, errors.New("invalid ip address")
	}

	resp, err := m.client.GetCountryCode(ip.String())
	if err != nil {
		return
	}

	return resp, nil
}
