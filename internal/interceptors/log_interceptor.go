package interceptors

import (
	"context"

	"connectrpc.com/connect"
	"github.com/rs/zerolog/log"
)

type logInterceptor struct{}

func NewLogInterceptor() *logInterceptor {
	return &logInterceptor{}
}

func (i *logInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(
		ctx context.Context,
		req connect.AnyRequest,
	) (connect.AnyResponse, error) {
		res, err := next(ctx, req)

		logger := log.Info()
		if err != nil {
			logger = log.Error().Err(err)
		} else {
			logger.Interface("response", res.Any())
		}

		logger.
			Str("procedure", req.Spec().Procedure).
			Interface("request", req.Any()).
			Interface("header", req.Header()).
			Msg("Log unary interceptor")

		return res, err
	}
}

func (i *logInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	return func(
		ctx context.Context,
		spec connect.Spec,
	) connect.StreamingClientConn {
		conn := next(ctx, spec)
		return conn
	}
}

func (i *logInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return func(
		ctx context.Context,
		conn connect.StreamingHandlerConn,
	) error {
		return next(ctx, conn)
	}
}
