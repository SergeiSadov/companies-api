package auth

import (
	"net/http"

	"companies-api/internal/pkg/error_adapters/http_adapter"
	"companies-api/internal/pkg/error_encoders"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func SetRoutes(svc IService, errorAdapter http_adapter.IErrorAdapter, r *mux.Router, logger *zap.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(error_encoders.NewErrorEncoder(errorAdapter, logger)),
	}

	r.Handle("/auth", kithttp.NewServer(
		makeAuthEndpoint(svc),
		decodeAuthRequest,
		encodeAuthResponse,
		opts...,
	)).Methods(http.MethodPost)
}
