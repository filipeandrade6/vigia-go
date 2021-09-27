package service

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type AuthInterceptor struct {
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		fmt.Println("--> unary interceptor: ", info.FullMethod)
		return handler(ctx, req)
	}
}
