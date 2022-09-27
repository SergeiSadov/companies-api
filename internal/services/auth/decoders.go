package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"companies-api/internal/entities/api"
	"companies-api/internal/pkg/errors"
)

func decodeAuthRequest(ctx context.Context, r *http.Request) (resp interface{}, err error) {
	var req *api.AuthRequest
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error: %v, %w", err, errors.ErrBadRequest)
	}
	if err = json.Unmarshal(body, &req); err != nil {
		return nil, fmt.Errorf("error: %v, %w", err, errors.ErrBadRequest)
	}

	return req, nil
}
