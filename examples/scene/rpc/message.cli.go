// Code generated by dol build. Only Generate by tools if not existed.
// source: MessageSrv.cli.go

package rpc

import (
	"scene/rpc/proto"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// MessageSrvClient defined
var MessageSrvClient proto.MessageSrvClient

func init() {
	opt := grpc.WithInsecure()
	conn, err := grpc.Dial(viper.GetString("rpc.message_srv"), opt)
	if err != nil {
		logrus.Errorf("grpc dial failed: %v", err)
	}
	MessageSrvClient = proto.NewMessageSrvClient(conn)
}
