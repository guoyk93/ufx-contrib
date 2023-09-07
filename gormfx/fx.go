package gormfx

import (
	"go.uber.org/fx"
)

var ModuleMySQL = fx.Module(
	"ufx_gormfx_mysql",
	fx.Provide(
		ParamsFromConf,
		MySQLParamsFromConf,
		NewMySQLDialector,
		NewConfig,
		NewClient,
	),
	fx.Invoke(AddCheckerForClient),
)
