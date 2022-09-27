package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upCreateTableCompanies, downCreateTableCompanies)
}

func upCreateTableCompanies(tx *sql.Tx) error {
	if _, err := tx.Exec(`
CREATE TABLE IF NOT EXISTS companies
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    name    VARCHAR(255),
    code    VARCHAR(255),
    country    VARCHAR(255),
    website    VARCHAR(255),
    phone    VARCHAR(13),
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT name_UNIQUE UNIQUE (name),
    CONSTRAINT code_UNIQUE UNIQUE (code)
);
`); err != nil {
		return err
	}

	return nil
}

func downCreateTableCompanies(tx *sql.Tx) error {
	if _, err := tx.Exec(`DROP TABLE companies IF EXISTS;`); err != nil {
		return err
	}

	return nil
}
