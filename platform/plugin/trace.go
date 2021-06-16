package plugin

import (
	"context"
	"os"

	"github.com/2637309949/dolphin/packages/trace"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var hostname string

func init() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = trace.RandId()
	}
}

// Hostname returns the name of the host, if no hostname, a random id is returned.
func Hostname() string {
	return hostname
}

func HttpTrace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		carrier, err := trace.Extract(trace.HttpFormat, ctx.Request.Header)
		if err != nil && err != trace.ErrInvalidCarrier {
			logrus.Error(err)
		}
		c, span := trace.StartServerSpan(ctx.Request.Context(), carrier, Hostname(), ctx.Request.RequestURI)
		defer span.Finish()
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	}
}

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
