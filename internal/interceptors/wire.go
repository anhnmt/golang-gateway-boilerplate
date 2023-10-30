package interceptors

import (
	"github.com/google/wire"
)

// ProviderInterceptorSet is Interceptor providers.
var ProviderInterceptorSet = wire.NewSet(
	New,
)
