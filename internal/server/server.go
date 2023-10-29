package server

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	return g.Wait()
}

func (s *Server) Close(ctx context.Context) error {
	g, _ := errgroup.WithContext(ctx)

	return g.Wait()
}
