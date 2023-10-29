package userservice

import (
	"context"

	"connectrpc.com/connect"

	userv1 "github.com/anhnmt/golang-gateway-boilerplate/proto/gengo/user/v1"
	"github.com/anhnmt/golang-gateway-boilerplate/proto/gengo/user/v1/userv1connect"
)

var _ userv1connect.UserServiceHandler = &Service{}

type Service struct {
	userv1connect.UnimplementedUserServiceHandler
}

func New() *Service {
	return &Service{}
}

func (s *Service) List(context.Context, *connect.Request[userv1.ListRequest]) (*connect.Response[userv1.ListResponse], error) {
	return connect.NewResponse(&userv1.ListResponse{
		Data: nil,
	}), nil
}
