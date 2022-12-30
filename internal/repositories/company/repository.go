//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

package company

import (
	"context"

	"companies-api/internal/entities/repository"
)

type IRepository interface {
	Create(ctx context.Context, req *repository.Company) (response *repository.Company, err error)
	Get(ctx context.Context, id string) (response *repository.Company, err error)
	Count(ctx context.Context, req *repository.ListCompanyParams) (response int, err error)
	List(ctx context.Context, req *repository.ListCompanyParams) (response []repository.Company, err error)
	Update(ctx context.Context, req *repository.Company) (response *repository.Company, err error)
	Delete(ctx context.Context, id string) (err error)
}
