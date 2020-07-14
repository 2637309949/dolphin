// Code generated by dol build. Only Generate by tools if not existed.
// source: DomainSrv.cli.go

package rpc

import (
	"github.com/2637309949/dolphin/platform/rpc/proto"

	"github.com/2637309949/dolphin/packages/logrus"
	"github.com/2637309949/dolphin/packages/viper"
	"google.golang.org/grpc"
)

// DomainSrvClient defined
var DomainSrvClient proto.DomainSrvClient

func init() {
	opt := grpc.WithInsecure()
	conn, err := grpc.Dial(viper.GetString("rpc.domain_srv"), opt)
	if err != nil {
		logrus.Error("grpc dial failed: %v", err)
	}
	DomainSrvClient = proto.NewDomainSrvClient(conn)
}