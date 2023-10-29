package gateway

import (
	"net/http"

	"connectrpc.com/vanguard"
	"github.com/rs/zerolog/log"

	"github.com/anhnmt/golang-gateway-boilerplate/internal/service/userservice"
	"github.com/anhnmt/golang-gateway-boilerplate/proto/gengo/user/v1/userv1connect"
)

func New(
	userService *userservice.Service,
) *vanguard.Transcoder {
	svc := []*vanguard.Service{
		vanguard.NewService(userv1connect.NewUserServiceHandler(userService)),
	}

	// Using Vanguard, the server can also accept RESTful requests. The Vanguard
	// Transcoder handles both REST and RPC traffic, so there's no need to mount
	// the RPC-only handler.
	transcoder, err := vanguard.NewTranscoder(
		svc,
		vanguard.WithUnknownHandler(custom404handler()),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create transcoder")
	}

	return transcoder
}

func custom404handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "custom 404 error", http.StatusNotFound)
	})
}
