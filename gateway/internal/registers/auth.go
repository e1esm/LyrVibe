package registers

import (
	"context"
	"github.com/e1esm/LyrVibe/auth-service/api/v1/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterAuth(ctx context.Context, mux *runtime.ServeMux, addr string) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	return proto.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, addr, opts)
}
