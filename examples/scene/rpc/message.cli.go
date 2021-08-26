// Code generated by dol build. Only Generate by tools if not existed.
// source: MessageSrv.cli.go

package rpc

import (
	"context"
	"time"

	"scene/rpc/proto"

	"github.com/2637309949/dolphin/platform/util/trace"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// MessageSrvClient defined
var MessageSrvClient proto.MessageSrvClient

// NewMessageSrvClient defined TODO
func NewMessageSrvClient(target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	options := append(opts, []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithChainUnaryInterceptor(trace.RpcSrvTrace),
	}...)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, viper.GetString("rpc.domain_srv"), options...)
	return conn, err
}

func init() {
	go func() {
		time.Sleep(1 * time.Second)
		conn, err := NewMessageSrvClient(viper.GetString("rpc.message_srv"))
		if err != nil {
			logrus.Errorf("grpc dial rpc.message_srv failed: %v", err)
		}
		MessageSrvClient = proto.NewMessageSrvClient(conn)
	}()
}
