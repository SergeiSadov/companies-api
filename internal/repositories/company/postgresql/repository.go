package postgresql

import (
	"context"

	"companies-api/internal/entities/repository"
	"companies-api/internal/pkg/db"
	"companies-api/internal/pkg/error_adapters/sql_adapter"
	"companies-api/internal/repositories/company/postgresql/commands"
)

type Company struct {
	interactor   db.IClient
	errorAdapter sql_adapter.IErrorAdapter
}

func NewCompanyRepository(
	interactor db.IClient,
	errorAdapter sql_adapter.IErrorAdapter,
) *Company {
	return &Company{
		interactor:   interactor,
		errorAdapter: errorAdapter,
	}
}

func (c *Company) Create(ctx context.Context, req *repository.Company) (response *repository.Company, err error) {
	cmd := &commands.Create{
		Ctx:     ctx,
		Name:    req.Name,
		Code:    req.Code,
		Country: req.Country,
		Website: req.Website,
		Phone:   req.Phone,
	}

	return cmd.Res, c.errorAdapter.AdaptSqlErr(c.interactor.Perform(cmd))
}

func (c *Company) Get(ctx context.Context, id int) (response *repository.Company, err error) {
	cmd := &commands.Get{
		Ctx: ctx,
		ID:  id,
	}

	return cmd.Res, c.errorAdapter.AdaptSqlErr(c.interactor.Perform(cmd))
}

func (c *Company) Count(ctx context.Context, req *repository.ListCompanyParams) (response int, err error) {
	cmd := &commands.Count{
		Ctx:     ctx,
		Name:    req.Name,
		Code:    req.Code,
		Country: req.Country,
		Website: req.Website,
		Phone:   req.Phone,
	}

	return cmd.Res, c.errorAdapter.AdaptSqlErr(c.interactor.Perform(cmd))
}

func (c *Company) List(ctx context.Context, req *repository.ListCompanyParams) (response []repository.Company, err error) {
	cmd := &commands.List{
		Ctx:     ctx,
		Name:    req.Name,
		Code:    req.Code,
		Country: req.Country,
		Website: req.Website,
		Phone:   req.Phone,
		Page:    req.Page,
		Size:    req.Size,
	}

	return cmd.Res, c.errorAdapter.AdaptSqlErr(c.interactor.Perform(cmd))
}

func (c *Company) Update(ctx context.Context, req *repository.Company) (response *repository.Company, err error) {
	cmd := &commands.Update{
		Ctx:     ctx,
		ID:      req.ID,
		Name:    req.Name,
		Code:    req.Code,
		Country: req.Country,
		Website: req.Website,
		Phone:   req.Phone,
	}

	return cmd.Res, c.errorAdapter.AdaptSqlErr(c.interactor.Perform(cmd))
}

func (c *Company) Delete(ctx context.Context, id int) (err error) {
	cmd := &commands.Delete{
		Ctx: ctx,
		ID:  id,
	}

	return c.errorAdapter.AdaptSqlErr(c.interactor.Perform(cmd))
}
