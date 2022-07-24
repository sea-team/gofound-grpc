package interceptor

import (
	"context"
	"gofound-grpc/global"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func VerifyAuth() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		m, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "")
		}

		auths := make([]string, 0, 2)
		for _, v := range m[global.CONFIG.Auth.Header] {
			auths = strings.Split(v, global.CONFIG.Auth.Separator)
		}

		if len(auths) < 2 {
			return nil, status.Error(codes.Unauthenticated, "")
		}

		// "admin:123456"
		if auths[0] != global.CONFIG.Auth.Username && auths[1] != global.CONFIG.Auth.Password {
			return nil, status.Error(codes.Unauthenticated, "")
		}

		return handler(ctx, req)
	}
}
