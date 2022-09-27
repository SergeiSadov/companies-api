package configs

import (
	"fmt"

	"companies-api/internal/pkg/constants"
)

type Database struct {
	Host         string `envconfig:"DB_HOST" required:"true"`
	Port         string `envconfig:"DB_PORT" required:"true"`
	Username     string `envconfig:"DB_USERNAME" required:"true"`
	Password     string `envconfig:"DB_PASSWORD" required:"true"`
	Name         string `envconfig:"DB_NAME" required:"true"`
	Dialect      string `envconfig:"DB_DIALECT" required:"true"`
	MigrationDir string `envconfig:"DB_MIGRATION_DIR" required:"true"`
}

func (c *Database) PrepareDSN() (dsn string) {
	return fmt.Sprintf(constants.DSNTemplate, c.Host, c.Username, c.Password, c.Name, c.Port)
}
