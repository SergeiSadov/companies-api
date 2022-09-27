package definitions

import (
	"companies-api/internal/configs"
	"companies-api/internal/pkg/kafka_writer"
	"companies-api/internal/repositories/company"
	authsvc "companies-api/internal/services/auth"
	companysvc "companies-api/internal/services/company"
	"companies-api/internal/services/company/adapter"

	"github.com/sarulabs/di"
)

const (
	CompanyServiceDef = "company_service"
	AuthServiceDef    = "auth_service"
)

func GetCompanyServiceDef() di.Def {
	return di.Def{
		Name:  CompanyServiceDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			companyRepo := ctn.Get(CompanyRepoDef).(company.IRepository)
			createWriter := ctn.Get(CreateCompanyKafkaWriterDef).(kafka_writer.IKafkaWriter)
			updateWriter := ctn.Get(UpdateCompanyKafkaWriterDef).(kafka_writer.IKafkaWriter)
			deleteWriter := ctn.Get(DeleteCompanyKafkaWriterDef).(kafka_writer.IKafkaWriter)

			return companysvc.NewService(
				companyRepo,
				createWriter,
				updateWriter,
				deleteWriter,
				adapter.NewAdapter(),
			), nil
		},
	}
}

func GetAuthServiceDef() di.Def {
	return di.Def{
		Name:  AuthServiceDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			config := ctn.Get(CfgDef).(configs.Config)
			return authsvc.NewService(
				config.App.JWTSecret,
			), nil
		},
	}
}
