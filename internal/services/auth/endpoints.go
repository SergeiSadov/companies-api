package auth

import (
	"context"

	"companies-api/internal/entities/api"

	"github.com/go-kit/kit/endpoint"
)

func makeAuthEndpoint(svc IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*api.AuthRequest)

		token, err := svc.Auth(ctx, req)
		if err != nil {
			return
		}

		return &api.AuthResponse{
			Token: token,
		}, nil
	}
}
