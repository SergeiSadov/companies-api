package definitions

import (
	"companies-api/internal/configs"
	"companies-api/internal/db"
	dbinteractor "companies-api/internal/pkg/db"
	"companies-api/internal/pkg/error_adapters/sql_adapter"
	"companies-api/internal/pkg/errors"
	companyrepo "companies-api/internal/repositories/company/postgresql"

	"github.com/sarulabs/di"
)

const (
	ConnectionDef  = "connection"
	CompanyRepoDef = "company_repo"
)

func GetConnectionDef() di.Def {
	return di.Def{
		Name:  ConnectionDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(CfgDef).(configs.Config)

			conn, err := db.New(cfg.Database.Dialect, cfg.Database.PrepareDSN())
			if err != nil {
				return nil, err
			}

			return conn, nil
		},
	}
}

func GetCompanyRepoDef() di.Def {
	return di.Def{
		Name:  CompanyRepoDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			conn := ctn.Get(ConnectionDef).(db.IDatabase)
			gormDB, err := conn.GetConn()
			if err != nil {
				return nil, err
			}

			return companyrepo.NewCompanyRepository(
				dbinteractor.NewInteractor(gormDB),
				sql_adapter.New(errors.PreparedPostrgesErrorsMap),
			), nil
		},
	}
}
