// The build tag makes sure the stub is not built in the final build.
//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/google/wire"

	"github.com/anhnmt/golang-gateway-boilerplate/internal/gateway"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/interceptors"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/server"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/service"
	"github.com/anhnmt/golang-gateway-boilerplate/pkg/redis"
)

func InitServer(ctx context.Context) (*server.Server, error) {
	wire.Build(
		redis.ProviderRedisSet,
		service.ProviderServiceSet,
		interceptors.ProviderInterceptorSet,
		gateway.ProviderGatewaySet,
		server.ProviderServerSet,
	)

	return &server.Server{}, nil
}
