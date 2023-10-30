package gateway

import (
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/vanguard"

	"github.com/anhnmt/golang-gateway-boilerplate/internal/service/userservice"
	"github.com/anhnmt/golang-gateway-boilerplate/proto/gengo/user/v1/userv1connect"
)

func New(
	interceptors connect.Option,
	userService *userservice.Service,
) (*vanguard.Transcoder, error) {
	services := []*vanguard.Service{
		vanguard.NewService(userv1connect.NewUserServiceHandler(userService, interceptors)),
	}

	transcoderOptions := []vanguard.TranscoderOption{
		vanguard.WithUnknownHandler(custom404handler()),
	}

	// Using Vanguard, the server can also accept RESTful requests. The Vanguard
	// Transcoder handles both REST and RPC traffic, so there's no need to mount
	// the RPC-only handler.
	transcoder, err := vanguard.NewTranscoder(services, transcoderOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to create transcoder: %w", err)
	}

	return transcoder, nil
}

func custom404handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "custom 404 error", http.StatusNotFound)
	})
}
