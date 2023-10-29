package server

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/vanguard"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"

	"github.com/anhnmt/golang-gateway-boilerplate/pkg/config"
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

	// Serve the http server on the http listener.
	g.TryGo(func() error {
		addr := fmt.Sprintf(":%d", config.AppPort())
		log.Info().Msgf("Starting application http://localhost%s", addr)

		// create new http server
		srv := &http.Server{
			Addr:    addr,
			Handler: s.gateway,
			// ReadHeaderTimeout: 10 * time.Second,
			// ReadTimeout:       1 * time.Minute,
			// WriteTimeout:      1 * time.Minute,
			// MaxHeaderBytes:    8 * 1024, // 8KiB
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
