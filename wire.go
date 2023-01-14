//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"doc-api/api"
	"doc-api/env"

	"github.com/google/wire"
)

func InitializeServer() api.Server {
	wire.Build(
		// env
		env.NewPort,
		// server
		api.NewServer,
	)
	return api.Server{}
}
