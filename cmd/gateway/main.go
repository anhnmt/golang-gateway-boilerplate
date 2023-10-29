package main

import (
	"context"
	"flag"

	"github.com/rs/zerolog/log"

	"github.com/anhnmt/golang-gateway-boilerplate/internal/bootstrap/config"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/bootstrap/logger"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/di"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/server"
	"github.com/anhnmt/golang-gateway-boilerplate/utils"
)

var (
	env     string
	logFile string
)

func init() {
	flag.StringVar(&env, "env", "", "environment")
	flag.StringVar(&logFile, "log-file", "", "log file path, ex: logs/data.log")
	flag.Parse()

	// bootstrap
	logger.New(logFile)
	config.New(env)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv, err := di.InitServer(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("initial server failed")
	}

	go func(_ctx context.Context, _srv *server.Server) {
		if err = _srv.Start(_ctx); err != nil {
			log.Fatal().Err(err).
				Msg("start server failed")
		}
	}(ctx, srv)

	// wait for termination signal
	wait := utils.GracefulShutdown(ctx, utils.DefaultShutdownTimeout, map[string]utils.Operation{
		"server": func(newCtx context.Context) error {
			return srv.Close(newCtx)
		}})
	<-wait

	log.Info().Msg("graceful shutdown complete")
}
