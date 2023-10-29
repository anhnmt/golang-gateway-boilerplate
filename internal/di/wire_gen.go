// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"context"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/gateway"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/server"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/service/userservice"
)

// Injectors from wire.go:

func InitServer(ctx context.Context) (*server.Server, error) {
	service := userservice.New()
	transcoder := gateway.New(service)
	serverServer := server.New(transcoder)
	return serverServer, nil
}
