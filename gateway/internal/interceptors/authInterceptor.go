package interceptors

import (
	"context"
	"fmt"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/e1esm/LyrVibe/auth-service/pkg/logger"
	"github.com/e1esm/LyrVibe/gateway/pkg/config"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	authorizationHeader = runtime.MetadataHeaderPrefix + "authorization"
	userIDHeader        = "user_id"
	userRoleHeader      = "user_role"
)

type AuthInterceptor struct {
	AuthClient proto.AuthServiceClient
}

func NewAuthInterceptor(cfg config.Config) (*AuthInterceptor, error) {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s",
		cfg.AuthService.Address, cfg.AuthService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Error("Couldn't have set up a connection with the service",
			zap.String("err", err.Error()))
		return nil, err
	}

	return &AuthInterceptor{AuthClient: proto.NewAuthServiceClient(conn)}, nil
}

func (ai *AuthInterceptor) identify(ctx context.Context) (context.Context, error) {
	val := metautils.ExtractOutgoing(ctx).Get(authorizationHeader)
	reply, err := ai.AuthClient.Identify(ctx, &proto.IdentifyRequest{AccessToken: strings.TrimPrefix(val, "Bearer")})
	if err != nil {
		return nil, err
	}
	return metadata.AppendToOutgoingContext(ctx, userIDHeader, reply.UserId, userRoleHeader, reply.Role), nil
}

func (i *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption) error {
		ctx, err := i.identify(ctx)
		if err != nil {
			return err
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
