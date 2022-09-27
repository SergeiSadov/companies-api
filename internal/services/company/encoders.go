package company

import (
	"context"
	"encoding/json"
	"net/http"

	"companies-api/internal/entities/api"
)

func encodeListCompaniesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	resp := response.(*api.ListCompanyResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func encodeCreateCompaniesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	resp := response.(*api.CreateCompanyResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func encodeGetCompaniesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	resp := response.(*api.GetCompanyResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func encodeUpdateCompaniesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	resp := response.(*api.UpdateCompanyResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func encodeDeleteCompaniesResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}
