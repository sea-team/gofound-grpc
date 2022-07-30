package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() error
}

// Validator 验证器
func Validator() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if v, ok := req.(validator); ok {
			if err := v.Validate(); err != nil {
				return nil, err
			}
		}
		return handler(ctx, req)
	}
}
