package auth

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type Middleware struct {
	key    string
	logger *zap.Logger
}

func NewMiddleware(
	key string,
	logger *zap.Logger,
) *Middleware {
	return &Middleware{
		key:    key,
		logger: logger,
	}
}

func (m *Middleware) HTTPMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tkn, err := m.getAuth(r)
		if err != nil {
			m.logger.Error(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		r = r.WithContext(context.WithValue(r.Context(), Authorization, tkn))

		h.ServeHTTP(w, r)
	})
}

func (m *Middleware) getAuth(r *http.Request) (token *jwt.Token, err error) {
	token, err = jwt.Parse(r.Header.Get(Authorization), func(token *jwt.Token) (interface{}, error) {
		return []byte(m.key), nil
	})

	return
}
