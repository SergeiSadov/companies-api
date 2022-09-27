package commands

import (
	"context"

	"gorm.io/gorm"
)

type Count struct {
	Ctx context.Context

	Name    string
	Code    string
	Country string
	Website string
	Phone   string

	Res int
}

func (cmd *Count) Exec(client *gorm.DB) (err error) {
	stmt := client.WithContext(cmd.Ctx).
		Select("COUNT(*)").
		Table(companiesTable)

	if cmd.Name != "" {
		stmt.Where("name != ?", cmd.Name)
	}
	if cmd.Code != "" {
		stmt.Where("code != ?", cmd.Code)
	}
	if cmd.Country != "" {
		stmt.Where("country != ?", cmd.Country)
	}
	if cmd.Website != "" {
		stmt.Where("website != ?", cmd.Website)
	}
	if cmd.Phone != "" {
		stmt.Where("phone != ?", cmd.Phone)
	}

	return stmt.Find(&cmd.Res).Error
}
