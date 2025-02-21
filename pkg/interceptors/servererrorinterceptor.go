package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"inquora/pkg/xcode"
)

func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		return resp, xcode.FromError(err).Err()
	}
}
