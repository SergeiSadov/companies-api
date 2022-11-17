package definitions

import (
	"net/http"

	"companies-api/internal/configs"
	"companies-api/internal/pkg/error_adapters/http_adapter"
	"companies-api/internal/pkg/http_clients/ipapi"
	"companies-api/internal/pkg/middlewares/auth"
	"companies-api/internal/pkg/middlewares/ip"
	authsvc "companies-api/internal/services/auth"
	companysvc "companies-api/internal/services/company"
	"companies-api/internal/services/company/validators"

	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
)

const (
	RouterDef = "router"
)

func GetRouterDef() di.Def {
	return di.Def{
		Name:  RouterDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			r := mux.NewRouter()
			companyService := ctn.Get(CompanyServiceDef).(companysvc.IService)
			authService := ctn.Get(AuthServiceDef).(authsvc.IService)
			logger := ctn.Get(LoggerDef).(*zap.Logger)
			config := ctn.Get(CfgDef).(configs.Config)

			httpErrAdapter := http_adapter.New(http.StatusInternalServerError, http_adapter.AdaptBadRequestError)
			cachedClient := ctn.Get(IPAPICachedClientDef).(ipapi.IClient)
			authMW := auth.NewMiddleware(config.App.JWTSecret, logger)
			ipMW := ip.NewMiddleware(cachedClient, logger, ip.NewErrorAdapter(ip.PreparedErrorMapping), config.Countries.AlloweCountryCode)

			companysvc.SetRoutes(companyService, validators.PreparedValidators, httpErrAdapter, r, logger, authMW, ipMW)
			authsvc.SetRoutes(authService, httpErrAdapter, r, logger)

			return r, nil
		},
	}
}
