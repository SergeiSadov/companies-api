package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"

	"companies-api/internal/entities/repository"

	"gorm.io/gorm"
)

const (
	createStmt = "INSERT INTO companies (id, name, code, country, website, phone) VALUES (?,?,?,?,?,?) RETURNING id, name, code, country, website, phone "
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
	return client.Transaction(func(tx *gorm.DB) error {
		id, err := uuid.NewUUID()
		if err != nil {
			return fmt.Errorf("generate uuid: %w", err)
		}

		if err = tx.WithContext(cmd.Ctx).Raw(createStmt, id, cmd.Name, cmd.Code, cmd.Country, cmd.Website, cmd.Phone).
			Scan(&cmd.Res).Error; err != nil {
			return fmt.Errorf("create company: %w", err)

		}

		payload, err := json.Marshal(cmd.Res)
		if err != nil {
			return fmt.Errorf("encode payload: %w", err)
		}

		if err = tx.WithContext(cmd.Ctx).Exec(createOutbox, id, id, typeCreate, string(payload)).Error; err != nil {
			return fmt.Errorf("create outbox: %w", err)
		}

		return nil
	})
}
