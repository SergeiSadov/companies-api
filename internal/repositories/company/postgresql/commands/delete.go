package commands

import (
	"companies-api/internal/entities/event"
	"context"
	"encoding/json"
	"fmt"

	"companies-api/internal/pkg/errors"

	"gorm.io/gorm"
)

const (
	deleteStmt = "DELETE FROM companies WHERE id = ? "
)

type Delete struct {
	Ctx context.Context

	ID string
}

func (cmd *Delete) Exec(client *gorm.DB) (err error) {
	return client.Transaction(func(tx *gorm.DB) error {
		stmt := tx.WithContext(cmd.Ctx).Exec(deleteStmt, cmd.ID)
		if stmt.RowsAffected == 0 {
			return errors.ErrNotFound
		}

		if err = stmt.Error; err != nil {
			return fmt.Errorf("delete company: %w", err)
		}

		payload, err := json.Marshal(event.IDEvent{ID: cmd.ID})
		if err != nil {
			return fmt.Errorf("encode payload: %w", err)
		}

		if err = tx.WithContext(cmd.Ctx).Exec(createOutbox, cmd.ID, cmd.ID, typDelete, string(payload)).Error; err != nil {
			return fmt.Errorf("create outbox: %w", err)
		}

		return nil
	})
}
