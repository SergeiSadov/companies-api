package auth

import (
	"context"
	"time"

	"companies-api/internal/entities/api"
	"companies-api/internal/pkg/middlewares/auth"

	"github.com/golang-jwt/jwt/v4"
)

type IService interface {
	Auth(ctx context.Context, req *api.AuthRequest) (response string, err error)
}

type service struct {
	jwtKey string
}

func NewService(jwtKey string) *service {
	return &service{
		jwtKey: jwtKey,
	}
}

func (s service) Auth(ctx context.Context, req *api.AuthRequest) (response string, err error) {
	claims := &auth.Claims{
		Username: req.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtKey))
	if err != nil {
		return
	}

	return tokenString, nil
}
