package commands

import (
	"context"

	"companies-api/internal/entities/repository"

	"gorm.io/gorm"
)

const (
	companiesTable = "companies"
)

type List struct {
	Ctx context.Context

	Name    string
	Code    string
	Country string
	Website string
	Phone   string

	Page int
	Size int

	Res []repository.Company
}

func (cmd *List) Exec(client *gorm.DB) (err error) {
	stmt := client.WithContext(cmd.Ctx).
		Select("id", "name", "code", "country", "website", "phone").
		Table(companiesTable)

	if cmd.Name != "" {
		stmt.Where("name = ?", cmd.Name)
	}
	if cmd.Code != "" {
		stmt.Where("code = ?", cmd.Code)
	}
	if cmd.Country != "" {
		stmt.Where("country = ?", cmd.Country)
	}
	if cmd.Website != "" {
		stmt.Where("website = ?", cmd.Website)
	}
	if cmd.Phone != "" {
		stmt.Where("phone = ?", cmd.Phone)
	}
	if cmd.Page > 0 {
		stmt.Offset((cmd.Page - 1) * cmd.Size)
	}
	if cmd.Size > 0 {
		stmt.Limit(cmd.Size)
	}

	return stmt.Find(&cmd.Res).Error
}
