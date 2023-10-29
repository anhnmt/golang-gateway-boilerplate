// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"context"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/server"
)

// Injectors from wire.go:

func InitServer(ctx context.Context) (*server.Server, error) {
	serverServer := server.New()
	return serverServer, nil
}