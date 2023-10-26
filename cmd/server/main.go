package main

import (
	"context"
	"flag"

	"github.com/rs/zerolog/log"

	"github.com/anhnmt/golang-gateway-boilerplate/pkg/config"
	"github.com/anhnmt/golang-gateway-boilerplate/pkg/logger"
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
}

func main() {
	logger.New(logFile)
	config.New(env)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Info().Msg("Hello world")

	// wait for termination signal
	wait := utils.GracefulShutdown(ctx, utils.DefaultShutdownTimeout, map[string]utils.Operation{})
	<-wait

	log.Info().Msg("graceful shutdown complete")
}
