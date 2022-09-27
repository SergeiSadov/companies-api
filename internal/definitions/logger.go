package definitions

import (
	"companies-api/internal/configs"

	"github.com/sarulabs/di"
	"go.uber.org/zap"
)

const (
	LoggerDef = "logger"
)

func GetLoggerDef() di.Def {
	return di.Def{
		Name:  LoggerDef,
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			cfg := ctn.Get(CfgDef).(configs.Config)

			var loggerCfg zap.Config

			switch cfg.App.Environment {
			case "prod":
				loggerCfg = zap.NewProductionConfig()
			default:
				loggerCfg = zap.NewDevelopmentConfig()
			}
			loggerCfg.DisableStacktrace = true

			logger, err := loggerCfg.Build()
			if err != nil {
				return nil, err
			}
			zap.ReplaceGlobals(logger)

			return logger, nil
		},
	}
}
