package company

import (
	"net/http"

	"companies-api/internal/pkg/error_adapters/http_adapter"
	"companies-api/internal/pkg/error_encoders"
	"companies-api/internal/pkg/middlewares"
	"companies-api/internal/services/company/validators"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func SetRoutes(svc IService,
	validator validators.IValidator,
	errorAdapter http_adapter.IErrorAdapter,
	r *mux.Router,
	logger *zap.Logger,
	authMW middlewares.IMiddleware,
	ipMW middlewares.IMiddleware,
) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(error_encoders.NewErrorEncoder(errorAdapter, logger)),
	}

	r.Handle("/companies", ipMW.HTTPMiddleware(authMW.HTTPMiddleware(kithttp.NewServer(
		makeListCompaniesEndpoint(svc),
		decodeListCompaniesRequest,
		encodeListCompaniesResponse,
		opts...,
	)))).Methods(http.MethodGet)

	r.Handle("/companies", ipMW.HTTPMiddleware(authMW.HTTPMiddleware(kithttp.NewServer(
		makeCreateCompanyEndpoint(svc, validator),
		decodeCreateCompanyRequest,
		encodeCreateCompaniesResponse,
		opts...,
	)))).Methods(http.MethodPost)

	r.Handle("/companies/{id}", ipMW.HTTPMiddleware(authMW.HTTPMiddleware(kithttp.NewServer(
		makeGetCompanyEndpoint(svc, validator),
		decodeGetCompanyRequest,
		encodeGetCompaniesResponse,
		opts...,
	)))).Methods(http.MethodGet)

	r.Handle("/companies/{id}", ipMW.HTTPMiddleware(authMW.HTTPMiddleware(kithttp.NewServer(
		makeUpdateCompanyEndpoint(svc, validator),
		decodeUpdateCompanyRequest,
		encodeUpdateCompaniesResponse,
		opts...,
	)))).Methods(http.MethodPut)

	r.Handle("/companies/{id}", ipMW.HTTPMiddleware(authMW.HTTPMiddleware(kithttp.NewServer(
		makeDeleteCompanyEndpoint(svc, validator),
		decodeDeleteCompanyRequest,
		encodeDeleteCompaniesResponse,
		opts...,
	)))).Methods(http.MethodDelete)
}
