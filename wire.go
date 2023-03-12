//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"doc-api/api"
	"doc-api/api/controller"
	"doc-api/api/db"
	"doc-api/api/entity"
	"doc-api/api/gateway"
	"doc-api/api/usecase"
	"doc-api/api/web"
	"doc-api/env"

	"github.com/google/wire"
)

func InitializeServer() api.Server {
	wire.Build(
		// env
		env.NewPort,
		env.NewJwtSecretKey,
		env.NewHashSalt,
		env.NewHashStretch,
		env.NewDBHost,
		env.NewDBPort,
		env.NewDBName,
		env.NewDBPass,
		env.NewDBUser,
		// server
		api.NewServer,
		// web
		web.NewRouter,
		web.NewCustomLogger,
		// controller
		controller.NewAuthUserMiddleware,
		controller.NewHelloHandler,
		controller.NewLoginHandler,
		controller.NewDocHandler,
		controller.NewUserHandler,
		// gateway
		gateway.NewLoginRepository,
		gateway.NewAuthRepository,
		gateway.NewDocRepository,
		gateway.NewUserRepository,
		// usecase
		usecase.NewLoginService,
		usecase.NewAuthUserService,
		usecase.NewDocService,
		usecase.NewUserService,
		// entity
		entity.NewHashing,
		entity.NewJwtToken,
		// db
		db.NewDBConnection,
	)
	return api.Server{}
}
