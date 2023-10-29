package service

import (
	"github.com/google/wire"

	"github.com/anhnmt/golang-gateway-boilerplate/internal/service/userservice"
)

// ProviderServiceSet is Service providers.
var ProviderServiceSet = wire.NewSet(
	userservice.New,
)
