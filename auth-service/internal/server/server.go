package server

import (
	"context"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/auth-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	Server      *grpc.Server
	AuthService service.Service
	proto.UnimplementedAuthServiceServer
}

func (s Server) SignUp(ctx context.Context, request *proto.SignUpRequest) (*proto.SignUpResponse, error) {

}

func (s Server) SignIn(ctx context.Context, request *proto.SignInRequest) (*proto.SignInResponse, error) {

}

func (s Server) RefreshToken(ctx context.Context, request *proto.RefreshRequest) (*emptypb.Empty, error) {

}
