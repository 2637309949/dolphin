package rpc

import (
	"demo/rpc/proto"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"google.golang.org/grpc"
)

// MessageClient defined
var MessageClient proto.MessageClient

func init() {
	opt := grpc.WithInsecure()
	conn, err := grpc.Dial(viper.GetString("rpc.message"), opt)
	if err != nil {
		logrus.Error("grpc dial failed: %v", err)
	}
	MessageClient = proto.NewMessageClient(conn)
}
