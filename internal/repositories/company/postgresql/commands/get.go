package commands

import (
	"context"

	"companies-api/internal/entities/repository"

	"gorm.io/gorm"
)

type Get struct {
	Ctx context.Context

	ID int

	Res *repository.Company
}

func (cmd *Get) Exec(client *gorm.DB) (err error) {
	stmt := client.WithContext(cmd.Ctx).
		Select("id", "name", "code", "country", "website", "phone").
		Where("id = ?", cmd.ID).
		Table(companiesTable)

	return stmt.Scan(&cmd.Res).Error
}
