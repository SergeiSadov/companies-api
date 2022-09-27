package definitions

import (
	"companies-api/internal/configs"

	"github.com/sarulabs/di"
)

const (
	CfgDef = "config"
)

func GetConfigDef() di.Def {
	return di.Def{
		Name:  CfgDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return configs.Setup()
		},
	}
}
