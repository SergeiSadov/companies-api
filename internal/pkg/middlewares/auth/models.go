package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

const (
	Authorization = "Authorization"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
