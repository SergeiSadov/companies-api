package definitions

import (
	"fmt"

	"github.com/sarulabs/di"
)

func Build() (container di.Container, err error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, fmt.Errorf("can't create container builder: %s", err)
	}

	err = builder.Add([]di.Def{
		GetConfigDef(),
		GetLoggerDef(),
		GetIPAPIClientDef(),
		GetIPAPICachedClientDef(),
		GetConnectionDef(),
		GetCompanyRepoDef(),
		GetCompanyServiceDef(),
		GetAuthServiceDef(),
		GetRouterDef(),
		GetHttpDef(),
	}...)
	if err != nil {
		return
	}

	container = builder.Build()

	return container, nil
}
