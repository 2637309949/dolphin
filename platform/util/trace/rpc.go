package trace

import (
	"context"

	"github.com/2637309949/dolphin/packages/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// RpcSrvTrace is an interceptor that handles tracing.
func RpcSrvTrace(ctx context.Context, method string, req, reply interface{},
	cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, span := trace.StartClientSpan(ctx, cc.Target(), method)
	defer span.Finish()

	var pairs []string
	span.Visit(func(key, val string) bool {
		pairs = append(pairs, key, val)
		return true
	})
	ctx = metadata.AppendToOutgoingContext(ctx, pairs...)

	return invoker(ctx, method, req, reply, cc, opts...)
}

// RpcCliTrace returns a func that handles tracing with given service name.
func RpcCliTrace(serviceName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return handler(ctx, req)
		}

		carrier, err := trace.Extract(trace.GrpcFormat, md)
		if err != nil {
			return handler(ctx, req)
		}

		ctx, span := trace.StartServerSpan(ctx, carrier, serviceName, info.FullMethod)
		defer span.Finish()
		return handler(ctx, req)
	}
}
