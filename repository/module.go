package repository

import(
	"go.uber.org/fx"
)

var DBClientModule fx.Options(
	fx.Provide(func() (*Config, error){
		
	})
)