package postgresql

import (
	"companies-api/internal/entities/repository"
	"companies-api/internal/pkg/db"
)

type Auth struct {
	interactor db.IClient
}

func NewAuthRepository(
	interactor db.IClient,
) *Auth {
	return &Auth{
		interactor: interactor,
	}
}

func (a Auth) Auth(auth *repository.Auth) (tkn *repository.Token, err error) {
	return nil, nil
}
