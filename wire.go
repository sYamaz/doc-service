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
		// server
		api.NewServer,
		// web
		web.NewRouter,
		// controller
		controller.NewAuthUserMiddleware,
		controller.NewHelloHandler,
		controller.NewLoginHandler,
		controller.NewSignupHandler,
		// gateway
		gateway.NewUserRepository,
		gateway.NewSignupRepository,
		gateway.NewAuthRepository,
		// usecase
		usecase.NewLoginService,
		usecase.NewSignupService,
		usecase.NewAuthUserService,
		// entity
		entity.NewHashing,
		entity.NewJwtToken,
		// db
		db.NewDBConnection,
	)
	return api.Server{}
}
