package gateway

import (
	"github.com/google/wire"
)

// ProviderGatewaySet is Gateway providers.
var ProviderGatewaySet = wire.NewSet(
	New,
)
