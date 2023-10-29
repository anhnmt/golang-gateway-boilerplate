// The build tag makes sure the stub is not built in the final build.
//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/google/wire"

	"github.com/anhnmt/golang-gateway-boilerplate/internal/server"
)

func InitServer(ctx context.Context) (*server.Server, error) {
	wire.Build(
		server.ProviderServerSet,
	)

	return &server.Server{}, nil
}
