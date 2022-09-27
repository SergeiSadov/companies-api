package commands

import (
	"context"

	"companies-api/internal/entities/repository"

	"gorm.io/gorm"
)

type Update struct {
	Ctx context.Context

	ID      int
	Name    string
	Code    string
	Country string
	Website string
	Phone   string

	Res *repository.Company
}

func (cmd *Update) Exec(client *gorm.DB) (err error) {
	return client.WithContext(cmd.Ctx).Table(companiesTable).Where("id = ?", cmd.ID).Updates(&repository.Company{
		Name:    cmd.Name,
		Code:    cmd.Code,
		Country: cmd.Country,
		Website: cmd.Website,
		Phone:   cmd.Phone,
	}).Find(&cmd.Res).Error
}
