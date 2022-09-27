package auth

import (
	"context"
	"encoding/json"
	"net/http"

	"companies-api/internal/entities/api"
)

func encodeAuthResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	resp := response.(*api.AuthResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}
