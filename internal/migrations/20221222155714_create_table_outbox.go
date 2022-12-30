package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableOutbox, downCreateTableOutbox)
}

func upCreateTableOutbox(tx *sql.Tx) error {
	if _, err := tx.Exec(`
CREATE TABLE IF NOT EXISTS outbox
(
    id               uuid         not null,
    aggregateid      varchar(255) not null,
    type             varchar(255) not null,
    payload          jsonb        not null
);

`); err != nil {
		return err
	}

	return nil
}

func downCreateTableOutbox(tx *sql.Tx) error {
	if _, err := tx.Exec(`
DROP TABLE IF EXISTS outbox;`); err != nil {
		return err
	}
	return nil
}
