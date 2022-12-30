package company

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"companies-api/internal/entities/api"
	"companies-api/internal/pkg/errors"

	"github.com/gorilla/mux"
)

func decodeListCompaniesRequest(ctx context.Context, r *http.Request) (resp interface{}, err error) {
	var page int
	pageParam := r.FormValue("page")
	if pageParam != "" {
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			return nil, fmt.Errorf("error: %v, %w", err, errors.ErrBadRequest)
		}
	}
	if page <= 0 {
		page = 1
	}

	var size int
	pageSize := r.FormValue("size")
	if pageSize != "" {
		size, err = strconv.Atoi(pageSize)
		if err != nil {
			return nil, fmt.Errorf("error: %v, %w", err, errors.ErrBadRequest)
		}
	}
	if size <= 0 {
		size = 5
	}

	return &api.ListCompanyRequest{
		Name:    r.FormValue("name"),
		Code:    r.FormValue("code"),
		Country: r.FormValue("country"),
		Website: r.FormValue("website"),
		Phone:   r.FormValue("phone"),
		Page:    page,
		Size:    size,
	}, nil
}

func decodeCreateCompanyRequest(ctx context.Context, r *http.Request) (resp interface{}, err error) {
	var req api.CreateCompanyRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error: %v, %w", err, errors.ErrBadRequest)
	}
	if err = json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("error: %v, %w", err, errors.ErrBadRequest)
	}

	return &req, nil
}

func decodeGetCompanyRequest(ctx context.Context, r *http.Request) (resp interface{}, err error) {
	return &api.GetCompanyRequest{ID: mux.Vars(r)["id"]}, nil
}

func decodeUpdateCompanyRequest(ctx context.Context, r *http.Request) (resp interface{}, err error) {
	var req api.UpdateCompanyRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error: %v, %w", err, errors.ErrBadRequest)
	}
	if err = json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("error: %v, %w", err, errors.ErrBadRequest)
	}
	req.ID = mux.Vars(r)["id"]

	return &req, nil
}

func decodeDeleteCompanyRequest(ctx context.Context, r *http.Request) (resp interface{}, err error) {
	return &api.DeleteCompanyRequest{ID: mux.Vars(r)["id"]}, nil
}
