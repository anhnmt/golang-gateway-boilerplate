package server

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"connectrpc.com/vanguard"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"

	"github.com/anhnmt/golang-gateway-boilerplate/internal/bootstrap/config"
)

type Server struct {
	gateway *vanguard.Transcoder
}

func New(
	gateway *vanguard.Transcoder,
) *Server {
	return &Server{
		gateway: gateway,
	}
}

func (s *Server) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	if config.PprofEnabled() {
		g.TryGo(func() error {
			addr := fmt.Sprintf(":%d", config.PprofPort())
			log.Info().Msgf("starting pprof on http://localhost%s", addr)

			return http.ListenAndServe(addr, nil)
		})
	}

	// Serve the http server on the http listener.
	g.TryGo(func() error {
		addr := fmt.Sprintf(":%d", config.AppPort())
		log.Info().
			Str("app_name", config.AppName()).
			Msgf("starting server on http://localhost%s", addr)

		// create new http server
		srv := &http.Server{
			Addr: addr,
			// We use the h2c package in order to support HTTP/2 without TLS,
			// so we can handle gRPC requests, which requires HTTP/2, in
			// addition to Connect and gRPC-Web (which work with HTTP 1.1).
			Handler: h2c.NewHandler(
				s.gateway,
				&http2.Server{},
			),
		}

		// run the server
		return srv.ListenAndServe()
	})

	return g.Wait()
}

func (s *Server) Close(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	return g.Wait()
}
