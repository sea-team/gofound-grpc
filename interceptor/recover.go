package interceptor

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
)

// GrpcRecover 防止panic导致主程序崩溃
func GrpcRecover() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			if e := recover(); e != nil {
				err = errors.New(fmt.Sprintf("panic:%v", e))
			}
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}
