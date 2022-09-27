package definitions

import (
	"fmt"
	"net/http"

	"companies-api/internal/configs"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sarulabs/di"
)

const (
	HttpDef = "http"
)

func GetHttpDef() di.Def {
	return di.Def{
		Name:  HttpDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(CfgDef).(configs.Config)
			r := ctn.Get(RouterDef).(*mux.Router)

			return &http.Server{
				Handler: handlers.RecoveryHandler()(r),
				Addr:    fmt.Sprintf(":%d", cfg.App.Port),
			}, nil
		},
	}
}
