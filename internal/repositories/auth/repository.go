//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package auth

import (
	"companies-api/internal/entities/repository"
)

type IRepository interface {
	Auth(auth *repository.Auth) (tkn *repository.Token, err error)
}
