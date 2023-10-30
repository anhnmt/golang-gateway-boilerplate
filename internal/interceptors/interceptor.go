package interceptors

import (
	"connectrpc.com/connect"

	"github.com/anhnmt/golang-gateway-boilerplate/internal/bootstrap/config"
)

func New() connect.Option {
	interceptors := make([]connect.Interceptor, 0)

	if config.LogPayload() {
		interceptors = append(interceptors, NewLogInterceptor())
	}

	return connect.WithInterceptors(interceptors...)
}
