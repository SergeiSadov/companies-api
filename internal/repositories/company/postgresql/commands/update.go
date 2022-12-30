package commands

import (
	"companies-api/internal/entities/repository"
	"context"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type Update struct {
	Ctx context.Context

	ID      string
	Name    string
	Code    string
	Country string
	Website string
	Phone   string

	Res *repository.Company
}

func (cmd *Update) Exec(client *gorm.DB) (err error) {
	return client.Transaction(func(tx *gorm.DB) error {
		if err = tx.WithContext(cmd.Ctx).Table(companiesTable).Where("id = ?", cmd.ID).Updates(&repository.Company{
			Name:    cmd.Name,
			Code:    cmd.Code,
			Country: cmd.Country,
			Website: cmd.Website,
			Phone:   cmd.Phone,
		}).Find(&cmd.Res).Error; err != nil {
			return fmt.Errorf("update company: %w", err)

		}

		payload, err := json.Marshal(cmd.Res)
		if err != nil {
			return fmt.Errorf("encode payload: %w", err)
		}

		if err = tx.WithContext(cmd.Ctx).Exec(createOutbox, cmd.Res.ID, cmd.Res.ID, typeUpdate, string(payload)).Error; err != nil {
			return fmt.Errorf("create outbox: %w", err)
		}

		return nil
	})
}
