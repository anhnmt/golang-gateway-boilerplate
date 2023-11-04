package db

import (
	"github.com/google/wire"
)

// ProviderDatabaseSet is database providers.
var ProviderDatabaseSet = wire.NewSet(
	NewDatabase,
)
