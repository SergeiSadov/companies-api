package commands

import (
	"context"

	"companies-api/internal/entities/repository"

	"gorm.io/gorm"
)

const (
	createStmt = "INSERT INTO companies (name, code, country, website, phone) VALUES (?,?,?,?,?) RETURNING id, name, code, country, website, phone "
)

type Create struct {
	Ctx context.Context

	Name    string
	Code    string
	Country string
	Website string
	Phone   string

	Res *repository.Company
}

func (cmd *Create) Exec(client *gorm.DB) (err error) {
	return client.WithContext(cmd.Ctx).Raw(createStmt, cmd.Name, cmd.Code, cmd.Country, cmd.Website, cmd.Phone).
		Scan(&cmd.Res).Error
}
