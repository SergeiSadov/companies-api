package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAlterTableCompaniesAddIndex, downAlterTableCompaniesAddIndex)
}

func upAlterTableCompaniesAddIndex(tx *sql.Tx) error {
	if _, err := tx.Exec(`CREATE INDEX companies_common_query_idx ON companies (country, website, phone);`); err != nil {
		return err
	}

	return nil
}

func downAlterTableCompaniesAddIndex(tx *sql.Tx) error {
	if _, err := tx.Exec(`DROP INDEX companies_common_query_idx;`); err != nil {
		return err
	}

	return nil
}
