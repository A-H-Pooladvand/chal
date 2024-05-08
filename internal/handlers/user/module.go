package user

import (
	"go.uber.org/fx"
	"theList/internal/repositories"
	"theList/internal/services"
)

var Module = fx.Module(
	"user",
	fx.Provide(
		NewHandler,
		services.NewUser,
		repositories.NewUser,
	),
)
