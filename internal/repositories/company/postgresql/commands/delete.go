package commands

import (
	"context"

	"companies-api/internal/pkg/errors"

	"gorm.io/gorm"
)

const (
	deleteStmt = "DELETE FROM companies WHERE id = ? "
)

type Delete struct {
	Ctx context.Context

	ID int
}

func (cmd *Delete) Exec(client *gorm.DB) (err error) {
	stmt := client.WithContext(cmd.Ctx).Exec(deleteStmt, cmd.ID)
	if stmt.RowsAffected == 0 {
		return errors.ErrNotFound
	}

	return client.WithContext(cmd.Ctx).Exec(deleteStmt, cmd.ID).Error
}
